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

__all__ = ['RealmGameServerClusterArgs', 'RealmGameServerCluster']

@pulumi.input_type
class RealmGameServerClusterArgs:
    def __init__(__self__, *,
                 game_server_clusters_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 realms_id: pulumi.Input[str],
                 connection_info: Optional[pulumi.Input['GameServerClusterConnectionInfoArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RealmGameServerCluster resource.
        :param pulumi.Input['GameServerClusterConnectionInfoArgs'] connection_info: The game server cluster connection information. This information is used to manage game server clusters.
        :param pulumi.Input[str] description: Human readable description of the cluster.
        :param pulumi.Input[str] etag: ETag of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this game server cluster. Each label is a key-value pair.
        :param pulumi.Input[str] name: Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
        """
        pulumi.set(__self__, "game_server_clusters_id", game_server_clusters_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        pulumi.set(__self__, "realms_id", realms_id)
        if connection_info is not None:
            pulumi.set(__self__, "connection_info", connection_info)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="gameServerClustersId")
    def game_server_clusters_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "game_server_clusters_id")

    @game_server_clusters_id.setter
    def game_server_clusters_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "game_server_clusters_id", value)

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
    @pulumi.getter(name="realmsId")
    def realms_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "realms_id")

    @realms_id.setter
    def realms_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realms_id", value)

    @property
    @pulumi.getter(name="connectionInfo")
    def connection_info(self) -> Optional[pulumi.Input['GameServerClusterConnectionInfoArgs']]:
        """
        The game server cluster connection information. This information is used to manage game server clusters.
        """
        return pulumi.get(self, "connection_info")

    @connection_info.setter
    def connection_info(self, value: Optional[pulumi.Input['GameServerClusterConnectionInfoArgs']]):
        pulumi.set(self, "connection_info", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable description of the cluster.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        ETag of the resource.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels associated with this game server cluster. Each label is a key-value pair.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class RealmGameServerCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_info: Optional[pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 game_server_clusters_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 realms_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new game server cluster in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']] connection_info: The game server cluster connection information. This information is used to manage game server clusters.
        :param pulumi.Input[str] description: Human readable description of the cluster.
        :param pulumi.Input[str] etag: ETag of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this game server cluster. Each label is a key-value pair.
        :param pulumi.Input[str] name: Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealmGameServerClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new game server cluster in a given project and location.

        :param str resource_name: The name of the resource.
        :param RealmGameServerClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealmGameServerClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_info: Optional[pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 game_server_clusters_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 realms_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RealmGameServerClusterArgs.__new__(RealmGameServerClusterArgs)

            __props__.__dict__["connection_info"] = connection_info
            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            if game_server_clusters_id is None and not opts.urn:
                raise TypeError("Missing required property 'game_server_clusters_id'")
            __props__.__dict__["game_server_clusters_id"] = game_server_clusters_id
            __props__.__dict__["labels"] = labels
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["name"] = name
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            if realms_id is None and not opts.urn:
                raise TypeError("Missing required property 'realms_id'")
            __props__.__dict__["realms_id"] = realms_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(RealmGameServerCluster, __self__).__init__(
            'gcp-native:gameservices/v1:RealmGameServerCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RealmGameServerCluster':
        """
        Get an existing RealmGameServerCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RealmGameServerClusterArgs.__new__(RealmGameServerClusterArgs)

        __props__.__dict__["connection_info"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["update_time"] = None
        return RealmGameServerCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionInfo")
    def connection_info(self) -> pulumi.Output['outputs.GameServerClusterConnectionInfoResponse']:
        """
        The game server cluster connection information. This information is used to manage game server clusters.
        """
        return pulumi.get(self, "connection_info")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Human readable description of the cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        ETag of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The labels associated with this game server cluster. Each label is a key-value pair.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last-modified time.
        """
        return pulumi.get(self, "update_time")

