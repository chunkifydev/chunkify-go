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
	"github.com/chunkifydev/chunkify-go/shared"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// UploadService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadService] method instead.
type UploadService struct {
	Options []option.RequestOption
}

// NewUploadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUploadService(opts ...option.RequestOption) (r UploadService) {
	r = UploadService{}
	r.Options = opts
	return
}

// Create an Upload session for a direct transfer to storage. Multipart creation
// returns every signed part URL. Retry parts with their original URLs and complete
// before expiry. Customer storage requires long-lived credentials and an
// incomplete-multipart lifecycle rule. Start a new Upload after abandoning a
// session.
func (r *UploadService) New(ctx context.Context, body UploadNewParams, opts ...option.RequestOption) (res *Upload, err error) {
	var env UploadNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "uploads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve details of a specific upload by its ID, including metadata, status, and
// associated source.
func (r *UploadService) Get(ctx context.Context, uploadID string, opts ...option.RequestOption) (res *Upload, err error) {
	var env UploadGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return nil, err
	}
	path := fmt.Sprintf("uploads/%s", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve a list of all uploads with optional filtering and pagination.
func (r *UploadService) List(ctx context.Context, query UploadListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Upload], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "uploads"
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

// Retrieve a list of all uploads with optional filtering and pagination.
func (r *UploadService) ListAutoPaging(ctx context.Context, query UploadListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Upload] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete an upload.
func (r *UploadService) Delete(ctx context.Context, uploadID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return err
	}
	path := fmt.Sprintf("uploads/%s", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// POST the returned completion_url before expires_at. Single-PUT Uploads send no
// body; multipart Uploads send every part number and storage ETag. Chunkify
// verifies the assembled object and atomically creates the Source and optional
// Asset. An Asset destination conflict returns 409 and fails the Upload without
// committing Source, Asset, or usage changes. Valid retries return 204 without
// duplicate side effects. Retry network errors, 429, and 5xx responses with
// bounded backoff.
func (r *UploadService) Complete(ctx context.Context, token string, body UploadCompleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return err
	}
	path := fmt.Sprintf("uploads/completion/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type Upload struct {
	// Unique identifier of the upload
	ID string `json:"id" api:"required"`
	// Whether this Upload creates an Asset on completion.
	CreateAsset bool `json:"create_asset" api:"required"`
	// Timestamp when the upload was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Timestamp when the upload will expire
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Current status of the upload
	//
	// Any of "waiting", "completed", "failed", "expired".
	Status UploadStatus `json:"status" api:"required"`
	// Timestamp when the upload was updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Resolved transfer mode for this Upload.
	//
	// Any of "single", "multipart".
	UploadMode UploadUploadMode `json:"upload_mode" api:"required"`
	// Selected Workflow ID, or null.
	WorkflowID string `json:"workflow_id" api:"required"`
	// Pinned Workflow revision, or null.
	WorkflowRevision int64 `json:"workflow_revision" api:"required"`
	// Automatically created Asset. May be absent before completion or after record
	// deletion.
	AssetID string `json:"asset_id"`
	// Completion capability, returned only on creation. POST after a successful single
	// PUT or with all multipart ETags before expires_at. Requires no API key. Repeated
	// valid calls are idempotent.
	CompletionURL string `json:"completion_url"`
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
	CreatedVia UploadCreatedVia `json:"created_via"`
	// Error message of the upload
	Error shared.ChunkifyError `json:"error"`
	// Declared file size in bytes, when supplied at creation.
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Additional metadata for the upload
	Metadata map[string]string `json:"metadata"`
	// Multipart transfer details, returned only on creation for multipart Uploads.
	Multipart UploadMultipart `json:"multipart"`
	// SourceId is the id of the source that was created from the upload
	SourceID string `json:"source_id"`
	// Resolved Storage selected when the Upload was created. Absent for historical
	// uploads.
	StorageID string `json:"storage_id"`
	// Presigned PUT URL, returned only when creating a single-PUT Upload session. Call
	// completion_url after the PUT succeeds.
	UploadURL string `json:"upload_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreateAsset      respjson.Field
		CreatedAt        respjson.Field
		ExpiresAt        respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		UploadMode       respjson.Field
		WorkflowID       respjson.Field
		WorkflowRevision respjson.Field
		AssetID          respjson.Field
		CompletionURL    respjson.Field
		CreatedByTokenID respjson.Field
		CreatedByUserID  respjson.Field
		CreatedVia       respjson.Field
		Error            respjson.Field
		FileSizeBytes    respjson.Field
		Metadata         respjson.Field
		Multipart        respjson.Field
		SourceID         respjson.Field
		StorageID        respjson.Field
		UploadURL        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Upload) RawJSON() string { return r.JSON.raw }
func (r *Upload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the upload
type UploadStatus string

const (
	UploadStatusWaiting   UploadStatus = "waiting"
	UploadStatusCompleted UploadStatus = "completed"
	UploadStatusFailed    UploadStatus = "failed"
	UploadStatusExpired   UploadStatus = "expired"
)

// Resolved transfer mode for this Upload.
type UploadUploadMode string

const (
	UploadUploadModeSingle    UploadUploadMode = "single"
	UploadUploadModeMultipart UploadUploadMode = "multipart"
)

// CreatedVia specifies how the resource was created: api for API requests,
// dashboard for the Chunkify dashboard, mcp for MCP clients, and workflow for
// Workflow actions. Omitted when the creator is unknown.
type UploadCreatedVia string

const (
	UploadCreatedViaAPI       UploadCreatedVia = "api"
	UploadCreatedViaDashboard UploadCreatedVia = "dashboard"
	UploadCreatedViaMcp       UploadCreatedVia = "mcp"
	UploadCreatedViaWorkflow  UploadCreatedVia = "workflow"
)

// Multipart transfer details, returned only on creation for multipart Uploads.
type UploadMultipart struct {
	// ExpiresAt specifies when the signed part URLs expire. URLs cannot be refreshed;
	// create a new Upload after expiry.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// PartSizeBytes specifies the size of each part except the final part, in bytes.
	PartSizeBytes int64 `json:"part_size_bytes" api:"required"`
	// Parts contains every signed upload URL in ascending part_number order. Upload
	// file slices directly to these URLs.
	Parts []UploadMultipartPart `json:"parts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt     respjson.Field
		PartSizeBytes respjson.Field
		Parts         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadMultipart) RawJSON() string { return r.JSON.raw }
func (r *UploadMultipart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadMultipartPart struct {
	// PartNumber specifies the one-based part number for this URL.
	PartNumber int64 `json:"part_number" api:"required"`
	// UploadUrl specifies the signed PUT URL for this part. Retry the same file slice
	// with this URL before multipart.expires_at.
	UploadURL string `json:"upload_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PartNumber  respjson.Field
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadMultipartPart) RawJSON() string { return r.JSON.raw }
func (r *UploadMultipartPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadNewParams struct {
	// WorkflowId specifies the Workflow to select. Omit to use the Project default.
	// Set null to disable automatic actions. The current revision is pinned when
	// creation is accepted.
	WorkflowID param.Opt[string] `json:"workflow_id,omitzero"`
	// CreateAsset specifies whether to create an Asset after Upload completion.
	// Defaults to false. Snapshots the Project output settings, including the path
	// template. Referenced metadata must be non-empty in this request. Asset
	// processing is billed separately.
	CreateAsset param.Opt[bool] `json:"create_asset,omitzero"`
	// FileSizeBytes specifies the exact file size in bytes. Required when UploadMode
	// is omitted, auto, or multipart. Optional for explicit single mode. Must fit the
	// team's Source size limit. Single PUT supports up to 5 GiB, subject to the
	// provider's limit; multipart supports up to 5 TiB.
	FileSizeBytes param.Opt[int64] `json:"file_size_bytes,omitzero"`
	// ValidityTimeout specifies the session and signed URL lifetime in seconds. Single
	// PUT defaults to 7200 and allows 300–36000. Multipart defaults to twice the
	// transfer time at 5 Mbps plus 1800 seconds, bounded to 7200–172800; an explicit
	// value allows 300–172800. Complete before expiry. Storage credentials and
	// policies can shorten URL validity. Expired URLs cannot be refreshed.
	ValidityTimeout param.Opt[int64] `json:"validity_timeout,omitzero"`
	// Metadata allows for additional information to be attached to the upload, with a
	// maximum size of 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Optional Storage override. Omit id to use the Project default.
	// Customer-connected Storage requires path; Chunkify Storage generates its own
	// path.
	Storage UploadNewParamsStorage `json:"storage,omitzero"`
	// UploadMode specifies the transfer method. Defaults to auto. Auto chooses
	// multipart above 100 MiB and single PUT at or below 100 MiB for every Storage
	// provider. Multipart requires a provider that supports the required multipart
	// operations and signed part URLs. Provider failures return an error without
	// falling back to single PUT.
	//
	// Any of "single", "multipart", "auto".
	UploadMode UploadNewParamsUploadMode `json:"upload_mode,omitzero"`
	paramObj
}

func (r UploadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UploadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UploadNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional Storage override. Omit id to use the Project default.
// Customer-connected Storage requires path; Chunkify Storage generates its own
// path.
type UploadNewParamsStorage struct {
	// Storage belonging to this Project. Omit to use the Project default.
	ID param.Opt[string] `json:"id,omitzero"`
	// Exact object key including filename, required for customer Storage and forbidden
	// for Chunkify Storage. The output base_prefix is not added. Existing keys may be
	// overwritten.
	Path param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r UploadNewParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow UploadNewParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UploadNewParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UploadMode specifies the transfer method. Defaults to auto. Auto chooses
// multipart above 100 MiB and single PUT at or below 100 MiB for every Storage
// provider. Multipart requires a provider that supports the required multipart
// operations and signed part URLs. Provider failures return an error without
// falling back to single PUT.
type UploadNewParamsUploadMode string

const (
	UploadNewParamsUploadModeSingle    UploadNewParamsUploadMode = "single"
	UploadNewParamsUploadModeMultipart UploadNewParamsUploadMode = "multipart"
	UploadNewParamsUploadModeAuto      UploadNewParamsUploadMode = "auto"
)

type UploadNewResponseEnvelope struct {
	// Data contains the response object
	Data Upload `json:"data" api:"required"`
	// Status indicates the response status "success"
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
func (r UploadNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *UploadNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadGetResponseEnvelope struct {
	// Data contains the response object
	Data Upload `json:"data" api:"required"`
	// Status indicates the response status "success"
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
func (r UploadGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *UploadGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadListParams struct {
	// Filter by upload ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Pagination limit (max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by source ID
	SourceID param.Opt[string]       `query:"source_id,omitzero" json:"-"`
	Created  UploadListParamsCreated `query:"created,omitzero" json:"-"`
	// Filter by metadata
	Metadata [][]string `query:"metadata,omitzero" json:"-"`
	// Filter by status (pending, completed, error)
	//
	// Any of "waiting", "completed", "failed", "expired".
	Status UploadListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UploadListParams]'s query parameters as `url.Values`.
func (r UploadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UploadListParamsCreated struct {
	// Filter by creation date greater than or equal (UNIX epoch time)
	Gte param.Opt[int64] `query:"gte,omitzero" format:"epoch" json:"-"`
	// Filter by creation date less than or equal (UNIX epoch time)
	Lte param.Opt[int64] `query:"lte,omitzero" format:"epoch" json:"-"`
	// Sort by creation date (asc/desc)
	//
	// Any of "asc", "desc".
	Sort string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UploadListParamsCreated]'s query parameters as
// `url.Values`.
func (r UploadListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status (pending, completed, error)
type UploadListParamsStatus string

const (
	UploadListParamsStatusWaiting   UploadListParamsStatus = "waiting"
	UploadListParamsStatusCompleted UploadListParamsStatus = "completed"
	UploadListParamsStatusFailed    UploadListParamsStatus = "failed"
	UploadListParamsStatusExpired   UploadListParamsStatus = "expired"
)

type UploadCompleteParams struct {
	// All uploaded parts in ascending part_number order, with the ETags returned by
	// storage.
	Parts []UploadCompleteParamsPart `json:"parts,omitzero"`
	paramObj
}

func (r UploadCompleteParams) MarshalJSON() (data []byte, err error) {
	type shadow UploadCompleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UploadCompleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Etag, PartNumber are required.
type UploadCompleteParamsPart struct {
	// ETag returned by storage for this part, including its quotation marks when
	// present.
	Etag string `json:"etag" api:"required"`
	// One-based part number.
	PartNumber int64 `json:"part_number" api:"required"`
	paramObj
}

func (r UploadCompleteParamsPart) MarshalJSON() (data []byte, err error) {
	type shadow UploadCompleteParamsPart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UploadCompleteParamsPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
