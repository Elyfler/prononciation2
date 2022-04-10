package models

// We may need to add location later
// City is a city obviously
type City struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	PostCode      string `json:"post_code,omitempty"`
	Prononciation string `json:"prononciation,omitempty"`
}

// Exists checks if a city exists
func (c City) Exists() bool {
	return c.ID != ""
}
