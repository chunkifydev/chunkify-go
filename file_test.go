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

func TestFileGet(t *testing.T) {
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
	_, err := client.Files.Get(context.TODO(), "fileId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.List(context.TODO(), chunkify.FileListParams{
		ID:         chunkify.String("id"),
		AudioCodec: chunkify.String("audio_codec"),
		Created: chunkify.FileListParamsCreated{
			Gte:  chunkify.String("gte"),
			Lte:  chunkify.String("lte"),
			Sort: chunkify.String("sort"),
		},
		Duration: chunkify.FileListParamsDuration{
			Eq:  chunkify.Float(0),
			Gt:  chunkify.Float(0),
			Gte: chunkify.Float(0),
			Lt:  chunkify.Float(0),
			Lte: chunkify.Float(0),
		},
		Height: chunkify.FileListParamsHeight{
			Eq:  chunkify.Int(0),
			Gt:  chunkify.Int(0),
			Gte: chunkify.Int(0),
			Lt:  chunkify.Int(0),
			Lte: chunkify.Int(0),
		},
		JobID:    chunkify.String("job_id"),
		Limit:    chunkify.Int(0),
		MimeType: chunkify.String("mime_type"),
		Offset:   chunkify.Int(0),
		Path: chunkify.FileListParamsPath{
			Eq:    chunkify.String("eq"),
			Ilike: chunkify.String("ilike"),
		},
		Size: chunkify.FileListParamsSize{
			Eq:  chunkify.Int(0),
			Gt:  chunkify.Int(0),
			Gte: chunkify.Int(0),
			Lt:  chunkify.Int(0),
			Lte: chunkify.Int(0),
		},
		StorageID:  chunkify.String("storage_id"),
		VideoCodec: chunkify.String("video_codec"),
		Width: chunkify.FileListParamsWidth{
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

func TestFileDelete(t *testing.T) {
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
	err := client.Files.Delete(context.TODO(), "fileId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
