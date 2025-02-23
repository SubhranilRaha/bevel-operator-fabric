/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelEtcdRaftApplyConfiguration represents an declarative configuration of the FabricMainChannelEtcdRaft type for use
// with apply.
type FabricMainChannelEtcdRaftApplyConfiguration struct {
	Options *FabricMainChannelEtcdRaftOptionsApplyConfiguration `json:"options,omitempty"`
}

// FabricMainChannelEtcdRaftApplyConfiguration constructs an declarative configuration of the FabricMainChannelEtcdRaft type for use with
// apply.
func FabricMainChannelEtcdRaft() *FabricMainChannelEtcdRaftApplyConfiguration {
	return &FabricMainChannelEtcdRaftApplyConfiguration{}
}

// WithOptions sets the Options field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Options field is set to the value of the last call.
func (b *FabricMainChannelEtcdRaftApplyConfiguration) WithOptions(value *FabricMainChannelEtcdRaftOptionsApplyConfiguration) *FabricMainChannelEtcdRaftApplyConfiguration {
	b.Options = value
	return b
}
