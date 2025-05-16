/*
Copyright 2024.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ReplicasSpec struct {
	Min int32 `json:"min"`
	Max int32 `json:"max"`
}

type LocationConfig struct {
	Tag      string       `json:"tag"`
	Replicas ReplicasSpec `json:"replicas"`
}

type PersistentVolumeConfig struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`

	//Capacity    string                              `json:"capacity"`
	//AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
}

type RegistryAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistryConfig struct {
	RegistryType string       `json:"registryType"`
	URL          string       `json:"url"`
	Auth         RegistryAuth `json:"auth"`
}

// PortConfig defines the structure of a port, including its privacy status
type PortConfig struct {
	corev1.ServicePort `json:",inline"`
	Private            bool `json:"private,omitempty"`
}

// AppSpec defines the desired state of App
type AppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Locations []LocationConfig            `json:"locations"`
	Image     string                      `json:"image"`
	Args      []string                    `json:"args,omitempty"`
	Env       []corev1.EnvVar             `json:"env,omitempty"`
	Ports     []PortConfig                `json:"ports"`
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	Volumes   []PersistentVolumeConfig    `json:"volumes,omitempty"`
	Private   bool                        `json:"private,omitempty"`
	Registry  RegistryConfig              `json:"registry,omitempty"`
	//easyjson:json
}

// AppStatus defines the observed state of App
type AppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// App is the Schema for the apps API
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`
	Status AppStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppList contains a list of App
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
