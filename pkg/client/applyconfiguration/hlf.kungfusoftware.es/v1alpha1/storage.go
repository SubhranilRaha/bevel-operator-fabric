/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// StorageApplyConfiguration represents an declarative configuration of the Storage type for use
// with apply.
type StorageApplyConfiguration struct {
	Size         *string                        `json:"size,omitempty"`
	StorageClass *string                        `json:"storageClass,omitempty"`
	AccessMode   *v1.PersistentVolumeAccessMode `json:"accessMode,omitempty"`
}

// StorageApplyConfiguration constructs an declarative configuration of the Storage type for use with
// apply.
func Storage() *StorageApplyConfiguration {
	return &StorageApplyConfiguration{}
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *StorageApplyConfiguration) WithSize(value string) *StorageApplyConfiguration {
	b.Size = &value
	return b
}

// WithStorageClass sets the StorageClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClass field is set to the value of the last call.
func (b *StorageApplyConfiguration) WithStorageClass(value string) *StorageApplyConfiguration {
	b.StorageClass = &value
	return b
}

// WithAccessMode sets the AccessMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AccessMode field is set to the value of the last call.
func (b *StorageApplyConfiguration) WithAccessMode(value v1.PersistentVolumeAccessMode) *StorageApplyConfiguration {
	b.AccessMode = &value
	return b
}
