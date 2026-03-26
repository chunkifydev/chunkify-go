// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"encoding/json"
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
func (r *StorageService) New(ctx context.Context, body StorageNewParams, opts ...option.RequestOption) (res *StorageUnion, err error) {
	var env StorageNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "api/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve details of a specific storage configuration by its id.
func (r *StorageService) Get(ctx context.Context, storageID string, opts ...option.RequestOption) (res *StorageUnion, err error) {
	var env StorageGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if storageID == "" {
		err = errors.New("missing required storageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/storages/%s", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve a list of all storage configurations for the current project.
func (r *StorageService) List(ctx context.Context, opts ...option.RequestOption) (res *StorageListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "api/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a storage configuration. The storage must not be currently attached to
// the project.
func (r *StorageService) Delete(ctx context.Context, storageID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if storageID == "" {
		err = errors.New("missing required storageId parameter")
		return err
	}
	path := fmt.Sprintf("api/storages/%s", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// StorageUnion contains all possible properties and values from [StorageChunkify],
// [StorageCloudflare], [StorageAws].
//
// Use the [StorageUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type StorageUnion struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// Any of "chunkify", "cloudflare", "aws".
	Provider string `json:"provider"`
	Region   string `json:"region"`
	Slug     string `json:"slug"`
	Bucket   string `json:"bucket"`
	// This field is from variant [StorageCloudflare].
	Endpoint string `json:"endpoint"`
	// This field is from variant [StorageCloudflare].
	Location string `json:"location"`
	Public   bool   `json:"public"`
	JSON     struct {
		ID        respjson.Field
		CreatedAt respjson.Field
		Provider  respjson.Field
		Region    respjson.Field
		Slug      respjson.Field
		Bucket    respjson.Field
		Endpoint  respjson.Field
		Location  respjson.Field
		Public    respjson.Field
		raw       string
	} `json:"-"`
}

// anyStorage is implemented by each variant of [StorageUnion] to add type safety
// for the return type of [StorageUnion.AsAny]
type anyStorage interface {
	implStorageUnion()
}

func (StorageChunkify) implStorageUnion()   {}
func (StorageCloudflare) implStorageUnion() {}
func (StorageAws) implStorageUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := StorageUnion.AsAny().(type) {
//	case chunkify.StorageChunkify:
//	case chunkify.StorageCloudflare:
//	case chunkify.StorageAws:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u StorageUnion) AsAny() anyStorage {
	switch u.Provider {
	case "chunkify":
		return u.AsChunkify()
	case "cloudflare":
		return u.AsCloudflare()
	case "aws":
		return u.AsAws()
	}
	return nil
}

func (u StorageUnion) AsChunkify() (v StorageChunkify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StorageUnion) AsCloudflare() (v StorageCloudflare) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StorageUnion) AsAws() (v StorageAws) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StorageUnion) RawJSON() string { return u.JSON.raw }

func (r *StorageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageChunkify struct {
	// Unique identifier of the storage configuration
	ID string `json:"id" api:"required"`
	// Created at timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Provider specifies the storage provider.
	Provider constant.Chunkify `json:"provider" default:"chunkify"`
	// Region specifies the region of the storage provider.
	//
	// Any of "us-east-1", "us-east-2", "us-west-1", "us-west-2", "eu-west-1",
	// "eu-west-2", "ap-northeast-1", "ap-southeast-1".
	Region string `json:"region" api:"required"`
	// Unique identifier of the storage configuration
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Provider    respjson.Field
		Region      respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageChunkify) RawJSON() string { return r.JSON.raw }
func (r *StorageChunkify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageCloudflare struct {
	// Unique identifier of the storage configuration
	ID string `json:"id" api:"required"`
	// Bucket is the name of the storage bucket.
	Bucket string `json:"bucket" api:"required"`
	// Created at timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Endpoint is the endpoint of the storage provider.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Location specifies the location of the storage provider.
	//
	// Any of "US", "EU", "ASIA".
	Location string `json:"location" api:"required"`
	// Provider specifies the storage provider.
	Provider constant.Cloudflare `json:"provider" default:"cloudflare"`
	// Public indicates whether the storage is publicly accessible.
	Public bool `json:"public" api:"required"`
	// Region specifies the region of the storage provider.
	Region constant.Auto `json:"region" default:"auto"`
	// Unique identifier of the storage configuration
	Slug string `json:"slug" api:"required"`
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
func (r StorageCloudflare) RawJSON() string { return r.JSON.raw }
func (r *StorageCloudflare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageAws struct {
	// Unique identifier of the storage configuration
	ID string `json:"id" api:"required"`
	// Bucket is the name of the storage bucket.
	Bucket string `json:"bucket" api:"required"`
	// Created at timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Provider specifies the storage provider.
	Provider constant.Aws `json:"provider" default:"aws"`
	// Public indicates whether the storage is publicly accessible.
	Public bool `json:"public" api:"required"`
	// Region specifies the region of the storage provider.
	//
	// Any of "us-east-1", "us-east-2", "us-central-1", "us-west-1", "us-west-2",
	// "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-north-1",
	// "ap-east-1", "ap-east-2", "ap-northeast-1", "ap-northeast-2", "ap-south-1",
	// "ap-southeast-1", "ap-southeast-2".
	Region string `json:"region" api:"required"`
	// Unique identifier of the storage configuration
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bucket      respjson.Field
		CreatedAt   respjson.Field
		Provider    respjson.Field
		Public      respjson.Field
		Region      respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageAws) RawJSON() string { return r.JSON.raw }
func (r *StorageAws) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing the list of storages configurations for a project
type StorageListResponse struct {
	// Data contains the storage items
	Data []StorageUnion `json:"data" api:"required"`
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
	OfAws *StorageNewParamsStorageAws `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Storage
	// parameters for Chunkify ephemeral storage.
	OfChunkify *StorageNewParamsStorageChunkify `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Storage
	// parameters for Cloudflare R2 storage.
	OfCloudflare *StorageNewParamsStorageCloudflare `json:",inline"`

	paramObj
}

func (u StorageNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAws, u.OfChunkify, u.OfCloudflare)
}
func (r *StorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage parameters for AWS S3 storage.
//
// The properties AccessKeyID, Bucket, Provider, Region, SecretAccessKey are
// required.
type StorageNewParamsStorageAws struct {
	// AccessKeyId is the access key for the storage provider. Required if not using
	// Chunkify storage.
	AccessKeyID string `json:"access_key_id" api:"required"`
	// Bucket is the name of the storage bucket.
	Bucket string `json:"bucket" api:"required"`
	// Region specifies the region of the storage provider.
	//
	// Any of "us-east-1", "us-east-2", "us-central-1", "us-west-1", "us-west-2",
	// "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-north-1",
	// "ap-east-1", "ap-east-2", "ap-northeast-1", "ap-northeast-2", "ap-south-1",
	// "ap-southeast-1", "ap-southeast-2".
	Region string `json:"region,omitzero" api:"required"`
	// SecretAccessKey is the secret key for the storage provider. Required if not
	// using Chunkify storage.
	SecretAccessKey string `json:"secret_access_key" api:"required"`
	// Public indicates whether the storage is publicly accessible.
	Public param.Opt[bool] `json:"public,omitzero"`
	// Provider specifies the storage provider.
	//
	// This field can be elided, and will marshal its zero value as "aws".
	Provider constant.Aws `json:"provider" default:"aws"`
	paramObj
}

func (r StorageNewParamsStorageAws) MarshalJSON() (data []byte, err error) {
	type shadow StorageNewParamsStorageAws
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageNewParamsStorageAws) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StorageNewParamsStorageAws](
		"region", "us-east-1", "us-east-2", "us-central-1", "us-west-1", "us-west-2", "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "eu-north-1", "ap-east-1", "ap-east-2", "ap-northeast-1", "ap-northeast-2", "ap-south-1", "ap-southeast-1", "ap-southeast-2",
	)
}

// Storage parameters for Chunkify ephemeral storage.
//
// The properties Provider, Region are required.
type StorageNewParamsStorageChunkify struct {
	// Region specifies the region of the storage provider.
	//
	// Any of "us-east-1", "us-east-2", "us-west-1", "us-west-2", "eu-west-1",
	// "eu-west-2", "ap-northeast-1", "ap-southeast-1".
	Region string `json:"region,omitzero" api:"required"`
	// Provider specifies the storage provider.
	//
	// This field can be elided, and will marshal its zero value as "chunkify".
	Provider constant.Chunkify `json:"provider" default:"chunkify"`
	paramObj
}

func (r StorageNewParamsStorageChunkify) MarshalJSON() (data []byte, err error) {
	type shadow StorageNewParamsStorageChunkify
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageNewParamsStorageChunkify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StorageNewParamsStorageChunkify](
		"region", "us-east-1", "us-east-2", "us-west-1", "us-west-2", "eu-west-1", "eu-west-2", "ap-northeast-1", "ap-southeast-1",
	)
}

// Storage parameters for Cloudflare R2 storage.
//
// The properties AccessKeyID, Bucket, Endpoint, Location, Provider, Region,
// SecretAccessKey are required.
type StorageNewParamsStorageCloudflare struct {
	// AccessKeyId is the access key for the storage provider.
	AccessKeyID string `json:"access_key_id" api:"required"`
	// Bucket is the name of the storage bucket.
	Bucket string `json:"bucket" api:"required"`
	// Endpoint is the endpoint of the storage provider.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Location specifies the location of the storage provider.
	//
	// Any of "US", "EU", "ASIA".
	Location string `json:"location,omitzero" api:"required"`
	// SecretAccessKey is the secret key for the storage provider.
	SecretAccessKey string `json:"secret_access_key" api:"required"`
	// Public indicates whether the storage is publicly accessible.
	Public param.Opt[bool] `json:"public,omitzero"`
	// Provider specifies the storage provider.
	//
	// This field can be elided, and will marshal its zero value as "cloudflare".
	Provider constant.Cloudflare `json:"provider" default:"cloudflare"`
	// Region must be set to 'auto'.
	//
	// This field can be elided, and will marshal its zero value as "auto".
	Region constant.Auto `json:"region" default:"auto"`
	paramObj
}

func (r StorageNewParamsStorageCloudflare) MarshalJSON() (data []byte, err error) {
	type shadow StorageNewParamsStorageCloudflare
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageNewParamsStorageCloudflare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StorageNewParamsStorageCloudflare](
		"location", "US", "EU", "ASIA",
	)
}

type StorageNewResponseEnvelope struct {
	// Data contains the response object
	Data StorageUnion `json:"data" api:"required"`
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
func (r StorageNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *StorageNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageGetResponseEnvelope struct {
	// Data contains the response object
	Data StorageUnion `json:"data" api:"required"`
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
func (r StorageGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *StorageGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
