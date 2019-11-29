package worker

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	corev1 "k8s.io/api/core/v1"

	"github.com/gardener/gardener-extensions/pkg/controller/healthcheck"
)

// DefaultWorkerHealthChecker all the information for the Worker HealthCheck
// This check assumes that the MachineControllerManager (https://github.com/gardener/machine-controller-manager) has been deployed by the Worker extension controller
type DefaultWorkerHealthChecker struct {
	logger logr.Logger
	// Needs to be set by actuator before calling the Check function
	seedClient client.Client
	// make sure shoot client is instantiated
	shootClient client.Client
}

// SufficientNodesAvailable is a healthCheck function to check if there are a sufficient amount of nodes registered in the cluster
// Checks if all machines created by the machine deployment joinend the cluster
func SufficientNodesAvailable() healthcheck.HealthCheck {
	return &DefaultWorkerHealthChecker{}
}

// SetSeedClient injects the seed client
func (healthChecker *DefaultWorkerHealthChecker) SetSeedClient(seedClient client.Client) {
	healthChecker.seedClient = seedClient
}

// SetShootClient injects the shoot client
func (healthChecker *DefaultWorkerHealthChecker) SetShootClient(shootClient client.Client) {
	healthChecker.shootClient = shootClient
}

// SetLoggerSuffix injects the logger
func (healthChecker *DefaultWorkerHealthChecker) SetLoggerSuffix(provider, extension string) {
	healthChecker.logger = log.Log.WithName(fmt.Sprintf("%s-%s-healthcheck-sufficient-nodes", provider, extension))
}

// Check executes the health check
func (healthChecker *DefaultWorkerHealthChecker) Check(ctx context.Context, request types.NamespacedName) (*healthcheck.HealthCheckResult, error) {
	machineDeploymentList := &machinev1alpha1.MachineDeploymentList{}
	// use seed seedClient
	if err := healthChecker.seedClient.List(ctx, machineDeploymentList, client.InNamespace(request.Namespace)); err != nil {
		err := fmt.Errorf("check for sufficient nodes failed. Failed to list machine deployments in namespace %s: %v'", request.Namespace, err)
		healthChecker.logger.Error(err, "Health check failed")
		return nil, err
	}

	if isHealthy, reason, err := machineDeploymentsAreHealthy(machineDeploymentList.Items); !isHealthy {
		err := fmt.Errorf("check for sufficient nodes failed: %v'", err)
		healthChecker.logger.Error(err, "Health check failed")
		return &healthcheck.HealthCheckResult{
			IsHealthy: false,
			Detail:    err.Error(),
			Reason:    *reason,
		}, nil
	}

	nodeList := &corev1.NodeList{}
	if err := healthChecker.shootClient.List(ctx, nodeList); err != nil {
		err := fmt.Errorf("check for sufficient nodes failed. Failed to list shoot nodes: %v'", err)
		healthChecker.logger.Error(err, "Health check failed")
		return nil, err
	}

	if isHealthy, reason, err := checkSufficientNodesAvailable(nodeList, machineDeploymentList); !isHealthy {
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
