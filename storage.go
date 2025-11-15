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
)

// StorageService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageService] method instead.
type StorageService struct {
	Options []option.RequestOption
}

// NewStorageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStorageService(opts ...option.RequestOption) (r StorageService) {
	r = StorageService{}
	r.Options = opts
	return
}

// Create a new storage configuration for cloud storage providers like AWS S3,
// Cloudflare R2, etc. The storage credentials will be validated before saving.
func (r *StorageService) New(ctx context.Context, body StorageNewParams, opts ...option.RequestOption) (res *StorageNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific storage configuration by its id.
func (r *StorageService) Get(ctx context.Context, storageID string, opts ...option.RequestOption) (res *StorageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if storageID == "" {
		err = errors.New("missing required storageId parameter")
		return
	}
	path := fmt.Sprintf("api/storages/%s", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of all storage configurations for the current project.
func (r *StorageService) List(ctx context.Context, opts ...option.RequestOption) (res *StorageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a storage configuration. The storage must not be currently attached to
// the project.
func (r *StorageService) Delete(ctx context.Context, storageID string, opts ...option.RequestOption) (res *StorageDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if storageID == "" {
		err = errors.New("missing required storageId parameter")
		return
	}
	path := fmt.Sprintf("api/storages/%s", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Storage struct {
	// Unique identifier of the storage configuration
	ID string `json:"id"`
	// Name of the storage bucket
	Bucket string `json:"bucket"`
	// Created at timestamp
	CreatedAt string `json:"created_at"`
	// Endpoint of the storage provider
	Endpoint string `json:"endpoint"`
	// Continent location of the storage (eg. US, EU, ASIA)
	Location string `json:"location"`
	// Name of the storage provider (e.g. AWS, GCP)
	Provider string `json:"provider"`
	// Whether the storage bucket is publicly accessible
	Public bool `json:"public"`
	// Geographic region where the storage is located
	Region string `json:"region"`
	// Unique identifier of the storage configuration
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bucket      respjson.Field
		CreatedAt   respjson.Field
		Endpoint    respjson.Field
		Location    respjson.Field
		Provider    respjson.Field
		Public      respjson.Field
		Region      respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Storage) RawJSON() string { return r.JSON.raw }
func (r *Storage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type StorageNewResponse struct {
	Data Storage `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r StorageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type StorageGetResponse struct {
	Data Storage `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r StorageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type StorageListResponse struct {
	Data []Storage `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ResponseOk
}

// Returns the unmodified JSON received from the API
func (r StorageListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageDeleteResponse = any

type StorageNewParams struct {
	// Provider specifies the storage provider.
	//
	// Any of "chunkify", "aws", "cloudflare".
	Provider               StorageNewParamsProvider `json:"provider,omitzero,required"`
	StorageStorageProvider any                      `json:"storage.StorageProvider,omitzero"`
	paramObj
}

func (r StorageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider specifies the storage provider.
type StorageNewParamsProvider string

const (
	StorageNewParamsProviderChunkify   StorageNewParamsProvider = "chunkify"
	StorageNewParamsProviderAws        StorageNewParamsProvider = "aws"
	StorageNewParamsProviderCloudflare StorageNewParamsProvider = "cloudflare"
)
