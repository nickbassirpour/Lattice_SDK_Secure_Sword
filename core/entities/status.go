package entities

type Status struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
	Details string[] `json:"details"`
}
