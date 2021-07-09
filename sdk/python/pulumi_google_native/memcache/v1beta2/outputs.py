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

__all__ = [
    'InstanceMessageResponse',
    'MemcacheParametersResponse',
    'NodeConfigResponse',
    'NodeResponse',
]

@pulumi.output_type
class InstanceMessageResponse(dict):
    def __init__(__self__, *,
                 code: str,
                 message: str):
        """
        :param str code: A code that correspond to one type of user-facing message.
        :param str message: Message on memcached instance which will be exposed to users.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        A code that correspond to one type of user-facing message.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        Message on memcached instance which will be exposed to users.
        """
        return pulumi.get(self, "message")


@pulumi.output_type
class MemcacheParametersResponse(dict):
    """
    The unique ID associated with this set of parameters. Users can use this id to determine if the parameters associated with the instance differ from the parameters associated with the nodes. A discrepancy between parameter ids can inform users that they may need to take action to apply parameters on nodes.
    """
    def __init__(__self__, *,
                 params: Mapping[str, str]):
        """
        The unique ID associated with this set of parameters. Users can use this id to determine if the parameters associated with the instance differ from the parameters associated with the nodes. A discrepancy between parameter ids can inform users that they may need to take action to apply parameters on nodes.
        :param Mapping[str, str] params: User defined set of parameters to use in the memcached process.
        """
        pulumi.set(__self__, "params", params)

    @property
    @pulumi.getter
    def params(self) -> Mapping[str, str]:
        """
        User defined set of parameters to use in the memcached process.
        """
        return pulumi.get(self, "params")


@pulumi.output_type
class NodeConfigResponse(dict):
    """
    Configuration for a Memcached Node.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cpuCount":
            suggest = "cpu_count"
        elif key == "memorySizeMb":
            suggest = "memory_size_mb"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NodeConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NodeConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NodeConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cpu_count: int,
                 memory_size_mb: int):
        """
        Configuration for a Memcached Node.
        :param int cpu_count: Number of cpus per Memcached node.
        :param int memory_size_mb: Memory size in MiB for each Memcached node.
        """
        pulumi.set(__self__, "cpu_count", cpu_count)
        pulumi.set(__self__, "memory_size_mb", memory_size_mb)

    @property
    @pulumi.getter(name="cpuCount")
    def cpu_count(self) -> int:
        """
        Number of cpus per Memcached node.
        """
        return pulumi.get(self, "cpu_count")

    @property
    @pulumi.getter(name="memorySizeMb")
    def memory_size_mb(self) -> int:
        """
        Memory size in MiB for each Memcached node.
        """
        return pulumi.get(self, "memory_size_mb")


@pulumi.output_type
class NodeResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nodeId":
            suggest = "node_id"
        elif key == "updateAvailable":
            suggest = "update_available"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NodeResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NodeResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NodeResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 host: str,
                 node_id: str,
                 parameters: 'outputs.MemcacheParametersResponse',
                 port: int,
                 state: str,
                 update_available: bool,
                 zone: str):
        """
        :param str host: Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
        :param str node_id: Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
        :param 'MemcacheParametersResponse' parameters: User defined parameters currently applied to the node.
        :param int port: The port number of the Memcached server on this node.
        :param str state: Current state of the Memcached node.
        :param bool update_available: Returns true if there is an update waiting to be applied
        :param str zone: Location (GCP Zone) for the Memcached node.
        """
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "node_id", node_id)
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "state", state)
        pulumi.set(__self__, "update_available", update_available)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> str:
        """
        Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
        """
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter
    def parameters(self) -> 'outputs.MemcacheParametersResponse':
        """
        User defined parameters currently applied to the node.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port number of the Memcached server on this node.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the Memcached node.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateAvailable")
    def update_available(self) -> bool:
        """
        Returns true if there is an update waiting to be applied
        """
        return pulumi.get(self, "update_available")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        Location (GCP Zone) for the Memcached node.
        """
        return pulumi.get(self, "zone")


