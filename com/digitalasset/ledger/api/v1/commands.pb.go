// Code generated by protoc-gen-go. DO NOT EDIT.
// source: com/digitalasset/ledger/api/v1/commands.proto

package com_digitalasset_ledger_api_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protobuf "google/protobuf"
	math "math"
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

// A composite command that groups multiple commands together.
type Commands struct {
	// Must correspond to the ledger ID reported by the Ledger Identification Service.
	// Required
	LedgerId string `protobuf:"bytes,1,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	// Identifier of the on-ledger workflow that this command is a part of.
	// Optional
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Uniquely identifies the application (or its part) that issued the command. This is used in tracing
	// across different components and to let applications subscribe to their own submissions only.
	// Required
	ApplicationId string `protobuf:"bytes,3,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// Unique command ID. This number should be unique for each new command within an application domain. It can be used for matching
	// the requests with their respective completions.
	// Required
	CommandId string `protobuf:"bytes,4,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// Party on whose behalf the command should be executed. It is up to the server to verify that the
	// authorisation can be granted and that the connection has been authenticated for that party.
	// Required
	Party string `protobuf:"bytes,5,opt,name=party,proto3" json:"party,omitempty"`
	// MUST be an approximation of the wall clock time on the ledger server.
	// Required
	LedgerEffectiveTime *protobuf.Timestamp `protobuf:"bytes,6,opt,name=ledger_effective_time,json=ledgerEffectiveTime,proto3" json:"ledger_effective_time,omitempty"`
	// The deadline for observing this command in the completion stream before it can be considered to have timed out.
	// Required
	MaximumRecordTime *protobuf.Timestamp `protobuf:"bytes,7,opt,name=maximum_record_time,json=maximumRecordTime,proto3" json:"maximum_record_time,omitempty"`
	// Individual elements of this atomic command. Must be non-empty.
	// Required
	Commands             []*Command `protobuf:"bytes,8,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Commands) Reset()         { *m = Commands{} }
func (m *Commands) String() string { return proto.CompactTextString(m) }
func (*Commands) ProtoMessage()    {}
func (*Commands) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0c48e86bcd20d63, []int{0}
}

func (m *Commands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commands.Unmarshal(m, b)
}
func (m *Commands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commands.Marshal(b, m, deterministic)
}
func (m *Commands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commands.Merge(m, src)
}
func (m *Commands) XXX_Size() int {
	return xxx_messageInfo_Commands.Size(m)
}
func (m *Commands) XXX_DiscardUnknown() {
	xxx_messageInfo_Commands.DiscardUnknown(m)
}

var xxx_messageInfo_Commands proto.InternalMessageInfo

func (m *Commands) GetLedgerId() string {
	if m != nil {
		return m.LedgerId
	}
	return ""
}

func (m *Commands) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *Commands) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *Commands) GetCommandId() string {
	if m != nil {
		return m.CommandId
	}
	return ""
}

func (m *Commands) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *Commands) GetLedgerEffectiveTime() *protobuf.Timestamp {
	if m != nil {
		return m.LedgerEffectiveTime
	}
	return nil
}

func (m *Commands) GetMaximumRecordTime() *protobuf.Timestamp {
	if m != nil {
		return m.MaximumRecordTime
	}
	return nil
}

func (m *Commands) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

// A command can either create a new contract or exercise a choice on an existing contract.
type Command struct {
	// Types that are valid to be assigned to Command:
	//	*Command_Create
	//	*Command_Exercise
	//	*Command_CreateAndExercise
	Command              isCommand_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0c48e86bcd20d63, []int{1}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

type isCommand_Command interface {
	isCommand_Command()
}

type Command_Create struct {
	Create *CreateCommand `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type Command_Exercise struct {
	Exercise *ExerciseCommand `protobuf:"bytes,2,opt,name=exercise,proto3,oneof"`
}

type Command_CreateAndExercise struct {
	CreateAndExercise *CreateAndExerciseCommand `protobuf:"bytes,3,opt,name=createAndExercise,proto3,oneof"`
}

func (*Command_Create) isCommand_Command() {}

func (*Command_Exercise) isCommand_Command() {}

func (*Command_CreateAndExercise) isCommand_Command() {}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Command) GetCreate() *CreateCommand {
	if x, ok := m.GetCommand().(*Command_Create); ok {
		return x.Create
	}
	return nil
}

func (m *Command) GetExercise() *ExerciseCommand {
	if x, ok := m.GetCommand().(*Command_Exercise); ok {
		return x.Exercise
	}
	return nil
}

func (m *Command) GetCreateAndExercise() *CreateAndExerciseCommand {
	if x, ok := m.GetCommand().(*Command_CreateAndExercise); ok {
		return x.CreateAndExercise
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Command) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Command_Create)(nil),
		(*Command_Exercise)(nil),
		(*Command_CreateAndExercise)(nil),
	}
}

// Create a new contract instance based on a template.
type CreateCommand struct {
	// The template of contract the client wants to create.
	// Required
	TemplateId *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The arguments required for creating a contract from this template.
	// Required
	CreateArguments      *Record  `protobuf:"bytes,2,opt,name=create_arguments,json=createArguments,proto3" json:"create_arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCommand) Reset()         { *m = CreateCommand{} }
func (m *CreateCommand) String() string { return proto.CompactTextString(m) }
func (*CreateCommand) ProtoMessage()    {}
func (*CreateCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0c48e86bcd20d63, []int{2}
}

func (m *CreateCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCommand.Unmarshal(m, b)
}
func (m *CreateCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCommand.Marshal(b, m, deterministic)
}
func (m *CreateCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCommand.Merge(m, src)
}
func (m *CreateCommand) XXX_Size() int {
	return xxx_messageInfo_CreateCommand.Size(m)
}
func (m *CreateCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCommand.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCommand proto.InternalMessageInfo

func (m *CreateCommand) GetTemplateId() *Identifier {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func (m *CreateCommand) GetCreateArguments() *Record {
	if m != nil {
		return m.CreateArguments
	}
	return nil
}

// Exercise a choice on an existing contract.
type ExerciseCommand struct {
	// The template of contract the client wants to exercise.
	// Required
	TemplateId *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The ID of the contract the client wants to exercise upon.
	// Required
	ContractId string `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	// The name of the choice the client wants to exercise.
	// Required
	Choice string `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	// The argument for this choice.
	// Required
	ChoiceArgument       *Value   `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExerciseCommand) Reset()         { *m = ExerciseCommand{} }
func (m *ExerciseCommand) String() string { return proto.CompactTextString(m) }
func (*ExerciseCommand) ProtoMessage()    {}
func (*ExerciseCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0c48e86bcd20d63, []int{3}
}

func (m *ExerciseCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExerciseCommand.Unmarshal(m, b)
}
func (m *ExerciseCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExerciseCommand.Marshal(b, m, deterministic)
}
func (m *ExerciseCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExerciseCommand.Merge(m, src)
}
func (m *ExerciseCommand) XXX_Size() int {
	return xxx_messageInfo_ExerciseCommand.Size(m)
}
func (m *ExerciseCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ExerciseCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ExerciseCommand proto.InternalMessageInfo

func (m *ExerciseCommand) GetTemplateId() *Identifier {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func (m *ExerciseCommand) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *ExerciseCommand) GetChoice() string {
	if m != nil {
		return m.Choice
	}
	return ""
}

func (m *ExerciseCommand) GetChoiceArgument() *Value {
	if m != nil {
		return m.ChoiceArgument
	}
	return nil
}

// Create a contract and exercise a choice on it in the same transaction.
type CreateAndExerciseCommand struct {
	// The template of the contract the client wants to create
	// Required
	TemplateId *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The arguments required for creating a contract from this template.
	// Required
	CreateArguments *Record `protobuf:"bytes,2,opt,name=create_arguments,json=createArguments,proto3" json:"create_arguments,omitempty"`
	// The name of the choice the client wants to exercise.
	// Required
	Choice string `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	// The argument for this choice.
	// Required
	ChoiceArgument       *Value   `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAndExerciseCommand) Reset()         { *m = CreateAndExerciseCommand{} }
func (m *CreateAndExerciseCommand) String() string { return proto.CompactTextString(m) }
func (*CreateAndExerciseCommand) ProtoMessage()    {}
func (*CreateAndExerciseCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0c48e86bcd20d63, []int{4}
}

func (m *CreateAndExerciseCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAndExerciseCommand.Unmarshal(m, b)
}
func (m *CreateAndExerciseCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAndExerciseCommand.Marshal(b, m, deterministic)
}
func (m *CreateAndExerciseCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAndExerciseCommand.Merge(m, src)
}
func (m *CreateAndExerciseCommand) XXX_Size() int {
	return xxx_messageInfo_CreateAndExerciseCommand.Size(m)
}
func (m *CreateAndExerciseCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAndExerciseCommand.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAndExerciseCommand proto.InternalMessageInfo

func (m *CreateAndExerciseCommand) GetTemplateId() *Identifier {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func (m *CreateAndExerciseCommand) GetCreateArguments() *Record {
	if m != nil {
		return m.CreateArguments
	}
	return nil
}

func (m *CreateAndExerciseCommand) GetChoice() string {
	if m != nil {
		return m.Choice
	}
	return ""
}

func (m *CreateAndExerciseCommand) GetChoiceArgument() *Value {
	if m != nil {
		return m.ChoiceArgument
	}
	return nil
}

func init() {
	proto.RegisterType((*Commands)(nil), "com.digitalasset.ledger.api.v1.Commands")
	proto.RegisterType((*Command)(nil), "com.digitalasset.ledger.api.v1.Command")
	proto.RegisterType((*CreateCommand)(nil), "com.digitalasset.ledger.api.v1.CreateCommand")
	proto.RegisterType((*ExerciseCommand)(nil), "com.digitalasset.ledger.api.v1.ExerciseCommand")
	proto.RegisterType((*CreateAndExerciseCommand)(nil), "com.digitalasset.ledger.api.v1.CreateAndExerciseCommand")
}

func init() {
	proto.RegisterFile("com/digitalasset/ledger/api/v1/commands.proto", fileDescriptor_d0c48e86bcd20d63)
}

var fileDescriptor_d0c48e86bcd20d63 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xd2, 0x7e, 0xcd, 0xcf, 0x58, 0x6d, 0xe8, 0x16, 0x50, 0x14, 0x04, 0x89, 0x22, 0x15,
	0xa2, 0x4a, 0xb5, 0x55, 0xc3, 0x05, 0xb7, 0x24, 0xaa, 0xc0, 0x20, 0x8a, 0xb0, 0x10, 0xb7, 0xd1,
	0x76, 0x77, 0x93, 0xae, 0xf0, 0x7a, 0xad, 0xf5, 0x26, 0x2d, 0x0f, 0xc0, 0x2b, 0xf0, 0x00, 0xbc,
	0x00, 0x8f, 0xc5, 0x6b, 0x20, 0xef, 0x4f, 0x1a, 0x40, 0x60, 0x2e, 0x10, 0xe2, 0xce, 0x3b, 0x73,
	0xce, 0xd9, 0x39, 0x33, 0xde, 0x81, 0x63, 0x22, 0x45, 0x44, 0xf9, 0x82, 0x6b, 0x9c, 0xe1, 0xb2,
	0x64, 0x3a, 0xca, 0x18, 0x5d, 0x30, 0x15, 0xe1, 0x82, 0x47, 0xab, 0x93, 0x88, 0x48, 0x21, 0x70,
	0x4e, 0xcb, 0xb0, 0x50, 0x52, 0x4b, 0x74, 0x8f, 0x48, 0x11, 0x6e, 0xc2, 0x43, 0x0b, 0x0f, 0x71,
	0xc1, 0xc3, 0xd5, 0x49, 0xff, 0xa8, 0x46, 0x6e, 0x85, 0xb3, 0x25, 0xb3, 0x5a, 0xfd, 0xc1, 0x42,
	0xca, 0x45, 0xc6, 0x22, 0x73, 0x3a, 0x5f, 0xce, 0x23, 0xcd, 0x05, 0x2b, 0x35, 0x16, 0x85, 0x05,
	0x8c, 0x3e, 0x6c, 0x43, 0x7b, 0xea, 0xee, 0x47, 0x77, 0xa0, 0x63, 0xa5, 0x66, 0x9c, 0xf6, 0x1a,
	0xc3, 0xc6, 0xb8, 0x93, 0xb6, 0x6d, 0x20, 0xa1, 0x68, 0x00, 0xc1, 0xa5, 0x54, 0xef, 0xe6, 0x99,
	0xbc, 0xac, 0xd2, 0x5b, 0x26, 0x0d, 0x3e, 0x94, 0x50, 0x74, 0x08, 0x7b, 0xb8, 0x28, 0x32, 0x4e,
	0xb0, 0xe6, 0x32, 0xaf, 0x30, 0xdb, 0x06, 0xb3, 0xbb, 0x11, 0x4d, 0x28, 0xba, 0x0b, 0xe0, 0x0c,
	0x57, 0x90, 0xff, 0x0d, 0xa4, 0xe3, 0x22, 0x09, 0x45, 0x37, 0x61, 0xa7, 0xc0, 0x4a, 0xbf, 0xef,
	0xed, 0x98, 0x8c, 0x3d, 0xa0, 0x33, 0xb8, 0xe5, 0x2a, 0x63, 0xf3, 0x39, 0x23, 0x9a, 0xaf, 0xd8,
	0xac, 0xb2, 0xd2, 0x6b, 0x0e, 0x1b, 0xe3, 0x20, 0xee, 0x87, 0xd6, 0x67, 0xe8, 0x7d, 0x86, 0x6f,
	0xbc, 0xcf, 0xf4, 0xc0, 0x12, 0x4f, 0x3d, 0xaf, 0xca, 0xa0, 0xe7, 0x70, 0x20, 0xf0, 0x15, 0x17,
	0x4b, 0x31, 0x53, 0x8c, 0x48, 0x45, 0xad, 0x5a, 0xab, 0x56, 0x6d, 0xdf, 0xd1, 0x52, 0xc3, 0x32,
	0x5a, 0x53, 0x68, 0xfb, 0x09, 0xf6, 0xda, 0xc3, 0xed, 0x71, 0x10, 0x3f, 0x08, 0x7f, 0x3d, 0xc2,
	0xd0, 0x75, 0x3c, 0x5d, 0x13, 0x47, 0x1f, 0xb7, 0xa0, 0xe5, 0xa2, 0xe8, 0x29, 0x34, 0x89, 0x62,
	0x58, 0x33, 0x33, 0x83, 0x20, 0x3e, 0xae, 0x95, 0x33, 0x68, 0x47, 0x7f, 0xf6, 0x5f, 0xea, 0xe8,
	0xe8, 0x25, 0xb4, 0xd9, 0x15, 0x53, 0x84, 0x97, 0xcc, 0xcc, 0x2b, 0x88, 0xa3, 0x3a, 0xa9, 0x53,
	0x87, 0xbf, 0x16, 0x5b, 0x4b, 0xa0, 0x0b, 0xd8, 0xb7, 0xc2, 0x4f, 0x72, 0xea, 0x71, 0x66, 0xc6,
	0x41, 0xfc, 0xf8, 0xf7, 0x4a, 0xdc, 0x20, 0x5e, 0x5f, 0xf0, 0xa3, 0xe8, 0xa4, 0x03, 0x2d, 0xd7,
	0x99, 0xd1, 0xe7, 0x06, 0xec, 0x7e, 0xe3, 0x0f, 0xbd, 0x80, 0x40, 0x33, 0x51, 0x64, 0x58, 0x33,
	0xff, 0x9f, 0x06, 0xf1, 0x51, 0x5d, 0x01, 0x09, 0x65, 0xb9, 0xe6, 0x73, 0xce, 0x54, 0x0a, 0x9e,
	0x9e, 0x50, 0xf4, 0x1a, 0x6e, 0xd8, 0xeb, 0x67, 0x58, 0x2d, 0x96, 0x82, 0xe5, 0xba, 0x74, 0xad,
	0xba, 0x5f, 0xa7, 0x68, 0x7f, 0x81, 0xb4, 0xeb, 0xca, 0xf7, 0xf4, 0xd1, 0x97, 0x06, 0x74, 0xbf,
	0x73, 0xf9, 0x67, 0x6b, 0x1e, 0x40, 0x40, 0x64, 0xae, 0x15, 0x26, 0x7a, 0xe3, 0x25, 0xfa, 0x50,
	0x42, 0xd1, 0x6d, 0x68, 0x92, 0x0b, 0xc9, 0x09, 0x73, 0x2f, 0xd0, 0x9d, 0xd0, 0x19, 0x74, 0xed,
	0xd7, 0xda, 0xac, 0x79, 0x7f, 0x41, 0x7c, 0x58, 0x57, 0xc9, 0xdb, 0x6a, 0xa7, 0xa4, 0x7b, 0x96,
	0xed, 0xad, 0x8e, 0x3e, 0x6d, 0x41, 0xef, 0x67, 0x83, 0xfd, 0xd7, 0xc7, 0xf4, 0xb7, 0x9a, 0x34,
	0x79, 0x04, 0x35, 0x0b, 0x7d, 0x82, 0xfc, 0x02, 0x7e, 0xb5, 0xd4, 0x4c, 0x4d, 0x2b, 0xcc, 0x79,
	0xd3, 0xec, 0x9e, 0x87, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x20, 0xbb, 0x3d, 0x3c, 0x06,
	0x00, 0x00,
}