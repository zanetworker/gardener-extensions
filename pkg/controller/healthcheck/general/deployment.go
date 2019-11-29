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

// DeploymentHealthChecker contains all the information for the Deployment HealthCheck
type DeploymentHealthChecker struct {
	logger         logr.Logger
	seedClient     client.Client
	shootClient    client.Client
	deploymentName string
	checkType      DeploymentCheckType
}

// DeploymentCheckType in which cluster the check will be executed
type DeploymentCheckType string

const (
	DeploymentCheckTypeSeed  DeploymentCheckType = "Seed"
	DeploymentCheckTypeShoot DeploymentCheckType = "Shoot"
)

// CheckSeedDeployment is a healthCheck function to check Deployments in the Seed cluster
func CheckSeedDeployment(deploymentName string) healthcheck.HealthCheck {
	return &DeploymentHealthChecker{
		deploymentName: deploymentName,
		checkType:      DeploymentCheckTypeSeed,
	}
}

// CheckSeedDeployment is a healthCheck function to check Deployments in the Shoot cluster
func CheckShootDeployment(deploymentName string) healthcheck.HealthCheck {
	return &DeploymentHealthChecker{
		deploymentName: deploymentName,
		checkType:      DeploymentCheckTypeShoot,
	}
}

// SetSeedClient injects the seed client
func (healthChecker *DeploymentHealthChecker) SetSeedClient(seedClient client.Client) {
	healthChecker.seedClient = seedClient
}

// SetShootClient injects the shoot client
func (healthChecker *DeploymentHealthChecker) SetShootClient(shootClient client.Client) {
	healthChecker.shootClient = shootClient
}

// SetLoggerSuffix injects the logger
func (healthChecker *DeploymentHealthChecker) SetLoggerSuffix(provider, extension string) {
	healthChecker.logger = log.Log.WithName(fmt.Sprintf("%s-%s-healthcheck-deployment", provider, extension))
}

// Check executes the health check
func (healthChecker *DeploymentHealthChecker) Check(ctx context.Context, request types.NamespacedName) (*healthcheck.HealthCheckResult, error) {
	deployment := &v1.Deployment{}

	var err error
	if healthChecker.checkType == DeploymentCheckTypeSeed {
		err = healthChecker.seedClient.Get(ctx, client.ObjectKey{Namespace: request.Namespace, Name: healthChecker.deploymentName}, deployment)
	} else {
		err = healthChecker.shootClient.Get(ctx, client.ObjectKey{Namespace: request.Namespace, Name: healthChecker.deploymentName}, deployment)
	}
	if err != nil {
		err := fmt.Errorf("failed to retrieve deployment '%s' in namespace '%s': %v", healthChecker.deploymentName, request.Namespace, err)
		healthChecker.logger.Error(err, "Health check failed")
		return nil, err
	}

	if isHealthy, reason, err := deploymentIsHealthy(deployment); !isHealthy {
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

func deploymentIsHealthy(deployment *v1.Deployment) (bool, *string, error) {
	if err := health.CheckDeployment(deployment); err != nil {
		reason := "DeploymentUnhealthy"
		err := fmt.Errorf("deployment %s in namespace %s is unhealthy: %v", deployment.Name, deployment.Namespace, err)
		return false, &reason, err
	}
	return true, nil, nil
}
