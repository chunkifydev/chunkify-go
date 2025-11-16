// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/chunkify-go/internal/apijson"
	"github.com/stainless-sdks/chunkify-go/internal/apiquery"
	"github.com/stainless-sdks/chunkify-go/internal/requestconfig"
	"github.com/stainless-sdks/chunkify-go/option"
	"github.com/stainless-sdks/chunkify-go/packages/param"
	"github.com/stainless-sdks/chunkify-go/packages/respjson"
	"github.com/stainless-sdks/chunkify-go/shared"
)

// JobLogService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobLogService] method instead.
type JobLogService struct {
	Options []option.RequestOption
}

// NewJobLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobLogService(opts ...option.RequestOption) (r JobLogService) {
	r = JobLogService{}
	r.Options = opts
	return
}

// Retrieve logs for a specific job, either from the transcoder or manager service
func (r *JobLogService) List(ctx context.Context, jobID string, query JobLogListParams, opts ...option.RequestOption) (res *JobLogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("api/jobs/%s/logs", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Successful response
type JobLogListResponse struct {
	Data []JobLogListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.ResponseOk
}

// Returns the unmodified JSON received from the API
func (r JobLogListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobLogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobLogListResponseData struct {
	// Additional structured data attached to the log
	Attributes any `json:"attributes"`
	// Optional ID of the job this log is associated with
	JobID string `json:"job_id"`
	// Log level (e.g. "info", "error", "debug")
	Level string `json:"level"`
	// The log message content
	Msg string `json:"msg"`
	// Name of the service that generated the log
	Service string `json:"service"`
	// Timestamp when the log was created
	Time string `json:"time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		JobID       respjson.Field
		Level       respjson.Field
		Msg         respjson.Field
		Service     respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobLogListResponseData) RawJSON() string { return r.JSON.raw }
func (r *JobLogListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobLogListParams struct {
	// Service type (transcoder or manager)
	//
	// Any of "transcoder", "manager".
	Service JobLogListParamsService `query:"service,omitzero,required" json:"-"`
	// Transcoder ID (required if service is transcoder)
	TranscoderID param.Opt[int64] `query:"transcoder_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobLogListParams]'s query parameters as `url.Values`.
func (r JobLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Service type (transcoder or manager)
type JobLogListParamsService string

const (
	JobLogListParamsServiceTranscoder JobLogListParamsService = "transcoder"
	JobLogListParamsServiceManager    JobLogListParamsService = "manager"
)
