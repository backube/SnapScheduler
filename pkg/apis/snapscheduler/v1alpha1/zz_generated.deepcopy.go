// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotRetentionSpec) DeepCopyInto(out *SnapshotRetentionSpec) {
	*out = *in
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotRetentionSpec.
func (in *SnapshotRetentionSpec) DeepCopy() *SnapshotRetentionSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotRetentionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSchedule) DeepCopyInto(out *SnapshotSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSchedule.
func (in *SnapshotSchedule) DeepCopy() *SnapshotSchedule {
	if in == nil {
		return nil
	}
	out := new(SnapshotSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotScheduleList) DeepCopyInto(out *SnapshotScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnapshotSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotScheduleList.
func (in *SnapshotScheduleList) DeepCopy() *SnapshotScheduleList {
	if in == nil {
		return nil
	}
	out := new(SnapshotScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotScheduleSpec) DeepCopyInto(out *SnapshotScheduleSpec) {
	*out = *in
	in.ClaimSelector.DeepCopyInto(&out.ClaimSelector)
	in.Retention.DeepCopyInto(&out.Retention)
	if in.SnapshotTemplate != nil {
		in, out := &in.SnapshotTemplate, &out.SnapshotTemplate
		*out = new(SnapshotTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotScheduleSpec.
func (in *SnapshotScheduleSpec) DeepCopy() *SnapshotScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotScheduleStatus) DeepCopyInto(out *SnapshotScheduleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastSnapshotTime != nil {
		in, out := &in.LastSnapshotTime, &out.LastSnapshotTime
		*out = (*in).DeepCopy()
	}
	if in.NextSnapshotTime != nil {
		in, out := &in.NextSnapshotTime, &out.NextSnapshotTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotScheduleStatus.
func (in *SnapshotScheduleStatus) DeepCopy() *SnapshotScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotTemplateSpec) DeepCopyInto(out *SnapshotTemplateSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SnapshotClassName != nil {
		in, out := &in.SnapshotClassName, &out.SnapshotClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotTemplateSpec.
func (in *SnapshotTemplateSpec) DeepCopy() *SnapshotTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
