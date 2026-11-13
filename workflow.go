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

// WorkflowActionUnion contains all possible properties and values from
// [WorkflowStaticAction], [WorkflowCaptionAction], [WorkflowTranslationAction],
// [WorkflowThumbnailAction], [WorkflowTranscriptAction].
//
// Use the [WorkflowActionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowActionUnion struct {
	ID string `json:"id"`
	// This field is a union of [WorkflowStatic], [WorkflowCaption],
	// [WorkflowTranslation], [WorkflowThumbnail], [WorkflowTranscript]
	Params WorkflowActionUnionParams `json:"params"`
	// Any of "create_static_rendition", "create_caption", "translate_caption",
	// "create_thumbnails", "create_transcript".
	Type string `json:"type"`
	// This field is from variant [WorkflowCaptionAction].
	TranscriptID WorkflowOutputBinding `json:"transcript_id"`
	// This field is from variant [WorkflowTranslationAction].
	CaptionID WorkflowOutputBinding `json:"caption_id"`
	JSON      struct {
		ID           respjson.Field
		Params       respjson.Field
		Type         respjson.Field
		TranscriptID respjson.Field
		CaptionID    respjson.Field
		raw          string
	} `json:"-"`
}

// anyWorkflowAction is implemented by each variant of [WorkflowActionUnion] to add
// type safety for the return type of [WorkflowActionUnion.AsAny]
type anyWorkflowAction interface {
	implWorkflowActionUnion()
}

func (WorkflowStaticAction) implWorkflowActionUnion()      {}
func (WorkflowCaptionAction) implWorkflowActionUnion()     {}
func (WorkflowTranslationAction) implWorkflowActionUnion() {}
func (WorkflowThumbnailAction) implWorkflowActionUnion()   {}
func (WorkflowTranscriptAction) implWorkflowActionUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WorkflowActionUnion.AsAny().(type) {
//	case chunkify.WorkflowStaticAction:
//	case chunkify.WorkflowCaptionAction:
//	case chunkify.WorkflowTranslationAction:
//	case chunkify.WorkflowThumbnailAction:
//	case chunkify.WorkflowTranscriptAction:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WorkflowActionUnion) AsAny() anyWorkflowAction {
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

func (u WorkflowActionUnion) AsCreateStaticRendition() (v WorkflowStaticAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowActionUnion) AsCreateCaption() (v WorkflowCaptionAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowActionUnion) AsTranslateCaption() (v WorkflowTranslationAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowActionUnion) AsCreateThumbnails() (v WorkflowThumbnailAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowActionUnion) AsCreateTranscript() (v WorkflowTranscriptAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowActionUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowActionUnionParams is an implicit subunion of [WorkflowActionUnion].
// WorkflowActionUnionParams provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowActionUnion].
type WorkflowActionUnionParams struct {
	// This field is from variant [WorkflowStatic].
	Resolution WorkflowStaticResolution `json:"resolution"`
	// This field is from variant [WorkflowTranslation].
	LanguageCode WorkflowTranslationLanguageCode `json:"language_code"`
	// This field is from variant [WorkflowThumbnail].
	Count int64 `json:"count"`
	// This field is from variant [WorkflowThumbnail].
	Format ThumbnailFormat `json:"format"`
	// This field is from variant [WorkflowThumbnail].
	Metadata map[string]string `json:"metadata"`
	// This field is from variant [WorkflowThumbnail].
	Output ThumbnailOutput `json:"output"`
	// This field is from variant [WorkflowThumbnail].
	Range WorkflowThumbnailRange `json:"range"`
	// This field is from variant [WorkflowThumbnail].
	Size int64 `json:"size"`
	// This field is from variant [WorkflowThumbnail].
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

func (r *WorkflowActionUnionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Counts of action outcomes in this run.
type WorkflowActionCounts struct {
	// Number of actions cancelled before completion.
	Cancelled int64 `json:"cancelled" api:"required"`
	// Number of actions that completed successfully, including reused outputs.
	Completed int64 `json:"completed" api:"required"`
	// Number of actions that failed.
	Failed int64 `json:"failed" api:"required"`
	// Number of completed actions that reused an existing output. Included in
	// completed.
	Reused int64 `json:"reused" api:"required"`
	// Number of actions skipped without producing an output.
	Skipped int64 `json:"skipped" api:"required"`
	// Total number of actions in the run, including unfinished actions.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		Completed   respjson.Field
		Failed      respjson.Field
		Reused      respjson.Field
		Skipped     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionCounts) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func WorkflowActionParamsOfCreateStaticRendition(id string, params WorkflowStaticParams) WorkflowActionParamsUnion {
	var createStaticRendition WorkflowStaticActionParams
	createStaticRendition.ID = id
	createStaticRendition.Params = params
	return WorkflowActionParamsUnion{OfCreateStaticRendition: &createStaticRendition}
}

func WorkflowActionParamsOfCreateCaption(id string, params WorkflowCaptionParams) WorkflowActionParamsUnion {
	var createCaption WorkflowCaptionActionParams
	createCaption.ID = id
	createCaption.Params = params
	return WorkflowActionParamsUnion{OfCreateCaption: &createCaption}
}

func WorkflowActionParamsOfTranslateCaption(captionID WorkflowOutputBindingParams, id string, params WorkflowTranslationParams) WorkflowActionParamsUnion {
	var translateCaption WorkflowTranslationActionParams
	translateCaption.CaptionID = captionID
	translateCaption.ID = id
	translateCaption.Params = params
	return WorkflowActionParamsUnion{OfTranslateCaption: &translateCaption}
}

func WorkflowActionParamsOfCreateThumbnails(id string, params WorkflowThumbnailParams) WorkflowActionParamsUnion {
	var createThumbnails WorkflowThumbnailActionParams
	createThumbnails.ID = id
	createThumbnails.Params = params
	return WorkflowActionParamsUnion{OfCreateThumbnails: &createThumbnails}
}

func WorkflowActionParamsOfCreateTranscript(id string, params WorkflowTranscriptParams) WorkflowActionParamsUnion {
	var createTranscript WorkflowTranscriptActionParams
	createTranscript.ID = id
	createTranscript.Params = params
	return WorkflowActionParamsUnion{OfCreateTranscript: &createTranscript}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowActionParamsUnion struct {
	OfCreateStaticRendition *WorkflowStaticActionParams      `json:",omitzero,inline"`
	OfCreateCaption         *WorkflowCaptionActionParams     `json:",omitzero,inline"`
	OfTranslateCaption      *WorkflowTranslationActionParams `json:",omitzero,inline"`
	OfCreateThumbnails      *WorkflowThumbnailActionParams   `json:",omitzero,inline"`
	OfCreateTranscript      *WorkflowTranscriptActionParams  `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowActionParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateStaticRendition,
		u.OfCreateCaption,
		u.OfTranslateCaption,
		u.OfCreateThumbnails,
		u.OfCreateTranscript)
}
func (u *WorkflowActionParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowActionParamsUnion) asAny() any {
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
func (u WorkflowActionParamsUnion) GetTranscriptID() *WorkflowOutputBindingParams {
	if vt := u.OfCreateCaption; vt != nil {
		return &vt.TranscriptID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowActionParamsUnion) GetCaptionID() *WorkflowOutputBindingParams {
	if vt := u.OfTranslateCaption; vt != nil {
		return &vt.CaptionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowActionParamsUnion) GetID() *string {
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
func (u WorkflowActionParamsUnion) GetType() *string {
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
func (u WorkflowActionParamsUnion) GetParams() (res workflowActionParamsUnionParams) {
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

// Can have the runtime types [*WorkflowStaticParams], [*WorkflowCaptionParams],
// [*WorkflowTranslationParams], [*WorkflowThumbnailParams],
// [*WorkflowTranscriptParams]
type workflowActionParamsUnionParams struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *chunkify.WorkflowStaticParams:
//	case *chunkify.WorkflowCaptionParams:
//	case *chunkify.WorkflowTranslationParams:
//	case *chunkify.WorkflowThumbnailParams:
//	case *chunkify.WorkflowTranscriptParams:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowActionParamsUnionParams) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[WorkflowActionParamsUnion](
		"type",
		apijson.Discriminator[WorkflowStaticActionParams]("create_static_rendition"),
		apijson.Discriminator[WorkflowCaptionActionParams]("create_caption"),
		apijson.Discriminator[WorkflowTranslationActionParams]("translate_caption"),
		apijson.Discriminator[WorkflowThumbnailActionParams]("create_thumbnails"),
		apijson.Discriminator[WorkflowTranscriptActionParams]("create_transcript"),
	)
}

// Parameters for generating original captions. Pass an empty object; the spoken
// language is detected automatically.
type WorkflowCaption struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowCaption) RawJSON() string { return r.JSON.raw }
func (r *WorkflowCaption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
type WorkflowCaptionAction struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params WorkflowCaption `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_caption.
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowOutputBinding `json:"transcript_id"`
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
func (r WorkflowCaptionAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowCaptionAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
//
// The properties ID, Params, Type are required.
type WorkflowCaptionActionParams struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params WorkflowCaptionParams `json:"params" api:"required"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowOutputBindingParams `json:"transcript_id,omitzero"`
	// Type specifies the action to perform. Must be create_caption.
	//
	// This field can be elided, and will marshal its zero value as "create_caption".
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	paramObj
}

func (r WorkflowCaptionActionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCaptionActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCaptionActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for generating original captions. Pass an empty object; the spoken
// language is detected automatically.
type WorkflowCaptionParams struct {
	paramObj
}

func (r WorkflowCaptionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCaptionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCaptionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connects this rule to the output created or reused by an action in the same
// Workflow revision.
type WorkflowOrigin struct {
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
func (r WorkflowOrigin) RawJSON() string { return r.JSON.raw }
func (r *WorkflowOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connects this rule to the output created or reused by an action in the same
// Workflow revision.
//
// The properties ActionID, RuleID are required.
type WorkflowOriginParams struct {
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

func (r WorkflowOriginParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowOriginParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowOriginParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Selects the output referenced by the triggering event.
type WorkflowOutputBinding struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From WorkflowOutputBindingFrom `json:"from" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowOutputBinding) RawJSON() string { return r.JSON.raw }
func (r *WorkflowOutputBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowOutputBindingFrom string

const (
	WorkflowOutputBindingFromEventOutputID WorkflowOutputBindingFrom = "event.output.id"
)

// Selects the output referenced by the triggering event.
//
// The property From is required.
type WorkflowOutputBindingParams struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From WorkflowOutputBindingFrom `json:"from,omitzero" api:"required"`
	paramObj
}

func (r WorkflowOutputBindingParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowOutputBindingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowOutputBindingParams) UnmarshalJSON(data []byte) error {
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
	Actions []WorkflowActionUnion `json:"actions" api:"required"`
	// Event specifies when this rule starts. Use asset.ready for the single starting
	// rule. Output ready events require an origin caption or static rendition action
	// in this Workflow.
	//
	// Any of "asset.ready", "asset.output.ready".
	Event WorkflowRuleEvent `json:"event" api:"required"`
	// Origin specifies the action whose ready output starts this rule. Required for
	// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
	// create_caption action. References must point to a compatible action in this
	// definition and must not form a cycle.
	Origin WorkflowOrigin `json:"origin"`
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

type WorkflowRuleEvent string

const (
	WorkflowRuleEventAssetReady       WorkflowRuleEvent = "asset.ready"
	WorkflowRuleEventAssetOutputReady WorkflowRuleEvent = "asset.output.ready"
)

// A starting or connected step with one or more independent actions.
//
// The properties ID, Actions, Event are required.
type WorkflowRuleParams struct {
	// Id specifies a stable rule key, unique within this Workflow definition. Use 1 to
	// 64 characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Actions specifies the work to start when the rule matches. Actions can run
	// concurrently; array order does not define execution order. Provide 1 to 32
	// actions per rule, with at most 32 across the definition.
	Actions []WorkflowActionParamsUnion `json:"actions,omitzero" api:"required"`
	// Event specifies when this rule starts. Use asset.ready for the single starting
	// rule. Output ready events require an origin caption or static rendition action
	// in this Workflow.
	//
	// Any of "asset.ready", "asset.output.ready".
	Event WorkflowRuleEvent `json:"event,omitzero" api:"required"`
	// Origin specifies the action whose ready output starts this rule. Required for
	// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
	// create_caption action. References must point to a compatible action in this
	// definition and must not form a cycle.
	Origin WorkflowOriginParams `json:"origin,omitzero"`
	paramObj
}

func (r WorkflowRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current run status. pending means its actions have not started; running means
// work remains. After all actions finish, failed takes precedence over cancelled.
// Otherwise the run is completed, including when actions were skipped.
type WorkflowRunStatus string

const (
	WorkflowRunStatusPending   WorkflowRunStatus = "pending"
	WorkflowRunStatusRunning   WorkflowRunStatus = "running"
	WorkflowRunStatusCompleted WorkflowRunStatus = "completed"
	WorkflowRunStatusFailed    WorkflowRunStatus = "failed"
	WorkflowRunStatusCancelled WorkflowRunStatus = "cancelled"
)

// Parameters for creating one static rendition.
type WorkflowStatic struct {
	// Resolution specifies the H.264 MP4 output resolution. highest selects the
	// largest supported output without upscaling. Hobby supports up to 1080p; other
	// plans support up to 2160p. Streaming resolution limits do not apply.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution WorkflowStaticResolution `json:"resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Resolution  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStatic) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses one static rendition for the Asset.
type WorkflowStaticAction struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies the static rendition resolution.
	Params WorkflowStatic `json:"params" api:"required"`
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
func (r WorkflowStaticAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStaticAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses one static rendition for the Asset.
//
// The properties ID, Params, Type are required.
type WorkflowStaticActionParams struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies the static rendition resolution.
	Params WorkflowStaticParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_static_rendition.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_static_rendition".
	Type constant.CreateStaticRendition `json:"type" default:"create_static_rendition"`
	paramObj
}

func (r WorkflowStaticActionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStaticActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStaticActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating one static rendition.
//
// The property Resolution is required.
type WorkflowStaticParams struct {
	// Resolution specifies the H.264 MP4 output resolution. highest selects the
	// largest supported output without upscaling. Hobby supports up to 1080p; other
	// plans support up to 2160p. Streaming resolution limits do not apply.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution WorkflowStaticResolution `json:"resolution,omitzero" api:"required"`
	paramObj
}

func (r WorkflowStaticParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStaticParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStaticParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowStaticResolution string

const (
	WorkflowStaticResolutionHighest WorkflowStaticResolution = "highest"
	WorkflowStaticResolution240p    WorkflowStaticResolution = "240p"
	WorkflowStaticResolution360p    WorkflowStaticResolution = "360p"
	WorkflowStaticResolution480p    WorkflowStaticResolution = "480p"
	WorkflowStaticResolution540p    WorkflowStaticResolution = "540p"
	WorkflowStaticResolution720p    WorkflowStaticResolution = "720p"
	WorkflowStaticResolution1080p   WorkflowStaticResolution = "1080p"
	WorkflowStaticResolution1440p   WorkflowStaticResolution = "1440p"
	WorkflowStaticResolution2160p   WorkflowStaticResolution = "2160p"
)

// Generate additional thumbnails in the Asset folder and Storage. Each layout,
// format and size combination can be requested once. Specify exactly one of
// timestamps or count.
type WorkflowThumbnail struct {
	// Count specifies 1 to 100 thumbnails at the middle of equal sections of the
	// selected range.
	Count int64 `json:"count"`
	// Format selects JPEG or WebP. Defaults to JPEG.
	//
	// Any of "jpeg", "webp".
	Format ThumbnailFormat `json:"format"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output ThumbnailOutput `json:"output"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range WorkflowThumbnailRange `json:"range"`
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
func (r WorkflowThumbnail) RawJSON() string { return r.JSON.raw }
func (r *WorkflowThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
type WorkflowThumbnailAction struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowThumbnail `json:"params" api:"required"`
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
func (r WorkflowThumbnailAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowThumbnailAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
//
// The properties ID, Params, Type are required.
type WorkflowThumbnailActionParams struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowThumbnailParams `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_thumbnails.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_thumbnails".
	Type constant.CreateThumbnails `json:"type" default:"create_thumbnails"`
	paramObj
}

func (r WorkflowThumbnailActionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThumbnailActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowThumbnailActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generate additional thumbnails in the Asset folder and Storage. Each layout,
// format and size combination can be requested once. Specify exactly one of
// timestamps or count.
type WorkflowThumbnailParams struct {
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
	Format ThumbnailFormat `json:"format,omitzero"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output ThumbnailOutput `json:"output,omitzero"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range WorkflowThumbnailRangeParams `json:"range,omitzero"`
	// Timestamps specifies 1 to 100 Source-relative target times in seconds. Times
	// outside the selected range are skipped when the action runs. If none remain, the
	// action is skipped without a charge. Repeated targets produce separate outputs.
	Timestamps []float64 `json:"timestamps,omitzero"`
	paramObj
}

func (r WorkflowThumbnailParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThumbnailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowThumbnailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional Source-relative sampling boundaries.
type WorkflowThumbnailRange struct {
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
func (r WorkflowThumbnailRange) RawJSON() string { return r.JSON.raw }
func (r *WorkflowThumbnailRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional Source-relative sampling boundaries.
type WorkflowThumbnailRangeParams struct {
	// End specifies the exclusive upper bound in seconds. Defaults to the Source
	// duration.
	End param.Opt[float64] `json:"end,omitzero"`
	// Start specifies the inclusive lower bound in seconds. Defaults to zero.
	Start param.Opt[float64] `json:"start,omitzero"`
	paramObj
}

func (r WorkflowThumbnailRangeParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThumbnailRangeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowThumbnailRangeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for an original-language transcript. Pass an empty object; the spoken
// language is detected automatically.
type WorkflowTranscript struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTranscript) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses the original-language transcript. Can run alongside caption
// generation in the same asset.ready rule.
type WorkflowTranscriptAction struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies transcript settings. Pass an empty object.
	Params WorkflowTranscript `json:"params" api:"required"`
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
func (r WorkflowTranscriptAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTranscriptAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses the original-language transcript. Can run alongside caption
// generation in the same asset.ready rule.
//
// The properties ID, Params, Type are required.
type WorkflowTranscriptActionParams struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies transcript settings. Pass an empty object.
	Params WorkflowTranscriptParams `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_transcript.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_transcript".
	Type constant.CreateTranscript `json:"type" default:"create_transcript"`
	paramObj
}

func (r WorkflowTranscriptActionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTranscriptActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTranscriptActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for an original-language transcript. Pass an empty object; the spoken
// language is detected automatically.
type WorkflowTranscriptParams struct {
	paramObj
}

func (r WorkflowTranscriptParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTranscriptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTranscriptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for translating one original caption into another language.
type WorkflowTranslation struct {
	// LanguageCode specifies the target language for the translated subtitle track,
	// such as en or fr. The same code appears in the caption and HLS manifest.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode WorkflowTranslationLanguageCode `json:"language_code" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTranslation) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translates or reuses an original caption in one target language. Requires an
// asset.output.ready rule connected to an original caption creation action.
type WorkflowTranslationAction struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// CaptionId specifies the binding to the original caption in the triggering event.
	CaptionID WorkflowOutputBinding `json:"caption_id" api:"required"`
	// Params specifies the target language.
	Params WorkflowTranslation `json:"params" api:"required"`
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
func (r WorkflowTranslationAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTranslationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translates or reuses an original caption in one target language. Requires an
// asset.output.ready rule connected to an original caption creation action.
//
// The properties ID, CaptionID, Params, Type are required.
type WorkflowTranslationActionParams struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// CaptionId specifies the binding to the original caption in the triggering event.
	CaptionID WorkflowOutputBindingParams `json:"caption_id,omitzero" api:"required"`
	// Params specifies the target language.
	Params WorkflowTranslationParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be translate_caption.
	//
	// This field can be elided, and will marshal its zero value as
	// "translate_caption".
	Type constant.TranslateCaption `json:"type" default:"translate_caption"`
	paramObj
}

func (r WorkflowTranslationActionParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTranslationActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTranslationActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTranslationLanguageCode string

const (
	WorkflowTranslationLanguageCodeEn WorkflowTranslationLanguageCode = "en"
	WorkflowTranslationLanguageCodeFr WorkflowTranslationLanguageCode = "fr"
	WorkflowTranslationLanguageCodeEs WorkflowTranslationLanguageCode = "es"
	WorkflowTranslationLanguageCodeDe WorkflowTranslationLanguageCode = "de"
	WorkflowTranslationLanguageCodeIt WorkflowTranslationLanguageCode = "it"
	WorkflowTranslationLanguageCodePt WorkflowTranslationLanguageCode = "pt"
	WorkflowTranslationLanguageCodeNl WorkflowTranslationLanguageCode = "nl"
	WorkflowTranslationLanguageCodeJa WorkflowTranslationLanguageCode = "ja"
	WorkflowTranslationLanguageCodeKo WorkflowTranslationLanguageCode = "ko"
	WorkflowTranslationLanguageCodeZh WorkflowTranslationLanguageCode = "zh"
	WorkflowTranslationLanguageCodeAr WorkflowTranslationLanguageCode = "ar"
	WorkflowTranslationLanguageCodeHi WorkflowTranslationLanguageCode = "hi"
	WorkflowTranslationLanguageCodeRu WorkflowTranslationLanguageCode = "ru"
	WorkflowTranslationLanguageCodeTr WorkflowTranslationLanguageCode = "tr"
	WorkflowTranslationLanguageCodePl WorkflowTranslationLanguageCode = "pl"
	WorkflowTranslationLanguageCodeSv WorkflowTranslationLanguageCode = "sv"
)

// Parameters for translating one original caption into another language.
//
// The property LanguageCode is required.
type WorkflowTranslationParams struct {
	// LanguageCode specifies the target language for the translated subtitle track,
	// such as en or fr. The same code appears in the caption and HLS manifest.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode WorkflowTranslationLanguageCode `json:"language_code,omitzero" api:"required"`
	paramObj
}

func (r WorkflowTranslationParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTranslationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTranslationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewParams struct {
	// Name specifies the Workflow name. Use 1 to 128 characters.
	Name string `json:"name" api:"required"`
	// Rules specifies the event subscriptions and actions in this definition. Provide
	// exactly one asset.ready rule and up to 15 connected rules, with at most 32
	// actions and 65536 encoded bytes across all rules. Array order affects display
	// only.
	Rules []WorkflowRuleParams `json:"rules,omitzero" api:"required"`
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
	Rules []WorkflowRuleParams `json:"rules,omitzero"`
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
