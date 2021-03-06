// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec.proto

/*
Package transformation is a generated protocol buffer package.

It is generated from these files:
	spec.proto

It has these top-level messages:
	RouteExtension
	Parameters
	TransformationSpec
*/
package transformation

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The REST Route Extension contains two components:
// * parameters for calling REST functions
// * Response Transformation
type RouteExtension struct {
	// If specified, these parameters will be used as inputs for REST templates for
	// the destination function for the route
	// (if the route destination is a functional destination that has a REST transformation)
	Parameters *Parameters `protobuf:"bytes,1,opt,name=parameters" json:"parameters,omitempty"`
	// If specified, responses on this route will be transformed according to the template(s) provided
	// in the transformation spec here
	ResponseTransformation *TransformationSpec `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation" json:"response_transformation,omitempty"`
	// If specified, paremeters for the response transformation will be extracted from these sources
	ResponseParams *Parameters `protobuf:"bytes,3,opt,name=response_params,json=responseParams" json:"response_params,omitempty"`
}

func (m *RouteExtension) Reset()                    { *m = RouteExtension{} }
func (m *RouteExtension) String() string            { return proto.CompactTextString(m) }
func (*RouteExtension) ProtoMessage()               {}
func (*RouteExtension) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{0} }

func (m *RouteExtension) GetParameters() *Parameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *RouteExtension) GetResponseTransformation() *TransformationSpec {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

func (m *RouteExtension) GetResponseParams() *Parameters {
	if m != nil {
		return m.ResponseParams
	}
	return nil
}

// Parameters define a set of parameters for REST Transformations
// Parameters can be extracted from HTTP Headers and Request Path
// Parameters can also be extracted from the HTTP Body, provided that it is
// valid JSON-encoded
// Gloo will search for parameters by their name in strings, enclosed in single
// curly braces, and attempt to match them to the variables in REST Function Templates
// for example:
//   # route
//   match: {...}
//   destination: {...}
//   extensions:
//     parameters:
//         headers:
//           x-user-id: { userId }
//   ---
//   # function
//   name: myfunc
//   spec:
//     body: |
//     {
//       "id": {{ userId }}
//     }
type Parameters struct {
	// headers that will be used to extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         headers:
	//           x-user-id: { userId }
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// part of the (or the entire) path that will be used extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         path: /users/{ userId }
	Path *google_protobuf1.StringValue `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *Parameters) Reset()                    { *m = Parameters{} }
func (m *Parameters) String() string            { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()               {}
func (*Parameters) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{1} }

func (m *Parameters) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Parameters) GetPath() *google_protobuf1.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

// TransformationSpec can act as part of a Route Extension (as a Response Transformation), or as
// a FunctionSpec (as a Request Transformation).
// Use TransformationSpec as the Function Spec for REST Services (where `Upstream.ServiceInfo.Type == "REST"`)
// TransformationSpec contains a set of templates that will be used to modify the Path, Headers, and Body
// Parameters for the tempalte come from the following sources:
// path: HTTP Request path (if present)
// method: HTTP Request method (if present)
// parameters specified in the RouteExtension.Parameters (or, in the case of ResponseTransformation, RouteExtension.ResponseParams)
// Parameters can also be extracted from the Request / Response Body provided that they are JSON
// To do so, specify the field using JSONPath syntax
// any field from the request body, assuming it's json (http://goessner.net/articles/JsonPath/index.html#e2)
type TransformationSpec struct {
	// a Jinja-style Template string for the outbound request path. Only useful for request transformation
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// a map of keys to Jinja-style Template strings HTTP Headers. Useful for request and response transformations
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// a Jinja-style Template string for the outbound HTTP Body. Useful for request and response transformations
	// If this is nil, the body will be passed through unmodified. If set to an empty string, the body will be removed
	// from the HTTP message.
	Body *google_protobuf1.StringValue `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *TransformationSpec) Reset()                    { *m = TransformationSpec{} }
func (m *TransformationSpec) String() string            { return proto.CompactTextString(m) }
func (*TransformationSpec) ProtoMessage()               {}
func (*TransformationSpec) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{2} }

func (m *TransformationSpec) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *TransformationSpec) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *TransformationSpec) GetBody() *google_protobuf1.StringValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteExtension)(nil), "gloo.api.common.v1.RouteExtension")
	proto.RegisterType((*Parameters)(nil), "gloo.api.common.v1.Parameters")
	proto.RegisterType((*TransformationSpec)(nil), "gloo.api.common.v1.TransformationSpec")
}
func (this *RouteExtension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteExtension)
	if !ok {
		that2, ok := that.(RouteExtension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !this.ResponseParams.Equal(that1.ResponseParams) {
		return false
	}
	return true
}
func (this *Parameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Parameters)
	if !ok {
		that2, ok := that.(Parameters)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if !this.Path.Equal(that1.Path) {
		return false
	}
	return true
}
func (this *TransformationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationSpec)
	if !ok {
		that2, ok := that.(TransformationSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if !this.Body.Equal(that1.Body) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("spec.proto", fileDescriptorSpec) }

var fileDescriptorSpec = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x6a, 0x1b, 0x31,
	0x10, 0xc7, 0x59, 0xbb, 0x1f, 0x58, 0x2e, 0x6e, 0x11, 0xa6, 0x35, 0xa6, 0x98, 0xe2, 0x43, 0x31,
	0x94, 0x4a, 0xad, 0x7d, 0x69, 0x7d, 0x68, 0xa1, 0x60, 0x9a, 0x4b, 0x20, 0xac, 0x43, 0x0e, 0xb9,
	0x18, 0xad, 0x2d, 0xcb, 0x8b, 0x77, 0x35, 0x42, 0xd2, 0x3a, 0xf1, 0x1b, 0xe5, 0x35, 0xf2, 0x2a,
	0x81, 0xbc, 0x44, 0x4e, 0x61, 0x25, 0xaf, 0x3f, 0x70, 0x3e, 0x0c, 0xb9, 0x8d, 0x76, 0xfe, 0xf3,
	0x9f, 0xf9, 0xcd, 0x2c, 0x42, 0x46, 0xf1, 0x31, 0x51, 0x1a, 0x2c, 0x60, 0x2c, 0x12, 0x00, 0xc2,
	0x54, 0x4c, 0xc6, 0x90, 0xa6, 0x20, 0xc9, 0xe2, 0x67, 0xb3, 0x2e, 0x40, 0x80, 0x4b, 0xd3, 0x3c,
	0xf2, 0xca, 0x66, 0x4b, 0x00, 0x88, 0x84, 0x53, 0xf7, 0x8a, 0xb2, 0x29, 0xbd, 0xd0, 0x4c, 0x29,
	0xae, 0x8d, 0xcf, 0xb7, 0xef, 0x02, 0x54, 0x0b, 0x21, 0xb3, 0x7c, 0x70, 0x69, 0xb9, 0x34, 0x31,
	0x48, 0xfc, 0x07, 0x21, 0xc5, 0x34, 0x4b, 0xb9, 0xe5, 0xda, 0x34, 0x82, 0x2f, 0x41, 0xa7, 0xda,
	0x6d, 0x91, 0xfd, 0x8e, 0xe4, 0x64, 0xad, 0x0a, 0xb7, 0x2a, 0xf0, 0x08, 0x7d, 0xd2, 0xdc, 0x28,
	0x90, 0x86, 0x8f, 0xac, 0x66, 0xd2, 0x4c, 0x41, 0xa7, 0xcc, 0xc6, 0x20, 0x1b, 0x25, 0x67, 0xf6,
	0xf5, 0x21, 0xb3, 0xd3, 0x1d, 0xe5, 0x50, 0xf1, 0x71, 0xf8, 0xb1, 0xb0, 0xd9, 0xcd, 0xe1, 0xff,
	0xe8, 0xfd, 0xba, 0x81, 0xeb, 0x6b, 0x1a, 0xe5, 0x83, 0xa6, 0xac, 0x15, 0x65, 0xee, 0x9b, 0x69,
	0x5f, 0x07, 0x08, 0x6d, 0xd2, 0x78, 0x80, 0xde, 0xce, 0x38, 0x9b, 0x78, 0xea, 0x72, 0xa7, 0xda,
	0xfd, 0xf6, 0xb4, 0x1f, 0x39, 0xf2, 0xea, 0x81, 0xb4, 0x7a, 0x19, 0x16, 0xb5, 0xf8, 0x07, 0x7a,
	0xa5, 0x98, 0x9d, 0xad, 0x60, 0x3f, 0x13, 0x7f, 0x01, 0x52, 0x5c, 0x80, 0x0c, 0xad, 0x8e, 0xa5,
	0x38, 0x63, 0x49, 0xc6, 0x43, 0xa7, 0x6c, 0xf6, 0xd1, 0xbb, 0x6d, 0x2b, 0xfc, 0x01, 0x95, 0xe7,
	0x7c, 0xe9, 0x56, 0x5f, 0x09, 0xf3, 0x10, 0xd7, 0xd1, 0xeb, 0x45, 0x5e, 0xe0, 0x4c, 0x2b, 0xa1,
	0x7f, 0xf4, 0x4b, 0xbf, 0x82, 0xf6, 0x6d, 0x80, 0xf0, 0xfe, 0xee, 0x30, 0x5e, 0x0d, 0xe1, 0x3d,
	0x5c, 0x8c, 0x8f, 0x37, 0x7c, 0x25, 0xc7, 0xd7, 0x3b, 0xec, 0x10, 0x8f, 0x73, 0x46, 0x30, 0x59,
	0xae, 0x76, 0xff, 0x0c, 0x67, 0xae, 0x7c, 0x09, 0xe7, 0xbf, 0xbf, 0x57, 0x37, 0xad, 0xe0, 0xfc,
	0xb7, 0x88, 0xed, 0x2c, 0x8b, 0xf2, 0x71, 0xa9, 0x81, 0x04, 0xbe, 0xc7, 0x40, 0x73, 0x06, 0xaa,
	0xe6, 0x82, 0xaa, 0x24, 0x13, 0xb1, 0x34, 0xd4, 0xb3, 0xd0, 0xdd, 0x7f, 0x2f, 0x7a, 0xe3, 0x06,
	0xeb, 0xdd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x5a, 0x60, 0x04, 0x48, 0x03, 0x00, 0x00,
}
