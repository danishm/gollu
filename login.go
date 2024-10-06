package gollu

type LLULoginResponse struct {
	Status int64
	Data   LLULoginResponseData
	Error  *LLULoginError
}

type LLULoginResponseData struct {
	AuthTicket LLLULoginResponseAuthTicket
}

type LLULoginError struct {
	Message string
}
