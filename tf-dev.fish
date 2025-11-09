#!/usr/bin/env fish

# Helper script to run Terraform with local provider development

set -x TF_CLI_CONFIG_FILE (realpath (dirname (status --current-filename))/.terraformrc)

echo "Using local provider override from: $TF_CLI_CONFIG_FILE"
terraform $argv
