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

package controllers

import (
	"reflect"
	"testing"

	gatewaytemplatev1alpha1 "github.com/takumakume/gateway-template-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatewayv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func Test_httpRouteTemplateToHTTPRoute(t *testing.T) {
	type args struct {
		httpRouteTemplate *gatewaytemplatev1alpha1.HTTPRouteTemplate
	}
	tests := []struct {
		name    string
		args    args
		want    *gatewayv1b1.HTTPRoute
		wantErr bool
	}{
		{
			name: "generate HTTPRoute",
			args: args{
				httpRouteTemplate: &gatewaytemplatev1alpha1.HTTPRouteTemplate{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "name1",
						Namespace: "ns1",
					},
					Spec: gatewaytemplatev1alpha1.HTTPRouteTemplateSpec{
						HTTPRouteAnnotations: map[string]string{
							"annotation-key1": "annotation-value1",
						},
						HTTPRouteLabels: map[string]string{
							"label-key1": "label-value1",
						},
						HTTPRouteSpecTemplate: gatewayv1b1.HTTPRouteSpec{
							Hostnames: []gatewayv1b1.Hostname{
								"hoge1.{{ .Metadata.Namespace }}.example.com",
								"hoge2.{{ .Metadata.Namespace }}.example.com",
							},
						},
					},
				},
			},
			want: &gatewayv1b1.HTTPRoute{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "name1",
					Namespace: "ns1",
					Annotations: map[string]string{
						"annotation-key1": "annotation-value1",
					},
					Labels: map[string]string{
						"label-key1": "label-value1",
					},
				},
				Spec: gatewayv1b1.HTTPRouteSpec{
					Hostnames: []gatewayv1b1.Hostname{
						"hoge1.ns1.example.com",
						"hoge2.ns1.example.com",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := httpRouteTemplateToHTTPRoute(tt.args.httpRouteTemplate)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpRouteTemplateToHTTPRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpRouteTemplateToHTTPRoute() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
