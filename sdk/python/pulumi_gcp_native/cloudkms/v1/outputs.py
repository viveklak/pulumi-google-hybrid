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
    'AuditConfigResponse',
    'AuditLogConfigResponse',
    'BindingResponse',
    'CertificateChainsResponse',
    'CryptoKeyVersionResponse',
    'CryptoKeyVersionTemplateResponse',
    'ExprResponse',
    'ExternalProtectionLevelOptionsResponse',
    'KeyOperationAttestationResponse',
    'WrappingPublicKeyResponse',
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
        :param Sequence['AuditLogConfigResponseArgs'] audit_log_configs: The configuration for logging of each type of permission.
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
        :param 'ExprResponseArgs' condition: The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the members in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
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
class CertificateChainsResponse(dict):
    """
    Certificate chains needed to verify the attestation. Certificates in chains are PEM-encoded and are ordered based on https://tools.ietf.org/html/rfc5246#section-7.4.2.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "caviumCerts":
            suggest = "cavium_certs"
        elif key == "googleCardCerts":
            suggest = "google_card_certs"
        elif key == "googlePartitionCerts":
            suggest = "google_partition_certs"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CertificateChainsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CertificateChainsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CertificateChainsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cavium_certs: Sequence[str],
                 google_card_certs: Sequence[str],
                 google_partition_certs: Sequence[str]):
        """
        Certificate chains needed to verify the attestation. Certificates in chains are PEM-encoded and are ordered based on https://tools.ietf.org/html/rfc5246#section-7.4.2.
        :param Sequence[str] cavium_certs: Cavium certificate chain corresponding to the attestation.
        :param Sequence[str] google_card_certs: Google card certificate chain corresponding to the attestation.
        :param Sequence[str] google_partition_certs: Google partition certificate chain corresponding to the attestation.
        """
        pulumi.set(__self__, "cavium_certs", cavium_certs)
        pulumi.set(__self__, "google_card_certs", google_card_certs)
        pulumi.set(__self__, "google_partition_certs", google_partition_certs)

    @property
    @pulumi.getter(name="caviumCerts")
    def cavium_certs(self) -> Sequence[str]:
        """
        Cavium certificate chain corresponding to the attestation.
        """
        return pulumi.get(self, "cavium_certs")

    @property
    @pulumi.getter(name="googleCardCerts")
    def google_card_certs(self) -> Sequence[str]:
        """
        Google card certificate chain corresponding to the attestation.
        """
        return pulumi.get(self, "google_card_certs")

    @property
    @pulumi.getter(name="googlePartitionCerts")
    def google_partition_certs(self) -> Sequence[str]:
        """
        Google partition certificate chain corresponding to the attestation.
        """
        return pulumi.get(self, "google_partition_certs")


@pulumi.output_type
class CryptoKeyVersionResponse(dict):
    """
    A CryptoKeyVersion represents an individual cryptographic key, and the associated key material. An ENABLED version can be used for cryptographic operations. For security reasons, the raw cryptographic key material represented by a CryptoKeyVersion can never be viewed or exported. It can only be used to encrypt, decrypt, or sign data when an authorized user or application invokes Cloud KMS.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createTime":
            suggest = "create_time"
        elif key == "destroyEventTime":
            suggest = "destroy_event_time"
        elif key == "destroyTime":
            suggest = "destroy_time"
        elif key == "externalProtectionLevelOptions":
            suggest = "external_protection_level_options"
        elif key == "generateTime":
            suggest = "generate_time"
        elif key == "importFailureReason":
            suggest = "import_failure_reason"
        elif key == "importJob":
            suggest = "import_job"
        elif key == "importTime":
            suggest = "import_time"
        elif key == "protectionLevel":
            suggest = "protection_level"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CryptoKeyVersionResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CryptoKeyVersionResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CryptoKeyVersionResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 algorithm: str,
                 attestation: 'outputs.KeyOperationAttestationResponse',
                 create_time: str,
                 destroy_event_time: str,
                 destroy_time: str,
                 external_protection_level_options: 'outputs.ExternalProtectionLevelOptionsResponse',
                 generate_time: str,
                 import_failure_reason: str,
                 import_job: str,
                 import_time: str,
                 name: str,
                 protection_level: str,
                 state: str):
        """
        A CryptoKeyVersion represents an individual cryptographic key, and the associated key material. An ENABLED version can be used for cryptographic operations. For security reasons, the raw cryptographic key material represented by a CryptoKeyVersion can never be viewed or exported. It can only be used to encrypt, decrypt, or sign data when an authorized user or application invokes Cloud KMS.
        :param str algorithm: The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        :param 'KeyOperationAttestationResponseArgs' attestation: Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.
        :param str create_time: The time at which this CryptoKeyVersion was created.
        :param str destroy_event_time: The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.
        :param str destroy_time: The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.
        :param 'ExternalProtectionLevelOptionsResponseArgs' external_protection_level_options: ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
        :param str generate_time: The time this CryptoKeyVersion's key material was generated.
        :param str import_failure_reason: The root cause of an import failure. Only present if state is IMPORT_FAILED.
        :param str import_job: The name of the ImportJob used to import this CryptoKeyVersion. Only present if the underlying key material was imported.
        :param str import_time: The time at which this CryptoKeyVersion's key material was imported.
        :param str name: The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
        :param str protection_level: The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        :param str state: The current state of the CryptoKeyVersion.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "attestation", attestation)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "destroy_event_time", destroy_event_time)
        pulumi.set(__self__, "destroy_time", destroy_time)
        pulumi.set(__self__, "external_protection_level_options", external_protection_level_options)
        pulumi.set(__self__, "generate_time", generate_time)
        pulumi.set(__self__, "import_failure_reason", import_failure_reason)
        pulumi.set(__self__, "import_job", import_job)
        pulumi.set(__self__, "import_time", import_time)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "protection_level", protection_level)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def attestation(self) -> 'outputs.KeyOperationAttestationResponse':
        """
        Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.
        """
        return pulumi.get(self, "attestation")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time at which this CryptoKeyVersion was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="destroyEventTime")
    def destroy_event_time(self) -> str:
        """
        The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.
        """
        return pulumi.get(self, "destroy_event_time")

    @property
    @pulumi.getter(name="destroyTime")
    def destroy_time(self) -> str:
        """
        The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.
        """
        return pulumi.get(self, "destroy_time")

    @property
    @pulumi.getter(name="externalProtectionLevelOptions")
    def external_protection_level_options(self) -> 'outputs.ExternalProtectionLevelOptionsResponse':
        """
        ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
        """
        return pulumi.get(self, "external_protection_level_options")

    @property
    @pulumi.getter(name="generateTime")
    def generate_time(self) -> str:
        """
        The time this CryptoKeyVersion's key material was generated.
        """
        return pulumi.get(self, "generate_time")

    @property
    @pulumi.getter(name="importFailureReason")
    def import_failure_reason(self) -> str:
        """
        The root cause of an import failure. Only present if state is IMPORT_FAILED.
        """
        return pulumi.get(self, "import_failure_reason")

    @property
    @pulumi.getter(name="importJob")
    def import_job(self) -> str:
        """
        The name of the ImportJob used to import this CryptoKeyVersion. Only present if the underlying key material was imported.
        """
        return pulumi.get(self, "import_job")

    @property
    @pulumi.getter(name="importTime")
    def import_time(self) -> str:
        """
        The time at which this CryptoKeyVersion's key material was imported.
        """
        return pulumi.get(self, "import_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> str:
        """
        The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        """
        return pulumi.get(self, "protection_level")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the CryptoKeyVersion.
        """
        return pulumi.get(self, "state")


@pulumi.output_type
class CryptoKeyVersionTemplateResponse(dict):
    """
    A CryptoKeyVersionTemplate specifies the properties to use when creating a new CryptoKeyVersion, either manually with CreateCryptoKeyVersion or automatically as a result of auto-rotation.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "protectionLevel":
            suggest = "protection_level"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CryptoKeyVersionTemplateResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CryptoKeyVersionTemplateResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CryptoKeyVersionTemplateResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 algorithm: str,
                 protection_level: str):
        """
        A CryptoKeyVersionTemplate specifies the properties to use when creating a new CryptoKeyVersion, either manually with CreateCryptoKeyVersion or automatically as a result of auto-rotation.
        :param str algorithm: Required. Algorithm to use when creating a CryptoKeyVersion based on this template. For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION is implied if both this field is omitted and CryptoKey.purpose is ENCRYPT_DECRYPT.
        :param str protection_level: ProtectionLevel to use when creating a CryptoKeyVersion based on this template. Immutable. Defaults to SOFTWARE.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "protection_level", protection_level)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        Required. Algorithm to use when creating a CryptoKeyVersion based on this template. For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION is implied if both this field is omitted and CryptoKey.purpose is ENCRYPT_DECRYPT.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> str:
        """
        ProtectionLevel to use when creating a CryptoKeyVersion based on this template. Immutable. Defaults to SOFTWARE.
        """
        return pulumi.get(self, "protection_level")


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
class ExternalProtectionLevelOptionsResponse(dict):
    """
    ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "externalKeyUri":
            suggest = "external_key_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExternalProtectionLevelOptionsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExternalProtectionLevelOptionsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExternalProtectionLevelOptionsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 external_key_uri: str):
        """
        ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
        :param str external_key_uri: The URI for an external resource that this CryptoKeyVersion represents.
        """
        pulumi.set(__self__, "external_key_uri", external_key_uri)

    @property
    @pulumi.getter(name="externalKeyUri")
    def external_key_uri(self) -> str:
        """
        The URI for an external resource that this CryptoKeyVersion represents.
        """
        return pulumi.get(self, "external_key_uri")


@pulumi.output_type
class KeyOperationAttestationResponse(dict):
    """
    Contains an HSM-generated attestation about a key operation. For more information, see [Verifying attestations] (https://cloud.google.com/kms/docs/attest-key).
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certChains":
            suggest = "cert_chains"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KeyOperationAttestationResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KeyOperationAttestationResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KeyOperationAttestationResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cert_chains: 'outputs.CertificateChainsResponse',
                 content: str,
                 format: str):
        """
        Contains an HSM-generated attestation about a key operation. For more information, see [Verifying attestations] (https://cloud.google.com/kms/docs/attest-key).
        :param 'CertificateChainsResponseArgs' cert_chains: The certificate chains needed to validate the attestation
        :param str content: The attestation data provided by the HSM when the key operation was performed.
        :param str format: The format of the attestation data.
        """
        pulumi.set(__self__, "cert_chains", cert_chains)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter(name="certChains")
    def cert_chains(self) -> 'outputs.CertificateChainsResponse':
        """
        The certificate chains needed to validate the attestation
        """
        return pulumi.get(self, "cert_chains")

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The attestation data provided by the HSM when the key operation was performed.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def format(self) -> str:
        """
        The format of the attestation data.
        """
        return pulumi.get(self, "format")


@pulumi.output_type
class WrappingPublicKeyResponse(dict):
    """
    The public key component of the wrapping key. For details of the type of key this public key corresponds to, see the ImportMethod.
    """
    def __init__(__self__, *,
                 pem: str):
        """
        The public key component of the wrapping key. For details of the type of key this public key corresponds to, see the ImportMethod.
        :param str pem: The public key, encoded in PEM format. For more information, see the [RFC 7468](https://tools.ietf.org/html/rfc7468) sections for [General Considerations](https://tools.ietf.org/html/rfc7468#section-2) and [Textual Encoding of Subject Public Key Info] (https://tools.ietf.org/html/rfc7468#section-13).
        """
        pulumi.set(__self__, "pem", pem)

    @property
    @pulumi.getter
    def pem(self) -> str:
        """
        The public key, encoded in PEM format. For more information, see the [RFC 7468](https://tools.ietf.org/html/rfc7468) sections for [General Considerations](https://tools.ietf.org/html/rfc7468#section-2) and [Textual Encoding of Subject Public Key Info] (https://tools.ietf.org/html/rfc7468#section-13).
        """
        return pulumi.get(self, "pem")


