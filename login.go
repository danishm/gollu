package gollu

type LLULoginResponse struct {
	Status int64
	Data   LLULoginResponseData
	Error  *LLULoginError
}

type LLULoginResponseData struct {
	User       LLUUserInfo
	AuthTicket LLLULoginResponseAuthTicket
}

type LLULoginError struct {
	Message string
}

type LLUUserInfo struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}
