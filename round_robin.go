package roundrobin

import (
	"errors"
	"net/url"
	"sync/atomic"
)

// RoundRobin is an interface which represents the round-robin balancing.
type RoundRobin interface {
	Next() *url.URL
}

type roundrobin struct {
	urls  []*url.URL
	index uint32
}

// New returns a RoundRobin implementation
func New(urls ...*url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, errors.New("no urls provided")
	}

	return &roundrobin{
		urls: urls,
	}, nil
}

// Next returns the next url
func (r *roundrobin) Next() *url.URL {
	var next uint32

	for {
		prev := atomic.LoadUint32(&r.index)
		next = prev + 1

		if next > uint32(len(r.urls)) {
			next = 1
		}

		if atomic.CompareAndSwapUint32(&r.index, prev, next) {
			break
		}
	}

	return r.urls[next-1]
}
