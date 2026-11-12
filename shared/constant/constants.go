// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/chunkifydev/chunkify-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Auto string                  // Always "auto"
type Aws string                   // Always "aws"
type Caption string               // Always "caption"
type Chunkify string              // Always "chunkify"
type Cloudflare string            // Always "cloudflare"
type CreateCaption string         // Always "create_caption"
type CreateStaticRendition string // Always "create_static_rendition"
type CreateThumbnails string      // Always "create_thumbnails"
type CreateTranscript string      // Always "create_transcript"
type H264 string                  // Always "h264"
type HlsAv1 string                // Always "hls_av1"
type HlsH264 string               // Always "hls_h264"
type HlsH265 string               // Always "hls_h265"
type Jpg string                   // Always "jpg"
type M4aAac string                // Always "m4a_aac"
type MP4 string                   // Always "mp4"
type MP4Av1 string                // Always "mp4_av1"
type MP4H264 string               // Always "mp4_h264"
type MP4H265 string               // Always "mp4_h265"
type S3Compatible string          // Always "s3_compatible"
type Sprite string                // Always "sprite"
type StaticRendition string       // Always "static_rendition"
type Success string               // Always "success"
type Thumbnail string             // Always "thumbnail"
type Transcript string            // Always "transcript"
type TranslateCaption string      // Always "translate_caption"
type WebmVp9 string               // Always "webm_vp9"

func (c Auto) Default() Auto                                   { return "auto" }
func (c Aws) Default() Aws                                     { return "aws" }
func (c Caption) Default() Caption                             { return "caption" }
func (c Chunkify) Default() Chunkify                           { return "chunkify" }
func (c Cloudflare) Default() Cloudflare                       { return "cloudflare" }
func (c CreateCaption) Default() CreateCaption                 { return "create_caption" }
func (c CreateStaticRendition) Default() CreateStaticRendition { return "create_static_rendition" }
func (c CreateThumbnails) Default() CreateThumbnails           { return "create_thumbnails" }
func (c CreateTranscript) Default() CreateTranscript           { return "create_transcript" }
func (c H264) Default() H264                                   { return "h264" }
func (c HlsAv1) Default() HlsAv1                               { return "hls_av1" }
func (c HlsH264) Default() HlsH264                             { return "hls_h264" }
func (c HlsH265) Default() HlsH265                             { return "hls_h265" }
func (c Jpg) Default() Jpg                                     { return "jpg" }
func (c M4aAac) Default() M4aAac                               { return "m4a_aac" }
func (c MP4) Default() MP4                                     { return "mp4" }
func (c MP4Av1) Default() MP4Av1                               { return "mp4_av1" }
func (c MP4H264) Default() MP4H264                             { return "mp4_h264" }
func (c MP4H265) Default() MP4H265                             { return "mp4_h265" }
func (c S3Compatible) Default() S3Compatible                   { return "s3_compatible" }
func (c Sprite) Default() Sprite                               { return "sprite" }
func (c StaticRendition) Default() StaticRendition             { return "static_rendition" }
func (c Success) Default() Success                             { return "success" }
func (c Thumbnail) Default() Thumbnail                         { return "thumbnail" }
func (c Transcript) Default() Transcript                       { return "transcript" }
func (c TranslateCaption) Default() TranslateCaption           { return "translate_caption" }
func (c WebmVp9) Default() WebmVp9                             { return "webm_vp9" }

func (c Auto) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Aws) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Caption) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Chunkify) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Cloudflare) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c CreateCaption) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c CreateStaticRendition) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c CreateThumbnails) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c CreateTranscript) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c H264) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c HlsAv1) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c HlsH264) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c HlsH265) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Jpg) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c M4aAac) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c MP4) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c MP4Av1) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c MP4H264) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c MP4H265) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c S3Compatible) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Sprite) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c StaticRendition) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Success) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Thumbnail) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Transcript) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c TranslateCaption) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c WebmVp9) MarshalJSON() ([]byte, error)               { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
