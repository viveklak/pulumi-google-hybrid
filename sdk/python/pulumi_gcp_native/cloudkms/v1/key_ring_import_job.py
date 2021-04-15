# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = ['KeyRingImportJobArgs', 'KeyRingImportJob']

@pulumi.input_type
class KeyRingImportJobArgs:
    def __init__(__self__, *,
                 import_jobs_id: pulumi.Input[str],
                 key_rings_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 import_method: Optional[pulumi.Input[str]] = None,
                 protection_level: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a KeyRingImportJob resource.
        :param pulumi.Input[str] import_method: Required. Immutable. The wrapping method to be used for incoming key material.
        :param pulumi.Input[str] protection_level: Required. Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        """
        pulumi.set(__self__, "import_jobs_id", import_jobs_id)
        pulumi.set(__self__, "key_rings_id", key_rings_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if import_method is not None:
            pulumi.set(__self__, "import_method", import_method)
        if protection_level is not None:
            pulumi.set(__self__, "protection_level", protection_level)

    @property
    @pulumi.getter(name="importJobsId")
    def import_jobs_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "import_jobs_id")

    @import_jobs_id.setter
    def import_jobs_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "import_jobs_id", value)

    @property
    @pulumi.getter(name="keyRingsId")
    def key_rings_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key_rings_id")

    @key_rings_id.setter
    def key_rings_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_rings_id", value)

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
    @pulumi.getter(name="importMethod")
    def import_method(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Immutable. The wrapping method to be used for incoming key material.
        """
        return pulumi.get(self, "import_method")

    @import_method.setter
    def import_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "import_method", value)

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        """
        return pulumi.get(self, "protection_level")

    @protection_level.setter
    def protection_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protection_level", value)


class KeyRingImportJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 import_jobs_id: Optional[pulumi.Input[str]] = None,
                 import_method: Optional[pulumi.Input[str]] = None,
                 key_rings_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 protection_level: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a new ImportJob within a KeyRing. ImportJob.import_method is required.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] import_method: Required. Immutable. The wrapping method to be used for incoming key material.
        :param pulumi.Input[str] protection_level: Required. Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyRingImportJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a new ImportJob within a KeyRing. ImportJob.import_method is required.

        :param str resource_name: The name of the resource.
        :param KeyRingImportJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyRingImportJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 import_jobs_id: Optional[pulumi.Input[str]] = None,
                 import_method: Optional[pulumi.Input[str]] = None,
                 key_rings_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 protection_level: Optional[pulumi.Input[str]] = None,
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
            __props__ = KeyRingImportJobArgs.__new__(KeyRingImportJobArgs)

            if import_jobs_id is None and not opts.urn:
                raise TypeError("Missing required property 'import_jobs_id'")
            __props__.__dict__["import_jobs_id"] = import_jobs_id
            __props__.__dict__["import_method"] = import_method
            if key_rings_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_rings_id'")
            __props__.__dict__["key_rings_id"] = key_rings_id
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["protection_level"] = protection_level
            __props__.__dict__["attestation"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["expire_event_time"] = None
            __props__.__dict__["expire_time"] = None
            __props__.__dict__["generate_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["public_key"] = None
            __props__.__dict__["state"] = None
        super(KeyRingImportJob, __self__).__init__(
            'gcp-native:cloudkms/v1:KeyRingImportJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'KeyRingImportJob':
        """
        Get an existing KeyRingImportJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyRingImportJobArgs.__new__(KeyRingImportJobArgs)

        __props__.__dict__["attestation"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["expire_event_time"] = None
        __props__.__dict__["expire_time"] = None
        __props__.__dict__["generate_time"] = None
        __props__.__dict__["import_method"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["protection_level"] = None
        __props__.__dict__["public_key"] = None
        __props__.__dict__["state"] = None
        return KeyRingImportJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attestation(self) -> pulumi.Output['outputs.KeyOperationAttestationResponse']:
        """
        Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen ImportMethod is one with a protection level of HSM.
        """
        return pulumi.get(self, "attestation")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time at which this ImportJob was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expireEventTime")
    def expire_event_time(self) -> pulumi.Output[str]:
        """
        The time this ImportJob expired. Only present if state is EXPIRED.
        """
        return pulumi.get(self, "expire_event_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        The time at which this ImportJob is scheduled for expiration and can no longer be used to import key material.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="generateTime")
    def generate_time(self) -> pulumi.Output[str]:
        """
        The time this ImportJob's key material was generated.
        """
        return pulumi.get(self, "generate_time")

    @property
    @pulumi.getter(name="importMethod")
    def import_method(self) -> pulumi.Output[str]:
        """
        Required. Immutable. The wrapping method to be used for incoming key material.
        """
        return pulumi.get(self, "import_method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for this ImportJob in the format `projects/*/locations/*/keyRings/*/importJobs/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> pulumi.Output[str]:
        """
        Required. Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        """
        return pulumi.get(self, "protection_level")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output['outputs.WrappingPublicKeyResponse']:
        """
        The public key with which to wrap key material prior to import. Only returned if state is ACTIVE.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the ImportJob, indicating if it can be used.
        """
        return pulumi.get(self, "state")

