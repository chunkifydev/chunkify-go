// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify_test

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stainless-sdks/chunkify-go"
	"github.com/stainless-sdks/chunkify-go/internal/testutil"
	"github.com/stainless-sdks/chunkify-go/option"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestWebhookNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.New(context.TODO(), chunkify.WebhookNewParams{
		URL:     "https://example.com/webhook",
		Enabled: chunkify.Bool(true),
		Events:  []string{"job.completed"},
	})
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookGet(t *testing.T) {
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
	_, err := client.Webhooks.Get(context.TODO(), "webhookId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Webhooks.Update(
		context.TODO(),
		"webhookId",
		chunkify.WebhookUpdateParams{
			Enabled: chunkify.Bool(true),
			Events:  []string{"job.completed"},
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

func TestWebhookList(t *testing.T) {
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
	_, err := client.Webhooks.List(context.TODO())
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookDelete(t *testing.T) {
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
	err := client.Webhooks.Delete(context.TODO(), "webhookId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookUnwrap(t *testing.T) {
	client := chunkify.NewClient(
		option.WithWebhookKey("whsec_c2VjcmV0Cg=="),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	payload := []byte(`{"id":"notf_2G6MJiNz71bHQGNzGwKx5cJwPFS","data":{"files":[{"id":"file_2G6MJiNz71bHQGNzGwKx5cJwPFS","audio_bitrate":128000,"audio_codec":"aac","created_at":"2025-01-01T12:00:00Z","duration":120,"height":1080,"job_id":"job_2G6MJiNz71bHQGNzGwKx5cJwPFS","mime_type":"video/mp4","path":"path/to/file.mp4","size":1234567,"storage_id":"stor_chunkify_2wLmj1fp8neUaFAWwwxvzKAT0Fa","url":"https://my-bucket.s3.us-east-1.amazonaws.com/path/to/file.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256","video_bitrate":20000000,"video_codec":"h264","video_framerate":29.97,"width":1920}],"job":{"id":"job_2G6MJiNz71bHQGNzGwKx5cJwPFS","billable_time":120,"created_at":"2025-01-01T12:00:00Z","format":{"id":"mp4_h264","audio_bitrate":32000,"bufsize":100000,"channels":1,"crf":35,"disable_audio":true,"disable_video":true,"duration":1,"framerate":15,"gop":1,"height":-2,"level":41,"maxrate":100000,"minrate":100000,"movflags":"movflags","pixfmt":"yuv410p","preset":"10","profilev":"main10","seek":1,"video_bitrate":100000,"width":-2},"progress":45.5,"source_id":"src_2G6MJiNz71bHQGNzGwKx5cJwPFS","status":"transcoding","storage":{"id":"stor_aws_S1cce6120E56e7Tu9ioP09Nhjk1","path":"path/to/video.mp4"},"transcoder":{"auto":true,"quantity":10,"type":"4vCPU"},"updated_at":"2025-01-01T12:05:00Z","error":{"detail":"detail","message":"message","type":"setup"},"hls_manifest_id":"hls_2v6EIgcNAycdS5g0IUm0TXBjvHV","metadata":{"key1":"value1","key2":"value2"},"started_at":"2025-01-01T12:01:00Z"}},"date":"2025-01-01T12:00:00Z","event":"job.completed"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Error("Failed to sign test webhook message")
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Error("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Error("Failed to unwrap webhook:", err)
	}
}
