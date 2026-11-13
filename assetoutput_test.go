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

func TestAssetOutputNew(t *testing.T) {
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
	_, err := client.Assets.Outputs.New(
		context.TODO(),
		"assetId",
		chunkify.AssetOutputNewParams{
			OfStaticRendition: &chunkify.AssetStaticRenditionOutputCreateParams{
				Params: chunkify.AssetStaticRenditionCreateParams{
					Resolution: chunkify.AssetStaticRenditionCreateParamsResolution1080p,
				},
			},
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

func TestAssetOutputGet(t *testing.T) {
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
	_, err := client.Assets.Outputs.Get(
		context.TODO(),
		"outputId",
		chunkify.AssetOutputGetParams{
			AssetID: "assetId",
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

func TestAssetOutputListWithOptionalParams(t *testing.T) {
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
	_, err := client.Assets.Outputs.List(
		context.TODO(),
		"assetId",
		chunkify.AssetOutputListParams{
			Limit:  chunkify.Int(1),
			Offset: chunkify.Int(0),
			Status: chunkify.AssetOutputListParamsStatusPreparing,
			Type:   chunkify.AssetOutputListParamsTypeCaption,
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

func TestAssetOutputDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Assets.Outputs.Delete(
		context.TODO(),
		"outputId",
		chunkify.AssetOutputDeleteParams{
			AssetID:     "assetId",
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
