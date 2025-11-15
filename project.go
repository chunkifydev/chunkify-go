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
)

// ProjectService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	return
}

// Create a new project with the specified name. The project will be created with
// default Chunkify storage settings.
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific project by its slug
func (r *ProjectService) Get(ctx context.Context, projectID string, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("api/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a project's name or storage settings. Only team owners can update
// projects.
func (r *ProjectService) Update(ctx context.Context, projectID string, body ProjectUpdateParams, opts ...option.RequestOption) (res *ProjectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("api/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve a list of all projects with optional filtering and pagination
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *ProjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a project and revoke all associated access tokens. Only team owners can
// delete projects.
func (r *ProjectService) Delete(ctx context.Context, projectID string, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("api/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Project struct {
	// Id is the unique identifier for the project.
	ID string `json:"id"`
	// Timestamp when the project was created
	CreatedAt string `json:"created_at"`
	// Name of the project
	Name string `json:"name"`
	// Slug is the slug for the project.
	Slug string `json:"slug"`
	// StorageId identifier where project files are stored
	StorageID string `json:"storage_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		StorageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type ProjectNewResponse struct {
	Data Project `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r ProjectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type ProjectGetResponse struct {
	Data Project `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r ProjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateResponse = any

// Successful response
type ProjectListResponse struct {
	Data []Project `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDeleteResponse = any

type ProjectNewParams struct {
	// Name is the name of the project, which must be between 4 and 32 characters.
	Name string `json:"name,required"`
	paramObj
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateParams struct {
	// Name is the new name for the project, which must be between 4 and 32 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// StorageId specifies the storage configuration for the project, which must be
	// between 4 and 64 characters.
	StorageID param.Opt[string] `json:"storage_id,omitzero"`
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListParams struct {
	// Pagination limit (max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
