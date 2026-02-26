// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/chunkifydev/chunkify-go/internal/apijson"
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ChunkifyError struct {
	// Additional error details or output
	Detail string `json:"detail" api:"required"`
	// Main error message
	Message string `json:"message" api:"required"`
	// Type of error
	//
	// Any of "setup", "ffmpeg", "source", "upload", "download", "ingest", "job",
	// "unexpected", "permission", "timeout", "cancelled".
	Type ChunkifyErrorType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkifyError) RawJSON() string { return r.JSON.raw }
func (r *ChunkifyError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of error
type ChunkifyErrorType string

const (
	ChunkifyErrorTypeSetup      ChunkifyErrorType = "setup"
	ChunkifyErrorTypeFfmpeg     ChunkifyErrorType = "ffmpeg"
	ChunkifyErrorTypeSource     ChunkifyErrorType = "source"
	ChunkifyErrorTypeUpload     ChunkifyErrorType = "upload"
	ChunkifyErrorTypeDownload   ChunkifyErrorType = "download"
	ChunkifyErrorTypeIngest     ChunkifyErrorType = "ingest"
	ChunkifyErrorTypeJob        ChunkifyErrorType = "job"
	ChunkifyErrorTypeUnexpected ChunkifyErrorType = "unexpected"
	ChunkifyErrorTypePermission ChunkifyErrorType = "permission"
	ChunkifyErrorTypeTimeout    ChunkifyErrorType = "timeout"
	ChunkifyErrorTypeCancelled  ChunkifyErrorType = "cancelled"
)
