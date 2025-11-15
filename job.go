// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/chunkify-go/internal/apijson"
	"github.com/stainless-sdks/chunkify-go/internal/apiquery"
	"github.com/stainless-sdks/chunkify-go/internal/requestconfig"
	"github.com/stainless-sdks/chunkify-go/option"
	"github.com/stainless-sdks/chunkify-go/packages/pagination"
	"github.com/stainless-sdks/chunkify-go/packages/param"
	"github.com/stainless-sdks/chunkify-go/packages/respjson"
)

// JobService contains methods and other services that help with interacting with
// the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r JobService) {
	r = JobService{}
	r.Options = opts
	return
}

// Create a new video processing job with specified parameters
func (r *JobService) New(ctx context.Context, body JobNewParams, opts ...option.RequestOption) (res *JobNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific job
func (r *JobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of jobs with optional filtering and pagination
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *pagination.MyOffsetPage[Job], err error) {
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
func (r *JobService) ListAutoPaging(ctx context.Context, query JobListParams, opts ...option.RequestOption) *pagination.MyOffsetPageAutoPager[Job] {
	return pagination.NewMyOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a job.
func (r *JobService) Delete(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Cancel a job.
func (r *JobService) Cancel(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/cancel", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve all files associated with a specific job
func (r *JobService) GetFiles(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetFilesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/files", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve logs for a specific job, either from the transcoder or manager service
func (r *JobService) GetLogs(ctx context.Context, jobID string, query JobGetLogsParams, opts ...option.RequestOption) (res *JobGetLogsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/logs", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve all the transcoders statuses for a specific job
func (r *JobService) GetTranscoders(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetTranscodersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/transcoders", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// FFmpeg encoding parameters specific to AV1 encoding.
type Av1Param struct {
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 63. Recommended
	// values: 16-35 for high quality, 35-45 for good quality, 45-63 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// Level specifies the AV1 profile level. Valid values: 30-31 (main), 41 (main10).
	// Higher levels support higher resolutions and bitrates but require more
	// processing power.
	Level    int64             `json:"level,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
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
	Preset string `json:"preset,omitzero"`
	// Profilev specifies the AV1 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	Profilev string `json:"profilev,omitzero"`
	VideoCommonParam
}

func (r Av1Param) MarshalJSON() (data []byte, err error) {
	type shadow Av1Param
	return param.MarshalObject(r, (*shadow)(&r))
}

type ChunkifyError struct {
	// Additional error details or output
	Detail string `json:"detail"`
	// Main error message
	Message string `json:"message"`
	// Type of error (e.g., "ffmpeg", "network", "storage", etc.)
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkifyError) RawJSON() string { return r.JSON.raw }
func (r *ChunkifyError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to H.264/AVC encoding.
type H264Param struct {
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// Level specifies the H.264 profile level. Valid values: 10-13 (baseline), 20-22
	// (main), 30-32 (high), 40-42 (high), 50-51 (high). Higher levels support higher
	// resolutions and bitrates but require more processing power.
	Level    int64             `json:"level,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	Preset string `json:"preset,omitzero"`
	// Profilev specifies the H.264 profile. Valid values:
	//
	// - baseline: Basic profile, good for mobile devices
	// - main: Main profile, good for most applications
	// - high: High profile, best quality but requires more processing
	// - high10: High 10-bit profile, supports 10-bit color
	// - high422: High 4:2:2 profile, supports 4:2:2 color sampling
	// - high444: High 4:4:4 profile, supports 4:4:4 color sampling
	Profilev string `json:"profilev,omitzero"`
	// X264KeyInt specifies the maximum number of frames between keyframes for H.264
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X264Keyint param.Opt[int64] `json:"x264_keyint,omitzero"`
	VideoCommonParam
}

func (r H264Param) MarshalJSON() (data []byte, err error) {
	type shadow H264Param
	return param.MarshalObject(r, (*shadow)(&r))
}

// FFmpeg encoding parameters specific to H.265/HEVC encoding. It extends
// FfmpegCommon with H.265-specific options for quality control and encoding
// profiles.
type H265Param struct {
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// Level specifies the H.265 profile level. Valid values: 30-31 (main), 41
	// (main10). Higher levels support higher resolutions and bitrates but require more
	// processing power.
	Level    int64             `json:"level,omitzero"`
	Movflags param.Opt[string] `json:"movflags,omitzero"`
	// Preset specifies the encoding speed preset. Valid values (from fastest to
	// slowest):
	//
	// - ultrafast: Fastest encoding, lowest quality
	// - superfast: Very fast encoding, lower quality
	// - veryfast: Fast encoding, moderate quality
	// - faster: Faster encoding, good quality
	// - fast: Fast encoding, better quality
	// - medium: Balanced preset, best quality
	Preset string `json:"preset,omitzero"`
	// Profilev specifies the H.265 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	Profilev string `json:"profilev,omitzero"`
	// X265KeyInt specifies the maximum number of frames between keyframes for H.265
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X265Keyint param.Opt[int64] `json:"x265_keyint,omitzero"`
	VideoCommonParam
}

func (r H265Param) MarshalJSON() (data []byte, err error) {
	type shadow H265Param
	return param.MarshalObject(r, (*shadow)(&r))
}

// FFmpeg encoding parameters specific to HLS packaging.
//
// The properties AudioBitrate, VideoBitrate are required.
type HlsParam struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate int64 `json:"audio_bitrate,required"`
	// VideoBitrate specifies the video bitrate in bits per second. Must be between
	// 100Kbps and 50Mbps.
	VideoBitrate int64 `json:"video_bitrate,required"`
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
	// HlsSegmentType specifies the type of HLS segments. Valid values:
	//
	// - mpegts: Traditional MPEG-TS segments, better compatibility
	// - fmp4: Fragmented MP4 segments, better efficiency
	//
	// Any of "mpegts", "fmp4".
	HlsSegmentType HlsHlsSegmentType `json:"hls_segment_type,omitzero"`
	paramObj
}

func (r HlsParam) MarshalJSON() (data []byte, err error) {
	type shadow HlsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HlsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HlsSegmentType specifies the type of HLS segments. Valid values:
//
// - mpegts: Traditional MPEG-TS segments, better compatibility
// - fmp4: Fragmented MP4 segments, better efficiency
type HlsHlsSegmentType string

const (
	HlsHlsSegmentTypeMpegts HlsHlsSegmentType = "mpegts"
	HlsHlsSegmentTypeFmp4   HlsHlsSegmentType = "fmp4"
)

type Job struct {
	// Unique identifier for the job
	ID string `json:"id"`
	// Billable time in seconds
	BillableTime int64 `json:"billable_time"`
	// Creation timestamp
	CreatedAt string `json:"created_at"`
	// Error message for the job
	Error ChunkifyError `json:"error"`
	// A template defines the transcoding parameters and settings for a job
	Format JobFormat `json:"format"`
	// HLS manifest ID
	HlsManifestID string `json:"hls_manifest_id"`
	// Additional metadata for the job
	Metadata map[string]string `json:"metadata"`
	// Progress percentage of the job (0-100)
	Progress float64 `json:"progress"`
	// ID of the source video being transcoded
	SourceID string `json:"source_id"`
	// When the job started processing
	StartedAt string `json:"started_at"`
	// Current status of the job (e.g., "queued", "ingesting","transcoding",
	// "downloading", "merging", "uploading", "failed", "completed")
	Status string `json:"status"`
	// Storage settings for where the job output will be saved
	Storage JobStorage `json:"storage"`
	// The transcoder configuration for a job
	Transcoder JobTranscoder `json:"transcoder"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BillableTime  respjson.Field
		CreatedAt     respjson.Field
		Error         respjson.Field
		Format        respjson.Field
		HlsManifestID respjson.Field
		Metadata      respjson.Field
		Progress      respjson.Field
		SourceID      respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		Storage       respjson.Field
		Transcoder    respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A template defines the transcoding parameters and settings for a job
type JobFormat struct {
	// Configuration parameters for the template. A map of configuration values
	// specific to the format For example, for mp4/h264 format this includes parameters
	// like crf, preset, profile etc.
	Config map[string]any `json:"config"`
	// Name of the transcoding template.The format to use for transcoding. Valid
	// formats are: mp4/h264, mp4/h265, mp4/av1, webm/vp9, hls/h264, hls/h265, hls/av1,
	// jpg
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobFormat) RawJSON() string { return r.JSON.raw }
func (r *JobFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage settings for where the job output will be saved
type JobStorage struct {
	// ID of the storage
	ID string `json:"id"`
	// Path where the output will be stored
	Path string `json:"path"`
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
	// Number of instances allocated
	Quantity int64 `json:"quantity"`
	// Type of transcoder instance
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

// FFmpeg encoding parameters common to all video formats.
type VideoCommonParam struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
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
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt VideoCommonPixfmt `json:"pixfmt,omitzero"`
	paramObj
}

func (r VideoCommonParam) MarshalJSON() (data []byte, err error) {
	type shadow VideoCommonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoCommonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VideoCommonParam](
		"channels", 1, 2, 5, 7,
	)
}

// PixFmt specifies the pixel format. Valid value: yuv420p
type VideoCommonPixfmt string

const (
	VideoCommonPixfmtYuv410p     VideoCommonPixfmt = "yuv410p"
	VideoCommonPixfmtYuv411p     VideoCommonPixfmt = "yuv411p"
	VideoCommonPixfmtYuv420p     VideoCommonPixfmt = "yuv420p"
	VideoCommonPixfmtYuv422p     VideoCommonPixfmt = "yuv422p"
	VideoCommonPixfmtYuv440p     VideoCommonPixfmt = "yuv440p"
	VideoCommonPixfmtYuv444p     VideoCommonPixfmt = "yuv444p"
	VideoCommonPixfmtYuvJ411p    VideoCommonPixfmt = "yuvJ411p"
	VideoCommonPixfmtYuvJ420p    VideoCommonPixfmt = "yuvJ420p"
	VideoCommonPixfmtYuvJ422p    VideoCommonPixfmt = "yuvJ422p"
	VideoCommonPixfmtYuvJ440p    VideoCommonPixfmt = "yuvJ440p"
	VideoCommonPixfmtYuvJ444p    VideoCommonPixfmt = "yuvJ444p"
	VideoCommonPixfmtYuv420p10le VideoCommonPixfmt = "yuv420p10le"
	VideoCommonPixfmtYuv422p10le VideoCommonPixfmt = "yuv422p10le"
	VideoCommonPixfmtYuv440p10le VideoCommonPixfmt = "yuv440p10le"
	VideoCommonPixfmtYuv444p10le VideoCommonPixfmt = "yuv444p10le"
	VideoCommonPixfmtYuv420p12le VideoCommonPixfmt = "yuv420p12le"
	VideoCommonPixfmtYuv422p12le VideoCommonPixfmt = "yuv422p12le"
	VideoCommonPixfmtYuv440p12le VideoCommonPixfmt = "yuv440p12le"
	VideoCommonPixfmtYuv444p12le VideoCommonPixfmt = "yuv444p12le"
	VideoCommonPixfmtYuv420p10be VideoCommonPixfmt = "yuv420p10be"
	VideoCommonPixfmtYuv422p10be VideoCommonPixfmt = "yuv422p10be"
	VideoCommonPixfmtYuv440p10be VideoCommonPixfmt = "yuv440p10be"
	VideoCommonPixfmtYuv444p10be VideoCommonPixfmt = "yuv444p10be"
	VideoCommonPixfmtYuv420p12be VideoCommonPixfmt = "yuv420p12be"
	VideoCommonPixfmtYuv422p12be VideoCommonPixfmt = "yuv422p12be"
	VideoCommonPixfmtYuv440p12be VideoCommonPixfmt = "yuv440p12be"
	VideoCommonPixfmtYuv444p12be VideoCommonPixfmt = "yuv444p12be"
)

// Successful response
type JobNewResponse struct {
	Data Job `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r JobNewResponse) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type JobGetResponse struct {
	Data Job `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r JobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobDeleteResponse = any

type JobCancelResponse = any

// Successful response
type JobGetFilesResponse struct {
	Data []APIFile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r JobGetFilesResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetFilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type JobGetLogsResponse struct {
	Data []JobGetLogsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r JobGetLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetLogsResponseData struct {
	// Additional structured data attached to the log
	Attributes map[string]any `json:"attributes"`
	// Optional ID of the job this log is associated with
	JobID string `json:"job_id"`
	// Log level (e.g. "info", "error", "debug")
	Level string `json:"level"`
	// The log message content
	Msg string `json:"msg"`
	// Name of the service that generated the log
	Service string `json:"service"`
	// Timestamp when the log was created
	Time string `json:"time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		JobID       respjson.Field
		Level       respjson.Field
		Msg         respjson.Field
		Service     respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetLogsResponseData) RawJSON() string { return r.JSON.raw }
func (r *JobGetLogsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type JobGetTranscodersResponse struct {
	Data []JobGetTranscodersResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r JobGetTranscodersResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetTranscodersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetTranscodersResponseData struct {
	// Unique identifier of the transcoder
	ID string `json:"id"`
	// Billable time in seconds
	BillableTime int64 `json:"billable_time"`
	// End time of the current chunk in seconds
	ChunkEndTime float64 `json:"chunk_end_time"`
	// Number of the chunk being processed
	ChunkNumber int64 `json:"chunk_number"`
	// Start time of the current chunk in seconds
	ChunkStartTime float64 `json:"chunk_start_time"`
	// CPU time used for transcoding in seconds
	CPUTime float64 `json:"cpu_time"`
	// Timestamp when the status was created
	CreatedAt string `json:"created_at"`
	// Error message if the transcoding failed
	Error ChunkifyError `json:"error"`
	// Current frames per second being processed
	Fps float64 `json:"fps"`
	// Current frame number being processed
	Frame int64 `json:"frame"`
	// Unique identifier of the job
	JobID string `json:"job_id"`
	// Current output time in seconds
	OutTime int64 `json:"out_time"`
	// Progress percentage of the transcoding operation (0-100)
	Progress float64 `json:"progress"`
	// Current processing speed multiplier
	Speed float64 `json:"speed"`
	// Current status of the transcoder (starting, transcoding, finished, error)
	Status string `json:"status"`
	// Unique identifier of the transcoder instance (generated by the transcoder)
	TranscoderInstanceID string `json:"transcoder_instance_id"`
	// Timestamp when the status was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		BillableTime         respjson.Field
		ChunkEndTime         respjson.Field
		ChunkNumber          respjson.Field
		ChunkStartTime       respjson.Field
		CPUTime              respjson.Field
		CreatedAt            respjson.Field
		Error                respjson.Field
		Fps                  respjson.Field
		Frame                respjson.Field
		JobID                respjson.Field
		OutTime              respjson.Field
		Progress             respjson.Field
		Speed                respjson.Field
		Status               respjson.Field
		TranscoderInstanceID respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetTranscodersResponseData) RawJSON() string { return r.JSON.raw }
func (r *JobGetTranscodersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobNewParams struct {
	// Required format configuration, one and only one valid format configuration must
	// be provided. If you want to use a format without specifying any configuration,
	// use an empty object in the corresponding field.
	Format JobNewParamsFormat `json:"format,omitzero,required"`
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

// Required format configuration, one and only one valid format configuration must
// be provided. If you want to use a format without specifying any configuration,
// use an empty object in the corresponding field.
type JobNewParamsFormat struct {
	// HLS AV1 configuration
	HlsAv1 JobNewParamsFormatHlsAv1 `json:"hls_av1,omitzero"`
	// HLS H264 configuration
	HlsH264 JobNewParamsFormatHlsH264 `json:"hls_h264,omitzero"`
	// HLS H265 configuration
	HlsH265 JobNewParamsFormatHlsH265 `json:"hls_h265,omitzero"`
	// FFmpeg encoding parameters specific to JPEG image extraction.
	Jpg JobNewParamsFormatJpg `json:"jpg,omitzero"`
	// FFmpeg encoding parameters specific to AV1 encoding.
	MP4Av1 Av1Param `json:"mp4_av1,omitzero"`
	// FFmpeg encoding parameters specific to H.264/AVC encoding.
	MP4H264 H264Param `json:"mp4_h264,omitzero"`
	// FFmpeg encoding parameters specific to H.265/HEVC encoding. It extends
	// FfmpegCommon with H.265-specific options for quality control and encoding
	// profiles.
	MP4H265 H265Param `json:"mp4_h265,omitzero"`
	// FFmpeg encoding parameters specific to VP9 encoding.
	WebmVp9 JobNewParamsFormatWebmVp9 `json:"webm_vp9,omitzero"`
	paramObj
}

func (r JobNewParamsFormat) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HLS AV1 configuration
type JobNewParamsFormatHlsAv1 struct {
	HlsParam
	Av1Param
	paramObj
}

func (r JobNewParamsFormatHlsAv1) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatHlsAv1
	return param.MarshalObject(r, (*shadow)(&r))
}

// HLS H264 configuration
type JobNewParamsFormatHlsH264 struct {
	HlsParam
	H264Param
	paramObj
}

func (r JobNewParamsFormatHlsH264) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatHlsH264
	return param.MarshalObject(r, (*shadow)(&r))
}

// HLS H265 configuration
type JobNewParamsFormatHlsH265 struct {
	HlsParam
	H265Param
	paramObj
}

func (r JobNewParamsFormatHlsH265) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatHlsH265
	return param.MarshalObject(r, (*shadow)(&r))
}

// FFmpeg encoding parameters specific to JPEG image extraction.
//
// The property Interval is required.
type JobNewParamsFormatJpg struct {
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
	paramObj
}

func (r JobNewParamsFormatJpg) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatJpg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormatJpg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FFmpeg encoding parameters specific to VP9 encoding.
type JobNewParamsFormatWebmVp9 struct {
	// AudioBitrate specifies the audio bitrate in bits per second. Must be between
	// 32Kbps and 512Kbps.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Bufsize specifies the video buffer size in bits. Must be between 100Kbps and
	// 50Mbps.
	Bufsize param.Opt[int64] `json:"bufsize,omitzero"`
	// CpuUsed specifies the CPU usage level for VP9 encoding. Range: 0 to 8. Lower
	// values mean better quality but slower encoding, higher values mean faster
	// encoding but lower quality. Recommended values: 0-2 for high quality, 2-4 for
	// good quality, 4-6 for balanced, 6-8 for speed
	CPUUsed param.Opt[string] `json:"cpu_used,omitzero"`
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
	// PixFmt specifies the pixel format. Valid value: yuv420p
	//
	// Any of "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p",
	// "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le",
	// "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le",
	// "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be",
	// "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be".
	Pixfmt string `json:"pixfmt,omitzero"`
	// Quality specifies the VP9 encoding quality preset. Valid values:
	//
	// - good: Balanced quality preset, good for most applications
	// - best: Best quality preset, slower encoding
	// - realtime: Fast encoding preset, suitable for live streaming
	//
	// Any of "good", "best", "realtime".
	Quality string `json:"quality,omitzero"`
	paramObj
}

func (r JobNewParamsFormatWebmVp9) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatWebmVp9
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormatWebmVp9) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsFormatWebmVp9](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatWebmVp9](
		"pixfmt", "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le", "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be", "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatWebmVp9](
		"quality", "good", "best", "realtime",
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
	// Quantity specifies the number of transcoder instances to use (1-50). Required if
	// Type is set.
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

type JobListParams struct {
	// Filter by job ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Filter by format name
	FormatName param.Opt[string] `query:"format_name,omitzero" json:"-"`
	// Filter by hls manifest ID
	HlsManifestID param.Opt[string] `query:"hls_manifest_id,omitzero" json:"-"`
	// Pagination limit
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by metadata (format: key:value,key2:value2)
	Metadata param.Opt[string] `query:"metadata,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by source ID
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// Filter by job status
	Status  param.Opt[string]    `query:"status,omitzero" json:"-"`
	Created JobListParamsCreated `query:"created,omitzero" json:"-"`
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
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobListParamsCreated]'s query parameters as `url.Values`.
func (r JobListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JobGetLogsParams struct {
	// Service type (transcoder or manager)
	//
	// Any of "transcoder", "manager".
	Service JobGetLogsParamsService `query:"service,omitzero,required" json:"-"`
	// Transcoder ID (required if service is transcoder)
	TranscoderID param.Opt[int64] `query:"transcoder_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobGetLogsParams]'s query parameters as `url.Values`.
func (r JobGetLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Service type (transcoder or manager)
type JobGetLogsParamsService string

const (
	JobGetLogsParamsServiceTranscoder JobGetLogsParamsService = "transcoder"
	JobGetLogsParamsServiceManager    JobGetLogsParamsService = "manager"
)
