// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	aws "github.com/gardener/gardener-extensions/controllers/provider-aws/pkg/apis/aws"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*EC2)(nil), (*aws.EC2)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EC2_To_aws_EC2(a.(*EC2), b.(*aws.EC2), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.EC2)(nil), (*EC2)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_EC2_To_v1alpha1_EC2(a.(*aws.EC2), b.(*EC2), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IAM)(nil), (*aws.IAM)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IAM_To_aws_IAM(a.(*IAM), b.(*aws.IAM), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.IAM)(nil), (*IAM)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_IAM_To_v1alpha1_IAM(a.(*aws.IAM), b.(*IAM), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureConfig)(nil), (*aws.InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureConfig_To_aws_InfrastructureConfig(a.(*InfrastructureConfig), b.(*aws.InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.InfrastructureConfig)(nil), (*InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(a.(*aws.InfrastructureConfig), b.(*InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureStatus)(nil), (*aws.InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureStatus_To_aws_InfrastructureStatus(a.(*InfrastructureStatus), b.(*aws.InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.InfrastructureStatus)(nil), (*InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(a.(*aws.InfrastructureStatus), b.(*InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InstanceProfile)(nil), (*aws.InstanceProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InstanceProfile_To_aws_InstanceProfile(a.(*InstanceProfile), b.(*aws.InstanceProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.InstanceProfile)(nil), (*InstanceProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_InstanceProfile_To_v1alpha1_InstanceProfile(a.(*aws.InstanceProfile), b.(*InstanceProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Networks)(nil), (*aws.Networks)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Networks_To_aws_Networks(a.(*Networks), b.(*aws.Networks), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.Networks)(nil), (*Networks)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_Networks_To_v1alpha1_Networks(a.(*aws.Networks), b.(*Networks), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Role)(nil), (*aws.Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Role_To_aws_Role(a.(*Role), b.(*aws.Role), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.Role)(nil), (*Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_Role_To_v1alpha1_Role(a.(*aws.Role), b.(*Role), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SecurityGroup)(nil), (*aws.SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SecurityGroup_To_aws_SecurityGroup(a.(*SecurityGroup), b.(*aws.SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.SecurityGroup)(nil), (*SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_SecurityGroup_To_v1alpha1_SecurityGroup(a.(*aws.SecurityGroup), b.(*SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Subnet)(nil), (*aws.Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Subnet_To_aws_Subnet(a.(*Subnet), b.(*aws.Subnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.Subnet)(nil), (*Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_Subnet_To_v1alpha1_Subnet(a.(*aws.Subnet), b.(*Subnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VPC)(nil), (*aws.VPC)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VPC_To_aws_VPC(a.(*VPC), b.(*aws.VPC), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.VPC)(nil), (*VPC)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_VPC_To_v1alpha1_VPC(a.(*aws.VPC), b.(*VPC), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VPCStatus)(nil), (*aws.VPCStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VPCStatus_To_aws_VPCStatus(a.(*VPCStatus), b.(*aws.VPCStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.VPCStatus)(nil), (*VPCStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_VPCStatus_To_v1alpha1_VPCStatus(a.(*aws.VPCStatus), b.(*VPCStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Zone)(nil), (*aws.Zone)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Zone_To_aws_Zone(a.(*Zone), b.(*aws.Zone), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aws.Zone)(nil), (*Zone)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aws_Zone_To_v1alpha1_Zone(a.(*aws.Zone), b.(*Zone), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_EC2_To_aws_EC2(in *EC2, out *aws.EC2, s conversion.Scope) error {
	out.KeyName = in.KeyName
	return nil
}

// Convert_v1alpha1_EC2_To_aws_EC2 is an autogenerated conversion function.
func Convert_v1alpha1_EC2_To_aws_EC2(in *EC2, out *aws.EC2, s conversion.Scope) error {
	return autoConvert_v1alpha1_EC2_To_aws_EC2(in, out, s)
}

func autoConvert_aws_EC2_To_v1alpha1_EC2(in *aws.EC2, out *EC2, s conversion.Scope) error {
	out.KeyName = in.KeyName
	return nil
}

// Convert_aws_EC2_To_v1alpha1_EC2 is an autogenerated conversion function.
func Convert_aws_EC2_To_v1alpha1_EC2(in *aws.EC2, out *EC2, s conversion.Scope) error {
	return autoConvert_aws_EC2_To_v1alpha1_EC2(in, out, s)
}

func autoConvert_v1alpha1_IAM_To_aws_IAM(in *IAM, out *aws.IAM, s conversion.Scope) error {
	out.InstanceProfiles = *(*[]aws.InstanceProfile)(unsafe.Pointer(&in.InstanceProfiles))
	out.Roles = *(*[]aws.Role)(unsafe.Pointer(&in.Roles))
	return nil
}

// Convert_v1alpha1_IAM_To_aws_IAM is an autogenerated conversion function.
func Convert_v1alpha1_IAM_To_aws_IAM(in *IAM, out *aws.IAM, s conversion.Scope) error {
	return autoConvert_v1alpha1_IAM_To_aws_IAM(in, out, s)
}

func autoConvert_aws_IAM_To_v1alpha1_IAM(in *aws.IAM, out *IAM, s conversion.Scope) error {
	out.InstanceProfiles = *(*[]InstanceProfile)(unsafe.Pointer(&in.InstanceProfiles))
	out.Roles = *(*[]Role)(unsafe.Pointer(&in.Roles))
	return nil
}

// Convert_aws_IAM_To_v1alpha1_IAM is an autogenerated conversion function.
func Convert_aws_IAM_To_v1alpha1_IAM(in *aws.IAM, out *IAM, s conversion.Scope) error {
	return autoConvert_aws_IAM_To_v1alpha1_IAM(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureConfig_To_aws_InfrastructureConfig(in *InfrastructureConfig, out *aws.InfrastructureConfig, s conversion.Scope) error {
	if err := Convert_v1alpha1_Networks_To_aws_Networks(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_InfrastructureConfig_To_aws_InfrastructureConfig is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureConfig_To_aws_InfrastructureConfig(in *InfrastructureConfig, out *aws.InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureConfig_To_aws_InfrastructureConfig(in, out, s)
}

func autoConvert_aws_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *aws.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	if err := Convert_aws_Networks_To_v1alpha1_Networks(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_aws_InfrastructureConfig_To_v1alpha1_InfrastructureConfig is an autogenerated conversion function.
func Convert_aws_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *aws.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_aws_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureStatus_To_aws_InfrastructureStatus(in *InfrastructureStatus, out *aws.InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_EC2_To_aws_EC2(&in.EC2, &out.EC2, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_IAM_To_aws_IAM(&in.IAM, &out.IAM, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VPCStatus_To_aws_VPCStatus(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_InfrastructureStatus_To_aws_InfrastructureStatus is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureStatus_To_aws_InfrastructureStatus(in *InfrastructureStatus, out *aws.InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureStatus_To_aws_InfrastructureStatus(in, out, s)
}

func autoConvert_aws_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *aws.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_aws_EC2_To_v1alpha1_EC2(&in.EC2, &out.EC2, s); err != nil {
		return err
	}
	if err := Convert_aws_IAM_To_v1alpha1_IAM(&in.IAM, &out.IAM, s); err != nil {
		return err
	}
	if err := Convert_aws_VPCStatus_To_v1alpha1_VPCStatus(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	return nil
}

// Convert_aws_InfrastructureStatus_To_v1alpha1_InfrastructureStatus is an autogenerated conversion function.
func Convert_aws_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *aws.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_aws_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in, out, s)
}

func autoConvert_v1alpha1_InstanceProfile_To_aws_InstanceProfile(in *InstanceProfile, out *aws.InstanceProfile, s conversion.Scope) error {
	out.Purpose = (*string)(unsafe.Pointer(in.Purpose))
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_InstanceProfile_To_aws_InstanceProfile is an autogenerated conversion function.
func Convert_v1alpha1_InstanceProfile_To_aws_InstanceProfile(in *InstanceProfile, out *aws.InstanceProfile, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceProfile_To_aws_InstanceProfile(in, out, s)
}

func autoConvert_aws_InstanceProfile_To_v1alpha1_InstanceProfile(in *aws.InstanceProfile, out *InstanceProfile, s conversion.Scope) error {
	out.Purpose = (*string)(unsafe.Pointer(in.Purpose))
	out.Name = in.Name
	return nil
}

// Convert_aws_InstanceProfile_To_v1alpha1_InstanceProfile is an autogenerated conversion function.
func Convert_aws_InstanceProfile_To_v1alpha1_InstanceProfile(in *aws.InstanceProfile, out *InstanceProfile, s conversion.Scope) error {
	return autoConvert_aws_InstanceProfile_To_v1alpha1_InstanceProfile(in, out, s)
}

func autoConvert_v1alpha1_Networks_To_aws_Networks(in *Networks, out *aws.Networks, s conversion.Scope) error {
	if err := Convert_v1alpha1_VPC_To_aws_VPC(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	out.Zones = *(*[]aws.Zone)(unsafe.Pointer(&in.Zones))
	return nil
}

// Convert_v1alpha1_Networks_To_aws_Networks is an autogenerated conversion function.
func Convert_v1alpha1_Networks_To_aws_Networks(in *Networks, out *aws.Networks, s conversion.Scope) error {
	return autoConvert_v1alpha1_Networks_To_aws_Networks(in, out, s)
}

func autoConvert_aws_Networks_To_v1alpha1_Networks(in *aws.Networks, out *Networks, s conversion.Scope) error {
	if err := Convert_aws_VPC_To_v1alpha1_VPC(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	out.Zones = *(*[]Zone)(unsafe.Pointer(&in.Zones))
	return nil
}

// Convert_aws_Networks_To_v1alpha1_Networks is an autogenerated conversion function.
func Convert_aws_Networks_To_v1alpha1_Networks(in *aws.Networks, out *Networks, s conversion.Scope) error {
	return autoConvert_aws_Networks_To_v1alpha1_Networks(in, out, s)
}

func autoConvert_v1alpha1_Role_To_aws_Role(in *Role, out *aws.Role, s conversion.Scope) error {
	out.Purpose = (*string)(unsafe.Pointer(in.Purpose))
	out.ARN = in.ARN
	return nil
}

// Convert_v1alpha1_Role_To_aws_Role is an autogenerated conversion function.
func Convert_v1alpha1_Role_To_aws_Role(in *Role, out *aws.Role, s conversion.Scope) error {
	return autoConvert_v1alpha1_Role_To_aws_Role(in, out, s)
}

func autoConvert_aws_Role_To_v1alpha1_Role(in *aws.Role, out *Role, s conversion.Scope) error {
	out.Purpose = (*string)(unsafe.Pointer(in.Purpose))
	out.ARN = in.ARN
	return nil
}

// Convert_aws_Role_To_v1alpha1_Role is an autogenerated conversion function.
func Convert_aws_Role_To_v1alpha1_Role(in *aws.Role, out *Role, s conversion.Scope) error {
	return autoConvert_aws_Role_To_v1alpha1_Role(in, out, s)
}

func autoConvert_v1alpha1_SecurityGroup_To_aws_SecurityGroup(in *SecurityGroup, out *aws.SecurityGroup, s conversion.Scope) error {
	if err := v1.Convert_Pointer_string_To_string(&in.Purpose, &out.Purpose, s); err != nil {
		return err
	}
	out.Name = in.Name
	out.ID = in.ID
	return nil
}

// Convert_v1alpha1_SecurityGroup_To_aws_SecurityGroup is an autogenerated conversion function.
func Convert_v1alpha1_SecurityGroup_To_aws_SecurityGroup(in *SecurityGroup, out *aws.SecurityGroup, s conversion.Scope) error {
	return autoConvert_v1alpha1_SecurityGroup_To_aws_SecurityGroup(in, out, s)
}

func autoConvert_aws_SecurityGroup_To_v1alpha1_SecurityGroup(in *aws.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	if err := v1.Convert_string_To_Pointer_string(&in.Purpose, &out.Purpose, s); err != nil {
		return err
	}
	out.Name = in.Name
	out.ID = in.ID
	return nil
}

// Convert_aws_SecurityGroup_To_v1alpha1_SecurityGroup is an autogenerated conversion function.
func Convert_aws_SecurityGroup_To_v1alpha1_SecurityGroup(in *aws.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	return autoConvert_aws_SecurityGroup_To_v1alpha1_SecurityGroup(in, out, s)
}

func autoConvert_v1alpha1_Subnet_To_aws_Subnet(in *Subnet, out *aws.Subnet, s conversion.Scope) error {
	out.Name = in.Name
	out.ID = in.ID
	out.Zone = in.Zone
	return nil
}

// Convert_v1alpha1_Subnet_To_aws_Subnet is an autogenerated conversion function.
func Convert_v1alpha1_Subnet_To_aws_Subnet(in *Subnet, out *aws.Subnet, s conversion.Scope) error {
	return autoConvert_v1alpha1_Subnet_To_aws_Subnet(in, out, s)
}

func autoConvert_aws_Subnet_To_v1alpha1_Subnet(in *aws.Subnet, out *Subnet, s conversion.Scope) error {
	out.Name = in.Name
	out.ID = in.ID
	out.Zone = in.Zone
	return nil
}

// Convert_aws_Subnet_To_v1alpha1_Subnet is an autogenerated conversion function.
func Convert_aws_Subnet_To_v1alpha1_Subnet(in *aws.Subnet, out *Subnet, s conversion.Scope) error {
	return autoConvert_aws_Subnet_To_v1alpha1_Subnet(in, out, s)
}

func autoConvert_v1alpha1_VPC_To_aws_VPC(in *VPC, out *aws.VPC, s conversion.Scope) error {
	if err := v1.Convert_Pointer_string_To_string(&in.ID, &out.ID, s); err != nil {
		return err
	}
	out.CIDR = (*aws.CIDR)(unsafe.Pointer(in.CIDR))
	return nil
}

// Convert_v1alpha1_VPC_To_aws_VPC is an autogenerated conversion function.
func Convert_v1alpha1_VPC_To_aws_VPC(in *VPC, out *aws.VPC, s conversion.Scope) error {
	return autoConvert_v1alpha1_VPC_To_aws_VPC(in, out, s)
}

func autoConvert_aws_VPC_To_v1alpha1_VPC(in *aws.VPC, out *VPC, s conversion.Scope) error {
	if err := v1.Convert_string_To_Pointer_string(&in.ID, &out.ID, s); err != nil {
		return err
	}
	out.CIDR = (*CIDR)(unsafe.Pointer(in.CIDR))
	return nil
}

// Convert_aws_VPC_To_v1alpha1_VPC is an autogenerated conversion function.
func Convert_aws_VPC_To_v1alpha1_VPC(in *aws.VPC, out *VPC, s conversion.Scope) error {
	return autoConvert_aws_VPC_To_v1alpha1_VPC(in, out, s)
}

func autoConvert_v1alpha1_VPCStatus_To_aws_VPCStatus(in *VPCStatus, out *aws.VPCStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.Subnets = *(*[]aws.Subnet)(unsafe.Pointer(&in.Subnets))
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]aws.SecurityGroup, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_SecurityGroup_To_aws_SecurityGroup(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.SecurityGroups = nil
	}
	return nil
}

// Convert_v1alpha1_VPCStatus_To_aws_VPCStatus is an autogenerated conversion function.
func Convert_v1alpha1_VPCStatus_To_aws_VPCStatus(in *VPCStatus, out *aws.VPCStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VPCStatus_To_aws_VPCStatus(in, out, s)
}

func autoConvert_aws_VPCStatus_To_v1alpha1_VPCStatus(in *aws.VPCStatus, out *VPCStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.Subnets = *(*[]Subnet)(unsafe.Pointer(&in.Subnets))
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]SecurityGroup, len(*in))
		for i := range *in {
			if err := Convert_aws_SecurityGroup_To_v1alpha1_SecurityGroup(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.SecurityGroups = nil
	}
	return nil
}

// Convert_aws_VPCStatus_To_v1alpha1_VPCStatus is an autogenerated conversion function.
func Convert_aws_VPCStatus_To_v1alpha1_VPCStatus(in *aws.VPCStatus, out *VPCStatus, s conversion.Scope) error {
	return autoConvert_aws_VPCStatus_To_v1alpha1_VPCStatus(in, out, s)
}

func autoConvert_v1alpha1_Zone_To_aws_Zone(in *Zone, out *aws.Zone, s conversion.Scope) error {
	out.Name = in.Name
	out.Internal = aws.CIDR(in.Internal)
	out.Public = aws.CIDR(in.Public)
	out.Workers = aws.CIDR(in.Workers)
	return nil
}

// Convert_v1alpha1_Zone_To_aws_Zone is an autogenerated conversion function.
func Convert_v1alpha1_Zone_To_aws_Zone(in *Zone, out *aws.Zone, s conversion.Scope) error {
	return autoConvert_v1alpha1_Zone_To_aws_Zone(in, out, s)
}

func autoConvert_aws_Zone_To_v1alpha1_Zone(in *aws.Zone, out *Zone, s conversion.Scope) error {
	out.Name = in.Name
	out.Internal = CIDR(in.Internal)
	out.Public = CIDR(in.Public)
	out.Workers = CIDR(in.Workers)
	return nil
}

// Convert_aws_Zone_To_v1alpha1_Zone is an autogenerated conversion function.
func Convert_aws_Zone_To_v1alpha1_Zone(in *aws.Zone, out *Zone, s conversion.Scope) error {
	return autoConvert_aws_Zone_To_v1alpha1_Zone(in, out, s)
}
