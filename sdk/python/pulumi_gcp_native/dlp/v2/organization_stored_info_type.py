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

__all__ = ['OrganizationStoredInfoTypeArgs', 'OrganizationStoredInfoType']

@pulumi.input_type
class OrganizationStoredInfoTypeArgs:
    def __init__(__self__, *,
                 organizations_id: pulumi.Input[str],
                 stored_info_types_id: pulumi.Input[str],
                 config: Optional[pulumi.Input['GooglePrivacyDlpV2StoredInfoTypeConfigArgs']] = None,
                 stored_info_type_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationStoredInfoType resource.
        :param pulumi.Input['GooglePrivacyDlpV2StoredInfoTypeConfigArgs'] config: Required. Configuration of the storedInfoType to create.
        :param pulumi.Input[str] stored_info_type_id: The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        """
        pulumi.set(__self__, "organizations_id", organizations_id)
        pulumi.set(__self__, "stored_info_types_id", stored_info_types_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if stored_info_type_id is not None:
            pulumi.set(__self__, "stored_info_type_id", stored_info_type_id)

    @property
    @pulumi.getter(name="organizationsId")
    def organizations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organizations_id")

    @organizations_id.setter
    def organizations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organizations_id", value)

    @property
    @pulumi.getter(name="storedInfoTypesId")
    def stored_info_types_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "stored_info_types_id")

    @stored_info_types_id.setter
    def stored_info_types_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "stored_info_types_id", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['GooglePrivacyDlpV2StoredInfoTypeConfigArgs']]:
        """
        Required. Configuration of the storedInfoType to create.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['GooglePrivacyDlpV2StoredInfoTypeConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="storedInfoTypeId")
    def stored_info_type_id(self) -> Optional[pulumi.Input[str]]:
        """
        The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        """
        return pulumi.get(self, "stored_info_type_id")

    @stored_info_type_id.setter
    def stored_info_type_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stored_info_type_id", value)


class OrganizationStoredInfoType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2StoredInfoTypeConfigArgs']]] = None,
                 organizations_id: Optional[pulumi.Input[str]] = None,
                 stored_info_type_id: Optional[pulumi.Input[str]] = None,
                 stored_info_types_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a pre-built stored infoType to be used for inspection. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2StoredInfoTypeConfigArgs']] config: Required. Configuration of the storedInfoType to create.
        :param pulumi.Input[str] stored_info_type_id: The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationStoredInfoTypeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a pre-built stored infoType to be used for inspection. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.

        :param str resource_name: The name of the resource.
        :param OrganizationStoredInfoTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationStoredInfoTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2StoredInfoTypeConfigArgs']]] = None,
                 organizations_id: Optional[pulumi.Input[str]] = None,
                 stored_info_type_id: Optional[pulumi.Input[str]] = None,
                 stored_info_types_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = OrganizationStoredInfoTypeArgs.__new__(OrganizationStoredInfoTypeArgs)

            __props__.__dict__["config"] = config
            if organizations_id is None and not opts.urn:
                raise TypeError("Missing required property 'organizations_id'")
            __props__.__dict__["organizations_id"] = organizations_id
            __props__.__dict__["stored_info_type_id"] = stored_info_type_id
            if stored_info_types_id is None and not opts.urn:
                raise TypeError("Missing required property 'stored_info_types_id'")
            __props__.__dict__["stored_info_types_id"] = stored_info_types_id
            __props__.__dict__["current_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["pending_versions"] = None
        super(OrganizationStoredInfoType, __self__).__init__(
            'gcp-native:dlp/v2:OrganizationStoredInfoType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationStoredInfoType':
        """
        Get an existing OrganizationStoredInfoType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationStoredInfoTypeArgs.__new__(OrganizationStoredInfoTypeArgs)

        __props__.__dict__["current_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["pending_versions"] = None
        return OrganizationStoredInfoType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> pulumi.Output['outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse']:
        """
        Current version of the stored info type.
        """
        return pulumi.get(self, "current_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pendingVersions")
    def pending_versions(self) -> pulumi.Output[Sequence['outputs.GooglePrivacyDlpV2StoredInfoTypeVersionResponse']]:
        """
        Pending versions of the stored info type. Empty if no versions are pending.
        """
        return pulumi.get(self, "pending_versions")

