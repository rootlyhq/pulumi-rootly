# Rootly Resource Provider

The Rootly Resource Provider lets you manage [Rootly](http://rootly.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @rootly/pulumi
```

or `yarn`:

```bash
yarn add @rootly/pulumi
```

### Python, Go, & .NET

*TBA*

## Configuration

The following configuration points are available for the `rootly` provider:

- `rootly:apiToken` (environment: `ROOTLY_API_TOKEN`) - the API token for `rootly`

## Creating resources

```javascript
const rootly = require("@rootly/pulumi")

new rootly.Severity("sev0", {
  name: "SEV0",
  color: "#FF0000"
})

new rootly.Service("elasticsearch_prod", {
  name: "elasticsearch-prod",
  color: "#800080"
})

new rootly.Functionality("add_items_to_card", {
  name: "Add items to card",
  color: "#FFFFFF"
})
```

## Syncing resources

Run the regular `pulumi up` command.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/rootly/api-docs/)

## Development

New resources/data sources must be added in the terraform bridge in `provider/resources.go`.

Update the SDKs with `make build_sdks` or just the NodeJS SDK with `make build_nodejs`.

Packages are automatically published on new version tags via GitHub Actions.
