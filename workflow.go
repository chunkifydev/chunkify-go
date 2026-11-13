// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"encoding/json"
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

// WorkflowService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options   []option.RequestOption
	Revisions WorkflowRevisionService
	Runs      WorkflowRunService
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.Options = opts
	r.Revisions = NewWorkflowRevisionService(opts...)
	r.Runs = NewWorkflowRunService(opts...)
	return
}

// Create a Workflow with its first definition revision. Select it when creating an
// Asset or Upload, or set it as the Project default.
func (r *WorkflowService) New(ctx context.Context, body WorkflowNewParams, opts ...option.RequestOption) (res *Workflow, err error) {
	var env WorkflowNewResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "workflows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve a Workflow and its current definition, including archived Workflows.
func (r *WorkflowService) Get(ctx context.Context, workflowID string, opts ...option.RequestOption) (res *Workflow, err error) {
	var env WorkflowGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workflows/%s", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update the Workflow name or replace its rules. Replacing rules requires the
// current revision. A changed definition creates a new revision; unchanged rules
// and name-only edits preserve it. Existing Assets and Uploads keep their selected
// revision until a ready Asset explicitly updates to the latest revision. Archived
// Workflows cannot be updated.
func (r *WorkflowService) Update(ctx context.Context, workflowID string, body WorkflowUpdateParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return err
	}
	path := fmt.Sprintf("workflows/%s", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// List Workflows in the Project, newest first. Archived Workflows are excluded
// unless archived is true.
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *pagination.PaginatedResults[Workflow], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "workflows"
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

// List Workflows in the Project, newest first. Archived Workflows are excluded
// unless archived is true.
func (r *WorkflowService) ListAutoPaging(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) *pagination.PaginatedResultsAutoPager[Workflow] {
	return pagination.NewPaginatedResultsAutoPager(r.List(ctx, query, opts...))
}

// Archive the Workflow and clear it from the Project default. New Assets cannot
// select it and new events do not start runs. Accepted actions finish, and outputs
// and retained history remain available.
func (r *WorkflowService) Delete(ctx context.Context, workflowID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return err
	}
	path := fmt.Sprintf("workflows/%s", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A reusable definition of actions triggered by Asset events.
type Workflow struct {
	// Unique Workflow ID.
	ID string `json:"id" api:"required"`
	// When the Workflow was archived, or null if it is available for selection.
	// Archiving stops new runs; accepted actions finish.
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	// When the Workflow was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name specifies the Workflow name. Use 1 to 128 characters.
	Name string `json:"name" api:"required"`
	// Current definition revision. Starts at 1 and increases when rules are replaced.
	Revision int64 `json:"revision" api:"required"`
	// Rules specifies the event subscriptions and actions in this definition. Provide
	// exactly one asset.ready rule and up to 15 connected rules, with at most 32
	// actions and 65536 encoded bytes across all rules. Array order affects display
	// only.
	Rules []WorkflowRule `json:"rules" api:"required"`
	// SchemaVersion specifies the Workflow definition format. The supported version
	// is 1.
	SchemaVersion int64 `json:"schema_version" api:"required"`
	// When the Workflow was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		Name          respjson.Field
		Revision      respjson.Field
		Rules         respjson.Field
		SchemaVersion respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A starting or connected step with one or more independent actions.
type WorkflowRule struct {
	// Id specifies a stable rule key, unique within this Workflow definition. Use 1 to
	// 64 characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Actions specifies the work to start when the rule matches. Actions can run
	// concurrently; array order does not define execution order. Provide 1 to 32
	// actions per rule, with at most 32 across the definition.
	Actions []WorkflowRuleActionUnion `json:"actions" api:"required"`
	// Event specifies when this rule starts. Use asset.ready for the single starting
	// rule. Output ready events require an origin caption or static rendition action
	// in this Workflow.
	//
	// Any of "asset.ready", "asset.output.ready".
	Event string `json:"event" api:"required"`
	// Origin specifies the action whose ready output starts this rule. Required for
	// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
	// create_caption action. References must point to a compatible action in this
	// definition and must not form a cycle.
	Origin WorkflowRuleOrigin `json:"origin"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		Event       respjson.Field
		Origin      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRule) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowRuleActionUnion contains all possible properties and values from
// [WorkflowRuleActionCreateStaticRendition], [WorkflowRuleActionCreateCaption],
// [WorkflowRuleActionTranslateCaption], [WorkflowRuleActionCreateThumbnails],
// [WorkflowRuleActionCreateTranscript].
//
// Use the [WorkflowRuleActionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowRuleActionUnion struct {
	ID string `json:"id"`
	// This field is a union of [WorkflowRuleActionCreateStaticRenditionParams], [any],
	// [WorkflowRuleActionTranslateCaptionParams],
	// [WorkflowRuleActionCreateThumbnailsParams], [any]
	Params WorkflowRuleActionUnionParams `json:"params"`
	// Any of "create_static_rendition", "create_caption", "translate_caption",
	// "create_thumbnails", "create_transcript".
	Type string `json:"type"`
	// This field is from variant [WorkflowRuleActionCreateCaption].
	TranscriptID WorkflowRuleActionCreateCaptionTranscriptID `json:"transcript_id"`
	// This field is from variant [WorkflowRuleActionTranslateCaption].
	CaptionID WorkflowRuleActionTranslateCaptionCaptionID `json:"caption_id"`
	JSON      struct {
		ID           respjson.Field
		Params       respjson.Field
		Type         respjson.Field
		TranscriptID respjson.Field
		CaptionID    respjson.Field
		raw          string
	} `json:"-"`
}

// anyWorkflowRuleAction is implemented by each variant of
// [WorkflowRuleActionUnion] to add type safety for the return type of
// [WorkflowRuleActionUnion.AsAny]
type anyWorkflowRuleAction interface {
	implWorkflowRuleActionUnion()
}

func (WorkflowRuleActionCreateStaticRendition) implWorkflowRuleActionUnion() {}
func (WorkflowRuleActionCreateCaption) implWorkflowRuleActionUnion()         {}
func (WorkflowRuleActionTranslateCaption) implWorkflowRuleActionUnion()      {}
func (WorkflowRuleActionCreateThumbnails) implWorkflowRuleActionUnion()      {}
func (WorkflowRuleActionCreateTranscript) implWorkflowRuleActionUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WorkflowRuleActionUnion.AsAny().(type) {
//	case chunkify.WorkflowRuleActionCreateStaticRendition:
//	case chunkify.WorkflowRuleActionCreateCaption:
//	case chunkify.WorkflowRuleActionTranslateCaption:
//	case chunkify.WorkflowRuleActionCreateThumbnails:
//	case chunkify.WorkflowRuleActionCreateTranscript:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WorkflowRuleActionUnion) AsAny() anyWorkflowRuleAction {
	switch u.Type {
	case "create_static_rendition":
		return u.AsCreateStaticRendition()
	case "create_caption":
		return u.AsCreateCaption()
	case "translate_caption":
		return u.AsTranslateCaption()
	case "create_thumbnails":
		return u.AsCreateThumbnails()
	case "create_transcript":
		return u.AsCreateTranscript()
	}
	return nil
}

func (u WorkflowRuleActionUnion) AsCreateStaticRendition() (v WorkflowRuleActionCreateStaticRendition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRuleActionUnion) AsCreateCaption() (v WorkflowRuleActionCreateCaption) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRuleActionUnion) AsTranslateCaption() (v WorkflowRuleActionTranslateCaption) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRuleActionUnion) AsCreateThumbnails() (v WorkflowRuleActionCreateThumbnails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRuleActionUnion) AsCreateTranscript() (v WorkflowRuleActionCreateTranscript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowRuleActionUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowRuleActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowRuleActionUnionParams is an implicit subunion of
// [WorkflowRuleActionUnion]. WorkflowRuleActionUnionParams provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowRuleActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfWorkflowRuleActionCreateTranscripts]
type WorkflowRuleActionUnionParams struct {
	// This field will be present if the value is a [any] instead of an object.
	OfWorkflowRuleActionCreateTranscripts any `json:",inline"`
	// This field is from variant [WorkflowRuleActionCreateStaticRenditionParams].
	Resolution string `json:"resolution"`
	// This field is from variant [WorkflowRuleActionTranslateCaptionParams].
	LanguageCode string `json:"language_code"`
	// This field is from variant [WorkflowRuleActionCreateThumbnailsParams].
	Count int64 `json:"count"`
	// This field is from variant [WorkflowRuleActionCreateThumbnailsParams].
	Format string `json:"format"`
	// This field is from variant [WorkflowRuleActionCreateThumbnailsParams].
	Metadata map[string]string `json:"metadata"`
	// This field is from variant [WorkflowRuleActionCreateThumbnailsParams].
	Output string `json:"output"`
	// This field is from variant [WorkflowRuleActionCreateThumbnailsParams].
	Range WorkflowRuleActionCreateThumbnailsParamsRange `json:"range"`
	// This field is from variant [WorkflowRuleActionCreateThumbnailsParams].
	Size int64 `json:"size"`
	// This field is from variant [WorkflowRuleActionCreateThumbnailsParams].
	Timestamps []float64 `json:"timestamps"`
	JSON       struct {
		OfWorkflowRuleActionCreateTranscripts respjson.Field
		Resolution                            respjson.Field
		LanguageCode                          respjson.Field
		Count                                 respjson.Field
		Format                                respjson.Field
		Metadata                              respjson.Field
		Output                                respjson.Field
		Range                                 respjson.Field
		Size                                  respjson.Field
		Timestamps                            respjson.Field
		raw                                   string
	} `json:"-"`
}

func (r *WorkflowRuleActionUnionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses one static rendition for the Asset.
type WorkflowRuleActionCreateStaticRendition struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies the static rendition resolution.
	Params WorkflowRuleActionCreateStaticRenditionParams `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_static_rendition.
	Type constant.CreateStaticRendition `json:"type" default:"create_static_rendition"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateStaticRendition) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateStaticRendition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the static rendition resolution.
type WorkflowRuleActionCreateStaticRenditionParams struct {
	// Resolution specifies the H.264 MP4 output resolution. highest selects the
	// largest supported output without upscaling. Hobby supports up to 1080p; other
	// plans support up to 2160p. Streaming resolution limits do not apply.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution string `json:"resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Resolution  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateStaticRenditionParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateStaticRenditionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
type WorkflowRuleActionCreateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params any `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_caption.
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowRuleActionCreateCaptionTranscriptID `json:"transcript_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Params       respjson.Field
		Type         respjson.Field
		TranscriptID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateCaption) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TranscriptId binds caption publication to the transcript in the triggering
// event. Use it only when connected to create_transcript. When omitted, creates or
// reuses the Asset’s original transcript.
type WorkflowRuleActionCreateCaptionTranscriptID struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateCaptionTranscriptID) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateCaptionTranscriptID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translates or reuses an original caption in one target language. Requires an
// asset.output.ready rule connected to an original caption creation action.
type WorkflowRuleActionTranslateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// CaptionId specifies the binding to the original caption in the triggering event.
	CaptionID WorkflowRuleActionTranslateCaptionCaptionID `json:"caption_id" api:"required"`
	// Params specifies the target language.
	Params WorkflowRuleActionTranslateCaptionParams `json:"params" api:"required"`
	// Type specifies the action to perform. Must be translate_caption.
	Type constant.TranslateCaption `json:"type" default:"translate_caption"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CaptionID   respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionTranslateCaption) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionTranslateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CaptionId specifies the binding to the original caption in the triggering event.
type WorkflowRuleActionTranslateCaptionCaptionID struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionTranslateCaptionCaptionID) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionTranslateCaptionCaptionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the target language.
type WorkflowRuleActionTranslateCaptionParams struct {
	// LanguageCode specifies the target language for the translated subtitle track,
	// such as en or fr. The same code appears in the caption and HLS manifest.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode string `json:"language_code" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionTranslateCaptionParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionTranslateCaptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
type WorkflowRuleActionCreateThumbnails struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowRuleActionCreateThumbnailsParams `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_thumbnails.
	Type constant.CreateThumbnails `json:"type" default:"create_thumbnails"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateThumbnails) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateThumbnails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies sampling and image settings. Outputs use the Asset folder and
// Storage.
type WorkflowRuleActionCreateThumbnailsParams struct {
	// Count specifies 1 to 100 thumbnails at the middle of equal sections of the
	// selected range.
	Count int64 `json:"count"`
	// Format selects JPEG or WebP. Defaults to JPEG.
	//
	// Any of "jpeg", "webp".
	Format string `json:"format"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output string `json:"output"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range WorkflowRuleActionCreateThumbnailsParamsRange `json:"range"`
	// Size specifies the maximum image or tile edge in pixels, from 64 to 1920.
	// Defaults to 320. Odd sizes are supported. Preserves display aspect ratio without
	// cropping or upscaling.
	Size int64 `json:"size"`
	// Timestamps specifies 1 to 100 Source-relative target times in seconds. Times
	// outside the selected range are skipped when the action runs. If none remain, the
	// action is skipped without a charge. Repeated targets produce separate outputs.
	Timestamps []float64 `json:"timestamps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Format      respjson.Field
		Metadata    respjson.Field
		Output      respjson.Field
		Range       respjson.Field
		Size        respjson.Field
		Timestamps  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateThumbnailsParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateThumbnailsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type WorkflowRuleActionCreateThumbnailsParamsRange struct {
	// End specifies the exclusive upper bound in seconds. Defaults to the Source
	// duration.
	End float64 `json:"end"`
	// Start specifies the inclusive lower bound in seconds. Defaults to zero.
	Start float64 `json:"start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateThumbnailsParamsRange) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateThumbnailsParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses the original-language transcript. Can run alongside caption
// generation in the same asset.ready rule.
type WorkflowRuleActionCreateTranscript struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies transcript settings. Pass an empty object.
	Params any `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_transcript.
	Type constant.CreateTranscript `json:"type" default:"create_transcript"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleActionCreateTranscript) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleActionCreateTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Origin specifies the action whose ready output starts this rule. Required for
// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
// create_caption action. References must point to a compatible action in this
// definition and must not form a cycle.
type WorkflowRuleOrigin struct {
	// ActionId specifies the ID of the producing action within that rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ActionID string `json:"action_id" api:"required"`
	// RuleId specifies the ID of the rule containing the producing action. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	RuleID string `json:"rule_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		RuleID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRuleOrigin) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRuleOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewParams struct {
	// Name specifies the Workflow name. Use 1 to 128 characters.
	Name string `json:"name" api:"required"`
	// Rules specifies the event subscriptions and actions in this definition. Provide
	// exactly one asset.ready rule and up to 15 connected rules, with at most 32
	// actions and 65536 encoded bytes across all rules. Array order affects display
	// only.
	Rules []WorkflowNewParamsRule `json:"rules,omitzero" api:"required"`
	// SchemaVersion specifies the Workflow definition format. The supported version
	// is 1.
	//
	// Any of 1.
	SchemaVersion int64 `json:"schema_version,omitzero" api:"required"`
	paramObj
}

func (r WorkflowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A starting or connected step with one or more independent actions.
//
// The properties ID, Actions, Event are required.
type WorkflowNewParamsRule struct {
	// Id specifies a stable rule key, unique within this Workflow definition. Use 1 to
	// 64 characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Actions specifies the work to start when the rule matches. Actions can run
	// concurrently; array order does not define execution order. Provide 1 to 32
	// actions per rule, with at most 32 across the definition.
	Actions []WorkflowNewParamsRuleActionUnion `json:"actions,omitzero" api:"required"`
	// Event specifies when this rule starts. Use asset.ready for the single starting
	// rule. Output ready events require an origin caption or static rendition action
	// in this Workflow.
	//
	// Any of "asset.ready", "asset.output.ready".
	Event string `json:"event,omitzero" api:"required"`
	// Origin specifies the action whose ready output starts this rule. Required for
	// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
	// create_caption action. References must point to a compatible action in this
	// definition and must not form a cycle.
	Origin WorkflowNewParamsRuleOrigin `json:"origin,omitzero"`
	paramObj
}

func (r WorkflowNewParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsRule](
		"event", "asset.ready", "asset.output.ready",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowNewParamsRuleActionUnion struct {
	OfCreateStaticRendition *WorkflowNewParamsRuleActionCreateStaticRendition `json:",omitzero,inline"`
	OfCreateCaption         *WorkflowNewParamsRuleActionCreateCaption         `json:",omitzero,inline"`
	OfTranslateCaption      *WorkflowNewParamsRuleActionTranslateCaption      `json:",omitzero,inline"`
	OfCreateThumbnails      *WorkflowNewParamsRuleActionCreateThumbnails      `json:",omitzero,inline"`
	OfCreateTranscript      *WorkflowNewParamsRuleActionCreateTranscript      `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsRuleActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateStaticRendition,
		u.OfCreateCaption,
		u.OfTranslateCaption,
		u.OfCreateThumbnails,
		u.OfCreateTranscript)
}
func (u *WorkflowNewParamsRuleActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowNewParamsRuleActionUnion) asAny() any {
	if !param.IsOmitted(u.OfCreateStaticRendition) {
		return u.OfCreateStaticRendition
	} else if !param.IsOmitted(u.OfCreateCaption) {
		return u.OfCreateCaption
	} else if !param.IsOmitted(u.OfTranslateCaption) {
		return u.OfTranslateCaption
	} else if !param.IsOmitted(u.OfCreateThumbnails) {
		return u.OfCreateThumbnails
	} else if !param.IsOmitted(u.OfCreateTranscript) {
		return u.OfCreateTranscript
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowNewParamsRuleActionUnion) GetTranscriptID() *WorkflowNewParamsRuleActionCreateCaptionTranscriptID {
	if vt := u.OfCreateCaption; vt != nil {
		return &vt.TranscriptID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowNewParamsRuleActionUnion) GetCaptionID() *WorkflowNewParamsRuleActionTranslateCaptionCaptionID {
	if vt := u.OfTranslateCaption; vt != nil {
		return &vt.CaptionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowNewParamsRuleActionUnion) GetID() *string {
	if vt := u.OfCreateStaticRendition; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCreateCaption; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfTranslateCaption; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCreateThumbnails; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCreateTranscript; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowNewParamsRuleActionUnion) GetType() *string {
	if vt := u.OfCreateStaticRendition; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCreateCaption; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTranslateCaption; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCreateThumbnails; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCreateTranscript; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowNewParamsRuleActionUnion) GetParams() (res workflowNewParamsRuleActionUnionParams) {
	if vt := u.OfCreateStaticRendition; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfCreateCaption; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfTranslateCaption; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfCreateThumbnails; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfCreateTranscript; vt != nil {
		res.any = &vt.Params
	}
	return
}

// Can have the runtime types
// [*WorkflowNewParamsRuleActionCreateStaticRenditionParams], [*any],
// [*WorkflowNewParamsRuleActionTranslateCaptionParams],
// [*WorkflowNewParamsRuleActionCreateThumbnailsParams]
type workflowNewParamsRuleActionUnionParams struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *chunkify.WorkflowNewParamsRuleActionCreateStaticRenditionParams:
//	case *any:
//	case *chunkify.WorkflowNewParamsRuleActionTranslateCaptionParams:
//	case *chunkify.WorkflowNewParamsRuleActionCreateThumbnailsParams:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowNewParamsRuleActionUnionParams) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[WorkflowNewParamsRuleActionUnion](
		"type",
		apijson.Discriminator[WorkflowNewParamsRuleActionCreateStaticRendition]("create_static_rendition"),
		apijson.Discriminator[WorkflowNewParamsRuleActionCreateCaption]("create_caption"),
		apijson.Discriminator[WorkflowNewParamsRuleActionTranslateCaption]("translate_caption"),
		apijson.Discriminator[WorkflowNewParamsRuleActionCreateThumbnails]("create_thumbnails"),
		apijson.Discriminator[WorkflowNewParamsRuleActionCreateTranscript]("create_transcript"),
	)
}

// Creates or reuses one static rendition for the Asset.
//
// The properties ID, Params, Type are required.
type WorkflowNewParamsRuleActionCreateStaticRendition struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies the static rendition resolution.
	Params WorkflowNewParamsRuleActionCreateStaticRenditionParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_static_rendition.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_static_rendition".
	Type constant.CreateStaticRendition `json:"type" default:"create_static_rendition"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateStaticRendition) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateStaticRendition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateStaticRendition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the static rendition resolution.
//
// The property Resolution is required.
type WorkflowNewParamsRuleActionCreateStaticRenditionParams struct {
	// Resolution specifies the H.264 MP4 output resolution. highest selects the
	// largest supported output without upscaling. Hobby supports up to 1080p; other
	// plans support up to 2160p. Streaming resolution limits do not apply.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution string `json:"resolution,omitzero" api:"required"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateStaticRenditionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateStaticRenditionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateStaticRenditionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsRuleActionCreateStaticRenditionParams](
		"resolution", "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p", "2160p",
	)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
//
// The properties ID, Params, Type are required.
type WorkflowNewParamsRuleActionCreateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params any `json:"params,omitzero" api:"required"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowNewParamsRuleActionCreateCaptionTranscriptID `json:"transcript_id,omitzero"`
	// Type specifies the action to perform. Must be create_caption.
	//
	// This field can be elided, and will marshal its zero value as "create_caption".
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateCaption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateCaption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TranscriptId binds caption publication to the transcript in the triggering
// event. Use it only when connected to create_transcript. When omitted, creates or
// reuses the Asset’s original transcript.
//
// The property From is required.
type WorkflowNewParamsRuleActionCreateCaptionTranscriptID struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from,omitzero" api:"required"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateCaptionTranscriptID) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateCaptionTranscriptID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateCaptionTranscriptID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsRuleActionCreateCaptionTranscriptID](
		"from", "event.output.id",
	)
}

// Translates or reuses an original caption in one target language. Requires an
// asset.output.ready rule connected to an original caption creation action.
//
// The properties ID, CaptionID, Params, Type are required.
type WorkflowNewParamsRuleActionTranslateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// CaptionId specifies the binding to the original caption in the triggering event.
	CaptionID WorkflowNewParamsRuleActionTranslateCaptionCaptionID `json:"caption_id,omitzero" api:"required"`
	// Params specifies the target language.
	Params WorkflowNewParamsRuleActionTranslateCaptionParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be translate_caption.
	//
	// This field can be elided, and will marshal its zero value as
	// "translate_caption".
	Type constant.TranslateCaption `json:"type" default:"translate_caption"`
	paramObj
}

func (r WorkflowNewParamsRuleActionTranslateCaption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionTranslateCaption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionTranslateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CaptionId specifies the binding to the original caption in the triggering event.
//
// The property From is required.
type WorkflowNewParamsRuleActionTranslateCaptionCaptionID struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from,omitzero" api:"required"`
	paramObj
}

func (r WorkflowNewParamsRuleActionTranslateCaptionCaptionID) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionTranslateCaptionCaptionID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionTranslateCaptionCaptionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsRuleActionTranslateCaptionCaptionID](
		"from", "event.output.id",
	)
}

// Params specifies the target language.
//
// The property LanguageCode is required.
type WorkflowNewParamsRuleActionTranslateCaptionParams struct {
	// LanguageCode specifies the target language for the translated subtitle track,
	// such as en or fr. The same code appears in the caption and HLS manifest.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode string `json:"language_code,omitzero" api:"required"`
	paramObj
}

func (r WorkflowNewParamsRuleActionTranslateCaptionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionTranslateCaptionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionTranslateCaptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsRuleActionTranslateCaptionParams](
		"language_code", "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi", "ru", "tr", "pl", "sv",
	)
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
//
// The properties ID, Params, Type are required.
type WorkflowNewParamsRuleActionCreateThumbnails struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowNewParamsRuleActionCreateThumbnailsParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_thumbnails.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_thumbnails".
	Type constant.CreateThumbnails `json:"type" default:"create_thumbnails"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateThumbnails) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateThumbnails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateThumbnails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies sampling and image settings. Outputs use the Asset folder and
// Storage.
type WorkflowNewParamsRuleActionCreateThumbnailsParams struct {
	// Count specifies 1 to 100 thumbnails at the middle of equal sections of the
	// selected range.
	Count param.Opt[int64] `json:"count,omitzero"`
	// Size specifies the maximum image or tile edge in pixels, from 64 to 1920.
	// Defaults to 320. Odd sizes are supported. Preserves display aspect ratio without
	// cropping or upscaling.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Format selects JPEG or WebP. Defaults to JPEG.
	//
	// Any of "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output string `json:"output,omitzero"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range WorkflowNewParamsRuleActionCreateThumbnailsParamsRange `json:"range,omitzero"`
	// Timestamps specifies 1 to 100 Source-relative target times in seconds. Times
	// outside the selected range are skipped when the action runs. If none remain, the
	// action is skipped without a charge. Repeated targets produce separate outputs.
	Timestamps []float64 `json:"timestamps,omitzero"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateThumbnailsParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateThumbnailsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateThumbnailsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsRuleActionCreateThumbnailsParams](
		"format", "jpeg", "webp",
	)
	apijson.RegisterFieldValidator[WorkflowNewParamsRuleActionCreateThumbnailsParams](
		"output", "images", "sprites",
	)
}

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type WorkflowNewParamsRuleActionCreateThumbnailsParamsRange struct {
	// End specifies the exclusive upper bound in seconds. Defaults to the Source
	// duration.
	End param.Opt[float64] `json:"end,omitzero"`
	// Start specifies the inclusive lower bound in seconds. Defaults to zero.
	Start param.Opt[float64] `json:"start,omitzero"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateThumbnailsParamsRange) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateThumbnailsParamsRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateThumbnailsParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses the original-language transcript. Can run alongside caption
// generation in the same asset.ready rule.
//
// The properties ID, Params, Type are required.
type WorkflowNewParamsRuleActionCreateTranscript struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies transcript settings. Pass an empty object.
	Params any `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_transcript.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_transcript".
	Type constant.CreateTranscript `json:"type" default:"create_transcript"`
	paramObj
}

func (r WorkflowNewParamsRuleActionCreateTranscript) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleActionCreateTranscript
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleActionCreateTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Origin specifies the action whose ready output starts this rule. Required for
// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
// create_caption action. References must point to a compatible action in this
// definition and must not form a cycle.
//
// The properties ActionID, RuleID are required.
type WorkflowNewParamsRuleOrigin struct {
	// ActionId specifies the ID of the producing action within that rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ActionID string `json:"action_id" api:"required"`
	// RuleId specifies the ID of the rule containing the producing action. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	RuleID string `json:"rule_id" api:"required"`
	paramObj
}

func (r WorkflowNewParamsRuleOrigin) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsRuleOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsRuleOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing one Workflow.
type WorkflowNewResponseEnvelope struct {
	// The requested Workflow.
	Data Workflow `json:"data" api:"required"`
	// Status of the API request. Always success for this response.
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
func (r WorkflowNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing one Workflow.
type WorkflowGetResponseEnvelope struct {
	// The requested Workflow.
	Data Workflow `json:"data" api:"required"`
	// Status of the API request. Always success for this response.
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
func (r WorkflowGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowUpdateParams struct {
	// ExpectedRevision specifies the current Workflow revision when replacing rules.
	// The update returns a conflict if it no longer matches. Required when rules is
	// provided. Must be at least 1.
	ExpectedRevision param.Opt[int64] `json:"expected_revision,omitzero"`
	// Name specifies the Workflow name. Use 1 to 128 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// Rules specifies the event subscriptions and actions in this definition. Provide
	// exactly one asset.ready rule and up to 15 connected rules, with at most 32
	// actions and 65536 encoded bytes across all rules. Array order affects display
	// only. Replaces the complete definition and requires schema_version and
	// expected_revision. Existing Assets and Uploads keep their pinned revision.
	Rules []WorkflowUpdateParamsRule `json:"rules,omitzero"`
	// SchemaVersion specifies the Workflow definition format. The supported version
	// is 1. Required when rules is provided.
	//
	// Any of 1.
	SchemaVersion int64 `json:"schema_version,omitzero"`
	paramObj
}

func (r WorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A starting or connected step with one or more independent actions.
//
// The properties ID, Actions, Event are required.
type WorkflowUpdateParamsRule struct {
	// Id specifies a stable rule key, unique within this Workflow definition. Use 1 to
	// 64 characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Actions specifies the work to start when the rule matches. Actions can run
	// concurrently; array order does not define execution order. Provide 1 to 32
	// actions per rule, with at most 32 across the definition.
	Actions []WorkflowUpdateParamsRuleActionUnion `json:"actions,omitzero" api:"required"`
	// Event specifies when this rule starts. Use asset.ready for the single starting
	// rule. Output ready events require an origin caption or static rendition action
	// in this Workflow.
	//
	// Any of "asset.ready", "asset.output.ready".
	Event string `json:"event,omitzero" api:"required"`
	// Origin specifies the action whose ready output starts this rule. Required for
	// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
	// create_caption action. References must point to a compatible action in this
	// definition and must not form a cycle.
	Origin WorkflowUpdateParamsRuleOrigin `json:"origin,omitzero"`
	paramObj
}

func (r WorkflowUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateParamsRule](
		"event", "asset.ready", "asset.output.ready",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowUpdateParamsRuleActionUnion struct {
	OfCreateStaticRendition *WorkflowUpdateParamsRuleActionCreateStaticRendition `json:",omitzero,inline"`
	OfCreateCaption         *WorkflowUpdateParamsRuleActionCreateCaption         `json:",omitzero,inline"`
	OfTranslateCaption      *WorkflowUpdateParamsRuleActionTranslateCaption      `json:",omitzero,inline"`
	OfCreateThumbnails      *WorkflowUpdateParamsRuleActionCreateThumbnails      `json:",omitzero,inline"`
	OfCreateTranscript      *WorkflowUpdateParamsRuleActionCreateTranscript      `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowUpdateParamsRuleActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateStaticRendition,
		u.OfCreateCaption,
		u.OfTranslateCaption,
		u.OfCreateThumbnails,
		u.OfCreateTranscript)
}
func (u *WorkflowUpdateParamsRuleActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowUpdateParamsRuleActionUnion) asAny() any {
	if !param.IsOmitted(u.OfCreateStaticRendition) {
		return u.OfCreateStaticRendition
	} else if !param.IsOmitted(u.OfCreateCaption) {
		return u.OfCreateCaption
	} else if !param.IsOmitted(u.OfTranslateCaption) {
		return u.OfTranslateCaption
	} else if !param.IsOmitted(u.OfCreateThumbnails) {
		return u.OfCreateThumbnails
	} else if !param.IsOmitted(u.OfCreateTranscript) {
		return u.OfCreateTranscript
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowUpdateParamsRuleActionUnion) GetTranscriptID() *WorkflowUpdateParamsRuleActionCreateCaptionTranscriptID {
	if vt := u.OfCreateCaption; vt != nil {
		return &vt.TranscriptID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowUpdateParamsRuleActionUnion) GetCaptionID() *WorkflowUpdateParamsRuleActionTranslateCaptionCaptionID {
	if vt := u.OfTranslateCaption; vt != nil {
		return &vt.CaptionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowUpdateParamsRuleActionUnion) GetID() *string {
	if vt := u.OfCreateStaticRendition; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCreateCaption; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfTranslateCaption; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCreateThumbnails; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfCreateTranscript; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowUpdateParamsRuleActionUnion) GetType() *string {
	if vt := u.OfCreateStaticRendition; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCreateCaption; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTranslateCaption; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCreateThumbnails; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCreateTranscript; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowUpdateParamsRuleActionUnion) GetParams() (res workflowUpdateParamsRuleActionUnionParams) {
	if vt := u.OfCreateStaticRendition; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfCreateCaption; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfTranslateCaption; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfCreateThumbnails; vt != nil {
		res.any = &vt.Params
	} else if vt := u.OfCreateTranscript; vt != nil {
		res.any = &vt.Params
	}
	return
}

// Can have the runtime types
// [*WorkflowUpdateParamsRuleActionCreateStaticRenditionParams], [*any],
// [*WorkflowUpdateParamsRuleActionTranslateCaptionParams],
// [*WorkflowUpdateParamsRuleActionCreateThumbnailsParams]
type workflowUpdateParamsRuleActionUnionParams struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *chunkify.WorkflowUpdateParamsRuleActionCreateStaticRenditionParams:
//	case *any:
//	case *chunkify.WorkflowUpdateParamsRuleActionTranslateCaptionParams:
//	case *chunkify.WorkflowUpdateParamsRuleActionCreateThumbnailsParams:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowUpdateParamsRuleActionUnionParams) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[WorkflowUpdateParamsRuleActionUnion](
		"type",
		apijson.Discriminator[WorkflowUpdateParamsRuleActionCreateStaticRendition]("create_static_rendition"),
		apijson.Discriminator[WorkflowUpdateParamsRuleActionCreateCaption]("create_caption"),
		apijson.Discriminator[WorkflowUpdateParamsRuleActionTranslateCaption]("translate_caption"),
		apijson.Discriminator[WorkflowUpdateParamsRuleActionCreateThumbnails]("create_thumbnails"),
		apijson.Discriminator[WorkflowUpdateParamsRuleActionCreateTranscript]("create_transcript"),
	)
}

// Creates or reuses one static rendition for the Asset.
//
// The properties ID, Params, Type are required.
type WorkflowUpdateParamsRuleActionCreateStaticRendition struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies the static rendition resolution.
	Params WorkflowUpdateParamsRuleActionCreateStaticRenditionParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_static_rendition.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_static_rendition".
	Type constant.CreateStaticRendition `json:"type" default:"create_static_rendition"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateStaticRendition) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateStaticRendition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateStaticRendition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the static rendition resolution.
//
// The property Resolution is required.
type WorkflowUpdateParamsRuleActionCreateStaticRenditionParams struct {
	// Resolution specifies the H.264 MP4 output resolution. highest selects the
	// largest supported output without upscaling. Hobby supports up to 1080p; other
	// plans support up to 2160p. Streaming resolution limits do not apply.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution string `json:"resolution,omitzero" api:"required"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateStaticRenditionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateStaticRenditionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateStaticRenditionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateParamsRuleActionCreateStaticRenditionParams](
		"resolution", "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p", "2160p",
	)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
//
// The properties ID, Params, Type are required.
type WorkflowUpdateParamsRuleActionCreateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params any `json:"params,omitzero" api:"required"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowUpdateParamsRuleActionCreateCaptionTranscriptID `json:"transcript_id,omitzero"`
	// Type specifies the action to perform. Must be create_caption.
	//
	// This field can be elided, and will marshal its zero value as "create_caption".
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateCaption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateCaption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TranscriptId binds caption publication to the transcript in the triggering
// event. Use it only when connected to create_transcript. When omitted, creates or
// reuses the Asset’s original transcript.
//
// The property From is required.
type WorkflowUpdateParamsRuleActionCreateCaptionTranscriptID struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from,omitzero" api:"required"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateCaptionTranscriptID) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateCaptionTranscriptID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateCaptionTranscriptID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateParamsRuleActionCreateCaptionTranscriptID](
		"from", "event.output.id",
	)
}

// Translates or reuses an original caption in one target language. Requires an
// asset.output.ready rule connected to an original caption creation action.
//
// The properties ID, CaptionID, Params, Type are required.
type WorkflowUpdateParamsRuleActionTranslateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// CaptionId specifies the binding to the original caption in the triggering event.
	CaptionID WorkflowUpdateParamsRuleActionTranslateCaptionCaptionID `json:"caption_id,omitzero" api:"required"`
	// Params specifies the target language.
	Params WorkflowUpdateParamsRuleActionTranslateCaptionParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be translate_caption.
	//
	// This field can be elided, and will marshal its zero value as
	// "translate_caption".
	Type constant.TranslateCaption `json:"type" default:"translate_caption"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionTranslateCaption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionTranslateCaption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionTranslateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CaptionId specifies the binding to the original caption in the triggering event.
//
// The property From is required.
type WorkflowUpdateParamsRuleActionTranslateCaptionCaptionID struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from,omitzero" api:"required"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionTranslateCaptionCaptionID) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionTranslateCaptionCaptionID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionTranslateCaptionCaptionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateParamsRuleActionTranslateCaptionCaptionID](
		"from", "event.output.id",
	)
}

// Params specifies the target language.
//
// The property LanguageCode is required.
type WorkflowUpdateParamsRuleActionTranslateCaptionParams struct {
	// LanguageCode specifies the target language for the translated subtitle track,
	// such as en or fr. The same code appears in the caption and HLS manifest.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode string `json:"language_code,omitzero" api:"required"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionTranslateCaptionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionTranslateCaptionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionTranslateCaptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateParamsRuleActionTranslateCaptionParams](
		"language_code", "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi", "ru", "tr", "pl", "sv",
	)
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
//
// The properties ID, Params, Type are required.
type WorkflowUpdateParamsRuleActionCreateThumbnails struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowUpdateParamsRuleActionCreateThumbnailsParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_thumbnails.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_thumbnails".
	Type constant.CreateThumbnails `json:"type" default:"create_thumbnails"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateThumbnails) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateThumbnails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateThumbnails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies sampling and image settings. Outputs use the Asset folder and
// Storage.
type WorkflowUpdateParamsRuleActionCreateThumbnailsParams struct {
	// Count specifies 1 to 100 thumbnails at the middle of equal sections of the
	// selected range.
	Count param.Opt[int64] `json:"count,omitzero"`
	// Size specifies the maximum image or tile edge in pixels, from 64 to 1920.
	// Defaults to 320. Odd sizes are supported. Preserves display aspect ratio without
	// cropping or upscaling.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Format selects JPEG or WebP. Defaults to JPEG.
	//
	// Any of "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output string `json:"output,omitzero"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range WorkflowUpdateParamsRuleActionCreateThumbnailsParamsRange `json:"range,omitzero"`
	// Timestamps specifies 1 to 100 Source-relative target times in seconds. Times
	// outside the selected range are skipped when the action runs. If none remain, the
	// action is skipped without a charge. Repeated targets produce separate outputs.
	Timestamps []float64 `json:"timestamps,omitzero"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateThumbnailsParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateThumbnailsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateThumbnailsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateParamsRuleActionCreateThumbnailsParams](
		"format", "jpeg", "webp",
	)
	apijson.RegisterFieldValidator[WorkflowUpdateParamsRuleActionCreateThumbnailsParams](
		"output", "images", "sprites",
	)
}

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type WorkflowUpdateParamsRuleActionCreateThumbnailsParamsRange struct {
	// End specifies the exclusive upper bound in seconds. Defaults to the Source
	// duration.
	End param.Opt[float64] `json:"end,omitzero"`
	// Start specifies the inclusive lower bound in seconds. Defaults to zero.
	Start param.Opt[float64] `json:"start,omitzero"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateThumbnailsParamsRange) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateThumbnailsParamsRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateThumbnailsParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses the original-language transcript. Can run alongside caption
// generation in the same asset.ready rule.
//
// The properties ID, Params, Type are required.
type WorkflowUpdateParamsRuleActionCreateTranscript struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies transcript settings. Pass an empty object.
	Params any `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_transcript.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_transcript".
	Type constant.CreateTranscript `json:"type" default:"create_transcript"`
	paramObj
}

func (r WorkflowUpdateParamsRuleActionCreateTranscript) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleActionCreateTranscript
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleActionCreateTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Origin specifies the action whose ready output starts this rule. Required for
// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
// create_caption action. References must point to a compatible action in this
// definition and must not form a cycle.
//
// The properties ActionID, RuleID are required.
type WorkflowUpdateParamsRuleOrigin struct {
	// ActionId specifies the ID of the producing action within that rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ActionID string `json:"action_id" api:"required"`
	// RuleId specifies the ID of the rule containing the producing action. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	RuleID string `json:"rule_id" api:"required"`
	paramObj
}

func (r WorkflowUpdateParamsRuleOrigin) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsRuleOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsRuleOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowListParams struct {
	// Archived specifies whether to include archived Workflows. Defaults to false.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// Limit specifies the maximum number of results to return. Accepts 1 to 100.
	// Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset specifies the number of matching results to skip. Must be at least 0.
	// Defaults to 0.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
