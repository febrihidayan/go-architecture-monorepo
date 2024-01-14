package request

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required|min:3|email"`
	Password string `json:"password" validate:"required|min:6"`
}

type AuthRegisterRequest struct {
	Name            string `json:"name" validate:"required|min:3"`
	Email           string `json:"email" validate:"required|min:3|email"`
	Password        string `json:"password" validate:"required|min:6"`
	ConfirmPassword string `json:"confirm_password" validate:"required|min:6|same:password"`
}

type AuthSendEmailVerifiedRequest struct {
	Email string `json:"email" validate:"required|min:3|email"`
}

type AuthEmailRequest struct {
	Email string `json:"email" validate:"required|min:3|email"`
}

type AuthPasswordResetRequest struct {
	Token           string `json:"token" validate:"required"`
	Password        string `json:"password" validate:"required|min:6"`
	ConfirmPassword string `json:"confirm_password" validate:"required|min:6|same:password"`
}
