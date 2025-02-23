/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCASigningApplyConfiguration represents an declarative configuration of the FabricCASigning type for use
// with apply.
type FabricCASigningApplyConfiguration struct {
	Default  *FabricCASigningDefaultApplyConfiguration  `json:"default,omitempty"`
	Profiles *FabricCASigningProfilesApplyConfiguration `json:"profiles,omitempty"`
}

// FabricCASigningApplyConfiguration constructs an declarative configuration of the FabricCASigning type for use with
// apply.
func FabricCASigning() *FabricCASigningApplyConfiguration {
	return &FabricCASigningApplyConfiguration{}
}

// WithDefault sets the Default field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Default field is set to the value of the last call.
func (b *FabricCASigningApplyConfiguration) WithDefault(value *FabricCASigningDefaultApplyConfiguration) *FabricCASigningApplyConfiguration {
	b.Default = value
	return b
}

// WithProfiles sets the Profiles field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Profiles field is set to the value of the last call.
func (b *FabricCASigningApplyConfiguration) WithProfiles(value *FabricCASigningProfilesApplyConfiguration) *FabricCASigningApplyConfiguration {
	b.Profiles = value
	return b
}
