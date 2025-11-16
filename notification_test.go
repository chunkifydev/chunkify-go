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

func TestNotificationNew(t *testing.T) {
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
	_, err := client.Notifications.New(context.TODO(), chunkify.NotificationNewParams{
		Event:     chunkify.NotificationNewParamsEventJobCompleted,
		ObjectID:  "job_A1cce6120E56e7Tu9ioP09Nhjk9",
		WebhookID: "wh_A1cce6120E56e7Tu9ioP09Nhjk9",
	})
	if err != nil {
		var apierr *chunkify.Error
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
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Notifications.Get(context.TODO(), "notificationId")
	if err != nil {
		var apierr *chunkify.Error
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
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	_, err := client.Notifications.List(context.TODO(), chunkify.NotificationListParams{
		Created: chunkify.NotificationListParamsCreated{
			Gte:  chunkify.String("gte"),
			Lte:  chunkify.String("lte"),
			Sort: chunkify.String("sort"),
		},
		Events:   []string{"string"},
		Limit:    chunkify.Int(0),
		ObjectID: chunkify.String("object_id"),
		Offset:   chunkify.Int(0),
		ResponseStatusCode: chunkify.NotificationListParamsResponseStatusCode{
			Eq:  chunkify.Int(0),
			Gte: chunkify.Int(0),
			Lte: chunkify.Int(0),
		},
		WebhookID: chunkify.String("webhook_id"),
	})
	if err != nil {
		var apierr *chunkify.Error
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
	client := chunkify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithProjectAccessToken("My Project Access Token"),
	)
	err := client.Notifications.Delete(context.TODO(), "notificationId")
	if err != nil {
		var apierr *chunkify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
