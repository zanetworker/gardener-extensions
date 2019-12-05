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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener-extensions/controllers/networking-calico/pkg/calico"
	networkcontroller "github.com/gardener/gardener-extensions/controllers/networking-calico/pkg/controller"
	extensionshealthcheckcontroller "github.com/gardener/gardener-extensions/pkg/controller/healthcheck"
	extensionhealthcheck "github.com/gardener/gardener-extensions/pkg/controller/healthcheck/config"
	"github.com/gardener/gardener-extensions/pkg/controller/healthcheck/general"
)

var (
	defaultSyncPeriod = time.Second * 30
	// AddOptions are the default DefaultAddArgs for AddToManager.
	AddOptions = extensionshealthcheckcontroller.DefaultAddArgs{
		HealthCheckConfig: extensionhealthcheck.HealthCheckConfig{SyncPeriod: metav1.Duration{Duration: defaultSyncPeriod}},
	}
)

// RegisterHealthChecks adds a controller with the given Options to the manager.
// The opts.Reconciler is being set with a newly instantiated HealthActuator.
func RegisterHealthChecks(mgr manager.Manager, opts extensionshealthcheckcontroller.DefaultAddArgs) error {
	return extensionshealthcheckcontroller.DefaultRegisterExtensionForHealthCheck(
		calico.Type,
		extensionsv1alpha1.SchemeGroupVersion.WithKind(extensionsv1alpha1.NetworkResource),
		func() runtime.Object { return &extensionsv1alpha1.Network{} },
		mgr,
		opts,
		nil,
		map[extensionshealthcheckcontroller.HealthCheck]string{
			general.CheckManagedResource(networkcontroller.CalicoConfigSecretName): string(gardencorev1alpha1.ShootSystemComponentsHealthy),
		},
	)
}

// AddToManager adds a controller with the default Options.
func AddToManager(mgr manager.Manager) error {
	return RegisterHealthChecks(mgr, AddOptions)
}
