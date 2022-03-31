# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['FeedArgs', 'Feed']

@pulumi.input_type
class FeedArgs:
    def __init__(__self__, *,
                 feed_id: pulumi.Input[str],
                 feed_output_config: pulumi.Input['FeedOutputConfigArgs'],
                 name: pulumi.Input[str],
                 v1_id: pulumi.Input[str],
                 v1_id1: pulumi.Input[str],
                 asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 condition: Optional[pulumi.Input['ExprArgs']] = None,
                 content_type: Optional[pulumi.Input['FeedContentType']] = None,
                 relationship_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Feed resource.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent project/folder/organization.
        :param pulumi.Input['FeedOutputConfigArgs'] feed_output_config: Feed output configuration defining where the asset updates are published to.
        :param pulumi.Input[str] name: The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
        :param pulumi.Input['ExprArgs'] condition: A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition) for detailed instructions.
        :param pulumi.Input['FeedContentType'] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] relationship_types: A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
        """
        pulumi.set(__self__, "feed_id", feed_id)
        pulumi.set(__self__, "feed_output_config", feed_output_config)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "v1_id", v1_id)
        pulumi.set(__self__, "v1_id1", v1_id1)
        if asset_names is not None:
            pulumi.set(__self__, "asset_names", asset_names)
        if asset_types is not None:
            pulumi.set(__self__, "asset_types", asset_types)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if relationship_types is not None:
            pulumi.set(__self__, "relationship_types", relationship_types)

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> pulumi.Input[str]:
        """
        This is the client-assigned asset feed identifier and it needs to be unique under a specific parent project/folder/organization.
        """
        return pulumi.get(self, "feed_id")

    @feed_id.setter
    def feed_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "feed_id", value)

    @property
    @pulumi.getter(name="feedOutputConfig")
    def feed_output_config(self) -> pulumi.Input['FeedOutputConfigArgs']:
        """
        Feed output configuration defining where the asset updates are published to.
        """
        return pulumi.get(self, "feed_output_config")

    @feed_output_config.setter
    def feed_output_config(self, value: pulumi.Input['FeedOutputConfigArgs']):
        pulumi.set(self, "feed_output_config", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="v1Id")
    def v1_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "v1_id")

    @v1_id.setter
    def v1_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "v1_id", value)

    @property
    @pulumi.getter(name="v1Id1")
    def v1_id1(self) -> pulumi.Input[str]:
        return pulumi.get(self, "v1_id1")

    @v1_id1.setter
    def v1_id1(self, value: pulumi.Input[str]):
        pulumi.set(self, "v1_id1", value)

    @property
    @pulumi.getter(name="assetNames")
    def asset_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
        """
        return pulumi.get(self, "asset_names")

    @asset_names.setter
    def asset_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "asset_names", value)

    @property
    @pulumi.getter(name="assetTypes")
    def asset_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
        """
        return pulumi.get(self, "asset_types")

    @asset_types.setter
    def asset_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "asset_types", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['ExprArgs']]:
        """
        A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition) for detailed instructions.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['ExprArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input['FeedContentType']]:
        """
        Asset content type. If not specified, no content but the asset name and type will be returned.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input['FeedContentType']]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="relationshipTypes")
    def relationship_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
        """
        return pulumi.get(self, "relationship_types")

    @relationship_types.setter
    def relationship_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "relationship_types", value)


class Feed(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['ExprArgs']]] = None,
                 content_type: Optional[pulumi.Input['FeedContentType']] = None,
                 feed_id: Optional[pulumi.Input[str]] = None,
                 feed_output_config: Optional[pulumi.Input[pulumi.InputType['FeedOutputConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 relationship_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 v1_id: Optional[pulumi.Input[str]] = None,
                 v1_id1: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a feed in a parent project/folder/organization to listen to its asset updates.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
        :param pulumi.Input[pulumi.InputType['ExprArgs']] condition: A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition) for detailed instructions.
        :param pulumi.Input['FeedContentType'] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent project/folder/organization.
        :param pulumi.Input[pulumi.InputType['FeedOutputConfigArgs']] feed_output_config: Feed output configuration defining where the asset updates are published to.
        :param pulumi.Input[str] name: The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] relationship_types: A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FeedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a feed in a parent project/folder/organization to listen to its asset updates.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param FeedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FeedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['ExprArgs']]] = None,
                 content_type: Optional[pulumi.Input['FeedContentType']] = None,
                 feed_id: Optional[pulumi.Input[str]] = None,
                 feed_output_config: Optional[pulumi.Input[pulumi.InputType['FeedOutputConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 relationship_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 v1_id: Optional[pulumi.Input[str]] = None,
                 v1_id1: Optional[pulumi.Input[str]] = None,
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
            __props__ = FeedArgs.__new__(FeedArgs)

            __props__.__dict__["asset_names"] = asset_names
            __props__.__dict__["asset_types"] = asset_types
            __props__.__dict__["condition"] = condition
            __props__.__dict__["content_type"] = content_type
            if feed_id is None and not opts.urn:
                raise TypeError("Missing required property 'feed_id'")
            __props__.__dict__["feed_id"] = feed_id
            if feed_output_config is None and not opts.urn:
                raise TypeError("Missing required property 'feed_output_config'")
            __props__.__dict__["feed_output_config"] = feed_output_config
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["relationship_types"] = relationship_types
            if v1_id is None and not opts.urn:
                raise TypeError("Missing required property 'v1_id'")
            __props__.__dict__["v1_id"] = v1_id
            if v1_id1 is None and not opts.urn:
                raise TypeError("Missing required property 'v1_id1'")
            __props__.__dict__["v1_id1"] = v1_id1
        super(Feed, __self__).__init__(
            'google-native:cloudasset/v1:Feed',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Feed':
        """
        Get an existing Feed resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FeedArgs.__new__(FeedArgs)

        __props__.__dict__["asset_names"] = None
        __props__.__dict__["asset_types"] = None
        __props__.__dict__["condition"] = None
        __props__.__dict__["content_type"] = None
        __props__.__dict__["feed_output_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["relationship_types"] = None
        return Feed(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assetNames")
    def asset_names(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
        """
        return pulumi.get(self, "asset_names")

    @property
    @pulumi.getter(name="assetTypes")
    def asset_types(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
        """
        return pulumi.get(self, "asset_types")

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output['outputs.ExprResponse']:
        """
        A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition) for detailed instructions.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[str]:
        """
        Asset content type. If not specified, no content but the asset name and type will be returned.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="feedOutputConfig")
    def feed_output_config(self) -> pulumi.Output['outputs.FeedOutputConfigResponse']:
        """
        Feed output configuration defining where the asset updates are published to.
        """
        return pulumi.get(self, "feed_output_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="relationshipTypes")
    def relationship_types(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
        """
        return pulumi.get(self, "relationship_types")

