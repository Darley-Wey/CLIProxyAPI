package claude

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestParallelToolCallsWithDisableParameter(t *testing.T) {
	tests := []struct {
		name                      string
		requestJSON               string
		expectedParallelToolCalls bool
	}{
		{
			name: "disable_parallel_tool_use=true should set parallel_tool_calls to false",
			requestJSON: `{
				"model": "claude-3-5-sonnet-20241022",
				"disable_parallel_tool_use": true,
				"messages": [{"role": "user", "content": "Hello"}]
			}`,
			expectedParallelToolCalls: false,
		},
		{
			name: "disable_parallel_tool_use=false should set parallel_tool_calls to true",
			requestJSON: `{
				"model": "claude-3-5-sonnet-20241022",
				"disable_parallel_tool_use": false,
				"messages": [{"role": "user", "content": "Hello"}]
			}`,
			expectedParallelToolCalls: true,
		},
		{
			name: "no disable_parallel_tool_use should default parallel_tool_calls to true",
			requestJSON: `{
				"model": "claude-3-5-sonnet-20241022",
				"messages": [{"role": "user", "content": "Hello"}]
			}`,
			expectedParallelToolCalls: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertClaudeRequestToCodex("claude-3-5-sonnet-20241022", []byte(tt.requestJSON), false)
			parallelToolCalls := gjson.GetBytes(result, "parallel_tool_calls").Bool()

			if parallelToolCalls != tt.expectedParallelToolCalls {
				t.Errorf("Expected parallel_tool_calls=%v, got %v", tt.expectedParallelToolCalls, parallelToolCalls)
			}
		})
	}
}
