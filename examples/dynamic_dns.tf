# Dynamic DNS Configuration Example

variable "international_domains" {
  description = "List of international domain names"
  type        = list(string)
  default = [
    "münchen.de",
    "日本.jp",
    "россия.рф"
  ]
}

locals {
  # Convert all international domains to ASCII-compatible encoding
  ascii_domains = [
    for domain in var.international_domains :
    provider::punycode::encode(domain)
  ]
}

output "dns_records" {
  value = {
    for idx, domain in var.international_domains :
    domain => local.ascii_domains[idx]
  }
  description = "Mapping of Unicode domains to their ASCII-encoded versions"
}
