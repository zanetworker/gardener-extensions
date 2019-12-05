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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/gardener/gardener-extensions/controllers/provider-azure/pkg/azure"
	"github.com/gardener/gardener-extensions/pkg/controller/controlplane"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	extensionshealthcheckcontroller "github.com/gardener/gardener-extensions/pkg/controller/healthcheck"
	"github.com/gardener/gardener-extensions/pkg/controller/healthcheck/general"
	"github.com/gardener/gardener-extensions/pkg/controller/healthcheck/worker"
	"github.com/gardener/gardener-extensions/pkg/controller/worker/genericactuator"

	controlplaneActuator "github.com/gardener/gardener-extensions/pkg/controller/controlplane/genericactuator"
	extensionhealthcheck "github.com/gardener/gardener-extensions/pkg/controller/healthcheck/config"
)

var (
	defaultSyncPeriod = time.Second * 30
	// DefaultAddOptions are the default DefaultAddArgs for AddToManager.
	DefaultAddOptions = extensionshealthcheckcontroller.DefaultAddArgs{
		HealthCheckConfig: extensionhealthcheck.HealthCheckConfig{SyncPeriod: metav1.Duration{Duration: defaultSyncPeriod}},
	}
)

// RegisterHealthChecks registers health checks for each extension resource
// HealthChecks are grouped by extension (e.g worker), extension.type (e.g azure) and  Health Check Type (e.g SystemComponentsHealthy)
func RegisterHealthChecks(mgr manager.Manager, opts extensionshealthcheckcontroller.DefaultAddArgs) error {
	normalPredicates := []predicate.Predicate{controlplane.HasPurpose(extensionsv1alpha1.Normal)}
	if err := extensionshealthcheckcontroller.DefaultRegisterExtensionForHealthCheck(
		azure.Type,
		extensionsv1alpha1.SchemeGroupVersion.WithKind(extensionsv1alpha1.ControlPlaneResource),
		func() runtime.Object { return &extensionsv1alpha1.ControlPlane{} },
		mgr,
		opts,
		normalPredicates,
		map[extensionshealthcheckcontroller.HealthCheck]string{
			general.CheckSeedDeployment(azure.CloudControllerManagerName):                         string(gardencorev1alpha1.ShootControlPlaneHealthy),
			general.CheckManagedResource(controlplaneActuator.ControlPlaneShootChartResourceName): string(gardencorev1alpha1.ShootSystemComponentsHealthy),
			general.CheckManagedResource(controlplaneActuator.StorageClassesChartResourceName):    string(gardencorev1alpha1.ShootSystemComponentsHealthy),
			general.CheckManagedResource(controlplaneActuator.ShootWebhooksResourceName):          string(gardencorev1alpha1.ShootSystemComponentsHealthy),
		}); err != nil {
		return err
	}

	return extensionshealthcheckcontroller.DefaultRegisterExtensionForHealthCheck(
		azure.Type,
		extensionsv1alpha1.SchemeGroupVersion.WithKind(extensionsv1alpha1.WorkerResource),
		func() runtime.Object { return &extensionsv1alpha1.Worker{} },
		mgr,
		opts,
		nil,
		map[extensionshealthcheckcontroller.HealthCheck]string{
			general.CheckManagedResource(genericactuator.McmShootResourceName): string(gardencorev1alpha1.ShootSystemComponentsHealthy),
			general.CheckSeedDeployment(azure.MachineControllerManagerName):    string(gardencorev1alpha1.ShootEveryNodeReady),
			worker.SufficientNodesAvailable():                                  string(gardencorev1alpha1.ShootEveryNodeReady),
		})
}

// AddToManager adds a controller with the default Options.
func AddToManager(mgr manager.Manager) error {
	return RegisterHealthChecks(mgr, DefaultAddOptions)
}
