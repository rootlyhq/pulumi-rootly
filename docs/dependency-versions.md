# Dependency baseline

The provider, Terraform shim, Go SDK, examples, and Node.js SDK are maintained
together. Dependabot covers all four Go modules, the Node.js SDK, and GitHub
Actions. Updates must include regenerated schema, mux metadata, and SDKs whenever
generation changes.

| Component | Version |
| --- | --- |
| Terraform Rootly provider and shim | 5.25.2 |
| Pulumi Terraform bridge | 3.143.0 |
| Pulumi CLI and Go packages | 3.274.1 |
| Node.js SDK Pulumi dependency minimum | 3.274.1 |
| Python SDK Pulumi dependency minimum | 3.262.0 |
| Go JOSE dependency minimum | 4.1.6 |
| .NET Pulumi dependency baseline | 3.113.2 |
| Go | 1.27.4 |
| Node.js | 24.22.3 |
| TypeScript / Node.js types | 7.0.2 / 24.13.4 |
| Python / .NET SDK | 3.14.7 / 10.0.400 |
| Yarn Classic | 1.22.22 |
| golangci-lint / govulncheck | 2.14.0 / 1.8.0 |
| pulumictl / GoReleaser | 0.0.50 / 2.18.1 |
| Random plugin used by the build | 4.21.1 |

## Compatibility constraints

- Keep the Terraform SDK replacement aligned with Pulumi's bridge:
  `github.com/pulumi/terraform-plugin-sdk/v2` at
  `v2.0.0-20260318212141-5525259d096b`.
- The examples retain mpb 8.15.2. Check Pulumi integration helper compatibility
  before updating it; earlier updates changed the `ProxyReader` signature.
- Node.js and its type definitions stay on the Node 24 LTS line. The published
  npm package declares `engines.node: >=24`.
- Yarn remains on Classic 1.22.22 for the build and SDK linking workflow.
  Transitive npm dependencies follow upstream-supported version ranges.

## Security validation

Run schema/SDK generation, provider tests and lint, Go SDK and shim compilation,
example helper compilation with `-tags=all`, and a frozen Node.js SDK build.
Run govulncheck 1.8.0 built with the repository's Go toolchain in each Go module,
and audit the Node.js dependency lockfile.

Credential regressions exercise provider Check inputs (including nested headers
and passwords), heartbeat Read outputs, schema secret flags, and unchanged public
sibling values. Existing tests cover mux dispatch and schedule rotation state
compatibility. Secret tagging applies to new provider results; it does not
retroactively rewrite historical checkpoints.
