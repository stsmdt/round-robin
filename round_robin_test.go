package roundrobin_test

import (
	"fmt"
	"net/url"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	roundrobin "github.com/stsmdt/round-robin"
)

func TestNew_NoURLsProvided(t *testing.T) {
	a := assert.New(t)

	rr, err := roundrobin.New([]url.URL{})
	a.Nil(rr)
	a.Error(err)
	a.ErrorIs(err, roundrobin.ErrNoURLsProvided)
}

func TestNew_URLsProvided(t *testing.T) {
	a := assert.New(t)

	rr, err := roundrobin.New([]url.URL{{Host: "127.0.0.1"}})
	a.NotNil(rr)
	a.NoError(err)
}

func TestNext_NoMutation(t *testing.T) {
	a := assert.New(t)

	givenURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
	}
	wantURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
	}

	rr, err := roundrobin.New(givenURLs)

	gotURLs := make([]url.URL, 0, len(wantURLs))
	for i := 0; i < len(wantURLs); i++ {
		gotURLs = append(gotURLs, rr.Next())
	}

	a.NoError(err)
	a.Equal(wantURLs, gotURLs)
}

func TestNext_GivenURLsMutation(t *testing.T) {
	a := assert.New(t)

	givenURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
	}
	wantURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
	}

	rr, err := roundrobin.New(givenURLs)

	givenURLs[0] = url.URL{Host: "127.0.0.6"}

	gotURLs := make([]url.URL, 0, len(wantURLs))
	for i := 0; i < len(wantURLs); i++ {
		gotURLs = append(gotURLs, rr.Next())
	}

	a.NoError(err)
	a.Equal(wantURLs, gotURLs)
}

func TestNext_GotURLMutation(t *testing.T) {
	a := assert.New(t)

	givenURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
	}
	wantURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
	}

	rr, err := roundrobin.New(givenURLs)

	gotURLs := make([]url.URL, 0, len(wantURLs))
	for i := 0; i < len(wantURLs); i++ {
		gotURL := rr.Next()
		gotURLs = append(gotURLs, gotURL)
		gotURL.Host = "127.0.0.6"
	}

	a.NoError(err)
	a.Equal(wantURLs, gotURLs)
}

func BenchmarkNext_Sync(b *testing.B) {
	givenURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
		{Host: "127.0.0.6"},
		{Host: "127.0.0.7"},
		{Host: "127.0.0.8"},
		{Host: "127.0.0.9"},
		{Host: "127.0.0.10"},
	}

	for i := 1; i <= len(givenURLs); i++ {
		b.Run(fmt.Sprintf("slice_len_%d", i), func(b *testing.B) {
			rr, err := roundrobin.New(givenURLs[:i])
			if err != nil {
				b.Fatal(err)
			}

			wg := &sync.WaitGroup{}
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				defer wg.Done()
				rr.Next()
			}
		})
	}
}

func BenchmarkNext_Async(b *testing.B) {
	givenURLs := []url.URL{
		{Host: "127.0.0.1"},
		{Host: "127.0.0.2"},
		{Host: "127.0.0.3"},
		{Host: "127.0.0.4"},
		{Host: "127.0.0.5"},
		{Host: "127.0.0.6"},
		{Host: "127.0.0.7"},
		{Host: "127.0.0.8"},
		{Host: "127.0.0.9"},
		{Host: "127.0.0.10"},
	}

	for i := 1; i <= len(givenURLs); i++ {
		b.Run(fmt.Sprintf("slice_len_%d", i), func(b *testing.B) {
			rr, err := roundrobin.New(givenURLs[:i])
			if err != nil {
				b.Fatal(err)
			}

			wg := &sync.WaitGroup{}
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					rr.Next()
				}()
			}
			wg.Wait()
		})
	}
}
