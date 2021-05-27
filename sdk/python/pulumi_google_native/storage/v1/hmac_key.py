# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['HmacKeyArgs', 'HmacKey']

@pulumi.input_type
class HmacKeyArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 service_account_email: pulumi.Input[str],
                 user_project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HmacKey resource.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "service_account_email", service_account_email)
        if user_project is not None:
            pulumi.set(__self__, "user_project", user_project)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter(name="userProject")
    def user_project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_project")

    @user_project.setter
    def user_project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_project", value)


class HmacKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 user_project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new HMAC key for the specified service account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HmacKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new HMAC key for the specified service account.

        :param str resource_name: The name of the resource.
        :param HmacKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HmacKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 user_project: Optional[pulumi.Input[str]] = None,
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
            __props__ = HmacKeyArgs.__new__(HmacKeyArgs)

            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if service_account_email is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_email'")
            __props__.__dict__["service_account_email"] = service_account_email
            __props__.__dict__["user_project"] = user_project
            __props__.__dict__["access_id"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["time_created"] = None
            __props__.__dict__["updated"] = None
        super(HmacKey, __self__).__init__(
            'google-native:storage/v1:HmacKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HmacKey':
        """
        Get an existing HmacKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HmacKeyArgs.__new__(HmacKeyArgs)

        __props__.__dict__["access_id"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["service_account_email"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["time_created"] = None
        __props__.__dict__["updated"] = None
        return HmacKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessId")
    def access_id(self) -> pulumi.Output[str]:
        """
        The ID of the HMAC Key.
        """
        return pulumi.get(self, "access_id")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        HTTP 1.1 Entity tag for the HMAC key.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The kind of item this is. For HMAC Key metadata, this is always storage#hmacKeyMetadata.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        Project ID owning the service account to which the key authenticates.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The link to this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[str]:
        """
        The email address of the key's associated service account.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the key. Can be one of ACTIVE, INACTIVE, or DELETED.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[str]:
        """
        The creation time of the HMAC key in RFC 3339 format.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[str]:
        """
        The last modification time of the HMAC key metadata in RFC 3339 format.
        """
        return pulumi.get(self, "updated")

