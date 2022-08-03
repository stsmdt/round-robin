package roundrobin_test

import (
	"net/url"
	"reflect"
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
