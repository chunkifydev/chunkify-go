// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/chunkifydev/chunkify-go/internal/apijson"
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A customer-facing Asset output processing failure.
type AssetOutputError struct {
	// Machine-readable processing error code.
	Code string `json:"code" api:"required"`
	// Description of the processing failure.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetOutputError) RawJSON() string { return r.JSON.raw }
func (r *AssetOutputError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// CreatedVia specifies how the resource was created: api for API requests,
// dashboard for the Chunkify dashboard, mcp for MCP clients, and workflow for
// Workflow actions. Omitted when the creator is unknown.
type CreatedVia string

const (
	CreatedViaAPI       CreatedVia = "api"
	CreatedViaDashboard CreatedVia = "dashboard"
	CreatedViaMcp       CreatedVia = "mcp"
	CreatedViaWorkflow  CreatedVia = "workflow"
)

type InstanceType string

const (
	InstanceType4vCPU  InstanceType = "4vCPU"
	InstanceType8vCPU  InstanceType = "8vCPU"
	InstanceType16vCPU InstanceType = "16vCPU"
)

type Log struct {
	// Additional structured data attached to the log
	Attributes map[string]any `json:"attributes" api:"required"`
	// Log level
	//
	// Any of "info", "error", "debug".
	Level LogLevel `json:"level" api:"required"`
	// The log message content
	Msg string `json:"msg" api:"required"`
	// Name of the service that generated the log
	//
	// Any of "transcoder", "manager".
	Service LogService `json:"service" api:"required"`
	// Timestamp when the log was created
	Time time.Time `json:"time" api:"required" format:"date-time"`
	// Optional ID of the job this log is associated with
	JobID string `json:"job_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes  respjson.Field
		Level       respjson.Field
		Msg         respjson.Field
		Service     respjson.Field
		Time        respjson.Field
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Log) RawJSON() string { return r.JSON.raw }
func (r *Log) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Log level
type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelError LogLevel = "error"
	LogLevelDebug LogLevel = "debug"
)

// Name of the service that generated the log
type LogService string

const (
	LogServiceTranscoder LogService = "transcoder"
	LogServiceManager    LogService = "manager"
)

type TranscoderStatus struct {
	// Unique identifier of the transcoder
	ID string `json:"id" api:"required"`
	// Billable time in seconds
	BillableTime int64 `json:"billable_time" api:"required"`
	// End time of the current chunk in seconds
	ChunkEndTime float64 `json:"chunk_end_time" api:"required"`
	// Number of the chunk being processed
	ChunkNumber int64 `json:"chunk_number" api:"required"`
	// Start time of the current chunk in seconds
	ChunkStartTime float64 `json:"chunk_start_time" api:"required"`
	// CPU time used for transcoding in seconds
	CPUTime float64 `json:"cpu_time" api:"required"`
	// Timestamp when the status was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Current frames per second being processed
	Fps float64 `json:"fps" api:"required"`
	// Current frame number being processed
	Frame int64 `json:"frame" api:"required"`
	// Unique identifier of the job
	JobID string `json:"job_id" api:"required"`
	// Current output time in seconds
	OutTime int64 `json:"out_time" api:"required"`
	// Progress percentage of the transcoding operation (0-100)
	Progress float64 `json:"progress" api:"required"`
	// Current processing speed multiplier
	Speed float64 `json:"speed" api:"required"`
	// Current status of the transcoder (starting, transcoding, finished, failed)
	//
	// Any of "starting", "pending", "transcoding", "completed", "failed", "cancelled".
	Status TranscoderStatusStatus `json:"status" api:"required"`
	// Unique identifier of the transcoder instance (generated by the transcoder)
	TranscoderInstanceID string `json:"transcoder_instance_id" api:"required"`
	// Timestamp when the status was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Error message if the transcoding failed
	Error ChunkifyError `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		BillableTime         respjson.Field
		ChunkEndTime         respjson.Field
		ChunkNumber          respjson.Field
		ChunkStartTime       respjson.Field
		CPUTime              respjson.Field
		CreatedAt            respjson.Field
		Fps                  respjson.Field
		Frame                respjson.Field
		JobID                respjson.Field
		OutTime              respjson.Field
		Progress             respjson.Field
		Speed                respjson.Field
		Status               respjson.Field
		TranscoderInstanceID respjson.Field
		UpdatedAt            respjson.Field
		Error                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscoderStatus) RawJSON() string { return r.JSON.raw }
func (r *TranscoderStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the transcoder (starting, transcoding, finished, failed)
type TranscoderStatusStatus string

const (
	TranscoderStatusStatusStarting    TranscoderStatusStatus = "starting"
	TranscoderStatusStatusPending     TranscoderStatusStatus = "pending"
	TranscoderStatusStatusTranscoding TranscoderStatusStatus = "transcoding"
	TranscoderStatusStatusCompleted   TranscoderStatusStatus = "completed"
	TranscoderStatusStatusFailed      TranscoderStatusStatus = "failed"
	TranscoderStatusStatusCancelled   TranscoderStatusStatus = "cancelled"
)
