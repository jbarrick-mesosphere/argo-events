// +build !ignore_autogenerated

/*
Copyright 2018 BlackRock, Inc.

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
// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package kafka

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/gateways/core/stream/kafka.Kafka": schema_gateways_core_stream_kafka_Kafka(ref),
	}
}

func schema_gateways_core_stream_kafka_Kafka(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kafka defines configuration required to connect to kafka cluster",
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL to kafka cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"partition": {
						SchemaProps: spec.SchemaProps{
							Description: "Partition name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"topic": {
						SchemaProps: spec.SchemaProps{
							Description: "Topic name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"url", "partition", "topic"},
			},
		},
		Dependencies: []string{},
	}
}