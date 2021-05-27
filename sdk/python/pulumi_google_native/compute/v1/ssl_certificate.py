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

__all__ = ['SslCertificateArgs', 'SslCertificate']

@pulumi.input_type
class SslCertificateArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 certificate: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input['SslCertificateManagedSslCertificateArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_managed: Optional[pulumi.Input['SslCertificateSelfManagedSslCertificateArgs']] = None,
                 subject_alternative_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SslCertificate resource.
        :param pulumi.Input[str] certificate: A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] expire_time: [Output Only] Expire time of the certificate. RFC3339
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#sslCertificate for SSL certificates.
        :param pulumi.Input['SslCertificateManagedSslCertificateArgs'] managed: Configuration and status of a managed SSL certificate.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] private_key: A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
        :param pulumi.Input[str] self_link: [Output only] Server-defined URL for the resource.
        :param pulumi.Input['SslCertificateSelfManagedSslCertificateArgs'] self_managed: Configuration and status of a self-managed SSL certificate.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subject_alternative_names: [Output Only] Domains associated with the certificate via Subject Alternative Name.
        :param pulumi.Input[str] type: (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
        """
        pulumi.set(__self__, "project", project)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if managed is not None:
            pulumi.set(__self__, "managed", managed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if self_managed is not None:
            pulumi.set(__self__, "self_managed", self_managed)
        if subject_alternative_names is not None:
            pulumi.set(__self__, "subject_alternative_names", subject_alternative_names)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Expire time of the certificate. RFC3339
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Type of the resource. Always compute#sslCertificate for SSL certificates.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def managed(self) -> Optional[pulumi.Input['SslCertificateManagedSslCertificateArgs']]:
        """
        Configuration and status of a managed SSL certificate.
        """
        return pulumi.get(self, "managed")

    @managed.setter
    def managed(self, value: Optional[pulumi.Input['SslCertificateManagedSslCertificateArgs']]):
        pulumi.set(self, "managed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="selfManaged")
    def self_managed(self) -> Optional[pulumi.Input['SslCertificateSelfManagedSslCertificateArgs']]:
        """
        Configuration and status of a self-managed SSL certificate.
        """
        return pulumi.get(self, "self_managed")

    @self_managed.setter
    def self_managed(self, value: Optional[pulumi.Input['SslCertificateSelfManagedSslCertificateArgs']]):
        pulumi.set(self, "self_managed", value)

    @property
    @pulumi.getter(name="subjectAlternativeNames")
    def subject_alternative_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        [Output Only] Domains associated with the certificate via Subject Alternative Name.
        """
        return pulumi.get(self, "subject_alternative_names")

    @subject_alternative_names.setter
    def subject_alternative_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subject_alternative_names", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class SslCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[pulumi.InputType['SslCertificateManagedSslCertificateArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_managed: Optional[pulumi.Input[pulumi.InputType['SslCertificateSelfManagedSslCertificateArgs']]] = None,
                 subject_alternative_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a SslCertificate resource in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] expire_time: [Output Only] Expire time of the certificate. RFC3339
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#sslCertificate for SSL certificates.
        :param pulumi.Input[pulumi.InputType['SslCertificateManagedSslCertificateArgs']] managed: Configuration and status of a managed SSL certificate.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] private_key: A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
        :param pulumi.Input[str] self_link: [Output only] Server-defined URL for the resource.
        :param pulumi.Input[pulumi.InputType['SslCertificateSelfManagedSslCertificateArgs']] self_managed: Configuration and status of a self-managed SSL certificate.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subject_alternative_names: [Output Only] Domains associated with the certificate via Subject Alternative Name.
        :param pulumi.Input[str] type: (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SslCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a SslCertificate resource in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param SslCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SslCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[pulumi.InputType['SslCertificateManagedSslCertificateArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_managed: Optional[pulumi.Input[pulumi.InputType['SslCertificateSelfManagedSslCertificateArgs']]] = None,
                 subject_alternative_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = SslCertificateArgs.__new__(SslCertificateArgs)

            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["description"] = description
            __props__.__dict__["expire_time"] = expire_time
            __props__.__dict__["id"] = id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["managed"] = managed
            __props__.__dict__["name"] = name
            __props__.__dict__["private_key"] = private_key
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["self_managed"] = self_managed
            __props__.__dict__["subject_alternative_names"] = subject_alternative_names
            __props__.__dict__["type"] = type
        super(SslCertificate, __self__).__init__(
            'google-native:compute/v1:SslCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SslCertificate':
        """
        Get an existing SslCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SslCertificateArgs.__new__(SslCertificateArgs)

        __props__.__dict__["certificate"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["expire_time"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["managed"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_key"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_managed"] = None
        __props__.__dict__["subject_alternative_names"] = None
        __props__.__dict__["type"] = None
        return SslCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        [Output Only] Expire time of the certificate. RFC3339
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output Only] Type of the resource. Always compute#sslCertificate for SSL certificates.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def managed(self) -> pulumi.Output['outputs.SslCertificateManagedSslCertificateResponse']:
        """
        Configuration and status of a managed SSL certificate.
        """
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        [Output Only] URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfManaged")
    def self_managed(self) -> pulumi.Output['outputs.SslCertificateSelfManagedSslCertificateResponse']:
        """
        Configuration and status of a self-managed SSL certificate.
        """
        return pulumi.get(self, "self_managed")

    @property
    @pulumi.getter(name="subjectAlternativeNames")
    def subject_alternative_names(self) -> pulumi.Output[Sequence[str]]:
        """
        [Output Only] Domains associated with the certificate via Subject Alternative Name.
        """
        return pulumi.get(self, "subject_alternative_names")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
        """
        return pulumi.get(self, "type")

