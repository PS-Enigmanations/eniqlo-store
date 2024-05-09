package response

// StaffResponse represents a staff response.
type StaffResponse struct {
	ID          string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
