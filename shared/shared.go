// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/chunkify-go/internal/apijson"
	"github.com/stainless-sdks/chunkify-go/packages/param"
	"github.com/stainless-sdks/chunkify-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ChunkifyError struct {
	// Additional error details or output
	Detail string `json:"detail"`
	// Main error message
	Message string `json:"message"`
	// Type of error (e.g., "ffmpeg", "network", "storage", etc.)
	Type string `json:"type"`
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

// Successful response
type ResponseOk struct {
	// Data contains the response object
	Data any `json:"data"`
	// Status indicates the response status "success"
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseOk) RawJSON() string { return r.JSON.raw }
func (r *ResponseOk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
