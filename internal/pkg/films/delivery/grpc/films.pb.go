// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: films.proto

package grpc

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

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{0}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{1}
}

type KeyWord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word string `protobuf:"bytes,1,opt,name=Word,proto3" json:"Word,omitempty"`
}

func (x *KeyWord) Reset() {
	*x = KeyWord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyWord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyWord) ProtoMessage() {}

func (x *KeyWord) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyWord.ProtoReflect.Descriptor instead.
func (*KeyWord) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{2}
}

func (x *KeyWord) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type Season struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Season) Reset() {
	*x = Season{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Season) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Season) ProtoMessage() {}

func (x *Season) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Season.ProtoReflect.Descriptor instead.
func (*Season) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{3}
}

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string     `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title              string     `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Genres             []string   `protobuf:"bytes,3,rep,name=Genres,proto3" json:"Genres,omitempty"`
	Country            string     `protobuf:"bytes,4,opt,name=Country,proto3" json:"Country,omitempty"`
	Year               int64      `protobuf:"varint,5,opt,name=Year,proto3" json:"Year,omitempty"`
	ReleaseRus         *Timestamp `protobuf:"bytes,6,opt,name=ReleaseRus,proto3" json:"ReleaseRus,omitempty"`
	Director           []string   `protobuf:"bytes,7,rep,name=Director,proto3" json:"Director,omitempty"`
	Authors            []string   `protobuf:"bytes,8,rep,name=Authors,proto3" json:"Authors,omitempty"`
	Actors             []string   `protobuf:"bytes,9,rep,name=Actors,proto3" json:"Actors,omitempty"`
	Release            *Timestamp `protobuf:"bytes,10,opt,name=Release,proto3" json:"Release,omitempty"`
	Duration           int64      `protobuf:"varint,11,opt,name=Duration,proto3" json:"Duration,omitempty"`
	ReleaseRusLanguage string     `protobuf:"bytes,12,opt,name=ReleaseRusLanguage,proto3" json:"ReleaseRusLanguage,omitempty"`
	Budget             string     `protobuf:"bytes,13,opt,name=Budget,proto3" json:"Budget,omitempty"`
	Age                int64      `protobuf:"varint,14,opt,name=Age,proto3" json:"Age,omitempty"`
	Pic                []string   `protobuf:"bytes,15,rep,name=Pic,proto3" json:"Pic,omitempty"`
	Src                []string   `protobuf:"bytes,16,rep,name=Src,proto3" json:"Src,omitempty"`
	Description        string     `protobuf:"bytes,17,opt,name=Description,proto3" json:"Description,omitempty"`
	IsSeries           bool       `protobuf:"varint,18,opt,name=IsSeries,proto3" json:"IsSeries,omitempty"`
	Seasons            []*Season  `protobuf:"bytes,19,rep,name=Seasons,proto3" json:"Seasons,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{4}
}

func (x *Film) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Film) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Film) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Film) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Film) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Film) GetReleaseRus() *Timestamp {
	if x != nil {
		return x.ReleaseRus
	}
	return nil
}

func (x *Film) GetDirector() []string {
	if x != nil {
		return x.Director
	}
	return nil
}

func (x *Film) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Film) GetActors() []string {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *Film) GetRelease() *Timestamp {
	if x != nil {
		return x.Release
	}
	return nil
}

func (x *Film) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Film) GetReleaseRusLanguage() string {
	if x != nil {
		return x.ReleaseRusLanguage
	}
	return ""
}

func (x *Film) GetBudget() string {
	if x != nil {
		return x.Budget
	}
	return ""
}

func (x *Film) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Film) GetPic() []string {
	if x != nil {
		return x.Pic
	}
	return nil
}

func (x *Film) GetSrc() []string {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *Film) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Film) GetIsSeries() bool {
	if x != nil {
		return x.IsSeries
	}
	return false
}

func (x *Film) GetSeasons() []*Season {
	if x != nil {
		return x.Seasons
	}
	return nil
}

type Films struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Film `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Films) Reset() {
	*x = Films{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Films) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Films) ProtoMessage() {}

func (x *Films) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Films.ProtoReflect.Descriptor instead.
func (*Films) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{5}
}

func (x *Films) GetData() []*Film {
	if x != nil {
		return x.Data
	}
	return nil
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmUUID string `protobuf:"bytes,1,opt,name=FilmUUID,proto3" json:"FilmUUID,omitempty"`
	UserUUID string `protobuf:"bytes,2,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{6}
}

func (x *Pair) GetFilmUUID() string {
	if x != nil {
		return x.FilmUUID
	}
	return ""
}

func (x *Pair) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{7}
}

func (x *UUID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_films_proto protoreflect.FileDescriptor

var file_films_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x1d, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x57, 0x6f, 0x72, 0x64, 0x22, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x8d,
	0x04, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x2a, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x24, 0x0a,
	0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x75, 0x73, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x52, 0x75, 0x73, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x63,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x50, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x72, 0x63, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x53, 0x72, 0x63, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x73, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x49, 0x73, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x53,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x22,
	0x0a, 0x05, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x3e, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x6d, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69,
	0x6c, 0x6d, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc1, 0x03, 0x0a, 0x0c, 0x46,
	0x69, 0x6c, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x46,
	0x69, 0x6c, 0x6d, 0x42, 0x79, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x08, 0x2e, 0x4b, 0x65, 0x79,
	0x57, 0x6f, 0x72, 0x64, 0x1a, 0x06, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x25,
	0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x6d, 0x42, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x1a, 0x06, 0x2e, 0x46, 0x69,
	0x6c, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x42, 0x79,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x46,
	0x69, 0x6c, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x6d, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x6d,
	0x22, 0x00, 0x12, 0x1e, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x73,
	0x22, 0x00, 0x12, 0x25, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a,
	0x06, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x53, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x12, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x08,
	0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0d, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x12, 0x05, 0x2e, 0x50, 0x61,
	0x69, 0x72, 0x1a, 0x08, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x21,
	0x0a, 0x0c, 0x41, 0x64, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x05,
	0x2e, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x08, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22,
	0x00, 0x12, 0x24, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x08, 0x2e, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x72,
	0x65, 0x64, 0x12, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x46, 0x69, 0x6c, 0x6d,
	0x73, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x22,
	0x00, 0x12, 0x1b, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x08, 0x2e, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x22, 0x00, 0x42, 0x22,
	0x5a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_films_proto_rawDescOnce sync.Once
	file_films_proto_rawDescData = file_films_proto_rawDesc
)

func file_films_proto_rawDescGZIP() []byte {
	file_films_proto_rawDescOnce.Do(func() {
		file_films_proto_rawDescData = protoimpl.X.CompressGZIP(file_films_proto_rawDescData)
	})
	return file_films_proto_rawDescData
}

var file_films_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_films_proto_goTypes = []interface{}{
	(*Timestamp)(nil), // 0: Timestamp
	(*Nothing)(nil),   // 1: Nothing
	(*KeyWord)(nil),   // 2: KeyWord
	(*Season)(nil),    // 3: Season
	(*Film)(nil),      // 4: Film
	(*Films)(nil),     // 5: Films
	(*Pair)(nil),      // 6: Pair
	(*UUID)(nil),      // 7: UUID
}
var file_films_proto_depIdxs = []int32{
	0,  // 0: Film.ReleaseRus:type_name -> Timestamp
	0,  // 1: Film.Release:type_name -> Timestamp
	3,  // 2: Film.Seasons:type_name -> Season
	4,  // 3: Films.Data:type_name -> Film
	2,  // 4: FilmsService.FilmByGenre:input_type -> KeyWord
	2,  // 5: FilmsService.FilmBySelection:input_type -> KeyWord
	7,  // 6: FilmsService.FilmsByActor:input_type -> UUID
	7,  // 7: FilmsService.FilmById:input_type -> UUID
	7,  // 8: FilmsService.FilmsByUser:input_type -> UUID
	7,  // 9: FilmsService.FilmStartSelection:input_type -> UUID
	6,  // 10: FilmsService.AddStarred:input_type -> Pair
	6,  // 11: FilmsService.RemoveStarred:input_type -> Pair
	6,  // 12: FilmsService.AddWatchList:input_type -> Pair
	6,  // 13: FilmsService.RemoveWatchList:input_type -> Pair
	7,  // 14: FilmsService.Starred:input_type -> UUID
	7,  // 15: FilmsService.WatchList:input_type -> UUID
	1,  // 16: FilmsService.Random:input_type -> Nothing
	5,  // 17: FilmsService.FilmByGenre:output_type -> Films
	5,  // 18: FilmsService.FilmBySelection:output_type -> Films
	5,  // 19: FilmsService.FilmsByActor:output_type -> Films
	4,  // 20: FilmsService.FilmById:output_type -> Film
	5,  // 21: FilmsService.FilmsByUser:output_type -> Films
	5,  // 22: FilmsService.FilmStartSelection:output_type -> Films
	1,  // 23: FilmsService.AddStarred:output_type -> Nothing
	1,  // 24: FilmsService.RemoveStarred:output_type -> Nothing
	1,  // 25: FilmsService.AddWatchList:output_type -> Nothing
	1,  // 26: FilmsService.RemoveWatchList:output_type -> Nothing
	5,  // 27: FilmsService.Starred:output_type -> Films
	5,  // 28: FilmsService.WatchList:output_type -> Films
	4,  // 29: FilmsService.Random:output_type -> Film
	17, // [17:30] is the sub-list for method output_type
	4,  // [4:17] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_films_proto_init() }
func file_films_proto_init() {
	if File_films_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_films_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
		file_films_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
		file_films_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyWord); i {
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
		file_films_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Season); i {
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
		file_films_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Film); i {
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
		file_films_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Films); i {
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
		file_films_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_films_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
			RawDescriptor: file_films_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_films_proto_goTypes,
		DependencyIndexes: file_films_proto_depIdxs,
		MessageInfos:      file_films_proto_msgTypes,
	}.Build()
	File_films_proto = out.File
	file_films_proto_rawDesc = nil
	file_films_proto_goTypes = nil
	file_films_proto_depIdxs = nil
}
