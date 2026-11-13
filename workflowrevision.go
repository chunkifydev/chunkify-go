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
	"github.com/chunkifydev/chunkify-go/packages/respjson"
	"github.com/chunkifydev/chunkify-go/shared/constant"
)

// WorkflowRevisionService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowRevisionService] method instead.
type WorkflowRevisionService struct {
	Options []option.RequestOption
}

// NewWorkflowRevisionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowRevisionService(opts ...option.RequestOption) (r WorkflowRevisionService) {
	r = WorkflowRevisionService{}
	r.Options = opts
	return
}

// Retrieve an immutable definition revision, including revisions of archived
// Workflows.
func (r *WorkflowRevisionService) Get(ctx context.Context, revision int64, query WorkflowRevisionGetParams, opts ...option.RequestOption) (res *WorkflowRevision, err error) {
	var env WorkflowRevisionGetResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.WorkflowID == "" {
		err = errors.New("missing required workflowId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workflows/%s/revisions/%v", query.WorkflowID, revision)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// An immutable Workflow definition pinned to Assets and Uploads at creation.
type WorkflowRevision struct {
	// When this Workflow revision was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Definition revision number, starting at 1.
	Revision int64 `json:"revision" api:"required"`
	// Rules specifies the event subscriptions and actions in this definition. Provide
	// exactly one asset.ready rule and up to 15 connected rules, with at most 32
	// actions and 65536 encoded bytes across all rules. Array order affects display
	// only.
	Rules []WorkflowRevisionRule `json:"rules" api:"required"`
	// SchemaVersion specifies the Workflow definition format. The supported version
	// is 1.
	SchemaVersion int64 `json:"schema_version" api:"required"`
	// ID of the Workflow that owns this revision.
	WorkflowID string `json:"workflow_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		Revision      respjson.Field
		Rules         respjson.Field
		SchemaVersion respjson.Field
		WorkflowID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRevision) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevision) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A starting or connected step with one or more independent actions.
type WorkflowRevisionRule struct {
	// Id specifies a stable rule key, unique within this Workflow definition. Use 1 to
	// 64 characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Actions specifies the work to start when the rule matches. Actions can run
	// concurrently; array order does not define execution order. Provide 1 to 32
	// actions per rule, with at most 32 across the definition.
	Actions []WorkflowRevisionRuleActionUnion `json:"actions" api:"required"`
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
	Origin WorkflowRevisionRuleOrigin `json:"origin"`
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
func (r WorkflowRevisionRule) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowRevisionRuleActionUnion contains all possible properties and values from
// [WorkflowRevisionRuleActionCreateStaticRendition],
// [WorkflowRevisionRuleActionCreateCaption],
// [WorkflowRevisionRuleActionTranslateCaption],
// [WorkflowRevisionRuleActionCreateThumbnails],
// [WorkflowRevisionRuleActionCreateTranscript].
//
// Use the [WorkflowRevisionRuleActionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowRevisionRuleActionUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [WorkflowRevisionRuleActionCreateStaticRenditionParams],
	// [WorkflowRevisionRuleActionCreateCaptionParams],
	// [WorkflowRevisionRuleActionTranslateCaptionParams],
	// [WorkflowRevisionRuleActionCreateThumbnailsParams],
	// [WorkflowRevisionRuleActionCreateTranscriptParams]
	Params WorkflowRevisionRuleActionUnionParams `json:"params"`
	// Any of "create_static_rendition", "create_caption", "translate_caption",
	// "create_thumbnails", "create_transcript".
	Type string `json:"type"`
	// This field is from variant [WorkflowRevisionRuleActionCreateCaption].
	TranscriptID WorkflowRevisionRuleActionCreateCaptionTranscriptID `json:"transcript_id"`
	// This field is from variant [WorkflowRevisionRuleActionTranslateCaption].
	CaptionID WorkflowRevisionRuleActionTranslateCaptionCaptionID `json:"caption_id"`
	JSON      struct {
		ID           respjson.Field
		Params       respjson.Field
		Type         respjson.Field
		TranscriptID respjson.Field
		CaptionID    respjson.Field
		raw          string
	} `json:"-"`
}

// anyWorkflowRevisionRuleAction is implemented by each variant of
// [WorkflowRevisionRuleActionUnion] to add type safety for the return type of
// [WorkflowRevisionRuleActionUnion.AsAny]
type anyWorkflowRevisionRuleAction interface {
	implWorkflowRevisionRuleActionUnion()
}

func (WorkflowRevisionRuleActionCreateStaticRendition) implWorkflowRevisionRuleActionUnion() {}
func (WorkflowRevisionRuleActionCreateCaption) implWorkflowRevisionRuleActionUnion()         {}
func (WorkflowRevisionRuleActionTranslateCaption) implWorkflowRevisionRuleActionUnion()      {}
func (WorkflowRevisionRuleActionCreateThumbnails) implWorkflowRevisionRuleActionUnion()      {}
func (WorkflowRevisionRuleActionCreateTranscript) implWorkflowRevisionRuleActionUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WorkflowRevisionRuleActionUnion.AsAny().(type) {
//	case chunkify.WorkflowRevisionRuleActionCreateStaticRendition:
//	case chunkify.WorkflowRevisionRuleActionCreateCaption:
//	case chunkify.WorkflowRevisionRuleActionTranslateCaption:
//	case chunkify.WorkflowRevisionRuleActionCreateThumbnails:
//	case chunkify.WorkflowRevisionRuleActionCreateTranscript:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WorkflowRevisionRuleActionUnion) AsAny() anyWorkflowRevisionRuleAction {
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

func (u WorkflowRevisionRuleActionUnion) AsCreateStaticRendition() (v WorkflowRevisionRuleActionCreateStaticRendition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRevisionRuleActionUnion) AsCreateCaption() (v WorkflowRevisionRuleActionCreateCaption) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRevisionRuleActionUnion) AsTranslateCaption() (v WorkflowRevisionRuleActionTranslateCaption) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRevisionRuleActionUnion) AsCreateThumbnails() (v WorkflowRevisionRuleActionCreateThumbnails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowRevisionRuleActionUnion) AsCreateTranscript() (v WorkflowRevisionRuleActionCreateTranscript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowRevisionRuleActionUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowRevisionRuleActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowRevisionRuleActionUnionParams is an implicit subunion of
// [WorkflowRevisionRuleActionUnion]. WorkflowRevisionRuleActionUnionParams
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowRevisionRuleActionUnion].
type WorkflowRevisionRuleActionUnionParams struct {
	// This field is from variant
	// [WorkflowRevisionRuleActionCreateStaticRenditionParams].
	Resolution string `json:"resolution"`
	// This field is from variant [WorkflowRevisionRuleActionTranslateCaptionParams].
	LanguageCode string `json:"language_code"`
	// This field is from variant [WorkflowRevisionRuleActionCreateThumbnailsParams].
	Count int64 `json:"count"`
	// This field is from variant [WorkflowRevisionRuleActionCreateThumbnailsParams].
	Format string `json:"format"`
	// This field is from variant [WorkflowRevisionRuleActionCreateThumbnailsParams].
	Metadata map[string]string `json:"metadata"`
	// This field is from variant [WorkflowRevisionRuleActionCreateThumbnailsParams].
	Output string `json:"output"`
	// This field is from variant [WorkflowRevisionRuleActionCreateThumbnailsParams].
	Range WorkflowRevisionRuleActionCreateThumbnailsParamsRange `json:"range"`
	// This field is from variant [WorkflowRevisionRuleActionCreateThumbnailsParams].
	Size int64 `json:"size"`
	// This field is from variant [WorkflowRevisionRuleActionCreateThumbnailsParams].
	Timestamps []float64 `json:"timestamps"`
	JSON       struct {
		Resolution   respjson.Field
		LanguageCode respjson.Field
		Count        respjson.Field
		Format       respjson.Field
		Metadata     respjson.Field
		Output       respjson.Field
		Range        respjson.Field
		Size         respjson.Field
		Timestamps   respjson.Field
		raw          string
	} `json:"-"`
}

func (r *WorkflowRevisionRuleActionUnionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses one static rendition for the Asset.
type WorkflowRevisionRuleActionCreateStaticRendition struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies the static rendition resolution.
	Params WorkflowRevisionRuleActionCreateStaticRenditionParams `json:"params" api:"required"`
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
func (r WorkflowRevisionRuleActionCreateStaticRendition) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateStaticRendition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the static rendition resolution.
type WorkflowRevisionRuleActionCreateStaticRenditionParams struct {
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
func (r WorkflowRevisionRuleActionCreateStaticRenditionParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateStaticRenditionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
type WorkflowRevisionRuleActionCreateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params WorkflowRevisionRuleActionCreateCaptionParams `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_caption.
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowRevisionRuleActionCreateCaptionTranscriptID `json:"transcript_id"`
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
func (r WorkflowRevisionRuleActionCreateCaption) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies caption generation settings. Pass an empty object.
type WorkflowRevisionRuleActionCreateCaptionParams struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRevisionRuleActionCreateCaptionParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateCaptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TranscriptId binds caption publication to the transcript in the triggering
// event. Use it only when connected to create_transcript. When omitted, creates or
// reuses the Asset’s original transcript.
type WorkflowRevisionRuleActionCreateCaptionTranscriptID struct {
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
func (r WorkflowRevisionRuleActionCreateCaptionTranscriptID) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateCaptionTranscriptID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translates or reuses an original caption in one target language. Requires an
// asset.output.ready rule connected to an original caption creation action.
type WorkflowRevisionRuleActionTranslateCaption struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// CaptionId specifies the binding to the original caption in the triggering event.
	CaptionID WorkflowRevisionRuleActionTranslateCaptionCaptionID `json:"caption_id" api:"required"`
	// Params specifies the target language.
	Params WorkflowRevisionRuleActionTranslateCaptionParams `json:"params" api:"required"`
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
func (r WorkflowRevisionRuleActionTranslateCaption) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionTranslateCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CaptionId specifies the binding to the original caption in the triggering event.
type WorkflowRevisionRuleActionTranslateCaptionCaptionID struct {
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
func (r WorkflowRevisionRuleActionTranslateCaptionCaptionID) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionTranslateCaptionCaptionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies the target language.
type WorkflowRevisionRuleActionTranslateCaptionParams struct {
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
func (r WorkflowRevisionRuleActionTranslateCaptionParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionTranslateCaptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
type WorkflowRevisionRuleActionCreateThumbnails struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowRevisionRuleActionCreateThumbnailsParams `json:"params" api:"required"`
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
func (r WorkflowRevisionRuleActionCreateThumbnails) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateThumbnails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies sampling and image settings. Outputs use the Asset folder and
// Storage.
type WorkflowRevisionRuleActionCreateThumbnailsParams struct {
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
	Range WorkflowRevisionRuleActionCreateThumbnailsParamsRange `json:"range"`
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
func (r WorkflowRevisionRuleActionCreateThumbnailsParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateThumbnailsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type WorkflowRevisionRuleActionCreateThumbnailsParamsRange struct {
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
func (r WorkflowRevisionRuleActionCreateThumbnailsParamsRange) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateThumbnailsParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses the original-language transcript. Can run alongside caption
// generation in the same asset.ready rule.
type WorkflowRevisionRuleActionCreateTranscript struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies transcript settings. Pass an empty object.
	Params WorkflowRevisionRuleActionCreateTranscriptParams `json:"params" api:"required"`
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
func (r WorkflowRevisionRuleActionCreateTranscript) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Params specifies transcript settings. Pass an empty object.
type WorkflowRevisionRuleActionCreateTranscriptParams struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRevisionRuleActionCreateTranscriptParams) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleActionCreateTranscriptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Origin specifies the action whose ready output starts this rule. Required for
// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
// create_caption action. References must point to a compatible action in this
// definition and must not form a cycle.
type WorkflowRevisionRuleOrigin struct {
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
func (r WorkflowRevisionRuleOrigin) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionRuleOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRevisionGetParams struct {
	WorkflowID string `path:"workflowId" api:"required" json:"-"`
	paramObj
}

// Response containing one Workflow revision.
type WorkflowRevisionGetResponseEnvelope struct {
	// The requested Workflow revision.
	Data WorkflowRevision `json:"data" api:"required"`
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
func (r WorkflowRevisionGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRevisionGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
