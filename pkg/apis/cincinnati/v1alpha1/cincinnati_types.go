package v1alpha1

import (
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CincinnatiSpec defines the desired state of Cincinnati
type CincinnatiSpec struct {
	// +kubebuilder:validation:Minimum=1
	// Replicas is the number of pods to run. When >=2, a PodDisruptionBudget
	// will ensure that voluntary disruption leaves at least one Pod running at
	// all times.
	Replicas int32 `json:"replicas"`

	// Registry is the container registry to use, such as "quay.io".
	Registry string `json:"registry"`

	// Repository is the repository to use in the Registry, such as
	// "openshift-release-dev/ocp-release"
	Repository string `json:"repository"`

	// GitHubOrg is the organization to use on GitHub for retrieving additional
	// graph data.
	GitHubOrg string `json:"gitHubOrg,omitempty"`

	// GitHubRepo is the repository to use on GitHub for retrieving additional
	// graph data.
	GitHubRepo string `json:"gitHubRepo,omitempty"`

	// Branch is the git branch to use on GitHub for retrieving additional graph
	// data.
	Branch string `json:"branch,omitempty"`

	// GraphDataImage is an init container image that contains the
	// Cincinnati graph data and copies it to /var/cincinnati/graph-data. If
	// specified, it takes precedence over GitHubOrg and GitHubRepo and disables
	// retrieval of graph data from GitHub.
	GraphDataImage string `json:"graphDataImage,omitempty"`

	// CofigmapKey refers to the ConfigMap key that holds a CA cert to the registry
	// Cincinnati will contact to build the upgrade graph.  The name of the
	// ConfigMap holding the key comes from the field additionalTrustedCA
	// in the image.config.openshift.io API.
	ConfigmapKey string `json:"configmapKey"`
}

// CincinnatiStatus defines the observed state of Cincinnati
type CincinnatiStatus struct {
	// Conditions describe the state of the Cincinnati resource.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +optional
	Conditions []conditionsv1.Condition `json:"conditions,omitempty"  patchStrategy:"merge" patchMergeKey:"type"`
}

// Condition Types
const (
	// ConditionReconcileCompleted reports whether all required resources have been created
	// in the cluster and reflect the specified state.
	ConditionReconcileCompleted conditionsv1.ConditionType = "ReconcileCompleted"

	// ConditionConfigurationConflict reports conflicts in the graph data configuration.
	// There are two ways to configure the graph data source. If both ways are configured,
	// this condition will indicate that.
	ConditionConfigurationConflict conditionsv1.ConditionType = "ConfigurationConflict"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cincinnati is the Schema for a Cincinnati service.
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cincinnatis,scope=Namespaced
type Cincinnati struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CincinnatiSpec   `json:"spec,omitempty"`
	Status CincinnatiStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CincinnatiList contains a list of Cincinnati
type CincinnatiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cincinnati `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cincinnati{}, &CincinnatiList{})
}
