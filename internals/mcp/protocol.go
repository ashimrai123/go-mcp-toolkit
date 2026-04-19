package mcp

import (
	"github.com/ashimrai123/go-mcp-toolkit/internals/types"
)

func HandleInitialize(req types.Request) types.Response {

	result := map[string]any{
		"protocolVersion": "2024-11-05",
		"serverInfo": map[string]any{
			"name":    "go-mcp-server",
			"version": "0.1.0",
		},
		"capabilities": map[string]any{
			"tools": map[string]any{},
		},
	}

	return types.Response{JSONRPC: "2.0", ID: req.ID, Result: result}

}

func HandleInitialized(r types.Request) types.Response {

	return types.Response{}
}
