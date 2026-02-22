package routing

import (
	"net"
	"net/http"
	"net/url"
	"testing"

	corerouting "github.com/go-gost/core/routing"
)

var routingBenchSink bool

func BenchmarkMatcherBuild(b *testing.B) {
	rule := "Host(`api.example.com`) && Method(`GET`) && PathPrefix(`/api`) && Header(`X-Tenant`,`prod`) && Query(`debug`,`0`)"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		matcher, err := NewMatcher(rule)
		if err != nil {
			b.Fatalf("new matcher: %v", err)
		}
		routingBenchSink = matcher != nil
	}
}

func BenchmarkMatcherMatch(b *testing.B) {
	rule := "Host(`api.example.com`) && Method(`GET`) && PathPrefix(`/api`) && Header(`X-Tenant`,`prod`) && Query(`debug`,`0`)"
	matcher, err := NewMatcher(rule)
	if err != nil {
		b.Fatalf("new matcher: %v", err)
	}

	req := &corerouting.Request{
		ClientIP: net.ParseIP("10.0.0.8"),
		Host:     "api.example.com:443",
		Method:   http.MethodGet,
		Path:     "/api/v1/forwards",
		Query:    url.Values{"debug": {"0"}},
		Header:   http.Header{"X-Tenant": {"prod"}},
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		routingBenchSink = matcher.Match(req)
	}

	if !routingBenchSink {
		b.Fatal("unexpected benchmark setup: matcher should match request")
	}
}

func BenchmarkMatcherMatchRegexp(b *testing.B) {
	rule := "HostRegexp(`^api\\.example\\.com$`) && Method(`POST`) && PathRegexp(`^/api/v1/tunnel/[0-9]+$`)"
	matcher, err := NewMatcher(rule)
	if err != nil {
		b.Fatalf("new matcher: %v", err)
	}

	req := &corerouting.Request{
		ClientIP: net.ParseIP("10.0.0.8"),
		Host:     "api.example.com",
		Method:   http.MethodPost,
		Path:     "/api/v1/tunnel/42",
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		routingBenchSink = matcher.Match(req)
	}

	if !routingBenchSink {
		b.Fatal("unexpected benchmark setup: matcher should match request")
	}
}
