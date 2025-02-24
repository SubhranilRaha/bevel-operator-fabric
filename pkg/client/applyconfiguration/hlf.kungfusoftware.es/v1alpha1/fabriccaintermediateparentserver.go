/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCAIntermediateParentServerApplyConfiguration represents an declarative configuration of the FabricCAIntermediateParentServer type for use
// with apply.
type FabricCAIntermediateParentServerApplyConfiguration struct {
	URL    *string `json:"url,omitempty"`
	CAName *string `json:"caName,omitempty"`
}

// FabricCAIntermediateParentServerApplyConfiguration constructs an declarative configuration of the FabricCAIntermediateParentServer type for use with
// apply.
func FabricCAIntermediateParentServer() *FabricCAIntermediateParentServerApplyConfiguration {
	return &FabricCAIntermediateParentServerApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *FabricCAIntermediateParentServerApplyConfiguration) WithURL(value string) *FabricCAIntermediateParentServerApplyConfiguration {
	b.URL = &value
	return b
}

// WithCAName sets the CAName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CAName field is set to the value of the last call.
func (b *FabricCAIntermediateParentServerApplyConfiguration) WithCAName(value string) *FabricCAIntermediateParentServerApplyConfiguration {
	b.CAName = &value
	return b
}
