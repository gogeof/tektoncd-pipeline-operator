// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.Config":          schema_pkg_apis_operator_v1alpha1_Config(ref),
		"github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigCondition": schema_pkg_apis_operator_v1alpha1_ConfigCondition(ref),
		"github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigSpec":      schema_pkg_apis_operator_v1alpha1_ConfigSpec(ref),
		"github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigStatus":    schema_pkg_apis_operator_v1alpha1_ConfigStatus(ref),
	}
}

func schema_pkg_apis_operator_v1alpha1_Config(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Config is the Schema for the configs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigSpec", "github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1alpha1_ConfigCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigCondition defines the observed state of installation at a point in time",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"code": {
						SchemaProps: spec.SchemaProps{
							Description: "Code indicates the status of installation of pipeline resources Valid values are:\n  - \"error\"\n  - \"installing\"\n  - \"installed\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"details": {
						SchemaProps: spec.SchemaProps{
							Description: "Additional details about the Code",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "The version of OpenShift pipelines operator",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pipelineVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "The version of OpenShift pipelines",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"triggersVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "The version of OpenShift triggers",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"code", "version"},
			},
		},
	}
}

func schema_pkg_apis_operator_v1alpha1_ConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigSpec defines the desired state of Config",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"targetNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "namespace where OpenShift pipelines will be installed",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"targetNamespace"},
			},
		},
	}
}

func schema_pkg_apis_operator_v1alpha1_ConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigStatus defines the observed state of Config",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"operatorUUID": {
						SchemaProps: spec.SchemaProps{
							Description: "OperatorUUID is the  uuid (auto-generated) of the operator that installed the pipeline",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "installation status sorted in reverse chronological order",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigCondition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/openshift-pipelines-operator/pkg/apis/operator/v1alpha1.ConfigCondition"},
	}
}
