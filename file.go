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

	"github.com/stainless-sdks/chunkify-go/internal/apijson"
	"github.com/stainless-sdks/chunkify-go/internal/apiquery"
	"github.com/stainless-sdks/chunkify-go/internal/requestconfig"
	"github.com/stainless-sdks/chunkify-go/option"
	"github.com/stainless-sdks/chunkify-go/packages/pagination"
	"github.com/stainless-sdks/chunkify-go/packages/param"
	"github.com/stainless-sdks/chunkify-go/packages/respjson"
	"github.com/stainless-sdks/chunkify-go/shared/constant"
)

// FileService contains methods and other services that help with interacting with
// the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// Retrieve details of a specific file by its ID, including metadata, media
// properties, and associated jobs.
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *APIFile, err error) {
	var env FileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("api/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve a list of files with optional filtering and pagination
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[APIFile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/files"
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

// Retrieve a list of files with optional filtering and pagination
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[APIFile] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete a file. It will fail if there are processing jobs using this file.
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("api/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type APIFile struct {
	// Unique identifier of the file
	ID string `json:"id,required"`
	// Audio bitrate in bits per second
	AudioBitrate int64 `json:"audio_bitrate,required"`
	// Audio codec used (e.g. aac, mp3)
	AudioCodec string `json:"audio_codec,required"`
	// Timestamp when the file was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Duration of the video in seconds
	Duration int64 `json:"duration,required"`
	// Height of the video in pixels
	Height int64 `json:"height,required"`
	// ID of the job that created this file
	JobID string `json:"job_id,required"`
	// MIME type of the file
	MimeType string `json:"mime_type,required"`
	// Path to the file in storage
	Path string `json:"path,required"`
	// Size of the file in bytes
	Size int64 `json:"size,required"`
	// StorageId identifier where the file is stored
	StorageID string `json:"storage_id,required"`
	// Pre-signed URL to directly access the file (only included when available)
	URL string `json:"url,required"`
	// Video bitrate in bits per second
	VideoBitrate int64 `json:"video_bitrate,required"`
	// Video codec used (e.g. h264, h265)
	VideoCodec string `json:"video_codec,required"`
	// Video framerate in frames per second
	VideoFramerate float64 `json:"video_framerate,required"`
	// Width of the video in pixels
	Width int64 `json:"width,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AudioBitrate   respjson.Field
		AudioCodec     respjson.Field
		CreatedAt      respjson.Field
		Duration       respjson.Field
		Height         respjson.Field
		JobID          respjson.Field
		MimeType       respjson.Field
		Path           respjson.Field
		Size           respjson.Field
		StorageID      respjson.Field
		URL            respjson.Field
		VideoBitrate   respjson.Field
		VideoCodec     respjson.Field
		VideoFramerate respjson.Field
		Width          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIFile) RawJSON() string { return r.JSON.raw }
func (r *APIFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type FileGetResponseEnvelope struct {
	// Data contains the response object
	Data APIFile `json:"data,required"`
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
func (r FileGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *FileGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListParams struct {
	// Filter by file ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Filter by audio codec
	AudioCodec param.Opt[string] `query:"audio_codec,omitzero" json:"-"`
	// Filter by job ID
	JobID param.Opt[string] `query:"job_id,omitzero" json:"-"`
	// Pagination limit (max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by mime type
	MimeType param.Opt[string] `query:"mime_type,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by storage ID
	StorageID param.Opt[string] `query:"storage_id,omitzero" json:"-"`
	// Filter by video codec
	VideoCodec param.Opt[string]      `query:"video_codec,omitzero" json:"-"`
	Created    FileListParamsCreated  `query:"created,omitzero" json:"-"`
	Duration   FileListParamsDuration `query:"duration,omitzero" json:"-"`
	Height     FileListParamsHeight   `query:"height,omitzero" json:"-"`
	Path       FileListParamsPath     `query:"path,omitzero" json:"-"`
	Size       FileListParamsSize     `query:"size,omitzero" json:"-"`
	Width      FileListParamsWidth    `query:"width,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParamsCreated struct {
	// Filter by creation date greater than or equal (RFC3339)
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Filter by creation date less than or equal (RFC3339)
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	// Sort by creation date (asc/desc)
	//
	// Any of "asc", "desc".
	Sort string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParamsCreated]'s query parameters as `url.Values`.
func (r FileListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParamsDuration struct {
	// Filter by exact duration
	Eq param.Opt[float64] `query:"eq,omitzero" json:"-"`
	// Filter by duration greater than
	Gt param.Opt[float64] `query:"gt,omitzero" json:"-"`
	// Filter by duration greater than or equal
	Gte param.Opt[float64] `query:"gte,omitzero" json:"-"`
	// Filter by duration less than
	Lt param.Opt[float64] `query:"lt,omitzero" json:"-"`
	// Filter by duration less than or equal
	Lte param.Opt[float64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParamsDuration]'s query parameters as `url.Values`.
func (r FileListParamsDuration) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParamsHeight struct {
	// Filter by exact height
	Eq param.Opt[int64] `query:"eq,omitzero" json:"-"`
	// Filter by height greater than
	Gt param.Opt[int64] `query:"gt,omitzero" json:"-"`
	// Filter by height greater than or equal
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Filter by height less than
	Lt param.Opt[int64] `query:"lt,omitzero" json:"-"`
	// Filter by height less than or equal
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParamsHeight]'s query parameters as `url.Values`.
func (r FileListParamsHeight) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParamsPath struct {
	// Filter by path
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Filter by path (case insensitive)
	Ilike param.Opt[string] `query:"ilike,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParamsPath]'s query parameters as `url.Values`.
func (r FileListParamsPath) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParamsSize struct {
	// Filter by exact file size
	Eq param.Opt[int64] `query:"eq,omitzero" json:"-"`
	// Filter by file size greater than
	Gt param.Opt[int64] `query:"gt,omitzero" json:"-"`
	// Filter by file size greater than or equal
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Filter by file size less than
	Lt param.Opt[int64] `query:"lt,omitzero" json:"-"`
	// Filter by file size less than or equal
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParamsSize]'s query parameters as `url.Values`.
func (r FileListParamsSize) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParamsWidth struct {
	// Filter by exact width
	Eq param.Opt[int64] `query:"eq,omitzero" json:"-"`
	// Filter by width greater than
	Gt param.Opt[int64] `query:"gt,omitzero" json:"-"`
	// Filter by width greater than or equal
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Filter by width less than
	Lt param.Opt[int64] `query:"lt,omitzero" json:"-"`
	// Filter by width less than or equal
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParamsWidth]'s query parameters as `url.Values`.
func (r FileListParamsWidth) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
