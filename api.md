# Shared Response Types

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/shared#ChunkifyError">ChunkifyError</a>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFile">JobFile</a>

Methods:

- <code title="get /api/files/{fileId}">client.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFile">JobFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/files">client.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileListParams">FileListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFile">JobFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/files/{fileId}">client.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

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

- <code title="post /api/jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobNewParams">JobNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobListParams">JobListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Job">Job</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/jobs/{jobId}/cancel">client.Jobs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Files

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFileListResponse">JobFileListResponse</a>

Methods:

- <code title="get /api/jobs/{jobId}/files">client.Jobs.Files.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobFileListResponse">JobFileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogListResponse">JobLogListResponse</a>

Methods:

- <code title="get /api/jobs/{jobId}/logs">client.Jobs.Logs.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogListParams">JobLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobLogListResponse">JobLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcoders

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobTranscoderListResponse">JobTranscoderListResponse</a>

Methods:

- <code title="get /api/jobs/{jobId}/transcoders">client.Jobs.Transcoders.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobTranscoderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#JobTranscoderListResponse">JobTranscoderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Notifications

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>

Methods:

- <code title="post /api/notifications">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationNewParams">NotificationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/notifications/{notificationId}">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, notificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/notifications">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationListParams">NotificationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Notification">Notification</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/notifications/{notificationId}">client.Notifications.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NotificationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, notificationID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /api/projects">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectNewParams">ProjectNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectUpdateParams">ProjectUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/projects">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Sources

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>

Methods:

- <code title="post /api/sources">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceNewParams">SourceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/sources/{sourceId}">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/sources">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceListParams">SourceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Source">Source</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/sources/{sourceId}">client.Sources.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#SourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sourceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Storages

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageUnion">StorageUnion</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageListResponse">StorageListResponse</a>

Methods:

- <code title="post /api/storages">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageNewParams">StorageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageUnion">StorageUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/storages/{storageId}">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageUnion">StorageUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/storages">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageListResponse">StorageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/storages/{storageId}">client.Storages.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#StorageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Tokens

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Token">Token</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenListResponse">TokenListResponse</a>

Methods:

- <code title="post /api/tokens">client.Tokens.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenNewParams">TokenNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Token">Token</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/tokens">client.Tokens.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenListResponse">TokenListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/tokens/{tokenId}">client.Tokens.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#TokenService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Uploads

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>

Methods:

- <code title="post /api/uploads">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadNewParams">UploadNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/uploads/{uploadId}">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/uploads">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadListParams">UploadListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go/packages/pagination#PaginatedResults">PaginatedResults</a>[<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Upload">Upload</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/uploads/{uploadId}">client.Uploads.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UploadService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Webhook">Webhook</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#NewEventWebhookEvent">NewEventWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#UnwrapWebhookEvent">UnwrapWebhookEvent</a>

Methods:

- <code title="post /api/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookUpdateParams">WebhookUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go">chunkify</a>.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/webhooks/{webhookId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/chunkifydev/chunkify-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
