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
	case "tools/call":
		name, _ := req.Params["name"].(string)
		params, _ := req.Params["arguments"].(map[string]any)
		tool, ok := tools.Get(name)
		if !ok {
			writeJSON(w, types.ErrorResponse(req.ID, -32601, "unknown tool: "+name))
			return
		}
		result, err := tool.Execute(r.Context(), params)
		if err != nil {
			writeJSON(w, types.ErrorResponse(req.ID, -32603, err.Error()))
			return
		}
		writeJSON(w, types.Response{
			JSONRPC: "2.0",
			ID:      req.ID,
			Result: map[string]any{
				"content": []map[string]any{
					{"type": "text", "text": result},
				},
			},
		})
	default:
		writeJSON(w, types.ErrorResponse(req.ID, -32601, "method not found"))
	}

}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
