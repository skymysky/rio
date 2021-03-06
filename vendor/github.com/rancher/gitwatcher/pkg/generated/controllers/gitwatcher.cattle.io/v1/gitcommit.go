/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/rancher/gitwatcher/pkg/apis/gitwatcher.cattle.io/v1"
	clientset "github.com/rancher/gitwatcher/pkg/generated/clientset/versioned/typed/gitwatcher.cattle.io/v1"
	informers "github.com/rancher/gitwatcher/pkg/generated/informers/externalversions/gitwatcher.cattle.io/v1"
	listers "github.com/rancher/gitwatcher/pkg/generated/listers/gitwatcher.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type GitCommitHandler func(string, *v1.GitCommit) (*v1.GitCommit, error)

type GitCommitController interface {
	GitCommitClient

	OnChange(ctx context.Context, name string, sync GitCommitHandler)
	OnRemove(ctx context.Context, name string, sync GitCommitHandler)
	Enqueue(namespace, name string)

	Cache() GitCommitCache

	Informer() cache.SharedIndexInformer
	GroupVersionKind() schema.GroupVersionKind

	AddGenericHandler(ctx context.Context, name string, handler generic.Handler)
	AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler)
	Updater() generic.Updater
}

type GitCommitClient interface {
	Create(*v1.GitCommit) (*v1.GitCommit, error)
	Update(*v1.GitCommit) (*v1.GitCommit, error)
	UpdateStatus(*v1.GitCommit) (*v1.GitCommit, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.GitCommit, error)
	List(namespace string, opts metav1.ListOptions) (*v1.GitCommitList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.GitCommit, err error)
}

type GitCommitCache interface {
	Get(namespace, name string) (*v1.GitCommit, error)
	List(namespace string, selector labels.Selector) ([]*v1.GitCommit, error)

	AddIndexer(indexName string, indexer GitCommitIndexer)
	GetByIndex(indexName, key string) ([]*v1.GitCommit, error)
}

type GitCommitIndexer func(obj *v1.GitCommit) ([]string, error)

type gitCommitController struct {
	controllerManager *generic.ControllerManager
	clientGetter      clientset.GitCommitsGetter
	informer          informers.GitCommitInformer
	gvk               schema.GroupVersionKind
}

func NewGitCommitController(gvk schema.GroupVersionKind, controllerManager *generic.ControllerManager, clientGetter clientset.GitCommitsGetter, informer informers.GitCommitInformer) GitCommitController {
	return &gitCommitController{
		controllerManager: controllerManager,
		clientGetter:      clientGetter,
		informer:          informer,
		gvk:               gvk,
	}
}

func FromGitCommitHandlerToHandler(sync GitCommitHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.GitCommit
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.GitCommit))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *gitCommitController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.GitCommit))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateGitCommitOnChange(updater generic.Updater, handler GitCommitHandler) GitCommitHandler {
	return func(key string, obj *v1.GitCommit) (*v1.GitCommit, error) {
		if obj == nil {
			return handler(key, nil)
		}

		copyObj := obj.DeepCopy()
		newObj, err := handler(key, copyObj)
		if newObj != nil {
			copyObj = newObj
		}
		if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
			newObj, err := updater(copyObj)
			if newObj != nil && err == nil {
				copyObj = newObj.(*v1.GitCommit)
			}
		}

		return copyObj, err
	}
}

func (c *gitCommitController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, handler)
}

func (c *gitCommitController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), handler)
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, removeHandler)
}

func (c *gitCommitController) OnChange(ctx context.Context, name string, sync GitCommitHandler) {
	c.AddGenericHandler(ctx, name, FromGitCommitHandlerToHandler(sync))
}

func (c *gitCommitController) OnRemove(ctx context.Context, name string, sync GitCommitHandler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), FromGitCommitHandlerToHandler(sync))
	c.AddGenericHandler(ctx, name, removeHandler)
}

func (c *gitCommitController) Enqueue(namespace, name string) {
	c.controllerManager.Enqueue(c.gvk, c.informer.Informer(), namespace, name)
}

func (c *gitCommitController) Informer() cache.SharedIndexInformer {
	return c.informer.Informer()
}

func (c *gitCommitController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *gitCommitController) Cache() GitCommitCache {
	return &gitCommitCache{
		lister:  c.informer.Lister(),
		indexer: c.informer.Informer().GetIndexer(),
	}
}

func (c *gitCommitController) Create(obj *v1.GitCommit) (*v1.GitCommit, error) {
	return c.clientGetter.GitCommits(obj.Namespace).Create(obj)
}

func (c *gitCommitController) Update(obj *v1.GitCommit) (*v1.GitCommit, error) {
	return c.clientGetter.GitCommits(obj.Namespace).Update(obj)
}

func (c *gitCommitController) UpdateStatus(obj *v1.GitCommit) (*v1.GitCommit, error) {
	return c.clientGetter.GitCommits(obj.Namespace).UpdateStatus(obj)
}

func (c *gitCommitController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return c.clientGetter.GitCommits(namespace).Delete(name, options)
}

func (c *gitCommitController) Get(namespace, name string, options metav1.GetOptions) (*v1.GitCommit, error) {
	return c.clientGetter.GitCommits(namespace).Get(name, options)
}

func (c *gitCommitController) List(namespace string, opts metav1.ListOptions) (*v1.GitCommitList, error) {
	return c.clientGetter.GitCommits(namespace).List(opts)
}

func (c *gitCommitController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientGetter.GitCommits(namespace).Watch(opts)
}

func (c *gitCommitController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.GitCommit, err error) {
	return c.clientGetter.GitCommits(namespace).Patch(name, pt, data, subresources...)
}

type gitCommitCache struct {
	lister  listers.GitCommitLister
	indexer cache.Indexer
}

func (c *gitCommitCache) Get(namespace, name string) (*v1.GitCommit, error) {
	return c.lister.GitCommits(namespace).Get(name)
}

func (c *gitCommitCache) List(namespace string, selector labels.Selector) ([]*v1.GitCommit, error) {
	return c.lister.GitCommits(namespace).List(selector)
}

func (c *gitCommitCache) AddIndexer(indexName string, indexer GitCommitIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.GitCommit))
		},
	}))
}

func (c *gitCommitCache) GetByIndex(indexName, key string) (result []*v1.GitCommit, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		result = append(result, obj.(*v1.GitCommit))
	}
	return result, nil
}
