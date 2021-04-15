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

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 instances_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 accelerators: Optional[pulumi.Input[Sequence[pulumi.Input['AcceleratorArgs']]]] = None,
                 available_version: Optional[pulumi.Input[Sequence[pulumi.Input['VersionArgs']]]] = None,
                 dataproc_service_account: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_rbac: Optional[pulumi.Input[bool]] = None,
                 enable_stackdriver_logging: Optional[pulumi.Input[bool]] = None,
                 enable_stackdriver_monitoring: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 network_config: Optional[pulumi.Input['NetworkConfigArgs']] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 private_instance: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[Sequence[pulumi.Input['AcceleratorArgs']]] accelerators: List of accelerators enabled for this CDF instance.
        :param pulumi.Input[Sequence[pulumi.Input['VersionArgs']]] available_version: Available versions that the instance can be upgraded to using UpdateInstanceRequest.
        :param pulumi.Input[str] dataproc_service_account: User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
        :param pulumi.Input[str] description: A description of this instance.
        :param pulumi.Input[str] display_name: Display name for an instance.
        :param pulumi.Input[bool] enable_rbac: Option to enable granular role-based access control.
        :param pulumi.Input[bool] enable_stackdriver_logging: Option to enable Stackdriver Logging.
        :param pulumi.Input[bool] enable_stackdriver_monitoring: Option to enable Stackdriver Monitoring.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
        :param pulumi.Input['NetworkConfigArgs'] network_config: Network configuration options. These are required when a private Data Fusion instance is to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: Map of additional options used to configure the behavior of Data Fusion instance.
        :param pulumi.Input[bool] private_instance: Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
        :param pulumi.Input[str] type: Required. Instance type.
        :param pulumi.Input[str] version: Current version of Data Fusion.
        :param pulumi.Input[str] zone: Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
        """
        pulumi.set(__self__, "instances_id", instances_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if accelerators is not None:
            pulumi.set(__self__, "accelerators", accelerators)
        if available_version is not None:
            pulumi.set(__self__, "available_version", available_version)
        if dataproc_service_account is not None:
            pulumi.set(__self__, "dataproc_service_account", dataproc_service_account)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if enable_rbac is not None:
            pulumi.set(__self__, "enable_rbac", enable_rbac)
        if enable_stackdriver_logging is not None:
            pulumi.set(__self__, "enable_stackdriver_logging", enable_stackdriver_logging)
        if enable_stackdriver_monitoring is not None:
            pulumi.set(__self__, "enable_stackdriver_monitoring", enable_stackdriver_monitoring)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if network_config is not None:
            pulumi.set(__self__, "network_config", network_config)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if private_instance is not None:
            pulumi.set(__self__, "private_instance", private_instance)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="instancesId")
    def instances_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instances_id")

    @instances_id.setter
    def instances_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instances_id", value)

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
    @pulumi.getter
    def accelerators(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AcceleratorArgs']]]]:
        """
        List of accelerators enabled for this CDF instance.
        """
        return pulumi.get(self, "accelerators")

    @accelerators.setter
    def accelerators(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AcceleratorArgs']]]]):
        pulumi.set(self, "accelerators", value)

    @property
    @pulumi.getter(name="availableVersion")
    def available_version(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VersionArgs']]]]:
        """
        Available versions that the instance can be upgraded to using UpdateInstanceRequest.
        """
        return pulumi.get(self, "available_version")

    @available_version.setter
    def available_version(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VersionArgs']]]]):
        pulumi.set(self, "available_version", value)

    @property
    @pulumi.getter(name="dataprocServiceAccount")
    def dataproc_service_account(self) -> Optional[pulumi.Input[str]]:
        """
        User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
        """
        return pulumi.get(self, "dataproc_service_account")

    @dataproc_service_account.setter
    def dataproc_service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataproc_service_account", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of this instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for an instance.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="enableRbac")
    def enable_rbac(self) -> Optional[pulumi.Input[bool]]:
        """
        Option to enable granular role-based access control.
        """
        return pulumi.get(self, "enable_rbac")

    @enable_rbac.setter
    def enable_rbac(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_rbac", value)

    @property
    @pulumi.getter(name="enableStackdriverLogging")
    def enable_stackdriver_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Option to enable Stackdriver Logging.
        """
        return pulumi.get(self, "enable_stackdriver_logging")

    @enable_stackdriver_logging.setter
    def enable_stackdriver_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_stackdriver_logging", value)

    @property
    @pulumi.getter(name="enableStackdriverMonitoring")
    def enable_stackdriver_monitoring(self) -> Optional[pulumi.Input[bool]]:
        """
        Option to enable Stackdriver Monitoring.
        """
        return pulumi.get(self, "enable_stackdriver_monitoring")

    @enable_stackdriver_monitoring.setter
    def enable_stackdriver_monitoring(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_stackdriver_monitoring", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> Optional[pulumi.Input['NetworkConfigArgs']]:
        """
        Network configuration options. These are required when a private Data Fusion instance is to be created.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: Optional[pulumi.Input['NetworkConfigArgs']]):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of additional options used to configure the behavior of Data Fusion instance.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter(name="privateInstance")
    def private_instance(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
        """
        return pulumi.get(self, "private_instance")

    @private_instance.setter
    def private_instance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "private_instance", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Instance type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Current version of Data Fusion.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcceleratorArgs']]]]] = None,
                 available_version: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VersionArgs']]]]] = None,
                 dataproc_service_account: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_rbac: Optional[pulumi.Input[bool]] = None,
                 enable_stackdriver_logging: Optional[pulumi.Input[bool]] = None,
                 enable_stackdriver_monitoring: Optional[pulumi.Input[bool]] = None,
                 instances_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 private_instance: Optional[pulumi.Input[bool]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new Data Fusion instance in the specified project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcceleratorArgs']]]] accelerators: List of accelerators enabled for this CDF instance.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VersionArgs']]]] available_version: Available versions that the instance can be upgraded to using UpdateInstanceRequest.
        :param pulumi.Input[str] dataproc_service_account: User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
        :param pulumi.Input[str] description: A description of this instance.
        :param pulumi.Input[str] display_name: Display name for an instance.
        :param pulumi.Input[bool] enable_rbac: Option to enable granular role-based access control.
        :param pulumi.Input[bool] enable_stackdriver_logging: Option to enable Stackdriver Logging.
        :param pulumi.Input[bool] enable_stackdriver_monitoring: Option to enable Stackdriver Monitoring.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
        :param pulumi.Input[pulumi.InputType['NetworkConfigArgs']] network_config: Network configuration options. These are required when a private Data Fusion instance is to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: Map of additional options used to configure the behavior of Data Fusion instance.
        :param pulumi.Input[bool] private_instance: Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
        :param pulumi.Input[str] type: Required. Instance type.
        :param pulumi.Input[str] version: Current version of Data Fusion.
        :param pulumi.Input[str] zone: Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Data Fusion instance in the specified project and location.

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AcceleratorArgs']]]]] = None,
                 available_version: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VersionArgs']]]]] = None,
                 dataproc_service_account: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_rbac: Optional[pulumi.Input[bool]] = None,
                 enable_stackdriver_logging: Optional[pulumi.Input[bool]] = None,
                 enable_stackdriver_monitoring: Optional[pulumi.Input[bool]] = None,
                 instances_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 private_instance: Optional[pulumi.Input[bool]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
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
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["accelerators"] = accelerators
            __props__.__dict__["available_version"] = available_version
            __props__.__dict__["dataproc_service_account"] = dataproc_service_account
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["enable_rbac"] = enable_rbac
            __props__.__dict__["enable_stackdriver_logging"] = enable_stackdriver_logging
            __props__.__dict__["enable_stackdriver_monitoring"] = enable_stackdriver_monitoring
            if instances_id is None and not opts.urn:
                raise TypeError("Missing required property 'instances_id'")
            __props__.__dict__["instances_id"] = instances_id
            __props__.__dict__["labels"] = labels
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["network_config"] = network_config
            __props__.__dict__["options"] = options
            __props__.__dict__["private_instance"] = private_instance
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["type"] = type
            __props__.__dict__["version"] = version
            __props__.__dict__["zone"] = zone
            __props__.__dict__["api_endpoint"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["gcs_bucket"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["p4_service_account"] = None
            __props__.__dict__["service_endpoint"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["state_message"] = None
            __props__.__dict__["tenant_project_id"] = None
            __props__.__dict__["update_time"] = None
        super(Instance, __self__).__init__(
            'gcp-native:datafusion/v1beta1:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstanceArgs.__new__(InstanceArgs)

        __props__.__dict__["accelerators"] = None
        __props__.__dict__["api_endpoint"] = None
        __props__.__dict__["available_version"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["dataproc_service_account"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["enable_rbac"] = None
        __props__.__dict__["enable_stackdriver_logging"] = None
        __props__.__dict__["enable_stackdriver_monitoring"] = None
        __props__.__dict__["gcs_bucket"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_config"] = None
        __props__.__dict__["options"] = None
        __props__.__dict__["p4_service_account"] = None
        __props__.__dict__["private_instance"] = None
        __props__.__dict__["service_endpoint"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["state_message"] = None
        __props__.__dict__["tenant_project_id"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["version"] = None
        __props__.__dict__["zone"] = None
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accelerators(self) -> pulumi.Output[Sequence['outputs.AcceleratorResponse']]:
        """
        List of accelerators enabled for this CDF instance.
        """
        return pulumi.get(self, "accelerators")

    @property
    @pulumi.getter(name="apiEndpoint")
    def api_endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint on which the REST APIs is accessible.
        """
        return pulumi.get(self, "api_endpoint")

    @property
    @pulumi.getter(name="availableVersion")
    def available_version(self) -> pulumi.Output[Sequence['outputs.VersionResponse']]:
        """
        Available versions that the instance can be upgraded to using UpdateInstanceRequest.
        """
        return pulumi.get(self, "available_version")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dataprocServiceAccount")
    def dataproc_service_account(self) -> pulumi.Output[str]:
        """
        User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
        """
        return pulumi.get(self, "dataproc_service_account")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description of this instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name for an instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enableRbac")
    def enable_rbac(self) -> pulumi.Output[bool]:
        """
        Option to enable granular role-based access control.
        """
        return pulumi.get(self, "enable_rbac")

    @property
    @pulumi.getter(name="enableStackdriverLogging")
    def enable_stackdriver_logging(self) -> pulumi.Output[bool]:
        """
        Option to enable Stackdriver Logging.
        """
        return pulumi.get(self, "enable_stackdriver_logging")

    @property
    @pulumi.getter(name="enableStackdriverMonitoring")
    def enable_stackdriver_monitoring(self) -> pulumi.Output[bool]:
        """
        Option to enable Stackdriver Monitoring.
        """
        return pulumi.get(self, "enable_stackdriver_monitoring")

    @property
    @pulumi.getter(name="gcsBucket")
    def gcs_bucket(self) -> pulumi.Output[str]:
        """
        Cloud Storage bucket generated by Data Fusion in the customer project.
        """
        return pulumi.get(self, "gcs_bucket")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output['outputs.NetworkConfigResponse']:
        """
        Network configuration options. These are required when a private Data Fusion instance is to be created.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of additional options used to configure the behavior of Data Fusion instance.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="p4ServiceAccount")
    def p4_service_account(self) -> pulumi.Output[str]:
        """
        P4 service account for the customer project.
        """
        return pulumi.get(self, "p4_service_account")

    @property
    @pulumi.getter(name="privateInstance")
    def private_instance(self) -> pulumi.Output[bool]:
        """
        Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
        """
        return pulumi.get(self, "private_instance")

    @property
    @pulumi.getter(name="serviceEndpoint")
    def service_endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint on which the Data Fusion UI is accessible.
        """
        return pulumi.get(self, "service_endpoint")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of this Data Fusion instance.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateMessage")
    def state_message(self) -> pulumi.Output[str]:
        """
        Additional information about the current state of this Data Fusion instance if available.
        """
        return pulumi.get(self, "state_message")

    @property
    @pulumi.getter(name="tenantProjectId")
    def tenant_project_id(self) -> pulumi.Output[str]:
        """
        The name of the tenant project.
        """
        return pulumi.get(self, "tenant_project_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Required. Instance type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time the instance was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Current version of Data Fusion.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
        """
        return pulumi.get(self, "zone")

