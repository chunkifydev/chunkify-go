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
// [AssetOutputCaption], [AssetOutputStaticRendition], [AssetOutputThumbnail],
// [AssetOutputTranscript].
//
// Use the [AssetOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssetOutputUnion struct {
	ID      string `json:"id"`
	AssetID string `json:"asset_id"`
	// This field is a union of [AssetOutputCaptionDetails],
	// [AssetOutputStaticRenditionDetails], [Thumbnail], [AssetOutputTranscriptDetails]
	Details AssetOutputUnionDetails `json:"details"`
	Status  string                  `json:"status"`
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

func (AssetOutputCaption) implAssetOutputUnion()         {}
func (AssetOutputStaticRendition) implAssetOutputUnion() {}
func (AssetOutputThumbnail) implAssetOutputUnion()       {}
func (AssetOutputTranscript) implAssetOutputUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AssetOutputUnion.AsAny().(type) {
//	case chunkify.AssetOutputCaption:
//	case chunkify.AssetOutputStaticRendition:
//	case chunkify.AssetOutputThumbnail:
//	case chunkify.AssetOutputTranscript:
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

func (u AssetOutputUnion) AsCaption() (v AssetOutputCaption) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssetOutputUnion) AsStaticRendition() (v AssetOutputStaticRendition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssetOutputUnion) AsThumbnail() (v AssetOutputThumbnail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssetOutputUnion) AsTranscript() (v AssetOutputTranscript) {
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
	// This field is from variant [AssetOutputCaptionDetails].
	PlaylistPath string `json:"playlist_path"`
	Size         int64  `json:"size"`
	// This field is from variant [AssetOutputCaptionDetails].
	SourceCaptionID string `json:"source_caption_id"`
	// This field is from variant [AssetOutputCaptionDetails].
	TranscriptID string `json:"transcript_id"`
	// This field is from variant [AssetOutputCaptionDetails].
	VttPath string `json:"vtt_path"`
	AssetID string `json:"asset_id"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	Container constant.MP4 `json:"container"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	Height int64 `json:"height"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	Path string `json:"path"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	RequestedResolution string `json:"requested_resolution"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	Resolution string    `json:"resolution"`
	StorageID  string    `json:"storage_id"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	VideoCodec constant.H264 `json:"video_codec"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	Width int64 `json:"width"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	AudioBitrate int64 `json:"audio_bitrate"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	AudioCodec string `json:"audio_codec"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	AudioSamplerate int64 `json:"audio_samplerate"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	CdnURL string `json:"cdn_url"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	Duration float64 `json:"duration"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	Progress AssetOutputStaticRenditionDetailsProgress `json:"progress"`
	ReadyAt  time.Time                                 `json:"ready_at"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	URL string `json:"url"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
	VideoBitrate int64 `json:"video_bitrate"`
	// This field is from variant [AssetOutputStaticRenditionDetails].
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
	CreatedVia ThumbnailCreatedVia `json:"created_via"`
	// This field is from variant [Thumbnail].
	DeleteAt time.Time `json:"delete_at"`
	// This field is from variant [Thumbnail].
	ExpiresAt time.Time `json:"expires_at"`
	// This field is from variant [AssetOutputTranscriptDetails].
	Content AssetOutputTranscriptDetailsContent `json:"content"`
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

// An Asset-owned caption output and its canonical resource details.
type AssetOutputCaption struct {
	// Id identifies the Asset output, prefixed with aop\_. The type field identifies
	// the output kind.
	ID string `json:"id" api:"required"`
	// AssetId identifies the owning Asset.
	AssetID string `json:"asset_id" api:"required"`
	// Details contains the canonical typed output resource, including its available
	// delivery fields.
	Details AssetOutputCaptionDetails `json:"details" api:"required"`
	// Status specifies the output processing state. Expired means temporary output
	// files are no longer available.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status string `json:"status" api:"required"`
	// Type identifies the output kind.
	Type constant.Caption `json:"type" default:"caption"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetID     respjson.Field
		Details     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputCaption) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details contains the canonical typed output resource, including its available
// delivery fields.
type AssetOutputCaptionDetails struct {
	// Unique Asset output ID, prefixed with aop\_.
	ID string `json:"id" api:"required"`
	// Caption processing status.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status string `json:"status" api:"required"`
	// Time this caption output was requested.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Reason caption processing failed.
	Error shared.AssetOutputError `json:"error"`
	// Primary language of this subtitle track, such as en or fr. Omitted when no
	// speech was detected or processing has not finished.
	LanguageCode string `json:"language_code"`
	// Subtitle media playlist path in Storage. Available when speech was detected.
	PlaylistPath string `json:"playlist_path"`
	// Combined subtitle playlist and WebVTT size in bytes. Available when ready.
	Size int64 `json:"size"`
	// ID of the caption track translated to create this track. Omitted for speech
	// transcription.
	SourceCaptionID string `json:"source_caption_id"`
	// TranscriptId identifies the original-language transcript used to generate this
	// caption. Omitted for translated captions.
	TranscriptID string `json:"transcript_id"`
	// WebVTT subtitle segment path in Storage. Available when speech was detected.
	VttPath string `json:"vtt_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Status          respjson.Field
		CreatedAt       respjson.Field
		Error           respjson.Field
		LanguageCode    respjson.Field
		PlaylistPath    respjson.Field
		Size            respjson.Field
		SourceCaptionID respjson.Field
		TranscriptID    respjson.Field
		VttPath         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputCaptionDetails) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputCaptionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Asset-owned static rendition output and its canonical resource details.
type AssetOutputStaticRendition struct {
	// Id identifies the Asset output, prefixed with aop\_. The type field identifies
	// the output kind.
	ID string `json:"id" api:"required"`
	// AssetId identifies the owning Asset.
	AssetID string `json:"asset_id" api:"required"`
	// Details contains the canonical typed output resource, including its available
	// delivery fields.
	Details AssetOutputStaticRenditionDetails `json:"details" api:"required"`
	// Status specifies the output processing state. Expired means temporary output
	// files are no longer available.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status string `json:"status" api:"required"`
	// Type identifies the output kind.
	Type constant.StaticRendition `json:"type" default:"static_rendition"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetID     respjson.Field
		Details     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputStaticRendition) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputStaticRendition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details contains the canonical typed output resource, including its available
// delivery fields.
type AssetOutputStaticRenditionDetails struct {
	// Unique Asset output ID, prefixed with aop\_.
	ID string `json:"id" api:"required"`
	// Parent Asset ID.
	AssetID string `json:"asset_id" api:"required"`
	// Media container.
	Container constant.MP4 `json:"container" default:"mp4"`
	// When the rendition was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Output height in pixels.
	Height int64 `json:"height" api:"required"`
	// MP4 path in Storage. Recreating the same resolution reuses this path and
	// overwrites any retained file.
	Path string `json:"path" api:"required"`
	// Requested resolution selector.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	RequestedResolution string `json:"requested_resolution" api:"required"`
	// Selected output resolution. May be smaller than 240p for small sources. Use
	// width and height for pixel dimensions.
	Resolution string `json:"resolution" api:"required"`
	// Processing status of this MP4, independent of streaming outputs.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status string `json:"status" api:"required"`
	// Storage containing this rendition, inherited from its Asset.
	StorageID string `json:"storage_id" api:"required"`
	// When the rendition was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Video codec.
	VideoCodec constant.H264 `json:"video_codec" default:"h264"`
	// Output width in pixels.
	Width int64 `json:"width" api:"required"`
	// Measured audio bitrate in bits per second. Omitted until known or when there is
	// no audio.
	AudioBitrate int64 `json:"audio_bitrate"`
	// Measured audio codec. Omitted until known or when there is no audio.
	AudioCodec string `json:"audio_codec"`
	// Measured audio sample rate in Hz. Omitted until known or when there is no audio.
	AudioSamplerate int64 `json:"audio_samplerate"`
	// Unsigned MP4 CDN URL. Available when ready and unexpired and Storage has a CDN
	// URL.
	CdnURL string `json:"cdn_url" format:"uri"`
	// Output duration in seconds. Omitted until ready.
	Duration float64 `json:"duration"`
	// A customer-facing Asset output processing failure.
	Error shared.AssetOutputError `json:"error"`
	// Progress specifies the current processing stage and encoding percentage. Omitted
	// after processing ends.
	Progress AssetOutputStaticRenditionDetailsProgress `json:"progress"`
	// When the MP4 became available.
	ReadyAt time.Time `json:"ready_at" format:"date-time"`
	// File size in bytes. Omitted until ready.
	Size int64 `json:"size"`
	// Signed MP4 download URL, available when ready and unexpired. Valid for up to ten
	// minutes. Retrieve the output to refresh it. Independent of preview_enabled.
	URL string `json:"url" format:"uri"`
	// Measured video bitrate in bits per second. Omitted until ready.
	VideoBitrate int64 `json:"video_bitrate"`
	// Measured video frame rate in frames per second. Omitted until known.
	VideoFramerate float64 `json:"video_framerate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AssetID             respjson.Field
		Container           respjson.Field
		CreatedAt           respjson.Field
		Height              respjson.Field
		Path                respjson.Field
		RequestedResolution respjson.Field
		Resolution          respjson.Field
		Status              respjson.Field
		StorageID           respjson.Field
		UpdatedAt           respjson.Field
		VideoCodec          respjson.Field
		Width               respjson.Field
		AudioBitrate        respjson.Field
		AudioCodec          respjson.Field
		AudioSamplerate     respjson.Field
		CdnURL              respjson.Field
		Duration            respjson.Field
		Error               respjson.Field
		Progress            respjson.Field
		ReadyAt             respjson.Field
		Size                respjson.Field
		URL                 respjson.Field
		VideoBitrate        respjson.Field
		VideoFramerate      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputStaticRenditionDetails) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputStaticRenditionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Progress specifies the current processing stage and encoding percentage. Omitted
// after processing ends.
type AssetOutputStaticRenditionDetailsProgress struct {
	// Stage specifies the current preparation activity. Finalization and streaming
	// publication can overlap encoding; finalizing is shown once encoding has
	// completed.
	//
	// Any of "queued", "starting", "encoding", "finalizing", "uploading".
	Stage string `json:"stage" api:"required"`
	// Percent specifies the percentage of the video encoded so far. Present only
	// during encoding when video duration is known. It may decrease if processing
	// restarts. It does not estimate total preparation time.
	Percent float64 `json:"percent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stage       respjson.Field
		Percent     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputStaticRenditionDetailsProgress) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputStaticRenditionDetailsProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Asset-owned thumbnail output and its canonical resource details.
type AssetOutputThumbnail struct {
	// Id identifies the Asset output, prefixed with aop\_. The type field identifies
	// the output kind.
	ID string `json:"id" api:"required"`
	// AssetId identifies the owning Asset.
	AssetID string `json:"asset_id" api:"required"`
	// Details contains the canonical typed output resource, including its available
	// delivery fields.
	Details Thumbnail `json:"details" api:"required"`
	// Status specifies the output processing state. Expired means temporary output
	// files are no longer available.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status string `json:"status" api:"required"`
	// Type identifies the output kind.
	Type constant.Thumbnail `json:"type" default:"thumbnail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetID     respjson.Field
		Details     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputThumbnail) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Asset-owned transcript output and its content.
type AssetOutputTranscript struct {
	// Id identifies the Asset output, prefixed with aop\_.
	ID string `json:"id" api:"required"`
	// AssetId identifies the owning Asset.
	AssetID string `json:"asset_id" api:"required"`
	// An original-language transcript. Content is included in ready output responses
	// and lifecycle event snapshots. Transcripts have no download or CDN URL.
	Details AssetOutputTranscriptDetails `json:"details" api:"required"`
	// Status specifies the output processing state. Expired means temporary output
	// files are no longer available.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status string `json:"status" api:"required"`
	// Type identifies the output kind.
	Type constant.Transcript `json:"type" default:"transcript"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetID     respjson.Field
		Details     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputTranscript) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An original-language transcript. Content is included in ready output responses
// and lifecycle event snapshots. Transcripts have no download or CDN URL.
type AssetOutputTranscriptDetails struct {
	// Id identifies the Asset output, prefixed with aop\_.
	ID string `json:"id" api:"required"`
	// Status specifies the output processing state. Expired means temporary output
	// files are no longer available.
	//
	// Any of "preparing", "ready", "failed", "expired".
	Status string `json:"status" api:"required"`
	// Content contains the complete text and timed segments. Included when ready;
	// omitted while preparing, failed or expired.
	Content AssetOutputTranscriptDetailsContent `json:"content"`
	// A customer-facing Asset output processing failure.
	Error shared.AssetOutputError `json:"error"`
	// LanguageCode identifies the detected primary language, such as en. Omitted
	// before recognition or when no speech is detected.
	LanguageCode string `json:"language_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Status       respjson.Field
		Content      respjson.Field
		Error        respjson.Field
		LanguageCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputTranscriptDetails) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputTranscriptDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content contains the complete text and timed segments. Included when ready;
// omitted while preparing, failed or expired.
type AssetOutputTranscriptDetailsContent struct {
	// Segments contains timed passages in chronological order. Individual word
	// timestamps are not included.
	Segments []AssetOutputTranscriptDetailsContentSegment `json:"segments" api:"required"`
	// Text contains the complete recognized speech in the original language.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Segments    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputTranscriptDetailsContent) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputTranscriptDetailsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A timed passage of recognized speech. Times are seconds from the start of the
// Asset.
type AssetOutputTranscriptDetailsContentSegment struct {
	// End specifies the passage end time in seconds.
	End float64 `json:"end" api:"required"`
	// Start specifies the passage start time in seconds.
	Start float64 `json:"start" api:"required"`
	// Text contains the recognized passage, including punctuation.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputTranscriptDetailsContentSegment) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputTranscriptDetailsContentSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetOutputNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Create
	// an Asset-owned caption output.
	OfCaption *AssetOutputNewParamsOutputCaption `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// an Asset-owned static rendition output.
	OfStaticRendition *AssetOutputNewParamsOutputStaticRendition `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// an Asset-owned thumbnail output.
	OfThumbnail *AssetOutputNewParamsOutputThumbnail `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// or reuse the original-language transcript for an Asset. Caption generation
	// shares this transcript and its speech-recognition charge.
	OfTranscript *AssetOutputNewParamsOutputTranscript `json:",inline"`

	paramObj
}

func (u AssetOutputNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCaption, u.OfStaticRendition, u.OfThumbnail, u.OfTranscript)
}
func (r *AssetOutputNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create an Asset-owned caption output.
//
// The properties Params, Type are required.
type AssetOutputNewParamsOutputCaption struct {
	// Params specifies the settings for this output type.
	Params AssetOutputNewParamsOutputCaptionParams `json:"params,omitzero" api:"required"`
	// Type specifies the output kind to create.
	//
	// This field can be elided, and will marshal its zero value as "caption".
	Type constant.Caption `json:"type" default:"caption"`
	paramObj
}

func (r AssetOutputNewParamsOutputCaption) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputCaption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the settings for this output type.
type AssetOutputNewParamsOutputCaptionParams struct {
	// SourceCaptionId specifies the caption on this Asset to translate. Supply
	// language_code with this field. Omit both fields to generate original captions.
	SourceCaptionID param.Opt[string] `json:"source_caption_id,omitzero"`
	// LanguageCode specifies the target language for a translated caption. Supply
	// source_caption_id with this field. Omit both fields to transcribe the Asset
	// audio with automatic language detection.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode string `json:"language_code,omitzero"`
	paramObj
}

func (r AssetOutputNewParamsOutputCaptionParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputCaptionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputCaptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssetOutputNewParamsOutputCaptionParams](
		"language_code", "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi", "ru", "tr", "pl", "sv",
	)
}

// Create an Asset-owned static rendition output.
//
// The properties Params, Type are required.
type AssetOutputNewParamsOutputStaticRendition struct {
	// Params specifies the settings for this output type.
	Params AssetOutputNewParamsOutputStaticRenditionParams `json:"params,omitzero" api:"required"`
	// Type specifies the output kind to create.
	//
	// This field can be elided, and will marshal its zero value as "static_rendition".
	Type constant.StaticRendition `json:"type" default:"static_rendition"`
	paramObj
}

func (r AssetOutputNewParamsOutputStaticRendition) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputStaticRendition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputStaticRendition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the settings for this output type.
//
// The property Resolution is required.
type AssetOutputNewParamsOutputStaticRenditionParams struct {
	// H.264 MP4 resolution. highest selects the largest supported output without
	// upscaling. Hobby supports up to 1080p; other plans support up to 2160p.
	// Independent of the streaming resolution limits.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution string `json:"resolution,omitzero" api:"required"`
	paramObj
}

func (r AssetOutputNewParamsOutputStaticRenditionParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputStaticRenditionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputStaticRenditionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssetOutputNewParamsOutputStaticRenditionParams](
		"resolution", "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p", "2160p",
	)
}

// Create an Asset-owned thumbnail output.
//
// The properties Params, Type are required.
type AssetOutputNewParamsOutputThumbnail struct {
	// Params specifies the settings for this output type.
	Params AssetOutputNewParamsOutputThumbnailParams `json:"params,omitzero" api:"required"`
	// Type specifies the output kind to create.
	//
	// This field can be elided, and will marshal its zero value as "thumbnail".
	Type constant.Thumbnail `json:"type" default:"thumbnail"`
	paramObj
}

func (r AssetOutputNewParamsOutputThumbnail) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputThumbnail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the settings for this output type.
type AssetOutputNewParamsOutputThumbnailParams struct {
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
	Format string `json:"format,omitzero"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output string `json:"output,omitzero"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range AssetOutputNewParamsOutputThumbnailParamsRange `json:"range,omitzero"`
	// Timestamps specifies 1 to 100 Source-relative target times in seconds. Every
	// timestamp must be within the selected range, including its start and excluding
	// its end. Repeated targets produce separate outputs.
	Timestamps []float64 `json:"timestamps,omitzero"`
	paramObj
}

func (r AssetOutputNewParamsOutputThumbnailParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputThumbnailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputThumbnailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssetOutputNewParamsOutputThumbnailParams](
		"format", "jpeg", "webp",
	)
	apijson.RegisterFieldValidator[AssetOutputNewParamsOutputThumbnailParams](
		"output", "images", "sprites",
	)
}

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type AssetOutputNewParamsOutputThumbnailParamsRange struct {
	// End specifies the exclusive upper bound in seconds. Defaults to the Source
	// duration.
	End param.Opt[float64] `json:"end,omitzero"`
	// Start specifies the inclusive lower bound in seconds. Defaults to zero.
	Start param.Opt[float64] `json:"start,omitzero"`
	paramObj
}

func (r AssetOutputNewParamsOutputThumbnailParamsRange) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputThumbnailParamsRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputThumbnailParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create or reuse the original-language transcript for an Asset. Caption
// generation shares this transcript and its speech-recognition charge.
//
// The properties Params, Type are required.
type AssetOutputNewParamsOutputTranscript struct {
	// Parameters for an original-language transcript. Pass an empty object; the spoken
	// language is detected automatically.
	Params any `json:"params,omitzero" api:"required"`
	// Type specifies the output kind.
	//
	// This field can be elided, and will marshal its zero value as "transcript".
	Type constant.Transcript `json:"type" default:"transcript"`
	paramObj
}

func (r AssetOutputNewParamsOutputTranscript) MarshalJSON() (data []byte, err error) {
	type shadow AssetOutputNewParamsOutputTranscript
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetOutputNewParamsOutputTranscript) UnmarshalJSON(data []byte) error {
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
