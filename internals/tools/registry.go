package tools

import "context"

type Tool interface {
	Name() string
	Description() string
	InputSchema() map[string]any
	Execute(ctx context.Context, params map[string]any) (string, error)
}

var registry = map[string]Tool{}

func Register(t Tool) {
	registry[t.Name()] = t
}

func List() []map[string]any {
	out := []map[string]any{}
	for _, t := range registry {
		out = append(out, map[string]any{
			"name":        t.Name(),
			"description": t.Description(),
			"inputSchema": t.InputSchema(),
		})
	}
	return out
}

func Get(name string) (Tool, bool) {
	t, ok := registry[name]
	return t, ok
}
