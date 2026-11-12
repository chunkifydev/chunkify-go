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
	// This field is a union of [WorkflowStaticParamsResp], [WorkflowCaptionParams],
	// [WorkflowTranslationParamsResp], [WorkflowThumbnailParamsResp], [any]
	Params WorkflowActionUnionParams `json:"params"`
	// Any of "create_static_rendition", "create_caption", "translate_caption",
	// "create_thumbnails", "create_transcript".
	Type string `json:"type"`
	// This field is from variant [WorkflowCaptionAction].
	TranscriptID WorkflowCaptionActionTranscriptID `json:"transcript_id"`
	// This field is from variant [WorkflowTranslationAction].
	CaptionID WorkflowTranslationActionCaptionID `json:"caption_id"`
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
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfWorkflowCaptions OfWorkflowTranscriptActions]
type WorkflowActionUnionParams struct {
	// This field will be present if the value is a [WorkflowCaptionParams] instead of
	// an object.
	OfWorkflowCaptions WorkflowCaptionParams `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfWorkflowTranscriptActions any `json:",inline"`
	// This field is from variant [WorkflowStaticParamsResp].
	Resolution WorkflowStaticParamsResolution `json:"resolution"`
	// This field is from variant [WorkflowTranslationParamsResp].
	LanguageCode WorkflowTranslationParamsLanguageCode `json:"language_code"`
	// This field is from variant [WorkflowThumbnailParamsResp].
	Count int64 `json:"count"`
	// This field is from variant [WorkflowThumbnailParamsResp].
	Format WorkflowThumbnailParamsFormat `json:"format"`
	// This field is from variant [WorkflowThumbnailParamsResp].
	Metadata map[string]string `json:"metadata"`
	// This field is from variant [WorkflowThumbnailParamsResp].
	Output WorkflowThumbnailParamsOutput `json:"output"`
	// This field is from variant [WorkflowThumbnailParamsResp].
	Range WorkflowThumbnailParamsRangeResp `json:"range"`
	// This field is from variant [WorkflowThumbnailParamsResp].
	Size int64 `json:"size"`
	// This field is from variant [WorkflowThumbnailParamsResp].
	Timestamps []float64 `json:"timestamps"`
	JSON       struct {
		OfWorkflowCaptions          respjson.Field
		OfWorkflowTranscriptActions respjson.Field
		Resolution                  respjson.Field
		LanguageCode                respjson.Field
		Count                       respjson.Field
		Format                      respjson.Field
		Metadata                    respjson.Field
		Output                      respjson.Field
		Range                       respjson.Field
		Size                        respjson.Field
		Timestamps                  respjson.Field
		raw                         string
	} `json:"-"`
}

func (r *WorkflowActionUnionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowActionUnion to a WorkflowActionUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowActionUnionParam.Overrides()
func (r WorkflowActionUnion) ToParam() WorkflowActionUnionParam {
	return param.Override[WorkflowActionUnionParam](json.RawMessage(r.RawJSON()))
}

func WorkflowActionParamOfCreateStaticRendition(id string, params WorkflowStaticParams) WorkflowActionUnionParam {
	var createStaticRendition WorkflowStaticActionParam
	createStaticRendition.ID = id
	createStaticRendition.Params = params
	return WorkflowActionUnionParam{OfCreateStaticRendition: &createStaticRendition}
}

func WorkflowActionParamOfCreateCaption(id string, params WorkflowCaptionParams) WorkflowActionUnionParam {
	var createCaption WorkflowCaptionActionParam
	createCaption.ID = id
	createCaption.Params = params
	return WorkflowActionUnionParam{OfCreateCaption: &createCaption}
}

func WorkflowActionParamOfTranslateCaption(captionID WorkflowTranslationActionCaptionIDParam, id string, params WorkflowTranslationParams) WorkflowActionUnionParam {
	var translateCaption WorkflowTranslationActionParam
	translateCaption.CaptionID = captionID
	translateCaption.ID = id
	translateCaption.Params = params
	return WorkflowActionUnionParam{OfTranslateCaption: &translateCaption}
}

func WorkflowActionParamOfCreateThumbnails(id string, params WorkflowThumbnailParams) WorkflowActionUnionParam {
	var createThumbnails WorkflowThumbnailActionParam
	createThumbnails.ID = id
	createThumbnails.Params = params
	return WorkflowActionUnionParam{OfCreateThumbnails: &createThumbnails}
}

func WorkflowActionParamOfCreateTranscript(id string, params any) WorkflowActionUnionParam {
	var createTranscript WorkflowTranscriptActionParam
	createTranscript.ID = id
	createTranscript.Params = params
	return WorkflowActionUnionParam{OfCreateTranscript: &createTranscript}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowActionUnionParam struct {
	OfCreateStaticRendition *WorkflowStaticActionParam      `json:",omitzero,inline"`
	OfCreateCaption         *WorkflowCaptionActionParam     `json:",omitzero,inline"`
	OfTranslateCaption      *WorkflowTranslationActionParam `json:",omitzero,inline"`
	OfCreateThumbnails      *WorkflowThumbnailActionParam   `json:",omitzero,inline"`
	OfCreateTranscript      *WorkflowTranscriptActionParam  `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowActionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateStaticRendition,
		u.OfCreateCaption,
		u.OfTranslateCaption,
		u.OfCreateThumbnails,
		u.OfCreateTranscript)
}
func (u *WorkflowActionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowActionUnionParam) asAny() any {
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
func (u WorkflowActionUnionParam) GetTranscriptID() *WorkflowCaptionActionTranscriptIDParam {
	if vt := u.OfCreateCaption; vt != nil {
		return &vt.TranscriptID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowActionUnionParam) GetCaptionID() *WorkflowTranslationActionCaptionIDParam {
	if vt := u.OfTranslateCaption; vt != nil {
		return &vt.CaptionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowActionUnionParam) GetID() *string {
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
func (u WorkflowActionUnionParam) GetType() *string {
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
func (u WorkflowActionUnionParam) GetParams() (res workflowActionUnionParamParams) {
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
// [*WorkflowTranslationParams], [*WorkflowThumbnailParams], [*any]
type workflowActionUnionParamParams struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *chunkify.WorkflowStaticParams:
//	case *chunkify.WorkflowCaptionParams:
//	case *chunkify.WorkflowTranslationParams:
//	case *chunkify.WorkflowThumbnailParams:
//	case *any:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowActionUnionParamParams) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[WorkflowActionUnionParam](
		"type",
		apijson.Discriminator[WorkflowStaticActionParam]("create_static_rendition"),
		apijson.Discriminator[WorkflowCaptionActionParam]("create_caption"),
		apijson.Discriminator[WorkflowTranslationActionParam]("translate_caption"),
		apijson.Discriminator[WorkflowThumbnailActionParam]("create_thumbnails"),
		apijson.Discriminator[WorkflowTranscriptActionParam]("create_transcript"),
	)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
type WorkflowCaptionAction struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params WorkflowCaptionParams `json:"params" api:"required"`
	// Type specifies the action to perform. Must be create_caption.
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowCaptionActionTranscriptID `json:"transcript_id"`
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

// ToParam converts this WorkflowCaptionAction to a WorkflowCaptionActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowCaptionActionParam.Overrides()
func (r WorkflowCaptionAction) ToParam() WorkflowCaptionActionParam {
	return param.Override[WorkflowCaptionActionParam](json.RawMessage(r.RawJSON()))
}

// TranscriptId binds caption publication to the transcript in the triggering
// event. Use it only when connected to create_transcript. When omitted, creates or
// reuses the Asset’s original transcript.
type WorkflowCaptionActionTranscriptID struct {
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
func (r WorkflowCaptionActionTranscriptID) RawJSON() string { return r.JSON.raw }
func (r *WorkflowCaptionActionTranscriptID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generates or reuses original captions and their transcript from the Asset's
// speech.
//
// The properties ID, Params, Type are required.
type WorkflowCaptionActionParam struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies caption generation settings. Pass an empty object.
	Params WorkflowCaptionParams `json:"params,omitzero" api:"required"`
	// TranscriptId binds caption publication to the transcript in the triggering
	// event. Use it only when connected to create_transcript. When omitted, creates or
	// reuses the Asset’s original transcript.
	TranscriptID WorkflowCaptionActionTranscriptIDParam `json:"transcript_id,omitzero"`
	// Type specifies the action to perform. Must be create_caption.
	//
	// This field can be elided, and will marshal its zero value as "create_caption".
	Type constant.CreateCaption `json:"type" default:"create_caption"`
	paramObj
}

func (r WorkflowCaptionActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCaptionActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCaptionActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TranscriptId binds caption publication to the transcript in the triggering
// event. Use it only when connected to create_transcript. When omitted, creates or
// reuses the Asset’s original transcript.
//
// The property From is required.
type WorkflowCaptionActionTranscriptIDParam struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from,omitzero" api:"required"`
	paramObj
}

func (r WorkflowCaptionActionTranscriptIDParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCaptionActionTranscriptIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCaptionActionTranscriptIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowCaptionActionTranscriptIDParam](
		"from", "event.output.id",
	)
}

type WorkflowCaptionParams = any

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

// ToParam converts this WorkflowRule to a WorkflowRuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowRuleParam.Overrides()
func (r WorkflowRule) ToParam() WorkflowRuleParam {
	return param.Override[WorkflowRuleParam](json.RawMessage(r.RawJSON()))
}

// Event specifies when this rule starts. Use asset.ready for the single starting
// rule. Output ready events require an origin caption or static rendition action
// in this Workflow.
type WorkflowRuleEvent string

const (
	WorkflowRuleEventAssetReady       WorkflowRuleEvent = "asset.ready"
	WorkflowRuleEventAssetOutputReady WorkflowRuleEvent = "asset.output.ready"
)

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

// A starting or connected step with one or more independent actions.
//
// The properties ID, Actions, Event are required.
type WorkflowRuleParam struct {
	// Id specifies a stable rule key, unique within this Workflow definition. Use 1 to
	// 64 characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Actions specifies the work to start when the rule matches. Actions can run
	// concurrently; array order does not define execution order. Provide 1 to 32
	// actions per rule, with at most 32 across the definition.
	Actions []WorkflowActionUnionParam `json:"actions,omitzero" api:"required"`
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
	Origin WorkflowRuleOriginParam `json:"origin,omitzero"`
	paramObj
}

func (r WorkflowRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Origin specifies the action whose ready output starts this rule. Required for
// asset.output.ready; omit it for asset.ready. Translation steps must connect to a
// create_caption action. References must point to a compatible action in this
// definition and must not form a cycle.
//
// The properties ActionID, RuleID are required.
type WorkflowRuleOriginParam struct {
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

func (r WorkflowRuleOriginParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRuleOriginParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRuleOriginParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates or reuses one static rendition for the Asset.
type WorkflowStaticAction struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// Params specifies the static rendition resolution.
	Params WorkflowStaticParamsResp `json:"params" api:"required"`
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

// ToParam converts this WorkflowStaticAction to a WorkflowStaticActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowStaticActionParam.Overrides()
func (r WorkflowStaticAction) ToParam() WorkflowStaticActionParam {
	return param.Override[WorkflowStaticActionParam](json.RawMessage(r.RawJSON()))
}

// Creates or reuses one static rendition for the Asset.
//
// The properties ID, Params, Type are required.
type WorkflowStaticActionParam struct {
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

func (r WorkflowStaticActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStaticActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStaticActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for creating one static rendition.
type WorkflowStaticParamsResp struct {
	// Resolution specifies the H.264 MP4 output resolution. highest selects the
	// largest supported output without upscaling. Hobby supports up to 1080p; other
	// plans support up to 2160p. Streaming resolution limits do not apply.
	//
	// Any of "highest", "240p", "360p", "480p", "540p", "720p", "1080p", "1440p",
	// "2160p".
	Resolution WorkflowStaticParamsResolution `json:"resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Resolution  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStaticParamsResp) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStaticParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowStaticParamsResp to a WorkflowStaticParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowStaticParams.Overrides()
func (r WorkflowStaticParamsResp) ToParam() WorkflowStaticParams {
	return param.Override[WorkflowStaticParams](json.RawMessage(r.RawJSON()))
}

// Resolution specifies the H.264 MP4 output resolution. highest selects the
// largest supported output without upscaling. Hobby supports up to 1080p; other
// plans support up to 2160p. Streaming resolution limits do not apply.
type WorkflowStaticParamsResolution string

const (
	WorkflowStaticParamsResolutionHighest WorkflowStaticParamsResolution = "highest"
	WorkflowStaticParamsResolution240p    WorkflowStaticParamsResolution = "240p"
	WorkflowStaticParamsResolution360p    WorkflowStaticParamsResolution = "360p"
	WorkflowStaticParamsResolution480p    WorkflowStaticParamsResolution = "480p"
	WorkflowStaticParamsResolution540p    WorkflowStaticParamsResolution = "540p"
	WorkflowStaticParamsResolution720p    WorkflowStaticParamsResolution = "720p"
	WorkflowStaticParamsResolution1080p   WorkflowStaticParamsResolution = "1080p"
	WorkflowStaticParamsResolution1440p   WorkflowStaticParamsResolution = "1440p"
	WorkflowStaticParamsResolution2160p   WorkflowStaticParamsResolution = "2160p"
)

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
	Resolution WorkflowStaticParamsResolution `json:"resolution,omitzero" api:"required"`
	paramObj
}

func (r WorkflowStaticParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStaticParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStaticParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
type WorkflowThumbnailAction struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowThumbnailParamsResp `json:"params" api:"required"`
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

// ToParam converts this WorkflowThumbnailAction to a WorkflowThumbnailActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowThumbnailActionParam.Overrides()
func (r WorkflowThumbnailAction) ToParam() WorkflowThumbnailActionParam {
	return param.Override[WorkflowThumbnailActionParam](json.RawMessage(r.RawJSON()))
}

// Creates an independent thumbnail request after asset.ready. Replaying this
// action uses its accepted request; separate actions create separate requests.
//
// The properties ID, Params, Type are required.
type WorkflowThumbnailActionParam struct {
	// Id specifies a stable action key, unique within its rule.
	ID string `json:"id" api:"required"`
	// Params specifies sampling and image settings. Outputs use the Asset folder and
	// Storage.
	Params WorkflowThumbnailParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be create_thumbnails.
	//
	// This field can be elided, and will marshal its zero value as
	// "create_thumbnails".
	Type constant.CreateThumbnails `json:"type" default:"create_thumbnails"`
	paramObj
}

func (r WorkflowThumbnailActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThumbnailActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowThumbnailActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generate additional thumbnails in the Asset folder and Storage. Each layout,
// format and size combination can be requested once. Specify exactly one of
// timestamps or count.
type WorkflowThumbnailParamsResp struct {
	// Count specifies 1 to 100 thumbnails at the middle of equal sections of the
	// selected range.
	Count int64 `json:"count"`
	// Format selects JPEG or WebP. Defaults to JPEG.
	//
	// Any of "jpeg", "webp".
	Format WorkflowThumbnailParamsFormat `json:"format"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output WorkflowThumbnailParamsOutput `json:"output"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range WorkflowThumbnailParamsRangeResp `json:"range"`
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
func (r WorkflowThumbnailParamsResp) RawJSON() string { return r.JSON.raw }
func (r *WorkflowThumbnailParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowThumbnailParamsResp to a WorkflowThumbnailParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowThumbnailParams.Overrides()
func (r WorkflowThumbnailParamsResp) ToParam() WorkflowThumbnailParams {
	return param.Override[WorkflowThumbnailParams](json.RawMessage(r.RawJSON()))
}

// Format selects JPEG or WebP. Defaults to JPEG.
type WorkflowThumbnailParamsFormat string

const (
	WorkflowThumbnailParamsFormatJpeg WorkflowThumbnailParamsFormat = "jpeg"
	WorkflowThumbnailParamsFormatWebp WorkflowThumbnailParamsFormat = "webp"
)

// Output selects images or sprites. Defaults to images.
type WorkflowThumbnailParamsOutput string

const (
	WorkflowThumbnailParamsOutputImages  WorkflowThumbnailParamsOutput = "images"
	WorkflowThumbnailParamsOutputSprites WorkflowThumbnailParamsOutput = "sprites"
)

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type WorkflowThumbnailParamsRangeResp struct {
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
func (r WorkflowThumbnailParamsRangeResp) RawJSON() string { return r.JSON.raw }
func (r *WorkflowThumbnailParamsRangeResp) UnmarshalJSON(data []byte) error {
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
	Format WorkflowThumbnailParamsFormat `json:"format,omitzero"`
	// Metadata attaches string values to this request. The serialized object may
	// contain at most 2048 bytes.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Output selects images or sprites. Defaults to images.
	//
	// Any of "images", "sprites".
	Output WorkflowThumbnailParamsOutput `json:"output,omitzero"`
	// Range limits sampling to part of the Source. Defaults to the whole Source
	// duration.
	Range WorkflowThumbnailParamsRange `json:"range,omitzero"`
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

// Range limits sampling to part of the Source. Defaults to the whole Source
// duration.
type WorkflowThumbnailParamsRange struct {
	// End specifies the exclusive upper bound in seconds. Defaults to the Source
	// duration.
	End param.Opt[float64] `json:"end,omitzero"`
	// Start specifies the inclusive lower bound in seconds. Defaults to zero.
	Start param.Opt[float64] `json:"start,omitzero"`
	paramObj
}

func (r WorkflowThumbnailParamsRange) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThumbnailParamsRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowThumbnailParamsRange) UnmarshalJSON(data []byte) error {
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
func (r WorkflowTranscriptAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTranscriptAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowTranscriptAction to a
// WorkflowTranscriptActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowTranscriptActionParam.Overrides()
func (r WorkflowTranscriptAction) ToParam() WorkflowTranscriptActionParam {
	return param.Override[WorkflowTranscriptActionParam](json.RawMessage(r.RawJSON()))
}

// Creates or reuses the original-language transcript. Can run alongside caption
// generation in the same asset.ready rule.
//
// The properties ID, Params, Type are required.
type WorkflowTranscriptActionParam struct {
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

func (r WorkflowTranscriptActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTranscriptActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTranscriptActionParam) UnmarshalJSON(data []byte) error {
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
	CaptionID WorkflowTranslationActionCaptionID `json:"caption_id" api:"required"`
	// Params specifies the target language.
	Params WorkflowTranslationParamsResp `json:"params" api:"required"`
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

// ToParam converts this WorkflowTranslationAction to a
// WorkflowTranslationActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowTranslationActionParam.Overrides()
func (r WorkflowTranslationAction) ToParam() WorkflowTranslationActionParam {
	return param.Override[WorkflowTranslationActionParam](json.RawMessage(r.RawJSON()))
}

// CaptionId specifies the binding to the original caption in the triggering event.
type WorkflowTranslationActionCaptionID struct {
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
func (r WorkflowTranslationActionCaptionID) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTranslationActionCaptionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translates or reuses an original caption in one target language. Requires an
// asset.output.ready rule connected to an original caption creation action.
//
// The properties ID, CaptionID, Params, Type are required.
type WorkflowTranslationActionParam struct {
	// Id specifies a stable action key, unique within its rule. Use 1 to 64
	// characters, starting with a letter and containing only letters, digits, or
	// underscores.
	ID string `json:"id" api:"required"`
	// CaptionId specifies the binding to the original caption in the triggering event.
	CaptionID WorkflowTranslationActionCaptionIDParam `json:"caption_id,omitzero" api:"required"`
	// Params specifies the target language.
	Params WorkflowTranslationParams `json:"params,omitzero" api:"required"`
	// Type specifies the action to perform. Must be translate_caption.
	//
	// This field can be elided, and will marshal its zero value as
	// "translate_caption".
	Type constant.TranslateCaption `json:"type" default:"translate_caption"`
	paramObj
}

func (r WorkflowTranslationActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTranslationActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTranslationActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CaptionId specifies the binding to the original caption in the triggering event.
//
// The property From is required.
type WorkflowTranslationActionCaptionIDParam struct {
	// From specifies where to read the input output ID. Use event.output.id to consume
	// the output in the triggering asset.output.ready event.
	//
	// Any of "event.output.id".
	From string `json:"from,omitzero" api:"required"`
	paramObj
}

func (r WorkflowTranslationActionCaptionIDParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTranslationActionCaptionIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTranslationActionCaptionIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowTranslationActionCaptionIDParam](
		"from", "event.output.id",
	)
}

// Parameters for translating one original caption into another language.
type WorkflowTranslationParamsResp struct {
	// LanguageCode specifies the target language for the translated subtitle track,
	// such as en or fr. The same code appears in the caption and HLS manifest.
	//
	// Any of "en", "fr", "es", "de", "it", "pt", "nl", "ja", "ko", "zh", "ar", "hi",
	// "ru", "tr", "pl", "sv".
	LanguageCode WorkflowTranslationParamsLanguageCode `json:"language_code" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTranslationParamsResp) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTranslationParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowTranslationParamsResp to a
// WorkflowTranslationParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowTranslationParams.Overrides()
func (r WorkflowTranslationParamsResp) ToParam() WorkflowTranslationParams {
	return param.Override[WorkflowTranslationParams](json.RawMessage(r.RawJSON()))
}

// LanguageCode specifies the target language for the translated subtitle track,
// such as en or fr. The same code appears in the caption and HLS manifest.
type WorkflowTranslationParamsLanguageCode string

const (
	WorkflowTranslationParamsLanguageCodeEn WorkflowTranslationParamsLanguageCode = "en"
	WorkflowTranslationParamsLanguageCodeFr WorkflowTranslationParamsLanguageCode = "fr"
	WorkflowTranslationParamsLanguageCodeEs WorkflowTranslationParamsLanguageCode = "es"
	WorkflowTranslationParamsLanguageCodeDe WorkflowTranslationParamsLanguageCode = "de"
	WorkflowTranslationParamsLanguageCodeIt WorkflowTranslationParamsLanguageCode = "it"
	WorkflowTranslationParamsLanguageCodePt WorkflowTranslationParamsLanguageCode = "pt"
	WorkflowTranslationParamsLanguageCodeNl WorkflowTranslationParamsLanguageCode = "nl"
	WorkflowTranslationParamsLanguageCodeJa WorkflowTranslationParamsLanguageCode = "ja"
	WorkflowTranslationParamsLanguageCodeKo WorkflowTranslationParamsLanguageCode = "ko"
	WorkflowTranslationParamsLanguageCodeZh WorkflowTranslationParamsLanguageCode = "zh"
	WorkflowTranslationParamsLanguageCodeAr WorkflowTranslationParamsLanguageCode = "ar"
	WorkflowTranslationParamsLanguageCodeHi WorkflowTranslationParamsLanguageCode = "hi"
	WorkflowTranslationParamsLanguageCodeRu WorkflowTranslationParamsLanguageCode = "ru"
	WorkflowTranslationParamsLanguageCodeTr WorkflowTranslationParamsLanguageCode = "tr"
	WorkflowTranslationParamsLanguageCodePl WorkflowTranslationParamsLanguageCode = "pl"
	WorkflowTranslationParamsLanguageCodeSv WorkflowTranslationParamsLanguageCode = "sv"
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
	LanguageCode WorkflowTranslationParamsLanguageCode `json:"language_code,omitzero" api:"required"`
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
	Rules []WorkflowRuleParam `json:"rules,omitzero" api:"required"`
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
	Rules []WorkflowRuleParam `json:"rules,omitzero"`
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
