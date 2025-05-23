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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	cainjector "github.com/cert-manager/cert-manager/internal/apis/config/cainjector"
	sharedv1alpha1 "github.com/cert-manager/cert-manager/internal/apis/config/shared/v1alpha1"
	cainjectorv1alpha1 "github.com/cert-manager/cert-manager/pkg/apis/config/cainjector/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*cainjectorv1alpha1.CAInjectorConfiguration)(nil), (*cainjector.CAInjectorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CAInjectorConfiguration_To_cainjector_CAInjectorConfiguration(a.(*cainjectorv1alpha1.CAInjectorConfiguration), b.(*cainjector.CAInjectorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cainjector.CAInjectorConfiguration)(nil), (*cainjectorv1alpha1.CAInjectorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cainjector_CAInjectorConfiguration_To_v1alpha1_CAInjectorConfiguration(a.(*cainjector.CAInjectorConfiguration), b.(*cainjectorv1alpha1.CAInjectorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cainjectorv1alpha1.EnableDataSourceConfig)(nil), (*cainjector.EnableDataSourceConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EnableDataSourceConfig_To_cainjector_EnableDataSourceConfig(a.(*cainjectorv1alpha1.EnableDataSourceConfig), b.(*cainjector.EnableDataSourceConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cainjector.EnableDataSourceConfig)(nil), (*cainjectorv1alpha1.EnableDataSourceConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cainjector_EnableDataSourceConfig_To_v1alpha1_EnableDataSourceConfig(a.(*cainjector.EnableDataSourceConfig), b.(*cainjectorv1alpha1.EnableDataSourceConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cainjectorv1alpha1.EnableInjectableConfig)(nil), (*cainjector.EnableInjectableConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EnableInjectableConfig_To_cainjector_EnableInjectableConfig(a.(*cainjectorv1alpha1.EnableInjectableConfig), b.(*cainjector.EnableInjectableConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cainjector.EnableInjectableConfig)(nil), (*cainjectorv1alpha1.EnableInjectableConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cainjector_EnableInjectableConfig_To_v1alpha1_EnableInjectableConfig(a.(*cainjector.EnableInjectableConfig), b.(*cainjectorv1alpha1.EnableInjectableConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CAInjectorConfiguration_To_cainjector_CAInjectorConfiguration(in *cainjectorv1alpha1.CAInjectorConfiguration, out *cainjector.CAInjectorConfiguration, s conversion.Scope) error {
	out.KubeConfig = in.KubeConfig
	out.Namespace = in.Namespace
	if err := sharedv1alpha1.Convert_v1alpha1_LeaderElectionConfig_To_shared_LeaderElectionConfig(&in.LeaderElectionConfig, &out.LeaderElectionConfig, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_EnableDataSourceConfig_To_cainjector_EnableDataSourceConfig(&in.EnableDataSourceConfig, &out.EnableDataSourceConfig, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_EnableInjectableConfig_To_cainjector_EnableInjectableConfig(&in.EnableInjectableConfig, &out.EnableInjectableConfig, s); err != nil {
		return err
	}
	out.EnablePprof = in.EnablePprof
	out.PprofAddress = in.PprofAddress
	out.Logging = in.Logging
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.MetricsListenAddress = in.MetricsListenAddress
	if err := sharedv1alpha1.Convert_v1alpha1_TLSConfig_To_shared_TLSConfig(&in.MetricsTLSConfig, &out.MetricsTLSConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_CAInjectorConfiguration_To_cainjector_CAInjectorConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_CAInjectorConfiguration_To_cainjector_CAInjectorConfiguration(in *cainjectorv1alpha1.CAInjectorConfiguration, out *cainjector.CAInjectorConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_CAInjectorConfiguration_To_cainjector_CAInjectorConfiguration(in, out, s)
}

func autoConvert_cainjector_CAInjectorConfiguration_To_v1alpha1_CAInjectorConfiguration(in *cainjector.CAInjectorConfiguration, out *cainjectorv1alpha1.CAInjectorConfiguration, s conversion.Scope) error {
	out.KubeConfig = in.KubeConfig
	out.Namespace = in.Namespace
	if err := sharedv1alpha1.Convert_shared_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(&in.LeaderElectionConfig, &out.LeaderElectionConfig, s); err != nil {
		return err
	}
	if err := Convert_cainjector_EnableDataSourceConfig_To_v1alpha1_EnableDataSourceConfig(&in.EnableDataSourceConfig, &out.EnableDataSourceConfig, s); err != nil {
		return err
	}
	if err := Convert_cainjector_EnableInjectableConfig_To_v1alpha1_EnableInjectableConfig(&in.EnableInjectableConfig, &out.EnableInjectableConfig, s); err != nil {
		return err
	}
	out.EnablePprof = in.EnablePprof
	out.PprofAddress = in.PprofAddress
	out.Logging = in.Logging
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.MetricsListenAddress = in.MetricsListenAddress
	if err := sharedv1alpha1.Convert_shared_TLSConfig_To_v1alpha1_TLSConfig(&in.MetricsTLSConfig, &out.MetricsTLSConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_cainjector_CAInjectorConfiguration_To_v1alpha1_CAInjectorConfiguration is an autogenerated conversion function.
func Convert_cainjector_CAInjectorConfiguration_To_v1alpha1_CAInjectorConfiguration(in *cainjector.CAInjectorConfiguration, out *cainjectorv1alpha1.CAInjectorConfiguration, s conversion.Scope) error {
	return autoConvert_cainjector_CAInjectorConfiguration_To_v1alpha1_CAInjectorConfiguration(in, out, s)
}

func autoConvert_v1alpha1_EnableDataSourceConfig_To_cainjector_EnableDataSourceConfig(in *cainjectorv1alpha1.EnableDataSourceConfig, out *cainjector.EnableDataSourceConfig, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Certificates, &out.Certificates, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_EnableDataSourceConfig_To_cainjector_EnableDataSourceConfig is an autogenerated conversion function.
func Convert_v1alpha1_EnableDataSourceConfig_To_cainjector_EnableDataSourceConfig(in *cainjectorv1alpha1.EnableDataSourceConfig, out *cainjector.EnableDataSourceConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_EnableDataSourceConfig_To_cainjector_EnableDataSourceConfig(in, out, s)
}

func autoConvert_cainjector_EnableDataSourceConfig_To_v1alpha1_EnableDataSourceConfig(in *cainjector.EnableDataSourceConfig, out *cainjectorv1alpha1.EnableDataSourceConfig, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Certificates, &out.Certificates, s); err != nil {
		return err
	}
	return nil
}

// Convert_cainjector_EnableDataSourceConfig_To_v1alpha1_EnableDataSourceConfig is an autogenerated conversion function.
func Convert_cainjector_EnableDataSourceConfig_To_v1alpha1_EnableDataSourceConfig(in *cainjector.EnableDataSourceConfig, out *cainjectorv1alpha1.EnableDataSourceConfig, s conversion.Scope) error {
	return autoConvert_cainjector_EnableDataSourceConfig_To_v1alpha1_EnableDataSourceConfig(in, out, s)
}

func autoConvert_v1alpha1_EnableInjectableConfig_To_cainjector_EnableInjectableConfig(in *cainjectorv1alpha1.EnableInjectableConfig, out *cainjector.EnableInjectableConfig, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.ValidatingWebhookConfigurations, &out.ValidatingWebhookConfigurations, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.MutatingWebhookConfigurations, &out.MutatingWebhookConfigurations, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.CustomResourceDefinitions, &out.CustomResourceDefinitions, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.APIServices, &out.APIServices, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_EnableInjectableConfig_To_cainjector_EnableInjectableConfig is an autogenerated conversion function.
func Convert_v1alpha1_EnableInjectableConfig_To_cainjector_EnableInjectableConfig(in *cainjectorv1alpha1.EnableInjectableConfig, out *cainjector.EnableInjectableConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_EnableInjectableConfig_To_cainjector_EnableInjectableConfig(in, out, s)
}

func autoConvert_cainjector_EnableInjectableConfig_To_v1alpha1_EnableInjectableConfig(in *cainjector.EnableInjectableConfig, out *cainjectorv1alpha1.EnableInjectableConfig, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.ValidatingWebhookConfigurations, &out.ValidatingWebhookConfigurations, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.MutatingWebhookConfigurations, &out.MutatingWebhookConfigurations, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.CustomResourceDefinitions, &out.CustomResourceDefinitions, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.APIServices, &out.APIServices, s); err != nil {
		return err
	}
	return nil
}

// Convert_cainjector_EnableInjectableConfig_To_v1alpha1_EnableInjectableConfig is an autogenerated conversion function.
func Convert_cainjector_EnableInjectableConfig_To_v1alpha1_EnableInjectableConfig(in *cainjector.EnableInjectableConfig, out *cainjectorv1alpha1.EnableInjectableConfig, s conversion.Scope) error {
	return autoConvert_cainjector_EnableInjectableConfig_To_v1alpha1_EnableInjectableConfig(in, out, s)
}
