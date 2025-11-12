# TruffleHog May 2025 Release Features

## Overview
Features released throughout May 2025 in TruffleHog, organized by capability area.

---

## 🚀 Improve Response
Features that help teams act faster and more effectively when secrets are found.

### 1. Analyze Data Filter
**Feature**: View all identified Secrets that contain Analyze data
**How to Use**:
- Open the Analyze filter option
- Select "All analyzed" as the Analyzer Type
- View permissions on resources to aid prioritization

**Availability**: Limited Availability in Enterprise Edition with Analyze add-on

**Configuration**:
```yaml
filters:
  analyze:
    type: "all_analyzed"
    show_permissions: true
```

### 2. Jira Ticket Integration
**Feature**: Jira Ticket IDs added to Secret details
**Benefits**: Promotes team collaboration and communication

**Availability**: Available in Enterprise Edition

**Configuration**:
```yaml
integrations:
  jira:
    enabled: true
    show_ticket_id: true
    link_to_secrets: true
```

### 3. Secret ID Search
**Feature**: Search by Secret ID
**Benefits**: Simplifies team work and communication

**Availability**: Available in Enterprise Edition

**Usage**:
```bash
# Search by Secret ID
trufflehog search --secret-id "SECRET_ID_HERE"
```

---

## 🔍 Find More Secrets
Features that expand detection coverage and secret types.

### New Detectors (7 Total)

#### 1. Deepseek Detector
**Type**: AI/ML Platform
**Availability**: Open Source and Enterprise Edition

```yaml
detectors:
  - name: deepseek
    enabled: true
    patterns:
      - "deepseek_api_key_[a-zA-Z0-9]{32}"
    verification: true
```

#### 2. xAI Detector
**Type**: AI Platform (Grok/X.AI)
**Availability**: Open Source and Enterprise Edition

```yaml
detectors:
  - name: xai
    enabled: true
    patterns:
      - "xai_[a-zA-Z0-9]{40}"
      - "grok_api_[a-zA-Z0-9]{32}"
    verification: true
```

#### 3. Harness Detector
**Type**: CI/CD Platform
**Availability**: Open Source and Enterprise Edition

```yaml
detectors:
  - name: harness
    enabled: true
    patterns:
      - "harness_api_key_[a-zA-Z0-9]{32}"
      - "sat\\.[a-zA-Z0-9]{22}\\.[a-zA-Z0-9]{40}"
    verification: true
```

### New Key Types (8 Total)
- Deepseek API Keys
- xAI API Keys
- Grok API Keys
- Harness API Keys
- Harness Service Account Tokens
- Additional platform-specific keys (3 more)

**Benefits**:
- Broader range of exposed secrets detection
- Improved coverage
- Reduced risks

---

## 🛠️ Ease Administration
Features that simplify ongoing management of TruffleHog.

### 1. On-Premise Scanner Management
**Feature**: Link to scanner management when setting up on-premise sources
**Availability**: Enterprise Edition

### 2. Enhanced Notifications Drawer
**Feature**: View most recently created notifications for a particular notifier
**Benefits**: Easier notification consumption

**Availability**: Enterprise Edition

**Configuration**:
```yaml
notifications:
  drawer:
    show_recent: true
    sort_by: "created_at"
    limit: 50
```

### 3. SAML SSO Error Handling
**Feature**: Unauthorized page when SAML SSO login fails
**Benefits**: Reduces user confusion and support requests

**Availability**: Enterprise Edition

### 4. Scan Job Reports UI Improvements
**Features**:
- Pagination
- Error modal
- Real-time scanning support
- Deeper visibility
- Smarter control over scan execution

**Availability**: Limited Availability in Enterprise Edition

### 5. Proactive Email Notifications
**Feature**: Emails sent directly to user who enabled the feature
**Availability**: Enterprise Edition (upon request)

**Configuration**:
```yaml
notifications:
  email:
    send_to_enabler: true
    proactive_alerts: true
```

### 6. Secret Locator Terminology
**Change**: Updated verbiage from "Secret ID" to "Secret Locator"
**Benefits**: Better understanding of where secrets were found

**Availability**: Enterprise Edition

### 7. Integration Management at Scale
**Features**:
- Pagination (for 25+ integrations)
- Filtering
- Searching
- No clutter for smaller setups

**Availability**: Enterprise Edition

**Usage**:
```bash
# List integrations with pagination
trufflehog integrations list --page 1 --per-page 25

# Filter integrations
trufflehog integrations list --type source --filter "github"

# Search integrations
trufflehog integrations search "starbucks"
```

### 8. Conditional Filtering
**Feature**: Search + integration type filtering
**Requirements**: 25 or more integrations
**Benefits**: Quickly find and manage specific integrations

**Availability**: Enterprise Edition

**Example**:
```yaml
integrations:
  filters:
    - type: "source"
      search: "github"
    - type: "notification"
      search: "slack"
```

### 9. Location Type Filter Performance
**Feature**: Improved performance of Location Type filter on Secrets page
**Benefits**: Faster and more responsive filtering

**Availability**: Enterprise Edition

---

## 📋 Configuration Examples

### Complete TruffleHog Configuration with May 2025 Features

```yaml
# .trufflehog/config.yml
version: "2.0"

# New Detectors
detectors:
  - name: deepseek
    enabled: true
    verification: true
  
  - name: xai
    enabled: true
    verification: true
  
  - name: harness
    enabled: true
    verification: true

# Analyze Features
analyze:
  enabled: true
  filter_type: "all_analyzed"
  show_permissions: true

# Integrations
integrations:
  jira:
    enabled: true
    show_ticket_id: true
    base_url: "https://your-org.atlassian.net"
  
  notifications:
    drawer:
      show_recent: true
      sort_by: "created_at"
    email:
      send_to_enabler: true
      proactive_alerts: true

# Search
search:
  by_secret_id: true
  by_secret_locator: true

# UI Preferences
ui:
  terminology:
    use_secret_locator: true
  pagination:
    enabled: true
    per_page: 25
  filters:
    location_type:
      performance_mode: true

# Scan Jobs
scan_jobs:
  reports:
    pagination: true
    error_modal: true
    realtime_support: true
```

---

## 🎯 Usage Examples for Starbucks App

### 1. Enable New Detectors
```bash
# Scan with new AI platform detectors
trufflehog git file://. \
  --detectors deepseek,xai,harness \
  --only-verified \
  --json

# Expected output for Deepseek API key
{
  "DetectorName": "Deepseek",
  "Verified": true,
  "Raw": "deepseek_api_key_...",
  "ExtraData": {
    "permissions": ["api:read", "api:write"]
  }
}
```

### 2. Search by Secret ID
```bash
# Search for specific secret
trufflehog search --secret-id "abc123def456"

# Search by secret locator
trufflehog search --locator "github.com/jkzilla/starbucks/server.go:42"
```

### 3. Analyze Permissions
```bash
# Scan and analyze permissions
trufflehog git file://. \
  --analyze \
  --filter "all_analyzed" \
  --json

# View secrets with analyze data
trufflehog secrets list --has-analyze-data
```

### 4. Integration Management
```bash
# List all integrations (paginated)
trufflehog integrations list --page 1

# Filter by type
trufflehog integrations list --type source --filter "github"

# Search integrations
trufflehog integrations search "circleci"
```

---

## 🔗 Integration with Existing Setup

### Update CircleCI Configuration
```yaml
# .circleci/config.yml
jobs:
  scan-secrets:
    docker:
      - image: trufflesecurity/trufflehog:latest
    steps:
      - checkout
      - run:
          name: "Scan with new detectors"
          command: |
            trufflehog git file://. \
              --since-commit main \
              --branch "$CIRCLE_BRANCH" \
              --detectors deepseek,xai,harness \
              --analyze \
              --only-verified \
              --json > results.json
            
            # Check for secrets with analyze data
            if jq -e '.[] | select(.ExtraData.permissions)' results.json; then
              echo "⚠️ Secrets with elevated permissions detected!"
              exit 1
            fi
```

### Update GitHub Actions
```yaml
# .github/workflows/trufflehog-scan.yml
- name: TruffleHog with May 2025 Features
  run: |
    trufflehog git file://. \
      --detectors deepseek,xai,harness \
      --analyze \
      --only-verified \
      --json | tee trufflehog-results.json
    
    # Search by secret locator if needed
    if [ -f secret-locators.txt ]; then
      while read locator; do
        trufflehog search --locator "$locator"
      done < secret-locators.txt
    fi
```

---

## 📊 Feature Availability Matrix

| Feature | Open Source | Enterprise | Enterprise + Analyze |
|---------|-------------|------------|---------------------|
| Deepseek Detector | ✅ | ✅ | ✅ |
| xAI Detector | ✅ | ✅ | ✅ |
| Harness Detector | ✅ | ✅ | ✅ |
| Analyze Data Filter | ❌ | ❌ | ✅ |
| Jira Integration | ❌ | ✅ | ✅ |
| Secret ID Search | ❌ | ✅ | ✅ |
| Enhanced Notifications | ❌ | ✅ | ✅ |
| Integration Management | ❌ | ✅ | ✅ |
| Scan Job Reports | ❌ | ✅ (Limited) | ✅ |

---

## 🎓 Best Practices

### 1. Enable All New Detectors
```yaml
detectors:
  deepseek: true
  xai: true
  harness: true
```

### 2. Use Analyze for Prioritization
- Filter by "all_analyzed"
- Review permissions data
- Prioritize secrets with elevated permissions

### 3. Leverage Jira Integration
- Link secrets to Jira tickets
- Track remediation progress
- Improve team collaboration

### 4. Optimize for Scale
- Enable pagination for 25+ integrations
- Use conditional filtering
- Leverage improved performance

### 5. Search Effectively
- Use Secret ID for quick lookups
- Use Secret Locator for file-based searches
- Combine with filters for precision

---

## 📝 Notes

- All features are available as of May 2025
- Some features require specific TruffleHog editions
- Analyze add-on required for permission analysis
- Contact TruffleHog support for Limited Availability features

---

**Last Updated**: May 2025
**TruffleHog Version**: Latest (May 2025 release)
