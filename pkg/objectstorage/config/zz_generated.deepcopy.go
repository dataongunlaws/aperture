//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package config

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageBackoffConfig) DeepCopyInto(out *ObjectStorageBackoffConfig) {
	*out = *in
	in.Initial.DeepCopyInto(&out.Initial)
	in.Maximum.DeepCopyInto(&out.Maximum)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageBackoffConfig.
func (in *ObjectStorageBackoffConfig) DeepCopy() *ObjectStorageBackoffConfig {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageBackoffConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageConfig) DeepCopyInto(out *ObjectStorageConfig) {
	*out = *in
	in.RetryPolicy.DeepCopyInto(&out.RetryPolicy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageConfig.
func (in *ObjectStorageConfig) DeepCopy() *ObjectStorageConfig {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageRetryPolicy) DeepCopyInto(out *ObjectStorageRetryPolicy) {
	*out = *in
	in.Timeout.DeepCopyInto(&out.Timeout)
	in.Backoff.DeepCopyInto(&out.Backoff)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageRetryPolicy.
func (in *ObjectStorageRetryPolicy) DeepCopy() *ObjectStorageRetryPolicy {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageRetryPolicy)
	in.DeepCopyInto(out)
	return out
}