# TruffleHog June 2025 Release Features

## 🚀 Improve Response

### New Analyzers (11 Total)
**Availability**: Enterprise Edition and Open Source

1. **Plaid** - Financial services API analyzer
2. **Netlify** - Deployment platform analyzer
3. **Fastly Personal Token** - CDN service analyzer
4. **Monday** - Project management analyzer
5. **Datadog Token** - Monitoring platform analyzer
6. **Ngrok** - Tunneling service analyzer
7. **Mux** - Video streaming analyzer
8. **PostHog** - Product analytics analyzer
9. **Dropbox** - File storage analyzer
10. **Databricks** - Data platform analyzer
11. **Jira** - Issue tracking analyzer

**Benefits**: Expanded analyzer coverage enables faster, broader response to potential exposures

**Configuration**:
```yaml
analyzers:
  plaid: true
  netlify: true
  fastly_personal_token: true
  monday: true
  datadog_token: true
  ngrok: true
  mux: true
  posthog: true
  dropbox: true
  databricks: true
  jira: true
```

---

## 🔍 Find More Secrets

### New Detectors and Key Types
**Total**: 806 active detectors supporting 848 different key types

**Availability**: Enterprise Edition and Open Source

#### New Detectors:

1. **Salesforce OAuth2**
```yaml
detectors:
  - name: salesforce_oauth2
    enabled: true
    patterns:
      - "00D[a-zA-Z0-9]{15}!.*"
    verification: true
```

2. **Lokalise**
```yaml
detectors:
  - name: lokalise
    enabled: true
    patterns:
      - "lokalise_[a-zA-Z0-9]{64}"
    verification: true
```

3. **Bannerbear**
```yaml
detectors:
  - name: bannerbear
    enabled: true
    patterns:
      - "bb_[a-z]{2}_[a-zA-Z0-9]{32}"
    verification: true
```

4. **Langsmith**
```yaml
detectors:
  - name: langsmith
    enabled: true
    patterns:
      - "ls__[a-zA-Z0-9]{48}"
    verification: true
```

**Benefits**: Expanded detection capabilities improve coverage and reduce credential exposure risk

---

## 🛠️ Ease Administration

### 1. Manual Jira Notification Control
**Availability**: Generally Available in Enterprise Edition

**Features**:
- Manually resend Jira notifications from Remediations tab
- Sync Jira notifications on demand
- Reopen closed Jira tickets if secret remains LIVE

**Limitations**:
- No bulk actions
- No auto-syncing
- Jira Cloud with autoclose not included

**Usage**:
```bash
# From Secret Details → Remediations tab
# Click "Resend to Jira" or "Sync with Jira"
```

### 2. Config Files in Open Source
**Availability**: Generally Available in Open Source

**Benefits**:
- Run multiple scans concurrently
- Troubleshoot Enterprise integration issues
- Better configuration management

**Example Config**:
```yaml
# .trufflehog/config.yml
version: "2.0"

detectors:
  - salesforce_oauth2
  - lokalise
  - bannerbear
  - langsmith

analyzers:
  - plaid
  - netlify
  - datadog_token

scan:
  paths:
    - "/path/to/repo1"
    - "/path/to/repo2"
  concurrent: true
  max_workers: 4
```

**Usage**:
```bash
# Run with config file
trufflehog git file://. --config .trufflehog/config.yml

# Run multiple concurrent scans
trufflehog git file://repo1 --config config1.yml &
trufflehog git file://repo2 --config config2.yml &
```

### 3. Always-On Integration Filters
**Availability**: Generally Available in Enterprise Edition

**Change**: Filters and sorting now always visible (previously only for 25+ integrations)

**Benefits**: Consistent filtering and sorting regardless of integration count

### 4. Live Secret Notifications
**Availability**: Generally Available in Enterprise Edition

**Feature**: Notifications when secrets transition from not-live to live

**Benefits**: Real-time awareness of newly active secrets for faster response

**Configuration**:
```yaml
notifications:
  live_secret_transitions:
    enabled: true
    channels:
      - slack
      - email
      - jira
```

### 5. Real-Time Scanning Animation
**Availability**: Generally Available in Enterprise Edition

**Feature**: Animated sync icon shows when real-time scanning is active

**Benefits**: Clear visual indicator of active scanning

### 6. Smarter Scans for File Systems
**Availability**: Limited Availability in Enterprise Edition

**Features**:
- Detailed scan progress visibility
- Asset-level scan status
- Active scan monitoring
- Last scan runtime information
- Early scan starts (before full asset enumeration)
- Automatic parallelization (multiple scanners per integration)
- Resume capability (continues from last checkpoint)
- Already available in GitHub scans

**Benefits**:
- Clearer visibility into scan progress
- Faster scan results
- Greater reliability
- Reduced friction
- Improved coverage

**Configuration**:
```yaml
scans:
  file_system:
    smart_scans: true
    parallelization: true
    max_workers: 8
    resume_on_failure: true
    early_start: true
```

**Parallelization**:
```yaml
# Multiple scanners work on different paths simultaneously
scans:
  paths:
    - "/app/backend"
    - "/app/frontend"
    - "/app/config"
    - "/app/scripts"
  # 4 parallel workers (one per path)
  parallel_workers: 4
```

### 7. Active Scan Sorting
**Availability**: Limited Availability in Enterprise Edition

**Location**: Dashboard → Active scans (expand scan)

**Sort Options**:
- Repository name
- Status
- Duration
- Verified secrets count

**Benefits**: Easily find and prioritize active scans

**Usage**:
```bash
# View active scans sorted by verified secrets
# Navigate to Dashboard → Active scans
# Click column header to sort
```

---

## 📊 Feature Summary

| Feature | Type | Availability |
|---------|------|--------------|
| 11 New Analyzers | Improve Response | OSS + Enterprise |
| 4 New Detectors | Find More Secrets | OSS + Enterprise |
| Manual Jira Control | Administration | Enterprise (GA) |
| Config Files in OSS | Administration | OSS (GA) |
| Always-On Filters | Administration | Enterprise (GA) |
| Live Secret Notifications | Administration | Enterprise (GA) |
| Scan Animation | Administration | Enterprise (GA) |
| Smarter Scans | Administration | Enterprise (LA) |
| Active Scan Sorting | Administration | Enterprise (LA) |

**Legend**: GA = Generally Available, LA = Limited Availability

---

## 🎯 Configuration for Starbucks App

### Complete June 2025 Configuration

```yaml
# .trufflehog/config.yml
version: "2.0"

# New Detectors (June 2025)
detectors:
  salesforce_oauth2:
    enabled: true
    verification: true
  
  lokalise:
    enabled: true
    verification: true
  
  bannerbear:
    enabled: true
    verification: true
  
  langsmith:
    enabled: true
    verification: true

# New Analyzers (June 2025)
analyzers:
  plaid: true
  netlify: true
  fastly_personal_token: true
  monday: true
  datadog_token: true
  ngrok: true
  mux: true
  posthog: true
  dropbox: true
  databricks: true
  jira: true

# Notifications
notifications:
  live_secret_transitions:
    enabled: true
    channels:
      - slack
      - email

# Jira Integration
jira:
  manual_resend: true
  manual_sync: true
  autoclose: true

# Smarter Scans
scans:
  file_system:
    smart_scans: true
    parallelization: true
    max_workers: 4
    resume_on_failure: true
    early_start: true
  
  paths:
    - "/"
    - "/frontend"
    - "/graph"
    - "/.circleci"
  
  sorting:
    enabled: true
    default: "verified_secrets_count"
```

---

## 💻 Usage Examples

### 1. Scan with New Detectors
```bash
# Scan with June 2025 detectors
trufflehog git file://. \
  --detectors salesforce_oauth2,lokalise,bannerbear,langsmith \
  --only-verified \
  --json

# Expected output
{
  "DetectorName": "Salesforce OAuth2",
  "Verified": true,
  "Raw": "00D...!...",
  "ExtraData": {
    "instance_url": "https://your-org.salesforce.com",
    "user_id": "005..."
  }
}
```

### 2. Use Config File (OSS)
```bash
# Create config
cat > .trufflehog/starbucks-config.yml <<EOF
version: "2.0"
detectors:
  - salesforce_oauth2
  - lokalise
analyzers:
  - netlify
  - datadog_token
EOF

# Run with config
trufflehog git file://. --config .trufflehog/starbucks-config.yml
```

### 3. Concurrent Scans (OSS)
```bash
# Scan multiple repos concurrently
trufflehog git file://. --config backend-config.yml &
trufflehog filesystem ./frontend --config frontend-config.yml &
wait
```

### 4. Monitor Active Scans
```bash
# View in TruffleHog Enterprise UI:
# Dashboard → Active scans → Expand scan
# Sort by: verified_secrets_count (descending)
```

---

## 🔄 Update CircleCI Configuration

```yaml
# .circleci/config.yml
jobs:
  scan-secrets-june-2025:
    docker:
      - image: trufflesecurity/trufflehog:latest
    steps:
      - checkout
      - run:
          name: "Scan with June 2025 features"
          command: |
            # Use config file
            trufflehog git file://. \
              --config .trufflehog/config.yml \
              --only-verified \
              --json > results.json
            
            # Check for new detector types
            if jq -e '.[] | select(.DetectorName | 
              test("Salesforce|Lokalise|Bannerbear|Langsmith"))' results.json; then
              echo "⚠️ New platform secrets detected!"
              exit 1
            fi
```

---

## 🔗 Update GitHub Actions

```yaml
# .github/workflows/trufflehog-june-2025.yml
name: TruffleHog June 2025 Scan

on: [push, pull_request]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - name: TruffleHog with June 2025 Features
        run: |
          trufflehog git file://. \
            --config .trufflehog/config.yml \
            --detectors salesforce_oauth2,lokalise,bannerbear,langsmith \
            --only-verified
```

---

## 📈 Detector Statistics

**June 2025 Totals**:
- **806 active detectors**
- **848 different key types**
- **11 new analyzers**
- **4 new detectors**

**Coverage Increase**:
- May 2025: ~795 detectors
- June 2025: 806 detectors
- Growth: +11 detectors (+1.4%)

---

## 🎓 Best Practices

### 1. Enable All New Analyzers
```yaml
analyzers:
  plaid: true
  netlify: true
  datadog_token: true
  jira: true
  # ... enable all 11
```

### 2. Use Config Files for Consistency
- Create environment-specific configs
- Version control your configs
- Share configs across team

### 3. Leverage Smarter Scans
- Enable parallelization for faster scans
- Use resume capability for reliability
- Monitor active scans dashboard

### 4. Set Up Live Secret Notifications
- Get alerted when secrets become active
- Respond faster to new threats
- Reduce exposure window

### 5. Use Jira Integration Effectively
- Manually resend when needed
- Sync to reopen closed tickets
- Track remediation progress

---

## 📝 Migration Guide

### From May 2025 to June 2025

**Step 1: Update Detectors**
```yaml
# Add to existing config
detectors:
  # May 2025
  - deepseek
  - xai
  - harness
  # June 2025 (new)
  - salesforce_oauth2
  - lokalise
  - bannerbear
  - langsmith
```

**Step 2: Enable Analyzers**
```yaml
analyzers:
  # June 2025 (all new)
  - plaid
  - netlify
  - fastly_personal_token
  - monday
  - datadog_token
  - ngrok
  - mux
  - posthog
  - dropbox
  - databricks
  - jira
```

**Step 3: Configure Notifications**
```yaml
notifications:
  live_secret_transitions:
    enabled: true
```

**Step 4: Enable Smarter Scans**
```yaml
scans:
  file_system:
    smart_scans: true
    parallelization: true
```

---

## 📞 Support

- **Documentation**: https://docs.trufflesecurity.com/
- **GitHub**: https://github.com/trufflesecurity/trufflehog
- **Community**: https://community.trufflesecurity.com/

---

**Last Updated**: June 2025
**TruffleHog Version**: Latest (June 2025 release)
**Total Detectors**: 806
**Total Key Types**: 848
