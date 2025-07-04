# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EscalationPolicyArgs', 'EscalationPolicy']

@pulumi.input_type
class EscalationPolicyArgs:
    def __init__(__self__, *,
                 created_by_user_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 last_updated_by_user_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_count: Optional[pulumi.Input[int]] = None,
                 service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a EscalationPolicy resource.
        :param pulumi.Input[int] created_by_user_id: User who created the escalation policy
        :param pulumi.Input[str] description: The description of the escalation policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_ids: Associated groups (alerting the group will trigger escalation policy)
        :param pulumi.Input[int] last_updated_by_user_id: User who updated the escalation policy
        :param pulumi.Input[str] name: The name of the escalation policy
        :param pulumi.Input[int] repeat_count: The number of times this policy will be executed until someone acknowledges the alert
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_ids: Associated services (alerting the service will trigger escalation policy)
        """
        if created_by_user_id is not None:
            pulumi.set(__self__, "created_by_user_id", created_by_user_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if last_updated_by_user_id is not None:
            pulumi.set(__self__, "last_updated_by_user_id", last_updated_by_user_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repeat_count is not None:
            pulumi.set(__self__, "repeat_count", repeat_count)
        if service_ids is not None:
            pulumi.set(__self__, "service_ids", service_ids)

    @property
    @pulumi.getter(name="createdByUserId")
    def created_by_user_id(self) -> Optional[pulumi.Input[int]]:
        """
        User who created the escalation policy
        """
        return pulumi.get(self, "created_by_user_id")

    @created_by_user_id.setter
    def created_by_user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "created_by_user_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the escalation policy
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Associated groups (alerting the group will trigger escalation policy)
        """
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter(name="lastUpdatedByUserId")
    def last_updated_by_user_id(self) -> Optional[pulumi.Input[int]]:
        """
        User who updated the escalation policy
        """
        return pulumi.get(self, "last_updated_by_user_id")

    @last_updated_by_user_id.setter
    def last_updated_by_user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_updated_by_user_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the escalation policy
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="repeatCount")
    def repeat_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of times this policy will be executed until someone acknowledges the alert
        """
        return pulumi.get(self, "repeat_count")

    @repeat_count.setter
    def repeat_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat_count", value)

    @property
    @pulumi.getter(name="serviceIds")
    def service_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Associated services (alerting the service will trigger escalation policy)
        """
        return pulumi.get(self, "service_ids")

    @service_ids.setter
    def service_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "service_ids", value)


@pulumi.input_type
class _EscalationPolicyState:
    def __init__(__self__, *,
                 created_by_user_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 last_updated_by_user_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_count: Optional[pulumi.Input[int]] = None,
                 service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering EscalationPolicy resources.
        :param pulumi.Input[int] created_by_user_id: User who created the escalation policy
        :param pulumi.Input[str] description: The description of the escalation policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_ids: Associated groups (alerting the group will trigger escalation policy)
        :param pulumi.Input[int] last_updated_by_user_id: User who updated the escalation policy
        :param pulumi.Input[str] name: The name of the escalation policy
        :param pulumi.Input[int] repeat_count: The number of times this policy will be executed until someone acknowledges the alert
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_ids: Associated services (alerting the service will trigger escalation policy)
        """
        if created_by_user_id is not None:
            pulumi.set(__self__, "created_by_user_id", created_by_user_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if last_updated_by_user_id is not None:
            pulumi.set(__self__, "last_updated_by_user_id", last_updated_by_user_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repeat_count is not None:
            pulumi.set(__self__, "repeat_count", repeat_count)
        if service_ids is not None:
            pulumi.set(__self__, "service_ids", service_ids)

    @property
    @pulumi.getter(name="createdByUserId")
    def created_by_user_id(self) -> Optional[pulumi.Input[int]]:
        """
        User who created the escalation policy
        """
        return pulumi.get(self, "created_by_user_id")

    @created_by_user_id.setter
    def created_by_user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "created_by_user_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the escalation policy
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Associated groups (alerting the group will trigger escalation policy)
        """
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter(name="lastUpdatedByUserId")
    def last_updated_by_user_id(self) -> Optional[pulumi.Input[int]]:
        """
        User who updated the escalation policy
        """
        return pulumi.get(self, "last_updated_by_user_id")

    @last_updated_by_user_id.setter
    def last_updated_by_user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_updated_by_user_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the escalation policy
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="repeatCount")
    def repeat_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of times this policy will be executed until someone acknowledges the alert
        """
        return pulumi.get(self, "repeat_count")

    @repeat_count.setter
    def repeat_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat_count", value)

    @property
    @pulumi.getter(name="serviceIds")
    def service_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Associated services (alerting the service will trigger escalation policy)
        """
        return pulumi.get(self, "service_ids")

    @service_ids.setter
    def service_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "service_ids", value)


class EscalationPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created_by_user_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 last_updated_by_user_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_count: Optional[pulumi.Input[int]] = None,
                 service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a EscalationPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] created_by_user_id: User who created the escalation policy
        :param pulumi.Input[str] description: The description of the escalation policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_ids: Associated groups (alerting the group will trigger escalation policy)
        :param pulumi.Input[int] last_updated_by_user_id: User who updated the escalation policy
        :param pulumi.Input[str] name: The name of the escalation policy
        :param pulumi.Input[int] repeat_count: The number of times this policy will be executed until someone acknowledges the alert
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_ids: Associated services (alerting the service will trigger escalation policy)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EscalationPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EscalationPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EscalationPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EscalationPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created_by_user_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 last_updated_by_user_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repeat_count: Optional[pulumi.Input[int]] = None,
                 service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EscalationPolicyArgs.__new__(EscalationPolicyArgs)

            __props__.__dict__["created_by_user_id"] = created_by_user_id
            __props__.__dict__["description"] = description
            __props__.__dict__["group_ids"] = group_ids
            __props__.__dict__["last_updated_by_user_id"] = last_updated_by_user_id
            __props__.__dict__["name"] = name
            __props__.__dict__["repeat_count"] = repeat_count
            __props__.__dict__["service_ids"] = service_ids
        super(EscalationPolicy, __self__).__init__(
            'rootly:index/escalationPolicy:EscalationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_by_user_id: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            last_updated_by_user_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            repeat_count: Optional[pulumi.Input[int]] = None,
            service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'EscalationPolicy':
        """
        Get an existing EscalationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] created_by_user_id: User who created the escalation policy
        :param pulumi.Input[str] description: The description of the escalation policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_ids: Associated groups (alerting the group will trigger escalation policy)
        :param pulumi.Input[int] last_updated_by_user_id: User who updated the escalation policy
        :param pulumi.Input[str] name: The name of the escalation policy
        :param pulumi.Input[int] repeat_count: The number of times this policy will be executed until someone acknowledges the alert
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_ids: Associated services (alerting the service will trigger escalation policy)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EscalationPolicyState.__new__(_EscalationPolicyState)

        __props__.__dict__["created_by_user_id"] = created_by_user_id
        __props__.__dict__["description"] = description
        __props__.__dict__["group_ids"] = group_ids
        __props__.__dict__["last_updated_by_user_id"] = last_updated_by_user_id
        __props__.__dict__["name"] = name
        __props__.__dict__["repeat_count"] = repeat_count
        __props__.__dict__["service_ids"] = service_ids
        return EscalationPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdByUserId")
    def created_by_user_id(self) -> pulumi.Output[int]:
        """
        User who created the escalation policy
        """
        return pulumi.get(self, "created_by_user_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the escalation policy
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Associated groups (alerting the group will trigger escalation policy)
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter(name="lastUpdatedByUserId")
    def last_updated_by_user_id(self) -> pulumi.Output[int]:
        """
        User who updated the escalation policy
        """
        return pulumi.get(self, "last_updated_by_user_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the escalation policy
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="repeatCount")
    def repeat_count(self) -> pulumi.Output[int]:
        """
        The number of times this policy will be executed until someone acknowledges the alert
        """
        return pulumi.get(self, "repeat_count")

    @property
    @pulumi.getter(name="serviceIds")
    def service_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Associated services (alerting the service will trigger escalation policy)
        """
        return pulumi.get(self, "service_ids")

