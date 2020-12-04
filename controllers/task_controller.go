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
