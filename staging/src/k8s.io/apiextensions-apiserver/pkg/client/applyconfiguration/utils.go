/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/applyconfiguration/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/client/applyconfiguration/apiextensions/v1beta1"
	internal "k8s.io/apiextensions-apiserver/pkg/client/applyconfiguration/internal"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=apiextensions.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("CustomResourceColumnDefinition"):
		return &apiextensionsv1.CustomResourceColumnDefinitionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceConversion"):
		return &apiextensionsv1.CustomResourceConversionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinition"):
		return &apiextensionsv1.CustomResourceDefinitionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionCondition"):
		return &apiextensionsv1.CustomResourceDefinitionConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionNames"):
		return &apiextensionsv1.CustomResourceDefinitionNamesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionSpec"):
		return &apiextensionsv1.CustomResourceDefinitionSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionStatus"):
		return &apiextensionsv1.CustomResourceDefinitionStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionVersion"):
		return &apiextensionsv1.CustomResourceDefinitionVersionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceSubresources"):
		return &apiextensionsv1.CustomResourceSubresourcesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceSubresourceScale"):
		return &apiextensionsv1.CustomResourceSubresourceScaleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceValidation"):
		return &apiextensionsv1.CustomResourceValidationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternalDocumentation"):
		return &apiextensionsv1.ExternalDocumentationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JSONSchemaProps"):
		return &apiextensionsv1.JSONSchemaPropsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SelectableField"):
		return &apiextensionsv1.SelectableFieldApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceReference"):
		return &apiextensionsv1.ServiceReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ValidationRule"):
		return &apiextensionsv1.ValidationRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebhookClientConfig"):
		return &apiextensionsv1.WebhookClientConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebhookConversion"):
		return &apiextensionsv1.WebhookConversionApplyConfiguration{}

		// Group=apiextensions.k8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceColumnDefinition"):
		return &apiextensionsv1beta1.CustomResourceColumnDefinitionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceConversion"):
		return &apiextensionsv1beta1.CustomResourceConversionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceDefinition"):
		return &apiextensionsv1beta1.CustomResourceDefinitionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceDefinitionCondition"):
		return &apiextensionsv1beta1.CustomResourceDefinitionConditionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceDefinitionNames"):
		return &apiextensionsv1beta1.CustomResourceDefinitionNamesApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceDefinitionSpec"):
		return &apiextensionsv1beta1.CustomResourceDefinitionSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceDefinitionStatus"):
		return &apiextensionsv1beta1.CustomResourceDefinitionStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceDefinitionVersion"):
		return &apiextensionsv1beta1.CustomResourceDefinitionVersionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceSubresources"):
		return &apiextensionsv1beta1.CustomResourceSubresourcesApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceSubresourceScale"):
		return &apiextensionsv1beta1.CustomResourceSubresourceScaleApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("CustomResourceValidation"):
		return &apiextensionsv1beta1.CustomResourceValidationApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ExternalDocumentation"):
		return &apiextensionsv1beta1.ExternalDocumentationApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("JSONSchemaProps"):
		return &apiextensionsv1beta1.JSONSchemaPropsApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SelectableField"):
		return &apiextensionsv1beta1.SelectableFieldApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ServiceReference"):
		return &apiextensionsv1beta1.ServiceReferenceApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ValidationRule"):
		return &apiextensionsv1beta1.ValidationRuleApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WebhookClientConfig"):
		return &apiextensionsv1beta1.WebhookClientConfigApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) managedfields.TypeConverter {
	return managedfields.NewSchemeTypeConverter(scheme, internal.Parser())
}
