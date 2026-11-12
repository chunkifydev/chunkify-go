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

func TestMetricAssetsWithOptionalParams(t *testing.T) {
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
	_, err := client.Metrics.Assets(context.TODO(), chunkify.MetricAssetsParams{
		Format: chunkify.MetricAssetsParamsFormatJson,
		From:   chunkify.String("7321-69-10"),
		To:     chunkify.String("7321-69-10"),
	})
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMetricStorageWithOptionalParams(t *testing.T) {
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
	_, err := client.Metrics.Storage(context.TODO(), chunkify.MetricStorageParams{
		Format:    chunkify.MetricStorageParamsFormatJson,
		From:      chunkify.String("7321-69-10"),
		StorageID: chunkify.String("x"),
		To:        chunkify.String("7321-69-10"),
	})
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMetricUsageWithOptionalParams(t *testing.T) {
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
	_, err := client.Metrics.Usage(context.TODO(), chunkify.MetricUsageParams{
		Format: chunkify.MetricUsageParamsFormatJson,
		From:   chunkify.String("7321-69-10"),
		To:     chunkify.String("7321-69-10"),
	})
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
