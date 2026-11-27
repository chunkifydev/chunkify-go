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

// Create a Project in your team with a name of 4 to 32 characters. New Projects
// start with Chunkify storage until you connect your own Storage.
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

// Retrieve a Project in your team by its ID, including its name, slug, default
// Storage, Asset defaults, and default Workflow.
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

// Update a Project's name, default Storage, Asset defaults, or default Workflow.
// Only team owners can update Projects, and project names must be unique.
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

// List the Projects in the team that owns the access token, with each Project's
// name, slug, default Storage, and Asset defaults.
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

type PlayerAppearance struct {
	// Dark specifies the player colours in dark mode.
	Dark PlayerPalette `json:"dark" api:"required"`
	// HiddenControls lists the player buttons to hide. Other buttons appear when the
	// browser and video support them.
	HiddenControls []PlayerControl `json:"hidden_controls" api:"required"`
	// Light specifies the player colours in light mode.
	Light PlayerPalette `json:"light" api:"required"`
	// Mode specifies which palette the player uses. auto follows the viewer's system
	// setting.
	//
	// Any of "auto", "light", "dark".
	Mode PlayerThemeMode `json:"mode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dark           respjson.Field
		HiddenControls respjson.Field
		Light          respjson.Field
		Mode           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlayerAppearance) RawJSON() string { return r.JSON.raw }
func (r *PlayerAppearance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerAppearanceParams struct {
	// Dark specifies the player colours in dark mode. Omit to use the default dark
	// colours.
	Dark PlayerPaletteParams `json:"dark,omitzero"`
	// HiddenControls lists the player buttons to hide. Defaults to an empty list.
	// Other buttons appear when the browser and video support them.
	HiddenControls []PlayerControl `json:"hidden_controls,omitzero"`
	// Light specifies the player colours in light mode. Omit to use the default light
	// colours.
	Light PlayerPaletteParams `json:"light,omitzero"`
	// Mode specifies which palette the player uses. auto follows the viewer's system
	// setting. An embed can override it with its theme URL parameter. Defaults to
	// auto.
	//
	// Any of "auto", "light", "dark".
	Mode PlayerThemeMode `json:"mode,omitzero"`
	paramObj
}

func (r PlayerAppearanceParams) MarshalJSON() (data []byte, err error) {
	type shadow PlayerAppearanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlayerAppearanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerControl string

const (
	PlayerControlCaptions   PlayerControl = "captions"
	PlayerControlSettings   PlayerControl = "settings"
	PlayerControlPip        PlayerControl = "pip"
	PlayerControlFullscreen PlayerControl = "fullscreen"
	PlayerControlCast       PlayerControl = "cast"
	PlayerControlAirplay    PlayerControl = "airplay"
	PlayerControlMute       PlayerControl = "mute"
)

type PlayerPalette struct {
	// Accent specifies the colour of progress and active controls.
	Accent string `json:"accent" api:"required"`
	// Primary specifies the colour of text and control icons.
	Primary string `json:"primary" api:"required"`
	// Secondary specifies the colour of player buttons.
	Secondary string `json:"secondary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accent      respjson.Field
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlayerPalette) RawJSON() string { return r.JSON.raw }
func (r *PlayerPalette) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Accent, Primary, Secondary are required.
type PlayerPaletteParams struct {
	// Accent specifies the colour of progress and active controls, as #rrggbb.
	Accent string `json:"accent" api:"required"`
	// Primary specifies the colour of text and control icons, as #rrggbb.
	Primary string `json:"primary" api:"required"`
	// Secondary specifies the colour of player buttons, as #rrggbb.
	Secondary string `json:"secondary" api:"required"`
	paramObj
}

func (r PlayerPaletteParams) MarshalJSON() (data []byte, err error) {
	type shadow PlayerPaletteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlayerPaletteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerThemeMode string

const (
	PlayerThemeModeAuto  PlayerThemeMode = "auto"
	PlayerThemeModeLight PlayerThemeMode = "light"
	PlayerThemeModeDark  PlayerThemeMode = "dark"
)

type Project struct {
	// Id is the unique identifier for the project.
	ID string `json:"id" api:"required"`
	// Defaults for new Assets, including Assets created from Uploads.
	AssetDefaults AssetDefaults `json:"asset_defaults" api:"required"`
	// Timestamp when the project was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the project
	Name string `json:"name" api:"required"`
	// Player colours and controls used by Chunkify-hosted playback pages, such as
	// embeds.
	PlayerAppearance PlayerAppearance `json:"player_appearance" api:"required"`
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
		PlayerAppearance  respjson.Field
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
	AssetDefaults AssetDefaultsParams `json:"asset_defaults,omitzero"`
	// Player colours and controls for Chunkify-hosted playback pages, such as embeds.
	// Replaces the whole object. Omitted fields reset to their defaults. Existing
	// embeds use the new values.
	PlayerAppearance PlayerAppearanceParams `json:"player_appearance,omitzero"`
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
