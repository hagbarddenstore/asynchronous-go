#!/bin/bash

# This bash script outputs the current, desired content of version.go, using
# git describe. For best effort, pipe this to the target file. Generally, this
# only needs to be updated for releases. The actual value of will be replaced
# during build time of the makefile is used.

set -e

cat <<EOF
package version

// Package is the overall, canocial project import path under which the
// package was built.
var Package = "$(go list)"

// Version indicates which version of the binary is running. This is set to
// the latest release tag by hand, always suffixed by "+unknown". During
// build, it will be replaced by the actual version. The value here will be
// used if the application is run after a go get based installed.
var Version = "$(git describe --match 'v[0-9]*' --dirty='.m' --always)+unknown"
EOF
