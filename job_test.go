// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/chunkify-go"
	"github.com/stainless-sdks/chunkify-go/internal/testutil"
	"github.com/stainless-sdks/chunkify-go/option"
)

func TestJobNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.New(context.TODO(), chunkify.JobNewParams{
		Format: chunkify.JobNewParamsFormat{
			HlsAv1: chunkify.JobNewParamsFormatHlsAv1{
				HlsParam: chunkify.HlsParam{
					AudioBitrate:   32000,
					VideoBitrate:   100000,
					HlsEnc:         chunkify.Bool(false),
					HlsEncIv:       chunkify.String("0123456789abcdef"),
					HlsEncKey:      chunkify.String("0123456789abcdef"),
					HlsEncKeyURL:   chunkify.String("https://example.com/key"),
					HlsSegmentType: chunkify.HlsHlsSegmentTypeMpegts,
					HlsTime:        chunkify.Int(6),
				},
				Av1Param: chunkify.Av1Param{
					VideoCommonParam: chunkify.VideoCommonParam{},
				},
			},
			HlsH264: chunkify.JobNewParamsFormatHlsH264{
				HlsParam: chunkify.HlsParam{
					AudioBitrate:   32000,
					VideoBitrate:   100000,
					HlsEnc:         chunkify.Bool(false),
					HlsEncIv:       chunkify.String("0123456789abcdef"),
					HlsEncKey:      chunkify.String("0123456789abcdef"),
					HlsEncKeyURL:   chunkify.String("https://example.com/key"),
					HlsSegmentType: chunkify.HlsHlsSegmentTypeMpegts,
					HlsTime:        chunkify.Int(6),
				},
				H264Param: chunkify.H264Param{
					VideoCommonParam: chunkify.VideoCommonParam{},
				},
			},
			HlsH265: chunkify.JobNewParamsFormatHlsH265{
				HlsParam: chunkify.HlsParam{
					AudioBitrate:   32000,
					VideoBitrate:   100000,
					HlsEnc:         chunkify.Bool(false),
					HlsEncIv:       chunkify.String("0123456789abcdef"),
					HlsEncKey:      chunkify.String("0123456789abcdef"),
					HlsEncKeyURL:   chunkify.String("https://example.com/key"),
					HlsSegmentType: chunkify.HlsHlsSegmentTypeMpegts,
					HlsTime:        chunkify.Int(6),
				},
				H265Param: chunkify.H265Param{
					VideoCommonParam: chunkify.VideoCommonParam{},
				},
			},
			Jpg: chunkify.JobNewParamsFormatJpg{
				Interval:      1,
				ChunkDuration: chunkify.Int(1),
				Duration:      chunkify.Int(1),
				Frames:        chunkify.Int(1),
				Height:        chunkify.Int(-2),
				Seek:          chunkify.Int(1),
				Sprite:        chunkify.Bool(true),
				Width:         chunkify.Int(-2),
			},
			MP4Av1: chunkify.Av1Param{
				VideoCommonParam: chunkify.VideoCommonParam{
					AudioBitrate: chunkify.Int(32000),
					Bufsize:      chunkify.Int(100000),
					Channels:     1,
					DisableAudio: chunkify.Bool(true),
					DisableVideo: chunkify.Bool(true),
					Duration:     chunkify.Int(1),
					Framerate:    chunkify.Float(15),
					Gop:          chunkify.Int(1),
					Height:       chunkify.Int(-2),
					Maxrate:      chunkify.Int(100000),
					Minrate:      chunkify.Int(100000),
					Pixfmt:       chunkify.VideoCommonPixfmtYuv410p,
					Seek:         chunkify.Int(1),
					VideoBitrate: chunkify.Int(100000),
					Width:        chunkify.Int(-2),
				},
				Crf:      chunkify.Int(35),
				Level:    41,
				Movflags: chunkify.String("movflags"),
				Preset:   "10",
				Profilev: "main10",
			},
			MP4H264: chunkify.H264Param{
				VideoCommonParam: chunkify.VideoCommonParam{
					AudioBitrate: chunkify.Int(32000),
					Bufsize:      chunkify.Int(100000),
					Channels:     1,
					DisableAudio: chunkify.Bool(true),
					DisableVideo: chunkify.Bool(true),
					Duration:     chunkify.Int(1),
					Framerate:    chunkify.Float(15),
					Gop:          chunkify.Int(1),
					Height:       chunkify.Int(-2),
					Maxrate:      chunkify.Int(100000),
					Minrate:      chunkify.Int(100000),
					Pixfmt:       chunkify.VideoCommonPixfmtYuv410p,
					Seek:         chunkify.Int(1),
					VideoBitrate: chunkify.Int(100000),
					Width:        chunkify.Int(-2),
				},
				Crf:        chunkify.Int(23),
				Level:      41,
				Movflags:   chunkify.String("movflags"),
				Preset:     "medium",
				Profilev:   "high",
				X264Keyint: chunkify.Int(60),
			},
			MP4H265: chunkify.H265Param{
				VideoCommonParam: chunkify.VideoCommonParam{
					AudioBitrate: chunkify.Int(32000),
					Bufsize:      chunkify.Int(100000),
					Channels:     1,
					DisableAudio: chunkify.Bool(true),
					DisableVideo: chunkify.Bool(true),
					Duration:     chunkify.Int(1),
					Framerate:    chunkify.Float(15),
					Gop:          chunkify.Int(1),
					Height:       chunkify.Int(-2),
					Maxrate:      chunkify.Int(100000),
					Minrate:      chunkify.Int(100000),
					Pixfmt:       chunkify.VideoCommonPixfmtYuv410p,
					Seek:         chunkify.Int(1),
					VideoBitrate: chunkify.Int(100000),
					Width:        chunkify.Int(-2),
				},
				Crf:        chunkify.Int(23),
				Level:      41,
				Movflags:   chunkify.String("movflags"),
				Preset:     "medium",
				Profilev:   "main10",
				X265Keyint: chunkify.Int(60),
			},
			WebmVp9: chunkify.VideoCommonParam{
				AudioBitrate: chunkify.Int(32000),
				Bufsize:      chunkify.Int(100000),
				Channels:     1,
				DisableAudio: chunkify.Bool(true),
				DisableVideo: chunkify.Bool(true),
				Duration:     chunkify.Int(1),
				Framerate:    chunkify.Float(15),
				Gop:          chunkify.Int(1),
				Height:       chunkify.Int(-2),
				Maxrate:      chunkify.Int(100000),
				Minrate:      chunkify.Int(100000),
				Pixfmt:       chunkify.VideoCommonPixfmtYuv410p,
				Seek:         chunkify.Int(1),
				VideoBitrate: chunkify.Int(100000),
				Width:        chunkify.Int(-2),
			},
		},
		SourceID:      "src_UioP9I876hjKlNBH78ILp0mo56t",
		HlsManifestID: chunkify.String("hls_2v6EIgcNAycdS5g0IUm0TXBjvHV"),
		Metadata: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		Storage: chunkify.JobNewParamsStorage{
			ID:   chunkify.String("aws-my-storage"),
			Path: chunkify.String("/path/to/video.mp4"),
		},
		Transcoder: chunkify.JobNewParamsTranscoder{
			Quantity: chunkify.Int(2),
			Type:     "8vCPU",
		},
	})
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.Get(context.TODO(), "jobId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.List(context.TODO(), chunkify.JobListParams{
		ID: chunkify.String("id"),
		Created: chunkify.JobListParamsCreated{
			Gte:  chunkify.String("gte"),
			Lte:  chunkify.String("lte"),
			Sort: chunkify.String("sort"),
		},
		FormatName:    chunkify.String("format_name"),
		HlsManifestID: chunkify.String("hls_manifest_id"),
		Limit:         chunkify.Int(0),
		Metadata:      chunkify.String("metadata"),
		Offset:        chunkify.Int(0),
		SourceID:      chunkify.String("source_id"),
		Status:        chunkify.String("status"),
	})
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.Delete(context.TODO(), "jobId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobCancel(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.Cancel(context.TODO(), "jobId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobGetFiles(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.GetFiles(context.TODO(), "jobId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobGetLogsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.GetLogs(
		context.TODO(),
		"jobId",
		chunkify.JobGetLogsParams{
			Service:      chunkify.JobGetLogsParamsServiceTranscoder,
			TranscoderID: chunkify.Int(0),
		},
	)
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobGetTranscoders(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.GetTranscoders(context.TODO(), "jobId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
