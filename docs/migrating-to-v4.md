# Migrating to v4

Version 4 updates the Rootly Terraform provider from 5.18.0 to 5.21.0 and the
Pulumi Terraform bridge to 3.139.0. It includes both upstream Terraform provider
frameworks, retaining all existing resource tokens. The upstream service lookup
and schedule rotation schema changes require a major SDK version.

## Service data sources

`getService` now requires the service ID and returns its full details:

```typescript
const service = await rootly.getService({ id: "existing-service-id" });
```

The previous name, slug, date, integration ID, and broadcast filters are no
longer accepted. `getServices()` now returns all services without input filters.
To locate a service by slug, filter that result explicitly:

```typescript
const { services } = await rootly.getServices();
const service = services.find(service => service.slug === "api");
if (!service) {
    throw new Error("Service api was not found");
}
```

`getServices` no longer returns the synthetic `id` or echoed filter fields.
`getService` no longer returns the old `createdAt` filter map. Both functions
return additional service details from the upstream provider.

## Schedule rotations

`scheduleRotationableAttributes` is now a typed object. Replace snake_case map
keys with SDK property names and use a number for `shiftLength`:

```typescript
scheduleRotationableAttributes: {
    handoffTime: "09:00",
    shiftLength: 12,
    shiftLengthUnit: "hours",
}
```

`handoffTime` is required; `handoffDay`, `shiftLength`, and `shiftLengthUnit` are
optional, depending on the rotation type. Go callers use
`ScheduleRotationScheduleRotationableAttributesArgs` instead of `pulumi.StringMap`.

The provider migrates saved v3 rotation attribute maps when reading prior state,
including string shift lengths and secret values. Keep existing resource names
and IDs, update program inputs, then run `pulumi preview --refresh` before applying.
The framework migration itself does not require replacing rotations. Some rotation
outputs are now optional in the generated SDK and callers must handle absent values.

## Go and Node.js SDKs

Update Go imports from `github.com/rootlyhq/pulumi-rootly/sdk/v3/go/rootly` to
`github.com/rootlyhq/pulumi-rootly/sdk/v4/go/rootly`. The Go SDK now declares its
dependencies and requires Go 1.26 or later. Provider development uses Go 1.26.8.

The npm package remains `@rootly/pulumi`. The generated SDK is built with
TypeScript 5 and current Node.js type definitions. Repository builds use Yarn
1.22.22 and commit the SDK lockfile.

## Additions and security fixes

New resources: `StatusPageComponent`, `StatusPageComponentGroup`, and
`UserOnCallRole`. New data source: `getUsers`.

Dependency updates address GitHub alerts 22–26 with gRPC 1.83.2,
go-git 6.0.0-alpha.5, and OpenTelemetry OpenTracing bridge 1.45.0.
The update also includes `golang.org/x/crypto` 0.56.0 for the SSH denial-of-service
fixes reported by govulncheck.

Release this change as **v4.0.0**, with the matching **sdk/v4.0.0** Go module tag.
