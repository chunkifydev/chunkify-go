// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/chunkify-go/internal/apijson"
	"github.com/stainless-sdks/chunkify-go/internal/requestconfig"
	"github.com/stainless-sdks/chunkify-go/option"
	"github.com/stainless-sdks/chunkify-go/packages/param"
	"github.com/stainless-sdks/chunkify-go/packages/respjson"
	"github.com/stainless-sdks/chunkify-go/shared"
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
func (r *TokenService) New(ctx context.Context, body TokenNewParams, opts ...option.RequestOption) (res *TokenNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a list of all API tokens for your account, including both
// account-scoped and project-scoped tokens. For each token, the response includes
// its name, scope, creation date, and usage statistics. The token values are not
// included in the response for security reasons.
func (r *TokenService) List(ctx context.Context, opts ...option.RequestOption) (res *TokenListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revoke an access token by its ID. This action is irreversible and will
// immediately invalidate the token.
func (r *TokenService) Revoke(ctx context.Context, tokenID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return
	}
	path := fmt.Sprintf("api/tokens/%s", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Token struct {
	// Unique identifier of the token
	ID string `json:"id,required"`
	// The actual token value (only returned on creation)
	Token string `json:"token,required"`
	// Name given to the token
	Name string `json:"name,required"`
	// ID of the project this token belongs to
	ProjectID string `json:"project_id,required"`
	// Access scope of the token (e.g.project, team)
	Scope string `json:"scope,required"`
	// Timestamp when the token was created
	CreatedAt string `json:"created_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Token       respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		Scope       respjson.Field
		CreatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Token) RawJSON() string { return r.JSON.raw }
func (r *Token) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type TokenNewResponse struct {
	Data Token `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.ResponseOk
}

// Returns the unmodified JSON received from the API
func (r TokenNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TokenNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type TokenListResponse struct {
	Data []Token `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.ResponseOk
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
	// Any of "team", "project".
	Scope TokenNewParamsScope `json:"scope,omitzero,required"`
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
	TokenNewParamsScopeTeam    TokenNewParamsScope = "team"
	TokenNewParamsScopeProject TokenNewParamsScope = "project"
)
