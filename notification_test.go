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

func TestNotificationNew(t *testing.T) {
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
	_, err := client.Notifications.New(context.TODO(), githubcomchunkifydevchunkifygo.NotificationNewParams{
		Event:     githubcomchunkifydevchunkifygo.NotificationNewParamsEventJobCompleted,
		ObjectID:  "job_A1cce6120E56e7Tu9ioP09Nhjk9",
		WebhookID: "wh_A1cce6120E56e7Tu9ioP09Nhjk9",
	})
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationGet(t *testing.T) {
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
	_, err := client.Notifications.Get(context.TODO(), "notificationId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.List(context.TODO(), githubcomchunkifydevchunkifygo.NotificationListParams{
		Created: githubcomchunkifydevchunkifygo.NotificationListParamsCreated{
			Gte:  githubcomchunkifydevchunkifygo.String("2102-57-32"),
			Lte:  githubcomchunkifydevchunkifygo.String("2025-01-01"),
			Sort: "asc",
		},
		Events:   []string{"job.completed"},
		Limit:    githubcomchunkifydevchunkifygo.Int(1),
		ObjectID: githubcomchunkifydevchunkifygo.String("object_id"),
		Offset:   githubcomchunkifydevchunkifygo.Int(0),
		ResponseStatusCode: githubcomchunkifydevchunkifygo.NotificationListParamsResponseStatusCode{
			Eq:  githubcomchunkifydevchunkifygo.Int(100),
			Gte: githubcomchunkifydevchunkifygo.Int(100),
			Lte: githubcomchunkifydevchunkifygo.Int(100),
		},
		WebhookID: githubcomchunkifydevchunkifygo.String("webhook_id"),
	})
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationDelete(t *testing.T) {
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
	err := client.Notifications.Delete(context.TODO(), "notificationId")
	if err != nil {
		var apierr *githubcomchunkifydevchunkifygo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
