/*
Copyright 2021 VMware VDP Team.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/vmware/kube-fluentd-operator/config-reloader/datasource/kubedatasource/fluentdconfig/apis/logs.vdp.vmware.com/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FluentdConfigLister helps list FluentdConfigs.
// All objects returned here must be treated as read-only.
type FluentdConfigLister interface {
	// List lists all FluentdConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FluentdConfig, err error)
	// FluentdConfigs returns an object that can list and get FluentdConfigs.
	FluentdConfigs(namespace string) FluentdConfigNamespaceLister
	FluentdConfigListerExpansion
}

// fluentdConfigLister implements the FluentdConfigLister interface.
type fluentdConfigLister struct {
	indexer cache.Indexer
}

// NewFluentdConfigLister returns a new FluentdConfigLister.
func NewFluentdConfigLister(indexer cache.Indexer) FluentdConfigLister {
	return &fluentdConfigLister{indexer: indexer}
}

// List lists all FluentdConfigs in the indexer.
func (s *fluentdConfigLister) List(selector labels.Selector) (ret []*v1beta1.FluentdConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FluentdConfig))
	})
	return ret, err
}

// FluentdConfigs returns an object that can list and get FluentdConfigs.
func (s *fluentdConfigLister) FluentdConfigs(namespace string) FluentdConfigNamespaceLister {
	return fluentdConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FluentdConfigNamespaceLister helps list and get FluentdConfigs.
// All objects returned here must be treated as read-only.
type FluentdConfigNamespaceLister interface {
	// List lists all FluentdConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FluentdConfig, err error)
	// Get retrieves the FluentdConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.FluentdConfig, error)
	FluentdConfigNamespaceListerExpansion
}

// fluentdConfigNamespaceLister implements the FluentdConfigNamespaceLister
// interface.
type fluentdConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FluentdConfigs in the indexer for a given namespace.
func (s fluentdConfigNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FluentdConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FluentdConfig))
	})
	return ret, err
}

// Get retrieves the FluentdConfig from the indexer for a given namespace and name.
func (s fluentdConfigNamespaceLister) Get(name string) (*v1beta1.FluentdConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("fluentdconfig"), name)
	}
	return obj.(*v1beta1.FluentdConfig), nil
}
