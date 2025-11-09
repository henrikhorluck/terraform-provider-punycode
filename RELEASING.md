# Releasing the Provider

This document describes the process for releasing a new version of the terraform-provider-punycode.

## Prerequisites

Before creating a release, you need to:

1. **Generate a GPG key** for signing releases (if you haven't already):
   ```bash
   gpg --full-generate-key
   ```
   - Choose RSA or DSA (not ECC)
   - Key size: 4096 bits recommended
   - Set an expiration date
   - Provide your name and email

2. **Export your GPG public key**:
   ```bash
   gpg --armor --export "{Key ID or email address}"
   ```
   
3. **Add the public key to the Terraform Registry**:
   - Go to [Terraform Registry - Signing Keys](https://registry.terraform.io/settings/gpg-keys)
   - Add your ASCII-armored public key

4. **Add secrets to your GitHub repository**:
   - Go to your repository Settings > Secrets and variables > Actions
   - Add the following secrets:
     - `GPG_PRIVATE_KEY`: Export with `gpg --armor --export-secret-keys [key ID or email]`
     - `PASSPHRASE`: The passphrase for your GPG private key

## Release Process

### Using GitHub (Recommended)

The repository is configured with GitHub Actions to automatically create releases:

1. **Ensure all changes are committed and pushed** to the `main` branch

2. **Create and push a version tag**:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **GitHub Actions will automatically**:
   - Build binaries for multiple platforms
   - Sign the release
   - Create a GitHub release
   - Upload all assets

4. **Monitor the release workflow**:
   - Go to the "Actions" tab in your GitHub repository
   - Check the "Release" workflow status

### Using GoReleaser Locally

If you prefer to create releases locally:

1. **Install GoReleaser**:
   ```bash
   # macOS
   brew install goreleaser
   
   # Other platforms: https://goreleaser.com/install/
   ```

2. **Cache your GPG passphrase**:
   ```bash
   gpg --armor --detach-sign
   ```
   (Enter your passphrase when prompted)

3. **Set your GitHub token**:
   ```bash
   export GITHUB_TOKEN="your_personal_access_token"
   export GPG_FINGERPRINT="your_gpg_fingerprint"
   ```

4. **Create and push the tag**:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

5. **Run GoReleaser**:
   ```bash
   goreleaser release --clean
   ```

## Version Numbering

Follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version (v2.0.0): Incompatible API changes
- **MINOR** version (v1.1.0): New functionality in a backward compatible manner
- **PATCH** version (v1.0.1): Backward compatible bug fixes

### Pre-release versions

Use hyphens for pre-release versions:
- `v1.0.0-alpha`
- `v1.0.0-beta.1`
- `v1.0.0-rc.1`

## What Gets Released

Each release includes:

1. **Binaries for multiple platforms**:
   - Linux: amd64, 386, arm, arm64
   - macOS (Darwin): amd64, arm64
   - Windows: amd64, 386, arm, arm64
   - FreeBSD: amd64, 386, arm, arm64

2. **Required files**:
   - `terraform-provider-punycode_{VERSION}_{OS}_{ARCH}.zip` - Binary archives
   - `terraform-provider-punycode_{VERSION}_manifest.json` - Registry manifest
   - `terraform-provider-punycode_{VERSION}_SHA256SUMS` - Checksums file
   - `terraform-provider-punycode_{VERSION}_SHA256SUMS.sig` - GPG signature

## Publishing to Terraform Registry

After creating your first release:

1. **Sign in to the Terraform Registry** with your GitHub account

2. **Navigate to** [Publish > Provider](https://registry.terraform.io/publish/provider)

3. **Select your organization and repository**

4. **The registry will**:
   - Create a webhook on your repository
   - Automatically ingest new releases when tags are pushed
   - Validate signatures for each release

## Troubleshooting

### Webhook Issues

If releases aren't appearing in the registry:

1. Check for existing webhooks at `https://github.com/{owner}/{repo}/settings/hooks`
2. Remove any broken webhooks for `registry.terraform.io`
3. Use the "Resync" button on the provider settings page in the Terraform Registry

### GPG Signature Errors

If you get GPG errors:

- Ensure your key type is RSA or DSA (not ECC)
- Verify the passphrase is correct
- Make sure the secrets are properly set in GitHub

### Build Failures

If the release workflow fails:

1. Check the Actions logs for specific errors
2. Verify `go.mod` is up to date (`go mod tidy`)
3. Ensure all tests pass (`go test ./...`)

## Additional Resources

- [Terraform Provider Publishing Guide](https://developer.hashicorp.com/terraform/registry/providers/publishing)
- [GoReleaser Documentation](https://goreleaser.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
