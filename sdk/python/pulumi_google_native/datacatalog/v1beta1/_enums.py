# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'EntryType',
    'TaxonomyActivatedPolicyTypesItem',
]


class EntryType(str, Enum):
    """
    The type of the entry. Only used for Entries with types in the EntryType enum.
    """
    ENTRY_TYPE_UNSPECIFIED = "ENTRY_TYPE_UNSPECIFIED"
    """Default unknown type."""
    TABLE = "TABLE"
    """Output only. The type of entry that has a GoogleSQL schema, including logical views."""
    MODEL = "MODEL"
    """Output only. The type of models. https://cloud.google.com/bigquery-ml/docs/bigqueryml-intro"""
    DATA_STREAM = "DATA_STREAM"
    """Output only. An entry type which is used for streaming entries. Example: Pub/Sub topic."""
    FILESET = "FILESET"
    """An entry type which is a set of files or objects. Example: Cloud Storage fileset."""


class TaxonomyActivatedPolicyTypesItem(str, Enum):
    POLICY_TYPE_UNSPECIFIED = "POLICY_TYPE_UNSPECIFIED"
    """Unspecified policy type."""
    FINE_GRAINED_ACCESS_CONTROL = "FINE_GRAINED_ACCESS_CONTROL"
    """Fine grained access control policy, which enables access control on tagged resources."""
