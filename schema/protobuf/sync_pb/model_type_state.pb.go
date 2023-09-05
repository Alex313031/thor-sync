// Copyright 2016 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.1
// source: model_type_state.proto

package sync_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ModelTypeState_InitialSyncState int32

const (
	// Default value, typically used when there is no metadata (e.g. because
	// Sync is disabled).
	ModelTypeState_INITIAL_SYNC_STATE_UNSPECIFIED ModelTypeState_InitialSyncState = 0
	// Indicates that syncing has started and some, but not all updates from the
	// initial download/sync have been delivered. This is only used for data
	// types in ApplyUpdatesImmediatelyTypes().
	ModelTypeState_INITIAL_SYNC_PARTIALLY_DONE ModelTypeState_InitialSyncState = 2
	// Indicates that the initial sync (download + merge) has been completed.
	ModelTypeState_INITIAL_SYNC_DONE ModelTypeState_InitialSyncState = 3
	// Indicates that no initial sync is necessary, used for CommitOnlyTypes().
	ModelTypeState_INITIAL_SYNC_UNNECESSARY ModelTypeState_InitialSyncState = 4
)

// Enum value maps for ModelTypeState_InitialSyncState.
var (
	ModelTypeState_InitialSyncState_name = map[int32]string{
		0: "INITIAL_SYNC_STATE_UNSPECIFIED",
		2: "INITIAL_SYNC_PARTIALLY_DONE",
		3: "INITIAL_SYNC_DONE",
		4: "INITIAL_SYNC_UNNECESSARY",
	}
	ModelTypeState_InitialSyncState_value = map[string]int32{
		"INITIAL_SYNC_STATE_UNSPECIFIED": 0,
		"INITIAL_SYNC_PARTIALLY_DONE":    2,
		"INITIAL_SYNC_DONE":              3,
		"INITIAL_SYNC_UNNECESSARY":       4,
	}
)

func (x ModelTypeState_InitialSyncState) Enum() *ModelTypeState_InitialSyncState {
	p := new(ModelTypeState_InitialSyncState)
	*p = x
	return p
}

func (x ModelTypeState_InitialSyncState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModelTypeState_InitialSyncState) Descriptor() protoreflect.EnumDescriptor {
	return file_model_type_state_proto_enumTypes[0].Descriptor()
}

func (ModelTypeState_InitialSyncState) Type() protoreflect.EnumType {
	return &file_model_type_state_proto_enumTypes[0]
}

func (x ModelTypeState_InitialSyncState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ModelTypeState_InitialSyncState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ModelTypeState_InitialSyncState(num)
	return nil
}

// Deprecated: Use ModelTypeState_InitialSyncState.Descriptor instead.
func (ModelTypeState_InitialSyncState) EnumDescriptor() ([]byte, []int) {
	return file_model_type_state_proto_rawDescGZIP(), []int{0, 0}
}

// Sync proto to store data type global metadata in model type storage.
type ModelTypeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The latest progress markers received from the server.
	ProgressMarker *DataTypeProgressMarker `protobuf:"bytes,1,opt,name=progress_marker,json=progressMarker" json:"progress_marker,omitempty"`
	// A data type context.  Sent to the server in every commit or update
	// request.  May be updated by either responses from the server or requests
	// made on the model thread.  The interpretation of this value may be
	// data-type specific.  Many data types ignore it.
	TypeContext *DataTypeContext `protobuf:"bytes,2,opt,name=type_context,json=typeContext" json:"type_context,omitempty"`
	// This value is set if this type's data should be encrypted on the server.
	// If this key changes, the client will need to re-commit all of its local
	// data to the server using the new encryption key.
	EncryptionKeyName *string `protobuf:"bytes,3,opt,name=encryption_key_name,json=encryptionKeyName" json:"encryption_key_name,omitempty"`
	// Deprecated in M113 and replaced by `initial_sync_state`.
	// As of M115, this is not populated anymore, but existing values are still
	// migrated to `initial_sync_state`.
	// TODO(crbug.com/1423338): Fully remove this field after a migration period.
	//
	// Deprecated: Marked as deprecated in model_type_state.proto.
	InitialSyncDoneDeprecated *bool `protobuf:"varint,4,opt,name=initial_sync_done_deprecated,json=initialSyncDoneDeprecated" json:"initial_sync_done_deprecated,omitempty"`
	// Indicates the status of "initial sync", i.e. whether the first download
	// cycle and initial merge are complete.
	InitialSyncState *ModelTypeState_InitialSyncState `protobuf:"varint,9,opt,name=initial_sync_state,json=initialSyncState,enum=sync_pb.ModelTypeState_InitialSyncState" json:"initial_sync_state,omitempty"`
	// A GUID that identifies the committing sync client. It's persisted within
	// the sync metadata and should be used to check the integrity of the
	// metadata. Mismatches with the guid of the running client indicates invalid
	// persisted sync metadata, because cache_guid is reset when sync is disabled,
	// and disabling sync is supposed to clear sync metadata.
	CacheGuid *string `protobuf:"bytes,5,opt,name=cache_guid,json=cacheGuid" json:"cache_guid,omitempty"`
	// Syncing account ID, representing the user.
	AuthenticatedAccountId *string `protobuf:"bytes,6,opt,name=authenticated_account_id,json=authenticatedAccountId" json:"authenticated_account_id,omitempty"`
	// The latest unprocessed invalidations received from the server.
	// All incoming invalidations are stored in this message and persist until
	// they are used in GetUpdate() message.
	Invalidations []*ModelTypeState_Invalidation `protobuf:"bytes,7,rep,name=invalidations" json:"invalidations,omitempty"`
	// This is relevant for the passwords datatype. This indicates that the
	// initial sync flow (downling all passwords from the server) has been run at
	// least once after the password notes features is enabled. It is used to
	// enforce redownload of passwords upon upgrading the browser to a version
	// that supports password notes if necessary. It is false by default and set
	// to true upon downloading passwords to make sure this download is executed
	// only once.
	NotesEnabledBeforeInitialSyncForPasswords *bool `protobuf:"varint,8,opt,name=notes_enabled_before_initial_sync_for_passwords,json=notesEnabledBeforeInitialSyncForPasswords" json:"notes_enabled_before_initial_sync_for_passwords,omitempty"`
}

func (x *ModelTypeState) Reset() {
	*x = ModelTypeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_type_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTypeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTypeState) ProtoMessage() {}

func (x *ModelTypeState) ProtoReflect() protoreflect.Message {
	mi := &file_model_type_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTypeState.ProtoReflect.Descriptor instead.
func (*ModelTypeState) Descriptor() ([]byte, []int) {
	return file_model_type_state_proto_rawDescGZIP(), []int{0}
}

func (x *ModelTypeState) GetProgressMarker() *DataTypeProgressMarker {
	if x != nil {
		return x.ProgressMarker
	}
	return nil
}

func (x *ModelTypeState) GetTypeContext() *DataTypeContext {
	if x != nil {
		return x.TypeContext
	}
	return nil
}

func (x *ModelTypeState) GetEncryptionKeyName() string {
	if x != nil && x.EncryptionKeyName != nil {
		return *x.EncryptionKeyName
	}
	return ""
}

// Deprecated: Marked as deprecated in model_type_state.proto.
func (x *ModelTypeState) GetInitialSyncDoneDeprecated() bool {
	if x != nil && x.InitialSyncDoneDeprecated != nil {
		return *x.InitialSyncDoneDeprecated
	}
	return false
}

func (x *ModelTypeState) GetInitialSyncState() ModelTypeState_InitialSyncState {
	if x != nil && x.InitialSyncState != nil {
		return *x.InitialSyncState
	}
	return ModelTypeState_INITIAL_SYNC_STATE_UNSPECIFIED
}

func (x *ModelTypeState) GetCacheGuid() string {
	if x != nil && x.CacheGuid != nil {
		return *x.CacheGuid
	}
	return ""
}

func (x *ModelTypeState) GetAuthenticatedAccountId() string {
	if x != nil && x.AuthenticatedAccountId != nil {
		return *x.AuthenticatedAccountId
	}
	return ""
}

func (x *ModelTypeState) GetInvalidations() []*ModelTypeState_Invalidation {
	if x != nil {
		return x.Invalidations
	}
	return nil
}

func (x *ModelTypeState) GetNotesEnabledBeforeInitialSyncForPasswords() bool {
	if x != nil && x.NotesEnabledBeforeInitialSyncForPasswords != nil {
		return *x.NotesEnabledBeforeInitialSyncForPasswords
	}
	return false
}

type ModelTypeState_Invalidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque field, which has to be provided as part of resulting GetUpdates
	// back to the server.
	Hint []byte `protobuf:"bytes,1,opt,name=hint" json:"hint,omitempty"`
	// Version of invalidation, used to order incoming invalidations.
	Version *int64 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (x *ModelTypeState_Invalidation) Reset() {
	*x = ModelTypeState_Invalidation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_type_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTypeState_Invalidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTypeState_Invalidation) ProtoMessage() {}

func (x *ModelTypeState_Invalidation) ProtoReflect() protoreflect.Message {
	mi := &file_model_type_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTypeState_Invalidation.ProtoReflect.Descriptor instead.
func (*ModelTypeState_Invalidation) Descriptor() ([]byte, []int) {
	return file_model_type_state_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ModelTypeState_Invalidation) GetHint() []byte {
	if x != nil {
		return x.Hint
	}
	return nil
}

func (x *ModelTypeState_Invalidation) GetVersion() int64 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

var File_model_type_state_proto protoreflect.FileDescriptor

var file_model_type_state_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70,
	0x62, 0x1a, 0x1f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xba, 0x06, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x3b, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x0b, 0x74, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x13,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x1c,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x64, 0x6f, 0x6e,
	0x65, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53,
	0x79, 0x6e, 0x63, 0x44, 0x6f, 0x6e, 0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x56, 0x0a, 0x12, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x47, 0x75, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x4a, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x62,
	0x0a, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x29, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x6f, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x1a, 0x3c, 0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x8c, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x79, 0x6e, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c,
	0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x49,
	0x54, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41,
	0x4c, 0x4c, 0x59, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e,
	0x49, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10,
	0x03, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x59, 0x4e,
	0x43, 0x5f, 0x55, 0x4e, 0x4e, 0x45, 0x43, 0x45, 0x53, 0x53, 0x41, 0x52, 0x59, 0x10, 0x04, 0x42,
	0x36, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01, 0x5a, 0x09, 0x2e, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
}

var (
	file_model_type_state_proto_rawDescOnce sync.Once
	file_model_type_state_proto_rawDescData = file_model_type_state_proto_rawDesc
)

func file_model_type_state_proto_rawDescGZIP() []byte {
	file_model_type_state_proto_rawDescOnce.Do(func() {
		file_model_type_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_type_state_proto_rawDescData)
	})
	return file_model_type_state_proto_rawDescData
}

var file_model_type_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_type_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_type_state_proto_goTypes = []interface{}{
	(ModelTypeState_InitialSyncState)(0), // 0: sync_pb.ModelTypeState.InitialSyncState
	(*ModelTypeState)(nil),               // 1: sync_pb.ModelTypeState
	(*ModelTypeState_Invalidation)(nil),  // 2: sync_pb.ModelTypeState.Invalidation
	(*DataTypeProgressMarker)(nil),       // 3: sync_pb.DataTypeProgressMarker
	(*DataTypeContext)(nil),              // 4: sync_pb.DataTypeContext
}
var file_model_type_state_proto_depIdxs = []int32{
	3, // 0: sync_pb.ModelTypeState.progress_marker:type_name -> sync_pb.DataTypeProgressMarker
	4, // 1: sync_pb.ModelTypeState.type_context:type_name -> sync_pb.DataTypeContext
	0, // 2: sync_pb.ModelTypeState.initial_sync_state:type_name -> sync_pb.ModelTypeState.InitialSyncState
	2, // 3: sync_pb.ModelTypeState.invalidations:type_name -> sync_pb.ModelTypeState.Invalidation
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_model_type_state_proto_init() }
func file_model_type_state_proto_init() {
	if File_model_type_state_proto != nil {
		return
	}
	file_data_type_progress_marker_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_type_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTypeState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_type_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTypeState_Invalidation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_type_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_type_state_proto_goTypes,
		DependencyIndexes: file_model_type_state_proto_depIdxs,
		EnumInfos:         file_model_type_state_proto_enumTypes,
		MessageInfos:      file_model_type_state_proto_msgTypes,
	}.Build()
	File_model_type_state_proto = out.File
	file_model_type_state_proto_rawDesc = nil
	file_model_type_state_proto_goTypes = nil
	file_model_type_state_proto_depIdxs = nil
}
