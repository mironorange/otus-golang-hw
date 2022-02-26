#!/usr/bin/env bash
set -xeuo pipefail

expected='!SUTO ,olleH'
result=$(go run main.go | sed 's/^ *//;s/ *$//')
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

# # Checking module installation
# # TODO: Replace branch after merge to master
# docker run --rm -i golang:1.17-alpine sh << 'EOF'
# go install github.com/mironorange/otus-golang-hw/hw01_hello_otus@hw01_hello_otus \
#     && [ -d "$(go env GOMODCACHE)/github.com/mironorange/otus-golang-hw" ] \
#         && echo "Successful"
# EOF

echo "PASS"
