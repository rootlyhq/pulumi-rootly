# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TeamArgs', 'Team']

@pulumi.input_type
class TeamArgs:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 opsgenie_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_service_id: Optional[pulumi.Input[str]] = None,
                 pagertree_id: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 slack_aliases: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]]] = None,
                 slack_channels: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 victor_ops_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Team resource.
        :param pulumi.Input[str] color: The hex color of the team
        :param pulumi.Input[str] description: The description of the team
        :param pulumi.Input[str] name: The name of the team
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notify_emails: Emails to attach to the team
        :param pulumi.Input[str] opsgenie_id: The Opsgenie group id associated to this team
        :param pulumi.Input[str] pagerduty_id: The PagerDuty group id associated to this team
        :param pulumi.Input[str] pagerduty_service_id: The PagerDuty service id associated to this team
        :param pulumi.Input[str] pagertree_id: The PagerTree group id associated to this team
        :param pulumi.Input[int] position: Position of the team
        :param pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]] slack_aliases: Slack Aliases associated with this service
        :param pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]] slack_channels: Slack Channels associated with this service
        :param pulumi.Input[Sequence[pulumi.Input[int]]] user_ids: The User ID's members of this team
        :param pulumi.Input[str] victor_ops_id: The VictorOps group id associated to this team
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notify_emails is not None:
            pulumi.set(__self__, "notify_emails", notify_emails)
        if opsgenie_id is not None:
            pulumi.set(__self__, "opsgenie_id", opsgenie_id)
        if pagerduty_id is not None:
            pulumi.set(__self__, "pagerduty_id", pagerduty_id)
        if pagerduty_service_id is not None:
            pulumi.set(__self__, "pagerduty_service_id", pagerduty_service_id)
        if pagertree_id is not None:
            pulumi.set(__self__, "pagertree_id", pagertree_id)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if slack_aliases is not None:
            pulumi.set(__self__, "slack_aliases", slack_aliases)
        if slack_channels is not None:
            pulumi.set(__self__, "slack_channels", slack_channels)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if user_ids is not None:
            pulumi.set(__self__, "user_ids", user_ids)
        if victor_ops_id is not None:
            pulumi.set(__self__, "victor_ops_id", victor_ops_id)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        The hex color of the team
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the team
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the team
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notifyEmails")
    def notify_emails(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Emails to attach to the team
        """
        return pulumi.get(self, "notify_emails")

    @notify_emails.setter
    def notify_emails(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notify_emails", value)

    @property
    @pulumi.getter(name="opsgenieId")
    def opsgenie_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Opsgenie group id associated to this team
        """
        return pulumi.get(self, "opsgenie_id")

    @opsgenie_id.setter
    def opsgenie_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "opsgenie_id", value)

    @property
    @pulumi.getter(name="pagerdutyId")
    def pagerduty_id(self) -> Optional[pulumi.Input[str]]:
        """
        The PagerDuty group id associated to this team
        """
        return pulumi.get(self, "pagerduty_id")

    @pagerduty_id.setter
    def pagerduty_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pagerduty_id", value)

    @property
    @pulumi.getter(name="pagerdutyServiceId")
    def pagerduty_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The PagerDuty service id associated to this team
        """
        return pulumi.get(self, "pagerduty_service_id")

    @pagerduty_service_id.setter
    def pagerduty_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pagerduty_service_id", value)

    @property
    @pulumi.getter(name="pagertreeId")
    def pagertree_id(self) -> Optional[pulumi.Input[str]]:
        """
        The PagerTree group id associated to this team
        """
        return pulumi.get(self, "pagertree_id")

    @pagertree_id.setter
    def pagertree_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pagertree_id", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Position of the team
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="slackAliases")
    def slack_aliases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]]]:
        """
        Slack Aliases associated with this service
        """
        return pulumi.get(self, "slack_aliases")

    @slack_aliases.setter
    def slack_aliases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]]]):
        pulumi.set(self, "slack_aliases", value)

    @property
    @pulumi.getter(name="slackChannels")
    def slack_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]]]:
        """
        Slack Channels associated with this service
        """
        return pulumi.get(self, "slack_channels")

    @slack_channels.setter
    def slack_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]]]):
        pulumi.set(self, "slack_channels", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        The User ID's members of this team
        """
        return pulumi.get(self, "user_ids")

    @user_ids.setter
    def user_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "user_ids", value)

    @property
    @pulumi.getter(name="victorOpsId")
    def victor_ops_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VictorOps group id associated to this team
        """
        return pulumi.get(self, "victor_ops_id")

    @victor_ops_id.setter
    def victor_ops_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "victor_ops_id", value)


@pulumi.input_type
class _TeamState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 opsgenie_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_service_id: Optional[pulumi.Input[str]] = None,
                 pagertree_id: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 slack_aliases: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]]] = None,
                 slack_channels: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 victor_ops_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Team resources.
        :param pulumi.Input[str] color: The hex color of the team
        :param pulumi.Input[str] description: The description of the team
        :param pulumi.Input[str] name: The name of the team
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notify_emails: Emails to attach to the team
        :param pulumi.Input[str] opsgenie_id: The Opsgenie group id associated to this team
        :param pulumi.Input[str] pagerduty_id: The PagerDuty group id associated to this team
        :param pulumi.Input[str] pagerduty_service_id: The PagerDuty service id associated to this team
        :param pulumi.Input[str] pagertree_id: The PagerTree group id associated to this team
        :param pulumi.Input[int] position: Position of the team
        :param pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]] slack_aliases: Slack Aliases associated with this service
        :param pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]] slack_channels: Slack Channels associated with this service
        :param pulumi.Input[Sequence[pulumi.Input[int]]] user_ids: The User ID's members of this team
        :param pulumi.Input[str] victor_ops_id: The VictorOps group id associated to this team
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notify_emails is not None:
            pulumi.set(__self__, "notify_emails", notify_emails)
        if opsgenie_id is not None:
            pulumi.set(__self__, "opsgenie_id", opsgenie_id)
        if pagerduty_id is not None:
            pulumi.set(__self__, "pagerduty_id", pagerduty_id)
        if pagerduty_service_id is not None:
            pulumi.set(__self__, "pagerduty_service_id", pagerduty_service_id)
        if pagertree_id is not None:
            pulumi.set(__self__, "pagertree_id", pagertree_id)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if slack_aliases is not None:
            pulumi.set(__self__, "slack_aliases", slack_aliases)
        if slack_channels is not None:
            pulumi.set(__self__, "slack_channels", slack_channels)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if user_ids is not None:
            pulumi.set(__self__, "user_ids", user_ids)
        if victor_ops_id is not None:
            pulumi.set(__self__, "victor_ops_id", victor_ops_id)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        The hex color of the team
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the team
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the team
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notifyEmails")
    def notify_emails(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Emails to attach to the team
        """
        return pulumi.get(self, "notify_emails")

    @notify_emails.setter
    def notify_emails(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notify_emails", value)

    @property
    @pulumi.getter(name="opsgenieId")
    def opsgenie_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Opsgenie group id associated to this team
        """
        return pulumi.get(self, "opsgenie_id")

    @opsgenie_id.setter
    def opsgenie_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "opsgenie_id", value)

    @property
    @pulumi.getter(name="pagerdutyId")
    def pagerduty_id(self) -> Optional[pulumi.Input[str]]:
        """
        The PagerDuty group id associated to this team
        """
        return pulumi.get(self, "pagerduty_id")

    @pagerduty_id.setter
    def pagerduty_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pagerduty_id", value)

    @property
    @pulumi.getter(name="pagerdutyServiceId")
    def pagerduty_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The PagerDuty service id associated to this team
        """
        return pulumi.get(self, "pagerduty_service_id")

    @pagerduty_service_id.setter
    def pagerduty_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pagerduty_service_id", value)

    @property
    @pulumi.getter(name="pagertreeId")
    def pagertree_id(self) -> Optional[pulumi.Input[str]]:
        """
        The PagerTree group id associated to this team
        """
        return pulumi.get(self, "pagertree_id")

    @pagertree_id.setter
    def pagertree_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pagertree_id", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Position of the team
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="slackAliases")
    def slack_aliases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]]]:
        """
        Slack Aliases associated with this service
        """
        return pulumi.get(self, "slack_aliases")

    @slack_aliases.setter
    def slack_aliases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackAliasArgs']]]]):
        pulumi.set(self, "slack_aliases", value)

    @property
    @pulumi.getter(name="slackChannels")
    def slack_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]]]:
        """
        Slack Channels associated with this service
        """
        return pulumi.get(self, "slack_channels")

    @slack_channels.setter
    def slack_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSlackChannelArgs']]]]):
        pulumi.set(self, "slack_channels", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        The User ID's members of this team
        """
        return pulumi.get(self, "user_ids")

    @user_ids.setter
    def user_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "user_ids", value)

    @property
    @pulumi.getter(name="victorOpsId")
    def victor_ops_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VictorOps group id associated to this team
        """
        return pulumi.get(self, "victor_ops_id")

    @victor_ops_id.setter
    def victor_ops_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "victor_ops_id", value)


class Team(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 opsgenie_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_service_id: Optional[pulumi.Input[str]] = None,
                 pagertree_id: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 slack_aliases: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackAliasArgs', 'TeamSlackAliasArgsDict']]]]] = None,
                 slack_channels: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackChannelArgs', 'TeamSlackChannelArgsDict']]]]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 victor_ops_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Team resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: The hex color of the team
        :param pulumi.Input[str] description: The description of the team
        :param pulumi.Input[str] name: The name of the team
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notify_emails: Emails to attach to the team
        :param pulumi.Input[str] opsgenie_id: The Opsgenie group id associated to this team
        :param pulumi.Input[str] pagerduty_id: The PagerDuty group id associated to this team
        :param pulumi.Input[str] pagerduty_service_id: The PagerDuty service id associated to this team
        :param pulumi.Input[str] pagertree_id: The PagerTree group id associated to this team
        :param pulumi.Input[int] position: Position of the team
        :param pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackAliasArgs', 'TeamSlackAliasArgsDict']]]] slack_aliases: Slack Aliases associated with this service
        :param pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackChannelArgs', 'TeamSlackChannelArgsDict']]]] slack_channels: Slack Channels associated with this service
        :param pulumi.Input[Sequence[pulumi.Input[int]]] user_ids: The User ID's members of this team
        :param pulumi.Input[str] victor_ops_id: The VictorOps group id associated to this team
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TeamArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Team resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TeamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 opsgenie_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_id: Optional[pulumi.Input[str]] = None,
                 pagerduty_service_id: Optional[pulumi.Input[str]] = None,
                 pagertree_id: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 slack_aliases: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackAliasArgs', 'TeamSlackAliasArgsDict']]]]] = None,
                 slack_channels: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackChannelArgs', 'TeamSlackChannelArgsDict']]]]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 victor_ops_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamArgs.__new__(TeamArgs)

            __props__.__dict__["color"] = color
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["notify_emails"] = notify_emails
            __props__.__dict__["opsgenie_id"] = opsgenie_id
            __props__.__dict__["pagerduty_id"] = pagerduty_id
            __props__.__dict__["pagerduty_service_id"] = pagerduty_service_id
            __props__.__dict__["pagertree_id"] = pagertree_id
            __props__.__dict__["position"] = position
            __props__.__dict__["slack_aliases"] = slack_aliases
            __props__.__dict__["slack_channels"] = slack_channels
            __props__.__dict__["slug"] = slug
            __props__.__dict__["user_ids"] = user_ids
            __props__.__dict__["victor_ops_id"] = victor_ops_id
        super(Team, __self__).__init__(
            'rootly:index/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notify_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            opsgenie_id: Optional[pulumi.Input[str]] = None,
            pagerduty_id: Optional[pulumi.Input[str]] = None,
            pagerduty_service_id: Optional[pulumi.Input[str]] = None,
            pagertree_id: Optional[pulumi.Input[str]] = None,
            position: Optional[pulumi.Input[int]] = None,
            slack_aliases: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackAliasArgs', 'TeamSlackAliasArgsDict']]]]] = None,
            slack_channels: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackChannelArgs', 'TeamSlackChannelArgsDict']]]]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            victor_ops_id: Optional[pulumi.Input[str]] = None) -> 'Team':
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: The hex color of the team
        :param pulumi.Input[str] description: The description of the team
        :param pulumi.Input[str] name: The name of the team
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notify_emails: Emails to attach to the team
        :param pulumi.Input[str] opsgenie_id: The Opsgenie group id associated to this team
        :param pulumi.Input[str] pagerduty_id: The PagerDuty group id associated to this team
        :param pulumi.Input[str] pagerduty_service_id: The PagerDuty service id associated to this team
        :param pulumi.Input[str] pagertree_id: The PagerTree group id associated to this team
        :param pulumi.Input[int] position: Position of the team
        :param pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackAliasArgs', 'TeamSlackAliasArgsDict']]]] slack_aliases: Slack Aliases associated with this service
        :param pulumi.Input[Sequence[pulumi.Input[Union['TeamSlackChannelArgs', 'TeamSlackChannelArgsDict']]]] slack_channels: Slack Channels associated with this service
        :param pulumi.Input[Sequence[pulumi.Input[int]]] user_ids: The User ID's members of this team
        :param pulumi.Input[str] victor_ops_id: The VictorOps group id associated to this team
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamState.__new__(_TeamState)

        __props__.__dict__["color"] = color
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["notify_emails"] = notify_emails
        __props__.__dict__["opsgenie_id"] = opsgenie_id
        __props__.__dict__["pagerduty_id"] = pagerduty_id
        __props__.__dict__["pagerduty_service_id"] = pagerduty_service_id
        __props__.__dict__["pagertree_id"] = pagertree_id
        __props__.__dict__["position"] = position
        __props__.__dict__["slack_aliases"] = slack_aliases
        __props__.__dict__["slack_channels"] = slack_channels
        __props__.__dict__["slug"] = slug
        __props__.__dict__["user_ids"] = user_ids
        __props__.__dict__["victor_ops_id"] = victor_ops_id
        return Team(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[str]:
        """
        The hex color of the team
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the team
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the team
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyEmails")
    def notify_emails(self) -> pulumi.Output[Sequence[str]]:
        """
        Emails to attach to the team
        """
        return pulumi.get(self, "notify_emails")

    @property
    @pulumi.getter(name="opsgenieId")
    def opsgenie_id(self) -> pulumi.Output[str]:
        """
        The Opsgenie group id associated to this team
        """
        return pulumi.get(self, "opsgenie_id")

    @property
    @pulumi.getter(name="pagerdutyId")
    def pagerduty_id(self) -> pulumi.Output[str]:
        """
        The PagerDuty group id associated to this team
        """
        return pulumi.get(self, "pagerduty_id")

    @property
    @pulumi.getter(name="pagerdutyServiceId")
    def pagerduty_service_id(self) -> pulumi.Output[str]:
        """
        The PagerDuty service id associated to this team
        """
        return pulumi.get(self, "pagerduty_service_id")

    @property
    @pulumi.getter(name="pagertreeId")
    def pagertree_id(self) -> pulumi.Output[str]:
        """
        The PagerTree group id associated to this team
        """
        return pulumi.get(self, "pagertree_id")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[int]:
        """
        Position of the team
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter(name="slackAliases")
    def slack_aliases(self) -> pulumi.Output[Sequence['outputs.TeamSlackAlias']]:
        """
        Slack Aliases associated with this service
        """
        return pulumi.get(self, "slack_aliases")

    @property
    @pulumi.getter(name="slackChannels")
    def slack_channels(self) -> pulumi.Output[Sequence['outputs.TeamSlackChannel']]:
        """
        Slack Channels associated with this service
        """
        return pulumi.get(self, "slack_channels")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> pulumi.Output[Sequence[int]]:
        """
        The User ID's members of this team
        """
        return pulumi.get(self, "user_ids")

    @property
    @pulumi.getter(name="victorOpsId")
    def victor_ops_id(self) -> pulumi.Output[str]:
        """
        The VictorOps group id associated to this team
        """
        return pulumi.get(self, "victor_ops_id")

