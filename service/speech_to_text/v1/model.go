// Package speech_to_text code generated by oapi sdk gen
package speech_to_text

import (
	"github.com/larksuite/oapi-sdk-go/core"
)

/**生成枚举值 **/

/**生成数据类型 **/

type FileConfig struct {
	FileId     *string `json:"file_id,omitempty"`
	Format     *string `json:"format,omitempty"`
	EngineType *string `json:"engine_type,omitempty"`
}

type Speech struct {
	Speech    *string `json:"speech,omitempty"`
	SpeechKey *string `json:"speech_key,omitempty"`
}

type StreamConfig struct {
	StreamId   *string `json:"stream_id,omitempty"`
	SequenceId *int    `json:"sequence_id,omitempty"`
	Action     *int    `json:"action,omitempty"`
	Format     *string `json:"format,omitempty"`
	EngineType *string `json:"engine_type,omitempty"`
}

/**生成请求和响应结果类型，以及请求对象的Builder构造器 **/

type FileRecognizeSpeechReqBodyBuilder struct {
	speech     *Speech
	speechFlag bool
	config     *FileConfig
	configFlag bool
}

// 生成body的New构造器
func NewFileRecognizeSpeechReqBodyBuilder() *FileRecognizeSpeechReqBodyBuilder {
	builder := &FileRecognizeSpeechReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *FileRecognizeSpeechReqBodyBuilder) Speech(speech *Speech) *FileRecognizeSpeechReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}
func (builder *FileRecognizeSpeechReqBodyBuilder) Config(config *FileConfig) *FileRecognizeSpeechReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *FileRecognizeSpeechReqBodyBuilder) Build() *FileRecognizeSpeechReqBody {
	req := &FileRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech

	}
	if builder.configFlag {
		req.Config = builder.config

	}
	return req
}

/**上传文件path开始**/
type FileRecognizeSpeechPathReqBodyBuilder struct {
	speech     *Speech
	speechFlag bool
	config     *FileConfig
	configFlag bool
}

func NewFileRecognizeSpeechPathReqBodyBuilder() *FileRecognizeSpeechPathReqBodyBuilder {
	builder := &FileRecognizeSpeechPathReqBodyBuilder{}
	return builder
}
func (builder *FileRecognizeSpeechPathReqBodyBuilder) Speech(speech *Speech) *FileRecognizeSpeechPathReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}
func (builder *FileRecognizeSpeechPathReqBodyBuilder) Config(config *FileConfig) *FileRecognizeSpeechPathReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

func (builder *FileRecognizeSpeechPathReqBodyBuilder) Build() (*FileRecognizeSpeechReqBody, error) {
	req := &FileRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech
	}
	if builder.configFlag {
		req.Config = builder.config
	}
	return req, nil
}

/**上传文件path结束**/

/*1.4 生成请求的builder结构体*/
type FileRecognizeSpeechReqBuilder struct {
	body     *FileRecognizeSpeechReqBody
	bodyFlag bool
}

// 生成请求的New构造器
func NewFileRecognizeSpeechReqBuilder() *FileRecognizeSpeechReqBuilder {
	builder := &FileRecognizeSpeechReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *FileRecognizeSpeechReqBuilder) Body(body *FileRecognizeSpeechReqBody) *FileRecognizeSpeechReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *FileRecognizeSpeechReqBuilder) Build() *FileRecognizeSpeechReq {
	req := &FileRecognizeSpeechReq{}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type FileRecognizeSpeechReqBody struct {
	Speech *Speech     `json:"speech,omitempty"`
	Config *FileConfig `json:"config,omitempty"`
}

type FileRecognizeSpeechReq struct {
	Body *FileRecognizeSpeechReqBody `body:""`
}

type FileRecognizeSpeechRespData struct {
	RecognitionText *string `json:"recognition_text,omitempty"`
}

type FileRecognizeSpeechResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *FileRecognizeSpeechRespData `json:"data"`
}

func (resp *FileRecognizeSpeechResp) Success() bool {
	return resp.Code == 0
}

type StreamRecognizeSpeechReqBodyBuilder struct {
	speech     *Speech
	speechFlag bool
	config     *StreamConfig
	configFlag bool
}

// 生成body的New构造器
func NewStreamRecognizeSpeechReqBodyBuilder() *StreamRecognizeSpeechReqBodyBuilder {
	builder := &StreamRecognizeSpeechReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *StreamRecognizeSpeechReqBodyBuilder) Speech(speech *Speech) *StreamRecognizeSpeechReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}
func (builder *StreamRecognizeSpeechReqBodyBuilder) Config(config *StreamConfig) *StreamRecognizeSpeechReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *StreamRecognizeSpeechReqBodyBuilder) Build() *StreamRecognizeSpeechReqBody {
	req := &StreamRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech

	}
	if builder.configFlag {
		req.Config = builder.config

	}
	return req
}

/**上传文件path开始**/
type StreamRecognizeSpeechPathReqBodyBuilder struct {
	speech     *Speech
	speechFlag bool
	config     *StreamConfig
	configFlag bool
}

func NewStreamRecognizeSpeechPathReqBodyBuilder() *StreamRecognizeSpeechPathReqBodyBuilder {
	builder := &StreamRecognizeSpeechPathReqBodyBuilder{}
	return builder
}
func (builder *StreamRecognizeSpeechPathReqBodyBuilder) Speech(speech *Speech) *StreamRecognizeSpeechPathReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}
func (builder *StreamRecognizeSpeechPathReqBodyBuilder) Config(config *StreamConfig) *StreamRecognizeSpeechPathReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

func (builder *StreamRecognizeSpeechPathReqBodyBuilder) Build() (*StreamRecognizeSpeechReqBody, error) {
	req := &StreamRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech
	}
	if builder.configFlag {
		req.Config = builder.config
	}
	return req, nil
}

/**上传文件path结束**/

/*1.4 生成请求的builder结构体*/
type StreamRecognizeSpeechReqBuilder struct {
	body     *StreamRecognizeSpeechReqBody
	bodyFlag bool
}

// 生成请求的New构造器
func NewStreamRecognizeSpeechReqBuilder() *StreamRecognizeSpeechReqBuilder {
	builder := &StreamRecognizeSpeechReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *StreamRecognizeSpeechReqBuilder) Body(body *StreamRecognizeSpeechReqBody) *StreamRecognizeSpeechReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *StreamRecognizeSpeechReqBuilder) Build() *StreamRecognizeSpeechReq {
	req := &StreamRecognizeSpeechReq{}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type StreamRecognizeSpeechReqBody struct {
	Speech *Speech       `json:"speech,omitempty"`
	Config *StreamConfig `json:"config,omitempty"`
}

type StreamRecognizeSpeechReq struct {
	Body *StreamRecognizeSpeechReqBody `body:""`
}

type StreamRecognizeSpeechRespData struct {
	StreamId        *string `json:"stream_id,omitempty"`
	SequenceId      *int    `json:"sequence_id,omitempty"`
	RecognitionText *string `json:"recognition_text,omitempty"`
}

type StreamRecognizeSpeechResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *StreamRecognizeSpeechRespData `json:"data"`
}

func (resp *StreamRecognizeSpeechResp) Success() bool {
	return resp.Code == 0
}

/**生成消息事件结构体 **/

/* 生成请求的builder构造器*/
/*1.1 生成body的builder结构体*/
