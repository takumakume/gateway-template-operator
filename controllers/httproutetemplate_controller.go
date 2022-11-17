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
	"context"
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	gatewaytemplatev1alpha1 "github.com/takumakume/gateway-template-operator/api/v1alpha1"
	gatewayv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// HTTPRouteTemplateReconciler reconciles a HTTPRouteTemplate object
type HTTPRouteTemplateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=gateway-template.takumakume.github.io,resources=httproutetemplates,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=gateway-template.takumakume.github.io,resources=httproutetemplates/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=gateway-template.takumakume.github.io,resources=httproutetemplates/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HTTPRouteTemplate object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *HTTPRouteTemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("HTTPRouteTemplate", req.NamespacedName.String())

	httpRouteTemplate := &gatewaytemplatev1alpha1.HTTPRouteTemplate{}
	if err := r.Get(ctx, req.NamespacedName, httpRouteTemplate); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		log.Error(err, "unable to fetch IngressTemplate")
		return ctrl.Result{}, err
	}

	log.Info("starting reconcile loop")
	defer log.Info("finish reconcile loop")

	if !httpRouteTemplate.GetDeletionTimestamp().IsZero() {
		return ctrl.Result{}, nil
	}

	log.Info("run create or update Ingress")

	httpRoute, err := httpRouteTemplateToHTTPRoute(httpRouteTemplate)
	if err != nil {
		return ctrl.Result{}, err
	}
	ownerRef := metav1.NewControllerRef(
		&httpRoute.ObjectMeta,
		schema.GroupVersionKind{
			Group:   gatewaytemplatev1alpha1.GroupVersion.Group,
			Version: gatewaytemplatev1alpha1.GroupVersion.Version,
			Kind:    "HTTPRouteTemplate",
		})
	ownerRef.Name = httpRouteTemplate.Name
	ownerRef.UID = httpRouteTemplate.GetUID()
	httpRoute.ObjectMeta.SetOwnerReferences([]metav1.OwnerReference{*ownerRef})

	createdHTTPRoute := &gatewayv1b1.HTTPRoute{}
	if err := r.Get(ctx, req.NamespacedName, createdHTTPRoute); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("run create Ingress")
			if createErr := r.Create(ctx, httpRoute); createErr != nil {
				return ctrl.Result{}, createErr
			}
			log.Info("create HTTPRoute successful")

			httpRouteTemplate.Status.Ready = corev1.ConditionTrue
			if statusUpdateErr := r.Status().Update(ctx, httpRouteTemplate); statusUpdateErr != nil {
				return ctrl.Result{}, statusUpdateErr
			}
		}

		log.Error(err, "unable to fetch HTTPRoute")
		return ctrl.Result{}, err
	} else {
		needUpdate := false
		if !reflect.DeepEqual(createdHTTPRoute.ObjectMeta.Labels, httpRouteTemplate.ObjectMeta.Labels) {
			log.Info(fmt.Sprintf("detects changes ObjectMeta.Label: %+v, %+v", createdHTTPRoute.ObjectMeta.Labels, ingress.ObjectMeta.Labels))
			needUpdate = true
		}
		if !reflect.DeepEqual(createdHTTPRoute.ObjectMeta.Annotations, httpRouteTemplate.ObjectMeta.Annotations) {
			log.Info(fmt.Sprintf("detects changes ObjectMeta.Annotations: %+v, %+v", createdHTTPRoute.ObjectMeta.Annotations, ingress.ObjectMeta.Annotations))
			needUpdate = true
		}
		if !reflect.DeepEqual(createdHTTPRoute.Spec, httpRouteTemplate.Spec) {
			log.Info(fmt.Sprintf("detects changes Spec: %+v, %+v", createdHTTPRoute.Spec, httpRouteTemplate.Spec))
			needUpdate = true
		}

		if needUpdate {
			log.Info("run update HTTPRoute")
			if err := r.Update(ctx, httpRouteTemplate); err != nil {
				return ctrl.Result{}, err
			}
			log.Info("update HTTPRoute successful")
		}
	}

	if err != nil {
		log.Error(err, "unable to create or update HTTPRoute")
		if statusUpdateErr := r.Update(ctx, httpRouteTemplate); statusUpdateErr != nil {
			return ctrl.Result{}, statusUpdateErr
		}
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *HTTPRouteTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&gatewaytemplatev1alpha1.HTTPRouteTemplate{}).
		Complete(r)
}

func httpRouteTemplateToHTTPRoute(httpRouteTemplate *gatewaytemplatev1alpha1.HTTPRouteTemplate) (*gatewayv1b1.HTTPRoute, error) {
	generated := &gatewayv1b1.HTTPRoute{}
	// TODO
	return generated, nil
}
