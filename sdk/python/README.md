# Rootly Pulumi Provider

The Rootly Pulumi Provider enables you to manage [Rootly](https://rootly.com) resources using infrastructure-as-code. Rootly is a comprehensive incident management and workflow automation platform that helps teams respond to and learn from incidents.

## Installing

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @rootly/pulumi
```

or `yarn`:

```bash
yarn add @rootly/pulumi
```

### Other Languages

Python, Go, and .NET SDKs are planned for future releases.

## Configuration

The following configuration points are available for the `rootly` provider:

- `rootly:apiToken` (environment: `ROOTLY_API_TOKEN`) - the API token for Rootly. You can obtain this from your Rootly account settings.
- `rootly:baseUrl` (environment: `ROOTLY_BASE_URL`) - the base URL for the Rootly API (optional, defaults to https://api.rootly.io)

## Usage Examples

### Basic Resources

```javascript
const rootly = require("@rootly/pulumi");

// Create severity levels
const sev0 = new rootly.Severity("sev0", {
    name: "SEV0",
    description: "Critical - Complete service outage",
    color: "#FF0000",
    position: 0
});

const sev1 = new rootly.Severity("sev1", {
    name: "SEV1", 
    description: "High - Significant service degradation",
    color: "#FF8C00",
    position: 1
});

// Create services
const apiService = new rootly.Service("api-service", {
    name: "API Service",
    description: "Core API backend service",
    color: "#4285F4"
});

const webService = new rootly.Service("web-service", {
    name: "Web Frontend",
    description: "Customer-facing web application", 
    color: "#34A853"
});

// Create teams
const sre = new rootly.Team("sre", {
    name: "Site Reliability Engineering",
    description: "SRE team responsible for platform reliability"
});
```

### Workflow Automation

```javascript
// Create an incident workflow
const incidentWorkflow = new rootly.WorkflowIncident("critical-incident-workflow", {
    name: "Critical Incident Response",
    description: "Automated workflow for SEV0/SEV1 incidents",
    triggerParams: {
        triggerType: "incident_created",
        severityIds: [sev0.id, sev1.id],
    },
    enabled: true
});

// Add workflow tasks
const slackNotification = new rootly.WorkflowTaskSendSlackMessage("slack-notification", {
    workflowId: incidentWorkflow.id,
    name: "Notify SRE Team",
    position: 1,
    taskParams: {
        channels: ["#incidents", "#sre-oncall"],
        message: "🚨 New {{ incident.severity.name }} incident: {{ incident.title }}"
    }
});
```

## Deploying Resources

After defining your Rootly resources, deploy them using standard Pulumi commands:

```bash
# Preview changes
pulumi preview

# Deploy resources
pulumi up

# View current stack resources
pulumi stack
```

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/rootly/api-docs/)

## Resource Types

This provider supports managing a comprehensive set of Rootly resources:

- **Incident Management**: Severities, Services, Teams, Incident Types, Incident Roles
- **Workflow Automation**: Workflows, Workflow Tasks (50+ task types)
- **Configuration**: Custom Fields, Forms, Environments, Functionalities  
- **Integrations**: Webhooks, Secrets, Heartbeats
- **Access Control**: Roles, Permissions
- **Dashboards**: Dashboard Panels, Status Pages

For a complete list of available resources and their properties, see the [Pulumi Registry documentation](https://www.pulumi.com/registry/packages/rootly/api-docs/).

## Development

### Building the Provider

```bash
# Update the upstream Terraform provider
make update_provider

# Build the provider and all SDKs
make development

# Build specific SDK
make build_nodejs
make build_python 
make build_go
make build_dotnet
```

### Testing

```bash
# Run integration tests (requires ROOTLY_API_TOKEN)
make test

# Lint provider code
make lint_provider
```

### Releases

Releases are automated via GitHub Actions when version tags are created. The CI/CD pipeline builds and publishes packages to all supported package registries.
