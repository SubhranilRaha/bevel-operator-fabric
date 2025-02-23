/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricOperatorAPIHLFConfigApplyConfiguration represents an declarative configuration of the FabricOperatorAPIHLFConfig type for use
// with apply.
type FabricOperatorAPIHLFConfigApplyConfiguration struct {
	MSPID         *string                                           `json:"mspID,omitempty"`
	User          *string                                           `json:"user,omitempty"`
	NetworkConfig *FabricOperatorAPINetworkConfigApplyConfiguration `json:"networkConfig,omitempty"`
}

// FabricOperatorAPIHLFConfigApplyConfiguration constructs an declarative configuration of the FabricOperatorAPIHLFConfig type for use with
// apply.
func FabricOperatorAPIHLFConfig() *FabricOperatorAPIHLFConfigApplyConfiguration {
	return &FabricOperatorAPIHLFConfigApplyConfiguration{}
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricOperatorAPIHLFConfigApplyConfiguration) WithMSPID(value string) *FabricOperatorAPIHLFConfigApplyConfiguration {
	b.MSPID = &value
	return b
}

// WithUser sets the User field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the User field is set to the value of the last call.
func (b *FabricOperatorAPIHLFConfigApplyConfiguration) WithUser(value string) *FabricOperatorAPIHLFConfigApplyConfiguration {
	b.User = &value
	return b
}

// WithNetworkConfig sets the NetworkConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkConfig field is set to the value of the last call.
func (b *FabricOperatorAPIHLFConfigApplyConfiguration) WithNetworkConfig(value *FabricOperatorAPINetworkConfigApplyConfiguration) *FabricOperatorAPIHLFConfigApplyConfiguration {
	b.NetworkConfig = value
	return b
}
