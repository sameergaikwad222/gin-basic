package utility

type UserQueryParams struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	IdList   []int  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
