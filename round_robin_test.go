package roundrobin_test

import (
	"fmt"
	"net/url"
	"reflect"
	"sync"
	"testing"

	roundrobin "github.com/stsmdt/round-robin"
)

func TestRoundRobin(t *testing.T) {
	testcases := []struct {
		desc        string
		urls        []*url.URL
		expected    []*url.URL
		expectedErr string
	}{
		{
			desc:        "No urls provided",
			urls:        []*url.URL{},
			expected:    []*url.URL{},
			expectedErr: "no urls provided",
		},
		{
			desc: "Everything provided",
			urls: []*url.URL{
				{Host: "127.0.0.1"},
				{Host: "127.0.0.2"},
				{Host: "127.0.0.3"},
				{Host: "127.0.0.4"},
			},
			expected: []*url.URL{
				{Host: "127.0.0.1"},
				{Host: "127.0.0.2"},
				{Host: "127.0.0.3"},
				{Host: "127.0.0.4"},
				{Host: "127.0.0.1"},
				{Host: "127.0.0.2"},
				{Host: "127.0.0.3"},
				{Host: "127.0.0.4"},
				{Host: "127.0.0.1"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			roundRobin, err := roundrobin.New(tc.urls...)

			gots := make([]*url.URL, 0, len(tc.expected))
			for j := 0; j < len(tc.expected); j++ {
				gots = append(gots, roundRobin.Next())
			}

			switch {
			case err == nil && len(tc.expectedErr) > 0:
				t.Errorf("Expected error %q but got nil", tc.expectedErr)
			case err != nil && err.Error() != tc.expectedErr:
				t.Errorf("Expected error %q but got %q", tc.expectedErr, err.Error())
			}

			if got, want := gots, tc.expected; !reflect.DeepEqual(got, want) {
				t.Errorf("Expected %v but got %v", want, got)
			}
		})
	}
}

func BenchmarkRoundRobinSync(b *testing.B) {
	urls := []*url.URL{
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
		{Host: "127.0.0.11"},
		{Host: "127.0.0.12"},
		{Host: "127.0.0.13"},
		{Host: "127.0.0.14"},
		{Host: "127.0.0.15"},
	}

	for i := 1; i <= len(urls); i++ {
		b.Run(fmt.Sprintf("RoundRobinSliceOf(%d)", i), func(b *testing.B) {
			roundrobin, err := roundrobin.New(urls[:i]...)
			if err != nil {
				b.Fatal(err)
			}

			waitGroup := &sync.WaitGroup{}
			for i := 0; i < b.N; i++ {
				waitGroup.Add(1)
				defer waitGroup.Done()
				roundrobin.Next()
			}
		})
	}
}

func BenchmarkRoundRobinAsync(b *testing.B) {
	urls := []*url.URL{
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
		{Host: "127.0.0.11"},
		{Host: "127.0.0.12"},
		{Host: "127.0.0.13"},
		{Host: "127.0.0.14"},
		{Host: "127.0.0.15"},
	}

	for i := 1; i <= len(urls); i++ {
		b.Run(fmt.Sprintf("RoundRobinSliceOf(%d)", i), func(b *testing.B) {
			roundrobin, err := roundrobin.New(urls[:i]...)
			if err != nil {
				b.Fatal(err)
			}

			waitGroup := &sync.WaitGroup{}
			for i := 0; i < b.N; i++ {
				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					roundrobin.Next()
				}()
			}
			waitGroup.Wait()
		})
	}
}
