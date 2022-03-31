package models

// We may need to add location later
// City is a city obviously
type City struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	PostCode      string `json:"post_code"`
	Prononciation string `json:"prononciation"`
}

// Exists checks if a city exists
func (c City) Exists() bool {
	return c.ID != ""
}
