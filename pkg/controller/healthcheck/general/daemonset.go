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

// DaemonSetHealthChecker contains all the information for the DaemonSet HealthCheck
type DaemonSetHealthChecker struct {
	logger         logr.Logger
	seedClient     client.Client
	shootClient    client.Client
	deploymentName string
}

// CheckStatefulSet is a healthCheck function to check DaemonSets
func CheckDaemonSet(name string) healthcheck.HealthCheck {
	return &DaemonSetHealthChecker{
		deploymentName: name,
	}
}

// SetSeedClient injects the seed client
func (healthChecker *DaemonSetHealthChecker) SetSeedClient(seedClient client.Client) {
	healthChecker.seedClient = seedClient
}

// SetShootClient injects the shoot client
func (healthChecker *DaemonSetHealthChecker) SetShootClient(shootClient client.Client) {
	healthChecker.shootClient = shootClient
}

// SetLoggerSuffix injects the logger
func (healthChecker *DaemonSetHealthChecker) SetLoggerSuffix(provider, extension string) {
	healthChecker.logger = log.Log.WithName(fmt.Sprintf("%s-%s-healthcheck-deployment", provider, extension))
}

// Check executes the health check
func (healthChecker *DaemonSetHealthChecker) Check(ctx context.Context, request types.NamespacedName) (*healthcheck.HealthCheckResult, error) {
	daemonSet := &v1.DaemonSet{}
	if err := healthChecker.seedClient.Get(ctx, client.ObjectKey{Namespace: request.Namespace, Name: healthChecker.deploymentName}, daemonSet); err != nil {
		err := fmt.Errorf("failed to retrieve DeamonSet '%s' in namespace '%s': %v", healthChecker.deploymentName, request.Namespace, err)
		healthChecker.logger.Error(err, "Health check failed")
		return nil, err
	}
	if isHealthy, reason, err := DaemonSetIsHealthy(daemonSet); !isHealthy {
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

func DaemonSetIsHealthy(daemonSet *v1.DaemonSet) (bool, *string, error) {
	if err := health.CheckDaemonSet(daemonSet); err != nil {
		reason := "DeploymentUnhealthy"
		err := fmt.Errorf("daemonSet %s in namespace %s is unhealthy: %v", daemonSet.Name, daemonSet.Namespace, err)
		return false, &reason, err
	}
	return true, nil, nil
}
