/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
)

// FabricMainChannelStatusApplyConfiguration represents an declarative configuration of the FabricMainChannelStatus type for use
// with apply.
type FabricMainChannelStatusApplyConfiguration struct {
	Conditions *status.Conditions         `json:"conditions,omitempty"`
	Message    *string                    `json:"message,omitempty"`
	Status     *v1alpha1.DeploymentStatus `json:"status,omitempty"`
}

// FabricMainChannelStatusApplyConfiguration constructs an declarative configuration of the FabricMainChannelStatus type for use with
// apply.
func FabricMainChannelStatus() *FabricMainChannelStatusApplyConfiguration {
	return &FabricMainChannelStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricMainChannelStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricMainChannelStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FabricMainChannelStatusApplyConfiguration) WithMessage(value string) *FabricMainChannelStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricMainChannelStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricMainChannelStatusApplyConfiguration {
	b.Status = &value
	return b
}
