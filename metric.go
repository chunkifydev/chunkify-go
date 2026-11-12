// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/chunkifydev/chunkify-go/internal/apijson"
	"github.com/chunkifydev/chunkify-go/internal/apiquery"
	"github.com/chunkifydev/chunkify-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-go/option"
	"github.com/chunkifydev/chunkify-go/packages/param"
	"github.com/chunkifydev/chunkify-go/packages/respjson"
)

// MetricService contains methods and other services that help with interacting
// with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetricService] method instead.
type MetricService struct {
	Options []option.RequestOption
}

// NewMetricService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMetricService(opts ...option.RequestOption) (r MetricService) {
	r = MetricService{}
	r.Options = opts
	return
}

// Retrieve Asset outcomes, timings, output sizes and source profiles by UTC
// creation date.
func (r *MetricService) Assets(ctx context.Context, query MetricAssetsParams, opts ...option.RequestOption) (res *MetricsAssets, err error) {
	var env MetricAssetsResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "metrics/assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve recorded additions and successful removals by storage and UTC date.
// Historical coverage depends on event retention.
func (r *MetricService) Storage(ctx context.Context, query MetricStorageParams, opts ...option.RequestOption) (res *MetricsStorage, err error) {
	var env MetricStorageResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "metrics/storage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Retrieve project usage totals and daily values for a UTC date range.
func (r *MetricService) Usage(ctx context.Context, query MetricUsageParams, opts ...option.RequestOption) (res *MetricsUsage, err error) {
	var env MetricUsageResponseEnvelope
	var preClientOpts = []option.RequestOption{requestconfig.WithProjectAccessTokenSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "metrics/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Measurements saved on retained Assets. Explicit deletion can change historical
// results; temporary media cleanup does not erase saved Asset measurements.
type MetricsAssets struct {
	// Daily specifies one record for each selected creation day, oldest first.
	Daily []MetricsAssetsDaily `json:"daily" api:"required"`
	// From specifies the inclusive UTC start date.
	From string `json:"from" api:"required"`
	// Summary specifies the complete creation-date cohort.
	Summary MetricsAssetsSummary `json:"summary" api:"required"`
	// To specifies the exclusive UTC end date.
	To string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Daily       respjson.Field
		From        respjson.Field
		Summary     respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssets) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Asset measurements grouped by creation day.
type MetricsAssetsDaily struct {
	// Asset outcomes grouped by Asset creation date.
	Assets MetricsAssetsDailyAssets `json:"assets" api:"required"`
	// Date specifies the UTC Asset creation day.
	Date string `json:"date" api:"required"`
	// Measured output sizes for completed Assets. Sample counts describe measurement
	// coverage.
	Outputs MetricsAssetsDailyOutputs `json:"outputs" api:"required"`
	// Source profiles count Assets, not unique Sources.
	Sources MetricsAssetsDailySources `json:"sources" api:"required"`
	// Streaming output distribution on completed Assets.
	Streaming MetricsAssetsDailyStreaming `json:"streaming" api:"required"`
	// Lifecycle timing measurements. Period medians use all samples, not averages of
	// daily medians.
	Timings MetricsAssetsDailyTimings `json:"timings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		Date        respjson.Field
		Outputs     respjson.Field
		Sources     respjson.Field
		Streaming   respjson.Field
		Timings     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsDaily) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsDaily) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Asset outcomes grouped by Asset creation date.
type MetricsAssetsDailyAssets struct {
	// Completed specifies Assets with a recorded ready timestamp.
	Completed int64 `json:"completed" api:"required"`
	// Failed specifies retained Assets currently in failed status.
	Failed int64 `json:"failed" api:"required"`
	// ProcessedSeconds specifies measured media duration of completed Assets,
	// preserving fractional seconds.
	ProcessedSeconds float64 `json:"processed_seconds" api:"required"`
	// Saved specifies retained Assets created in the period.
	Saved int64 `json:"saved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed        respjson.Field
		Failed           respjson.Field
		ProcessedSeconds respjson.Field
		Saved            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsDailyAssets) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsDailyAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Measured output sizes for completed Assets. Sample counts describe measurement
// coverage.
type MetricsAssetsDailyOutputs struct {
	// MeasuredAssetCount specifies completed Assets with at least one measured
	// rendition.
	MeasuredAssetCount int64 `json:"measured_asset_count" api:"required"`
	// StaticBytes specifies the sum of measured ready static rendition sizes. Null
	// when no measurements exist.
	StaticBytes int64 `json:"static_bytes" api:"required"`
	// StaticSampleCount specifies ready static renditions with measured sizes.
	StaticSampleCount int64 `json:"static_sample_count" api:"required"`
	// StreamingBytes specifies the sum of measured streaming rendition sizes. Null
	// when no measurements exist.
	StreamingBytes int64 `json:"streaming_bytes" api:"required"`
	// StreamingSampleCount specifies streaming renditions with measured sizes.
	StreamingSampleCount int64 `json:"streaming_sample_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MeasuredAssetCount   respjson.Field
		StaticBytes          respjson.Field
		StaticSampleCount    respjson.Field
		StreamingBytes       respjson.Field
		StreamingSampleCount respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsDailyOutputs) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsDailyOutputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source profiles count Assets, not unique Sources.
type MetricsAssetsDailySources struct {
	// AssetsByCodec specifies Asset counts by normalized source video codec, including
	// unknown. All recorded codecs are returned.
	AssetsByCodec map[string]int64 `json:"assets_by_codec" api:"required"`
	// AssetsByDuration specifies Asset counts by source duration. Keys are under_1m,
	// 1_to_5m, 5_to_10m, 10_to_30m, 30_to_60m, over_60m and unknown. Boundaries are
	// 60, 300, 600, 1800 and 3600 seconds; exactly 3600 belongs to 30_to_60m.
	AssetsByDuration map[string]int64 `json:"assets_by_duration" api:"required"`
	// AssetsByResolution specifies Asset counts by source resolution. Keys are
	// up_to_576, 577_to_720, 721_to_1080, 1081_to_2159, 2160_plus and unknown, using
	// the shorter pixel edge.
	AssetsByResolution map[string]int64 `json:"assets_by_resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetsByCodec      respjson.Field
		AssetsByDuration   respjson.Field
		AssetsByResolution respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsDailySources) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsDailySources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming output distribution on completed Assets.
type MetricsAssetsDailyStreaming struct {
	// AssetCount specifies completed Assets, including those with no streaming
	// renditions.
	AssetCount int64 `json:"asset_count" api:"required"`
	// AssetsByHighestResolution specifies counts by highest measured streaming
	// resolution. Keys are up_to_576, 577_to_720, 721_to_1080, 1081_to_2159, 2160_plus
	// and unknown, using the shorter pixel edge.
	AssetsByHighestResolution map[string]int64 `json:"assets_by_highest_resolution" api:"required"`
	// RenditionCount specifies streaming renditions on completed Assets.
	RenditionCount int64 `json:"rendition_count" api:"required"`
	// RenditionsPerAsset specifies rendition_count divided by asset_count. Null when
	// no completed Assets exist.
	RenditionsPerAsset float64 `json:"renditions_per_asset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetCount                respjson.Field
		AssetsByHighestResolution respjson.Field
		RenditionCount            respjson.Field
		RenditionsPerAsset        respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsDailyStreaming) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsDailyStreaming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle timing measurements. Period medians use all samples, not averages of
// daily medians.
type MetricsAssetsDailyTimings struct {
	// CompletionMedianSeconds specifies the median elapsed time from creation to full
	// readiness. Null when no samples exist.
	CompletionMedianSeconds float64 `json:"completion_median_seconds" api:"required"`
	// CompletionSampleCount specifies Assets with a full readiness measurement.
	CompletionSampleCount int64 `json:"completion_sample_count" api:"required"`
	// PlaybackMedianSeconds specifies the median elapsed time from creation to
	// playback readiness. Null when no samples exist.
	PlaybackMedianSeconds float64 `json:"playback_median_seconds" api:"required"`
	// PlaybackSampleCount specifies Assets with a playback readiness measurement.
	PlaybackSampleCount int64 `json:"playback_sample_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionMedianSeconds respjson.Field
		CompletionSampleCount   respjson.Field
		PlaybackMedianSeconds   respjson.Field
		PlaybackSampleCount     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsDailyTimings) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsDailyTimings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary specifies the complete creation-date cohort.
type MetricsAssetsSummary struct {
	// Asset outcomes grouped by Asset creation date.
	Assets MetricsAssetsSummaryAssets `json:"assets" api:"required"`
	// Measured output sizes for completed Assets. Sample counts describe measurement
	// coverage.
	Outputs MetricsAssetsSummaryOutputs `json:"outputs" api:"required"`
	// Source profiles count Assets, not unique Sources.
	Sources MetricsAssetsSummarySources `json:"sources" api:"required"`
	// Streaming output distribution on completed Assets.
	Streaming MetricsAssetsSummaryStreaming `json:"streaming" api:"required"`
	// Lifecycle timing measurements. Period medians use all samples, not averages of
	// daily medians.
	Timings MetricsAssetsSummaryTimings `json:"timings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		Outputs     respjson.Field
		Sources     respjson.Field
		Streaming   respjson.Field
		Timings     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsSummary) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Asset outcomes grouped by Asset creation date.
type MetricsAssetsSummaryAssets struct {
	// Completed specifies Assets with a recorded ready timestamp.
	Completed int64 `json:"completed" api:"required"`
	// Failed specifies retained Assets currently in failed status.
	Failed int64 `json:"failed" api:"required"`
	// ProcessedSeconds specifies measured media duration of completed Assets,
	// preserving fractional seconds.
	ProcessedSeconds float64 `json:"processed_seconds" api:"required"`
	// Saved specifies retained Assets created in the period.
	Saved int64 `json:"saved" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed        respjson.Field
		Failed           respjson.Field
		ProcessedSeconds respjson.Field
		Saved            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsSummaryAssets) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsSummaryAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Measured output sizes for completed Assets. Sample counts describe measurement
// coverage.
type MetricsAssetsSummaryOutputs struct {
	// MeasuredAssetCount specifies completed Assets with at least one measured
	// rendition.
	MeasuredAssetCount int64 `json:"measured_asset_count" api:"required"`
	// StaticBytes specifies the sum of measured ready static rendition sizes. Null
	// when no measurements exist.
	StaticBytes int64 `json:"static_bytes" api:"required"`
	// StaticSampleCount specifies ready static renditions with measured sizes.
	StaticSampleCount int64 `json:"static_sample_count" api:"required"`
	// StreamingBytes specifies the sum of measured streaming rendition sizes. Null
	// when no measurements exist.
	StreamingBytes int64 `json:"streaming_bytes" api:"required"`
	// StreamingSampleCount specifies streaming renditions with measured sizes.
	StreamingSampleCount int64 `json:"streaming_sample_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MeasuredAssetCount   respjson.Field
		StaticBytes          respjson.Field
		StaticSampleCount    respjson.Field
		StreamingBytes       respjson.Field
		StreamingSampleCount respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsSummaryOutputs) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsSummaryOutputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source profiles count Assets, not unique Sources.
type MetricsAssetsSummarySources struct {
	// AssetsByCodec specifies Asset counts by normalized source video codec, including
	// unknown. All recorded codecs are returned.
	AssetsByCodec map[string]int64 `json:"assets_by_codec" api:"required"`
	// AssetsByDuration specifies Asset counts by source duration. Keys are under_1m,
	// 1_to_5m, 5_to_10m, 10_to_30m, 30_to_60m, over_60m and unknown. Boundaries are
	// 60, 300, 600, 1800 and 3600 seconds; exactly 3600 belongs to 30_to_60m.
	AssetsByDuration map[string]int64 `json:"assets_by_duration" api:"required"`
	// AssetsByResolution specifies Asset counts by source resolution. Keys are
	// up_to_576, 577_to_720, 721_to_1080, 1081_to_2159, 2160_plus and unknown, using
	// the shorter pixel edge.
	AssetsByResolution map[string]int64 `json:"assets_by_resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetsByCodec      respjson.Field
		AssetsByDuration   respjson.Field
		AssetsByResolution respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsSummarySources) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsSummarySources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming output distribution on completed Assets.
type MetricsAssetsSummaryStreaming struct {
	// AssetCount specifies completed Assets, including those with no streaming
	// renditions.
	AssetCount int64 `json:"asset_count" api:"required"`
	// AssetsByHighestResolution specifies counts by highest measured streaming
	// resolution. Keys are up_to_576, 577_to_720, 721_to_1080, 1081_to_2159, 2160_plus
	// and unknown, using the shorter pixel edge.
	AssetsByHighestResolution map[string]int64 `json:"assets_by_highest_resolution" api:"required"`
	// RenditionCount specifies streaming renditions on completed Assets.
	RenditionCount int64 `json:"rendition_count" api:"required"`
	// RenditionsPerAsset specifies rendition_count divided by asset_count. Null when
	// no completed Assets exist.
	RenditionsPerAsset float64 `json:"renditions_per_asset" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetCount                respjson.Field
		AssetsByHighestResolution respjson.Field
		RenditionCount            respjson.Field
		RenditionsPerAsset        respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsSummaryStreaming) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsSummaryStreaming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle timing measurements. Period medians use all samples, not averages of
// daily medians.
type MetricsAssetsSummaryTimings struct {
	// CompletionMedianSeconds specifies the median elapsed time from creation to full
	// readiness. Null when no samples exist.
	CompletionMedianSeconds float64 `json:"completion_median_seconds" api:"required"`
	// CompletionSampleCount specifies Assets with a full readiness measurement.
	CompletionSampleCount int64 `json:"completion_sample_count" api:"required"`
	// PlaybackMedianSeconds specifies the median elapsed time from creation to
	// playback readiness. Null when no samples exist.
	PlaybackMedianSeconds float64 `json:"playback_median_seconds" api:"required"`
	// PlaybackSampleCount specifies Assets with a playback readiness measurement.
	PlaybackSampleCount int64 `json:"playback_sample_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionMedianSeconds respjson.Field
		CompletionSampleCount   respjson.Field
		PlaybackMedianSeconds   respjson.Field
		PlaybackSampleCount     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsAssetsSummaryTimings) RawJSON() string { return r.JSON.raw }
func (r *MetricsAssetsSummaryTimings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recorded additions and successful removals within event retention. Excludes
// external storage changes and does not measure current storage size.
type MetricsStorage struct {
	// From specifies the inclusive UTC start date.
	From string `json:"from" api:"required"`
	// Storages specifies persistent project storage configurations, including those
	// without recorded activity. Chunkify temporary storage is excluded.
	Storages []MetricsStorageStorage `json:"storages" api:"required"`
	// To specifies the exclusive UTC end date.
	To string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Storages    respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsStorage) RawJSON() string { return r.JSON.raw }
func (r *MetricsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Activity for one persistent storage configuration.
type MetricsStorageStorage struct {
	// Daily specifies one record for each selected event day, oldest first.
	Daily []MetricsStorageStorageDaily `json:"daily" api:"required"`
	// StorageId identifies the persistent storage configuration.
	StorageID string `json:"storage_id" api:"required"`
	// Summary specifies recorded activity in the selected period.
	Summary MetricsStorageStorageSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Daily       respjson.Field
		StorageID   respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsStorageStorage) RawJSON() string { return r.JSON.raw }
func (r *MetricsStorageStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recorded storage activity grouped by event day.
type MetricsStorageStorageDaily struct {
	// AssetBytesAdded specifies recorded Asset package additions.
	AssetBytesAdded int64 `json:"asset_bytes_added" api:"required"`
	// BytesRemoved specifies recorded successful file removals as a positive amount.
	BytesRemoved int64 `json:"bytes_removed" api:"required"`
	// Date specifies the UTC event day.
	Date string `json:"date" api:"required"`
	// NetBytesChange specifies additions minus successful removals. This may be
	// negative and is not bucket occupancy.
	NetBytesChange int64 `json:"net_bytes_change" api:"required"`
	// SourceBytesAdded specifies recorded Source additions.
	SourceBytesAdded int64 `json:"source_bytes_added" api:"required"`
	// StaticRenditionBytesAdded specifies recorded static rendition additions.
	StaticRenditionBytesAdded int64 `json:"static_rendition_bytes_added" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetBytesAdded           respjson.Field
		BytesRemoved              respjson.Field
		Date                      respjson.Field
		NetBytesChange            respjson.Field
		SourceBytesAdded          respjson.Field
		StaticRenditionBytesAdded respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsStorageStorageDaily) RawJSON() string { return r.JSON.raw }
func (r *MetricsStorageStorageDaily) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary specifies recorded activity in the selected period.
type MetricsStorageStorageSummary struct {
	// AssetBytesAdded specifies recorded Asset package additions.
	AssetBytesAdded int64 `json:"asset_bytes_added" api:"required"`
	// BytesRemoved specifies recorded successful file removals as a positive amount.
	BytesRemoved int64 `json:"bytes_removed" api:"required"`
	// NetBytesChange specifies additions minus successful removals. This may be
	// negative and is not bucket occupancy.
	NetBytesChange int64 `json:"net_bytes_change" api:"required"`
	// SourceBytesAdded specifies recorded Source additions.
	SourceBytesAdded int64 `json:"source_bytes_added" api:"required"`
	// StaticRenditionBytesAdded specifies recorded static rendition additions.
	StaticRenditionBytesAdded int64 `json:"static_rendition_bytes_added" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetBytesAdded           respjson.Field
		BytesRemoved              respjson.Field
		NetBytesChange            respjson.Field
		SourceBytesAdded          respjson.Field
		StaticRenditionBytesAdded respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsStorageStorageSummary) RawJSON() string { return r.JSON.raw }
func (r *MetricsStorageStorageSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Project usage for the selected UTC accounting period.
type MetricsUsage struct {
	// Daily specifies one record for every selected day, oldest first.
	Daily []MetricsUsageDaily `json:"daily" api:"required"`
	// From specifies the inclusive UTC start date.
	From string `json:"from" api:"required"`
	// Summary specifies totals across the selected period.
	Summary MetricsUsageSummary `json:"summary" api:"required"`
	// To specifies the exclusive UTC end date.
	To string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Daily       respjson.Field
		From        respjson.Field
		Summary     respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsage) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Daily recorded usage. Inactive days contain zero counters.
type MetricsUsageDaily struct {
	// Asset processing usage.
	Assets MetricsUsageDailyAssets `json:"assets" api:"required"`
	// Caption processing usage.
	Captions MetricsUsageDailyCaptions `json:"captions" api:"required"`
	// Settled credits in integer units, with 10000 units equal to USD 1 of credits.
	// Category values are parts of total, not additional charges.
	CreditsSpent MetricsUsageDailyCreditsSpent `json:"credits_spent" api:"required"`
	// Date specifies the UTC accounting day.
	Date string `json:"date" api:"required"`
	// Recorded output egress eligible for the transfer tariff. Provider-exempt
	// transfers are excluded.
	Egress MetricsUsageDailyEgress `json:"egress" api:"required"`
	// Webhook delivery activity.
	Notifications MetricsUsageDailyNotifications `json:"notifications" api:"required"`
	// Source ingestion activity.
	Sources MetricsUsageDailySources `json:"sources" api:"required"`
	// Static rendition usage.
	StaticRenditions MetricsUsageDailyStaticRenditions `json:"static_renditions" api:"required"`
	// Thumbnail processing usage.
	Thumbnails MetricsUsageDailyThumbnails `json:"thumbnails" api:"required"`
	// Standalone Transcoding usage.
	Transcodings MetricsUsageDailyTranscodings `json:"transcodings" api:"required"`
	// Upload activity.
	Uploads MetricsUsageDailyUploads `json:"uploads" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets           respjson.Field
		Captions         respjson.Field
		CreditsSpent     respjson.Field
		Date             respjson.Field
		Egress           respjson.Field
		Notifications    respjson.Field
		Sources          respjson.Field
		StaticRenditions respjson.Field
		Thumbnails       respjson.Field
		Transcodings     respjson.Field
		Uploads          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDaily) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDaily) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Asset processing usage.
type MetricsUsageDailyAssets struct {
	// Failed specifies recorded failed Asset outcomes.
	Failed int64 `json:"failed" api:"required"`
	// InputSeconds specifies settled billable Asset input duration.
	InputSeconds int64 `json:"input_seconds" api:"required"`
	// Video duration in seconds by pricing resolution tier: sd, hd, 2k and uhd.
	InputSecondsByResolution map[string]int64 `json:"input_seconds_by_resolution" api:"required"`
	// Ready specifies recorded successful Asset outcomes.
	Ready int64 `json:"ready" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed                   respjson.Field
		InputSeconds             respjson.Field
		InputSecondsByResolution respjson.Field
		Ready                    respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyAssets) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caption processing usage.
type MetricsUsageDailyCaptions struct {
	// TranscriptionSeconds specifies audio duration transcribed.
	TranscriptionSeconds int64 `json:"transcription_seconds" api:"required"`
	// TranslationSeconds specifies caption duration translated, counted once per
	// target language.
	TranslationSeconds int64 `json:"translation_seconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TranscriptionSeconds respjson.Field
		TranslationSeconds   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyCaptions) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyCaptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settled credits in integer units, with 10000 units equal to USD 1 of credits.
// Category values are parts of total, not additional charges.
type MetricsUsageDailyCreditsSpent struct {
	// Amount specifies settled credits spent on assets. Integer units; 10000 units
	// equal USD 1 of credits.
	Assets int64 `json:"assets" api:"required"`
	// Amount specifies settled credits spent on egress. Integer units; 10000 units
	// equal USD 1 of credits.
	Egress int64 `json:"egress" api:"required"`
	// Amount specifies settled credits spent on static renditions. Integer units;
	// 10000 units equal USD 1 of credits.
	StaticRenditions int64 `json:"static_renditions" api:"required"`
	// Amount specifies settled credits spent on thumbnails. Integer units; 10000 units
	// equal USD 1 of credits.
	Thumbnails int64 `json:"thumbnails" api:"required"`
	// Total specifies the sum of these media-processing and output-egress charges.
	// Integer units; 10000 units equal USD 1 of credits.
	Total int64 `json:"total" api:"required"`
	// Amount specifies settled credits spent on transcodings. Integer units; 10000
	// units equal USD 1 of credits.
	Transcodings int64 `json:"transcodings" api:"required"`
	// Amount specifies settled credits spent on transcription. Integer units; 10000
	// units equal USD 1 of credits.
	Transcription int64 `json:"transcription" api:"required"`
	// Amount specifies settled credits spent on translation. Integer units; 10000
	// units equal USD 1 of credits.
	Translation int64 `json:"translation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets           respjson.Field
		Egress           respjson.Field
		StaticRenditions respjson.Field
		Thumbnails       respjson.Field
		Total            respjson.Field
		Transcodings     respjson.Field
		Transcription    respjson.Field
		Translation      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyCreditsSpent) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyCreditsSpent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recorded output egress eligible for the transfer tariff. Provider-exempt
// transfers are excluded.
type MetricsUsageDailyEgress struct {
	// Bytes specifies tariff-eligible output transfer bytes for Asset preparation.
	AssetPreparationBytes int64 `json:"asset_preparation_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for streaming renditions.
	AssetRenditionBytes int64 `json:"asset_rendition_bytes" api:"required"`
	// BillableBytes specifies external output bytes charged after the allowance.
	BillableBytes int64 `json:"billable_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for captions.
	CaptionBytes int64 `json:"caption_bytes" api:"required"`
	// IncludedBytes specifies external output bytes covered by the allowance.
	IncludedBytes int64 `json:"included_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for static renditions.
	StaticRenditionBytes int64 `json:"static_rendition_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for Thumbnails.
	ThumbnailBytes int64 `json:"thumbnail_bytes" api:"required"`
	// TotalBytes specifies the sum of tariff-eligible output transfers.
	TotalBytes int64 `json:"total_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for Transcodings.
	TranscodingBytes int64 `json:"transcoding_bytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetPreparationBytes respjson.Field
		AssetRenditionBytes   respjson.Field
		BillableBytes         respjson.Field
		CaptionBytes          respjson.Field
		IncludedBytes         respjson.Field
		StaticRenditionBytes  respjson.Field
		ThumbnailBytes        respjson.Field
		TotalBytes            respjson.Field
		TranscodingBytes      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyEgress) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyEgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook delivery activity.
type MetricsUsageDailyNotifications struct {
	// Failed specifies recorded failed webhook delivery attempts.
	Failed int64 `json:"failed" api:"required"`
	// Sent specifies recorded successful webhook deliveries.
	Sent int64 `json:"sent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Sent        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyNotifications) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyNotifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source ingestion activity.
type MetricsUsageDailySources struct {
	// Bytes specifies cumulative ingested Source bytes, not current storage occupancy.
	Bytes int64 `json:"bytes" api:"required"`
	// Created specifies the number of ingested Sources.
	Created int64 `json:"created" api:"required"`
	// IngestedSeconds specifies ingested media duration in seconds.
	IngestedSeconds int64 `json:"ingested_seconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes           respjson.Field
		Created         respjson.Field
		IngestedSeconds respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailySources) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailySources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Static rendition usage.
type MetricsUsageDailyStaticRenditions struct {
	// AudioOutputSeconds specifies the audio-only subset of output_seconds.
	AudioOutputSeconds int64 `json:"audio_output_seconds" api:"required"`
	// OutputSeconds specifies settled billable output duration, including audio-only
	// output.
	OutputSeconds int64 `json:"output_seconds" api:"required"`
	// Video duration in seconds by pricing resolution tier. Contains sd, hd, 2k and
	// uhd keys; audio is reported separately.
	OutputSecondsByResolution map[string]int64 `json:"output_seconds_by_resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioOutputSeconds        respjson.Field
		OutputSeconds             respjson.Field
		OutputSecondsByResolution respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyStaticRenditions) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyStaticRenditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Thumbnail processing usage.
type MetricsUsageDailyThumbnails struct {
	// ImagesPublished specifies the number of published thumbnail images.
	ImagesPublished int64 `json:"images_published" api:"required"`
	// SetsFailed specifies recorded failed Thumbnail requests.
	SetsFailed int64 `json:"sets_failed" api:"required"`
	// SetsReady specifies recorded successful Thumbnail requests.
	SetsReady int64 `json:"sets_ready" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImagesPublished respjson.Field
		SetsFailed      respjson.Field
		SetsReady       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyThumbnails) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyThumbnails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standalone Transcoding usage.
type MetricsUsageDailyTranscodings struct {
	// AudioOutputSeconds specifies the audio-only subset of output_seconds.
	AudioOutputSeconds int64 `json:"audio_output_seconds" api:"required"`
	// Failed specifies recorded failed Transcoding outcomes.
	Failed int64 `json:"failed" api:"required"`
	// OutputSeconds specifies settled billable output duration, including audio-only
	// output.
	OutputSeconds int64 `json:"output_seconds" api:"required"`
	// Video duration in seconds by pricing resolution tier. Contains sd, hd, 2k and
	// uhd keys; audio is reported separately.
	OutputSecondsByResolution map[string]int64 `json:"output_seconds_by_resolution" api:"required"`
	// Ready specifies recorded successful Transcoding outcomes.
	Ready int64 `json:"ready" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioOutputSeconds        respjson.Field
		Failed                    respjson.Field
		OutputSeconds             respjson.Field
		OutputSecondsByResolution respjson.Field
		Ready                     respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyTranscodings) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyTranscodings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upload activity.
type MetricsUsageDailyUploads struct {
	// Bytes specifies successfully uploaded bytes.
	Bytes int64 `json:"bytes" api:"required"`
	// Completed specifies successfully completed Uploads.
	Completed int64 `json:"completed" api:"required"`
	// Failed specifies recorded failed Uploads.
	Failed int64 `json:"failed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Completed   respjson.Field
		Failed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageDailyUploads) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageDailyUploads) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary specifies totals across the selected period.
type MetricsUsageSummary struct {
	// Asset processing usage.
	Assets MetricsUsageSummaryAssets `json:"assets" api:"required"`
	// Caption processing usage.
	Captions MetricsUsageSummaryCaptions `json:"captions" api:"required"`
	// Settled credits in integer units, with 10000 units equal to USD 1 of credits.
	// Category values are parts of total, not additional charges.
	CreditsSpent MetricsUsageSummaryCreditsSpent `json:"credits_spent" api:"required"`
	// Recorded output egress eligible for the transfer tariff. Provider-exempt
	// transfers are excluded.
	Egress MetricsUsageSummaryEgress `json:"egress" api:"required"`
	// Webhook delivery activity.
	Notifications MetricsUsageSummaryNotifications `json:"notifications" api:"required"`
	// Source ingestion activity.
	Sources MetricsUsageSummarySources `json:"sources" api:"required"`
	// Static rendition usage.
	StaticRenditions MetricsUsageSummaryStaticRenditions `json:"static_renditions" api:"required"`
	// Thumbnail processing usage.
	Thumbnails MetricsUsageSummaryThumbnails `json:"thumbnails" api:"required"`
	// Standalone Transcoding usage.
	Transcodings MetricsUsageSummaryTranscodings `json:"transcodings" api:"required"`
	// Upload activity.
	Uploads MetricsUsageSummaryUploads `json:"uploads" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets           respjson.Field
		Captions         respjson.Field
		CreditsSpent     respjson.Field
		Egress           respjson.Field
		Notifications    respjson.Field
		Sources          respjson.Field
		StaticRenditions respjson.Field
		Thumbnails       respjson.Field
		Transcodings     respjson.Field
		Uploads          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummary) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Asset processing usage.
type MetricsUsageSummaryAssets struct {
	// Failed specifies recorded failed Asset outcomes.
	Failed int64 `json:"failed" api:"required"`
	// InputSeconds specifies settled billable Asset input duration.
	InputSeconds int64 `json:"input_seconds" api:"required"`
	// Video duration in seconds by pricing resolution tier: sd, hd, 2k and uhd.
	InputSecondsByResolution map[string]int64 `json:"input_seconds_by_resolution" api:"required"`
	// Ready specifies recorded successful Asset outcomes.
	Ready int64 `json:"ready" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed                   respjson.Field
		InputSeconds             respjson.Field
		InputSecondsByResolution respjson.Field
		Ready                    respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryAssets) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Caption processing usage.
type MetricsUsageSummaryCaptions struct {
	// TranscriptionSeconds specifies audio duration transcribed.
	TranscriptionSeconds int64 `json:"transcription_seconds" api:"required"`
	// TranslationSeconds specifies caption duration translated, counted once per
	// target language.
	TranslationSeconds int64 `json:"translation_seconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TranscriptionSeconds respjson.Field
		TranslationSeconds   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryCaptions) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryCaptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settled credits in integer units, with 10000 units equal to USD 1 of credits.
// Category values are parts of total, not additional charges.
type MetricsUsageSummaryCreditsSpent struct {
	// Amount specifies settled credits spent on assets. Integer units; 10000 units
	// equal USD 1 of credits.
	Assets int64 `json:"assets" api:"required"`
	// Amount specifies settled credits spent on egress. Integer units; 10000 units
	// equal USD 1 of credits.
	Egress int64 `json:"egress" api:"required"`
	// Amount specifies settled credits spent on static renditions. Integer units;
	// 10000 units equal USD 1 of credits.
	StaticRenditions int64 `json:"static_renditions" api:"required"`
	// Amount specifies settled credits spent on thumbnails. Integer units; 10000 units
	// equal USD 1 of credits.
	Thumbnails int64 `json:"thumbnails" api:"required"`
	// Total specifies the sum of these media-processing and output-egress charges.
	// Integer units; 10000 units equal USD 1 of credits.
	Total int64 `json:"total" api:"required"`
	// Amount specifies settled credits spent on transcodings. Integer units; 10000
	// units equal USD 1 of credits.
	Transcodings int64 `json:"transcodings" api:"required"`
	// Amount specifies settled credits spent on transcription. Integer units; 10000
	// units equal USD 1 of credits.
	Transcription int64 `json:"transcription" api:"required"`
	// Amount specifies settled credits spent on translation. Integer units; 10000
	// units equal USD 1 of credits.
	Translation int64 `json:"translation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets           respjson.Field
		Egress           respjson.Field
		StaticRenditions respjson.Field
		Thumbnails       respjson.Field
		Total            respjson.Field
		Transcodings     respjson.Field
		Transcription    respjson.Field
		Translation      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryCreditsSpent) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryCreditsSpent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recorded output egress eligible for the transfer tariff. Provider-exempt
// transfers are excluded.
type MetricsUsageSummaryEgress struct {
	// Bytes specifies tariff-eligible output transfer bytes for Asset preparation.
	AssetPreparationBytes int64 `json:"asset_preparation_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for streaming renditions.
	AssetRenditionBytes int64 `json:"asset_rendition_bytes" api:"required"`
	// BillableBytes specifies external output bytes charged after the allowance.
	BillableBytes int64 `json:"billable_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for captions.
	CaptionBytes int64 `json:"caption_bytes" api:"required"`
	// IncludedBytes specifies external output bytes covered by the allowance.
	IncludedBytes int64 `json:"included_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for static renditions.
	StaticRenditionBytes int64 `json:"static_rendition_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for Thumbnails.
	ThumbnailBytes int64 `json:"thumbnail_bytes" api:"required"`
	// TotalBytes specifies the sum of tariff-eligible output transfers.
	TotalBytes int64 `json:"total_bytes" api:"required"`
	// Bytes specifies tariff-eligible output transfer bytes for Transcodings.
	TranscodingBytes int64 `json:"transcoding_bytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetPreparationBytes respjson.Field
		AssetRenditionBytes   respjson.Field
		BillableBytes         respjson.Field
		CaptionBytes          respjson.Field
		IncludedBytes         respjson.Field
		StaticRenditionBytes  respjson.Field
		ThumbnailBytes        respjson.Field
		TotalBytes            respjson.Field
		TranscodingBytes      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryEgress) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryEgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook delivery activity.
type MetricsUsageSummaryNotifications struct {
	// Failed specifies recorded failed webhook delivery attempts.
	Failed int64 `json:"failed" api:"required"`
	// Sent specifies recorded successful webhook deliveries.
	Sent int64 `json:"sent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Sent        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryNotifications) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryNotifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source ingestion activity.
type MetricsUsageSummarySources struct {
	// Bytes specifies cumulative ingested Source bytes, not current storage occupancy.
	Bytes int64 `json:"bytes" api:"required"`
	// Created specifies the number of ingested Sources.
	Created int64 `json:"created" api:"required"`
	// IngestedSeconds specifies ingested media duration in seconds.
	IngestedSeconds int64 `json:"ingested_seconds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes           respjson.Field
		Created         respjson.Field
		IngestedSeconds respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummarySources) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummarySources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Static rendition usage.
type MetricsUsageSummaryStaticRenditions struct {
	// AudioOutputSeconds specifies the audio-only subset of output_seconds.
	AudioOutputSeconds int64 `json:"audio_output_seconds" api:"required"`
	// OutputSeconds specifies settled billable output duration, including audio-only
	// output.
	OutputSeconds int64 `json:"output_seconds" api:"required"`
	// Video duration in seconds by pricing resolution tier. Contains sd, hd, 2k and
	// uhd keys; audio is reported separately.
	OutputSecondsByResolution map[string]int64 `json:"output_seconds_by_resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioOutputSeconds        respjson.Field
		OutputSeconds             respjson.Field
		OutputSecondsByResolution respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryStaticRenditions) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryStaticRenditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Thumbnail processing usage.
type MetricsUsageSummaryThumbnails struct {
	// ImagesPublished specifies the number of published thumbnail images.
	ImagesPublished int64 `json:"images_published" api:"required"`
	// SetsFailed specifies recorded failed Thumbnail requests.
	SetsFailed int64 `json:"sets_failed" api:"required"`
	// SetsReady specifies recorded successful Thumbnail requests.
	SetsReady int64 `json:"sets_ready" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImagesPublished respjson.Field
		SetsFailed      respjson.Field
		SetsReady       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryThumbnails) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryThumbnails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standalone Transcoding usage.
type MetricsUsageSummaryTranscodings struct {
	// AudioOutputSeconds specifies the audio-only subset of output_seconds.
	AudioOutputSeconds int64 `json:"audio_output_seconds" api:"required"`
	// Failed specifies recorded failed Transcoding outcomes.
	Failed int64 `json:"failed" api:"required"`
	// OutputSeconds specifies settled billable output duration, including audio-only
	// output.
	OutputSeconds int64 `json:"output_seconds" api:"required"`
	// Video duration in seconds by pricing resolution tier. Contains sd, hd, 2k and
	// uhd keys; audio is reported separately.
	OutputSecondsByResolution map[string]int64 `json:"output_seconds_by_resolution" api:"required"`
	// Ready specifies recorded successful Transcoding outcomes.
	Ready int64 `json:"ready" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioOutputSeconds        respjson.Field
		Failed                    respjson.Field
		OutputSeconds             respjson.Field
		OutputSecondsByResolution respjson.Field
		Ready                     respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryTranscodings) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryTranscodings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upload activity.
type MetricsUsageSummaryUploads struct {
	// Bytes specifies successfully uploaded bytes.
	Bytes int64 `json:"bytes" api:"required"`
	// Completed specifies successfully completed Uploads.
	Completed int64 `json:"completed" api:"required"`
	// Failed specifies recorded failed Uploads.
	Failed int64 `json:"failed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bytes       respjson.Field
		Completed   respjson.Field
		Failed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsUsageSummaryUploads) RawJSON() string { return r.JSON.raw }
func (r *MetricsUsageSummaryUploads) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetricAssetsParams struct {
	// Inclusive UTC start date, YYYY-MM-DD. Supply both dates for a custom range of at
	// most 366 days. Omit both for the past 7 UTC calendar days including today.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Exclusive UTC end date, YYYY-MM-DD. Must be after from and at most 366 days
	// later.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Response format. Defaults to JSON unless Accept requests CSV. Overrides Accept
	// when supplied; errors remain JSON.
	//
	// Any of "json", "csv".
	Format MetricAssetsParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MetricAssetsParams]'s query parameters as `url.Values`.
func (r MetricAssetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response format. Defaults to JSON unless Accept requests CSV. Overrides Accept
// when supplied; errors remain JSON.
type MetricAssetsParamsFormat string

const (
	MetricAssetsParamsFormatJson MetricAssetsParamsFormat = "json"
	MetricAssetsParamsFormatCsv  MetricAssetsParamsFormat = "csv"
)

// Successful metrics response.
type MetricAssetsResponseEnvelope struct {
	// Data contains the requested report.
	Data MetricsAssets `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricAssetsResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *MetricAssetsResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetricStorageParams struct {
	// Inclusive UTC start date, YYYY-MM-DD. Supply both dates for a custom range of at
	// most 366 days. Omit both for the past 7 UTC calendar days including today.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Filter to one persistent storage configuration in the authenticated project.
	StorageID param.Opt[string] `query:"storage_id,omitzero" json:"-"`
	// Exclusive UTC end date, YYYY-MM-DD. Must be after from and at most 366 days
	// later.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Response format. Defaults to JSON unless Accept requests CSV. Overrides Accept
	// when supplied; errors remain JSON.
	//
	// Any of "json", "csv".
	Format MetricStorageParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MetricStorageParams]'s query parameters as `url.Values`.
func (r MetricStorageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response format. Defaults to JSON unless Accept requests CSV. Overrides Accept
// when supplied; errors remain JSON.
type MetricStorageParamsFormat string

const (
	MetricStorageParamsFormatJson MetricStorageParamsFormat = "json"
	MetricStorageParamsFormatCsv  MetricStorageParamsFormat = "csv"
)

// Successful metrics response.
type MetricStorageResponseEnvelope struct {
	// Data contains the requested report.
	Data MetricsStorage `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricStorageResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *MetricStorageResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetricUsageParams struct {
	// Inclusive UTC start date, YYYY-MM-DD. Supply both dates for a custom range of at
	// most 366 days. Omit both for the past 7 UTC calendar days including today.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Exclusive UTC end date, YYYY-MM-DD. Must be after from and at most 366 days
	// later.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Response format. Defaults to JSON unless Accept requests CSV. Overrides Accept
	// when supplied; errors remain JSON.
	//
	// Any of "json", "csv".
	Format MetricUsageParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MetricUsageParams]'s query parameters as `url.Values`.
func (r MetricUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response format. Defaults to JSON unless Accept requests CSV. Overrides Accept
// when supplied; errors remain JSON.
type MetricUsageParamsFormat string

const (
	MetricUsageParamsFormatJson MetricUsageParamsFormat = "json"
	MetricUsageParamsFormatCsv  MetricUsageParamsFormat = "csv"
)

// Successful metrics response.
type MetricUsageResponseEnvelope struct {
	// Data contains the requested report.
	Data MetricsUsage `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricUsageResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *MetricUsageResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
