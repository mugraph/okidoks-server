package utils

import (
	"fmt"
	"net/url"
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

// Normalize a DOI
func NormalizeDOI(doi string, sandbox bool) string {
	doiStr, _ := ValidateDOI(doi)
	if doiStr == "" {
		return ""
	}
	return DOIResolver(doiStr, sandbox) + strings.ToLower(doiStr)
}

// NormalizeID cekcs for valid DOI or HTTP(S) URL
func NormalizeID(pid string, sandbox bool) string {
	if pid == "" {
		return ""
	}

	// Check for valid DOI
	doi := NormalizeDOI(pid, sandbox)
	if doi != "" {
		return doi
	}

	// Check for valid URL
	URL, err := url.Parse(pid)
	if err != nil {
		// invalid
		log.Warn(fmt.Sprintf("could not parse %s as url", pid),
			"input", pid,
			"error", err)
	} else {
		if URL.Scheme == "" || strings.Contains(URL.Host, "http") || strings.Contains(URL.Host, "https") {
			return ""
		} else {
			pid = strings.Replace(pid, HTTP_SCHEME, HTTPS_SCHEME, 1)
		}
	}

	pid = strings.TrimSuffix(pid, "/")

	return pid
}

// NormalizeCCURL takes a URL string and returns a normalized URL string
func NormalizeCCURL(url string) string {
	secureURL := NormalizeURL(url, true, true)
	return NormalizedCCLicensesMap.GetVal(secureURL, false, false)
}

// NormalizeORCID
func NormalizeORCID(orcid string) string {
	if orcid == "" {
		return ""
	}
	orcid, err := ValidateORCID(orcid)
	if err != nil {
		log.Warn("could not validate ORCID", "error", err)
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
