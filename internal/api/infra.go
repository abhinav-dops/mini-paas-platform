package api

type InfraStatus struct {
	Status string `json:"status"`
	Ip     string `json:"ip,omitempty"`
	Error  string `json:"error,omitempty"`
}

var Infra = InfraStatus{
	Status: "not_created",
}
