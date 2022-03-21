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
    'GetCryptoKeyVersionResult',
    'AwaitableGetCryptoKeyVersionResult',
    'get_crypto_key_version',
    'get_crypto_key_version_output',
]

@pulumi.output_type
class GetCryptoKeyVersionResult:
    def __init__(__self__, algorithm=None, attestation=None, create_time=None, destroy_event_time=None, destroy_time=None, external_protection_level_options=None, generate_time=None, import_failure_reason=None, import_job=None, import_time=None, name=None, protection_level=None, reimport_eligible=None, state=None):
        if algorithm and not isinstance(algorithm, str):
            raise TypeError("Expected argument 'algorithm' to be a str")
        pulumi.set(__self__, "algorithm", algorithm)
        if attestation and not isinstance(attestation, dict):
            raise TypeError("Expected argument 'attestation' to be a dict")
        pulumi.set(__self__, "attestation", attestation)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if destroy_event_time and not isinstance(destroy_event_time, str):
            raise TypeError("Expected argument 'destroy_event_time' to be a str")
        pulumi.set(__self__, "destroy_event_time", destroy_event_time)
        if destroy_time and not isinstance(destroy_time, str):
            raise TypeError("Expected argument 'destroy_time' to be a str")
        pulumi.set(__self__, "destroy_time", destroy_time)
        if external_protection_level_options and not isinstance(external_protection_level_options, dict):
            raise TypeError("Expected argument 'external_protection_level_options' to be a dict")
        pulumi.set(__self__, "external_protection_level_options", external_protection_level_options)
        if generate_time and not isinstance(generate_time, str):
            raise TypeError("Expected argument 'generate_time' to be a str")
        pulumi.set(__self__, "generate_time", generate_time)
        if import_failure_reason and not isinstance(import_failure_reason, str):
            raise TypeError("Expected argument 'import_failure_reason' to be a str")
        pulumi.set(__self__, "import_failure_reason", import_failure_reason)
        if import_job and not isinstance(import_job, str):
            raise TypeError("Expected argument 'import_job' to be a str")
        pulumi.set(__self__, "import_job", import_job)
        if import_time and not isinstance(import_time, str):
            raise TypeError("Expected argument 'import_time' to be a str")
        pulumi.set(__self__, "import_time", import_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if protection_level and not isinstance(protection_level, str):
            raise TypeError("Expected argument 'protection_level' to be a str")
        pulumi.set(__self__, "protection_level", protection_level)
        if reimport_eligible and not isinstance(reimport_eligible, bool):
            raise TypeError("Expected argument 'reimport_eligible' to be a bool")
        pulumi.set(__self__, "reimport_eligible", reimport_eligible)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def attestation(self) -> 'outputs.KeyOperationAttestationResponse':
        """
        Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.
        """
        return pulumi.get(self, "attestation")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time at which this CryptoKeyVersion was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="destroyEventTime")
    def destroy_event_time(self) -> str:
        """
        The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.
        """
        return pulumi.get(self, "destroy_event_time")

    @property
    @pulumi.getter(name="destroyTime")
    def destroy_time(self) -> str:
        """
        The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.
        """
        return pulumi.get(self, "destroy_time")

    @property
    @pulumi.getter(name="externalProtectionLevelOptions")
    def external_protection_level_options(self) -> 'outputs.ExternalProtectionLevelOptionsResponse':
        """
        ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels.
        """
        return pulumi.get(self, "external_protection_level_options")

    @property
    @pulumi.getter(name="generateTime")
    def generate_time(self) -> str:
        """
        The time this CryptoKeyVersion's key material was generated.
        """
        return pulumi.get(self, "generate_time")

    @property
    @pulumi.getter(name="importFailureReason")
    def import_failure_reason(self) -> str:
        """
        The root cause of the most recent import failure. Only present if state is IMPORT_FAILED.
        """
        return pulumi.get(self, "import_failure_reason")

    @property
    @pulumi.getter(name="importJob")
    def import_job(self) -> str:
        """
        The name of the ImportJob used in the most recent import of this CryptoKeyVersion. Only present if the underlying key material was imported.
        """
        return pulumi.get(self, "import_job")

    @property
    @pulumi.getter(name="importTime")
    def import_time(self) -> str:
        """
        The time at which this CryptoKeyVersion's key material was most recently imported.
        """
        return pulumi.get(self, "import_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> str:
        """
        The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        """
        return pulumi.get(self, "protection_level")

    @property
    @pulumi.getter(name="reimportEligible")
    def reimport_eligible(self) -> bool:
        """
        Whether or not this key version is eligible for reimport, by being specified as a target in ImportCryptoKeyVersionRequest.crypto_key_version.
        """
        return pulumi.get(self, "reimport_eligible")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the CryptoKeyVersion.
        """
        return pulumi.get(self, "state")


class AwaitableGetCryptoKeyVersionResult(GetCryptoKeyVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCryptoKeyVersionResult(
            algorithm=self.algorithm,
            attestation=self.attestation,
            create_time=self.create_time,
            destroy_event_time=self.destroy_event_time,
            destroy_time=self.destroy_time,
            external_protection_level_options=self.external_protection_level_options,
            generate_time=self.generate_time,
            import_failure_reason=self.import_failure_reason,
            import_job=self.import_job,
            import_time=self.import_time,
            name=self.name,
            protection_level=self.protection_level,
            reimport_eligible=self.reimport_eligible,
            state=self.state)


def get_crypto_key_version(crypto_key_id: Optional[str] = None,
                           crypto_key_version_id: Optional[str] = None,
                           key_ring_id: Optional[str] = None,
                           location: Optional[str] = None,
                           project: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCryptoKeyVersionResult:
    """
    Returns metadata for a given CryptoKeyVersion.
    """
    __args__ = dict()
    __args__['cryptoKeyId'] = crypto_key_id
    __args__['cryptoKeyVersionId'] = crypto_key_version_id
    __args__['keyRingId'] = key_ring_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:cloudkms/v1:getCryptoKeyVersion', __args__, opts=opts, typ=GetCryptoKeyVersionResult).value

    return AwaitableGetCryptoKeyVersionResult(
        algorithm=__ret__.algorithm,
        attestation=__ret__.attestation,
        create_time=__ret__.create_time,
        destroy_event_time=__ret__.destroy_event_time,
        destroy_time=__ret__.destroy_time,
        external_protection_level_options=__ret__.external_protection_level_options,
        generate_time=__ret__.generate_time,
        import_failure_reason=__ret__.import_failure_reason,
        import_job=__ret__.import_job,
        import_time=__ret__.import_time,
        name=__ret__.name,
        protection_level=__ret__.protection_level,
        reimport_eligible=__ret__.reimport_eligible,
        state=__ret__.state)


@_utilities.lift_output_func(get_crypto_key_version)
def get_crypto_key_version_output(crypto_key_id: Optional[pulumi.Input[str]] = None,
                                  crypto_key_version_id: Optional[pulumi.Input[str]] = None,
                                  key_ring_id: Optional[pulumi.Input[str]] = None,
                                  location: Optional[pulumi.Input[str]] = None,
                                  project: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCryptoKeyVersionResult]:
    """
    Returns metadata for a given CryptoKeyVersion.
    """
    ...
