# Terraform Provider Punycode

This provider was essentially one-shot from

> We are going to make a terraform provider that provides a punycode-enconding and decoding function

To Claude Sonnet 4.5, use at your own risk!

A Terraform provider that provides punycode encoding and decoding functions.

## Overview

This provider implements RFC 3492 punycode encoding/decoding using Terraform's provider-defined functions. It allows you to encode Unicode domain names to ASCII-compatible encoding and decode them back.

## Requirements

- Terraform >= 1.8.0 (for provider-defined functions support)
- Go >= 1.21 (for development)

## Installation

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    punycode = {
      source = "henrikhorluck/punycode"
    }
  }
}

provider "punycode" {}
```

## Usage

### Encode Function

Converts a Unicode string to punycode (ASCII-compatible encoding):

```hcl
output "encoded" {
  value = provider::punycode::encode("münchen.de")
}
# Output: "xn--mnchen-3ya.de"

output "encoded_japanese" {
  value = provider::punycode::encode("日本.jp")
}
# Output: "xn--wgv71a.jp"
```

### Decode Function

Converts a punycode string back to Unicode:

```hcl
output "decoded" {
  value = provider::punycode::decode("xn--mnchen-3ya.de")
}
# Output: "münchen.de"

output "decoded_japanese" {
  value = provider::punycode::decode("xn--wgv71a.jp")
}
# Output: "日本.jp"
```

## Development

### Building the Provider

```bash
go build -o terraform-provider-punycode
```

### Running Tests

```bash
go test ./...
```

### Local Installation

To use the provider locally during development:

1. Build the provider:
   ```bash
   go build -o terraform-provider-punycode
   ```

2. Set the environment variable to use the local dev override:
   ```bash
   export TF_CLI_CONFIG_FILE="$(pwd)/.terraformrc"
   ```

3. Run Terraform (skip `terraform init` when using dev overrides):
   ```bash
   cd examples
   terraform plan
   terraform apply
   ```

Alternatively, use the provided helper script:
```bash
./tf-dev.fish plan
./tf-dev.fish apply
```

The `.terraformrc` file configures Terraform to use the local provider binary instead of downloading from the registry.

## What is Punycode?

Punycode is a representation of Unicode with the limited ASCII character subset used for Internet host names. It's used to encode internationalized domain names (IDN) into ASCII-compatible encoding (ACE).

For example:

- `münchen.de` → `xn--mnchen-3ya.de`
- `日本.jp` → `xn--wgv71a.jp`
- `россия.рф` → `xn--h1alffa9f.xn--p1ai`

## Releasing

See [RELEASING.md](RELEASING.md) for detailed instructions on creating and publishing releases.

Quick steps:

1. Ensure GPG key is configured (see RELEASING.md)
2. Create and push a version tag: `git tag v1.0.0 && git push origin v1.0.0`
3. GitHub Actions will automatically build and release

## License

This provider is distributed under the Mozilla Public License 2.0. See `LICENSE` for more information.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
