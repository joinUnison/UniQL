package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"testing"
)

type MockExecutableSchema struct {
	SchemaFunc     func() *ast.Schema
	ComplexityFunc func(typeName, fieldName string, childComplexity int, args map[string]interface{}) (int, bool)
	ExecFunc       func(ctx context.Context) graphql.ResponseHandler
}

func (m MockExecutableSchema) Schema() *ast.Schema {
	return m.SchemaFunc()
}

func (m MockExecutableSchema) Complexity(typeName, fieldName string, childComplexity int, args map[string]interface{}) (int, bool) {
	return m.ComplexityFunc(typeName, fieldName, childComplexity, args)
}

func (m MockExecutableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	return m.ExecFunc(ctx)
}

func TestNewServer_Gqlgen_Success(t *testing.T) {
	schema := &MockExecutableSchema{}

	config := ServerConfig{
		Gqlgen: schema,
	}

	server := NewServer(config)

	if server == nil {
		t.Fatal("NewServer returned nil, expected a server instance")
	}

	if server.Gqlgen == nil {
		t.Fatal("server.gqlgen is nil, expected a handler.Server instance")
	}
}
