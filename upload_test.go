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

func TestUploadNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Uploads.New(context.TODO(), githubcomchunkifydevchunkifygo.UploadNewParams{
		Metadata: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		ValidityTimeout: githubcomchunkifydevchunkifygo.Int(3600),
	})
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUploadGet(t *testing.T) {
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
	_, err := client.Uploads.Get(context.TODO(), "uploadId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUploadListWithOptionalParams(t *testing.T) {
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
	_, err := client.Uploads.List(context.TODO(), githubcomchunkifydevchunkifygo.UploadListParams{
		ID: githubcomchunkifydevchunkifygo.String("id"),
		Created: githubcomchunkifydevchunkifygo.UploadListParamsCreated{
			Gte:  githubcomchunkifydevchunkifygo.String("gte"),
			Lte:  githubcomchunkifydevchunkifygo.String("lte"),
			Sort: "asc",
		},
		Limit:    githubcomchunkifydevchunkifygo.Int(1),
		Metadata: [][]string{{"J!Q0Ok0bzJb7:pro"}},
		Offset:   githubcomchunkifydevchunkifygo.Int(0),
		SourceID: githubcomchunkifydevchunkifygo.String("source_id"),
		Status:   githubcomchunkifydevchunkifygo.UploadListParamsStatusWaiting,
	})
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUploadDelete(t *testing.T) {
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
	err := client.Uploads.Delete(context.TODO(), "uploadId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
