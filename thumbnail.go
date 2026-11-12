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

// ThumbnailService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThumbnailService] method instead.
type ThumbnailService struct {
	Options []option.RequestOption
}

// NewThumbnailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThumbnailService(opts ...option.RequestOption) (r ThumbnailService) {
	r = ThumbnailService{}
	r.Options = opts
	return
}

// Generate individual images or sprite sheets from a Source. Each accepted request
// creates a new resource. Processing continues after the response.
func (r *ThumbnailService) New(ctx context.Context, body ThumbnailNewParams, opts ...option.RequestOption) (res *Thumbnail, err error) {
	var env ThumbnailNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "thumbnails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve the accepted settings, published files and current state of one
// Thumbnail.
func (r *ThumbnailService) Get(ctx context.Context, thumbnailID string, opts ...option.RequestOption) (res *Thumbnail, err error) {
	var env ThumbnailGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if thumbnailID == "" {
		err = errors.New("missing required thumbnailId parameter")
		return nil, err
	}
	path := fmt.Sprintf("thumbnails/%s", thumbnailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List Thumbnails in the authenticated Project, newest first by default.
func (r *ThumbnailService) List(ctx context.Context, query ThumbnailListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Thumbnail], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "thumbnails"
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

// List Thumbnails in the authenticated Project, newest first by default.
func (r *ThumbnailService) ListAutoPaging(ctx context.Context, query ThumbnailListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Thumbnail] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete a ready, failed or expired Thumbnail after all workers have stopped. The
// Source and settled usage are preserved. Files remain unless delete_files is
// true; temporary files retain their original 24-hour expiry. Deletion remains
// available for inactive teams.
func (r *ThumbnailService) Delete(ctx context.Context, thumbnailID string, body ThumbnailDeleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if thumbnailID == "" {
		err = errors.New("missing required thumbnailId parameter")
		return err
	}
	path := fmt.Sprintf("thumbnails/%s", thumbnailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// One thumbnail generation request. At least one sample and its companion files
// must be published before the resource is ready.
type Thumbnail struct {
	// Id identifies this thumbnail request. Standalone requests use thm* IDs;
	// Asset-owned outputs use aop* IDs.
	ID string `json:"id" api:"required"`
	// AssetId identifies the owning Asset for Workflow outputs, or null for standalone
	// requests.
	AssetID string `json:"asset_id" api:"required"`
	// CreatedAt records when the request was accepted.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Files contains all published images or sheets and the WebVTT index. Empty while
	// preparing.
	Files []ThumbnailFile `json:"files" api:"required"`
	// Metadata contains the accepted string metadata.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Samples contains published thumbnails ordered by timestamp. Empty while
	// preparing. A ready resource may contain fewer samples than requested.
	Samples []ThumbnailSample `json:"samples" api:"required"`
	// Settings contains the accepted sampling and output settings.
	Settings ThumbnailSettings `json:"settings" api:"required"`
	// SourceId identifies the input Source, or null if it has been deleted.
	SourceID string `json:"source_id" api:"required"`
	// Status reports processing or output expiry.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status ThumbnailStatus `json:"status" api:"required"`
	// StorageId identifies the accepted output Storage.
	StorageID string `json:"storage_id" api:"required"`
	// StoragePath specifies the accepted filename template relative to the Storage
	// base_prefix, or null when Chunkify generated the file names.
	StoragePath string `json:"storage_path" api:"required"`
	// StoragePrefix specifies the complete output folder saved at creation. File and
	// sample paths are relative to this folder.
	StoragePrefix string `json:"storage_prefix" api:"required"`
	// UpdatedAt records the last lifecycle update.
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
	CreatedVia ThumbnailCreatedVia `json:"created_via"`
	// DeleteAt specifies when this standalone request is deleted: three days after its
	// temporary output expires, otherwise when the team's data retention period ends,
	// counted from when the request was accepted, and no sooner than three days after
	// it finished. It does not delete files in customer-owned Storage. Omitted while
	// preparing and for Asset-owned thumbnails, which follow their Asset.
	DeleteAt time.Time `json:"delete_at" format:"date-time"`
	// Error explains a failed request. Omitted otherwise.
	Error ThumbnailError `json:"error"`
	// ExpiresAt specifies when temporary output files expire. Omitted for customer
	// output Storage.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// ReadyAt records when the complete available set was published. Omitted before
	// publication.
	ReadyAt time.Time `json:"ready_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AssetID          respjson.Field
		CreatedAt        respjson.Field
		Files            respjson.Field
		Metadata         respjson.Field
		Samples          respjson.Field
		Settings         respjson.Field
		SourceID         respjson.Field
		Status           respjson.Field
		StorageID        respjson.Field
		StoragePath      respjson.Field
		StoragePrefix    respjson.Field
		UpdatedAt        respjson.Field
		CreatedByTokenID respjson.Field
		CreatedByUserID  respjson.Field
		CreatedVia       respjson.Field
		DeleteAt         respjson.Field
		Error            respjson.Field
		ExpiresAt        respjson.Field
		ReadyAt          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Thumbnail) RawJSON() string { return r.JSON.raw }
func (r *Thumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One published image, sprite sheet, or WebVTT file.
type ThumbnailFile struct {
	// Kind identifies an individual image, sprite sheet, or WebVTT index.
	//
	// Any of "image", "sprite", "vtt".
	Kind string `json:"kind" api:"required"`
	// Path specifies the file relative to storage_prefix.
	Path string `json:"path" api:"required"`
	// Size specifies the measured encoded file size in bytes.
	Size int64 `json:"size" api:"required"`
	// CdnUrl provides delivery through the configured CDN. Omitted without a CDN or
	// after expiry.
	CdnURL string `json:"cdn_url"`
	// Height specifies the measured image or complete sprite canvas height. Omitted
	// for WebVTT.
	Height int64 `json:"height"`
	// Url provides temporary authorized delivery. Omitted after expiry or when
	// unavailable.
	URL string `json:"url"`
	// Width specifies the measured image or complete sprite canvas width. Omitted for
	// WebVTT.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Path        respjson.Field
		Size        respjson.Field
		CdnURL      respjson.Field
		Height      respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThumbnailFile) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One successfully published requested sample. Repeated samples remain separate
// entries.
type ThumbnailSample struct {
	// Height specifies this image or tile height in pixels.
	Height int64 `json:"height" api:"required"`
	// Path identifies the image or sprite sheet in files.
	Path string `json:"path" api:"required"`
	// Timestamp specifies the measured Source-relative presentation time of the
	// selected frame in seconds. Compatible cached inputs use millisecond timestamps,
	// which can select a neighboring frame compared with the original Source.
	Timestamp float64 `json:"timestamp" api:"required"`
	// Width specifies this image or tile width in pixels.
	Width int64 `json:"width" api:"required"`
	// Crop specifies the tile origin in its sprite sheet. Omitted for individual
	// images; width and height give the tile dimensions.
	Crop ThumbnailSampleCrop `json:"crop"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Path        respjson.Field
		Timestamp   respjson.Field
		Width       respjson.Field
		Crop        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThumbnailSample) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crop specifies the tile origin in its sprite sheet. Omitted for individual
// images; width and height give the tile dimensions.
type ThumbnailSampleCrop struct {
	// X specifies the left edge in pixels.
	X int64 `json:"x" api:"required"`
	// Y specifies the top edge in pixels.
	Y int64 `json:"y" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThumbnailSampleCrop) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailSampleCrop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings contains the accepted sampling and output settings.
type ThumbnailSettings struct {
	// Format specifies the accepted image encoding.
	//
	// Any of "jpeg", "webp".
	Format string `json:"format" api:"required"`
	// Output specifies the accepted output mode.
	//
	// Any of "images", "sprites".
	Output string `json:"output" api:"required"`
	// Range contains the accepted sampling boundaries.
	Range ThumbnailSettingsRange `json:"range" api:"required"`
	// Size specifies the accepted maximum image or tile edge in pixels.
	Size int64 `json:"size" api:"required"`
	// Count specifies 1 to 100 thumbnails at the middle of equal sections of the
	// selected range.
	Count int64 `json:"count"`
	// Timestamps specifies 1 to 100 Source-relative target times in seconds. Every
	// time must be within the selected range. Repeated targets produce separate
	// outputs.
	Timestamps []float64 `json:"timestamps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		Output      respjson.Field
		Range       respjson.Field
		Size        respjson.Field
		Count       respjson.Field
		Timestamps  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThumbnailSettings) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Range contains the accepted sampling boundaries.
type ThumbnailSettingsRange struct {
	// End specifies the exclusive upper bound in seconds.
	End float64 `json:"end" api:"required"`
	// Start specifies the inclusive lower bound in seconds.
	Start float64 `json:"start" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThumbnailSettingsRange) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailSettingsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status reports processing or output expiry.
type ThumbnailStatus string

const (
	ThumbnailStatusPreparing ThumbnailStatus = "preparing"
	ThumbnailStatusReady     ThumbnailStatus = "ready"
	ThumbnailStatusFailed    ThumbnailStatus = "failed"
	ThumbnailStatusExpired   ThumbnailStatus = "expired"
)

// CreatedVia specifies how the resource was created: api for API requests,
// dashboard for the Chunkify dashboard, mcp for MCP clients, and workflow for
// Workflow actions. Omitted when the creator is unknown.
type ThumbnailCreatedVia string

const (
	ThumbnailCreatedViaAPI       ThumbnailCreatedVia = "api"
	ThumbnailCreatedViaDashboard ThumbnailCreatedVia = "dashboard"
	ThumbnailCreatedViaMcp       ThumbnailCreatedVia = "mcp"
	ThumbnailCreatedViaWorkflow  ThumbnailCreatedVia = "workflow"
)

// Error explains a failed request. Omitted otherwise.
type ThumbnailError struct {
	// Code identifies the processing failure.
	Code string `json:"code" api:"required"`
	// Message explains the processing failure.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThumbnailError) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output Storage selection, independent of Source storage.
type ThumbnailCreateStorageParams struct {
	// Id identifies output Storage in this Project. Omit to use the Project default.
	ID param.Opt[string] `json:"id,omitzero"`
	// Path specifies the output filename template relative to the selected Storage
	// base_prefix. Required for customer storage; generated for Chunkify temporary
	// storage. The filename must contain one %05d placeholder, replaced by the image
	// or sprite sheet number starting at 1, and end in .jpg for jpeg or .webp for
	// webp. The WebVTT file uses the filename without the placeholder and its adjacent
	// separator, such as thumb.vtt for thumb-%05d.jpg. The combined path may contain
	// at most 1024 bytes.
	Path param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r ThumbnailCreateStorageParams) MarshalJSON() (data []byte, err error) {
	type shadow ThumbnailCreateStorageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThumbnailCreateStorageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThumbnailNewParams struct {
	// SourceId identifies a video Source in this Project.
	SourceID string `json:"source_id" api:"required"`
	// Count specifies 1 to 100 thumbnails at the middle of equal sections of the
	// selected range.
	Count param.Opt[int64] `json:"count,omitzero"`
	// Size specifies the maximum image or tile edge in pixels, from 64 to 1920.
	// Defaults to 320. Odd sizes are supported. Preserves display aspect ratio without
	// cropping or upscaling.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Format selects JPEG or WebP. Defaults to JPEG.
	//
	// Any of "jpeg", "webp".
	Format ThumbnailNewParamsFormat `json:"format,omitzero"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output ThumbnailNewParamsOutput `json:"output,omitzero"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range ThumbnailNewParamsRange `json:"range,omitzero"`
	// Storage overrides the Project output destination and names the output files.
	// Source storage remains independent.
	Storage ThumbnailCreateStorageParams `json:"storage,omitzero"`
	// Timestamps specifies 1 to 100 Source-relative target times in seconds. Every
	// time must be within the selected range. Repeated targets produce separate
	// outputs.
	Timestamps []float64 `json:"timestamps,omitzero"`
	paramObj
}

func (r ThumbnailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ThumbnailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThumbnailNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format selects JPEG or WebP. Defaults to JPEG.
type ThumbnailNewParamsFormat string

const (
	ThumbnailNewParamsFormatJpeg ThumbnailNewParamsFormat = "jpeg"
	ThumbnailNewParamsFormatWebp ThumbnailNewParamsFormat = "webp"
)

// Output selects images or sprites. Defaults to images.
type ThumbnailNewParamsOutput string

const (
	ThumbnailNewParamsOutputImages  ThumbnailNewParamsOutput = "images"
	ThumbnailNewParamsOutputSprites ThumbnailNewParamsOutput = "sprites"
)

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type ThumbnailNewParamsRange struct {
	// End specifies the exclusive upper bound in seconds. Defaults to the Source
	// duration.
	End param.Opt[float64] `json:"end,omitzero"`
	// Start specifies the inclusive lower bound in seconds. Defaults to zero.
	Start param.Opt[float64] `json:"start,omitzero"`
	paramObj
}

func (r ThumbnailNewParamsRange) MarshalJSON() (data []byte, err error) {
	type shadow ThumbnailNewParamsRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThumbnailNewParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A thumbnail resource response.
type ThumbnailNewResponseEnvelope struct {
	// Data contains the thumbnail resource.
	Data Thumbnail `json:"data" api:"required"`
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
func (r ThumbnailNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A thumbnail resource response.
type ThumbnailGetResponseEnvelope struct {
	// Data contains the thumbnail resource.
	Data Thumbnail `json:"data" api:"required"`
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
func (r ThumbnailGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ThumbnailGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThumbnailListParams struct {
	// Id filters by an exact Thumbnail ID. Omit to include all matching resources.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Maximum number of Thumbnails to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of Thumbnails to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter Thumbnails linked to this Source. Omit to include Thumbnails from all
	// Sources in the Project.
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// Filter by output Storage ID.
	StorageID param.Opt[string]          `query:"storage_id,omitzero" json:"-"`
	Created   ThumbnailListParamsCreated `query:"created,omitzero" json:"-"`
	// Filter by Thumbnail status.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status ThumbnailListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ThumbnailListParams]'s query parameters as `url.Values`.
func (r ThumbnailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThumbnailListParamsCreated struct {
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

// URLQuery serializes [ThumbnailListParamsCreated]'s query parameters as
// `url.Values`.
func (r ThumbnailListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by Thumbnail status.
type ThumbnailListParamsStatus string

const (
	ThumbnailListParamsStatusPreparing ThumbnailListParamsStatus = "preparing"
	ThumbnailListParamsStatusReady     ThumbnailListParamsStatus = "ready"
	ThumbnailListParamsStatusFailed    ThumbnailListParamsStatus = "failed"
	ThumbnailListParamsStatusExpired   ThumbnailListParamsStatus = "expired"
)

type ThumbnailDeleteParams struct {
	// DeleteFiles specifies whether to permanently delete this Thumbnail's owned
	// images, sprite sheets, VTT and partial outputs at their saved storage prefix.
	// Defaults to false. The Source and unrelated files are preserved. If deletion
	// fails, records remain for retry, but some files may already be deleted.
	// Temporary outputs still expire when this option is false.
	DeleteFiles param.Opt[bool] `query:"delete_files,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ThumbnailDeleteParams]'s query parameters as `url.Values`.
func (r ThumbnailDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
