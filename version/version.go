package version

// Package is the overall, canocial project import path under which the
// package was built.
var Package = ""

// Version indicates which version of the binary is running. This is set to
// the latest release tag by hand, always suffixed by "+unknown". During
// build, it will be replaced by the actual version. The value here will be
// used if the application is run after a go get based installed.
var Version = "c7cc331+unknown"
