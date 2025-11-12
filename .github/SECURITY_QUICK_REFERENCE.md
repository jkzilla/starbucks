# GitHub Advanced Security - Quick Reference

## 🚨 Security Alerts Dashboard

```
https://github.com/jkzilla/starbucks/security
```

## 📋 Quick Commands

### View Security Status
```bash
# Check all security features
gh api /repos/jkzilla/starbucks | jq '.security_and_analysis'

# List CodeQL alerts
gh api /repos/jkzilla/starbucks/code-scanning/alerts

# List secret scanning alerts
gh api /repos/jkzilla/starbucks/secret-scanning/alerts

# List Dependabot alerts
gh api /repos/jkzilla/starbucks/dependabot/alerts
```

### Trigger Workflows
```bash
# Run CodeQL scan
gh workflow run codeql.yml

# Run secret scanning
gh workflow run secret-scanning.yml

# Run OSSF Scorecard
gh workflow run security-scorecard.yml
```

### View Workflow Results
```bash
# List recent runs
gh run list --limit 10

# View specific workflow
gh run view <run-id> --log

# Watch workflow
gh run watch <run-id>
```

## 🔍 What Each Tool Does

| Tool | Purpose | When It Runs |
|------|---------|--------------|
| **CodeQL** | Finds security vulnerabilities in code | Push, PR, Daily |
| **TruffleHog** | Detects leaked secrets | Push, PR, Weekly |
| **Gitleaks** | Additional secret detection | Push, PR, Weekly |
| **Dependabot** | Updates dependencies | Weekly (Monday) |
| **Dependency Review** | Blocks vulnerable deps in PRs | Every PR |
| **OSSF Scorecard** | Evaluates security posture | Weekly (Saturday) |

## ⚡ Common Tasks

### Responding to CodeQL Alert
1. Go to Security → Code scanning
2. Click on the alert
3. Review the code path
4. Fix the vulnerability
5. Push the fix
6. Verify alert closes

### Responding to Secret Alert
1. **IMMEDIATELY** revoke the secret
2. Go to Security → Secret scanning
3. Click on the alert
4. Remove secret from code
5. Mark as resolved
6. Update secret in secure location

### Reviewing Dependabot PR
1. Check the changelog
2. Review security advisory (if any)
3. Run tests locally
4. Approve and merge if safe
5. Monitor for issues

## 🛡️ Security Checklist

### Before Committing
- [ ] No secrets in code
- [ ] No hardcoded credentials
- [ ] Dependencies updated
- [ ] Tests passing
- [ ] Code reviewed

### Before Merging PR
- [ ] All status checks pass
- [ ] CodeQL scan clean
- [ ] No new secrets detected
- [ ] Dependency review passed
- [ ] Code review approved

### Weekly Review
- [ ] Check security dashboard
- [ ] Review open alerts
- [ ] Merge Dependabot PRs
- [ ] Check OSSF Scorecard
- [ ] Update documentation

## 🚫 What NOT to Commit

```bash
# Secrets
AWS_ACCESS_KEY_ID=AKIA...
DATABASE_URL=postgres://user:pass@host/db
API_KEY=sk_live_...

# Credentials
password=secret123
token=ghp_...
private_key=-----BEGIN RSA PRIVATE KEY-----

# Configuration
.env
.env.local
config/secrets.yml
credentials.json
```

## ✅ What TO Use Instead

```bash
# Use environment variables
export AWS_ACCESS_KEY_ID=...

# Use GitHub Secrets
${{ secrets.AWS_ACCESS_KEY_ID }}

# Use .env files (gitignored)
# .env
AWS_ACCESS_KEY_ID=...

# Use secret management
aws secretsmanager get-secret-value
```

## 📞 Emergency Contacts

| Issue | Contact | Response Time |
|-------|---------|---------------|
| Critical vulnerability | @jkzilla | < 24 hours |
| Leaked secret | @jkzilla | Immediate |
| Security question | @jkzilla | < 48 hours |

## 🔗 Quick Links

- [Full Setup Guide](../GITHUB_ADVANCED_SECURITY_SETUP.md)
- [Security Policy](../SECURITY.md)
- [Testing Guide](../TESTING.md)
- [GitHub Security Docs](https://docs.github.com/en/code-security)

---

**Keep this handy!** Bookmark this page for quick reference.
