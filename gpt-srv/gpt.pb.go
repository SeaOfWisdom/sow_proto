// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.2
// source: gpt.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextRequest) Reset() {
	*x = TextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextRequest) ProtoMessage() {}

func (x *TextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextRequest.ProtoReflect.Descriptor instead.
func (*TextRequest) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{0}
}

func (x *TextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type RepeatedTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *RepeatedTextRequest) Reset() {
	*x = RepeatedTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedTextRequest) ProtoMessage() {}

func (x *RepeatedTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedTextRequest.ProtoReflect.Descriptor instead.
func (*RepeatedTextRequest) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{1}
}

func (x *RepeatedTextRequest) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

type DetectLanguageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang  string `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DetectLanguageResponse) Reset() {
	*x = DetectLanguageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectLanguageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectLanguageResponse) ProtoMessage() {}

func (x *DetectLanguageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectLanguageResponse.ProtoReflect.Descriptor instead.
func (*DetectLanguageResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{2}
}

func (x *DetectLanguageResponse) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *DetectLanguageResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DetectLanguageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ExtractSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Authors  []string          `protobuf:"bytes,2,rep,name=authors,proto3" json:"authors,omitempty"`
	Abstract string            `protobuf:"bytes,3,opt,name=abstract,proto3" json:"abstract,omitempty"`
	Keywords []string          `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Main     map[string]string `protobuf:"bytes,5,rep,name=main,proto3" json:"main,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExtractSectionResponse) Reset() {
	*x = ExtractSectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractSectionResponse) ProtoMessage() {}

func (x *ExtractSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractSectionResponse.ProtoReflect.Descriptor instead.
func (*ExtractSectionResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{3}
}

func (x *ExtractSectionResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ExtractSectionResponse) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *ExtractSectionResponse) GetAbstract() string {
	if x != nil {
		return x.Abstract
	}
	return ""
}

func (x *ExtractSectionResponse) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *ExtractSectionResponse) GetMain() map[string]string {
	if x != nil {
		return x.Main
	}
	return nil
}

type ExtractDiplomaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackText string `protobuf:"bytes,1,opt,name=back_text,json=backText,proto3" json:"back_text,omitempty"`
}

func (x *ExtractDiplomaRequest) Reset() {
	*x = ExtractDiplomaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractDiplomaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractDiplomaRequest) ProtoMessage() {}

func (x *ExtractDiplomaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractDiplomaRequest.ProtoReflect.Descriptor instead.
func (*ExtractDiplomaRequest) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{4}
}

func (x *ExtractDiplomaRequest) GetBackText() string {
	if x != nil {
		return x.BackText
	}
	return ""
}

type ExtractDiplomaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Back *DiplomaBackResponse `protobuf:"bytes,1,opt,name=back,proto3" json:"back,omitempty"`
}

func (x *ExtractDiplomaResponse) Reset() {
	*x = ExtractDiplomaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractDiplomaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractDiplomaResponse) ProtoMessage() {}

func (x *ExtractDiplomaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractDiplomaResponse.ProtoReflect.Descriptor instead.
func (*ExtractDiplomaResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{5}
}

func (x *ExtractDiplomaResponse) GetBack() *DiplomaBackResponse {
	if x != nil {
		return x.Back
	}
	return nil
}

type DiplomaBackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Series string                 `protobuf:"bytes,1,opt,name=series,proto3" json:"series,omitempty"`
	Number string                 `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Date   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *DiplomaBackResponse) Reset() {
	*x = DiplomaBackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiplomaBackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiplomaBackResponse) ProtoMessage() {}

func (x *DiplomaBackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiplomaBackResponse.ProtoReflect.Descriptor instead.
func (*DiplomaBackResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{6}
}

func (x *DiplomaBackResponse) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

func (x *DiplomaBackResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *DiplomaBackResponse) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type FilterKeyWordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyWords []string `protobuf:"bytes,1,rep,name=key_words,json=keyWords,proto3" json:"key_words,omitempty"`
}

func (x *FilterKeyWordsResponse) Reset() {
	*x = FilterKeyWordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterKeyWordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterKeyWordsResponse) ProtoMessage() {}

func (x *FilterKeyWordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterKeyWordsResponse.ProtoReflect.Descriptor instead.
func (*FilterKeyWordsResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{7}
}

func (x *FilterKeyWordsResponse) GetKeyWords() []string {
	if x != nil {
		return x.KeyWords
	}
	return nil
}

type ReformattedReferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	References string `protobuf:"bytes,1,opt,name=references,proto3" json:"references,omitempty"`
}

func (x *ReformattedReferenceResponse) Reset() {
	*x = ReformattedReferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReformattedReferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReformattedReferenceResponse) ProtoMessage() {}

func (x *ReformattedReferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReformattedReferenceResponse.ProtoReflect.Descriptor instead.
func (*ReformattedReferenceResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{8}
}

func (x *ReformattedReferenceResponse) GetReferences() string {
	if x != nil {
		return x.References
	}
	return ""
}

type ScientificCriteriaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperText string `protobuf:"bytes,1,opt,name=paper_text,json=paperText,proto3" json:"paper_text,omitempty"`
}

func (x *ScientificCriteriaRequest) Reset() {
	*x = ScientificCriteriaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScientificCriteriaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScientificCriteriaRequest) ProtoMessage() {}

func (x *ScientificCriteriaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScientificCriteriaRequest.ProtoReflect.Descriptor instead.
func (*ScientificCriteriaRequest) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{9}
}

func (x *ScientificCriteriaRequest) GetPaperText() string {
	if x != nil {
		return x.PaperText
	}
	return ""
}

type ScientificReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemRelevance       int64  `protobuf:"varint,1,opt,name=problem_relevance,json=problemRelevance,proto3" json:"problem_relevance,omitempty"`
	ArticleQuality         int64  `protobuf:"varint,2,opt,name=article_quality,json=articleQuality,proto3" json:"article_quality,omitempty"`
	MethodsAppropriateness int64  `protobuf:"varint,3,opt,name=methods_appropriateness,json=methodsAppropriateness,proto3" json:"methods_appropriateness,omitempty"`
	ScientificNovelty      int64  `protobuf:"varint,4,opt,name=scientific_novelty,json=scientificNovelty,proto3" json:"scientific_novelty,omitempty"`
	LiteratureRelevance    int64  `protobuf:"varint,5,opt,name=literature_relevance,json=literatureRelevance,proto3" json:"literature_relevance,omitempty"`
	DetailedConclusion     string `protobuf:"bytes,6,opt,name=detailed_conclusion,json=detailedConclusion,proto3" json:"detailed_conclusion,omitempty"`
	Decision               string `protobuf:"bytes,7,opt,name=decision,proto3" json:"decision,omitempty"`
}

func (x *ScientificReviewResponse) Reset() {
	*x = ScientificReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScientificReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScientificReviewResponse) ProtoMessage() {}

func (x *ScientificReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScientificReviewResponse.ProtoReflect.Descriptor instead.
func (*ScientificReviewResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{10}
}

func (x *ScientificReviewResponse) GetProblemRelevance() int64 {
	if x != nil {
		return x.ProblemRelevance
	}
	return 0
}

func (x *ScientificReviewResponse) GetArticleQuality() int64 {
	if x != nil {
		return x.ArticleQuality
	}
	return 0
}

func (x *ScientificReviewResponse) GetMethodsAppropriateness() int64 {
	if x != nil {
		return x.MethodsAppropriateness
	}
	return 0
}

func (x *ScientificReviewResponse) GetScientificNovelty() int64 {
	if x != nil {
		return x.ScientificNovelty
	}
	return 0
}

func (x *ScientificReviewResponse) GetLiteratureRelevance() int64 {
	if x != nil {
		return x.LiteratureRelevance
	}
	return 0
}

func (x *ScientificReviewResponse) GetDetailedConclusion() string {
	if x != nil {
		return x.DetailedConclusion
	}
	return ""
}

func (x *ScientificReviewResponse) GetDecision() string {
	if x != nil {
		return x.Decision
	}
	return ""
}

var File_gpt_proto protoreflect.FileDescriptor

var file_gpt_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x70, 0x2e,
	0x67, 0x70, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x29, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x56, 0x0a, 0x16, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xf7, 0x01, 0x0a, 0x16, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x3c, 0x0a,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x70,
	0x2e, 0x67, 0x70, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x37, 0x0a, 0x09, 0x4d,
	0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x34, 0x0a, 0x15, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44,
	0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x22, 0x49, 0x0a, 0x16, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x70, 0x2e, 0x67, 0x70, 0x74, 0x2e, 0x44, 0x69, 0x70, 0x6c,
	0x6f, 0x6d, 0x61, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x75, 0x0a, 0x13, 0x44, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61,
	0x42, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x16,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x1c, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x19, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x70, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x22,
	0xd8, 0x02, 0x0a, 0x18, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x52, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x37, 0x0a, 0x17, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x5f, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x70, 0x72, 0x69, 0x61, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x16, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x70, 0x72, 0x69, 0x61, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x73,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x14, 0x6c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x89, 0x03, 0x0a, 0x0a, 0x47,
	0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x70,
	0x2e, 0x67, 0x70, 0x74, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x70, 0x2e, 0x67, 0x70, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x70, 0x2e, 0x67, 0x70, 0x74, 0x2e, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x70, 0x2e, 0x67, 0x70, 0x74,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x44, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x12, 0x1d, 0x2e, 0x70, 0x70, 0x2e, 0x67,
	0x70, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x69, 0x70, 0x6c, 0x6f, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x70, 0x2e, 0x67, 0x70,
	0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x70, 0x2e,
	0x67, 0x70, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x70, 0x2e, 0x67, 0x70, 0x74,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x70,
	0x2e, 0x67, 0x70, 0x74, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x70, 0x2e, 0x67, 0x70, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gpt_proto_rawDescOnce sync.Once
	file_gpt_proto_rawDescData = file_gpt_proto_rawDesc
)

func file_gpt_proto_rawDescGZIP() []byte {
	file_gpt_proto_rawDescOnce.Do(func() {
		file_gpt_proto_rawDescData = protoimpl.X.CompressGZIP(file_gpt_proto_rawDescData)
	})
	return file_gpt_proto_rawDescData
}

var file_gpt_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_gpt_proto_goTypes = []interface{}{
	(*TextRequest)(nil),                  // 0: pp.gpt.TextRequest
	(*RepeatedTextRequest)(nil),          // 1: pp.gpt.RepeatedTextRequest
	(*DetectLanguageResponse)(nil),       // 2: pp.gpt.DetectLanguageResponse
	(*ExtractSectionResponse)(nil),       // 3: pp.gpt.ExtractSectionResponse
	(*ExtractDiplomaRequest)(nil),        // 4: pp.gpt.ExtractDiplomaRequest
	(*ExtractDiplomaResponse)(nil),       // 5: pp.gpt.ExtractDiplomaResponse
	(*DiplomaBackResponse)(nil),          // 6: pp.gpt.DiplomaBackResponse
	(*FilterKeyWordsResponse)(nil),       // 7: pp.gpt.FilterKeyWordsResponse
	(*ReformattedReferenceResponse)(nil), // 8: pp.gpt.ReformattedReferenceResponse
	(*ScientificCriteriaRequest)(nil),    // 9: pp.gpt.ScientificCriteriaRequest
	(*ScientificReviewResponse)(nil),     // 10: pp.gpt.ScientificReviewResponse
	nil,                                  // 11: pp.gpt.ExtractSectionResponse.MainEntry
	(*timestamppb.Timestamp)(nil),        // 12: google.protobuf.Timestamp
}
var file_gpt_proto_depIdxs = []int32{
	11, // 0: pp.gpt.ExtractSectionResponse.main:type_name -> pp.gpt.ExtractSectionResponse.MainEntry
	6,  // 1: pp.gpt.ExtractDiplomaResponse.back:type_name -> pp.gpt.DiplomaBackResponse
	12, // 2: pp.gpt.DiplomaBackResponse.date:type_name -> google.protobuf.Timestamp
	0,  // 3: pp.gpt.GptService.DetectLanguage:input_type -> pp.gpt.TextRequest
	0,  // 4: pp.gpt.GptService.ExtractSection:input_type -> pp.gpt.TextRequest
	4,  // 5: pp.gpt.GptService.ExtractDiploma:input_type -> pp.gpt.ExtractDiplomaRequest
	1,  // 6: pp.gpt.GptService.FilterKeyWords:input_type -> pp.gpt.RepeatedTextRequest
	0,  // 7: pp.gpt.GptService.FormatReferences:input_type -> pp.gpt.TextRequest
	2,  // 8: pp.gpt.GptService.DetectLanguage:output_type -> pp.gpt.DetectLanguageResponse
	3,  // 9: pp.gpt.GptService.ExtractSection:output_type -> pp.gpt.ExtractSectionResponse
	5,  // 10: pp.gpt.GptService.ExtractDiploma:output_type -> pp.gpt.ExtractDiplomaResponse
	7,  // 11: pp.gpt.GptService.FilterKeyWords:output_type -> pp.gpt.FilterKeyWordsResponse
	8,  // 12: pp.gpt.GptService.FormatReferences:output_type -> pp.gpt.ReformattedReferenceResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_gpt_proto_init() }
func file_gpt_proto_init() {
	if File_gpt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gpt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextRequest); i {
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
		file_gpt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedTextRequest); i {
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
		file_gpt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectLanguageResponse); i {
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
		file_gpt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractSectionResponse); i {
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
		file_gpt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractDiplomaRequest); i {
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
		file_gpt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractDiplomaResponse); i {
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
		file_gpt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiplomaBackResponse); i {
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
		file_gpt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterKeyWordsResponse); i {
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
		file_gpt_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReformattedReferenceResponse); i {
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
		file_gpt_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScientificCriteriaRequest); i {
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
		file_gpt_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScientificReviewResponse); i {
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
			RawDescriptor: file_gpt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gpt_proto_goTypes,
		DependencyIndexes: file_gpt_proto_depIdxs,
		MessageInfos:      file_gpt_proto_msgTypes,
	}.Build()
	File_gpt_proto = out.File
	file_gpt_proto_rawDesc = nil
	file_gpt_proto_goTypes = nil
	file_gpt_proto_depIdxs = nil
}
