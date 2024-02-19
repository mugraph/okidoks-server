package utils

import (
	"fmt"
	"strings"

	"github.com/mugraph/okidoks-server/logger"
)

var log = logger.Log

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
		url = url[:len(url)-1]
	}

	if secure && strings.HasPrefix(url, HTTP_SCHEME) {
		url = strings.Replace(url, HTTP_SCHEME, HTTPS_SCHEME, 1)
	}

	if lower {
		return strings.ToLower(url)
	}

	return url
}

// NormalizeCCURL takes a URL string and returns a normalized URL string
func NormalizeCCURL(url string) string {
	secureURL := NormalizeURL(url, true, true)
	fmt.Printf("%v\n", secureURL)
	return NormalizedCCLicensesMap.GetVal(secureURL, false, false)
}

// NormalizeORCID
func NormalizeORCID(orcid string) string {
	if orcid == "" {
		return ""
	}
	orcid, err := ValidateORCID(orcid)
	if err != nil {
		log.Warn("Could not validate ORCID", "error", err)
		return ""
	}
	return "https://orcid.org/" + orcid
}

func NormalizeROR(ror string) string {
	if ror == "" {
		return ""
	}
	ror, err := ValidateROR(ror)
	if err != nil {
		log.Warn("Could not validate ROR", "error", err)
		return ""
	}
	return "https://ror.org/" + ror
}

func NormalizeISNI(isni string) string {
	if isni == "" {
		return ""
	}
	ror, err := ValidateISNI(isni)
	if err != nil {
		log.Warn("Could not validate ISNI", "error", err)
		return ""
	}
	return "https://isni.org/isni/" + ror
}
