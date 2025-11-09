# Examples for encode function

output "example_german" {
  value = provider::punycode::encode("münchen.de")
  # Result: "xn--mnchen-3ya.de"
}

output "example_japanese" {
  value = provider::punycode::encode("日本.jp")
  # Result: "xn--wgv71a.jp"
}

output "example_russian" {
  value = provider::punycode::encode("россия.рф")
  # Result: "xn--h1alffa9f.xn--p1ai"
}
