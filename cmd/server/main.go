package main

import (
	"log"
	"net/http"

	"github.com/ashimrai123/go-mcp-toolkit/internals/rpc"
	"github.com/ashimrai123/go-mcp-toolkit/internals/tools"
	"github.com/ashimrai123/go-mcp-toolkit/internals/tools/echo"
	"github.com/ashimrai123/go-mcp-toolkit/internals/transport"
)

func main() {
	tools.Register(echo.Skill{})
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/message", rpc.Handle)

	mux.HandleFunc("/sse", transport.SSEHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
