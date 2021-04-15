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

__all__ = ['FeatureArgs', 'Feature']

@pulumi.input_type
class FeatureArgs:
    def __init__(__self__, *,
                 features_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 spec: Optional[pulumi.Input['CommonFeatureSpecArgs']] = None):
        """
        The set of arguments for constructing a Feature resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: GCP labels for this Feature.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] membership_specs: Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: projects/{p}/locations/{l}/memberships/{m} Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        :param pulumi.Input['CommonFeatureSpecArgs'] spec: Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        pulumi.set(__self__, "features_id", features_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if membership_specs is not None:
            pulumi.set(__self__, "membership_specs", membership_specs)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="featuresId")
    def features_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "features_id")

    @features_id.setter
    def features_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "features_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        GCP labels for this Feature.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="membershipSpecs")
    def membership_specs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: projects/{p}/locations/{l}/memberships/{m} Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        """
        return pulumi.get(self, "membership_specs")

    @membership_specs.setter
    def membership_specs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "membership_specs", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['CommonFeatureSpecArgs']]:
        """
        Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['CommonFeatureSpecArgs']]):
        pulumi.set(self, "spec", value)


class Feature(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 features_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['CommonFeatureSpecArgs']]] = None,
                 __props__=None):
        """
        Adds a new Feature.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: GCP labels for this Feature.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] membership_specs: Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: projects/{p}/locations/{l}/memberships/{m} Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        :param pulumi.Input[pulumi.InputType['CommonFeatureSpecArgs']] spec: Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FeatureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Adds a new Feature.

        :param str resource_name: The name of the resource.
        :param FeatureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FeatureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 features_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['CommonFeatureSpecArgs']]] = None,
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
            __props__ = FeatureArgs.__new__(FeatureArgs)

            if features_id is None and not opts.urn:
                raise TypeError("Missing required property 'features_id'")
            __props__.__dict__["features_id"] = features_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["membership_specs"] = membership_specs
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["spec"] = spec
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["membership_states"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["resource_state"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        super(Feature, __self__).__init__(
            'gcp-native:gkehub/v1beta:Feature',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Feature':
        """
        Get an existing Feature resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FeatureArgs.__new__(FeatureArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["delete_time"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["membership_specs"] = None
        __props__.__dict__["membership_states"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["resource_state"] = None
        __props__.__dict__["spec"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return Feature(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        When the Feature resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        When the Feature resource was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        GCP labels for this Feature.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="membershipSpecs")
    def membership_specs(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: projects/{p}/locations/{l}/memberships/{m} Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        """
        return pulumi.get(self, "membership_specs")

    @property
    @pulumi.getter(name="membershipStates")
    def membership_states(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: projects/{p}/locations/{l}/memberships/{m} Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        """
        return pulumi.get(self, "membership_states")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The full, unique name of this Feature resource in the format `projects/*/locations/global/features/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> pulumi.Output['outputs.FeatureResourceStateResponse']:
        """
        State of the Feature resource itself.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output['outputs.CommonFeatureSpecResponse']:
        """
        Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output['outputs.CommonFeatureStateResponse']:
        """
        The Hub-wide Feature state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        When the Feature resource was last updated.
        """
        return pulumi.get(self, "update_time")

