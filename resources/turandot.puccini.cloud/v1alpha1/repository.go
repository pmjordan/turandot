package v1alpha1

import (
	"fmt"

	"github.com/tliron/kutil/ard"
	group "github.com/tliron/turandot/resources/turandot.puccini.cloud"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RepositoryGVK = SchemeGroupVersion.WithKind(RepositoryKind)

type RepositoryType string

const (
	RepositoryKind     = "Repository"
	RepositoryListKind = "RepositoryList"

	RepositorySingular  = "repository"
	RepositoryPlural    = "repositories"
	RepositoryShortName = "repo"

	RepositoryTypeRegistry RepositoryType = "registry"
)

//
// Repository
//

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Repository struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepositorySpec   `json:"spec"`
	Status RepositoryStatus `json:"status"`
}

type RepositorySpec struct {
	Type     RepositoryType      `json:"type"`
	Direct   *RepositoryDirect   `json:"direct,omitempty"`
	Indirect *RepositoryIndirect `json:"indirect,omitempty"`
	Secret   string              `json:"secret,omitempty"`
}

type RepositoryDirect struct {
	Address string `json:"address"`
}

type RepositoryIndirect struct {
	Namespace string `json:"namespace,omitempty"`
	Service   string `json:"service"`
	Port      uint64 `json:"port"`
}

type RepositoryStatus struct {
	SpoolerPod string `json:"spoolerPod"`
}

//
// RepositoryList
//

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type RepositoryList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata"`

	Items []Repository `json:"items"`
}

//
// RepositoryCustomResourceDefinition
//

// See: assets/custom-resource-definitions.yaml

var RepositoryResourcesName = fmt.Sprintf("%s.%s", RepositoryPlural, group.GroupName)

var RepositoryCustomResourceDefinition = apiextensions.CustomResourceDefinition{
	ObjectMeta: meta.ObjectMeta{
		Name: RepositoryResourcesName,
	},
	Spec: apiextensions.CustomResourceDefinitionSpec{
		Group: group.GroupName,
		Names: apiextensions.CustomResourceDefinitionNames{
			Singular: RepositorySingular,
			Plural:   RepositoryPlural,
			Kind:     RepositoryKind,
			ListKind: RepositoryListKind,
			ShortNames: []string{
				RepositoryShortName,
			},
			Categories: []string{
				"all", // will appear in "kubectl get all"
			},
		},
		Scope: apiextensions.NamespaceScoped,
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    Version,
				Served:  true,
				Storage: true, // one and only one version must be marked with storage=true
				Subresources: &apiextensions.CustomResourceSubresources{ // requires CustomResourceSubresources feature gate enabled
					Status: &apiextensions.CustomResourceSubresourceStatus{},
				},
				Schema: &apiextensions.CustomResourceValidation{
					OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
						Type:     "object",
						Required: []string{"spec"},
						Properties: map[string]apiextensions.JSONSchemaProps{
							"spec": {
								Type:     "object",
								Required: []string{"type"},
								Properties: map[string]apiextensions.JSONSchemaProps{
									"type": {
										Type: "string",
										Enum: []apiextensions.JSON{
											{Raw: []byte(fmt.Sprintf("%q", RepositoryTypeRegistry))},
										},
									},
									"direct": {
										Type:     "object",
										Required: []string{"address"},
										Properties: map[string]apiextensions.JSONSchemaProps{
											"address": {
												Type: "string",
											},
										},
									},
									"indirect": {
										Type:     "object",
										Required: []string{"service", "port"},
										Properties: map[string]apiextensions.JSONSchemaProps{
											"namespace": {
												Type: "string",
											},
											"service": {
												Type: "string",
											},
											"port": {
												Type: "integer",
											},
										},
									},
									"secret": {
										Type: "string",
									},
								},
								OneOf: []apiextensions.JSONSchemaProps{
									{
										Required: []string{"direct"},
									},
									{
										Required: []string{"indirect"},
									},
								},
							},
							"status": {
								Type: "object",
								Properties: map[string]apiextensions.JSONSchemaProps{
									"spoolerPod": {
										Type: "string",
									},
								},
							},
						},
					},
				},
				AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
					{
						Name:     "Type",
						Type:     "string",
						JSONPath: ".spec.type",
					},
					{
						Name:     "SpoolerPod",
						Type:     "string",
						JSONPath: ".status.spoolerPod",
					},
				},
			},
		},
	},
}

func RepositoryToARD(repository *Repository) ard.StringMap {
	map_ := make(ard.StringMap)
	map_["Name"] = repository.Name
	if repository.Spec.Direct != nil {
		map_["Direct"] = ard.StringMap{
			"Address": repository.Spec.Direct.Address,
		}
	}
	if repository.Spec.Indirect != nil {
		map_["Indirect"] = ard.StringMap{
			"Namespace": repository.Spec.Indirect.Namespace,
			"Service":   repository.Spec.Indirect.Service,
			"Port":      repository.Spec.Indirect.Port,
		}
	}
	map_["Secret"] = repository.Spec.Secret
	map_["SpoolerPod"] = repository.Status.SpoolerPod
	return map_
}
