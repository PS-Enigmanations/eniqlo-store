package request

type StaffRegisterRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,min=10,max=16,regex=^\+.*$"`
	Name        string `json:"name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
