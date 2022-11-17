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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatewayv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HTTPRouteTemplateSpec defines the desired state of HTTPRouteTemplate
type HTTPRouteTemplateSpec struct {
	// HTTPRouteSpec Template for HTTPRoute.Spec
	// +kubebuilder:validation:Required
	HTTPRouteSpecTemplate gatewayv1b1.HTTPRouteSpec `json:"httpRouteSpecTemplate"`

	// Annotations This annotation is generated in HTTPRoute
	// +optional
	HTTPRouteAnnotations map[string]string `json:"httpRouteAnnotations,omitempty"`

	// Labels This labels is generated in HTTPRoute
	// +optional
	HTTPRouteLabels map[string]string `json:"httpRouteLabels,omitempty"`
}

// HTTPRouteTemplateStatus defines the observed state of HTTPRouteTemplate
type HTTPRouteTemplateStatus struct {
	// Ready HTTPRoute generation status
	Ready corev1.ConditionStatus `json:"ready,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HTTPRouteTemplate is the Schema for the httproutetemplates API
type HTTPRouteTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPRouteTemplateSpec   `json:"spec,omitempty"`
	Status HTTPRouteTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HTTPRouteTemplateList contains a list of HTTPRouteTemplate
type HTTPRouteTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPRouteTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HTTPRouteTemplate{}, &HTTPRouteTemplateList{})
}
