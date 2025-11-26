// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomchunkifydevchunkifygo

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

// JobFileService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobFileService] method instead.
type JobFileService struct {
	Options []option.RequestOption
}

// NewJobFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobFileService(opts ...option.RequestOption) (r JobFileService) {
	r = JobFileService{}
	r.Options = opts
	return
}

// Retrieve all files associated with a specific job
func (r *JobFileService) List(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobFileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/files", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing a list of files for a job
type JobFileListResponse struct {
	Data []APIFile `json:"data,required"`
	// Status indicates the response status "success"
	Status constant.Success `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
