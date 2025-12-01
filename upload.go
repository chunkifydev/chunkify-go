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
	opts = slices.Concat(r.Options, opts)
	path := "api/uploads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve details of a specific upload by its ID, including metadata, status, and
// associated source.
func (r *UploadService) Get(ctx context.Context, uploadID string, opts ...option.RequestOption) (res *Upload, err error) {
	var env UploadGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return
	}
	path := fmt.Sprintf("api/uploads/%s", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve a list of all uploads with optional filtering and pagination.
func (r *UploadService) List(ctx context.Context, query UploadListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Upload], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return
	}
	path := fmt.Sprintf("api/uploads/%s", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Upload struct {
	// Unique identifier of the upload
	ID string `json:"id,required"`
	// Timestamp when the upload was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp when the upload will expire
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Current status of the upload
	//
	// Any of "waiting", "completed", "failed", "expired".
	Status UploadStatus `json:"status,required"`
	// Timestamp when the upload was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Pre-signed URL where the file should be uploaded to
	UploadURL string `json:"upload_url,required"`
	// Error message of the upload
	Error shared.ChunkifyError `json:"error"`
	// Additional metadata for the upload
	Metadata map[string]string `json:"metadata"`
	// SourceId is the id of the source that was created from the upload
	SourceID string `json:"source_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UploadURL   respjson.Field
		Error       respjson.Field
		Metadata    respjson.Field
		SourceID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// The upload URL will be valid for the given timeout in seconds
	ValidityTimeout param.Opt[int64] `json:"validity_timeout,omitzero"`
	// Metadata allows for additional information to be attached to the upload, with a
	// maximum size of 1024 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r UploadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UploadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UploadNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadNewResponseEnvelope struct {
	// Data contains the response object
	Data Upload `json:"data,required"`
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
func (r UploadNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *UploadNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadGetResponseEnvelope struct {
	// Data contains the response object
	Data Upload `json:"data,required"`
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
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Filter by creation date less than or equal (UNIX epoch time)
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
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
