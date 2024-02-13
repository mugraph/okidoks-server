package utils

import (
	"strings"
)

const (
	HTTP_SCHEME  = "http://"
	HTTPS_SCHEME = "https://"
)

// NormalizeURL takes a URL string and returns a normalized URL string based on the 'secure' and 'lower' paramters 
func NormalizeURL(url string, secure, lower bool) string {
	if url == "" {
		return ""
	}

	if url[len(url)-1:] == "/" {
		url = url[len(url)-1:]
	}

	if secure && strings.HasPrefix(url, HTTP_SCHEME) {
		url = strings.Replace(url, HTTP_SCHEME, HTTPS_SCHEME, 1)
	}

	if lower {
		return strings.ToLower(url)
	}

	return url
}

// Normalize Creative Commons URL takes a URL string and return a normalized URL string
func NormalizeCCURL(url string) string {
	secureURL := NormalizeURL(url, true, true)
	return NormalizedCCLicensesMap.GetVal(secureURL, false)
}
