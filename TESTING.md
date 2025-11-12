# Testing Guide

This document describes the comprehensive testing strategy for the Starbucks Coffee Shop application.

## Testing Architecture

The project uses a **fan-in/fan-out** testing pattern in CircleCI that runs tests in parallel and gates deployment on all tests passing.

### Test Pyramid

```
        /\
       /  \
      / E2E \
     /--------\
    /Integration\
   /--------------\
  /   Unit Tests   \
 /------------------\
```

## Test Types

### 1. Security Scanning
- **Tool**: TruffleHog
- **Purpose**: Detect leaked secrets and credentials
- **Runs**: On every commit (except main branch)
- **Blocks**: All other tests if secrets are found

### 2. Backend Tests

#### Unit Tests
```bash
# Run all Go tests
go test -v ./...

# Run with coverage
go test -v -race -coverprofile=coverage.out ./...

# View coverage report
go tool cover -html=coverage.out
```

**Location**: `graph/resolver_test.go`

**Tests**:
- Resolver functionality
- GraphQL queries and mutations
- Concurrent operations
- Error handling

#### Linting
```bash
# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s

# Run linters
golangci-lint run

# Check formatting
gofmt -l .
```

**Configuration**: `.golangci.yml`

### 3. Frontend Tests

#### Unit Tests (Vitest)
```bash
cd frontend

# Run tests
npm test

# Run with UI
npm run test:ui

# Run with coverage
npm run test:coverage
```

**Configuration**: `frontend/vitest.config.ts`

#### Type Checking
```bash
cd frontend

# TypeScript type checking
npx tsc --noEmit
```

#### Linting
```bash
cd frontend

# ESLint
npm run lint
```

### 4. Integration Tests

Tests the full stack integration between frontend and backend.

```bash
# Start backend
go run .

# Run integration tests (in CircleCI)
# Tests GraphQL API endpoints
# Tests data flow
# Tests error handling
```

**Tests**:
- GraphQL query execution
- Mutation operations
- Data persistence
- API response validation

### 5. E2E Tests (Playwright)

End-to-end tests that simulate real user interactions.

```bash
cd frontend

# Install Playwright browsers
npx playwright install

# Run E2E tests
npm run test:e2e

# Run with UI
npm run test:e2e:ui

# Debug mode
npm run test:e2e:debug

# Run specific browser
npx playwright test --project=chromium
```

**Location**: `frontend/e2e/coffee-shop.spec.ts`

**Test Scenarios**:
- ✅ Display header and branding
- ✅ Load and display coffee items
- ✅ Add items to cart
- ✅ Update cart quantities
- ✅ Remove items from cart
- ✅ Complete purchase flow
- ✅ Handle out-of-stock scenarios
- ✅ Responsive design (mobile/desktop)
- ✅ GraphQL API integration

**Configuration**: `frontend/playwright.config.ts`

### 6. Performance Tests (k6)

Load testing to ensure the application can handle traffic.

```bash
# Install k6
brew install k6  # macOS
# or follow: https://k6.io/docs/getting-started/installation/

# Run load tests (example)
k6 run performance-test.js
```

**Test Scenarios**:
- Ramp up to 10 concurrent users
- Sustained load for 1 minute
- Response time thresholds (p95 < 500ms)
- Error rate monitoring

## CircleCI Workflow

### Fan-Out Phase (Parallel Execution)

1. **Security Gate**: TruffleHog scans for secrets
2. **Backend Jobs** (parallel):
   - Linting
   - Unit tests
   - Build verification
3. **Frontend Jobs** (parallel):
   - Linting
   - Type checking
   - Unit tests
   - Build

### Fan-In Phase (Integration)

4. **Integration Tests**: Requires backend + frontend builds
5. **E2E Tests**: Requires backend + frontend builds
6. **Performance Tests**: Requires backend + frontend builds

### Deployment Gate

7. **Deploy**: Only runs if ALL tests pass

### Workflow Diagram

```
                    [scan-secrets]
                          |
        +-----------------+------------------+
        |                                    |
    [Backend]                           [Frontend]
    ├─ lint                             ├─ lint
    ├─ test                             ├─ typecheck
    └─ build                            ├─ test
        |                               └─ build
        |                                    |
        +----------------+-------------------+
                         |
            +------------+------------+
            |            |            |
      [integration]   [e2e]    [performance]
            |            |            |
            +------------+------------+
                         |
                    [deploy] ✅
```

## Running Tests Locally

### Quick Test Run
```bash
# Backend
go test ./...

# Frontend
cd frontend && npm test

# E2E (requires running server)
go run . &
cd frontend && npm run test:e2e
```

### Full Test Suite
```bash
# 1. Security scan
trufflehog git file://. --since-commit main

# 2. Backend
golangci-lint run
go test -v -race -coverprofile=coverage.out ./...

# 3. Frontend
cd frontend
npm run lint
npx tsc --noEmit
npm run test:coverage

# 4. Build
npm run build
cd ..
go build -o starbucks .

# 5. Integration & E2E
./starbucks &
SERVER_PID=$!
cd frontend
npm run test:e2e
kill $SERVER_PID
```

## Test Data Management

### Backend Test Data
- Defined in `graph/resolver_test.go`
- Uses in-memory data structures
- Isolated per test

### Frontend Test Data
- Mock GraphQL responses
- Playwright fixtures
- Test-specific data in E2E tests

## Continuous Integration

### On Pull Request
- ✅ Security scan
- ✅ All linting
- ✅ All unit tests
- ✅ Integration tests
- ✅ E2E tests
- ✅ Performance tests

### On Merge to Main
- ✅ Full test suite
- ✅ Deploy to production (if all pass)

### Nightly
- ✅ Comprehensive test run
- ✅ Extended performance tests
- ✅ Security audit

## Test Coverage Goals

- **Backend**: > 80% code coverage
- **Frontend**: > 70% code coverage
- **E2E**: All critical user paths
- **Integration**: All API endpoints

## Adding New Tests

### Backend Unit Test
```go
// graph/resolver_test.go
func TestNewFeature(t *testing.T) {
    resolver := NewResolver()
    // ... test implementation
}
```

### Frontend Unit Test
```typescript
// frontend/src/components/__tests__/Component.test.tsx
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/react';

describe('Component', () => {
  it('should render', () => {
    // ... test implementation
  });
});
```

### E2E Test
```typescript
// frontend/e2e/new-feature.spec.ts
import { test, expect } from '@playwright/test';

test('new feature works', async ({ page }) => {
  await page.goto('/');
  // ... test implementation
});
```

## Troubleshooting

### Tests Failing Locally But Passing in CI
- Check Node.js/Go versions match CI
- Clear caches: `go clean -cache`, `npm ci`
- Check for environment-specific issues

### E2E Tests Timing Out
- Increase timeout in `playwright.config.ts`
- Check if server is starting correctly
- Verify network conditions

### Flaky Tests
- Add explicit waits
- Use `waitForSelector` instead of fixed timeouts
- Check for race conditions

## Resources

- [Vitest Documentation](https://vitest.dev/)
- [Playwright Documentation](https://playwright.dev/)
- [k6 Documentation](https://k6.io/docs/)
- [golangci-lint](https://golangci-lint.run/)
- [CircleCI Documentation](https://circleci.com/docs/)
