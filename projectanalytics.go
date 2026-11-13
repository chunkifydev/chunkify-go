// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chunkify

import (
	"github.com/chunkifydev/chunkify-go/option"
)

// ProjectAnalyticsService contains methods and other services that help with
// interacting with the chunkify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectAnalyticsService] method instead.
type ProjectAnalyticsService struct {
	Options     []option.RequestOption
	Connections ProjectAnalyticsConnectionService
}

// NewProjectAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectAnalyticsService(opts ...option.RequestOption) (r ProjectAnalyticsService) {
	r = ProjectAnalyticsService{}
	r.Options = opts
	r.Connections = NewProjectAnalyticsConnectionService(opts...)
	return
}
