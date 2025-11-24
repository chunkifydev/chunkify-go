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

// SourceService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSourceService] method instead.
type SourceService struct {
	Options []option.RequestOption
}

// NewSourceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSourceService(opts ...option.RequestOption) (r SourceService) {
	r = SourceService{}
	r.Options = opts
	return
}

// Create a new source from a media URL. The source will be analyzed to extract
// metadata and generate a thumbnail. The source will be automatically deleted
// after the data retention period.
func (r *SourceService) New(ctx context.Context, body SourceNewParams, opts ...option.RequestOption) (res *Source, err error) {
	var env SourceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "api/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve details of a specific source by its ID, including metadata, media
// properties, and associated jobs.
func (r *SourceService) Get(ctx context.Context, sourceID string, opts ...option.RequestOption) (res *Source, err error) {
	var env SourceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return
	}
	path := fmt.Sprintf("api/sources/%s", sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve a list of all sources with optional filtering and pagination. Supports
// filtering by various media properties like duration, dimensions, codecs, etc.
func (r *SourceService) List(ctx context.Context, query SourceListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Source], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/sources"
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

// Retrieve a list of all sources with optional filtering and pagination. Supports
// filtering by various media properties like duration, dimensions, codecs, etc.
func (r *SourceService) ListAutoPaging(ctx context.Context, query SourceListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Source] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete a source. It will fail if there are processing jobs using this source.
func (r *SourceService) Delete(ctx context.Context, sourceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return
	}
	path := fmt.Sprintf("api/sources/%s", sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Source struct {
	// Unique identifier of the source
	ID string `json:"id,required"`
	// Audio bitrate in bits per second
	AudioBitrate int64 `json:"audio_bitrate,required"`
	// Audio codec used
	AudioCodec string `json:"audio_codec,required"`
	// Timestamp when the source was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Device used to record the video
	Device string `json:"device,required"`
	// Duration of the video in seconds
	Duration int64 `json:"duration,required"`
	// Height of the video in pixels
	Height int64 `json:"height,required"`
	// Additional metadata for the source
	Metadata map[string]string `json:"metadata,required"`
	// Size of the source file in bytes
	Size int64 `json:"size,required"`
	// URL where the source video can be accessed
	URL string `json:"url,required"`
	// Video bitrate in bits per second
	VideoBitrate int64 `json:"video_bitrate,required"`
	// Video codec used
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
		Device         respjson.Field
		Duration       respjson.Field
		Height         respjson.Field
		Metadata       respjson.Field
		Size           respjson.Field
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
func (r Source) RawJSON() string { return r.JSON.raw }
func (r *Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceNewParams struct {
	// Url is the URL of the source, which must be a valid HTTP URL.
	URL string `json:"url,required"`
	// Metadata allows for additional information to be attached to the source, with a
	// maximum size of 1024 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r SourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceNewResponseEnvelope struct {
	// Data contains the response object
	Data Source `json:"data,required"`
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
func (r SourceNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *SourceNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceGetResponseEnvelope struct {
	// Data contains the response object
	Data Source `json:"data,required"`
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
func (r SourceGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *SourceGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListParams struct {
	// Filter by source ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Filter by audio codec
	AudioCodec param.Opt[string] `query:"audio_codec,omitzero" json:"-"`
	// Pagination limit (max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by video codec
	VideoCodec param.Opt[string]       `query:"video_codec,omitzero" json:"-"`
	Created    SourceListParamsCreated `query:"created,omitzero" json:"-"`
	// Filter by device (apple/android)
	//
	// Any of "apple", "android", "unknown".
	Device   SourceListParamsDevice   `query:"device,omitzero" json:"-"`
	Duration SourceListParamsDuration `query:"duration,omitzero" json:"-"`
	Height   SourceListParamsHeight   `query:"height,omitzero" json:"-"`
	// Filter by metadata (format: key:value,key:value)
	Metadata [][]string            `query:"metadata,omitzero" json:"-"`
	Size     SourceListParamsSize  `query:"size,omitzero" json:"-"`
	Width    SourceListParamsWidth `query:"width,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SourceListParams]'s query parameters as `url.Values`.
func (r SourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SourceListParamsCreated struct {
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

// URLQuery serializes [SourceListParamsCreated]'s query parameters as
// `url.Values`.
func (r SourceListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by device (apple/android)
type SourceListParamsDevice string

const (
	SourceListParamsDeviceApple   SourceListParamsDevice = "apple"
	SourceListParamsDeviceAndroid SourceListParamsDevice = "android"
	SourceListParamsDeviceUnknown SourceListParamsDevice = "unknown"
)

type SourceListParamsDuration struct {
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

// URLQuery serializes [SourceListParamsDuration]'s query parameters as
// `url.Values`.
func (r SourceListParamsDuration) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SourceListParamsHeight struct {
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

// URLQuery serializes [SourceListParamsHeight]'s query parameters as `url.Values`.
func (r SourceListParamsHeight) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SourceListParamsSize struct {
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

// URLQuery serializes [SourceListParamsSize]'s query parameters as `url.Values`.
func (r SourceListParamsSize) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SourceListParamsWidth struct {
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

// URLQuery serializes [SourceListParamsWidth]'s query parameters as `url.Values`.
func (r SourceListParamsWidth) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
