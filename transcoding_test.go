// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/chunkifydev/chunkify-go"
	"github.com/chunkifydev/chunkify-go/internal/testutil"
	"github.com/chunkifydev/chunkify-go/option"
)

func TestTranscodingNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		option.WithTeamAccessToken("My Team Access Token"),
	)
	_, err := client.Transcodings.New(context.TODO(), chunkify.TranscodingNewParams{
		Format: chunkify.TranscodingNewParamsFormatUnion{
			OfMP4H264: &chunkify.TranscodingMP4H264Param{
				AudioBitrate:    chunkify.Int(32000),
				AudioSamplerate: 8000,
				Bufsize:         chunkify.Int(100000),
				Channels:        1,
				Crf:             chunkify.Int(23),
				DisableAudio:    chunkify.Bool(true),
				DisableVideo:    chunkify.Bool(true),
				Duration:        chunkify.Float(1),
				Framerate:       chunkify.Float(0.001),
				Gop:             chunkify.Int(1),
				Height:          chunkify.Int(720),
				Level:           41,
				Maxrate:         chunkify.Int(100000),
				Movflags:        chunkify.TranscodingMP4H264MovflagsFaststart,
				PerTitle:        chunkify.Bool(true),
				Pixfmt:          chunkify.TranscodingMP4H264PixfmtYuv420p,
				Preset:          chunkify.TranscodingMP4H264PresetMedium,
				Profilev:        chunkify.TranscodingMP4H264ProfilevHigh,
				Seek:            chunkify.Float(0),
				VideoBitrate:    chunkify.Int(1000),
				Width:           chunkify.Int(-2),
				X264Keyint:      chunkify.Int(60),
			},
		},
		SourceID: "src_example",
		Metadata: map[string]string{
			"foo": "string",
		},
		Storage: chunkify.TranscodingNewParamsStorage{
			ID:   chunkify.String("stor_example"),
			Path: chunkify.String("exports/video.mp4"),
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

func TestTranscodingGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		option.WithTeamAccessToken("My Team Access Token"),
	)
	_, err := client.Transcodings.Get(context.TODO(), "transcodingId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTranscodingListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		option.WithTeamAccessToken("My Team Access Token"),
	)
	_, err := client.Transcodings.List(context.TODO(), chunkify.TranscodingListParams{
		ID: chunkify.String("x"),
		Created: chunkify.TranscodingListParamsCreated{
			Gte:  chunkify.Int(0),
			Lte:  chunkify.Int(0),
			Sort: "asc",
		},
		Limit:     chunkify.Int(1),
		Offset:    chunkify.Int(0),
		SourceID:  chunkify.String("x"),
		Status:    chunkify.TranscodingListParamsStatusPreparing,
		StorageID: chunkify.String("x"),
	})
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTranscodingDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		option.WithTeamAccessToken("My Team Access Token"),
	)
	err := client.Transcodings.Delete(
		context.TODO(),
		"transcodingId",
		chunkify.TranscodingDeleteParams{
			DeleteFiles: chunkify.Bool(true),
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

func TestTranscodingCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
		option.WithTeamAccessToken("My Team Access Token"),
	)
	err := client.Transcodings.Cancel(context.TODO(), "transcodingId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
