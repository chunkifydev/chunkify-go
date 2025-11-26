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
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
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
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "api/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve details of a specific project by its slug
func (r *ProjectService) Get(ctx context.Context, projectID string, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("api/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update a project's name or storage settings. Only team owners can update
// projects.
func (r *ProjectService) Update(ctx context.Context, projectID string, body ProjectUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("api/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
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
func (r *ProjectService) Delete(ctx context.Context, projectID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("api/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Project struct {
	// Id is the unique identifier for the project.
	ID string `json:"id,required"`
	// Timestamp when the project was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the project
	Name string `json:"name,required"`
	// Slug is the slug for the project.
	Slug string `json:"slug,required"`
	// StorageId identifier where project files are stored
	StorageID string `json:"storage_id,required"`
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

type ProjectListResponse struct {
	Data []Project `json:"data,required"`
	// Status indicates the response status "success"
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type ProjectNewResponseEnvelope struct {
	// Data contains the response object
	Data Project `json:"data,required"`
	// Status indicates the response status "success"
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ProjectNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetResponseEnvelope struct {
	// Data contains the response object
	Data Project `json:"data,required"`
	// Status indicates the response status "success"
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateParams struct {
	// Name is the name of the project. Required when storage_id is not provided.
	Name param.Opt[string] `json:"name,omitzero"`
	// StorageId is the storage id of the project. Required when name is not provided.
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
