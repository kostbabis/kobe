// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KobeFederator) DeepCopyInto(out *KobeFederator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KobeFederator.
func (in *KobeFederator) DeepCopy() *KobeFederator {
	if in == nil {
		return nil
	}
	out := new(KobeFederator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KobeFederator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KobeFederatorList) DeepCopyInto(out *KobeFederatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KobeFederator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KobeFederatorList.
func (in *KobeFederatorList) DeepCopy() *KobeFederatorList {
	if in == nil {
		return nil
	}
	out := new(KobeFederatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KobeFederatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KobeFederatorSpec) DeepCopyInto(out *KobeFederatorSpec) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KobeFederatorSpec.
func (in *KobeFederatorSpec) DeepCopy() *KobeFederatorSpec {
	if in == nil {
		return nil
	}
	out := new(KobeFederatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KobeFederatorStatus) DeepCopyInto(out *KobeFederatorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KobeFederatorStatus.
func (in *KobeFederatorStatus) DeepCopy() *KobeFederatorStatus {
	if in == nil {
		return nil
	}
	out := new(KobeFederatorStatus)
	in.DeepCopyInto(out)
	return out
}
