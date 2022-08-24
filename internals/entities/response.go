package entities

type Response struct {
	Status     string `json:"status"`
	StatusCode uint32 `json:"status_code"`
	Message    string `json:"message"`
	Result     Result `json:"result"`
}

type Result struct {
	Data any `json:"data"`
}
