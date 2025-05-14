package handlers

var jwtSecret = []byte("secret")

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//func login(w http.ResponseWriter, r *http.Request) {
//	var req loginRequest
//
//	user, err := GetUserByEmail(req.Email)
//	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password)
//}
