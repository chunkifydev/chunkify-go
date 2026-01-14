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
	opts = slices.Concat(r.Options, opts)
	path := "api/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve details of a specific webhook configuration by its ID. The webhook must
// belong to the current project.
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *Webhook, err error) {
	var env WebhookGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("api/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update the enabled status of a webhook. The webhook must belong to the current
// project.
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("api/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of all webhooks configured for the current project. Each webhook
// includes its URL, enabled status, and subscribed events.
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently delete a webhook configuration. The webhook must belong to the
// current project. This action cannot be undone.
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("api/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
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
	ID string `json:"id,required"`
	// Whether the webhook is currently enabled
	Enabled bool `json:"enabled,required"`
	// Array of event types this webhook subscribes to
	//
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
	Events []string `json:"events,required"`
	// ID of the project this webhook belongs to
	ProjectID string `json:"project_id,required"`
	// URL where webhook events will be sent
	URL string `json:"url,required"`
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
	Data []Webhook `json:"data,required"`
	// Status indicates the response status "success"
	Status constant.Success `json:"status,required"`
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
	ID string `json:"id,required"`
	// Event-specific payload data
	Data NewEventWebhookEventDataUnion `json:"data,required"`
	// Timestamp when the notification was sent
	Date time.Time `json:"date,required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
	Event NewEventWebhookEventEvent `json:"event,required"`
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
// [NewEventWebhookEventDataNotificationPayloadUploadCompleted],
// [NewEventWebhookEventDataNotificationPayloadUploadFailed].
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
	// [NewEventWebhookEventDataNotificationPayloadUploadCompleted].
	Source Source `json:"source"`
	// This field is from variant
	// [NewEventWebhookEventDataNotificationPayloadUploadCompleted].
	Upload Upload `json:"upload"`
	JSON   struct {
		Files  respjson.Field
		Job    respjson.Field
		Source respjson.Field
		Upload respjson.Field
		raw    string
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

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadUploadCompleted() (v NewEventWebhookEventDataNotificationPayloadUploadCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NewEventWebhookEventDataUnion) AsNewEventWebhookEventDataNotificationPayloadUploadFailed() (v NewEventWebhookEventDataNotificationPayloadUploadFailed) {
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
	Files []JobFile `json:"files,required"`
	Job   Job       `json:"job,required"`
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
	Job Job `json:"job,required"`
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

// Payload data structure for upload.completed events
type NewEventWebhookEventDataNotificationPayloadUploadCompleted struct {
	Source Source `json:"source,required"`
	Upload Upload `json:"upload,required"`
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
	Upload Upload `json:"upload,required"`
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

// Type of event that triggered the notification.
type NewEventWebhookEventEvent string

const (
	NewEventWebhookEventEventJobCompleted    NewEventWebhookEventEvent = "job.completed"
	NewEventWebhookEventEventJobFailed       NewEventWebhookEventEvent = "job.failed"
	NewEventWebhookEventEventJobCancelled    NewEventWebhookEventEvent = "job.cancelled"
	NewEventWebhookEventEventUploadCompleted NewEventWebhookEventEvent = "upload.completed"
	NewEventWebhookEventEventUploadFailed    NewEventWebhookEventEvent = "upload.failed"
	NewEventWebhookEventEventUploadExpired   NewEventWebhookEventEvent = "upload.expired"
)

type UnwrapWebhookEvent struct {
	// Unique identifier of the notification
	ID string `json:"id,required"`
	// Event-specific payload data
	Data UnwrapWebhookEventDataUnion `json:"data,required"`
	// Timestamp when the notification was sent
	Date time.Time `json:"date,required" format:"date-time"`
	// Type of event that triggered the notification.
	//
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
	Event UnwrapWebhookEventEvent `json:"event,required"`
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
// [UnwrapWebhookEventDataNotificationPayloadUploadCompleted],
// [UnwrapWebhookEventDataNotificationPayloadUploadFailed].
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
	// [UnwrapWebhookEventDataNotificationPayloadUploadCompleted].
	Source Source `json:"source"`
	// This field is from variant
	// [UnwrapWebhookEventDataNotificationPayloadUploadCompleted].
	Upload Upload `json:"upload"`
	JSON   struct {
		Files  respjson.Field
		Job    respjson.Field
		Source respjson.Field
		Upload respjson.Field
		raw    string
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

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadUploadCompleted() (v UnwrapWebhookEventDataNotificationPayloadUploadCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventDataUnion) AsUnwrapWebhookEventDataNotificationPayloadUploadFailed() (v UnwrapWebhookEventDataNotificationPayloadUploadFailed) {
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
	Files []JobFile `json:"files,required"`
	Job   Job       `json:"job,required"`
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
	Job Job `json:"job,required"`
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

// Payload data structure for upload.completed events
type UnwrapWebhookEventDataNotificationPayloadUploadCompleted struct {
	Source Source `json:"source,required"`
	Upload Upload `json:"upload,required"`
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
	Upload Upload `json:"upload,required"`
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

// Type of event that triggered the notification.
type UnwrapWebhookEventEvent string

const (
	UnwrapWebhookEventEventJobCompleted    UnwrapWebhookEventEvent = "job.completed"
	UnwrapWebhookEventEventJobFailed       UnwrapWebhookEventEvent = "job.failed"
	UnwrapWebhookEventEventJobCancelled    UnwrapWebhookEventEvent = "job.cancelled"
	UnwrapWebhookEventEventUploadCompleted UnwrapWebhookEventEvent = "upload.completed"
	UnwrapWebhookEventEventUploadFailed    UnwrapWebhookEventEvent = "upload.failed"
	UnwrapWebhookEventEventUploadExpired   UnwrapWebhookEventEvent = "upload.expired"
)

type WebhookNewParams struct {
	// Url is the endpoint that will receive webhook notifications, which must be a
	// valid HTTP URL.
	URL string `json:"url,required"`
	// Enabled indicates whether the webhook is active.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Events specifies the types of events that will trigger the webhook.
	//
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
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
	Data Webhook `json:"data,required"`
	// Status indicates the response status "success"
	Status constant.Success `json:"status,required"`
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
	Data Webhook `json:"data,required"`
	// Status indicates the response status "success"
	Status constant.Success `json:"status,required"`
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
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
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
