package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
)

type Server struct {
	gqlgen *handler.Server
}

type ServerConfig struct {
	gqlgen graphql.ExecutableSchema
}

func NewServer(config ServerConfig) *Server {
	server := handler.NewDefaultServer(config.gqlgen)

	return &Server{
		gqlgen: server,
	}
}
