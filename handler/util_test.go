package handler

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQuery(t *testing.T) {
	urlValues := url.Values{}
	urlValues["searchTerm"] = []string{"camera"}
	urlValues["lat"] = []string{"12.12312"}
	urlValues["lng"] = []string{"15.5726"}
	reqParams, err := parseQuery(urlValues)
	assert.NoError(t, err)
	assert.Equal(t, "camera", reqParams.SearchTerm())
	assert.Equal(t, 12.12312, reqParams.Lat())
	assert.Equal(t, 15.5726, reqParams.Lng())

}

func TestParseQueryWithInvalidLongitude(t *testing.T) {
	urlValues := url.Values{}
	urlValues["searchTerm"] = []string{"camera"}
	urlValues["lat"] = []string{"12.12312"}
	urlValues["lng"] = []string{"invalid"}
	_, err := parseQuery(urlValues)
	assert.Error(t, err)
}

func TestParseQueryWithInvalidLatitude(t *testing.T) {
	urlValues := url.Values{}
	urlValues["searchTerm"] = []string{"camera"}
	urlValues["lat"] = []string{"invalid"}
	urlValues["lng"] = []string{"12.1212"}
	_, err := parseQuery(urlValues)
	assert.Error(t, err)
}
