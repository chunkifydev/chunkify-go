// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/chunkifydev/chunkify-go/internal/apijson"
	"github.com/chunkifydev/chunkify-go/internal/apiquery"
	"github.com/chunkifydev/chunkify-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-go/option"
	"github.com/chunkifydev/chunkify-go/packages/pagination"
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// WorkflowRunActionService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowRunActionService] method instead.
type WorkflowRunActionService struct {
	Options []option.RequestOption
}

// NewWorkflowRunActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowRunActionService(opts ...option.RequestOption) (r WorkflowRunActionService) {
	r = WorkflowRunActionService{}
	r.Options = opts
	return
}

// Retrieve one action execution in a retained Workflow run, including its resolved
// inputs, status, and result.
func (r *WorkflowRunActionService) Get(ctx context.Context, actionRunID string, query WorkflowRunActionGetParams, opts ...option.RequestOption) (res *AssetActionRun, err error) {
	var env WorkflowRunActionGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.WorkflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return nil, err
	}
	if query.RunID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	if actionRunID == "" {
		err = errors.New("missing required actionRunId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workflows/%s/runs/%s/actions/%s", query.WorkflowID, query.RunID, actionRunID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List actions in a retained Workflow run, newest first. Includes resolved inputs,
// statuses, output references, and skip or failure reasons.
func (r *WorkflowRunActionService) List(ctx context.Context, runID string, params WorkflowRunActionListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[AssetActionRun], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.WorkflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workflows/%s/runs/%s/actions", params.WorkflowID, runID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List actions in a retained Workflow run, newest first. Includes resolved inputs,
// statuses, output references, and skip or failure reasons.
func (r *WorkflowRunActionService) ListAutoPaging(ctx context.Context, runID string, params WorkflowRunActionListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[AssetActionRun] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, runID, params, opts...))
}

type WorkflowRunActionGetParams struct {
	WorkflowID string `path:"workflowId" api:"required" json:"-"`
	RunID      string `path:"runId" api:"required" json:"-"`
	paramObj
}

// Response containing one Asset action run.
type WorkflowRunActionGetResponseEnvelope struct {
	// The requested Asset action run.
	Data AssetActionRun `json:"data" api:"required"`
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
func (r WorkflowRunActionGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunActionGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRunActionListParams struct {
	WorkflowID string `path:"workflowId" api:"required" json:"-"`
	// AssetId specifies the Asset whose results to return. Omit it to include all
	// Assets.
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// Limit specifies the maximum number of results to return. Accepts 1 to 100.
	// Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset specifies the number of matching results to skip. Must be at least 0.
	// Defaults to 0.
	Offset  param.Opt[int64]                   `query:"offset,omitzero" json:"-"`
	Created WorkflowRunActionListParamsCreated `query:"created,omitzero" json:"-"`
	// Status specifies the action status to filter by. Omit it to include all
	// statuses.
	//
	// Any of "pending", "running", "waiting", "completed", "failed", "skipped",
	// "cancelled".
	Status WorkflowRunActionListParamsStatus `query:"status,omitzero" json:"-"`
	// Type specifies the action type to filter by. Omit it to include all action
	// types.
	//
	// Any of "create_static_rendition", "create_caption", "create_transcript",
	// "translate_caption".
	Type WorkflowRunActionListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowRunActionListParams]'s query parameters as
// `url.Values`.
func (r WorkflowRunActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRunActionListParamsCreated struct {
	// Created at or after this Unix timestamp in seconds.
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Created at or before this Unix timestamp in seconds.
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowRunActionListParamsCreated]'s query parameters as
// `url.Values`.
func (r WorkflowRunActionListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status specifies the action status to filter by. Omit it to include all
// statuses.
type WorkflowRunActionListParamsStatus string

const (
	WorkflowRunActionListParamsStatusPending   WorkflowRunActionListParamsStatus = "pending"
	WorkflowRunActionListParamsStatusRunning   WorkflowRunActionListParamsStatus = "running"
	WorkflowRunActionListParamsStatusWaiting   WorkflowRunActionListParamsStatus = "waiting"
	WorkflowRunActionListParamsStatusCompleted WorkflowRunActionListParamsStatus = "completed"
	WorkflowRunActionListParamsStatusFailed    WorkflowRunActionListParamsStatus = "failed"
	WorkflowRunActionListParamsStatusSkipped   WorkflowRunActionListParamsStatus = "skipped"
	WorkflowRunActionListParamsStatusCancelled WorkflowRunActionListParamsStatus = "cancelled"
)

// Type specifies the action type to filter by. Omit it to include all action
// types.
type WorkflowRunActionListParamsType string

const (
	WorkflowRunActionListParamsTypeCreateStaticRendition WorkflowRunActionListParamsType = "create_static_rendition"
	WorkflowRunActionListParamsTypeCreateCaption         WorkflowRunActionListParamsType = "create_caption"
	WorkflowRunActionListParamsTypeCreateTranscript      WorkflowRunActionListParamsType = "create_transcript"
	WorkflowRunActionListParamsTypeTranslateCaption      WorkflowRunActionListParamsType = "translate_caption"
)
