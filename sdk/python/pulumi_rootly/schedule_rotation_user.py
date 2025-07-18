# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ScheduleRotationUserArgs', 'ScheduleRotationUser']

@pulumi.input_type
class ScheduleRotationUserArgs:
    def __init__(__self__, *,
                 user_id: pulumi.Input[int],
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ScheduleRotationUser resource.
        :param pulumi.Input[int] user_id: Schedule rotation user
        :param pulumi.Input[int] position: Position of the user inside rotation
        """
        pulumi.set(__self__, "user_id", user_id)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if schedule_rotation_id is not None:
            pulumi.set(__self__, "schedule_rotation_id", schedule_rotation_id)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[int]:
        """
        Schedule rotation user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Position of the user inside rotation
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="scheduleRotationId")
    def schedule_rotation_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule_rotation_id")

    @schedule_rotation_id.setter
    def schedule_rotation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_rotation_id", value)


@pulumi.input_type
class _ScheduleRotationUserState:
    def __init__(__self__, *,
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ScheduleRotationUser resources.
        :param pulumi.Input[int] position: Position of the user inside rotation
        :param pulumi.Input[int] user_id: Schedule rotation user
        """
        if position is not None:
            pulumi.set(__self__, "position", position)
        if schedule_rotation_id is not None:
            pulumi.set(__self__, "schedule_rotation_id", schedule_rotation_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Position of the user inside rotation
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="scheduleRotationId")
    def schedule_rotation_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule_rotation_id")

    @schedule_rotation_id.setter
    def schedule_rotation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_rotation_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[int]]:
        """
        Schedule rotation user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "user_id", value)


class ScheduleRotationUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a ScheduleRotationUser resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] position: Position of the user inside rotation
        :param pulumi.Input[int] user_id: Schedule rotation user
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScheduleRotationUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ScheduleRotationUser resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ScheduleRotationUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduleRotationUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScheduleRotationUserArgs.__new__(ScheduleRotationUserArgs)

            __props__.__dict__["position"] = position
            __props__.__dict__["schedule_rotation_id"] = schedule_rotation_id
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(ScheduleRotationUser, __self__).__init__(
            'rootly:index/scheduleRotationUser:ScheduleRotationUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            position: Optional[pulumi.Input[int]] = None,
            schedule_rotation_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[int]] = None) -> 'ScheduleRotationUser':
        """
        Get an existing ScheduleRotationUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] position: Position of the user inside rotation
        :param pulumi.Input[int] user_id: Schedule rotation user
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduleRotationUserState.__new__(_ScheduleRotationUserState)

        __props__.__dict__["position"] = position
        __props__.__dict__["schedule_rotation_id"] = schedule_rotation_id
        __props__.__dict__["user_id"] = user_id
        return ScheduleRotationUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[int]:
        """
        Position of the user inside rotation
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter(name="scheduleRotationId")
    def schedule_rotation_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "schedule_rotation_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[int]:
        """
        Schedule rotation user
        """
        return pulumi.get(self, "user_id")

