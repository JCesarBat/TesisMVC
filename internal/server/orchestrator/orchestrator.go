package orchestrator

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/internal/server/asociado"
	"TesisMVC/internal/server/auth"
	"TesisMVC/internal/server/common_data"
	"TesisMVC/pkg/util"
)

type Orchestrator struct {
	Auth     *auth.Server
	Asociado *asociado.Server
}

func NewOrchestrator(store database.Store, config util.Config) (*Orchestrator, error) {
	server, err := common_data.NewServer(store, config)
	if err != nil {
		return nil, err
	}
	return &Orchestrator{
		Auth: &auth.Server{
			GinServer: server,
			Data:      make(map[string]any),
		},
		Asociado: &asociado.Server{
			GinServer: server,
			Data:      make(map[string]any),
		},
	}, nil
}
