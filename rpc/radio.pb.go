// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/radio.proto

package rpc // import "github.com/R-a-dio/valkyrie/rpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Song struct {
	// song identifier (esong.id)
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// sha1 hexdigest of metadata contents
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// short metadata
	Metadata string `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// song length
	Length *duration.Duration `protobuf:"bytes,4,opt,name=length,proto3" json:"length,omitempty"`
	// last time this song was played
	LastPlayed *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_played,json=lastPlayed,proto3" json:"last_played,omitempty"`
	// DatabaseTrack fields
	TrackId              int32                `protobuf:"varint,16,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
	Artist               string               `protobuf:"bytes,17,opt,name=artist,proto3" json:"artist,omitempty"`
	Title                string               `protobuf:"bytes,18,opt,name=title,proto3" json:"title,omitempty"`
	Album                string               `protobuf:"bytes,19,opt,name=album,proto3" json:"album,omitempty"`
	FilePath             string               `protobuf:"bytes,20,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Tags                 string               `protobuf:"bytes,21,opt,name=tags,proto3" json:"tags,omitempty"`
	Acceptor             string               `protobuf:"bytes,22,opt,name=acceptor,proto3" json:"acceptor,omitempty"`
	LastEditor           string               `protobuf:"bytes,23,opt,name=last_editor,json=lastEditor,proto3" json:"last_editor,omitempty"`
	Priority             int32                `protobuf:"varint,24,opt,name=priority,proto3" json:"priority,omitempty"`
	Usable               bool                 `protobuf:"varint,25,opt,name=usable,proto3" json:"usable,omitempty"`
	LastRequested        *timestamp.Timestamp `protobuf:"bytes,26,opt,name=last_requested,json=lastRequested,proto3" json:"last_requested,omitempty"`
	RequestCount         int32                `protobuf:"varint,27,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	RequestDelay         *duration.Duration   `protobuf:"bytes,28,opt,name=request_delay,json=requestDelay,proto3" json:"request_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{0}
}
func (m *Song) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Song.Unmarshal(m, b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Song.Marshal(b, m, deterministic)
}
func (dst *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(dst, src)
}
func (m *Song) XXX_Size() int {
	return xxx_messageInfo_Song.Size(m)
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Song) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Song) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Song) GetLength() *duration.Duration {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *Song) GetLastPlayed() *timestamp.Timestamp {
	if m != nil {
		return m.LastPlayed
	}
	return nil
}

func (m *Song) GetTrackId() int32 {
	if m != nil {
		return m.TrackId
	}
	return 0
}

func (m *Song) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *Song) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Song) GetAlbum() string {
	if m != nil {
		return m.Album
	}
	return ""
}

func (m *Song) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *Song) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Song) GetAcceptor() string {
	if m != nil {
		return m.Acceptor
	}
	return ""
}

func (m *Song) GetLastEditor() string {
	if m != nil {
		return m.LastEditor
	}
	return ""
}

func (m *Song) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Song) GetUsable() bool {
	if m != nil {
		return m.Usable
	}
	return false
}

func (m *Song) GetLastRequested() *timestamp.Timestamp {
	if m != nil {
		return m.LastRequested
	}
	return nil
}

func (m *Song) GetRequestCount() int32 {
	if m != nil {
		return m.RequestCount
	}
	return 0
}

func (m *Song) GetRequestDelay() *duration.Duration {
	if m != nil {
		return m.RequestDelay
	}
	return nil
}

type StatusResponse struct {
	// the current user that is streaming
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// the current song being played
	Song *Song `protobuf:"bytes,2,opt,name=song,proto3" json:"song,omitempty"`
	// information about the current song
	Info *SongInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	// information about the current listeners
	ListenerInfo *ListenerInfo `protobuf:"bytes,4,opt,name=listener_info,json=listenerInfo,proto3" json:"listener_info,omitempty"`
	// the current thread to be shown on the website or elsewhere
	Thread string `protobuf:"bytes,5,opt,name=thread,proto3" json:"thread,omitempty"`
	// the current configuration of the streamer
	StreamerConfig       *StreamerConfig `protobuf:"bytes,6,opt,name=streamer_config,json=streamerConfig,proto3" json:"streamer_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{1}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (dst *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(dst, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *StatusResponse) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *StatusResponse) GetInfo() *SongInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *StatusResponse) GetListenerInfo() *ListenerInfo {
	if m != nil {
		return m.ListenerInfo
	}
	return nil
}

func (m *StatusResponse) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *StatusResponse) GetStreamerConfig() *StreamerConfig {
	if m != nil {
		return m.StreamerConfig
	}
	return nil
}

type SongUpdate struct {
	Song                 *Song     `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	Info                 *SongInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SongUpdate) Reset()         { *m = SongUpdate{} }
func (m *SongUpdate) String() string { return proto.CompactTextString(m) }
func (*SongUpdate) ProtoMessage()    {}
func (*SongUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{2}
}
func (m *SongUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongUpdate.Unmarshal(m, b)
}
func (m *SongUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongUpdate.Marshal(b, m, deterministic)
}
func (dst *SongUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongUpdate.Merge(dst, src)
}
func (m *SongUpdate) XXX_Size() int {
	return xxx_messageInfo_SongUpdate.Size(m)
}
func (m *SongUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_SongUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_SongUpdate proto.InternalMessageInfo

func (m *SongUpdate) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *SongUpdate) GetInfo() *SongInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type SongInfo struct {
	// the time this song started playing
	StartTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// the time this song will end playing
	EndTime              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SongInfo) Reset()         { *m = SongInfo{} }
func (m *SongInfo) String() string { return proto.CompactTextString(m) }
func (*SongInfo) ProtoMessage()    {}
func (*SongInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{3}
}
func (m *SongInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongInfo.Unmarshal(m, b)
}
func (m *SongInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongInfo.Marshal(b, m, deterministic)
}
func (dst *SongInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongInfo.Merge(dst, src)
}
func (m *SongInfo) XXX_Size() int {
	return xxx_messageInfo_SongInfo.Size(m)
}
func (m *SongInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SongInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SongInfo proto.InternalMessageInfo

func (m *SongInfo) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *SongInfo) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type StreamerConfig struct {
	// can users request songs to be played right now
	RequestsEnabled bool `protobuf:"varint,1,opt,name=requests_enabled,json=requestsEnabled,proto3" json:"requests_enabled,omitempty"`
	// the queue implementation to use for the streamer
	QueueUsed            string   `protobuf:"bytes,2,opt,name=queue_used,json=queueUsed,proto3" json:"queue_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamerConfig) Reset()         { *m = StreamerConfig{} }
func (m *StreamerConfig) String() string { return proto.CompactTextString(m) }
func (*StreamerConfig) ProtoMessage()    {}
func (*StreamerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{4}
}
func (m *StreamerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamerConfig.Unmarshal(m, b)
}
func (m *StreamerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamerConfig.Marshal(b, m, deterministic)
}
func (dst *StreamerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamerConfig.Merge(dst, src)
}
func (m *StreamerConfig) XXX_Size() int {
	return xxx_messageInfo_StreamerConfig.Size(m)
}
func (m *StreamerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StreamerConfig proto.InternalMessageInfo

func (m *StreamerConfig) GetRequestsEnabled() bool {
	if m != nil {
		return m.RequestsEnabled
	}
	return false
}

func (m *StreamerConfig) GetQueueUsed() string {
	if m != nil {
		return m.QueueUsed
	}
	return ""
}

type User struct {
	// user identifier
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user nickname, this is only a display-name
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// indicates if this user is a robot or not
	IsRobot              bool     `protobuf:"varint,3,opt,name=is_robot,json=isRobot,proto3" json:"is_robot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{5}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetIsRobot() bool {
	if m != nil {
		return m.IsRobot
	}
	return false
}

type ListenerInfo struct {
	// the amount of listeners to the stream
	Listeners            int64    `protobuf:"varint,1,opt,name=listeners,proto3" json:"listeners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenerInfo) Reset()         { *m = ListenerInfo{} }
func (m *ListenerInfo) String() string { return proto.CompactTextString(m) }
func (*ListenerInfo) ProtoMessage()    {}
func (*ListenerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{6}
}
func (m *ListenerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerInfo.Unmarshal(m, b)
}
func (m *ListenerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerInfo.Marshal(b, m, deterministic)
}
func (dst *ListenerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerInfo.Merge(dst, src)
}
func (m *ListenerInfo) XXX_Size() int {
	return xxx_messageInfo_ListenerInfo.Size(m)
}
func (m *ListenerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerInfo proto.InternalMessageInfo

func (m *ListenerInfo) GetListeners() int64 {
	if m != nil {
		return m.Listeners
	}
	return 0
}

type SongAnnouncement struct {
	Song                 *Song         `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	Info                 *SongInfo     `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	ListenerInfo         *ListenerInfo `protobuf:"bytes,3,opt,name=listener_info,json=listenerInfo,proto3" json:"listener_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SongAnnouncement) Reset()         { *m = SongAnnouncement{} }
func (m *SongAnnouncement) String() string { return proto.CompactTextString(m) }
func (*SongAnnouncement) ProtoMessage()    {}
func (*SongAnnouncement) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{7}
}
func (m *SongAnnouncement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongAnnouncement.Unmarshal(m, b)
}
func (m *SongAnnouncement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongAnnouncement.Marshal(b, m, deterministic)
}
func (dst *SongAnnouncement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongAnnouncement.Merge(dst, src)
}
func (m *SongAnnouncement) XXX_Size() int {
	return xxx_messageInfo_SongAnnouncement.Size(m)
}
func (m *SongAnnouncement) XXX_DiscardUnknown() {
	xxx_messageInfo_SongAnnouncement.DiscardUnknown(m)
}

var xxx_messageInfo_SongAnnouncement proto.InternalMessageInfo

func (m *SongAnnouncement) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *SongAnnouncement) GetInfo() *SongInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SongAnnouncement) GetListenerInfo() *ListenerInfo {
	if m != nil {
		return m.ListenerInfo
	}
	return nil
}

type SongRequestAnnouncement struct {
	Song                 *Song    `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SongRequestAnnouncement) Reset()         { *m = SongRequestAnnouncement{} }
func (m *SongRequestAnnouncement) String() string { return proto.CompactTextString(m) }
func (*SongRequestAnnouncement) ProtoMessage()    {}
func (*SongRequestAnnouncement) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{8}
}
func (m *SongRequestAnnouncement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongRequestAnnouncement.Unmarshal(m, b)
}
func (m *SongRequestAnnouncement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongRequestAnnouncement.Marshal(b, m, deterministic)
}
func (dst *SongRequestAnnouncement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongRequestAnnouncement.Merge(dst, src)
}
func (m *SongRequestAnnouncement) XXX_Size() int {
	return xxx_messageInfo_SongRequestAnnouncement.Size(m)
}
func (m *SongRequestAnnouncement) XXX_DiscardUnknown() {
	xxx_messageInfo_SongRequestAnnouncement.DiscardUnknown(m)
}

var xxx_messageInfo_SongRequestAnnouncement proto.InternalMessageInfo

func (m *SongRequestAnnouncement) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

type QueueEntry struct {
	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	// is_user_request indicates if this was a request made by a human
	IsUserRequest bool `protobuf:"varint,2,opt,name=is_user_request,json=isUserRequest,proto3" json:"is_user_request,omitempty"`
	// user_identifier is the way we identify the user that added this to the
	// queue; This can be anything that uniquely identifies a user
	UserIdentifier string `protobuf:"bytes,3,opt,name=user_identifier,json=userIdentifier,proto3" json:"user_identifier,omitempty"`
	// expected_start_time is the expected time this song will start playing
	ExpectedStartTime    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expected_start_time,json=expectedStartTime,proto3" json:"expected_start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueueEntry) Reset()         { *m = QueueEntry{} }
func (m *QueueEntry) String() string { return proto.CompactTextString(m) }
func (*QueueEntry) ProtoMessage()    {}
func (*QueueEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{9}
}
func (m *QueueEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueEntry.Unmarshal(m, b)
}
func (m *QueueEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueEntry.Marshal(b, m, deterministic)
}
func (dst *QueueEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueEntry.Merge(dst, src)
}
func (m *QueueEntry) XXX_Size() int {
	return xxx_messageInfo_QueueEntry.Size(m)
}
func (m *QueueEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueEntry.DiscardUnknown(m)
}

var xxx_messageInfo_QueueEntry proto.InternalMessageInfo

func (m *QueueEntry) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *QueueEntry) GetIsUserRequest() bool {
	if m != nil {
		return m.IsUserRequest
	}
	return false
}

func (m *QueueEntry) GetUserIdentifier() string {
	if m != nil {
		return m.UserIdentifier
	}
	return ""
}

func (m *QueueEntry) GetExpectedStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpectedStartTime
	}
	return nil
}

type QueueInfo struct {
	// the name of the queue implementation
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the entries in the queue
	Entries              []*QueueEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueueInfo) Reset()         { *m = QueueInfo{} }
func (m *QueueInfo) String() string { return proto.CompactTextString(m) }
func (*QueueInfo) ProtoMessage()    {}
func (*QueueInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{10}
}
func (m *QueueInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueInfo.Unmarshal(m, b)
}
func (m *QueueInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueInfo.Marshal(b, m, deterministic)
}
func (dst *QueueInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueInfo.Merge(dst, src)
}
func (m *QueueInfo) XXX_Size() int {
	return xxx_messageInfo_QueueInfo.Size(m)
}
func (m *QueueInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QueueInfo proto.InternalMessageInfo

func (m *QueueInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueueInfo) GetEntries() []*QueueEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type SongRequest struct {
	UserIdentifier       string   `protobuf:"bytes,1,opt,name=user_identifier,json=userIdentifier,proto3" json:"user_identifier,omitempty"`
	Song                 *Song    `protobuf:"bytes,2,opt,name=song,proto3" json:"song,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SongRequest) Reset()         { *m = SongRequest{} }
func (m *SongRequest) String() string { return proto.CompactTextString(m) }
func (*SongRequest) ProtoMessage()    {}
func (*SongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{11}
}
func (m *SongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongRequest.Unmarshal(m, b)
}
func (m *SongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongRequest.Marshal(b, m, deterministic)
}
func (dst *SongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongRequest.Merge(dst, src)
}
func (m *SongRequest) XXX_Size() int {
	return xxx_messageInfo_SongRequest.Size(m)
}
func (m *SongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SongRequest proto.InternalMessageInfo

func (m *SongRequest) GetUserIdentifier() string {
	if m != nil {
		return m.UserIdentifier
	}
	return ""
}

func (m *SongRequest) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

type RequestResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestResponse) Reset()         { *m = RequestResponse{} }
func (m *RequestResponse) String() string { return proto.CompactTextString(m) }
func (*RequestResponse) ProtoMessage()    {}
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_radio_c0de4a8355798ad9, []int{12}
}
func (m *RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestResponse.Unmarshal(m, b)
}
func (m *RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestResponse.Marshal(b, m, deterministic)
}
func (dst *RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestResponse.Merge(dst, src)
}
func (m *RequestResponse) XXX_Size() int {
	return xxx_messageInfo_RequestResponse.Size(m)
}
func (m *RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestResponse proto.InternalMessageInfo

func (m *RequestResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RequestResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Song)(nil), "radio.Song")
	proto.RegisterType((*StatusResponse)(nil), "radio.StatusResponse")
	proto.RegisterType((*SongUpdate)(nil), "radio.SongUpdate")
	proto.RegisterType((*SongInfo)(nil), "radio.SongInfo")
	proto.RegisterType((*StreamerConfig)(nil), "radio.StreamerConfig")
	proto.RegisterType((*User)(nil), "radio.User")
	proto.RegisterType((*ListenerInfo)(nil), "radio.ListenerInfo")
	proto.RegisterType((*SongAnnouncement)(nil), "radio.SongAnnouncement")
	proto.RegisterType((*SongRequestAnnouncement)(nil), "radio.SongRequestAnnouncement")
	proto.RegisterType((*QueueEntry)(nil), "radio.QueueEntry")
	proto.RegisterType((*QueueInfo)(nil), "radio.QueueInfo")
	proto.RegisterType((*SongRequest)(nil), "radio.SongRequest")
	proto.RegisterType((*RequestResponse)(nil), "radio.RequestResponse")
}

func init() { proto.RegisterFile("rpc/radio.proto", fileDescriptor_radio_c0de4a8355798ad9) }

var fileDescriptor_radio_c0de4a8355798ad9 = []byte{
	// 1196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0x96, 0x1d, 0x27, 0xb6, 0x8f, 0x93, 0x38, 0x99, 0xb6, 0xe9, 0xd6, 0xed, 0xaf, 0xc9, 0x6f,
	0x2b, 0x41, 0x10, 0xd4, 0x51, 0x5d, 0x50, 0xff, 0x20, 0x8a, 0xd2, 0x36, 0x17, 0x41, 0x2d, 0x2a,
	0xeb, 0x16, 0xa4, 0xde, 0x58, 0x93, 0xdd, 0x13, 0x7b, 0xd4, 0xf5, 0xcc, 0x76, 0x66, 0xb6, 0x34,
	0x12, 0x2f, 0xc0, 0x2d, 0x37, 0x3c, 0x03, 0xaf, 0xc2, 0x4b, 0x70, 0xcb, 0x1d, 0xaf, 0x80, 0xe6,
	0xec, 0x6c, 0xe2, 0xc4, 0x71, 0x5c, 0x10, 0x77, 0x7b, 0xfe, 0xcd, 0x9c, 0xef, 0x9b, 0xef, 0x1c,
	0x1b, 0xda, 0x3a, 0x8b, 0x77, 0x34, 0x4f, 0x84, 0xea, 0x66, 0x5a, 0x59, 0xc5, 0x16, 0xc9, 0xe8,
	0xdc, 0x1c, 0x2a, 0x35, 0x4c, 0x71, 0x87, 0x9c, 0x07, 0xf9, 0xe1, 0x4e, 0x92, 0x6b, 0x6e, 0x85,
	0x92, 0x45, 0x5a, 0xe7, 0xfa, 0xd9, 0x38, 0x8e, 0x33, 0x7b, 0xe4, 0x83, 0x53, 0xc5, 0x3f, 0x6a,
	0x9e, 0x65, 0xa8, 0x8d, 0x8f, 0x6f, 0x9e, 0x8d, 0x5b, 0x31, 0x46, 0x63, 0xf9, 0x38, 0x2b, 0x12,
	0xc2, 0x3f, 0x6a, 0x50, 0xeb, 0x2b, 0x39, 0x64, 0xab, 0x50, 0x15, 0x49, 0x50, 0xd9, 0xaa, 0x6c,
	0x2f, 0x46, 0x55, 0x91, 0x30, 0x06, 0xb5, 0x11, 0x37, 0xa3, 0xa0, 0xba, 0x55, 0xd9, 0x6e, 0x46,
	0xf4, 0xcd, 0x3a, 0xd0, 0x18, 0xa3, 0xe5, 0x09, 0xb7, 0x3c, 0x58, 0x20, 0xff, 0xb1, 0xcd, 0xee,
	0xc0, 0x52, 0x8a, 0x72, 0x68, 0x47, 0x41, 0x6d, 0xab, 0xb2, 0xdd, 0xea, 0x5d, 0xeb, 0x16, 0x57,
	0x77, 0xcb, 0xab, 0xbb, 0x4f, 0x3d, 0xae, 0xc8, 0x27, 0xb2, 0x2f, 0xa1, 0x95, 0x72, 0x63, 0x07,
	0x59, 0xca, 0x8f, 0x30, 0x09, 0x16, 0xa9, 0xae, 0x33, 0x55, 0xf7, 0xb2, 0x6c, 0x39, 0x02, 0x97,
	0xfe, 0x82, 0xb2, 0xd9, 0x35, 0x68, 0x58, 0xcd, 0xe3, 0x37, 0x03, 0x91, 0x04, 0x6b, 0xd4, 0x75,
	0x9d, 0xec, 0xfd, 0x84, 0x6d, 0xc0, 0x12, 0xd7, 0x56, 0x18, 0x1b, 0xac, 0x53, 0x93, 0xde, 0x62,
	0x97, 0x61, 0xd1, 0x0a, 0x9b, 0x62, 0xc0, 0xc8, 0x5d, 0x18, 0xce, 0xcb, 0xd3, 0x83, 0x7c, 0x1c,
	0x5c, 0x2a, 0xbc, 0x64, 0xb0, 0xeb, 0xd0, 0x3c, 0x14, 0x29, 0x0e, 0x32, 0x6e, 0x47, 0xc1, 0xe5,
	0x02, 0xab, 0x73, 0xbc, 0xe0, 0x76, 0xe4, 0xb8, 0xb1, 0x7c, 0x68, 0x82, 0x2b, 0x05, 0x37, 0xee,
	0xdb, 0x71, 0xc3, 0xe3, 0x18, 0x33, 0xab, 0x74, 0xb0, 0x51, 0xe4, 0x97, 0x36, 0xdb, 0xf4, 0x40,
	0x31, 0x11, 0x2e, 0x7c, 0x95, 0xc2, 0x04, 0x66, 0x8f, 0x3c, 0xae, 0x38, 0xd3, 0x42, 0x69, 0x61,
	0x8f, 0x82, 0x80, 0xc0, 0x1c, 0xdb, 0x0e, 0x4d, 0x6e, 0xf8, 0x41, 0x8a, 0xc1, 0xb5, 0xad, 0xca,
	0x76, 0x23, 0xf2, 0x16, 0xdb, 0x85, 0x55, 0x3a, 0x54, 0xe3, 0xdb, 0x1c, 0x8d, 0xc5, 0x24, 0xe8,
	0xcc, 0x25, 0x70, 0xc5, 0x55, 0x44, 0x65, 0x01, 0xbb, 0x05, 0x2b, 0xbe, 0x7a, 0x10, 0xab, 0x5c,
	0xda, 0xe0, 0x3a, 0xdd, 0xbd, 0xec, 0x9d, 0x4f, 0x9c, 0x8f, 0x3d, 0x3a, 0x49, 0x4a, 0x30, 0xe5,
	0x47, 0xc1, 0x8d, 0x79, 0xef, 0x5b, 0xd6, 0x3f, 0x75, 0xe9, 0xe1, 0xcf, 0x55, 0x58, 0xed, 0x5b,
	0x6e, 0x73, 0x13, 0xa1, 0xc9, 0x94, 0x34, 0xc8, 0x36, 0xa1, 0x96, 0x1b, 0xd4, 0xa4, 0xb6, 0x56,
	0xaf, 0xd5, 0x2d, 0xa6, 0xe2, 0x95, 0x41, 0x1d, 0x51, 0xc0, 0x25, 0x18, 0x25, 0x87, 0x24, 0xbe,
	0x93, 0x04, 0xa7, 0xd3, 0x88, 0x02, 0xec, 0x16, 0xd4, 0x84, 0x3c, 0x54, 0xa4, 0xc2, 0x56, 0xaf,
	0x3d, 0x91, 0xb0, 0x2f, 0x0f, 0x55, 0x44, 0x41, 0x76, 0x1f, 0x56, 0x52, 0x61, 0x2c, 0x4a, 0xd4,
	0x03, 0xca, 0x2e, 0x94, 0x79, 0xc9, 0x67, 0x3f, 0xf3, 0x31, 0xaa, 0x58, 0x4e, 0x27, 0x2c, 0xc7,
	0xb9, 0x1d, 0x69, 0xe4, 0x85, 0x28, 0x9b, 0x91, 0xb7, 0xd8, 0x23, 0x68, 0x1b, 0xab, 0x91, 0x8f,
	0x51, 0x0f, 0x62, 0x25, 0x0f, 0xc5, 0x30, 0x58, 0xa2, 0x33, 0xaf, 0x94, 0x1d, 0xf8, 0xe8, 0x13,
	0x0a, 0x46, 0xab, 0xe6, 0x94, 0x1d, 0x46, 0x00, 0xae, 0xc7, 0x57, 0x59, 0xc2, 0x2d, 0x1e, 0xa3,
	0xac, 0xcc, 0x43, 0x59, 0xbd, 0x00, 0x65, 0xf8, 0x13, 0x34, 0x4a, 0x0f, 0x7b, 0x00, 0x60, 0x2c,
	0xd7, 0x76, 0xe0, 0xc6, 0xdc, 0x93, 0x73, 0x91, 0x1e, 0x9a, 0x94, 0xed, 0x6c, 0xf6, 0x05, 0x34,
	0x50, 0x26, 0x45, 0x61, 0x6d, 0x6e, 0x61, 0x1d, 0x65, 0xe2, 0xac, 0xf0, 0xb5, 0x7b, 0xdc, 0x49,
	0x8c, 0xec, 0x13, 0x58, 0xf3, 0xef, 0x6f, 0x06, 0x28, 0x9d, 0x54, 0x8b, 0xb5, 0xd2, 0x88, 0xda,
	0xa5, 0x7f, 0xaf, 0x70, 0xb3, 0xff, 0x01, 0xbc, 0xcd, 0x31, 0xc7, 0x41, 0x6e, 0x30, 0xf1, 0x9b,
	0xa6, 0x49, 0x9e, 0x57, 0x06, 0x93, 0xf0, 0x39, 0xd4, 0x9c, 0x26, 0xa6, 0x56, 0x53, 0x07, 0x1a,
	0x52, 0xc4, 0x6f, 0x24, 0x1f, 0xa3, 0x2f, 0x3a, 0xb6, 0xdd, 0x5a, 0x10, 0x66, 0xa0, 0xd5, 0x81,
	0xb2, 0x84, 0xbf, 0x11, 0xd5, 0x85, 0x89, 0x9c, 0x19, 0x7e, 0x06, 0xcb, 0x93, 0x4f, 0xce, 0x6e,
	0x40, 0xb3, 0x7c, 0x74, 0x43, 0xa7, 0x2f, 0x44, 0x27, 0x8e, 0xf0, 0x97, 0x0a, 0xac, 0x39, 0x5e,
	0x77, 0xa5, 0x54, 0xb9, 0x8c, 0x71, 0x8c, 0xd2, 0xfe, 0x37, 0x2f, 0x36, 0xad, 0xcb, 0x85, 0x0f,
	0xd4, 0x65, 0xf8, 0x10, 0xae, 0xd2, 0x65, 0x05, 0x8f, 0xff, 0xa8, 0xb5, 0xf0, 0xf7, 0x0a, 0xc0,
	0x77, 0x8e, 0xdb, 0x3d, 0x69, 0xf5, 0xd1, 0x7c, 0x28, 0x1f, 0x41, 0x5b, 0x18, 0xf7, 0x32, 0xba,
	0x5c, 0x31, 0x84, 0xaa, 0x11, 0xad, 0x08, 0x43, 0xa3, 0x5a, 0x38, 0xd9, 0xc7, 0xd0, 0xa6, 0x24,
	0x91, 0xa0, 0xb4, 0xe2, 0x50, 0xa0, 0xf6, 0xbf, 0x0d, 0xab, 0xce, 0xbd, 0x7f, 0xec, 0x65, 0xdf,
	0xc0, 0x25, 0x7c, 0x9f, 0x61, 0x6c, 0x31, 0x19, 0x4c, 0xa8, 0x74, 0xbe, 0xd8, 0xd6, 0xcb, 0xb2,
	0x7e, 0xa9, 0xd6, 0xf0, 0x19, 0x34, 0x09, 0x0b, 0x3d, 0x24, 0x83, 0x1a, 0x69, 0xa1, 0x52, 0xac,
	0x63, 0xd2, 0xc1, 0xa7, 0x50, 0x47, 0x69, 0xb5, 0x40, 0x13, 0x54, 0xb7, 0x16, 0xb6, 0x5b, 0xbd,
	0x75, 0x8f, 0xf0, 0x84, 0x82, 0xa8, 0xcc, 0x08, 0x7f, 0x80, 0xd6, 0x04, 0xad, 0xe7, 0x21, 0xaa,
	0x9c, 0x8b, 0x68, 0xde, 0x9a, 0x0a, 0xbf, 0x82, 0xb6, 0x3f, 0xf4, 0x78, 0xf7, 0x05, 0x50, 0x37,
	0x79, 0x1c, 0xa3, 0x31, 0x7e, 0x2a, 0x4a, 0x93, 0xad, 0xc1, 0xc2, 0xd8, 0x0c, 0xbd, 0xa2, 0xdd,
	0x67, 0xef, 0xaf, 0x2a, 0xd4, 0x9f, 0x73, 0xc9, 0x87, 0xa8, 0xd9, 0x3d, 0x58, 0x2a, 0xb6, 0x28,
	0xdb, 0x98, 0xa2, 0x6a, 0xcf, 0xfd, 0x23, 0xe8, 0x9c, 0xec, 0xa0, 0x53, 0xcb, 0xb6, 0x0b, 0xf5,
	0x3e, 0x5a, 0x1a, 0xa4, 0xc9, 0x4d, 0xdb, 0x99, 0x71, 0x0c, 0xfb, 0x9c, 0xf2, 0xe9, 0x3f, 0xc1,
	0xfa, 0x04, 0xa2, 0x62, 0x67, 0xcd, 0xac, 0x7a, 0x0c, 0xeb, 0xae, 0xea, 0xf4, 0x2a, 0x38, 0x7f,
	0x2b, 0xce, 0x3c, 0x63, 0x17, 0x9a, 0x7d, 0xb4, 0x2f, 0x8b, 0x55, 0x7b, 0x63, 0x2a, 0xa9, 0x6f,
	0xb5, 0x90, 0xc3, 0xef, 0x79, 0x9a, 0xcf, 0x6e, 0xe3, 0x11, 0xb4, 0xfb, 0x68, 0x4f, 0x8d, 0xf9,
	0x79, 0x63, 0x35, 0xab, 0xbe, 0xf7, 0x6b, 0x05, 0x9a, 0xe5, 0x58, 0x69, 0xf6, 0x35, 0x2c, 0x97,
	0x06, 0xf1, 0x71, 0x75, 0x82, 0x8f, 0xc9, 0xe1, 0x9b, 0xd9, 0xce, 0x3e, 0xb4, 0xcb, 0xbc, 0x52,
	0x5c, 0x37, 0x27, 0x55, 0x32, 0x3d, 0xc7, 0x33, 0x3b, 0xfb, 0xad, 0x0a, 0x8d, 0x92, 0x47, 0x76,
	0x0f, 0x16, 0x69, 0x16, 0x66, 0x6a, 0x61, 0x56, 0x43, 0x0f, 0xa1, 0xd6, 0xb7, 0x2a, 0x63, 0xd3,
	0xe3, 0xf6, 0x58, 0xa9, 0xf4, 0x62, 0x6e, 0x1f, 0x40, 0xcb, 0x37, 0x4c, 0x64, 0xb0, 0x69, 0x20,
	0x9d, 0x0d, 0xef, 0x3b, 0x2b, 0xfa, 0x87, 0xf4, 0xb2, 0xff, 0x4e, 0x15, 0x77, 0x60, 0x91, 0x66,
	0x76, 0x26, 0xd6, 0xb5, 0xc9, 0xc9, 0x76, 0xaf, 0xdb, 0xfb, 0xb3, 0x52, 0xd6, 0xdc, 0x03, 0xd8,
	0x4d, 0x92, 0x92, 0xfb, 0xe9, 0x1d, 0x30, 0xf3, 0xd6, 0xfb, 0x0e, 0xac, 0x41, 0xfd, 0x0e, 0xbf,
	0xc5, 0xf7, 0xb3, 0x79, 0x9e, 0x3e, 0xd1, 0x0d, 0x6a, 0x84, 0x63, 0xf5, 0x0e, 0xcf, 0xbb, 0xee,
	0x02, 0xde, 0xd9, 0x5d, 0xa8, 0xef, 0x15, 0x0b, 0xe9, 0xc3, 0xa1, 0x3e, 0xfe, 0xff, 0xeb, 0xcd,
	0xa1, 0xb0, 0xa3, 0xfc, 0xa0, 0x1b, 0xab, 0xf1, 0x4e, 0x74, 0x9b, 0xdf, 0x4e, 0x84, 0xda, 0x79,
	0xc7, 0xd3, 0x37, 0x47, 0x5a, 0xe0, 0x8e, 0xce, 0xe2, 0x83, 0x25, 0x3a, 0xe4, 0xee, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x52, 0xbf, 0xdf, 0x11, 0x81, 0x0c, 0x00, 0x00,
}
