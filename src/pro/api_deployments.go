package pro

import (
	"net/http"
	"sync"

	"github.com/azukaar/cosmos-server/src/docker"
	"github.com/azukaar/cosmos-server/src/utils"
	"github.com/nats-io/nats.go"
)

type Deployment struct {
	Name     string                            `json:"name" validate:"required,min=3,max=64,alphanum"`
	Replicas int                               `json:"replicas" validate:"required,min=1"`
	Strategy string                            `json:"strategy" validate:"omitempty,oneof=round-robin least-busy"`
	Tags     []string                          `json:"tags,omitempty" validate:"omitempty,dive,min=1,max=64"`
	Storage  []string                          `json:"storage,omitempty" validate:"omitempty,dive,min=1,max=64"`
	Compose  docker.DockerServiceCreateRequest `json:"compose" validate:"required"`
}

func DeploymentsRoute(w http.ResponseWriter, req *http.Request, lock *sync.RWMutex, js nats.JetStreamContext) {
	utils.Error("This is a pro and is not currently available on your server. Please upgrade to Cosmos Pro to access this feature.", nil)
	utils.HTTPError(w, "This feature is only available in Cosmos Pro", http.StatusForbidden, "PRO001")
	return
}

func DeploymentsIdRoute(w http.ResponseWriter, req *http.Request, lock *sync.RWMutex, js nats.JetStreamContext) {
	utils.Error("This is a pro and is not currently available on your server. Please upgrade to Cosmos Pro to access this feature.", nil)
	utils.HTTPError(w, "This feature is only available in Cosmos Pro", http.StatusForbidden, "PRO001")
	return
}
