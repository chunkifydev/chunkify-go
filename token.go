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

// TokenService contains methods and other services that help with interacting with
// the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r TokenService) {
	r = TokenService{}
	r.Options = opts
	return
}

// Create a new access token for either account-wide or project-specific access.
// Project tokens require a valid project slug.
func (r *TokenService) New(ctx context.Context, body TokenNewParams, opts ...option.RequestOption) (res *Token, err error) {
	var env TokenNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "api/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve a list of all API tokens for your account, including both team-scoped
// and project-scoped tokens. For each token, the response includes its name,
// scope, creation date, and usage statistics. The token values are not included in
// the response for security reasons.
func (r *TokenService) List(ctx context.Context, opts ...option.RequestOption) (res *TokenListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "api/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke an access token by its ID. This action is irreversible and will
// immediately invalidate the token.
func (r *TokenService) Revoke(ctx context.Context, tokenID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithTeamAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return err
	}
	path := fmt.Sprintf("api/tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Token struct {
	// Unique identifier of the token
	ID string `json:"id" api:"required"`
	// The actual token value (only returned on creation)
	Token string `json:"token" api:"required"`
	// Timestamp when the token was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name given to the token
	Name string `json:"name" api:"required"`
	// ID of the project this token belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Access scope of the token
	//
	// Any of "project", "team".
	Scope TokenScope `json:"scope" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Token       respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		Scope       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Token) RawJSON() string { return r.JSON.raw }
func (r *Token) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Access scope of the token
type TokenScope string

const (
	TokenScopeProject TokenScope = "project"
	TokenScopeTeam    TokenScope = "team"
)

// Response containing the list of all tokens for a team. Including project and
// team tokens.
type TokenListResponse struct {
	// Data contains the token items
	Data []Token `json:"data" api:"required"`
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
func (r TokenListResponse) RawJSON() string { return r.JSON.raw }
func (r *TokenListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenNewParams struct {
	// Scope specifies the scope of the token, which must be either "team" or
	// "project".
	//
	// Any of "project", "team".
	Scope TokenNewParamsScope `json:"scope,omitzero" api:"required"`
	// Name is the name of the token, which can be up to 64 characters long.
	Name param.Opt[string] `json:"name,omitzero"`
	// ProjectId is required if the scope is set to "project".
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	paramObj
}

func (r TokenNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TokenNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scope specifies the scope of the token, which must be either "team" or
// "project".
type TokenNewParamsScope string

const (
	TokenNewParamsScopeProject TokenNewParamsScope = "project"
	TokenNewParamsScopeTeam    TokenNewParamsScope = "team"
)

type TokenNewResponseEnvelope struct {
	// Data contains the response object
	Data Token `json:"data" api:"required"`
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
func (r TokenNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TokenNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
