# Upstream framework adapter

Rootly 5.20+ serves both Terraform SDKv2 and Plugin Framework resources. Its
framework constructor is inside an `internal` package. This small local module
uses an upstream-prefixed module path so Go permits the import; it delegates
directly to the upstream provider and copies no implementation code.

The parent provider replaces this module with `./shim` and combines both
providers using Pulumi's mux bridge. Keep the Rootly version in both `go.mod`
files aligned when updating, run `go mod tidy` in both modules, and regenerate
the schema, bridge metadata, and SDKs from the repository root.
