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

func TestFileGet(t *testing.T) {
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
	_, err := client.Files.Get(context.TODO(), "fileId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
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
	client := githubcomchunkifydevchunkifygo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Files.List(context.TODO(), githubcomchunkifydevchunkifygo.FileListParams{
		ID:         githubcomchunkifydevchunkifygo.String("id"),
		AudioCodec: githubcomchunkifydevchunkifygo.String("audio_codec"),
		Created: githubcomchunkifydevchunkifygo.FileListParamsCreated{
			Gte:  githubcomchunkifydevchunkifygo.String("gte"),
			Lte:  githubcomchunkifydevchunkifygo.String("lte"),
			Sort: "asc",
		},
		Duration: githubcomchunkifydevchunkifygo.FileListParamsDuration{
			Eq:  githubcomchunkifydevchunkifygo.Float(0),
			Gt:  githubcomchunkifydevchunkifygo.Float(0),
			Gte: githubcomchunkifydevchunkifygo.Float(0),
			Lt:  githubcomchunkifydevchunkifygo.Float(0),
			Lte: githubcomchunkifydevchunkifygo.Float(0),
		},
		Height: githubcomchunkifydevchunkifygo.FileListParamsHeight{
			Eq:  githubcomchunkifydevchunkifygo.Int(0),
			Gt:  githubcomchunkifydevchunkifygo.Int(0),
			Gte: githubcomchunkifydevchunkifygo.Int(0),
			Lt:  githubcomchunkifydevchunkifygo.Int(0),
			Lte: githubcomchunkifydevchunkifygo.Int(0),
		},
		JobID:    githubcomchunkifydevchunkifygo.String("job_id"),
		Limit:    githubcomchunkifydevchunkifygo.Int(1),
		MimeType: githubcomchunkifydevchunkifygo.String("mime_type"),
		Offset:   githubcomchunkifydevchunkifygo.Int(0),
		Path: githubcomchunkifydevchunkifygo.FileListParamsPath{
			Eq:    githubcomchunkifydevchunkifygo.String("eq"),
			Ilike: githubcomchunkifydevchunkifygo.String("ilike"),
		},
		Size: githubcomchunkifydevchunkifygo.FileListParamsSize{
			Eq:  githubcomchunkifydevchunkifygo.Int(0),
			Gt:  githubcomchunkifydevchunkifygo.Int(0),
			Gte: githubcomchunkifydevchunkifygo.Int(0),
			Lt:  githubcomchunkifydevchunkifygo.Int(0),
			Lte: githubcomchunkifydevchunkifygo.Int(0),
		},
		StorageID:  githubcomchunkifydevchunkifygo.String("storage_id"),
		VideoCodec: githubcomchunkifydevchunkifygo.String("video_codec"),
		Width: githubcomchunkifydevchunkifygo.FileListParamsWidth{
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

func TestFileDelete(t *testing.T) {
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
	err := client.Files.Delete(context.TODO(), "fileId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
