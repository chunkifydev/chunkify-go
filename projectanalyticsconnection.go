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

// ProjectAnalyticsConnectionService contains methods and other services that help
// with interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectAnalyticsConnectionService] method instead.
type ProjectAnalyticsConnectionService struct {
	Options []option.RequestOption
	Keys    ProjectAnalyticsConnectionKeyService
}

// NewProjectAnalyticsConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProjectAnalyticsConnectionService(opts ...option.RequestOption) (r ProjectAnalyticsConnectionService) {
	r = ProjectAnalyticsConnectionService{}
	r.Options = opts
	r.Keys = NewProjectAnalyticsConnectionKeyService(opts...)
	return
}

// Add a hosted connection or validate and save a custom server. Only team owners
// can create connections. Retry a known failed hosted setup with the same name to
// resume its saved project; an unknown creation outcome requires operator
// reconciliation.
func (r *ProjectAnalyticsConnectionService) New(ctx context.Context, projectID string, body ProjectAnalyticsConnectionNewParams, opts ...option.RequestOption) (res *ProjectAnalyticsConnectionNewResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/analytics/connections", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve non-secret connection metadata within its owning project.
func (r *ProjectAnalyticsConnectionService) Get(ctx context.Context, connectionID string, query ProjectAnalyticsConnectionGetParams, opts ...option.RequestOption) (res *AnalyticsConnection, err error) {
	var env ProjectAnalyticsConnectionGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/analytics/connections/%s", query.ProjectID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Rename a connection or validate replacement custom credentials. Only team owners
// can update connections. Existing settings remain unchanged if validation fails.
func (r *ProjectAnalyticsConnectionService) Update(ctx context.Context, connectionID string, params ProjectAnalyticsConnectionUpdateParams, opts ...option.RequestOption) (res *AnalyticsConnection, err error) {
	var env ProjectAnalyticsConnectionUpdateResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/analytics/connections/%s", params.ProjectID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List the saved connections for this project. No connection is activated by this
// request.
func (r *ProjectAnalyticsConnectionService) List(ctx context.Context, projectID string, opts ...option.RequestOption) (res *ProjectAnalyticsConnectionListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/analytics/connections", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// For hosted connections, delete the linked analytics project and invalidate its
// keys before removing the saved connection. Stored events expire under the
// collector retention policy. For custom connections, remove only the saved
// credentials and connection; the remote server, project and data remain
// unchanged. Only team owners can delete connections. An uncertain remote result
// preserves the connection for reconciliation.
func (r *ProjectAnalyticsConnectionService) Delete(ctx context.Context, connectionID string, body ProjectAnalyticsConnectionDeleteParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return err
	}
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return err
	}
	path := fmt.Sprintf("projects/%s/analytics/connections/%s", body.ProjectID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A saved source for analytics reports. The dashboard player default is
// independent of the connection selected in reports. External players keep their
// own configuration. Only the public ingest key is included; read and management
// credentials stay private.
type AnalyticsConnection struct {
	// Id identifies this saved analytics connection.
	ID string `json:"id" api:"required"`
	// CreatedAt records when this connection was added.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// DefaultForDashboard selects this connection for Chunkify dashboard players,
	// public previews, the playground and newly generated snippets. It does not change
	// existing external players. At most one connection per project is selected.
	DefaultForDashboard bool `json:"default_for_dashboard" api:"required"`
	// Endpoint specifies the analytics server base URL. It does not configure the
	// player.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Name identifies this connection in the dashboard.
	Name string `json:"name" api:"required"`
	// ProjectId identifies the owning Chunkify project.
	ProjectID string `json:"project_id" api:"required"`
	// Status indicates whether the connection is ready to read or requires setup
	// repair.
	//
	// Any of "preparing", "ready", "failed".
	Status AnalyticsConnectionStatus `json:"status" api:"required"`
	// Type specifies whether Chunkify manages the analytics project.
	//
	// Any of "hosted", "custom".
	Type AnalyticsConnectionType `json:"type" api:"required"`
	// UpdatedAt records when this connection was last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// AllowedOrigins lists additional website origins permitted to ingest events for a
	// hosted connection. Shared installation origins remain allowed. An empty list
	// clears project-specific origins. Custom servers manage origins on their own
	// infrastructure.
	AllowedOrigins []string `json:"allowed_origins"`
	// LastCheckedAt records the most recent successful connection check.
	LastCheckedAt time.Time `json:"last_checked_at" format:"date-time"`
	// PublicIngestKey is the public browser ingestion key used by Chunkify previews
	// and player examples. It cannot read or manage analytics. Hosted connections use
	// their primary public key. Omitted when preview tracking is not configured.
	PublicIngestKey string `json:"public_ingest_key"`
	// RetentionSeconds is the last retention applied to a hosted project, in seconds.
	// Omitted for custom connections.
	RetentionSeconds int64 `json:"retention_seconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		DefaultForDashboard respjson.Field
		Endpoint            respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		Status              respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		AllowedOrigins      respjson.Field
		LastCheckedAt       respjson.Field
		PublicIngestKey     respjson.Field
		RetentionSeconds    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsConnection) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status indicates whether the connection is ready to read or requires setup
// repair.
type AnalyticsConnectionStatus string

const (
	AnalyticsConnectionStatusPreparing AnalyticsConnectionStatus = "preparing"
	AnalyticsConnectionStatusReady     AnalyticsConnectionStatus = "ready"
	AnalyticsConnectionStatusFailed    AnalyticsConnectionStatus = "failed"
)

// Type specifies whether Chunkify manages the analytics project.
type AnalyticsConnectionType string

const (
	AnalyticsConnectionTypeHosted AnalyticsConnectionType = "hosted"
	AnalyticsConnectionTypeCustom AnalyticsConnectionType = "custom"
)

type ProjectAnalyticsConnectionNewResponse struct {
	// A saved source for analytics reports. The dashboard player default is
	// independent of the connection selected in reports. External players keep their
	// own configuration. Only the public ingest key is included; read and management
	// credentials stay private.
	Data   AnalyticsConnection `json:"data" api:"required"`
	Status constant.Success    `json:"status" default:"success"`
	// PublicKey is the hosted ingestion key. It remains available as public_ingest_key
	// on the connection. Omitted for custom connections.
	PublicKey string `json:"public_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAnalyticsConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectAnalyticsConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionListResponse struct {
	// Saved connections in creation order.
	Data   []AnalyticsConnection `json:"data" api:"required"`
	Status constant.Success      `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAnalyticsConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectAnalyticsConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionNewParams struct {
	// Name identifies this connection in the dashboard.
	Name string `json:"name" api:"required"`
	// Type specifies whether Chunkify manages the analytics project.
	//
	// Any of "hosted", "custom".
	Type ProjectAnalyticsConnectionNewParamsType `json:"type,omitzero" api:"required"`
	// Endpoint specifies the custom server base URL. Required for custom connections
	// and omitted for hosted connections.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// PublicIngestKey optionally enables preview tracking for a custom connection.
	// Supply the public ingestion key for the same analytics project as read_key. It
	// is returned in connection reads and used in browser code. Omit for hosted
	// connections.
	PublicIngestKey param.Opt[string] `json:"public_ingest_key,omitzero"`
	// ReadKey supplies a read-only project credential for a custom connection.
	// Required for custom connections and omitted for hosted connections. Stored
	// encrypted and never returned.
	ReadKey param.Opt[string] `json:"read_key,omitzero"`
	paramObj
}

func (r ProjectAnalyticsConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectAnalyticsConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectAnalyticsConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type specifies whether Chunkify manages the analytics project.
type ProjectAnalyticsConnectionNewParamsType string

const (
	ProjectAnalyticsConnectionNewParamsTypeHosted ProjectAnalyticsConnectionNewParamsType = "hosted"
	ProjectAnalyticsConnectionNewParamsTypeCustom ProjectAnalyticsConnectionNewParamsType = "custom"
)

type ProjectAnalyticsConnectionGetParams struct {
	ProjectID string `path:"projectId" api:"required" json:"-"`
	paramObj
}

type ProjectAnalyticsConnectionGetResponseEnvelope struct {
	// A saved source for analytics reports. The dashboard player default is
	// independent of the connection selected in reports. External players keep their
	// own configuration. Only the public ingest key is included; read and management
	// credentials stay private.
	Data   AnalyticsConnection `json:"data" api:"required"`
	Status constant.Success    `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAnalyticsConnectionGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ProjectAnalyticsConnectionGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionUpdateParams struct {
	ProjectID string `path:"projectId" api:"required" json:"-"`
	// DefaultForDashboard selects this ready connection for Chunkify dashboard
	// players, public previews, the playground and newly generated snippets. Requires
	// a public ingest key. Setting true replaces the previous default atomically.
	// Setting false disables dashboard tracking if no default remains; another
	// connection is not selected automatically. Existing external players keep their
	// own configuration.
	DefaultForDashboard param.Opt[bool] `json:"default_for_dashboard,omitzero"`
	// Endpoint specifies a replacement custom server base URL. Supply read_key
	// together with it.
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Name identifies this connection in the dashboard.
	Name param.Opt[string] `json:"name,omitzero"`
	// PublicIngestKey optionally enables preview tracking for a custom connection.
	// Supply the public ingestion key for the same analytics project as read_key. It
	// is returned in connection reads and used in browser code. Omit for hosted
	// connections. An empty string disables preview tracking. A replacement endpoint
	// clears the saved public key unless a replacement is supplied.
	PublicIngestKey param.Opt[string] `json:"public_ingest_key,omitzero"`
	// ReadKey replaces the custom connection credential after validation. Its value is
	// never returned.
	ReadKey param.Opt[string] `json:"read_key,omitzero"`
	// AllowedOrigins lists additional website origins permitted to ingest events for a
	// hosted connection. Shared installation origins remain allowed. An empty list
	// clears project-specific origins. Custom servers manage origins on their own
	// infrastructure.
	AllowedOrigins []string `json:"allowed_origins,omitzero"`
	paramObj
}

func (r ProjectAnalyticsConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectAnalyticsConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectAnalyticsConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionUpdateResponseEnvelope struct {
	// A saved source for analytics reports. The dashboard player default is
	// independent of the connection selected in reports. External players keep their
	// own configuration. Only the public ingest key is included; read and management
	// credentials stay private.
	Data   AnalyticsConnection `json:"data" api:"required"`
	Status constant.Success    `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAnalyticsConnectionUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ProjectAnalyticsConnectionUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionDeleteParams struct {
	ProjectID string `path:"projectId" api:"required" json:"-"`
	paramObj
}
