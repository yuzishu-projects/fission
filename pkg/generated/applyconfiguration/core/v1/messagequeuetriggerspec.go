/*
Copyright The Fission Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "github.com/fission/fission/pkg/apis/core/v1"
	apicorev1 "k8s.io/api/core/v1"
)

// MessageQueueTriggerSpecApplyConfiguration represents an declarative configuration of the MessageQueueTriggerSpec type for use
// with apply.
type MessageQueueTriggerSpecApplyConfiguration struct {
	FunctionReference *FunctionReferenceApplyConfiguration `json:"functionref,omitempty"`
	MessageQueueType  *corev1.MessageQueueType             `json:"messageQueueType,omitempty"`
	Topic             *string                              `json:"topic,omitempty"`
	ResponseTopic     *string                              `json:"respTopic,omitempty"`
	ErrorTopic        *string                              `json:"errorTopic,omitempty"`
	MaxRetries        *int                                 `json:"maxRetries,omitempty"`
	ContentType       *string                              `json:"contentType,omitempty"`
	PollingInterval   *int32                               `json:"pollingInterval,omitempty"`
	CooldownPeriod    *int32                               `json:"cooldownPeriod,omitempty"`
	MinReplicaCount   *int32                               `json:"minReplicaCount,omitempty"`
	MaxReplicaCount   *int32                               `json:"maxReplicaCount,omitempty"`
	Metadata          map[string]string                    `json:"metadata,omitempty"`
	Secret            *string                              `json:"secret,omitempty"`
	MqtKind           *string                              `json:"mqtkind,omitempty"`
	PodSpec           *apicorev1.PodSpec                   `json:"podspec,omitempty"`
}

// MessageQueueTriggerSpecApplyConfiguration constructs an declarative configuration of the MessageQueueTriggerSpec type for use with
// apply.
func MessageQueueTriggerSpec() *MessageQueueTriggerSpecApplyConfiguration {
	return &MessageQueueTriggerSpecApplyConfiguration{}
}

// WithFunctionReference sets the FunctionReference field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FunctionReference field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithFunctionReference(value *FunctionReferenceApplyConfiguration) *MessageQueueTriggerSpecApplyConfiguration {
	b.FunctionReference = value
	return b
}

// WithMessageQueueType sets the MessageQueueType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MessageQueueType field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithMessageQueueType(value corev1.MessageQueueType) *MessageQueueTriggerSpecApplyConfiguration {
	b.MessageQueueType = &value
	return b
}

// WithTopic sets the Topic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Topic field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithTopic(value string) *MessageQueueTriggerSpecApplyConfiguration {
	b.Topic = &value
	return b
}

// WithResponseTopic sets the ResponseTopic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResponseTopic field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithResponseTopic(value string) *MessageQueueTriggerSpecApplyConfiguration {
	b.ResponseTopic = &value
	return b
}

// WithErrorTopic sets the ErrorTopic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ErrorTopic field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithErrorTopic(value string) *MessageQueueTriggerSpecApplyConfiguration {
	b.ErrorTopic = &value
	return b
}

// WithMaxRetries sets the MaxRetries field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxRetries field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithMaxRetries(value int) *MessageQueueTriggerSpecApplyConfiguration {
	b.MaxRetries = &value
	return b
}

// WithContentType sets the ContentType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContentType field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithContentType(value string) *MessageQueueTriggerSpecApplyConfiguration {
	b.ContentType = &value
	return b
}

// WithPollingInterval sets the PollingInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PollingInterval field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithPollingInterval(value int32) *MessageQueueTriggerSpecApplyConfiguration {
	b.PollingInterval = &value
	return b
}

// WithCooldownPeriod sets the CooldownPeriod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CooldownPeriod field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithCooldownPeriod(value int32) *MessageQueueTriggerSpecApplyConfiguration {
	b.CooldownPeriod = &value
	return b
}

// WithMinReplicaCount sets the MinReplicaCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReplicaCount field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithMinReplicaCount(value int32) *MessageQueueTriggerSpecApplyConfiguration {
	b.MinReplicaCount = &value
	return b
}

// WithMaxReplicaCount sets the MaxReplicaCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxReplicaCount field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithMaxReplicaCount(value int32) *MessageQueueTriggerSpecApplyConfiguration {
	b.MaxReplicaCount = &value
	return b
}

// WithMetadata puts the entries into the Metadata field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Metadata field,
// overwriting an existing map entries in Metadata field with the same key.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithMetadata(entries map[string]string) *MessageQueueTriggerSpecApplyConfiguration {
	if b.Metadata == nil && len(entries) > 0 {
		b.Metadata = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Metadata[k] = v
	}
	return b
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithSecret(value string) *MessageQueueTriggerSpecApplyConfiguration {
	b.Secret = &value
	return b
}

// WithMqtKind sets the MqtKind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MqtKind field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithMqtKind(value string) *MessageQueueTriggerSpecApplyConfiguration {
	b.MqtKind = &value
	return b
}

// WithPodSpec sets the PodSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodSpec field is set to the value of the last call.
func (b *MessageQueueTriggerSpecApplyConfiguration) WithPodSpec(value apicorev1.PodSpec) *MessageQueueTriggerSpecApplyConfiguration {
	b.PodSpec = &value
	return b
}
