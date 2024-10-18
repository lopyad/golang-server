package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// -----CREATE-----

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type CreateUserResponse struct {
	*ApiResponse
	*User
}

// -----GET-----

type GetRequest struct {
}
type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

// -----UPDATE-----

type UpdateRequest struct {
	Name      string `json:"name" binding:"required"`
	UpdateAge int64  `json:"updateAge" binding:"required"`
}
type UpdateUserResponse struct {
	*ApiResponse
	*User
}

// -----DELETE-----

type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *DeleteRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}

type DeleteUserResponse struct {
	*ApiResponse
}
