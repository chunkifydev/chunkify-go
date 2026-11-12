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
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

// WebhookService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Create a new webhook for a project. The webhook will receive notifications for
// specified events.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *Webhook, err error) {
	var env WebhookNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve details of a specific webhook configuration by its ID. The webhook must
// belong to the current project.
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *Webhook, err error) {
	var env WebhookGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update the enabled status of a webhook. The webhook must belong to the current
// project.
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Retrieve a list of all webhooks configured for the current project. Each webhook
// includes its URL, enabled status, and subscribed events.
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Permanently delete a webhook configuration. The webhook must belong to the
// current project. This action cannot be undone.
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEvent, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookKey
	if key == "" {
		return nil, errors.New("The WebhookKey option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &UnwrapWebhookEvent{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type Webhook struct {
	// Unique identifier of the webhook
	ID string `json:"id" api:"required"`
	// Whether the webhook is currently enabled
	Enabled bool `json:"enabled" api:"required"`
	// Array of event types this webhook subscribes to
	//
	// Any of "job.created", "job.completed", "job.failed", "job.cancelled",
	// "source.created", "source.deleted", "upload.created", "upload.completed",
	// "upload.failed", "upload.expired", "asset.created", "asset.ready",
	// "asset.playback_ready", "asset.failed", "asset.expired", "asset.deleted",
	// "asset.output.created", "asset.output.ready", "asset.output.failed",
	// "asset.output.deleted", "workflow_run.created", "workflow_run.completed",
	// "workflow_run.failed", "workflow_run.cancelled", "transcoding.created",
	// "transcoding.ready", "transcoding.failed", "transcoding.cancelled",
	// "transcoding.expired", "transcoding.deleted", "thumbnail.created",
	// "thumbnail.ready", "thumbnail.failed", "thumbnail.expired", "thumbnail.deleted".
	Events []string `json:"events" api:"required"`
	// ID of the project this webhook belongs to
	ProjectID string `json:"project_id" api:"required"`
	// URL where webhook events will be sent
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		ProjectID   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing the list of all webhooks for a project
type WebhookListResponse struct {
	// Data contains the webhook items
	Data []Webhook `json:"data" api:"required"`
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
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewEventWebhookEvent struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventDataUnion `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.created", "job.completed", "job.failed", "job.cancelled",
	// "source.created", "source.deleted", "upload.created", "upload.completed",
	// "upload.failed", "upload.expired", "asset.created", "asset.ready",
	// "asset.playback_ready", "asset.failed", "asset.expired", "asset.deleted",
	// "asset.output.created", "asset.output.ready", "asset.output.failed",
	// "asset.output.deleted", "workflow_run.created", "workflow_run.completed",
	// "workflow_run.failed", "workflow_run.cancelled", "transcoding.created",
	// "transcoding.ready", "transcoding.failed", "transcoding.cancelled",
	// "transcoding.expired", "transcoding.deleted", "thumbnail.created",
	// "thumbnail.ready", "thumbnail.failed", "thumbnail.expired", "thumbnail.deleted".
	Event NewEventWebhookEventEvent `json:"event" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		Date        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NewEventWebhookEventDataUnion contains all possible properties and values from
// [NewEventWebhookEventDataNotificationPayloadJobCompleted],
// [NewEventWebhookEventDataNotificationPayloadJobFailed],
// [NewEventWebhookEventDataNotificationPayloadJobCreated],
// [NewEventWebhookEventDataNotificationPayloadSourceCreated],
// [NewEventWebhookEventDataNotificationPayloadUploadCreated],
// [NewEventWebhookEventDataNotificationPayloadUploadCompleted],
// [NewEventWebhookEventDataNotificationPayloadUploadFailed],
// [NewEventWebhookEventDataNotificationPayloadAsset],
// [NewEventWebhookEventDataNotificationPayloadAssetOutput],
// [NewEventWebhookEventDataNotificationPayloadAssetOutputDeleted],
// [NewEventWebhookEventDataNotificationPayloadAssetDeleted],
// [NewEventWebhookEventDataNotificationPayloadWorkflowRun],
// [NewEventWebhookEventDataNotificationPayloadTranscoding],
// [NewEventWebhookEventDataThumbnailEventData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type NewEventWebhookEventDataUnion struct {
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadJobCompleted].
	Files []JobFile `json:"files"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadJobCompleted].
	Job Job `json:"job"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadSourceCreated].
	Source Source `json:"source"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadUploadCreated].
	Upload Upload `json:"upload"`
	// This field is from variant [NewEventWebhookEventDataNotificationPayloadAsset].
	Asset Asset `json:"asset"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadAssetOutput].
	Output       AssetOutputUnion `json:"output"`
	FilesDeleted bool             `json:"files_deleted"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadAssetDeleted].
	DeletedSize int64 `json:"deleted_size"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadWorkflowRun].
	ActionRuns []AssetActionRun `json:"action_runs"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadWorkflowRun].
	WorkflowRun WorkflowRun `json:"workflow_run"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadTranscoding].
	Transcoding Transcoding `json:"transcoding"`
	// This field is from variant [NewEventWebhookEventDataThumbnailEventData].
	Thumbnail Thumbnail `json:"thumbnail"`
	JSON      struct {
		Files        respjson.Field
		Job          respjson.Field
		Source       respjson.Field
		Upload       respjson.Field
		Asset        respjson.Field
		Output       respjson.Field
		FilesDeleted respjson.Field
		DeletedSize  respjson.Field
		ActionRuns   respjson.Field
		WorkflowRun  respjson.Field
		Transcoding  respjson.Field
		Thumbnail    respjson.Field
		raw          string
	} `json:"-"`
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadJobCompleted() (v NewEventWebhookEventDataNotificationPayloadJobCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadJobFailed() (v NewEventWebhookEventDataNotificationPayloadJobFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadJobCreated() (v NewEventWebhookEventDataNotificationPayloadJobCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadSourceCreated() (v NewEventWebhookEventDataNotificationPayloadSourceCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadUploadCreated() (v NewEventWebhookEventDataNotificationPayloadUploadCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadUploadCompleted() (v NewEventWebhookEventDataNotificationPayloadUploadCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadUploadFailed() (v NewEventWebhookEventDataNotificationPayloadUploadFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadAsset() (v NewEventWebhookEventDataNotificationPayloadAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadAssetOutput() (v NewEventWebhookEventDataNotificationPayloadAssetOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadAssetOutputDeleted() (v NewEventWebhookEventDataNotificationPayloadAssetOutputDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadAssetDeleted() (v NewEventWebhookEventDataNotificationPayloadAssetDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadWorkflowRun() (v NewEventWebhookEventDataNotificationPayloadWorkflowRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadTranscoding() (v NewEventWebhookEventDataNotificationPayloadTranscoding) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataThumbnailEventData() (v NewEventWebhookEventDataThumbnailEventData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NewEventWebhookEventDataUnion) RawJSON() string { return u.JSON.raw }

func (r *NewEventWebhookEventDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for job.completed events
type NewEventWebhookEventDataNotificationPayloadJobCompleted struct {
	// List of files generated by the job
	Files []JobFile `json:"files" api:"required"`
	Job   Job       `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadJobCompleted) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadJobCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for job.failed and job.cancelled events
type NewEventWebhookEventDataNotificationPayloadJobFailed struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadJobFailed) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadJobFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for job.created. A standalone Job has been accepted for processing;
// execution may not have started.
type NewEventWebhookEventDataNotificationPayloadJobCreated struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadJobCreated) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadJobCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for source.created. The Source has been validated and saved. Thumbnails
// may still be pending.
type NewEventWebhookEventDataNotificationPayloadSourceCreated struct {
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadSourceCreated) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadSourceCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for upload.created. A waiting upload session has been created; its
// Source is created only after successful completion. Transfer credentials are not
// included.
type NewEventWebhookEventDataNotificationPayloadUploadCreated struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadUploadCreated) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadUploadCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for upload.completed events
type NewEventWebhookEventDataNotificationPayloadUploadCompleted struct {
	Source Source `json:"source" api:"required"`
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadUploadCompleted) RawJSON() string {
	return r.JSON.raw
}
func (r *NewEventWebhookEventDataNotificationPayloadUploadCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for upload.failed and upload.expired events
type NewEventWebhookEventDataNotificationPayloadUploadFailed struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadUploadFailed) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadUploadFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewEventWebhookEventDataNotificationPayloadAsset struct {
	Asset Asset `json:"asset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadAsset) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for Asset output creation, readiness, and failure events.
type NewEventWebhookEventDataNotificationPayloadAssetOutput struct {
	// Output contains the canonical typed output snapshot at the event time, matching
	// the output API.
	Output AssetOutputUnion `json:"output" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadAssetOutput) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadAssetOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for asset.output.deleted. The output snapshot is preserved after
// removal.
type NewEventWebhookEventDataNotificationPayloadAssetOutputDeleted struct {
	// FilesDeleted indicates whether the requested cleanup of the output files
	// completed without error. Missing files count as successful cleanup. False means
	// these files were not removed at event time. Private transcript text and saved
	// subtitle copies are always scheduled for asynchronous cleanup.
	FilesDeleted bool `json:"files_deleted" api:"required"`
	// Output contains the canonical typed output snapshot before removal.
	Output AssetOutputUnion `json:"output" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilesDeleted respjson.Field
		Output       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadAssetOutputDeleted) RawJSON() string {
	return r.JSON.raw
}
func (r *NewEventWebhookEventDataNotificationPayloadAssetOutputDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for asset.deleted. The Asset snapshot is preserved after record
// deletion.
type NewEventWebhookEventDataNotificationPayloadAssetDeleted struct {
	Asset Asset `json:"asset" api:"required"`
	// FilesDeleted indicates whether the requested cleanup of the Asset's generated
	// files completed without error. Missing files count as successful cleanup. False
	// means these files were not removed at event time. Private transcript text and
	// saved subtitle copies are always scheduled for asynchronous cleanup.
	FilesDeleted bool `json:"files_deleted" api:"required"`
	// DeletedSize specifies the recorded generated-file bytes removed with the Asset,
	// including attached outputs. Omitted when cleanup was not requested or the total
	// is unavailable. Missing storage objects still contribute their recorded sizes.
	DeletedSize int64 `json:"deleted_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset        respjson.Field
		FilesDeleted respjson.Field
		DeletedSize  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadAssetDeleted) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadAssetDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for workflow_run.created, workflow_run.completed, workflow_run.failed,
// and workflow_run.cancelled. The run and its actions are snapshots taken when the
// event was recorded.
type NewEventWebhookEventDataNotificationPayloadWorkflowRun struct {
	// ActionRuns specifies every action admitted by this run, including its rule and
	// action IDs, type, input, status, output reference, and any skip or failure
	// reason. Actions from other runs are excluded.
	ActionRuns []AssetActionRun `json:"action_runs" api:"required"`
	// Snapshot of the Workflow run when the event was recorded.
	WorkflowRun WorkflowRun `json:"workflow_run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionRuns  respjson.Field
		WorkflowRun respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadWorkflowRun) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadWorkflowRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for transcoding.created, transcoding.ready, transcoding.failed,
// transcoding.cancelled, transcoding.expired and transcoding.deleted events.
type NewEventWebhookEventDataNotificationPayloadTranscoding struct {
	// Transcoding contains the canonical resource snapshot at the event time.
	Transcoding Transcoding `json:"transcoding" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transcoding respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataNotificationPayloadTranscoding) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataNotificationPayloadTranscoding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A canonical thumbnail snapshot for standalone thumbnail.\* events. Asset-owned
// thumbnails use the Asset output event payload.
type NewEventWebhookEventDataThumbnailEventData struct {
	// Thumbnail contains the resource at the event time.
	Thumbnail Thumbnail `json:"thumbnail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Thumbnail   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventDataThumbnailEventData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventDataThumbnailEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NewEventWebhookEventEvent string

const (
	NewEventWebhookEventEventJobCreated           NewEventWebhookEventEvent = "job.created"
	NewEventWebhookEventEventJobCompleted         NewEventWebhookEventEvent = "job.completed"
	NewEventWebhookEventEventJobFailed            NewEventWebhookEventEvent = "job.failed"
	NewEventWebhookEventEventJobCancelled         NewEventWebhookEventEvent = "job.cancelled"
	NewEventWebhookEventEventSourceCreated        NewEventWebhookEventEvent = "source.created"
	NewEventWebhookEventEventSourceDeleted        NewEventWebhookEventEvent = "source.deleted"
	NewEventWebhookEventEventUploadCreated        NewEventWebhookEventEvent = "upload.created"
	NewEventWebhookEventEventUploadCompleted      NewEventWebhookEventEvent = "upload.completed"
	NewEventWebhookEventEventUploadFailed         NewEventWebhookEventEvent = "upload.failed"
	NewEventWebhookEventEventUploadExpired        NewEventWebhookEventEvent = "upload.expired"
	NewEventWebhookEventEventAssetCreated         NewEventWebhookEventEvent = "asset.created"
	NewEventWebhookEventEventAssetReady           NewEventWebhookEventEvent = "asset.ready"
	NewEventWebhookEventEventAssetPlaybackReady   NewEventWebhookEventEvent = "asset.playback_ready"
	NewEventWebhookEventEventAssetFailed          NewEventWebhookEventEvent = "asset.failed"
	NewEventWebhookEventEventAssetExpired         NewEventWebhookEventEvent = "asset.expired"
	NewEventWebhookEventEventAssetDeleted         NewEventWebhookEventEvent = "asset.deleted"
	NewEventWebhookEventEventAssetOutputCreated   NewEventWebhookEventEvent = "asset.output.created"
	NewEventWebhookEventEventAssetOutputReady     NewEventWebhookEventEvent = "asset.output.ready"
	NewEventWebhookEventEventAssetOutputFailed    NewEventWebhookEventEvent = "asset.output.failed"
	NewEventWebhookEventEventAssetOutputDeleted   NewEventWebhookEventEvent = "asset.output.deleted"
	NewEventWebhookEventEventWorkflowRunCreated   NewEventWebhookEventEvent = "workflow_run.created"
	NewEventWebhookEventEventWorkflowRunCompleted NewEventWebhookEventEvent = "workflow_run.completed"
	NewEventWebhookEventEventWorkflowRunFailed    NewEventWebhookEventEvent = "workflow_run.failed"
	NewEventWebhookEventEventWorkflowRunCancelled NewEventWebhookEventEvent = "workflow_run.cancelled"
	NewEventWebhookEventEventTranscodingCreated   NewEventWebhookEventEvent = "transcoding.created"
	NewEventWebhookEventEventTranscodingReady     NewEventWebhookEventEvent = "transcoding.ready"
	NewEventWebhookEventEventTranscodingFailed    NewEventWebhookEventEvent = "transcoding.failed"
	NewEventWebhookEventEventTranscodingCancelled NewEventWebhookEventEvent = "transcoding.cancelled"
	NewEventWebhookEventEventTranscodingExpired   NewEventWebhookEventEvent = "transcoding.expired"
	NewEventWebhookEventEventTranscodingDeleted   NewEventWebhookEventEvent = "transcoding.deleted"
	NewEventWebhookEventEventThumbnailCreated     NewEventWebhookEventEvent = "thumbnail.created"
	NewEventWebhookEventEventThumbnailReady       NewEventWebhookEventEvent = "thumbnail.ready"
	NewEventWebhookEventEventThumbnailFailed      NewEventWebhookEventEvent = "thumbnail.failed"
	NewEventWebhookEventEventThumbnailExpired     NewEventWebhookEventEvent = "thumbnail.expired"
	NewEventWebhookEventEventThumbnailDeleted     NewEventWebhookEventEvent = "thumbnail.deleted"
)

type UnwrapWebhookEvent struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventDataUnion `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.created", "job.completed", "job.failed", "job.cancelled",
	// "source.created", "source.deleted", "upload.created", "upload.completed",
	// "upload.failed", "upload.expired", "asset.created", "asset.ready",
	// "asset.playback_ready", "asset.failed", "asset.expired", "asset.deleted",
	// "asset.output.created", "asset.output.ready", "asset.output.failed",
	// "asset.output.deleted", "workflow_run.created", "workflow_run.completed",
	// "workflow_run.failed", "workflow_run.cancelled", "transcoding.created",
	// "transcoding.ready", "transcoding.failed", "transcoding.cancelled",
	// "transcoding.expired", "transcoding.deleted", "thumbnail.created",
	// "thumbnail.ready", "thumbnail.failed", "thumbnail.expired", "thumbnail.deleted".
	Event UnwrapWebhookEventEvent `json:"event" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		Date        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventDataUnion contains all possible properties and values from
// [UnwrapWebhookEventDataNotificationPayloadJobCompleted],
// [UnwrapWebhookEventDataNotificationPayloadJobFailed],
// [UnwrapWebhookEventDataNotificationPayloadJobCreated],
// [UnwrapWebhookEventDataNotificationPayloadSourceCreated],
// [UnwrapWebhookEventDataNotificationPayloadUploadCreated],
// [UnwrapWebhookEventDataNotificationPayloadUploadCompleted],
// [UnwrapWebhookEventDataNotificationPayloadUploadFailed],
// [UnwrapWebhookEventDataNotificationPayloadAsset],
// [UnwrapWebhookEventDataNotificationPayloadAssetOutput],
// [UnwrapWebhookEventDataNotificationPayloadAssetOutputDeleted],
// [UnwrapWebhookEventDataNotificationPayloadAssetDeleted],
// [UnwrapWebhookEventDataNotificationPayloadWorkflowRun],
// [UnwrapWebhookEventDataNotificationPayloadTranscoding],
// [UnwrapWebhookEventDataThumbnailEventData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventDataUnion struct {
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadJobCompleted].
	Files []JobFile `json:"files"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadJobCompleted].
	Job Job `json:"job"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadSourceCreated].
	Source Source `json:"source"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadUploadCreated].
	Upload Upload `json:"upload"`
	// This field is from variant [UnwrapWebhookEventDataNotificationPayloadAsset].
	Asset Asset `json:"asset"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadAssetOutput].
	Output       AssetOutputUnion `json:"output"`
	FilesDeleted bool             `json:"files_deleted"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadAssetDeleted].
	DeletedSize int64 `json:"deleted_size"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadWorkflowRun].
	ActionRuns []AssetActionRun `json:"action_runs"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadWorkflowRun].
	WorkflowRun WorkflowRun `json:"workflow_run"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadTranscoding].
	Transcoding Transcoding `json:"transcoding"`
	// This field is from variant [UnwrapWebhookEventDataThumbnailEventData].
	Thumbnail Thumbnail `json:"thumbnail"`
	JSON      struct {
		Files        respjson.Field
		Job          respjson.Field
		Source       respjson.Field
		Upload       respjson.Field
		Asset        respjson.Field
		Output       respjson.Field
		FilesDeleted respjson.Field
		DeletedSize  respjson.Field
		ActionRuns   respjson.Field
		WorkflowRun  respjson.Field
		Transcoding  respjson.Field
		Thumbnail    respjson.Field
		raw          string
	} `json:"-"`
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadJobCompleted() (v UnwrapWebhookEventDataNotificationPayloadJobCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadJobFailed() (v UnwrapWebhookEventDataNotificationPayloadJobFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadJobCreated() (v UnwrapWebhookEventDataNotificationPayloadJobCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadSourceCreated() (v UnwrapWebhookEventDataNotificationPayloadSourceCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadUploadCreated() (v UnwrapWebhookEventDataNotificationPayloadUploadCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadUploadCompleted() (v UnwrapWebhookEventDataNotificationPayloadUploadCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadUploadFailed() (v UnwrapWebhookEventDataNotificationPayloadUploadFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadAsset() (v UnwrapWebhookEventDataNotificationPayloadAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadAssetOutput() (v UnwrapWebhookEventDataNotificationPayloadAssetOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadAssetOutputDeleted() (v UnwrapWebhookEventDataNotificationPayloadAssetOutputDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadAssetDeleted() (v UnwrapWebhookEventDataNotificationPayloadAssetDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadWorkflowRun() (v UnwrapWebhookEventDataNotificationPayloadWorkflowRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadTranscoding() (v UnwrapWebhookEventDataNotificationPayloadTranscoding) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataThumbnailEventData() (v UnwrapWebhookEventDataThumbnailEventData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventDataUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for job.completed events
type UnwrapWebhookEventDataNotificationPayloadJobCompleted struct {
	// List of files generated by the job
	Files []JobFile `json:"files" api:"required"`
	Job   Job       `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadJobCompleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadJobCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for job.failed and job.cancelled events
type UnwrapWebhookEventDataNotificationPayloadJobFailed struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadJobFailed) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadJobFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for job.created. A standalone Job has been accepted for processing;
// execution may not have started.
type UnwrapWebhookEventDataNotificationPayloadJobCreated struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadJobCreated) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadJobCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for source.created. The Source has been validated and saved. Thumbnails
// may still be pending.
type UnwrapWebhookEventDataNotificationPayloadSourceCreated struct {
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadSourceCreated) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadSourceCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for upload.created. A waiting upload session has been created; its
// Source is created only after successful completion. Transfer credentials are not
// included.
type UnwrapWebhookEventDataNotificationPayloadUploadCreated struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadUploadCreated) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadUploadCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for upload.completed events
type UnwrapWebhookEventDataNotificationPayloadUploadCompleted struct {
	Source Source `json:"source" api:"required"`
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadUploadCompleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadUploadCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload data structure for upload.failed and upload.expired events
type UnwrapWebhookEventDataNotificationPayloadUploadFailed struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadUploadFailed) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadUploadFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnwrapWebhookEventDataNotificationPayloadAsset struct {
	Asset Asset `json:"asset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadAsset) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for Asset output creation, readiness, and failure events.
type UnwrapWebhookEventDataNotificationPayloadAssetOutput struct {
	// Output contains the canonical typed output snapshot at the event time, matching
	// the output API.
	Output AssetOutputUnion `json:"output" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadAssetOutput) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadAssetOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for asset.output.deleted. The output snapshot is preserved after
// removal.
type UnwrapWebhookEventDataNotificationPayloadAssetOutputDeleted struct {
	// FilesDeleted indicates whether the requested cleanup of the output files
	// completed without error. Missing files count as successful cleanup. False means
	// these files were not removed at event time. Private transcript text and saved
	// subtitle copies are always scheduled for asynchronous cleanup.
	FilesDeleted bool `json:"files_deleted" api:"required"`
	// Output contains the canonical typed output snapshot before removal.
	Output AssetOutputUnion `json:"output" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilesDeleted respjson.Field
		Output       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadAssetOutputDeleted) RawJSON() string {
	return r.JSON.raw
}
func (r *UnwrapWebhookEventDataNotificationPayloadAssetOutputDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for asset.deleted. The Asset snapshot is preserved after record
// deletion.
type UnwrapWebhookEventDataNotificationPayloadAssetDeleted struct {
	Asset Asset `json:"asset" api:"required"`
	// FilesDeleted indicates whether the requested cleanup of the Asset's generated
	// files completed without error. Missing files count as successful cleanup. False
	// means these files were not removed at event time. Private transcript text and
	// saved subtitle copies are always scheduled for asynchronous cleanup.
	FilesDeleted bool `json:"files_deleted" api:"required"`
	// DeletedSize specifies the recorded generated-file bytes removed with the Asset,
	// including attached outputs. Omitted when cleanup was not requested or the total
	// is unavailable. Missing storage objects still contribute their recorded sizes.
	DeletedSize int64 `json:"deleted_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset        respjson.Field
		FilesDeleted respjson.Field
		DeletedSize  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadAssetDeleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadAssetDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for workflow_run.created, workflow_run.completed, workflow_run.failed,
// and workflow_run.cancelled. The run and its actions are snapshots taken when the
// event was recorded.
type UnwrapWebhookEventDataNotificationPayloadWorkflowRun struct {
	// ActionRuns specifies every action admitted by this run, including its rule and
	// action IDs, type, input, status, output reference, and any skip or failure
	// reason. Actions from other runs are excluded.
	ActionRuns []AssetActionRun `json:"action_runs" api:"required"`
	// Snapshot of the Workflow run when the event was recorded.
	WorkflowRun WorkflowRun `json:"workflow_run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionRuns  respjson.Field
		WorkflowRun respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadWorkflowRun) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadWorkflowRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for transcoding.created, transcoding.ready, transcoding.failed,
// transcoding.cancelled, transcoding.expired and transcoding.deleted events.
type UnwrapWebhookEventDataNotificationPayloadTranscoding struct {
	// Transcoding contains the canonical resource snapshot at the event time.
	Transcoding Transcoding `json:"transcoding" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transcoding respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataNotificationPayloadTranscoding) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataNotificationPayloadTranscoding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A canonical thumbnail snapshot for standalone thumbnail.\* events. Asset-owned
// thumbnails use the Asset output event payload.
type UnwrapWebhookEventDataThumbnailEventData struct {
	// Thumbnail contains the resource at the event time.
	Thumbnail Thumbnail `json:"thumbnail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Thumbnail   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventDataThumbnailEventData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventDataThumbnailEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type UnwrapWebhookEventEvent string

const (
	UnwrapWebhookEventEventJobCreated           UnwrapWebhookEventEvent = "job.created"
	UnwrapWebhookEventEventJobCompleted         UnwrapWebhookEventEvent = "job.completed"
	UnwrapWebhookEventEventJobFailed            UnwrapWebhookEventEvent = "job.failed"
	UnwrapWebhookEventEventJobCancelled         UnwrapWebhookEventEvent = "job.cancelled"
	UnwrapWebhookEventEventSourceCreated        UnwrapWebhookEventEvent = "source.created"
	UnwrapWebhookEventEventSourceDeleted        UnwrapWebhookEventEvent = "source.deleted"
	UnwrapWebhookEventEventUploadCreated        UnwrapWebhookEventEvent = "upload.created"
	UnwrapWebhookEventEventUploadCompleted      UnwrapWebhookEventEvent = "upload.completed"
	UnwrapWebhookEventEventUploadFailed         UnwrapWebhookEventEvent = "upload.failed"
	UnwrapWebhookEventEventUploadExpired        UnwrapWebhookEventEvent = "upload.expired"
	UnwrapWebhookEventEventAssetCreated         UnwrapWebhookEventEvent = "asset.created"
	UnwrapWebhookEventEventAssetReady           UnwrapWebhookEventEvent = "asset.ready"
	UnwrapWebhookEventEventAssetPlaybackReady   UnwrapWebhookEventEvent = "asset.playback_ready"
	UnwrapWebhookEventEventAssetFailed          UnwrapWebhookEventEvent = "asset.failed"
	UnwrapWebhookEventEventAssetExpired         UnwrapWebhookEventEvent = "asset.expired"
	UnwrapWebhookEventEventAssetDeleted         UnwrapWebhookEventEvent = "asset.deleted"
	UnwrapWebhookEventEventAssetOutputCreated   UnwrapWebhookEventEvent = "asset.output.created"
	UnwrapWebhookEventEventAssetOutputReady     UnwrapWebhookEventEvent = "asset.output.ready"
	UnwrapWebhookEventEventAssetOutputFailed    UnwrapWebhookEventEvent = "asset.output.failed"
	UnwrapWebhookEventEventAssetOutputDeleted   UnwrapWebhookEventEvent = "asset.output.deleted"
	UnwrapWebhookEventEventWorkflowRunCreated   UnwrapWebhookEventEvent = "workflow_run.created"
	UnwrapWebhookEventEventWorkflowRunCompleted UnwrapWebhookEventEvent = "workflow_run.completed"
	UnwrapWebhookEventEventWorkflowRunFailed    UnwrapWebhookEventEvent = "workflow_run.failed"
	UnwrapWebhookEventEventWorkflowRunCancelled UnwrapWebhookEventEvent = "workflow_run.cancelled"
	UnwrapWebhookEventEventTranscodingCreated   UnwrapWebhookEventEvent = "transcoding.created"
	UnwrapWebhookEventEventTranscodingReady     UnwrapWebhookEventEvent = "transcoding.ready"
	UnwrapWebhookEventEventTranscodingFailed    UnwrapWebhookEventEvent = "transcoding.failed"
	UnwrapWebhookEventEventTranscodingCancelled UnwrapWebhookEventEvent = "transcoding.cancelled"
	UnwrapWebhookEventEventTranscodingExpired   UnwrapWebhookEventEvent = "transcoding.expired"
	UnwrapWebhookEventEventTranscodingDeleted   UnwrapWebhookEventEvent = "transcoding.deleted"
	UnwrapWebhookEventEventThumbnailCreated     UnwrapWebhookEventEvent = "thumbnail.created"
	UnwrapWebhookEventEventThumbnailReady       UnwrapWebhookEventEvent = "thumbnail.ready"
	UnwrapWebhookEventEventThumbnailFailed      UnwrapWebhookEventEvent = "thumbnail.failed"
	UnwrapWebhookEventEventThumbnailExpired     UnwrapWebhookEventEvent = "thumbnail.expired"
	UnwrapWebhookEventEventThumbnailDeleted     UnwrapWebhookEventEvent = "thumbnail.deleted"
)

type WebhookNewParams struct {
	// Url is the endpoint that will receive webhook notifications, which must be a
	// valid HTTP URL.
	URL string `json:"url" api:"required"`
	// Enabled indicates whether the webhook is active.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Events specifies the types of events that will trigger the webhook.
	//
	// Any of "job.created", "job.completed", "job.failed", "job.cancelled",
	// "source.created", "source.deleted", "upload.created", "upload.completed",
	// "upload.failed", "upload.expired", "asset.created", "asset.ready",
	// "asset.playback_ready", "asset.failed", "asset.expired", "asset.deleted",
	// "asset.output.created", "asset.output.ready", "asset.output.failed",
	// "asset.output.deleted", "workflow_run.created", "workflow_run.completed",
	// "workflow_run.failed", "workflow_run.cancelled", "transcoding.created",
	// "transcoding.ready", "transcoding.failed", "transcoding.cancelled",
	// "transcoding.expired", "transcoding.deleted", "thumbnail.created",
	// "thumbnail.ready", "thumbnail.failed", "thumbnail.expired", "thumbnail.deleted".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseEnvelope struct {
	// Data contains the response object
	Data Webhook `json:"data" api:"required"`
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
func (r WebhookNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetResponseEnvelope struct {
	// Data contains the response object
	Data Webhook `json:"data" api:"required"`
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
func (r WebhookGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	// Enabled indicates whether the webhook should be enabled or disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Events specifies the types of events that will trigger the webhook.
	//
	// Any of "job.created", "job.completed", "job.failed", "job.cancelled",
	// "source.created", "source.deleted", "upload.created", "upload.completed",
	// "upload.failed", "upload.expired", "asset.created", "asset.ready",
	// "asset.playback_ready", "asset.failed", "asset.expired", "asset.deleted",
	// "asset.output.created", "asset.output.ready", "asset.output.failed",
	// "asset.output.deleted", "workflow_run.created", "workflow_run.completed",
	// "workflow_run.failed", "workflow_run.cancelled", "transcoding.created",
	// "transcoding.ready", "transcoding.failed", "transcoding.cancelled",
	// "transcoding.expired", "transcoding.deleted", "thumbnail.created",
	// "thumbnail.ready", "thumbnail.failed", "thumbnail.expired", "thumbnail.deleted".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
