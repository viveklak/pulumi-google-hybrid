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
    'GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponse',
    'GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse',
    'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
    'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse',
    'GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse',
]

@pulumi.output_type
class GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponse(dict):
    """
    AcceleratorConfig defines the accelerator cards to attach to the VM.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "acceleratorCount":
            suggest = "accelerator_count"
        elif key == "acceleratorType":
            suggest = "accelerator_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 accelerator_count: str,
                 accelerator_type: str):
        """
        AcceleratorConfig defines the accelerator cards to attach to the VM.
        :param str accelerator_count: The number of guest accelerator cards exposed to each VM.
        :param str accelerator_type: The type of accelerator to attach to each VM, e.g. "nvidia-tesla-k80" for nVidia Tesla K80.
        """
        pulumi.set(__self__, "accelerator_count", accelerator_count)
        pulumi.set(__self__, "accelerator_type", accelerator_type)

    @property
    @pulumi.getter(name="acceleratorCount")
    def accelerator_count(self) -> str:
        """
        The number of guest accelerator cards exposed to each VM.
        """
        return pulumi.get(self, "accelerator_count")

    @property
    @pulumi.getter(name="acceleratorType")
    def accelerator_type(self) -> str:
        """
        The type of accelerator to attach to each VM, e.g. "nvidia-tesla-k80" for nVidia Tesla K80.
        """
        return pulumi.get(self, "accelerator_type")


@pulumi.output_type
class GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse(dict):
    """
    Autoscale defines the autoscaling policy of a worker pool.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxSize":
            suggest = "max_size"
        elif key == "minSize":
            suggest = "min_size"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_size: str,
                 min_size: str):
        """
        Autoscale defines the autoscaling policy of a worker pool.
        :param str max_size: The maximal number of workers. Must be equal to or greater than min_size.
        :param str min_size: The minimal number of workers. Must be greater than 0.
        """
        pulumi.set(__self__, "max_size", max_size)
        pulumi.set(__self__, "min_size", min_size)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> str:
        """
        The maximal number of workers. Must be equal to or greater than min_size.
        """
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> str:
        """
        The minimal number of workers. Must be greater than 0.
        """
        return pulumi.get(self, "min_size")


@pulumi.output_type
class GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse(dict):
    """
    Defines whether a feature can be used or what values are accepted.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedValues":
            suggest = "allowed_values"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed_values: Sequence[str],
                 policy: str):
        """
        Defines whether a feature can be used or what values are accepted.
        :param Sequence[str] allowed_values: A list of acceptable values. Only effective when the policy is `RESTRICTED`.
        :param str policy: The policy of the feature.
        """
        pulumi.set(__self__, "allowed_values", allowed_values)
        pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> Sequence[str]:
        """
        A list of acceptable values. Only effective when the policy is `RESTRICTED`.
        """
        return pulumi.get(self, "allowed_values")

    @property
    @pulumi.getter
    def policy(self) -> str:
        """
        The policy of the feature.
        """
        return pulumi.get(self, "policy")


@pulumi.output_type
class GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse(dict):
    """
    FeaturePolicy defines features allowed to be used on RBE instances, as well as instance-wide behavior changes that take effect without opt-in or opt-out at usage time.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "containerImageSources":
            suggest = "container_image_sources"
        elif key == "dockerAddCapabilities":
            suggest = "docker_add_capabilities"
        elif key == "dockerChrootPath":
            suggest = "docker_chroot_path"
        elif key == "dockerNetwork":
            suggest = "docker_network"
        elif key == "dockerPrivileged":
            suggest = "docker_privileged"
        elif key == "dockerRunAsRoot":
            suggest = "docker_run_as_root"
        elif key == "dockerRuntime":
            suggest = "docker_runtime"
        elif key == "dockerSiblingContainers":
            suggest = "docker_sibling_containers"
        elif key == "linuxIsolation":
            suggest = "linux_isolation"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 container_image_sources: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 docker_add_capabilities: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 docker_chroot_path: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 docker_network: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 docker_privileged: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 docker_run_as_root: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 docker_runtime: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 docker_sibling_containers: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse',
                 linux_isolation: str):
        """
        FeaturePolicy defines features allowed to be used on RBE instances, as well as instance-wide behavior changes that take effect without opt-in or opt-out at usage time.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' container_image_sources: Which container image sources are allowed. Currently only RBE-supported registry (gcr.io) is allowed. One can allow all repositories under a project or one specific repository only. E.g. container_image_sources { policy: RESTRICTED allowed_values: [ "gcr.io/project-foo", "gcr.io/project-bar/repo-baz", ] } will allow any repositories under "gcr.io/project-foo" plus the repository "gcr.io/project-bar/repo-baz". Default (UNSPECIFIED) is equivalent to any source is allowed.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' docker_add_capabilities: Whether dockerAddCapabilities can be used or what capabilities are allowed.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' docker_chroot_path: Whether dockerChrootPath can be used.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' docker_network: Whether dockerNetwork can be used or what network modes are allowed. E.g. one may allow `off` value only via `allowed_values`.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' docker_privileged: Whether dockerPrivileged can be used.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' docker_run_as_root: Whether dockerRunAsRoot can be used.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' docker_runtime: Whether dockerRuntime is allowed to be set or what runtimes are allowed. Note linux_isolation takes precedence, and if set, docker_runtime values may be rejected if they are incompatible with the selected isolation.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponseArgs' docker_sibling_containers: Whether dockerSiblingContainers can be used.
        :param str linux_isolation: linux_isolation allows overriding the docker runtime used for containers started on Linux.
        """
        pulumi.set(__self__, "container_image_sources", container_image_sources)
        pulumi.set(__self__, "docker_add_capabilities", docker_add_capabilities)
        pulumi.set(__self__, "docker_chroot_path", docker_chroot_path)
        pulumi.set(__self__, "docker_network", docker_network)
        pulumi.set(__self__, "docker_privileged", docker_privileged)
        pulumi.set(__self__, "docker_run_as_root", docker_run_as_root)
        pulumi.set(__self__, "docker_runtime", docker_runtime)
        pulumi.set(__self__, "docker_sibling_containers", docker_sibling_containers)
        pulumi.set(__self__, "linux_isolation", linux_isolation)

    @property
    @pulumi.getter(name="containerImageSources")
    def container_image_sources(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Which container image sources are allowed. Currently only RBE-supported registry (gcr.io) is allowed. One can allow all repositories under a project or one specific repository only. E.g. container_image_sources { policy: RESTRICTED allowed_values: [ "gcr.io/project-foo", "gcr.io/project-bar/repo-baz", ] } will allow any repositories under "gcr.io/project-foo" plus the repository "gcr.io/project-bar/repo-baz". Default (UNSPECIFIED) is equivalent to any source is allowed.
        """
        return pulumi.get(self, "container_image_sources")

    @property
    @pulumi.getter(name="dockerAddCapabilities")
    def docker_add_capabilities(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Whether dockerAddCapabilities can be used or what capabilities are allowed.
        """
        return pulumi.get(self, "docker_add_capabilities")

    @property
    @pulumi.getter(name="dockerChrootPath")
    def docker_chroot_path(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Whether dockerChrootPath can be used.
        """
        return pulumi.get(self, "docker_chroot_path")

    @property
    @pulumi.getter(name="dockerNetwork")
    def docker_network(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Whether dockerNetwork can be used or what network modes are allowed. E.g. one may allow `off` value only via `allowed_values`.
        """
        return pulumi.get(self, "docker_network")

    @property
    @pulumi.getter(name="dockerPrivileged")
    def docker_privileged(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Whether dockerPrivileged can be used.
        """
        return pulumi.get(self, "docker_privileged")

    @property
    @pulumi.getter(name="dockerRunAsRoot")
    def docker_run_as_root(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Whether dockerRunAsRoot can be used.
        """
        return pulumi.get(self, "docker_run_as_root")

    @property
    @pulumi.getter(name="dockerRuntime")
    def docker_runtime(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Whether dockerRuntime is allowed to be set or what runtimes are allowed. Note linux_isolation takes precedence, and if set, docker_runtime values may be rejected if they are incompatible with the selected isolation.
        """
        return pulumi.get(self, "docker_runtime")

    @property
    @pulumi.getter(name="dockerSiblingContainers")
    def docker_sibling_containers(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyFeatureResponse':
        """
        Whether dockerSiblingContainers can be used.
        """
        return pulumi.get(self, "docker_sibling_containers")

    @property
    @pulumi.getter(name="linuxIsolation")
    def linux_isolation(self) -> str:
        """
        linux_isolation allows overriding the docker runtime used for containers started on Linux.
        """
        return pulumi.get(self, "linux_isolation")


@pulumi.output_type
class GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse(dict):
    """
    Defines the configuration to be used for creating workers in the worker pool.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "diskSizeGb":
            suggest = "disk_size_gb"
        elif key == "diskType":
            suggest = "disk_type"
        elif key == "machineType":
            suggest = "machine_type"
        elif key == "maxConcurrentActions":
            suggest = "max_concurrent_actions"
        elif key == "minCpuPlatform":
            suggest = "min_cpu_platform"
        elif key == "networkAccess":
            suggest = "network_access"
        elif key == "soleTenantNodeType":
            suggest = "sole_tenant_node_type"
        elif key == "vmImage":
            suggest = "vm_image"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 accelerator: 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponse',
                 disk_size_gb: str,
                 disk_type: str,
                 labels: Mapping[str, str],
                 machine_type: str,
                 max_concurrent_actions: str,
                 min_cpu_platform: str,
                 network_access: str,
                 reserved: bool,
                 sole_tenant_node_type: str,
                 vm_image: str):
        """
        Defines the configuration to be used for creating workers in the worker pool.
        :param 'GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponseArgs' accelerator: The accelerator card attached to each VM.
        :param str disk_size_gb: Required. Size of the disk attached to the worker, in GB. See https://cloud.google.com/compute/docs/disks/
        :param str disk_type: Required. Disk Type to use for the worker. See [Storage options](https://cloud.google.com/compute/docs/disks/#introduction). Currently only `pd-standard` and `pd-ssd` are supported.
        :param Mapping[str, str] labels: Labels associated with the workers. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International letters are permitted. Label keys must start with a letter. Label values are optional. There can not be more than 64 labels per resource.
        :param str machine_type: Required. Machine type of the worker, such as `e2-standard-2`. See https://cloud.google.com/compute/docs/machine-types for a list of supported machine types. Note that `f1-micro` and `g1-small` are not yet supported.
        :param str max_concurrent_actions: The maximum number of actions a worker can execute concurrently.
        :param str min_cpu_platform: Minimum CPU platform to use when creating the worker. See [CPU Platforms](https://cloud.google.com/compute/docs/cpu-platforms).
        :param str network_access: Determines the type of network access granted to workers. Possible values: - "public": Workers can connect to the public internet. - "private": Workers can only connect to Google APIs and services. - "restricted-private": Workers can only connect to Google APIs that are reachable through `restricted.googleapis.com` (`199.36.153.4/30`).
        :param bool reserved: Determines whether the worker is reserved (equivalent to a Compute Engine on-demand VM and therefore won't be preempted). See [Preemptible VMs](https://cloud.google.com/preemptible-vms/) for more details.
        :param str sole_tenant_node_type: The node type name to be used for sole-tenant nodes.
        :param str vm_image: The name of the image used by each VM.
        """
        pulumi.set(__self__, "accelerator", accelerator)
        pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        pulumi.set(__self__, "disk_type", disk_type)
        pulumi.set(__self__, "labels", labels)
        pulumi.set(__self__, "machine_type", machine_type)
        pulumi.set(__self__, "max_concurrent_actions", max_concurrent_actions)
        pulumi.set(__self__, "min_cpu_platform", min_cpu_platform)
        pulumi.set(__self__, "network_access", network_access)
        pulumi.set(__self__, "reserved", reserved)
        pulumi.set(__self__, "sole_tenant_node_type", sole_tenant_node_type)
        pulumi.set(__self__, "vm_image", vm_image)

    @property
    @pulumi.getter
    def accelerator(self) -> 'outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigResponse':
        """
        The accelerator card attached to each VM.
        """
        return pulumi.get(self, "accelerator")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> str:
        """
        Required. Size of the disk attached to the worker, in GB. See https://cloud.google.com/compute/docs/disks/
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> str:
        """
        Required. Disk Type to use for the worker. See [Storage options](https://cloud.google.com/compute/docs/disks/#introduction). Currently only `pd-standard` and `pd-ssd` are supported.
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels associated with the workers. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International letters are permitted. Label keys must start with a letter. Label values are optional. There can not be more than 64 labels per resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> str:
        """
        Required. Machine type of the worker, such as `e2-standard-2`. See https://cloud.google.com/compute/docs/machine-types for a list of supported machine types. Note that `f1-micro` and `g1-small` are not yet supported.
        """
        return pulumi.get(self, "machine_type")

    @property
    @pulumi.getter(name="maxConcurrentActions")
    def max_concurrent_actions(self) -> str:
        """
        The maximum number of actions a worker can execute concurrently.
        """
        return pulumi.get(self, "max_concurrent_actions")

    @property
    @pulumi.getter(name="minCpuPlatform")
    def min_cpu_platform(self) -> str:
        """
        Minimum CPU platform to use when creating the worker. See [CPU Platforms](https://cloud.google.com/compute/docs/cpu-platforms).
        """
        return pulumi.get(self, "min_cpu_platform")

    @property
    @pulumi.getter(name="networkAccess")
    def network_access(self) -> str:
        """
        Determines the type of network access granted to workers. Possible values: - "public": Workers can connect to the public internet. - "private": Workers can only connect to Google APIs and services. - "restricted-private": Workers can only connect to Google APIs that are reachable through `restricted.googleapis.com` (`199.36.153.4/30`).
        """
        return pulumi.get(self, "network_access")

    @property
    @pulumi.getter
    def reserved(self) -> bool:
        """
        Determines whether the worker is reserved (equivalent to a Compute Engine on-demand VM and therefore won't be preempted). See [Preemptible VMs](https://cloud.google.com/preemptible-vms/) for more details.
        """
        return pulumi.get(self, "reserved")

    @property
    @pulumi.getter(name="soleTenantNodeType")
    def sole_tenant_node_type(self) -> str:
        """
        The node type name to be used for sole-tenant nodes.
        """
        return pulumi.get(self, "sole_tenant_node_type")

    @property
    @pulumi.getter(name="vmImage")
    def vm_image(self) -> str:
        """
        The name of the image used by each VM.
        """
        return pulumi.get(self, "vm_image")


