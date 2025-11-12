# GitHub Advanced Security Setup Guide

This guide walks you through setting up GitHub Advanced Security (GHAS) for the Starbucks Coffee Shop application.

## Prerequisites

- GitHub Enterprise Cloud or GitHub Enterprise Server with Advanced Security license
- Repository admin access
- GitHub Actions enabled

## Step 1: Enable GitHub Advanced Security

### For Organization Owners

1. Go to your organization settings
2. Navigate to **Code security and analysis**
3. Enable **GitHub Advanced Security** for all repositories or specific ones

### For Repository Admins

1. Go to repository **Settings**
2. Navigate to **Code security and analysis**
3. Enable the following features:

```
✅ Dependency graph
✅ Dependabot alerts
✅ Dependabot security updates
✅ Grouped security updates
✅ Code scanning
✅ Secret scanning
✅ Push protection
```

## Step 2: Configure Code Scanning (CodeQL)

### Automatic Setup (Recommended)

1. Go to **Security** tab → **Code scanning**
2. Click **Set up code scanning**
3. Choose **Default** setup
4. Select languages: **Go**, **JavaScript/TypeScript**
5. Click **Enable CodeQL**

### Manual Setup (Already Configured)

The repository includes `.github/workflows/codeql.yml` with:
- Multi-language support (Go, JavaScript/TypeScript)
- Extended security queries
- Scheduled daily scans
- Custom path exclusions

**Verify it's working:**
```bash
# Check workflow runs
gh run list --workflow=codeql.yml

# View latest run
gh run view --workflow=codeql.yml
```

## Step 3: Configure Secret Scanning

### Enable Push Protection

1. Go to **Settings** → **Code security and analysis**
2. Under **Secret scanning**, enable:
   - ✅ **Secret scanning**
   - ✅ **Push protection**
   - ✅ **Validity checks** (verifies if secrets are active)

### Configure Custom Patterns

1. Go to **Settings** → **Code security and analysis** → **Secret scanning**
2. Click **New pattern**
3. Add custom patterns for:
   - Internal API keys
   - Database connection strings
   - Custom tokens

**Example custom pattern:**
```regex
# Internal API Key Pattern
starbucks_api_key_[a-zA-Z0-9]{32}
```

### Verify Secret Scanning

```bash
# Check for any detected secrets
gh api /repos/jkzilla/starbucks/secret-scanning/alerts

# View secret scanning settings
gh api /repos/jkzilla/starbucks/secret-scanning
```

## Step 4: Configure Dependabot

### Enable Dependabot (Already Configured)

The repository includes `.github/dependabot.yml` with:
- Go module updates (weekly)
- NPM package updates (weekly)
- GitHub Actions updates (weekly)
- Grouped updates for easier review
- Security-first update strategy

### Configure Dependabot Secrets

If you need private registry access:

1. Go to **Settings** → **Secrets and variables** → **Dependabot**
2. Add secrets:
   - `NPM_TOKEN` (if using private npm registry)
   - `GO_PRIVATE_KEY` (if using private Go modules)

### Review Dependabot Alerts

```bash
# List Dependabot alerts
gh api /repos/jkzilla/starbucks/dependabot/alerts

# View specific alert
gh api /repos/jkzilla/starbucks/dependabot/alerts/1
```

## Step 5: Configure Dependency Review

### Enable in Pull Requests

The workflow `.github/workflows/dependency-review.yml` automatically:
- Scans PRs for vulnerable dependencies
- Checks license compliance
- Blocks PRs with critical/high vulnerabilities
- Comments on PRs with findings

**Configuration:**
- Fails on: `moderate` severity or higher
- Denied licenses: GPL-3.0, AGPL-3.0
- Allowed licenses: MIT, Apache-2.0, BSD-3-Clause, ISC

## Step 6: Configure Branch Protection

### Recommended Settings

1. Go to **Settings** → **Branches**
2. Add rule for `main` branch:

```yaml
Branch name pattern: main

Protect matching branches:
✅ Require a pull request before merging
  ✅ Require approvals: 1
  ✅ Dismiss stale pull request approvals when new commits are pushed
  ✅ Require review from Code Owners

✅ Require status checks to pass before merging
  ✅ Require branches to be up to date before merging
  Required status checks:
    - CodeQL / Analyze (go)
    - CodeQL / Analyze (javascript-typescript)
    - Dependency Review
    - TruffleHog Secret Scan
    - Gitleaks Secret Scan

✅ Require conversation resolution before merging
✅ Require signed commits
✅ Require linear history
✅ Include administrators

Do not allow bypassing the above settings:
✅ Do not allow bypassing the above settings
```

## Step 7: Configure Security Policies

### Create CODEOWNERS File

Create `.github/CODEOWNERS`:
```
# Global owners
* @jkzilla

# Security-sensitive files
/SECURITY.md @jkzilla
/.github/workflows/ @jkzilla
/graph/schema.resolvers.go @jkzilla
/server.go @jkzilla

# Frontend security
/frontend/src/graphql/ @jkzilla
```

### Configure Security Advisories

1. Go to **Security** tab → **Advisories**
2. Enable **Private vulnerability reporting**
3. Set up security contacts

## Step 8: Set Up Notifications

### Configure Alert Notifications

1. Go to **Settings** → **Notifications**
2. Configure:
   - Email notifications for security alerts
   - Slack/Teams integration (if available)
   - Custom webhooks

### GitHub CLI Setup

```bash
# Watch repository for security alerts
gh repo set-default jkzilla/starbucks

# Get notified of security alerts
gh api /repos/jkzilla/starbucks/vulnerability-alerts
```

## Step 9: Verify Setup

### Run Security Checks

```bash
# Trigger CodeQL scan
gh workflow run codeql.yml

# Trigger secret scanning
gh workflow run secret-scanning.yml

# Trigger OSSF Scorecard
gh workflow run security-scorecard.yml

# Check all security features
gh api /repos/jkzilla/starbucks | jq '.security_and_analysis'
```

### Expected Output

```json
{
  "advanced_security": {
    "status": "enabled"
  },
  "secret_scanning": {
    "status": "enabled"
  },
  "secret_scanning_push_protection": {
    "status": "enabled"
  },
  "dependabot_security_updates": {
    "status": "enabled"
  }
}
```

## Step 10: Monitor and Maintain

### Daily Tasks

- Review new security alerts
- Check CodeQL findings
- Monitor secret scanning alerts

### Weekly Tasks

- Review Dependabot PRs
- Check OSSF Scorecard results
- Update security documentation

### Monthly Tasks

- Review security policy
- Audit access controls
- Update security training
- Review incident response procedures

## Security Dashboard

Access your security overview:

```
https://github.com/jkzilla/starbucks/security
```

This shows:
- Code scanning alerts
- Secret scanning alerts
- Dependabot alerts
- Security advisories
- Security policy

## Troubleshooting

### CodeQL Not Running

```bash
# Check workflow file syntax
gh workflow view codeql.yml

# Check recent runs
gh run list --workflow=codeql.yml --limit 5

# View logs
gh run view <run-id> --log
```

### Secret Scanning Not Detecting Secrets

1. Verify push protection is enabled
2. Check custom patterns are correct
3. Ensure secrets match known patterns
4. Review `.gitignore` for excluded files

### Dependabot Not Creating PRs

1. Check Dependabot configuration syntax
2. Verify package ecosystems are correct
3. Check for rate limiting
4. Review Dependabot logs in Security tab

## Best Practices

### 1. Regular Reviews

- Review security alerts daily
- Triage and prioritize vulnerabilities
- Document remediation steps

### 2. Automation

- Auto-merge low-risk Dependabot PRs
- Use GitHub Actions for security checks
- Implement automated testing

### 3. Training

- Train team on security best practices
- Conduct security awareness sessions
- Share security updates

### 4. Documentation

- Keep security policy updated
- Document security procedures
- Maintain runbooks

## Advanced Configuration

### Custom CodeQL Queries

Create `.github/codeql/custom-queries.ql`:

```ql
/**
 * @name Hardcoded credentials
 * @description Detects hardcoded credentials in code
 * @kind problem
 * @id go/hardcoded-credentials
 * @problem.severity error
 */

import go

from StringLit s
where s.getValue().regexpMatch("(?i)(password|secret|key|token)\\s*=\\s*['\"][^'\"]+['\"]")
select s, "Potential hardcoded credential detected"
```

### Custom Secret Patterns

Add to repository settings or use `.github/secret_scanning.yml`:

```yaml
patterns:
  - name: "Starbucks API Key"
    regex: "starbucks_[a-zA-Z0-9]{32}"
    
  - name: "Internal Database URL"
    regex: "postgres://[^:]+:[^@]+@[^/]+/starbucks"
```

## Resources

- [GitHub Advanced Security Documentation](https://docs.github.com/en/code-security)
- [CodeQL Documentation](https://codeql.github.com/docs/)
- [Dependabot Documentation](https://docs.github.com/en/code-security/dependabot)
- [Secret Scanning Documentation](https://docs.github.com/en/code-security/secret-scanning)
- [OSSF Scorecard](https://github.com/ossf/scorecard)

## Support

For issues with GitHub Advanced Security:
- GitHub Support: https://support.github.com
- Community Forum: https://github.community
- Documentation: https://docs.github.com

---

**Setup completed!** Your repository is now protected by GitHub Advanced Security. 🔒
