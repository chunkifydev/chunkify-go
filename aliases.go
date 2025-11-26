// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomchunkifydevchunkifygo

import (
	"github.com/chunkifydev/chunkify-go/internal/apierror"
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type ChunkifyError = shared.ChunkifyError

// Type of error
//
// This is an alias to an internal type.
type ChunkifyErrorType = shared.ChunkifyErrorType

// Equals "setup"
const ChunkifyErrorTypeSetup = shared.ChunkifyErrorTypeSetup

// Equals "ffmpeg"
const ChunkifyErrorTypeFfmpeg = shared.ChunkifyErrorTypeFfmpeg

// Equals "source"
const ChunkifyErrorTypeSource = shared.ChunkifyErrorTypeSource

// Equals "upload"
const ChunkifyErrorTypeUpload = shared.ChunkifyErrorTypeUpload

// Equals "download"
const ChunkifyErrorTypeDownload = shared.ChunkifyErrorTypeDownload

// Equals "ingest"
const ChunkifyErrorTypeIngest = shared.ChunkifyErrorTypeIngest

// Equals "job"
const ChunkifyErrorTypeJob = shared.ChunkifyErrorTypeJob

// Equals "unexpected"
const ChunkifyErrorTypeUnexpected = shared.ChunkifyErrorTypeUnexpected

// Equals "permission"
const ChunkifyErrorTypePermission = shared.ChunkifyErrorTypePermission

// Equals "timeout"
const ChunkifyErrorTypeTimeout = shared.ChunkifyErrorTypeTimeout

// Equals "cancelled"
const ChunkifyErrorTypeCancelled = shared.ChunkifyErrorTypeCancelled
