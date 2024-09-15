package trace

import (
	"context"
	"net/http"
	"sync"
)

const (
	TagDescription = "Description"
)

type Span interface {
	Context() context.Context
	AddTags(tags map[string]string)
	AddExtraData(extraData map[string]any)
	Finish()
}

type SpanFactory interface {
	StartSpan(ctx context.Context, name string) (context.Context, Span)
}

func StartSpanRequest(r *http.Request, name string) (*http.Request, Span) {
	ctx, span := StartSpan(r.Context(), name)
	return r.WithContext(ctx), span
}

func StartSpan(ctx context.Context, name string) (context.Context, Span) {
	return getFactory().StartSpan(ctx, name)
}

var (
	lock    sync.Mutex
	factory SpanFactory
)

func Register(spanFactory SpanFactory) {
	lock.Lock()
	defer lock.Unlock()
	factory = spanFactory
}

func getFactory() SpanFactory {
	lock.Lock()
	defer lock.Unlock()
	if factory == nil {
		return dummyFactory{}
	}
	return factory
}

type dummyFactory struct{}

func (dummyFactory) StartSpan(ctx context.Context, name string) (context.Context, Span) {
	return ctx, dummySpan{ctx}
}

type dummySpan struct {
	ctx context.Context
}

func (s dummySpan) Context() context.Context {
	return s.ctx
}

func (s dummySpan) AddTags(_ map[string]string) {
}

func (s dummySpan) AddExtraData(_ map[string]any) {
}

func (s dummySpan) Finish() {
}
