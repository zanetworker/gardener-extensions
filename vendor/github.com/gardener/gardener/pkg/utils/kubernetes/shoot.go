// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package kubernetes

import (
	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	gardencore "github.com/gardener/gardener/pkg/client/core/clientset/versioned"
	"github.com/gardener/gardener/pkg/logger"

	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

func tryUpdateShoot(
	g gardencore.Interface,
	backoff wait.Backoff,
	meta metav1.ObjectMeta,
	transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error),
	updateFunc func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error),
	equalFunc func(cur, updated *gardencorev1alpha1.Shoot) bool,
) (*gardencorev1alpha1.Shoot, error) {

	var (
		result  *gardencorev1alpha1.Shoot
		attempt int
	)
	err := retry.RetryOnConflict(backoff, func() (err error) {
		attempt++
		cur, err := g.CoreV1alpha1().Shoots(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}

		updated, err := transform(cur.DeepCopy())
		if err != nil {
			return err
		}

		if equalFunc(cur, updated) {
			result = cur
			return nil
		}

		result, err = updateFunc(g, updated)
		if err != nil {
			logger.Logger.Errorf("Attempt %d failed to update Shoot %s/%s due to %v", attempt, cur.Namespace, cur.Name, err)
		}
		return
	})
	if err != nil {
		logger.Logger.Errorf("Failed to updated Shoot %s/%s after %d attempts due to %v", meta.Namespace, meta.Name, attempt, err)
	}
	return result, err
}

// TryUpdateShoot tries to update the shoot matching the given <meta>.
// It retries with the given <backoff> characteristics as long as it gets Conflict errors.
// The transformation function is applied to the current state of the Shoot object. If the transformation
// yields a semantically equal Shoot, no update is done and the operation returns normally.
func TryUpdateShoot(g gardencore.Interface, backoff wait.Backoff, meta metav1.ObjectMeta, transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error)) (*gardencorev1alpha1.Shoot, error) {
	return tryUpdateShoot(g, backoff, meta, transform, func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error) {
		return g.CoreV1alpha1().Shoots(shoot.Namespace).Update(shoot)
	}, func(cur, updated *gardencorev1alpha1.Shoot) bool {
		return equality.Semantic.DeepEqual(cur, updated)
	})
}

// TryUpdateShootHibernation tries to update the status of the shoot matching the given <meta>.
// It retries with the given <backoff> characteristics as long as it gets Conflict errors.
// The transformation function is applied to the current state of the Shoot object. If the transformation
// yields a semantically equal Shoot, no update is done and the operation returns normally.
func TryUpdateShootHibernation(g gardencore.Interface, backoff wait.Backoff, meta metav1.ObjectMeta, transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error)) (*gardencorev1alpha1.Shoot, error) {
	return tryUpdateShoot(g, backoff, meta, transform, func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error) {
		return g.CoreV1alpha1().Shoots(shoot.Namespace).Update(shoot)
	}, func(cur, updated *gardencorev1alpha1.Shoot) bool {
		return equality.Semantic.DeepEqual(cur.Spec.Hibernation, updated.Spec.Hibernation)
	})
}

// TryUpdateShootStatus tries to update the status of the shoot matching the given <meta>.
// It retries with the given <backoff> characteristics as long as it gets Conflict errors.
// The transformation function is applied to the current state of the Shoot object. If the transformation
// yields a semantically equal Shoot (regarding Status), no update is done and the operation returns normally.
func TryUpdateShootStatus(g gardencore.Interface, backoff wait.Backoff, meta metav1.ObjectMeta, transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error)) (*gardencorev1alpha1.Shoot, error) {
	return tryUpdateShoot(g, backoff, meta, transform, func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error) {
		return g.CoreV1alpha1().Shoots(shoot.Namespace).UpdateStatus(shoot)
	}, func(cur, updated *gardencorev1alpha1.Shoot) bool {
		return equality.Semantic.DeepEqual(cur.Status, updated.Status)
	})
}

// TryUpdateShootLabels tries to update the status of the shoot matching the given <meta>.
// It retries with the given <backoff> characteristics as long as it gets Conflict errors.
// The transformation function is applied to the current state of the Shoot object. If the transformation
// yields a semantically equal Shoot (regarding labels), no update is done and the operation returns normally.
func TryUpdateShootLabels(g gardencore.Interface, backoff wait.Backoff, meta metav1.ObjectMeta, transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error)) (*gardencorev1alpha1.Shoot, error) {
	return tryUpdateShoot(g, backoff, meta, transform, func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error) {
		return g.CoreV1alpha1().Shoots(shoot.Namespace).Update(shoot)
	}, func(cur, updated *gardencorev1alpha1.Shoot) bool {
		return equality.Semantic.DeepEqual(cur.Labels, updated.Labels)
	})
}

// TryUpdateShootConditions tries to update the status of the shoot matching the given <meta>.
// It retries with the given <backoff> characteristics as long as it gets Conflict errors.
// The transformation function is applied to the current state of the Shoot object. If the transformation
// yields a semantically equal Shoot (regarding conditions), no update is done and the operation returns normally.
func TryUpdateShootConditions(g gardencore.Interface, backoff wait.Backoff, meta metav1.ObjectMeta, transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error)) (*gardencorev1alpha1.Shoot, error) {
	return tryUpdateShoot(g, backoff, meta, transform, func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error) {
		return g.CoreV1alpha1().Shoots(shoot.Namespace).UpdateStatus(shoot)
	}, func(cur, updated *gardencorev1alpha1.Shoot) bool {
		return equality.Semantic.DeepEqual(cur.Status.Conditions, updated.Status.Conditions)
	})
}

// TryUpdateShootConstraints tries to update the status of the shoot matching the given <meta>.
// It retries with the given <backoff> characteristics as long as it gets Conflict errors.
// The transformation function is applied to the current state of the Shoot object. If the transformation
// yields a semantically equal Shoot (regarding conditions), no update is done and the operation returns normally.
func TryUpdateShootConstraints(g gardencore.Interface, backoff wait.Backoff, meta metav1.ObjectMeta, transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error)) (*gardencorev1alpha1.Shoot, error) {
	return tryUpdateShoot(g, backoff, meta, transform, func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error) {
		return g.CoreV1alpha1().Shoots(shoot.Namespace).UpdateStatus(shoot)
	}, func(cur, updated *gardencorev1alpha1.Shoot) bool {
		return equality.Semantic.DeepEqual(cur.Status.Constraints, updated.Status.Constraints)
	})
}

// TryUpdateShootAnnotations tries to update the annotations of the shoot matching the given <meta>.
// It retries with the given <backoff> characteristics as long as it gets Conflict errors.
// The transformation function is applied to the current state of the Shoot object. If the transformation
// yields a semantically equal Shoot (regarding conditions), no update is done and the operation returns normally.
func TryUpdateShootAnnotations(g gardencore.Interface, backoff wait.Backoff, meta metav1.ObjectMeta, transform func(*gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error)) (*gardencorev1alpha1.Shoot, error) {
	return tryUpdateShoot(g, backoff, meta, transform, func(g gardencore.Interface, shoot *gardencorev1alpha1.Shoot) (*gardencorev1alpha1.Shoot, error) {
		return g.CoreV1alpha1().Shoots(shoot.Namespace).Update(shoot)
	}, func(cur, updated *gardencorev1alpha1.Shoot) bool {
		return equality.Semantic.DeepEqual(cur.Annotations, updated.Annotations)
	})
}
