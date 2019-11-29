package general

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	resourcesv1alpha1 "github.com/gardener/gardener-resource-manager/pkg/apis/resources/v1alpha1"

	"github.com/gardener/gardener-extensions/pkg/controller/healthcheck"
)

// ManagedResourceHealthChecker contains all the information for the ManagedResource HealthCheck
type ManagedResourceHealthChecker struct {
	logger              logr.Logger
	seedClient          client.Client
	shootClient         client.Client
	managedResourceName string
}

// CheckManagedResource is a healthCheck function to check ManagedResources
func CheckManagedResource(managedResourceName string) healthcheck.HealthCheck {
	return &ManagedResourceHealthChecker{
		managedResourceName: managedResourceName,
	}
}

// SetSeedClient injects the seed client
func (healthChecker *ManagedResourceHealthChecker) SetSeedClient(seedClient client.Client) {
	healthChecker.seedClient = seedClient
}

// SetShootClient injects the shoot client
func (healthChecker *ManagedResourceHealthChecker) SetShootClient(shootClient client.Client) {
	healthChecker.shootClient = shootClient
}

// SetLoggerSuffix injects the logger
func (healthChecker *ManagedResourceHealthChecker) SetLoggerSuffix(provider, extension string) {
	healthChecker.logger = log.Log.WithName(fmt.Sprintf("%s-%s-healthcheck-managed-resource", provider, extension))
}

// Check executes the health check
func (healthChecker *ManagedResourceHealthChecker) Check(ctx context.Context, request types.NamespacedName) (*healthcheck.HealthCheckResult, error) {
	mcmDeployment := &resourcesv1alpha1.ManagedResource{}

	if err := healthChecker.seedClient.Get(ctx, client.ObjectKey{Namespace: request.Namespace, Name: healthChecker.managedResourceName}, mcmDeployment); err != nil {
		err := fmt.Errorf("check Managed Resource failed. Unable to retrieve managed resource '%s' in namespace '%s': %v", healthChecker.managedResourceName, request.Namespace, err)
		healthChecker.logger.Error(err, "Health check failed")
		return nil, err
	}
	if isHealthy, reason, err := managedResourceIsHealthy(mcmDeployment); !isHealthy {
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

func managedResourceIsHealthy(resource *resourcesv1alpha1.ManagedResource) (bool, *string, error) {
	if err := checkManagedResourceIsHealthy(resource); err != nil {
		reason := "ManagedResourceUnhealthy"
		err := fmt.Errorf("managed resource %s in namespace %s is unhealthy: %v", resource.Name, resource.Namespace, err)
		return false, &reason, err
	}
	return true, nil, nil
}
