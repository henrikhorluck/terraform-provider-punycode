package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"golang.org/x/net/idna"
)

var _ function.Function = &DecodeFunction{}

type DecodeFunction struct{}

func NewDecodeFunction() function.Function {
	return &DecodeFunction{}
}

func (f *DecodeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "decode"
}

func (f *DecodeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Decode a punycode string",
		Description: "Takes an ASCII-compatible punycode string and converts it to Unicode (RFC 3492).",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The punycode string to decode",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *DecodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = req.Arguments.Get(ctx, &input)
	if resp.Error != nil {
		return
	}

	// Use the IDNA profile for punycode decoding
	decoded, err := idna.ToUnicode(input)
	if err != nil {
		resp.Error = function.NewFuncError("Failed to decode string: " + err.Error())
		return
	}

	resp.Error = resp.Result.Set(ctx, decoded)
}
