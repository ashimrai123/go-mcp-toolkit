package echo

import (
	"context"
	"fmt"
)

type Skill struct{}

func (s Skill) Name() string { return "echo" }

func (s Skill) Description() string { return "Returns whatever text you send it" }

func (s Skill) InputSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"message": map[string]any{
				"type":        "string",
				"description": "Text to echo back",
			},
		},
		"required": []string{"message"},
	}
}

func (s Skill) Execute(c context.Context, params map[string]any) (string, error) {
	msg, ok := params["message"].(string)
	if !ok {
		return "", fmt.Errorf("message param missing or not a string")
	}
	return msg, nil
}
