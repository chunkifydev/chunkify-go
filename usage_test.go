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

func TestUsage(t *testing.T) {
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
	job, err := client.Jobs.New(context.TODO(), githubcomchunkifydevchunkifygo.JobNewParams{
		Format: githubcomchunkifydevchunkifygo.JobNewParamsFormatUnion{
			OfMP4H264: &githubcomchunkifydevchunkifygo.MP4H264Param{
				Width:  githubcomchunkifydevchunkifygo.Int(1920),
				Height: githubcomchunkifydevchunkifygo.Int(1080),
				Crf:    githubcomchunkifydevchunkifygo.Int(21),
			},
		},
		SourceID: "src_2G6MJiNz71bHQGNzGwKx5cJwPFS",
		Transcoder: githubcomchunkifydevchunkifygo.JobNewParamsTranscoder{
			Quantity: githubcomchunkifydevchunkifygo.Int(4),
			Type:     "8vCPU",
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", job.ID)
}
