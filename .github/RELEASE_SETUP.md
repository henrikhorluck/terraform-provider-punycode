# Release Workflow Setup Summary

## ✅ Files Created

The following files have been added to configure automatic releases:

### 1. `.goreleaser.yml`

- Configures GoReleaser to build binaries for multiple platforms
- Creates checksums and signs releases with GPG
- Builds for: Linux, macOS, Windows, FreeBSD (multiple architectures each)

### 2. `.github/workflows/release.yml`

- GitHub Actions workflow that triggers on version tags (v*)
- Automatically builds and releases when you push a tag
- Handles GPG signing and uploads to GitHub Releases

### 3. `terraform-registry-manifest.json`

- Required by Terraform Registry
- Specifies protocol version 6.0 (Terraform Plugin Framework)

### 4. `RELEASING.md`

- Complete documentation for the release process
- Instructions for GPG key setup
- Troubleshooting guide

## 🔧 Setup Required

Before you can create releases, you need to:

### 1. Generate a GPG Key

```bash
gpg --full-generate-key
```

Choose:
- Type: RSA (not ECC)
- Size: 4096 bits
- Set an expiration date
- Use your real name and email

### 2. Add GPG Public Key to Terraform Registry

```bash
# Export your public key
gpg --armor --export "your-email@example.com"
```

Then:
1. Go to https://registry.terraform.io/settings/gpg-keys
2. Sign in with GitHub
3. Paste your public key

### 3. Add Secrets to GitHub Repository

Go to: `https://github.com/henrikhorluck/terraform-provider-punycode/settings/secrets/actions`

Add these secrets:

**GPG_PRIVATE_KEY:**
```bash
gpg --armor --export-secret-keys "your-email@example.com"
```

**PASSPHRASE:**
- The passphrase you set for your GPG key

## 🚀 Creating a Release

Once setup is complete:

```bash
# Tag a version
git tag v0.1.0

# Push the tag
git push origin v0.1.0
```

GitHub Actions will automatically:
1. Build binaries for all platforms
2. Create SHA256 checksums
3. Sign the checksums with your GPG key
4. Create a GitHub Release with all assets
5. The Terraform Registry will detect the release via webhook

## 📋 Checklist

Before your first release:

- [ ] Generate GPG key (RSA or DSA, not ECC)
- [ ] Add GPG public key to Terraform Registry
- [ ] Add `GPG_PRIVATE_KEY` secret to GitHub
- [ ] Add `PASSPHRASE` secret to GitHub
- [ ] Verify Actions are enabled in repository settings
- [ ] Create and push a test tag (e.g., v0.1.0)
- [ ] Verify release appears in GitHub Releases
- [ ] Publish provider on Terraform Registry (first time only)

## 🔗 Publishing to Terraform Registry (First Time)

After your first successful GitHub release:

1. Go to https://registry.terraform.io/publish/provider
2. Sign in with GitHub
3. Select your organization and the `terraform-provider-punycode` repository
4. The registry will create a webhook for automatic future releases

## 📚 Additional Resources

- [Terraform Provider Publishing Guide](https://developer.hashicorp.com/terraform/registry/providers/publishing)
- [GoReleaser Documentation](https://goreleaser.com/)
- [GitHub Actions for Go](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)

## 🎯 Next Steps

1. Complete the checklist above
2. Test the release workflow with v0.1.0
3. If successful, publish to Terraform Registry
4. Future releases will be automatic!
