# Examples for decode function

output "example_german" {
  value = provider::punycode::decode("xn--mnchen-3ya.de")
  # Result: "münchen.de"
}

output "example_japanese" {
  value = provider::punycode::decode("xn--wgv71a.jp")
  # Result: "日本.jp"
}

output "example_russian" {
  value = provider::punycode::decode("xn--h1alffa9f.xn--p1ai")
  # Result: "россия.рф"
}
