package models



type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      int `json:"expiry"`
}

type Response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          int `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset int`json:"rate_limit_reset"`
}
