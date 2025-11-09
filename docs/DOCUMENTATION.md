# Documentation Structure

This provider follows the [Terraform Registry Provider Documentation](https://developer.hashicorp.com/terraform/registry/providers/docs) format.

## Directory Structure

```
docs/
├── index.md                    # Provider overview and configuration
└── functions/
    ├── encode.md              # Documentation for encode function
    └── decode.md              # Documentation for decode function

templates/                      # Templates for automatic doc generation
├── index.md.tmpl              # Provider index template
└── functions.md.tmpl          # Function documentation template

examples/
├── basic.tf                   # Basic usage examples
├── dynamic_dns.tf            # Advanced use case
└── functions/
    ├── encode/
    │   └── function.tf       # Encode function examples
    └── decode/
        └── function.tf       # Decode function examples
```

## Generating Documentation

The documentation is automatically generated using [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs).

### Manual Generation

To regenerate documentation manually:

```bash
go generate ./...
```

This will:
1. Read the provider schema from the code
2. Use templates in `templates/` directory
3. Include examples from `examples/` directory
4. Generate markdown files in `docs/` directory

### Automatic Generation

Documentation is automatically generated during the release process via the `go:generate` directive in `main.go`.

## Documentation Guidelines

### Writing Function Documentation

Each function should have:

1. **Description**: Clear explanation of what the function does
2. **Example Usage**: Multiple examples showing common use cases
3. **Signature**: The function signature showing parameters and return type
4. **Arguments**: Detailed description of each parameter
5. **Return Value**: Description of what the function returns
6. **Notes**: Any important information about behavior or limitations

### Adding New Examples

To add new examples:

1. Create a `.tf` file in the appropriate `examples/` subdirectory
2. Reference it in the template or documentation
3. Run `go generate` to update the docs

### Testing Documentation

Preview documentation before publishing:

1. Use the [Terraform Registry Doc Preview Tool](https://registry.terraform.io/tools/doc-preview)
2. Copy and paste your markdown content
3. Verify formatting and rendering

## Publishing

Documentation is published automatically with each provider release. The Terraform Registry reads the `docs/` directory from the Git tag and displays it on the provider page.

### Storage Limits

- Each document can contain up to 500KB
- Documents exceeding this limit will be truncated
- Keep documentation concise and well-organized

## Frontmatter

Documentation files can include YAML frontmatter for metadata:

```yaml
---
page_title: "encode Function - Punycode Provider"
description: |-
  Brief description of the function
---
```

Supported attributes:
- `page_title`: The title displayed in navigation
- `description`: Brief description for SEO and display
- `subcategory`: Optional grouping for navigation (not used in this provider)

## Resources

- [Provider Documentation Format](https://developer.hashicorp.com/terraform/registry/providers/docs)
- [terraform-plugin-docs Tool](https://github.com/hashicorp/terraform-plugin-docs)
- [Doc Preview Tool](https://registry.terraform.io/tools/doc-preview)
