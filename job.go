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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/cancel", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
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

type Job struct {
	// Unique identifier for the job
	ID string `json:"id"`
	// Billable time in seconds
	BillableTime int64 `json:"billable_time"`
	// Creation timestamp
	CreatedAt string `json:"created_at"`
	// Error message for the job
	Error JobError `json:"error"`
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

// Error message for the job
type JobError struct {
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
func (r JobError) RawJSON() string { return r.JSON.raw }
func (r *JobError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A template defines the transcoding parameters and settings for a job
type JobFormat struct {
	// Configuration parameters for the template. A map of configuration values
	// specific to the format For example, for mp4/h264 format this includes parameters
	// like crf, preset, profile etc.
	Config any `json:"config"`
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

// FFmpeg encoding parameters specific to MP4 with H.264 encoding.
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

type JobNewResponse struct {
	Data Job `json:"data"`
	// Status indicates the response status "success"
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobNewResponse) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetResponse struct {
	Data Job `json:"data"`
	// Status indicates the response status "success"
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetFilesResponse struct {
	Data any `json:"data"`
	// Status indicates the response status "success"
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetFilesResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetFilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetLogsResponse struct {
	Data any `json:"data"`
	// Status indicates the response status "success"
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetTranscodersResponse struct {
	Data any `json:"data"`
	// Status indicates the response status "success"
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetTranscodersResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetTranscodersResponse) UnmarshalJSON(data []byte) error {
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
	// AV1 configuration
	MP4Av1 JobNewParamsFormatMP4Av1 `json:"mp4_av1,omitzero"`
	// FFmpeg encoding parameters specific to MP4 with H.264 encoding.
	MP4H264 MP4H264Param `json:"mp4_h264,omitzero"`
	// H265 configuration
	MP4H265 JobNewParamsFormatMP4H265 `json:"mp4_h265,omitzero"`
	// VP9 configuration
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
	HlsSegmentType string `json:"hls_segment_type,omitzero"`
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
	Pixfmt string `json:"pixfmt,omitzero"`
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
	Preset string `json:"preset,omitzero"`
	// Profilev specifies the AV1 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev string `json:"profilev,omitzero"`
	paramObj
}

func (r JobNewParamsFormatHlsAv1) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatHlsAv1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormatHlsAv1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsAv1](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsAv1](
		"hls_segment_type", "mpegts", "fmp4",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsAv1](
		"level", 30, 31, 41,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsAv1](
		"pixfmt", "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le", "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be", "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsAv1](
		"preset", "6", "7", "8", "9", "10", "11", "12", "13",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsAv1](
		"profilev", "main", "main10", "mainstillpicture",
	)
}

// HLS H264 configuration
type JobNewParamsFormatHlsH264 struct {
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
	HlsSegmentType string `json:"hls_segment_type,omitzero"`
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
	Pixfmt string `json:"pixfmt,omitzero"`
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
	Preset string `json:"preset,omitzero"`
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
	Profilev string `json:"profilev,omitzero"`
	paramObj
}

func (r JobNewParamsFormatHlsH264) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatHlsH264
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormatHlsH264) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH264](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH264](
		"hls_segment_type", "mpegts", "fmp4",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH264](
		"level", 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH264](
		"pixfmt", "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le", "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be", "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH264](
		"preset", "ultrafast", "superfast", "veryfast", "faster", "fast", "medium",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH264](
		"profilev", "baseline", "main", "high", "high10", "high422", "high444",
	)
}

// HLS H265 configuration
type JobNewParamsFormatHlsH265 struct {
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
	HlsSegmentType string `json:"hls_segment_type,omitzero"`
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
	Pixfmt string `json:"pixfmt,omitzero"`
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
	Preset string `json:"preset,omitzero"`
	// Profilev specifies the H.265 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev string `json:"profilev,omitzero"`
	paramObj
}

func (r JobNewParamsFormatHlsH265) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatHlsH265
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormatHlsH265) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH265](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH265](
		"hls_segment_type", "mpegts", "fmp4",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH265](
		"level", 30, 31, 41,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH265](
		"pixfmt", "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le", "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be", "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH265](
		"preset", "ultrafast", "superfast", "veryfast", "faster", "fast", "medium",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatHlsH265](
		"profilev", "main", "main10", "mainstillpicture",
	)
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

// AV1 configuration
type JobNewParamsFormatMP4Av1 struct {
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
	Pixfmt string `json:"pixfmt,omitzero"`
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
	Preset string `json:"preset,omitzero"`
	// Profilev specifies the AV1 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev string `json:"profilev,omitzero"`
	paramObj
}

func (r JobNewParamsFormatMP4Av1) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatMP4Av1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormatMP4Av1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4Av1](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4Av1](
		"level", 30, 31, 41,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4Av1](
		"pixfmt", "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le", "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be", "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4Av1](
		"preset", "6", "7", "8", "9", "10", "11", "12", "13",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4Av1](
		"profilev", "main", "main10", "mainstillpicture",
	)
}

// H265 configuration
type JobNewParamsFormatMP4H265 struct {
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
	Pixfmt string `json:"pixfmt,omitzero"`
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
	Preset string `json:"preset,omitzero"`
	// Profilev specifies the H.265 profile. Valid values:
	//
	// - main: Main profile, good for most applications
	// - main10: Main 10-bit profile, supports 10-bit color
	// - mainstillpicture: Still picture profile, optimized for single images
	//
	// Any of "main", "main10", "mainstillpicture".
	Profilev string `json:"profilev,omitzero"`
	paramObj
}

func (r JobNewParamsFormatMP4H265) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsFormatMP4H265
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsFormatMP4H265) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4H265](
		"channels", 1, 2, 5, 7,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4H265](
		"level", 30, 31, 41,
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4H265](
		"pixfmt", "yuv410p", "yuv411p", "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuvJ411p", "yuvJ420p", "yuvJ422p", "yuvJ440p", "yuvJ444p", "yuv420p10le", "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le", "yuv444p12le", "yuv420p10be", "yuv422p10be", "yuv440p10be", "yuv444p10be", "yuv420p12be", "yuv422p12be", "yuv440p12be", "yuv444p12be",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4H265](
		"preset", "ultrafast", "superfast", "veryfast", "faster", "fast", "medium",
	)
	apijson.RegisterFieldValidator[JobNewParamsFormatMP4H265](
		"profilev", "main", "main10", "mainstillpicture",
	)
}

// VP9 configuration
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
