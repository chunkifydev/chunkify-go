// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/stainless-sdks/chunkify-go/internal/apijson"
	"github.com/stainless-sdks/chunkify-go/internal/requestconfig"
	"github.com/stainless-sdks/chunkify-go/option"
	"github.com/stainless-sdks/chunkify-go/packages/param"
	"github.com/stainless-sdks/chunkify-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PaginatedResults[T any] struct {
	Data   []T   `json:"data"`
	Total  int64 `json:"total"`
	Offset int64 `json:"offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Total       respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedResults[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedResults[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedResults[T]) GetNextPage() (res *PaginatedResults[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Offset

	if next < r.Total && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedResults[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedResults[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedResultsAutoPager[T any] struct {
	page *PaginatedResults[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedResultsAutoPager[T any](page *PaginatedResults[T], err error) *PaginatedResultsAutoPager[T] {
	return &PaginatedResultsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedResultsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedResultsAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedResultsAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedResultsAutoPager[T]) Index() int {
	return r.run
}
