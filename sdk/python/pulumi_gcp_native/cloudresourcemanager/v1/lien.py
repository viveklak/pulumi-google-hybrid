# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['LienArgs', 'Lien']

@pulumi.input_type
class LienArgs:
    def __init__(__self__, *,
                 liens_id: pulumi.Input[str],
                 create_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Lien resource.
        :param pulumi.Input[str] create_time: The creation time of this Lien.
        :param pulumi.Input[str] name: A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
        :param pulumi.Input[str] origin: A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
        :param pulumi.Input[str] parent: A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
        :param pulumi.Input[str] reason: Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] restrictions: The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
        """
        pulumi.set(__self__, "liens_id", liens_id)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if origin is not None:
            pulumi.set(__self__, "origin", origin)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if reason is not None:
            pulumi.set(__self__, "reason", reason)
        if restrictions is not None:
            pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter(name="liensId")
    def liens_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "liens_id")

    @liens_id.setter
    def liens_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "liens_id", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of this Lien.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def origin(self) -> Optional[pulumi.Input[str]]:
        """
        A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
        """
        return pulumi.get(self, "origin")

    @origin.setter
    def origin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter
    def reason(self) -> Optional[pulumi.Input[str]]:
        """
        Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
        """
        return pulumi.get(self, "reason")

    @reason.setter
    def reason(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reason", value)

    @property
    @pulumi.getter
    def restrictions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
        """
        return pulumi.get(self, "restrictions")

    @restrictions.setter
    def restrictions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "restrictions", value)


class Lien(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 liens_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a Lien which applies to the resource denoted by the `parent` field. Callers of this method will require permission on the `parent` resource. For example, applying to `projects/1234` requires permission `resourcemanager.projects.updateLiens`. NOTE: Some resources may limit the number of Liens which may be applied.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The creation time of this Lien.
        :param pulumi.Input[str] name: A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
        :param pulumi.Input[str] origin: A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
        :param pulumi.Input[str] parent: A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
        :param pulumi.Input[str] reason: Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] restrictions: The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LienArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Lien which applies to the resource denoted by the `parent` field. Callers of this method will require permission on the `parent` resource. For example, applying to `projects/1234` requires permission `resourcemanager.projects.updateLiens`. NOTE: Some resources may limit the number of Liens which may be applied.

        :param str resource_name: The name of the resource.
        :param LienArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LienArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 liens_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 origin: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LienArgs.__new__(LienArgs)

            __props__.__dict__["create_time"] = create_time
            if liens_id is None and not opts.urn:
                raise TypeError("Missing required property 'liens_id'")
            __props__.__dict__["liens_id"] = liens_id
            __props__.__dict__["name"] = name
            __props__.__dict__["origin"] = origin
            __props__.__dict__["parent"] = parent
            __props__.__dict__["reason"] = reason
            __props__.__dict__["restrictions"] = restrictions
        super(Lien, __self__).__init__(
            'gcp-native:cloudresourcemanager/v1:Lien',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Lien':
        """
        Get an existing Lien resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LienArgs.__new__(LienArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["origin"] = None
        __props__.__dict__["parent"] = None
        __props__.__dict__["reason"] = None
        __props__.__dict__["restrictions"] = None
        return Lien(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of this Lien.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def origin(self) -> pulumi.Output[str]:
        """
        A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter
    def reason(self) -> pulumi.Output[str]:
        """
        Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
        """
        return pulumi.get(self, "reason")

    @property
    @pulumi.getter
    def restrictions(self) -> pulumi.Output[Sequence[str]]:
        """
        The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
        """
        return pulumi.get(self, "restrictions")

