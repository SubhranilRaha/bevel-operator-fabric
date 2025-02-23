/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricNetworkConfigLister helps list FabricNetworkConfigs.
// All objects returned here must be treated as read-only.
type FabricNetworkConfigLister interface {
	// List lists all FabricNetworkConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricNetworkConfig, err error)
	// FabricNetworkConfigs returns an object that can list and get FabricNetworkConfigs.
	FabricNetworkConfigs(namespace string) FabricNetworkConfigNamespaceLister
	FabricNetworkConfigListerExpansion
}

// fabricNetworkConfigLister implements the FabricNetworkConfigLister interface.
type fabricNetworkConfigLister struct {
	indexer cache.Indexer
}

// NewFabricNetworkConfigLister returns a new FabricNetworkConfigLister.
func NewFabricNetworkConfigLister(indexer cache.Indexer) FabricNetworkConfigLister {
	return &fabricNetworkConfigLister{indexer: indexer}
}

// List lists all FabricNetworkConfigs in the indexer.
func (s *fabricNetworkConfigLister) List(selector labels.Selector) (ret []*v1alpha1.FabricNetworkConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricNetworkConfig))
	})
	return ret, err
}

// FabricNetworkConfigs returns an object that can list and get FabricNetworkConfigs.
func (s *fabricNetworkConfigLister) FabricNetworkConfigs(namespace string) FabricNetworkConfigNamespaceLister {
	return fabricNetworkConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricNetworkConfigNamespaceLister helps list and get FabricNetworkConfigs.
// All objects returned here must be treated as read-only.
type FabricNetworkConfigNamespaceLister interface {
	// List lists all FabricNetworkConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricNetworkConfig, err error)
	// Get retrieves the FabricNetworkConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricNetworkConfig, error)
	FabricNetworkConfigNamespaceListerExpansion
}

// fabricNetworkConfigNamespaceLister implements the FabricNetworkConfigNamespaceLister
// interface.
type fabricNetworkConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricNetworkConfigs in the indexer for a given namespace.
func (s fabricNetworkConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricNetworkConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricNetworkConfig))
	})
	return ret, err
}

// Get retrieves the FabricNetworkConfig from the indexer for a given namespace and name.
func (s fabricNetworkConfigNamespaceLister) Get(name string) (*v1alpha1.FabricNetworkConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricnetworkconfig"), name)
	}
	return obj.(*v1alpha1.FabricNetworkConfig), nil
}
