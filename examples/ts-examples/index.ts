import * as rootly from "@rootly/pulumi";

// -----------------------------------------------------------------------------
// Severities
// -----------------------------------------------------------------------------
const sev0 = new rootly.Severity("sev0", {
    name: "SEV0",
    color: "#FF0000",
    notifyEmails: ["oncall@acme.com"],
    slackChannels: [{ id: "C06A4RZR9", name: "incidents-critical" }],
});

const sev1 = new rootly.Severity("sev1", {
    name: "SEV1",
    color: "#FFA500",
    notifyEmails: ["oncall@acme.com"],
    slackChannels: [{ id: "C06A4RZR9", name: "incidents-critical" }],
});

const sev2 = new rootly.Severity("sev2", {
    name: "SEV2",
    color: "#FFA500",
});

// -----------------------------------------------------------------------------
// Environments
// -----------------------------------------------------------------------------
const production = new rootly.Environment("production", {
    name: "Production",
    color: "#FF0000",
    notifyEmails: ["oncall@acme.com"],
    slackChannels: [{ id: "C06A4RZR9", name: "incidents-prod" }],
});

const staging = new rootly.Environment("staging", {
    name: "Staging",
    color: "#FFA500",
});

// -----------------------------------------------------------------------------
// Services
// -----------------------------------------------------------------------------
const api = new rootly.Service("api", {
    name: "api-gateway",
    color: "#800080",
    notifyEmails: ["platform@acme.com"],
    slackChannels: [{ id: "C06A4RZR9", name: "platform-alerts" }],
});

const database = new rootly.Service("database", {
    name: "customer-postgresql",
    color: "#008000",
    notifyEmails: ["data@acme.com"],
});

// -----------------------------------------------------------------------------
// Functionalities
// -----------------------------------------------------------------------------
const checkout = new rootly.Functionality("checkout", {
    name: "Checkout Flow",
    color: "#800080",
});

const login = new rootly.Functionality("login", {
    name: "User Login",
    color: "#0000FF",
});

// -----------------------------------------------------------------------------
// Teams
// -----------------------------------------------------------------------------
const sre = new rootly.Team("sre", {
    name: "SRE",
});

// -----------------------------------------------------------------------------
// Custom form fields
// -----------------------------------------------------------------------------
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

new rootly.FormFieldOption("north-america", {
    formFieldId: regionsAffected.id,
    value: "North America",
});

// -----------------------------------------------------------------------------
// Dashboards
// -----------------------------------------------------------------------------
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

// -----------------------------------------------------------------------------
// Workflow: Create Jira issue on incident
// -----------------------------------------------------------------------------
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

// -----------------------------------------------------------------------------
// Workflow: Notify Slack channels
// -----------------------------------------------------------------------------
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

// -----------------------------------------------------------------------------
// Workflow: Page PagerDuty responders
// -----------------------------------------------------------------------------
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

// -----------------------------------------------------------------------------
// Workflow: Send SMS for critical incidents (using data source)
// -----------------------------------------------------------------------------
const criticalSeverity = rootly.getSeverityOutput({ slug: "sev0" });

const smsWorkflow = new rootly.WorkflowIncident("sms-on-critical", {
    name: "SMS on-call for critical incidents",
    description: "Send SMS when critical incident starts",
    triggerParams: {
        triggers: ["incident_created"],
        incidentConditionKind: "IS",
        incidentKinds: ["normal"],
        incidentConditionStatus: "IS",
        incidentStatuses: ["started"],
        severityIds: [criticalSeverity.id],
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

// -----------------------------------------------------------------------------
// Workflow: Auto-archive Slack channels after resolution
// -----------------------------------------------------------------------------
const archiveWorkflow = new rootly.WorkflowIncident("auto-archive", {
    name: "Archive Slack channel after resolution",
    description: "Automatically archive incident Slack channel 30 days after resolution",
    triggerParams: {
        triggers: ["incident_updated"],
        incidentConditionStatus: "IS",
        incidentStatuses: ["resolved"],
    },
    enabled: true,
});

new rootly.WorkflowTaskArchiveSlackChannels("archive-task", {
    workflowId: archiveWorkflow.id,
    taskParams: {
        channels: [{
            id: "{{ incident.slack_channel_id }}",
            name: "{{ incident.slack_channel_id }}",
        }],
    },
});

// -----------------------------------------------------------------------------
// Workflow: Update status page
// -----------------------------------------------------------------------------
const statusPageWorkflow = new rootly.WorkflowIncident("publish-to-status-page", {
    name: "Publish incident to status page",
    description: "Publish incident when started",
    triggerParams: {
        triggers: ["incident_created"],
        incidentConditionStatus: "IS",
        incidentStatuses: ["started"],
    },
    enabled: true,
});

new rootly.WorkflowTaskPublishIncident("publish-task", {
    workflowId: statusPageWorkflow.id,
    taskParams: {},
});

// -----------------------------------------------------------------------------
// Exports
// -----------------------------------------------------------------------------
export const severityIds = {
    sev0: sev0.id,
    sev1: sev1.id,
    sev2: sev2.id,
};

export const serviceIds = {
    api: api.id,
    database: database.id,
};
