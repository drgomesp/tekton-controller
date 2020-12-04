package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tektondevv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

// TaskRunReconciler reconciles a TaskRun object
type TaskRunReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tekton.dev.my.domain,resources=taskruns,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tekton.dev.my.domain,resources=taskruns/status,verbs=get;update;patch

func (r *TaskRunReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("taskrun", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *TaskRunReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tektondevv1beta1.TaskRun{}).
		Complete(r)
}
