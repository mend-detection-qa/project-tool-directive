module github.com/example/tool-test

go 1.24.0

// tool directive (Go 1.24+) - replaces the legacy tools.go blank-import pattern
tool (
	golang.org/x/tools/cmd/stringer
	mvdan.cc/gofumpt
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/tools v0.28.0 // indirect
	mvdan.cc/gofumpt v0.7.0 // indirect
)
