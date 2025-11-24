// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/chunkify-go/internal/encoding/json"
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

type Auto string       // Always "auto"
type Aws string        // Always "aws"
type Chunkify string   // Always "chunkify"
type Cloudflare string // Always "cloudflare"
type HlsAv1 string     // Always "hls_av1"
type HlsH264 string    // Always "hls_h264"
type HlsH265 string    // Always "hls_h265"
type Jpg string        // Always "jpg"
type MP4Av1 string     // Always "mp4_av1"
type MP4H264 string    // Always "mp4_h264"
type MP4H265 string    // Always "mp4_h265"
type Success string    // Always "success"
type WebmVp9 string    // Always "webm_vp9"

func (c Auto) Default() Auto             { return "auto" }
func (c Aws) Default() Aws               { return "aws" }
func (c Chunkify) Default() Chunkify     { return "chunkify" }
func (c Cloudflare) Default() Cloudflare { return "cloudflare" }
func (c HlsAv1) Default() HlsAv1         { return "hls_av1" }
func (c HlsH264) Default() HlsH264       { return "hls_h264" }
func (c HlsH265) Default() HlsH265       { return "hls_h265" }
func (c Jpg) Default() Jpg               { return "jpg" }
func (c MP4Av1) Default() MP4Av1         { return "mp4_av1" }
func (c MP4H264) Default() MP4H264       { return "mp4_h264" }
func (c MP4H265) Default() MP4H265       { return "mp4_h265" }
func (c Success) Default() Success       { return "success" }
func (c WebmVp9) Default() WebmVp9       { return "webm_vp9" }

func (c Auto) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Aws) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Chunkify) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Cloudflare) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c HlsAv1) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c HlsH264) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c HlsH265) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Jpg) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c MP4Av1) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c MP4H264) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c MP4H265) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Success) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c WebmVp9) MarshalJSON() ([]byte, error)    { return marshalString(c) }

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
