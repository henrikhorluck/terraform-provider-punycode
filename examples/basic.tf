# Examples for terraform-provider-punycode
# For local development, set TF_CLI_CONFIG_FILE to point to ../.terraformrc

terraform {
  required_providers {
    punycode = {
      source = "henrikhorluck/punycode"
    }
  }
}

provider "punycode" {}

# Encode Unicode domain names
output "encoded_german" {
  value       = provider::punycode::encode("münchen.de")
  description = "German domain encoded to punycode"
}

output "encoded_japanese" {
  value       = provider::punycode::encode("日本.jp")
  description = "Japanese domain encoded to punycode"
}

output "encoded_russian" {
  value       = provider::punycode::encode("россия.рф")
  description = "Russian domain encoded to punycode"
}

# Decode punycode to Unicode
output "decoded_german" {
  value       = provider::punycode::decode("xn--mnchen-3ya.de")
  description = "Punycode decoded to German"
}

output "decoded_japanese" {
  value       = provider::punycode::decode("xn--wgv71a.jp")
  description = "Punycode decoded to Japanese"
}

