# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 keys_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 restrictions: Optional[pulumi.Input['V2RestrictionsArgs']] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[str] display_name: Human-readable display name of this key that you can modify. The maximum length is 63 characters.
        :param pulumi.Input['V2RestrictionsArgs'] restrictions: Key restrictions.
        """
        pulumi.set(__self__, "keys_id", keys_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if restrictions is not None:
            pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter(name="keysId")
    def keys_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "keys_id")

    @keys_id.setter
    def keys_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "keys_id", value)

    @property
    @pulumi.getter(name="locationsId")
    def locations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "locations_id")

    @locations_id.setter
    def locations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "locations_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable display name of this key that you can modify. The maximum length is 63 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def restrictions(self) -> Optional[pulumi.Input['V2RestrictionsArgs']]:
        """
        Key restrictions.
        """
        return pulumi.get(self, "restrictions")

    @restrictions.setter
    def restrictions(self, value: Optional[pulumi.Input['V2RestrictionsArgs']]):
        pulumi.set(self, "restrictions", value)


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 keys_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 restrictions: Optional[pulumi.Input[pulumi.InputType['V2RestrictionsArgs']]] = None,
                 __props__=None):
        """
        Creates a new API key. NOTE: Key is a global resource; hence the only supported value for location is `global`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human-readable display name of this key that you can modify. The maximum length is 63 characters.
        :param pulumi.Input[pulumi.InputType['V2RestrictionsArgs']] restrictions: Key restrictions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new API key. NOTE: Key is a global resource; hence the only supported value for location is `global`.

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 keys_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 restrictions: Optional[pulumi.Input[pulumi.InputType['V2RestrictionsArgs']]] = None,
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
            __props__ = KeyArgs.__new__(KeyArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["key_id"] = key_id
            if keys_id is None and not opts.urn:
                raise TypeError("Missing required property 'keys_id'")
            __props__.__dict__["keys_id"] = keys_id
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["restrictions"] = restrictions
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["key_string"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        super(Key, __self__).__init__(
            'google-native:apikeys/v2:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyArgs.__new__(KeyArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["delete_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["key_string"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["restrictions"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        A timestamp identifying the time this key was originally created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        A timestamp when this key was deleted. If the resource is not deleted, this must be empty.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Human-readable display name of this key that you can modify. The maximum length is 63 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A checksum computed by the server based on the current value of the Key resource. This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="keyString")
    def key_string(self) -> pulumi.Output[str]:
        """
        An encrypted and signed value held by this key. This field can be accessed only through the `GetKeyString` method.
        """
        return pulumi.get(self, "key_string")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the key. The `name` has the form: `projects//locations/global/keys/`. For example: `projects/123456867718/locations/global/keys/b7ff1f9f-8275-410a-94dd-3855ee9b5dd2` NOTE: Key is a global resource; hence the only supported value for location is `global`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def restrictions(self) -> pulumi.Output['outputs.V2RestrictionsResponse']:
        """
        Key restrictions.
        """
        return pulumi.get(self, "restrictions")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Unique id in UUID4 format.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        A timestamp identifying the time this key was last updated.
        """
        return pulumi.get(self, "update_time")
