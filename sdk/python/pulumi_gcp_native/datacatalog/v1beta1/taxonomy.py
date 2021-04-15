# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = ['TaxonomyArgs', 'Taxonomy']

@pulumi.input_type
class TaxonomyArgs:
    def __init__(__self__, *,
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 taxonomies_id: pulumi.Input[str],
                 activated_policy_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Taxonomy resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] activated_policy_types: Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
        :param pulumi.Input[str] description: Optional. Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
        :param pulumi.Input[str] display_name: Required. User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        """
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        pulumi.set(__self__, "taxonomies_id", taxonomies_id)
        if activated_policy_types is not None:
            pulumi.set(__self__, "activated_policy_types", activated_policy_types)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

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
    @pulumi.getter(name="taxonomiesId")
    def taxonomies_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "taxonomies_id")

    @taxonomies_id.setter
    def taxonomies_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "taxonomies_id", value)

    @property
    @pulumi.getter(name="activatedPolicyTypes")
    def activated_policy_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
        """
        return pulumi.get(self, "activated_policy_types")

    @activated_policy_types.setter
    def activated_policy_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "activated_policy_types", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


class Taxonomy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activated_policy_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 taxonomies_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a taxonomy in the specified project.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] activated_policy_types: Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
        :param pulumi.Input[str] description: Optional. Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
        :param pulumi.Input[str] display_name: Required. User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaxonomyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a taxonomy in the specified project.

        :param str resource_name: The name of the resource.
        :param TaxonomyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaxonomyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activated_policy_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 taxonomies_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TaxonomyArgs.__new__(TaxonomyArgs)

            __props__.__dict__["activated_policy_types"] = activated_policy_types
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            if taxonomies_id is None and not opts.urn:
                raise TypeError("Missing required property 'taxonomies_id'")
            __props__.__dict__["taxonomies_id"] = taxonomies_id
            __props__.__dict__["name"] = None
            __props__.__dict__["policy_tag_count"] = None
            __props__.__dict__["taxonomy_timestamps"] = None
        super(Taxonomy, __self__).__init__(
            'gcp-native:datacatalog/v1beta1:Taxonomy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Taxonomy':
        """
        Get an existing Taxonomy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TaxonomyArgs.__new__(TaxonomyArgs)

        __props__.__dict__["activated_policy_types"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["policy_tag_count"] = None
        __props__.__dict__["taxonomy_timestamps"] = None
        return Taxonomy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activatedPolicyTypes")
    def activated_policy_types(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
        """
        return pulumi.get(self, "activated_policy_types")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of this taxonomy, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{id}".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyTagCount")
    def policy_tag_count(self) -> pulumi.Output[int]:
        """
        Number of policy tags contained in this taxonomy.
        """
        return pulumi.get(self, "policy_tag_count")

    @property
    @pulumi.getter(name="taxonomyTimestamps")
    def taxonomy_timestamps(self) -> pulumi.Output['outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse']:
        """
        Timestamps about this taxonomy. Only create_time and update_time are used.
        """
        return pulumi.get(self, "taxonomy_timestamps")

