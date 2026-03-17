# Pull Request: Dynamic parallel_tool_calls based on disable_parallel_tool_use

## Summary

This PR implements dynamic control of the `parallel_tool_calls` parameter based on Claude's `disable_parallel_tool_use` parameter in the Claude to Codex request translator.

## Problem Statement

根据claude的disable_parallel_tool_use参数动态决定parallel_tool_calls，客户端没有值时才使用默认值true

(Dynamically decide parallel_tool_calls based on Claude's disable_parallel_tool_use parameter, using default value true only when client doesn't provide a value)

## Changes

### Modified Files:
1. **`internal/translator/codex/claude/codex_claude_request.go`** (Lines 217-225)
   - Added logic to read `disable_parallel_tool_use` from incoming Claude API requests
   - Dynamically sets `parallel_tool_calls` based on the parameter value
   - Maintains backward compatibility with default value `true`

2. **`internal/translator/codex/claude/codex_claude_request_test.go`** (New file)
   - Added comprehensive unit tests covering all scenarios

3. **`go.mod`**
   - Minor dependency reordering from `go mod tidy`

### Implementation Details:

```go
// Set parallel_tool_calls based on disable_parallel_tool_use parameter.
// If disable_parallel_tool_use is provided, use its inverse value.
// Otherwise, default to true.
parallelToolCalls := true
if disableParallelToolUse := rootResult.Get("disable_parallel_tool_use"); disableParallelToolUse.Exists() {
    parallelToolCalls = !disableParallelToolUse.Bool()
}
template, _ = sjson.Set(template, "parallel_tool_calls", parallelToolCalls)
```

**Behavior:**
- When `disable_parallel_tool_use=true` → `parallel_tool_calls=false`
- When `disable_parallel_tool_use=false` → `parallel_tool_calls=true`
- When `disable_parallel_tool_use` is not provided → `parallel_tool_calls=true` (default)

## Testing

Added comprehensive unit tests in `codex_claude_request_test.go`:

```go
func TestParallelToolCallsWithDisableParameter(t *testing.T) {
    // Test 1: disable_parallel_tool_use=true should set parallel_tool_calls to false
    // Test 2: disable_parallel_tool_use=false should set parallel_tool_calls to true
    // Test 3: no disable_parallel_tool_use should default parallel_tool_calls to true
}
```

**Test Results:**
```
=== RUN   TestParallelToolCallsWithDisableParameter
=== RUN   TestParallelToolCallsWithDisableParameter/disable_parallel_tool_use=true_should_set_parallel_tool_calls_to_false
=== RUN   TestParallelToolCallsWithDisableParameter/disable_parallel_tool_use=false_should_set_parallel_tool_calls_to_true
=== RUN   TestParallelToolCallsWithDisableParameter/no_disable_parallel_tool_use_should_default_parallel_tool_calls_to_true
--- PASS: TestParallelToolCallsWithDisableParameter (0.00s)
PASS
ok  	github.com/router-for-me/CLIProxyAPI/v6/internal/translator/codex/claude	0.003s
```

✅ All tests pass
✅ Code builds successfully
✅ Code formatted with `gofmt`

## Backward Compatibility

This change maintains full backward compatibility:
- Existing clients that don't send `disable_parallel_tool_use` will continue to get `parallel_tool_calls=true` (the previous hardcoded behavior)
- Only clients that explicitly send `disable_parallel_tool_use=true` will get `parallel_tool_calls=false`

## Related Information

- Branch: `claude/update-parallel-tool-calls`
- Commits:
  - `66e509e` - feat(translator): dynamically set parallel_tool_calls based on disable_parallel_tool_use parameter
  - `cd93a65` - Initial plan

## Checklist

- [x] Code changes implemented
- [x] Unit tests added and passing
- [x] Code builds successfully
- [x] Code formatted with gofmt
- [x] Backward compatibility maintained
- [x] Documentation (comments) added

---

## How to Create the PR

Since direct API access is restricted in this environment, please create the PR manually using one of these methods:

### Method 1: GitHub Web Interface
1. Go to https://github.com/Darley-Wey/CLIProxyAPI
2. Click "Contribute" → "Open pull request"
3. Set base repository to `router-for-me/CLIProxyAPI` and base branch to `main`
4. Set head repository to `Darley-Wey/CLIProxyAPI` and compare branch to `claude/update-parallel-tool-calls`
5. Copy the content above as the PR description

### Method 2: GitHub CLI (from local environment)
```bash
gh pr create \
  --repo router-for-me/CLIProxyAPI \
  --base main \
  --head Darley-Wey:claude/update-parallel-tool-calls \
  --title "feat: dynamically set parallel_tool_calls based on disable_parallel_tool_use parameter" \
  --body-file PR_DESCRIPTION.md
```

### Method 3: Direct URL
Visit: https://github.com/router-for-me/CLIProxyAPI/compare/main...Darley-Wey:CLIProxyAPI:claude/update-parallel-tool-calls
