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

package v1

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"strings"
)

// log is for logging in this package.
var guestbooklog = logf.Log.WithName("guestbook-resource")

func (r *GuestBook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:webhookVersions={v1beta1},path=/mutate-webapp-example-org-v1-guestbook,mutating=true,failurePolicy=fail,sideEffects=None,groups=webapp.example.org,resources=guestbooks,verbs=create;update,versions=v1,name=mguestbook.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &GuestBook{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *GuestBook) Default() {
	guestbooklog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	if r.Spec.Foo == "" {
		r.Spec.Foo = "default"
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:webhookVersions={v1beta1},path=/validate-webapp-example-org-v1-guestbook,mutating=false,failurePolicy=fail,sideEffects=None,groups=webapp.example.org,resources=guestbooks,verbs=create;update,versions=v1,name=vguestbook.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &GuestBook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *GuestBook) ValidateCreate() error {
	guestbooklog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.

	return r.ValidateSpec()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *GuestBook) ValidateUpdate(old runtime.Object) error {
	guestbooklog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return r.ValidateSpec()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *GuestBook) ValidateDelete() error {
	guestbooklog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
func (r *GuestBook) ValidateSpec() error {
	if strings.Contains(r.Spec.Foo, "black") {
		return errors.New("spec.foo can not contain black")
	}
	return nil
}
