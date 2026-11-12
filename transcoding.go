// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/chunkifydev/chunkify-go/internal/apijson"
	"github.com/chunkifydev/chunkify-go/internal/apiquery"
	"github.com/chunkifydev/chunkify-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-go/option"
	"github.com/chunkifydev/chunkify-go/packages/pagination"
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// TranscodingService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTranscodingService] method instead.
type TranscodingService struct {
	Options []option.RequestOption
}

// NewTranscodingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTranscodingService(opts ...option.RequestOption) (r TranscodingService) {
	r = TranscodingService{}
	r.Options = opts
	return
}

// Create a standalone file export with resolved encoding settings.
func (r *TranscodingService) New(ctx context.Context, body TranscodingNewParams, opts ...option.RequestOption) (res *Transcoding, err error) {
	var env TranscodingNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "transcodings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve the accepted settings, published files and current state of one
// Transcoding.
func (r *TranscodingService) Get(ctx context.Context, transcodingID string, opts ...option.RequestOption) (res *Transcoding, err error) {
	var env TranscodingGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if transcodingID == "" {
		err = errors.New("missing required transcodingId parameter")
		return nil, err
	}
	path := fmt.Sprintf("transcodings/%s", transcodingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List Transcodings in the authenticated Project, newest first by default.
func (r *TranscodingService) List(ctx context.Context, query TranscodingListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Transcoding], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "transcodings"
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

// List Transcodings in the authenticated Project, newest first by default.
func (r *TranscodingService) ListAutoPaging(ctx context.Context, query TranscodingListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Transcoding] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete a ready, failed, cancelled or expired Transcoding and its file records.
// The Source and settled usage are preserved. Customer output files remain unless
// delete_files is true; temporary files retain their original 24-hour expiry.
// Deletion remains available for inactive teams.
func (r *TranscodingService) Delete(ctx context.Context, transcodingID string, body TranscodingDeleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transcodingID == "" {
		err = errors.New("missing required transcodingId parameter")
		return err
	}
	path := fmt.Sprintf("transcodings/%s", transcodingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Cancel a preparing Transcoding and release its unpaid reservation. Repeating an
// accepted cancellation succeeds.
func (r *TranscodingService) Cancel(ctx context.Context, transcodingID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transcodingID == "" {
		err = errors.New("missing required transcodingId parameter")
		return err
	}
	path := fmt.Sprintf("transcodings/%s/cancel", transcodingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// A standalone file-encoding request with its accepted settings and published
// files.
type Transcoding struct {
	// Id uniquely identifies the Transcoding.
	ID string `json:"id" api:"required"`
	// CreatedAt specifies when the request was accepted.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Files contains the published output files and their measured properties, ordered
	// by creation time. Empty until publication. Presigned URLs are refreshed on each
	// read and included in lifecycle webhooks when available.
	Files []TranscodingFile `json:"files" api:"required"`
	// Format contains the resolved encoding settings accepted at creation, including
	// defaults, dimensions, frame rate and normalized seek and duration. Per-title
	// rate control is resolved during processing. Measured media properties are
	// available in Files.
	Format TranscodingFormat `json:"format" api:"required"`
	// Metadata contains the string values attached to the Transcoding.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Status reports processing or output expiry.
	//
	// Any of "preparing", "ready", "failed", "cancelled", "expired".
	Status TranscodingStatus `json:"status" api:"required"`
	// Storage contains the resolved output destination.
	Storage TranscodingStorage `json:"storage" api:"required"`
	// UpdatedAt specifies when the request was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// CreatedByTokenId specifies the Project token used to create the resource. It
	// remains set after the token is revoked. Omitted for dashboard requests, MCP
	// OAuth connections, Workflow actions, and unknown creators.
	CreatedByTokenID string `json:"created_by_token_id"`
	// CreatedByUserId specifies the team member who created the resource. Omitted for
	// Workflow actions, unknown creators, and deleted accounts.
	CreatedByUserID string `json:"created_by_user_id"`
	// CreatedVia specifies how the resource was created: api for API requests,
	// dashboard for the Chunkify dashboard, mcp for MCP clients, and workflow for
	// Workflow actions. Omitted when the creator is unknown.
	//
	// Any of "api", "dashboard", "mcp", "workflow".
	CreatedVia TranscodingCreatedVia `json:"created_via"`
	// DeleteAt specifies when this record is deleted: three days after its temporary
	// output expires, otherwise when the team's data retention period ends, counted
	// from when the request was accepted, and no sooner than three days after it
	// finished. It does not delete files in customer-owned Storage. Omitted while the
	// Transcoding is preparing.
	DeleteAt time.Time `json:"delete_at" format:"date-time"`
	// Error describes why processing failed. Omitted when no processing error exists.
	Error TranscodingError `json:"error"`
	// ExpiresAt specifies when temporary output files expire. Omitted for customer
	// output Storage and before the Transcoding is ready.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// ProcessingStage specifies the current activity while the status is preparing.
	// Omitted otherwise.
	//
	// Any of "queued", "starting", "encoding", "finalizing", "uploading".
	ProcessingStage TranscodingProcessingStage `json:"processing_stage"`
	// Progress specifies the video encoding percentage. Available once processing
	// starts and after the output is ready. Omitted for audio-only output; 100
	// requires published output.
	Progress float64 `json:"progress"`
	// ReadyAt specifies when the output was published. Omitted before the Transcoding
	// is ready.
	ReadyAt time.Time `json:"ready_at" format:"date-time"`
	// SourceId identifies the original Source. Omitted after Source deletion.
	SourceID string `json:"source_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Files            respjson.Field
		Format           respjson.Field
		Metadata         respjson.Field
		Status           respjson.Field
		Storage          respjson.Field
		UpdatedAt        respjson.Field
		CreatedByTokenID respjson.Field
		CreatedByUserID  respjson.Field
		CreatedVia       respjson.Field
		DeleteAt         respjson.Field
		Error            respjson.Field
		ExpiresAt        respjson.Field
		ProcessingStage  respjson.Field
		Progress         respjson.Field
		ReadyAt          respjson.Field
		SourceID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transcoding) RawJSON() string { return r.JSON.raw }
func (r *Transcoding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A validated media file published by a Transcoding.
type TranscodingFile struct {
	// Id identifies this published Transcoding file.
	ID string `json:"id" api:"required"`
	// CreatedAt records when the validated file was published.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Duration specifies the measured output duration in seconds.
	Duration float64 `json:"duration" api:"required"`
	// MimeType specifies the media content type.
	MimeType string `json:"mime_type" api:"required"`
	// Path specifies the output path relative to the Storage base_prefix saved when
	// the Transcoding was created. Later Storage prefix changes do not move this file.
	Path string `json:"path" api:"required"`
	// Size specifies the measured file size in bytes.
	Size int64 `json:"size" api:"required"`
	// StorageId identifies the output Storage.
	StorageID string `json:"storage_id" api:"required"`
	// Thumbnails contains available dashboard preview images. Empty for audio-only
	// files, failed previews or expired files.
	Thumbnails []TranscodingFileThumbnail `json:"thumbnails" api:"required"`
	// TranscodingId identifies the Transcoding that produced this file.
	TranscodingID string `json:"transcoding_id" api:"required"`
	// AudioBitrate specifies measured audio bits per second, or zero when unavailable.
	// Omitted without audio.
	AudioBitrate int64 `json:"audio_bitrate"`
	// AudioCodec specifies the measured audio codec. Omitted without audio.
	AudioCodec string `json:"audio_codec"`
	// AudioSamplerate specifies measured audio samples per second. Omitted without
	// audio.
	AudioSamplerate int64 `json:"audio_samplerate"`
	// CdnUrl provides the URL from the Storage CDN configuration. Omitted without a
	// configured CDN or after file expiry.
	CdnURL string `json:"cdn_url"`
	// ExpiresAt specifies when a temporary output expires. Omitted for customer-owned
	// output storage.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Height specifies the measured video height in pixels. Omitted for audio-only
	// files.
	Height int64 `json:"height"`
	// ThumbnailStatus reports the optional dashboard preview result. Failure does not
	// affect the media file. Omitted for audio-only files.
	//
	// Any of "ready", "failed".
	ThumbnailStatus string `json:"thumbnail_status"`
	// Url provides a signed download valid for at most ten minutes. Omitted when the
	// file has expired or signing is unavailable.
	URL string `json:"url"`
	// VideoBitrate specifies measured video bits per second, or zero when unavailable.
	// Omitted without video.
	VideoBitrate int64 `json:"video_bitrate"`
	// VideoCodec specifies the measured video codec. Omitted for audio-only files.
	VideoCodec string `json:"video_codec"`
	// VideoFramerate specifies measured video frames per second. Omitted without
	// video.
	VideoFramerate float64 `json:"video_framerate"`
	// Width specifies the measured video width in pixels. Omitted for audio-only
	// files.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Duration        respjson.Field
		MimeType        respjson.Field
		Path            respjson.Field
		Size            respjson.Field
		StorageID       respjson.Field
		Thumbnails      respjson.Field
		TranscodingID   respjson.Field
		AudioBitrate    respjson.Field
		AudioCodec      respjson.Field
		AudioSamplerate respjson.Field
		CdnURL          respjson.Field
		ExpiresAt       respjson.Field
		Height          respjson.Field
		ThumbnailStatus respjson.Field
		URL             respjson.Field
		VideoBitrate    respjson.Field
		VideoCodec      respjson.Field
		VideoFramerate  respjson.Field
		Width           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscodingFile) RawJSON() string { return r.JSON.raw }
func (r *TranscodingFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An optional dashboard preview of the completed output.
type TranscodingFileThumbnail struct {
	// Format specifies the image encoding.
	//
	// Any of "jpeg", "webp".
	Format string `json:"format" api:"required"`
	// Height specifies the measured image height in pixels.
	Height int64 `json:"height" api:"required"`
	// Size specifies the measured image size in bytes.
	Size int64 `json:"size" api:"required"`
	// Url provides a signed image URL valid for at most ten minutes.
	URL string `json:"url" api:"required"`
	// Width specifies the measured image width in pixels.
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		Height      respjson.Field
		Size        respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscodingFileThumbnail) RawJSON() string { return r.JSON.raw }
func (r *TranscodingFileThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format contains the resolved encoding settings accepted at creation, including
// defaults, dimensions, frame rate and normalized seek and duration. Per-title
// rate control is resolved during processing. Measured media properties are
// available in Files.
type TranscodingFormat struct {
	// Id specifies the resolved container and codec selection.
	//
	// Any of "mp4_h264", "mp4_h265", "mp4_av1", "webm_vp9", "m4a_aac".
	ID string `json:"id" api:"required"`
	// Duration specifies the accepted output duration in seconds, including normalized
	// trims.
	Duration float64 `json:"duration" api:"required"`
	// Seek specifies the accepted start position in source seconds, normalized to the
	// output timeline.
	Seek float64 `json:"seek" api:"required"`
	// AudioBitrate specifies the resolved audio bitrate in bits per second. Omitted
	// for video-only output.
	AudioBitrate int64 `json:"audio_bitrate"`
	// AudioSamplerate specifies the resolved audio samples per second. Omitted for
	// video-only output.
	AudioSamplerate int64 `json:"audio_samplerate"`
	// Bufsize specifies the resolved video buffer size in bits, including codec-level
	// defaults. Omitted when no buffer applies.
	Bufsize int64 `json:"bufsize"`
	// Channels specifies the resolved audio channel count: 1, 2, 6 or 8. Omitted for
	// video-only output.
	Channels int64 `json:"channels"`
	// CpuUsed specifies the accepted VP9 encoder speed setting from 0 through 8.
	// Omitted for other codecs.
	CPUUsed string `json:"cpu_used"`
	// Crf specifies the accepted constant quality. Omitted for bitrate-controlled,
	// per-title and audio-only encoding.
	Crf int64 `json:"crf"`
	// DisableAudio is true for video-only output. Omitted when audio is selected.
	DisableAudio bool `json:"disable_audio"`
	// Framerate specifies the resolved video frames per second. Omitted for audio-only
	// output.
	Framerate float64 `json:"framerate"`
	// Gop specifies the resolved maximum keyframe interval in frames. Omitted for
	// audio-only output.
	Gop int64 `json:"gop"`
	// Height specifies the resolved video height in pixels. Omitted for audio-only
	// output.
	Height int64 `json:"height"`
	// Level specifies the accepted codec level when explicitly selected.
	Level int64 `json:"level"`
	// Maxrate specifies the resolved maximum video bitrate in bits per second,
	// including codec-level defaults. Omitted when no maximum applies.
	Maxrate int64 `json:"maxrate"`
	// Minrate specifies the accepted minimum video bitrate in bits per second, when
	// configured.
	Minrate int64 `json:"minrate"`
	// Movflags specifies MP4 packaging. MP4 video exports use faststart.
	Movflags string `json:"movflags"`
	// PerTitle specifies whether video rate control is selected during content
	// analysis. Omitted for audio-only output.
	PerTitle bool `json:"per_title"`
	// Pixfmt specifies the accepted video pixel format. Omitted for audio-only output.
	Pixfmt string `json:"pixfmt"`
	// Preset specifies the accepted codec preset. Omitted when the codec does not use
	// presets.
	Preset string `json:"preset"`
	// Profilev specifies the accepted codec profile when explicitly selected.
	Profilev string `json:"profilev"`
	// Quality specifies the accepted VP9 deadline policy: good, best or realtime.
	// Omitted for other codecs.
	Quality string `json:"quality"`
	// VideoBitrate specifies the accepted video bitrate in bits per second. Omitted
	// for CRF and per-title encoding.
	VideoBitrate int64 `json:"video_bitrate"`
	// Width specifies the resolved video width in pixels. Omitted for audio-only
	// output.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Duration        respjson.Field
		Seek            respjson.Field
		AudioBitrate    respjson.Field
		AudioSamplerate respjson.Field
		Bufsize         respjson.Field
		Channels        respjson.Field
		CPUUsed         respjson.Field
		Crf             respjson.Field
		DisableAudio    respjson.Field
		Framerate       respjson.Field
		Gop             respjson.Field
		Height          respjson.Field
		Level           respjson.Field
		Maxrate         respjson.Field
		Minrate         respjson.Field
		Movflags        respjson.Field
		PerTitle        respjson.Field
		Pixfmt          respjson.Field
		Preset          respjson.Field
		Profilev        respjson.Field
		Quality         respjson.Field
		VideoBitrate    respjson.Field
		Width           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscodingFormat) RawJSON() string { return r.JSON.raw }
func (r *TranscodingFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status reports processing or output expiry.
type TranscodingStatus string

const (
	TranscodingStatusPreparing TranscodingStatus = "preparing"
	TranscodingStatusReady     TranscodingStatus = "ready"
	TranscodingStatusFailed    TranscodingStatus = "failed"
	TranscodingStatusCancelled TranscodingStatus = "cancelled"
	TranscodingStatusExpired   TranscodingStatus = "expired"
)

// Storage contains the resolved output destination.
type TranscodingStorage struct {
	// Id identifies the output Storage.
	ID string `json:"id" api:"required"`
	// Path specifies the output file path relative to the Storage base_prefix saved at
	// creation.
	Path string `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscodingStorage) RawJSON() string { return r.JSON.raw }
func (r *TranscodingStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreatedVia specifies how the resource was created: api for API requests,
// dashboard for the Chunkify dashboard, mcp for MCP clients, and workflow for
// Workflow actions. Omitted when the creator is unknown.
type TranscodingCreatedVia string

const (
	TranscodingCreatedViaAPI       TranscodingCreatedVia = "api"
	TranscodingCreatedViaDashboard TranscodingCreatedVia = "dashboard"
	TranscodingCreatedViaMcp       TranscodingCreatedVia = "mcp"
	TranscodingCreatedViaWorkflow  TranscodingCreatedVia = "workflow"
)

// Error describes why processing failed. Omitted when no processing error exists.
type TranscodingError struct {
	// Code identifies the processing failure.
	Code string `json:"code" api:"required"`
	// Message describes the failure without internal worker or provider diagnostics.
	Message string `json:"message" api:"required"`
	// Stage identifies where processing failed.
	//
	// Any of "queued", "starting", "encoding", "finalizing", "uploading".
	Stage string `json:"stage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Stage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscodingError) RawJSON() string { return r.JSON.raw }
func (r *TranscodingError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProcessingStage specifies the current activity while the status is preparing.
// Omitted otherwise.
type TranscodingProcessingStage string

const (
	TranscodingProcessingStageQueued     TranscodingProcessingStage = "queued"
	TranscodingProcessingStageStarting   TranscodingProcessingStage = "starting"
	TranscodingProcessingStageEncoding   TranscodingProcessingStage = "encoding"
	TranscodingProcessingStageFinalizing TranscodingProcessingStage = "finalizing"
	TranscodingProcessingStageUploading  TranscodingProcessingStage = "uploading"
)

type TranscodingNewParams struct {
	// Format specifies the requested encoding settings.
	Format TranscodingNewParamsFormatUnion `json:"format,omitzero" api:"required"`
	// SourceId identifies a Source in this Project.
	SourceID string `json:"source_id" api:"required"`
	// Metadata attaches string values to the Transcoding. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Storage overrides the Project output destination. Source storage remains
	// independent.
	Storage TranscodingNewParamsStorage `json:"storage,omitzero"`
	paramObj
}

func (r TranscodingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TranscodingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscodingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TranscodingNewParamsFormatUnion struct {
	OfMP4H264 *TranscodingNewParamsFormatMP4H264 `json:",omitzero,inline"`
	OfMP4H265 *TranscodingNewParamsFormatMP4H265 `json:",omitzero,inline"`
	OfMP4Av1  *TranscodingNewParamsFormatMP4Av1  `json:",omitzero,inline"`
	OfWebmVp9 *TranscodingNewParamsFormatWebmVp9 `json:",omitzero,inline"`
	OfM4aAac  *TranscodingNewParamsFormatM4aAac  `json:",omitzero,inline"`
	paramUnion
}

func (u TranscodingNewParamsFormatUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMP4H264,
		u.OfMP4H265,
		u.OfMP4Av1,
		u.OfWebmVp9,
		u.OfM4aAac)
}
func (u *TranscodingNewParamsFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TranscodingNewParamsFormatUnion) asAny() any {
	if !param.IsOmitted(u.OfMP4H264) {
		return u.OfMP4H264
	} else if !param.IsOmitted(u.OfMP4H265) {
		return u.OfMP4H265
	} else if !param.IsOmitted(u.OfMP4Av1) {
		return u.OfMP4Av1
	} else if !param.IsOmitted(u.OfWebmVp9) {
		return u.OfWebmVp9
	} else if !param.IsOmitted(u.OfM4aAac) {
		return u.OfM4aAac
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetX264Keyint() *int64 {
	if vt := u.OfMP4H264; vt != nil && vt.X264Keyint.Valid() {
		return &vt.X264Keyint.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetX265Keyint() *int64 {
	if vt := u.OfMP4H265; vt != nil && vt.X265Keyint.Valid() {
		return &vt.X265Keyint.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetCPUUsed() *string {
	if vt := u.OfWebmVp9; vt != nil {
		return &vt.CPUUsed
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetMinrate() *int64 {
	if vt := u.OfWebmVp9; vt != nil && vt.Minrate.Valid() {
		return &vt.Minrate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetQuality() *string {
	if vt := u.OfWebmVp9; vt != nil {
		return &vt.Quality
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetID() *string {
	if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfWebmVp9; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfM4aAac; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetAudioBitrate() *int64 {
	if vt := u.OfMP4H264; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfMP4Av1; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	} else if vt := u.OfM4aAac; vt != nil && vt.AudioBitrate.Valid() {
		return &vt.AudioBitrate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetAudioSamplerate() *int64 {
	if vt := u.OfMP4H264; vt != nil {
		return (*int64)(&vt.AudioSamplerate)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*int64)(&vt.AudioSamplerate)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*int64)(&vt.AudioSamplerate)
	} else if vt := u.OfWebmVp9; vt != nil {
		return (*int64)(&vt.AudioSamplerate)
	} else if vt := u.OfM4aAac; vt != nil {
		return (*int64)(&vt.AudioSamplerate)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetChannels() *int64 {
	if vt := u.OfMP4H264; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfWebmVp9; vt != nil {
		return (*int64)(&vt.Channels)
	} else if vt := u.OfM4aAac; vt != nil {
		return (*int64)(&vt.Channels)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetCrf() *int64 {
	if vt := u.OfMP4H264; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfMP4Av1; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Crf.Valid() {
		return &vt.Crf.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetDuration() *float64 {
	if vt := u.OfMP4H264; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfMP4Av1; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	} else if vt := u.OfM4aAac; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetLevel() *int64 {
	if vt := u.OfMP4H264; vt != nil {
		return (*int64)(&vt.Level)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*int64)(&vt.Level)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*int64)(&vt.Level)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetMovflags() *string {
	if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.Movflags)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.Movflags)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.Movflags)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetPixfmt() *string {
	if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.Pixfmt)
	} else if vt := u.OfWebmVp9; vt != nil {
		return (*string)(&vt.Pixfmt)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetPreset() *string {
	if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.Preset)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.Preset)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.Preset)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetProfilev() *string {
	if vt := u.OfMP4H264; vt != nil {
		return (*string)(&vt.Profilev)
	} else if vt := u.OfMP4H265; vt != nil {
		return (*string)(&vt.Profilev)
	} else if vt := u.OfMP4Av1; vt != nil {
		return (*string)(&vt.Profilev)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TranscodingNewParamsFormatUnion) GetSeek() *float64 {
	if vt := u.OfMP4H264; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfMP4H265; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfMP4Av1; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfWebmVp9; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	} else if vt := u.OfM4aAac; vt != nil && vt.Seek.Valid() {
		return &vt.Seek.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[TranscodingNewParamsFormatUnion](
		"id",
		apijson.Discriminator[TranscodingNewParamsFormatMP4H264]("mp4_h264"),
		apijson.Discriminator[TranscodingNewParamsFormatMP4H265]("mp4_h265"),
		apijson.Discriminator[TranscodingNewParamsFormatMP4Av1]("mp4_av1"),
		apijson.Discriminator[TranscodingNewParamsFormatWebmVp9]("webm_vp9"),
		apijson.Discriminator[TranscodingNewParamsFormatM4aAac]("m4a_aac"),
	)
}

// The property ID is required.
type TranscodingNewParamsFormatMP4H264 struct {
	// AudioBitrate specifies the output audio bitrate in bits per second. Defaults
	// to 128000.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// Duration specifies the requested output length in seconds. Defaults to the
	// remaining source duration and is capped at the source end. Audio-only output
	// must resolve to at least 0.001 seconds after trimming and sample rounding. The
	// accepted interval determines the fixed price.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Seek specifies the start in source seconds. Defaults to zero and must precede
	// the source end. The accepted interval uses the output frame or sample timeline.
	// Audio-only output must retain at least 0.001 seconds after sample rounding.
	Seek param.Opt[float64] `json:"seek,omitzero"`
	// X264KeyInt specifies the maximum number of frames between keyframes for H.264
	// encoding. Range: 1 to 960. Higher values can improve compression but may affect
	// seeking.
	X264Keyint param.Opt[int64] `json:"x264_keyint,omitzero"`
	// AudioSamplerate specifies the output sample rate in Hz. Defaults to 48000. WebM
	// Opus audio uses 48000.
	//
	// Any of 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000,
	// 88200, 96000.
	AudioSamplerate int64 `json:"audio_samplerate,omitzero"`
	// Channels specifies the output channel count. Defaults to stereo. Six channels
	// means 5.1; eight means 7.1.
	//
	// Any of 1, 2, 6, 8.
	Channels int64 `json:"channels,omitzero"`
	// Level specifies the H.264 decoder level using its two-digit representation, such
	// as 41 for level 4.1. The dimensions, frame rate and rate controls must fit this
	// level and the effective profile. Omitted rate limits resolve to level-compatible
	// defaults.
	//
	// Any of 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51.
	Level int64 `json:"level,omitzero"`
	// Movflags specifies MP4 packaging. Defaults to faststart, which places playback
	// metadata before the media.
	//
	// Any of "faststart".
	Movflags string `json:"movflags,omitzero"`
	// Pixfmt specifies the output pixel format. Defaults to yuv420p. The selected
	// profile must support it.
	//
	// Any of "yuv420p", "yuv422p", "yuv444p", "yuv420p10le", "yuv422p10le",
	// "yuv444p10le".
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
	// Id specifies the output format.
	//
	// This field can be elided, and will marshal its zero value as "mp4_h264".
	ID constant.MP4H264 `json:"id" default:"mp4_h264"`
	paramObj
}

func (r TranscodingNewParamsFormatMP4H264) MarshalJSON() (data []byte, err error) {
	type shadow TranscodingNewParamsFormatMP4H264
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscodingNewParamsFormatMP4H264) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H264](
		"audio_samplerate", 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H264](
		"channels", 1, 2, 6, 8,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H264](
		"level", 10, 11, 12, 13, 20, 21, 22, 30, 31, 32, 40, 41, 42, 50, 51,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H264](
		"movflags", "faststart",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H264](
		"pixfmt", "yuv420p", "yuv422p", "yuv444p", "yuv420p10le", "yuv422p10le", "yuv444p10le",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H264](
		"preset", "ultrafast", "superfast", "veryfast", "faster", "fast", "medium",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H264](
		"profilev", "baseline", "main", "high", "high10", "high422", "high444",
	)
}

// The property ID is required.
type TranscodingNewParamsFormatMP4H265 struct {
	// AudioBitrate specifies the output audio bitrate in bits per second. Defaults
	// to 128000.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// Duration specifies the requested output length in seconds. Defaults to the
	// remaining source duration and is capped at the source end. Audio-only output
	// must resolve to at least 0.001 seconds after trimming and sample rounding. The
	// accepted interval determines the fixed price.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Seek specifies the start in source seconds. Defaults to zero and must precede
	// the source end. The accepted interval uses the output frame or sample timeline.
	// Audio-only output must retain at least 0.001 seconds after sample rounding.
	Seek param.Opt[float64] `json:"seek,omitzero"`
	// X265KeyInt specifies the maximum number of frames between keyframes for H.265
	// encoding. Range: 1 to 300. Higher values can improve compression but may affect
	// seeking.
	X265Keyint param.Opt[int64] `json:"x265_keyint,omitzero"`
	// AudioSamplerate specifies the output sample rate in Hz. Defaults to 48000. WebM
	// Opus audio uses 48000.
	//
	// Any of 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000,
	// 88200, 96000.
	AudioSamplerate int64 `json:"audio_samplerate,omitzero"`
	// Channels specifies the output channel count. Defaults to stereo. Six channels
	// means 5.1; eight means 7.1.
	//
	// Any of 1, 2, 6, 8.
	Channels int64 `json:"channels,omitzero"`
	// Level specifies the HEVC level using its two-digit representation, such as 31
	// for level 3.1. The dimensions, frame rate and rate controls must fit this level.
	// Omitted rate limits resolve to level-compatible defaults.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level,omitzero"`
	// Movflags specifies MP4 packaging. Defaults to faststart, which places playback
	// metadata before the media.
	//
	// Any of "faststart".
	Movflags string `json:"movflags,omitzero"`
	// Pixfmt specifies the output pixel format. Defaults to yuv420p. The selected
	// profile must support it.
	//
	// Any of "yuv420p", "yuv422p", "yuv444p", "yuv420p10le", "yuv422p10le",
	// "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv444p12le".
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
	// Profilev selects the HEVC profile. It must support the requested pixel format.
	//
	// Any of "main", "main10", "main12", "main422-10", "main422-12", "main444-8",
	// "main444-10", "main444-12".
	Profilev string `json:"profilev,omitzero"`
	// Id specifies the output format.
	//
	// This field can be elided, and will marshal its zero value as "mp4_h265".
	ID constant.MP4H265 `json:"id" default:"mp4_h265"`
	paramObj
}

func (r TranscodingNewParamsFormatMP4H265) MarshalJSON() (data []byte, err error) {
	type shadow TranscodingNewParamsFormatMP4H265
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscodingNewParamsFormatMP4H265) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H265](
		"audio_samplerate", 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H265](
		"channels", 1, 2, 6, 8,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H265](
		"level", 30, 31, 41,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H265](
		"movflags", "faststart",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H265](
		"pixfmt", "yuv420p", "yuv422p", "yuv444p", "yuv420p10le", "yuv422p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv444p12le",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H265](
		"preset", "ultrafast", "superfast", "veryfast", "faster", "fast", "medium",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4H265](
		"profilev", "main", "main10", "main12", "main422-10", "main422-12", "main444-8", "main444-10", "main444-12",
	)
}

// The property ID is required.
type TranscodingNewParamsFormatMP4Av1 struct {
	// AudioBitrate specifies the output audio bitrate in bits per second. Defaults
	// to 128000.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 16 to 63. Recommended
	// values: 16-35 for high quality, 35-45 for good quality, 45-63 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// Duration specifies the requested output length in seconds. Defaults to the
	// remaining source duration and is capped at the source end. Audio-only output
	// must resolve to at least 0.001 seconds after trimming and sample rounding. The
	// accepted interval determines the fixed price.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Seek specifies the start in source seconds. Defaults to zero and must precede
	// the source end. The accepted interval uses the output frame or sample timeline.
	// Audio-only output must retain at least 0.001 seconds after sample rounding.
	Seek param.Opt[float64] `json:"seek,omitzero"`
	// AudioSamplerate specifies the output sample rate in Hz. Defaults to 48000. WebM
	// Opus audio uses 48000.
	//
	// Any of 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000,
	// 88200, 96000.
	AudioSamplerate int64 `json:"audio_samplerate,omitzero"`
	// Channels specifies the output channel count. Defaults to stereo. Six channels
	// means 5.1; eight means 7.1.
	//
	// Any of 1, 2, 6, 8.
	Channels int64 `json:"channels,omitzero"`
	// Level specifies the AV1 level using its two-digit representation, such as 31 for
	// level 3.1. The dimensions, frame rate and rate controls must fit this level.
	// Omitted rate limits resolve to level-compatible defaults.
	//
	// Any of 30, 31, 41.
	Level int64 `json:"level,omitzero"`
	// Movflags specifies MP4 packaging. Defaults to faststart, which places playback
	// metadata before the media.
	//
	// Any of "faststart".
	Movflags string `json:"movflags,omitzero"`
	// Pixfmt specifies the output pixel format. Defaults to yuv420p. The selected
	// profile must support it.
	//
	// Any of "yuv420p", "yuv420p10le".
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
	// Profilev selects the AV1 Main profile. Use pixfmt to select 8-bit or 10-bit
	// output.
	//
	// Any of "main".
	Profilev string `json:"profilev,omitzero"`
	// Id specifies the output format.
	//
	// This field can be elided, and will marshal its zero value as "mp4_av1".
	ID constant.MP4Av1 `json:"id" default:"mp4_av1"`
	paramObj
}

func (r TranscodingNewParamsFormatMP4Av1) MarshalJSON() (data []byte, err error) {
	type shadow TranscodingNewParamsFormatMP4Av1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscodingNewParamsFormatMP4Av1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4Av1](
		"audio_samplerate", 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4Av1](
		"channels", 1, 2, 6, 8,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4Av1](
		"level", 30, 31, 41,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4Av1](
		"movflags", "faststart",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4Av1](
		"pixfmt", "yuv420p", "yuv420p10le",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4Av1](
		"preset", "6", "7", "8", "9", "10", "11", "12", "13",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatMP4Av1](
		"profilev", "main",
	)
}

// The property ID is required.
type TranscodingNewParamsFormatWebmVp9 struct {
	// AudioBitrate specifies the output audio bitrate in bits per second. Defaults
	// to 128000.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Crf (Constant Rate Factor) controls the quality of the output video. Lower
	// values mean better quality but larger file size. Range: 15 to 35. Recommended
	// values: 18-28 for high quality, 23-28 for good quality, 28-35 for acceptable
	// quality.
	Crf param.Opt[int64] `json:"crf,omitzero"`
	// Duration specifies the requested output length in seconds. Defaults to the
	// remaining source duration and is capped at the source end. Audio-only output
	// must resolve to at least 0.001 seconds after trimming and sample rounding. The
	// accepted interval determines the fixed price.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Minrate selects VP9 constant bitrate when equal to video_bitrate and maxrate.
	// Other minimum-rate combinations are unsupported.
	Minrate param.Opt[int64] `json:"minrate,omitzero"`
	// Seek specifies the start in source seconds. Defaults to zero and must precede
	// the source end. The accepted interval uses the output frame or sample timeline.
	// Audio-only output must retain at least 0.001 seconds after sample rounding.
	Seek param.Opt[float64] `json:"seek,omitzero"`
	// AudioSamplerate specifies the output sample rate in Hz. Defaults to 48000. WebM
	// Opus audio uses 48000.
	//
	// Any of 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000,
	// 88200, 96000.
	AudioSamplerate int64 `json:"audio_samplerate,omitzero"`
	// Channels specifies the output channel count. Defaults to stereo. Six channels
	// means 5.1; eight means 7.1.
	//
	// Any of 1, 2, 6, 8.
	Channels int64 `json:"channels,omitzero"`
	// CpuUsed specifies the CPU usage level for VP9 encoding. Range: 0 to 8. Lower
	// values mean better quality but slower encoding, higher values mean faster
	// encoding but lower quality. Recommended values: 0-2 for high quality, 2-4 for
	// good quality, 4-6 for balanced, 6-8 for speed
	//
	// Any of "0", "1", "2", "3", "4", "5", "6", "7", "8".
	CPUUsed string `json:"cpu_used,omitzero"`
	// Pixfmt specifies the output pixel format. Defaults to yuv420p. The selected
	// profile must support it.
	//
	// Any of "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuv420p10le", "yuv422p10le",
	// "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le",
	// "yuv444p12le".
	Pixfmt string `json:"pixfmt,omitzero"`
	// Quality specifies the VP9 encoding quality preset. Valid values:
	//
	// - good: Balanced quality preset, good for most applications
	// - best: Best quality preset, slower encoding
	// - realtime: Fast encoding preset, suitable for live streaming
	//
	// Any of "good", "best", "realtime".
	Quality string `json:"quality,omitzero"`
	// Id specifies the output format.
	//
	// This field can be elided, and will marshal its zero value as "webm_vp9".
	ID constant.WebmVp9 `json:"id" default:"webm_vp9"`
	paramObj
}

func (r TranscodingNewParamsFormatWebmVp9) MarshalJSON() (data []byte, err error) {
	type shadow TranscodingNewParamsFormatWebmVp9
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscodingNewParamsFormatWebmVp9) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatWebmVp9](
		"audio_samplerate", 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatWebmVp9](
		"channels", 1, 2, 6, 8,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatWebmVp9](
		"cpu_used", "0", "1", "2", "3", "4", "5", "6", "7", "8",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatWebmVp9](
		"pixfmt", "yuv420p", "yuv422p", "yuv440p", "yuv444p", "yuv420p10le", "yuv422p10le", "yuv440p10le", "yuv444p10le", "yuv420p12le", "yuv422p12le", "yuv440p12le", "yuv444p12le",
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatWebmVp9](
		"quality", "good", "best", "realtime",
	)
}

// The property ID is required.
type TranscodingNewParamsFormatM4aAac struct {
	// AudioBitrate specifies the output audio bitrate in bits per second. Defaults
	// to 128000.
	AudioBitrate param.Opt[int64] `json:"audio_bitrate,omitzero"`
	// Duration specifies the requested output length in seconds. Defaults to the
	// remaining source duration and is capped at the source end. Audio-only output
	// must resolve to at least 0.001 seconds after trimming and sample rounding. The
	// accepted interval determines the fixed price.
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Seek specifies the start in source seconds. Defaults to zero and must precede
	// the source end. The accepted interval uses the output frame or sample timeline.
	// Audio-only output must retain at least 0.001 seconds after sample rounding.
	Seek param.Opt[float64] `json:"seek,omitzero"`
	// AudioSamplerate specifies the output sample rate in Hz. Defaults to 48000. WebM
	// Opus audio uses 48000.
	//
	// Any of 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000,
	// 88200, 96000.
	AudioSamplerate int64 `json:"audio_samplerate,omitzero"`
	// Channels specifies the output channel count. Defaults to stereo. Six channels
	// means 5.1; eight means 7.1.
	//
	// Any of 1, 2, 6, 8.
	Channels int64 `json:"channels,omitzero"`
	// Id specifies an AAC M4A export.
	//
	// This field can be elided, and will marshal its zero value as "m4a_aac".
	ID constant.M4aAac `json:"id" default:"m4a_aac"`
	paramObj
}

func (r TranscodingNewParamsFormatM4aAac) MarshalJSON() (data []byte, err error) {
	type shadow TranscodingNewParamsFormatM4aAac
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscodingNewParamsFormatM4aAac) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatM4aAac](
		"audio_samplerate", 8000, 11025, 12000, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000,
	)
	apijson.RegisterFieldValidator[TranscodingNewParamsFormatM4aAac](
		"channels", 1, 2, 6, 8,
	)
}

// Storage overrides the Project output destination. Source storage remains
// independent.
type TranscodingNewParamsStorage struct {
	// Id specifies a Storage in this Project. Defaults to the Project output storage.
	ID param.Opt[string] `json:"id,omitzero"`
	// Path specifies the output object path relative to the selected Storage
	// base_prefix. Required for customer storage; generated for Chunkify temporary
	// storage. The extension must match the resolved output format. The combined path
	// may contain at most 1024 bytes.
	Path param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r TranscodingNewParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow TranscodingNewParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscodingNewParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Project-scoped Transcoding response.
type TranscodingNewResponseEnvelope struct {
	// Data contains the Transcoding.
	Data Transcoding `json:"data" api:"required"`
	// Status is success for a successful request.
	Status constant.Success `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscodingNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TranscodingNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Project-scoped Transcoding response.
type TranscodingGetResponseEnvelope struct {
	// Data contains the Transcoding.
	Data Transcoding `json:"data" api:"required"`
	// Status is success for a successful request.
	Status constant.Success `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscodingGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TranscodingGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscodingListParams struct {
	// Id filters by an exact Transcoding ID. Omit to include all matching resources.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Maximum number of Transcodings to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of Transcodings to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter Transcodings linked to this Source. Omit to include Transcodings from all
	// Sources in the Project.
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// Filter by output Storage ID.
	StorageID param.Opt[string]            `query:"storage_id,omitzero" json:"-"`
	Created   TranscodingListParamsCreated `query:"created,omitzero" json:"-"`
	// Filter by Transcoding status.
	//
	// Any of "preparing", "ready", "failed", "cancelled", "expired".
	Status TranscodingListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TranscodingListParams]'s query parameters as `url.Values`.
func (r TranscodingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TranscodingListParamsCreated struct {
	// Created at or after this Unix timestamp in seconds.
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Created at or before this Unix timestamp in seconds.
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	// Creation-time ordering. Defaults to newest first.
	//
	// Any of "asc", "desc".
	Sort string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TranscodingListParamsCreated]'s query parameters as
// `url.Values`.
func (r TranscodingListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by Transcoding status.
type TranscodingListParamsStatus string

const (
	TranscodingListParamsStatusPreparing TranscodingListParamsStatus = "preparing"
	TranscodingListParamsStatusReady     TranscodingListParamsStatus = "ready"
	TranscodingListParamsStatusFailed    TranscodingListParamsStatus = "failed"
	TranscodingListParamsStatusCancelled TranscodingListParamsStatus = "cancelled"
	TranscodingListParamsStatusExpired   TranscodingListParamsStatus = "expired"
)

type TranscodingDeleteParams struct {
	// DeleteFiles specifies whether to permanently delete this Transcoding's output
	// files at their saved storage paths. Defaults to false. The original Source and
	// unrelated files are preserved. If deletion fails, records remain for retry, but
	// some files may already be deleted. Managed dashboard images are removed with the
	// resource; temporary outputs still expire after 24 hours when this option is
	// false.
	DeleteFiles param.Opt[bool] `query:"delete_files,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TranscodingDeleteParams]'s query parameters as
// `url.Values`.
func (r TranscodingDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
