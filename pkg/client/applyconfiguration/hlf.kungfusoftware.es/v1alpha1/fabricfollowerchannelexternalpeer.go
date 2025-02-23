/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricFollowerChannelExternalPeerApplyConfiguration represents an declarative configuration of the FabricFollowerChannelExternalPeer type for use
// with apply.
type FabricFollowerChannelExternalPeerApplyConfiguration struct {
	URL       *string `json:"url,omitempty"`
	TLSCACert *string `json:"tlsCACert,omitempty"`
}

// FabricFollowerChannelExternalPeerApplyConfiguration constructs an declarative configuration of the FabricFollowerChannelExternalPeer type for use with
// apply.
func FabricFollowerChannelExternalPeer() *FabricFollowerChannelExternalPeerApplyConfiguration {
	return &FabricFollowerChannelExternalPeerApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *FabricFollowerChannelExternalPeerApplyConfiguration) WithURL(value string) *FabricFollowerChannelExternalPeerApplyConfiguration {
	b.URL = &value
	return b
}

// WithTLSCACert sets the TLSCACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSCACert field is set to the value of the last call.
func (b *FabricFollowerChannelExternalPeerApplyConfiguration) WithTLSCACert(value string) *FabricFollowerChannelExternalPeerApplyConfiguration {
	b.TLSCACert = &value
	return b
}
