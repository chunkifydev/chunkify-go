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

// Notification for asset.created, asset.ready, asset.playback_ready, asset.failed
// and asset.expired.
type NotificationPayloadAsset struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadAssetData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.created", "asset.ready", "asset.playback_ready", "asset.failed",
	// "asset.expired".
	Event NotificationPayloadAssetEvent `json:"event" api:"required"`
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
func (r NotificationPayloadAsset) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadAssetEvent string

const (
	NotificationPayloadAssetEventAssetCreated       NotificationPayloadAssetEvent = "asset.created"
	NotificationPayloadAssetEventAssetReady         NotificationPayloadAssetEvent = "asset.ready"
	NotificationPayloadAssetEventAssetPlaybackReady NotificationPayloadAssetEvent = "asset.playback_ready"
	NotificationPayloadAssetEventAssetFailed        NotificationPayloadAssetEvent = "asset.failed"
	NotificationPayloadAssetEventAssetExpired       NotificationPayloadAssetEvent = "asset.expired"
)

type NotificationPayloadAssetData struct {
	Asset Asset `json:"asset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPayloadAssetData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAssetData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.deleted.
type NotificationPayloadAssetDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadAssetDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.deleted".
	Event NotificationPayloadAssetDeletedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadAssetDeleted) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAssetDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadAssetDeletedEvent string

const (
	NotificationPayloadAssetDeletedEventAssetDeleted NotificationPayloadAssetDeletedEvent = "asset.deleted"
)

// Payload for asset.deleted. The Asset snapshot is preserved after record
// deletion.
type NotificationPayloadAssetDeletedData struct {
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
func (r NotificationPayloadAssetDeletedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAssetDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.output.created, asset.output.ready and
// asset.output.failed.
type NotificationPayloadAssetOutput struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadAssetOutputData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.output.created", "asset.output.ready", "asset.output.failed".
	Event NotificationPayloadAssetOutputEvent `json:"event" api:"required"`
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
func (r NotificationPayloadAssetOutput) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAssetOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadAssetOutputEvent string

const (
	NotificationPayloadAssetOutputEventAssetOutputCreated NotificationPayloadAssetOutputEvent = "asset.output.created"
	NotificationPayloadAssetOutputEventAssetOutputReady   NotificationPayloadAssetOutputEvent = "asset.output.ready"
	NotificationPayloadAssetOutputEventAssetOutputFailed  NotificationPayloadAssetOutputEvent = "asset.output.failed"
)

// Payload for Asset output creation, readiness, and failure events.
type NotificationPayloadAssetOutputData struct {
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
func (r NotificationPayloadAssetOutputData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAssetOutputData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for asset.output.deleted.
type NotificationPayloadAssetOutputDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadAssetOutputDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "asset.output.deleted".
	Event NotificationPayloadAssetOutputDeletedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadAssetOutputDeleted) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAssetOutputDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadAssetOutputDeletedEvent string

const (
	NotificationPayloadAssetOutputDeletedEventAssetOutputDeleted NotificationPayloadAssetOutputDeletedEvent = "asset.output.deleted"
)

// Payload for asset.output.deleted. The output snapshot is preserved after
// removal.
type NotificationPayloadAssetOutputDeletedData struct {
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
func (r NotificationPayloadAssetOutputDeletedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadAssetOutputDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.completed.
type NotificationPayloadJobCompleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadJobCompletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.completed".
	Event NotificationPayloadJobCompletedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadJobCompleted) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadJobCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadJobCompletedEvent string

const (
	NotificationPayloadJobCompletedEventJobCompleted NotificationPayloadJobCompletedEvent = "job.completed"
)

// Payload data structure for job.completed events
type NotificationPayloadJobCompletedData struct {
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
func (r NotificationPayloadJobCompletedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadJobCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.created.
type NotificationPayloadJobCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadJobCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.created".
	Event NotificationPayloadJobCreatedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadJobCreated) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadJobCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadJobCreatedEvent string

const (
	NotificationPayloadJobCreatedEventJobCreated NotificationPayloadJobCreatedEvent = "job.created"
)

// Payload for job.created. A standalone Job has been accepted for processing;
// execution may not have started.
type NotificationPayloadJobCreatedData struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPayloadJobCreatedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadJobCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for job.failed and job.cancelled.
type NotificationPayloadJobFailed struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadJobFailedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.failed", "job.cancelled".
	Event NotificationPayloadJobFailedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadJobFailed) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadJobFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadJobFailedEvent string

const (
	NotificationPayloadJobFailedEventJobFailed    NotificationPayloadJobFailedEvent = "job.failed"
	NotificationPayloadJobFailedEventJobCancelled NotificationPayloadJobFailedEvent = "job.cancelled"
)

// Payload data structure for job.failed and job.cancelled events
type NotificationPayloadJobFailedData struct {
	Job Job `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPayloadJobFailedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadJobFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for source.created.
type NotificationPayloadSourceCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadSourceCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "source.created".
	Event NotificationPayloadSourceCreatedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadSourceCreated) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadSourceCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadSourceCreatedEvent string

const (
	NotificationPayloadSourceCreatedEventSourceCreated NotificationPayloadSourceCreatedEvent = "source.created"
)

// Payload for source.created. The Source has been validated and saved. Thumbnails
// may still be pending.
type NotificationPayloadSourceCreatedData struct {
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPayloadSourceCreatedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadSourceCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for source.deleted.
type NotificationPayloadSourceDeleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadSourceDeletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "source.deleted".
	Event NotificationPayloadSourceDeletedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadSourceDeleted) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadSourceDeleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadSourceDeletedEvent string

const (
	NotificationPayloadSourceDeletedEventSourceDeleted NotificationPayloadSourceDeletedEvent = "source.deleted"
)

// Payload for source.deleted. The Source was deleted on request or after the
// team's data retention period. Resources created from it are preserved and no
// longer report its source_id.
type NotificationPayloadSourceDeletedData struct {
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
func (r NotificationPayloadSourceDeletedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadSourceDeletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for thumbnail.created, thumbnail.ready, thumbnail.failed,
// thumbnail.expired and thumbnail.deleted.
type NotificationPayloadThumbnail struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadThumbnailData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "thumbnail.created", "thumbnail.ready", "thumbnail.failed",
	// "thumbnail.expired", "thumbnail.deleted".
	Event NotificationPayloadThumbnailEvent `json:"event" api:"required"`
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
func (r NotificationPayloadThumbnail) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadThumbnailEvent string

const (
	NotificationPayloadThumbnailEventThumbnailCreated NotificationPayloadThumbnailEvent = "thumbnail.created"
	NotificationPayloadThumbnailEventThumbnailReady   NotificationPayloadThumbnailEvent = "thumbnail.ready"
	NotificationPayloadThumbnailEventThumbnailFailed  NotificationPayloadThumbnailEvent = "thumbnail.failed"
	NotificationPayloadThumbnailEventThumbnailExpired NotificationPayloadThumbnailEvent = "thumbnail.expired"
	NotificationPayloadThumbnailEventThumbnailDeleted NotificationPayloadThumbnailEvent = "thumbnail.deleted"
)

// A canonical thumbnail snapshot for standalone thumbnail.\* events. Asset-owned
// thumbnails use the Asset output event payload.
type NotificationPayloadThumbnailData struct {
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
func (r NotificationPayloadThumbnailData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadThumbnailData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for transcoding.created, transcoding.ready, transcoding.failed,
// transcoding.cancelled, transcoding.expired and transcoding.deleted.
type NotificationPayloadTranscoding struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadTranscodingData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "transcoding.created", "transcoding.ready", "transcoding.failed",
	// "transcoding.cancelled", "transcoding.expired", "transcoding.deleted".
	Event NotificationPayloadTranscodingEvent `json:"event" api:"required"`
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
func (r NotificationPayloadTranscoding) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadTranscoding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadTranscodingEvent string

const (
	NotificationPayloadTranscodingEventTranscodingCreated   NotificationPayloadTranscodingEvent = "transcoding.created"
	NotificationPayloadTranscodingEventTranscodingReady     NotificationPayloadTranscodingEvent = "transcoding.ready"
	NotificationPayloadTranscodingEventTranscodingFailed    NotificationPayloadTranscodingEvent = "transcoding.failed"
	NotificationPayloadTranscodingEventTranscodingCancelled NotificationPayloadTranscodingEvent = "transcoding.cancelled"
	NotificationPayloadTranscodingEventTranscodingExpired   NotificationPayloadTranscodingEvent = "transcoding.expired"
	NotificationPayloadTranscodingEventTranscodingDeleted   NotificationPayloadTranscodingEvent = "transcoding.deleted"
)

// Payload for transcoding.created, transcoding.ready, transcoding.failed,
// transcoding.cancelled, transcoding.expired and transcoding.deleted events.
type NotificationPayloadTranscodingData struct {
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
func (r NotificationPayloadTranscodingData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadTranscodingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.completed.
type NotificationPayloadUploadCompleted struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadUploadCompletedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "upload.completed".
	Event NotificationPayloadUploadCompletedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadUploadCompleted) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadUploadCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadUploadCompletedEvent string

const (
	NotificationPayloadUploadCompletedEventUploadCompleted NotificationPayloadUploadCompletedEvent = "upload.completed"
)

// Payload data structure for upload.completed events
type NotificationPayloadUploadCompletedData struct {
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
func (r NotificationPayloadUploadCompletedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadUploadCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.created.
type NotificationPayloadUploadCreated struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadUploadCreatedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "upload.created".
	Event NotificationPayloadUploadCreatedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadUploadCreated) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadUploadCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadUploadCreatedEvent string

const (
	NotificationPayloadUploadCreatedEventUploadCreated NotificationPayloadUploadCreatedEvent = "upload.created"
)

// Payload for upload.created. A waiting upload session has been created; its
// Source is created only after successful completion. Transfer credentials are not
// included.
type NotificationPayloadUploadCreatedData struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPayloadUploadCreatedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadUploadCreatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for upload.failed and upload.expired.
type NotificationPayloadUploadFailed struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadUploadFailedData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "upload.failed", "upload.expired".
	Event NotificationPayloadUploadFailedEvent `json:"event" api:"required"`
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
func (r NotificationPayloadUploadFailed) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadUploadFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadUploadFailedEvent string

const (
	NotificationPayloadUploadFailedEventUploadFailed  NotificationPayloadUploadFailedEvent = "upload.failed"
	NotificationPayloadUploadFailedEventUploadExpired NotificationPayloadUploadFailedEvent = "upload.expired"
)

// Payload data structure for upload.failed and upload.expired events
type NotificationPayloadUploadFailedData struct {
	Upload Upload `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPayloadUploadFailedData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadUploadFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification for workflow_run.created, workflow_run.completed,
// workflow_run.failed and workflow_run.cancelled.
type NotificationPayloadWorkflowRun struct {
	// Unique identifier of the notification
	ID string `json:"id" api:"required"`
	// Payload for the event named in event. Replays preserve the original data;
	// retrieve the resource for current state and fresh signed URLs.
	Data NotificationPayloadWorkflowRunData `json:"data" api:"required"`
	// Timestamp of the event payload.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "workflow_run.created", "workflow_run.completed", "workflow_run.failed",
	// "workflow_run.cancelled".
	Event NotificationPayloadWorkflowRunEvent `json:"event" api:"required"`
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
func (r NotificationPayloadWorkflowRun) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadWorkflowRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered the notification.
type NotificationPayloadWorkflowRunEvent string

const (
	NotificationPayloadWorkflowRunEventWorkflowRunCreated   NotificationPayloadWorkflowRunEvent = "workflow_run.created"
	NotificationPayloadWorkflowRunEventWorkflowRunCompleted NotificationPayloadWorkflowRunEvent = "workflow_run.completed"
	NotificationPayloadWorkflowRunEventWorkflowRunFailed    NotificationPayloadWorkflowRunEvent = "workflow_run.failed"
	NotificationPayloadWorkflowRunEventWorkflowRunCancelled NotificationPayloadWorkflowRunEvent = "workflow_run.cancelled"
)

// Payload for workflow_run.created, workflow_run.completed, workflow_run.failed,
// and workflow_run.cancelled. The run and its actions are snapshots taken when the
// event was recorded.
type NotificationPayloadWorkflowRunData struct {
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
func (r NotificationPayloadWorkflowRunData) RawJSON() string { return r.JSON.raw }
func (r *NotificationPayloadWorkflowRunData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Webhook struct {
	// Unique identifier of the webhook
	ID string `json:"id" api:"required"`
	// Whether the webhook is currently enabled
	Enabled bool `json:"enabled" api:"required"`
	// Array of event types this webhook subscribes to
	Events []WebhookEvent `json:"events" api:"required"`
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

type WebhookEvent string

const (
	WebhookEventJobCreated           WebhookEvent = "job.created"
	WebhookEventJobCompleted         WebhookEvent = "job.completed"
	WebhookEventJobFailed            WebhookEvent = "job.failed"
	WebhookEventJobCancelled         WebhookEvent = "job.cancelled"
	WebhookEventSourceCreated        WebhookEvent = "source.created"
	WebhookEventSourceDeleted        WebhookEvent = "source.deleted"
	WebhookEventUploadCreated        WebhookEvent = "upload.created"
	WebhookEventUploadCompleted      WebhookEvent = "upload.completed"
	WebhookEventUploadFailed         WebhookEvent = "upload.failed"
	WebhookEventUploadExpired        WebhookEvent = "upload.expired"
	WebhookEventAssetCreated         WebhookEvent = "asset.created"
	WebhookEventAssetReady           WebhookEvent = "asset.ready"
	WebhookEventAssetPlaybackReady   WebhookEvent = "asset.playback_ready"
	WebhookEventAssetFailed          WebhookEvent = "asset.failed"
	WebhookEventAssetExpired         WebhookEvent = "asset.expired"
	WebhookEventAssetDeleted         WebhookEvent = "asset.deleted"
	WebhookEventAssetOutputCreated   WebhookEvent = "asset.output.created"
	WebhookEventAssetOutputReady     WebhookEvent = "asset.output.ready"
	WebhookEventAssetOutputFailed    WebhookEvent = "asset.output.failed"
	WebhookEventAssetOutputDeleted   WebhookEvent = "asset.output.deleted"
	WebhookEventWorkflowRunCreated   WebhookEvent = "workflow_run.created"
	WebhookEventWorkflowRunCompleted WebhookEvent = "workflow_run.completed"
	WebhookEventWorkflowRunFailed    WebhookEvent = "workflow_run.failed"
	WebhookEventWorkflowRunCancelled WebhookEvent = "workflow_run.cancelled"
	WebhookEventTranscodingCreated   WebhookEvent = "transcoding.created"
	WebhookEventTranscodingReady     WebhookEvent = "transcoding.ready"
	WebhookEventTranscodingFailed    WebhookEvent = "transcoding.failed"
	WebhookEventTranscodingCancelled WebhookEvent = "transcoding.cancelled"
	WebhookEventTranscodingExpired   WebhookEvent = "transcoding.expired"
	WebhookEventTranscodingDeleted   WebhookEvent = "transcoding.deleted"
	WebhookEventThumbnailCreated     WebhookEvent = "thumbnail.created"
	WebhookEventThumbnailReady       WebhookEvent = "thumbnail.ready"
	WebhookEventThumbnailFailed      WebhookEvent = "thumbnail.failed"
	WebhookEventThumbnailExpired     WebhookEvent = "thumbnail.expired"
	WebhookEventThumbnailDeleted     WebhookEvent = "thumbnail.deleted"
)

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

// UnwrapWebhookEventUnion contains all possible properties and values from
// [NotificationPayloadJobCompleted], [NotificationPayloadJobFailed],
// [NotificationPayloadJobCreated], [NotificationPayloadSourceCreated],
// [NotificationPayloadSourceDeleted], [NotificationPayloadUploadCreated],
// [NotificationPayloadUploadCompleted], [NotificationPayloadUploadFailed],
// [NotificationPayloadAsset], [NotificationPayloadAssetOutput],
// [NotificationPayloadAssetOutputDeleted], [NotificationPayloadAssetDeleted],
// [NotificationPayloadWorkflowRun], [NotificationPayloadTranscoding],
// [NotificationPayloadThumbnail].
//
// Use the [UnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	ID string `json:"id"`
	// This field is a union of [NotificationPayloadJobCompletedData],
	// [NotificationPayloadJobFailedData], [NotificationPayloadJobCreatedData],
	// [NotificationPayloadSourceCreatedData], [NotificationPayloadSourceDeletedData],
	// [NotificationPayloadUploadCreatedData],
	// [NotificationPayloadUploadCompletedData], [NotificationPayloadUploadFailedData],
	// [NotificationPayloadAssetData], [NotificationPayloadAssetOutputData],
	// [NotificationPayloadAssetOutputDeletedData],
	// [NotificationPayloadAssetDeletedData], [NotificationPayloadWorkflowRunData],
	// [NotificationPayloadTranscodingData], [NotificationPayloadThumbnailData]
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

func (u UnwrapWebhookEventUnion) AsJobCompleted() (v NotificationPayloadJobCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNotificationPayloadJobFailed() (v NotificationPayloadJobFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsJobCreated() (v NotificationPayloadJobCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSourceCreated() (v NotificationPayloadSourceCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSourceDeleted() (v NotificationPayloadSourceDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadCreated() (v NotificationPayloadUploadCreated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadCompleted() (v NotificationPayloadUploadCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNotificationPayloadUploadFailed() (v NotificationPayloadUploadFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNotificationPayloadAsset() (v NotificationPayloadAsset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNotificationPayloadAssetOutput() (v NotificationPayloadAssetOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsAssetOutputDeleted() (v NotificationPayloadAssetOutputDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsAssetDeleted() (v NotificationPayloadAssetDeleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNotificationPayloadWorkflowRun() (v NotificationPayloadWorkflowRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNotificationPayloadTranscoding() (v NotificationPayloadTranscoding) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNotificationPayloadThumbnail() (v NotificationPayloadThumbnail) {
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
	// This field is from variant [NotificationPayloadJobCompletedData].
	Files []JobFile `json:"files"`
	// This field is from variant [NotificationPayloadJobCompletedData].
	Job Job `json:"job"`
	// This field is from variant [NotificationPayloadSourceCreatedData].
	Source Source `json:"source"`
	// This field is from variant [NotificationPayloadUploadCreatedData].
	Upload Upload `json:"upload"`
	// This field is from variant [NotificationPayloadAssetData].
	Asset Asset `json:"asset"`
	// This field is from variant [NotificationPayloadAssetOutputData].
	Output       AssetOutputUnion `json:"output"`
	FilesDeleted bool             `json:"files_deleted"`
	// This field is from variant [NotificationPayloadAssetDeletedData].
	DeletedSize int64 `json:"deleted_size"`
	// This field is from variant [NotificationPayloadWorkflowRunData].
	ActionRuns []AssetActionRun `json:"action_runs"`
	// This field is from variant [NotificationPayloadWorkflowRunData].
	WorkflowRun WorkflowRun `json:"workflow_run"`
	// This field is from variant [NotificationPayloadTranscodingData].
	Transcoding Transcoding `json:"transcoding"`
	// This field is from variant [NotificationPayloadThumbnailData].
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

type WebhookNewParams struct {
	// Url is the endpoint that will receive webhook notifications, which must be a
	// valid HTTP URL.
	URL string `json:"url" api:"required"`
	// Enabled indicates whether the webhook is active.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Events specifies the types of events that will trigger the webhook.
	Events []WebhookEvent `json:"events,omitzero"`
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
	Events []WebhookEvent `json:"events,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
