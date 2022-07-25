package roundrobin_test

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	roundrobin "github.com/stsmdt/round-robin"
)

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		desc          string
		urls          []*url.URL
		expected      []*url.URL
		expectedError error
	}{
		{
			desc:          "No urls provided",
			urls:          []*url.URL{},
			expected:      []*url.URL{},
			expectedError: errors.New("no urls provided"),
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

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			roundRobin, err := roundrobin.New(test.urls...)

			gots := make([]*url.URL, 0, len(test.expected))
			for j := 0; j < len(test.expected); j++ {
				gots = append(gots, roundRobin.Next())
			}

			if test.expectedError != nil && assert.Error(t, err) {
				assert.Equal(t, test.expectedError, err)
			}

			assert.Equal(t, test.expected, gots)
		})
	}
}
