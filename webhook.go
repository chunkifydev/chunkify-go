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

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
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
	res := &UnwrapWebhookEventUnion{}
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

// NewEventWebhookEventUnion contains all possible properties and values from
// [NewEventWebhookEventJobCompleted],
// [NewEventWebhookEventNotificationPayloadJobFailed],
// [NewEventWebhookEventJobCreated], [NewEventWebhookEventSourceCreated],
// [NewEventWebhookEventSourceDeleted], [NewEventWebhookEventUploadCreated],
// [NewEventWebhookEventUploadCompleted],
// [NewEventWebhookEventNotificationPayloadUploadFailed],
// [NewEventWebhookEventNotificationPayloadAsset],
// [NewEventWebhookEventNotificationPayloadAssetOutput],
// [NewEventWebhookEventAssetOutputDeleted], [NewEventWebhookEventAssetDeleted],
// [NewEventWebhookEventNotificationPayloadWorkflowRun],
// [NewEventWebhookEventNotificationPayloadTranscoding],
// [NewEventWebhookEventNotificationPayloadThumbnail].
//
// Use the [NewEventWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type NewEventWebhookEventUnion struct {
	ID string `json:"id"`
	// This field is a union of [NewEventWebhookEventJobCompletedData],
	// [NewEventWebhookEventNotificationPayloadJobFailedData],
	// [NewEventWebhookEventJobCreatedData], [NewEventWebhookEventSourceCreatedData],
	// [NewEventWebhookEventSourceDeletedData],
	// [NewEventWebhookEventUploadCreatedData],
	// [NewEventWebhookEventUploadCompletedData],
	// [NewEventWebhookEventNotificationPayloadUploadFailedData],
	// [NewEventWebhookEventNotificationPayloadAssetData],
	// [NewEventWebhookEventNotificationPayloadAssetOutputData],
	// [NewEventWebhookEventAssetOutputDeletedData],
	// [NewEventWebhookEventAssetDeletedData],
	// [NewEventWebhookEventNotificationPayloadWorkflowRunData],
	// [NewEventWebhookEventNotificationPayloadTranscodingData],
	// [NewEventWebhookEventNotificationPayloadThumbnailData]
	Data NewEventWebhookEventUnionData `json:"data"`
	Date time.Time                     `json:"date"`
	// Any of "job.completed", nil, "job.created", "source.created", "source.deleted",
	// "upload.created", "upload.completed", nil, nil, nil, "asset.output.deleted",
	// "asset.deleted", nil, nil, nil.
	Event string `json:"event"`
	JSON  struct {
		ID    respjson.Field
		Data  respjson.Field
		Date  respjson.Field
		Event respjson.Field
		raw   string
	} `json:"-"`
}

func (u NewEventWebhookEventUnion) AsJobCompleted() (v NewEventWebhookEventJobCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsNewEventWebhookEventNotificationPayloadJobFailed() (v NewEventWebhookEventNotificationPayloadJobFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsJobCreated() (v NewEventWebhookEventJobCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsSourceCreated() (v NewEventWebhookEventSourceCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsSourceDeleted() (v NewEventWebhookEventSourceDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsUploadCreated() (v NewEventWebhookEventUploadCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsUploadCompleted() (v NewEventWebhookEventUploadCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsNewEventWebhookEventNotificationPayloadUploadFailed() (v NewEventWebhookEventNotificationPayloadUploadFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsNewEventWebhookEventNotificationPayloadAsset() (v NewEventWebhookEventNotificationPayloadAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsNewEventWebhookEventNotificationPayloadAssetOutput() (v NewEventWebhookEventNotificationPayloadAssetOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsAssetOutputDeleted() (v NewEventWebhookEventAssetOutputDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsAssetDeleted() (v NewEventWebhookEventAssetDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsNewEventWebhookEventNotificationPayloadWorkflowRun() (v NewEventWebhookEventNotificationPayloadWorkflowRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsNewEventWebhookEventNotificationPayloadTranscoding() (v NewEventWebhookEventNotificationPayloadTranscoding) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventUnion) AsNewEventWebhookEventNotificationPayloadThumbnail() (v NewEventWebhookEventNotificationPayloadThumbnail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NewEventWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *NewEventWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NewEventWebhookEventUnionData is an implicit subunion of
// [NewEventWebhookEventUnion]. NewEventWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [NewEventWebhookEventUnion].
type NewEventWebhookEventUnionData struct {
	// This field is from variant [NewEventWebhookEventJobCompletedData].
	Files []JobFile `json:"files"`
	// This field is from variant [NewEventWebhookEventJobCompletedData].
	Job Job `json:"job"`
	// This field is from variant [NewEventWebhookEventSourceCreatedData].
	Source Source `json:"source"`
	// This field is from variant [NewEventWebhookEventUploadCreatedData].
	Upload Upload `json:"upload"`
	// This field is from variant [NewEventWebhookEventNotificationPayloadAssetData].
	Asset Asset `json:"asset"`
	// This field is from variant
	// [NewEventWebhookEventNotificationPayloadAssetOutputData].
	Output       AssetOutputUnion `json:"output"`
	FilesDeleted bool             `json:"files_deleted"`
	// This field is from variant [NewEventWebhookEventAssetDeletedData].
	DeletedSize int64 `json:"deleted_size"`
	// This field is from variant
	// [NewEventWebhookEventNotificationPayloadWorkflowRunData].
	ActionRuns []AssetActionRun `json:"action_runs"`
	// This field is from variant
	// [NewEventWebhookEventNotificationPayloadWorkflowRunData].
	WorkflowRun WorkflowRun `json:"workflow_run"`
	// This field is from variant
	// [NewEventWebhookEventNotificationPayloadTranscodingData].
	Transcoding Transcoding `json:"transcoding"`
	// This field is from variant
	// [NewEventWebhookEventNotificationPayloadThumbnailData].
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

func (r *NewEventWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.completed.
type NewEventWebhookEventJobCompleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventJobCompletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.JobCompleted `json:"event" default:"job.completed"`
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
func (r NewEventWebhookEventJobCompleted) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventJobCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventJobCompletedData struct {
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
func (r NewEventWebhookEventJobCompletedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventJobCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.failed and job.cancelled.
type NewEventWebhookEventNotificationPayloadJobFailed struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventNotificationPayloadJobFailedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.failed", "job.cancelled".
	Event string `json:"event" api:"required"`
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
func (r NewEventWebhookEventNotificationPayloadJobFailed) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadJobFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventNotificationPayloadJobFailedData struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventNotificationPayloadJobFailedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadJobFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.created.
type NewEventWebhookEventJobCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventJobCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.JobCreated `json:"event" default:"job.created"`
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
func (r NewEventWebhookEventJobCreated) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventJobCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventJobCreatedData struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventJobCreatedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventJobCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for source.created.
type NewEventWebhookEventSourceCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventSourceCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.SourceCreated `json:"event" default:"source.created"`
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
func (r NewEventWebhookEventSourceCreated) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventSourceCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventSourceCreatedData struct {
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventSourceCreatedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventSourceCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for source.deleted.
type NewEventWebhookEventSourceDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventSourceDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.SourceDeleted `json:"event" default:"source.deleted"`
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
func (r NewEventWebhookEventSourceDeleted) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventSourceDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventSourceDeletedData struct {
	// Source contains the last snapshot before deletion.
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventSourceDeletedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventSourceDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.created.
type NewEventWebhookEventUploadCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventUploadCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.UploadCreated `json:"event" default:"upload.created"`
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
func (r NewEventWebhookEventUploadCreated) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventUploadCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventUploadCreatedData struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventUploadCreatedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventUploadCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.completed.
type NewEventWebhookEventUploadCompleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventUploadCompletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.UploadCompleted `json:"event" default:"upload.completed"`
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
func (r NewEventWebhookEventUploadCompleted) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventUploadCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventUploadCompletedData struct {
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
func (r NewEventWebhookEventUploadCompletedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventUploadCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.failed and upload.expired.
type NewEventWebhookEventNotificationPayloadUploadFailed struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventNotificationPayloadUploadFailedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "upload.failed", "upload.expired".
	Event string `json:"event" api:"required"`
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
func (r NewEventWebhookEventNotificationPayloadUploadFailed) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadUploadFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventNotificationPayloadUploadFailedData struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventNotificationPayloadUploadFailedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadUploadFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.created, asset.ready, asset.playback_ready, asset.failed
// and asset.expired.
type NewEventWebhookEventNotificationPayloadAsset struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventNotificationPayloadAssetData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.created", "asset.ready", "asset.playback_ready", "asset.failed",
	// "asset.expired".
	Event string `json:"event" api:"required"`
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
func (r NewEventWebhookEventNotificationPayloadAsset) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventNotificationPayloadAssetData struct {
	Asset Asset `json:"asset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewEventWebhookEventNotificationPayloadAssetData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadAssetData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.output.created, asset.output.ready and
// asset.output.failed.
type NewEventWebhookEventNotificationPayloadAssetOutput struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventNotificationPayloadAssetOutputData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.output.created", "asset.output.ready", "asset.output.failed".
	Event string `json:"event" api:"required"`
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
func (r NewEventWebhookEventNotificationPayloadAssetOutput) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadAssetOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventNotificationPayloadAssetOutputData struct {
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
func (r NewEventWebhookEventNotificationPayloadAssetOutputData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadAssetOutputData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.output.deleted.
type NewEventWebhookEventAssetOutputDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventAssetOutputDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.AssetOutputDeleted `json:"event" default:"asset.output.deleted"`
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
func (r NewEventWebhookEventAssetOutputDeleted) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventAssetOutputDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventAssetOutputDeletedData struct {
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
func (r NewEventWebhookEventAssetOutputDeletedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventAssetOutputDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.deleted.
type NewEventWebhookEventAssetDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventAssetDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.AssetDeleted `json:"event" default:"asset.deleted"`
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
func (r NewEventWebhookEventAssetDeleted) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventAssetDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventAssetDeletedData struct {
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
func (r NewEventWebhookEventAssetDeletedData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventAssetDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for workflow_run.created, workflow_run.completed,
// workflow_run.failed and workflow_run.cancelled.
type NewEventWebhookEventNotificationPayloadWorkflowRun struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventNotificationPayloadWorkflowRunData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "workflow_run.created", "workflow_run.completed", "workflow_run.failed",
	// "workflow_run.cancelled".
	Event string `json:"event" api:"required"`
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
func (r NewEventWebhookEventNotificationPayloadWorkflowRun) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadWorkflowRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventNotificationPayloadWorkflowRunData struct {
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
func (r NewEventWebhookEventNotificationPayloadWorkflowRunData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadWorkflowRunData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for transcoding.created, transcoding.ready, transcoding.failed,
// transcoding.cancelled, transcoding.expired and transcoding.deleted.
type NewEventWebhookEventNotificationPayloadTranscoding struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventNotificationPayloadTranscodingData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "transcoding.created", "transcoding.ready", "transcoding.failed",
	// "transcoding.cancelled", "transcoding.expired", "transcoding.deleted".
	Event string `json:"event" api:"required"`
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
func (r NewEventWebhookEventNotificationPayloadTranscoding) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadTranscoding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventNotificationPayloadTranscodingData struct {
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
func (r NewEventWebhookEventNotificationPayloadTranscodingData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadTranscodingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for thumbnail.created, thumbnail.ready, thumbnail.failed,
// thumbnail.expired and thumbnail.deleted.
type NewEventWebhookEventNotificationPayloadThumbnail struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NewEventWebhookEventNotificationPayloadThumbnailData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "thumbnail.created", "thumbnail.ready", "thumbnail.failed",
	// "thumbnail.expired", "thumbnail.deleted".
	Event string `json:"event" api:"required"`
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
func (r NewEventWebhookEventNotificationPayloadThumbnail) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type NewEventWebhookEventNotificationPayloadThumbnailData struct {
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
func (r NewEventWebhookEventNotificationPayloadThumbnailData) RawJSON() string { return r.JSON.raw }
func (r *NewEventWebhookEventNotificationPayloadThumbnailData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnion contains all possible properties and values from
// [UnwrapWebhookEventJobCompleted],
// [UnwrapWebhookEventNotificationPayloadJobFailed],
// [UnwrapWebhookEventJobCreated], [UnwrapWebhookEventSourceCreated],
// [UnwrapWebhookEventSourceDeleted], [UnwrapWebhookEventUploadCreated],
// [UnwrapWebhookEventUploadCompleted],
// [UnwrapWebhookEventNotificationPayloadUploadFailed],
// [UnwrapWebhookEventNotificationPayloadAsset],
// [UnwrapWebhookEventNotificationPayloadAssetOutput],
// [UnwrapWebhookEventAssetOutputDeleted], [UnwrapWebhookEventAssetDeleted],
// [UnwrapWebhookEventNotificationPayloadWorkflowRun],
// [UnwrapWebhookEventNotificationPayloadTranscoding],
// [UnwrapWebhookEventNotificationPayloadThumbnail].
//
// Use the [UnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	ID string `json:"id"`
	// This field is a union of [UnwrapWebhookEventJobCompletedData],
	// [UnwrapWebhookEventNotificationPayloadJobFailedData],
	// [UnwrapWebhookEventJobCreatedData], [UnwrapWebhookEventSourceCreatedData],
	// [UnwrapWebhookEventSourceDeletedData], [UnwrapWebhookEventUploadCreatedData],
	// [UnwrapWebhookEventUploadCompletedData],
	// [UnwrapWebhookEventNotificationPayloadUploadFailedData],
	// [UnwrapWebhookEventNotificationPayloadAssetData],
	// [UnwrapWebhookEventNotificationPayloadAssetOutputData],
	// [UnwrapWebhookEventAssetOutputDeletedData],
	// [UnwrapWebhookEventAssetDeletedData],
	// [UnwrapWebhookEventNotificationPayloadWorkflowRunData],
	// [UnwrapWebhookEventNotificationPayloadTranscodingData],
	// [UnwrapWebhookEventNotificationPayloadThumbnailData]
	Data UnwrapWebhookEventUnionData `json:"data"`
	Date time.Time                   `json:"date"`
	// Any of "job.completed", nil, "job.created", "source.created", "source.deleted",
	// "upload.created", "upload.completed", nil, nil, nil, "asset.output.deleted",
	// "asset.deleted", nil, nil, nil.
	Event string `json:"event"`
	JSON  struct {
		ID    respjson.Field
		Data  respjson.Field
		Date  respjson.Field
		Event respjson.Field
		raw   string
	} `json:"-"`
}

func (u UnwrapWebhookEventUnion) AsJobCompleted() (v UnwrapWebhookEventJobCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUnwrapWebhookEventNotificationPayloadJobFailed() (v UnwrapWebhookEventNotificationPayloadJobFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsJobCreated() (v UnwrapWebhookEventJobCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSourceCreated() (v UnwrapWebhookEventSourceCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSourceDeleted() (v UnwrapWebhookEventSourceDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadCreated() (v UnwrapWebhookEventUploadCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadCompleted() (v UnwrapWebhookEventUploadCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUnwrapWebhookEventNotificationPayloadUploadFailed() (v UnwrapWebhookEventNotificationPayloadUploadFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUnwrapWebhookEventNotificationPayloadAsset() (v UnwrapWebhookEventNotificationPayloadAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUnwrapWebhookEventNotificationPayloadAssetOutput() (v UnwrapWebhookEventNotificationPayloadAssetOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsAssetOutputDeleted() (v UnwrapWebhookEventAssetOutputDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsAssetDeleted() (v UnwrapWebhookEventAssetDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUnwrapWebhookEventNotificationPayloadWorkflowRun() (v UnwrapWebhookEventNotificationPayloadWorkflowRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUnwrapWebhookEventNotificationPayloadTranscoding() (v UnwrapWebhookEventNotificationPayloadTranscoding) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUnwrapWebhookEventNotificationPayloadThumbnail() (v UnwrapWebhookEventNotificationPayloadThumbnail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionData is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionData struct {
	// This field is from variant [UnwrapWebhookEventJobCompletedData].
	Files []JobFile `json:"files"`
	// This field is from variant [UnwrapWebhookEventJobCompletedData].
	Job Job `json:"job"`
	// This field is from variant [UnwrapWebhookEventSourceCreatedData].
	Source Source `json:"source"`
	// This field is from variant [UnwrapWebhookEventUploadCreatedData].
	Upload Upload `json:"upload"`
	// This field is from variant [UnwrapWebhookEventNotificationPayloadAssetData].
	Asset Asset `json:"asset"`
	// This field is from variant
	// [UnwrapWebhookEventNotificationPayloadAssetOutputData].
	Output       AssetOutputUnion `json:"output"`
	FilesDeleted bool             `json:"files_deleted"`
	// This field is from variant [UnwrapWebhookEventAssetDeletedData].
	DeletedSize int64 `json:"deleted_size"`
	// This field is from variant
	// [UnwrapWebhookEventNotificationPayloadWorkflowRunData].
	ActionRuns []AssetActionRun `json:"action_runs"`
	// This field is from variant
	// [UnwrapWebhookEventNotificationPayloadWorkflowRunData].
	WorkflowRun WorkflowRun `json:"workflow_run"`
	// This field is from variant
	// [UnwrapWebhookEventNotificationPayloadTranscodingData].
	Transcoding Transcoding `json:"transcoding"`
	// This field is from variant [UnwrapWebhookEventNotificationPayloadThumbnailData].
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

func (r *UnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.completed.
type UnwrapWebhookEventJobCompleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventJobCompletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.JobCompleted `json:"event" default:"job.completed"`
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
func (r UnwrapWebhookEventJobCompleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventJobCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventJobCompletedData struct {
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
func (r UnwrapWebhookEventJobCompletedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventJobCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.failed and job.cancelled.
type UnwrapWebhookEventNotificationPayloadJobFailed struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventNotificationPayloadJobFailedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.failed", "job.cancelled".
	Event string `json:"event" api:"required"`
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
func (r UnwrapWebhookEventNotificationPayloadJobFailed) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadJobFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventNotificationPayloadJobFailedData struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventNotificationPayloadJobFailedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadJobFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.created.
type UnwrapWebhookEventJobCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventJobCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.JobCreated `json:"event" default:"job.created"`
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
func (r UnwrapWebhookEventJobCreated) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventJobCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventJobCreatedData struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventJobCreatedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventJobCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for source.created.
type UnwrapWebhookEventSourceCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventSourceCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.SourceCreated `json:"event" default:"source.created"`
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
func (r UnwrapWebhookEventSourceCreated) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventSourceCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventSourceCreatedData struct {
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventSourceCreatedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventSourceCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for source.deleted.
type UnwrapWebhookEventSourceDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventSourceDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.SourceDeleted `json:"event" default:"source.deleted"`
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
func (r UnwrapWebhookEventSourceDeleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventSourceDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventSourceDeletedData struct {
	// Source contains the last snapshot before deletion.
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventSourceDeletedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventSourceDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.created.
type UnwrapWebhookEventUploadCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventUploadCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.UploadCreated `json:"event" default:"upload.created"`
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
func (r UnwrapWebhookEventUploadCreated) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventUploadCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventUploadCreatedData struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventUploadCreatedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventUploadCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.completed.
type UnwrapWebhookEventUploadCompleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventUploadCompletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.UploadCompleted `json:"event" default:"upload.completed"`
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
func (r UnwrapWebhookEventUploadCompleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventUploadCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventUploadCompletedData struct {
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
func (r UnwrapWebhookEventUploadCompletedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventUploadCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.failed and upload.expired.
type UnwrapWebhookEventNotificationPayloadUploadFailed struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventNotificationPayloadUploadFailedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "upload.failed", "upload.expired".
	Event string `json:"event" api:"required"`
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
func (r UnwrapWebhookEventNotificationPayloadUploadFailed) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadUploadFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventNotificationPayloadUploadFailedData struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventNotificationPayloadUploadFailedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadUploadFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.created, asset.ready, asset.playback_ready, asset.failed
// and asset.expired.
type UnwrapWebhookEventNotificationPayloadAsset struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventNotificationPayloadAssetData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.created", "asset.ready", "asset.playback_ready", "asset.failed",
	// "asset.expired".
	Event string `json:"event" api:"required"`
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
func (r UnwrapWebhookEventNotificationPayloadAsset) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventNotificationPayloadAssetData struct {
	Asset Asset `json:"asset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventNotificationPayloadAssetData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadAssetData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.output.created, asset.output.ready and
// asset.output.failed.
type UnwrapWebhookEventNotificationPayloadAssetOutput struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventNotificationPayloadAssetOutputData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.output.created", "asset.output.ready", "asset.output.failed".
	Event string `json:"event" api:"required"`
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
func (r UnwrapWebhookEventNotificationPayloadAssetOutput) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadAssetOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventNotificationPayloadAssetOutputData struct {
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
func (r UnwrapWebhookEventNotificationPayloadAssetOutputData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadAssetOutputData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.output.deleted.
type UnwrapWebhookEventAssetOutputDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventAssetOutputDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.AssetOutputDeleted `json:"event" default:"asset.output.deleted"`
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
func (r UnwrapWebhookEventAssetOutputDeleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventAssetOutputDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventAssetOutputDeletedData struct {
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
func (r UnwrapWebhookEventAssetOutputDeletedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventAssetOutputDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.deleted.
type UnwrapWebhookEventAssetDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventAssetDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	Event constant.AssetDeleted `json:"event" default:"asset.deleted"`
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
func (r UnwrapWebhookEventAssetDeleted) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventAssetDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventAssetDeletedData struct {
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
func (r UnwrapWebhookEventAssetDeletedData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventAssetDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for workflow_run.created, workflow_run.completed,
// workflow_run.failed and workflow_run.cancelled.
type UnwrapWebhookEventNotificationPayloadWorkflowRun struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventNotificationPayloadWorkflowRunData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "workflow_run.created", "workflow_run.completed", "workflow_run.failed",
	// "workflow_run.cancelled".
	Event string `json:"event" api:"required"`
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
func (r UnwrapWebhookEventNotificationPayloadWorkflowRun) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadWorkflowRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventNotificationPayloadWorkflowRunData struct {
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
func (r UnwrapWebhookEventNotificationPayloadWorkflowRunData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadWorkflowRunData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for transcoding.created, transcoding.ready, transcoding.failed,
// transcoding.cancelled, transcoding.expired and transcoding.deleted.
type UnwrapWebhookEventNotificationPayloadTranscoding struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventNotificationPayloadTranscodingData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "transcoding.created", "transcoding.ready", "transcoding.failed",
	// "transcoding.cancelled", "transcoding.expired", "transcoding.deleted".
	Event string `json:"event" api:"required"`
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
func (r UnwrapWebhookEventNotificationPayloadTranscoding) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadTranscoding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventNotificationPayloadTranscodingData struct {
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
func (r UnwrapWebhookEventNotificationPayloadTranscodingData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadTranscodingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for thumbnail.created, thumbnail.ready, thumbnail.failed,
// thumbnail.expired and thumbnail.deleted.
type UnwrapWebhookEventNotificationPayloadThumbnail struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data UnwrapWebhookEventNotificationPayloadThumbnailData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "thumbnail.created", "thumbnail.ready", "thumbnail.failed",
	// "thumbnail.expired", "thumbnail.deleted".
	Event string `json:"event" api:"required"`
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
func (r UnwrapWebhookEventNotificationPayloadThumbnail) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the event named in event. Replays preserve the original data;
// retrieve the resource for current state and fresh signed URLs.
type UnwrapWebhookEventNotificationPayloadThumbnailData struct {
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
func (r UnwrapWebhookEventNotificationPayloadThumbnailData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventNotificationPayloadThumbnailData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
