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
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// ProjectAnalyticsConnectionKeyService contains methods and other services that
// help with interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectAnalyticsConnectionKeyService] method instead.
type ProjectAnalyticsConnectionKeyService struct {
	Options []option.RequestOption
}

// NewProjectAnalyticsConnectionKeyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProjectAnalyticsConnectionKeyService(opts ...option.RequestOption) (r ProjectAnalyticsConnectionKeyService) {
	r = ProjectAnalyticsConnectionKeyService{}
	r.Options = opts
	return
}

// Generate a public key for players to send events, or a read-only secret key for
// a dashboard or MCP client to query a hosted Chunkify Analytics server. Save
// secret keys when returned. Requires a team owner and a Team Access Key.
func (r *ProjectAnalyticsConnectionKeyService) New(ctx context.Context, connectionID string, params ProjectAnalyticsConnectionKeyNewParams, opts ...option.RequestOption) (res *ProjectAnalyticsConnectionKeyNewResponse, err error) {
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
	path := fmt.Sprintf("projects/%s/analytics/connections/%s/keys", params.ProjectID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List keys for a hosted Chunkify Analytics server: the public key used by players
// to send events, and read-only secret key records for dashboard and MCP access.
// Secret values are not returned. Requires a team owner and a Team Access Key.
func (r *ProjectAnalyticsConnectionKeyService) List(ctx context.Context, connectionID string, query ProjectAnalyticsConnectionKeyListParams, opts ...option.RequestOption) (res *ProjectAnalyticsConnectionKeyListResponse, err error) {
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
	path := fmt.Sprintf("projects/%s/analytics/connections/%s/keys", query.ProjectID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rotate a public or read-only secret key for a hosted Chunkify Analytics server.
// Update the player, dashboard, or MCP client with the new key; the old key stops
// working immediately. Requires a team owner and a Team Access Key.
func (r *ProjectAnalyticsConnectionKeyService) Rotate(ctx context.Context, keyID string, body ProjectAnalyticsConnectionKeyRotateParams, opts ...option.RequestOption) (res *ProjectAnalyticsConnectionKeyRotateResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if body.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if body.ConnectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	if keyID == "" {
		err = errors.New("missing required keyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/analytics/connections/%s/keys/%s/rotate", body.ProjectID, body.ConnectionID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// One key per purpose. Public player values remain available; read and MCP values
// are returned only on creation or rotation.
type AnalyticsKey struct {
	// Id identifies this key in the collector.
	ID string `json:"id" api:"required"`
	// Purpose distinguishes player ingestion, read-only API access and a separate
	// read-only MCP credential. Read and MCP keys have the same collector permissions.
	//
	// Any of "player", "read", "mcp".
	Purpose AnalyticsKeyPurpose `json:"purpose" api:"required"`
	// PublicKey is the browser ingestion credential for a player key. Omitted for
	// read-only API and MCP keys, whose secret values are shown once.
	PublicKey string `json:"public_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Purpose     respjson.Field
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsKey) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purpose distinguishes player ingestion, read-only API access and a separate
// read-only MCP credential. Read and MCP keys have the same collector permissions.
type AnalyticsKeyPurpose string

const (
	AnalyticsKeyPurposePlayer AnalyticsKeyPurpose = "player"
	AnalyticsKeyPurposeRead   AnalyticsKeyPurpose = "read"
	AnalyticsKeyPurposeMcp    AnalyticsKeyPurpose = "mcp"
)

type ProjectAnalyticsConnectionKeyNewResponse struct {
	// One key per purpose. Public player values remain available; read and MCP values
	// are returned only on creation or rotation.
	Data AnalyticsKey `json:"data" api:"required"`
	// Key is the new credential. Save read and MCP keys now; later reads omit their
	// values. Public keys remain available.
	Key    string           `json:"key" api:"required"`
	Status constant.Success `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Key         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAnalyticsConnectionKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectAnalyticsConnectionKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionKeyListResponse struct {
	Data   []AnalyticsKey   `json:"data" api:"required"`
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
func (r ProjectAnalyticsConnectionKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectAnalyticsConnectionKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionKeyRotateResponse struct {
	// One key per purpose. Public player values remain available; read and MCP values
	// are returned only on creation or rotation.
	Data AnalyticsKey `json:"data" api:"required"`
	// Key is the new credential. Save read and MCP keys now; later reads omit their
	// values. Public keys remain available.
	Key    string           `json:"key" api:"required"`
	Status constant.Success `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Key         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAnalyticsConnectionKeyRotateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectAnalyticsConnectionKeyRotateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAnalyticsConnectionKeyNewParams struct {
	ProjectID string `path:"projectId" api:"required" json:"-"`
	// Purpose distinguishes player ingestion, read-only API access and a separate
	// read-only MCP credential. Read and MCP keys have the same collector permissions.
	//
	// Any of "player", "read", "mcp".
	Purpose ProjectAnalyticsConnectionKeyNewParamsPurpose `json:"purpose,omitzero" api:"required"`
	paramObj
}

func (r ProjectAnalyticsConnectionKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectAnalyticsConnectionKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectAnalyticsConnectionKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purpose distinguishes player ingestion, read-only API access and a separate
// read-only MCP credential. Read and MCP keys have the same collector permissions.
type ProjectAnalyticsConnectionKeyNewParamsPurpose string

const (
	ProjectAnalyticsConnectionKeyNewParamsPurposePlayer ProjectAnalyticsConnectionKeyNewParamsPurpose = "player"
	ProjectAnalyticsConnectionKeyNewParamsPurposeRead   ProjectAnalyticsConnectionKeyNewParamsPurpose = "read"
	ProjectAnalyticsConnectionKeyNewParamsPurposeMcp    ProjectAnalyticsConnectionKeyNewParamsPurpose = "mcp"
)

type ProjectAnalyticsConnectionKeyListParams struct {
	ProjectID string `path:"projectId" api:"required" json:"-"`
	paramObj
}

type ProjectAnalyticsConnectionKeyRotateParams struct {
	ProjectID    string `path:"projectId" api:"required" json:"-"`
	ConnectionID string `path:"connectionId" api:"required" json:"-"`
	paramObj
}
