# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a Pulumi provider for Rootly, a platform for incident management and workflow automation. The provider acts as a bridge between Pulumi and the Rootly Terraform provider, allowing users to manage Rootly resources (services, severities, workflows, etc.) through Pulumi infrastructure-as-code.

## Prerequisites

This project uses [mise](https://mise.jdx.dev) for development tool versioning (Go, Node.js, Python, Yarn, Pulumi, golangci-lint). Run `make install_deps` to install all required tools.

## Key Commands

### Development Workflow
- `make development` - Complete development build (provider + SDKs + install + cleanup)
- `make build` - Build provider and all SDKs
- `make provider` - Build only the provider binary
- `make build_nodejs` - Build Node.js SDK specifically
- `make tfgen` - Generate Pulumi schema from Terraform provider

### Provider Updates
- `make update_provider` - Update the upstream Terraform provider dependency

### Testing and Quality
- `make test` - Run integration tests (in examples/ directory)
- `make lint_provider` - Run golangci-lint on provider code
- `cd provider && golangci-lint run -c ../.golangci.yml` - Manual lint execution

### Version Management and Releases
- `make version-show` - Show current and next versions
- `make release-patch` - Bump patch version and push tag (triggers CI release)
- `make release-minor` - Bump minor version and push tag (triggers CI release)
- `make release-major` - Bump major version and push tag (triggers CI release)

### Cleanup
- `make cleanup` - Remove build artifacts and temporary files
- `make clean` - Remove all SDK build outputs

## Architecture

### High-Level Structure
The repository follows Pulumi's standard terraform-bridge provider pattern:

1. **Provider Bridge** (`provider/`): Go code that bridges Terraform provider to Pulumi
   - `resources.go` - Main provider configuration and resource mapping
   - `cmd/pulumi-resource-rootly/` - Provider binary entry point
   - `cmd/pulumi-tfgen-rootly/` - Schema generation tool

2. **Generated SDKs** (`sdk/`): Language-specific SDKs generated from Terraform schema
   - `nodejs/` - TypeScript/JavaScript SDK with comprehensive Rootly resource types
   - Python, Go, and .NET SDKs (build targets available)

3. **Examples and Tests** (`examples/`): Integration tests in Go that verify provider functionality

### Key Components

- **Terraform Provider Integration**: Uses `github.com/rootlyhq/terraform-provider-rootly/v2` as the upstream provider
- **Bridge Configuration**: Single module mapping (`tokens.SingleModule`) puts all resources in main module
- **Auto-aliasing**: Automatically handles resource name changes during provider updates
- **Multi-language Support**: Generates SDKs for Node.js, Python, Go, and .NET

### Resource Organization
All Rootly resources are mapped to the main `rootly` module with consistent naming:
- Terraform resources `rootly_*` become Pulumi resources in the main module
- Data sources follow the same pattern with `get*` functions

## Development Notes

### Building Locally
1. Provider must be built before SDKs: `make provider`
2. SDKs depend on generated schema: `make tfgen` → `make build_nodejs`
3. Use `make development` for complete local development setup

### Environment Variables
- `ROOTLY_API_TOKEN` - Required for testing; API token from Rootly account settings
- `ROOTLY_BASE_URL` - Optional; defaults to https://api.rootly.io

### Testing
- Integration tests require `ROOTLY_API_TOKEN` environment variable
- Tests run against live Rootly API (not mocked)
- Test parallelism set to 4 concurrent tests

### Provider Updates
When updating the upstream Terraform provider:
1. `make update_provider` - pulls latest terraform-provider-rootly
2. `make tfgen` - regenerates schema
3. `make build_sdks` - rebuilds all SDKs
4. Commit generated SDK changes (required for all PRs with code changes)

### Release Process
Releases are handled via GitHub Actions when version tags are created:
1. Use `make release-patch|minor|major` to bump version and push tag
2. CI automatically builds and publishes to NPM and GitHub Releases