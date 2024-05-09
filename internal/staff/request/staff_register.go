package request

type StaffRegisterRequest struct {
	PhoneNumber string `form:"phoneNumber" binding:"required,min=10,max=16,startswith=+"`
	Name        string `form:"name" binding:"required"`
	Password    string `form:"password" binding:"required"`
}
