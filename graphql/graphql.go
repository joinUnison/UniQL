package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
)

type Server struct {
	Gqlgen *handler.Server
}

type ServerConfig struct {
	Gqlgen graphql.ExecutableSchema
}

func NewServer(config ServerConfig) *Server {
	server := handler.NewDefaultServer(config.Gqlgen)

	return &Server{
		Gqlgen: server,
	}
}
