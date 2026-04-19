package rpc

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ashimrai123/go-mcp-toolkit/internals/mcp"
	"github.com/ashimrai123/go-mcp-toolkit/internals/tools"
	"github.com/ashimrai123/go-mcp-toolkit/internals/types"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, types.ErrorResponse(nil, -32700, "parse error"))
		return
	}

	log.Printf("method=%s id=%v", req.Method, req.ID)

	switch req.Method {
	case "initialize":
		writeJSON(w, mcp.HandleInitialize(req))
	case "notifications/initialized":
		mcp.HandleInitialized(req) // no response
	case "tools/list":
		writeJSON(w, types.Response{
			JSONRPC: "2.0",
			ID:      req.ID,
			Result:  map[string]any{"tools": tools.List()},
		})
	default:
		writeJSON(w, types.ErrorResponse(req.ID, -32601, "method not found"))
	}

}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
