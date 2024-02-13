package datacite

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mugraph/okidoks-server/models/commonmeta"
	"github.com/mugraph/okidoks-server/utils"
)

// Takes a datacite DOI string, validates it and fetches it's metadata via the datacite API.
// Returns a datacite.Resource and an error.
func GetDataCite(doi string) (Resource, error) {
	doi, err := utils.ValidateDOI(doi)
	if err != nil {
		return Resource{}, fmt.Errorf("failed to validate DOI: %v", err)
	}

	// Construct the API URL for fetching DOI attrdata
	apiURL := "https://api.datacite.org/dois/" + doi +
		"?publisher=true&affiliation=true"

	// Make the HTTP GET request with NewRequest and DefaultClient.Do
	// for use with custom header
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("accept", "application/vnd.api+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Resource{}, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Return on non-200 status codes
	// DataCite API complies with JSON:API v1.1 where a server MUST respond with
	// 200 OK for a successful request:
	// https://jsonapi.org/format/#fetching-resources-responses
	if resp.StatusCode != http.StatusOK {
		return Resource{}, fmt.Errorf("HTTP response status: %v %v",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Resource{}, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal JSON response
	r, err := UnmarshalResource(body)
	if err != nil {
		return r, fmt.Errorf("failed to unmarshal response.Data: %v", err)
	}

	return r, nil
}

// Takes a slice of datacite.ResourceContributors struct.
// Returns s slice of commonmeta.Contributor pointers.
func contributors(c []ResourceContributor) (con []*commonmeta.Contributor) {
	for _, v := range c {
		var roles []commonmeta.ContributorRole
		roles = append(roles, commonmeta.ContributorRole{Role: commonmeta.Role(v.ContributorType)})
		// if v.ContributorType is Person
		con = append(con, &commonmeta.Contributor{
			GivenName:        v.GivenName,
			FamilyName:       v.FamilyName,
			ContributorRoles: roles,
		})
	}
	return con
}

// Takes a datacite.Types struct by value.
// Returns two commonmeta.ResourceType values.
func types(t Types) (gt, at commonmeta.ResourceType) {
	rtg := string(t.ResourceTypeGeneral)
	rt := t.ResourceType

	gt = commonmeta.ResourceType(utils.DataciteToCommonMeta.GetVal(rtg, true))
	at = commonmeta.ResourceType(utils.DataciteToCommonMeta.GetVal(rt, false))

	// If ResourceType is one of the new ResourceTypeGeneral types
	// introduced in schema 4.3, use it.
	if at != "" {
		gt = at
		at = ""
	} else {
		at = commonmeta.ResourceType(rt)
	}
	return gt, at
}

// Takes a slice of datacite.ResourceTitle structs.
// Returns a slice of commonmeta.Title pointers.
func titles(t []ResourceTitle) (titles []*commonmeta.Title) {
	for _, v := range t {
		var tt commonmeta.TitleType
		var lang string
		if v.TitleType != nil {
			tts := utils.TitleTypeMap.GetVal(string(*v.TitleType), false)
			tt = commonmeta.TitleType(tts)
		}
		if v.Lang != nil {
			lang = *v.Lang
		}

		titles = append(titles, &commonmeta.Title{
			Title:    v.Title,
			Type:     &tt,
			Language: lang,
		})
	}
	return titles
}

// Takes a slice of datacite.RightsList structs.
// Returns a commonmeta.License pointer.
func license(rl []RightsList) *commonmeta.License {
	var lic *commonmeta.License
	if len(rl) != 0 && rl[0].RightsURI != nil && rl[0].RightsIdentifier != nil {
		URL := utils.NormalizeCCURL(*rl[0].RightsURI)
		lic = &commonmeta.License{
			URL: URL,
			ID: rl[0].RightsIdentifier,
		}
	}
	if len(rl) >= 2 && rl[1].RightsURI != nil && rl[1].RightsIdentifier != nil {
		URL := utils.NormalizeCCURL(*rl[1].RightsURI)
		lic = &commonmeta.License{
			URL: URL,
			ID: rl[1].RightsIdentifier,
		}
	}
	return lic
}

// Takes a datacite.Publisher struct by value.
// Returns a commonmeta.Publisher pointer.
func publisher(p Publisher) *commonmeta.Publisher {
	pub := &commonmeta.Publisher{Name: p.Name}
	return pub
}

// Takes a datacite.Resource struct by value.
// Returns the corresponding commonmeta.Resource. 
func ReadDataCite(r Resource) (rs commonmeta.Resource, err error) {
	id, err := utils.DOIAsURL(r.Doi)
	if err != nil {
		return commonmeta.Resource{}, fmt.Errorf("could not get DOI as URL: %v", err)
	}

	gt, at := types(r.Types)

	rs = commonmeta.Resource{
		ID:           id,
		Type:         &gt,
		URL:          utils.NormalizeURL(*r.URL, true, true),
		Contributors: contributors(r.Contributors),
		Titles:       titles(r.Titles),
		Publisher:    publisher(r.Publisher),
		// Date
		// Recommended and optional properties
		AdditionalType: &at,
		License:        license(r.RightsList),
	}

	return rs, nil
}
