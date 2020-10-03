// +build !ignore_autogenerated

/*


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
	"k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicRole) DeepCopyInto(out *DynamicRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicRole.
func (in *DynamicRole) DeepCopy() *DynamicRole {
	if in == nil {
		return nil
	}
	out := new(DynamicRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicRoleList) DeepCopyInto(out *DynamicRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamicRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicRoleList.
func (in *DynamicRoleList) DeepCopy() *DynamicRoleList {
	if in == nil {
		return nil
	}
	out := new(DynamicRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicRoleSpec) DeepCopyInto(out *DynamicRoleSpec) {
	*out = *in
	if in.Inherit != nil {
		in, out := &in.Inherit, &out.Inherit
		*out = new(InheritedRole)
		**out = **in
	}
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = new([]v1.PolicyRule)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.PolicyRule, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = new([]v1.PolicyRule)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.PolicyRule, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicRoleSpec.
func (in *DynamicRoleSpec) DeepCopy() *DynamicRoleSpec {
	if in == nil {
		return nil
	}
	out := new(DynamicRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicRoleStatus) DeepCopyInto(out *DynamicRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicRoleStatus.
func (in *DynamicRoleStatus) DeepCopy() *DynamicRoleStatus {
	if in == nil {
		return nil
	}
	out := new(DynamicRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InheritedRole) DeepCopyInto(out *InheritedRole) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InheritedRole.
func (in *InheritedRole) DeepCopy() *InheritedRole {
	if in == nil {
		return nil
	}
	out := new(InheritedRole)
	in.DeepCopyInto(out)
	return out
}