package auth

type PhoneAuthRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type PhoneAuthResponce struct {
	SessionId string `json:"session_id" validate:"required"`
}

type PhoneAuthCodeRequest struct {
	SessionId string `json:"session_id" validate:"required"`
	Code      string `json:"code" validate:"required,numeric"`
}

type PhoneAuthCodeResponce struct {
	Token string `json:"token" validate:"required,jwt"`
}
