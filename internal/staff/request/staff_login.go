package request

type StaffLoginRequest struct {
	PhoneNumber string `form:"phoneNumber" binding:"required,min=10,max=16"`
	Password    string `form:"password" binding:"required"`
}
