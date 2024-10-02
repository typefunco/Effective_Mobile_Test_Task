package entity

// Admin can delete, add, edit songs. Usual users only can only read

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	isAdmin  bool   // false = not Admin
}
