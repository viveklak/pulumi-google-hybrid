# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'AuditConfigArgs',
    'AuditLogConfigArgs',
    'BindingArgs',
    'EventTriggerArgs',
    'ExprArgs',
    'FailurePolicyArgs',
    'HttpsTriggerArgs',
    'RetryArgs',
    'SourceRepositoryArgs',
]

@pulumi.input_type
class AuditConfigArgs:
    def __init__(__self__, *,
                 audit_log_configs: Optional[pulumi.Input[Sequence[pulumi.Input['AuditLogConfigArgs']]]] = None,
                 service: Optional[pulumi.Input[str]] = None):
        """
        Specifies the audit configuration for a service. The configuration determines which permission types are logged, and what identities, if any, are exempted from logging. An AuditConfig must have one or more AuditLogConfigs. If there are AuditConfigs for both `allServices` and a specific service, the union of the two AuditConfigs is used for that service: the log_types specified in each AuditConfig are enabled, and the exempted_members in each AuditLogConfig are exempted. Example Policy with multiple AuditConfigs: { "audit_configs": [ { "service": "allServices", "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" }, { "log_type": "ADMIN_READ" } ] }, { "service": "sampleservice.googleapis.com", "audit_log_configs": [ { "log_type": "DATA_READ" }, { "log_type": "DATA_WRITE", "exempted_members": [ "user:aliya@example.com" ] } ] } ] } For sampleservice, this policy enables DATA_READ, DATA_WRITE and ADMIN_READ logging. It also exempts jose@example.com from DATA_READ logging, and aliya@example.com from DATA_WRITE logging.
        :param pulumi.Input[Sequence[pulumi.Input['AuditLogConfigArgs']]] audit_log_configs: The configuration for logging of each type of permission.
        :param pulumi.Input[str] service: Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        if audit_log_configs is not None:
            pulumi.set(__self__, "audit_log_configs", audit_log_configs)
        if service is not None:
            pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter(name="auditLogConfigs")
    def audit_log_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AuditLogConfigArgs']]]]:
        """
        The configuration for logging of each type of permission.
        """
        return pulumi.get(self, "audit_log_configs")

    @audit_log_configs.setter
    def audit_log_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AuditLogConfigArgs']]]]):
        pulumi.set(self, "audit_log_configs", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)


@pulumi.input_type
class AuditLogConfigArgs:
    def __init__(__self__, *,
                 exempted_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_type: Optional[pulumi.Input[str]] = None):
        """
        Provides the configuration for logging a type of permissions. Example: { "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" } ] } This enables 'DATA_READ' and 'DATA_WRITE' logging, while exempting jose@example.com from DATA_READ logging.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exempted_members: Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
        :param pulumi.Input[str] log_type: The log type that this config enables.
        """
        if exempted_members is not None:
            pulumi.set(__self__, "exempted_members", exempted_members)
        if log_type is not None:
            pulumi.set(__self__, "log_type", log_type)

    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
        """
        return pulumi.get(self, "exempted_members")

    @exempted_members.setter
    def exempted_members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exempted_members", value)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> Optional[pulumi.Input[str]]:
        """
        The log type that this config enables.
        """
        return pulumi.get(self, "log_type")

    @log_type.setter
    def log_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_type", value)


@pulumi.input_type
class BindingArgs:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['ExprArgs']] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Associates `members` with a `role`.
        :param pulumi.Input['ExprArgs'] condition: The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the members in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Specifies the identities requesting access for a Cloud Platform resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        :param pulumi.Input[str] role: Role that is assigned to `members`. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['ExprArgs']]:
        """
        The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the members in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['ExprArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the identities requesting access for a Cloud Platform resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Role that is assigned to `members`. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class EventTriggerArgs:
    def __init__(__self__, *,
                 event_type: Optional[pulumi.Input[str]] = None,
                 failure_policy: Optional[pulumi.Input['FailurePolicyArgs']] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None):
        """
        Describes EventTrigger, used to request events be sent from another service.
        :param pulumi.Input[str] event_type: Required. The type of event to observe. For example: `providers/cloud.storage/eventTypes/object.change` and `providers/cloud.pubsub/eventTypes/topic.publish`. Event types match pattern `providers/*/eventTypes/*.*`. The pattern contains: 1. namespace: For example, `cloud.storage` and `google.firebase.analytics`. 2. resource type: The type of resource on which event occurs. For example, the Google Cloud Storage API includes the type `object`. 3. action: The action that generates the event. For example, action for a Google Cloud Storage Object is 'change'. These parts are lower case.
        :param pulumi.Input['FailurePolicyArgs'] failure_policy: Specifies policy for failed executions.
        :param pulumi.Input[str] resource: Required. The resource(s) from which to observe events, for example, `projects/_/buckets/myBucket`. Not all syntactically correct values are accepted by all services. For example: 1. The authorization model must support it. Google Cloud Functions only allows EventTriggers to be deployed that observe resources in the same project as the `CloudFunction`. 2. The resource type must match the pattern expected for an `event_type`. For example, an `EventTrigger` that has an `event_type` of "google.pubsub.topic.publish" should have a resource that matches Google Cloud Pub/Sub topics. Additionally, some services may support short names when creating an `EventTrigger`. These will always be returned in the normalized "long" format. See each *service's* documentation for supported formats.
        :param pulumi.Input[str] service: The hostname of the service that should be observed. If no string is provided, the default service implementing the API will be used. For example, `storage.googleapis.com` is the default for all event types in the `google.storage` namespace.
        """
        if event_type is not None:
            pulumi.set(__self__, "event_type", event_type)
        if failure_policy is not None:
            pulumi.set(__self__, "failure_policy", failure_policy)
        if resource is not None:
            pulumi.set(__self__, "resource", resource)
        if service is not None:
            pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter(name="eventType")
    def event_type(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The type of event to observe. For example: `providers/cloud.storage/eventTypes/object.change` and `providers/cloud.pubsub/eventTypes/topic.publish`. Event types match pattern `providers/*/eventTypes/*.*`. The pattern contains: 1. namespace: For example, `cloud.storage` and `google.firebase.analytics`. 2. resource type: The type of resource on which event occurs. For example, the Google Cloud Storage API includes the type `object`. 3. action: The action that generates the event. For example, action for a Google Cloud Storage Object is 'change'. These parts are lower case.
        """
        return pulumi.get(self, "event_type")

    @event_type.setter
    def event_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_type", value)

    @property
    @pulumi.getter(name="failurePolicy")
    def failure_policy(self) -> Optional[pulumi.Input['FailurePolicyArgs']]:
        """
        Specifies policy for failed executions.
        """
        return pulumi.get(self, "failure_policy")

    @failure_policy.setter
    def failure_policy(self, value: Optional[pulumi.Input['FailurePolicyArgs']]):
        pulumi.set(self, "failure_policy", value)

    @property
    @pulumi.getter
    def resource(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The resource(s) from which to observe events, for example, `projects/_/buckets/myBucket`. Not all syntactically correct values are accepted by all services. For example: 1. The authorization model must support it. Google Cloud Functions only allows EventTriggers to be deployed that observe resources in the same project as the `CloudFunction`. 2. The resource type must match the pattern expected for an `event_type`. For example, an `EventTrigger` that has an `event_type` of "google.pubsub.topic.publish" should have a resource that matches Google Cloud Pub/Sub topics. Additionally, some services may support short names when creating an `EventTrigger`. These will always be returned in the normalized "long" format. See each *service's* documentation for supported formats.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        The hostname of the service that should be observed. If no string is provided, the default service implementing the API will be used. For example, `storage.googleapis.com` is the default for all event types in the `google.storage` namespace.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)


@pulumi.input_type
class ExprArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
        :param pulumi.Input[str] description: Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] location: Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param pulumi.Input[str] title: Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class FailurePolicyArgs:
    def __init__(__self__, *,
                 retry: Optional[pulumi.Input['RetryArgs']] = None):
        """
        Describes the policy in case of function's execution failure. If empty, then defaults to ignoring failures (i.e. not retrying them).
        :param pulumi.Input['RetryArgs'] retry: If specified, then the function will be retried in case of a failure.
        """
        if retry is not None:
            pulumi.set(__self__, "retry", retry)

    @property
    @pulumi.getter
    def retry(self) -> Optional[pulumi.Input['RetryArgs']]:
        """
        If specified, then the function will be retried in case of a failure.
        """
        return pulumi.get(self, "retry")

    @retry.setter
    def retry(self, value: Optional[pulumi.Input['RetryArgs']]):
        pulumi.set(self, "retry", value)


@pulumi.input_type
class HttpsTriggerArgs:
    def __init__(__self__, *,
                 security_level: Optional[pulumi.Input[str]] = None):
        """
        Describes HttpsTrigger, could be used to connect web hooks to function.
        :param pulumi.Input[str] security_level: The security level for the function.
        """
        if security_level is not None:
            pulumi.set(__self__, "security_level", security_level)

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> Optional[pulumi.Input[str]]:
        """
        The security level for the function.
        """
        return pulumi.get(self, "security_level")

    @security_level.setter
    def security_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_level", value)


@pulumi.input_type
class RetryArgs:
    def __init__(__self__):
        """
        Describes the retry policy in case of function's execution failure. A function execution will be retried on any failure. A failed execution will be retried up to 7 days with an exponential backoff (capped at 10 seconds). Retried execution is charged as any other execution.
        """
        pass


@pulumi.input_type
class SourceRepositoryArgs:
    def __init__(__self__, *,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Describes SourceRepository, used to represent parameters related to source repository where a function is hosted.
        :param pulumi.Input[str] url: The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats: To refer to a specific commit: `https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*` To refer to a moveable alias (branch): `https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*` In particular, to refer to HEAD use `master` moveable alias. To refer to a specific fixed alias (tag): `https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*` You may omit `paths/*` if you want to use the main directory.
        """
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats: To refer to a specific commit: `https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*` To refer to a moveable alias (branch): `https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*` In particular, to refer to HEAD use `master` moveable alias. To refer to a specific fixed alias (tag): `https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*` You may omit `paths/*` if you want to use the main directory.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


