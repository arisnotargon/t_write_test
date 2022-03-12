package response

type OkResponse struct {
	Ok bool `json:"ok"`
}

type AgeResponse struct {
	Age int32 `json:"age"`
}

type CarResponse struct {
	Car string `json:"car"`
}

type RateResponse struct {
	Rate int `json:"rate"`
}

type BufferResponse struct {
	Buffer int `json:"buffer"`
}
