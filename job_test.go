// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomchunkifydevchunkifygo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/chunkifydev/chunkify-go"
	"github.com/chunkifydev/chunkify-go/internal/testutil"
	"github.com/chunkifydev/chunkify-go/option"
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
	client := githubcomchunkifydevchunkifygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.New(context.TODO(), githubcomchunkifydevchunkifygo.JobNewParams{
		Format: githubcomchunkifydevchunkifygo.JobNewParamsFormatUnion{
			OfMP4Av1: &githubcomchunkifydevchunkifygo.MP4Av1Param{
				AudioBitrate: githubcomchunkifydevchunkifygo.Int(32000),
				Bufsize:      githubcomchunkifydevchunkifygo.Int(100000),
				Channels:     1,
				Crf:          githubcomchunkifydevchunkifygo.Int(35),
				DisableAudio: githubcomchunkifydevchunkifygo.Bool(true),
				DisableVideo: githubcomchunkifydevchunkifygo.Bool(true),
				Duration:     githubcomchunkifydevchunkifygo.Int(1),
				Framerate:    githubcomchunkifydevchunkifygo.Float(15),
				Gop:          githubcomchunkifydevchunkifygo.Int(1),
				Height:       githubcomchunkifydevchunkifygo.Int(-2),
				Level:        41,
				Maxrate:      githubcomchunkifydevchunkifygo.Int(100000),
				Minrate:      githubcomchunkifydevchunkifygo.Int(100000),
				Movflags:     githubcomchunkifydevchunkifygo.String("movflags"),
				Pixfmt:       githubcomchunkifydevchunkifygo.MP4Av1PixfmtYuv410p,
				Preset:       githubcomchunkifydevchunkifygo.MP4Av1Preset10,
				Profilev:     githubcomchunkifydevchunkifygo.MP4Av1ProfilevMain10,
				Seek:         githubcomchunkifydevchunkifygo.Int(1),
				VideoBitrate: githubcomchunkifydevchunkifygo.Int(100000),
				Width:        githubcomchunkifydevchunkifygo.Int(-2),
			},
		},
		SourceID:      "src_UioP9I876hjKlNBH78ILp0mo56t",
		HlsManifestID: githubcomchunkifydevchunkifygo.String("hls_2v6EIgcNAycdS5g0IUm0TXBjvHV"),
		Metadata: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		Storage: githubcomchunkifydevchunkifygo.JobNewParamsStorage{
			ID:   githubcomchunkifydevchunkifygo.String("aws-my-storage"),
			Path: githubcomchunkifydevchunkifygo.String("/path/to/video.mp4"),
		},
		Transcoder: githubcomchunkifydevchunkifygo.JobNewParamsTranscoder{
			Quantity: githubcomchunkifydevchunkifygo.Int(2),
			Type:     "4vCPU",
		},
	})
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
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
	client := githubcomchunkifydevchunkifygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.Get(context.TODO(), "jobId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
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
	client := githubcomchunkifydevchunkifygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Jobs.List(context.TODO(), githubcomchunkifydevchunkifygo.JobListParams{
		ID: githubcomchunkifydevchunkifygo.String("id"),
		Created: githubcomchunkifydevchunkifygo.JobListParamsCreated{
			Gte:  githubcomchunkifydevchunkifygo.String("2025-01-01"),
			Lte:  githubcomchunkifydevchunkifygo.String("2025-01-01"),
			Sort: "asc",
		},
		FormatID:      githubcomchunkifydevchunkifygo.JobListParamsFormatIDMP4H264,
		HlsManifestID: githubcomchunkifydevchunkifygo.String("hls_manifest_id"),
		Limit:         githubcomchunkifydevchunkifygo.Int(1),
		Metadata:      [][]string{{"key1:value1"}},
		Offset:        githubcomchunkifydevchunkifygo.Int(0),
		SourceID:      githubcomchunkifydevchunkifygo.String("source_id"),
		Status:        githubcomchunkifydevchunkifygo.JobListParamsStatusCompleted,
	})
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
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
	client := githubcomchunkifydevchunkifygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	err := client.Jobs.Delete(context.TODO(), "jobId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
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
	client := githubcomchunkifydevchunkifygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	err := client.Jobs.Cancel(context.TODO(), "jobId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
