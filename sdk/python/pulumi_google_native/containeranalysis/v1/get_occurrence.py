# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetOccurrenceResult',
    'AwaitableGetOccurrenceResult',
    'get_occurrence',
    'get_occurrence_output',
]

@pulumi.output_type
class GetOccurrenceResult:
    def __init__(__self__, attestation=None, build=None, compliance=None, create_time=None, deployment=None, discovery=None, dsse_attestation=None, envelope=None, image=None, kind=None, name=None, note_name=None, package=None, remediation=None, resource_uri=None, update_time=None, upgrade=None, vulnerability=None):
        if attestation and not isinstance(attestation, dict):
            raise TypeError("Expected argument 'attestation' to be a dict")
        pulumi.set(__self__, "attestation", attestation)
        if build and not isinstance(build, dict):
            raise TypeError("Expected argument 'build' to be a dict")
        pulumi.set(__self__, "build", build)
        if compliance and not isinstance(compliance, dict):
            raise TypeError("Expected argument 'compliance' to be a dict")
        pulumi.set(__self__, "compliance", compliance)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if deployment and not isinstance(deployment, dict):
            raise TypeError("Expected argument 'deployment' to be a dict")
        pulumi.set(__self__, "deployment", deployment)
        if discovery and not isinstance(discovery, dict):
            raise TypeError("Expected argument 'discovery' to be a dict")
        pulumi.set(__self__, "discovery", discovery)
        if dsse_attestation and not isinstance(dsse_attestation, dict):
            raise TypeError("Expected argument 'dsse_attestation' to be a dict")
        pulumi.set(__self__, "dsse_attestation", dsse_attestation)
        if envelope and not isinstance(envelope, dict):
            raise TypeError("Expected argument 'envelope' to be a dict")
        pulumi.set(__self__, "envelope", envelope)
        if image and not isinstance(image, dict):
            raise TypeError("Expected argument 'image' to be a dict")
        pulumi.set(__self__, "image", image)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if note_name and not isinstance(note_name, str):
            raise TypeError("Expected argument 'note_name' to be a str")
        pulumi.set(__self__, "note_name", note_name)
        if package and not isinstance(package, dict):
            raise TypeError("Expected argument 'package' to be a dict")
        pulumi.set(__self__, "package", package)
        if remediation and not isinstance(remediation, str):
            raise TypeError("Expected argument 'remediation' to be a str")
        pulumi.set(__self__, "remediation", remediation)
        if resource_uri and not isinstance(resource_uri, str):
            raise TypeError("Expected argument 'resource_uri' to be a str")
        pulumi.set(__self__, "resource_uri", resource_uri)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if upgrade and not isinstance(upgrade, dict):
            raise TypeError("Expected argument 'upgrade' to be a dict")
        pulumi.set(__self__, "upgrade", upgrade)
        if vulnerability and not isinstance(vulnerability, dict):
            raise TypeError("Expected argument 'vulnerability' to be a dict")
        pulumi.set(__self__, "vulnerability", vulnerability)

    @property
    @pulumi.getter
    def attestation(self) -> 'outputs.AttestationOccurrenceResponse':
        """
        Describes an attestation of an artifact.
        """
        return pulumi.get(self, "attestation")

    @property
    @pulumi.getter
    def build(self) -> 'outputs.BuildOccurrenceResponse':
        """
        Describes a verifiable build.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter
    def compliance(self) -> 'outputs.ComplianceOccurrenceResponse':
        """
        Describes a compliance violation on a linked resource.
        """
        return pulumi.get(self, "compliance")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time this occurrence was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def deployment(self) -> 'outputs.DeploymentOccurrenceResponse':
        """
        Describes the deployment of an artifact on a runtime.
        """
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter
    def discovery(self) -> 'outputs.DiscoveryOccurrenceResponse':
        """
        Describes when a resource was discovered.
        """
        return pulumi.get(self, "discovery")

    @property
    @pulumi.getter(name="dsseAttestation")
    def dsse_attestation(self) -> 'outputs.DSSEAttestationOccurrenceResponse':
        """
        Describes an attestation of an artifact using dsse.
        """
        return pulumi.get(self, "dsse_attestation")

    @property
    @pulumi.getter
    def envelope(self) -> 'outputs.EnvelopeResponse':
        """
        https://github.com/secure-systems-lab/dsse
        """
        return pulumi.get(self, "envelope")

    @property
    @pulumi.getter
    def image(self) -> 'outputs.ImageOccurrenceResponse':
        """
        Describes how this resource derives from the basis in the associated note.
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noteName")
    def note_name(self) -> str:
        """
        Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "note_name")

    @property
    @pulumi.getter
    def package(self) -> 'outputs.PackageOccurrenceResponse':
        """
        Describes the installation of a package on the linked resource.
        """
        return pulumi.get(self, "package")

    @property
    @pulumi.getter
    def remediation(self) -> str:
        """
        A description of actions that can be taken to remedy the note.
        """
        return pulumi.get(self, "remediation")

    @property
    @pulumi.getter(name="resourceUri")
    def resource_uri(self) -> str:
        """
        Immutable. A URI that represents the resource for which the occurrence applies. For example, `https://gcr.io/project/image@sha256:123abc` for a Docker image.
        """
        return pulumi.get(self, "resource_uri")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time this occurrence was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def upgrade(self) -> 'outputs.UpgradeOccurrenceResponse':
        """
        Describes an available package upgrade on the linked resource.
        """
        return pulumi.get(self, "upgrade")

    @property
    @pulumi.getter
    def vulnerability(self) -> 'outputs.VulnerabilityOccurrenceResponse':
        """
        Describes a security vulnerability.
        """
        return pulumi.get(self, "vulnerability")


class AwaitableGetOccurrenceResult(GetOccurrenceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOccurrenceResult(
            attestation=self.attestation,
            build=self.build,
            compliance=self.compliance,
            create_time=self.create_time,
            deployment=self.deployment,
            discovery=self.discovery,
            dsse_attestation=self.dsse_attestation,
            envelope=self.envelope,
            image=self.image,
            kind=self.kind,
            name=self.name,
            note_name=self.note_name,
            package=self.package,
            remediation=self.remediation,
            resource_uri=self.resource_uri,
            update_time=self.update_time,
            upgrade=self.upgrade,
            vulnerability=self.vulnerability)


def get_occurrence(occurrence_id: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOccurrenceResult:
    """
    Gets the specified occurrence.
    """
    __args__ = dict()
    __args__['occurrenceId'] = occurrence_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:containeranalysis/v1:getOccurrence', __args__, opts=opts, typ=GetOccurrenceResult).value

    return AwaitableGetOccurrenceResult(
        attestation=__ret__.attestation,
        build=__ret__.build,
        compliance=__ret__.compliance,
        create_time=__ret__.create_time,
        deployment=__ret__.deployment,
        discovery=__ret__.discovery,
        dsse_attestation=__ret__.dsse_attestation,
        envelope=__ret__.envelope,
        image=__ret__.image,
        kind=__ret__.kind,
        name=__ret__.name,
        note_name=__ret__.note_name,
        package=__ret__.package,
        remediation=__ret__.remediation,
        resource_uri=__ret__.resource_uri,
        update_time=__ret__.update_time,
        upgrade=__ret__.upgrade,
        vulnerability=__ret__.vulnerability)


@_utilities.lift_output_func(get_occurrence)
def get_occurrence_output(occurrence_id: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOccurrenceResult]:
    """
    Gets the specified occurrence.
    """
    ...
