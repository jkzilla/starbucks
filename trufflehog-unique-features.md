# TruffleHog Features Not Available in JFrog Xray

## Secret Detection & Credential Scanning

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Secret scanning in code | ✅ | ❌ |
| Credential leak detection | ✅ | ❌ |
| API key detection | ✅ | ❌ |
| Private key scanning | ✅ | ❌ |
| Token detection | ✅ | ❌ |
| Password detection | ✅ | ❌ |
| AWS credentials scanning | ✅ | ❌ |
| Database connection strings | ✅ | ❌ |
| OAuth tokens detection | ✅ | ❌ |

## Git History Analysis

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Full git history scanning | ✅ | ❌ |
| Commit-by-commit analysis | ✅ | ❌ |
| Branch scanning | ✅ | ❌ |
| Historical secret detection | ✅ | ❌ |
| Deleted file scanning | ✅ | ❌ |
| Git diff analysis | ✅ | ❌ |
| Scan all branches simultaneously | ✅ | ❌ |
| Scan from specific commit range | ✅ | ❌ |

## Secret Verification

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Active secret verification | ✅ | ❌ |
| Live credential validation | ✅ | ❌ |
| API endpoint testing | ✅ | ❌ |
| Reduce false positives via verification | ✅ | ❌ |
| 700+ detector types with verification | ✅ | ❌ |
| Real-time credential checking | ✅ | ❌ |

## Detection Capabilities

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Entropy-based detection | ✅ | ❌ |
| Regex pattern matching | ✅ | ❌ |
| High entropy string detection | ✅ | ❌ |
| Custom regex patterns | ✅ | ❌ |
| 700+ built-in secret detectors | ✅ | ❌ |
| Cloud provider credentials (AWS, Azure, GCP) | ✅ | ❌ |
| SaaS platform tokens (GitHub, Slack, etc.) | ✅ | ❌ |
| Database credentials | ✅ | ❌ |
| Generic API keys | ✅ | ❌ |

## Source Control Integration

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| GitHub repository scanning | ✅ | Limited |
| GitLab repository scanning | ✅ | Limited |
| Bitbucket scanning | ✅ | Limited |
| Local git repository scanning | ✅ | ❌ |
| GitHub Enterprise support | ✅ | ❌ |
| GitLab self-hosted support | ✅ | ❌ |
| Pull request scanning | ✅ | ❌ |
| Pre-commit hooks | ✅ | ❌ |
| Pre-push hooks | ✅ | ❌ |

## File System & Archive Scanning

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Local file system scanning | ✅ | Limited |
| Archive file scanning (zip, tar, etc.) | ✅ | ✅ |
| Compressed file analysis | ✅ | ✅ |
| Nested archive scanning | ✅ | ✅ |
| S3 bucket scanning | ✅ | ❌ |
| GCS bucket scanning | ✅ | ❌ |
| Azure Blob storage scanning | ✅ | ❌ |
| Filesystem directory scanning | ✅ | ❌ |

## Scanning Modes & Flexibility

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| CLI-based scanning | ✅ | ✅ |
| Standalone operation (no server required) | ✅ | ❌ |
| Lightweight deployment | ✅ | ❌ |
| Docker image scanning | ✅ | ✅ |
| Binary file scanning | ✅ | ✅ |
| Base64 encoded data scanning | ✅ | ✅ |
| JSON output format | ✅ | ✅ |
| SARIF output format | ✅ | Limited |
| Custom output formats | ✅ | Limited |

## Performance & Optimization

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Parallel scanning | ✅ | ✅ |
| Incremental scanning | ✅ | ✅ |
| Resume interrupted scans | ✅ | ✅ |
| Scan only changed files | ✅ | Limited |
| Optimized for large repositories | ✅ | ❌ |
| Memory-efficient scanning | ✅ | ❌ |
| Fast git history traversal | ✅ | ❌ |

## Filtering & Configuration

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Include/exclude file patterns | ✅ | ✅ |
| Path-based filtering | ✅ | ✅ |
| Regex-based exclusions | ✅ | Limited |
| Allowlist configuration | ✅ | ❌ |
| Custom detector configuration | ✅ | ❌ |
| Severity filtering | ✅ | ✅ |
| Time-based filtering (scan since date) | ✅ | ❌ |

## Open Source & Community

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Open source version available | ✅ | ❌ |
| Community-driven development | ✅ | ❌ |
| Free for unlimited use | ✅ (OSS version) | ❌ |
| Self-hosted without licensing | ✅ | ❌ |
| Transparent detection logic | ✅ | ❌ |

## Enterprise Features (TruffleHog Enterprise)

| Feature | TruffleHog Enterprise | JFrog Xray |
|---------|---------------------|------------|
| Secret sprawl detection | ✅ | ❌ |
| Cross-repository secret tracking | ✅ | ❌ |
| Secret lifecycle management | ✅ | ❌ |
| Automated secret rotation alerts | ✅ | ❌ |
| Secret exposure timeline | ✅ | ❌ |
| Developer attribution | ✅ | Limited |
| Secret remediation workflow | ✅ | ❌ |

## Specialized Detection

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| Slack tokens | ✅ | ❌ |
| Discord webhooks | ✅ | ❌ |
| Stripe API keys | ✅ | ❌ |
| Twilio credentials | ✅ | ❌ |
| SendGrid API keys | ✅ | ❌ |
| Mailgun credentials | ✅ | ❌ |
| JWT tokens | ✅ | ❌ |
| SSH private keys | ✅ | ❌ |
| PGP private keys | ✅ | ❌ |
| RSA private keys | ✅ | ❌ |

## CI/CD Integration

| Feature | TruffleHog | JFrog Xray |
|---------|-----------|------------|
| GitHub Actions native support | ✅ | ✅ |
| GitLab CI integration | ✅ | ✅ |
| CircleCI integration | ✅ | ✅ |
| Jenkins plugin | ✅ | ✅ |
| Pre-commit framework | ✅ | ❌ |
| Git hooks integration | ✅ | ❌ |
| Fail on secret detection | ✅ | ✅ |
| Exit code customization | ✅ | Limited |

## Summary

**TruffleHog** specializes in:
- Secret and credential detection across code, git history, and file systems
- Active verification of discovered secrets
- Lightweight, standalone operation
- Git-native scanning capabilities
- 700+ built-in secret detectors
- Open source availability

**JFrog Xray** specializes in:
- Vulnerability and license compliance
- Software composition analysis
- Artifact security management
- Enterprise governance and policy enforcement

## Use Case Comparison

| Use Case | Best Tool |
|----------|-----------|
| Find leaked API keys in code | TruffleHog |
| Scan git history for credentials | TruffleHog |
| Verify if secrets are still active | TruffleHog |
| Detect CVEs in dependencies | JFrog Xray |
| License compliance checking | JFrog Xray |
| Container vulnerability scanning | JFrog Xray |
| Policy enforcement | JFrog Xray |
| SBOM generation | JFrog Xray |

**Recommendation**: Use both tools together for comprehensive security coverage - TruffleHog for secret detection and JFrog Xray for vulnerability and compliance management.
