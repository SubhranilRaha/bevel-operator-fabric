/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCASubjectApplyConfiguration represents an declarative configuration of the FabricCASubject type for use
// with apply.
type FabricCASubjectApplyConfiguration struct {
	CN *string `json:"cn,omitempty"`
	C  *string `json:"C,omitempty"`
	ST *string `json:"ST,omitempty"`
	O  *string `json:"O,omitempty"`
	L  *string `json:"L,omitempty"`
	OU *string `json:"OU,omitempty"`
}

// FabricCASubjectApplyConfiguration constructs an declarative configuration of the FabricCASubject type for use with
// apply.
func FabricCASubject() *FabricCASubjectApplyConfiguration {
	return &FabricCASubjectApplyConfiguration{}
}

// WithCN sets the CN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CN field is set to the value of the last call.
func (b *FabricCASubjectApplyConfiguration) WithCN(value string) *FabricCASubjectApplyConfiguration {
	b.CN = &value
	return b
}

// WithC sets the C field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the C field is set to the value of the last call.
func (b *FabricCASubjectApplyConfiguration) WithC(value string) *FabricCASubjectApplyConfiguration {
	b.C = &value
	return b
}

// WithST sets the ST field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ST field is set to the value of the last call.
func (b *FabricCASubjectApplyConfiguration) WithST(value string) *FabricCASubjectApplyConfiguration {
	b.ST = &value
	return b
}

// WithO sets the O field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the O field is set to the value of the last call.
func (b *FabricCASubjectApplyConfiguration) WithO(value string) *FabricCASubjectApplyConfiguration {
	b.O = &value
	return b
}

// WithL sets the L field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the L field is set to the value of the last call.
func (b *FabricCASubjectApplyConfiguration) WithL(value string) *FabricCASubjectApplyConfiguration {
	b.L = &value
	return b
}

// WithOU sets the OU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OU field is set to the value of the last call.
func (b *FabricCASubjectApplyConfiguration) WithOU(value string) *FabricCASubjectApplyConfiguration {
	b.OU = &value
	return b
}
