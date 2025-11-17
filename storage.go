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
func (r *StorageService) Delete(ctx context.Context, storageID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if storageID == "" {
		err = errors.New("missing required storageId parameter")
		return
	}
	path := fmt.Sprintf("api/storages/%s", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Storage struct {
	// Unique identifier of the storage configuration
	ID string `json:"id,required"`
	// Name of the storage bucket
	Bucket string `json:"bucket,required"`
	// Created at timestamp
	CreatedAt string `json:"created_at,required"`
	// Endpoint of the storage provider
	Endpoint string `json:"endpoint,required"`
	// Continent location of the storage (eg. US, EU, ASIA)
	Location string `json:"location,required"`
	// Name of the storage provider (e.g. AWS, GCP)
	Provider string `json:"provider,required"`
	// Whether the storage bucket is publicly accessible
	Public bool `json:"public,required"`
	// Geographic region where the storage is located
	Region string `json:"region,required"`
	// Unique identifier of the storage configuration
	Slug string `json:"slug,required"`
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
	shared.ResponseOk
}

// Returns the unmodified JSON received from the API
func (r StorageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Successful response
type StorageGetResponse struct {
	Data Storage `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.ResponseOk
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
	shared.ResponseOk
}

// Returns the unmodified JSON received from the API
func (r StorageListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Storage
	// parameters for AWS S3 storage.
	OfObject *StorageNewParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Storage
	// parameters for Chunkify ephemeral storage.
	OfStorageNewsBodyObject *StorageNewParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Storage
	// parameters for Cloudflare R2 storage.
	OfVariant2 *StorageNewParamsBodyObject `json:",inline"`

	paramObj
}

func (u StorageNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject, u.OfStorageNewsBodyObject, u.OfVariant2)
}
func (r *StorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage parameters for AWS S3 storage.
//
// The properties AccessKeyID, Bucket, Provider, Region, SecretAccessKey are
// required.
type StorageNewParamsBodyObject struct {
	// AccessKeyId is the access key for the storage provider. Required if not using
	// Chunkify storage.
	AccessKeyID string `json:"access_key_id,required"`
	// Bucket is the name of the storage bucket.
	Bucket string `json:"bucket,required"`
	// Provider specifies the storage provider.
	//
	// Any of "aws".
	Provider string `json:"provider,omitzero,required"`
	// Region specifies the region of the storage provider.
	//
	// Any of "us-east-1", "us-east-2", "us-central-1", "us-west-1", "us-west-2",
	// "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-north-1",
	// "ap-east-1", "ap-east-2", "ap-northeast-1", "ap-northeast-2", "ap-south-1",
	// "ap-southeast-1", "ap-southeast-2".
	Region string `json:"region,omitzero,required"`
	// SecretAccessKey is the secret key for the storage provider. Required if not
	// using Chunkify storage.
	SecretAccessKey string `json:"secret_access_key,required"`
	// Public indicates whether the storage is publicly accessible.
	Public param.Opt[bool] `json:"public,omitzero"`
	paramObj
}

func (r StorageNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow StorageNewParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageNewParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StorageNewParamsBodyObject](
		"provider", "aws",
	)
	apijson.RegisterFieldValidator[StorageNewParamsBodyObject](
		"region", "us-east-1", "us-east-2", "us-central-1", "us-west-1", "us-west-2", "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-north-1", "ap-east-1", "ap-east-2", "ap-northeast-1", "ap-northeast-2", "ap-south-1", "ap-southeast-1", "ap-southeast-2",
	)
}
