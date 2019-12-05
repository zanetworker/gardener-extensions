// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package healthcheck

import (
	"context"
	"fmt"
	"github.com/gardener/gardener/pkg/api/extensions"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	"github.com/gardener/gardener-extensions/pkg/util"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	gardencorev1alpha1helper "github.com/gardener/gardener/pkg/apis/core/v1alpha1/helper"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
)

type reconciler struct {
	logger              logr.Logger
	actuator            HealthCheckActuator
	ctx                 context.Context
	client              client.Client
	recorder            record.EventRecorder
	registeredExtension RegisteredExtension
	syncPeriod          metav1.Duration
}

// NewReconciler creates a new performHealthCheck.Reconciler that reconciles
// the registered extension resources (Gardener's `extensions.gardener.cloud` API group).
func NewReconciler(mgr manager.Manager, actuator HealthCheckActuator, registeredExtension RegisteredExtension, syncPeriod metav1.Duration) reconcile.Reconciler {
	return &reconciler{
		logger:              log.Log.WithName(ControllerName),
		actuator:            actuator,
		recorder:            mgr.GetEventRecorderFor(ControllerName),
		registeredExtension: registeredExtension,
		syncPeriod:          syncPeriod,
	}
}

func (r *reconciler) InjectFunc(f inject.Func) error {
	return f(r.actuator)
}

func (r *reconciler) InjectClient(client client.Client) error {
	r.client = client
	return nil
}

func (r *reconciler) InjectStopChannel(stopCh <-chan struct{}) error {
	r.ctx = util.ContextFromStopChannel(stopCh)
	return nil
}

func (r *reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	rawExtension := unstructured.Unstructured{}
	rawExtension.SetGroupVersionKind(r.registeredExtension.groupVersionKind)

	if err := r.client.Get(r.ctx, request.NamespacedName, &rawExtension); err != nil {
		if errors.IsNotFound(err) {
			return r.resultWithRequeue(), nil
		}
		return r.resultWithRequeue(), err
	}

	acc, err := extensions.Accessor(rawExtension.DeepCopyObject())
	if err != nil {
		return r.resultWithRequeue(), err
	}

	if acc.GetDeletionTimestamp() != nil {
		r.logger.Info("Do not perform HealthCheck for extension resource. extension is being deleted.", "name", acc.GetName(), "Namespace", acc.GetNamespace())
		return reconcile.Result{}, nil
	}

	return r.performHealthCheck(r.ctx, request, rawExtension)
}

func (r *reconciler) performHealthCheck(ctx context.Context, request reconcile.Request, rawExtension unstructured.Unstructured) (reconcile.Result, error) {
	healthCheckResults, err := r.actuator.ExecuteHealthCheckFunctions(ctx, types.NamespacedName{Namespace: request.Namespace, Name: request.Name})
	if err != nil {
		r.logger.Info("Failed to execute healthChecks. Updating each HealthCheckCondition for the extension resource to ConditionCheckError.", "kind", r.registeredExtension.groupVersionKind.Kind, "health condition type", r.registeredExtension.healthConditionType, "name", request.Name, "Namespace", request.Namespace, "Error", err.Error())
		for healthConditionType := range r.registeredExtension.healthConditionType {
			healthCondition := gardencorev1alpha1helper.GetOrInitCondition(r.registeredExtension.extension.GetExtensionStatus().GetConditions(), gardencorev1alpha1.ConditionType(healthConditionType))
			if err := r.updateExtensionConditionFailedToExecute(ctx, r.registeredExtension.extension, &rawExtension, healthCondition, r.registeredExtension.groupVersionKind.Kind, err); err != nil {
				return r.resultWithRequeue(), err
			}
		}
		return r.resultWithRequeue(), nil
	}

	for _, healthCheckResult := range *healthCheckResults {
		// get or init conditions on extension resource
		healthCondition := gardencorev1alpha1helper.GetOrInitCondition(r.registeredExtension.extension.GetExtensionStatus().GetConditions(), gardencorev1alpha1.ConditionType(healthCheckResult.HealthConditionType))
		if !healthCheckResult.IsHealthy && healthCheckResult.FailedChecks > 0 {
			r.logger.Info("Updating HealthCheckCondition for extension resource to ConditionCheckError.", "Kind", r.registeredExtension.groupVersionKind.Kind, "health condition type", healthCheckResult.HealthConditionType, "name", request.Name, "Namespace", request.Namespace)
			if err := r.updateExtensionConditionToConditionCheckError(ctx, r.registeredExtension.extension, &rawExtension, healthCondition, r.registeredExtension.groupVersionKind.Kind, healthCheckResult); err != nil {
				return r.resultWithRequeue(), err
			}
			continue
		}

		if !healthCheckResult.IsHealthy {
			r.logger.Info("Health check for extension resource unsuccessful.", "kind", fmt.Sprintf("%s.%s.%s", r.registeredExtension.groupVersionKind.Kind, r.registeredExtension.groupVersionKind.Group, r.registeredExtension.groupVersionKind.Version), "name", request.Name, "Namespace", request.Namespace, "failed", healthCheckResult.FailedChecks, "successful", healthCheckResult.SuccessfulChecks, "details", healthCheckResult.GetDetails())
			if err := r.updateExtensionConditionToError(ctx, r.registeredExtension.extension, &rawExtension, healthCondition, healthCheckResult); err != nil {
				return r.resultWithRequeue(), err
			}
			continue
		}

		r.logger.Info("Health check for extension resource successful.", "kind", r.registeredExtension.groupVersionKind.Kind, "health condition type", healthCheckResult.HealthConditionType, "name", request.Name, "Namespace", request.Namespace)
		if err := r.updateExtensionConditionToSuccessful(ctx, r.registeredExtension.extension, &rawExtension, healthCondition, healthCheckResult); err != nil {
			return r.resultWithRequeue(), err
		}
	}
	return r.resultWithRequeue(), nil
}

func (r *reconciler) updateExtensionConditionFailedToExecute(ctx context.Context, extensionResource extensionsv1alpha1.Object, extension *unstructured.Unstructured, condition gardencorev1alpha1.Condition, kind string, err error) error {
	healthCondition := gardencorev1alpha1helper.UpdatedCondition(condition, gardencorev1alpha1.ConditionUnknown, gardencorev1alpha1.ConditionCheckError, fmt.Sprintf("failed to execute health checks for '%s': %v", kind, err.Error()))
	return r.updateExtensionCondition(ctx, extension, condition, extensionResource, healthCondition)
}

func (r *reconciler) updateExtensionConditionToConditionCheckError(ctx context.Context, extensionResource extensionsv1alpha1.Object, extension *unstructured.Unstructured, condition gardencorev1alpha1.Condition, kind string, healthCheckResult ExtensionHealthCheckResult) error {
	healthCondition := gardencorev1alpha1helper.UpdatedCondition(condition, gardencorev1alpha1.ConditionUnknown, gardencorev1alpha1.ConditionCheckError, fmt.Sprintf("failed to execute %d/%d health checks for '%s': %v", healthCheckResult.FailedChecks, healthCheckResult.SuccessfulChecks+healthCheckResult.UnsuccessfulChecks+healthCheckResult.FailedChecks, kind, healthCheckResult.GetDetails()))
	return r.updateExtensionCondition(ctx, extension, condition, extensionResource, healthCondition)
}

func (r *reconciler) updateExtensionConditionToError(ctx context.Context, extensionResource extensionsv1alpha1.Object, extension *unstructured.Unstructured, condition gardencorev1alpha1.Condition, healthCheckResult ExtensionHealthCheckResult) error {
	detail := fmt.Sprintf("Health check for %d/%d component(s) unsuccessful. ", healthCheckResult.UnsuccessfulChecks, healthCheckResult.UnsuccessfulChecks+healthCheckResult.SuccessfulChecks)
	healthCondition := gardencorev1alpha1helper.UpdatedCondition(condition, gardencorev1alpha1.ConditionFalse, "HealthCheckUnsuccessful", detail+healthCheckResult.GetDetails())
	return r.updateExtensionCondition(ctx, extension, condition, extensionResource, healthCondition)
}

func (r *reconciler) updateExtensionConditionToSuccessful(ctx context.Context, extensionResource extensionsv1alpha1.Object, extension *unstructured.Unstructured, condition gardencorev1alpha1.Condition, healthCheckResult ExtensionHealthCheckResult) error {
	healthCondition := gardencorev1alpha1helper.UpdatedCondition(condition, gardencorev1alpha1.ConditionTrue, "HealthCheckSuccessful", fmt.Sprintf("(%d/%d) Health checks successful", healthCheckResult.SuccessfulChecks, healthCheckResult.SuccessfulChecks))
	return r.updateExtensionCondition(ctx, extension, condition, extensionResource, healthCondition)
}

func (r *reconciler) updateExtensionCondition(ctx context.Context, extension *unstructured.Unstructured, condition gardencorev1alpha1.Condition, extensionResource extensionsv1alpha1.Object, healthCondition gardencorev1alpha1.Condition) error {
	return extensionscontroller.TryUpdateStatus(ctx, retry.DefaultBackoff, r.client, extension, func() error {
		acc, err := extensions.Accessor(extension.DeepCopyObject())
		if err != nil {
			return fmt.Errorf("error updating health check condition (type: %s, name: %s, ns %s) - failed to create an extensionsv1alpha1.Object from the extension object: %v", condition.Type, extensionResource.GetName(), extensionResource.GetNamespace(), err)
		}
		conditions := gardencorev1alpha1helper.MergeConditions(acc.GetExtensionStatus().GetConditions(), healthCondition)
		acc.GetExtensionStatus().SetConditions(conditions)
		unstrc, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&acc)
		if err != nil {
			return fmt.Errorf("error writing health check condition to error (type: %s, name: %s, ns %s) - failed to convert extensionsv1alpha1.Object back to extension object: %v", condition.Type, extensionResource.GetName(), extensionResource.GetNamespace(), err)
		}
		extension.Object = unstrc
		return nil
	})
}

func (r *reconciler) resultWithRequeue() reconcile.Result {
	return reconcile.Result{RequeueAfter: r.syncPeriod.Duration}
}
