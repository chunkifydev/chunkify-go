// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify_test

import (
	"context"
	"os"
	"testing"

	"github.com/chunkifydev/chunkify-go"
	"github.com/chunkifydev/chunkify-go/internal/testutil"
	"github.com/chunkifydev/chunkify-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	job, err := client.Jobs.New(context.TODO(), chunkify.JobNewParams{
		Format: chunkify.JobNewParamsFormatUnion{
			OfMP4H264: &chunkify.MP4H264Param{
				Width:  chunkify.Int(1920),
				Height: chunkify.Int(1080),
				Crf:    chunkify.Int(21),
			},
		},
		SourceID: "src_2G6MJiNz71bHQGNzGwKx5cJwPFS",
		Transcoder: chunkify.JobNewParamsTranscoder{
			Quantity: chunkify.Int(4),
			Type:     "8vCPU",
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", job.ID)
}
