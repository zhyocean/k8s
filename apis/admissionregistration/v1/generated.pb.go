// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generated.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	generated.proto

It has these top-level messages:
	MutatingWebhook
	MutatingWebhookConfiguration
	MutatingWebhookConfigurationList
	Rule
	RuleWithOperations
	ServiceReference
	ValidatingWebhook
	ValidatingWebhookConfiguration
	ValidatingWebhookConfigurationList
	WebhookClientConfig
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import k8s_io_apimachinery_pkg_apis_meta_v1 "github.com/zhyocean/k8s/apis/meta/v1"
import _ "github.com/zhyocean/k8s/runtime"
import _ "github.com/zhyocean/k8s/runtime/schema"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MutatingWebhook describes an admission webhook and the resources and operations it applies to.
type MutatingWebhook struct {
	// The name of the admission webhook.
	// Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where
	// "imagepolicy" is the name of the webhook, and kubernetes.io is the name
	// of the organization.
	// Required.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// ClientConfig defines how to communicate with the hook.
	// Required
	ClientConfig *WebhookClientConfig `protobuf:"bytes,2,opt,name=clientConfig" json:"clientConfig,omitempty"`
	// Rules describes what operations on what resources/subresources the webhook cares about.
	// The webhook cares about an operation if it matches _any_ Rule.
	// However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks
	// from putting the cluster in a state which cannot be recovered from without completely
	// disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called
	// on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
	Rules []*RuleWithOperations `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
	// FailurePolicy defines how unrecognized errors from the admission endpoint are handled -
	// allowed values are Ignore or Fail. Defaults to Fail.
	// +optional
	FailurePolicy *string `protobuf:"bytes,4,opt,name=failurePolicy" json:"failurePolicy,omitempty"`
	// matchPolicy defines how the "rules" list is used to match incoming requests.
	// Allowed values are "Exact" or "Equivalent".
	//
	// - Exact: match a request only if it exactly matches a specified rule.
	// For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1,
	// but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`,
	// a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.
	//
	// - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version.
	// For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1,
	// and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`,
	// a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.
	//
	// Defaults to "Equivalent"
	// +optional
	MatchPolicy *string `protobuf:"bytes,9,opt,name=matchPolicy" json:"matchPolicy,omitempty"`
	// NamespaceSelector decides whether to run the webhook on an object based
	// on whether the namespace for that object matches the selector. If the
	// object itself is a namespace, the matching is performed on
	// object.metadata.labels. If the object is another cluster scoped resource,
	// it never skips the webhook.
	//
	// For example, to run the webhook on any objects whose namespace is not
	// associated with "runlevel" of "0" or "1";  you will set the selector as
	// follows:
	// "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "runlevel",
	//       "operator": "NotIn",
	//       "values": [
	//         "0",
	//         "1"
	//       ]
	//     }
	//   ]
	// }
	//
	// If instead you want to only run the webhook on any objects whose
	// namespace is associated with the "environment" of "prod" or "staging";
	// you will set the selector as follows:
	// "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "environment",
	//       "operator": "In",
	//       "values": [
	//         "prod",
	//         "staging"
	//       ]
	//     }
	//   ]
	// }
	//
	// See
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	// for more examples of label selectors.
	//
	// Default to the empty LabelSelector, which matches everything.
	// +optional
	NamespaceSelector *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector `protobuf:"bytes,5,opt,name=namespaceSelector" json:"namespaceSelector,omitempty"`
	// ObjectSelector decides whether to run the webhook based on if the
	// object has matching labels. objectSelector is evaluated against both
	// the oldObject and newObject that would be sent to the webhook, and
	// is considered to match if either object matches the selector. A null
	// object (oldObject in the case of create, or newObject in the case of
	// delete) or an object that cannot have labels (like a
	// DeploymentRollback or a PodProxyOptions object) is not considered to
	// match.
	// Use the object selector only if the webhook is opt-in, because end
	// users may skip the admission webhook by setting the labels.
	// Default to the empty LabelSelector, which matches everything.
	// +optional
	ObjectSelector *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector `protobuf:"bytes,11,opt,name=objectSelector" json:"objectSelector,omitempty"`
	// SideEffects states whether this webhook has side effects.
	// Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown).
	// Webhooks with side effects MUST implement a reconciliation system, since a request may be
	// rejected by a future step in the admission change and the side effects therefore need to be undone.
	// Requests with the dryRun attribute will be auto-rejected if they match a webhook with
	// sideEffects == Unknown or Some.
	SideEffects *string `protobuf:"bytes,6,opt,name=sideEffects" json:"sideEffects,omitempty"`
	// TimeoutSeconds specifies the timeout for this webhook. After the timeout passes,
	// the webhook call will be ignored or the API call will fail based on the
	// failure policy.
	// The timeout value must be between 1 and 30 seconds.
	// Default to 10 seconds.
	// +optional
	TimeoutSeconds *int32 `protobuf:"varint,7,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	// AdmissionReviewVersions is an ordered list of preferred `AdmissionReview`
	// versions the Webhook expects. API server will try to use first version in
	// the list which it supports. If none of the versions specified in this list
	// supported by API server, validation will fail for this object.
	// If a persisted webhook configuration specifies allowed versions and does not
	// include any versions known to the API Server, calls to the webhook will fail
	// and be subject to the failure policy.
	AdmissionReviewVersions []string `protobuf:"bytes,8,rep,name=admissionReviewVersions" json:"admissionReviewVersions,omitempty"`
	// reinvocationPolicy indicates whether this webhook should be called multiple times as part of a single admission evaluation.
	// Allowed values are "Never" and "IfNeeded".
	//
	// Never: the webhook will not be called more than once in a single admission evaluation.
	//
	// IfNeeded: the webhook will be called at least one additional time as part of the admission evaluation
	// if the object being admitted is modified by other admission plugins after the initial webhook call.
	// Webhooks that specify this option *must* be idempotent, able to process objects they previously admitted.
	// Note:
	// * the number of additional invocations is not guaranteed to be exactly one.
	// * if additional invocations result in further modifications to the object, webhooks are not guaranteed to be invoked again.
	// * webhooks that use this option may be reordered to minimize the number of additional invocations.
	// * to validate an object after all mutations are guaranteed complete, use a validating admission webhook instead.
	//
	// Defaults to "Never".
	// +optional
	ReinvocationPolicy *string `protobuf:"bytes,10,opt,name=reinvocationPolicy" json:"reinvocationPolicy,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *MutatingWebhook) Reset()                    { *m = MutatingWebhook{} }
func (m *MutatingWebhook) String() string            { return proto.CompactTextString(m) }
func (*MutatingWebhook) ProtoMessage()               {}
func (*MutatingWebhook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MutatingWebhook) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MutatingWebhook) GetClientConfig() *WebhookClientConfig {
	if m != nil {
		return m.ClientConfig
	}
	return nil
}

func (m *MutatingWebhook) GetRules() []*RuleWithOperations {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *MutatingWebhook) GetFailurePolicy() string {
	if m != nil && m.FailurePolicy != nil {
		return *m.FailurePolicy
	}
	return ""
}

func (m *MutatingWebhook) GetMatchPolicy() string {
	if m != nil && m.MatchPolicy != nil {
		return *m.MatchPolicy
	}
	return ""
}

func (m *MutatingWebhook) GetNamespaceSelector() *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector {
	if m != nil {
		return m.NamespaceSelector
	}
	return nil
}

func (m *MutatingWebhook) GetObjectSelector() *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector {
	if m != nil {
		return m.ObjectSelector
	}
	return nil
}

func (m *MutatingWebhook) GetSideEffects() string {
	if m != nil && m.SideEffects != nil {
		return *m.SideEffects
	}
	return ""
}

func (m *MutatingWebhook) GetTimeoutSeconds() int32 {
	if m != nil && m.TimeoutSeconds != nil {
		return *m.TimeoutSeconds
	}
	return 0
}

func (m *MutatingWebhook) GetAdmissionReviewVersions() []string {
	if m != nil {
		return m.AdmissionReviewVersions
	}
	return nil
}

func (m *MutatingWebhook) GetReinvocationPolicy() string {
	if m != nil && m.ReinvocationPolicy != nil {
		return *m.ReinvocationPolicy
	}
	return ""
}

// MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object.
type MutatingWebhookConfiguration struct {
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	Webhooks         []*MutatingWebhook `protobuf:"bytes,2,rep,name=Webhooks" json:"Webhooks,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *MutatingWebhookConfiguration) Reset()                    { *m = MutatingWebhookConfiguration{} }
func (m *MutatingWebhookConfiguration) String() string            { return proto.CompactTextString(m) }
func (*MutatingWebhookConfiguration) ProtoMessage()               {}
func (*MutatingWebhookConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MutatingWebhookConfiguration) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MutatingWebhookConfiguration) GetWebhooks() []*MutatingWebhook {
	if m != nil {
		return m.Webhooks
	}
	return nil
}

// MutatingWebhookConfigurationList is a list of MutatingWebhookConfiguration.
type MutatingWebhookConfigurationList struct {
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// List of MutatingWebhookConfiguration.
	Items            []*MutatingWebhookConfiguration `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *MutatingWebhookConfigurationList) Reset()         { *m = MutatingWebhookConfigurationList{} }
func (m *MutatingWebhookConfigurationList) String() string { return proto.CompactTextString(m) }
func (*MutatingWebhookConfigurationList) ProtoMessage()    {}
func (*MutatingWebhookConfigurationList) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

func (m *MutatingWebhookConfigurationList) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MutatingWebhookConfigurationList) GetItems() []*MutatingWebhookConfiguration {
	if m != nil {
		return m.Items
	}
	return nil
}

// Rule is a tuple of APIGroups, APIVersion, and Resources.It is recommended
// to make sure that all the tuple expansions are valid.
type Rule struct {
	// APIGroups is the API groups the resources belong to. '*' is all groups.
	// If '*' is present, the length of the slice must be one.
	// Required.
	ApiGroups []string `protobuf:"bytes,1,rep,name=apiGroups" json:"apiGroups,omitempty"`
	// APIVersions is the API versions the resources belong to. '*' is all versions.
	// If '*' is present, the length of the slice must be one.
	// Required.
	ApiVersions []string `protobuf:"bytes,2,rep,name=apiVersions" json:"apiVersions,omitempty"`
	// Resources is a list of resources this rule applies to.
	//
	// For example:
	// 'pods' means pods.
	// 'pods/log' means the log subresource of pods.
	// '*' means all resources, but not subresources.
	// 'pods/*' means all subresources of pods.
	// '*/scale' means all scale subresources.
	// '*/*' means all resources and their subresources.
	//
	// If wildcard is present, the validation rule will ensure resources do not
	// overlap with each other.
	//
	// Depending on the enclosing object, subresources might not be allowed.
	// Required.
	Resources []string `protobuf:"bytes,3,rep,name=resources" json:"resources,omitempty"`
	// scope specifies the scope of this rule.
	// Valid values are "Cluster", "Namespaced", and "*"
	// "Cluster" means that only cluster-scoped resources will match this rule.
	// Namespace API objects are cluster-scoped.
	// "Namespaced" means that only namespaced resources will match this rule.
	// "*" means that there are no scope restrictions.
	// Subresources match the scope of their parent resource.
	// Default is "*".
	//
	// +optional
	Scope            *string `protobuf:"bytes,4,opt,name=scope" json:"scope,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Rule) GetApiGroups() []string {
	if m != nil {
		return m.ApiGroups
	}
	return nil
}

func (m *Rule) GetApiVersions() []string {
	if m != nil {
		return m.ApiVersions
	}
	return nil
}

func (m *Rule) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Rule) GetScope() string {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return ""
}

// RuleWithOperations is a tuple of Operations and Resources. It is recommended to make
// sure that all the tuple expansions are valid.
type RuleWithOperations struct {
	// Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or *
	// for all of those operations and any future admission operations that are added.
	// If '*' is present, the length of the slice must be one.
	// Required.
	Operations []string `protobuf:"bytes,1,rep,name=operations" json:"operations,omitempty"`
	// Rule is embedded, it describes other criteria of the rule, like
	// APIGroups, APIVersions, Resources, etc.
	Rule             *Rule  `protobuf:"bytes,2,opt,name=rule" json:"rule,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RuleWithOperations) Reset()                    { *m = RuleWithOperations{} }
func (m *RuleWithOperations) String() string            { return proto.CompactTextString(m) }
func (*RuleWithOperations) ProtoMessage()               {}
func (*RuleWithOperations) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RuleWithOperations) GetOperations() []string {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *RuleWithOperations) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
	// `namespace` is the namespace of the service.
	// Required
	Namespace *string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	// `name` is the name of the service.
	// Required
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// `path` is an optional URL path which will be sent in any request to
	// this service.
	// +optional
	Path *string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// If specified, the port on the service that hosting webhook.
	// Default to 443 for backward compatibility.
	// `port` should be a valid port number (1-65535, inclusive).
	// +optional
	Port             *int32 `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ServiceReference) Reset()                    { *m = ServiceReference{} }
func (m *ServiceReference) String() string            { return proto.CompactTextString(m) }
func (*ServiceReference) ProtoMessage()               {}
func (*ServiceReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServiceReference) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

func (m *ServiceReference) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ServiceReference) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *ServiceReference) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

// ValidatingWebhook describes an admission webhook and the resources and operations it applies to.
type ValidatingWebhook struct {
	// The name of the admission webhook.
	// Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where
	// "imagepolicy" is the name of the webhook, and kubernetes.io is the name
	// of the organization.
	// Required.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// ClientConfig defines how to communicate with the hook.
	// Required
	ClientConfig *WebhookClientConfig `protobuf:"bytes,2,opt,name=clientConfig" json:"clientConfig,omitempty"`
	// Rules describes what operations on what resources/subresources the webhook cares about.
	// The webhook cares about an operation if it matches _any_ Rule.
	// However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks
	// from putting the cluster in a state which cannot be recovered from without completely
	// disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called
	// on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
	Rules []*RuleWithOperations `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
	// FailurePolicy defines how unrecognized errors from the admission endpoint are handled -
	// allowed values are Ignore or Fail. Defaults to Fail.
	// +optional
	FailurePolicy *string `protobuf:"bytes,4,opt,name=failurePolicy" json:"failurePolicy,omitempty"`
	// matchPolicy defines how the "rules" list is used to match incoming requests.
	// Allowed values are "Exact" or "Equivalent".
	//
	// - Exact: match a request only if it exactly matches a specified rule.
	// For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1,
	// but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`,
	// a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.
	//
	// - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version.
	// For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1,
	// and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`,
	// a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.
	//
	// Defaults to "Equivalent"
	// +optional
	MatchPolicy *string `protobuf:"bytes,9,opt,name=matchPolicy" json:"matchPolicy,omitempty"`
	// NamespaceSelector decides whether to run the webhook on an object based
	// on whether the namespace for that object matches the selector. If the
	// object itself is a namespace, the matching is performed on
	// object.metadata.labels. If the object is another cluster scoped resource,
	// it never skips the webhook.
	//
	// For example, to run the webhook on any objects whose namespace is not
	// associated with "runlevel" of "0" or "1";  you will set the selector as
	// follows:
	// "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "runlevel",
	//       "operator": "NotIn",
	//       "values": [
	//         "0",
	//         "1"
	//       ]
	//     }
	//   ]
	// }
	//
	// If instead you want to only run the webhook on any objects whose
	// namespace is associated with the "environment" of "prod" or "staging";
	// you will set the selector as follows:
	// "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "environment",
	//       "operator": "In",
	//       "values": [
	//         "prod",
	//         "staging"
	//       ]
	//     }
	//   ]
	// }
	//
	// See
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// for more examples of label selectors.
	//
	// Default to the empty LabelSelector, which matches everything.
	// +optional
	NamespaceSelector *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector `protobuf:"bytes,5,opt,name=namespaceSelector" json:"namespaceSelector,omitempty"`
	// ObjectSelector decides whether to run the webhook based on if the
	// object has matching labels. objectSelector is evaluated against both
	// the oldObject and newObject that would be sent to the webhook, and
	// is considered to match if either object matches the selector. A null
	// object (oldObject in the case of create, or newObject in the case of
	// delete) or an object that cannot have labels (like a
	// DeploymentRollback or a PodProxyOptions object) is not considered to
	// match.
	// Use the object selector only if the webhook is opt-in, because end
	// users may skip the admission webhook by setting the labels.
	// Default to the empty LabelSelector, which matches everything.
	// +optional
	ObjectSelector *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector `protobuf:"bytes,10,opt,name=objectSelector" json:"objectSelector,omitempty"`
	// SideEffects states whether this webhook has side effects.
	// Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown).
	// Webhooks with side effects MUST implement a reconciliation system, since a request may be
	// rejected by a future step in the admission change and the side effects therefore need to be undone.
	// Requests with the dryRun attribute will be auto-rejected if they match a webhook with
	// sideEffects == Unknown or Some.
	SideEffects *string `protobuf:"bytes,6,opt,name=sideEffects" json:"sideEffects,omitempty"`
	// TimeoutSeconds specifies the timeout for this webhook. After the timeout passes,
	// the webhook call will be ignored or the API call will fail based on the
	// failure policy.
	// The timeout value must be between 1 and 30 seconds.
	// Default to 10 seconds.
	// +optional
	TimeoutSeconds *int32 `protobuf:"varint,7,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	// AdmissionReviewVersions is an ordered list of preferred `AdmissionReview`
	// versions the Webhook expects. API server will try to use first version in
	// the list which it supports. If none of the versions specified in this list
	// supported by API server, validation will fail for this object.
	// If a persisted webhook configuration specifies allowed versions and does not
	// include any versions known to the API Server, calls to the webhook will fail
	// and be subject to the failure policy.
	AdmissionReviewVersions []string `protobuf:"bytes,8,rep,name=admissionReviewVersions" json:"admissionReviewVersions,omitempty"`
	XXX_unrecognized        []byte   `json:"-"`
}

func (m *ValidatingWebhook) Reset()                    { *m = ValidatingWebhook{} }
func (m *ValidatingWebhook) String() string            { return proto.CompactTextString(m) }
func (*ValidatingWebhook) ProtoMessage()               {}
func (*ValidatingWebhook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ValidatingWebhook) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ValidatingWebhook) GetClientConfig() *WebhookClientConfig {
	if m != nil {
		return m.ClientConfig
	}
	return nil
}

func (m *ValidatingWebhook) GetRules() []*RuleWithOperations {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *ValidatingWebhook) GetFailurePolicy() string {
	if m != nil && m.FailurePolicy != nil {
		return *m.FailurePolicy
	}
	return ""
}

func (m *ValidatingWebhook) GetMatchPolicy() string {
	if m != nil && m.MatchPolicy != nil {
		return *m.MatchPolicy
	}
	return ""
}

func (m *ValidatingWebhook) GetNamespaceSelector() *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector {
	if m != nil {
		return m.NamespaceSelector
	}
	return nil
}

func (m *ValidatingWebhook) GetObjectSelector() *k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector {
	if m != nil {
		return m.ObjectSelector
	}
	return nil
}

func (m *ValidatingWebhook) GetSideEffects() string {
	if m != nil && m.SideEffects != nil {
		return *m.SideEffects
	}
	return ""
}

func (m *ValidatingWebhook) GetTimeoutSeconds() int32 {
	if m != nil && m.TimeoutSeconds != nil {
		return *m.TimeoutSeconds
	}
	return 0
}

func (m *ValidatingWebhook) GetAdmissionReviewVersions() []string {
	if m != nil {
		return m.AdmissionReviewVersions
	}
	return nil
}

// ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.
type ValidatingWebhookConfiguration struct {
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	Webhooks         []*ValidatingWebhook `protobuf:"bytes,2,rep,name=Webhooks" json:"Webhooks,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ValidatingWebhookConfiguration) Reset()                    { *m = ValidatingWebhookConfiguration{} }
func (m *ValidatingWebhookConfiguration) String() string            { return proto.CompactTextString(m) }
func (*ValidatingWebhookConfiguration) ProtoMessage()               {}
func (*ValidatingWebhookConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ValidatingWebhookConfiguration) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ValidatingWebhookConfiguration) GetWebhooks() []*ValidatingWebhook {
	if m != nil {
		return m.Webhooks
	}
	return nil
}

// ValidatingWebhookConfigurationList is a list of ValidatingWebhookConfiguration.
type ValidatingWebhookConfigurationList struct {
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// List of ValidatingWebhookConfiguration.
	Items            []*ValidatingWebhookConfiguration `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *ValidatingWebhookConfigurationList) Reset()         { *m = ValidatingWebhookConfigurationList{} }
func (m *ValidatingWebhookConfigurationList) String() string { return proto.CompactTextString(m) }
func (*ValidatingWebhookConfigurationList) ProtoMessage()    {}
func (*ValidatingWebhookConfigurationList) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8}
}

func (m *ValidatingWebhookConfigurationList) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ValidatingWebhookConfigurationList) GetItems() []*ValidatingWebhookConfiguration {
	if m != nil {
		return m.Items
	}
	return nil
}

// WebhookClientConfig contains the information to make a TLS
// connection with the webhook
type WebhookClientConfig struct {
	// `url` gives the location of the webhook, in standard URL form
	// (`scheme://host:port/path`). Exactly one of `url` or `service`
	// must be specified.
	//
	// The `host` should not refer to a service running in the cluster; use
	// the `service` field instead. The host might be resolved via external
	// DNS in some apiservers (e.g., `kube-apiserver` cannot resolve
	// in-cluster DNS as that would be a layering violation). `host` may
	// also be an IP address.
	//
	// Please note that using `localhost` or `127.0.0.1` as a `host` is
	// risky unless you take great care to run this webhook on all hosts
	// which run an apiserver which might need to make calls to this
	// webhook. Such installs are likely to be non-portable, i.e., not easy
	// to turn up in a new cluster.
	//
	// The scheme must be "https"; the URL must begin with "https://".
	//
	// A path is optional, and if present may be any string permissible in
	// a URL. You may use the path to pass an arbitrary string to the
	// webhook, for example, a cluster identifier.
	//
	// Attempting to use a user or basic auth e.g. "user:password@" is not
	// allowed. Fragments ("#...") and query parameters ("?...") are not
	// allowed, either.
	//
	// +optional
	Url *string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	// `service` is a reference to the service for this webhook. Either
	// `service` or `url` must be specified.
	//
	// If the webhook is running within the cluster, then you should use `service`.
	//
	// +optional
	Service *ServiceReference `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	// `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate.
	// If unspecified, system trust roots on the apiserver are used.
	// +optional
	CaBundle         []byte `protobuf:"bytes,2,opt,name=caBundle" json:"caBundle,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *WebhookClientConfig) Reset()                    { *m = WebhookClientConfig{} }
func (m *WebhookClientConfig) String() string            { return proto.CompactTextString(m) }
func (*WebhookClientConfig) ProtoMessage()               {}
func (*WebhookClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *WebhookClientConfig) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *WebhookClientConfig) GetService() *ServiceReference {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *WebhookClientConfig) GetCaBundle() []byte {
	if m != nil {
		return m.CaBundle
	}
	return nil
}

func init() {
	proto.RegisterType((*MutatingWebhook)(nil), "k8s.io.api.admissionregistration.v1.MutatingWebhook")
	proto.RegisterType((*MutatingWebhookConfiguration)(nil), "k8s.io.api.admissionregistration.v1.MutatingWebhookConfiguration")
	proto.RegisterType((*MutatingWebhookConfigurationList)(nil), "k8s.io.api.admissionregistration.v1.MutatingWebhookConfigurationList")
	proto.RegisterType((*Rule)(nil), "k8s.io.api.admissionregistration.v1.Rule")
	proto.RegisterType((*RuleWithOperations)(nil), "k8s.io.api.admissionregistration.v1.RuleWithOperations")
	proto.RegisterType((*ServiceReference)(nil), "k8s.io.api.admissionregistration.v1.ServiceReference")
	proto.RegisterType((*ValidatingWebhook)(nil), "k8s.io.api.admissionregistration.v1.ValidatingWebhook")
	proto.RegisterType((*ValidatingWebhookConfiguration)(nil), "k8s.io.api.admissionregistration.v1.ValidatingWebhookConfiguration")
	proto.RegisterType((*ValidatingWebhookConfigurationList)(nil), "k8s.io.api.admissionregistration.v1.ValidatingWebhookConfigurationList")
	proto.RegisterType((*WebhookClientConfig)(nil), "k8s.io.api.admissionregistration.v1.WebhookClientConfig")
}

func init() { proto.RegisterFile("generated.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xcd, 0x6e, 0xf3, 0x44,
	0x14, 0x95, 0xf3, 0xc3, 0x97, 0xdc, 0x7c, 0xf4, 0x67, 0x40, 0xc2, 0xaa, 0xaa, 0x2a, 0x32, 0x08,
	0x85, 0x8d, 0x43, 0x0a, 0x94, 0x6c, 0x58, 0xd0, 0x0a, 0x21, 0xa1, 0x56, 0xad, 0xa6, 0x52, 0x2b,
	0x7e, 0x36, 0x93, 0xc9, 0x4d, 0x3c, 0xc4, 0xf6, 0x58, 0x33, 0xe3, 0xa0, 0xf2, 0x10, 0xac, 0x78,
	0x1e, 0x16, 0x6c, 0x58, 0xf0, 0x12, 0x3c, 0x0a, 0xf2, 0xd8, 0x75, 0x1d, 0xa7, 0x2d, 0x56, 0x05,
	0x12, 0x42, 0xec, 0x66, 0x8e, 0x73, 0x4f, 0xee, 0x99, 0x7b, 0xee, 0x81, 0xdd, 0x25, 0xc6, 0xa8,
	0x98, 0xc1, 0xb9, 0x9f, 0x28, 0x69, 0x24, 0x79, 0x77, 0x35, 0xd5, 0xbe, 0x90, 0x3e, 0x4b, 0x84,
	0xcf, 0xe6, 0x91, 0xd0, 0x5a, 0xc8, 0x58, 0xe1, 0x52, 0x68, 0xa3, 0x98, 0x11, 0x32, 0xf6, 0xd7,
	0x93, 0x83, 0x93, 0xa5, 0x30, 0x41, 0x3a, 0xf3, 0xb9, 0x8c, 0xc6, 0xa8, 0x04, 0xe7, 0x81, 0x60,
	0xf1, 0x72, 0xbc, 0x9a, 0xea, 0x31, 0x4b, 0x84, 0x1e, 0x47, 0x68, 0xd8, 0x78, 0x3d, 0x19, 0xd7,
	0xc8, 0x0f, 0x26, 0x4f, 0xd7, 0xa9, 0x34, 0x36, 0x22, 0xc2, 0xad, 0x92, 0xe9, 0x5f, 0x97, 0x68,
	0x1e, 0x60, 0xc4, 0xea, 0x95, 0xde, 0x4f, 0x5d, 0xd8, 0xbd, 0x48, 0x0d, 0x33, 0x22, 0x5e, 0xde,
	0xe2, 0x2c, 0x90, 0x72, 0x45, 0x08, 0x74, 0x62, 0x16, 0xa1, 0xeb, 0x0c, 0x9d, 0x51, 0x9f, 0xda,
	0x33, 0xf9, 0x0e, 0x5e, 0xf3, 0x50, 0x60, 0x6c, 0xce, 0x64, 0xbc, 0x10, 0x4b, 0xb7, 0x35, 0x74,
	0x46, 0x83, 0xe3, 0xa9, 0xdf, 0xe0, 0x21, 0xfc, 0x82, 0xf7, 0xac, 0x52, 0x4f, 0x37, 0xd8, 0xc8,
	0x05, 0x74, 0x55, 0x1a, 0xa2, 0x76, 0xdb, 0xc3, 0xf6, 0x68, 0x70, 0xfc, 0x69, 0x23, 0x5a, 0x9a,
	0x86, 0x78, 0x2b, 0x4c, 0x70, 0x99, 0x60, 0x0e, 0x6a, 0x9a, 0xb3, 0x90, 0xf7, 0xe0, 0xcd, 0x05,
	0x13, 0x61, 0xaa, 0xf0, 0x4a, 0x86, 0x82, 0xdf, 0xb9, 0x1d, 0xab, 0x64, 0x13, 0x24, 0x43, 0x18,
	0x44, 0xcc, 0xf0, 0xa0, 0xf8, 0x4d, 0xdf, 0xfe, 0xa6, 0x0a, 0x11, 0x06, 0xfb, 0x99, 0x78, 0x9d,
	0x30, 0x8e, 0xd7, 0x18, 0x22, 0x37, 0x52, 0xb9, 0x5d, 0xab, 0xfc, 0xa3, 0x4a, 0x8b, 0x11, 0xe3,
	0x81, 0x88, 0x51, 0xdd, 0xf9, 0xc9, 0x6a, 0x99, 0x01, 0xda, 0xcf, 0x66, 0x9b, 0xf5, 0x78, 0xce,
	0x66, 0x18, 0xde, 0x97, 0xd2, 0x6d, 0x36, 0xf2, 0x2d, 0xec, 0xc8, 0xd9, 0xf7, 0xc8, 0x4d, 0xc9,
	0x3f, 0x78, 0x39, 0x7f, 0x8d, 0x2a, 0x53, 0xa8, 0xc5, 0x1c, 0xbf, 0x58, 0x2c, 0x90, 0x1b, 0xed,
	0xbe, 0x91, 0x2b, 0xac, 0x40, 0xe4, 0x7d, 0xd8, 0xc9, 0xdc, 0x21, 0x53, 0x73, 0x8d, 0x5c, 0xc6,
	0x73, 0xed, 0xbe, 0x1a, 0x3a, 0xa3, 0x2e, 0xad, 0xa1, 0x64, 0x0a, 0xef, 0x94, 0x73, 0xa0, 0xb8,
	0x16, 0xf8, 0xc3, 0x0d, 0xaa, 0xec, 0xa2, 0xdd, 0xde, 0xb0, 0x3d, 0xea, 0xd3, 0xa7, 0x3e, 0x13,
	0x1f, 0x88, 0x42, 0x11, 0xaf, 0x25, 0xb7, 0x33, 0x2a, 0x1e, 0x1b, 0x6c, 0x2b, 0x8f, 0x7c, 0xf1,
	0x7e, 0x71, 0xe0, 0xb0, 0x66, 0xc8, 0xdc, 0x24, 0x69, 0x3e, 0x64, 0x72, 0x0e, 0xbd, 0x4c, 0xfd,
	0x9c, 0x19, 0x66, 0x1d, 0x3a, 0x38, 0xfe, 0xb0, 0xd9, 0x5b, 0x5d, 0xda, 0xc7, 0xb9, 0x40, 0xc3,
	0x68, 0xc9, 0x40, 0xae, 0xa0, 0x57, 0xfc, 0x8b, 0x76, 0x5b, 0xd6, 0x7c, 0x1f, 0x37, 0x32, 0x5f,
	0xad, 0x45, 0x5a, 0xb2, 0x78, 0xbf, 0x39, 0x30, 0x7c, 0x4e, 0xc0, 0xb9, 0xd0, 0x86, 0x7c, 0xb5,
	0x25, 0xc2, 0x6f, 0x38, 0x70, 0xa1, 0xeb, 0x12, 0x6e, 0xa1, 0x2b, 0x0c, 0x46, 0xf7, 0xfd, 0x7f,
	0xfe, 0x92, 0xfe, 0x37, 0x3a, 0xa4, 0x39, 0x9f, 0xf7, 0x23, 0x74, 0xb2, 0x1d, 0x23, 0x87, 0xd0,
	0x67, 0x89, 0xf8, 0x52, 0xc9, 0x34, 0xd1, 0xae, 0x63, 0xc7, 0xfd, 0x00, 0x64, 0x26, 0x63, 0x89,
	0x28, 0xed, 0xd0, 0xb2, 0xdf, 0xab, 0x50, 0x56, 0xaf, 0x50, 0xcb, 0x54, 0xf1, 0x62, 0xc3, 0xfb,
	0xf4, 0x01, 0x20, 0x6f, 0x43, 0x57, 0x73, 0x99, 0x60, 0xb1, 0xa4, 0xf9, 0xc5, 0xd3, 0x40, 0xb6,
	0xf7, 0x9b, 0x1c, 0x01, 0xc8, 0xf2, 0x56, 0xb4, 0x52, 0x41, 0xc8, 0x67, 0xd0, 0xc9, 0x12, 0xa0,
	0x48, 0xa7, 0x0f, 0x1a, 0xc7, 0x08, 0xb5, 0x65, 0x5e, 0x08, 0x7b, 0xd7, 0xa8, 0xd6, 0x82, 0x23,
	0xc5, 0x05, 0x2a, 0x8c, 0xb9, 0x15, 0x5f, 0x6e, 0x6d, 0x91, 0x88, 0x0f, 0x40, 0x19, 0x95, 0xad,
	0x4a, 0x54, 0x12, 0xe8, 0x24, 0xcc, 0x04, 0x6e, 0x3b, 0xc7, 0xb2, 0xb3, 0xc5, 0xa4, 0x32, 0x56,
	0x63, 0x97, 0xda, 0xb3, 0xf7, 0x47, 0x07, 0xf6, 0x6f, 0x58, 0x28, 0xe6, 0xff, 0x87, 0xef, 0xbf,
	0x26, 0x7c, 0xe1, 0x3f, 0x15, 0xbe, 0xde, 0xaf, 0x0e, 0x1c, 0x6d, 0x59, 0xec, 0x9f, 0x8c, 0x53,
	0xba, 0x15, 0xa7, 0x27, 0x8d, 0xec, 0xb4, 0xd5, 0x64, 0x25, 0x50, 0x7f, 0x77, 0xc0, 0x7b, 0x5e,
	0xc4, 0xdf, 0x1e, 0xa9, 0x5f, 0x6f, 0x46, 0xea, 0xd9, 0xcb, 0x34, 0x3c, 0x1a, 0xaa, 0x3f, 0x3b,
	0xf0, 0xd6, 0x23, 0x3b, 0x49, 0xf6, 0xa0, 0x9d, 0xaa, 0xb0, 0x08, 0x8d, 0xec, 0x48, 0x2e, 0xe1,
	0x95, 0xce, 0xd3, 0xa8, 0xd0, 0xf3, 0x49, 0xa3, 0x36, 0xea, 0x09, 0x46, 0xef, 0x59, 0xc8, 0x01,
	0xf4, 0x38, 0x3b, 0x4d, 0xe3, 0x79, 0x91, 0x90, 0xaf, 0x69, 0x79, 0x3f, 0xed, 0x7c, 0xd3, 0x5a,
	0x4f, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xde, 0xa2, 0x8a, 0x50, 0xe9, 0x0a, 0x00, 0x00,
}
