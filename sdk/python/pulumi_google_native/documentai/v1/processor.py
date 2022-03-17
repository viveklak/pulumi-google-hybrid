# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ProcessorArgs', 'Processor']

@pulumi.input_type
class ProcessorArgs:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 default_processor_version: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Processor resource.
        :param pulumi.Input[str] create_time: The time the processor was created.
        :param pulumi.Input[str] default_processor_version: The default processor version.
        :param pulumi.Input[str] display_name: The display name of the processor.
        :param pulumi.Input[str] kms_key_name: The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
        :param pulumi.Input[str] type: The processor type, e.g., OCR_PROCESSOR, INVOICE_PROCESSOR, etc. To get a list of processors types, see FetchProcessorTypes.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if default_processor_version is not None:
            pulumi.set(__self__, "default_processor_version", default_processor_version)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time the processor was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="defaultProcessorVersion")
    def default_processor_version(self) -> Optional[pulumi.Input[str]]:
        """
        The default processor version.
        """
        return pulumi.get(self, "default_processor_version")

    @default_processor_version.setter
    def default_processor_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_processor_version", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the processor.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
        """
        return pulumi.get(self, "kms_key_name")

    @kms_key_name.setter
    def kms_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The processor type, e.g., OCR_PROCESSOR, INVOICE_PROCESSOR, etc. To get a list of processors types, see FetchProcessorTypes.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Processor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 default_processor_version: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a processor from the type processor that the user chose. The processor will be at "ENABLED" state by default after its creation.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time the processor was created.
        :param pulumi.Input[str] default_processor_version: The default processor version.
        :param pulumi.Input[str] display_name: The display name of the processor.
        :param pulumi.Input[str] kms_key_name: The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
        :param pulumi.Input[str] type: The processor type, e.g., OCR_PROCESSOR, INVOICE_PROCESSOR, etc. To get a list of processors types, see FetchProcessorTypes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProcessorArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a processor from the type processor that the user chose. The processor will be at "ENABLED" state by default after its creation.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ProcessorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProcessorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 default_processor_version: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = ProcessorArgs.__new__(ProcessorArgs)

            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["default_processor_version"] = default_processor_version
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["kms_key_name"] = kms_key_name
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["type"] = type
            __props__.__dict__["name"] = None
            __props__.__dict__["process_endpoint"] = None
            __props__.__dict__["state"] = None
        super(Processor, __self__).__init__(
            'google-native:documentai/v1:Processor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Processor':
        """
        Get an existing Processor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProcessorArgs.__new__(ProcessorArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["default_processor_version"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["kms_key_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["process_endpoint"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        return Processor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time the processor was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="defaultProcessorVersion")
    def default_processor_version(self) -> pulumi.Output[str]:
        """
        The default processor version.
        """
        return pulumi.get(self, "default_processor_version")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the processor.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> pulumi.Output[str]:
        """
        The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The resource name of the processor. Format: `projects/{project}/locations/{location}/processors/{processor}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="processEndpoint")
    def process_endpoint(self) -> pulumi.Output[str]:
        """
        Immutable. The http endpoint that can be called to invoke processing.
        """
        return pulumi.get(self, "process_endpoint")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the processor.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The processor type, e.g., OCR_PROCESSOR, INVOICE_PROCESSOR, etc. To get a list of processors types, see FetchProcessorTypes.
        """
        return pulumi.get(self, "type")

