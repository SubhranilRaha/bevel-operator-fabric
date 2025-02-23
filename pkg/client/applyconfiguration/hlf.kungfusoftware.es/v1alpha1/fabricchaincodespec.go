/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricChaincodeSpecApplyConfiguration represents an declarative configuration of the FabricChaincodeSpec type for use
// with apply.
type FabricChaincodeSpecApplyConfiguration struct {
	Image            *string                   `json:"image,omitempty"`
	ImagePullPolicy  *v1.PullPolicy            `json:"imagePullPolicy,omitempty"`
	PackageID        *string                   `json:"packageId,omitempty"`
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Affinity         *v1.Affinity              `json:"affinity,omitempty"`
	Tolerations      []v1.Toleration           `json:"tolerations,omitempty"`
	Resources        *v1.ResourceRequirements  `json:"resources,omitempty"`
	Credentials      *TLSApplyConfiguration    `json:"credentials,omitempty"`
	Replicas         *int                      `json:"replicas,omitempty"`
	Env              []v1.EnvVar               `json:"env,omitempty"`
}

// FabricChaincodeSpecApplyConfiguration constructs an declarative configuration of the FabricChaincodeSpec type for use with
// apply.
func FabricChaincodeSpec() *FabricChaincodeSpecApplyConfiguration {
	return &FabricChaincodeSpecApplyConfiguration{}
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *FabricChaincodeSpecApplyConfiguration) WithImage(value string) *FabricChaincodeSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullPolicy field is set to the value of the last call.
func (b *FabricChaincodeSpecApplyConfiguration) WithImagePullPolicy(value v1.PullPolicy) *FabricChaincodeSpecApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}

// WithPackageID sets the PackageID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PackageID field is set to the value of the last call.
func (b *FabricChaincodeSpecApplyConfiguration) WithPackageID(value string) *FabricChaincodeSpecApplyConfiguration {
	b.PackageID = &value
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *FabricChaincodeSpecApplyConfiguration) WithImagePullSecrets(values ...v1.LocalObjectReference) *FabricChaincodeSpecApplyConfiguration {
	for i := range values {
		b.ImagePullSecrets = append(b.ImagePullSecrets, values[i])
	}
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *FabricChaincodeSpecApplyConfiguration) WithAffinity(value v1.Affinity) *FabricChaincodeSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *FabricChaincodeSpecApplyConfiguration) WithTolerations(values ...v1.Toleration) *FabricChaincodeSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *FabricChaincodeSpecApplyConfiguration) WithResources(value v1.ResourceRequirements) *FabricChaincodeSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithCredentials sets the Credentials field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Credentials field is set to the value of the last call.
func (b *FabricChaincodeSpecApplyConfiguration) WithCredentials(value *TLSApplyConfiguration) *FabricChaincodeSpecApplyConfiguration {
	b.Credentials = value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *FabricChaincodeSpecApplyConfiguration) WithReplicas(value int) *FabricChaincodeSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *FabricChaincodeSpecApplyConfiguration) WithEnv(values ...v1.EnvVar) *FabricChaincodeSpecApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}
