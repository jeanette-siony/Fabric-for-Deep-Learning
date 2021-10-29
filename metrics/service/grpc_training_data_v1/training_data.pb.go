// Code generated by protoc-gen-go. DO NOT EDIT.
// source: training_data.proto

/*
Package grpc_training_data_v1 is a generated protocol buffer package.

It is generated from these files:
	training_data.proto

It has these top-level messages:
	MetaInfo
	LogLine
	LogLineBatch
	Any
	EMetrics
	EMetricsBatch
	Query
	DeleteQuery
	AddResponse
	DeleteResponse
	HelloResponse
	Empty
*/
package grpc_training_data_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Any_DataType int32

const (
	Any_STRING     Any_DataType = 0
	Any_JSONSTRING Any_DataType = 1
	Any_INT        Any_DataType = 2
	Any_FLOAT      Any_DataType = 3
)

var Any_DataType_name = map[int32]string{
	0: "STRING",
	1: "JSONSTRING",
	2: "INT",
	3: "FLOAT",
}
var Any_DataType_value = map[string]int32{
	"STRING":     0,
	"JSONSTRING": 1,
	"INT":        2,
	"FLOAT":      3,
}

func (x Any_DataType) String() string {
	return proto.EnumName(Any_DataType_name, int32(x))
}
func (Any_DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Query_SearchType int32

const (
	Query_TERM   Query_SearchType = 0
	Query_NESTED Query_SearchType = 1
	Query_MATCH  Query_SearchType = 2
	Query_ALL    Query_SearchType = 3
)

var Query_SearchType_name = map[int32]string{
	0: "TERM",
	1: "NESTED",
	2: "MATCH",
	3: "ALL",
}
var Query_SearchType_value = map[string]int32{
	"TERM":   0,
	"NESTED": 1,
	"MATCH":  2,
	"ALL":    3,
}

func (x Query_SearchType) String() string {
	return proto.EnumName(Query_SearchType_name, int32(x))
}
func (Query_SearchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type MetaInfo struct {
	// Unique id identifying the training job
	TrainingId string `protobuf:"bytes,1,opt,name=training_id,json=trainingId" json:"training_id,omitempty"`
	// Unique id identifying the user
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Time that the metric occured: representing the number of millisecond since midnight January 1, 1970.
	Time int64 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	// Sequential index
	Rindex int64 `protobuf:"varint,4,opt,name=rindex" json:"rindex,omitempty"`
	// Optional subid
	Subid string `protobuf:"bytes,5,opt,name=subid" json:"subid,omitempty"`
}

func (m *MetaInfo) Reset()                    { *m = MetaInfo{} }
func (m *MetaInfo) String() string            { return proto.CompactTextString(m) }
func (*MetaInfo) ProtoMessage()               {}
func (*MetaInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetaInfo) GetTrainingId() string {
	if m != nil {
		return m.TrainingId
	}
	return ""
}

func (m *MetaInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *MetaInfo) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MetaInfo) GetRindex() int64 {
	if m != nil {
		return m.Rindex
	}
	return 0
}

func (m *MetaInfo) GetSubid() string {
	if m != nil {
		return m.Subid
	}
	return ""
}

type LogLine struct {
	Meta *MetaInfo `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Line string    `protobuf:"bytes,2,opt,name=line" json:"line,omitempty"`
}

func (m *LogLine) Reset()                    { *m = LogLine{} }
func (m *LogLine) String() string            { return proto.CompactTextString(m) }
func (*LogLine) ProtoMessage()               {}
func (*LogLine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogLine) GetMeta() *MetaInfo {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *LogLine) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

type LogLineBatch struct {
	Force   bool       `protobuf:"varint,1,opt,name=force" json:"force,omitempty"`
	LogLine []*LogLine `protobuf:"bytes,2,rep,name=logLine" json:"logLine,omitempty"`
}

func (m *LogLineBatch) Reset()                    { *m = LogLineBatch{} }
func (m *LogLineBatch) String() string            { return proto.CompactTextString(m) }
func (*LogLineBatch) ProtoMessage()               {}
func (*LogLineBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogLineBatch) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *LogLineBatch) GetLogLine() []*LogLine {
	if m != nil {
		return m.LogLine
	}
	return nil
}

type Any struct {
	Type  Any_DataType `protobuf:"varint,1,opt,name=type,enum=grpc.training.data.v1.Any_DataType" json:"type,omitempty"`
	Value string       `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Any) Reset()                    { *m = Any{} }
func (m *Any) String() string            { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()               {}
func (*Any) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Any) GetType() Any_DataType {
	if m != nil {
		return m.Type
	}
	return Any_STRING
}

func (m *Any) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type EMetrics struct {
	Meta *MetaInfo `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	// Repeated, order-dependent list of temporal keys
	// Example: {"iteration": 209}
	Etimes map[string]*Any `protobuf:"bytes,2,rep,name=etimes" json:"etimes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Group label, such as test, train, or validate
	Grouplabel string `protobuf:"bytes,3,opt,name=grouplabel" json:"grouplabel,omitempty"`
	// / {"cross_entropy": 0.4430539906024933,	"accuracy": 0.8999999761581421}
	Values map[string]*Any `protobuf:"bytes,4,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EMetrics) Reset()                    { *m = EMetrics{} }
func (m *EMetrics) String() string            { return proto.CompactTextString(m) }
func (*EMetrics) ProtoMessage()               {}
func (*EMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EMetrics) GetMeta() *MetaInfo {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *EMetrics) GetEtimes() map[string]*Any {
	if m != nil {
		return m.Etimes
	}
	return nil
}

func (m *EMetrics) GetGrouplabel() string {
	if m != nil {
		return m.Grouplabel
	}
	return ""
}

func (m *EMetrics) GetValues() map[string]*Any {
	if m != nil {
		return m.Values
	}
	return nil
}

type EMetricsBatch struct {
	Force    bool        `protobuf:"varint,1,opt,name=force" json:"force,omitempty"`
	Emetrics []*EMetrics `protobuf:"bytes,2,rep,name=emetrics" json:"emetrics,omitempty"`
}

func (m *EMetricsBatch) Reset()                    { *m = EMetricsBatch{} }
func (m *EMetricsBatch) String() string            { return proto.CompactTextString(m) }
func (*EMetricsBatch) ProtoMessage()               {}
func (*EMetricsBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EMetricsBatch) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *EMetricsBatch) GetEmetrics() []*EMetrics {
	if m != nil {
		return m.Emetrics
	}
	return nil
}

// Playing with semi-generalized query request.
type Query struct {
	SearchType Query_SearchType `protobuf:"varint,1,opt,name=searchType,enum=grpc.training.data.v1.Query_SearchType" json:"searchType,omitempty"`
	// The following three options are exclusive
	Meta  *MetaInfo `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
	Since string    `protobuf:"bytes,4,opt,name=since" json:"since,omitempty"`
	// Only get this many records
	Pagesize int32 `protobuf:"varint,5,opt,name=pagesize" json:"pagesize,omitempty"`
	// The starting position.  If positive or zero, count from beginning, if negative, count from end.
	Pos int64 `protobuf:"varint,6,opt,name=pos" json:"pos,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Query) GetSearchType() Query_SearchType {
	if m != nil {
		return m.SearchType
	}
	return Query_TERM
}

func (m *Query) GetMeta() *MetaInfo {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Query) GetSince() string {
	if m != nil {
		return m.Since
	}
	return ""
}

func (m *Query) GetPagesize() int32 {
	if m != nil {
		return m.Pagesize
	}
	return 0
}

func (m *Query) GetPos() int64 {
	if m != nil {
		return m.Pos
	}
	return 0
}

type DeleteQuery struct {
	// The following two options are exclusive
	TrainingId string `protobuf:"bytes,1,opt,name=training_id,json=trainingId" json:"training_id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteQuery) Reset()                    { *m = DeleteQuery{} }
func (m *DeleteQuery) String() string            { return proto.CompactTextString(m) }
func (*DeleteQuery) ProtoMessage()               {}
func (*DeleteQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteQuery) GetTrainingId() string {
	if m != nil {
		return m.TrainingId
	}
	return ""
}

func (m *DeleteQuery) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// ***
type AddResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type DeleteResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type HelloResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *HelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*MetaInfo)(nil), "grpc.training.data.v1.MetaInfo")
	proto.RegisterType((*LogLine)(nil), "grpc.training.data.v1.LogLine")
	proto.RegisterType((*LogLineBatch)(nil), "grpc.training.data.v1.LogLineBatch")
	proto.RegisterType((*Any)(nil), "grpc.training.data.v1.Any")
	proto.RegisterType((*EMetrics)(nil), "grpc.training.data.v1.EMetrics")
	proto.RegisterType((*EMetricsBatch)(nil), "grpc.training.data.v1.EMetricsBatch")
	proto.RegisterType((*Query)(nil), "grpc.training.data.v1.Query")
	proto.RegisterType((*DeleteQuery)(nil), "grpc.training.data.v1.DeleteQuery")
	proto.RegisterType((*AddResponse)(nil), "grpc.training.data.v1.AddResponse")
	proto.RegisterType((*DeleteResponse)(nil), "grpc.training.data.v1.DeleteResponse")
	proto.RegisterType((*HelloResponse)(nil), "grpc.training.data.v1.HelloResponse")
	proto.RegisterType((*Empty)(nil), "grpc.training.data.v1.Empty")
	proto.RegisterEnum("grpc.training.data.v1.Any_DataType", Any_DataType_name, Any_DataType_value)
	proto.RegisterEnum("grpc.training.data.v1.Query_SearchType", Query_SearchType_name, Query_SearchType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TrainingData service

type TrainingDataClient interface {
	// Get loglines, based on query
	GetLogs(ctx context.Context, in *Query, opts ...grpc.CallOption) (TrainingData_GetLogsClient, error)
	// Get evaluation metrics records, based on query
	GetEMetrics(ctx context.Context, in *Query, opts ...grpc.CallOption) (TrainingData_GetEMetricsClient, error)
	// Add evaluation metrics record
	AddEMetrics(ctx context.Context, in *EMetrics, opts ...grpc.CallOption) (*AddResponse, error)
	// Add log line record
	AddLogLine(ctx context.Context, in *LogLine, opts ...grpc.CallOption) (*AddResponse, error)
	// Add evaluation metrics record
	AddEMetricsBatch(ctx context.Context, in *EMetricsBatch, opts ...grpc.CallOption) (*AddResponse, error)
	// Add log line record
	AddLogLineBatch(ctx context.Context, in *LogLineBatch, opts ...grpc.CallOption) (*AddResponse, error)
	// Delete all evaluation metrics belonging to a training job or user id
	DeleteEMetrics(ctx context.Context, in *Query, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Delete all log lines belonging to a training job or user id
	DeleteLogLines(ctx context.Context, in *Query, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Delete all log lines belonging to a training job or user id
	DeleteJob(ctx context.Context, in *Query, opts ...grpc.CallOption) (*DeleteResponse, error)
	// ===== In case you want it to say "Hello" =========
	Hello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloResponse, error)
}

type trainingDataClient struct {
	cc *grpc.ClientConn
}

func NewTrainingDataClient(cc *grpc.ClientConn) TrainingDataClient {
	return &trainingDataClient{cc}
}

func (c *trainingDataClient) GetLogs(ctx context.Context, in *Query, opts ...grpc.CallOption) (TrainingData_GetLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TrainingData_serviceDesc.Streams[0], c.cc, "/grpc.training.data.v1.TrainingData/GetLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &trainingDataGetLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrainingData_GetLogsClient interface {
	Recv() (*LogLine, error)
	grpc.ClientStream
}

type trainingDataGetLogsClient struct {
	grpc.ClientStream
}

func (x *trainingDataGetLogsClient) Recv() (*LogLine, error) {
	m := new(LogLine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trainingDataClient) GetEMetrics(ctx context.Context, in *Query, opts ...grpc.CallOption) (TrainingData_GetEMetricsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TrainingData_serviceDesc.Streams[1], c.cc, "/grpc.training.data.v1.TrainingData/GetEMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &trainingDataGetEMetricsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrainingData_GetEMetricsClient interface {
	Recv() (*EMetrics, error)
	grpc.ClientStream
}

type trainingDataGetEMetricsClient struct {
	grpc.ClientStream
}

func (x *trainingDataGetEMetricsClient) Recv() (*EMetrics, error) {
	m := new(EMetrics)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trainingDataClient) AddEMetrics(ctx context.Context, in *EMetrics, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/AddEMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingDataClient) AddLogLine(ctx context.Context, in *LogLine, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/AddLogLine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingDataClient) AddEMetricsBatch(ctx context.Context, in *EMetricsBatch, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/AddEMetricsBatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingDataClient) AddLogLineBatch(ctx context.Context, in *LogLineBatch, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/AddLogLineBatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingDataClient) DeleteEMetrics(ctx context.Context, in *Query, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/DeleteEMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingDataClient) DeleteLogLines(ctx context.Context, in *Query, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/DeleteLogLines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingDataClient) DeleteJob(ctx context.Context, in *Query, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/DeleteJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingDataClient) Hello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/grpc.training.data.v1.TrainingData/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrainingData service

type TrainingDataServer interface {
	// Get loglines, based on query
	GetLogs(*Query, TrainingData_GetLogsServer) error
	// Get evaluation metrics records, based on query
	GetEMetrics(*Query, TrainingData_GetEMetricsServer) error
	// Add evaluation metrics record
	AddEMetrics(context.Context, *EMetrics) (*AddResponse, error)
	// Add log line record
	AddLogLine(context.Context, *LogLine) (*AddResponse, error)
	// Add evaluation metrics record
	AddEMetricsBatch(context.Context, *EMetricsBatch) (*AddResponse, error)
	// Add log line record
	AddLogLineBatch(context.Context, *LogLineBatch) (*AddResponse, error)
	// Delete all evaluation metrics belonging to a training job or user id
	DeleteEMetrics(context.Context, *Query) (*DeleteResponse, error)
	// Delete all log lines belonging to a training job or user id
	DeleteLogLines(context.Context, *Query) (*DeleteResponse, error)
	// Delete all log lines belonging to a training job or user id
	DeleteJob(context.Context, *Query) (*DeleteResponse, error)
	// ===== In case you want it to say "Hello" =========
	Hello(context.Context, *Empty) (*HelloResponse, error)
}

func RegisterTrainingDataServer(s *grpc.Server, srv TrainingDataServer) {
	s.RegisterService(&_TrainingData_serviceDesc, srv)
}

func _TrainingData_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrainingDataServer).GetLogs(m, &trainingDataGetLogsServer{stream})
}

type TrainingData_GetLogsServer interface {
	Send(*LogLine) error
	grpc.ServerStream
}

type trainingDataGetLogsServer struct {
	grpc.ServerStream
}

func (x *trainingDataGetLogsServer) Send(m *LogLine) error {
	return x.ServerStream.SendMsg(m)
}

func _TrainingData_GetEMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrainingDataServer).GetEMetrics(m, &trainingDataGetEMetricsServer{stream})
}

type TrainingData_GetEMetricsServer interface {
	Send(*EMetrics) error
	grpc.ServerStream
}

type trainingDataGetEMetricsServer struct {
	grpc.ServerStream
}

func (x *trainingDataGetEMetricsServer) Send(m *EMetrics) error {
	return x.ServerStream.SendMsg(m)
}

func _TrainingData_AddEMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).AddEMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/AddEMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).AddEMetrics(ctx, req.(*EMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingData_AddLogLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).AddLogLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/AddLogLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).AddLogLine(ctx, req.(*LogLine))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingData_AddEMetricsBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EMetricsBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).AddEMetricsBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/AddEMetricsBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).AddEMetricsBatch(ctx, req.(*EMetricsBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingData_AddLogLineBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLineBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).AddLogLineBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/AddLogLineBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).AddLogLineBatch(ctx, req.(*LogLineBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingData_DeleteEMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).DeleteEMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/DeleteEMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).DeleteEMetrics(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingData_DeleteLogLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).DeleteLogLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/DeleteLogLines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).DeleteLogLines(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingData_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/DeleteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).DeleteJob(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingData_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingDataServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.training.data.v1.TrainingData/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingDataServer).Hello(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrainingData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.training.data.v1.TrainingData",
	HandlerType: (*TrainingDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEMetrics",
			Handler:    _TrainingData_AddEMetrics_Handler,
		},
		{
			MethodName: "AddLogLine",
			Handler:    _TrainingData_AddLogLine_Handler,
		},
		{
			MethodName: "AddEMetricsBatch",
			Handler:    _TrainingData_AddEMetricsBatch_Handler,
		},
		{
			MethodName: "AddLogLineBatch",
			Handler:    _TrainingData_AddLogLineBatch_Handler,
		},
		{
			MethodName: "DeleteEMetrics",
			Handler:    _TrainingData_DeleteEMetrics_Handler,
		},
		{
			MethodName: "DeleteLogLines",
			Handler:    _TrainingData_DeleteLogLines_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _TrainingData_DeleteJob_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _TrainingData_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogs",
			Handler:       _TrainingData_GetLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetEMetrics",
			Handler:       _TrainingData_GetEMetrics_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "training_data.proto",
}

func init() { proto.RegisterFile("training_data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x6f, 0xf3, 0x44,
	0x10, 0x8e, 0xe3, 0x7c, 0x8e, 0xdb, 0x62, 0x2d, 0x5f, 0x56, 0x84, 0xda, 0xb2, 0x05, 0xb5, 0x02,
	0xc9, 0x2a, 0xa9, 0x04, 0x55, 0x39, 0x85, 0x36, 0xa4, 0x29, 0x49, 0x0a, 0x1b, 0x03, 0x17, 0x54,
	0xe4, 0xd8, 0x5b, 0xd7, 0xc2, 0xb1, 0x2d, 0xaf, 0x53, 0x61, 0x6e, 0x5c, 0x38, 0x73, 0x7b, 0xff,
	0xe3, 0xfb, 0x2b, 0x5e, 0xed, 0xfa, 0x23, 0xa9, 0xf4, 0x3a, 0x49, 0xab, 0xde, 0x76, 0x76, 0x67,
	0x9e, 0x79, 0x3c, 0x33, 0xcf, 0x24, 0xf0, 0x61, 0x1c, 0x99, 0xae, 0xef, 0xfa, 0xce, 0x9f, 0xb6,
	0x19, 0x9b, 0x7a, 0x18, 0x05, 0x71, 0x80, 0x3e, 0x76, 0xa2, 0xd0, 0xd2, 0xf3, 0x17, 0x5d, 0xbc,
	0x3c, 0x7e, 0x83, 0xff, 0x93, 0xa0, 0x35, 0xa6, 0xb1, 0x39, 0xf4, 0xef, 0x03, 0x74, 0x00, 0x4a,
	0x11, 0xea, 0xda, 0x9a, 0x74, 0x28, 0x9d, 0xb4, 0x09, 0xe4, 0x57, 0x43, 0x1b, 0x7d, 0x0a, 0xcd,
	0x05, 0xa3, 0x11, 0x7f, 0xac, 0x8a, 0xc7, 0x06, 0x37, 0x87, 0x36, 0x42, 0x50, 0x8b, 0xdd, 0x39,
	0xd5, 0xe4, 0x43, 0xe9, 0x44, 0x26, 0xe2, 0x8c, 0x3e, 0x81, 0x46, 0xe4, 0xfa, 0x36, 0xfd, 0x5b,
	0xab, 0x89, 0xdb, 0xcc, 0x42, 0x1f, 0x41, 0x9d, 0x2d, 0x66, 0xae, 0xad, 0xd5, 0x05, 0x44, 0x6a,
	0x60, 0x02, 0xcd, 0x51, 0xe0, 0x8c, 0x5c, 0x9f, 0xa2, 0x33, 0xa8, 0xcd, 0x69, 0x6c, 0x8a, 0xfc,
	0x4a, 0xf7, 0x40, 0x7f, 0x2f, 0x73, 0x3d, 0x67, 0x4d, 0x84, 0x33, 0x67, 0xe0, 0xb9, 0x3e, 0xcd,
	0x78, 0x89, 0x33, 0xbe, 0x83, 0x9d, 0x0c, 0xf3, 0x07, 0x33, 0xb6, 0x1e, 0x78, 0xe6, 0xfb, 0x20,
	0xb2, 0xa8, 0x40, 0x6e, 0x91, 0xd4, 0x40, 0xe7, 0xd0, 0xf4, 0x52, 0x2f, 0xad, 0x7a, 0x28, 0x9f,
	0x28, 0xdd, 0xfd, 0x92, 0x8c, 0x19, 0x16, 0xc9, 0xdd, 0xf1, 0xff, 0x12, 0xc8, 0x3d, 0x3f, 0x41,
	0xdf, 0x41, 0x2d, 0x4e, 0xc2, 0x14, 0x76, 0xaf, 0x7b, 0x54, 0x12, 0xde, 0xf3, 0x13, 0xfd, 0xca,
	0x8c, 0x4d, 0x23, 0x09, 0x29, 0x11, 0x01, 0x9c, 0xd0, 0xa3, 0xe9, 0x2d, 0x72, 0xd6, 0xa9, 0x81,
	0x2f, 0xa0, 0x95, 0xfb, 0x21, 0x80, 0xc6, 0xd4, 0x20, 0xc3, 0xc9, 0x40, 0xad, 0xa0, 0x3d, 0x80,
	0x9b, 0xe9, 0xed, 0x24, 0xb3, 0x25, 0xd4, 0x04, 0x79, 0x38, 0x31, 0xd4, 0x2a, 0x6a, 0x43, 0xfd,
	0xc7, 0xd1, 0x6d, 0xcf, 0x50, 0x65, 0xfc, 0x46, 0x86, 0x56, 0x7f, 0x4c, 0xe3, 0xc8, 0xb5, 0xd8,
	0xcb, 0x0a, 0x79, 0x09, 0x0d, 0xca, 0xfb, 0xc7, 0xb2, 0x6a, 0x7c, 0x5d, 0x12, 0x96, 0x67, 0xd1,
	0xfb, 0xc2, 0xbb, 0xef, 0xc7, 0x51, 0x42, 0xb2, 0x50, 0xb4, 0x0f, 0xe0, 0x44, 0xc1, 0x22, 0xf4,
	0xcc, 0x19, 0xf5, 0xc4, 0x54, 0xb4, 0xc9, 0xca, 0x0d, 0x4f, 0x22, 0xbe, 0x95, 0x69, 0xb5, 0xed,
	0x92, 0xfc, 0x26, 0xbc, 0xb3, 0x24, 0x69, 0x68, 0xe7, 0x57, 0x50, 0x56, 0x72, 0x23, 0x15, 0xe4,
	0xbf, 0x68, 0x92, 0x4d, 0x2d, 0x3f, 0xa2, 0xd3, 0xd5, 0xf2, 0x2a, 0xdd, 0x4e, 0x79, 0x63, 0xb2,
	0xd2, 0x5f, 0x54, 0xcf, 0x25, 0x0e, 0xbb, 0x92, 0xed, 0xb5, 0x60, 0xf1, 0x0c, 0x76, 0xf3, 0xaf,
	0x59, 0x37, 0x8d, 0xdf, 0x43, 0x8b, 0xce, 0x53, 0xb7, 0xac, 0x01, 0x07, 0x1b, 0x6a, 0x43, 0x8a,
	0x00, 0xfc, 0x6f, 0x15, 0xea, 0xbf, 0x2c, 0x68, 0x94, 0xa0, 0x01, 0x00, 0xa3, 0x66, 0x64, 0x3d,
	0x18, 0xcb, 0xc1, 0x3c, 0x2e, 0x01, 0x12, 0x11, 0xfa, 0xb4, 0x70, 0x27, 0x2b, 0xa1, 0xc5, 0x0c,
	0xc9, 0xcf, 0x99, 0x21, 0x2e, 0x71, 0xd7, 0xb7, 0xa8, 0x50, 0x3e, 0x97, 0x38, 0x37, 0x50, 0x07,
	0x5a, 0xa1, 0xe9, 0x50, 0xe6, 0xfe, 0x43, 0x85, 0xf6, 0xeb, 0xa4, 0xb0, 0x79, 0x95, 0xc3, 0x80,
	0x69, 0x0d, 0xb1, 0x29, 0xf8, 0x11, 0x7f, 0x0b, 0xb0, 0xa4, 0x84, 0x5a, 0x50, 0x33, 0xfa, 0x64,
	0xac, 0x56, 0xb8, 0x22, 0x26, 0xfd, 0xa9, 0xd1, 0xbf, 0x52, 0x25, 0x3e, 0xf8, 0xe3, 0x9e, 0x71,
	0x79, 0xad, 0x56, 0xb9, 0x18, 0x7a, 0xa3, 0x91, 0x2a, 0xe3, 0x01, 0x28, 0x57, 0xd4, 0xa3, 0x31,
	0x4d, 0x0b, 0xf1, 0xe2, 0x9d, 0x86, 0x8f, 0x41, 0xe9, 0xd9, 0x36, 0xa1, 0x2c, 0x0c, 0x7c, 0x46,
	0x91, 0x06, 0x4d, 0xb6, 0xb0, 0x2c, 0xca, 0x58, 0xd6, 0xb0, 0xdc, 0xc4, 0x5f, 0xc1, 0x5e, 0x9a,
	0x71, 0x0b, 0xdf, 0xcf, 0x61, 0xf7, 0x9a, 0x7a, 0x5e, 0x50, 0xb8, 0xaa, 0x20, 0xcf, 0x99, 0x93,
	0x8f, 0xd7, 0x9c, 0x39, 0xb8, 0x09, 0xf5, 0xfe, 0x3c, 0x8c, 0x93, 0xee, 0xdb, 0x06, 0xec, 0x18,
	0x19, 0x51, 0xbe, 0x10, 0xd0, 0x4f, 0xd0, 0x1c, 0xd0, 0x78, 0x14, 0x38, 0x0c, 0x7d, 0xb6, 0xae,
	0x97, 0x9d, 0x0d, 0x1b, 0x0c, 0x57, 0x4e, 0x25, 0xf4, 0x33, 0x28, 0x03, 0x1a, 0x17, 0xbb, 0x62,
	0x3d, 0xe0, 0xa6, 0x19, 0x14, 0x88, 0x86, 0x28, 0x58, 0x81, 0xb8, 0x29, 0xa6, 0x83, 0xcb, 0x84,
	0xb3, 0xac, 0x3a, 0xae, 0x20, 0x02, 0xd0, 0xb3, 0xed, 0xfc, 0xb7, 0x61, 0xc3, 0x97, 0x6d, 0x89,
	0x79, 0x07, 0xea, 0x0a, 0xd3, 0x54, 0x8e, 0x5f, 0x6c, 0xa0, 0x2b, 0xbc, 0xb6, 0xc4, 0xff, 0x03,
	0x3e, 0x58, 0x72, 0x4e, 0xe1, 0x8f, 0xd6, 0x13, 0x7f, 0x0e, 0xfa, 0xef, 0xf9, 0xbc, 0x6d, 0xd9,
	0xbc, 0x2f, 0x4b, 0x5e, 0x9f, 0x0e, 0xed, 0x2a, 0x70, 0x46, 0xea, 0xd5, 0x80, 0x09, 0xb4, 0xd3,
	0xbb, 0x9b, 0x60, 0xf6, 0x5a, 0x98, 0x63, 0xa8, 0x0b, 0x25, 0x95, 0xe2, 0x09, 0x11, 0x75, 0xca,
	0xda, 0xfa, 0x44, 0x85, 0xb8, 0x32, 0x6b, 0x88, 0xbf, 0x49, 0x67, 0xef, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xf8, 0x9a, 0xc7, 0xc9, 0x3d, 0x09, 0x00, 0x00,
}