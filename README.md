# Rootly Pulumi Provider

The Rootly Pulumi Provider enables you to manage [Rootly](https://rootly.com) resources using infrastructure-as-code. Rootly is a comprehensive incident management and workflow automation platform that helps teams respond to and learn from incidents.

## Installing

### Node.js (JavaScript/TypeScript)

```bash
npm install @rootly/pulumi
```

or

```bash
yarn add @rootly/pulumi
```

## Configuration

- `rootly:apiToken` (environment: `ROOTLY_API_TOKEN`) - API token from your Rootly account settings
- `rootly:baseUrl` (environment: `ROOTLY_BASE_URL`) - Base URL for the Rootly API (optional, defaults to https://api.rootly.io)

## Examples

### Severities, Environments, and Services

```typescript
import * as rootly from "@rootly/pulumi";

const sev0 = new rootly.Severity("sev0", {
    name: "SEV0",
    color: "#FF0000",
    notifyEmails: ["oncall@acme.com"],
    slackChannels: [{ id: "C06A4RZR9", name: "incidents-critical" }],
});

const sev1 = new rootly.Severity("sev1", {
    name: "SEV1",
    color: "#FFA500",
});

const production = new rootly.Environment("production", {
    name: "Production",
    color: "#FF0000",
    slackChannels: [{ id: "C06A4RZR9", name: "incidents-prod" }],
});

const api = new rootly.Service("api", {
    name: "api-gateway",
    color: "#800080",
    notifyEmails: ["platform@acme.com"],
});

const checkout = new rootly.Functionality("checkout", {
    name: "Checkout Flow",
    color: "#800080",
});
```

### Teams

```typescript
const sre = new rootly.Team("sre", {
    name: "SRE",
});
```

### Custom Form Fields

```typescript
const regionsAffected = new rootly.FormField("regions-affected", {
    name: "Regions affected",
    kind: "custom",
    inputKind: "multi_select",
    showns: ["web_new_incident_form", "web_update_incident_form"],
    requireds: ["web_new_incident_form", "web_update_incident_form"],
});

new rootly.FormFieldOption("asia", {
    formFieldId: regionsAffected.id,
    value: "Asia",
});

new rootly.FormFieldOption("europe", {
    formFieldId: regionsAffected.id,
    value: "Europe",
});
```

### Dashboards

```typescript
const overview = new rootly.Dashboard("overview", {
    name: "Incident Overview",
});

new rootly.DashboardPanel("incidents-by-severity", {
    dashboardId: overview.id,
    name: "Incidents by Severity",
    params: {
        display: "line_chart",
        datasets: [{
            collection: "incidents",
            filter: {
                operation: "and",
                rules: [{
                    operation: "and",
                    condition: "=",
                    key: "status",
                    value: "started",
                }],
            },
            groupBy: "severity",
            aggregate: {
                cumulative: false,
                key: "results",
                operation: "count",
            },
        }],
    },
});
```

### Alert Sources

```typescript
const titleField = rootly.getAlertFieldOutput({ kind: "title" });
const descField = rootly.getAlertFieldOutput({ kind: "description" });
const urlField = rootly.getAlertFieldOutput({ kind: "external_url" });

new rootly.AlertsSource("webhook-source", {
    name: "Generic webhook source",
    sourceType: "generic_webhook",
    alertSourceFieldsAttributes: [
        { alertFieldId: titleField.id, templateBody: "Server alert" },
        { alertFieldId: descField.id, templateBody: "Alert triggered" },
        { alertFieldId: urlField.id, templateBody: "https://rootly.com" },
    ],
    sourceableAttributes: {
        autoResolve: true,
        resolveState: "$.status",
        fieldMappingsAttributes: [
            { field: "state", jsonPath: "$.my_group_attribute" },
            { field: "external_id", jsonPath: "$.my_id_attribute" },
        ],
    },
});
```

### Workflow: Create Jira Issue on Incident

```typescript
const jiraWorkflow = new rootly.WorkflowIncident("jira-on-incident", {
    name: "Create a Jira Issue",
    description: "Open Jira ticket whenever incident starts",
    triggerParams: {
        triggers: ["incident_created"],
        incidentConditionKind: "IS",
        incidentKinds: ["normal"],
        incidentConditionStatus: "IS",
        incidentStatuses: ["started"],
    },
    enabled: true,
});

new rootly.WorkflowTaskCreateJiraIssue("jira-task", {
    workflowId: jiraWorkflow.id,
    taskParams: {
        title: "{{ incident.title }}",
        description: "{{ incident.summary }}",
        projectKey: "ROOT",
        issueType: { id: "10001", name: "Task" },
        status: { id: "10000", name: "To Do" },
        labels: '{{ incident.environment_slugs | concat: incident.service_slugs | join: "," }}',
    },
});
```

### Workflow: Notify Slack Channels

```typescript
const notifyWorkflow = new rootly.WorkflowIncident("notify-slack", {
    name: "Notify teams on Slack",
    description: "Send a message to teams on Slack about the incident",
    triggerParams: {
        triggers: ["incident_created"],
        incidentConditionStatus: "IS",
        incidentStatuses: ["started"],
    },
    enabled: true,
});

new rootly.WorkflowTaskSendSlackMessage("notify-slack-task", {
    workflowId: notifyWorkflow.id,
    taskParams: {
        channels: [{
            id: "{{ incident.slack_channel_id }}",
            name: "{{ incident.slack_channel_id }}",
        }],
        text: "Heads up - we have an active incident.",
    },
});
```

### Workflow: Page PagerDuty Responders

```typescript
const pageWorkflow = new rootly.WorkflowIncident("page-pagerduty", {
    name: "Page responders via PagerDuty",
    description: "Automatically page responders when incident starts",
    triggerParams: {
        triggers: ["incident_created"],
        incidentConditionStatus: "IS",
        incidentStatuses: ["started"],
    },
    enabled: true,
});

new rootly.WorkflowTaskPagePagerdutyOnCallResponders("page-pd-task", {
    workflowId: pageWorkflow.id,
    taskParams: {},
});
```

### Workflow: SMS for Critical Incidents (Using Data Source)

```typescript
const critical = rootly.getSeverityOutput({ slug: "sev0" });

const smsWorkflow = new rootly.WorkflowIncident("sms-on-critical", {
    name: "SMS on-call for critical incidents",
    description: "Send SMS when critical incident starts",
    triggerParams: {
        triggers: ["incident_created"],
        incidentConditionKind: "IS",
        incidentKinds: ["normal"],
        incidentConditionStatus: "IS",
        incidentStatuses: ["started"],
        severityIds: [critical.id],
    },
    enabled: true,
});

new rootly.WorkflowTaskSendSms("sms-task", {
    workflowId: smsWorkflow.id,
    taskParams: {
        phoneNumbers: ["+11231231234"],
        content: "Critical incident started",
    },
});
```

### On-Call Schedules and Escalation Policies

```typescript
const schedule = new rootly.Schedule("primary", {
    name: "Primary On-Call Schedule",
});

const weekdays = new rootly.ScheduleRotation("weekdays", {
    scheduleId: schedule.id,
    name: "Weekdays",
    activeAllWeek: false,
    activeDays: ["M", "T", "W", "R", "F"],
    activeTimeType: "custom",
    position: 1,
    scheduleRotationableAttributes: { handoffTime: "10:00" },
    scheduleRotationableType: "ScheduleDailyRotation",
    timeZone: "America/Toronto",
});

const policy = new rootly.EscalationPolicy("primary", {
    name: "Primary",
});

const defaultPath = new rootly.EscalationPath("default", {
    name: "Default",
    default: true,
    escalationPolicyId: policy.id,
});

new rootly.EscalationLevel("first", {
    escalationPolicyPathId: defaultPath.id,
    escalationPolicyId: policy.id,
    position: 1,
    notificationTargetParams: [
        { type: "schedule", id: schedule.id, teamMembers: "all" },
    ],
});
```

## Deploying

```bash
pulumi preview    # Preview changes
pulumi up         # Deploy resources
pulumi stack      # View current stack
```

## Resource Types

- **Incident Management**: Severities, Services, Teams, Incident Types, Incident Roles, SLAs
- **Catalog**: Catalogs, Catalog Entities, Catalog Properties
- **Workflow Automation**: Workflows, Workflow Tasks (50+ task types including Jira, Slack, PagerDuty, email, SMS, and more)
- **On-Call**: Schedules, Rotations, Escalation Policies, Escalation Paths
- **Alerts**: Alert Sources, Alert Groups, Alert Routes, Alert Urgencies
- **Configuration**: Custom Fields, Forms, Environments, Functionalities
- **Communications**: Communications Groups, Stages, Templates, Types
- **Integrations**: Webhooks, Secrets, Heartbeats, Edge Connectors
- **Access Control**: Roles, Permissions, API Keys
- **Dashboards**: Dashboard Panels, Status Pages

## Development

```bash
make update_provider   # Update upstream Terraform provider
make development       # Build provider + SDKs + lint
make build_nodejs      # Build Node.js SDK only
make lint_provider     # Lint provider code
make test              # Run integration tests (requires ROOTLY_API_TOKEN)
```

## Reference

For detailed API documentation, visit the [Pulumi Registry](https://www.pulumi.com/registry/packages/rootly/api-docs/).
