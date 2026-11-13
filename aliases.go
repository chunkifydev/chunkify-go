// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

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

// A customer-facing Asset output processing failure.
//
// This is an alias to an internal type.
type AssetOutputError = shared.AssetOutputError

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

// CreatedVia specifies how the resource was created: api for API requests,
// dashboard for the Chunkify dashboard, mcp for MCP clients, and workflow for
// Workflow actions. Omitted when the creator is unknown.
//
// This is an alias to an internal type.
type CreatedVia = shared.CreatedVia

// Equals "api"
const CreatedViaAPI = shared.CreatedViaAPI

// Equals "dashboard"
const CreatedViaDashboard = shared.CreatedViaDashboard

// Equals "mcp"
const CreatedViaMcp = shared.CreatedViaMcp

// Equals "workflow"
const CreatedViaWorkflow = shared.CreatedViaWorkflow

// This is an alias to an internal type.
type InstanceType = shared.InstanceType

// Equals "4vCPU"
const InstanceType4vCPU = shared.InstanceType4vCPU

// Equals "8vCPU"
const InstanceType8vCPU = shared.InstanceType8vCPU

// Equals "16vCPU"
const InstanceType16vCPU = shared.InstanceType16vCPU

// This is an alias to an internal type.
type Log = shared.Log

// Log level
//
// This is an alias to an internal type.
type LogLevel = shared.LogLevel

// Equals "info"
const LogLevelInfo = shared.LogLevelInfo

// Equals "error"
const LogLevelError = shared.LogLevelError

// Equals "debug"
const LogLevelDebug = shared.LogLevelDebug

// Name of the service that generated the log
//
// This is an alias to an internal type.
type LogService = shared.LogService

// Equals "transcoder"
const LogServiceTranscoder = shared.LogServiceTranscoder

// Equals "manager"
const LogServiceManager = shared.LogServiceManager

// This is an alias to an internal type.
type TranscoderStatus = shared.TranscoderStatus

// Current status of the transcoder (starting, transcoding, finished, failed)
//
// This is an alias to an internal type.
type TranscoderStatusStatus = shared.TranscoderStatusStatus

// Equals "starting"
const TranscoderStatusStatusStarting = shared.TranscoderStatusStatusStarting

// Equals "pending"
const TranscoderStatusStatusPending = shared.TranscoderStatusStatusPending

// Equals "transcoding"
const TranscoderStatusStatusTranscoding = shared.TranscoderStatusStatusTranscoding

// Equals "completed"
const TranscoderStatusStatusCompleted = shared.TranscoderStatusStatusCompleted

// Equals "failed"
const TranscoderStatusStatusFailed = shared.TranscoderStatusStatusFailed

// Equals "cancelled"
const TranscoderStatusStatusCancelled = shared.TranscoderStatusStatusCancelled
