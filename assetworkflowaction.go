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
)

// Automatic video-on-demand Assets and independent H.264 MP4 static renditions.
//
// AssetWorkflowActionService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetWorkflowActionService] method instead.
type AssetWorkflowActionService struct {
	Options []option.RequestOption
}

// NewAssetWorkflowActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssetWorkflowActionService(opts ...option.RequestOption) (r AssetWorkflowActionService) {
	r = AssetWorkflowActionService{}
	r.Options = opts
	return
}

// List retained Workflow actions for an Asset, newest first. Direct output
// requests are excluded.
func (r *AssetWorkflowActionService) List(ctx context.Context, assetID string, query AssetWorkflowActionListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[AssetActionRun], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s/workflow/actions", assetID)
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

// List retained Workflow actions for an Asset, newest first. Direct output
// requests are excluded.
func (r *AssetWorkflowActionService) ListAutoPaging(ctx context.Context, assetID string, query AssetWorkflowActionListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[AssetActionRun] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, assetID, query, opts...))
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
	Input AssetActionInput `json:"input" api:"required"`
	// Current execution status of the action.
	//
	// Any of "pending", "running", "waiting", "completed", "failed", "skipped",
	// "cancelled".
	Status AssetActionStatus `json:"status" api:"required"`
	// Type of action performed.
	//
	// Any of "create_static_rendition", "create_caption", "create_transcript",
	// "translate_caption", "create_thumbnails".
	Type AssetActionType `json:"type" api:"required"`
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
	Result AssetActionResult `json:"result"`
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

type AssetWorkflowActionListParams struct {
	// LatestPerAction specifies whether to return only the latest execution of each
	// rule and action on the Asset. Defaults to false. Pagination and total count
	// apply after this selection.
	LatestPerAction param.Opt[bool] `query:"latest_per_action,omitzero" json:"-"`
	// Limit specifies the maximum number of results to return. Accepts 1 to 100.
	// Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset specifies the number of matching results to skip. Must be at least 0.
	// Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Revision specifies the Workflow revision to filter by. Omit it to include all
	// revisions.
	Revision param.Opt[int64]                     `query:"revision,omitzero" json:"-"`
	Created  AssetWorkflowActionListParamsCreated `query:"created,omitzero" json:"-"`
	// Status specifies the action status to filter by. Omit it to include all
	// statuses.
	//
	// Any of "pending", "running", "waiting", "completed", "failed", "skipped",
	// "cancelled".
	Status AssetWorkflowActionListParamsStatus `query:"status,omitzero" json:"-"`
	// Type specifies the action type to filter by. Omit it to include all action
	// types.
	//
	// Any of "create_static_rendition", "create_caption", "create_transcript",
	// "translate_caption".
	Type AssetWorkflowActionListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetWorkflowActionListParams]'s query parameters as
// `url.Values`.
func (r AssetWorkflowActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssetWorkflowActionListParamsCreated struct {
	// Created at or after this Unix timestamp in seconds.
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Created at or before this Unix timestamp in seconds.
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetWorkflowActionListParamsCreated]'s query parameters as
// `url.Values`.
func (r AssetWorkflowActionListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status specifies the action status to filter by. Omit it to include all
// statuses.
type AssetWorkflowActionListParamsStatus string

const (
	AssetWorkflowActionListParamsStatusPending   AssetWorkflowActionListParamsStatus = "pending"
	AssetWorkflowActionListParamsStatusRunning   AssetWorkflowActionListParamsStatus = "running"
	AssetWorkflowActionListParamsStatusWaiting   AssetWorkflowActionListParamsStatus = "waiting"
	AssetWorkflowActionListParamsStatusCompleted AssetWorkflowActionListParamsStatus = "completed"
	AssetWorkflowActionListParamsStatusFailed    AssetWorkflowActionListParamsStatus = "failed"
	AssetWorkflowActionListParamsStatusSkipped   AssetWorkflowActionListParamsStatus = "skipped"
	AssetWorkflowActionListParamsStatusCancelled AssetWorkflowActionListParamsStatus = "cancelled"
)

// Type specifies the action type to filter by. Omit it to include all action
// types.
type AssetWorkflowActionListParamsType string

const (
	AssetWorkflowActionListParamsTypeCreateStaticRendition AssetWorkflowActionListParamsType = "create_static_rendition"
	AssetWorkflowActionListParamsTypeCreateCaption         AssetWorkflowActionListParamsType = "create_caption"
	AssetWorkflowActionListParamsTypeCreateTranscript      AssetWorkflowActionListParamsType = "create_transcript"
	AssetWorkflowActionListParamsTypeTranslateCaption      AssetWorkflowActionListParamsType = "translate_caption"
)
