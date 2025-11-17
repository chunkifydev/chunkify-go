// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/chunkify-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Aws string        // Always "aws"
type Chunkify string   // Always "chunkify"
type Cloudflare string // Always "cloudflare"

func (c Aws) Default() Aws               { return "aws" }
func (c Chunkify) Default() Chunkify     { return "chunkify" }
func (c Cloudflare) Default() Cloudflare { return "cloudflare" }

func (c Aws) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Chunkify) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Cloudflare) MarshalJSON() ([]byte, error) { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
