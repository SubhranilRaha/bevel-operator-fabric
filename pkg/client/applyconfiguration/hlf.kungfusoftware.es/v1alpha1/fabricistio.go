/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricIstioApplyConfiguration represents an declarative configuration of the FabricIstio type for use
// with apply.
type FabricIstioApplyConfiguration struct {
	Port           *int     `json:"port,omitempty"`
	Hosts          []string `json:"hosts,omitempty"`
	IngressGateway *string  `json:"ingressGateway,omitempty"`
}

// FabricIstioApplyConfiguration constructs an declarative configuration of the FabricIstio type for use with
// apply.
func FabricIstio() *FabricIstioApplyConfiguration {
	return &FabricIstioApplyConfiguration{}
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *FabricIstioApplyConfiguration) WithPort(value int) *FabricIstioApplyConfiguration {
	b.Port = &value
	return b
}

// WithHosts adds the given value to the Hosts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Hosts field.
func (b *FabricIstioApplyConfiguration) WithHosts(values ...string) *FabricIstioApplyConfiguration {
	for i := range values {
		b.Hosts = append(b.Hosts, values[i])
	}
	return b
}

// WithIngressGateway sets the IngressGateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IngressGateway field is set to the value of the last call.
func (b *FabricIstioApplyConfiguration) WithIngressGateway(value string) *FabricIstioApplyConfiguration {
	b.IngressGateway = &value
	return b
}
