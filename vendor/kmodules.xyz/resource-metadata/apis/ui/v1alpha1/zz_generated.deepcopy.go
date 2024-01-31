//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apisshared "kmodules.xyz/resource-metadata/apis/shared"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	releasesv1alpha1 "x-helm.dev/apimachinery/apis/releases/v1alpha1"
	shared "x-helm.dev/apimachinery/apis/shared"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionTemplate) DeepCopyInto(out *ActionTemplate) {
	*out = *in
	out.ActionInfo = in.ActionInfo
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]shared.ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Editor != nil {
		in, out := &in.Editor, &out.Editor
		*out = new(releasesv1alpha1.ChartSourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionTemplate.
func (in *ActionTemplate) DeepCopy() *ActionTemplate {
	if in == nil {
		return nil
	}
	out := new(ActionTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionTemplateGroup) DeepCopyInto(out *ActionTemplateGroup) {
	*out = *in
	out.ActionInfo = in.ActionInfo
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActionTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionTemplateGroup.
func (in *ActionTemplateGroup) DeepCopy() *ActionTemplateGroup {
	if in == nil {
		return nil
	}
	out := new(ActionTemplateGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfo) DeepCopyInto(out *ChartInfo) {
	*out = *in
	out.SourceRef = in.SourceRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfo.
func (in *ChartInfo) DeepCopy() *ChartInfo {
	if in == nil {
		return nil
	}
	out := new(ChartInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	if in.Managed != nil {
		in, out := &in.Managed, &out.Managed
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]apisshared.DashboardVar, len(*in))
		copy(*out, *in)
	}
	if in.Panels != nil {
		in, out := &in.Panels, &out.Panels
		*out = make([]PanelLinkRequest, len(*in))
		copy(*out, *in)
	}
	if in.If != nil {
		in, out := &in.If, &out.If
		*out = new(apisshared.If)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependentFeatureSet) DeepCopyInto(out *DependentFeatureSet) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependentFeatureSet.
func (in *DependentFeatureSet) DeepCopy() *DependentFeatureSet {
	if in == nil {
		return nil
	}
	out := new(DependentFeatureSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependents) DeepCopyInto(out *Dependents) {
	*out = *in
	if in.FeatureSets != nil {
		in, out := &in.FeatureSets, &out.FeatureSets
		*out = make([]DependentFeatureSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependents.
func (in *Dependents) DeepCopy() *Dependents {
	if in == nil {
		return nil
	}
	out := new(Dependents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Feature) DeepCopyInto(out *Feature) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Feature.
func (in *Feature) DeepCopy() *Feature {
	if in == nil {
		return nil
	}
	out := new(Feature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Feature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureList) DeepCopyInto(out *FeatureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Feature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureList.
func (in *FeatureList) DeepCopy() *FeatureList {
	if in == nil {
		return nil
	}
	out := new(FeatureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSet) DeepCopyInto(out *FeatureSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSet.
func (in *FeatureSet) DeepCopy() *FeatureSet {
	if in == nil {
		return nil
	}
	out := new(FeatureSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSetList) DeepCopyInto(out *FeatureSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSetList.
func (in *FeatureSetList) DeepCopy() *FeatureSetList {
	if in == nil {
		return nil
	}
	out := new(FeatureSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSetSpec) DeepCopyInto(out *FeatureSetSpec) {
	*out = *in
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]shared.ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.RequiredFeatures != nil {
		in, out := &in.RequiredFeatures, &out.RequiredFeatures
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Chart = in.Chart
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSetSpec.
func (in *FeatureSetSpec) DeepCopy() *FeatureSetSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSetStatus) DeepCopyInto(out *FeatureSetStatus) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]ComponentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Dependents.DeepCopyInto(&out.Dependents)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSetStatus.
func (in *FeatureSetStatus) DeepCopy() *FeatureSetStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSpec) DeepCopyInto(out *FeatureSpec) {
	*out = *in
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]shared.ImageSpec, len(*in))
		copy(*out, *in)
	}
	in.Requirements.DeepCopyInto(&out.Requirements)
	in.ReadinessChecks.DeepCopyInto(&out.ReadinessChecks)
	out.Chart = in.Chart
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSpec.
func (in *FeatureSpec) DeepCopy() *FeatureSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStatus) DeepCopyInto(out *FeatureStatus) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Managed != nil {
		in, out := &in.Managed, &out.Managed
		*out = new(bool)
		**out = **in
	}
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStatus.
func (in *FeatureStatus) DeepCopy() *FeatureStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PanelLinkRequest) DeepCopyInto(out *PanelLinkRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PanelLinkRequest.
func (in *PanelLinkRequest) DeepCopy() *PanelLinkRequest {
	if in == nil {
		return nil
	}
	out := new(PanelLinkRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessChecks) DeepCopyInto(out *ReadinessChecks) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]metav1.GroupVersionKind, len(*in))
		copy(*out, *in)
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]WorkloadInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessChecks.
func (in *ReadinessChecks) DeepCopy() *ReadinessChecks {
	if in == nil {
		return nil
	}
	out := new(ReadinessChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requirements) DeepCopyInto(out *Requirements) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requirements.
func (in *Requirements) DeepCopy() *Requirements {
	if in == nil {
		return nil
	}
	out := new(Requirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDashboard) DeepCopyInto(out *ResourceDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDashboard.
func (in *ResourceDashboard) DeepCopy() *ResourceDashboard {
	if in == nil {
		return nil
	}
	out := new(ResourceDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDashboardList) DeepCopyInto(out *ResourceDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDashboardList.
func (in *ResourceDashboardList) DeepCopy() *ResourceDashboardList {
	if in == nil {
		return nil
	}
	out := new(ResourceDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDashboardSpec) DeepCopyInto(out *ResourceDashboardSpec) {
	*out = *in
	out.Resource = in.Resource
	if in.Dashboards != nil {
		in, out := &in.Dashboards, &out.Dashboards
		*out = make([]Dashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDashboardSpec.
func (in *ResourceDashboardSpec) DeepCopy() *ResourceDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEditor) DeepCopyInto(out *ResourceEditor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEditor.
func (in *ResourceEditor) DeepCopy() *ResourceEditor {
	if in == nil {
		return nil
	}
	out := new(ResourceEditor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceEditor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEditorList) DeepCopyInto(out *ResourceEditorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceEditor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEditorList.
func (in *ResourceEditorList) DeepCopy() *ResourceEditorList {
	if in == nil {
		return nil
	}
	out := new(ResourceEditorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceEditorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEditorSpec) DeepCopyInto(out *ResourceEditorSpec) {
	*out = *in
	out.Resource = in.Resource
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(UIParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]shared.ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Variants != nil {
		in, out := &in.Variants, &out.Variants
		*out = make([]VariantRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Installer != nil {
		in, out := &in.Installer, &out.Installer
		*out = new(apisshared.DeploymentParameters)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEditorSpec.
func (in *ResourceEditorSpec) DeepCopy() *ResourceEditorSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceEditorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UIParameters) DeepCopyInto(out *UIParameters) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(releasesv1alpha1.ChartSourceRef)
		**out = **in
	}
	if in.Editor != nil {
		in, out := &in.Editor, &out.Editor
		*out = new(releasesv1alpha1.ChartSourceRef)
		**out = **in
	}
	if in.InstanceLabelPaths != nil {
		in, out := &in.InstanceLabelPaths, &out.InstanceLabelPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*ActionTemplateGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ActionTemplateGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UIParameters.
func (in *UIParameters) DeepCopy() *UIParameters {
	if in == nil {
		return nil
	}
	out := new(UIParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VariantRef) DeepCopyInto(out *VariantRef) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]shared.ImageSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VariantRef.
func (in *VariantRef) DeepCopy() *VariantRef {
	if in == nil {
		return nil
	}
	out := new(VariantRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadInfo) DeepCopyInto(out *WorkloadInfo) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadInfo.
func (in *WorkloadInfo) DeepCopy() *WorkloadInfo {
	if in == nil {
		return nil
	}
	out := new(WorkloadInfo)
	in.DeepCopyInto(out)
	return out
}