# Quick Reference: Release Setup

## One-Time Setup (Do This First!)

### 1. Generate GPG Key
```bash
gpg --full-generate-key
# Choose: RSA, 4096 bits, set expiration
```

### 2. Get Your GPG Fingerprint
```bash
gpg --list-secret-keys --keyid-format=long
# Copy the fingerprint (40 character hex string)
```

### 3. Export Keys
```bash
# Public key (for Terraform Registry)
gpg --armor --export your-email@example.com > public.asc

# Private key (for GitHub Secrets)
gpg --armor --export-secret-keys your-email@example.com > private.asc
```

### 4. Add to Terraform Registry
1. Visit: https://registry.terraform.io/settings/gpg-keys
2. Sign in with GitHub
3. Paste contents of `public.asc`

### 5. Add to GitHub Secrets
Repository Settings → Secrets and variables → Actions → New repository secret

- **Name:** `GPG_PRIVATE_KEY`
  **Value:** Contents of `private.asc`

- **Name:** `PASSPHRASE`
  **Value:** Your GPG key passphrase

### 6. Clean up key files
```bash
rm public.asc private.asc
```

## Creating Releases

### Simple Release
```bash
git tag v0.1.0
git push origin v0.1.0
```

GitHub Actions automatically handles the rest!

### Check Release Status
Go to: https://github.com/henrikhorluck/terraform-provider-punycode/actions

## Publishing to Registry (First Time Only)

After your first successful release:

1. Go to: https://registry.terraform.io/publish/provider
2. Select organization: `henrikhorluck`
3. Select repository: `terraform-provider-punycode`
4. Click "Publish Provider"

Future releases will be automatic via webhook!

## Testing Locally (Before Release)

```bash
# Test GoReleaser without publishing
goreleaser release --snapshot --clean
```

## Files You Created

- ✅ `.goreleaser.yml` - Build configuration
- ✅ `.github/workflows/release.yml` - GitHub Actions workflow
- ✅ `terraform-registry-manifest.json` - Registry metadata
- ✅ `RELEASING.md` - Detailed documentation
- ✅ `.github/RELEASE_SETUP.md` - This summary

## Troubleshooting

**Q: Release fails with GPG error**
A: Check that secrets are set correctly and key is RSA/DSA (not ECC)

**Q: Release not appearing in registry**
A: Check webhook at github.com/{owner}/{repo}/settings/hooks

**Q: How do I test without publishing?**
A: Use `goreleaser release --snapshot --clean`

## Need Help?

See `RELEASING.md` for detailed documentation.
