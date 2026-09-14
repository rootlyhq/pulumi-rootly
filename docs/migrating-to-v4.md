# Migrating to v4

Version 4 updates the Rootly Terraform provider from 5.18.0 to 5.21.1 and the
Pulumi Terraform bridge to 3.139.0. It includes both upstream Terraform provider
frameworks, retaining all existing resource tokens. The upstream service lookup
and schedule rotation schema changes require a major SDK version.

## Service data sources

`getService` accepts either an ID or search filters and returns full service details:

```typescript
const byId = await rootly.getService({ id: "existing-service-id" });
const bySlug = await rootly.getService({ slug: "api" });
const matchingServices = await rootly.getServices({ slug: "api" });
```

Terraform provider 5.21.1 restores filters for both service data sources: `name`,
`slug`, `backstageId`, `cortexId`, `externalId`, `alertBroadcastEnabled`, and
`incidentBroadcastEnabled`. `getService` requires an ID or at least one filter;
combining an ID with filters, or matching multiple services, produces an error.
`getServices()` without filters still returns all services.

Compared with v3, `getService` no longer accepts or returns the `createdAt` filter
map. `getServices` no longer accepts the `opsgenieId` or `pagerdutyId` filters and
no longer returns a synthetic `id`. Both functions return additional service
details from the upstream provider.

## Workflow incident visibility

`triggerParams.incidentVisibilities` is now a list of booleans for `WorkflowActionItem`,
`WorkflowIncident`, and `WorkflowPostMortem`. Replace string values with booleans:

```typescript
triggerParams: {
    incidentVisibilities: [true, false],
}
```

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
dependencies and requires Go 1.27.1 or later. Provider development uses the same version.

The npm package remains `@rootly/pulumi`. The generated SDK is built with
TypeScript 7.0.2 and Node.js 26.5.1 type definitions, with Pulumi 3.262.0.
Repository builds use Node.js 26.8.2 and Yarn
1.22.22 and commit the SDK lockfile.

## Additions and security fixes

New resources: `StatusPageComponent`, `StatusPageComponentGroup`, and
`UserOnCallRole`. New data source: `getUsers`.

Dependency updates address GitHub alerts 22–27 with gRPC 1.83.2,
go-git 6.0.0-alpha.5, and OpenTelemetry OpenTracing bridge 1.46.0.
The update also includes `golang.org/x/crypto` 0.57.0 for the SSH denial-of-service
fixes reported by govulncheck.

Release this change as **v4.0.0**, with the matching **sdk/v4.0.0** Go module tag.
