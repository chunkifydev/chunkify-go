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

func TestSourceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sources.New(context.TODO(), chunkify.SourceNewParams{
		URL: "https://example.com/video.mp4",
		Metadata: map[string]string{
			"key":  "value",
			"key2": "value2",
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

func TestSourceGet(t *testing.T) {
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
	_, err := client.Sources.Get(context.TODO(), "sourceId")
	if err != nil {
		var apierr *chunkify.Error
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
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Sources.List(context.TODO(), chunkify.SourceListParams{
		ID:         chunkify.String("id"),
		AudioCodec: chunkify.String("audio_codec"),
		Created: chunkify.SourceListParamsCreated{
			Gte:  chunkify.String("gte"),
			Lte:  chunkify.String("lte"),
			Sort: chunkify.String("sort"),
		},
		Device: chunkify.String("device"),
		Duration: chunkify.SourceListParamsDuration{
			Eq:  chunkify.Float(0),
			Gt:  chunkify.Float(0),
			Gte: chunkify.Float(0),
			Lt:  chunkify.Float(0),
			Lte: chunkify.Float(0),
		},
		Height: chunkify.SourceListParamsHeight{
			Eq:  chunkify.Int(0),
			Gt:  chunkify.Int(0),
			Gte: chunkify.Int(0),
			Lt:  chunkify.Int(0),
			Lte: chunkify.Int(0),
		},
		Limit:    chunkify.Int(0),
		Metadata: chunkify.String("metadata"),
		Offset:   chunkify.Int(0),
		Size: chunkify.SourceListParamsSize{
			Eq:  chunkify.Int(0),
			Gt:  chunkify.Int(0),
			Gte: chunkify.Int(0),
			Lt:  chunkify.Int(0),
			Lte: chunkify.Int(0),
		},
		VideoCodec: chunkify.String("video_codec"),
		Width: chunkify.SourceListParamsWidth{
			Eq:  chunkify.Int(0),
			Gt:  chunkify.Int(0),
			Gte: chunkify.Int(0),
			Lt:  chunkify.Int(0),
			Lte: chunkify.Int(0),
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

func TestSourceDelete(t *testing.T) {
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
	_, err := client.Sources.Delete(context.TODO(), "sourceId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
