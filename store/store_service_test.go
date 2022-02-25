package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionAndRetrieval(t *testing.T) {

	testcase := []struct {
		desc        string
		initialLink string
		shortURL    string
		userId      string
	}{
		{
			desc:        "a",
			initialLink: "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
			shortURL:    "e0dba740-fc4b-4977-872c-d360239e6b1a",
			userId:      "Jsz4k57oAX",
		},
	}

	for _, tc := range testcase{
		t.Run(tc.desc, func(t *testing.T) {
			Store.InitializeStore()
			Store.SaveUrlMapping(tc.shortURL, tc.initialLink, tc.userId)
		})
		assert.Equal(t, tc.initialLink, Store.RetrieveInitialUrl(tc.shortURL))
	}



}
