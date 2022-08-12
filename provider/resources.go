// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rootly

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/rootlyhq/pulumi-rootly/provider/pkg/version"
	rootly "github.com/rootlyhq/terraform-provider-rootly/provider"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "rootly"
	// modules:
	mainMod = "index" // the rootly module
)

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(rootly.New("pulumi")())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "rootly",
		DisplayName:       "Rootly",
		Publisher:         "Rootly",
		LogoURL:           "https://raw.githubusercontent.com/rootlyhq/pulumi-rootly/v0.0.3/rootly.svg",
		PluginDownloadURL: "https://github.com/rootlyhq/pulumi-rootly/releases/v${VERSION}",
		Description:       "A Pulumi package for creating and managing Rootly cloud resources.",
		Keywords:          []string{"pulumi", "rootly", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://rootly.com",
		Repository:        "https://github.com/rootlyhq/pulumi-rootly",
		GitHubOrg:         "rootlyhq",
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"rootly_cause": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cause"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"cause": {
						CSharpName: "rootly_cause",
					},
				},
			},
			"rootly_custom_field": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CustomField"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"custom_field": {
						CSharpName: "rootly_custom_field",
					},
				},
			},
			"rootly_custom_field_option": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CustomFieldOption"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"custom_field_option": {
						CSharpName: "rootly_custom_field_option",
					},
				},
			},
			"rootly_dashboard": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Dashboard"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"dashboard": {
						CSharpName: "rootly_dashboard",
					},
				},
			},
			"rootly_dashboard_panel": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DashboardPanel"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"dashboard_panel": {
						CSharpName: "rootly_dashboard_panel",
					},
				},
			},
			"rootly_environment": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Environment"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"environment": {
						CSharpName: "rootly_environment",
					},
				},
			},
			"rootly_functionality": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Functionality"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"functionality": {
						CSharpName: "rootly_functionality",
					},
				},
			},
			"rootly_incident_role": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IncidentRole"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"incident_role": {
						CSharpName: "rootly_incident_role",
					},
				},
			},
			"rootly_incident_type": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IncidentType"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"incident_type": {
						CSharpName: "rootly_incident_type",
					},
				},
			},
			"rootly_service": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Service"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"service": {
						CSharpName: "rootly_service",
					},
				},
			},
			"rootly_severity": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Severity"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"severity": {
						CSharpName: "rootly_severity",
					},
				},
			},
			"rootly_team": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"team": {
						CSharpName: "rootly_team",
					},
				},
			},
      "rootly_workflow_incident": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowIncident"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_incident": {
            CSharpName: "rootly_workflow_incident",
          },
        },
      },
      "rootly_workflow_action_item": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowActionItem"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_action_item": {
            CSharpName: "rootly_workflow_action_item",
          },
        },
      },
      "rootly_workflow_alert": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowAlert"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_alert": {
            CSharpName: "rootly_workflow_alert",
          },
        },
      },
      "rootly_workflow_pulse": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowPulse"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_pulse": {
            CSharpName: "rootly_workflow_pulse",
          },
        },
      },
      "rootly_workflow_post_mortem": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowPostMortem"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_post_mortem": {
            CSharpName: "rootly_workflow_post_mortem",
          },
        },
      },
      "rootly_workflow_task_add_action_item": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAddActionItem"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_add_action_item": {
            CSharpName: "rootly_workflow_task_add_action_item",
          },
        },
      },
      "rootly_workflow_task_add_role": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAddRole"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_add_role": {
            CSharpName: "rootly_workflow_task_add_role",
          },
        },
      },
      "rootly_workflow_task_add_slack_bookmark": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAddSlackBookmark"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_add_slack_bookmark": {
            CSharpName: "rootly_workflow_task_add_slack_bookmark",
          },
        },
      },
      "rootly_workflow_task_add_team": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAddTeam"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_add_team": {
            CSharpName: "rootly_workflow_task_add_team",
          },
        },
      },
      "rootly_workflow_task_add_to_timeline": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAddToTimeline"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_add_to_timeline": {
            CSharpName: "rootly_workflow_task_add_to_timeline",
          },
        },
      },
      "rootly_workflow_task_archive_slack_channels": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskArchiveSlackChannels"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_archive_slack_channels": {
            CSharpName: "rootly_workflow_task_archive_slack_channels",
          },
        },
      },
      "rootly_workflow_task_attach_datadog_dashboards": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAttachDatadogDashboards"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_attach_datadog_dashboards": {
            CSharpName: "rootly_workflow_task_attach_datadog_dashboards",
          },
        },
      },
      "rootly_workflow_task_auto_assign_role_opsgenie": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAutoAssignRoleOpsgenie"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_auto_assign_role_opsgenie": {
            CSharpName: "rootly_workflow_task_auto_assign_role_opsgenie",
          },
        },
      },
      "rootly_workflow_task_auto_assign_role_pagerduty": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAutoAssignRolePagerduty"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_auto_assign_role_pagerduty": {
            CSharpName: "rootly_workflow_task_auto_assign_role_pagerduty",
          },
        },
      },
      "rootly_workflow_task_auto_assign_role_victor_ops": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskAutoAssignRoleVictorOps"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_auto_assign_role_victor_ops": {
            CSharpName: "rootly_workflow_task_auto_assign_role_victor_ops",
          },
        },
      },
      "rootly_workflow_task_call_people": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCallPeople"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_call_people": {
            CSharpName: "rootly_workflow_task_call_people",
          },
        },
      },
      "rootly_workflow_task_create_airtable_table_record": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateAirtableTableRecord"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_airtable_table_record": {
            CSharpName: "rootly_workflow_task_create_airtable_table_record",
          },
        },
      },
      "rootly_workflow_task_create_asana_subtask": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateAsanaSubtask"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_asana_subtask": {
            CSharpName: "rootly_workflow_task_create_asana_subtask",
          },
        },
      },
      "rootly_workflow_task_create_asana_task": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateAsanaTask"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_asana_task": {
            CSharpName: "rootly_workflow_task_create_asana_task",
          },
        },
      },
      "rootly_workflow_task_create_confluence_page": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateConfluencePage"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_confluence_page": {
            CSharpName: "rootly_workflow_task_create_confluence_page",
          },
        },
      },
      "rootly_workflow_task_create_datadog_notebook": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateDatadogNotebook"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_datadog_notebook": {
            CSharpName: "rootly_workflow_task_create_datadog_notebook",
          },
        },
      },
      "rootly_workflow_task_create_dropbox_paper_page": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateDropboxPaperPage"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_dropbox_paper_page": {
            CSharpName: "rootly_workflow_task_create_dropbox_paper_page",
          },
        },
      },
      "rootly_workflow_task_create_github_issue": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateGithubIssue"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_github_issue": {
            CSharpName: "rootly_workflow_task_create_github_issue",
          },
        },
      },
      "rootly_workflow_task_create_google_calendar_event": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateGoogleCalendarEvent"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_google_calendar_event": {
            CSharpName: "rootly_workflow_task_create_google_calendar_event",
          },
        },
      },
      "rootly_workflow_task_update_google_calendar_event": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateGoogleCalendarEvent"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_google_calendar_event": {
            CSharpName: "rootly_workflow_task_update_google_calendar_event",
          },
        },
      },
      "rootly_workflow_task_create_google_docs_page": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateGoogleDocsPage"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_google_docs_page": {
            CSharpName: "rootly_workflow_task_create_google_docs_page",
          },
        },
      },
      "rootly_workflow_task_create_google_meeting": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateGoogleMeeting"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_google_meeting": {
            CSharpName: "rootly_workflow_task_create_google_meeting",
          },
        },
      },
      "rootly_workflow_task_create_incident": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateIncident"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_incident": {
            CSharpName: "rootly_workflow_task_create_incident",
          },
        },
      },
      "rootly_workflow_task_create_jira_issue": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateJiraIssue"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_jira_issue": {
            CSharpName: "rootly_workflow_task_create_jira_issue",
          },
        },
      },
      "rootly_workflow_task_create_jira_subtask": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateJiraSubtask"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_jira_subtask": {
            CSharpName: "rootly_workflow_task_create_jira_subtask",
          },
        },
      },
      "rootly_workflow_task_create_linear_issue": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateLinearIssue"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_linear_issue": {
            CSharpName: "rootly_workflow_task_create_linear_issue",
          },
        },
      },
      "rootly_workflow_task_create_linear_subtask_issue": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateLinearSubtaskIssue"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_linear_subtask_issue": {
            CSharpName: "rootly_workflow_task_create_linear_subtask_issue",
          },
        },
      },
      "rootly_workflow_task_create_microsoft_teams_meeting": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateMicrosoftTeamsMeeting"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_microsoft_teams_meeting": {
            CSharpName: "rootly_workflow_task_create_microsoft_teams_meeting",
          },
        },
      },
      "rootly_workflow_task_create_notion_page": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateNotionPage"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_notion_page": {
            CSharpName: "rootly_workflow_task_create_notion_page",
          },
        },
      },
      "rootly_workflow_task_create_service_now_incident": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateServiceNowIncident"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_service_now_incident": {
            CSharpName: "rootly_workflow_task_create_service_now_incident",
          },
        },
      },
      "rootly_workflow_task_create_shortcut_story": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateShortcutStory"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_shortcut_story": {
            CSharpName: "rootly_workflow_task_create_shortcut_story",
          },
        },
      },
      "rootly_workflow_task_create_shortcut_task": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateShortcutTask"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_shortcut_task": {
            CSharpName: "rootly_workflow_task_create_shortcut_task",
          },
        },
      },
      "rootly_workflow_task_create_trello_card": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateTrelloCard"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_trello_card": {
            CSharpName: "rootly_workflow_task_create_trello_card",
          },
        },
      },
      "rootly_workflow_task_create_webex_meeting": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateWebexMeeting"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_webex_meeting": {
            CSharpName: "rootly_workflow_task_create_webex_meeting",
          },
        },
      },
      "rootly_workflow_task_create_zendesk_ticket": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateZendeskTicket"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_zendesk_ticket": {
            CSharpName: "rootly_workflow_task_create_zendesk_ticket",
          },
        },
      },
      "rootly_workflow_task_create_zoom_meeting": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskCreateZoomMeeting"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_create_zoom_meeting": {
            CSharpName: "rootly_workflow_task_create_zoom_meeting",
          },
        },
      },
      "rootly_workflow_task_get_github_commits": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskGetGithubCommits"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_get_github_commits": {
            CSharpName: "rootly_workflow_task_get_github_commits",
          },
        },
      },
      "rootly_workflow_task_get_gitlab_commits": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskGetGitlabCommits"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_get_gitlab_commits": {
            CSharpName: "rootly_workflow_task_get_gitlab_commits",
          },
        },
      },
      "rootly_workflow_task_get_pulses": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskGetPulses"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_get_pulses": {
            CSharpName: "rootly_workflow_task_get_pulses",
          },
        },
      },
      "rootly_workflow_task_http_client": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskHttpClient"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_http_client": {
            CSharpName: "rootly_workflow_task_http_client",
          },
        },
      },
      "rootly_workflow_task_invite_to_slack_channel_opsgenie": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskInviteToSlackChannelOpsgenie"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_invite_to_slack_channel_opsgenie": {
            CSharpName: "rootly_workflow_task_invite_to_slack_channel_opsgenie",
          },
        },
      },
      "rootly_workflow_task_invite_to_slack_channel_pagerduty": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskInviteToSlackChannelPagerduty"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_invite_to_slack_channel_pagerduty": {
            CSharpName: "rootly_workflow_task_invite_to_slack_channel_pagerduty",
          },
        },
      },
      "rootly_workflow_task_invite_to_slack_channel": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskInviteToSlackChannel"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_invite_to_slack_channel": {
            CSharpName: "rootly_workflow_task_invite_to_slack_channel",
          },
        },
      },
      "rootly_workflow_task_invite_to_slack_channel_victor_ops": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskInviteToSlackChannelVictorOps"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_invite_to_slack_channel_victor_ops": {
            CSharpName: "rootly_workflow_task_invite_to_slack_channel_victor_ops",
          },
        },
      },
      "rootly_workflow_task_page_opsgenie_on_call_responders": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskPageOpsgenieOnCallResponders"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_page_opsgenie_on_call_responders": {
            CSharpName: "rootly_workflow_task_page_opsgenie_on_call_responders",
          },
        },
      },
      "rootly_workflow_task_page_pagerduty_on_call_responders": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskPagePagerdutyOnCallResponders"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_page_pagerduty_on_call_responders": {
            CSharpName: "rootly_workflow_task_page_pagerduty_on_call_responders",
          },
        },
      },
      "rootly_workflow_task_page_victor_ops_on_call_responders": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskPageVictorOpsOnCallResponders"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_page_victor_ops_on_call_responders": {
            CSharpName: "rootly_workflow_task_page_victor_ops_on_call_responders",
          },
        },
      },
      "rootly_workflow_task_print": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskPrint"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_print": {
            CSharpName: "rootly_workflow_task_print",
          },
        },
      },
      "rootly_workflow_task_publish_incident": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskPublishIncident"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_publish_incident": {
            CSharpName: "rootly_workflow_task_publish_incident",
          },
        },
      },
      "rootly_workflow_task_redis_client": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskRedisClient"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_redis_client": {
            CSharpName: "rootly_workflow_task_redis_client",
          },
        },
      },
      "rootly_workflow_task_rename_slack_channel": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskRenameSlackChannel"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_rename_slack_channel": {
            CSharpName: "rootly_workflow_task_rename_slack_channel",
          },
        },
      },
      "rootly_workflow_task_run_command_heroku": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskRunCommandHeroku"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_run_command_heroku": {
            CSharpName: "rootly_workflow_task_run_command_heroku",
          },
        },
      },
      "rootly_workflow_task_send_email": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskSendEmail"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_send_email": {
            CSharpName: "rootly_workflow_task_send_email",
          },
        },
      },
      "rootly_workflow_task_send_slack_message": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskSendSlackMessage"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_send_slack_message": {
            CSharpName: "rootly_workflow_task_send_slack_message",
          },
        },
      },
      "rootly_workflow_task_send_sms": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskSendSms"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_send_sms": {
            CSharpName: "rootly_workflow_task_send_sms",
          },
        },
      },
      "rootly_workflow_task_snapshot_datadog_graph": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskSnapshotDatadogGraph"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_snapshot_datadog_graph": {
            CSharpName: "rootly_workflow_task_snapshot_datadog_graph",
          },
        },
      },
      "rootly_workflow_task_snapshot_grafana_dashboard": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskSnapshotGrafanaDashboard"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_snapshot_grafana_dashboard": {
            CSharpName: "rootly_workflow_task_snapshot_grafana_dashboard",
          },
        },
      },
      "rootly_workflow_task_snapshot_looker_look": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskSnapshotLookerLook"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_snapshot_looker_look": {
            CSharpName: "rootly_workflow_task_snapshot_looker_look",
          },
        },
      },
      "rootly_workflow_task_snapshot_new_relic_graph": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskSnapshotNewRelicGraph"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_snapshot_new_relic_graph": {
            CSharpName: "rootly_workflow_task_snapshot_new_relic_graph",
          },
        },
      },
      "rootly_workflow_task_tweet_twitter_message": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskTweetTwitterMessage"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_tweet_twitter_message": {
            CSharpName: "rootly_workflow_task_tweet_twitter_message",
          },
        },
      },
      "rootly_workflow_task_update_airtable_table_record": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateAirtableTableRecord"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_airtable_table_record": {
            CSharpName: "rootly_workflow_task_update_airtable_table_record",
          },
        },
      },
      "rootly_workflow_task_update_asana_task": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateAsanaTask"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_asana_task": {
            CSharpName: "rootly_workflow_task_update_asana_task",
          },
        },
      },
      "rootly_workflow_task_update_github_issue": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateGithubIssue"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_github_issue": {
            CSharpName: "rootly_workflow_task_update_github_issue",
          },
        },
      },
      "rootly_workflow_task_update_incident": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateIncident"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_incident": {
            CSharpName: "rootly_workflow_task_update_incident",
          },
        },
      },
      "rootly_workflow_task_update_jira_issue": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateJiraIssue"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_jira_issue": {
            CSharpName: "rootly_workflow_task_update_jira_issue",
          },
        },
      },
      "rootly_workflow_task_update_linear_issue": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateLinearIssue"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_linear_issue": {
            CSharpName: "rootly_workflow_task_update_linear_issue",
          },
        },
      },
      "rootly_workflow_task_update_service_now_incident": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateServiceNowIncident"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_service_now_incident": {
            CSharpName: "rootly_workflow_task_update_service_now_incident",
          },
        },
      },
      "rootly_workflow_task_update_shortcut_story": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateShortcutStory"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_shortcut_story": {
            CSharpName: "rootly_workflow_task_update_shortcut_story",
          },
        },
      },
      "rootly_workflow_task_update_shortcut_task": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateShortcutTask"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_shortcut_task": {
            CSharpName: "rootly_workflow_task_update_shortcut_task",
          },
        },
      },
      "rootly_workflow_task_update_slack_channel_topic": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateSlackChannelTopic"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_slack_channel_topic": {
            CSharpName: "rootly_workflow_task_update_slack_channel_topic",
          },
        },
      },
      "rootly_workflow_task_update_status": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateStatus"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_status": {
            CSharpName: "rootly_workflow_task_update_status",
          },
        },
      },
      "rootly_workflow_task_update_trello_card": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateTrelloCard"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_trello_card": {
            CSharpName: "rootly_workflow_task_update_trello_card",
          },
        },
      },
      "rootly_workflow_task_update_zendesk_ticket": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "WorkflowTaskUpdateZendeskTicket"),
        Fields: map[string]*tfbridge.SchemaInfo{
          "workflow_task_update_zendesk_ticket": {
            CSharpName: "rootly_workflow_task_update_zendesk_ticket",
          },
        },
      },
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			"rootly_causes": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCauses")},
			"rootly_custom_fields": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCustomFields")},
			"rootly_custom_field_options": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCustomFieldOptions")},
			"rootly_environments": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEnvironments")},
			"rootly_functionalities": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFunctionalities")},
			"rootly_incident_types": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIncidentTypes")},
			"rootly_incident_roles": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIncidentRoles")},
			"rootly_services": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getServices")},
			"rootly_severities": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSeverities")},
			"rootly_teams": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTeams")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
