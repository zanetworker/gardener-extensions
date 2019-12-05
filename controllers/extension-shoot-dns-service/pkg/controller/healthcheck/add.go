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
	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener-extensions/controllers/extension-shoot-dns-service/pkg/controller/config"
	extensionshealthcheckcontroller "github.com/gardener/gardener-extensions/pkg/controller/healthcheck"
	"github.com/gardener/gardener-extensions/pkg/controller/healthcheck/general"

	dnscontroller "github.com/gardener/gardener-extensions/controllers/extension-shoot-dns-service/pkg/controller"
	"github.com/gardener/gardener-extensions/controllers/extension-shoot-dns-service/pkg/service"
	extensionhealthcheck "github.com/gardener/gardener-extensions/pkg/controller/healthcheck/config"
)

// RegisterHealthChecks registers health checks for each extension resource
// HealthChecks are grouped by extension (e.g worker), extension.type (e.g aws) and  Health Check Type (e.g SystemComponentsHealthy)
func RegisterHealthChecks(mgr manager.Manager) error {
	opts := extensionshealthcheckcontroller.DefaultAddArgs{
		Controller:        config.HealthConfig.ControllerOptions,
		HealthCheckConfig: extensionhealthcheck.HealthCheckConfig{SyncPeriod: config.HealthConfig.Health.HealthCheckSyncPeriod},
	}

	return extensionshealthcheckcontroller.DefaultRegisterExtensionForHealthCheck(
		service.ExtensionType,
		extensionsv1alpha1.SchemeGroupVersion.WithKind(extensionsv1alpha1.ExtensionResource),
		func() runtime.Object { return &extensionsv1alpha1.Extension{} },
		mgr,
		opts,
		nil,
		map[extensionshealthcheckcontroller.HealthCheck]string{
			general.CheckManagedResource(dnscontroller.ShootResourcesName): string(gardencorev1alpha1.ShootSystemComponentsHealthy),
			general.CheckManagedResource(dnscontroller.SeedResourcesName):  string(gardencorev1alpha1.ShootControlPlaneHealthy),
		})
}
