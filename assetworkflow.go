// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/chunkifydev/chunkify-go/internal/apijson"
	"github.com/chunkifydev/chunkify-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-go/option"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// Automatic video-on-demand Assets and independent H.264 MP4 static renditions.
//
// AssetWorkflowService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetWorkflowService] method instead.
type AssetWorkflowService struct {
	Options []option.RequestOption
	// Automatic video-on-demand Assets and independent H.264 MP4 static renditions.
	Actions AssetWorkflowActionService
}

// NewAssetWorkflowService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssetWorkflowService(opts ...option.RequestOption) (r AssetWorkflowService) {
	r = AssetWorkflowService{}
	r.Options = opts
	r.Actions = NewAssetWorkflowActionService(opts...)
	return
}

// Update a ready Asset to the latest revision of its selected Workflow. Existing
// outputs are preserved and missing outputs are created.
func (r *AssetWorkflowService) Update(ctx context.Context, assetID string, opts ...option.RequestOption) (res *WorkflowRun, err error) {
	var env AssetWorkflowUpdateResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s/workflow/update", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// One execution of an Asset action, started by a Workflow or a direct output
// request.
type AssetActionRun struct {
	// Unique action execution ID, prefixed with acr\_.
	ID string `json:"id" api:"required"`
	// ID of the Asset processed by this action.
	AssetID string `json:"asset_id" api:"required"`
	// When the action completed, failed, was skipped, or was cancelled. Null while
	// work remains.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// When the action was accepted.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resolved parameters and input references for this execution.
	Input AssetActionRunInput `json:"input" api:"required"`
	// Current execution status of the action.
	//
	// Any of "pending", "running", "waiting", "completed", "failed", "skipped",
	// "cancelled".
	Status AssetActionRunStatus `json:"status" api:"required"`
	// Type of action performed.
	//
	// Any of "create_static_rendition", "create_caption", "create_transcript",
	// "translate_caption", "create_thumbnails".
	Type AssetActionRunType `json:"type" api:"required"`
	// When the action was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// ID of the Workflow run that started the action. Null for direct output requests.
	WorkflowRunID string `json:"workflow_run_id" api:"required"`
	// Action key within the rule. Omitted for direct output requests.
	ActionID string `json:"action_id"`
	// Reason the action failed. Omitted when no failure message is available.
	Error string `json:"error"`
	// Machine-readable reason for waiting, skipping, failure, or cancellation. Omitted
	// when no reason applies.
	Reason string `json:"reason"`
	// Output reference and reuse flag. Included when the action completes
	// successfully.
	Result AssetActionRunResult `json:"result"`
	// Rule key from the Workflow definition. Omitted for direct output requests.
	RuleID string `json:"rule_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AssetID       respjson.Field
		CompletedAt   respjson.Field
		CreatedAt     respjson.Field
		Input         respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		WorkflowRunID respjson.Field
		ActionID      respjson.Field
		Error         respjson.Field
		Reason        respjson.Field
		Result        respjson.Field
		RuleID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetActionRun) RawJSON() string { return r.JSON.raw }
func (r *AssetActionRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resolved parameters and input references for this execution.
type AssetActionRunInput struct {
	// LanguageCode specifies the target language for the translated subtitle track,
	// such as en or fr. The same code appears in the caption and HLS manifest.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode string `json:"language_code"`
	// Resolution specifies the H.264 MP4 output resolution. highest selects the
	// largest supported output without upscaling. Hobby supports up to 1080p; other
	// plans support up to 2160p. Streaming resolution limits do not apply.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution string `json:"resolution"`
	// ID of the original caption to translate. Included only for translation actions.
	SourceCaptionID string `json:"source_caption_id"`
	// Thumbnail settings. Included only for create_thumbnails actions.
	Thumbnails WorkflowThumbnailParamsResp `json:"thumbnails"`
	// TranscriptId specifies the original transcript consumed by caption publication.
	// Omitted for actions without a transcript input.
	TranscriptID string `json:"transcript_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageCode    respjson.Field
		Resolution      respjson.Field
		SourceCaptionID respjson.Field
		Thumbnails      respjson.Field
		TranscriptID    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetActionRunInput) RawJSON() string { return r.JSON.raw }
func (r *AssetActionRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current execution status of the action.
type AssetActionRunStatus string

const (
	AssetActionRunStatusPending   AssetActionRunStatus = "pending"
	AssetActionRunStatusRunning   AssetActionRunStatus = "running"
	AssetActionRunStatusWaiting   AssetActionRunStatus = "waiting"
	AssetActionRunStatusCompleted AssetActionRunStatus = "completed"
	AssetActionRunStatusFailed    AssetActionRunStatus = "failed"
	AssetActionRunStatusSkipped   AssetActionRunStatus = "skipped"
	AssetActionRunStatusCancelled AssetActionRunStatus = "cancelled"
)

// Type of action performed.
type AssetActionRunType string

const (
	AssetActionRunTypeCreateStaticRendition AssetActionRunType = "create_static_rendition"
	AssetActionRunTypeCreateCaption         AssetActionRunType = "create_caption"
	AssetActionRunTypeCreateTranscript      AssetActionRunType = "create_transcript"
	AssetActionRunTypeTranslateCaption      AssetActionRunType = "translate_caption"
	AssetActionRunTypeCreateThumbnails      AssetActionRunType = "create_thumbnails"
)

// Output reference and reuse flag. Included when the action completes
// successfully.
type AssetActionRunResult struct {
	// ID of the caption, static rendition, or thumbnail request produced by the
	// action.
	OutputID string `json:"output_id" api:"required"`
	// Type of output referenced by output_id. Caption generation and translation both
	// produce caption outputs.
	//
	// Any of "caption", "static_rendition", "thumbnail", "transcript".
	OutputType string `json:"output_type" api:"required"`
	// Whether the action reused an existing output instead of creating one.
	Reused bool `json:"reused" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputID    respjson.Field
		OutputType  respjson.Field
		Reused      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetActionRunResult) RawJSON() string { return r.JSON.raw }
func (r *AssetActionRunResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing one Workflow run.
type AssetWorkflowUpdateResponseEnvelope struct {
	// The requested Workflow run.
	Data WorkflowRun `json:"data" api:"required"`
	// Status of the API request. Always success for this response.
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
func (r AssetWorkflowUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AssetWorkflowUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
