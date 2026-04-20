package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	// send the endpoint event - tells client where to POST messages
	fmt.Fprintf(w, "event: endpoint\ndata: /message\n\n")
	flusher.Flush()
	// keep connection alive until client disconnects
	<-r.Context().Done()

}

func WriteEvent(w http.ResponseWriter, event string, data any) {
	b, _ := json.Marshal(data)
	fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, string(b))
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

}
