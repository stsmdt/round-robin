package roundrobin

import (
	"errors"
	"net/url"

	"go.uber.org/atomic"
)

// RoundRobin is an interface which represents the round-robin balancing.
type RoundRobin interface {
	Next() url.URL
}

type roundrobin struct {
	urls  []url.URL
	index atomic.Uint32
}

// ErrNoURLsProvided is the error that no URLs were provided
var ErrNoURLsProvided = errors.New("no urls provided")

// New returns a RoundRobin implementation
func New(urls []url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, ErrNoURLsProvided
	}

	rr := &roundrobin{}
	rr.urls = make([]url.URL, len(urls))
	copy(rr.urls, urls)

	return rr, nil
}

// Next returns the next url
func (r *roundrobin) Next() url.URL {
	var next uint32

	for {
		prev := r.index.Load()
		next = prev + 1

		if next > uint32(len(r.urls)) {
			next = 1
		}

		if r.index.CompareAndSwap(prev, next) {
			break
		}
	}

	return r.urls[next-1]
}
