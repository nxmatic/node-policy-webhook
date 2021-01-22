// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPAuthDatasource) DeepCopyInto(out *GCPAuthDatasource) {
	*out = *in
	out.GCPAuthSecretDatasource = in.GCPAuthSecretDatasource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPAuthDatasource.
func (in *GCPAuthDatasource) DeepCopy() *GCPAuthDatasource {
	if in == nil {
		return nil
	}
	out := new(GCPAuthDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPAuthPolicyProfile) DeepCopyInto(out *GCPAuthPolicyProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPAuthPolicyProfile.
func (in *GCPAuthPolicyProfile) DeepCopy() *GCPAuthPolicyProfile {
	if in == nil {
		return nil
	}
	out := new(GCPAuthPolicyProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPAuthPolicyProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPAuthPolicyProfileList) DeepCopyInto(out *GCPAuthPolicyProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCPAuthPolicyProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPAuthPolicyProfileList.
func (in *GCPAuthPolicyProfileList) DeepCopy() *GCPAuthPolicyProfileList {
	if in == nil {
		return nil
	}
	out := new(GCPAuthPolicyProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPAuthPolicyProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPAuthPolicyProfileSpec) DeepCopyInto(out *GCPAuthPolicyProfileSpec) {
	*out = *in
	out.GCPAuthDatasource = in.GCPAuthDatasource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPAuthPolicyProfileSpec.
func (in *GCPAuthPolicyProfileSpec) DeepCopy() *GCPAuthPolicyProfileSpec {
	if in == nil {
		return nil
	}
	out := new(GCPAuthPolicyProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPAuthPolicyProfileStatus) DeepCopyInto(out *GCPAuthPolicyProfileStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPAuthPolicyProfileStatus.
func (in *GCPAuthPolicyProfileStatus) DeepCopy() *GCPAuthPolicyProfileStatus {
	if in == nil {
		return nil
	}
	out := new(GCPAuthPolicyProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPAuthSecretDatasource) DeepCopyInto(out *GCPAuthSecretDatasource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPAuthSecretDatasource.
func (in *GCPAuthSecretDatasource) DeepCopy() *GCPAuthSecretDatasource {
	if in == nil {
		return nil
	}
	out := new(GCPAuthSecretDatasource)
	in.DeepCopyInto(out)
	return out
}
