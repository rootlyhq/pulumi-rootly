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

__all__ = ['ScheduleRotationActiveDayArgs', 'ScheduleRotationActiveDay']

@pulumi.input_type
class ScheduleRotationActiveDayArgs:
    def __init__(__self__, *,
                 active_time_attributes: pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]],
                 day_name: Optional[pulumi.Input[str]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ScheduleRotationActiveDay resource.
        :param pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]] active_time_attributes: Schedule rotation active times per day
        :param pulumi.Input[str] day_name: Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        pulumi.set(__self__, "active_time_attributes", active_time_attributes)
        if day_name is not None:
            pulumi.set(__self__, "day_name", day_name)
        if schedule_rotation_id is not None:
            pulumi.set(__self__, "schedule_rotation_id", schedule_rotation_id)

    @property
    @pulumi.getter(name="activeTimeAttributes")
    def active_time_attributes(self) -> pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]]:
        """
        Schedule rotation active times per day
        """
        return pulumi.get(self, "active_time_attributes")

    @active_time_attributes.setter
    def active_time_attributes(self, value: pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]]):
        pulumi.set(self, "active_time_attributes", value)

    @property
    @pulumi.getter(name="dayName")
    def day_name(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        return pulumi.get(self, "day_name")

    @day_name.setter
    def day_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_name", value)

    @property
    @pulumi.getter(name="scheduleRotationId")
    def schedule_rotation_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule_rotation_id")

    @schedule_rotation_id.setter
    def schedule_rotation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_rotation_id", value)


@pulumi.input_type
class _ScheduleRotationActiveDayState:
    def __init__(__self__, *,
                 active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]]] = None,
                 day_name: Optional[pulumi.Input[str]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ScheduleRotationActiveDay resources.
        :param pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]] active_time_attributes: Schedule rotation active times per day
        :param pulumi.Input[str] day_name: Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        if active_time_attributes is not None:
            pulumi.set(__self__, "active_time_attributes", active_time_attributes)
        if day_name is not None:
            pulumi.set(__self__, "day_name", day_name)
        if schedule_rotation_id is not None:
            pulumi.set(__self__, "schedule_rotation_id", schedule_rotation_id)

    @property
    @pulumi.getter(name="activeTimeAttributes")
    def active_time_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]]]:
        """
        Schedule rotation active times per day
        """
        return pulumi.get(self, "active_time_attributes")

    @active_time_attributes.setter
    def active_time_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveDayActiveTimeAttributeArgs']]]]):
        pulumi.set(self, "active_time_attributes", value)

    @property
    @pulumi.getter(name="dayName")
    def day_name(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        return pulumi.get(self, "day_name")

    @day_name.setter
    def day_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_name", value)

    @property
    @pulumi.getter(name="scheduleRotationId")
    def schedule_rotation_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule_rotation_id")

    @schedule_rotation_id.setter
    def schedule_rotation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_rotation_id", value)


class ScheduleRotationActiveDay(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveDayActiveTimeAttributeArgs', 'ScheduleRotationActiveDayActiveTimeAttributeArgsDict']]]]] = None,
                 day_name: Optional[pulumi.Input[str]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ScheduleRotationActiveDay resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveDayActiveTimeAttributeArgs', 'ScheduleRotationActiveDayActiveTimeAttributeArgsDict']]]] active_time_attributes: Schedule rotation active times per day
        :param pulumi.Input[str] day_name: Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScheduleRotationActiveDayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ScheduleRotationActiveDay resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ScheduleRotationActiveDayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduleRotationActiveDayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveDayActiveTimeAttributeArgs', 'ScheduleRotationActiveDayActiveTimeAttributeArgsDict']]]]] = None,
                 day_name: Optional[pulumi.Input[str]] = None,
                 schedule_rotation_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScheduleRotationActiveDayArgs.__new__(ScheduleRotationActiveDayArgs)

            if active_time_attributes is None and not opts.urn:
                raise TypeError("Missing required property 'active_time_attributes'")
            __props__.__dict__["active_time_attributes"] = active_time_attributes
            __props__.__dict__["day_name"] = day_name
            __props__.__dict__["schedule_rotation_id"] = schedule_rotation_id
        super(ScheduleRotationActiveDay, __self__).__init__(
            'rootly:index/scheduleRotationActiveDay:ScheduleRotationActiveDay',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveDayActiveTimeAttributeArgs', 'ScheduleRotationActiveDayActiveTimeAttributeArgsDict']]]]] = None,
            day_name: Optional[pulumi.Input[str]] = None,
            schedule_rotation_id: Optional[pulumi.Input[str]] = None) -> 'ScheduleRotationActiveDay':
        """
        Get an existing ScheduleRotationActiveDay resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveDayActiveTimeAttributeArgs', 'ScheduleRotationActiveDayActiveTimeAttributeArgsDict']]]] active_time_attributes: Schedule rotation active times per day
        :param pulumi.Input[str] day_name: Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduleRotationActiveDayState.__new__(_ScheduleRotationActiveDayState)

        __props__.__dict__["active_time_attributes"] = active_time_attributes
        __props__.__dict__["day_name"] = day_name
        __props__.__dict__["schedule_rotation_id"] = schedule_rotation_id
        return ScheduleRotationActiveDay(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeTimeAttributes")
    def active_time_attributes(self) -> pulumi.Output[Sequence['outputs.ScheduleRotationActiveDayActiveTimeAttribute']]:
        """
        Schedule rotation active times per day
        """
        return pulumi.get(self, "active_time_attributes")

    @property
    @pulumi.getter(name="dayName")
    def day_name(self) -> pulumi.Output[Optional[str]]:
        """
        Schedule rotation day name for which active times to be created. Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        return pulumi.get(self, "day_name")

    @property
    @pulumi.getter(name="scheduleRotationId")
    def schedule_rotation_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "schedule_rotation_id")

