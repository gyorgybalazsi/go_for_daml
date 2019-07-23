// Code generated by protoc-gen-go. DO NOT EDIT.
// source: com/digitalasset/ledger/api/v1/event.proto

package com_digitalasset_ledger_api_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// An event on the ledger can either be the creation or the archiving of
// a contract, or the exercise of a choice on a contract.
//
// The ``GetTransactionTrees`` response will only contain create and
// exercise events. Archive events correspond to consuming exercise
// events.
//
// In the transaction service the events are restricted to the events
// visible for the parties specified in the transaction filter. Each
// event message type below contains a ``witness_parties`` field which
// indicates the subset of the requested parties that can see the event
// in question.
//
// However, note that transaction _trees_ might contain events with
// _no_ witness parties, which were included simply because they were
// children of events which have witnesses. On the other hand, on
// the flat transaction stream you'll only receive events that have
// witnesses.
type Event struct {
	// Types that are valid to be assigned to Event:
	//	*Event_Created
	//	*Event_Exercised
	//	*Event_Archived
	Event                isEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27c3e7d1bcb7f68, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Created struct {
	Created *CreatedEvent `protobuf:"bytes,1,opt,name=created,proto3,oneof"`
}

type Event_Exercised struct {
	Exercised *ExercisedEvent `protobuf:"bytes,2,opt,name=exercised,proto3,oneof"`
}

type Event_Archived struct {
	Archived *ArchivedEvent `protobuf:"bytes,3,opt,name=archived,proto3,oneof"`
}

func (*Event_Created) isEvent_Event() {}

func (*Event_Exercised) isEvent_Event() {}

func (*Event_Archived) isEvent_Event() {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetCreated() *CreatedEvent {
	if x, ok := m.GetEvent().(*Event_Created); ok {
		return x.Created
	}
	return nil
}

func (m *Event) GetExercised() *ExercisedEvent {
	if x, ok := m.GetEvent().(*Event_Exercised); ok {
		return x.Exercised
	}
	return nil
}

func (m *Event) GetArchived() *ArchivedEvent {
	if x, ok := m.GetEvent().(*Event_Archived); ok {
		return x.Archived
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_Created)(nil),
		(*Event_Exercised)(nil),
		(*Event_Archived)(nil),
	}
}

// Records that a contract has been created, and choices may now be exercised on it.
type CreatedEvent struct {
	// The ID of this particular event.
	// Required
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// The ID of the created contract.
	// Required
	ContractId string `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	// The template of the created contract.
	// Required
	TemplateId *Identifier `protobuf:"bytes,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The arguments that have been used to create the contract.
	// Required
	CreateArguments *Record `protobuf:"bytes,4,opt,name=create_arguments,json=createArguments,proto3" json:"create_arguments,omitempty"`
	// The parties that are notified of this event. For `CreatedEvent`s,
	// these are the intersection of the stakeholders of the contract in
	// question and the parties specified in the `TransactionFilter`. The
	// stakeholders are the union of the signatories and the observers of
	// the contract.
	// Required
	WitnessParties       []string `protobuf:"bytes,5,rep,name=witness_parties,json=witnessParties,proto3" json:"witness_parties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatedEvent) Reset()         { *m = CreatedEvent{} }
func (m *CreatedEvent) String() string { return proto.CompactTextString(m) }
func (*CreatedEvent) ProtoMessage()    {}
func (*CreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27c3e7d1bcb7f68, []int{1}
}

func (m *CreatedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatedEvent.Unmarshal(m, b)
}
func (m *CreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatedEvent.Marshal(b, m, deterministic)
}
func (m *CreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatedEvent.Merge(m, src)
}
func (m *CreatedEvent) XXX_Size() int {
	return xxx_messageInfo_CreatedEvent.Size(m)
}
func (m *CreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CreatedEvent proto.InternalMessageInfo

func (m *CreatedEvent) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *CreatedEvent) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *CreatedEvent) GetTemplateId() *Identifier {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func (m *CreatedEvent) GetCreateArguments() *Record {
	if m != nil {
		return m.CreateArguments
	}
	return nil
}

func (m *CreatedEvent) GetWitnessParties() []string {
	if m != nil {
		return m.WitnessParties
	}
	return nil
}

// Records that a contract has been archived, and choices may no longer be exercised on it.
type ArchivedEvent struct {
	// The ID of this particular event.
	// Required
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// The ID of the archived contract.
	// Required
	ContractId string `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	// The template of the archived contract.
	// Required
	TemplateId *Identifier `protobuf:"bytes,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The parties that are notified of this event. For `ArchivedEvent`s,
	// these are the intersection of the stakeholders of the contract in
	// question and the parties specified in the `TransactionFilter`. The
	// stakeholders are the union of the signatories and the observers of
	// the contract.
	// Required
	WitnessParties       []string `protobuf:"bytes,4,rep,name=witness_parties,json=witnessParties,proto3" json:"witness_parties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchivedEvent) Reset()         { *m = ArchivedEvent{} }
func (m *ArchivedEvent) String() string { return proto.CompactTextString(m) }
func (*ArchivedEvent) ProtoMessage()    {}
func (*ArchivedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27c3e7d1bcb7f68, []int{2}
}

func (m *ArchivedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchivedEvent.Unmarshal(m, b)
}
func (m *ArchivedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchivedEvent.Marshal(b, m, deterministic)
}
func (m *ArchivedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchivedEvent.Merge(m, src)
}
func (m *ArchivedEvent) XXX_Size() int {
	return xxx_messageInfo_ArchivedEvent.Size(m)
}
func (m *ArchivedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchivedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ArchivedEvent proto.InternalMessageInfo

func (m *ArchivedEvent) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *ArchivedEvent) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *ArchivedEvent) GetTemplateId() *Identifier {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func (m *ArchivedEvent) GetWitnessParties() []string {
	if m != nil {
		return m.WitnessParties
	}
	return nil
}

// Records that a choice has been exercised on a target contract.
type ExercisedEvent struct {
	// The ID of this particular event.
	// Required
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// The ID of the target contract.
	// Required
	ContractId string `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	// The template of the target contract.
	// Required
	TemplateId *Identifier `protobuf:"bytes,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The ID of the event in which the target contract has been created.
	// Required
	ContractCreatingEventId string `protobuf:"bytes,4,opt,name=contract_creating_event_id,json=contractCreatingEventId,proto3" json:"contract_creating_event_id,omitempty"`
	// The choice that's been exercised on the target contract.
	// Required
	Choice string `protobuf:"bytes,5,opt,name=choice,proto3" json:"choice,omitempty"`
	// The argument the choice was made with.
	// Required
	ChoiceArgument *Value `protobuf:"bytes,6,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
	// The parties that made the choice.
	// Required
	ActingParties []string `protobuf:"bytes,7,rep,name=acting_parties,json=actingParties,proto3" json:"acting_parties,omitempty"`
	// If true, the target contract may no longer be exercised.
	// Required
	Consuming bool `protobuf:"varint,8,opt,name=consuming,proto3" json:"consuming,omitempty"`
	// The parties that are notified of this event. The witnesses of an exercise
	// node will depend on whether the exercise was consuming or not.
	//
	// If consuming, the witnesses are the union of the stakeholders and
	// the actors.
	//
	// If not consuming, the witnesses are the union of the signatories and
	// the actors. Note that the actors might not necessarily be observers
	// and thus signatories. This is the case when the controllers of a
	// choice are specified using "flexible controllers", using the
	// `choice ... controller` syntax, and said controllers are not
	// explicitly marked as observers.
	//
	// Required
	WitnessParties []string `protobuf:"bytes,10,rep,name=witness_parties,json=witnessParties,proto3" json:"witness_parties,omitempty"`
	// References to further events in the same transaction that appeared as a result of this ``ExercisedEvent``.
	// It contains only the immediate children of this event, not all members of the subtree rooted at this node.
	// Optional
	ChildEventIds []string `protobuf:"bytes,11,rep,name=child_event_ids,json=childEventIds,proto3" json:"child_event_ids,omitempty"`
	// The result of exercising the choice
	// Required
	ExerciseResult       *Value   `protobuf:"bytes,12,opt,name=exercise_result,json=exerciseResult,proto3" json:"exercise_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExercisedEvent) Reset()         { *m = ExercisedEvent{} }
func (m *ExercisedEvent) String() string { return proto.CompactTextString(m) }
func (*ExercisedEvent) ProtoMessage()    {}
func (*ExercisedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a27c3e7d1bcb7f68, []int{3}
}

func (m *ExercisedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExercisedEvent.Unmarshal(m, b)
}
func (m *ExercisedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExercisedEvent.Marshal(b, m, deterministic)
}
func (m *ExercisedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExercisedEvent.Merge(m, src)
}
func (m *ExercisedEvent) XXX_Size() int {
	return xxx_messageInfo_ExercisedEvent.Size(m)
}
func (m *ExercisedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ExercisedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ExercisedEvent proto.InternalMessageInfo

func (m *ExercisedEvent) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *ExercisedEvent) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *ExercisedEvent) GetTemplateId() *Identifier {
	if m != nil {
		return m.TemplateId
	}
	return nil
}

func (m *ExercisedEvent) GetContractCreatingEventId() string {
	if m != nil {
		return m.ContractCreatingEventId
	}
	return ""
}

func (m *ExercisedEvent) GetChoice() string {
	if m != nil {
		return m.Choice
	}
	return ""
}

func (m *ExercisedEvent) GetChoiceArgument() *Value {
	if m != nil {
		return m.ChoiceArgument
	}
	return nil
}

func (m *ExercisedEvent) GetActingParties() []string {
	if m != nil {
		return m.ActingParties
	}
	return nil
}

func (m *ExercisedEvent) GetConsuming() bool {
	if m != nil {
		return m.Consuming
	}
	return false
}

func (m *ExercisedEvent) GetWitnessParties() []string {
	if m != nil {
		return m.WitnessParties
	}
	return nil
}

func (m *ExercisedEvent) GetChildEventIds() []string {
	if m != nil {
		return m.ChildEventIds
	}
	return nil
}

func (m *ExercisedEvent) GetExerciseResult() *Value {
	if m != nil {
		return m.ExerciseResult
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "com.digitalasset.ledger.api.v1.Event")
	proto.RegisterType((*CreatedEvent)(nil), "com.digitalasset.ledger.api.v1.CreatedEvent")
	proto.RegisterType((*ArchivedEvent)(nil), "com.digitalasset.ledger.api.v1.ArchivedEvent")
	proto.RegisterType((*ExercisedEvent)(nil), "com.digitalasset.ledger.api.v1.ExercisedEvent")
}

func init() {
	proto.RegisterFile("com/digitalasset/ledger/api/v1/event.proto", fileDescriptor_a27c3e7d1bcb7f68)
}

var fileDescriptor_a27c3e7d1bcb7f68 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x25, 0x6d, 0x3e, 0x27, 0x6d, 0x52, 0xf9, 0x00, 0xa6, 0x42, 0x10, 0x45, 0x6a, 0x89, 0x2a,
	0xd8, 0x28, 0x70, 0xe4, 0xd4, 0x44, 0x91, 0x1a, 0x90, 0x0a, 0xec, 0x81, 0x6b, 0x64, 0xec, 0x21,
	0xb5, 0xb4, 0x5f, 0xb2, 0xbd, 0x0b, 0xff, 0x83, 0x7f, 0xc1, 0xbf, 0xe0, 0x57, 0x71, 0x45, 0x6b,
	0xaf, 0xb7, 0xad, 0x14, 0xb1, 0xe2, 0xd6, 0xdb, 0xee, 0xf8, 0xbd, 0x37, 0x6f, 0x9e, 0xad, 0x81,
	0x0b, 0x9e, 0xc6, 0x73, 0x21, 0x77, 0xd2, 0xb0, 0x88, 0x69, 0x8d, 0x66, 0x1e, 0xa1, 0xd8, 0xa1,
	0x9a, 0xb3, 0x4c, 0xce, 0x8b, 0xc5, 0x1c, 0x0b, 0x4c, 0x4c, 0x90, 0xa9, 0xd4, 0xa4, 0xe4, 0x39,
	0x4f, 0xe3, 0xe0, 0x2e, 0x36, 0x70, 0xd8, 0x80, 0x65, 0x32, 0x28, 0x16, 0xa7, 0x4d, 0x5a, 0x05,
	0x8b, 0x72, 0x74, 0x5a, 0xd3, 0x3f, 0x2d, 0xe8, 0xac, 0x4b, 0x6d, 0x72, 0x05, 0x3d, 0xae, 0x90,
	0x19, 0x14, 0xb4, 0x35, 0x69, 0xcd, 0x86, 0x6f, 0x5e, 0x05, 0xff, 0xee, 0x13, 0xac, 0x1c, 0xdc,
	0xd2, 0xaf, 0x1e, 0x85, 0x9e, 0x4e, 0xae, 0x61, 0x80, 0x3f, 0x50, 0x71, 0xa9, 0x51, 0xd0, 0x03,
	0xab, 0x15, 0x34, 0x69, 0xad, 0x3d, 0xc1, 0xab, 0xdd, 0x4a, 0x90, 0x0f, 0xd0, 0x67, 0x8a, 0xdf,
	0xc8, 0x02, 0x05, 0x3d, 0xb4, 0x72, 0xaf, 0x9b, 0xe4, 0x2e, 0x2b, 0xbc, 0x57, 0xab, 0x05, 0x96,
	0x3d, 0xe8, 0xd8, 0x2c, 0xa7, 0x3f, 0x0f, 0xe0, 0xe8, 0xee, 0x04, 0xe4, 0x29, 0xf4, 0xed, 0xc9,
	0x56, 0xba, 0x04, 0x06, 0x61, 0xcf, 0xfe, 0x6f, 0x04, 0x79, 0x01, 0x43, 0x9e, 0x26, 0x46, 0x31,
	0x6e, 0x4f, 0x0f, 0xec, 0x29, 0xf8, 0xd2, 0xa6, 0xb4, 0x38, 0x34, 0x18, 0x67, 0x11, 0x33, 0x58,
	0x02, 0x9c, 0xcb, 0x8b, 0x26, 0x97, 0x1b, 0x81, 0x89, 0x91, 0xdf, 0x24, 0xaa, 0x10, 0x3c, 0x7d,
	0x23, 0xc8, 0x67, 0x38, 0x71, 0x51, 0x6e, 0x99, 0xda, 0xe5, 0x31, 0x26, 0x46, 0xd3, 0xb6, 0x55,
	0x3c, 0x6f, 0x52, 0x0c, 0x91, 0xa7, 0x4a, 0x84, 0x63, 0xc7, 0xbf, 0xf4, 0x74, 0xf2, 0x12, 0xc6,
	0xdf, 0xa5, 0x49, 0x50, 0xeb, 0x6d, 0xc6, 0x94, 0x91, 0xa8, 0x69, 0x67, 0x72, 0x38, 0x1b, 0x84,
	0xa3, 0xaa, 0xfc, 0xc9, 0x55, 0xa7, 0xbf, 0x5b, 0x70, 0x7c, 0x2f, 0xbc, 0x87, 0x13, 0xcb, 0x9e,
	0x19, 0xda, 0x7b, 0x67, 0xf8, 0xd5, 0x86, 0xd1, 0xfd, 0xf7, 0xf4, 0x70, 0x86, 0x78, 0x07, 0xa7,
	0x75, 0x37, 0x7b, 0x49, 0x32, 0xd9, 0x6d, 0x6b, 0x6b, 0x6d, 0xdb, 0xfc, 0x89, 0x47, 0xac, 0x2a,
	0xc0, 0xba, 0xb2, 0xfa, 0x18, 0xba, 0xfc, 0x26, 0x95, 0x1c, 0x69, 0xc7, 0x02, 0xab, 0x3f, 0x72,
	0x0d, 0x63, 0xf7, 0x55, 0x3f, 0x18, 0xda, 0xb5, 0x2e, 0xcf, 0x9a, 0x5c, 0x7e, 0x29, 0x57, 0x41,
	0x38, 0x72, 0x6c, 0xff, 0x5c, 0xc8, 0x19, 0x8c, 0x18, 0xb7, 0xce, 0x7c, 0xd0, 0x3d, 0x1b, 0xf4,
	0xb1, 0xab, 0x56, 0x39, 0x93, 0x67, 0x30, 0xe0, 0x69, 0xa2, 0xf3, 0x58, 0x26, 0x3b, 0xda, 0x9f,
	0xb4, 0x66, 0xfd, 0xf0, 0xb6, 0xb0, 0xef, 0xba, 0x60, 0xdf, 0x75, 0x91, 0xf3, 0xd2, 0xbd, 0x8c,
	0x44, 0x1d, 0x83, 0xa6, 0x43, 0xd7, 0xce, 0x96, 0xab, 0xe1, 0x75, 0x39, 0xa5, 0xdf, 0x09, 0x5b,
	0x85, 0x3a, 0x8f, 0x0c, 0x3d, 0xfa, 0xaf, 0x29, 0x3d, 0x3b, 0xb4, 0xe4, 0xf7, 0xed, 0xfe, 0xe0,
	0x04, 0x96, 0x0b, 0x68, 0x58, 0xa7, 0xcb, 0xb1, 0x75, 0xf0, 0x31, 0x37, 0xa8, 0x56, 0x25, 0xe0,
	0x6b, 0xd7, 0xae, 0xce, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xda, 0xb6, 0x96, 0x43, 0xb4,
	0x05, 0x00, 0x00,
}
