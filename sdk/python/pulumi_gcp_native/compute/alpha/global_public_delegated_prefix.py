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

__all__ = ['GlobalPublicDelegatedPrefixArgs', 'GlobalPublicDelegatedPrefix']

@pulumi.input_type
class GlobalPublicDelegatedPrefixArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 public_delegated_prefix: pulumi.Input[str],
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ip_cidr_range: Optional[pulumi.Input[str]] = None,
                 is_live_migration: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_prefix: Optional[pulumi.Input[str]] = None,
                 public_delegated_sub_prefixs: Optional[pulumi.Input[Sequence[pulumi.Input['PublicDelegatedPrefixPublicDelegatedSubPrefixArgs']]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GlobalPublicDelegatedPrefix resource.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
               
               To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource type. The server generates this identifier.
        :param pulumi.Input[str] ip_cidr_range: The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        :param pulumi.Input[bool] is_live_migration: If true, the prefix will be live migrated.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] parent_prefix: The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        :param pulumi.Input[Sequence[pulumi.Input['PublicDelegatedPrefixPublicDelegatedSubPrefixArgs']]] public_delegated_sub_prefixs: The list of sub public delegated prefixes that exist for this public delegated prefix.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_id: [Output Only] Server-defined URL with id for the resource.
        :param pulumi.Input[str] status: [Output Only] The status of the public delegated prefix.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "public_delegated_prefix", public_delegated_prefix)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip_cidr_range is not None:
            pulumi.set(__self__, "ip_cidr_range", ip_cidr_range)
        if is_live_migration is not None:
            pulumi.set(__self__, "is_live_migration", is_live_migration)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_prefix is not None:
            pulumi.set(__self__, "parent_prefix", parent_prefix)
        if public_delegated_sub_prefixs is not None:
            pulumi.set(__self__, "public_delegated_sub_prefixs", public_delegated_sub_prefixs)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id is not None:
            pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="publicDelegatedPrefix")
    def public_delegated_prefix(self) -> pulumi.Input[str]:
        return pulumi.get(self, "public_delegated_prefix")

    @public_delegated_prefix.setter
    def public_delegated_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_delegated_prefix", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.

        To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The unique identifier for the resource type. The server generates this identifier.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="ipCidrRange")
    def ip_cidr_range(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        """
        return pulumi.get(self, "ip_cidr_range")

    @ip_cidr_range.setter
    def ip_cidr_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_cidr_range", value)

    @property
    @pulumi.getter(name="isLiveMigration")
    def is_live_migration(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, the prefix will be live migrated.
        """
        return pulumi.get(self, "is_live_migration")

    @is_live_migration.setter
    def is_live_migration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_live_migration", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentPrefix")
    def parent_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        """
        return pulumi.get(self, "parent_prefix")

    @parent_prefix.setter
    def parent_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_prefix", value)

    @property
    @pulumi.getter(name="publicDelegatedSubPrefixs")
    def public_delegated_sub_prefixs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PublicDelegatedPrefixPublicDelegatedSubPrefixArgs']]]]:
        """
        The list of sub public delegated prefixes that exist for this public delegated prefix.
        """
        return pulumi.get(self, "public_delegated_sub_prefixs")

    @public_delegated_sub_prefixs.setter
    def public_delegated_sub_prefixs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PublicDelegatedPrefixPublicDelegatedSubPrefixArgs']]]]):
        pulumi.set(self, "public_delegated_sub_prefixs", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL with id for the resource.
        """
        return pulumi.get(self, "self_link_with_id")

    @self_link_with_id.setter
    def self_link_with_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link_with_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The status of the public delegated prefix.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class GlobalPublicDelegatedPrefix(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ip_cidr_range: Optional[pulumi.Input[str]] = None,
                 is_live_migration: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 public_delegated_prefix: Optional[pulumi.Input[str]] = None,
                 public_delegated_sub_prefixs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PublicDelegatedPrefixPublicDelegatedSubPrefixArgs']]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
               
               To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource type. The server generates this identifier.
        :param pulumi.Input[str] ip_cidr_range: The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        :param pulumi.Input[bool] is_live_migration: If true, the prefix will be live migrated.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] parent_prefix: The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PublicDelegatedPrefixPublicDelegatedSubPrefixArgs']]]] public_delegated_sub_prefixs: The list of sub public delegated prefixes that exist for this public delegated prefix.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_id: [Output Only] Server-defined URL with id for the resource.
        :param pulumi.Input[str] status: [Output Only] The status of the public delegated prefix.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlobalPublicDelegatedPrefixArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param GlobalPublicDelegatedPrefixArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalPublicDelegatedPrefixArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ip_cidr_range: Optional[pulumi.Input[str]] = None,
                 is_live_migration: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 public_delegated_prefix: Optional[pulumi.Input[str]] = None,
                 public_delegated_sub_prefixs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PublicDelegatedPrefixPublicDelegatedSubPrefixArgs']]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = GlobalPublicDelegatedPrefixArgs.__new__(GlobalPublicDelegatedPrefixArgs)

            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["description"] = description
            __props__.__dict__["fingerprint"] = fingerprint
            __props__.__dict__["id"] = id
            __props__.__dict__["ip_cidr_range"] = ip_cidr_range
            __props__.__dict__["is_live_migration"] = is_live_migration
            __props__.__dict__["kind"] = kind
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_prefix"] = parent_prefix
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if public_delegated_prefix is None and not opts.urn:
                raise TypeError("Missing required property 'public_delegated_prefix'")
            __props__.__dict__["public_delegated_prefix"] = public_delegated_prefix
            __props__.__dict__["public_delegated_sub_prefixs"] = public_delegated_sub_prefixs
            __props__.__dict__["region"] = region
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["self_link_with_id"] = self_link_with_id
            __props__.__dict__["status"] = status
        super(GlobalPublicDelegatedPrefix, __self__).__init__(
            'gcp-native:compute/alpha:GlobalPublicDelegatedPrefix',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GlobalPublicDelegatedPrefix':
        """
        Get an existing GlobalPublicDelegatedPrefix resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GlobalPublicDelegatedPrefixArgs.__new__(GlobalPublicDelegatedPrefixArgs)

        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["fingerprint"] = None
        __props__.__dict__["ip_cidr_range"] = None
        __props__.__dict__["is_live_migration"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parent_prefix"] = None
        __props__.__dict__["public_delegated_sub_prefixs"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_link_with_id"] = None
        __props__.__dict__["status"] = None
        return GlobalPublicDelegatedPrefix(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.

        To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="ipCidrRange")
    def ip_cidr_range(self) -> pulumi.Output[str]:
        """
        The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        """
        return pulumi.get(self, "ip_cidr_range")

    @property
    @pulumi.getter(name="isLiveMigration")
    def is_live_migration(self) -> pulumi.Output[bool]:
        """
        If true, the prefix will be live migrated.
        """
        return pulumi.get(self, "is_live_migration")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output Only] Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentPrefix")
    def parent_prefix(self) -> pulumi.Output[str]:
        """
        The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        """
        return pulumi.get(self, "parent_prefix")

    @property
    @pulumi.getter(name="publicDelegatedSubPrefixs")
    def public_delegated_sub_prefixs(self) -> pulumi.Output[Sequence['outputs.PublicDelegatedPrefixPublicDelegatedSubPrefixResponse']]:
        """
        The list of sub public delegated prefixes that exist for this public delegated prefix.
        """
        return pulumi.get(self, "public_delegated_sub_prefixs")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        [Output Only] URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL with id for the resource.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        [Output Only] The status of the public delegated prefix.
        """
        return pulumi.get(self, "status")

