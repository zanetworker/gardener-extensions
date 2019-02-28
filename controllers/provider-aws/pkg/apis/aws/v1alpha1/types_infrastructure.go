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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

const (
	// EventReasonDestruction an event describing infrastructure destruction
	EventReasonDestruction string = "InfrastructureDestruction"
	// EventReasonCreation an event describing infrastructure creation
	EventReasonCreation string = "InfrastructureCreation"

	// PurposeNodes is the purpose for nodes
	PurposeNodes string = "nodes"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InfrastructureConfig infrastructure configuration resource
type InfrastructureConfig struct {
	metav1.TypeMeta `json:",inline"`
	Networks        Networks `json:"networks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InfrastructureStatus contains information about created infrastructure resources.
type InfrastructureStatus struct {
	metav1.TypeMeta `json:",inline"`

	// EC2 contains information about the created AWS EC2 resources.
	EC2 EC2 `json:"ec2"`
	// IAM contains information about the created AWS IAM resources.
	IAM IAM `json:"iam"`
	// VPC contains information about the created AWS VPC and some related resources.
	VPC VPCStatus `json:"vpc"`
}

// CIDR is a string alias.
type CIDR string

// Networks holds information about the Kubernetes and infrastructure networks.
type Networks struct {
	// VPC indicates whether to use an existing VPC or create a new one.
	VPC VPC `json:"vpc"`
	// Zones belonging to the same region
	Zones []Zone `json:"zones"`
}

// Zone describes the properties of a zone
type Zone struct {
	// Name is the name for this zone.
	Name string `json:"name"`
	// Internal is  the  private subnet range to create (used for internal load balancers).
	Internal CIDR `json:"internal"`
	// Public is the  public subnet range to create (used for bastion and load balancers).
	Public CIDR `json:"public"`
	// Workers is the  workers  subnet range  to create (used for the VMs).
	Workers CIDR `json:"workers"`
}

// EC2 contains information about the  AWS EC2 resources.
type EC2 struct {
	// KeyName is the name of the SSH key.
	KeyName string `json:"keyName"`
}

// IAM contains information about the AWS IAM resources.
type IAM struct {
	// InstanceProfiles is a list of AWS IAM instance profiles.
	InstanceProfiles []InstanceProfile `json:"instanceProfiles"`
	// Roles is a list of AWS IAM roles.
	Roles []Role `json:"roles"`
}

// VPC contains information about the AWS VPC and some related resources.
type VPC struct {
	// ID is the VPC id.
	// +optional
	ID *string `json:"id,omitempty"`
	// CIDR is the VPC CIDR
	// +optional
	CIDR *CIDR `json:"cidr,omitempty"`
}

// VPCStatus vpc operation results that will be part of the status
type VPCStatus struct {
	// ID is the VPC id.
	ID string `json:"id"`
	// Subnets is a list of subnets that have been created.
	Subnets []Subnet `json:"subnets"`
	// SecurityGroups is a list of security groups that have been created.
	SecurityGroups []SecurityGroup `json:"securityGroups"`
}

// InstanceProfile is an AWS IAM instance profile.
type InstanceProfile struct {
	// Purpose is a logical description of the instance profile.
	// +optional
	Purpose *string `json:"purpose"`
	// Name is the name for this instance profile.
	Name string `json:"name"`
}

// Role is an AWS IAM role.
type Role struct {
	// Purpose is a logical description of the role.
	// +optional
	Purpose *string `json:"purpose"`
	// ARN is the AWS Resource Name for this role.
	ARN string `json:"arn"`
}

// Subnet is an AWS subnet related to a VPC.
type Subnet struct {
	// Name is a logical name of the subnet.
	Name string `json:"name"`
	// ID is the subnet id.
	ID string `json:"id"`
	// Zone is the availability zone into which the subnet has been created.
	Zone string `json:"zone"`
}

// SecurityGroup is an AWS security group related to a VPC.
type SecurityGroup struct {
	// Purpose is a logical description of the security group.
	// +optional
	Purpose *string `json:"purpose"`
	// Name is a logical name of the subnet.
	Name string `json:"name"`
	// ID is the subnet id.
	ID string `json:"id"`
}
