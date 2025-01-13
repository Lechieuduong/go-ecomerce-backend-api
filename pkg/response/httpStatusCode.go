package response

const (
	ErrorCodeSusscess = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Email is invalid
	ErrorInvalidToken = 30001 // token is invalid
)

// message
var msg = map[int]string{
	ErrorCodeSusscess: "Success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorInvalidToken: "Token is invalid",
}