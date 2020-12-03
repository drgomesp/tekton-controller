/*
Copyright 2020 Daniel Ribeiro.

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

	"github.com/davecgh/go-spew/spew"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tektondevv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

var (
	jobOwnerKey = ".metadata.controller"
	apiGVStr    = tektondevv1beta1.SchemeGroupVersion.String()
)

// TaskReconciler reconciles a Task object
type TaskReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tekton.dev.my.domain,resources=tasks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tekton.dev.my.domain,resources=tasks/status,verbs=get;update;patch

func (r *TaskReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("task", req.NamespacedName)

	var task tektondevv1beta1.Task
	if err := r.Get(
		ctx,
		req.NamespacedName,
		&task,
	); err != nil {
		r.Log.Error(err, "unable to get task")
		return ctrl.Result{}, err
	}

	spew.Dump(task)

	return ctrl.Result{}, nil
}

func (r *TaskReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tektondevv1beta1.Task{}).
		Complete(r)
}
