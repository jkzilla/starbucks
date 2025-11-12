# TruffleHog Configuration

This directory contains TruffleHog configuration and rules files.

## Rules

### `rules/trufflehog-api-openapi.json`

OpenAPI specification for the TruffleHog API. This file can be used to:
- Understand the TruffleHog API endpoints
- Generate API clients
- Validate API requests
- Document available endpoints and parameters

### `rules/may-2025-features.md`

Documentation of TruffleHog features released in May 2025, including:
- **New Detectors**: Deepseek, xAI, Harness (7 new detectors, 8 new key types)
- **Analyze Data Filter**: View secrets with permission analysis
- **Jira Integration**: Ticket IDs in secret details
- **Secret ID Search**: Search by secret identifier
- **Enhanced Administration**: Improved UI, pagination, filtering
- **Performance Improvements**: Faster location type filtering

### `rules/june-2025-features.md`

Documentation of TruffleHog features released in June 2025, including:
- **11 New Analyzers**: Plaid, Netlify, Fastly, Monday, Datadog, Ngrok, Mux, PostHog, Dropbox, Databricks, Jira
- **4 New Detectors**: Salesforce OAuth2, Lokalise, Bannerbear, Langsmith
- **806 Total Detectors**: Supporting 848 different key types
- **Manual Jira Control**: Resend/sync notifications from Remediations tab
- **Config Files in OSS**: Run multiple concurrent scans
- **Smarter Scans**: Parallelization, resume capability, detailed progress
- **Live Secret Notifications**: Alerts when secrets become active
- **Active Scan Sorting**: Sort by repo, status, duration, verified count

### `rules/latest-release-features.md`

Documentation of TruffleHog latest release customer-facing features:
- **CSV Dashboard Links**: Direct links to secrets in exports
- **Source in Emails**: Notification emails include source name
- **3 New Detectors**: Webex Bot, CircleCI v2, Salesforce OAuth2 (enhanced)
- **810 Total Detectors**: Supporting 854 different key types
- **Bulk Queue API**: Queue up to 100 notifications per request
- **IdP Role Mapping**: Auto-sync users from Identity Provider
- **Share Link Enhancements**: New share viewer role with auditing
- **Unauthorized Page**: Clear access restriction messaging
- **Improved Terminology**: "Queue Notification" (was "Create")
- **Enhanced API Docs**: Better grouping and navigation
- **Alphabetized Filters**: Sorted integration dropdown

See the full documentation for configuration examples and usage.

## Usage with TruffleHog

To use custom rules with TruffleHog:

```bash
# Scan with custom rules
trufflehog git file://. --rules .trufflehog/rules/

# Scan with specific rule file
trufflehog git file://. --config .trufflehog/config.yml
```

## TruffleHog API

The OpenAPI spec documents the TruffleHog Enterprise API, which provides:
- Secret scanning endpoints
- Repository management
- Scan results retrieval
- Webhook configuration
- User and team management

For more information, see the [TruffleHog documentation](https://docs.trufflesecurity.com/).
