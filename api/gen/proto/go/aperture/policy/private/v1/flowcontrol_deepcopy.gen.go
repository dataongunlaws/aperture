// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package privatev1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using LoadActuator within kubernetes types, where deepcopy-gen is used.
func (in *LoadActuator) DeepCopyInto(out *LoadActuator) {
	p := proto.Clone(in).(*LoadActuator)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator. Required by controller-gen.
func (in *LoadActuator) DeepCopy() *LoadActuator {
	if in == nil {
		return nil
	}
	out := new(LoadActuator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator. Required by controller-gen.
func (in *LoadActuator) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadActuator_Ins within kubernetes types, where deepcopy-gen is used.
func (in *LoadActuator_Ins) DeepCopyInto(out *LoadActuator_Ins) {
	p := proto.Clone(in).(*LoadActuator_Ins)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator_Ins. Required by controller-gen.
func (in *LoadActuator_Ins) DeepCopy() *LoadActuator_Ins {
	if in == nil {
		return nil
	}
	out := new(LoadActuator_Ins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator_Ins. Required by controller-gen.
func (in *LoadActuator_Ins) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}