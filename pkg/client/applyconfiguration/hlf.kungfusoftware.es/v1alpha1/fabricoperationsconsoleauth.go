/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricOperationsConsoleAuthApplyConfiguration represents an declarative configuration of the FabricOperationsConsoleAuth type for use
// with apply.
type FabricOperationsConsoleAuthApplyConfiguration struct {
	Scheme   *string `json:"scheme,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// FabricOperationsConsoleAuthApplyConfiguration constructs an declarative configuration of the FabricOperationsConsoleAuth type for use with
// apply.
func FabricOperationsConsoleAuth() *FabricOperationsConsoleAuthApplyConfiguration {
	return &FabricOperationsConsoleAuthApplyConfiguration{}
}

// WithScheme sets the Scheme field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scheme field is set to the value of the last call.
func (b *FabricOperationsConsoleAuthApplyConfiguration) WithScheme(value string) *FabricOperationsConsoleAuthApplyConfiguration {
	b.Scheme = &value
	return b
}

// WithUsername sets the Username field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Username field is set to the value of the last call.
func (b *FabricOperationsConsoleAuthApplyConfiguration) WithUsername(value string) *FabricOperationsConsoleAuthApplyConfiguration {
	b.Username = &value
	return b
}

// WithPassword sets the Password field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Password field is set to the value of the last call.
func (b *FabricOperationsConsoleAuthApplyConfiguration) WithPassword(value string) *FabricOperationsConsoleAuthApplyConfiguration {
	b.Password = &value
	return b
}
