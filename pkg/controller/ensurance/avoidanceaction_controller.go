/*
Copyright 2021.

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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ensurancecraneiov1alpha1 "github.com/gocrane-io/api/ensurance/v1alpha1"
	"github.com/gocrane-io/crane/pkg/ensurance/cache"
)

// AvoidanceActionController reconciles a AvoidanceAction object
type AvoidanceActionController struct {
	client.Client
	Scheme         *runtime.Scheme
	Log            logr.Logger
	RestMapper     meta.RESTMapper
	Recorder       record.EventRecorder
	DetectionCache *cache.DetectionConditionCache
}

//+kubebuilder:rbac:groups=ensurance.crane.io.crane.io,resources=AvoidanceActiones,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ensurance.crane.io.crane.io,resources=AvoidanceActiones/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ensurance.crane.io.crane.io,resources=AvoidanceActiones/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AvoidanceAction object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *AvoidanceActionController) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AvoidanceActionController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ensurancecraneiov1alpha1.AvoidanceAction{}).
		Complete(r)
}
