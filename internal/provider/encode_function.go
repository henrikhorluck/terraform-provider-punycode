package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"golang.org/x/net/idna"
)

var _ function.Function = &EncodeFunction{}

type EncodeFunction struct{}

func NewEncodeFunction() function.Function {
	return &EncodeFunction{}
}

func (f *EncodeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "encode"
}

func (f *EncodeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Encode a string to punycode",
		Description: "Takes a Unicode string and converts it to ASCII-compatible punycode encoding (RFC 3492).",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The Unicode string to encode",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *EncodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = req.Arguments.Get(ctx, &input)
	if resp.Error != nil {
		return
	}

	// Use the IDNA profile for punycode encoding
	encoded, err := idna.ToASCII(input)
	if err != nil {
		resp.Error = function.NewFuncError("Failed to encode string: " + err.Error())
		return
	}

	resp.Error = resp.Result.Set(ctx, encoded)
}
