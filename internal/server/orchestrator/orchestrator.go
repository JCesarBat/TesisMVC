package orchestrator

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/internal/server/common_data"
	"TesisMVC/pkg/util"
)

type Orchestrator struct {
}

func NewOrchestrator(store database.Store, config util.Config) (*Orchestrator, error) {
	_, err := common_data.NewServer(store, config)
	if err != nil {
		return nil, err
	}
	return &Orchestrator{}, nil
}
