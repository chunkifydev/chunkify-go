// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

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
