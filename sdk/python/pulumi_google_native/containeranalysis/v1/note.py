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

__all__ = ['NoteArgs', 'Note']

@pulumi.input_type
class NoteArgs:
    def __init__(__self__, *,
                 note_id: pulumi.Input[str],
                 attestation: Optional[pulumi.Input['AttestationNoteArgs']] = None,
                 build: Optional[pulumi.Input['BuildNoteArgs']] = None,
                 compliance: Optional[pulumi.Input['ComplianceNoteArgs']] = None,
                 deployment: Optional[pulumi.Input['DeploymentNoteArgs']] = None,
                 discovery: Optional[pulumi.Input['DiscoveryNoteArgs']] = None,
                 dsse_attestation: Optional[pulumi.Input['DSSEAttestationNoteArgs']] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input['ImageNoteArgs']] = None,
                 long_description: Optional[pulumi.Input[str]] = None,
                 package: Optional[pulumi.Input['PackageNoteArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 related_note_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 related_url: Optional[pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]]] = None,
                 short_description: Optional[pulumi.Input[str]] = None,
                 upgrade: Optional[pulumi.Input['UpgradeNoteArgs']] = None,
                 vulnerability: Optional[pulumi.Input['VulnerabilityNoteArgs']] = None):
        """
        The set of arguments for constructing a Note resource.
        :param pulumi.Input[str] note_id: Required. The ID to use for this note.
        :param pulumi.Input['AttestationNoteArgs'] attestation: A note describing an attestation role.
        :param pulumi.Input['BuildNoteArgs'] build: A note describing build provenance for a verifiable build.
        :param pulumi.Input['ComplianceNoteArgs'] compliance: A note describing a compliance check.
        :param pulumi.Input['DeploymentNoteArgs'] deployment: A note describing something that can be deployed.
        :param pulumi.Input['DiscoveryNoteArgs'] discovery: A note describing the initial analysis of a resource.
        :param pulumi.Input['DSSEAttestationNoteArgs'] dsse_attestation: A note describing a dsse attestation note.
        :param pulumi.Input[str] expiration_time: Time of expiration for this note. Empty if note does not expire.
        :param pulumi.Input['ImageNoteArgs'] image: A note describing a base image.
        :param pulumi.Input[str] long_description: A detailed description of this note.
        :param pulumi.Input['PackageNoteArgs'] package: A note describing a package hosted by various package managers.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] related_note_names: Other notes related to this note.
        :param pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]] related_url: URLs associated with this note.
        :param pulumi.Input[str] short_description: A one sentence description of this note.
        :param pulumi.Input['UpgradeNoteArgs'] upgrade: A note describing available package upgrades.
        :param pulumi.Input['VulnerabilityNoteArgs'] vulnerability: A note describing a package vulnerability.
        """
        pulumi.set(__self__, "note_id", note_id)
        if attestation is not None:
            pulumi.set(__self__, "attestation", attestation)
        if build is not None:
            pulumi.set(__self__, "build", build)
        if compliance is not None:
            pulumi.set(__self__, "compliance", compliance)
        if deployment is not None:
            pulumi.set(__self__, "deployment", deployment)
        if discovery is not None:
            pulumi.set(__self__, "discovery", discovery)
        if dsse_attestation is not None:
            pulumi.set(__self__, "dsse_attestation", dsse_attestation)
        if expiration_time is not None:
            pulumi.set(__self__, "expiration_time", expiration_time)
        if image is not None:
            pulumi.set(__self__, "image", image)
        if long_description is not None:
            pulumi.set(__self__, "long_description", long_description)
        if package is not None:
            pulumi.set(__self__, "package", package)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if related_note_names is not None:
            pulumi.set(__self__, "related_note_names", related_note_names)
        if related_url is not None:
            pulumi.set(__self__, "related_url", related_url)
        if short_description is not None:
            pulumi.set(__self__, "short_description", short_description)
        if upgrade is not None:
            pulumi.set(__self__, "upgrade", upgrade)
        if vulnerability is not None:
            pulumi.set(__self__, "vulnerability", vulnerability)

    @property
    @pulumi.getter(name="noteId")
    def note_id(self) -> pulumi.Input[str]:
        """
        Required. The ID to use for this note.
        """
        return pulumi.get(self, "note_id")

    @note_id.setter
    def note_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "note_id", value)

    @property
    @pulumi.getter
    def attestation(self) -> Optional[pulumi.Input['AttestationNoteArgs']]:
        """
        A note describing an attestation role.
        """
        return pulumi.get(self, "attestation")

    @attestation.setter
    def attestation(self, value: Optional[pulumi.Input['AttestationNoteArgs']]):
        pulumi.set(self, "attestation", value)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['BuildNoteArgs']]:
        """
        A note describing build provenance for a verifiable build.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['BuildNoteArgs']]):
        pulumi.set(self, "build", value)

    @property
    @pulumi.getter
    def compliance(self) -> Optional[pulumi.Input['ComplianceNoteArgs']]:
        """
        A note describing a compliance check.
        """
        return pulumi.get(self, "compliance")

    @compliance.setter
    def compliance(self, value: Optional[pulumi.Input['ComplianceNoteArgs']]):
        pulumi.set(self, "compliance", value)

    @property
    @pulumi.getter
    def deployment(self) -> Optional[pulumi.Input['DeploymentNoteArgs']]:
        """
        A note describing something that can be deployed.
        """
        return pulumi.get(self, "deployment")

    @deployment.setter
    def deployment(self, value: Optional[pulumi.Input['DeploymentNoteArgs']]):
        pulumi.set(self, "deployment", value)

    @property
    @pulumi.getter
    def discovery(self) -> Optional[pulumi.Input['DiscoveryNoteArgs']]:
        """
        A note describing the initial analysis of a resource.
        """
        return pulumi.get(self, "discovery")

    @discovery.setter
    def discovery(self, value: Optional[pulumi.Input['DiscoveryNoteArgs']]):
        pulumi.set(self, "discovery", value)

    @property
    @pulumi.getter(name="dsseAttestation")
    def dsse_attestation(self) -> Optional[pulumi.Input['DSSEAttestationNoteArgs']]:
        """
        A note describing a dsse attestation note.
        """
        return pulumi.get(self, "dsse_attestation")

    @dsse_attestation.setter
    def dsse_attestation(self, value: Optional[pulumi.Input['DSSEAttestationNoteArgs']]):
        pulumi.set(self, "dsse_attestation", value)

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of expiration for this note. Empty if note does not expire.
        """
        return pulumi.get(self, "expiration_time")

    @expiration_time.setter
    def expiration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time", value)

    @property
    @pulumi.getter
    def image(self) -> Optional[pulumi.Input['ImageNoteArgs']]:
        """
        A note describing a base image.
        """
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: Optional[pulumi.Input['ImageNoteArgs']]):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter(name="longDescription")
    def long_description(self) -> Optional[pulumi.Input[str]]:
        """
        A detailed description of this note.
        """
        return pulumi.get(self, "long_description")

    @long_description.setter
    def long_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "long_description", value)

    @property
    @pulumi.getter
    def package(self) -> Optional[pulumi.Input['PackageNoteArgs']]:
        """
        A note describing a package hosted by various package managers.
        """
        return pulumi.get(self, "package")

    @package.setter
    def package(self, value: Optional[pulumi.Input['PackageNoteArgs']]):
        pulumi.set(self, "package", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="relatedNoteNames")
    def related_note_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Other notes related to this note.
        """
        return pulumi.get(self, "related_note_names")

    @related_note_names.setter
    def related_note_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "related_note_names", value)

    @property
    @pulumi.getter(name="relatedUrl")
    def related_url(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]]]:
        """
        URLs associated with this note.
        """
        return pulumi.get(self, "related_url")

    @related_url.setter
    def related_url(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]]]):
        pulumi.set(self, "related_url", value)

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> Optional[pulumi.Input[str]]:
        """
        A one sentence description of this note.
        """
        return pulumi.get(self, "short_description")

    @short_description.setter
    def short_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_description", value)

    @property
    @pulumi.getter
    def upgrade(self) -> Optional[pulumi.Input['UpgradeNoteArgs']]:
        """
        A note describing available package upgrades.
        """
        return pulumi.get(self, "upgrade")

    @upgrade.setter
    def upgrade(self, value: Optional[pulumi.Input['UpgradeNoteArgs']]):
        pulumi.set(self, "upgrade", value)

    @property
    @pulumi.getter
    def vulnerability(self) -> Optional[pulumi.Input['VulnerabilityNoteArgs']]:
        """
        A note describing a package vulnerability.
        """
        return pulumi.get(self, "vulnerability")

    @vulnerability.setter
    def vulnerability(self, value: Optional[pulumi.Input['VulnerabilityNoteArgs']]):
        pulumi.set(self, "vulnerability", value)


class Note(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attestation: Optional[pulumi.Input[pulumi.InputType['AttestationNoteArgs']]] = None,
                 build: Optional[pulumi.Input[pulumi.InputType['BuildNoteArgs']]] = None,
                 compliance: Optional[pulumi.Input[pulumi.InputType['ComplianceNoteArgs']]] = None,
                 deployment: Optional[pulumi.Input[pulumi.InputType['DeploymentNoteArgs']]] = None,
                 discovery: Optional[pulumi.Input[pulumi.InputType['DiscoveryNoteArgs']]] = None,
                 dsse_attestation: Optional[pulumi.Input[pulumi.InputType['DSSEAttestationNoteArgs']]] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[pulumi.InputType['ImageNoteArgs']]] = None,
                 long_description: Optional[pulumi.Input[str]] = None,
                 note_id: Optional[pulumi.Input[str]] = None,
                 package: Optional[pulumi.Input[pulumi.InputType['PackageNoteArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 related_note_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 related_url: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RelatedUrlArgs']]]]] = None,
                 short_description: Optional[pulumi.Input[str]] = None,
                 upgrade: Optional[pulumi.Input[pulumi.InputType['UpgradeNoteArgs']]] = None,
                 vulnerability: Optional[pulumi.Input[pulumi.InputType['VulnerabilityNoteArgs']]] = None,
                 __props__=None):
        """
        Creates a new note.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AttestationNoteArgs']] attestation: A note describing an attestation role.
        :param pulumi.Input[pulumi.InputType['BuildNoteArgs']] build: A note describing build provenance for a verifiable build.
        :param pulumi.Input[pulumi.InputType['ComplianceNoteArgs']] compliance: A note describing a compliance check.
        :param pulumi.Input[pulumi.InputType['DeploymentNoteArgs']] deployment: A note describing something that can be deployed.
        :param pulumi.Input[pulumi.InputType['DiscoveryNoteArgs']] discovery: A note describing the initial analysis of a resource.
        :param pulumi.Input[pulumi.InputType['DSSEAttestationNoteArgs']] dsse_attestation: A note describing a dsse attestation note.
        :param pulumi.Input[str] expiration_time: Time of expiration for this note. Empty if note does not expire.
        :param pulumi.Input[pulumi.InputType['ImageNoteArgs']] image: A note describing a base image.
        :param pulumi.Input[str] long_description: A detailed description of this note.
        :param pulumi.Input[str] note_id: Required. The ID to use for this note.
        :param pulumi.Input[pulumi.InputType['PackageNoteArgs']] package: A note describing a package hosted by various package managers.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] related_note_names: Other notes related to this note.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RelatedUrlArgs']]]] related_url: URLs associated with this note.
        :param pulumi.Input[str] short_description: A one sentence description of this note.
        :param pulumi.Input[pulumi.InputType['UpgradeNoteArgs']] upgrade: A note describing available package upgrades.
        :param pulumi.Input[pulumi.InputType['VulnerabilityNoteArgs']] vulnerability: A note describing a package vulnerability.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NoteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new note.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param NoteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NoteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attestation: Optional[pulumi.Input[pulumi.InputType['AttestationNoteArgs']]] = None,
                 build: Optional[pulumi.Input[pulumi.InputType['BuildNoteArgs']]] = None,
                 compliance: Optional[pulumi.Input[pulumi.InputType['ComplianceNoteArgs']]] = None,
                 deployment: Optional[pulumi.Input[pulumi.InputType['DeploymentNoteArgs']]] = None,
                 discovery: Optional[pulumi.Input[pulumi.InputType['DiscoveryNoteArgs']]] = None,
                 dsse_attestation: Optional[pulumi.Input[pulumi.InputType['DSSEAttestationNoteArgs']]] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[pulumi.InputType['ImageNoteArgs']]] = None,
                 long_description: Optional[pulumi.Input[str]] = None,
                 note_id: Optional[pulumi.Input[str]] = None,
                 package: Optional[pulumi.Input[pulumi.InputType['PackageNoteArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 related_note_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 related_url: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RelatedUrlArgs']]]]] = None,
                 short_description: Optional[pulumi.Input[str]] = None,
                 upgrade: Optional[pulumi.Input[pulumi.InputType['UpgradeNoteArgs']]] = None,
                 vulnerability: Optional[pulumi.Input[pulumi.InputType['VulnerabilityNoteArgs']]] = None,
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
            __props__ = NoteArgs.__new__(NoteArgs)

            __props__.__dict__["attestation"] = attestation
            __props__.__dict__["build"] = build
            __props__.__dict__["compliance"] = compliance
            __props__.__dict__["deployment"] = deployment
            __props__.__dict__["discovery"] = discovery
            __props__.__dict__["dsse_attestation"] = dsse_attestation
            __props__.__dict__["expiration_time"] = expiration_time
            __props__.__dict__["image"] = image
            __props__.__dict__["long_description"] = long_description
            if note_id is None and not opts.urn:
                raise TypeError("Missing required property 'note_id'")
            __props__.__dict__["note_id"] = note_id
            __props__.__dict__["package"] = package
            __props__.__dict__["project"] = project
            __props__.__dict__["related_note_names"] = related_note_names
            __props__.__dict__["related_url"] = related_url
            __props__.__dict__["short_description"] = short_description
            __props__.__dict__["upgrade"] = upgrade
            __props__.__dict__["vulnerability"] = vulnerability
            __props__.__dict__["create_time"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        super(Note, __self__).__init__(
            'google-native:containeranalysis/v1:Note',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Note':
        """
        Get an existing Note resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NoteArgs.__new__(NoteArgs)

        __props__.__dict__["attestation"] = None
        __props__.__dict__["build"] = None
        __props__.__dict__["compliance"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["deployment"] = None
        __props__.__dict__["discovery"] = None
        __props__.__dict__["dsse_attestation"] = None
        __props__.__dict__["expiration_time"] = None
        __props__.__dict__["image"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["long_description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["package"] = None
        __props__.__dict__["related_note_names"] = None
        __props__.__dict__["related_url"] = None
        __props__.__dict__["short_description"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["upgrade"] = None
        __props__.__dict__["vulnerability"] = None
        return Note(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attestation(self) -> pulumi.Output['outputs.AttestationNoteResponse']:
        """
        A note describing an attestation role.
        """
        return pulumi.get(self, "attestation")

    @property
    @pulumi.getter
    def build(self) -> pulumi.Output['outputs.BuildNoteResponse']:
        """
        A note describing build provenance for a verifiable build.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter
    def compliance(self) -> pulumi.Output['outputs.ComplianceNoteResponse']:
        """
        A note describing a compliance check.
        """
        return pulumi.get(self, "compliance")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time this note was created. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def deployment(self) -> pulumi.Output['outputs.DeploymentNoteResponse']:
        """
        A note describing something that can be deployed.
        """
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter
    def discovery(self) -> pulumi.Output['outputs.DiscoveryNoteResponse']:
        """
        A note describing the initial analysis of a resource.
        """
        return pulumi.get(self, "discovery")

    @property
    @pulumi.getter(name="dsseAttestation")
    def dsse_attestation(self) -> pulumi.Output['outputs.DSSEAttestationNoteResponse']:
        """
        A note describing a dsse attestation note.
        """
        return pulumi.get(self, "dsse_attestation")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> pulumi.Output[str]:
        """
        Time of expiration for this note. Empty if note does not expire.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def image(self) -> pulumi.Output['outputs.ImageNoteResponse']:
        """
        A note describing a base image.
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The type of analysis. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="longDescription")
    def long_description(self) -> pulumi.Output[str]:
        """
        A detailed description of this note.
        """
        return pulumi.get(self, "long_description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def package(self) -> pulumi.Output['outputs.PackageNoteResponse']:
        """
        A note describing a package hosted by various package managers.
        """
        return pulumi.get(self, "package")

    @property
    @pulumi.getter(name="relatedNoteNames")
    def related_note_names(self) -> pulumi.Output[Sequence[str]]:
        """
        Other notes related to this note.
        """
        return pulumi.get(self, "related_note_names")

    @property
    @pulumi.getter(name="relatedUrl")
    def related_url(self) -> pulumi.Output[Sequence['outputs.RelatedUrlResponse']]:
        """
        URLs associated with this note.
        """
        return pulumi.get(self, "related_url")

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> pulumi.Output[str]:
        """
        A one sentence description of this note.
        """
        return pulumi.get(self, "short_description")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time this note was last updated. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def upgrade(self) -> pulumi.Output['outputs.UpgradeNoteResponse']:
        """
        A note describing available package upgrades.
        """
        return pulumi.get(self, "upgrade")

    @property
    @pulumi.getter
    def vulnerability(self) -> pulumi.Output['outputs.VulnerabilityNoteResponse']:
        """
        A note describing a package vulnerability.
        """
        return pulumi.get(self, "vulnerability")

