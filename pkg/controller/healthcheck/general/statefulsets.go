package general

import (
	"context"
	"fmt"
	"github.com/gardener/gardener/pkg/utils/kubernetes/health"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/gardener/gardener-extensions/pkg/controller/healthcheck"
)

// StatefulSetHealthChecker contains all the information for the StatefulSet HealthCheck
type StatefulSetHealthChecker struct {
	logger         logr.Logger
	seedClient     client.Client
	shootClient    client.Client
	deploymentName string
}

// CheckStatefulSet is a healthCheck function to check StatefulSets
func CheckStatefulSet(name string) healthcheck.HealthCheck {
	return &StatefulSetHealthChecker{
		deploymentName: name,
	}
}

// SetSeedClient injects the seed client
func (healthChecker *StatefulSetHealthChecker) SetSeedClient(seedClient client.Client) {
	healthChecker.seedClient = seedClient
}

// SetShootClient injects the shoot client
func (healthChecker *StatefulSetHealthChecker) SetShootClient(shootClient client.Client) {
	healthChecker.shootClient = shootClient
}

// SetLoggerSuffix injects the logger
func (healthChecker *StatefulSetHealthChecker) SetLoggerSuffix(provider, extension string) {
	healthChecker.logger = log.Log.WithName(fmt.Sprintf("%s-%s-healthcheck-deployment", provider, extension))
}

// Check executes the health check
func (healthChecker *StatefulSetHealthChecker) Check(ctx context.Context, request types.NamespacedName) (*healthcheck.HealthCheckResult, error) {
	statefulSet := &v1.StatefulSet{}
	if err := healthChecker.seedClient.Get(ctx, client.ObjectKey{Namespace: request.Namespace, Name: healthChecker.deploymentName}, statefulSet); err != nil {
		err := fmt.Errorf("failed to retrieve StatefulSet '%s' in namespace '%s': %v", healthChecker.deploymentName, request.Namespace, err)
		healthChecker.logger.Error(err, "Health check failed")
		return nil, err
	}
	if isHealthy, reason, err := statefulSetIsHealthy(statefulSet); !isHealthy {
		healthChecker.logger.Error(err, "Health check failed")
		return &healthcheck.HealthCheckResult{
			IsHealthy: false,
			Detail:    err.Error(),
			Reason:    *reason,
		}, nil
	}

	return &healthcheck.HealthCheckResult{
		IsHealthy: true,
	}, nil
}

func statefulSetIsHealthy(statefulSet *v1.StatefulSet) (bool, *string, error) {
	if err := health.CheckStatefulSet(statefulSet); err != nil {
		reason := "DeploymentUnhealthy"
		err := fmt.Errorf("statefulSet %s in namespace %s is unhealthy: %v", statefulSet.Name, statefulSet.Namespace, err)
		return false, &reason, err
	}
	return true, nil, nil
}
