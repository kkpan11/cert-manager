//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The cert-manager Authors.

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

package controller

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEDNS01Config) DeepCopyInto(out *ACMEDNS01Config) {
	*out = *in
	if in.RecursiveNameservers != nil {
		in, out := &in.RecursiveNameservers, &out.RecursiveNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEDNS01Config.
func (in *ACMEDNS01Config) DeepCopy() *ACMEDNS01Config {
	if in == nil {
		return nil
	}
	out := new(ACMEDNS01Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEHTTP01Config) DeepCopyInto(out *ACMEHTTP01Config) {
	*out = *in
	if in.SolverNameservers != nil {
		in, out := &in.SolverNameservers, &out.SolverNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEHTTP01Config.
func (in *ACMEHTTP01Config) DeepCopy() *ACMEHTTP01Config {
	if in == nil {
		return nil
	}
	out := new(ACMEHTTP01Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfiguration) DeepCopyInto(out *ControllerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.LeaderElectionConfig = in.LeaderElectionConfig
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CopiedAnnotationPrefixes != nil {
		in, out := &in.CopiedAnnotationPrefixes, &out.CopiedAnnotationPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.MetricsTLSConfig.DeepCopyInto(&out.MetricsTLSConfig)
	in.Logging.DeepCopyInto(&out.Logging)
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.IngressShimConfig.DeepCopyInto(&out.IngressShimConfig)
	in.ACMEHTTP01Config.DeepCopyInto(&out.ACMEHTTP01Config)
	in.ACMEDNS01Config.DeepCopyInto(&out.ACMEDNS01Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfiguration.
func (in *ControllerConfiguration) DeepCopy() *ControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressShimConfig) DeepCopyInto(out *IngressShimConfig) {
	*out = *in
	if in.DefaultAutoCertificateAnnotations != nil {
		in, out := &in.DefaultAutoCertificateAnnotations, &out.DefaultAutoCertificateAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraCertificateAnnotations != nil {
		in, out := &in.ExtraCertificateAnnotations, &out.ExtraCertificateAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressShimConfig.
func (in *IngressShimConfig) DeepCopy() *IngressShimConfig {
	if in == nil {
		return nil
	}
	out := new(IngressShimConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfig) DeepCopyInto(out *LeaderElectionConfig) {
	*out = *in
	out.LeaderElectionConfig = in.LeaderElectionConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfig.
func (in *LeaderElectionConfig) DeepCopy() *LeaderElectionConfig {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfig)
	in.DeepCopyInto(out)
	return out
}
