// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/chunkify-go/internal/apijson"
	"github.com/stainless-sdks/chunkify-go/internal/apiquery"
	"github.com/stainless-sdks/chunkify-go/internal/requestconfig"
	"github.com/stainless-sdks/chunkify-go/option"
	"github.com/stainless-sdks/chunkify-go/packages/pagination"
	"github.com/stainless-sdks/chunkify-go/packages/param"
	"github.com/stainless-sdks/chunkify-go/packages/respjson"
	"github.com/stainless-sdks/chunkify-go/shared"
	"github.com/stainless-sdks/chunkify-go/shared/constant"
)

// JobService contains methods and other services that help with interacting with
// the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options     []option.RequestOption
	Files       JobFileService
	Logs        JobLogService
	Transcoders JobTranscoderService
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r JobService) {
	r = JobService{}
	r.Options = opts
	r.Files = NewJobFileService(opts...)
	r.Logs = NewJobLogService(opts...)
	r.Transcoders = NewJobTranscoderService(opts...)
	return
}

// Create a new video processing job with specified parameters
func (r *JobService) New(ctx context.Context, body JobNewParams, opts ...option.RequestOption) (res *Job, err error) {
	var env JobNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "api/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve details of a specific job
func (r *JobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *Job, err error) {
	var env JobGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve a list of jobs with optional filtering and pagination
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Job], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/jobs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a list of jobs with optional filtering and pagination
func (r *JobService) ListAutoPaging(ctx context.Context, query JobListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Job] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete a job.
func (r *JobService) Delete(ctx context.Context, jobID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Cancel a job.
func (r *JobService) Cancel(ctx context.Context, jobID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/cancel", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type HlsAv1 struct {
	ID constant.HlsAv1 `json:"id,required"`
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 63. Recommended
	// values: 16-35 for high quality, 35-45 for good quality, 45-63 for acceptable
	// quality.
	Crf int64 `json:"crf"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio bool `json:"disable_audio"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo bool `json:"disable_video"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate float64 `json:"framerate"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop int64 `json:"gop"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height int64 `json:"height"`
	// HlsEnc enables encryption for HLS segments when set to true.
	HlsEnc bool `json:"hls_enc"`
	// HlsEncIv specifies the initialization vector for encryption. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncIv string `json:"hls_enc_iv"`
	// HlsEncKey specifies the encryption key for HLS segments. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncKey string `json:"hls_enc_key"`
	// HlsEncKeyUrl specifies the URL where clients can fetch the encryption key.
	// Required when HlsEnc is true.
	HlsEncKeyURL string `json:"hls_enc_key_url"`
	// HlsSegmentType specifies the type of HLS segments. Valid values:
	//
	// - mpegts: Traditional MPEG-TS segments, better compatibility
	// - fmp4: Fragmented MP4 segments, better efficiency
	//
	// Any of "mpegts", "fmp4".
	HlsSegmentType HlsAv1HlsSegmentType `json:"hls_segment_type"`
	// HlsTime specifies the duration of each HLS segment in seconds. Range: 1 to 10.
	// Shorter segments provide faster startup but more overhead, longer segments are
	// more efficient.
	HlsTime int64 `json:"hls_time"`
	// Level specifies the AV1 profile level. Valid values: 30-31 (main), 41 (main10).
	// Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  int64  `json:"minrate"`
	Movflags string `json:"movflags"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt HlsAv1Pixfmt `json:"pixfmt"`
	// Preset controls the encoding efficiency and processing intensity. Lower presets
	// use more optimization features, creating smaller files with better quality but
	// requiring more compute time. Higher presets encode faster but produce larger
	// files.
	//
	// Preset ranges:
	//
	// - 6-7: Fast encoding for real-time applications (smaller files)
	// - 8-10: Balanced efficiency and speed for general use
	// - 11-13: Fastest encoding for real-time applications (larger files)
	//
	// Any of "6", "7", "8", "9", "10", "11", "12", "13".
	Preset HlsAv1Preset `json:"preset"`
	// Profilev specifies the AV1 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev HlsAv1Profilev `json:"profilev"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek int64 `json:"seek"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AudioBitrate   respjson.Field
		Bufsize        respjson.Field
		Channels       respjson.Field
		Crf            respjson.Field
		DisableAudio   respjson.Field
		DisableVideo   respjson.Field
		Duration       respjson.Field
		Framerate      respjson.Field
		Gop            respjson.Field
		Height         respjson.Field
		HlsEnc         respjson.Field
		HlsEncIv       respjson.Field
		HlsEncKey      respjson.Field
		HlsEncKeyURL   respjson.Field
		HlsSegmentType respjson.Field
		HlsTime        respjson.Field
		Level          respjson.Field
		Maxrate        respjson.Field
		Minrate        respjson.Field
		Movflags       respjson.Field
		Pixfmt         respjson.Field
		Preset         respjson.Field
		Profilev       respjson.Field
		Seek           respjson.Field
		VideoBitrate   respjson.Field
		Width          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HlsAv1) RawJSON() string { return r.JSON.raw }
func (r *HlsAv1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HlsAv1 to a HlsAv1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HlsAv1Param.Overrides()
func (r HlsAv1) ToParam() HlsAv1Param {
	return param.Override[HlsAv1Param](json.RawMessage(r.RawJSON()))
}

// HlsSegmentType specifies the type of HLS segments. Valid values:
//
// - mpegts: Traditional MPEG-TS segments, better compatibility
// - fmp4: Fragmented MP4 segments, better efficiency
type HlsAv1HlsSegmentType string

const (
	HlsAv1HlsSegmentTypeMpegts HlsAv1HlsSegmentType = "mpegts"
	HlsAv1HlsSegmentTypeFmp4   HlsAv1HlsSegmentType = "fmp4"
)

// PixFmt specifies the pixel format. Valid value: yuv420p
type HlsAv1Pixfmt string

const (
	HlsAv1PixfmtYuv410p     HlsAv1Pixfmt = "yuv410p"
	HlsAv1PixfmtYuv411p     HlsAv1Pixfmt = "yuv411p"
	HlsAv1PixfmtYuv420p     HlsAv1Pixfmt = "yuv420p"
	HlsAv1PixfmtYuv422p     HlsAv1Pixfmt = "yuv422p"
	HlsAv1PixfmtYuv440p     HlsAv1Pixfmt = "yuv440p"
	HlsAv1PixfmtYuv444p     HlsAv1Pixfmt = "yuv444p"
	HlsAv1PixfmtYuvJ411p    HlsAv1Pixfmt = "yuvJ411p"
	HlsAv1PixfmtYuvJ420p    HlsAv1Pixfmt = "yuvJ420p"
	HlsAv1PixfmtYuvJ422p    HlsAv1Pixfmt = "yuvJ422p"
	HlsAv1PixfmtYuvJ440p    HlsAv1Pixfmt = "yuvJ440p"
	HlsAv1PixfmtYuvJ444p    HlsAv1Pixfmt = "yuvJ444p"
	HlsAv1PixfmtYuv420p10le HlsAv1Pixfmt = "yuv420p10le"
	HlsAv1PixfmtYuv422p10le HlsAv1Pixfmt = "yuv422p10le"
	HlsAv1PixfmtYuv440p10le HlsAv1Pixfmt = "yuv440p10le"
	HlsAv1PixfmtYuv444p10le HlsAv1Pixfmt = "yuv444p10le"
	HlsAv1PixfmtYuv420p12le HlsAv1Pixfmt = "yuv420p12le"
	HlsAv1PixfmtYuv422p12le HlsAv1Pixfmt = "yuv422p12le"
	HlsAv1PixfmtYuv440p12le HlsAv1Pixfmt = "yuv440p12le"
	HlsAv1PixfmtYuv444p12le HlsAv1Pixfmt = "yuv444p12le"
	HlsAv1PixfmtYuv420p10be HlsAv1Pixfmt = "yuv420p10be"
	HlsAv1PixfmtYuv422p10be HlsAv1Pixfmt = "yuv422p10be"
	HlsAv1PixfmtYuv440p10be HlsAv1Pixfmt = "yuv440p10be"
	HlsAv1PixfmtYuv444p10be HlsAv1Pixfmt = "yuv444p10be"
	HlsAv1PixfmtYuv420p12be HlsAv1Pixfmt = "yuv420p12be"
	HlsAv1PixfmtYuv422p12be HlsAv1Pixfmt = "yuv422p12be"
	HlsAv1PixfmtYuv440p12be HlsAv1Pixfmt = "yuv440p12be"
	HlsAv1PixfmtYuv444p12be HlsAv1Pixfmt = "yuv444p12be"
)

// Preset controls the encoding efficiency and processing intensity. Lower presets
// use more optimization features, creating smaller files with better quality but
// requiring more compute time. Higher presets encode faster but produce larger
// files.
//
// Preset ranges:
//
// - 6-7: Fast encoding for real-time applications (smaller files)
// - 8-10: Balanced efficiency and speed for general use
// - 11-13: Fastest encoding for real-time applications (larger files)
type HlsAv1Preset string

const (
	HlsAv1Preset6  HlsAv1Preset = "6"
	HlsAv1Preset7  HlsAv1Preset = "7"
	HlsAv1Preset8  HlsAv1Preset = "8"
	HlsAv1Preset9  HlsAv1Preset = "9"
	HlsAv1Preset10 HlsAv1Preset = "10"
	HlsAv1Preset11 HlsAv1Preset = "11"
	HlsAv1Preset12 HlsAv1Preset = "12"
	HlsAv1Preset13 HlsAv1Preset = "13"
)

// Profilev specifies the AV1 profile. Valid values:
//
// - main: Main profile, good for most applications
// - main10: Main 10-bit profile, supports 10-bit color
// - mainstillpicture: Still picture profile, optimized for single images
type HlsAv1Profilev string

const (
	HlsAv1ProfilevMain             HlsAv1Profilev = "main"
	HlsAv1ProfilevMain10           HlsAv1Profilev = "main10"
	HlsAv1ProfilevMainstillpicture HlsAv1Profilev = "mainstillpicture"
)

// The property ID is required.
type HlsAv1Param struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 63. Recommended
	// values: 16-35 for high quality, 35-45 for good quality, 45-63 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio param.Opt[bool] `json:"disable_audio,omitzero"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo param.Opt[bool] `json:"disable_video,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate param.Opt[float64] `json:"framerate,omitzero"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop param.Opt[int64] `json:"gop,omitzero"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height param.Opt[int64] `json:"height,omitzero"`
	// HlsEnc enables encryption for HLS segments when set to true.
	HlsEnc param.Opt[bool] `json:"hls_enc,omitzero"`
	// HlsEncIv specifies the initialization vector for encryption. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncIv param.Opt[string] `json:"hls_enc_iv,omitzero"`
	// HlsEncKey specifies the encryption key for HLS segments. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncKey param.Opt[string] `json:"hls_enc_key,omitzero"`
	// HlsEncKeyUrl specifies the URL where clients can fetch the encryption key.
	// Required when HlsEnc is true.
	HlsEncKeyURL param.Opt[string] `json:"hls_enc_key_url,omitzero"`
	// HlsTime specifies the duration of each HLS segment in seconds. Range: 1 to 10.
	// Shorter segments provide faster startup but more overhead, longer segments are
	// more efficient.
	HlsTime param.Opt[int64] `json:"hls_time,omitzero"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate param.Opt[int64] `json:"maxrate,omitzero"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  param.Opt[int64]  `json:"minrate,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek param.Opt[int64] `json:"seek,omitzero"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate param.Opt[int64] `json:"video_bitrate,omitzero"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width param.Opt[int64] `json:"width,omitzero"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels,omitzero"`
	// HlsSegmentType specifies the type of HLS segments. Valid values:
	//
	// - mpegts: Traditional MPEG-TS segments, better compatibility
	// - fmp4: Fragmented MP4 segments, better efficiency
	//
	// Any of "mpegts", "fmp4".
	HlsSegmentType HlsAv1HlsSegmentType `json:"hls_segment_type,omitzero"`
	// Level specifies the AV1 profile level. Valid values: 30-31 (main), 41 (main10).
	// Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level,omitzero"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt HlsAv1Pixfmt `json:"pixfmt,omitzero"`
	// Preset controls the encoding efficiency and processing intensity. Lower presets
	// use more optimization features, creating smaller files with better quality but
	// requiring more compute time. Higher presets encode faster but produce larger
	// files.
	//
	// Preset ranges:
	//
	// - 6-7: Fast encoding for real-time applications (smaller files)
	// - 8-10: Balanced efficiency and speed for general use
	// - 11-13: Fastest encoding for real-time applications (larger files)
	//
	// Any of "6", "7", "8", "9", "10", "11", "12", "13".
	Preset HlsAv1Preset `json:"preset,omitzero"`
	// Profilev specifies the AV1 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev HlsAv1Profilev `json:"profilev,omitzero"`
	// This field can be elided, and will marshal its zero value as "hls_av1".
	ID constant.HlsAv1 `json:"id,required"`
	paramObj
}

func (r HlsAv1Param) MarshalJSON() (data []byte, err error) {
	type shadow HlsAv1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HlsAv1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[HlsAv1Param](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[HlsAv1Param](
		"level", 30, 31, 41,
	)
}

type HlsH264 struct {
	ID constant.HlsH264 `json:"id,required"`
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf int64 `json:"crf"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio bool `json:"disable_audio"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo bool `json:"disable_video"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate float64 `json:"framerate"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop int64 `json:"gop"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height int64 `json:"height"`
	// HlsEnc enables encryption for HLS segments when set to true.
	HlsEnc bool `json:"hls_enc"`
	// HlsEncIv specifies the initialization vector for encryption. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncIv string `json:"hls_enc_iv"`
	// HlsEncKey specifies the encryption key for HLS segments. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncKey string `json:"hls_enc_key"`
	// HlsEncKeyUrl specifies the URL where clients can fetch the encryption key.
	// Required when HlsEnc is true.
	HlsEncKeyURL string `json:"hls_enc_key_url"`
	// HlsSegmentType specifies the type of HLS segments. Valid values:
	//
	// - mpegts: Traditional MPEG-TS segments, better compatibility
	// - fmp4: Fragmented MP4 segments, better efficiency
	//
	// Any of "mpegts", "fmp4".
	HlsSegmentType HlsH264HlsSegmentType `json:"hls_segment_type"`
	// HlsTime specifies the duration of each HLS segment in seconds. Range: 1 to 10.
	// Shorter segments provide faster startup but more overhead, longer segments are
	// more efficient.
	HlsTime int64 `json:"hls_time"`
	// Level specifies the H.264 profile level. Valid values: 10-13 (baseline), 20-22
	// (main), 30-32 (high), 40-42 (high), 50-51 (high). Higher levels support higher
	// resolutions and bitrates but require more processing power.
	//
	// Any of 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51.
	Level int64 `json:"level"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  int64  `json:"minrate"`
	Movflags string `json:"movflags"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt HlsH264Pixfmt `json:"pixfmt"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset HlsH264Preset `json:"preset"`
	// Profilev specifies the H.264 profile. Valid values:
	//
	// - baseline: Basic profile, good for mobile devices
	// - main: Main profile, good for most applications
	// - high: High profile, best quality but requires more processing
	// - high10: High 10-bit profile, supports 10-bit color
	// - high422: High 4:2:2 profile, supports 4:2:2 color sampling
	// - high444: High 4:4:4 profile, supports 4:4:4 color sampling
	//
	// Any of "baseline", "main", "high", "high10", "high422", "high444".
	Profilev HlsH264Profilev `json:"profilev"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek int64 `json:"seek"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width int64 `json:"width"`
	// X264KeyInt specifies the maximum number of frames between keyframes for H.264
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X264Keyint int64 `json:"x264_keyint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AudioBitrate   respjson.Field
		Bufsize        respjson.Field
		Channels       respjson.Field
		Crf            respjson.Field
		DisableAudio   respjson.Field
		DisableVideo   respjson.Field
		Duration       respjson.Field
		Framerate      respjson.Field
		Gop            respjson.Field
		Height         respjson.Field
		HlsEnc         respjson.Field
		HlsEncIv       respjson.Field
		HlsEncKey      respjson.Field
		HlsEncKeyURL   respjson.Field
		HlsSegmentType respjson.Field
		HlsTime        respjson.Field
		Level          respjson.Field
		Maxrate        respjson.Field
		Minrate        respjson.Field
		Movflags       respjson.Field
		Pixfmt         respjson.Field
		Preset         respjson.Field
		Profilev       respjson.Field
		Seek           respjson.Field
		VideoBitrate   respjson.Field
		Width          respjson.Field
		X264Keyint     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HlsH264) RawJSON() string { return r.JSON.raw }
func (r *HlsH264) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HlsH264 to a HlsH264Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HlsH264Param.Overrides()
func (r HlsH264) ToParam() HlsH264Param {
	return param.Override[HlsH264Param](json.RawMessage(r.RawJSON()))
}

// HlsSegmentType specifies the type of HLS segments. Valid values:
//
// - mpegts: Traditional MPEG-TS segments, better compatibility
// - fmp4: Fragmented MP4 segments, better efficiency
type HlsH264HlsSegmentType string

const (
	HlsH264HlsSegmentTypeMpegts HlsH264HlsSegmentType = "mpegts"
	HlsH264HlsSegmentTypeFmp4   HlsH264HlsSegmentType = "fmp4"
)

// PixFmt specifies the pixel format. Valid value: yuv420p
type HlsH264Pixfmt string

const (
	HlsH264PixfmtYuv410p     HlsH264Pixfmt = "yuv410p"
	HlsH264PixfmtYuv411p     HlsH264Pixfmt = "yuv411p"
	HlsH264PixfmtYuv420p     HlsH264Pixfmt = "yuv420p"
	HlsH264PixfmtYuv422p     HlsH264Pixfmt = "yuv422p"
	HlsH264PixfmtYuv440p     HlsH264Pixfmt = "yuv440p"
	HlsH264PixfmtYuv444p     HlsH264Pixfmt = "yuv444p"
	HlsH264PixfmtYuvJ411p    HlsH264Pixfmt = "yuvJ411p"
	HlsH264PixfmtYuvJ420p    HlsH264Pixfmt = "yuvJ420p"
	HlsH264PixfmtYuvJ422p    HlsH264Pixfmt = "yuvJ422p"
	HlsH264PixfmtYuvJ440p    HlsH264Pixfmt = "yuvJ440p"
	HlsH264PixfmtYuvJ444p    HlsH264Pixfmt = "yuvJ444p"
	HlsH264PixfmtYuv420p10le HlsH264Pixfmt = "yuv420p10le"
	HlsH264PixfmtYuv422p10le HlsH264Pixfmt = "yuv422p10le"
	HlsH264PixfmtYuv440p10le HlsH264Pixfmt = "yuv440p10le"
	HlsH264PixfmtYuv444p10le HlsH264Pixfmt = "yuv444p10le"
	HlsH264PixfmtYuv420p12le HlsH264Pixfmt = "yuv420p12le"
	HlsH264PixfmtYuv422p12le HlsH264Pixfmt = "yuv422p12le"
	HlsH264PixfmtYuv440p12le HlsH264Pixfmt = "yuv440p12le"
	HlsH264PixfmtYuv444p12le HlsH264Pixfmt = "yuv444p12le"
	HlsH264PixfmtYuv420p10be HlsH264Pixfmt = "yuv420p10be"
	HlsH264PixfmtYuv422p10be HlsH264Pixfmt = "yuv422p10be"
	HlsH264PixfmtYuv440p10be HlsH264Pixfmt = "yuv440p10be"
	HlsH264PixfmtYuv444p10be HlsH264Pixfmt = "yuv444p10be"
	HlsH264PixfmtYuv420p12be HlsH264Pixfmt = "yuv420p12be"
	HlsH264PixfmtYuv422p12be HlsH264Pixfmt = "yuv422p12be"
	HlsH264PixfmtYuv440p12be HlsH264Pixfmt = "yuv440p12be"
	HlsH264PixfmtYuv444p12be HlsH264Pixfmt = "yuv444p12be"
)

// Preset specifies the encoding speed preset. Valid values (from fastest to
// slowest):
//
// - ultrafast: Fastest encoding, lowest quality
// - superfast: Very fast encoding, lower quality
// - veryfast: Fast encoding, moderate quality
// - faster: Faster encoding, good quality
// - fast: Fast encoding, better quality
// - medium: Balanced preset, best quality
type HlsH264Preset string

const (
	HlsH264PresetUltrafast HlsH264Preset = "ultrafast"
	HlsH264PresetSuperfast HlsH264Preset = "superfast"
	HlsH264PresetVeryfast  HlsH264Preset = "veryfast"
	HlsH264PresetFaster    HlsH264Preset = "faster"
	HlsH264PresetFast      HlsH264Preset = "fast"
	HlsH264PresetMedium    HlsH264Preset = "medium"
)

// Profilev specifies the H.264 profile. Valid values:
//
// - baseline: Basic profile, good for mobile devices
// - main: Main profile, good for most applications
// - high: High profile, best quality but requires more processing
// - high10: High 10-bit profile, supports 10-bit color
// - high422: High 4:2:2 profile, supports 4:2:2 color sampling
// - high444: High 4:4:4 profile, supports 4:4:4 color sampling
type HlsH264Profilev string

const (
	HlsH264ProfilevBaseline HlsH264Profilev = "baseline"
	HlsH264ProfilevMain     HlsH264Profilev = "main"
	HlsH264ProfilevHigh     HlsH264Profilev = "high"
	HlsH264ProfilevHigh10   HlsH264Profilev = "high10"
	HlsH264ProfilevHigh422  HlsH264Profilev = "high422"
	HlsH264ProfilevHigh444  HlsH264Profilev = "high444"
)

// The property ID is required.
type HlsH264Param struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio param.Opt[bool] `json:"disable_audio,omitzero"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo param.Opt[bool] `json:"disable_video,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate param.Opt[float64] `json:"framerate,omitzero"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop param.Opt[int64] `json:"gop,omitzero"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height param.Opt[int64] `json:"height,omitzero"`
	// HlsEnc enables encryption for HLS segments when set to true.
	HlsEnc param.Opt[bool] `json:"hls_enc,omitzero"`
	// HlsEncIv specifies the initialization vector for encryption. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncIv param.Opt[string] `json:"hls_enc_iv,omitzero"`
	// HlsEncKey specifies the encryption key for HLS segments. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncKey param.Opt[string] `json:"hls_enc_key,omitzero"`
	// HlsEncKeyUrl specifies the URL where clients can fetch the encryption key.
	// Required when HlsEnc is true.
	HlsEncKeyURL param.Opt[string] `json:"hls_enc_key_url,omitzero"`
	// HlsTime specifies the duration of each HLS segment in seconds. Range: 1 to 10.
	// Shorter segments provide faster startup but more overhead, longer segments are
	// more efficient.
	HlsTime param.Opt[int64] `json:"hls_time,omitzero"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate param.Opt[int64] `json:"maxrate,omitzero"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  param.Opt[int64]  `json:"minrate,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek param.Opt[int64] `json:"seek,omitzero"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate param.Opt[int64] `json:"video_bitrate,omitzero"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width param.Opt[int64] `json:"width,omitzero"`
	// X264KeyInt specifies the maximum number of frames between keyframes for H.264
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X264Keyint param.Opt[int64] `json:"x264_keyint,omitzero"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels,omitzero"`
	// HlsSegmentType specifies the type of HLS segments. Valid values:
	//
	// - mpegts: Traditional MPEG-TS segments, better compatibility
	// - fmp4: Fragmented MP4 segments, better efficiency
	//
	// Any of "mpegts", "fmp4".
	HlsSegmentType HlsH264HlsSegmentType `json:"hls_segment_type,omitzero"`
	// Level specifies the H.264 profile level. Valid values: 10-13 (baseline), 20-22
	// (main), 30-32 (high), 40-42 (high), 50-51 (high). Higher levels support higher
	// resolutions and bitrates but require more processing power.
	//
	// Any of 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51.
	Level int64 `json:"level,omitzero"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt HlsH264Pixfmt `json:"pixfmt,omitzero"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset HlsH264Preset `json:"preset,omitzero"`
	// Profilev specifies the H.264 profile. Valid values:
	//
	// - baseline: Basic profile, good for mobile devices
	// - main: Main profile, good for most applications
	// - high: High profile, best quality but requires more processing
	// - high10: High 10-bit profile, supports 10-bit color
	// - high422: High 4:2:2 profile, supports 4:2:2 color sampling
	// - high444: High 4:4:4 profile, supports 4:4:4 color sampling
	//
	// Any of "baseline", "main", "high", "high10", "high422", "high444".
	Profilev HlsH264Profilev `json:"profilev,omitzero"`
	// This field can be elided, and will marshal its zero value as "hls_h264".
	ID constant.HlsH264 `json:"id,required"`
	paramObj
}

func (r HlsH264Param) MarshalJSON() (data []byte, err error) {
	type shadow HlsH264Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HlsH264Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[HlsH264Param](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[HlsH264Param](
		"level", 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51,
	)
}

type HlsH265 struct {
	ID constant.HlsH265 `json:"id,required"`
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf int64 `json:"crf"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio bool `json:"disable_audio"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo bool `json:"disable_video"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate float64 `json:"framerate"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop int64 `json:"gop"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height int64 `json:"height"`
	// HlsEnc enables encryption for HLS segments when set to true.
	HlsEnc bool `json:"hls_enc"`
	// HlsEncIv specifies the initialization vector for encryption. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncIv string `json:"hls_enc_iv"`
	// HlsEncKey specifies the encryption key for HLS segments. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncKey string `json:"hls_enc_key"`
	// HlsEncKeyUrl specifies the URL where clients can fetch the encryption key.
	// Required when HlsEnc is true.
	HlsEncKeyURL string `json:"hls_enc_key_url"`
	// HlsSegmentType specifies the type of HLS segments. Valid values:
	//
	// - mpegts: Traditional MPEG-TS segments, better compatibility
	// - fmp4: Fragmented MP4 segments, better efficiency
	//
	// Any of "mpegts", "fmp4".
	HlsSegmentType HlsH265HlsSegmentType `json:"hls_segment_type"`
	// HlsTime specifies the duration of each HLS segment in seconds. Range: 1 to 10.
	// Shorter segments provide faster startup but more overhead, longer segments are
	// more efficient.
	HlsTime int64 `json:"hls_time"`
	// Level specifies the H.265 profile level. Valid values: 30-31 (main), 41
	// (main10). Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  int64  `json:"minrate"`
	Movflags string `json:"movflags"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt HlsH265Pixfmt `json:"pixfmt"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset HlsH265Preset `json:"preset"`
	// Profilev specifies the H.265 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev HlsH265Profilev `json:"profilev"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek int64 `json:"seek"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width int64 `json:"width"`
	// X265KeyInt specifies the maximum number of frames between keyframes for H.265
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X265Keyint int64 `json:"x265_keyint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AudioBitrate   respjson.Field
		Bufsize        respjson.Field
		Channels       respjson.Field
		Crf            respjson.Field
		DisableAudio   respjson.Field
		DisableVideo   respjson.Field
		Duration       respjson.Field
		Framerate      respjson.Field
		Gop            respjson.Field
		Height         respjson.Field
		HlsEnc         respjson.Field
		HlsEncIv       respjson.Field
		HlsEncKey      respjson.Field
		HlsEncKeyURL   respjson.Field
		HlsSegmentType respjson.Field
		HlsTime        respjson.Field
		Level          respjson.Field
		Maxrate        respjson.Field
		Minrate        respjson.Field
		Movflags       respjson.Field
		Pixfmt         respjson.Field
		Preset         respjson.Field
		Profilev       respjson.Field
		Seek           respjson.Field
		VideoBitrate   respjson.Field
		Width          respjson.Field
		X265Keyint     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HlsH265) RawJSON() string { return r.JSON.raw }
func (r *HlsH265) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HlsH265 to a HlsH265Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HlsH265Param.Overrides()
func (r HlsH265) ToParam() HlsH265Param {
	return param.Override[HlsH265Param](json.RawMessage(r.RawJSON()))
}

// HlsSegmentType specifies the type of HLS segments. Valid values:
//
// - mpegts: Traditional MPEG-TS segments, better compatibility
// - fmp4: Fragmented MP4 segments, better efficiency
type HlsH265HlsSegmentType string

const (
	HlsH265HlsSegmentTypeMpegts HlsH265HlsSegmentType = "mpegts"
	HlsH265HlsSegmentTypeFmp4   HlsH265HlsSegmentType = "fmp4"
)

// PixFmt specifies the pixel format. Valid value: yuv420p
type HlsH265Pixfmt string

const (
	HlsH265PixfmtYuv410p     HlsH265Pixfmt = "yuv410p"
	HlsH265PixfmtYuv411p     HlsH265Pixfmt = "yuv411p"
	HlsH265PixfmtYuv420p     HlsH265Pixfmt = "yuv420p"
	HlsH265PixfmtYuv422p     HlsH265Pixfmt = "yuv422p"
	HlsH265PixfmtYuv440p     HlsH265Pixfmt = "yuv440p"
	HlsH265PixfmtYuv444p     HlsH265Pixfmt = "yuv444p"
	HlsH265PixfmtYuvJ411p    HlsH265Pixfmt = "yuvJ411p"
	HlsH265PixfmtYuvJ420p    HlsH265Pixfmt = "yuvJ420p"
	HlsH265PixfmtYuvJ422p    HlsH265Pixfmt = "yuvJ422p"
	HlsH265PixfmtYuvJ440p    HlsH265Pixfmt = "yuvJ440p"
	HlsH265PixfmtYuvJ444p    HlsH265Pixfmt = "yuvJ444p"
	HlsH265PixfmtYuv420p10le HlsH265Pixfmt = "yuv420p10le"
	HlsH265PixfmtYuv422p10le HlsH265Pixfmt = "yuv422p10le"
	HlsH265PixfmtYuv440p10le HlsH265Pixfmt = "yuv440p10le"
	HlsH265PixfmtYuv444p10le HlsH265Pixfmt = "yuv444p10le"
	HlsH265PixfmtYuv420p12le HlsH265Pixfmt = "yuv420p12le"
	HlsH265PixfmtYuv422p12le HlsH265Pixfmt = "yuv422p12le"
	HlsH265PixfmtYuv440p12le HlsH265Pixfmt = "yuv440p12le"
	HlsH265PixfmtYuv444p12le HlsH265Pixfmt = "yuv444p12le"
	HlsH265PixfmtYuv420p10be HlsH265Pixfmt = "yuv420p10be"
	HlsH265PixfmtYuv422p10be HlsH265Pixfmt = "yuv422p10be"
	HlsH265PixfmtYuv440p10be HlsH265Pixfmt = "yuv440p10be"
	HlsH265PixfmtYuv444p10be HlsH265Pixfmt = "yuv444p10be"
	HlsH265PixfmtYuv420p12be HlsH265Pixfmt = "yuv420p12be"
	HlsH265PixfmtYuv422p12be HlsH265Pixfmt = "yuv422p12be"
	HlsH265PixfmtYuv440p12be HlsH265Pixfmt = "yuv440p12be"
	HlsH265PixfmtYuv444p12be HlsH265Pixfmt = "yuv444p12be"
)

// Preset specifies the encoding speed preset. Valid values (from fastest to
// slowest):
//
// - ultrafast: Fastest encoding, lowest quality
// - superfast: Very fast encoding, lower quality
// - veryfast: Fast encoding, moderate quality
// - faster: Faster encoding, good quality
// - fast: Fast encoding, better quality
// - medium: Balanced preset, best quality
type HlsH265Preset string

const (
	HlsH265PresetUltrafast HlsH265Preset = "ultrafast"
	HlsH265PresetSuperfast HlsH265Preset = "superfast"
	HlsH265PresetVeryfast  HlsH265Preset = "veryfast"
	HlsH265PresetFaster    HlsH265Preset = "faster"
	HlsH265PresetFast      HlsH265Preset = "fast"
	HlsH265PresetMedium    HlsH265Preset = "medium"
)

// Profilev specifies the H.265 profile. Valid values:
//
// - main: Main profile, good for most applications
// - main10: Main 10-bit profile, supports 10-bit color
// - mainstillpicture: Still picture profile, optimized for single images
type HlsH265Profilev string

const (
	HlsH265ProfilevMain             HlsH265Profilev = "main"
	HlsH265ProfilevMain10           HlsH265Profilev = "main10"
	HlsH265ProfilevMainstillpicture HlsH265Profilev = "mainstillpicture"
)

// The property ID is required.
type HlsH265Param struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio param.Opt[bool] `json:"disable_audio,omitzero"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo param.Opt[bool] `json:"disable_video,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate param.Opt[float64] `json:"framerate,omitzero"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop param.Opt[int64] `json:"gop,omitzero"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height param.Opt[int64] `json:"height,omitzero"`
	// HlsEnc enables encryption for HLS segments when set to true.
	HlsEnc param.Opt[bool] `json:"hls_enc,omitzero"`
	// HlsEncIv specifies the initialization vector for encryption. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncIv param.Opt[string] `json:"hls_enc_iv,omitzero"`
	// HlsEncKey specifies the encryption key for HLS segments. Maximum length: 64
	// characters. Required when HlsEnc is true.
	HlsEncKey param.Opt[string] `json:"hls_enc_key,omitzero"`
	// HlsEncKeyUrl specifies the URL where clients can fetch the encryption key.
	// Required when HlsEnc is true.
	HlsEncKeyURL param.Opt[string] `json:"hls_enc_key_url,omitzero"`
	// HlsTime specifies the duration of each HLS segment in seconds. Range: 1 to 10.
	// Shorter segments provide faster startup but more overhead, longer segments are
	// more efficient.
	HlsTime param.Opt[int64] `json:"hls_time,omitzero"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate param.Opt[int64] `json:"maxrate,omitzero"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  param.Opt[int64]  `json:"minrate,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek param.Opt[int64] `json:"seek,omitzero"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate param.Opt[int64] `json:"video_bitrate,omitzero"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width param.Opt[int64] `json:"width,omitzero"`
	// X265KeyInt specifies the maximum number of frames between keyframes for H.265
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X265Keyint param.Opt[int64] `json:"x265_keyint,omitzero"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels,omitzero"`
	// HlsSegmentType specifies the type of HLS segments. Valid values:
	//
	// - mpegts: Traditional MPEG-TS segments, better compatibility
	// - fmp4: Fragmented MP4 segments, better efficiency
	//
	// Any of "mpegts", "fmp4".
	HlsSegmentType HlsH265HlsSegmentType `json:"hls_segment_type,omitzero"`
	// Level specifies the H.265 profile level. Valid values: 30-31 (main), 41
	// (main10). Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level,omitzero"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt HlsH265Pixfmt `json:"pixfmt,omitzero"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset HlsH265Preset `json:"preset,omitzero"`
	// Profilev specifies the H.265 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev HlsH265Profilev `json:"profilev,omitzero"`
	// This field can be elided, and will marshal its zero value as "hls_h265".
	ID constant.HlsH265 `json:"id,required"`
	paramObj
}

func (r HlsH265Param) MarshalJSON() (data []byte, err error) {
	type shadow HlsH265Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HlsH265Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[HlsH265Param](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[HlsH265Param](
		"level", 30, 31, 41,
	)
}

type Job struct {
	// Unique identifier for the job
	ID string `json:"id,required"`
	// Billable time in seconds
	BillableTime int64 `json:"billable_time,required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A template defines the transcoding parameters and settings for a job
	Format JobFormatUnion `json:"format,required"`
	// Progress percentage of the job (0-100)
	Progress float64 `json:"progress,required"`
	// ID of the source video being transcoded
	SourceID string `json:"source_id,required"`
	// Current status of the job
	//
	// Any of "queued", "ingesting", "transcoding", "downloading", "merging",
	// "uploading", "failed", "completed", "cancelled", "merged", "downloaded",
	// "transcoded", "waiting".
	Status JobStatus `json:"status,required"`
	// Storage settings for where the job output will be saved
	Storage JobStorage `json:"storage,required"`
	// The transcoder configuration for a job
	Transcoder JobTranscoder `json:"transcoder,required"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Error message for the job
	Error shared.ChunkifyError `json:"error"`
	// HLS manifest ID
	HlsManifestID string `json:"hls_manifest_id"`
	// Additional metadata for the job
	Metadata map[string]string `json:"metadata"`
	// When the job started processing
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BillableTime  respjson.Field
		CreatedAt     respjson.Field
		Format        respjson.Field
		Progress      respjson.Field
		SourceID      respjson.Field
		Status        respjson.Field
		Storage       respjson.Field
		Transcoder    respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		HlsManifestID respjson.Field
		Metadata      respjson.Field
		StartedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JobFormatUnion contains all possible properties and values from
// [JobFormatMP4Av1], [JobFormatMP4H264], [JobFormatMP4H265], [JobFormatWebmVp9],
// [JobFormatHlsAv1], [JobFormatHlsH264], [JobFormatHlsH265], [JobFormatJpg].
//
// Use the [JobFormatUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type JobFormatUnion struct {
	// Any of "mp4_av1", "mp4_h264", "mp4_h265", "webm_vp9", "hls_av1", "hls_h264",
	// "hls_h265", "jpg".
	ID           string  `json:"id"`
	AudioBitrate int64   `json:"audio_bitrate"`
	Bufsize      int64   `json:"bufsize"`
	Channels     int64   `json:"channels"`
	Crf          int64   `json:"crf"`
	DisableAudio bool    `json:"disable_audio"`
	DisableVideo bool    `json:"disable_video"`
	Duration     int64   `json:"duration"`
	Framerate    float64 `json:"framerate"`
	Gop          int64   `json:"gop"`
	Height       int64   `json:"height"`
	Level        int64   `json:"level"`
	Maxrate      int64   `json:"maxrate"`
	Minrate      int64   `json:"minrate"`
	Movflags     string  `json:"movflags"`
	Pixfmt       string  `json:"pixfmt"`
	Preset       string  `json:"preset"`
	Profilev     string  `json:"profilev"`
	Seek         int64   `json:"seek"`
	VideoBitrate int64   `json:"video_bitrate"`
	Width        int64   `json:"width"`
	X264Keyint   int64   `json:"x264_keyint"`
	X265Keyint   int64   `json:"x265_keyint"`
	// This field is from variant [JobFormatWebmVp9].
	CPUUsed WebmVp9CPUUsed `json:"cpu_used"`
	// This field is from variant [JobFormatWebmVp9].
	Quality        WebmVp9Quality `json:"quality"`
	HlsEnc         bool           `json:"hls_enc"`
	HlsEncIv       string         `json:"hls_enc_iv"`
	HlsEncKey      string         `json:"hls_enc_key"`
	HlsEncKeyURL   string         `json:"hls_enc_key_url"`
	HlsSegmentType string         `json:"hls_segment_type"`
	HlsTime        int64          `json:"hls_time"`
	// This field is from variant [JobFormatJpg].
	Interval int64 `json:"interval"`
	// This field is from variant [JobFormatJpg].
	ChunkDuration int64 `json:"chunk_duration"`
	// This field is from variant [JobFormatJpg].
	Frames int64 `json:"frames"`
	// This field is from variant [JobFormatJpg].
	Sprite bool `json:"sprite"`
	JSON   struct {
		ID             respjson.Field
		AudioBitrate   respjson.Field
		Bufsize        respjson.Field
		Channels       respjson.Field
		Crf            respjson.Field
		DisableAudio   respjson.Field
		DisableVideo   respjson.Field
		Duration       respjson.Field
		Framerate      respjson.Field
		Gop            respjson.Field
		Height         respjson.Field
		Level          respjson.Field
		Maxrate        respjson.Field
		Minrate        respjson.Field
		Movflags       respjson.Field
		Pixfmt         respjson.Field
		Preset         respjson.Field
		Profilev       respjson.Field
		Seek           respjson.Field
		VideoBitrate   respjson.Field
		Width          respjson.Field
		X264Keyint     respjson.Field
		X265Keyint     respjson.Field
		CPUUsed        respjson.Field
		Quality        respjson.Field
		HlsEnc         respjson.Field
		HlsEncIv       respjson.Field
		HlsEncKey      respjson.Field
		HlsEncKeyURL   respjson.Field
		HlsSegmentType respjson.Field
		HlsTime        respjson.Field
		Interval       respjson.Field
		ChunkDuration  respjson.Field
		Frames         respjson.Field
		Sprite         respjson.Field
		raw            string
	} `json:"-"`
}

// anyJobFormat is implemented by each variant of [JobFormatUnion] to add type
// safety for the return type of [JobFormatUnion.AsAny]
type anyJobFormat interface {
	implJobFormatUnion()
}

func (JobFormatMP4Av1) implJobFormatUnion()  {}
func (JobFormatMP4H264) implJobFormatUnion() {}
func (JobFormatMP4H265) implJobFormatUnion() {}
func (JobFormatWebmVp9) implJobFormatUnion() {}
func (JobFormatHlsAv1) implJobFormatUnion()  {}
func (JobFormatHlsH264) implJobFormatUnion() {}
func (JobFormatHlsH265) implJobFormatUnion() {}
func (JobFormatJpg) implJobFormatUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := JobFormatUnion.AsAny().(type) {
//	case chunkify.JobFormatMP4Av1:
//	case chunkify.JobFormatMP4H264:
//	case chunkify.JobFormatMP4H265:
//	case chunkify.JobFormatWebmVp9:
//	case chunkify.JobFormatHlsAv1:
//	case chunkify.JobFormatHlsH264:
//	case chunkify.JobFormatHlsH265:
//	case chunkify.JobFormatJpg:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u JobFormatUnion) AsAny() anyJobFormat {
	switch u.ID {
	case "mp4_av1":
		return u.AsMP4Av1()
	case "mp4_h264":
		return u.AsMP4H264()
	case "mp4_h265":
		return u.AsMP4H265()
	case "webm_vp9":
		return u.AsWebmVp9()
	case "hls_av1":
		return u.AsHlsAv1()
	case "hls_h264":
		return u.AsHlsH264()
	case "hls_h265":
		return u.AsHlsH265()
	case "jpg":
		return u.AsJpg()
	}
	return nil
}

func (u JobFormatUnion) AsMP4Av1() (v JobFormatMP4Av1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JobFormatUnion) AsMP4H264() (v JobFormatMP4H264) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JobFormatUnion) AsMP4H265() (v JobFormatMP4H265) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JobFormatUnion) AsWebmVp9() (v JobFormatWebmVp9) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JobFormatUnion) AsHlsAv1() (v JobFormatHlsAv1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JobFormatUnion) AsHlsH264() (v JobFormatHlsH264) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JobFormatUnion) AsHlsH265() (v JobFormatHlsH265) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JobFormatUnion) AsJpg() (v JobFormatJpg) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u JobFormatUnion) RawJSON() string { return u.JSON.raw }

func (r *JobFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to MP4 with AV1 encoding.
type JobFormatMP4Av1 struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	MP4Av1
}

// Returns the unmodified JSON received from the API
func (r JobFormatMP4Av1) RawJSON() string { return r.JSON.raw }
func (r *JobFormatMP4Av1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to MP4 with H.264 encoding.
type JobFormatMP4H264 struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	MP4H264
}

// Returns the unmodified JSON received from the API
func (r JobFormatMP4H264) RawJSON() string { return r.JSON.raw }
func (r *JobFormatMP4H264) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to MP4 with H.265 encoding.
type JobFormatMP4H265 struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	MP4H265
}

// Returns the unmodified JSON received from the API
func (r JobFormatMP4H265) RawJSON() string { return r.JSON.raw }
func (r *JobFormatMP4H265) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to WebM with VP9 encoding.
type JobFormatWebmVp9 struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	WebmVp9
}

// Returns the unmodified JSON received from the API
func (r JobFormatWebmVp9) RawJSON() string { return r.JSON.raw }
func (r *JobFormatWebmVp9) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to HLS with AV1 encoding.
type JobFormatHlsAv1 struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	HlsAv1
}

// Returns the unmodified JSON received from the API
func (r JobFormatHlsAv1) RawJSON() string { return r.JSON.raw }
func (r *JobFormatHlsAv1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to HLS with H.264 encoding.
type JobFormatHlsH264 struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	HlsH264
}

// Returns the unmodified JSON received from the API
func (r JobFormatHlsH264) RawJSON() string { return r.JSON.raw }
func (r *JobFormatHlsH264) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to HLS with H.265 encoding.
type JobFormatHlsH265 struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	HlsH265
}

// Returns the unmodified JSON received from the API
func (r JobFormatHlsH265) RawJSON() string { return r.JSON.raw }
func (r *JobFormatHlsH265) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to JPEG image extraction.
type JobFormatJpg struct {
	// The format ID
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Jpg
}

// Returns the unmodified JSON received from the API
func (r JobFormatJpg) RawJSON() string { return r.JSON.raw }
func (r *JobFormatJpg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the job
type JobStatus string

const (
	JobStatusQueued      JobStatus = "queued"
	JobStatusIngesting   JobStatus = "ingesting"
	JobStatusTranscoding JobStatus = "transcoding"
	JobStatusDownloading JobStatus = "downloading"
	JobStatusMerging     JobStatus = "merging"
	JobStatusUploading   JobStatus = "uploading"
	JobStatusFailed      JobStatus = "failed"
	JobStatusCompleted   JobStatus = "completed"
	JobStatusCancelled   JobStatus = "cancelled"
	JobStatusMerged      JobStatus = "merged"
	JobStatusDownloaded  JobStatus = "downloaded"
	JobStatusTranscoded  JobStatus = "transcoded"
	JobStatusWaiting     JobStatus = "waiting"
)

// Storage settings for where the job output will be saved
type JobStorage struct {
	// ID of the storage
	ID string `json:"id,required"`
	// Path where the output will be stored
	Path string `json:"path,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobStorage) RawJSON() string { return r.JSON.raw }
func (r *JobStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The transcoder configuration for a job
type JobTranscoder struct {
	// Whether the transcoder configuration is automatically set by Chunkify
	Auto bool `json:"auto,required"`
	// Number of instances allocated
	Quantity int64 `json:"quantity,required"`
	// Type of transcoder instance
	//
	// Any of "4vCPU", "8vCPU", "16vCPU".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Auto        respjson.Field
		Quantity    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobTranscoder) RawJSON() string { return r.JSON.raw }
func (r *JobTranscoder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to JPEG image extraction.
type Jpg struct {
	ID constant.Jpg `json:"id,required"`
	// Time interval in seconds at which frames are extracted from the video (e.g.,
	// interval=10 extracts frames at 0s, 10s, 20s, etc.). Must be between 1 and 60
	// seconds.
	Interval      int64 `json:"interval,required"`
	ChunkDuration int64 `json:"chunk_duration"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	Frames   int64 `json:"frames"`
	Height   int64 `json:"height"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek   int64 `json:"seek"`
	Sprite bool  `json:"sprite"`
	Width  int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Interval      respjson.Field
		ChunkDuration respjson.Field
		Duration      respjson.Field
		Frames        respjson.Field
		Height        respjson.Field
		Seek          respjson.Field
		Sprite        respjson.Field
		Width         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Jpg) RawJSON() string { return r.JSON.raw }
func (r *Jpg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Jpg to a JpgParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JpgParam.Overrides()
func (r Jpg) ToParam() JpgParam {
	return param.Override[JpgParam](json.RawMessage(r.RawJSON()))
}

// FFmpeg encoding parameters specific to JPEG image extraction.
//
// The properties ID, Interval are required.
type JpgParam struct {
	// Time interval in seconds at which frames are extracted from the video (e.g.,
	// interval=10 extracts frames at 0s, 10s, 20s, etc.). Must be between 1 and 60
	// seconds.
	Interval      int64            `json:"interval,required"`
	ChunkDuration param.Opt[int64] `json:"chunk_duration,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	Frames   param.Opt[int64] `json:"frames,omitzero"`
	Height   param.Opt[int64] `json:"height,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek   param.Opt[int64] `json:"seek,omitzero"`
	Sprite param.Opt[bool]  `json:"sprite,omitzero"`
	Width  param.Opt[int64] `json:"width,omitzero"`
	// This field can be elided, and will marshal its zero value as "jpg".
	ID constant.Jpg `json:"id,required"`
	paramObj
}

func (r JpgParam) MarshalJSON() (data []byte, err error) {
	type shadow JpgParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JpgParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MP4Av1 struct {
	ID constant.MP4Av1 `json:"id,required"`
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 63. Recommended
	// values: 16-35 for high quality, 35-45 for good quality, 45-63 for acceptable
	// quality.
	Crf int64 `json:"crf"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio bool `json:"disable_audio"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo bool `json:"disable_video"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate float64 `json:"framerate"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop int64 `json:"gop"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height int64 `json:"height"`
	// Level specifies the AV1 profile level. Valid values: 30-31 (main), 41 (main10).
	// Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  int64  `json:"minrate"`
	Movflags string `json:"movflags"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt MP4Av1Pixfmt `json:"pixfmt"`
	// Preset controls the encoding efficiency and processing intensity. Lower presets
	// use more optimization features, creating smaller files with better quality but
	// requiring more compute time. Higher presets encode faster but produce larger
	// files.
	//
	// Preset ranges:
	//
	// - 6-7: Fast encoding for real-time applications (smaller files)
	// - 8-10: Balanced efficiency and speed for general use
	// - 11-13: Fastest encoding for real-time applications (larger files)
	//
	// Any of "6", "7", "8", "9", "10", "11", "12", "13".
	Preset MP4Av1Preset `json:"preset"`
	// Profilev specifies the AV1 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev MP4Av1Profilev `json:"profilev"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek int64 `json:"seek"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AudioBitrate respjson.Field
		Bufsize      respjson.Field
		Channels     respjson.Field
		Crf          respjson.Field
		DisableAudio respjson.Field
		DisableVideo respjson.Field
		Duration     respjson.Field
		Framerate    respjson.Field
		Gop          respjson.Field
		Height       respjson.Field
		Level        respjson.Field
		Maxrate      respjson.Field
		Minrate      respjson.Field
		Movflags     respjson.Field
		Pixfmt       respjson.Field
		Preset       respjson.Field
		Profilev     respjson.Field
		Seek         respjson.Field
		VideoBitrate respjson.Field
		Width        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MP4Av1) RawJSON() string { return r.JSON.raw }
func (r *MP4Av1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MP4Av1 to a MP4Av1Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MP4Av1Param.Overrides()
func (r MP4Av1) ToParam() MP4Av1Param {
	return param.Override[MP4Av1Param](json.RawMessage(r.RawJSON()))
}

// PixFmt specifies the pixel format. Valid value: yuv420p
type MP4Av1Pixfmt string

const (
	MP4Av1PixfmtYuv410p     MP4Av1Pixfmt = "yuv410p"
	MP4Av1PixfmtYuv411p     MP4Av1Pixfmt = "yuv411p"
	MP4Av1PixfmtYuv420p     MP4Av1Pixfmt = "yuv420p"
	MP4Av1PixfmtYuv422p     MP4Av1Pixfmt = "yuv422p"
	MP4Av1PixfmtYuv440p     MP4Av1Pixfmt = "yuv440p"
	MP4Av1PixfmtYuv444p     MP4Av1Pixfmt = "yuv444p"
	MP4Av1PixfmtYuvJ411p    MP4Av1Pixfmt = "yuvJ411p"
	MP4Av1PixfmtYuvJ420p    MP4Av1Pixfmt = "yuvJ420p"
	MP4Av1PixfmtYuvJ422p    MP4Av1Pixfmt = "yuvJ422p"
	MP4Av1PixfmtYuvJ440p    MP4Av1Pixfmt = "yuvJ440p"
	MP4Av1PixfmtYuvJ444p    MP4Av1Pixfmt = "yuvJ444p"
	MP4Av1PixfmtYuv420p10le MP4Av1Pixfmt = "yuv420p10le"
	MP4Av1PixfmtYuv422p10le MP4Av1Pixfmt = "yuv422p10le"
	MP4Av1PixfmtYuv440p10le MP4Av1Pixfmt = "yuv440p10le"
	MP4Av1PixfmtYuv444p10le MP4Av1Pixfmt = "yuv444p10le"
	MP4Av1PixfmtYuv420p12le MP4Av1Pixfmt = "yuv420p12le"
	MP4Av1PixfmtYuv422p12le MP4Av1Pixfmt = "yuv422p12le"
	MP4Av1PixfmtYuv440p12le MP4Av1Pixfmt = "yuv440p12le"
	MP4Av1PixfmtYuv444p12le MP4Av1Pixfmt = "yuv444p12le"
	MP4Av1PixfmtYuv420p10be MP4Av1Pixfmt = "yuv420p10be"
	MP4Av1PixfmtYuv422p10be MP4Av1Pixfmt = "yuv422p10be"
	MP4Av1PixfmtYuv440p10be MP4Av1Pixfmt = "yuv440p10be"
	MP4Av1PixfmtYuv444p10be MP4Av1Pixfmt = "yuv444p10be"
	MP4Av1PixfmtYuv420p12be MP4Av1Pixfmt = "yuv420p12be"
	MP4Av1PixfmtYuv422p12be MP4Av1Pixfmt = "yuv422p12be"
	MP4Av1PixfmtYuv440p12be MP4Av1Pixfmt = "yuv440p12be"
	MP4Av1PixfmtYuv444p12be MP4Av1Pixfmt = "yuv444p12be"
)

// Preset controls the encoding efficiency and processing intensity. Lower presets
// use more optimization features, creating smaller files with better quality but
// requiring more compute time. Higher presets encode faster but produce larger
// files.
//
// Preset ranges:
//
// - 6-7: Fast encoding for real-time applications (smaller files)
// - 8-10: Balanced efficiency and speed for general use
// - 11-13: Fastest encoding for real-time applications (larger files)
type MP4Av1Preset string

const (
	MP4Av1Preset6  MP4Av1Preset = "6"
	MP4Av1Preset7  MP4Av1Preset = "7"
	MP4Av1Preset8  MP4Av1Preset = "8"
	MP4Av1Preset9  MP4Av1Preset = "9"
	MP4Av1Preset10 MP4Av1Preset = "10"
	MP4Av1Preset11 MP4Av1Preset = "11"
	MP4Av1Preset12 MP4Av1Preset = "12"
	MP4Av1Preset13 MP4Av1Preset = "13"
)

// Profilev specifies the AV1 profile. Valid values:
//
// - main: Main profile, good for most applications
// - main10: Main 10-bit profile, supports 10-bit color
// - mainstillpicture: Still picture profile, optimized for single images
type MP4Av1Profilev string

const (
	MP4Av1ProfilevMain             MP4Av1Profilev = "main"
	MP4Av1ProfilevMain10           MP4Av1Profilev = "main10"
	MP4Av1ProfilevMainstillpicture MP4Av1Profilev = "mainstillpicture"
)

// The property ID is required.
type MP4Av1Param struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 63. Recommended
	// values: 16-35 for high quality, 35-45 for good quality, 45-63 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio param.Opt[bool] `json:"disable_audio,omitzero"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo param.Opt[bool] `json:"disable_video,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate param.Opt[float64] `json:"framerate,omitzero"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop param.Opt[int64] `json:"gop,omitzero"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate param.Opt[int64] `json:"maxrate,omitzero"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  param.Opt[int64]  `json:"minrate,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek param.Opt[int64] `json:"seek,omitzero"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate param.Opt[int64] `json:"video_bitrate,omitzero"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width param.Opt[int64] `json:"width,omitzero"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels,omitzero"`
	// Level specifies the AV1 profile level. Valid values: 30-31 (main), 41 (main10).
	// Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level,omitzero"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt MP4Av1Pixfmt `json:"pixfmt,omitzero"`
	// Preset controls the encoding efficiency and processing intensity. Lower presets
	// use more optimization features, creating smaller files with better quality but
	// requiring more compute time. Higher presets encode faster but produce larger
	// files.
	//
	// Preset ranges:
	//
	// - 6-7: Fast encoding for real-time applications (smaller files)
	// - 8-10: Balanced efficiency and speed for general use
	// - 11-13: Fastest encoding for real-time applications (larger files)
	//
	// Any of "6", "7", "8", "9", "10", "11", "12", "13".
	Preset MP4Av1Preset `json:"preset,omitzero"`
	// Profilev specifies the AV1 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev MP4Av1Profilev `json:"profilev,omitzero"`
	// This field can be elided, and will marshal its zero value as "mp4_av1".
	ID constant.MP4Av1 `json:"id,required"`
	paramObj
}

func (r MP4Av1Param) MarshalJSON() (data []byte, err error) {
	type shadow MP4Av1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MP4Av1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MP4Av1Param](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[MP4Av1Param](
		"level", 30, 31, 41,
	)
}

type MP4H264 struct {
	ID constant.MP4H264 `json:"id,required"`
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf int64 `json:"crf"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio bool `json:"disable_audio"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo bool `json:"disable_video"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate float64 `json:"framerate"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop int64 `json:"gop"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height int64 `json:"height"`
	// Level specifies the H.264 profile level. Valid values: 10-13 (baseline), 20-22
	// (main), 30-32 (high), 40-42 (high), 50-51 (high). Higher levels support higher
	// resolutions and bitrates but require more processing power.
	//
	// Any of 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51.
	Level int64 `json:"level"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  int64  `json:"minrate"`
	Movflags string `json:"movflags"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt MP4H264Pixfmt `json:"pixfmt"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset MP4H264Preset `json:"preset"`
	// Profilev specifies the H.264 profile. Valid values:
	//
	// - baseline: Basic profile, good for mobile devices
	// - main: Main profile, good for most applications
	// - high: High profile, best quality but requires more processing
	// - high10: High 10-bit profile, supports 10-bit color
	// - high422: High 4:2:2 profile, supports 4:2:2 color sampling
	// - high444: High 4:4:4 profile, supports 4:4:4 color sampling
	//
	// Any of "baseline", "main", "high", "high10", "high422", "high444".
	Profilev MP4H264Profilev `json:"profilev"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek int64 `json:"seek"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width int64 `json:"width"`
	// X264KeyInt specifies the maximum number of frames between keyframes for H.264
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X264Keyint int64 `json:"x264_keyint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AudioBitrate respjson.Field
		Bufsize      respjson.Field
		Channels     respjson.Field
		Crf          respjson.Field
		DisableAudio respjson.Field
		DisableVideo respjson.Field
		Duration     respjson.Field
		Framerate    respjson.Field
		Gop          respjson.Field
		Height       respjson.Field
		Level        respjson.Field
		Maxrate      respjson.Field
		Minrate      respjson.Field
		Movflags     respjson.Field
		Pixfmt       respjson.Field
		Preset       respjson.Field
		Profilev     respjson.Field
		Seek         respjson.Field
		VideoBitrate respjson.Field
		Width        respjson.Field
		X264Keyint   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MP4H264) RawJSON() string { return r.JSON.raw }
func (r *MP4H264) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MP4H264 to a MP4H264Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MP4H264Param.Overrides()
func (r MP4H264) ToParam() MP4H264Param {
	return param.Override[MP4H264Param](json.RawMessage(r.RawJSON()))
}

// PixFmt specifies the pixel format. Valid value: yuv420p
type MP4H264Pixfmt string

const (
	MP4H264PixfmtYuv410p     MP4H264Pixfmt = "yuv410p"
	MP4H264PixfmtYuv411p     MP4H264Pixfmt = "yuv411p"
	MP4H264PixfmtYuv420p     MP4H264Pixfmt = "yuv420p"
	MP4H264PixfmtYuv422p     MP4H264Pixfmt = "yuv422p"
	MP4H264PixfmtYuv440p     MP4H264Pixfmt = "yuv440p"
	MP4H264PixfmtYuv444p     MP4H264Pixfmt = "yuv444p"
	MP4H264PixfmtYuvJ411p    MP4H264Pixfmt = "yuvJ411p"
	MP4H264PixfmtYuvJ420p    MP4H264Pixfmt = "yuvJ420p"
	MP4H264PixfmtYuvJ422p    MP4H264Pixfmt = "yuvJ422p"
	MP4H264PixfmtYuvJ440p    MP4H264Pixfmt = "yuvJ440p"
	MP4H264PixfmtYuvJ444p    MP4H264Pixfmt = "yuvJ444p"
	MP4H264PixfmtYuv420p10le MP4H264Pixfmt = "yuv420p10le"
	MP4H264PixfmtYuv422p10le MP4H264Pixfmt = "yuv422p10le"
	MP4H264PixfmtYuv440p10le MP4H264Pixfmt = "yuv440p10le"
	MP4H264PixfmtYuv444p10le MP4H264Pixfmt = "yuv444p10le"
	MP4H264PixfmtYuv420p12le MP4H264Pixfmt = "yuv420p12le"
	MP4H264PixfmtYuv422p12le MP4H264Pixfmt = "yuv422p12le"
	MP4H264PixfmtYuv440p12le MP4H264Pixfmt = "yuv440p12le"
	MP4H264PixfmtYuv444p12le MP4H264Pixfmt = "yuv444p12le"
	MP4H264PixfmtYuv420p10be MP4H264Pixfmt = "yuv420p10be"
	MP4H264PixfmtYuv422p10be MP4H264Pixfmt = "yuv422p10be"
	MP4H264PixfmtYuv440p10be MP4H264Pixfmt = "yuv440p10be"
	MP4H264PixfmtYuv444p10be MP4H264Pixfmt = "yuv444p10be"
	MP4H264PixfmtYuv420p12be MP4H264Pixfmt = "yuv420p12be"
	MP4H264PixfmtYuv422p12be MP4H264Pixfmt = "yuv422p12be"
	MP4H264PixfmtYuv440p12be MP4H264Pixfmt = "yuv440p12be"
	MP4H264PixfmtYuv444p12be MP4H264Pixfmt = "yuv444p12be"
)

// Preset specifies the encoding speed preset. Valid values (from fastest to
// slowest):
//
// - ultrafast: Fastest encoding, lowest quality
// - superfast: Very fast encoding, lower quality
// - veryfast: Fast encoding, moderate quality
// - faster: Faster encoding, good quality
// - fast: Fast encoding, better quality
// - medium: Balanced preset, best quality
type MP4H264Preset string

const (
	MP4H264PresetUltrafast MP4H264Preset = "ultrafast"
	MP4H264PresetSuperfast MP4H264Preset = "superfast"
	MP4H264PresetVeryfast  MP4H264Preset = "veryfast"
	MP4H264PresetFaster    MP4H264Preset = "faster"
	MP4H264PresetFast      MP4H264Preset = "fast"
	MP4H264PresetMedium    MP4H264Preset = "medium"
)

// Profilev specifies the H.264 profile. Valid values:
//
// - baseline: Basic profile, good for mobile devices
// - main: Main profile, good for most applications
// - high: High profile, best quality but requires more processing
// - high10: High 10-bit profile, supports 10-bit color
// - high422: High 4:2:2 profile, supports 4:2:2 color sampling
// - high444: High 4:4:4 profile, supports 4:4:4 color sampling
type MP4H264Profilev string

const (
	MP4H264ProfilevBaseline MP4H264Profilev = "baseline"
	MP4H264ProfilevMain     MP4H264Profilev = "main"
	MP4H264ProfilevHigh     MP4H264Profilev = "high"
	MP4H264ProfilevHigh10   MP4H264Profilev = "high10"
	MP4H264ProfilevHigh422  MP4H264Profilev = "high422"
	MP4H264ProfilevHigh444  MP4H264Profilev = "high444"
)

// The property ID is required.
type MP4H264Param struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio param.Opt[bool] `json:"disable_audio,omitzero"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo param.Opt[bool] `json:"disable_video,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate param.Opt[float64] `json:"framerate,omitzero"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop param.Opt[int64] `json:"gop,omitzero"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate param.Opt[int64] `json:"maxrate,omitzero"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  param.Opt[int64]  `json:"minrate,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek param.Opt[int64] `json:"seek,omitzero"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate param.Opt[int64] `json:"video_bitrate,omitzero"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width param.Opt[int64] `json:"width,omitzero"`
	// X264KeyInt specifies the maximum number of frames between keyframes for H.264
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X264Keyint param.Opt[int64] `json:"x264_keyint,omitzero"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels,omitzero"`
	// Level specifies the H.264 profile level. Valid values: 10-13 (baseline), 20-22
	// (main), 30-32 (high), 40-42 (high), 50-51 (high). Higher levels support higher
	// resolutions and bitrates but require more processing power.
	//
	// Any of 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51.
	Level int64 `json:"level,omitzero"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt MP4H264Pixfmt `json:"pixfmt,omitzero"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset MP4H264Preset `json:"preset,omitzero"`
	// Profilev specifies the H.264 profile. Valid values:
	//
	// - baseline: Basic profile, good for mobile devices
	// - main: Main profile, good for most applications
	// - high: High profile, best quality but requires more processing
	// - high10: High 10-bit profile, supports 10-bit color
	// - high422: High 4:2:2 profile, supports 4:2:2 color sampling
	// - high444: High 4:4:4 profile, supports 4:4:4 color sampling
	//
	// Any of "baseline", "main", "high", "high10", "high422", "high444".
	Profilev MP4H264Profilev `json:"profilev,omitzero"`
	// This field can be elided, and will marshal its zero value as "mp4_h264".
	ID constant.MP4H264 `json:"id,required"`
	paramObj
}

func (r MP4H264Param) MarshalJSON() (data []byte, err error) {
	type shadow MP4H264Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MP4H264Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MP4H264Param](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[MP4H264Param](
		"level", 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51,
	)
}

type MP4H265 struct {
	ID constant.MP4H265 `json:"id,required"`
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf int64 `json:"crf"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio bool `json:"disable_audio"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo bool `json:"disable_video"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate float64 `json:"framerate"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop int64 `json:"gop"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height int64 `json:"height"`
	// Level specifies the H.265 profile level. Valid values: 30-31 (main), 41
	// (main10). Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  int64  `json:"minrate"`
	Movflags string `json:"movflags"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt MP4H265Pixfmt `json:"pixfmt"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset MP4H265Preset `json:"preset"`
	// Profilev specifies the H.265 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev MP4H265Profilev `json:"profilev"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek int64 `json:"seek"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width int64 `json:"width"`
	// X265KeyInt specifies the maximum number of frames between keyframes for H.265
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X265Keyint int64 `json:"x265_keyint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AudioBitrate respjson.Field
		Bufsize      respjson.Field
		Channels     respjson.Field
		Crf          respjson.Field
		DisableAudio respjson.Field
		DisableVideo respjson.Field
		Duration     respjson.Field
		Framerate    respjson.Field
		Gop          respjson.Field
		Height       respjson.Field
		Level        respjson.Field
		Maxrate      respjson.Field
		Minrate      respjson.Field
		Movflags     respjson.Field
		Pixfmt       respjson.Field
		Preset       respjson.Field
		Profilev     respjson.Field
		Seek         respjson.Field
		VideoBitrate respjson.Field
		Width        respjson.Field
		X265Keyint   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MP4H265) RawJSON() string { return r.JSON.raw }
func (r *MP4H265) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MP4H265 to a MP4H265Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MP4H265Param.Overrides()
func (r MP4H265) ToParam() MP4H265Param {
	return param.Override[MP4H265Param](json.RawMessage(r.RawJSON()))
}

// PixFmt specifies the pixel format. Valid value: yuv420p
type MP4H265Pixfmt string

const (
	MP4H265PixfmtYuv410p     MP4H265Pixfmt = "yuv410p"
	MP4H265PixfmtYuv411p     MP4H265Pixfmt = "yuv411p"
	MP4H265PixfmtYuv420p     MP4H265Pixfmt = "yuv420p"
	MP4H265PixfmtYuv422p     MP4H265Pixfmt = "yuv422p"
	MP4H265PixfmtYuv440p     MP4H265Pixfmt = "yuv440p"
	MP4H265PixfmtYuv444p     MP4H265Pixfmt = "yuv444p"
	MP4H265PixfmtYuvJ411p    MP4H265Pixfmt = "yuvJ411p"
	MP4H265PixfmtYuvJ420p    MP4H265Pixfmt = "yuvJ420p"
	MP4H265PixfmtYuvJ422p    MP4H265Pixfmt = "yuvJ422p"
	MP4H265PixfmtYuvJ440p    MP4H265Pixfmt = "yuvJ440p"
	MP4H265PixfmtYuvJ444p    MP4H265Pixfmt = "yuvJ444p"
	MP4H265PixfmtYuv420p10le MP4H265Pixfmt = "yuv420p10le"
	MP4H265PixfmtYuv422p10le MP4H265Pixfmt = "yuv422p10le"
	MP4H265PixfmtYuv440p10le MP4H265Pixfmt = "yuv440p10le"
	MP4H265PixfmtYuv444p10le MP4H265Pixfmt = "yuv444p10le"
	MP4H265PixfmtYuv420p12le MP4H265Pixfmt = "yuv420p12le"
	MP4H265PixfmtYuv422p12le MP4H265Pixfmt = "yuv422p12le"
	MP4H265PixfmtYuv440p12le MP4H265Pixfmt = "yuv440p12le"
	MP4H265PixfmtYuv444p12le MP4H265Pixfmt = "yuv444p12le"
	MP4H265PixfmtYuv420p10be MP4H265Pixfmt = "yuv420p10be"
	MP4H265PixfmtYuv422p10be MP4H265Pixfmt = "yuv422p10be"
	MP4H265PixfmtYuv440p10be MP4H265Pixfmt = "yuv440p10be"
	MP4H265PixfmtYuv444p10be MP4H265Pixfmt = "yuv444p10be"
	MP4H265PixfmtYuv420p12be MP4H265Pixfmt = "yuv420p12be"
	MP4H265PixfmtYuv422p12be MP4H265Pixfmt = "yuv422p12be"
	MP4H265PixfmtYuv440p12be MP4H265Pixfmt = "yuv440p12be"
	MP4H265PixfmtYuv444p12be MP4H265Pixfmt = "yuv444p12be"
)

// Preset specifies the encoding speed preset. Valid values (from fastest to
// slowest):
//
// - ultrafast: Fastest encoding, lowest quality
// - superfast: Very fast encoding, lower quality
// - veryfast: Fast encoding, moderate quality
// - faster: Faster encoding, good quality
// - fast: Fast encoding, better quality
// - medium: Balanced preset, best quality
type MP4H265Preset string

const (
	MP4H265PresetUltrafast MP4H265Preset = "ultrafast"
	MP4H265PresetSuperfast MP4H265Preset = "superfast"
	MP4H265PresetVeryfast  MP4H265Preset = "veryfast"
	MP4H265PresetFaster    MP4H265Preset = "faster"
	MP4H265PresetFast      MP4H265Preset = "fast"
	MP4H265PresetMedium    MP4H265Preset = "medium"
)

// Profilev specifies the H.265 profile. Valid values:
//
// - main: Main profile, good for most applications
// - main10: Main 10-bit profile, supports 10-bit color
// - mainstillpicture: Still picture profile, optimized for single images
type MP4H265Profilev string

const (
	MP4H265ProfilevMain             MP4H265Profilev = "main"
	MP4H265ProfilevMain10           MP4H265Profilev = "main10"
	MP4H265ProfilevMainstillpicture MP4H265Profilev = "mainstillpicture"
)

// The property ID is required.
type MP4H265Param struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio param.Opt[bool] `json:"disable_audio,omitzero"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo param.Opt[bool] `json:"disable_video,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate param.Opt[float64] `json:"framerate,omitzero"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop param.Opt[int64] `json:"gop,omitzero"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate param.Opt[int64] `json:"maxrate,omitzero"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate  param.Opt[int64]  `json:"minrate,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek param.Opt[int64] `json:"seek,omitzero"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate param.Opt[int64] `json:"video_bitrate,omitzero"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width param.Opt[int64] `json:"width,omitzero"`
	// X265KeyInt specifies the maximum number of frames between keyframes for H.265
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X265Keyint param.Opt[int64] `json:"x265_keyint,omitzero"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels,omitzero"`
	// Level specifies the H.265 profile level. Valid values: 30-31 (main), 41
	// (main10). Higher levels support higher resolutions and bitrates but require more
	// processing power.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level,omitzero"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt MP4H265Pixfmt `json:"pixfmt,omitzero"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	//
	// Any of "ultrafast", "superfast", "veryfast", "faster", "fast", "medium".
	Preset MP4H265Preset `json:"preset,omitzero"`
	// Profilev specifies the H.265 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev MP4H265Profilev `json:"profilev,omitzero"`
	// This field can be elided, and will marshal its zero value as "mp4_h265".
	ID constant.MP4H265 `json:"id,required"`
	paramObj
}

func (r MP4H265Param) MarshalJSON() (data []byte, err error) {
	type shadow MP4H265Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MP4H265Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MP4H265Param](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[MP4H265Param](
		"level", 30, 31, 41,
	)
}

type WebmVp9 struct {
	ID constant.WebmVp9 `json:"id,required"`
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels"`
	// CpuUsed specifies the CPU usage level for VP9 encoding. Range: 0 to 8. Lower
	// values mean better quality but slower encoding, higher values mean faster
	// encoding but lower quality. Recommended values: 0-2 for high quality, 2-4 for
	// good quality, 4-6 for balanced, 6-8 for speed
	//
	// Any of "0", "1", "2", "3", "4", "5", "6", "7", "8".
	CPUUsed WebmVp9CPUUsed `json:"cpu_used"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 15 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf int64 `json:"crf"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio bool `json:"disable_audio"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo bool `json:"disable_video"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration int64 `json:"duration"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate float64 `json:"framerate"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop int64 `json:"gop"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height int64 `json:"height"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate int64 `json:"minrate"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt WebmVp9Pixfmt `json:"pixfmt"`
	// Quality specifies the VP9 encoding quality preset. Valid values:
	//
	// - good: Balanced quality preset, good for most applications
	// - best: Best quality preset, slower encoding
	// - realtime: Fast encoding preset, suitable for live streaming
	//
	// Any of "good", "best", "realtime".
	Quality WebmVp9Quality `json:"quality"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek int64 `json:"seek"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AudioBitrate respjson.Field
		Bufsize      respjson.Field
		Channels     respjson.Field
		CPUUsed      respjson.Field
		Crf          respjson.Field
		DisableAudio respjson.Field
		DisableVideo respjson.Field
		Duration     respjson.Field
		Framerate    respjson.Field
		Gop          respjson.Field
		Height       respjson.Field
		Maxrate      respjson.Field
		Minrate      respjson.Field
		Pixfmt       respjson.Field
		Quality      respjson.Field
		Seek         respjson.Field
		VideoBitrate respjson.Field
		Width        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebmVp9) RawJSON() string { return r.JSON.raw }
func (r *WebmVp9) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WebmVp9 to a WebmVp9Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WebmVp9Param.Overrides()
func (r WebmVp9) ToParam() WebmVp9Param {
	return param.Override[WebmVp9Param](json.RawMessage(r.RawJSON()))
}

// CpuUsed specifies the CPU usage level for VP9 encoding. Range: 0 to 8. Lower
// values mean better quality but slower encoding, higher values mean faster
// encoding but lower quality. Recommended values: 0-2 for high quality, 2-4 for
// good quality, 4-6 for balanced, 6-8 for speed
type WebmVp9CPUUsed string

const (
	WebmVp9CPUUsed0 WebmVp9CPUUsed = "0"
	WebmVp9CPUUsed1 WebmVp9CPUUsed = "1"
	WebmVp9CPUUsed2 WebmVp9CPUUsed = "2"
	WebmVp9CPUUsed3 WebmVp9CPUUsed = "3"
	WebmVp9CPUUsed4 WebmVp9CPUUsed = "4"
	WebmVp9CPUUsed5 WebmVp9CPUUsed = "5"
	WebmVp9CPUUsed6 WebmVp9CPUUsed = "6"
	WebmVp9CPUUsed7 WebmVp9CPUUsed = "7"
	WebmVp9CPUUsed8 WebmVp9CPUUsed = "8"
)

// PixFmt specifies the pixel format. Valid value: yuv420p
type WebmVp9Pixfmt string

const (
	WebmVp9PixfmtYuv410p     WebmVp9Pixfmt = "yuv410p"
	WebmVp9PixfmtYuv411p     WebmVp9Pixfmt = "yuv411p"
	WebmVp9PixfmtYuv420p     WebmVp9Pixfmt = "yuv420p"
	WebmVp9PixfmtYuv422p     WebmVp9Pixfmt = "yuv422p"
	WebmVp9PixfmtYuv440p     WebmVp9Pixfmt = "yuv440p"
	WebmVp9PixfmtYuv444p     WebmVp9Pixfmt = "yuv444p"
	WebmVp9PixfmtYuvJ411p    WebmVp9Pixfmt = "yuvJ411p"
	WebmVp9PixfmtYuvJ420p    WebmVp9Pixfmt = "yuvJ420p"
	WebmVp9PixfmtYuvJ422p    WebmVp9Pixfmt = "yuvJ422p"
	WebmVp9PixfmtYuvJ440p    WebmVp9Pixfmt = "yuvJ440p"
	WebmVp9PixfmtYuvJ444p    WebmVp9Pixfmt = "yuvJ444p"
	WebmVp9PixfmtYuv420p10le WebmVp9Pixfmt = "yuv420p10le"
	WebmVp9PixfmtYuv422p10le WebmVp9Pixfmt = "yuv422p10le"
	WebmVp9PixfmtYuv440p10le WebmVp9Pixfmt = "yuv440p10le"
	WebmVp9PixfmtYuv444p10le WebmVp9Pixfmt = "yuv444p10le"
	WebmVp9PixfmtYuv420p12le WebmVp9Pixfmt = "yuv420p12le"
	WebmVp9PixfmtYuv422p12le WebmVp9Pixfmt = "yuv422p12le"
	WebmVp9PixfmtYuv440p12le WebmVp9Pixfmt = "yuv440p12le"
	WebmVp9PixfmtYuv444p12le WebmVp9Pixfmt = "yuv444p12le"
	WebmVp9PixfmtYuv420p10be WebmVp9Pixfmt = "yuv420p10be"
	WebmVp9PixfmtYuv422p10be WebmVp9Pixfmt = "yuv422p10be"
	WebmVp9PixfmtYuv440p10be WebmVp9Pixfmt = "yuv440p10be"
	WebmVp9PixfmtYuv444p10be WebmVp9Pixfmt = "yuv444p10be"
	WebmVp9PixfmtYuv420p12be WebmVp9Pixfmt = "yuv420p12be"
	WebmVp9PixfmtYuv422p12be WebmVp9Pixfmt = "yuv422p12be"
	WebmVp9PixfmtYuv440p12be WebmVp9Pixfmt = "yuv440p12be"
	WebmVp9PixfmtYuv444p12be WebmVp9Pixfmt = "yuv444p12be"
)

// Quality specifies the VP9 encoding quality preset. Valid values:
//
// - good: Balanced quality preset, good for most applications
// - best: Best quality preset, slower encoding
// - realtime: Fast encoding preset, suitable for live streaming
type WebmVp9Quality string

const (
	WebmVp9QualityGood     WebmVp9Quality = "good"
	WebmVp9QualityBest     WebmVp9Quality = "best"
	WebmVp9QualityRealtime WebmVp9Quality = "realtime"
)

// The property ID is required.
type WebmVp9Param struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 15 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// DisableAudio indicates whether to disable audio processing.
	DisableAudio param.Opt[bool] `json:"disable_audio,omitzero"`
	// DisableVideo indicates whether to disable video processing.
	DisableVideo param.Opt[bool] `json:"disable_video,omitzero"`
	// Duration specifies the duration to process in seconds. Must be a positive value.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Framerate specifies the output video frame rate. Must be between 15 and 120 fps.
	Framerate param.Opt[float64] `json:"framerate,omitzero"`
	// Gop specifies the Group of Pictures (GOP) size. Must be between 1 and 300.
	Gop param.Opt[int64] `json:"gop,omitzero"`
	// Height specifies the output video height in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Height param.Opt[int64] `json:"height,omitzero"`
	// Maxrate specifies the maximum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Maxrate param.Opt[int64] `json:"maxrate,omitzero"`
	// Minrate specifies the minimum video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	Minrate param.Opt[int64] `json:"minrate,omitzero"`
	// Seek specifies the timestamp to start processing from (in seconds). Must be a
	// positive value.
	Seek param.Opt[int64] `json:"seek,omitzero"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate param.Opt[int64] `json:"video_bitrate,omitzero"`
	// Width specifies the output video width in pixels. Must be between -2 and 7680.
	// Use -2 for automatic calculation while maintaining aspect ratio.
	Width param.Opt[int64] `json:"width,omitzero"`
	// Channels specifies the number of audio channels. Valid values: 1 (mono), 2
	// (stereo), 5 (5.1), 7 (7.1)
	//
	// Any of 1, 2, 5, 7.
	Channels int64 `json:"channels,omitzero"`
	// CpuUsed specifies the CPU usage level for VP9 encoding. Range: 0 to 8. Lower
	// values mean better quality but slower encoding, higher values mean faster
	// encoding but lower quality. Recommended values: 0-2 for high quality, 2-4 for
	// good quality, 4-6 for balanced, 6-8 for speed
	//
	// Any of "0", "1", "2", "3", "4", "5", "6", "7", "8".
	CPUUsed WebmVp9CPUUsed `json:"cpu_used,omitzero"`
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt WebmVp9Pixfmt `json:"pixfmt,omitzero"`
	// Quality specifies the VP9 encoding quality preset. Valid values:
	//
	// - good: Balanced quality preset, good for most applications
	// - best: Best quality preset, slower encoding
	// - realtime: Fast encoding preset, suitable for live streaming
	//
	// Any of "good", "best", "realtime".
	Quality WebmVp9Quality `json:"quality,omitzero"`
	// This field can be elided, and will marshal its zero value as "webm_vp9".
	ID constant.WebmVp9 `json:"id,required"`
	paramObj
}

func (r WebmVp9Param) MarshalJSON() (data []byte, err error) {
	type shadow WebmVp9Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebmVp9Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebmVp9Param](
		"channels", 1, 2, 5, 7,
	)
}

type JobNewParams struct {
	// Required format configuration, one and only one valid format configuration must
	// be provided. If you want to use a format without specifying any configuration,
	// use an empty object in the corresponding field.
	Format JobNewParamsFormatUnion `json:"format,omitzero,required"`
	// The ID of the source file to transcode
	SourceID string `json:"source_id,required"`
	// Optional HLS manifest configuration Use the same hls manifest ID to group
	// multiple jobs into a single HLS manifest By default, it's automatically
	// generated if no set for HLS jobs
	HlsManifestID param.Opt[string] `json:"hls_manifest_id,omitzero"`
	// Optional metadata to attach to the job (max 1024 bytes)
	Metadata map[string]string `json:"metadata,omitzero"`
	// Optional storage configuration
	Storage JobNewParamsStorage `json:"storage,omitzero"`
	// Optional transcoder configuration. If not provided, the system will
	// automatically calculate the optimal quantity and CPU type based on the source
	// file specifications and output requirements. This auto-scaling ensures efficient
	// resource utilization.
	Transcoder JobNewParamsTranscoder `json:"transcoder,omitzero"`
	paramObj
}

func (r JobNewParams) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JobNewParamsFormatUnion struct {
	OfMP4Av1  *MP4Av1Param  `json:",omitzero,inline"`
	OfMP4H264 *MP4H264Param `json:",omitzero,inline"`
	OfMP4H265 *MP4H265Param `json:",omitzero,inline"`
	OfWebmVp9 *WebmVp9Param `json:",omitzero,inline"`
	OfHlsAv1  *HlsAv1Param  `json:",omitzero,inline"`
	OfHlsH264 *HlsH264Param `json:",omitzero,inline"`
	OfHlsH265 *HlsH265Param `json:",omitzero,inline"`
	OfJpg     *JpgParam     `json:",omitzero,inline"`
	paramUnion
}

func (u JobNewParamsFormatUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMP4Av1,
		u.OfMP4H264,
		u.OfMP4H265,
		u.OfWebmVp9,
		u.OfHlsAv1,
		u.OfHlsH264,
		u.OfHlsH265,
		u.OfJpg)
}
func (u *JobNewParamsFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *JobNewParamsFormatUnion) asAny() any {
	if !param.IsOmitted(u.OfMP4Av1) {
		return u.OfMP4Av1
	} else if !param.IsOmitted(u.OfMP4H264) {
		return u.OfMP4H264
	} else if !param.IsOmitted(u.OfMP4H265) {
		return u.OfMP4H265
	} else if !param.IsOmitted(u.OfWebmVp9) {
		return u.OfWebmVp9
	} else if !param.IsOmitted(u.OfHlsAv1) {
		return u.OfHlsAv1
	} else if !param.IsOmitted(u.OfHlsH264) {
		return u.OfHlsH264
	} else if !param.IsOmitted(u.OfHlsH265) {
		return u.OfHlsH265
	} else if !param.IsOmitted(u.OfJpg) {
		return u.OfJpg
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetCPUUsed() *string {
	if vt := u.OfWebmVp9; vt != nil {
		return (*string)(&vt.CPUUsed)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetQuality() *string {
	if vt := u.OfWebmVp9; vt != nil {
		return (*string)(&vt.Quality)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetInterval() *int64 {
	if vt := u.OfJpg; vt != nil {
		return &vt.Interval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetChunkDuration() *int64 {
	if vt := u.OfJpg; vt != nil && vt.ChunkDuration.Valid() {
		return &vt.ChunkDuration.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetFrames() *int64 {
	if vt := u.OfJpg; vt != nil && vt.Frames.Valid() {
		return &vt.Frames.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetSprite() *bool {
	if vt := u.OfJpg; vt != nil && vt.Sprite.Valid() {
		return &vt.Sprite.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetID() *string {
	if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfWebmVp9; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfHlsAv1; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfHlsH264; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfHlsH265; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfJpg; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetAudioBitrate() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetBufsize() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Bufsize.Valid() {
		return &vt.Bufsize.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Bufsize.Valid() {
		return &vt.Bufsize.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Bufsize.Valid() {
		return &vt.Bufsize.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Bufsize.Valid() {
		return &vt.Bufsize.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Bufsize.Valid() {
		return &vt.Bufsize.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Bufsize.Valid() {
		return &vt.Bufsize.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Bufsize.Valid() {
		return &vt.Bufsize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetChannels() *int64 {
	if vt := u.OfMP4Av1; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfMP4H264; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfWebmVp9; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfHlsAv1; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfHlsH264; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfHlsH265; vt != nil {
		return (*int64)(&vt.Channels)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetCrf() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetDisableAudio() *bool {
	if vt := u.OfMP4Av1; vt != nil && vt.DisableAudio.Valid() {
		return &vt.DisableAudio.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.DisableAudio.Valid() {
		return &vt.DisableAudio.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.DisableAudio.Valid() {
		return &vt.DisableAudio.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.DisableAudio.Valid() {
		return &vt.DisableAudio.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.DisableAudio.Valid() {
		return &vt.DisableAudio.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.DisableAudio.Valid() {
		return &vt.DisableAudio.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.DisableAudio.Valid() {
		return &vt.DisableAudio.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetDisableVideo() *bool {
	if vt := u.OfMP4Av1; vt != nil && vt.DisableVideo.Valid() {
		return &vt.DisableVideo.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.DisableVideo.Valid() {
		return &vt.DisableVideo.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.DisableVideo.Valid() {
		return &vt.DisableVideo.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.DisableVideo.Valid() {
		return &vt.DisableVideo.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.DisableVideo.Valid() {
		return &vt.DisableVideo.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.DisableVideo.Valid() {
		return &vt.DisableVideo.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.DisableVideo.Valid() {
		return &vt.DisableVideo.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetDuration() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfJpg; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetFramerate() *float64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Framerate.Valid() {
		return &vt.Framerate.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Framerate.Valid() {
		return &vt.Framerate.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Framerate.Valid() {
		return &vt.Framerate.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Framerate.Valid() {
		return &vt.Framerate.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Framerate.Valid() {
		return &vt.Framerate.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Framerate.Valid() {
		return &vt.Framerate.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Framerate.Valid() {
		return &vt.Framerate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetGop() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Gop.Valid() {
		return &vt.Gop.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Gop.Valid() {
		return &vt.Gop.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Gop.Valid() {
		return &vt.Gop.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Gop.Valid() {
		return &vt.Gop.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Gop.Valid() {
		return &vt.Gop.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Gop.Valid() {
		return &vt.Gop.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Gop.Valid() {
		return &vt.Gop.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetHeight() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	} else if vt := u.OfJpg; vt != nil && vt.Height.Valid() {
		return &vt.Height.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetLevel() *int64 {
	if vt := u.OfMP4Av1; vt != nil {
		return (*int64)(&vt.Level)
	} else if vt := u.OfMP4H264; vt != nil {
		return (*int64)(&vt.Level)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*int64)(&vt.Level)
	} else if vt := u.OfHlsAv1; vt != nil {
		return (*int64)(&vt.Level)
	} else if vt := u.OfHlsH264; vt != nil {
		return (*int64)(&vt.Level)
	} else if vt := u.OfHlsH265; vt != nil {
		return (*int64)(&vt.Level)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetMaxrate() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Maxrate.Valid() {
		return &vt.Maxrate.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Maxrate.Valid() {
		return &vt.Maxrate.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Maxrate.Valid() {
		return &vt.Maxrate.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Maxrate.Valid() {
		return &vt.Maxrate.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Maxrate.Valid() {
		return &vt.Maxrate.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Maxrate.Valid() {
		return &vt.Maxrate.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Maxrate.Valid() {
		return &vt.Maxrate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetMinrate() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetMovflags() *string {
	if vt := u.OfMP4Av1; vt != nil && vt.Movflags.Valid() {
		return &vt.Movflags.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Movflags.Valid() {
		return &vt.Movflags.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Movflags.Valid() {
		return &vt.Movflags.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Movflags.Valid() {
		return &vt.Movflags.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Movflags.Valid() {
		return &vt.Movflags.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Movflags.Valid() {
		return &vt.Movflags.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetPixfmt() *string {
	if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfWebmVp9; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfHlsAv1; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfHlsH264; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfHlsH265; vt != nil {
		return (*string)(&vt.Pixfmt)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetPreset() *string {
	if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.Preset)
	} else if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.Preset)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.Preset)
	} else if vt := u.OfHlsAv1; vt != nil {
		return (*string)(&vt.Preset)
	} else if vt := u.OfHlsH264; vt != nil {
		return (*string)(&vt.Preset)
	} else if vt := u.OfHlsH265; vt != nil {
		return (*string)(&vt.Preset)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetProfilev() *string {
	if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.Profilev)
	} else if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.Profilev)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.Profilev)
	} else if vt := u.OfHlsAv1; vt != nil {
		return (*string)(&vt.Profilev)
	} else if vt := u.OfHlsH264; vt != nil {
		return (*string)(&vt.Profilev)
	} else if vt := u.OfHlsH265; vt != nil {
		return (*string)(&vt.Profilev)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetSeek() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfJpg; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetVideoBitrate() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.VideoBitrate.Valid() {
		return &vt.VideoBitrate.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.VideoBitrate.Valid() {
		return &vt.VideoBitrate.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.VideoBitrate.Valid() {
		return &vt.VideoBitrate.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.VideoBitrate.Valid() {
		return &vt.VideoBitrate.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.VideoBitrate.Valid() {
		return &vt.VideoBitrate.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.VideoBitrate.Valid() {
		return &vt.VideoBitrate.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.VideoBitrate.Valid() {
		return &vt.VideoBitrate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetWidth() *int64 {
	if vt := u.OfMP4Av1; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	} else if vt := u.OfMP4H264; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	} else if vt := u.OfHlsAv1; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	} else if vt := u.OfJpg; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetX264Keyint() *int64 {
	if vt := u.OfMP4H264; vt != nil && vt.X264Keyint.Valid() {
		return &vt.X264Keyint.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.X264Keyint.Valid() {
		return &vt.X264Keyint.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetX265Keyint() *int64 {
	if vt := u.OfMP4H265; vt != nil && vt.X265Keyint.Valid() {
		return &vt.X265Keyint.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.X265Keyint.Valid() {
		return &vt.X265Keyint.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetHlsEnc() *bool {
	if vt := u.OfHlsAv1; vt != nil && vt.HlsEnc.Valid() {
		return &vt.HlsEnc.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.HlsEnc.Valid() {
		return &vt.HlsEnc.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.HlsEnc.Valid() {
		return &vt.HlsEnc.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetHlsEncIv() *string {
	if vt := u.OfHlsAv1; vt != nil && vt.HlsEncIv.Valid() {
		return &vt.HlsEncIv.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.HlsEncIv.Valid() {
		return &vt.HlsEncIv.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.HlsEncIv.Valid() {
		return &vt.HlsEncIv.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetHlsEncKey() *string {
	if vt := u.OfHlsAv1; vt != nil && vt.HlsEncKey.Valid() {
		return &vt.HlsEncKey.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.HlsEncKey.Valid() {
		return &vt.HlsEncKey.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.HlsEncKey.Valid() {
		return &vt.HlsEncKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetHlsEncKeyURL() *string {
	if vt := u.OfHlsAv1; vt != nil && vt.HlsEncKeyURL.Valid() {
		return &vt.HlsEncKeyURL.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.HlsEncKeyURL.Valid() {
		return &vt.HlsEncKeyURL.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.HlsEncKeyURL.Valid() {
		return &vt.HlsEncKeyURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetHlsSegmentType() *string {
	if vt := u.OfHlsAv1; vt != nil {
		return (*string)(&vt.HlsSegmentType)
	} else if vt := u.OfHlsH264; vt != nil {
		return (*string)(&vt.HlsSegmentType)
	} else if vt := u.OfHlsH265; vt != nil {
		return (*string)(&vt.HlsSegmentType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JobNewParamsFormatUnion) GetHlsTime() *int64 {
	if vt := u.OfHlsAv1; vt != nil && vt.HlsTime.Valid() {
		return &vt.HlsTime.Value
	} else if vt := u.OfHlsH264; vt != nil && vt.HlsTime.Valid() {
		return &vt.HlsTime.Value
	} else if vt := u.OfHlsH265; vt != nil && vt.HlsTime.Valid() {
		return &vt.HlsTime.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[JobNewParamsFormatUnion](
		"id",
		apijson.Discriminator[MP4Av1Param]("mp4_av1"),
		apijson.Discriminator[MP4H264Param]("mp4_h264"),
		apijson.Discriminator[MP4H265Param]("mp4_h265"),
		apijson.Discriminator[WebmVp9Param]("webm_vp9"),
		apijson.Discriminator[HlsAv1Param]("hls_av1"),
		apijson.Discriminator[HlsH264Param]("hls_h264"),
		apijson.Discriminator[HlsH265Param]("hls_h265"),
		apijson.Discriminator[JpgParam]("jpg"),
	)
}

// Optional storage configuration
type JobNewParamsStorage struct {
	// Storage Id specifies the storage configuration to use from pre-configured
	// storage options. Must be 4-64 characters long and contain only alphanumeric
	// characters, underscores and hyphens. Optional if Storage Path is provided.
	ID param.Opt[string] `json:"id,omitzero"`
	// Storage Path specifies a custom storage path where processed files will be
	// stored. Must be a valid file path with max length of 1024 characters. Optional
	// if Storage Id is provided.
	Path param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r JobNewParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional transcoder configuration. If not provided, the system will
// automatically calculate the optimal quantity and CPU type based on the source
// file specifications and output requirements. This auto-scaling ensures efficient
// resource utilization.
type JobNewParamsTranscoder struct {
	// Quantity specifies the number of transcoder instances. Required if Type is set.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// Type specifies the CPU configuration for each transcoder instance. Required if
	// Quantity is set.
	//
	// Any of "4vCPU", "8vCPU", "16vCPU".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r JobNewParamsTranscoder) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsTranscoder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsTranscoder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsTranscoder](
		"type", "4vCPU", "8vCPU", "16vCPU",
	)
}

type JobNewResponseEnvelope struct {
	// Data contains the response object
	Data Job `json:"data,required"`
	// Status indicates the response status "success"
	Status constant.Success `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetResponseEnvelope struct {
	// Data contains the response object
	Data Job `json:"data,required"`
	// Status indicates the response status "success"
	Status constant.Success `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListParams struct {
	// Filter by job ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Filter by hls manifest ID
	HlsManifestID param.Opt[string] `query:"hls_manifest_id,omitzero" json:"-"`
	// Pagination limit
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by source ID
	SourceID param.Opt[string]    `query:"source_id,omitzero" json:"-"`
	Created  JobListParamsCreated `query:"created,omitzero" json:"-"`
	// Filter by format id
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "hls_h264", "hls_h265",
	// "hls_av1", "jpg".
	FormatID JobListParamsFormatID `query:"format_id,omitzero" json:"-"`
	// Filter by metadata (format: key:value,key2:value2)
	Metadata [][]string `query:"metadata,omitzero" json:"-"`
	// Filter by job status
	//
	// Any of "completed", "processing", "failed", "cancelled", "queued".
	Status JobListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobListParams]'s query parameters as `url.Values`.
func (r JobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JobListParamsCreated struct {
	// Filter by creation date greater than or equal
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Filter by creation date less than or equal
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	// Sort by creation date (asc/desc)
	//
	// Any of "asc", "desc".
	Sort string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobListParamsCreated]'s query parameters as `url.Values`.
func (r JobListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by format id
type JobListParamsFormatID string

const (
	JobListParamsFormatIDMP4H264 JobListParamsFormatID = "mp4_h264"
	JobListParamsFormatIDMP4H265 JobListParamsFormatID = "mp4_h265"
	JobListParamsFormatIDMP4Av1  JobListParamsFormatID = "mp4_av1"
	JobListParamsFormatIDWebmVp9 JobListParamsFormatID = "webm_vp9"
	JobListParamsFormatIDHlsH264 JobListParamsFormatID = "hls_h264"
	JobListParamsFormatIDHlsH265 JobListParamsFormatID = "hls_h265"
	JobListParamsFormatIDHlsAv1  JobListParamsFormatID = "hls_av1"
	JobListParamsFormatIDJpg     JobListParamsFormatID = "jpg"
)

// Filter by job status
type JobListParamsStatus string

const (
	JobListParamsStatusCompleted  JobListParamsStatus = "completed"
	JobListParamsStatusProcessing JobListParamsStatus = "processing"
	JobListParamsStatusFailed     JobListParamsStatus = "failed"
	JobListParamsStatusCancelled  JobListParamsStatus = "cancelled"
	JobListParamsStatusQueued     JobListParamsStatus = "queued"
)
