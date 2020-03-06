package test_test

import (
	"testing"

	"github.com/minio/httpcache"
	"github.com/minio/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
