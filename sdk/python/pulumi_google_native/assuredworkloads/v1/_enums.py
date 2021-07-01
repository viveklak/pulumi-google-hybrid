# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResourceType',
    'WorkloadComplianceRegime',
]


class GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResourceType(str, Enum):
    """
    Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT)
    """
    RESOURCE_TYPE_UNSPECIFIED = "RESOURCE_TYPE_UNSPECIFIED"
    """Unknown resource type."""
    CONSUMER_PROJECT = "CONSUMER_PROJECT"
    """Consumer project."""
    ENCRYPTION_KEYS_PROJECT = "ENCRYPTION_KEYS_PROJECT"
    """Consumer project containing encryption keys."""
    KEYRING = "KEYRING"
    """Keyring resource that hosts encryption keys."""


class WorkloadComplianceRegime(str, Enum):
    """
    Required. Immutable. Compliance Regime associated with this workload.
    """
    COMPLIANCE_REGIME_UNSPECIFIED = "COMPLIANCE_REGIME_UNSPECIFIED"
    """Unknown compliance regime."""
    IL4 = "IL4"
    """Information protection as per DoD IL4 requirements."""
    CJIS = "CJIS"
    """Criminal Justice Information Services (CJIS) Security policies."""
    FEDRAMP_HIGH = "FEDRAMP_HIGH"
    """FedRAMP High data protection controls"""
    FEDRAMP_MODERATE = "FEDRAMP_MODERATE"
    """FedRAMP Moderate data protection controls"""
    US_REGIONAL_ACCESS = "US_REGIONAL_ACCESS"
    """Assured Workloads For US Regions data protection controls"""
    HIPAA = "HIPAA"
    """Health Insurance Portability and Accountability Act controls"""
    HITRUST = "HITRUST"
    """Health Information Trust Alliance controls"""
