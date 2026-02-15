# CLAUDE.md - Gotenberg Go Client

This document provides comprehensive guidance for AI assistants working with the Gotenberg Go Client codebase.

## Project Overview

**Purpose**: A Go client library for interacting with the Gotenberg API (v7+), which provides PDF generation services.

**Repository**: github.com/commitsmart/gotenberg-go-client

**Language**: Go 1.19+

**Inspiration**: Based on [thecodingmachine/gotenberg-go-client](https://github.com/thecodingmachine/gotenberg-go-client)

## Architecture Overview

### Core Components

1. **Client** (`client.go`): HTTP client for communicating with Gotenberg API
   - `Client` struct: Main client with hostname and HTTP client
   - `Request` interface: Common interface for all request types
   - `Post()`: Sends requests to Gotenberg
   - `Store()`: Convenience method to save PDF to file

2. **Document** (`document.go`): Abstraction for file inputs
   - `Document` interface: Common interface for all document types
   - `NewDocumentFromPath()`: Create from file path
   - `NewDocumentFromString()`: Create from string content
   - `NewDocumentFromBytes()`: Create from byte array

3. **Chromium Module** (`chromium.go`, `html.go`): HTML to PDF conversion
   - `chromiumRequest`: Base request with Chromium-specific options
   - `ConvertHTMLRequest`: HTML conversion with headers, footers, assets
   - Supports page properties, margins, CSS, JavaScript options

4. **LibreOffice Module** (`office.go`): Office documents to PDF
   - `OfficeRequest`: Convert Word, Excel, PowerPoint to PDF
   - Supports landscape, page ranges, PDF formats, merging

5. **PDF Engines Module** (`merge.go`): PDF manipulation
   - `MergeRequest`: Merge multiple PDFs into one

### Request Flow

```
User Code
    ↓
Create Document(s) (NewDocumentFromPath/String/Bytes)
    ↓
Create Request (NewConvertHTMLRequest/NewOfficeRequest/NewMergeRequest)
    ↓
Configure Request (set options like margins, paper size, etc.)
    ↓
Client.Post() or Client.Store()
    ↓
multipartForm() - builds multipart/form-data
    ↓
HTTP POST to Gotenberg API
    ↓
Response (PDF bytes) or Error
```

## File Structure

```
gotenberg-go-client/
├── client.go           # Core client and request infrastructure
├── document.go         # Document abstraction (path/string/bytes)
├── chromium.go         # Chromium/HTML base request with options
├── html.go            # HTML to PDF conversion request
├── office.go          # Office documents to PDF conversion
├── merge.go           # PDF merging functionality
├── html_test.go       # HTML conversion tests
├── office_test.go     # Office conversion tests
├── merge_test.go      # Merge tests
├── go.mod             # Go module definition
├── go.sum             # Go dependency checksums
├── docker-compose.test.yml  # Testing environment setup
├── test.Dockerfile    # Docker image for running tests
├── test/
│   ├── test.go        # Test utilities
│   └── testdata/      # Test files (HTML, Office docs)
│       ├── html/      # HTML test files
│       └── office/    # Office test files
├── LICENSE            # License file
└── README.md          # User-facing documentation
```

## Development Conventions

### Code Organization

1. **Flat Structure**: Main code files are in the root directory
2. **One Request Type Per File**: Each conversion type has its own file
3. **Shared Base Types**: Common functionality in `client.go` and base request types
4. **Interface-Based Design**: `Request` and `Document` interfaces for extensibility

### Naming Conventions

1. **Constants**: camelCase for internal form field names (e.g., `paperWidth`, `marginTop`)
2. **Public API**: PascalCase for exported functions and types (e.g., `NewClient`, `PaperSize`)
3. **Files**: lowercase with underscores for tests (e.g., `html_test.go`)
4. **Request Types**: Descriptive names ending in "Request" (e.g., `ConvertHTMLRequest`)

### Go Idioms Used

1. **Compile-time Interface Checks**:
   ```go
   var _ Request = new(ConvertHTMLRequest)
   ```

2. **Functional Options Pattern**: Methods return nothing, modify request in place:
   ```go
   req.PaperSize(gotenberg.A4)
   req.Margins(gotenberg.NoMargins)
   ```

3. **Embedded Structs**: Shared functionality through composition:
   ```go
   type chromiumRequest struct {
       *request  // Embedded base request
   }
   ```

4. **Context Usage**: All API calls accept `context.Context` for cancellation

### Error Handling

1. **Return Errors**: Functions return `(value, error)` tuples
2. **Validate Inputs**: Check for invalid inputs (empty files, missing paths)
3. **Wrap Errors**: Add context using `fmt.Errorf()`
4. **Defer Close**: Always defer closing resources

### Testing Standards

1. **Framework**: Uses `github.com/stretchr/testify` for assertions
2. **Test Files**: Co-located with source (`*_test.go`)
3. **Test Names**: `TestFunctionName` or `TestFeatureDescription`
4. **Assertions**:
   - Use `require.Nil()` for critical checks that should stop the test
   - Use `assert.Nil()` for non-critical checks
   - Use `assert.FileExists()` to verify file creation

4. **Test Environment**:
   - Uses Docker Compose to run Gotenberg service
   - Client connects to `http://gotenberg:3000`
   - Tests create temp directories and clean up after

5. **Test Utilities**: Helper functions in `test/test.go`

## Development Workflow

### Setting Up Development Environment

1. **Prerequisites**:
   ```bash
   go version  # Should be 1.19+
   docker --version
   docker-compose --version
   ```

2. **Install Dependencies**:
   ```bash
   go mod download
   ```

3. **Install golangci-lint**:
   ```bash
   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
     sh -s -- -b $(go env GOPATH)/bin v1.50.1
   ```

### Running Tests

1. **Start Test Environment**:
   ```bash
   docker-compose -f docker-compose.test.yml up -d
   ```

2. **Run Tests**:
   ```bash
   go test ./...
   ```

3. **Run Tests in Docker** (recommended for CI/consistency):
   ```bash
   docker-compose -f docker-compose.test.yml run test-gotenberg-go-client
   ```

4. **Cleanup**:
   ```bash
   docker-compose -f docker-compose.test.yml down
   ```

### Linting

Run linter before committing:
```bash
golangci-lint run --timeout=2m ./...
```

The linter is also run automatically in the Docker test image.

### Adding New Features

1. **New Request Type**:
   - Create new file (e.g., `newfeature.go`)
   - Define request struct embedding appropriate base type
   - Implement `Request` interface methods:
     - `postURL()`: Return Gotenberg API endpoint
     - `formFiles()`: Return map of files to upload
     - `customHTTPHeaders()`: Inherited from base
     - `formValues()`: Inherited from base
   - Add compile-time interface check
   - Add constructor function
   - Create corresponding test file

2. **New Options**:
   - Add constant for form field name at top of file
   - Add method to set the value on appropriate request type
   - Document the option with a comment
   - Add test coverage

3. **Documentation**:
   - Update README.md with usage examples
   - Add godoc comments to exported functions/types
   - Reference official Gotenberg docs where applicable

## Key Patterns for AI Assistants

### When Reading Code

1. **Interface Satisfaction**: Check compile-time assertions at bottom of files
2. **Embedded Types**: Look for embedded structs to understand inheritance
3. **Constants**: Check top of files for available options
4. **Form Fields**: Constants map to Gotenberg API parameters

### When Writing Code

1. **Follow Builder Pattern**: Add methods that modify request and return nothing
2. **Use Existing Patterns**: Look at similar request types for consistency
3. **Add Tests**: Every new feature needs corresponding tests
4. **Validate Inputs**: Check for nil, empty, or invalid values
5. **Document Public API**: Add godoc comments to all exported items

### When Debugging

1. **Check Gotenberg Logs**: The Docker container logs show API errors
2. **Verify Form Data**: Print `req.formValues()` and `req.formFiles()`
3. **Check HTTP Response**: Inspect `resp.StatusCode` and `resp.Body`
4. **Test Isolation**: Ensure tests clean up (remove temp files)

### Common Gotcha's

1. **Nil HTTP Client**: Client creation accepts nil (uses default), but test it
2. **Webhook vs Store**: Can't use `Store()` with webhook URLs
3. **File Paths**: Must use absolute paths, not relative
4. **Context Cancellation**: Always pass context from caller, don't create new ones
5. **Defer Errors**: Marked with `// nolint: errcheck` where intentional

## API Modules Status

| Module            | Implementation | Files |
|:------------------|:--------------:|:------|
| Chromium          | ✅ | `chromium.go`, `html.go` |
| LibreOffice       | ✅ | `office.go` |
| PDF Engines       | Partial | `merge.go` (merge only) |
| Prometheus        | ❌ | - |
| Logging           | ❌ | - |
| Graceful Shutdown | ❌ | - |

## External Resources

- **Gotenberg Documentation**: https://gotenberg.dev/docs/about
- **Gotenberg API v7**: https://gotenberg.dev/docs/modules/api
- **Go Module**: github.com/commitsmart/gotenberg-go-client

## Git Workflow

### Branch Naming
- Feature branches should be descriptive: `feature/add-logging-support`
- Bug fixes: `fix/correct-margin-calculation`
- Always work on feature branches, never directly on main

### Commit Messages
- Follow conventional commits format
- Examples:
  - `feat: add support for PDF metadata`
  - `fix: correct paper size calculation for A4`
  - `test: add coverage for merge requests`
  - `docs: update README with new examples`
  - `chore: update dependencies`

### Testing Before Commit
1. Run linter: `golangci-lint run --timeout=2m ./...`
2. Run tests: `docker-compose -f docker-compose.test.yml run test-gotenberg-go-client`
3. Verify all tests pass before pushing

## Code Quality Standards

### Security
- Use `#nosec` comments sparingly and only when justified
- Validate all user inputs (file paths, URLs, form values)
- Be cautious with file permissions (0644 for files, 0755 for directories)
- Never log sensitive data (webhook URLs may contain tokens)

### Performance
- Reuse HTTP clients (don't create new ones per request)
- Use `defer` for cleanup but be aware of timing
- Buffer sizes are reasonable (uses Go stdlib defaults)
- No unnecessary allocations in hot paths

### Maintainability
- Keep functions small and focused
- Use meaningful variable names
- Add comments for non-obvious logic
- Follow Go proverbs (errors are values, interfaces are satisfied implicitly)

## Dependencies

### Direct Dependencies
- `github.com/stretchr/testify v1.8.1`: Testing framework

### Indirect Dependencies (test only)
- `github.com/davecgh/go-spew`: Pretty printing for tests
- `github.com/pmezard/go-difflib`: Diff output for tests
- `gopkg.in/yaml.v3`: YAML parsing for testify

### External Services
- **Gotenberg API**: Docker image `gotenberg/gotenberg:7`
  - Port: 3000
  - Required for running tests

## Future Considerations

When extending this library, consider:

1. **New Gotenberg Modules**: Prometheus, Logging, PDF Engines extensions
2. **Streaming Responses**: For very large PDF outputs
3. **Retry Logic**: For transient network failures
4. **Rate Limiting**: For high-volume usage
5. **Async/Webhook Support**: Better webhook handling utilities
6. **Configuration**: Builder pattern for Client initialization
7. **Mocking**: Test utilities for users to mock Gotenberg in their tests

## Questions to Ask When Making Changes

1. Does this change maintain backward compatibility?
2. Are there tests covering the new functionality?
3. Does the linter pass?
4. Is the documentation updated?
5. Does this follow existing patterns in the codebase?
6. Are errors properly wrapped with context?
7. Is cleanup properly handled (defer, close, remove)?
8. Are there any security implications?

---

**Last Updated**: 2025-11-14
**Go Version**: 1.19+
**Gotenberg Version**: 7+
