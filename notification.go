// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/chunkifydev/chunkify-go/internal/apijson"
	"github.com/chunkifydev/chunkify-go/internal/apiquery"
	"github.com/chunkifydev/chunkify-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-go/option"
	"github.com/chunkifydev/chunkify-go/packages/pagination"
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
)

// NotificationService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	Options []option.RequestOption
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r NotificationService) {
	r = NotificationService{}
	r.Options = opts
	return
}

// Create a new notification for a job event
func (r *NotificationService) New(ctx context.Context, body NotificationNewParams, opts ...option.RequestOption) (res *Notification, err error) {
	var env NotificationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "api/notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve details of a specific notification
func (r *NotificationService) Get(ctx context.Context, notificationID string, opts ...option.RequestOption) (res *Notification, err error) {
	var env NotificationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return
	}
	path := fmt.Sprintf("api/notifications/%s", notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Retrieve a list of notifications with optional filtering and pagination
func (r *NotificationService) List(ctx context.Context, query NotificationListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Notification], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/notifications"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a list of notifications with optional filtering and pagination
func (r *NotificationService) ListAutoPaging(ctx context.Context, query NotificationListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Notification] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Delete a notification.
func (r *NotificationService) Delete(ctx context.Context, notificationID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return
	}
	path := fmt.Sprintf("api/notifications/%s", notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Notification struct {
	// Unique identifier of the notification
	ID string `json:"id,required"`
	// Timestamp when the notification was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Type of event that triggered this notification
	//
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
	Event NotificationEvent `json:"event,required"`
	// ID of the object that triggered this notification
	ObjectID string `json:"object_id,required"`
	// JSON payload that was sent to the webhook endpoint
	Payload string `json:"payload,required"`
	// Webhook endpoint configuration that received this notification
	Webhook Webhook `json:"webhook,required"`
	// HTTP status code received from the webhook endpoint
	ResponseStatusCode int64 `json:"response_status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Event              respjson.Field
		ObjectID           respjson.Field
		Payload            respjson.Field
		Webhook            respjson.Field
		ResponseStatusCode respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Notification) RawJSON() string { return r.JSON.raw }
func (r *Notification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggered this notification
type NotificationEvent string

const (
	NotificationEventJobCompleted    NotificationEvent = "job.completed"
	NotificationEventJobFailed       NotificationEvent = "job.failed"
	NotificationEventJobCancelled    NotificationEvent = "job.cancelled"
	NotificationEventUploadCompleted NotificationEvent = "upload.completed"
	NotificationEventUploadFailed    NotificationEvent = "upload.failed"
	NotificationEventUploadExpired   NotificationEvent = "upload.expired"
)

type NotificationNewParams struct {
	// Event specifies the type of event that triggered the notification.
	//
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
	Event NotificationNewParamsEvent `json:"event,omitzero,required"`
	// ObjectId specifies the object that triggered this notification.
	ObjectID string `json:"object_id,required"`
	// WebhookId specifies the webhook endpoint that will receive the notification.
	WebhookID string `json:"webhook_id,required"`
	paramObj
}

func (r NotificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event specifies the type of event that triggered the notification.
type NotificationNewParamsEvent string

const (
	NotificationNewParamsEventJobCompleted    NotificationNewParamsEvent = "job.completed"
	NotificationNewParamsEventJobFailed       NotificationNewParamsEvent = "job.failed"
	NotificationNewParamsEventJobCancelled    NotificationNewParamsEvent = "job.cancelled"
	NotificationNewParamsEventUploadCompleted NotificationNewParamsEvent = "upload.completed"
	NotificationNewParamsEventUploadFailed    NotificationNewParamsEvent = "upload.failed"
	NotificationNewParamsEventUploadExpired   NotificationNewParamsEvent = "upload.expired"
)

type NotificationNewResponseEnvelope struct {
	// Data contains the response object
	Data Notification `json:"data,required"`
	// Status indicates the response status "success"
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *NotificationNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetResponseEnvelope struct {
	// Data contains the response object
	Data Notification `json:"data,required"`
	// Status indicates the response status "success"
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListParams struct {
	// Pagination limit (max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by object ID
	ObjectID param.Opt[string] `query:"object_id,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by webhook ID
	WebhookID param.Opt[string]             `query:"webhook_id,omitzero" json:"-"`
	Created   NotificationListParamsCreated `query:"created,omitzero" json:"-"`
	// Filter by events (e.g. job.completed, job.failed, upload.completed,
	// upload.failed, upload.expired)
	//
	// Any of "job.completed", "job.failed", "job.cancelled", "upload.completed",
	// "upload.failed", "upload.expired".
	Events             []string                                 `query:"events,omitzero" json:"-"`
	ResponseStatusCode NotificationListParamsResponseStatusCode `query:"response_status_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationListParams]'s query parameters as `url.Values`.
func (r NotificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationListParamsCreated struct {
	// Filter by creation date greater than or equal (RFC3339)
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Filter by creation date less than or equal (RFC3339)
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	// Sort by creation date (asc/desc)
	//
	// Any of "asc", "desc".
	Sort string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationListParamsCreated]'s query parameters as
// `url.Values`.
func (r NotificationListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationListParamsResponseStatusCode struct {
	// Filter by exact response status code
	Eq param.Opt[int64] `query:"eq,omitzero" json:"-"`
	// Filter by response status code greater than or equal
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Filter by response status code less than or equal
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationListParamsResponseStatusCode]'s query
// parameters as `url.Values`.
func (r NotificationListParamsResponseStatusCode) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
