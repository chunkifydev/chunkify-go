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
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// ProjectService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options   []option.RequestOption
	Analytics ProjectAnalyticsService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	r.Analytics = NewProjectAnalyticsService(opts...)
	return
}

// Create a new project with the specified name. The project will be created with
// default Chunkify storage settings.
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve details of a specific project by its slug
func (r *ProjectService) Get(ctx context.Context, projectID string, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update a project's name or storage settings. Only team owners can update
// projects.
func (r *ProjectService) Update(ctx context.Context, projectID string, body ProjectUpdateParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return err
	}
	path := fmt.Sprintf("projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Retrieve a list of all projects for a team
func (r *ProjectService) List(ctx context.Context, opts ...option.RequestOption) (res *ProjectListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a project and revoke all associated access tokens. Only team owners can
// delete projects. Customer output files remain in Storage; managed temporary
// files retain their original expiry.
func (r *ProjectService) Delete(ctx context.Context, projectID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return err
	}
	path := fmt.Sprintf("projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Project struct {
	// Id is the unique identifier for the project.
	ID string `json:"id" api:"required"`
	// Defaults for new Assets, including Assets created from Uploads.
	AssetDefaults ProjectAssetDefaults `json:"asset_defaults" api:"required"`
	// Timestamp when the project was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the project
	Name string `json:"name" api:"required"`
	// Slug is the slug for the project.
	Slug string `json:"slug" api:"required"`
	// StorageId identifier where project files are stored
	StorageID string `json:"storage_id" api:"required"`
	// DefaultWorkflowId specifies the default Workflow for new Assets. Set null to
	// clear it.
	DefaultWorkflowID string `json:"default_workflow_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AssetDefaults     respjson.Field
		CreatedAt         respjson.Field
		Name              respjson.Field
		Slug              respjson.Field
		StorageID         respjson.Field
		DefaultWorkflowID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defaults for new Assets, including Assets created from Uploads.
type ProjectAssetDefaults struct {
	// Delete generated and partial files from the output Storage when a new Asset
	// fails. The Asset record and original Source are preserved. Cleanup is attempted
	// once after a delay and requires delete permission on the Storage. Defaults to
	// false. Existing Assets are unchanged.
	DeleteFilesOnFailure bool `json:"delete_files_on_failure"`
	// Maximum streaming resolution. Defaults to 1080p on Hobby and 2160p on other
	// plans. 2160p requires Starter or above. Smaller sources are not upscaled. Does
	// not affect static renditions.
	//
	// Any of "1080p", "2160p".
	MaxResolution string `json:"max_resolution"`
	// Minimum streaming resolution. Defaults to auto, with no minimum. Must not exceed
	// max_resolution. Smaller sources are not upscaled. Does not affect static
	// renditions.
	//
	// Any of "auto", "240p", "360p", "540p", "720p", "1080p", "1440p", "2160p".
	MinResolution string `json:"min_resolution"`
	// PathTemplate specifies the Asset output folder relative to Storage base_prefix.
	// Supports {assetID} and case-sensitive {metadata.<key>}; referenced values must
	// be non-empty. Ignored for Chunkify storage, which always uses {assetID}.
	// Defaults to {assetID}. Maximum template size: 1024 UTF-8 bytes. Expanded prefix:
	// 768 bytes including base_prefix and trailing slash.
	PathTemplate string `json:"path_template"`
	// Allow public preview access for new Assets without an API token. Defaults to
	// true. Existing Assets are unchanged.
	PreviewEnabled bool `json:"preview_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeleteFilesOnFailure respjson.Field
		MaxResolution        respjson.Field
		MinResolution        respjson.Field
		PathTemplate         respjson.Field
		PreviewEnabled       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAssetDefaults) RawJSON() string { return r.JSON.raw }
func (r *ProjectAssetDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing the list of projects for a team
type ProjectListResponse struct {
	// Data contains the project items
	Data []Project `json:"data" api:"required"`
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
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewParams struct {
	// Name is the name of the project, which must be between 4 and 32 characters.
	Name string `json:"name" api:"required"`
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
	Data Project `json:"data" api:"required"`
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
func (r ProjectNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ProjectNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetResponseEnvelope struct {
	// Data contains the response object
	Data Project `json:"data" api:"required"`
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
func (r ProjectGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateParams struct {
	// DefaultWorkflowId specifies the default Workflow for new Assets. Set null to
	// clear it.
	DefaultWorkflowID param.Opt[string] `json:"default_workflow_id,omitzero"`
	// Project name.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of this Project's default Storage.
	StorageID param.Opt[string] `json:"storage_id,omitzero"`
	// Defaults for new Assets and automatic Uploads. Replaces the whole object.
	// Omitted fields reset to their defaults. Existing Assets are unchanged.
	AssetDefaults ProjectUpdateParamsAssetDefaults `json:"asset_defaults,omitzero"`
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defaults for new Assets and automatic Uploads. Replaces the whole object.
// Omitted fields reset to their defaults. Existing Assets are unchanged.
type ProjectUpdateParamsAssetDefaults struct {
	// Delete generated and partial files from the output Storage when a new Asset
	// fails. The Asset record and original Source are preserved. Cleanup is attempted
	// once after a delay and requires delete permission on the Storage. Defaults to
	// false. Existing Assets are unchanged.
	DeleteFilesOnFailure param.Opt[bool] `json:"delete_files_on_failure,omitzero"`
	// PathTemplate specifies the Asset output folder relative to Storage base_prefix.
	// Supports {assetID} and case-sensitive {metadata.<key>}; referenced values must
	// be non-empty. Ignored for Chunkify storage, which always uses {assetID}.
	// Defaults to {assetID}. Maximum template size: 1024 UTF-8 bytes. Expanded prefix:
	// 768 bytes including base_prefix and trailing slash.
	PathTemplate param.Opt[string] `json:"path_template,omitzero"`
	// Allow public preview access for new Assets without an API token. Defaults to
	// true. Existing Assets are unchanged.
	PreviewEnabled param.Opt[bool] `json:"preview_enabled,omitzero"`
	// Maximum streaming resolution. Defaults to 1080p on Hobby and 2160p on other
	// plans. 2160p requires Starter or above. Smaller sources are not upscaled. Does
	// not affect static renditions.
	//
	// Any of "1080p", "2160p".
	MaxResolution string `json:"max_resolution,omitzero"`
	// Minimum streaming resolution. Defaults to auto, with no minimum. Must not exceed
	// max_resolution. Smaller sources are not upscaled. Does not affect static
	// renditions.
	//
	// Any of "auto", "240p", "360p", "540p", "720p", "1080p", "1440p", "2160p".
	MinResolution string `json:"min_resolution,omitzero"`
	paramObj
}

func (r ProjectUpdateParamsAssetDefaults) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParamsAssetDefaults
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParamsAssetDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProjectUpdateParamsAssetDefaults](
		"max_resolution", "1080p", "2160p",
	)
	apijson.RegisterFieldValidator[ProjectUpdateParamsAssetDefaults](
		"min_resolution", "auto", "240p", "360p", "540p", "720p", "1080p", "1440p", "2160p",
	)
}
