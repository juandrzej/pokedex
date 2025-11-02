package pokeapi

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/juandrzej/pokedex/internal/pokecache"
)

type mockRT struct {
	count int
	body  []byte
	code  int
}

func (m *mockRT) RoundTrip(req *http.Request) (*http.Response, error) {
	m.count++
	return &http.Response{
		StatusCode: m.code,
		Body:       io.NopCloser(bytes.NewReader(m.body)),
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func newTestClient(rt http.RoundTripper) *Client {
	return &Client{
		httpClient: &http.Client{Transport: rt, Timeout: 2 * time.Second},
		cache:      pokecache.NewCache(50 * time.Millisecond),
	}
}

func TestFetchLocationAreas_DecodeOK(t *testing.T) {
	b := []byte(`{
        "count": 1,
        "next": "",
        "previous": "",
        "results": [{"name":"area-1","url":"u1"}]
    }`)
	rt := &mockRT{body: b, code: 200}
	c := newTestClient(rt)

	got, err := c.FetchLocationAreas(baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Count != 1 || len(got.Results) != 1 || got.Results[0].Name != "area-1" {
		t.Fatalf("bad decode: %+v", got)
	}
}

func TestFetchLocationAreas_UsesCacheOnSecondCall(t *testing.T) {
	b := []byte(`{"count":1,"next":"","previous":"","results":[{"name":"area-1","url":"u1"}]}`)
	rt := &mockRT{body: b, code: 200}
	c := newTestClient(rt)

	url := baseURL

	if _, err := c.FetchLocationAreas(url); err != nil {
		t.Fatalf("first call error: %v", err)
	}
	if rt.count != 1 {
		t.Fatalf("expected 1 network call, got %d", rt.count)
	}

	if _, err := c.FetchLocationAreas(url); err != nil {
		t.Fatalf("second call error: %v", err)
	}
	if rt.count != 1 {
		t.Fatalf("expected cache hit (no extra network), got %d calls", rt.count)
	}
}
