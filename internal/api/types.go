package api

type App struct {
	Name          string `json:"name"`
	Repo          string `json:"repo"`
	Port          int    `json:"port"`
	ContainerPort int    `json:"container_port"`
	Status        string `json:"status"`
	Error         string `json:"error,omitempty"`
}
