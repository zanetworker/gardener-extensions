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

package controlplane

import (
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	extensionspredicate "github.com/gardener/gardener-extensions/pkg/predicate"
)

// GenerationChangedPredicate is a predicate for generation changes.
func GenerationChangedPredicate() predicate.Predicate {
	return predicate.Funcs{
		UpdateFunc: func(event event.UpdateEvent) bool {
			return event.MetaOld.GetGeneration() != event.MetaNew.GetGeneration()
		},
	}
}

// HasPurpose filters the incoming Controlplanes  for the given spec.purpose
func HasPurpose(purpose extensionsv1alpha1.Purpose) predicate.Predicate {
	return extensionspredicate.FromMapper(extensionspredicate.MapperFunc(func(e event.GenericEvent) bool {
		controlPlane, ok := e.Object.(*extensionsv1alpha1.ControlPlane)
		if !ok {
			return false
		}

		// needed because ControlPlane of type "normal" has the spec.purpose field not set
		if controlPlane.Spec.Purpose == nil && purpose == extensionsv1alpha1.Normal {
			return true
		}

		if controlPlane.Spec.Purpose == nil {
			return false
		}

		return *controlPlane.Spec.Purpose == purpose
	}), extensionspredicate.CreateTrigger, extensionspredicate.UpdateNewTrigger, extensionspredicate.DeleteTrigger, extensionspredicate.GenericTrigger)
}
