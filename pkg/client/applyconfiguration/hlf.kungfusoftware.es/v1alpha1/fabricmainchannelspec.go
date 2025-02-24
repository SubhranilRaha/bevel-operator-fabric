/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelSpecApplyConfiguration represents an declarative configuration of the FabricMainChannelSpec type for use
// with apply.
type FabricMainChannelSpecApplyConfiguration struct {
	Name                         *string                                                           `json:"name,omitempty"`
	Identities                   map[string]FabricMainChannelIdentityApplyConfiguration            `json:"identities,omitempty"`
	AdminPeerOrganizations       []FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration    `json:"adminPeerOrganizations,omitempty"`
	PeerOrganizations            []FabricMainChannelPeerOrganizationApplyConfiguration             `json:"peerOrganizations,omitempty"`
	ExternalPeerOrganizations    []FabricMainChannelExternalPeerOrganizationApplyConfiguration     `json:"externalPeerOrganizations,omitempty"`
	ChannelConfig                *FabricMainChannelConfigApplyConfiguration                        `json:"channelConfig,omitempty"`
	AdminOrdererOrganizations    []FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration `json:"adminOrdererOrganizations,omitempty"`
	OrdererOrganizations         []FabricMainChannelOrdererOrganizationApplyConfiguration          `json:"ordererOrganizations,omitempty"`
	ExternalOrdererOrganizations []FabricMainChannelExternalOrdererOrganizationApplyConfiguration  `json:"externalOrdererOrganizations,omitempty"`
	Consenters                   []FabricMainChannelConsenterApplyConfiguration                    `json:"orderers,omitempty"`
}

// FabricMainChannelSpecApplyConfiguration constructs an declarative configuration of the FabricMainChannelSpec type for use with
// apply.
func FabricMainChannelSpec() *FabricMainChannelSpecApplyConfiguration {
	return &FabricMainChannelSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *FabricMainChannelSpecApplyConfiguration) WithName(value string) *FabricMainChannelSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithIdentities puts the entries into the Identities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Identities field,
// overwriting an existing map entries in Identities field with the same key.
func (b *FabricMainChannelSpecApplyConfiguration) WithIdentities(entries map[string]FabricMainChannelIdentityApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	if b.Identities == nil && len(entries) > 0 {
		b.Identities = make(map[string]FabricMainChannelIdentityApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Identities[k] = v
	}
	return b
}

// WithAdminPeerOrganizations adds the given value to the AdminPeerOrganizations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdminPeerOrganizations field.
func (b *FabricMainChannelSpecApplyConfiguration) WithAdminPeerOrganizations(values ...*FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdminPeerOrganizations")
		}
		b.AdminPeerOrganizations = append(b.AdminPeerOrganizations, *values[i])
	}
	return b
}

// WithPeerOrganizations adds the given value to the PeerOrganizations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PeerOrganizations field.
func (b *FabricMainChannelSpecApplyConfiguration) WithPeerOrganizations(values ...*FabricMainChannelPeerOrganizationApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPeerOrganizations")
		}
		b.PeerOrganizations = append(b.PeerOrganizations, *values[i])
	}
	return b
}

// WithExternalPeerOrganizations adds the given value to the ExternalPeerOrganizations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalPeerOrganizations field.
func (b *FabricMainChannelSpecApplyConfiguration) WithExternalPeerOrganizations(values ...*FabricMainChannelExternalPeerOrganizationApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExternalPeerOrganizations")
		}
		b.ExternalPeerOrganizations = append(b.ExternalPeerOrganizations, *values[i])
	}
	return b
}

// WithChannelConfig sets the ChannelConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ChannelConfig field is set to the value of the last call.
func (b *FabricMainChannelSpecApplyConfiguration) WithChannelConfig(value *FabricMainChannelConfigApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	b.ChannelConfig = value
	return b
}

// WithAdminOrdererOrganizations adds the given value to the AdminOrdererOrganizations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdminOrdererOrganizations field.
func (b *FabricMainChannelSpecApplyConfiguration) WithAdminOrdererOrganizations(values ...*FabricMainChannelAdminOrdererOrganizationSpecApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdminOrdererOrganizations")
		}
		b.AdminOrdererOrganizations = append(b.AdminOrdererOrganizations, *values[i])
	}
	return b
}

// WithOrdererOrganizations adds the given value to the OrdererOrganizations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OrdererOrganizations field.
func (b *FabricMainChannelSpecApplyConfiguration) WithOrdererOrganizations(values ...*FabricMainChannelOrdererOrganizationApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOrdererOrganizations")
		}
		b.OrdererOrganizations = append(b.OrdererOrganizations, *values[i])
	}
	return b
}

// WithExternalOrdererOrganizations adds the given value to the ExternalOrdererOrganizations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalOrdererOrganizations field.
func (b *FabricMainChannelSpecApplyConfiguration) WithExternalOrdererOrganizations(values ...*FabricMainChannelExternalOrdererOrganizationApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExternalOrdererOrganizations")
		}
		b.ExternalOrdererOrganizations = append(b.ExternalOrdererOrganizations, *values[i])
	}
	return b
}

// WithConsenters adds the given value to the Consenters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Consenters field.
func (b *FabricMainChannelSpecApplyConfiguration) WithConsenters(values ...*FabricMainChannelConsenterApplyConfiguration) *FabricMainChannelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConsenters")
		}
		b.Consenters = append(b.Consenters, *values[i])
	}
	return b
}
