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

// WorkflowRevisionService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowRevisionService] method instead.
type WorkflowRevisionService struct {
	Options []option.RequestOption
}

// NewWorkflowRevisionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowRevisionService(opts ...option.RequestOption) (r WorkflowRevisionService) {
	r = WorkflowRevisionService{}
	r.Options = opts
	return
}

// Retrieve an immutable definition revision, including revisions of archived
// Workflows.
func (r *WorkflowRevisionService) Get(ctx context.Context, revision int64, query WorkflowRevisionGetParams, opts ...option.RequestOption) (res *WorkflowRevision, err error) {
	var env WorkflowRevisionGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.WorkflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workflows/%s/revisions/%v", query.WorkflowID, revision)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// An immutable Workflow definition pinned to Assets and Uploads at creation.
type WorkflowRevision struct {
	// When this Workflow revision was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Definition revision number, starting at 1.
	Revision int64 `json:"revision" api:"required"`
	// Rules specifies the event subscriptions and actions in this definition. Provide
	// exactly one asset.ready rule and up to 15 connected rules, with at most 32
	// actions and 65536 encoded bytes across all rules. Array order affects display
	// only.
	Rules []WorkflowRule `json:"rules" api:"required"`
	// SchemaVersion specifies the Workflow definition format. The supported version
	// is 1.
	SchemaVersion int64 `json:"schema_version" api:"required"`
	// ID of the Workflow that owns this revision.
	WorkflowID string `json:"workflow_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		Revision      respjson.Field
		Rules         respjson.Field
		SchemaVersion respjson.Field
		WorkflowID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRevision) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevision) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRevisionGetParams struct {
	WorkflowID string `path:"workflowId" api:"required" json:"-"`
	paramObj
}

// Response containing one Workflow revision.
type WorkflowRevisionGetResponseEnvelope struct {
	// The requested Workflow revision.
	Data WorkflowRevision `json:"data" api:"required"`
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
func (r WorkflowRevisionGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
