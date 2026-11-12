// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/chunkifydev/chunkify-go/internal/apiquery"
	"github.com/chunkifydev/chunkify-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-go/option"
	"github.com/chunkifydev/chunkify-go/packages/pagination"
	"github.com/chunkifydev/chunkify-go/packages/param"
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
