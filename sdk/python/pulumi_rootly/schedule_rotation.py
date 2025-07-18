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

__all__ = ['ScheduleRotationArgs', 'ScheduleRotation']

@pulumi.input_type
class ScheduleRotationArgs:
    def __init__(__self__, *,
                 schedule_rotationable_attributes: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 active_all_week: Optional[pulumi.Input[bool]] = None,
                 active_days: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]]] = None,
                 active_time_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_id: Optional[pulumi.Input[str]] = None,
                 schedule_rotationable_type: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ScheduleRotation resource.
        :param pulumi.Input[bool] active_all_week: Schedule rotation active all week?. Value must be one of true or false
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_days: Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        :param pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]] active_time_attributes: Schedule rotation's active times
        :param pulumi.Input[str] name: The name of the schedule rotation
        :param pulumi.Input[int] position: Position of the schedule rotation
        :param pulumi.Input[str] schedule_id: The ID of parent schedule
        :param pulumi.Input[str] schedule_rotationable_type: Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        :param pulumi.Input[str] time_zone: A valid IANA time zone name.
        """
        pulumi.set(__self__, "schedule_rotationable_attributes", schedule_rotationable_attributes)
        if active_all_week is not None:
            pulumi.set(__self__, "active_all_week", active_all_week)
        if active_days is not None:
            pulumi.set(__self__, "active_days", active_days)
        if active_time_attributes is not None:
            pulumi.set(__self__, "active_time_attributes", active_time_attributes)
        if active_time_type is not None:
            pulumi.set(__self__, "active_time_type", active_time_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if schedule_id is not None:
            pulumi.set(__self__, "schedule_id", schedule_id)
        if schedule_rotationable_type is not None:
            pulumi.set(__self__, "schedule_rotationable_type", schedule_rotationable_type)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter(name="scheduleRotationableAttributes")
    def schedule_rotationable_attributes(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        return pulumi.get(self, "schedule_rotationable_attributes")

    @schedule_rotationable_attributes.setter
    def schedule_rotationable_attributes(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "schedule_rotationable_attributes", value)

    @property
    @pulumi.getter(name="activeAllWeek")
    def active_all_week(self) -> Optional[pulumi.Input[bool]]:
        """
        Schedule rotation active all week?. Value must be one of true or false
        """
        return pulumi.get(self, "active_all_week")

    @active_all_week.setter
    def active_all_week(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active_all_week", value)

    @property
    @pulumi.getter(name="activeDays")
    def active_days(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        return pulumi.get(self, "active_days")

    @active_days.setter
    def active_days(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "active_days", value)

    @property
    @pulumi.getter(name="activeTimeAttributes")
    def active_time_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]]]:
        """
        Schedule rotation's active times
        """
        return pulumi.get(self, "active_time_attributes")

    @active_time_attributes.setter
    def active_time_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]]]):
        pulumi.set(self, "active_time_attributes", value)

    @property
    @pulumi.getter(name="activeTimeType")
    def active_time_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "active_time_type")

    @active_time_type.setter
    def active_time_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "active_time_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the schedule rotation
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Position of the schedule rotation
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of parent schedule
        """
        return pulumi.get(self, "schedule_id")

    @schedule_id.setter
    def schedule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_id", value)

    @property
    @pulumi.getter(name="scheduleRotationableType")
    def schedule_rotationable_type(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        """
        return pulumi.get(self, "schedule_rotationable_type")

    @schedule_rotationable_type.setter
    def schedule_rotationable_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_rotationable_type", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[str]]:
        """
        A valid IANA time zone name.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_zone", value)


@pulumi.input_type
class _ScheduleRotationState:
    def __init__(__self__, *,
                 active_all_week: Optional[pulumi.Input[bool]] = None,
                 active_days: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]]] = None,
                 active_time_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_id: Optional[pulumi.Input[str]] = None,
                 schedule_rotationable_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 schedule_rotationable_type: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ScheduleRotation resources.
        :param pulumi.Input[bool] active_all_week: Schedule rotation active all week?. Value must be one of true or false
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_days: Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        :param pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]] active_time_attributes: Schedule rotation's active times
        :param pulumi.Input[str] name: The name of the schedule rotation
        :param pulumi.Input[int] position: Position of the schedule rotation
        :param pulumi.Input[str] schedule_id: The ID of parent schedule
        :param pulumi.Input[str] schedule_rotationable_type: Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        :param pulumi.Input[str] time_zone: A valid IANA time zone name.
        """
        if active_all_week is not None:
            pulumi.set(__self__, "active_all_week", active_all_week)
        if active_days is not None:
            pulumi.set(__self__, "active_days", active_days)
        if active_time_attributes is not None:
            pulumi.set(__self__, "active_time_attributes", active_time_attributes)
        if active_time_type is not None:
            pulumi.set(__self__, "active_time_type", active_time_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if schedule_id is not None:
            pulumi.set(__self__, "schedule_id", schedule_id)
        if schedule_rotationable_attributes is not None:
            pulumi.set(__self__, "schedule_rotationable_attributes", schedule_rotationable_attributes)
        if schedule_rotationable_type is not None:
            pulumi.set(__self__, "schedule_rotationable_type", schedule_rotationable_type)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter(name="activeAllWeek")
    def active_all_week(self) -> Optional[pulumi.Input[bool]]:
        """
        Schedule rotation active all week?. Value must be one of true or false
        """
        return pulumi.get(self, "active_all_week")

    @active_all_week.setter
    def active_all_week(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active_all_week", value)

    @property
    @pulumi.getter(name="activeDays")
    def active_days(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        return pulumi.get(self, "active_days")

    @active_days.setter
    def active_days(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "active_days", value)

    @property
    @pulumi.getter(name="activeTimeAttributes")
    def active_time_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]]]:
        """
        Schedule rotation's active times
        """
        return pulumi.get(self, "active_time_attributes")

    @active_time_attributes.setter
    def active_time_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ScheduleRotationActiveTimeAttributeArgs']]]]):
        pulumi.set(self, "active_time_attributes", value)

    @property
    @pulumi.getter(name="activeTimeType")
    def active_time_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "active_time_type")

    @active_time_type.setter
    def active_time_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "active_time_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the schedule rotation
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Position of the schedule rotation
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of parent schedule
        """
        return pulumi.get(self, "schedule_id")

    @schedule_id.setter
    def schedule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_id", value)

    @property
    @pulumi.getter(name="scheduleRotationableAttributes")
    def schedule_rotationable_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "schedule_rotationable_attributes")

    @schedule_rotationable_attributes.setter
    def schedule_rotationable_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "schedule_rotationable_attributes", value)

    @property
    @pulumi.getter(name="scheduleRotationableType")
    def schedule_rotationable_type(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        """
        return pulumi.get(self, "schedule_rotationable_type")

    @schedule_rotationable_type.setter
    def schedule_rotationable_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_rotationable_type", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[str]]:
        """
        A valid IANA time zone name.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_zone", value)


class ScheduleRotation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_all_week: Optional[pulumi.Input[bool]] = None,
                 active_days: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveTimeAttributeArgs', 'ScheduleRotationActiveTimeAttributeArgsDict']]]]] = None,
                 active_time_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_id: Optional[pulumi.Input[str]] = None,
                 schedule_rotationable_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 schedule_rotationable_type: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ScheduleRotation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active_all_week: Schedule rotation active all week?. Value must be one of true or false
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_days: Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveTimeAttributeArgs', 'ScheduleRotationActiveTimeAttributeArgsDict']]]] active_time_attributes: Schedule rotation's active times
        :param pulumi.Input[str] name: The name of the schedule rotation
        :param pulumi.Input[int] position: Position of the schedule rotation
        :param pulumi.Input[str] schedule_id: The ID of parent schedule
        :param pulumi.Input[str] schedule_rotationable_type: Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        :param pulumi.Input[str] time_zone: A valid IANA time zone name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScheduleRotationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ScheduleRotation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ScheduleRotationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduleRotationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_all_week: Optional[pulumi.Input[bool]] = None,
                 active_days: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveTimeAttributeArgs', 'ScheduleRotationActiveTimeAttributeArgsDict']]]]] = None,
                 active_time_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 schedule_id: Optional[pulumi.Input[str]] = None,
                 schedule_rotationable_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 schedule_rotationable_type: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScheduleRotationArgs.__new__(ScheduleRotationArgs)

            __props__.__dict__["active_all_week"] = active_all_week
            __props__.__dict__["active_days"] = active_days
            __props__.__dict__["active_time_attributes"] = active_time_attributes
            __props__.__dict__["active_time_type"] = active_time_type
            __props__.__dict__["name"] = name
            __props__.__dict__["position"] = position
            __props__.__dict__["schedule_id"] = schedule_id
            if schedule_rotationable_attributes is None and not opts.urn:
                raise TypeError("Missing required property 'schedule_rotationable_attributes'")
            __props__.__dict__["schedule_rotationable_attributes"] = schedule_rotationable_attributes
            __props__.__dict__["schedule_rotationable_type"] = schedule_rotationable_type
            __props__.__dict__["time_zone"] = time_zone
        super(ScheduleRotation, __self__).__init__(
            'rootly:index/scheduleRotation:ScheduleRotation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_all_week: Optional[pulumi.Input[bool]] = None,
            active_days: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            active_time_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveTimeAttributeArgs', 'ScheduleRotationActiveTimeAttributeArgsDict']]]]] = None,
            active_time_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            position: Optional[pulumi.Input[int]] = None,
            schedule_id: Optional[pulumi.Input[str]] = None,
            schedule_rotationable_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            schedule_rotationable_type: Optional[pulumi.Input[str]] = None,
            time_zone: Optional[pulumi.Input[str]] = None) -> 'ScheduleRotation':
        """
        Get an existing ScheduleRotation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active_all_week: Schedule rotation active all week?. Value must be one of true or false
        :param pulumi.Input[Sequence[pulumi.Input[str]]] active_days: Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ScheduleRotationActiveTimeAttributeArgs', 'ScheduleRotationActiveTimeAttributeArgsDict']]]] active_time_attributes: Schedule rotation's active times
        :param pulumi.Input[str] name: The name of the schedule rotation
        :param pulumi.Input[int] position: Position of the schedule rotation
        :param pulumi.Input[str] schedule_id: The ID of parent schedule
        :param pulumi.Input[str] schedule_rotationable_type: Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        :param pulumi.Input[str] time_zone: A valid IANA time zone name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduleRotationState.__new__(_ScheduleRotationState)

        __props__.__dict__["active_all_week"] = active_all_week
        __props__.__dict__["active_days"] = active_days
        __props__.__dict__["active_time_attributes"] = active_time_attributes
        __props__.__dict__["active_time_type"] = active_time_type
        __props__.__dict__["name"] = name
        __props__.__dict__["position"] = position
        __props__.__dict__["schedule_id"] = schedule_id
        __props__.__dict__["schedule_rotationable_attributes"] = schedule_rotationable_attributes
        __props__.__dict__["schedule_rotationable_type"] = schedule_rotationable_type
        __props__.__dict__["time_zone"] = time_zone
        return ScheduleRotation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeAllWeek")
    def active_all_week(self) -> pulumi.Output[bool]:
        """
        Schedule rotation active all week?. Value must be one of true or false
        """
        return pulumi.get(self, "active_all_week")

    @property
    @pulumi.getter(name="activeDays")
    def active_days(self) -> pulumi.Output[Sequence[str]]:
        """
        Value must be one of `S`, `M`, `T`, `W`, `R`, `F`, `U`.
        """
        return pulumi.get(self, "active_days")

    @property
    @pulumi.getter(name="activeTimeAttributes")
    def active_time_attributes(self) -> pulumi.Output[Sequence['outputs.ScheduleRotationActiveTimeAttribute']]:
        """
        Schedule rotation's active times
        """
        return pulumi.get(self, "active_time_attributes")

    @property
    @pulumi.getter(name="activeTimeType")
    def active_time_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "active_time_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the schedule rotation
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[int]:
        """
        Position of the schedule rotation
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> pulumi.Output[str]:
        """
        The ID of parent schedule
        """
        return pulumi.get(self, "schedule_id")

    @property
    @pulumi.getter(name="scheduleRotationableAttributes")
    def schedule_rotationable_attributes(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "schedule_rotationable_attributes")

    @property
    @pulumi.getter(name="scheduleRotationableType")
    def schedule_rotationable_type(self) -> pulumi.Output[Optional[str]]:
        """
        Schedule rotation type. Value must be one of `ScheduleDailyRotation`, `ScheduleWeeklyRotation`, `ScheduleBiweeklyRotation`, `ScheduleMonthlyRotation`, `ScheduleCustomRotation`.
        """
        return pulumi.get(self, "schedule_rotationable_type")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Output[Optional[str]]:
        """
        A valid IANA time zone name.
        """
        return pulumi.get(self, "time_zone")

