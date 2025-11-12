# TruffleHog Latest Release - Customer Facing Features

## Overview
Latest customer-facing features in TruffleHog, organized by capability area.

---

## 🚀 Improve Response
Features that help teams act faster and more effectively when secrets are found.

### 1. CSV Exports with Dashboard Links
**Availability**: Enterprise Edition

**Feature**: CSV exports now include direct links to each secret in the TruffleHog dashboard

**Benefits**:
- Instantly open and investigate live secrets
- No manual searching required
- Saves time and accelerates response

**Example CSV Output**:
```csv
Secret ID,Detector,Status,Verified,Dashboard Link
abc123,AWS,Live,true,https://app.trufflesecurity.com/secrets/abc123
def456,GitHub,Live,true,https://app.trufflesecurity.com/secrets/def456
```

**Usage**:
```bash
# Export secrets with dashboard links
# Navigate to Secrets page → Export → Download CSV
# CSV will include "Dashboard Link" column
```

### 2. Source Name in Notification Emails
**Availability**: Enterprise Edition

**Feature**: Notification emails now include the source name where the secret was detected

**Benefits**:
- Immediate context about leak location
- Faster triage
- Improved routing
- Reduced time to resolution

**Email Example**:
```
Subject: [TruffleHog] Live Secret Detected

Secret Type: AWS Access Key
Status: Verified Live
Source: github.com/jkzilla/starbucks
File: server.go
Line: 42

View in Dashboard: https://app.trufflesecurity.com/secrets/abc123
```

**Configuration**:
```yaml
notifications:
  email:
    include_source_name: true
    include_file_path: true
    include_dashboard_link: true
```

---

## 🔍 Find More Secrets
Features that expand detection coverage and secret types.

### New Detectors and Key Types
**Total**: 810 active detectors supporting 854 different key types

**Availability**: Enterprise Edition and Open Source

#### 1. Webex Bot Detector
```yaml
detectors:
  - name: webexbot
    enabled: true
    patterns:
      - "Bearer [A-Za-z0-9+/]{100,}={0,2}"
      - "webex_bot_[a-zA-Z0-9_-]{40,}"
    verification: true
    verify_endpoint: "https://webexapis.com/v1/people/me"
```

**Usage**:
```bash
# Scan for Webex bot tokens
trufflehog git file://. --detectors webexbot --only-verified
```

#### 2. CircleCI v2 Detector
```yaml
detectors:
  - name: circleci_v2
    enabled: true
    patterns:
      - "CIRCLE_TOKEN=[a-f0-9]{40}"
      - "circleci_[a-f0-9]{40}"
    verification: true
    verify_endpoint: "https://circleci.com/api/v2/me"
```

**Usage**:
```bash
# Scan for CircleCI v2 tokens
trufflehog git file://. --detectors circleci_v2 --only-verified
```

#### 3. Salesforce OAuth2 Detector (Enhanced)
```yaml
detectors:
  - name: salesforce_oauth2
    enabled: true
    patterns:
      - "00D[a-zA-Z0-9]{15}!.*"
      - "oauth_token=[a-zA-Z0-9!.]{100,}"
    verification: true
    verify_endpoint: "https://login.salesforce.com/services/oauth2/userinfo"
```

**Benefits**: Expanded detection capabilities improve coverage and reduce credential exposure risk

**Growth**: +4 detectors, +6 key types from previous release

---

## 🛠️ Ease Administration
Features that simplify ongoing management of TruffleHog.

### 1. Bulk Queue Notifications via API
**Availability**: Enterprise Edition

**Feature**: Queue notifications via external API for multiple secrets

**Limits**: 100 secrets max per request

**Benefits**:
- Prompt alerts for newly confirmed live secrets
- Batch processing for efficiency
- Maintain strong security coverage

**API Example**:
```bash
# Bulk queue notifications
curl -X POST https://api.trufflesecurity.com/v1/notifications/queue \
  -H "Authorization: Bearer $TRUFFLEHOG_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "notifier_id": "notifier_123",
    "secret_ids": [
      "secret_abc123",
      "secret_def456",
      "secret_ghi789"
    ]
  }'
```

**Response**:
```json
{
  "queued": 3,
  "failed": 0,
  "notifier_id": "notifier_123",
  "status": "success"
}
```

**Python Example**:
```python
import requests

def bulk_queue_notifications(notifier_id, secret_ids):
    url = "https://api.trufflesecurity.com/v1/notifications/queue"
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    
    # Process in batches of 100
    for i in range(0, len(secret_ids), 100):
        batch = secret_ids[i:i+100]
        payload = {
            "notifier_id": notifier_id,
            "secret_ids": batch
        }
        response = requests.post(url, headers=headers, json=payload)
        print(f"Queued {len(batch)} notifications")
```

### 2. IdP Role Mapping
**Availability**: Enterprise Edition

**Feature**: Automatically sync users from Identity Provider (IdP)

**Benefits**:
- Simplifies user setup
- Eliminates manual account management
- Automatically up-to-date user lists

**Supported IdPs**:
- Okta
- Azure AD
- Google Workspace
- OneLogin
- Auth0

**Configuration**:
```yaml
idp:
  enabled: true
  provider: "okta"
  sync_interval: "1h"
  
  role_mapping:
    - idp_group: "security-team"
      trufflehog_role: "admin"
    
    - idp_group: "developers"
      trufflehog_role: "viewer"
    
    - idp_group: "devops"
      trufflehog_role: "editor"
  
  auto_provision: true
  auto_deprovision: true
```

**Setup**:
```bash
# Configure IdP integration
curl -X POST https://api.trufflesecurity.com/v1/idp/configure \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{
    "provider": "okta",
    "domain": "your-org.okta.com",
    "client_id": "...",
    "client_secret": "...",
    "role_mapping": [...]
  }'
```

### 3. Share Link Enhancements
**Availability**: Enterprise Edition

**Features**:
- Authorization improvements
- New "share viewer" role
- Better visibility into access
- Quick permission updates
- Auditing support

**Benefits**:
- Secure collaboration
- Dedicated share viewer role
- Enhanced authorization controls
- Better access visibility
- Audit trail

**Share Viewer Role**:
```yaml
roles:
  share_viewer:
    permissions:
      - view_shared_secrets
      - view_shared_dashboards
    restrictions:
      - cannot_edit
      - cannot_delete
      - cannot_share_further
```

**Usage**:
```bash
# Create share link with viewer role
curl -X POST https://api.trufflesecurity.com/v1/shares \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "secret_id": "abc123",
    "role": "share_viewer",
    "expires_in": "7d"
  }'

# List share viewers for auditing
curl https://api.trufflesecurity.com/v1/shares/viewers \
  -H "Authorization: Bearer $TOKEN"
```

### 4. Unauthorized Access Page
**Availability**: Enterprise Edition

**Feature**: Dedicated "Unauthorized" page for insufficient permissions

**Benefits**:
- Improves user experience
- Clear communication of access restrictions
- Reduces confusion
- Faster identification of misconfigured permissions

**Before**:
```
Error: undefined
[Blank screen or unclear error]
```

**After**:
```
Unauthorized Access

You don't have permission to access this resource.

Required Permission: secrets:admin
Your Role: viewer

Contact your administrator to request access.
```

### 5. Improved Notification Terminology
**Availability**: Enterprise Edition

**Change**: "Create Notification" → "Queue Notification"

**Benefits**:
- More accurate reflection of action
- Improved clarity
- Reduced learning curve

**UI Updates**:
- Button: "Queue Notification" (was "Create Notification")
- API: `/notifications/queue` (was `/notifications/create`)
- Documentation: Updated terminology throughout

### 6. Enhanced API Documentation
**Availability**: Enterprise Edition

**Feature**: Restructured API docs with better grouping

**Benefits**:
- Clearer navigation
- Better understanding of API implementation
- Reduced learning curve

**New Grouping**:
```
API Documentation
├── Authentication
│   ├── OAuth2
│   └── API Keys
├── Secrets
│   ├── List Secrets
│   ├── Get Secret
│   └── Update Secret
├── Notifications
│   ├── Queue Notification
│   ├── Bulk Queue
│   └── List Notifications
├── Integrations
│   ├── Sources
│   └── Notifiers
└── Administration
    ├── Users
    ├── Roles
    └── IdP Sync
```

### 7. Alphabetized Integration Filter
**Availability**: Enterprise Edition

**Feature**: Integration filter dropdown sorted alphabetically

**Benefits**:
- Faster integration location
- Easier to find integrations
- Reduced learning curve

**Before**: Random order
**After**: A-Z sorted

```
Integrations Filter:
☐ AWS
☐ Azure
☐ CircleCI
☐ Datadog
☐ GitHub
☐ GitLab
☐ Jira
☐ Netlify
☐ Slack
```

---

## 📊 Feature Summary

| Feature | Category | Availability |
|---------|----------|--------------|
| CSV Dashboard Links | Improve Response | Enterprise |
| Source in Email | Improve Response | Enterprise |
| Webex Bot Detector | Find More Secrets | OSS + Enterprise |
| CircleCI v2 Detector | Find More Secrets | OSS + Enterprise |
| Salesforce OAuth2 | Find More Secrets | OSS + Enterprise |
| Bulk Queue API | Administration | Enterprise |
| IdP Role Mapping | Administration | Enterprise |
| Share Link Enhancements | Administration | Enterprise |
| Unauthorized Page | Administration | Enterprise |
| Queue Terminology | Administration | Enterprise |
| Enhanced API Docs | Administration | Enterprise |
| Alphabetized Filters | Administration | Enterprise |

---

## 🎯 Configuration for Starbucks App

### Complete Latest Release Configuration

```yaml
# .trufflehog/config.yml
version: "2.0"

# New Detectors
detectors:
  webexbot:
    enabled: true
    verification: true
  
  circleci_v2:
    enabled: true
    verification: true
  
  salesforce_oauth2:
    enabled: true
    verification: true

# Notifications
notifications:
  email:
    include_source_name: true
    include_file_path: true
    include_dashboard_link: true
  
  bulk_queue:
    enabled: true
    max_per_request: 100

# CSV Exports
exports:
  csv:
    include_dashboard_links: true
    include_source_name: true

# IdP Integration
idp:
  enabled: true
  provider: "okta"
  sync_interval: "1h"
  role_mapping:
    - idp_group: "security-team"
      trufflehog_role: "admin"

# Share Links
sharing:
  viewer_role_enabled: true
  authorization_enhanced: true
  audit_logging: true

# UI Preferences
ui:
  terminology:
    use_queue_notification: true
  filters:
    alphabetize_integrations: true
  
  unauthorized_page:
    enabled: true
    show_required_permission: true
```

---

## 💻 Usage Examples

### 1. Scan with New Detectors
```bash
# Scan for Webex, CircleCI, Salesforce
trufflehog git file://. \
  --detectors webexbot,circleci_v2,salesforce_oauth2 \
  --only-verified \
  --json

# Expected output
{
  "DetectorName": "CircleCI v2",
  "Verified": true,
  "Raw": "circleci_abc123...",
  "ExtraData": {
    "username": "jkzilla",
    "project": "starbucks"
  }
}
```

### 2. Bulk Queue Notifications
```bash
# Queue notifications for multiple secrets
curl -X POST https://api.trufflesecurity.com/v1/notifications/queue \
  -H "Authorization: Bearer $API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "notifier_id": "slack_notifier_123",
    "secret_ids": ["secret_1", "secret_2", "secret_3"]
  }'
```

### 3. Export CSV with Dashboard Links
```bash
# From TruffleHog UI:
# 1. Navigate to Secrets page
# 2. Click "Export"
# 3. Select "CSV with Dashboard Links"
# 4. Download file

# CSV will include:
# Secret ID, Detector, Status, Source, Dashboard Link
```

### 4. Configure IdP Sync
```bash
# Set up Okta integration
curl -X POST https://api.trufflesecurity.com/v1/idp/configure \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{
    "provider": "okta",
    "domain": "starbucks.okta.com",
    "role_mapping": [
      {"idp_group": "security", "trufflehog_role": "admin"}
    ]
  }'
```

---

## 🔄 Update CircleCI Configuration

```yaml
# .circleci/config.yml
jobs:
  scan-secrets-latest:
    docker:
      - image: trufflesecurity/trufflehog:latest
    steps:
      - checkout
      - run:
          name: "Scan with latest detectors"
          command: |
            trufflehog git file://. \
              --detectors webexbot,circleci_v2,salesforce_oauth2 \
              --only-verified \
              --json > results.json
            
            # Queue notifications if secrets found
            if [ -s results.json ]; then
              SECRET_IDS=$(jq -r '.[].SecretID' results.json)
              curl -X POST https://api.trufflesecurity.com/v1/notifications/queue \
                -H "Authorization: Bearer $TRUFFLEHOG_API_KEY" \
                -d "{\"notifier_id\": \"$NOTIFIER_ID\", \"secret_ids\": $SECRET_IDS}"
            fi
```

---

## 📈 Detector Statistics

**Latest Release Totals**:
- **810 active detectors** (+4 from June 2025)
- **854 different key types** (+6 from June 2025)
- **3 new detectors**: Webex Bot, CircleCI v2, Salesforce OAuth2 (enhanced)

**Growth Timeline**:
- May 2025: ~795 detectors
- June 2025: 806 detectors
- Latest: 810 detectors
- Total Growth: +15 detectors (+1.9%)

---

## 🎓 Best Practices

### 1. Enable Dashboard Links in Exports
```yaml
exports:
  csv:
    include_dashboard_links: true
```

### 2. Use Bulk Queue API Efficiently
```python
# Process in batches of 100
def queue_notifications_batch(secret_ids):
    for i in range(0, len(secret_ids), 100):
        batch = secret_ids[i:i+100]
        queue_notifications(notifier_id, batch)
```

### 3. Configure IdP for Automated User Management
```yaml
idp:
  auto_provision: true
  auto_deprovision: true
  sync_interval: "1h"
```

### 4. Use Share Viewer Role for External Collaboration
```bash
# Create time-limited share links
curl -X POST /v1/shares \
  -d '{"role": "share_viewer", "expires_in": "7d"}'
```

### 5. Monitor Unauthorized Access Attempts
```bash
# Review unauthorized access logs
curl https://api.trufflesecurity.com/v1/audit/unauthorized \
  -H "Authorization: Bearer $TOKEN"
```

---

## 📝 Migration Notes

### Terminology Updates
- Update scripts: `create_notification` → `queue_notification`
- Update documentation references
- Update API calls to new endpoints

### API Changes
```bash
# Old (deprecated)
POST /v1/notifications/create

# New (recommended)
POST /v1/notifications/queue
```

### CSV Export Format
```csv
# New columns added:
Dashboard Link, Source Name
```

---

## 📞 Support

- **Documentation**: https://docs.trufflesecurity.com/
- **API Docs**: https://api.trufflesecurity.com/docs (now with better grouping!)
- **GitHub**: https://github.com/trufflesecurity/trufflehog
- **Community**: https://community.trufflesecurity.com/

---

**Last Updated**: Latest Release
**TruffleHog Version**: Latest
**Total Detectors**: 810
**Total Key Types**: 854
