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

__all__ = ['WorkloadIdentityPoolProviderArgs', 'WorkloadIdentityPoolProvider']

@pulumi.input_type
class WorkloadIdentityPoolProviderArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 workload_identity_pool_id: pulumi.Input[str],
                 workload_identity_pool_provider_id: pulumi.Input[str],
                 attribute_condition: Optional[pulumi.Input[str]] = None,
                 attribute_mapping: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 aws: Optional[pulumi.Input['AwsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 oidc: Optional[pulumi.Input['OidcArgs']] = None):
        """
        The set of arguments for constructing a WorkloadIdentityPoolProvider resource.
        :param pulumi.Input[str] attribute_condition: [A Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credential are accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attribute_mapping: Maps attributes from authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. Cannot exceed 127 characters. * `google.groups`: Groups the external identity belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where `{custom_attribute}` is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workload to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 8KB. For AWS providers, if no attribute mapping is defined, the following default mapping applies: ``` { "google.subject":"assertion.arn", "attribute.aws_role": "assertion.arn.contains('assumed-role')" " ? assertion.arn.extract('{account_arn}assumed-role/')" " + 'assumed-role/'" " + assertion.arn.extract('assumed-role/{role_name}/')" " : assertion.arn", } ``` If any custom attribute mappings are defined, they must include a mapping to the `google.subject` attribute. For OIDC providers, you must supply a custom mapping, which must include the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        :param pulumi.Input['AwsArgs'] aws: An Amazon Web Services identity provider.
        :param pulumi.Input[str] description: A description for the provider. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        :param pulumi.Input[str] display_name: A display name for the provider. Cannot exceed 32 characters.
        :param pulumi.Input['OidcArgs'] oidc: An OpenId Connect 1.0 identity provider.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "workload_identity_pool_id", workload_identity_pool_id)
        pulumi.set(__self__, "workload_identity_pool_provider_id", workload_identity_pool_provider_id)
        if attribute_condition is not None:
            pulumi.set(__self__, "attribute_condition", attribute_condition)
        if attribute_mapping is not None:
            pulumi.set(__self__, "attribute_mapping", attribute_mapping)
        if aws is not None:
            pulumi.set(__self__, "aws", aws)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if oidc is not None:
            pulumi.set(__self__, "oidc", oidc)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="workloadIdentityPoolId")
    def workload_identity_pool_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workload_identity_pool_id")

    @workload_identity_pool_id.setter
    def workload_identity_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workload_identity_pool_id", value)

    @property
    @pulumi.getter(name="workloadIdentityPoolProviderId")
    def workload_identity_pool_provider_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workload_identity_pool_provider_id")

    @workload_identity_pool_provider_id.setter
    def workload_identity_pool_provider_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workload_identity_pool_provider_id", value)

    @property
    @pulumi.getter(name="attributeCondition")
    def attribute_condition(self) -> Optional[pulumi.Input[str]]:
        """
        [A Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credential are accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        """
        return pulumi.get(self, "attribute_condition")

    @attribute_condition.setter
    def attribute_condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_condition", value)

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Maps attributes from authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. Cannot exceed 127 characters. * `google.groups`: Groups the external identity belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where `{custom_attribute}` is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workload to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 8KB. For AWS providers, if no attribute mapping is defined, the following default mapping applies: ``` { "google.subject":"assertion.arn", "attribute.aws_role": "assertion.arn.contains('assumed-role')" " ? assertion.arn.extract('{account_arn}assumed-role/')" " + 'assumed-role/'" " + assertion.arn.extract('assumed-role/{role_name}/')" " : assertion.arn", } ``` If any custom attribute mappings are defined, they must include a mapping to the `google.subject` attribute. For OIDC providers, you must supply a custom mapping, which must include the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        """
        return pulumi.get(self, "attribute_mapping")

    @attribute_mapping.setter
    def attribute_mapping(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "attribute_mapping", value)

    @property
    @pulumi.getter
    def aws(self) -> Optional[pulumi.Input['AwsArgs']]:
        """
        An Amazon Web Services identity provider.
        """
        return pulumi.get(self, "aws")

    @aws.setter
    def aws(self, value: Optional[pulumi.Input['AwsArgs']]):
        pulumi.set(self, "aws", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the provider. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A display name for the provider. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def oidc(self) -> Optional[pulumi.Input['OidcArgs']]:
        """
        An OpenId Connect 1.0 identity provider.
        """
        return pulumi.get(self, "oidc")

    @oidc.setter
    def oidc(self, value: Optional[pulumi.Input['OidcArgs']]):
        pulumi.set(self, "oidc", value)


class WorkloadIdentityPoolProvider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_condition: Optional[pulumi.Input[str]] = None,
                 attribute_mapping: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 aws: Optional[pulumi.Input[pulumi.InputType['AwsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 oidc: Optional[pulumi.Input[pulumi.InputType['OidcArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_provider_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new WorkloadIdentityPoolProvider in a WorkloadIdentityPool. You cannot reuse the name of a deleted provider until 30 days after deletion.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_condition: [A Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credential are accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attribute_mapping: Maps attributes from authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. Cannot exceed 127 characters. * `google.groups`: Groups the external identity belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where `{custom_attribute}` is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workload to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 8KB. For AWS providers, if no attribute mapping is defined, the following default mapping applies: ``` { "google.subject":"assertion.arn", "attribute.aws_role": "assertion.arn.contains('assumed-role')" " ? assertion.arn.extract('{account_arn}assumed-role/')" " + 'assumed-role/'" " + assertion.arn.extract('assumed-role/{role_name}/')" " : assertion.arn", } ``` If any custom attribute mappings are defined, they must include a mapping to the `google.subject` attribute. For OIDC providers, you must supply a custom mapping, which must include the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        :param pulumi.Input[pulumi.InputType['AwsArgs']] aws: An Amazon Web Services identity provider.
        :param pulumi.Input[str] description: A description for the provider. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        :param pulumi.Input[str] display_name: A display name for the provider. Cannot exceed 32 characters.
        :param pulumi.Input[pulumi.InputType['OidcArgs']] oidc: An OpenId Connect 1.0 identity provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkloadIdentityPoolProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new WorkloadIdentityPoolProvider in a WorkloadIdentityPool. You cannot reuse the name of a deleted provider until 30 days after deletion.

        :param str resource_name: The name of the resource.
        :param WorkloadIdentityPoolProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkloadIdentityPoolProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_condition: Optional[pulumi.Input[str]] = None,
                 attribute_mapping: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 aws: Optional[pulumi.Input[pulumi.InputType['AwsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 oidc: Optional[pulumi.Input[pulumi.InputType['OidcArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_provider_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = WorkloadIdentityPoolProviderArgs.__new__(WorkloadIdentityPoolProviderArgs)

            __props__.__dict__["attribute_condition"] = attribute_condition
            __props__.__dict__["attribute_mapping"] = attribute_mapping
            __props__.__dict__["aws"] = aws
            __props__.__dict__["description"] = description
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["display_name"] = display_name
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["oidc"] = oidc
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if workload_identity_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'workload_identity_pool_id'")
            __props__.__dict__["workload_identity_pool_id"] = workload_identity_pool_id
            if workload_identity_pool_provider_id is None and not opts.urn:
                raise TypeError("Missing required property 'workload_identity_pool_provider_id'")
            __props__.__dict__["workload_identity_pool_provider_id"] = workload_identity_pool_provider_id
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
        super(WorkloadIdentityPoolProvider, __self__).__init__(
            'google-native:iam/v1:WorkloadIdentityPoolProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkloadIdentityPoolProvider':
        """
        Get an existing WorkloadIdentityPoolProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkloadIdentityPoolProviderArgs.__new__(WorkloadIdentityPoolProviderArgs)

        __props__.__dict__["attribute_condition"] = None
        __props__.__dict__["attribute_mapping"] = None
        __props__.__dict__["aws"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["disabled"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["oidc"] = None
        __props__.__dict__["state"] = None
        return WorkloadIdentityPoolProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attributeCondition")
    def attribute_condition(self) -> pulumi.Output[str]:
        """
        [A Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credential are accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ```
        """
        return pulumi.get(self, "attribute_condition")

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Maps attributes from authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. Cannot exceed 127 characters. * `google.groups`: Groups the external identity belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where `{custom_attribute}` is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workload to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 8KB. For AWS providers, if no attribute mapping is defined, the following default mapping applies: ``` { "google.subject":"assertion.arn", "attribute.aws_role": "assertion.arn.contains('assumed-role')" " ? assertion.arn.extract('{account_arn}assumed-role/')" " + 'assumed-role/'" " + assertion.arn.extract('assumed-role/{role_name}/')" " : assertion.arn", } ``` If any custom attribute mappings are defined, they must include a mapping to the `google.subject` attribute. For OIDC providers, you must supply a custom mapping, which must include the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {"google.subject": "assertion.sub"} ```
        """
        return pulumi.get(self, "attribute_mapping")

    @property
    @pulumi.getter
    def aws(self) -> pulumi.Output['outputs.AwsResponse']:
        """
        An Amazon Web Services identity provider.
        """
        return pulumi.get(self, "aws")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description for the provider. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A display name for the provider. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the provider.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def oidc(self) -> pulumi.Output['outputs.OidcResponse']:
        """
        An OpenId Connect 1.0 identity provider.
        """
        return pulumi.get(self, "oidc")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the provider.
        """
        return pulumi.get(self, "state")

