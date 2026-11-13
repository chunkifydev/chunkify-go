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

// Automatic video-on-demand Assets and independent H.264 MP4 static renditions.
//
// AssetService contains methods and other services that help with interacting with
// the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
	// Automatic video-on-demand Assets and independent H.264 MP4 static renditions.
	Outputs AssetOutputService
	// Automatic video-on-demand Assets and independent H.264 MP4 static renditions.
	Workflow AssetWorkflowService
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r AssetService) {
	r = AssetService{}
	r.Options = opts
	r.Outputs = NewAssetOutputService(opts...)
	r.Workflow = NewAssetWorkflowService(opts...)
	return
}

// Create an Asset from a video Source. Processing starts asynchronously.
func (r *AssetService) New(ctx context.Context, body AssetNewParams, opts ...option.RequestOption) (res *Asset, err error) {
	var env AssetNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve an Asset and its streaming, static, poster, and timeline outputs.
func (r *AssetService) Get(ctx context.Context, assetID string, opts ...option.RequestOption) (res *Asset, err error) {
	var env AssetGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update the title, metadata, or public preview access. Omitted fields remain
// unchanged.
func (r *AssetService) Update(ctx context.Context, assetID string, body AssetUpdateParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return err
	}
	path := fmt.Sprintf("assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// List Assets in the Project, with filtering and pagination.
func (r *AssetService) List(ctx context.Context, query AssetListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Asset], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "assets"
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

// List Assets in the Project, with filtering and pagination.
func (r *AssetService) ListAutoPaging(ctx context.Context, query AssetListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Asset] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete an Asset after all processing finishes. Files remain in Storage unless
// delete_files is true.
func (r *AssetService) Delete(ctx context.Context, assetID string, body AssetDeleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return err
	}
	path := fmt.Sprintf("assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

type Asset struct {
	// Unique Asset ID.
	ID string `json:"id" api:"required"`
	// Whether the Asset can accept new static rendition requests. Creation can still
	// fail if the original video is no longer accessible.
	CanAddStaticRenditions bool `json:"can_add_static_renditions" api:"required"`
	// When the Asset was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Maximum streaming resolution selected at creation. Does not limit static
	// renditions.
	//
	// Any of "1080p", "2160p".
	MaxResolution AssetMaxResolution `json:"max_resolution" api:"required"`
	// String metadata attached to the Asset.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Minimum streaming resolution selected at creation. auto means no minimum;
	// smaller sources are not upscaled.
	//
	// Any of "auto", "240p", "360p", "540p", "720p", "1080p", "1440p", "2160p".
	MinResolution AssetMinResolution `json:"min_resolution" api:"required"`
	// Whether HLS playback is available, including while the Asset is preparing.
	PlaybackReady bool `json:"playback_ready" api:"required"`
	// Allow public preview access without an API token. Does not control CDN or static
	// download access. Disabling it does not revoke signed URLs already issued.
	PreviewEnabled bool `json:"preview_enabled" api:"required"`
	// Original media metadata saved at Asset creation and preserved after Source
	// deletion.
	SourceMetadata AssetSourceMetadata `json:"source_metadata" api:"required"`
	// Status specifies the Asset preparation state. ready means the video is prepared
	// for playback, with its poster and timeline previews available. Static renditions
	// and captions have independent statuses.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status AssetStatus `json:"status" api:"required"`
	// ID of the output Storage.
	StorageID string `json:"storage_id" api:"required"`
	// Asset title.
	Title string `json:"title" api:"required"`
	// When the Asset was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Selected Workflow ID, or null.
	WorkflowID string `json:"workflow_id" api:"required"`
	// Pinned Workflow revision, or null.
	WorkflowRevision int64 `json:"workflow_revision" api:"required"`
	// Unsigned CDN URL of the HLS master playlist. Available when playback_ready is
	// true and Storage has a CDN URL.
	CdnURL string `json:"cdn_url" format:"uri"`
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
	CreatedVia AssetCreatedVia `json:"created_via"`
	Error      AssetError      `json:"error"`
	// When managed outputs expire, 24 hours after readiness. Omitted for
	// customer-connected Storage.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Path of the HLS master playlist in Storage. Available when playback_ready is
	// true.
	PlaybackPath string `json:"playback_path"`
	// When HLS playback first became available.
	PlaybackReadyAt time.Time `json:"playback_ready_at" format:"date-time"`
	// Poster image, available once generated.
	Poster AssetPoster `json:"poster"`
	// Public HLS preview URL. Available when playback_ready and preview_enabled are
	// true. No API token is required.
	PreviewURL string `json:"preview_url" format:"uri"`
	// When all required Asset outputs finished processing.
	ReadyAt time.Time `json:"ready_at" format:"date-time"`
	// Size specifies the total bytes of available streaming, playback preview, poster,
	// and timeline files. Excludes static renditions. Omitted until the Asset is
	// ready.
	Size int64 `json:"size"`
	// Original Source ID. Omitted after Source deletion.
	SourceID string `json:"source_id"`
	// Timeline thumbnails. The WebVTT file maps video timestamps to image URLs and
	// crop coordinates.
	TimelinePreviews AssetTimelinePreviews `json:"timeline_previews"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CanAddStaticRenditions respjson.Field
		CreatedAt              respjson.Field
		MaxResolution          respjson.Field
		Metadata               respjson.Field
		MinResolution          respjson.Field
		PlaybackReady          respjson.Field
		PreviewEnabled         respjson.Field
		SourceMetadata         respjson.Field
		Status                 respjson.Field
		StorageID              respjson.Field
		Title                  respjson.Field
		UpdatedAt              respjson.Field
		WorkflowID             respjson.Field
		WorkflowRevision       respjson.Field
		CdnURL                 respjson.Field
		CreatedByTokenID       respjson.Field
		CreatedByUserID        respjson.Field
		CreatedVia             respjson.Field
		Error                  respjson.Field
		ExpiresAt              respjson.Field
		PlaybackPath           respjson.Field
		PlaybackReadyAt        respjson.Field
		Poster                 respjson.Field
		PreviewURL             respjson.Field
		ReadyAt                respjson.Field
		Size                   respjson.Field
		SourceID               respjson.Field
		TimelinePreviews       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Asset) RawJSON() string { return r.JSON.raw }
func (r *Asset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum streaming resolution selected at creation. Does not limit static
// renditions.
type AssetMaxResolution string

const (
	AssetMaxResolution1080p AssetMaxResolution = "1080p"
	AssetMaxResolution2160p AssetMaxResolution = "2160p"
)

// Minimum streaming resolution selected at creation. auto means no minimum;
// smaller sources are not upscaled.
type AssetMinResolution string

const (
	AssetMinResolutionAuto  AssetMinResolution = "auto"
	AssetMinResolution240p  AssetMinResolution = "240p"
	AssetMinResolution360p  AssetMinResolution = "360p"
	AssetMinResolution540p  AssetMinResolution = "540p"
	AssetMinResolution720p  AssetMinResolution = "720p"
	AssetMinResolution1080p AssetMinResolution = "1080p"
	AssetMinResolution1440p AssetMinResolution = "1440p"
	AssetMinResolution2160p AssetMinResolution = "2160p"
)

// Original media metadata saved at Asset creation and preserved after Source
// deletion.
type AssetSourceMetadata struct {
	// Original audio bitrate in bits per second, or zero when unavailable.
	AudioBitrate int64 `json:"audio_bitrate" api:"required"`
	// Original audio codec, or an empty string when there is no audio.
	AudioCodec string `json:"audio_codec" api:"required"`
	// Original audio sample rate in Hz, or zero when unavailable.
	AudioSamplerate int64 `json:"audio_samplerate" api:"required"`
	// Original duration in seconds.
	Duration float64 `json:"duration" api:"required"`
	// Original video height in pixels.
	Height int64 `json:"height" api:"required"`
	// Original file size in bytes.
	Size int64 `json:"size" api:"required"`
	// Original video bitrate in bits per second, or zero when unknown.
	VideoBitrate int64 `json:"video_bitrate" api:"required"`
	// Original video codec.
	VideoCodec string `json:"video_codec" api:"required"`
	// Original video frame rate in frames per second.
	VideoFramerate float64 `json:"video_framerate" api:"required"`
	// Original video width in pixels.
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioBitrate    respjson.Field
		AudioCodec      respjson.Field
		AudioSamplerate respjson.Field
		Duration        respjson.Field
		Height          respjson.Field
		Size            respjson.Field
		VideoBitrate    respjson.Field
		VideoCodec      respjson.Field
		VideoFramerate  respjson.Field
		Width           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetSourceMetadata) RawJSON() string { return r.JSON.raw }
func (r *AssetSourceMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status specifies the Asset preparation state. ready means the video is prepared
// for playback, with its poster and timeline previews available. Static renditions
// and captions have independent statuses.
type AssetStatus string

const (
	AssetStatusPreparing AssetStatus = "preparing"
	AssetStatusReady     AssetStatus = "ready"
	AssetStatusFailed    AssetStatus = "failed"
	AssetStatusExpired   AssetStatus = "expired"
)

// CreatedVia specifies how the resource was created: api for API requests,
// dashboard for the Chunkify dashboard, mcp for MCP clients, and workflow for
// Workflow actions. Omitted when the creator is unknown.
type AssetCreatedVia string

const (
	AssetCreatedViaAPI       AssetCreatedVia = "api"
	AssetCreatedViaDashboard AssetCreatedVia = "dashboard"
	AssetCreatedViaMcp       AssetCreatedVia = "mcp"
	AssetCreatedViaWorkflow  AssetCreatedVia = "workflow"
)

type AssetError struct {
	// Machine-readable processing error code.
	Code string `json:"code" api:"required"`
	// Description of the processing failure.
	Message string `json:"message" api:"required"`
	// Processing stage that failed.
	//
	// Any of "ingest", "analysis", "encoding", "packaging", "upload", "validation",
	// "cleanup".
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
func (r AssetError) RawJSON() string { return r.JSON.raw }
func (r *AssetError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Poster image, available once generated.
type AssetPoster struct {
	// Image height in pixels.
	Height int64 `json:"height" api:"required"`
	// Image path in Storage.
	Path string `json:"path" api:"required"`
	// Image width in pixels.
	Width int64 `json:"width" api:"required"`
	// Unsigned image CDN URL. Available once the poster exists and Storage has a CDN
	// URL.
	CdnURL string `json:"cdn_url" format:"uri"`
	// Public poster preview URL. Available when the Asset is ready and preview_enabled
	// is true.
	PreviewURL string `json:"preview_url" format:"uri"`
	// Image size in bytes.
	Size int64 `json:"size"`
	// Position of the selected frame in seconds.
	TimeSeconds float64 `json:"time_seconds"`
	// Signed image URL, valid for up to ten minutes. Available once the poster exists,
	// regardless of preview_enabled.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Path        respjson.Field
		Width       respjson.Field
		CdnURL      respjson.Field
		PreviewURL  respjson.Field
		Size        respjson.Field
		TimeSeconds respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetPoster) RawJSON() string { return r.JSON.raw }
func (r *AssetPoster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Timeline thumbnails. The WebVTT file maps video timestamps to image URLs and
// crop coordinates.
type AssetTimelinePreviews struct {
	// Thumbnail image layout.
	Format constant.Sprite `json:"format" default:"sprite"`
	// WebVTT file path in Storage.
	VttPath string `json:"vtt_path" api:"required"`
	// Unsigned WebVTT CDN URL. Available when the Asset is ready and Storage has a CDN
	// URL.
	CdnURL string `json:"cdn_url" format:"uri"`
	// Number of thumbnails.
	ImageCount int64 `json:"image_count"`
	// Target thumbnail sampling interval in seconds. The WebVTT file provides actual
	// cue timestamps.
	IntervalSeconds float64 `json:"interval_seconds"`
	// Public WebVTT preview URL with signed image references. Available when the Asset
	// is ready and preview_enabled is true.
	PreviewURL string `json:"preview_url" format:"uri"`
	// Combined WebVTT and sprite image size in bytes.
	Size int64 `json:"size"`
	// Number of sprite images containing the thumbnails.
	SpriteCount int64 `json:"sprite_count"`
	// Thumbnail height in pixels.
	ThumbnailHeight int64 `json:"thumbnail_height"`
	// Thumbnail width in pixels.
	ThumbnailWidth int64 `json:"thumbnail_width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format          respjson.Field
		VttPath         respjson.Field
		CdnURL          respjson.Field
		ImageCount      respjson.Field
		IntervalSeconds respjson.Field
		PreviewURL      respjson.Field
		Size            respjson.Field
		SpriteCount     respjson.Field
		ThumbnailHeight respjson.Field
		ThumbnailWidth  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetTimelinePreviews) RawJSON() string { return r.JSON.raw }
func (r *AssetTimelinePreviews) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetNewParams struct {
	// ID of the video Source to process.
	SourceID string `json:"source_id" api:"required"`
	// WorkflowId specifies the Workflow to select. Omit to use the Project default.
	// Set null to disable automatic actions. The current revision is pinned when
	// creation is accepted.
	WorkflowID param.Opt[string] `json:"workflow_id,omitzero"`
	// PathTemplate specifies the Asset output folder relative to Storage base_prefix.
	// Supports {assetID} and case-sensitive {metadata.<key>}; referenced values must
	// be non-empty. Ignored for Chunkify storage, which always uses {assetID}.
	// Overrides the Project default. Defaults to {assetID} when neither is set.
	// Maximum template size: 1024 UTF-8 bytes. Expanded prefix: 768 bytes including
	// base_prefix and trailing slash.
	PathTemplate param.Opt[string] `json:"path_template,omitzero"`
	// Allow public preview access without an API token. Omit to use the Project
	// default. Does not control CDN or static download access.
	PreviewEnabled param.Opt[bool] `json:"preview_enabled,omitzero"`
	// Asset title, up to 256 characters. Defaults to an empty string.
	Title param.Opt[string] `json:"title,omitzero"`
	// Maximum streaming resolution. Omit to use the Project default, or the plan
	// maximum if unset. 2160p requires Starter or above. Smaller sources are not
	// upscaled. Does not affect static renditions.
	//
	// Any of "1080p", "2160p".
	MaxResolution AssetNewParamsMaxResolution `json:"max_resolution,omitzero"`
	// String metadata with a maximum encoded JSON size of 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Minimum streaming resolution. Omit to use the Project default; auto removes the
	// minimum. Must not exceed max_resolution. Smaller sources are not upscaled. Does
	// not affect static renditions.
	//
	// Any of "auto", "240p", "360p", "540p", "720p", "1080p", "1440p", "2160p".
	MinResolution AssetNewParamsMinResolution `json:"min_resolution,omitzero"`
	// Output Storage. Omit to use the Project default. The path_template controls the
	// folder for generated output paths.
	Storage AssetNewParamsStorage `json:"storage,omitzero"`
	paramObj
}

func (r AssetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum streaming resolution. Omit to use the Project default, or the plan
// maximum if unset. 2160p requires Starter or above. Smaller sources are not
// upscaled. Does not affect static renditions.
type AssetNewParamsMaxResolution string

const (
	AssetNewParamsMaxResolution1080p AssetNewParamsMaxResolution = "1080p"
	AssetNewParamsMaxResolution2160p AssetNewParamsMaxResolution = "2160p"
)

// Minimum streaming resolution. Omit to use the Project default; auto removes the
// minimum. Must not exceed max_resolution. Smaller sources are not upscaled. Does
// not affect static renditions.
type AssetNewParamsMinResolution string

const (
	AssetNewParamsMinResolutionAuto  AssetNewParamsMinResolution = "auto"
	AssetNewParamsMinResolution240p  AssetNewParamsMinResolution = "240p"
	AssetNewParamsMinResolution360p  AssetNewParamsMinResolution = "360p"
	AssetNewParamsMinResolution540p  AssetNewParamsMinResolution = "540p"
	AssetNewParamsMinResolution720p  AssetNewParamsMinResolution = "720p"
	AssetNewParamsMinResolution1080p AssetNewParamsMinResolution = "1080p"
	AssetNewParamsMinResolution1440p AssetNewParamsMinResolution = "1440p"
	AssetNewParamsMinResolution2160p AssetNewParamsMinResolution = "2160p"
)

// Output Storage. Omit to use the Project default. The path_template controls the
// folder for generated output paths.
type AssetNewParamsStorage struct {
	// ID of output Storage in this Project. Omit to use the Project default. Applies
	// to all outputs, including static renditions added later.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r AssetNewParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow AssetNewParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetNewParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetNewResponseEnvelope struct {
	Data   Asset            `json:"data" api:"required"`
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
func (r AssetNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AssetNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetGetResponseEnvelope struct {
	Data   Asset            `json:"data" api:"required"`
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
func (r AssetGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AssetGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetUpdateParams struct {
	// Allow public preview access without an API token. Does not control CDN or static
	// download access. Disabling it does not revoke signed URLs already issued.
	PreviewEnabled param.Opt[bool] `json:"preview_enabled,omitzero"`
	// Asset title, up to 256 characters. Send an empty string to clear it; omit to
	// leave unchanged.
	Title param.Opt[string] `json:"title,omitzero"`
	// Replaces all metadata. Send {} to clear it; omit to leave unchanged. Maximum
	// encoded JSON size: 2,048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r AssetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetListParams struct {
	// Maximum number of Assets to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of Assets to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Case-insensitive search of Asset IDs and titles.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter Assets linked to this Source. Omit to include Assets from all Sources in
	// the Project.
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// Filter by output Storage ID.
	StorageID param.Opt[string]      `query:"storage_id,omitzero" json:"-"`
	Created   AssetListParamsCreated `query:"created,omitzero" json:"-"`
	// Repeat key:value pairs to match all supplied metadata conditions.
	Metadata []string `query:"metadata,omitzero" json:"-"`
	// Filter by Asset status.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status AssetListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetListParams]'s query parameters as `url.Values`.
func (r AssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssetListParamsCreated struct {
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

// URLQuery serializes [AssetListParamsCreated]'s query parameters as `url.Values`.
func (r AssetListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by Asset status.
type AssetListParamsStatus string

const (
	AssetListParamsStatusPreparing AssetListParamsStatus = "preparing"
	AssetListParamsStatusReady     AssetListParamsStatus = "ready"
	AssetListParamsStatusFailed    AssetListParamsStatus = "failed"
	AssetListParamsStatusExpired   AssetListParamsStatus = "expired"
)

type AssetDeleteParams struct {
	// Also delete all files under the Asset's storage prefix, including streaming and
	// static outputs. Other files and the original input are preserved. Defaults to
	// false, which preserves published media files. Private transcript text and saved
	// subtitle copies are always scheduled for removal. File deletion is irreversible.
	// If cleanup fails, records remain for retry, but files may already be deleted.
	DeleteFiles param.Opt[bool] `query:"delete_files,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetDeleteParams]'s query parameters as `url.Values`.
func (r AssetDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
