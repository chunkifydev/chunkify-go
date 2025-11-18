# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/shared#ChunkifyError">ChunkifyError</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/shared#ResponseOk">ResponseOk</a>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#APIFile">APIFile</a>

Methods:

- <code title="get /api/files/{fileId}">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#APIFile">APIFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/files">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#APIFile">APIFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/files/{fileId}">client.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Jobs

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#HlsAv1Param">HlsAv1Param</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#HlsH264Param">HlsH264Param</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#HlsH265Param">HlsH265Param</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JpgParam">JpgParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#MP4Av1Param">MP4Av1Param</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#MP4H264Param">MP4H264Param</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#MP4H265Param">MP4H265Param</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebmVp9Param">WebmVp9Param</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobNewResponse">JobNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobGetResponse">JobGetResponse</a>

Methods:

- <code title="post /api/jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobNewParams">JobNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobNewResponse">JobNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobGetResponse">JobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobListParams">JobListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Job">Job</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/jobs/{jobId}/cancel">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Files

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobFileListResponse">JobFileListResponse</a>

Methods:

- <code title="get /api/jobs/{jobId}/files">client.Jobs.Files.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobFileListResponse">JobFileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobLogListResponse">JobLogListResponse</a>

Methods:

- <code title="get /api/jobs/{jobId}/logs">client.Jobs.Logs.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobLogListParams">JobLogListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobLogListResponse">JobLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcoders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobTranscoderListResponse">JobTranscoderListResponse</a>

Methods:

- <code title="get /api/jobs/{jobId}/transcoders">client.Jobs.Transcoders.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobTranscoderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#JobTranscoderListResponse">JobTranscoderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Notifications

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Notification">Notification</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationNewResponse">NotificationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationGetResponse">NotificationGetResponse</a>

Methods:

- <code title="post /api/notifications">client.Notifications.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationNewParams">NotificationNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationNewResponse">NotificationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/notifications/{notificationId}">client.Notifications.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, notificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationGetResponse">NotificationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/notifications">client.Notifications.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationListParams">NotificationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Notification">Notification</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/notifications/{notificationId}">client.Notifications.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NotificationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, notificationID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectNewResponse">ProjectNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectGetResponse">ProjectGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /api/projects">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectNewParams">ProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectNewResponse">ProjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectGetResponse">ProjectGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectUpdateParams">ProjectUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/projects">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Sources

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Source">Source</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceNewResponse">SourceNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceGetResponse">SourceGetResponse</a>

Methods:

- <code title="post /api/sources">client.Sources.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceNewParams">SourceNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceNewResponse">SourceNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/sources/{sourceId}">client.Sources.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceGetResponse">SourceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/sources">client.Sources.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceListParams">SourceListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Source">Source</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/sources/{sourceId}">client.Sources.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#SourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Storages

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Storage">Storage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageNewResponse">StorageNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageGetResponse">StorageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageListResponse">StorageListResponse</a>

Methods:

- <code title="post /api/storages">client.Storages.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageNewParams">StorageNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageNewResponse">StorageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/storages/{storageId}">client.Storages.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageGetResponse">StorageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/storages">client.Storages.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageListResponse">StorageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/storages/{storageId}">client.Storages.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#StorageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Tokens

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Token">Token</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenNewResponse">TokenNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenListResponse">TokenListResponse</a>

Methods:

- <code title="post /api/tokens">client.Tokens.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenNewParams">TokenNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenNewResponse">TokenNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/tokens">client.Tokens.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenListResponse">TokenListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/tokens/{tokenId}">client.Tokens.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#TokenService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Uploads

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Upload">Upload</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadNewResponse">UploadNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadGetResponse">UploadGetResponse</a>

Methods:

- <code title="post /api/uploads">client.Uploads.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadNewParams">UploadNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadNewResponse">UploadNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/uploads/{uploadId}">client.Uploads.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadGetResponse">UploadGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/uploads">client.Uploads.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadListParams">UploadListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Upload">Upload</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/uploads/{uploadId}">client.Uploads.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UploadService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#Webhook">Webhook</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookGetResponse">WebhookGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#NewEventWebhookEvent">NewEventWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#UnwrapWebhookEvent">UnwrapWebhookEvent</a>

Methods:

- <code title="post /api/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookNewParams">WebhookNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookGetResponse">WebhookGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookUpdateParams">WebhookUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/chunkify-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
