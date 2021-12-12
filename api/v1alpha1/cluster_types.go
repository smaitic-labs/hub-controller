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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ClusterGeoLocation struct {
	Latitude      string `json:"latitude,omitempty"`
	Longitude     string `json:"longitude,omitempty"`
	CloudProvider string `json:"cloudProvider,omitempty"`
	CloudRegion   string `json:"cloudRegion,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	Country       string `json:"country,omitempty"`
}

type ClusterMonitoringKubernetesDashboard struct {
	Enabled       bool   `json:"enabled,omitempty"`
	IngressPrefix string `json:"ingressPrefix,omitempty"`
	AccessToken   string `json:"accessToken,omitempty"`
}

type ClusterMonitoringSpec struct {
	KubernetesDashboard ClusterMonitoringKubernetesDashboard `json:"kubernetesDashboard,omitempty"`
}

type TelemetryProvider string

type ClusterTelemetrySpec struct {
	Enabled           bool              `json:"enabled,omitempty"`
	TelemetryProvider TelemetryProvider `json:"telemetryProvider,omitempty"`
	Endpoint          string            `json:"endpoint,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Cluster. Edit cluster_types.go to remove/update
	HubEndpoint string                `json:"hubEndpoint,omitempty"`
	AccessToken string                `json:"accessToken,omitempty"`
	Telemetry   ClusterTelemetrySpec  `json:"telemetry,omitempty"`
	Monitoring  ClusterMonitoringSpec `json:"monitoring,omitempty"`
	GeoLocation ClusterGeoLocation    `json:"geoLocation,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
