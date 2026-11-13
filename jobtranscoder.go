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
	"github.com/chunkifydev/chunkify-go/shared"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// JobTranscoderService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobTranscoderService] method instead.
type JobTranscoderService struct {
	Options []option.RequestOption
}

// NewJobTranscoderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobTranscoderService(opts ...option.RequestOption) (r JobTranscoderService) {
	r = JobTranscoderService{}
	r.Options = opts
	return
}

// Retrieve all the transcoders statuses for a specific job. Deprecated: Jobs are a
// legacy resource. Use Assets, Transcodings or Thumbnails instead.
//
// Deprecated: deprecated
func (r *JobTranscoderService) List(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobTranscoderListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s/transcoders", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response containing a list of transcoders for a job
type JobTranscoderListResponse struct {
	Data []shared.TranscoderStatus `json:"data" api:"required"`
	// Status indicates the response status "success"
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
func (r JobTranscoderListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobTranscoderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
