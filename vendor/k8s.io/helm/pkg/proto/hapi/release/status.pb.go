// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/release/status.proto

package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Status_Code int32

const (
	// Status_UNKNOWN indicates that a release is in an uncertain state.
	Status_UNKNOWN Status_Code = 0
	// Status_DEPLOYED indicates that the release has been pushed to Kubernetes.
	Status_DEPLOYED Status_Code = 1
	// Status_DELETED indicates that a release has been deleted from Kubermetes.
	Status_DELETED Status_Code = 2
	// Status_SUPERSEDED indicates that this release object is outdated and a newer one exists.
	Status_SUPERSEDED Status_Code = 3
	// Status_FAILED indicates that the release was not successfully deployed.
	Status_FAILED Status_Code = 4
	// Status_DELETING indicates that a delete operation is underway.
	Status_DELETING Status_Code = 5
	// Status_PENDING_INSTALL indicates that an install operation is underway.
	Status_PENDING_INSTALL Status_Code = 6
	// Status_PENDING_UPGRADE indicates that an upgrade operation is underway.
	Status_PENDING_UPGRADE Status_Code = 7
	// Status_PENDING_ROLLBACK indicates that an rollback operation is underway.
	Status_PENDING_ROLLBACK Status_Code = 8
)

var Status_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "DEPLOYED",
	2: "DELETED",
	3: "SUPERSEDED",
	4: "FAILED",
	5: "DELETING",
	6: "PENDING_INSTALL",
	7: "PENDING_UPGRADE",
	8: "PENDING_ROLLBACK",
}
var Status_Code_value = map[string]int32{
	"UNKNOWN":          0,
	"DEPLOYED":         1,
	"DELETED":          2,
	"SUPERSEDED":       3,
	"FAILED":           4,
	"DELETING":         5,
	"PENDING_INSTALL":  6,
	"PENDING_UPGRADE":  7,
	"PENDING_ROLLBACK": 8,
}

func (x Status_Code) String() string {
	return proto.EnumName(Status_Code_name, int32(x))
}
func (Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

// Status defines the status of a release.
type Status struct {
	Code Status_Code `protobuf:"varint,1,opt,name=code,enum=hapi.release.Status_Code" json:"code,omitempty"`
	// Cluster resources as kubectl would print them.
	Resources string `protobuf:"bytes,3,opt,name=resources" json:"resources,omitempty"`
	// Contains the rendered templates/NOTES.txt if available
	Notes string `protobuf:"bytes,4,opt,name=notes" json:"notes,omitempty"`
	// LastTestSuiteRun provides results on the last test run on a release
	LastTestSuiteRun *TestSuite `protobuf:"bytes,5,opt,name=last_test_suite_run,json=lastTestSuiteRun" json:"last_test_suite_run,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Status) GetCode() Status_Code {
	if m != nil {
		return m.Code
	}
	return Status_UNKNOWN
}

func (m *Status) GetResources() string {
	if m != nil {
		return m.Resources
	}
	return ""
}

func (m *Status) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Status) GetLastTestSuiteRun() *TestSuite {
	if m != nil {
		return m.LastTestSuiteRun
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "hapi.release.Status")
	proto.RegisterEnum("hapi.release.Status_Code", Status_Code_name, Status_Code_value)
}

func init() { proto.RegisterFile("hapi/release/status.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xd1, 0x6e, 0xa2, 0x40,
	0x14, 0x86, 0x17, 0x45, 0xd4, 0xa3, 0x71, 0x27, 0xa3, 0xc9, 0xa2, 0xd9, 0x4d, 0x8c, 0x57, 0xde,
	0x2c, 0x24, 0xf6, 0x09, 0xd0, 0x19, 0x0d, 0x71, 0x82, 0x04, 0x30, 0x4d, 0x7b, 0x43, 0x50, 0xa7,
	0xd6, 0xc4, 0x30, 0x86, 0x19, 0x2e, 0xfa, 0x26, 0x7d, 0xaa, 0x3e, 0x53, 0x03, 0xd8, 0xa8, 0x97,
	0xff, 0xff, 0x7d, 0x87, 0x73, 0x18, 0x18, 0xbe, 0x27, 0x97, 0x93, 0x9d, 0xf1, 0x33, 0x4f, 0x24,
	0xb7, 0xa5, 0x4a, 0x54, 0x2e, 0xad, 0x4b, 0x26, 0x94, 0xc0, 0xdd, 0x02, 0x59, 0x57, 0x34, 0xfa,
	0xf7, 0x20, 0x2a, 0x2e, 0x55, 0x2c, 0xf3, 0x93, 0xe2, 0x95, 0x3c, 0x1a, 0x1e, 0x85, 0x38, 0x9e,
	0xb9, 0x5d, 0xa6, 0x5d, 0xfe, 0x66, 0x27, 0xe9, 0x47, 0x85, 0x26, 0x5f, 0x35, 0x30, 0xc2, 0xf2,
	0xc3, 0xf8, 0x3f, 0xe8, 0x7b, 0x71, 0xe0, 0xa6, 0x36, 0xd6, 0xa6, 0xbd, 0xd9, 0xd0, 0xba, 0xdf,
	0x60, 0x55, 0x8e, 0xb5, 0x10, 0x07, 0x1e, 0x94, 0x1a, 0xfe, 0x0b, 0xed, 0x8c, 0x4b, 0x91, 0x67,
	0x7b, 0x2e, 0xcd, 0xfa, 0x58, 0x9b, 0xb6, 0x83, 0x5b, 0x81, 0x07, 0xd0, 0x48, 0x85, 0xe2, 0xd2,
	0xd4, 0x4b, 0x52, 0x05, 0xbc, 0x84, 0xfe, 0x39, 0x91, 0x2a, 0xbe, 0x5d, 0x18, 0x67, 0x79, 0x6a,
	0x36, 0xc6, 0xda, 0xb4, 0x33, 0xfb, 0xf3, 0xb8, 0x31, 0xe2, 0x52, 0x85, 0x85, 0x12, 0xa0, 0x62,
	0xe6, 0x16, 0xf3, 0x74, 0xf2, 0xa9, 0x81, 0x5e, 0x9c, 0x82, 0x3b, 0xd0, 0xdc, 0x7a, 0x6b, 0x6f,
	0xf3, 0xec, 0xa1, 0x5f, 0xb8, 0x0b, 0x2d, 0x42, 0x7d, 0xb6, 0x79, 0xa1, 0x04, 0x69, 0x05, 0x22,
	0x94, 0xd1, 0x88, 0x12, 0x54, 0xc3, 0x3d, 0x80, 0x70, 0xeb, 0xd3, 0x20, 0xa4, 0x84, 0x12, 0x54,
	0xc7, 0x00, 0xc6, 0xd2, 0x71, 0x19, 0x25, 0x48, 0xaf, 0xc6, 0x18, 0x8d, 0x5c, 0x6f, 0x85, 0x1a,
	0xb8, 0x0f, 0xbf, 0x7d, 0xea, 0x11, 0xd7, 0x5b, 0xc5, 0xae, 0x17, 0x46, 0x0e, 0x63, 0xc8, 0xb8,
	0x2f, 0xb7, 0xfe, 0x2a, 0x70, 0x08, 0x45, 0x4d, 0x3c, 0x00, 0xf4, 0x53, 0x06, 0x1b, 0xc6, 0xe6,
	0xce, 0x62, 0x8d, 0x5a, 0xf3, 0xf6, 0x6b, 0xf3, 0xfa, 0x07, 0x3b, 0xa3, 0x7c, 0xe2, 0xa7, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x48, 0x18, 0xba, 0xc7, 0x01, 0x00, 0x00,
}
