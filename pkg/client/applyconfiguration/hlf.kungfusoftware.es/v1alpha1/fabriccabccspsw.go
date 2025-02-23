/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCABCCSPSWApplyConfiguration represents an declarative configuration of the FabricCABCCSPSW type for use
// with apply.
type FabricCABCCSPSWApplyConfiguration struct {
	Hash     *string `json:"hash,omitempty"`
	Security *string `json:"security,omitempty"`
}

// FabricCABCCSPSWApplyConfiguration constructs an declarative configuration of the FabricCABCCSPSW type for use with
// apply.
func FabricCABCCSPSW() *FabricCABCCSPSWApplyConfiguration {
	return &FabricCABCCSPSWApplyConfiguration{}
}

// WithHash sets the Hash field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hash field is set to the value of the last call.
func (b *FabricCABCCSPSWApplyConfiguration) WithHash(value string) *FabricCABCCSPSWApplyConfiguration {
	b.Hash = &value
	return b
}

// WithSecurity sets the Security field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Security field is set to the value of the last call.
func (b *FabricCABCCSPSWApplyConfiguration) WithSecurity(value string) *FabricCABCCSPSWApplyConfiguration {
	b.Security = &value
	return b
}
