package models

// The publisher of many resources.
type Publisher struct {
	ID   string `json:"id,omitempty"` // The identifier for the publisher.
	Name string `json:"name"`         // The name of the publisher.
}
