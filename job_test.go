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
		Format: chunkify.JobNewParamsFormatUnion{
			OfJobsHlsAv1: &chunkify.JobNewParamsFormatJobsHlsAv1{
				HlsAv1Param: chunkify.HlsAv1Param{
					AudioBitrate:   chunkify.Int(32000),
					Bufsize:        chunkify.Int(100000),
					Channels:       1,
					Crf:            chunkify.Int(35),
					DisableAudio:   chunkify.Bool(true),
					DisableVideo:   chunkify.Bool(true),
					Duration:       chunkify.Int(1),
					Framerate:      chunkify.Float(15),
					Gop:            chunkify.Int(1),
					Height:         chunkify.Int(-2),
					HlsEnc:         chunkify.Bool(false),
					HlsEncIv:       chunkify.String("0123456789abcdef"),
					HlsEncKey:      chunkify.String("0123456789abcdef"),
					HlsEncKeyURL:   chunkify.String("https://example.com/key"),
					HlsSegmentType: chunkify.HlsAv1HlsSegmentTypeMpegts,
					HlsTime:        chunkify.Int(6),
					Level:          41,
					Maxrate:        chunkify.Int(100000),
					Minrate:        chunkify.Int(100000),
					Movflags:       chunkify.String("movflags"),
					Pixfmt:         chunkify.HlsAv1PixfmtYuv410p,
					Preset:         chunkify.HlsAv1Preset10,
					Profilev:       chunkify.HlsAv1ProfilevMain10,
					Seek:           chunkify.Int(1),
					VideoBitrate:   chunkify.Int(100000),
					Width:          chunkify.Int(-2),
				},
				Name: chunkify.String("mp4/h264"),
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
	err := client.Jobs.Delete(context.TODO(), "jobId")
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
	err := client.Jobs.Cancel(context.TODO(), "jobId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
