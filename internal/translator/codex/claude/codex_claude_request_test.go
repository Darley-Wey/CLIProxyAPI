package claude

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestConvertClaudeRequestToCodex_ParallelToolCalls(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "default true when disable_parallel_tool_use is absent",
			input: `{
				"model": "claude-3-opus",
				"messages": [{"role":"user","content":[{"type":"text","text":"hi"}]}]
			}`,
			expected: true,
		},
		{
			name: "false when disable_parallel_tool_use is true",
			input: `{
				"model": "claude-3-opus",
				"disable_parallel_tool_use": true,
				"messages": [{"role":"user","content":[{"type":"text","text":"hi"}]}]
			}`,
			expected: false,
		},
		{
			name: "true when disable_parallel_tool_use is false",
			input: `{
				"model": "claude-3-opus",
				"disable_parallel_tool_use": false,
				"messages": [{"role":"user","content":[{"type":"text","text":"hi"}]}]
			}`,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ConvertClaudeRequestToCodex("test-model", []byte(tt.input), false)
			got := gjson.GetBytes(output, "parallel_tool_calls")
			if !got.Exists() {
				t.Fatalf("parallel_tool_calls not found in output: %s", string(output))
			}
			if got.Bool() != tt.expected {
				t.Fatalf("parallel_tool_calls = %v, want %v; output: %s", got.Bool(), tt.expected, string(output))
			}
		})
	}
}
