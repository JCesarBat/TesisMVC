package common_data

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/internal/token"
	"TesisMVC/pkg/util"
	"fmt"
)

type GinServer interface {
	GetStore() database.Store
	GetTokenMaker() token.Maker
	GetConfig() util.Config
	SetStore(store database.Store)
	GetData() map[string]any
	GetCookie() util.Cookie
	ResetData()
}

type Server struct {
	store      database.Store
	tokenMaker token.Maker
	config     util.Config
	data       map[string]any
	cookie     util.Cookie
}

func NewServer(store database.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetrickey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
		cookie:     util.Cookie{},
		data:       make(map[string]any),
	}, nil
}

func (s *Server) GetStore() database.Store {
	return s.store
}
func (s *Server) GetTokenMaker() token.Maker {
	return s.tokenMaker
}
func (s *Server) GetConfig() util.Config {
	return s.config
}
func (s *Server) SetStore(store database.Store) {
	s.store = store

}

func (s *Server) GetCookie() util.Cookie {
	return s.cookie
}
func (s *Server) GetData() map[string]any {
	return s.data
}

func (s *Server) ResetData() {
	s.data = make(map[string]any)

}
