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
    'AuditConfigResponse',
    'AuditLogConfigResponse',
    'AwsResponse',
    'BindingResponse',
    'ExprResponse',
    'OidcResponse',
]

@pulumi.output_type
class AuditConfigResponse(dict):
    """
    Specifies the audit configuration for a service. The configuration determines which permission types are logged, and what identities, if any, are exempted from logging. An AuditConfig must have one or more AuditLogConfigs. If there are AuditConfigs for both `allServices` and a specific service, the union of the two AuditConfigs is used for that service: the log_types specified in each AuditConfig are enabled, and the exempted_members in each AuditLogConfig are exempted. Example Policy with multiple AuditConfigs: { "audit_configs": [ { "service": "allServices", "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" }, { "log_type": "ADMIN_READ" } ] }, { "service": "sampleservice.googleapis.com", "audit_log_configs": [ { "log_type": "DATA_READ" }, { "log_type": "DATA_WRITE", "exempted_members": [ "user:aliya@example.com" ] } ] } ] } For sampleservice, this policy enables DATA_READ, DATA_WRITE and ADMIN_READ logging. It also exempts jose@example.com from DATA_READ logging, and aliya@example.com from DATA_WRITE logging.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "auditLogConfigs":
            suggest = "audit_log_configs"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuditConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuditConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuditConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 audit_log_configs: Sequence['outputs.AuditLogConfigResponse'],
                 service: str):
        """
        Specifies the audit configuration for a service. The configuration determines which permission types are logged, and what identities, if any, are exempted from logging. An AuditConfig must have one or more AuditLogConfigs. If there are AuditConfigs for both `allServices` and a specific service, the union of the two AuditConfigs is used for that service: the log_types specified in each AuditConfig are enabled, and the exempted_members in each AuditLogConfig are exempted. Example Policy with multiple AuditConfigs: { "audit_configs": [ { "service": "allServices", "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" }, { "log_type": "ADMIN_READ" } ] }, { "service": "sampleservice.googleapis.com", "audit_log_configs": [ { "log_type": "DATA_READ" }, { "log_type": "DATA_WRITE", "exempted_members": [ "user:aliya@example.com" ] } ] } ] } For sampleservice, this policy enables DATA_READ, DATA_WRITE and ADMIN_READ logging. It also exempts jose@example.com from DATA_READ logging, and aliya@example.com from DATA_WRITE logging.
        :param Sequence['AuditLogConfigResponse'] audit_log_configs: The configuration for logging of each type of permission.
        :param str service: Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        pulumi.set(__self__, "audit_log_configs", audit_log_configs)
        pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter(name="auditLogConfigs")
    def audit_log_configs(self) -> Sequence['outputs.AuditLogConfigResponse']:
        """
        The configuration for logging of each type of permission.
        """
        return pulumi.get(self, "audit_log_configs")

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        return pulumi.get(self, "service")


@pulumi.output_type
class AuditLogConfigResponse(dict):
    """
    Provides the configuration for logging a type of permissions. Example: { "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" } ] } This enables 'DATA_READ' and 'DATA_WRITE' logging, while exempting jose@example.com from DATA_READ logging.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "exemptedMembers":
            suggest = "exempted_members"
        elif key == "logType":
            suggest = "log_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuditLogConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuditLogConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuditLogConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 exempted_members: Sequence[str],
                 log_type: str):
        """
        Provides the configuration for logging a type of permissions. Example: { "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" } ] } This enables 'DATA_READ' and 'DATA_WRITE' logging, while exempting jose@example.com from DATA_READ logging.
        :param Sequence[str] exempted_members: Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
        :param str log_type: The log type that this config enables.
        """
        pulumi.set(__self__, "exempted_members", exempted_members)
        pulumi.set(__self__, "log_type", log_type)

    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Sequence[str]:
        """
        Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
        """
        return pulumi.get(self, "exempted_members")

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> str:
        """
        The log type that this config enables.
        """
        return pulumi.get(self, "log_type")


@pulumi.output_type
class AwsResponse(dict):
    """
    Represents an Amazon Web Services identity provider.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountId":
            suggest = "account_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AwsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AwsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AwsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_id: str):
        """
        Represents an Amazon Web Services identity provider.
        :param str account_id: The AWS account ID.
        """
        pulumi.set(__self__, "account_id", account_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The AWS account ID.
        """
        return pulumi.get(self, "account_id")


@pulumi.output_type
class BindingResponse(dict):
    """
    Associates `members` with a `role`.
    """
    def __init__(__self__, *,
                 condition: 'outputs.ExprResponse',
                 members: Sequence[str],
                 role: str):
        """
        Associates `members` with a `role`.
        :param 'ExprResponse' condition: The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the members in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        :param Sequence[str] members: Specifies the identities requesting access for a Cloud Platform resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        :param str role: Role that is assigned to `members`. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        pulumi.set(__self__, "condition", condition)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> 'outputs.ExprResponse':
        """
        The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the members in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        """
        Specifies the identities requesting access for a Cloud Platform resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Role that is assigned to `members`. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        return pulumi.get(self, "role")


@pulumi.output_type
class ExprResponse(dict):
    """
    Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
    """
    def __init__(__self__, *,
                 description: str,
                 expression: str,
                 location: str,
                 title: str):
        """
        Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
        :param str description: Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        :param str expression: Textual representation of an expression in Common Expression Language syntax.
        :param str location: Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param str title: Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")


@pulumi.output_type
class OidcResponse(dict):
    """
    Represents an OpenId Connect 1.0 identity provider.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedAudiences":
            suggest = "allowed_audiences"
        elif key == "issuerUri":
            suggest = "issuer_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OidcResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OidcResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OidcResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed_audiences: Sequence[str],
                 issuer_uri: str):
        """
        Represents an OpenId Connect 1.0 identity provider.
        :param Sequence[str] allowed_audiences: Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange requests are rejected if the token audience does not match one of the configured values. Each audience may be at most 256 characters. A maximum of 10 audiences may be configured. If this list is empty, the OIDC token audience must be equal to the full canonical resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix. For example: ``` //iam.googleapis.com/projects//locations//workloadIdentityPools//providers/ https://iam.googleapis.com/projects//locations//workloadIdentityPools//providers/ ```
        :param str issuer_uri: The OIDC issuer URL.
        """
        pulumi.set(__self__, "allowed_audiences", allowed_audiences)
        pulumi.set(__self__, "issuer_uri", issuer_uri)

    @property
    @pulumi.getter(name="allowedAudiences")
    def allowed_audiences(self) -> Sequence[str]:
        """
        Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange requests are rejected if the token audience does not match one of the configured values. Each audience may be at most 256 characters. A maximum of 10 audiences may be configured. If this list is empty, the OIDC token audience must be equal to the full canonical resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix. For example: ``` //iam.googleapis.com/projects//locations//workloadIdentityPools//providers/ https://iam.googleapis.com/projects//locations//workloadIdentityPools//providers/ ```
        """
        return pulumi.get(self, "allowed_audiences")

    @property
    @pulumi.getter(name="issuerUri")
    def issuer_uri(self) -> str:
        """
        The OIDC issuer URL.
        """
        return pulumi.get(self, "issuer_uri")


