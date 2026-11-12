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
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// WorkflowRunService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowRunService] method instead.
type WorkflowRunService struct {
	Options []option.RequestOption
	Actions WorkflowRunActionService
}

// NewWorkflowRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowRunService(opts ...option.RequestOption) (r WorkflowRunService) {
	r = WorkflowRunService{}
	r.Options = opts
	r.Actions = NewWorkflowRunActionService(opts...)
	return
}

// Retrieve a retained run with its triggering event, status, and action counts.
func (r *WorkflowRunService) Get(ctx context.Context, runID string, query WorkflowRunGetParams, opts ...option.RequestOption) (res *WorkflowRun, err error) {
	var env WorkflowRunGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.WorkflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workflows/%s/runs/%s", query.WorkflowID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List retained runs for this Workflow, newest first. Each run corresponds to one
// matching Asset event.
func (r *WorkflowRunService) List(ctx context.Context, workflowID string, query WorkflowRunListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[WorkflowRun], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workflows/%s/runs", workflowID)
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

// List retained runs for this Workflow, newest first. Each run corresponds to one
// matching Asset event.
func (r *WorkflowRunService) ListAutoPaging(ctx context.Context, workflowID string, query WorkflowRunListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[WorkflowRun] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, workflowID, query, opts...))
}

// One execution of Workflow rules for an Asset, started by an event or by updating
// to a newer revision.
type WorkflowRun struct {
	// Unique Workflow run ID.
	ID string `json:"id" api:"required"`
	// Counts of action outcomes in this run. Reused actions are included in completed.
	Actions WorkflowRunActions `json:"actions" api:"required"`
	// AssetExecutionStatus specifies the combined status of this revision for the
	// Asset across retained runs. Any failed action or run makes it failed. Pending or
	// active work, or an Asset still preparing for an unarchived Workflow, makes it
	// running. Otherwise it is completed.
	//
	// Any of "pending", "running", "completed", "failed", "cancelled".
	AssetExecutionStatus WorkflowRunAssetExecutionStatus `json:"asset_execution_status" api:"required"`
	// ID of the Asset processed by this run.
	AssetID string `json:"asset_id" api:"required"`
	// When the run completed, failed, or was cancelled. Null while the run is pending
	// or running.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Event specifies the Asset lifecycle event that triggered this run, or manual
	// when the Asset was updated to a newer revision. Action details are available
	// from the run's actions endpoint.
	Event string `json:"event" api:"required"`
	// EventId specifies the event that triggered this run. Null when the Asset was
	// updated to a newer revision.
	EventID string `json:"event_id" api:"required"`
	// Pinned Workflow revision used by this run.
	Revision int64 `json:"revision" api:"required"`
	// ID of the first run in this event chain. Equals id when this run starts the
	// chain.
	RootRunID string `json:"root_run_id" api:"required"`
	// Current status of this event run, based on its own actions.
	//
	// Any of "pending", "running", "completed", "failed", "cancelled".
	Status WorkflowRunStatus `json:"status" api:"required"`
	// When the run was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// ID of the Workflow executed by this run.
	WorkflowID string `json:"workflow_id" api:"required"`
	// Reason the run failed. Omitted when no failure message is available.
	Error string `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Actions              respjson.Field
		AssetExecutionStatus respjson.Field
		AssetID              respjson.Field
		CompletedAt          respjson.Field
		CreatedAt            respjson.Field
		Event                respjson.Field
		EventID              respjson.Field
		Revision             respjson.Field
		RootRunID            respjson.Field
		Status               respjson.Field
		UpdatedAt            respjson.Field
		WorkflowID           respjson.Field
		Error                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRun) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Counts of action outcomes in this run. Reused actions are included in completed.
type WorkflowRunActions struct {
	// Number of actions cancelled before completion.
	Cancelled int64 `json:"cancelled" api:"required"`
	// Number of actions that completed successfully, including reused outputs.
	Completed int64 `json:"completed" api:"required"`
	// Number of actions that failed.
	Failed int64 `json:"failed" api:"required"`
	// Number of completed actions that reused an existing output. Included in
	// completed.
	Reused int64 `json:"reused" api:"required"`
	// Number of actions skipped without producing an output.
	Skipped int64 `json:"skipped" api:"required"`
	// Total number of actions in the run, including unfinished actions.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		Completed   respjson.Field
		Failed      respjson.Field
		Reused      respjson.Field
		Skipped     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunActions) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssetExecutionStatus specifies the combined status of this revision for the
// Asset across retained runs. Any failed action or run makes it failed. Pending or
// active work, or an Asset still preparing for an unarchived Workflow, makes it
// running. Otherwise it is completed.
type WorkflowRunAssetExecutionStatus string

const (
	WorkflowRunAssetExecutionStatusPending   WorkflowRunAssetExecutionStatus = "pending"
	WorkflowRunAssetExecutionStatusRunning   WorkflowRunAssetExecutionStatus = "running"
	WorkflowRunAssetExecutionStatusCompleted WorkflowRunAssetExecutionStatus = "completed"
	WorkflowRunAssetExecutionStatusFailed    WorkflowRunAssetExecutionStatus = "failed"
	WorkflowRunAssetExecutionStatusCancelled WorkflowRunAssetExecutionStatus = "cancelled"
)

// Current status of this event run, based on its own actions.
type WorkflowRunStatus string

const (
	WorkflowRunStatusPending   WorkflowRunStatus = "pending"
	WorkflowRunStatusRunning   WorkflowRunStatus = "running"
	WorkflowRunStatusCompleted WorkflowRunStatus = "completed"
	WorkflowRunStatusFailed    WorkflowRunStatus = "failed"
	WorkflowRunStatusCancelled WorkflowRunStatus = "cancelled"
)

type WorkflowRunGetParams struct {
	WorkflowID string `path:"workflowId" api:"required" json:"-"`
	paramObj
}

// Response containing one Workflow run.
type WorkflowRunGetResponseEnvelope struct {
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
func (r WorkflowRunGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRunListParams struct {
	// AssetId specifies the Asset whose results to return. Omit it to include all
	// Assets.
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// EventId specifies the triggering event ID to filter by. Omit it to include all
	// events.
	EventID param.Opt[string] `query:"event_id,omitzero" json:"-"`
	// LatestPerAsset specifies whether to return only the latest matching run for each
	// Asset. Defaults to false. Pagination and total count apply to Assets when true.
	LatestPerAsset param.Opt[bool] `query:"latest_per_asset,omitzero" json:"-"`
	// Limit specifies the maximum number of results to return. Accepts 1 to 100.
	// Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset specifies the number of matching results to skip. Must be at least 0.
	// Defaults to 0.
	Offset  param.Opt[int64]             `query:"offset,omitzero" json:"-"`
	Created WorkflowRunListParamsCreated `query:"created,omitzero" json:"-"`
	// Status specifies the run status to filter by. Omit it to include all statuses.
	//
	// Any of "pending", "running", "completed", "failed", "cancelled".
	Status WorkflowRunListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowRunListParams]'s query parameters as `url.Values`.
func (r WorkflowRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRunListParamsCreated struct {
	// Created at or after this Unix timestamp in seconds.
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Created at or before this Unix timestamp in seconds.
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowRunListParamsCreated]'s query parameters as
// `url.Values`.
func (r WorkflowRunListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status specifies the run status to filter by. Omit it to include all statuses.
type WorkflowRunListParamsStatus string

const (
	WorkflowRunListParamsStatusPending   WorkflowRunListParamsStatus = "pending"
	WorkflowRunListParamsStatusRunning   WorkflowRunListParamsStatus = "running"
	WorkflowRunListParamsStatusCompleted WorkflowRunListParamsStatus = "completed"
	WorkflowRunListParamsStatusFailed    WorkflowRunListParamsStatus = "failed"
	WorkflowRunListParamsStatusCancelled WorkflowRunListParamsStatus = "cancelled"
)
