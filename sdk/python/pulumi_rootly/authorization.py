# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AuthorizationArgs', 'Authorization']

@pulumi.input_type
class AuthorizationArgs:
    def __init__(__self__, *,
                 authorizable_id: pulumi.Input[str],
                 grantee_id: pulumi.Input[str],
                 permissions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 authorizable_type: Optional[pulumi.Input[str]] = None,
                 grantee_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Authorization resource.
        :param pulumi.Input[str] authorizable_id: The id of the resource being accessed.
        :param pulumi.Input[str] grantee_id: The resource id granted access.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: Value must be one of `read`, `update`, `authorize`, `destroy`.
        :param pulumi.Input[str] authorizable_type: The type of resource being accessed.. Value must be one of `Dashboard`.
        :param pulumi.Input[str] grantee_type: The type of resource granted access.. Value must be one of `User`, `Team`.
        """
        pulumi.set(__self__, "authorizable_id", authorizable_id)
        pulumi.set(__self__, "grantee_id", grantee_id)
        pulumi.set(__self__, "permissions", permissions)
        if authorizable_type is not None:
            pulumi.set(__self__, "authorizable_type", authorizable_type)
        if grantee_type is not None:
            pulumi.set(__self__, "grantee_type", grantee_type)

    @property
    @pulumi.getter(name="authorizableId")
    def authorizable_id(self) -> pulumi.Input[str]:
        """
        The id of the resource being accessed.
        """
        return pulumi.get(self, "authorizable_id")

    @authorizable_id.setter
    def authorizable_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorizable_id", value)

    @property
    @pulumi.getter(name="granteeId")
    def grantee_id(self) -> pulumi.Input[str]:
        """
        The resource id granted access.
        """
        return pulumi.get(self, "grantee_id")

    @grantee_id.setter
    def grantee_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "grantee_id", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Value must be one of `read`, `update`, `authorize`, `destroy`.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="authorizableType")
    def authorizable_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource being accessed.. Value must be one of `Dashboard`.
        """
        return pulumi.get(self, "authorizable_type")

    @authorizable_type.setter
    def authorizable_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizable_type", value)

    @property
    @pulumi.getter(name="granteeType")
    def grantee_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource granted access.. Value must be one of `User`, `Team`.
        """
        return pulumi.get(self, "grantee_type")

    @grantee_type.setter
    def grantee_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grantee_type", value)


@pulumi.input_type
class _AuthorizationState:
    def __init__(__self__, *,
                 authorizable_id: Optional[pulumi.Input[str]] = None,
                 authorizable_type: Optional[pulumi.Input[str]] = None,
                 grantee_id: Optional[pulumi.Input[str]] = None,
                 grantee_type: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Authorization resources.
        :param pulumi.Input[str] authorizable_id: The id of the resource being accessed.
        :param pulumi.Input[str] authorizable_type: The type of resource being accessed.. Value must be one of `Dashboard`.
        :param pulumi.Input[str] grantee_id: The resource id granted access.
        :param pulumi.Input[str] grantee_type: The type of resource granted access.. Value must be one of `User`, `Team`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: Value must be one of `read`, `update`, `authorize`, `destroy`.
        """
        if authorizable_id is not None:
            pulumi.set(__self__, "authorizable_id", authorizable_id)
        if authorizable_type is not None:
            pulumi.set(__self__, "authorizable_type", authorizable_type)
        if grantee_id is not None:
            pulumi.set(__self__, "grantee_id", grantee_id)
        if grantee_type is not None:
            pulumi.set(__self__, "grantee_type", grantee_type)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="authorizableId")
    def authorizable_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the resource being accessed.
        """
        return pulumi.get(self, "authorizable_id")

    @authorizable_id.setter
    def authorizable_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizable_id", value)

    @property
    @pulumi.getter(name="authorizableType")
    def authorizable_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource being accessed.. Value must be one of `Dashboard`.
        """
        return pulumi.get(self, "authorizable_type")

    @authorizable_type.setter
    def authorizable_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizable_type", value)

    @property
    @pulumi.getter(name="granteeId")
    def grantee_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id granted access.
        """
        return pulumi.get(self, "grantee_id")

    @grantee_id.setter
    def grantee_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grantee_id", value)

    @property
    @pulumi.getter(name="granteeType")
    def grantee_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource granted access.. Value must be one of `User`, `Team`.
        """
        return pulumi.get(self, "grantee_type")

    @grantee_type.setter
    def grantee_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grantee_type", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Value must be one of `read`, `update`, `authorize`, `destroy`.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)


class Authorization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorizable_id: Optional[pulumi.Input[str]] = None,
                 authorizable_type: Optional[pulumi.Input[str]] = None,
                 grantee_id: Optional[pulumi.Input[str]] = None,
                 grantee_type: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a Authorization resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorizable_id: The id of the resource being accessed.
        :param pulumi.Input[str] authorizable_type: The type of resource being accessed.. Value must be one of `Dashboard`.
        :param pulumi.Input[str] grantee_id: The resource id granted access.
        :param pulumi.Input[str] grantee_type: The type of resource granted access.. Value must be one of `User`, `Team`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: Value must be one of `read`, `update`, `authorize`, `destroy`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Authorization resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AuthorizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorizable_id: Optional[pulumi.Input[str]] = None,
                 authorizable_type: Optional[pulumi.Input[str]] = None,
                 grantee_id: Optional[pulumi.Input[str]] = None,
                 grantee_type: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthorizationArgs.__new__(AuthorizationArgs)

            if authorizable_id is None and not opts.urn:
                raise TypeError("Missing required property 'authorizable_id'")
            __props__.__dict__["authorizable_id"] = authorizable_id
            __props__.__dict__["authorizable_type"] = authorizable_type
            if grantee_id is None and not opts.urn:
                raise TypeError("Missing required property 'grantee_id'")
            __props__.__dict__["grantee_id"] = grantee_id
            __props__.__dict__["grantee_type"] = grantee_type
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
        super(Authorization, __self__).__init__(
            'rootly:index/authorization:Authorization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorizable_id: Optional[pulumi.Input[str]] = None,
            authorizable_type: Optional[pulumi.Input[str]] = None,
            grantee_id: Optional[pulumi.Input[str]] = None,
            grantee_type: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Authorization':
        """
        Get an existing Authorization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorizable_id: The id of the resource being accessed.
        :param pulumi.Input[str] authorizable_type: The type of resource being accessed.. Value must be one of `Dashboard`.
        :param pulumi.Input[str] grantee_id: The resource id granted access.
        :param pulumi.Input[str] grantee_type: The type of resource granted access.. Value must be one of `User`, `Team`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: Value must be one of `read`, `update`, `authorize`, `destroy`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthorizationState.__new__(_AuthorizationState)

        __props__.__dict__["authorizable_id"] = authorizable_id
        __props__.__dict__["authorizable_type"] = authorizable_type
        __props__.__dict__["grantee_id"] = grantee_id
        __props__.__dict__["grantee_type"] = grantee_type
        __props__.__dict__["permissions"] = permissions
        return Authorization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizableId")
    def authorizable_id(self) -> pulumi.Output[str]:
        """
        The id of the resource being accessed.
        """
        return pulumi.get(self, "authorizable_id")

    @property
    @pulumi.getter(name="authorizableType")
    def authorizable_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of resource being accessed.. Value must be one of `Dashboard`.
        """
        return pulumi.get(self, "authorizable_type")

    @property
    @pulumi.getter(name="granteeId")
    def grantee_id(self) -> pulumi.Output[str]:
        """
        The resource id granted access.
        """
        return pulumi.get(self, "grantee_id")

    @property
    @pulumi.getter(name="granteeType")
    def grantee_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of resource granted access.. Value must be one of `User`, `Team`.
        """
        return pulumi.get(self, "grantee_type")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence[str]]:
        """
        Value must be one of `read`, `update`, `authorize`, `destroy`.
        """
        return pulumi.get(self, "permissions")

