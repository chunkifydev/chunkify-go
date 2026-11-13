// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"encoding/json"
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

// Automatic video-on-demand Assets and independent H.264 MP4 static renditions.
//
// AssetOutputService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetOutputService] method instead.
type AssetOutputService struct {
	Options []option.RequestOption
}

// NewAssetOutputService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssetOutputService(opts ...option.RequestOption) (r AssetOutputService) {
	r = AssetOutputService{}
	r.Options = opts
	return
}

// Create one caption, static rendition or additional thumbnail output. The
// accepted output processes asynchronously and is returned in the same envelope as
// output reads.
func (r *AssetOutputService) New(ctx context.Context, assetID string, body AssetOutputNewParams, opts ...option.RequestOption) (res *AssetOutputUnion, err error) {
	var env AssetOutputNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s/outputs", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve one output belonging to the specified Asset and authenticated Project.
// Delivery URLs follow the output resource retention and expiry rules.
func (r *AssetOutputService) Get(ctx context.Context, outputID string, query AssetOutputGetParams, opts ...option.RequestOption) (res *AssetOutputUnion, err error) {
	var env AssetOutputGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.AssetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	if outputID == "" {
		err = errors.New("missing required outputId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s/outputs/%s", query.AssetID, outputID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List outputs requested directly or by a Workflow. Results are ordered by
// creation date, newest first, with descending output ID as a tie-breaker.
// Pagination and total apply after type and status filters. Core playback, poster
// and default timeline previews remain on the Asset.
func (r *AssetOutputService) List(ctx context.Context, assetID string, query AssetOutputListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[AssetOutputUnion], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s/outputs", assetID)
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

// List outputs requested directly or by a Workflow. Results are ordered by
// creation date, newest first, with descending output ID as a tie-breaker.
// Pagination and total apply after type and status filters. Core playback, poster
// and default timeline previews remain on the Asset.
func (r *AssetOutputService) ListAutoPaging(ctx context.Context, assetID string, query AssetOutputListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[AssetOutputUnion] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, assetID, query, opts...))
}

// Delete one output belonging to this Asset and Project. Its type is resolved from
// the output ID. Type-specific processing and dependency restrictions still apply.
func (r *AssetOutputService) Delete(ctx context.Context, outputID string, params AssetOutputDeleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AssetID == "" {
		err = errors.New("missing required assetId parameter")
		return err
	}
	if outputID == "" {
		err = errors.New("missing required outputId parameter")
		return err
	}
	path := fmt.Sprintf("assets/%s/outputs/%s", params.AssetID, outputID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// AssetOutputUnion contains all possible properties and values from
// [AssetCaptionOutput], [AssetStaticRenditionOutput], [AssetThumbnailOutput],
// [AssetTranscriptOutput].
//
// Use the [AssetOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssetOutputUnion struct {
	ID      string `json:"id"`
	AssetID string `json:"asset_id"`
	// This field is a union of [AssetCaption], [AssetStaticRendition], [Thumbnail],
	// [AssetTranscript]
	Details AssetOutputUnionDetails `json:"details"`
	// This field is from variant [AssetCaptionOutput].
	Status AssetOutputStatus `json:"status"`
	// Any of "caption", "static_rendition", "thumbnail", "transcript".
	Type string `json:"type"`
	JSON struct {
		ID      respjson.Field
		AssetID respjson.Field
		Details respjson.Field
		Status  respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

// anyAssetOutput is implemented by each variant of [AssetOutputUnion] to add type
// safety for the return type of [AssetOutputUnion.AsAny]
type anyAssetOutput interface {
	implAssetOutputUnion()
}

func (AssetCaptionOutput) implAssetOutputUnion()         {}
func (AssetStaticRenditionOutput) implAssetOutputUnion() {}
func (AssetThumbnailOutput) implAssetOutputUnion()       {}
func (AssetTranscriptOutput) implAssetOutputUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AssetOutputUnion.AsAny().(type) {
//	case chunkify.AssetCaptionOutput:
//	case chunkify.AssetStaticRenditionOutput:
//	case chunkify.AssetThumbnailOutput:
//	case chunkify.AssetTranscriptOutput:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AssetOutputUnion) AsAny() anyAssetOutput {
	switch u.Type {
	case "caption":
		return u.AsCaption()
	case "static_rendition":
		return u.AsStaticRendition()
	case "thumbnail":
		return u.AsThumbnail()
	case "transcript":
		return u.AsTranscript()
	}
	return nil
}

func (u AssetOutputUnion) AsCaption() (v AssetCaptionOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssetOutputUnion) AsStaticRendition() (v AssetStaticRenditionOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssetOutputUnion) AsThumbnail() (v AssetThumbnailOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssetOutputUnion) AsTranscript() (v AssetTranscriptOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssetOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *AssetOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssetOutputUnionDetails is an implicit subunion of [AssetOutputUnion].
// AssetOutputUnionDetails provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [AssetOutputUnion].
type AssetOutputUnionDetails struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	// This field is a union of [shared.AssetOutputError], [ThumbnailError]
	Error        AssetOutputUnionDetailsError `json:"error"`
	LanguageCode string                       `json:"language_code"`
	// This field is from variant [AssetCaption].
	PlaylistPath string `json:"playlist_path"`
	Size         int64  `json:"size"`
	// This field is from variant [AssetCaption].
	SourceCaptionID string `json:"source_caption_id"`
	// This field is from variant [AssetCaption].
	TranscriptID string `json:"transcript_id"`
	// This field is from variant [AssetCaption].
	VttPath string `json:"vtt_path"`
	AssetID string `json:"asset_id"`
	// This field is from variant [AssetStaticRendition].
	Container constant.MP4 `json:"container"`
	// This field is from variant [AssetStaticRendition].
	Height int64 `json:"height"`
	// This field is from variant [AssetStaticRendition].
	Path string `json:"path"`
	// This field is from variant [AssetStaticRendition].
	RequestedResolution AssetStaticRenditionRequestedResolution `json:"requested_resolution"`
	// This field is from variant [AssetStaticRendition].
	Resolution string    `json:"resolution"`
	StorageID  string    `json:"storage_id"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant [AssetStaticRendition].
	VideoCodec constant.H264 `json:"video_codec"`
	// This field is from variant [AssetStaticRendition].
	Width int64 `json:"width"`
	// This field is from variant [AssetStaticRendition].
	AudioBitrate int64 `json:"audio_bitrate"`
	// This field is from variant [AssetStaticRendition].
	AudioCodec string `json:"audio_codec"`
	// This field is from variant [AssetStaticRendition].
	AudioSamplerate int64 `json:"audio_samplerate"`
	// This field is from variant [AssetStaticRendition].
	CdnURL string `json:"cdn_url"`
	// This field is from variant [AssetStaticRendition].
	Duration float64 `json:"duration"`
	// This field is from variant [AssetStaticRendition].
	Progress AssetRenditionProgress `json:"progress"`
	ReadyAt  time.Time              `json:"ready_at"`
	// This field is from variant [AssetStaticRendition].
	URL string `json:"url"`
	// This field is from variant [AssetStaticRendition].
	VideoBitrate int64 `json:"video_bitrate"`
	// This field is from variant [AssetStaticRendition].
	VideoFramerate float64 `json:"video_framerate"`
	// This field is from variant [Thumbnail].
	Files []ThumbnailFile `json:"files"`
	// This field is from variant [Thumbnail].
	Metadata map[string]string `json:"metadata"`
	// This field is from variant [Thumbnail].
	Samples []ThumbnailSample `json:"samples"`
	// This field is from variant [Thumbnail].
	Settings ThumbnailSettings `json:"settings"`
	// This field is from variant [Thumbnail].
	SourceID string `json:"source_id"`
	// This field is from variant [Thumbnail].
	StoragePath string `json:"storage_path"`
	// This field is from variant [Thumbnail].
	StoragePrefix string `json:"storage_prefix"`
	// This field is from variant [Thumbnail].
	CreatedByTokenID string `json:"created_by_token_id"`
	// This field is from variant [Thumbnail].
	CreatedByUserID string `json:"created_by_user_id"`
	// This field is from variant [Thumbnail].
	CreatedVia shared.CreatedVia `json:"created_via"`
	// This field is from variant [Thumbnail].
	DeleteAt time.Time `json:"delete_at"`
	// This field is from variant [Thumbnail].
	ExpiresAt time.Time `json:"expires_at"`
	// This field is from variant [AssetTranscript].
	Content AssetTranscriptContent `json:"content"`
	JSON    struct {
		ID                  respjson.Field
		Status              respjson.Field
		CreatedAt           respjson.Field
		Error               respjson.Field
		LanguageCode        respjson.Field
		PlaylistPath        respjson.Field
		Size                respjson.Field
		SourceCaptionID     respjson.Field
		TranscriptID        respjson.Field
		VttPath             respjson.Field
		AssetID             respjson.Field
		Container           respjson.Field
		Height              respjson.Field
		Path                respjson.Field
		RequestedResolution respjson.Field
		Resolution          respjson.Field
		StorageID           respjson.Field
		UpdatedAt           respjson.Field
		VideoCodec          respjson.Field
		Width               respjson.Field
		AudioBitrate        respjson.Field
		AudioCodec          respjson.Field
		AudioSamplerate     respjson.Field
		CdnURL              respjson.Field
		Duration            respjson.Field
		Progress            respjson.Field
		ReadyAt             respjson.Field
		URL                 respjson.Field
		VideoBitrate        respjson.Field
		VideoFramerate      respjson.Field
		Files               respjson.Field
		Metadata            respjson.Field
		Samples             respjson.Field
		Settings            respjson.Field
		SourceID            respjson.Field
		StoragePath         respjson.Field
		StoragePrefix       respjson.Field
		CreatedByTokenID    respjson.Field
		CreatedByUserID     respjson.Field
		CreatedVia          respjson.Field
		DeleteAt            respjson.Field
		ExpiresAt           respjson.Field
		Content             respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *AssetOutputUnionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssetOutputUnionDetailsError is an implicit subunion of [AssetOutputUnion].
// AssetOutputUnionDetailsError provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [AssetOutputUnion].
type AssetOutputUnionDetailsError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	JSON    struct {
		Code    respjson.Field
		Message respjson.Field
		raw     string
	} `json:"-"`
}

func (r *AssetOutputUnionDetailsError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetOutputNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Create
	// an Asset-owned caption output.
	OfCaption *AssetCaptionOutputCreateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// an Asset-owned static rendition output.
	OfStaticRendition *AssetStaticRenditionOutputCreateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// an Asset-owned thumbnail output.
	OfThumbnail *AssetThumbnailOutputCreateParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// or reuse the original-language transcript for an Asset. Caption generation
	// shares this transcript and its speech-recognition charge.
	OfTranscript *AssetTranscriptOutputCreateParams `json:",inline"`

	paramObj
}

func (u AssetOutputNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCaption, u.OfStaticRendition, u.OfThumbnail, u.OfTranscript)
}
func (r *AssetOutputNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One Asset output.
type AssetOutputNewResponseEnvelope struct {
	// A resource produced for an Asset, requested directly or by a Workflow. Its type
	// selects the details schema. Output state is independent of Asset playback
	// readiness.
	Data AssetOutputUnion `json:"data" api:"required"`
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
func (r AssetOutputNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetOutputGetParams struct {
	AssetID string `path:"assetId" api:"required" json:"-"`
	paramObj
}

// One Asset output.
type AssetOutputGetResponseEnvelope struct {
	// A resource produced for an Asset, requested directly or by a Workflow. Its type
	// selects the details schema. Output state is independent of Asset playback
	// readiness.
	Data AssetOutputUnion `json:"data" api:"required"`
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
func (r AssetOutputGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetOutputListParams struct {
	// Limit specifies the maximum number of results to return. Accepts 1 to 100.
	// Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset specifies the number of matching results to skip. Must be at least 0.
	// Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Status specifies the current output state to include. Omit it to include all
	// states.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status AssetOutputListParamsStatus `query:"status,omitzero" json:"-"`
	// Type specifies the output kind to include. Omit it to include all kinds.
	//
	// Any of "caption", "transcript", "static_rendition", "thumbnail".
	Type AssetOutputListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetOutputListParams]'s query parameters as `url.Values`.
func (r AssetOutputListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status specifies the current output state to include. Omit it to include all
// states.
type AssetOutputListParamsStatus string

const (
	AssetOutputListParamsStatusPreparing AssetOutputListParamsStatus = "preparing"
	AssetOutputListParamsStatusReady     AssetOutputListParamsStatus = "ready"
	AssetOutputListParamsStatusFailed    AssetOutputListParamsStatus = "failed"
	AssetOutputListParamsStatusExpired   AssetOutputListParamsStatus = "expired"
)

// Type specifies the output kind to include. Omit it to include all kinds.
type AssetOutputListParamsType string

const (
	AssetOutputListParamsTypeCaption         AssetOutputListParamsType = "caption"
	AssetOutputListParamsTypeTranscript      AssetOutputListParamsType = "transcript"
	AssetOutputListParamsTypeStaticRendition AssetOutputListParamsType = "static_rendition"
	AssetOutputListParamsTypeThumbnail       AssetOutputListParamsType = "thumbnail"
)

type AssetOutputDeleteParams struct {
	AssetID string `path:"assetId" api:"required" json:"-"`
	// DeleteFiles specifies whether to remove this output's published media files from
	// Storage. Defaults to false, preserving those files. Transcript deletion always
	// schedules removal of its private text and saved subtitle copy, regardless of
	// this option. The Asset, Source, and other outputs are preserved. File cleanup
	// failures retain the output record for retry; files already removed cannot be
	// restored.
	DeleteFiles param.Opt[bool] `query:"delete_files,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetOutputDeleteParams]'s query parameters as
// `url.Values`.
func (r AssetOutputDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
