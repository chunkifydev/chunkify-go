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
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// EventService contains methods and other services that help with interacting with
// the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// Retrieve a retained event, its frozen payload, related resources, and up to 100
// delivery attempts. No webhook configuration is required.
func (r *EventService) Get(ctx context.Context, eventID string, opts ...option.RequestOption) (res *Event, err error) {
	var env EventGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("events/%s", eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List lifecycle events recorded with or without webhooks. Events and their
// deliveries expire according to the plan retention deadline. Historical
// notifications recorded before the event log remain available through
// Notifications.
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Event], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events"
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

// List lifecycle events recorded with or without webhooks. Events and their
// deliveries expire according to the plan retention deadline. Historical
// notifications recorded before the event log remain available through
// Notifications.
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Event] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// A retained lifecycle event, recorded independently of webhook subscriptions. The
// payload is an immutable JSON snapshot.
type Event struct {
	// Unique event ID.
	ID string `json:"id" api:"required"`
	// Time the event occurred.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Number of retained webhook delivery attempts.
	DeliveryCount int64 `json:"delivery_count" api:"required"`
	// Summary of the latest retained attempt for each webhook. none means no retained
	// attempts; pending means at least one latest attempt is awaiting delivery; failed
	// means at least one latest attempt failed and none are pending; delivered means
	// all latest attempts succeeded. Earlier failures superseded by successful retries
	// do not mark the event failed.
	//
	// Any of "none", "pending", "delivered", "failed".
	DeliveryStatus EventDeliveryStatus `json:"delivery_status" api:"required"`
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
	Event EventEvent `json:"event" api:"required"`
	// Retention deadline set from the project team plan when the event was recorded.
	// Deliveries and retries share this deadline.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// ID of the resource that caused this event.
	ObjectID string `json:"object_id" api:"required"`
	// Immutable JSON event envelope. Its ID is the event ID; webhook delivery
	// envelopes retain their notification IDs. Signed links in historical snapshots
	// may have expired.
	Payload string `json:"payload" api:"required"`
	// Related resource IDs retained independently of resource deletion.
	RelatedObjectIDs []string `json:"related_object_ids" api:"required"`
	// First 100 delivery attempts, oldest first; included on retrieval. List
	// notifications with event_id to paginate all attempts.
	Deliveries []Notification `json:"deliveries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DeliveryCount    respjson.Field
		DeliveryStatus   respjson.Field
		Event            respjson.Field
		ExpiresAt        respjson.Field
		ObjectID         respjson.Field
		Payload          respjson.Field
		RelatedObjectIDs respjson.Field
		Deliveries       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of the latest retained attempt for each webhook. none means no retained
// attempts; pending means at least one latest attempt is awaiting delivery; failed
// means at least one latest attempt failed and none are pending; delivered means
// all latest attempts succeeded. Earlier failures superseded by successful retries
// do not mark the event failed.
type EventDeliveryStatus string

const (
	EventDeliveryStatusNone      EventDeliveryStatus = "none"
	EventDeliveryStatusPending   EventDeliveryStatus = "pending"
	EventDeliveryStatusDelivered EventDeliveryStatus = "delivered"
	EventDeliveryStatusFailed    EventDeliveryStatus = "failed"
)

type EventEvent string

const (
	EventEventJobCreated           EventEvent = "job.created"
	EventEventJobCompleted         EventEvent = "job.completed"
	EventEventJobFailed            EventEvent = "job.failed"
	EventEventJobCancelled         EventEvent = "job.cancelled"
	EventEventSourceCreated        EventEvent = "source.created"
	EventEventSourceDeleted        EventEvent = "source.deleted"
	EventEventUploadCreated        EventEvent = "upload.created"
	EventEventUploadCompleted      EventEvent = "upload.completed"
	EventEventUploadFailed         EventEvent = "upload.failed"
	EventEventUploadExpired        EventEvent = "upload.expired"
	EventEventAssetCreated         EventEvent = "asset.created"
	EventEventAssetReady           EventEvent = "asset.ready"
	EventEventAssetPlaybackReady   EventEvent = "asset.playback_ready"
	EventEventAssetFailed          EventEvent = "asset.failed"
	EventEventAssetExpired         EventEvent = "asset.expired"
	EventEventAssetDeleted         EventEvent = "asset.deleted"
	EventEventAssetOutputCreated   EventEvent = "asset.output.created"
	EventEventAssetOutputReady     EventEvent = "asset.output.ready"
	EventEventAssetOutputFailed    EventEvent = "asset.output.failed"
	EventEventAssetOutputDeleted   EventEvent = "asset.output.deleted"
	EventEventWorkflowRunCreated   EventEvent = "workflow_run.created"
	EventEventWorkflowRunCompleted EventEvent = "workflow_run.completed"
	EventEventWorkflowRunFailed    EventEvent = "workflow_run.failed"
	EventEventWorkflowRunCancelled EventEvent = "workflow_run.cancelled"
	EventEventTranscodingCreated   EventEvent = "transcoding.created"
	EventEventTranscodingReady     EventEvent = "transcoding.ready"
	EventEventTranscodingFailed    EventEvent = "transcoding.failed"
	EventEventTranscodingCancelled EventEvent = "transcoding.cancelled"
	EventEventTranscodingExpired   EventEvent = "transcoding.expired"
	EventEventTranscodingDeleted   EventEvent = "transcoding.deleted"
	EventEventThumbnailCreated     EventEvent = "thumbnail.created"
	EventEventThumbnailReady       EventEvent = "thumbnail.ready"
	EventEventThumbnailFailed      EventEvent = "thumbnail.failed"
	EventEventThumbnailExpired     EventEvent = "thumbnail.expired"
	EventEventThumbnailDeleted     EventEvent = "thumbnail.deleted"
)

type EventGetResponseEnvelope struct {
	// A retained lifecycle event, recorded independently of webhook subscriptions. The
	// payload is an immutable JSON snapshot.
	Data Event `json:"data" api:"required"`
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
func (r EventGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *EventGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
	// Pagination limit (max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by the primary or a related resource ID, including retained associations
	// after resource deletion.
	ObjectID param.Opt[string] `query:"object_id,omitzero" json:"-"`
	// Pagination offset
	Offset  param.Opt[int64]       `query:"offset,omitzero" json:"-"`
	Created EventListParamsCreated `query:"created,omitzero" json:"-"`
	// Filter by events
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
	Events []string `query:"events,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventListParamsCreated struct {
	// Filter by creation date greater than or equal (UNIX epoch time)
	Gte param.Opt[int64] `query:"gte,omitzero" format:"epoch" json:"-"`
	// Filter by creation date less than or equal (UNIX epoch time)
	Lte param.Opt[int64] `query:"lte,omitzero" format:"epoch" json:"-"`
	// Sort by creation date (asc/desc)
	//
	// Any of "asc", "desc".
	Sort string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParamsCreated]'s query parameters as `url.Values`.
func (r EventListParamsCreated) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
