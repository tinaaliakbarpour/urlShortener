package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestShortLinkGenerator(t *testing.T) {

	testcases := []struct {
		desc string
		initialLink string
		shortLink string
	}{
		{
			desc: "a",
			initialLink: "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
			shortLink:   "EJSR38HU",
		},
		{
			desc: "b",
			initialLink: "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/",
			shortLink:   "Egp9XL2r",
		},
	}

	for _, tc := range testcases{
		t.Run(tc.desc, func(t *testing.T) {
			result := Shortener.GenerateShortLink(tc.initialLink, "1")
			assert.Equal(t, tc.shortLink,result)
		})


	}

}
