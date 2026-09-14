# Dependency refresh (2026-09-14)

The provider, Terraform shim, Go SDK, examples, and Node.js SDK were updated
together. Go module versions were checked with `go list -m -u -json all` against
each module's declared requirements after `go get -u` and `go mod tidy`.
Example test dependencies were included with `go get -tags=all -t -u ./...`.
The Node.js lockfile was refreshed with `yarn upgrade`; `yarn outdated` reports
no newer direct dependencies within the selected major versions.

| Component | Version |
| --- | --- |
| Terraform Rootly provider | 5.21.1 |
| Pulumi Terraform bridge | 3.139.0 |
| Pulumi CLI, Go packages, Node.js and Python dependency baselines | 3.262.0 |
| .NET Pulumi dependency baseline | 3.113.2 |
| Go | 1.27.1 |
| Node.js | 24.21.0 |
| TypeScript / Node.js types | 7.0.2 / 24.13.4 |
| Python / .NET SDK | 3.14.7 / 10.0.400 |
| Yarn Classic | 1.22.22 |
| golangci-lint / govulncheck | 2.13.2 / 1.8.0 |
| pulumictl / GoReleaser | 0.0.50 / 2.18.1 |
| Random plugin used by the build | 4.21.1 |

GitHub Actions are pinned to the current release commits. Local development and
the development container share `mise.toml`; CI and release templates use the
same Go and Node.js generation. The obsolete, unused `@types/mime` dependency
and `imdario/mergo` override were removed.

## Compatibility constraints

- The bridge requires Pulumi's Terraform SDK fork at
  `v2.0.0-20260318212141-5525259d096b`. Keep its `replace` directive aligned with
  the bridge rather than replacing it with HashiCorp's SDK.
- `github.com/pulumi-labs/pulumi-hcl` stays at 0.12.0. Later releases declare
  `github.com/pulumi/pulumi-hcl`, while bridge 3.139.0 still imports the old path.
  Go rejects those versions with a module-path mismatch.
- The examples use mpb 8.15.2. Both 8.16.0 and 8.16.1 change `ProxyReader` to
  return two values, breaking Pulumi 3.262.0's integration-test helper build.
- Node.js and its type definitions stay on the requested Node 24 LTS line.
  The published npm package declares `engines.node: >=24`.
- Yarn remains on the latest Classic release, 1.22.22, for the provider's build
  and SDK linking workflow. Transitive npm dependencies follow their upstream
  packages' supported version ranges.

The HCL and mpb constraints are the only newer versions reported for declared
Go requirements; the SDK and shim have none. Revisit these constraints when
upgrading the Pulumi bridge and packages.

## Validation

Provider and SDK generation/builds, provider race tests, lint, Go SDK and shim
compilation, and example helper compilation pass. The example helpers do not
contain live acceptance tests. `govulncheck` reports no vulnerable imported
packages or reachable vulnerabilities across all four Go modules. An unused
`x/crypto/openpgp` module-only advisory has no patched release. Yarn audit reports
zero vulnerabilities across 242 dependencies, and a frozen-lockfile build passes.

The development container was not built locally because the Docker daemon was
unavailable. Its configured tool versions were installed and exercised locally.
