package api

import (
	"encoding/json"
	"net/http"

	tf "github.com/abhinav-dops/mini-paas-platform/internal/terraform"
)

const terraformPath = "./terraform"

func ProvisionInfra(w http.ResponseWriter, _ *http.Request) {
	if Infra.Status == "ready" {
		json.NewEncoder(w).Encode(Infra)
		return
	}

	Infra.Status = "pending"

	go func() {
		if err := tf.Init(terraformPath); err != nil {
			Infra.Status = "failed"
			Infra.Error = err.Error()
			return
		}
		if err := tf.Apply(terraformPath); err != nil {
			Infra.Status = "failed"
			Infra.Error = err.Error()
			return
		}
		ip, err := tf.Output(terraformPath)
		if err != nil {
			Infra.Status = "failed"
			Infra.Error = err.Error()
			return
		}
		Infra.IP = ip
		Infra.Status = "ready"
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Infra)
}

func GetInfraStatus(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(Infra)
}

func DestroyInfra(w http.ResponseWriter, _ *http.Request) {
	go func() {
		tf.Destroy(terraformPath)
		Infra.Status = "not_created"
		Infra.IP = ""
	}()
	w.WriteHeader(http.StatusAccepted)
}
