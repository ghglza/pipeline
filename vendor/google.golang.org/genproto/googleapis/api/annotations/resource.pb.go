// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/resource.proto

package annotations

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A description of the historical or future-looking state of the
// resource pattern.
type ResourceDescriptor_History int32

const (
	// The "unset" value.
	ResourceDescriptor_HISTORY_UNSPECIFIED ResourceDescriptor_History = 0
	// The resource originally had one pattern and launched as such, and
	// additional patterns were added later.
	ResourceDescriptor_ORIGINALLY_SINGLE_PATTERN ResourceDescriptor_History = 1
	// The resource has one pattern, but the API owner expects to add more
	// later. (This is the inverse of ORIGINALLY_SINGLE_PATTERN, and prevents
	// that from being necessary once there are multiple patterns.)
	ResourceDescriptor_FUTURE_MULTI_PATTERN ResourceDescriptor_History = 2
)

var ResourceDescriptor_History_name = map[int32]string{
	0: "HISTORY_UNSPECIFIED",
	1: "ORIGINALLY_SINGLE_PATTERN",
	2: "FUTURE_MULTI_PATTERN",
}

var ResourceDescriptor_History_value = map[string]int32{
	"HISTORY_UNSPECIFIED":       0,
	"ORIGINALLY_SINGLE_PATTERN": 1,
	"FUTURE_MULTI_PATTERN":      2,
}

func (x ResourceDescriptor_History) String() string {
	return proto.EnumName(ResourceDescriptor_History_name, int32(x))
}

func (ResourceDescriptor_History) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_465e9122405d1bb5, []int{0, 0}
}

// A simple descriptor of a resource type.
//
// ResourceDescriptor annotates a resource message (either by means of a
// protobuf annotation or use in the service config), and associates the
// resource's schema, the resource type, and the pattern of the resource name.
//
// Example:
//
//     message Topic {
//       // Indicates this message defines a resource schema.
//       // Declares the resource type in the format of {service}/{kind}.
//       // For Kubernetes resources, the format is {api group}/{kind}.
//       option (google.api.resource) = {
//         type: "pubsub.googleapis.com/Topic"
//         name_descriptor: {
//           pattern: "projects/{project}/topics/{topic}"
//           parent_type: "cloudresourcemanager.googleapis.com/Project"
//           parent_name_extractor: "projects/{project}"
//         }
//       };
//     }
//
// The ResourceDescriptor Yaml config will look like:
//
//     resources:
//     - type: "pubsub.googleapis.com/Topic"
//       name_descriptor:
//         - pattern: "projects/{project}/topics/{topic}"
//           parent_type: "cloudresourcemanager.googleapis.com/Project"
//           parent_name_extractor: "projects/{project}"
//
// Sometimes, resources have multiple patterns, typically because they can
// live under multiple parents.
//
// Example:
//
//     message LogEntry {
//       option (google.api.resource) = {
//         type: "logging.googleapis.com/LogEntry"
//         name_descriptor: {
//           pattern: "projects/{project}/logs/{log}"
//           parent_type: "cloudresourcemanager.googleapis.com/Project"
//           parent_name_extractor: "projects/{project}"
//         }
//         name_descriptor: {
//           pattern: "folders/{folder}/logs/{log}"
//           parent_type: "cloudresourcemanager.googleapis.com/Folder"
//           parent_name_extractor: "folders/{folder}"
//         }
//         name_descriptor: {
//           pattern: "organizations/{organization}/logs/{log}"
//           parent_type: "cloudresourcemanager.googleapis.com/Organization"
//           parent_name_extractor: "organizations/{organization}"
//         }
//         name_descriptor: {
//           pattern: "billingAccounts/{billing_account}/logs/{log}"
//           parent_type: "billing.googleapis.com/BillingAccount"
//           parent_name_extractor: "billingAccounts/{billing_account}"
//         }
//       };
//     }
//
// The ResourceDescriptor Yaml config will look like:
//
//     resources:
//     - type: 'logging.googleapis.com/LogEntry'
//       name_descriptor:
//         - pattern: "projects/{project}/logs/{log}"
//           parent_type: "cloudresourcemanager.googleapis.com/Project"
//           parent_name_extractor: "projects/{project}"
//         - pattern: "folders/{folder}/logs/{log}"
//           parent_type: "cloudresourcemanager.googleapis.com/Folder"
//           parent_name_extractor: "folders/{folder}"
//         - pattern: "organizations/{organization}/logs/{log}"
//           parent_type: "cloudresourcemanager.googleapis.com/Organization"
//           parent_name_extractor: "organizations/{organization}"
//         - pattern: "billingAccounts/{billing_account}/logs/{log}"
//           parent_type: "billing.googleapis.com/BillingAccount"
//           parent_name_extractor: "billingAccounts/{billing_account}"
//
// For flexible resources, the resource name doesn't contain parent names, but
// the resource itself has parents for policy evaluation.
//
// Example:
//
//     message Shelf {
//       option (google.api.resource) = {
//         type: "library.googleapis.com/Shelf"
//         name_descriptor: {
//           pattern: "shelves/{shelf}"
//           parent_type: "cloudresourcemanager.googleapis.com/Project"
//         }
//         name_descriptor: {
//           pattern: "shelves/{shelf}"
//           parent_type: "cloudresourcemanager.googleapis.com/Folder"
//         }
//       };
//     }
//
// The ResourceDescriptor Yaml config will look like:
//
//     resources:
//     - type: 'library.googleapis.com/Shelf'
//       name_descriptor:
//         - pattern: "shelves/{shelf}"
//           parent_type: "cloudresourcemanager.googleapis.com/Project"
//         - pattern: "shelves/{shelf}"
//           parent_type: "cloudresourcemanager.googleapis.com/Folder"
type ResourceDescriptor struct {
	// The resource type. It must be in the format of
	// {service_name}/{resource_type_kind}. The `resource_type_kind` must be
	// singular and must not include version numbers.
	//
	// Example: `storage.googleapis.com/Bucket`
	//
	// The value of the resource_type_kind must follow the regular expression
	// /[A-Za-z][a-zA-Z0-9]+/. It should start with an upper case character and
	// should use PascalCase (UpperCamelCase). The maximum number of
	// characters allowed for the `resource_type_kind` is 100.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. The relative resource name pattern associated with this resource
	// type. The DNS prefix of the full resource name shouldn't be specified here.
	//
	// The path pattern must follow the syntax, which aligns with HTTP binding
	// syntax:
	//
	//     Template = Segment { "/" Segment } ;
	//     Segment = LITERAL | Variable ;
	//     Variable = "{" LITERAL "}" ;
	//
	// Examples:
	//
	//     - "projects/{project}/topics/{topic}"
	//     - "projects/{project}/knowledgeBases/{knowledge_base}"
	//
	// The components in braces correspond to the IDs for each resource in the
	// hierarchy. It is expected that, if multiple patterns are provided,
	// the same component name (e.g. "project") refers to IDs of the same
	// type of resource.
	Pattern []string `protobuf:"bytes,2,rep,name=pattern,proto3" json:"pattern,omitempty"`
	// Optional. The field on the resource that designates the resource name
	// field. If omitted, this is assumed to be "name".
	NameField string `protobuf:"bytes,3,opt,name=name_field,json=nameField,proto3" json:"name_field,omitempty"`
	// Optional. The historical or future-looking state of the resource pattern.
	//
	// Example:
	//
	//     // The InspectTemplate message originally only supported resource
	//     // names with organization, and project was added later.
	//     message InspectTemplate {
	//       option (google.api.resource) = {
	//         type: "dlp.googleapis.com/InspectTemplate"
	//         pattern:
	//         "organizations/{organization}/inspectTemplates/{inspect_template}"
	//         pattern: "projects/{project}/inspectTemplates/{inspect_template}"
	//         history: ORIGINALLY_SINGLE_PATTERN
	//       };
	//     }
	History ResourceDescriptor_History `protobuf:"varint,4,opt,name=history,proto3,enum=google.api.ResourceDescriptor_History" json:"history,omitempty"`
	// The plural name used in the resource name, such as 'projects' for
	// the name of 'projects/{project}'. It is the same concept of the `plural`
	// field in k8s CRD spec
	// https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/
	Plural string `protobuf:"bytes,5,opt,name=plural,proto3" json:"plural,omitempty"`
	// The same concept of the `singular` field in k8s CRD spec
	// https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/
	// Such as "project" for the `resourcemanager.googleapis.com/Project` type.
	Singular             string   `protobuf:"bytes,6,opt,name=singular,proto3" json:"singular,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceDescriptor) Reset()         { *m = ResourceDescriptor{} }
func (m *ResourceDescriptor) String() string { return proto.CompactTextString(m) }
func (*ResourceDescriptor) ProtoMessage()    {}
func (*ResourceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_465e9122405d1bb5, []int{0}
}

func (m *ResourceDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceDescriptor.Unmarshal(m, b)
}
func (m *ResourceDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceDescriptor.Marshal(b, m, deterministic)
}
func (m *ResourceDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDescriptor.Merge(m, src)
}
func (m *ResourceDescriptor) XXX_Size() int {
	return xxx_messageInfo_ResourceDescriptor.Size(m)
}
func (m *ResourceDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDescriptor proto.InternalMessageInfo

func (m *ResourceDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ResourceDescriptor) GetPattern() []string {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *ResourceDescriptor) GetNameField() string {
	if m != nil {
		return m.NameField
	}
	return ""
}

func (m *ResourceDescriptor) GetHistory() ResourceDescriptor_History {
	if m != nil {
		return m.History
	}
	return ResourceDescriptor_HISTORY_UNSPECIFIED
}

func (m *ResourceDescriptor) GetPlural() string {
	if m != nil {
		return m.Plural
	}
	return ""
}

func (m *ResourceDescriptor) GetSingular() string {
	if m != nil {
		return m.Singular
	}
	return ""
}

// Defines a proto annotation that describes a string field that refers to
// an API resource.
type ResourceReference struct {
	// The resource type that the annotated field references.
	//
	// Example:
	//
	//     message Subscription {
	//       string topic = 2 [(google.api.resource_reference) = {
	//         type: "pubsub.googleapis.com/Topic"
	//       }];
	//     }
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The resource type of a child collection that the annotated field
	// references. This is useful for annotating the `parent` field that
	// doesn't have a fixed resource type.
	//
	// Example:
	//
	//     message ListLogEntriesRequest {
	//       string parent = 1 [(google.api.resource_reference) = {
	//         child_type: "logging.googleapis.com/LogEntry"
	//       };
	//     }
	ChildType            string   `protobuf:"bytes,2,opt,name=child_type,json=childType,proto3" json:"child_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceReference) Reset()         { *m = ResourceReference{} }
func (m *ResourceReference) String() string { return proto.CompactTextString(m) }
func (*ResourceReference) ProtoMessage()    {}
func (*ResourceReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_465e9122405d1bb5, []int{1}
}

func (m *ResourceReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceReference.Unmarshal(m, b)
}
func (m *ResourceReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceReference.Marshal(b, m, deterministic)
}
func (m *ResourceReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceReference.Merge(m, src)
}
func (m *ResourceReference) XXX_Size() int {
	return xxx_messageInfo_ResourceReference.Size(m)
}
func (m *ResourceReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceReference.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceReference proto.InternalMessageInfo

func (m *ResourceReference) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ResourceReference) GetChildType() string {
	if m != nil {
		return m.ChildType
	}
	return ""
}

var E_ResourceReference = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*ResourceReference)(nil),
	Field:         1055,
	Name:          "google.api.resource_reference",
	Tag:           "bytes,1055,opt,name=resource_reference",
	Filename:      "google/api/resource.proto",
}

var E_ResourceDefinition = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: ([]*ResourceDescriptor)(nil),
	Field:         1053,
	Name:          "google.api.resource_definition",
	Tag:           "bytes,1053,rep,name=resource_definition",
	Filename:      "google/api/resource.proto",
}

var E_Resource = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*ResourceDescriptor)(nil),
	Field:         1053,
	Name:          "google.api.resource",
	Tag:           "bytes,1053,opt,name=resource",
	Filename:      "google/api/resource.proto",
}

func init() {
	proto.RegisterEnum("google.api.ResourceDescriptor_History", ResourceDescriptor_History_name, ResourceDescriptor_History_value)
	proto.RegisterType((*ResourceDescriptor)(nil), "google.api.ResourceDescriptor")
	proto.RegisterType((*ResourceReference)(nil), "google.api.ResourceReference")
	proto.RegisterExtension(E_ResourceReference)
	proto.RegisterExtension(E_ResourceDefinition)
	proto.RegisterExtension(E_Resource)
}

func init() {
	proto.RegisterFile("google/api/resource.proto", fileDescriptor_465e9122405d1bb5)
}

var fileDescriptor_465e9122405d1bb5 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x9c, 0xe4, 0xcb, 0xcf, 0xad, 0xa8, 0xda, 0x29, 0x02, 0xb7, 0x22, 0x60, 0x65, 0x81,
	0xb2, 0xb2, 0xa5, 0xb0, 0x0b, 0x1b, 0x52, 0xe2, 0xa4, 0x96, 0xd2, 0xc4, 0x9a, 0x38, 0x8b, 0x02,
	0x92, 0x35, 0x75, 0x26, 0xee, 0x48, 0xee, 0xcc, 0x68, 0xec, 0x2c, 0xf2, 0x30, 0x08, 0x89, 0x67,
	0xe0, 0xe1, 0x58, 0xa2, 0x8c, 0x7f, 0x88, 0x68, 0x84, 0xd8, 0xcd, 0xbd, 0xe7, 0xde, 0x73, 0x8e,
	0xcf, 0x95, 0xe1, 0x32, 0x16, 0x22, 0x4e, 0xa8, 0x43, 0x24, 0x73, 0x14, 0x4d, 0xc5, 0x56, 0x45,
	0xd4, 0x96, 0x4a, 0x64, 0x02, 0x41, 0x0e, 0xd9, 0x44, 0xb2, 0x2b, 0xab, 0x18, 0xd3, 0xc8, 0xfd,
	0x76, 0xe3, 0xac, 0x69, 0x1a, 0x29, 0x26, 0x33, 0xa1, 0xf2, 0xe9, 0xde, 0x8f, 0x1a, 0x20, 0x5c,
	0x10, 0x8c, 0x2b, 0x10, 0x21, 0x68, 0x64, 0x3b, 0x49, 0x4d, 0xc3, 0x32, 0xfa, 0x1d, 0xac, 0xdf,
	0xc8, 0x84, 0x96, 0x24, 0x59, 0x46, 0x15, 0x37, 0x6b, 0x56, 0xbd, 0xdf, 0xc1, 0x65, 0x89, 0xba,
	0x00, 0x9c, 0x3c, 0xd2, 0x70, 0xc3, 0x68, 0xb2, 0x36, 0xeb, 0x7a, 0xa7, 0xb3, 0xef, 0x4c, 0xf6,
	0x0d, 0xf4, 0x01, 0x5a, 0x0f, 0x2c, 0xcd, 0x84, 0xda, 0x99, 0x0d, 0xcb, 0xe8, 0x9f, 0x0e, 0xde,
	0xda, 0xbf, 0x3d, 0xda, 0x4f, 0xd5, 0xed, 0x9b, 0x7c, 0x1a, 0x97, 0x6b, 0xe8, 0x05, 0x34, 0x65,
	0xb2, 0x55, 0x24, 0x31, 0xff, 0xd7, 0xe4, 0x45, 0x85, 0xae, 0xa0, 0x9d, 0x32, 0x1e, 0x6f, 0x13,
	0xa2, 0xcc, 0xa6, 0x46, 0xaa, 0xba, 0xf7, 0x19, 0x5a, 0x05, 0x0f, 0x7a, 0x09, 0x17, 0x37, 0xde,
	0x32, 0x58, 0xe0, 0xbb, 0x70, 0x35, 0x5f, 0xfa, 0xee, 0x47, 0x6f, 0xe2, 0xb9, 0xe3, 0xb3, 0xff,
	0x50, 0x17, 0x2e, 0x17, 0xd8, 0x9b, 0x7a, 0xf3, 0xd1, 0x6c, 0x76, 0x17, 0x2e, 0xbd, 0xf9, 0x74,
	0xe6, 0x86, 0xfe, 0x28, 0x08, 0x5c, 0x3c, 0x3f, 0x33, 0x90, 0x09, 0xcf, 0x27, 0xab, 0x60, 0x85,
	0xdd, 0xf0, 0x76, 0x35, 0x0b, 0xbc, 0x0a, 0xa9, 0xf5, 0x26, 0x70, 0x5e, 0xfa, 0xc6, 0x74, 0x43,
	0x15, 0xe5, 0x11, 0x3d, 0x1a, 0x5a, 0x17, 0x20, 0x7a, 0x60, 0xc9, 0x3a, 0xd4, 0x48, 0x2d, 0x8f,
	0x46, 0x77, 0x82, 0x9d, 0xa4, 0xc3, 0x04, 0x50, 0x79, 0xbe, 0x50, 0x55, 0x44, 0xdd, 0x32, 0x9f,
	0xf2, 0x6e, 0xb6, 0x0e, 0x72, 0x21, 0x33, 0x26, 0x78, 0x6a, 0x7e, 0x6b, 0x5b, 0x46, 0xff, 0x64,
	0xd0, 0x3d, 0x96, 0x62, 0xe5, 0x06, 0x9f, 0xab, 0x3f, 0x5b, 0x43, 0x0e, 0x17, 0x95, 0xda, 0x9a,
	0x6e, 0x18, 0x67, 0x7b, 0x42, 0xf4, 0xea, 0x88, 0x5c, 0x42, 0x4b, 0xb5, 0xaf, 0x6d, 0xab, 0xde,
	0x3f, 0x19, 0xbc, 0xfe, 0xfb, 0xcd, 0x70, 0xf5, 0x1d, 0xe3, 0x8a, 0x78, 0xf8, 0x05, 0xda, 0x65,
	0x17, 0xbd, 0x79, 0x22, 0x72, 0x4b, 0xd3, 0x94, 0xc4, 0x87, 0x3a, 0xc6, 0x3f, 0xe8, 0x54, 0x8c,
	0xd7, 0x1c, 0x4e, 0x23, 0xf1, 0x78, 0x30, 0x7e, 0xfd, 0xac, 0x9c, 0xf7, 0xf7, 0x1a, 0xbe, 0xf1,
	0x69, 0x54, 0x80, 0xb1, 0x48, 0x08, 0x8f, 0x6d, 0xa1, 0x62, 0x27, 0xa6, 0x5c, 0x3b, 0x70, 0x72,
	0x88, 0x48, 0x96, 0xea, 0xbf, 0x88, 0x70, 0x2e, 0x32, 0xa2, 0xad, 0xbc, 0x3f, 0x78, 0xff, 0x34,
	0x8c, 0xef, 0xb5, 0xc6, 0x74, 0xe4, 0x7b, 0xf7, 0x4d, 0xbd, 0xf7, 0xee, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x75, 0x12, 0x53, 0xef, 0x7c, 0x03, 0x00, 0x00,
}
