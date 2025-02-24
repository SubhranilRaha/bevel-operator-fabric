/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCACSRCAApplyConfiguration represents an declarative configuration of the FabricCACSRCA type for use
// with apply.
type FabricCACSRCAApplyConfiguration struct {
	Expiry     *string `json:"expiry,omitempty"`
	PathLength *int    `json:"pathLength,omitempty"`
}

// FabricCACSRCAApplyConfiguration constructs an declarative configuration of the FabricCACSRCA type for use with
// apply.
func FabricCACSRCA() *FabricCACSRCAApplyConfiguration {
	return &FabricCACSRCAApplyConfiguration{}
}

// WithExpiry sets the Expiry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Expiry field is set to the value of the last call.
func (b *FabricCACSRCAApplyConfiguration) WithExpiry(value string) *FabricCACSRCAApplyConfiguration {
	b.Expiry = &value
	return b
}

// WithPathLength sets the PathLength field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PathLength field is set to the value of the last call.
func (b *FabricCACSRCAApplyConfiguration) WithPathLength(value int) *FabricCACSRCAApplyConfiguration {
	b.PathLength = &value
	return b
}
