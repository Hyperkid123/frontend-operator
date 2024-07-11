/*
Copyright 2021 RedHatInsights.

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
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NavTemplateSpec defines the desired state of NavTemplate
type NavTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ID          string   `json:"id" yaml:"id"`
	Title       string   `json:"title" yaml:"title"`
	EnvName     string   `json:"envName,omitempty" yaml:"envName,omitempty"`
	MountPoints []string `json:"mountPoints" yaml:"mountPoints"`
}

type GeneratedNavTemplate struct {
	ID       string           `json:"id" yaml:"id"`
	Title    string           `json:"title" yaml:"title"`
	NavItems []*ChromeNavItem `json:"navItems" yaml:"navItems"`
}

// NavTemplateStatus defines the observed state of NavTemplate
type NavTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NavTemplate is the Schema for the navtemplates API
type NavTemplate struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   NavTemplateSpec   `json:"spec,omitempty" yaml:"id,omitempty"`
	Status NavTemplateStatus `json:"status,omitempty" yaml:"id,omitempty"`
}

//+kubebuilder:object:root=true

// NavTemplateList contains a list of NavTemplate
type NavTemplateList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []NavTemplate `json:"items" yaml:"items"`
}

func (i *NavTemplate) GetOurEnv(ctx context.Context, pClient client.Client, env *FrontendEnvironment) error {
	return pClient.Get(ctx, types.NamespacedName{Name: i.Spec.EnvName}, env)
}

func init() {
	SchemeBuilder.Register(&NavTemplate{}, &NavTemplateList{})
}
