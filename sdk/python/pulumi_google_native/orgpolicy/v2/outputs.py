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
    'GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse',
    'GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse',
    'GoogleCloudOrgpolicyV2PolicySpecResponse',
    'GoogleTypeExprResponse',
]

@pulumi.output_type
class GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse(dict):
    """
    A rule used to express this policy.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowAll":
            suggest = "allow_all"
        elif key == "denyAll":
            suggest = "deny_all"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_all: bool,
                 condition: 'outputs.GoogleTypeExprResponse',
                 deny_all: bool,
                 enforce: bool,
                 values: 'outputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse'):
        """
        A rule used to express this policy.
        :param bool allow_all: Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.
        :param 'GoogleTypeExprResponse' condition: A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
        :param bool deny_all: Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.
        :param bool enforce: If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
        :param 'GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse' values: List of values to be used for this PolicyRule. This field can be set only in Policies for list constraints.
        """
        pulumi.set(__self__, "allow_all", allow_all)
        pulumi.set(__self__, "condition", condition)
        pulumi.set(__self__, "deny_all", deny_all)
        pulumi.set(__self__, "enforce", enforce)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="allowAll")
    def allow_all(self) -> bool:
        """
        Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.
        """
        return pulumi.get(self, "allow_all")

    @property
    @pulumi.getter
    def condition(self) -> 'outputs.GoogleTypeExprResponse':
        """
        A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="denyAll")
    def deny_all(self) -> bool:
        """
        Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.
        """
        return pulumi.get(self, "deny_all")

    @property
    @pulumi.getter
    def enforce(self) -> bool:
        """
        If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
        """
        return pulumi.get(self, "enforce")

    @property
    @pulumi.getter
    def values(self) -> 'outputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse':
        """
        List of values to be used for this PolicyRule. This field can be set only in Policies for list constraints.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse(dict):
    """
    A message that holds specific allowed and denied values. This message can define specific values and subtrees of Cloud Resource Manager resource hierarchy (`Organizations`, `Folders`, `Projects`) that are allowed or denied. This is achieved by using the `under:` and optional `is:` prefixes. The `under:` prefix is used to denote resource subtree values. The `is:` prefix is used to denote specific values, and is required only if the value contains a ":". Values prefixed with "is:" are treated the same as values with no prefix. Ancestry subtrees must be in one of the following formats: - "projects/", e.g. "projects/tokyo-rain-123" - "folders/", e.g. "folders/1234" - "organizations/", e.g. "organizations/1234" The `supports_under` field of the associated `Constraint` defines whether ancestry prefixes can be used.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedValues":
            suggest = "allowed_values"
        elif key == "deniedValues":
            suggest = "denied_values"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed_values: Sequence[str],
                 denied_values: Sequence[str]):
        """
        A message that holds specific allowed and denied values. This message can define specific values and subtrees of Cloud Resource Manager resource hierarchy (`Organizations`, `Folders`, `Projects`) that are allowed or denied. This is achieved by using the `under:` and optional `is:` prefixes. The `under:` prefix is used to denote resource subtree values. The `is:` prefix is used to denote specific values, and is required only if the value contains a ":". Values prefixed with "is:" are treated the same as values with no prefix. Ancestry subtrees must be in one of the following formats: - "projects/", e.g. "projects/tokyo-rain-123" - "folders/", e.g. "folders/1234" - "organizations/", e.g. "organizations/1234" The `supports_under` field of the associated `Constraint` defines whether ancestry prefixes can be used.
        :param Sequence[str] allowed_values: List of values allowed at this resource.
        :param Sequence[str] denied_values: List of values denied at this resource.
        """
        pulumi.set(__self__, "allowed_values", allowed_values)
        pulumi.set(__self__, "denied_values", denied_values)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> Sequence[str]:
        """
        List of values allowed at this resource.
        """
        return pulumi.get(self, "allowed_values")

    @property
    @pulumi.getter(name="deniedValues")
    def denied_values(self) -> Sequence[str]:
        """
        List of values denied at this resource.
        """
        return pulumi.get(self, "denied_values")


@pulumi.output_type
class GoogleCloudOrgpolicyV2PolicySpecResponse(dict):
    """
    Defines a Cloud Organization `PolicySpec` which is used to specify `Constraints` for configurations of Cloud Platform resources.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "inheritFromParent":
            suggest = "inherit_from_parent"
        elif key == "updateTime":
            suggest = "update_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudOrgpolicyV2PolicySpecResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudOrgpolicyV2PolicySpecResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudOrgpolicyV2PolicySpecResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 etag: str,
                 inherit_from_parent: bool,
                 reset: bool,
                 rules: Sequence['outputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse'],
                 update_time: str):
        """
        Defines a Cloud Organization `PolicySpec` which is used to specify `Constraints` for configurations of Cloud Platform resources.
        :param str etag: An opaque tag indicating the current version of the `Policy`, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the `Policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current `Policy` to use when executing a read-modify-write loop. When the `Policy` is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
        :param bool inherit_from_parent: Determines the inheritance behavior for this `Policy`. If `inherit_from_parent` is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
        :param bool reset: Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific `Constraint` at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
        :param Sequence['GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse'] rules: Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set `enforced` to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
        :param str update_time: The time stamp this was previously updated. This represents the last time a call to `CreatePolicy` or `UpdatePolicy` was made for that `Policy`.
        """
        pulumi.set(__self__, "etag", etag)
        pulumi.set(__self__, "inherit_from_parent", inherit_from_parent)
        pulumi.set(__self__, "reset", reset)
        pulumi.set(__self__, "rules", rules)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        An opaque tag indicating the current version of the `Policy`, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the `Policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current `Policy` to use when executing a read-modify-write loop. When the `Policy` is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> bool:
        """
        Determines the inheritance behavior for this `Policy`. If `inherit_from_parent` is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
        """
        return pulumi.get(self, "inherit_from_parent")

    @property
    @pulumi.getter
    def reset(self) -> bool:
        """
        Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific `Constraint` at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
        """
        return pulumi.get(self, "reset")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleResponse']:
        """
        Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set `enforced` to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time stamp this was previously updated. This represents the last time a call to `CreatePolicy` or `UpdatePolicy` was made for that `Policy`.
        """
        return pulumi.get(self, "update_time")


@pulumi.output_type
class GoogleTypeExprResponse(dict):
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

