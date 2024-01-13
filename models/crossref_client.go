package models

type CrossRefData struct {
	Resource Resource `json:"message"`
}

// func GetDOIAgency(doi string) (agencyID string, err error) {
// 	// Construct the API URL for fetching DOI agency information
// 	apiURL := "https://api.crossref.org/works/" + doi + "/agency"
//
// 	// Make the HTTP GET request
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		return "", fmt.Errorf("Failed to make HTTP request: %v", err)
// 	}
// 	defer response.Body.Close()
//
// 	// Read the response body
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("Failed to read response body: %v", err)
// 	}
//
// 	// Unmarshal JSON response
// 	var data map[string]interface{}
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return "", fmt.Errorf("Failed to unmarshal JSON: %v", err)
// 	}
//
// 	// Extract agency ID from the JSON response
// 	if message, ok := data["message"].(map[string]interface{}); ok {
// 		if agency, ok := message["agency"].(map[string]interface{}); ok {
// 			if agencyID, ok := agency["id"].(string); ok {
// 				return agencyID, nil
// 			}
// 		}
// 	}
//
// 	return "", fmt.Errorf("Agency information not found in the response.")
// }
//
// func GetDOIMetadataFromCrossref(doi string) (resource Resource, err error) {
// 	// Construct the API URL for fetching DOI metadata
// 	apiURL := "https://api.crossref.org/works/" + doi
//
// 	// Make the HTTP GET request
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		return Resource{}, fmt.Errorf("Failed to make HTTP request: %v", err)
// 	}
// 	defer response.Body.Close()
//
// 	// Read the response body
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return Resource{}, fmt.Errorf("Failed to read response body: %v", err)
// 	}
//
// 	// // Prettify the response data
// 	// prettyBody, err := utils.PrettyJSON(body)
// 	// if err != nil {
// 	// 	return Resource{}, fmt.Errorf("Failed to prettify response data: %v", err)
// 	// }
//
// 	// // Print the prettified response
// 	// fmt.Println(prettyBody + "\n")
//
// 	// Unmarshal JSON response
// 	var data CrossRefData
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return Resource{}, fmt.Errorf("Failed to unmarshal JSON: %v", err)
// 	}
//
// 	resource = data.Resource
// 	resource.DOIAgency = "crossref"
//
// 	return resource, nil
// }
