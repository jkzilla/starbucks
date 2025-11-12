# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| 0.0.x   | :white_check_mark: |

## Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, please report them via GitHub Security Advisories:

1. Go to the repository's Security tab
2. Click "Report a vulnerability"
3. Fill out the form with details about the vulnerability

Alternatively, you can email security concerns to: [your-security-email@example.com]

Please include the following information:

- Type of issue (e.g. buffer overflow, SQL injection, cross-site scripting, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit it

### Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 7 days
- **Fix Timeline**: Depends on severity
  - Critical: 24-48 hours
  - High: 7 days
  - Medium: 30 days
  - Low: 90 days

## Security Measures

### GitHub Advanced Security

This project uses GitHub Advanced Security features:

#### 1. **Code Scanning (CodeQL)**
- Automated scanning for security vulnerabilities
- Runs on every push and pull request
- Daily scheduled scans
- Languages: Go, JavaScript/TypeScript

#### 2. **Secret Scanning**
- Automatic detection of leaked secrets
- Push protection enabled
- Partner patterns and custom patterns
- TruffleHog integration for enhanced detection

#### 3. **Dependency Scanning (Dependabot)**
- Automatic dependency updates
- Security vulnerability alerts
- Weekly update schedule
- Grouped updates for easier review

#### 4. **Dependency Review**
- Blocks PRs with vulnerable dependencies
- License compliance checking
- Fails on moderate+ severity vulnerabilities

#### 5. **OSSF Scorecard**
- Weekly security posture evaluation
- Best practices compliance
- Automated recommendations

### Additional Security Tools

- **TruffleHog**: Secret scanning with verification
- **Gitleaks**: Additional secret detection
- **golangci-lint**: Go code quality and security
- **ESLint**: JavaScript/TypeScript security rules
- **Playwright**: E2E security testing

## Security Best Practices

### For Contributors

1. **Never commit secrets**
   - Use `.env` files (gitignored)
   - Use GitHub Secrets for CI/CD
   - Enable push protection

2. **Keep dependencies updated**
   - Review Dependabot PRs promptly
   - Check for security advisories
   - Use `npm audit` and `go mod tidy`

3. **Code review**
   - All changes require review
   - Security-sensitive changes need 2+ reviews
   - Use branch protection rules

4. **Testing**
   - Write security tests
   - Test authentication/authorization
   - Validate input sanitization

### For Maintainers

1. **Enable branch protection**
   - Require pull request reviews
   - Require status checks
   - Require signed commits
   - Restrict force pushes

2. **Configure security settings**
   - Enable push protection
   - Configure secret scanning
   - Set up security policies
   - Enable vulnerability alerts

3. **Monitor security alerts**
   - Review Dependabot alerts
   - Check CodeQL findings
   - Monitor secret scanning
   - Review OSSF Scorecard

4. **Incident response**
   - Have a security contact
   - Document response procedures
   - Maintain security changelog
   - Communicate vulnerabilities

## Security Checklist

### Repository Settings

- [x] GitHub Advanced Security enabled
- [x] Secret scanning enabled
- [x] Push protection enabled
- [x] Dependabot alerts enabled
- [x] Dependabot security updates enabled
- [x] Code scanning enabled
- [x] Dependency review enabled
- [ ] Branch protection rules configured
- [ ] Required reviews configured
- [ ] Signed commits required
- [ ] Security policy published

### CI/CD Security

- [x] CodeQL workflow configured
- [x] Secret scanning workflow
- [x] Dependency review workflow
- [x] OSSF Scorecard workflow
- [x] TruffleHog integration
- [ ] SAST tools configured
- [ ] Container scanning
- [ ] Infrastructure scanning

### Application Security

- [ ] Input validation implemented
- [ ] Output encoding implemented
- [ ] Authentication implemented
- [ ] Authorization implemented
- [ ] HTTPS enforced
- [ ] Security headers configured
- [ ] Rate limiting implemented
- [ ] Logging and monitoring
- [ ] Error handling (no info leakage)
- [ ] Secure session management

## Vulnerability Disclosure

We follow coordinated vulnerability disclosure:

1. **Private Disclosure**: Report to security team
2. **Acknowledgment**: Within 48 hours
3. **Investigation**: 7-30 days
4. **Fix Development**: Based on severity
5. **Testing**: Verify fix
6. **Release**: Deploy patched version
7. **Public Disclosure**: After fix is deployed
8. **Credit**: Reporter credited (if desired)

## Security Updates

Security updates are released as:

- **Patch releases** (0.0.x) for minor issues
- **Minor releases** (0.x.0) for moderate issues
- **Major releases** (x.0.0) for breaking security changes

Subscribe to:
- GitHub Security Advisories
- Release notifications
- Security mailing list (if available)

## Compliance

This project aims to comply with:

- OWASP Top 10
- CWE Top 25
- NIST Cybersecurity Framework
- OSSF Best Practices

## Resources

- [GitHub Security Best Practices](https://docs.github.com/en/code-security)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [CWE Top 25](https://cwe.mitre.org/top25/)
- [OSSF Scorecard](https://github.com/ossf/scorecard)

## Contact

For security concerns, contact:
- GitHub Security Advisories (preferred)
- Email: [your-security-email@example.com]
- Security team: @jkzilla

---

**Last Updated**: November 12, 2025
