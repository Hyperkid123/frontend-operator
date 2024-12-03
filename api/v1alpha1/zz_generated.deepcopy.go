//go:build !ignore_autogenerated

/*
Copyright 2021 RedHatInsights.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIInfo) DeepCopyInto(out *APIInfo) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIInfo.
func (in *APIInfo) DeepCopy() *APIInfo {
	if in == nil {
		return nil
	}
	out := new(APIInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Analytics) DeepCopyInto(out *Analytics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Analytics.
func (in *Analytics) DeepCopy() *Analytics {
	if in == nil {
		return nil
	}
	out := new(Analytics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bundle) DeepCopyInto(out *Bundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bundle.
func (in *Bundle) DeepCopy() *Bundle {
	if in == nil {
		return nil
	}
	out := new(Bundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleList) DeepCopyInto(out *BundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleList.
func (in *BundleList) DeepCopy() *BundleList {
	if in == nil {
		return nil
	}
	out := new(BundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleNavItem) DeepCopyInto(out *BundleNavItem) {
	*out = *in
	if in.NavItems != nil {
		in, out := &in.NavItems, &out.NavItems
		*out = make([]LeafBundleNavItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]BundlePermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]EmbeddedRoute, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleNavItem.
func (in *BundleNavItem) DeepCopy() *BundleNavItem {
	if in == nil {
		return nil
	}
	out := new(BundleNavItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundlePermission) DeepCopyInto(out *BundlePermission) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]BundlePermissionArg, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundlePermission.
func (in *BundlePermission) DeepCopy() *BundlePermission {
	if in == nil {
		return nil
	}
	out := new(BundlePermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleSpec) DeepCopyInto(out *BundleSpec) {
	*out = *in
	if in.AppList != nil {
		in, out := &in.AppList, &out.AppList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraNavItems != nil {
		in, out := &in.ExtraNavItems, &out.ExtraNavItems
		*out = make([]ExtraNavItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomNav != nil {
		in, out := &in.CustomNav, &out.CustomNav
		*out = make([]ChromeNavItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleSpec.
func (in *BundleSpec) DeepCopy() *BundleSpec {
	if in == nil {
		return nil
	}
	out := new(BundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleStatus) DeepCopyInto(out *BundleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleStatus.
func (in *BundleStatus) DeepCopy() *BundleStatus {
	if in == nil {
		return nil
	}
	out := new(BundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChromeNavItem) DeepCopyInto(out *ChromeNavItem) {
	*out = *in
	if in.NavItems != nil {
		in, out := &in.NavItems, &out.NavItems
		*out = make([]ChromeNavItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]ChromeNavItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Position != nil {
		in, out := &in.Position, &out.Position
		*out = new(uint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChromeNavItem.
func (in *ChromeNavItem) DeepCopy() *ChromeNavItem {
	if in == nil {
		return nil
	}
	out := new(ChromeNavItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputedBundle) DeepCopyInto(out *ComputedBundle) {
	*out = *in
	if in.NavItems != nil {
		in, out := &in.NavItems, &out.NavItems
		*out = make([]ChromeNavItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputedBundle.
func (in *ComputedBundle) DeepCopy() *ComputedBundle {
	if in == nil {
		return nil
	}
	out := new(ComputedBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedRoute) DeepCopyInto(out *EmbeddedRoute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedRoute.
func (in *EmbeddedRoute) DeepCopy() *EmbeddedRoute {
	if in == nil {
		return nil
	}
	out := new(EmbeddedRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraNavItem) DeepCopyInto(out *ExtraNavItem) {
	*out = *in
	in.NavItem.DeepCopyInto(&out.NavItem)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraNavItem.
func (in *ExtraNavItem) DeepCopy() *ExtraNavItem {
	if in == nil {
		return nil
	}
	out := new(ExtraNavItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FedModule) DeepCopyInto(out *FedModule) {
	*out = *in
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.ModuleConfig != nil {
		in, out := &in.ModuleConfig, &out.ModuleConfig
		*out = new(ModuleConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.FullProfile != nil {
		in, out := &in.FullProfile, &out.FullProfile
		*out = new(bool)
		**out = **in
	}
	if in.IsFedramp != nil {
		in, out := &in.IsFedramp, &out.IsFedramp
		*out = new(bool)
		**out = **in
	}
	if in.Analytics != nil {
		in, out := &in.Analytics, &out.Analytics
		*out = new(Analytics)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FedModule.
func (in *FedModule) DeepCopy() *FedModule {
	if in == nil {
		return nil
	}
	out := new(FedModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Frontend) DeepCopyInto(out *Frontend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Frontend.
func (in *Frontend) DeepCopy() *Frontend {
	if in == nil {
		return nil
	}
	out := new(Frontend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Frontend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendBundles) DeepCopyInto(out *FrontendBundles) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendBundles.
func (in *FrontendBundles) DeepCopy() *FrontendBundles {
	if in == nil {
		return nil
	}
	out := new(FrontendBundles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendBundlesGenerated) DeepCopyInto(out *FrontendBundlesGenerated) {
	*out = *in
	if in.NavItems != nil {
		in, out := &in.NavItems, &out.NavItems
		*out = new([]ChromeNavItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ChromeNavItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendBundlesGenerated.
func (in *FrontendBundlesGenerated) DeepCopy() *FrontendBundlesGenerated {
	if in == nil {
		return nil
	}
	out := new(FrontendBundlesGenerated)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendDeployments) DeepCopyInto(out *FrontendDeployments) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendDeployments.
func (in *FrontendDeployments) DeepCopy() *FrontendDeployments {
	if in == nil {
		return nil
	}
	out := new(FrontendDeployments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendEnvironment) DeepCopyInto(out *FrontendEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendEnvironment.
func (in *FrontendEnvironment) DeepCopy() *FrontendEnvironment {
	if in == nil {
		return nil
	}
	out := new(FrontendEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrontendEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendEnvironmentList) DeepCopyInto(out *FrontendEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FrontendEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendEnvironmentList.
func (in *FrontendEnvironmentList) DeepCopy() *FrontendEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(FrontendEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrontendEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendEnvironmentSpec) DeepCopyInto(out *FrontendEnvironmentSpec) {
	*out = *in
	if in.IngressAnnotations != nil {
		in, out := &in.IngressAnnotations, &out.IngressAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Whitelist != nil {
		in, out := &in.Whitelist, &out.Whitelist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(MonitoringConfig)
		**out = **in
	}
	if in.AkamaiCacheBustURLs != nil {
		in, out := &in.AkamaiCacheBustURLs, &out.AkamaiCacheBustURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TargetNamespaces != nil {
		in, out := &in.TargetNamespaces, &out.TargetNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceCategories != nil {
		in, out := &in.ServiceCategories, &out.ServiceCategories
		*out = new([]FrontendServiceCategory)
		if **in != nil {
			in, out := *in, *out
			*out = make([]FrontendServiceCategory, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultReplicas != nil {
		in, out := &in.DefaultReplicas, &out.DefaultReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Bundles != nil {
		in, out := &in.Bundles, &out.Bundles
		*out = new([]FrontendBundles)
		if **in != nil {
			in, out := *in, *out
			*out = make([]FrontendBundles, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendEnvironmentSpec.
func (in *FrontendEnvironmentSpec) DeepCopy() *FrontendEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(FrontendEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendEnvironmentStatus) DeepCopyInto(out *FrontendEnvironmentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendEnvironmentStatus.
func (in *FrontendEnvironmentStatus) DeepCopy() *FrontendEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(FrontendEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendInfo) DeepCopyInto(out *FrontendInfo) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendInfo.
func (in *FrontendInfo) DeepCopy() *FrontendInfo {
	if in == nil {
		return nil
	}
	out := new(FrontendInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendList) DeepCopyInto(out *FrontendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Frontend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendList.
func (in *FrontendList) DeepCopy() *FrontendList {
	if in == nil {
		return nil
	}
	out := new(FrontendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrontendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendServiceCategory) DeepCopyInto(out *FrontendServiceCategory) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]FrontendServiceCategoryGroup, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendServiceCategory.
func (in *FrontendServiceCategory) DeepCopy() *FrontendServiceCategory {
	if in == nil {
		return nil
	}
	out := new(FrontendServiceCategory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendServiceCategoryGenerated) DeepCopyInto(out *FrontendServiceCategoryGenerated) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]FrontendServiceCategoryGroupGenerated, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendServiceCategoryGenerated.
func (in *FrontendServiceCategoryGenerated) DeepCopy() *FrontendServiceCategoryGenerated {
	if in == nil {
		return nil
	}
	out := new(FrontendServiceCategoryGenerated)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendServiceCategoryGroup) DeepCopyInto(out *FrontendServiceCategoryGroup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendServiceCategoryGroup.
func (in *FrontendServiceCategoryGroup) DeepCopy() *FrontendServiceCategoryGroup {
	if in == nil {
		return nil
	}
	out := new(FrontendServiceCategoryGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendServiceCategoryGroupGenerated) DeepCopyInto(out *FrontendServiceCategoryGroupGenerated) {
	*out = *in
	if in.Tiles != nil {
		in, out := &in.Tiles, &out.Tiles
		*out = new([]ServiceTile)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ServiceTile, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendServiceCategoryGroupGenerated.
func (in *FrontendServiceCategoryGroupGenerated) DeepCopy() *FrontendServiceCategoryGroupGenerated {
	if in == nil {
		return nil
	}
	out := new(FrontendServiceCategoryGroupGenerated)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendSpec) DeepCopyInto(out *FrontendSpec) {
	*out = *in
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(APIInfo)
		(*in).DeepCopyInto(*out)
	}
	in.Frontend.DeepCopyInto(&out.Frontend)
	out.ServiceMonitor = in.ServiceMonitor
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = new(FedModule)
		(*in).DeepCopyInto(*out)
	}
	if in.NavItems != nil {
		in, out := &in.NavItems, &out.NavItems
		*out = make([]*BundleNavItem, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BundleNavItem)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NavigationSegments != nil {
		in, out := &in.NavigationSegments, &out.NavigationSegments
		*out = make([]*NavigationSegment, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NavigationSegment)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AkamaiCacheBustPaths != nil {
		in, out := &in.AkamaiCacheBustPaths, &out.AkamaiCacheBustPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SearchEntries != nil {
		in, out := &in.SearchEntries, &out.SearchEntries
		*out = make([]*SearchEntry, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SearchEntry)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ServiceTiles != nil {
		in, out := &in.ServiceTiles, &out.ServiceTiles
		*out = make([]*ServiceTile, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ServiceTile)
				**out = **in
			}
		}
	}
	if in.WidgetRegistry != nil {
		in, out := &in.WidgetRegistry, &out.WidgetRegistry
		*out = make([]*WidgetEntry, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(WidgetEntry)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendSpec.
func (in *FrontendSpec) DeepCopy() *FrontendSpec {
	if in == nil {
		return nil
	}
	out := new(FrontendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendStatus) DeepCopyInto(out *FrontendStatus) {
	*out = *in
	out.Deployments = in.Deployments
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendStatus.
func (in *FrontendStatus) DeepCopy() *FrontendStatus {
	if in == nil {
		return nil
	}
	out := new(FrontendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeafBundleNavItem) DeepCopyInto(out *LeafBundleNavItem) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]EmbeddedRoute, len(*in))
		copy(*out, *in)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]BundlePermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeafBundleNavItem.
func (in *LeafBundleNavItem) DeepCopy() *LeafBundleNavItem {
	if in == nil {
		return nil
	}
	out := new(LeafBundleNavItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OptionalDependencies != nil {
		in, out := &in.OptionalDependencies, &out.OptionalDependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleConfig) DeepCopyInto(out *ModuleConfig) {
	*out = *in
	out.SupportCaseData = in.SupportCaseData
	if in.SSOScopes != nil {
		in, out := &in.SSOScopes, &out.SSOScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleConfig.
func (in *ModuleConfig) DeepCopy() *ModuleConfig {
	if in == nil {
		return nil
	}
	out := new(ModuleConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfig) DeepCopyInto(out *MonitoringConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfig.
func (in *MonitoringConfig) DeepCopy() *MonitoringConfig {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NavigationSegment) DeepCopyInto(out *NavigationSegment) {
	*out = *in
	if in.NavItems != nil {
		in, out := &in.NavItems, &out.NavItems
		*out = new([]ChromeNavItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ChromeNavItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NavigationSegment.
func (in *NavigationSegment) DeepCopy() *NavigationSegment {
	if in == nil {
		return nil
	}
	out := new(NavigationSegment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permission) DeepCopyInto(out *Permission) {
	*out = *in
	if in.Apps != nil {
		in, out := &in.Apps, &out.Apps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permission.
func (in *Permission) DeepCopy() *Permission {
	if in == nil {
		return nil
	}
	out := new(Permission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.Props != nil {
		in, out := &in.Props, &out.Props
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.SupportCaseData != nil {
		in, out := &in.SupportCaseData, &out.SupportCaseData
		*out = new(SupportCaseData)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchEntry) DeepCopyInto(out *SearchEntry) {
	*out = *in
	if in.AltTitle != nil {
		in, out := &in.AltTitle, &out.AltTitle
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchEntry.
func (in *SearchEntry) DeepCopy() *SearchEntry {
	if in == nil {
		return nil
	}
	out := new(SearchEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorConfig) DeepCopyInto(out *ServiceMonitorConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorConfig.
func (in *ServiceMonitorConfig) DeepCopy() *ServiceMonitorConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTile) DeepCopyInto(out *ServiceTile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTile.
func (in *ServiceTile) DeepCopy() *ServiceTile {
	if in == nil {
		return nil
	}
	out := new(ServiceTile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportCaseData) DeepCopyInto(out *SupportCaseData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportCaseData.
func (in *SupportCaseData) DeepCopy() *SupportCaseData {
	if in == nil {
		return nil
	}
	out := new(SupportCaseData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetConfig) DeepCopyInto(out *WidgetConfig) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.HeaderLink = in.HeaderLink
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetConfig.
func (in *WidgetConfig) DeepCopy() *WidgetConfig {
	if in == nil {
		return nil
	}
	out := new(WidgetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetDefaultVariant) DeepCopyInto(out *WidgetDefaultVariant) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetDefaultVariant.
func (in *WidgetDefaultVariant) DeepCopy() *WidgetDefaultVariant {
	if in == nil {
		return nil
	}
	out := new(WidgetDefaultVariant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetDefaults) DeepCopyInto(out *WidgetDefaults) {
	*out = *in
	out.Small = in.Small
	out.Medium = in.Medium
	out.Large = in.Large
	out.XLarge = in.XLarge
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetDefaults.
func (in *WidgetDefaults) DeepCopy() *WidgetDefaults {
	if in == nil {
		return nil
	}
	out := new(WidgetDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetEntry) DeepCopyInto(out *WidgetEntry) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	out.Defaults = in.Defaults
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetEntry.
func (in *WidgetEntry) DeepCopy() *WidgetEntry {
	if in == nil {
		return nil
	}
	out := new(WidgetEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetHeaderLink) DeepCopyInto(out *WidgetHeaderLink) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetHeaderLink.
func (in *WidgetHeaderLink) DeepCopy() *WidgetHeaderLink {
	if in == nil {
		return nil
	}
	out := new(WidgetHeaderLink)
	in.DeepCopyInto(out)
	return out
}
