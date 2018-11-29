#!/bin/bash
set -eu

touch go.mod

CONTENT=$(cat <<-EOD
module github.com/ludw1gj/mathfever

require github.com/aws/aws-lambda-go v1.6.0
EOD
)

echo "$CONTENT" > go.mod
