package selector

import (
	"context"
	"testing"

	corelogger "github.com/go-gost/core/logger"
	ctxvalue "github.com/go-gost/x/ctx"
	xlogger "github.com/go-gost/x/logger"
)

var selectorBenchSink int

func benchCandidates(size int) []int {
	vs := make([]int, size)
	for i := range vs {
		vs[i] = i + 1
	}
	return vs
}

func BenchmarkRoundRobinStrategyApply(b *testing.B) {
	strategy := RoundRobinStrategy[int]()
	vs := benchCandidates(64)
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectorBenchSink = strategy.Apply(ctx, vs...)
	}
}

func BenchmarkRandomStrategyApply(b *testing.B) {
	strategy := RandomStrategy[int]()
	vs := benchCandidates(64)
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectorBenchSink = strategy.Apply(ctx, vs...)
	}
}

func BenchmarkHashStrategyApply(b *testing.B) {
	corelogger.SetDefault(xlogger.Nop())

	strategy := HashStrategy[int]()
	vs := benchCandidates(64)
	ctx := ctxvalue.ContextWithHash(context.Background(), &ctxvalue.Hash{Source: "bench-client"})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		selectorBenchSink = strategy.Apply(ctx, vs...)
	}
}
