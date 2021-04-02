# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Occurrence']


class Occurrence(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attestation: Optional[pulumi.Input[pulumi.InputType['AttestationArgs']]] = None,
                 build_details: Optional[pulumi.Input[pulumi.InputType['BuildDetailsArgs']]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 deployment: Optional[pulumi.Input[pulumi.InputType['DeploymentArgs']]] = None,
                 derived_image: Optional[pulumi.Input[pulumi.InputType['DerivedArgs']]] = None,
                 discovered: Optional[pulumi.Input[pulumi.InputType['DiscoveredArgs']]] = None,
                 installation: Optional[pulumi.Input[pulumi.InputType['InstallationArgs']]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 note_name: Optional[pulumi.Input[str]] = None,
                 occurrences_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 remediation: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[pulumi.InputType['ResourceArgs']]] = None,
                 resource_url: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 upgrade: Optional[pulumi.Input[pulumi.InputType['UpgradeOccurrenceArgs']]] = None,
                 vulnerability_details: Optional[pulumi.Input[pulumi.InputType['VulnerabilityDetailsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new `Occurrence`. Use this method to create `Occurrences` for a resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AttestationArgs']] attestation: Describes an attestation of an artifact.
        :param pulumi.Input[pulumi.InputType['BuildDetailsArgs']] build_details: Build details for a verifiable build.
        :param pulumi.Input[str] create_time: The time this `Occurrence` was created.
        :param pulumi.Input[pulumi.InputType['DeploymentArgs']] deployment: Describes the deployment of an artifact on a runtime.
        :param pulumi.Input[pulumi.InputType['DerivedArgs']] derived_image: Describes how this resource derives from the basis in the associated note.
        :param pulumi.Input[pulumi.InputType['DiscoveredArgs']] discovered: Describes the initial scan status for this resource.
        :param pulumi.Input[pulumi.InputType['InstallationArgs']] installation: Describes the installation of a package on the linked resource.
        :param pulumi.Input[str] kind: This explicitly denotes which of the `Occurrence` details are specified. This field can be used as a filter in list requests.
        :param pulumi.Input[str] name: The name of the `Occurrence` in the form "projects/{project_id}/occurrences/{OCCURRENCE_ID}"
        :param pulumi.Input[str] note_name: An analysis note associated with this image, in the form "providers/{provider_id}/notes/{NOTE_ID}" This field can be used as a filter in list requests.
        :param pulumi.Input[str] remediation: A description of actions that can be taken to remedy the `Note`
        :param pulumi.Input[pulumi.InputType['ResourceArgs']] resource:  The resource for which the `Occurrence` applies.
        :param pulumi.Input[str] resource_url: The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
        :param pulumi.Input[str] update_time: The time this `Occurrence` was last updated.
        :param pulumi.Input[pulumi.InputType['UpgradeOccurrenceArgs']] upgrade: Describes an upgrade.
        :param pulumi.Input[pulumi.InputType['VulnerabilityDetailsArgs']] vulnerability_details: Details of a security vulnerability note.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['attestation'] = attestation
            __props__['build_details'] = build_details
            __props__['create_time'] = create_time
            __props__['deployment'] = deployment
            __props__['derived_image'] = derived_image
            __props__['discovered'] = discovered
            __props__['installation'] = installation
            __props__['kind'] = kind
            __props__['name'] = name
            __props__['note_name'] = note_name
            if occurrences_id is None and not opts.urn:
                raise TypeError("Missing required property 'occurrences_id'")
            __props__['occurrences_id'] = occurrences_id
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__['projects_id'] = projects_id
            __props__['remediation'] = remediation
            __props__['resource'] = resource
            __props__['resource_url'] = resource_url
            __props__['update_time'] = update_time
            __props__['upgrade'] = upgrade
            __props__['vulnerability_details'] = vulnerability_details
        super(Occurrence, __self__).__init__(
            'google-cloud:containeranalysis/v1alpha1:Occurrence',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Occurrence':
        """
        Get an existing Occurrence resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attestation"] = None
        __props__["build_details"] = None
        __props__["create_time"] = None
        __props__["deployment"] = None
        __props__["derived_image"] = None
        __props__["discovered"] = None
        __props__["installation"] = None
        __props__["kind"] = None
        __props__["name"] = None
        __props__["note_name"] = None
        __props__["remediation"] = None
        __props__["resource"] = None
        __props__["resource_url"] = None
        __props__["update_time"] = None
        __props__["upgrade"] = None
        __props__["vulnerability_details"] = None
        return Occurrence(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attestation(self) -> pulumi.Output['outputs.AttestationResponse']:
        """
        Describes an attestation of an artifact.
        """
        return pulumi.get(self, "attestation")

    @property
    @pulumi.getter(name="buildDetails")
    def build_details(self) -> pulumi.Output['outputs.BuildDetailsResponse']:
        """
        Build details for a verifiable build.
        """
        return pulumi.get(self, "build_details")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time this `Occurrence` was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def deployment(self) -> pulumi.Output['outputs.DeploymentResponse']:
        """
        Describes the deployment of an artifact on a runtime.
        """
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter(name="derivedImage")
    def derived_image(self) -> pulumi.Output['outputs.DerivedResponse']:
        """
        Describes how this resource derives from the basis in the associated note.
        """
        return pulumi.get(self, "derived_image")

    @property
    @pulumi.getter
    def discovered(self) -> pulumi.Output['outputs.DiscoveredResponse']:
        """
        Describes the initial scan status for this resource.
        """
        return pulumi.get(self, "discovered")

    @property
    @pulumi.getter
    def installation(self) -> pulumi.Output['outputs.InstallationResponse']:
        """
        Describes the installation of a package on the linked resource.
        """
        return pulumi.get(self, "installation")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        This explicitly denotes which of the `Occurrence` details are specified. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the `Occurrence` in the form "projects/{project_id}/occurrences/{OCCURRENCE_ID}"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noteName")
    def note_name(self) -> pulumi.Output[str]:
        """
        An analysis note associated with this image, in the form "providers/{provider_id}/notes/{NOTE_ID}" This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "note_name")

    @property
    @pulumi.getter
    def remediation(self) -> pulumi.Output[str]:
        """
        A description of actions that can be taken to remedy the `Note`
        """
        return pulumi.get(self, "remediation")

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Output['outputs.ResourceResponse']:
        """
         The resource for which the `Occurrence` applies.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter(name="resourceUrl")
    def resource_url(self) -> pulumi.Output[str]:
        """
        The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "resource_url")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time this `Occurrence` was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def upgrade(self) -> pulumi.Output['outputs.UpgradeOccurrenceResponse']:
        """
        Describes an upgrade.
        """
        return pulumi.get(self, "upgrade")

    @property
    @pulumi.getter(name="vulnerabilityDetails")
    def vulnerability_details(self) -> pulumi.Output['outputs.VulnerabilityDetailsResponse']:
        """
        Details of a security vulnerability note.
        """
        return pulumi.get(self, "vulnerability_details")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
