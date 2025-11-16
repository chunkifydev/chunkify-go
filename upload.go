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
func (r *UploadService) New(ctx context.Context, body UploadNewParams, opts ...option.RequestOption) (res *UploadNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/uploads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific upload by its ID, including metadata, status, and
// associated source.
func (r *UploadService) Get(ctx context.Context, uploadID string, opts ...option.RequestOption) (res *UploadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return
	}
	path := fmt.Sprintf("api/uploads/%s", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	ID string `json:"id"`
	// Timestamp when the upload was created
	CreatedAt string `json:"created_at"`
	// Error message of the upload
	Error UploadError `json:"error"`
	// Timestamp when the upload will expire
	ExpiresAt string `json:"expires_at"`
	// Additional metadata for the upload
	Metadata map[string]string `json:"metadata"`
	// SourceId is the id of the source that was created from the upload
	SourceID string `json:"source_id"`
	// Current status of the upload (waiting, completed, failed, expired)
	Status string `json:"status"`
	// Timestamp when the upload was updated
	UpdatedAt string `json:"updated_at"`
	// Pre-signed URL where the file should be uploaded to
	UploadURL string `json:"upload_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Error       respjson.Field
		ExpiresAt   respjson.Field
		Metadata    respjson.Field
		SourceID    respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Upload) RawJSON() string { return r.JSON.raw }
func (r *Upload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error message of the upload
type UploadError struct {
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
func (r UploadError) RawJSON() string { return r.JSON.raw }
func (r *UploadError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadNewResponse struct {
	Data Upload `json:"data"`
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
func (r UploadNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UploadNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadGetResponse struct {
	Data Upload `json:"data"`
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
func (r UploadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UploadGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type UploadListParams struct {
	// Filter by upload ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Pagination limit (max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by metadata (format: key:value,key:value)
	Metadata param.Opt[string] `query:"metadata,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by source ID
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// Filter by status (pending, completed, error)
	Status  param.Opt[string]       `query:"status,omitzero" json:"-"`
	Created UploadListParamsCreated `query:"created,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UploadListParams]'s query parameters as `url.Values`.
func (r UploadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UploadListParamsCreated struct {
	// Filter by creation date greater than or equal (RFC3339)
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Filter by creation date less than or equal (RFC3339)
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	// Sort by creation date (asc/desc)
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UploadListParamsCreated]'s query parameters as
// `url.Values`.
func (r UploadListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
