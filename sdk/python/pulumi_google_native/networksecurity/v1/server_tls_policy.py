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

__all__ = ['ServerTlsPolicyArgs', 'ServerTlsPolicy']

@pulumi.input_type
class ServerTlsPolicyArgs:
    def __init__(__self__, *,
                 server_tls_policy_id: pulumi.Input[str],
                 allow_open: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mtls_policy: Optional[pulumi.Input['MTLSPolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 server_certificate: Optional[pulumi.Input['GoogleCloudNetworksecurityV1CertificateProviderArgs']] = None):
        """
        The set of arguments for constructing a ServerTlsPolicy resource.
        :param pulumi.Input[str] server_tls_policy_id: Required. Short name of the ServerTlsPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "server_mtls_policy".
        :param pulumi.Input[bool] allow_open:  Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
        :param pulumi.Input[str] description: Free-text description of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of label tags associated with the resource.
        :param pulumi.Input['MTLSPolicyArgs'] mtls_policy:  Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        :param pulumi.Input[str] name: Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        :param pulumi.Input['GoogleCloudNetworksecurityV1CertificateProviderArgs'] server_certificate:  Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        """
        pulumi.set(__self__, "server_tls_policy_id", server_tls_policy_id)
        if allow_open is not None:
            pulumi.set(__self__, "allow_open", allow_open)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if mtls_policy is not None:
            pulumi.set(__self__, "mtls_policy", mtls_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if server_certificate is not None:
            pulumi.set(__self__, "server_certificate", server_certificate)

    @property
    @pulumi.getter(name="serverTlsPolicyId")
    def server_tls_policy_id(self) -> pulumi.Input[str]:
        """
        Required. Short name of the ServerTlsPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "server_mtls_policy".
        """
        return pulumi.get(self, "server_tls_policy_id")

    @server_tls_policy_id.setter
    def server_tls_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_tls_policy_id", value)

    @property
    @pulumi.getter(name="allowOpen")
    def allow_open(self) -> Optional[pulumi.Input[bool]]:
        """
         Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
        """
        return pulumi.get(self, "allow_open")

    @allow_open.setter
    def allow_open(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_open", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Free-text description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set of label tags associated with the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="mtlsPolicy")
    def mtls_policy(self) -> Optional[pulumi.Input['MTLSPolicyArgs']]:
        """
         Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        """
        return pulumi.get(self, "mtls_policy")

    @mtls_policy.setter
    def mtls_policy(self, value: Optional[pulumi.Input['MTLSPolicyArgs']]):
        pulumi.set(self, "mtls_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> Optional[pulumi.Input['GoogleCloudNetworksecurityV1CertificateProviderArgs']]:
        """
         Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        """
        return pulumi.get(self, "server_certificate")

    @server_certificate.setter
    def server_certificate(self, value: Optional[pulumi.Input['GoogleCloudNetworksecurityV1CertificateProviderArgs']]):
        pulumi.set(self, "server_certificate", value)


class ServerTlsPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_open: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mtls_policy: Optional[pulumi.Input[pulumi.InputType['MTLSPolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 server_certificate: Optional[pulumi.Input[pulumi.InputType['GoogleCloudNetworksecurityV1CertificateProviderArgs']]] = None,
                 server_tls_policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new ServerTlsPolicy in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_open:  Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
        :param pulumi.Input[str] description: Free-text description of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of label tags associated with the resource.
        :param pulumi.Input[pulumi.InputType['MTLSPolicyArgs']] mtls_policy:  Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        :param pulumi.Input[str] name: Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        :param pulumi.Input[pulumi.InputType['GoogleCloudNetworksecurityV1CertificateProviderArgs']] server_certificate:  Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        :param pulumi.Input[str] server_tls_policy_id: Required. Short name of the ServerTlsPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "server_mtls_policy".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerTlsPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new ServerTlsPolicy in a given project and location.

        :param str resource_name: The name of the resource.
        :param ServerTlsPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerTlsPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_open: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mtls_policy: Optional[pulumi.Input[pulumi.InputType['MTLSPolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 server_certificate: Optional[pulumi.Input[pulumi.InputType['GoogleCloudNetworksecurityV1CertificateProviderArgs']]] = None,
                 server_tls_policy_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServerTlsPolicyArgs.__new__(ServerTlsPolicyArgs)

            __props__.__dict__["allow_open"] = allow_open
            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["mtls_policy"] = mtls_policy
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["server_certificate"] = server_certificate
            if server_tls_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_tls_policy_id'")
            __props__.__dict__["server_tls_policy_id"] = server_tls_policy_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(ServerTlsPolicy, __self__).__init__(
            'google-native:networksecurity/v1:ServerTlsPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServerTlsPolicy':
        """
        Get an existing ServerTlsPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServerTlsPolicyArgs.__new__(ServerTlsPolicyArgs)

        __props__.__dict__["allow_open"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["mtls_policy"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["server_certificate"] = None
        __props__.__dict__["update_time"] = None
        return ServerTlsPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowOpen")
    def allow_open(self) -> pulumi.Output[bool]:
        """
         Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
        """
        return pulumi.get(self, "allow_open")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Free-text description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Set of label tags associated with the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mtlsPolicy")
    def mtls_policy(self) -> pulumi.Output['outputs.MTLSPolicyResponse']:
        """
         Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        """
        return pulumi.get(self, "mtls_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> pulumi.Output['outputs.GoogleCloudNetworksecurityV1CertificateProviderResponse']:
        """
         Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        """
        return pulumi.get(self, "server_certificate")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the resource was updated.
        """
        return pulumi.get(self, "update_time")

