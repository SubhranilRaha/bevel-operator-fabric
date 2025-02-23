/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricOperatorUIAuthApplyConfiguration represents an declarative configuration of the FabricOperatorUIAuth type for use
// with apply.
type FabricOperatorUIAuthApplyConfiguration struct {
	OIDCAuthority *string `json:"oidcAuthority,omitempty"`
	OIDCClientId  *string `json:"oidcClientId,omitempty"`
	OIDCScope     *string `json:"oidcScope,omitempty"`
}

// FabricOperatorUIAuthApplyConfiguration constructs an declarative configuration of the FabricOperatorUIAuth type for use with
// apply.
func FabricOperatorUIAuth() *FabricOperatorUIAuthApplyConfiguration {
	return &FabricOperatorUIAuthApplyConfiguration{}
}

// WithOIDCAuthority sets the OIDCAuthority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCAuthority field is set to the value of the last call.
func (b *FabricOperatorUIAuthApplyConfiguration) WithOIDCAuthority(value string) *FabricOperatorUIAuthApplyConfiguration {
	b.OIDCAuthority = &value
	return b
}

// WithOIDCClientId sets the OIDCClientId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCClientId field is set to the value of the last call.
func (b *FabricOperatorUIAuthApplyConfiguration) WithOIDCClientId(value string) *FabricOperatorUIAuthApplyConfiguration {
	b.OIDCClientId = &value
	return b
}

// WithOIDCScope sets the OIDCScope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCScope field is set to the value of the last call.
func (b *FabricOperatorUIAuthApplyConfiguration) WithOIDCScope(value string) *FabricOperatorUIAuthApplyConfiguration {
	b.OIDCScope = &value
	return b
}
