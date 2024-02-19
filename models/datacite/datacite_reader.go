package datacite

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/mugraph/fullname_parser"
	"github.com/mugraph/okidoks-server/models/commonmeta"
	"github.com/mugraph/okidoks-server/utils"
)

// Takes a datacite DOI string, validates it and fetches it's metadata via the datacite API.
// Returns a datacite.Resource and an error.
func GetDataCite(doi string, test bool) (Resource, error) {
	doi, err := utils.ValidateDOI(doi)
	if err != nil {
		return Resource{}, fmt.Errorf("failed to validate DOI: %v", err)
	}

	var apiURL string
	if test {
		// Construct the API URL for fetching DOI attrdata
		apiURL = "https://api.test.datacite.org/dois/" + doi +
			"?publisher=true&affiliation=true"
	} else {
		// Construct the API URL for fetching DOI attrdata
		apiURL = "https://api.datacite.org/dois/" + doi +
			"?publisher=true&affiliation=true"

	}

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

func FormatNameIdentifier(ni NameIdentifier) string {
	if ni.NameIdentifierScheme == "ORCID" {
		return utils.NormalizeORCID(ni.NameIdentifier)
	}
	if ni.NameIdentifierScheme == "ROR" {
		return utils.NormalizeROR(ni.NameIdentifier)
	}
	if ni.NameIdentifierScheme == "ISNI" {
		return utils.NormalizeISNI(ni.NameIdentifier)
	}
	if utils.ValidateURL(ni.NameIdentifier) == "URL" {
		return ni.NameIdentifier
	}
	if ni.SchemeURI != nil && *ni.SchemeURI != "" {
		return *ni.SchemeURI + ni.NameIdentifier
	}
	return ""
}

func containsWord(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

func isPersonalName(name string) bool {
	// Personal Names are not allowed to contain semicolons
	if strings.Contains(name, ";") {
		return false
	}

	// Check if a name has no spaces, e.g. FamousOrganisation, not including commas
	if len(strings.Split(name, " ")) == 1 && !strings.Contains(name, ",") {
		return false
	}

	// Check if name contains words known to be used in organization names
	words := strings.Fields(name)
	notWanted := []string{
		"University",
		"College",
		"Institute",
		"School",
		"Center",
		"Department",
		"Laboratory",
		"Library",
		"Museum",
		"Foundation",
		"Society",
		"Association",
		"Company",
		"Corporation",
		"Collaboration",
		"Consortium",
		"Incorporated",
		"Inc.",
		"Institut",
		"Research",
		"Science",
		"Team",
	}
	for _, w := range words {
		if containsWord(notWanted, w) {
			return false
		}
	}

	// Check for suffixes, e.g. "John Smith, MD"
	splits := strings.Split(name, ", ")
	last := splits[len(splits)-1]
	notWanted = []string{"MD", "PhD", "BS"}
	if containsWord(notWanted, last) {
		return true
	}

	// Check if name can be parsed into given/family name
	names := fullname_parser.ParseFullname(name)
	if names.First != "" || names.Last != "" {
		return true
	}

	return true
}

func cleanupName(name string) string {
	if name == "" {
		return ""
	}

	// Detect pattern "Smith J.", but not "Smith, John K."
	if !strings.Contains(name, ",") {
		re := regexp.MustCompile(`/(?:[A-Z]\.)?(?:-?[A-Z]\.)/`)
		name = re.ReplaceAllString(name, ", ${1}${2}")
	}

	// Remove spaces around hyphens
	name = strings.ReplaceAll(name, " - ", "-")

	// Remove non-standard space characters
	re := regexp.MustCompile(`/[ \t\r\n\v\f]/`)
	name = re.ReplaceAllString(name, " ")

	return name
}

// Parse onc ResourceContributor to commonmeta.Contributor
func contributor(c ResourceContributor) (con *commonmeta.Contributor) {
	var name string
	if c.Name != "" {
		name = cleanupName(c.Name)
	}

	var givenName string
	if c.GivenName != nil {
		givenName = string(*c.GivenName)
	}

	var familyName string
	if c.FamilyName != nil {
		familyName = string(*c.FamilyName)
	}

	var id string
	for _, ni := range c.NameIdentifiers {
		id = FormatNameIdentifier(ni)
	}

	var ctype string
	if c.NameType != nil {
		ctype = string(*c.NameType)
		if ctype[len(ctype)-2:] == "al" {
			ctype = ctype[:len(ctype)-2]
		}
	}

	_, err := utils.ValidateROR(id)
	if err == nil && id != "" && ctype == "" {
		ctype = "Organization"
	}

	_, err = utils.ValidateORCID(id)
	if err == nil && id != "" && ctype == "" {
		ctype = "Person"
	}

	if ctype == "" && (c.GivenName != nil || c.FamilyName != nil) {
		ctype = "Person"
	}

	if ctype == "" && c.Name != "" && isPersonalName(c.Name) {
		ctype = "Person"
	}

	if ctype == "" && c.Name != "" {
		ctype = "Organization"
	}

	if ctype == "Person" && c.Name != "" && c.GivenName == nil && c.FamilyName == nil {
		// names = HumanName(name)
		names := fullname_parser.ParseFullname(c.Name)
		if names.First != "" {
			givenName = names.First
		}
		if names.First != "" && names.Middle != "" {
			givenName = names.First + " " + names.Middle
		}
		if names.Last != "" {
			familyName = names.Last
		}
	}

	// Parse contributor roles, checking for roles supported by commonmeta
	var roles []commonmeta.ContributorRole
	roles = append(
		roles,
		commonmeta.ContributorRole{
			Role: commonmeta.Role(
				utils.ContributorRoleMap.GetVal(string(c.ContributorType), true, true),
			),
		},
	)

	con = &commonmeta.Contributor{
		ContributorRoles: roles,
	}

	if id != "" {
		con.ID = &id
	}

	// Final check for valid ContributorTypes then assign to Output
	if ctype != "" {
		con.Type = commonmeta.ContributorType(utils.ContributorTypeMap.GetVal(ctype, false, false))
	}

	// If ContributorType is Person, keep Given/FamilyName
	if ctype == "Person" {
		con.GivenName = &givenName
		con.FamilyName = &familyName
	}

	// If ContributorType is Organization, keep Name
	if ctype == "Organization" {
		con.Name = &name
	}

	return con
}

// Takes a slice of datacite.ResourceContributors struct.
// Returns s slice of commonmeta.Contributor pointers.
func contributors(c []ResourceContributor) (con []*commonmeta.Contributor) {
	for _, v := range c {
		contributor := contributor(v)
		con = append(con, contributor)
	}
	return con
}

// Takes a datacite.Types struct by value.
// Returns two commonmeta.ResourceType values.
func types(t Types) (gt, at commonmeta.ResourceType) {
	rtg := string(t.ResourceTypeGeneral)
	rt := t.ResourceType

	gt = commonmeta.ResourceType(utils.DataciteToCommonMeta.GetVal(rtg, true, false))
	at = commonmeta.ResourceType(utils.DataciteToCommonMeta.GetVal(rt, false, false))

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
			tts := utils.TitleTypeMap.GetVal(string(*v.TitleType), false, false)
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
			ID:  rl[0].RightsIdentifier,
		}
	}
	if len(rl) >= 2 && rl[1].RightsURI != nil && rl[1].RightsIdentifier != nil {
		URL := utils.NormalizeCCURL(*rl[1].RightsURI)
		lic = &commonmeta.License{
			URL: URL,
			ID:  rl[1].RightsIdentifier,
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
