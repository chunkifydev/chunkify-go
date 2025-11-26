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

func TestSourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sources.New(context.TODO(), githubcomchunkifydevchunkifygo.SourceNewParams{
		URL: "https://example.com/video.mp4",
		Metadata: map[string]string{
			"key":  "value",
			"key2": "value2",
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

func TestSourceGet(t *testing.T) {
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
	_, err := client.Sources.Get(context.TODO(), "sourceId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sources.List(context.TODO(), githubcomchunkifydevchunkifygo.SourceListParams{
		ID:         githubcomchunkifydevchunkifygo.String("id"),
		AudioCodec: githubcomchunkifydevchunkifygo.String("audio_codec"),
		Created: githubcomchunkifydevchunkifygo.SourceListParamsCreated{
			Gte:  githubcomchunkifydevchunkifygo.String("2025-01-01"),
			Lte:  githubcomchunkifydevchunkifygo.String("2025-01-01"),
			Sort: "asc",
		},
		Device: githubcomchunkifydevchunkifygo.SourceListParamsDeviceApple,
		Duration: githubcomchunkifydevchunkifygo.SourceListParamsDuration{
			Eq:  githubcomchunkifydevchunkifygo.Float(0),
			Gt:  githubcomchunkifydevchunkifygo.Float(0),
			Gte: githubcomchunkifydevchunkifygo.Float(0),
			Lt:  githubcomchunkifydevchunkifygo.Float(0),
			Lte: githubcomchunkifydevchunkifygo.Float(0),
		},
		Height: githubcomchunkifydevchunkifygo.SourceListParamsHeight{
			Eq:  githubcomchunkifydevchunkifygo.Int(0),
			Gt:  githubcomchunkifydevchunkifygo.Int(0),
			Gte: githubcomchunkifydevchunkifygo.Int(0),
			Lt:  githubcomchunkifydevchunkifygo.Int(0),
			Lte: githubcomchunkifydevchunkifygo.Int(0),
		},
		Limit:    githubcomchunkifydevchunkifygo.Int(1),
		Metadata: [][]string{{"J!Q0Ok0bzJb7:pro"}},
		Offset:   githubcomchunkifydevchunkifygo.Int(0),
		Size: githubcomchunkifydevchunkifygo.SourceListParamsSize{
			Eq:  githubcomchunkifydevchunkifygo.Int(0),
			Gt:  githubcomchunkifydevchunkifygo.Int(0),
			Gte: githubcomchunkifydevchunkifygo.Int(0),
			Lt:  githubcomchunkifydevchunkifygo.Int(0),
			Lte: githubcomchunkifydevchunkifygo.Int(0),
		},
		VideoCodec: githubcomchunkifydevchunkifygo.String("video_codec"),
		Width: githubcomchunkifydevchunkifygo.SourceListParamsWidth{
			Eq:  githubcomchunkifydevchunkifygo.Int(0),
			Gt:  githubcomchunkifydevchunkifygo.Int(0),
			Gte: githubcomchunkifydevchunkifygo.Int(0),
			Lt:  githubcomchunkifydevchunkifygo.Int(0),
			Lte: githubcomchunkifydevchunkifygo.Int(0),
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

func TestSourceDelete(t *testing.T) {
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
	err := client.Sources.Delete(context.TODO(), "sourceId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
