// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomchunkifydevchunkifygo_test

import (
	"context"
	"os"
	"testing"

	"github.com/chunkifydev/chunkify-go"
	"github.com/chunkifydev/chunkify-go/internal/testutil"
	"github.com/chunkifydev/chunkify-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.Sources.List(context.TODO(), githubcomchunkifydevchunkifygo.SourceListParams{
		Limit: githubcomchunkifydevchunkifygo.Int(30),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, source := range page.Data {
		t.Logf("%+v\n", source.ID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, source := range page.Data {
			t.Logf("%+v\n", source.ID)
		}
	}
}
