# Shared Response Types

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/shared#ChunkifyError">ChunkifyError</a>

# Assets

Params Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetStaticRenditionCreateParams">AssetStaticRenditionCreateParams</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetStatus">AssetStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Asset">Asset</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetStatus">AssetStatus</a>

Methods:

- <code title="post /assets">client.Assets.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetNewParams">AssetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Asset">Asset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Asset">Asset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetUpdateParams">AssetUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /assets">client.Assets.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetListParams">AssetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Asset">Asset</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetDeleteParams">AssetDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Outputs

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputUnion">AssetOutputUnion</a>

Methods:

- <code title="post /assets/{assetId}/outputs">client.Assets.Outputs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputNewParams">AssetOutputNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputUnion">AssetOutputUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/{assetId}/outputs/{outputId}">client.Assets.Outputs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, outputID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputGetParams">AssetOutputGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputUnion">AssetOutputUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/{assetId}/outputs">client.Assets.Outputs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputListParams">AssetOutputListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputUnion">AssetOutputUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /assets/{assetId}/outputs/{outputId}">client.Assets.Outputs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, outputID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetOutputDeleteParams">AssetOutputDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Workflow

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetActionRun">AssetActionRun</a>

Methods:

- <code title="post /assets/{assetId}/workflow/update">client.Assets.Workflow.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetWorkflowService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRun">WorkflowRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="get /assets/{assetId}/workflow/actions">client.Assets.Workflow.Actions.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetWorkflowActionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetWorkflowActionListParams">AssetWorkflowActionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetActionRun">AssetActionRun</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Events

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Event">Event</a>

Methods:

- <code title="get /events/{eventId}">client.Events.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#EventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Event">Event</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /events">client.Events.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#EventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#EventListParams">EventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Event">Event</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFile">JobFile</a>

Methods:

- <code title="get /files/{fileId}">client.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFile">JobFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileListParams">FileListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFile">JobFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /files/{fileId}">client.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Jobs

Params Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#HlsAv1Param">HlsAv1Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#HlsH264Param">HlsH264Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#HlsH265Param">HlsH265Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JpgParam">JpgParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MP4Av1Param">MP4Av1Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MP4H264Param">MP4H264Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MP4H265Param">MP4H265Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebmVp9Param">WebmVp9Param</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#HlsAv1">HlsAv1</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#HlsH264">HlsH264</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#HlsH265">HlsH265</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Jpg">Jpg</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MP4Av1">MP4Av1</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MP4H264">MP4H264</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MP4H265">MP4H265</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebmVp9">WebmVp9</a>

Methods:

- <code title="post /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobNewParams">JobNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobListParams">JobListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Job">Job</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /jobs/{jobId}/cancel">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Files

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFileListResponse">JobFileListResponse</a>

Methods:

- <code title="get /jobs/{jobId}/files">client.Jobs.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFileListResponse">JobFileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogListResponse">JobLogListResponse</a>

Methods:

- <code title="get /jobs/{jobId}/logs">client.Jobs.Logs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogListParams">JobLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogListResponse">JobLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcoders

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobTranscoderListResponse">JobTranscoderListResponse</a>

Methods:

- <code title="get /jobs/{jobId}/transcoders">client.Jobs.Transcoders.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobTranscoderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobTranscoderListResponse">JobTranscoderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Metrics

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricsAssets">MetricsAssets</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricsStorage">MetricsStorage</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricsUsage">MetricsUsage</a>

Methods:

- <code title="get /metrics/assets">client.Metrics.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricService.Assets">Assets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricAssetsParams">MetricAssetsParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricsAssets">MetricsAssets</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /metrics/storage">client.Metrics.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricService.Storage">Storage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricStorageParams">MetricStorageParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricsStorage">MetricsStorage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /metrics/usage">client.Metrics.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricService.Usage">Usage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricUsageParams">MetricUsageParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#MetricsUsage">MetricsUsage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Notifications

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>

Methods:

- <code title="post /notifications">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationNewParams">NotificationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /notifications/{notificationId}">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, notificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /notifications">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationListParams">NotificationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /notifications/{notificationId}">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, notificationID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /projects">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectNewParams">ProjectNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectUpdateParams">ProjectUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /projects">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Analytics

### Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AnalyticsConnection">AnalyticsConnection</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionNewResponse">ProjectAnalyticsConnectionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionListResponse">ProjectAnalyticsConnectionListResponse</a>

Methods:

- <code title="post /projects/{projectId}/analytics/connections">client.Projects.Analytics.Connections.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionNewParams">ProjectAnalyticsConnectionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionNewResponse">ProjectAnalyticsConnectionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/analytics/connections/{connectionId}">client.Projects.Analytics.Connections.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionGetParams">ProjectAnalyticsConnectionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AnalyticsConnection">AnalyticsConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /projects/{projectId}/analytics/connections/{connectionId}">client.Projects.Analytics.Connections.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionUpdateParams">ProjectAnalyticsConnectionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AnalyticsConnection">AnalyticsConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/analytics/connections">client.Projects.Analytics.Connections.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionListResponse">ProjectAnalyticsConnectionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /projects/{projectId}/analytics/connections/{connectionId}">client.Projects.Analytics.Connections.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionDeleteParams">ProjectAnalyticsConnectionDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

#### Keys

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AnalyticsKey">AnalyticsKey</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyNewResponse">ProjectAnalyticsConnectionKeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyListResponse">ProjectAnalyticsConnectionKeyListResponse</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyRotateResponse">ProjectAnalyticsConnectionKeyRotateResponse</a>

Methods:

- <code title="post /projects/{projectId}/analytics/connections/{connectionId}/keys">client.Projects.Analytics.Connections.Keys.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyNewParams">ProjectAnalyticsConnectionKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyNewResponse">ProjectAnalyticsConnectionKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/analytics/connections/{connectionId}/keys">client.Projects.Analytics.Connections.Keys.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyListParams">ProjectAnalyticsConnectionKeyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyListResponse">ProjectAnalyticsConnectionKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /projects/{projectId}/analytics/connections/{connectionId}/keys/{keyId}/rotate">client.Projects.Analytics.Connections.Keys.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, keyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyRotateParams">ProjectAnalyticsConnectionKeyRotateParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectAnalyticsConnectionKeyRotateResponse">ProjectAnalyticsConnectionKeyRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sources

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>

Methods:

- <code title="post /sources">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceNewParams">SourceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sources/{sourceId}">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sources">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceListParams">SourceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sources/{sourceId}">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Storages

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageUnion">StorageUnion</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageListResponse">StorageListResponse</a>

Methods:

- <code title="post /storages">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageNewParams">StorageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageUnion">StorageUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storages/{storageId}">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageUnion">StorageUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /storages/{storageId}">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageUpdateParams">StorageUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /storages">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageListResponse">StorageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /storages/{storageId}">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Thumbnails

Params Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailCreateStorageParams">ThumbnailCreateStorageParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Thumbnail">Thumbnail</a>

Methods:

- <code title="post /thumbnails">client.Thumbnails.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailNewParams">ThumbnailNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Thumbnail">Thumbnail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /thumbnails/{thumbnailId}">client.Thumbnails.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, thumbnailID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Thumbnail">Thumbnail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /thumbnails">client.Thumbnails.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailListParams">ThumbnailListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Thumbnail">Thumbnail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /thumbnails/{thumbnailId}">client.Thumbnails.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, thumbnailID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ThumbnailDeleteParams">ThumbnailDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Tokens

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Token">Token</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenListResponse">TokenListResponse</a>

Methods:

- <code title="post /tokens">client.Tokens.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenNewParams">TokenNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Token">Token</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /tokens">client.Tokens.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenListResponse">TokenListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /tokens/{tokenId}">client.Tokens.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Transcodings

Params Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingM4aAacParam">TranscodingM4aAacParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingMP4Av1Param">TranscodingMP4Av1Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingMP4H264Param">TranscodingMP4H264Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingMP4H265Param">TranscodingMP4H265Param</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingWebmVp9Param">TranscodingWebmVp9Param</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Transcoding">Transcoding</a>

Methods:

- <code title="post /transcodings">client.Transcodings.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingNewParams">TranscodingNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Transcoding">Transcoding</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transcodings/{transcodingId}">client.Transcodings.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transcodingID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Transcoding">Transcoding</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transcodings">client.Transcodings.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingListParams">TranscodingListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Transcoding">Transcoding</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /transcodings/{transcodingId}">client.Transcodings.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transcodingID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingDeleteParams">TranscodingDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /transcodings/{transcodingId}/cancel">client.Transcodings.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TranscodingService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transcodingID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Uploads

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>

Methods:

- <code title="post /uploads">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadNewParams">UploadNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /uploads/{uploadId}">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /uploads">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadListParams">UploadListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /uploads/{uploadId}">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /uploads/completion/{token}">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadCompleteParams">UploadCompleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Webhook">Webhook</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NewEventWebhookEvent">NewEventWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UnwrapWebhookEvent">UnwrapWebhookEvent</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookUpdateParams">WebhookUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Workflows

Params Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowActionUnionParam">WorkflowActionUnionParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowCaptionActionParam">WorkflowCaptionActionParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowCaptionParams">WorkflowCaptionParams</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRuleParam">WorkflowRuleParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowStaticActionParam">WorkflowStaticActionParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowStaticParams">WorkflowStaticParams</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowThumbnailActionParam">WorkflowThumbnailActionParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowThumbnailParams">WorkflowThumbnailParams</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowTranscriptActionParam">WorkflowTranscriptActionParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowTranslationActionParam">WorkflowTranslationActionParam</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowTranslationParams">WorkflowTranslationParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Workflow">Workflow</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowActionUnion">WorkflowActionUnion</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowCaptionAction">WorkflowCaptionAction</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowCaptionParams">WorkflowCaptionParams</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRule">WorkflowRule</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowStaticAction">WorkflowStaticAction</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowStaticParamsResp">WorkflowStaticParamsResp</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowThumbnailAction">WorkflowThumbnailAction</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowThumbnailParamsResp">WorkflowThumbnailParamsResp</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowTranscriptAction">WorkflowTranscriptAction</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowTranslationAction">WorkflowTranslationAction</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowTranslationParamsResp">WorkflowTranslationParamsResp</a>

Methods:

- <code title="post /workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowNewParams">WorkflowNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workflows/{workflowId}">client.Workflows.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /workflows/{workflowId}">client.Workflows.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowUpdateParams">WorkflowUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowListParams">WorkflowListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Workflow">Workflow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /workflows/{workflowId}">client.Workflows.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Revisions

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRevision">WorkflowRevision</a>

Methods:

- <code title="get /workflows/{workflowId}/revisions/{revision}">client.Workflows.Revisions.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRevisionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, revision <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRevisionGetParams">WorkflowRevisionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRevision">WorkflowRevision</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRun">WorkflowRun</a>

Methods:

- <code title="get /workflows/{workflowId}/runs/{runId}">client.Workflows.Runs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunGetParams">WorkflowRunGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRun">WorkflowRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workflows/{workflowId}/runs">client.Workflows.Runs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunListParams">WorkflowRunListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRun">WorkflowRun</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="get /workflows/{workflowId}/runs/{runId}/actions/{actionRunId}">client.Workflows.Runs.Actions.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunActionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, actionRunID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunActionGetParams">WorkflowRunActionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetActionRun">AssetActionRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workflows/{workflowId}/runs/{runId}/actions">client.Workflows.Runs.Actions.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunActionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WorkflowRunActionListParams">WorkflowRunActionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#AssetActionRun">AssetActionRun</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
