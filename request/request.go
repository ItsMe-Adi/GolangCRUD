package request

type CreateUserRequest struct {
	Username string `form:"username"  binding:"required"`
	Role     string `form:"role" binding:"required,valid-role"`
	Password string `form:"password" binding:"required,min=5"`
}

type GetUserRequest struct {
	ID int `uri:"id" binding:"required,min=1,max=10"`
}

type GetUserListsRequest struct {
	PageNo   int `form:"page_no" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=10"`
}

type LoginRequest struct {
	Username string `form:"username"  binding:"required"`
	Password string `form:"password" binding:"required,min=5"`
}