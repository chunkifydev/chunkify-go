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

// Create a new upload with the specified name.
func (r *UploadService) New(ctx context.Context, body UploadNewParams, opts ...option.RequestOption) (res *Upload, err error) {
	var env UploadNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "api/uploads"
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
	path := fmt.Sprintf("api/uploads/%s", uploadID)
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
	path := "api/uploads"
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
	path := fmt.Sprintf("api/uploads/%s", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// After a successful PUT, POST the returned completion_url before expires_at. The
// token authorizes only this Upload; no API key, cookies, or request body is
// required. Verifies the stored object and commits one Source relationship. Valid
// retries return 204 without duplicate side effects. Retry network errors, 429,
// and 5xx responses with bounded backoff; never repeat the PUT just to retry
// completion.
func (r *UploadService) Complete(ctx context.Context, token string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return err
	}
	path := fmt.Sprintf("api/uploads/completion/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type Upload struct {
	// Unique identifier of the upload
	ID string `json:"id" api:"required"`
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
	// Short-lived completion capability, returned only on creation. POST after a
	// successful PUT before expires_at. Requires no API key. Repeated valid calls are
	// idempotent.
	CompletionURL string `json:"completion_url"`
	// Error message of the upload
	Error shared.ChunkifyError `json:"error"`
	// Additional metadata for the upload
	Metadata map[string]string `json:"metadata"`
	// SourceId is the id of the source that was created from the upload
	SourceID string `json:"source_id"`
	// Resolved Storage selected when the Upload was created. Absent for historical
	// uploads.
	StorageID string `json:"storage_id"`
	// Presigned PUT URL, returned only when creating an Upload session. Call
	// completion_url after the PUT succeeds.
	UploadURL string `json:"upload_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		ExpiresAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		CompletionURL respjson.Field
		Error         respjson.Field
		Metadata      respjson.Field
		SourceID      respjson.Field
		StorageID     respjson.Field
		UploadURL     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
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

type UploadNewParams struct {
	// Both the file PUT and completion POST must finish within this timeout in seconds
	ValidityTimeout param.Opt[int64] `json:"validity_timeout,omitzero"`
	// Metadata allows for additional information to be attached to the upload, with a
	// maximum size of 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Optional Storage override. Omit id to use the Project default.
	// Customer-connected Storage requires path; Chunkify Storage generates its own
	// path.
	Storage UploadNewParamsStorage `json:"storage,omitzero"`
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
