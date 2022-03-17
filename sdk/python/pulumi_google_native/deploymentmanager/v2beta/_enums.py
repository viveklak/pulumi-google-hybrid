# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'CompositeTypeStatus',
    'DiagnosticLevel',
    'InputMappingLocation',
    'TemplateContentsInterpreter',
    'ValidationOptionsSchemaValidation',
    'ValidationOptionsUndeclaredProperties',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class CompositeTypeStatus(str, Enum):
    UNKNOWN_STATUS = "UNKNOWN_STATUS"
    DEPRECATED = "DEPRECATED"
    EXPERIMENTAL = "EXPERIMENTAL"
    SUPPORTED = "SUPPORTED"


class DiagnosticLevel(str, Enum):
    """
    Level to record this diagnostic.
    """
    UNKNOWN = "UNKNOWN"
    INFORMATION = "INFORMATION"
    """
    If level is informational, it only gets displayed as part of the resource.
    """
    WARNING = "WARNING"
    """
    If level is warning, will end up in the resource as a warning.
    """
    ERROR = "ERROR"
    """
    If level is error, it will indicate an error occurred after finishCondition is set, and this field will populate resource errors and operation errors.
    """


class InputMappingLocation(str, Enum):
    """
    The location where this mapping applies.
    """
    UNKNOWN = "UNKNOWN"
    PATH = "PATH"
    QUERY = "QUERY"
    BODY = "BODY"
    HEADER = "HEADER"


class TemplateContentsInterpreter(str, Enum):
    """
    Which interpreter (python or jinja) should be used during expansion.
    """
    UNKNOWN_INTERPRETER = "UNKNOWN_INTERPRETER"
    PYTHON = "PYTHON"
    JINJA = "JINJA"


class ValidationOptionsSchemaValidation(str, Enum):
    """
    Customize how deployment manager will validate the resource against schema errors.
    """
    UNKNOWN = "UNKNOWN"
    IGNORE = "IGNORE"
    """
    Ignore schema failures.
    """
    IGNORE_WITH_WARNINGS = "IGNORE_WITH_WARNINGS"
    """
    Ignore schema failures but display them as warnings.
    """
    FAIL = "FAIL"
    """
    Fail the resource if the schema is not valid, this is the default behavior.
    """


class ValidationOptionsUndeclaredProperties(str, Enum):
    """
    Specify what to do with extra properties when executing a request.
    """
    UNKNOWN = "UNKNOWN"
    INCLUDE = "INCLUDE"
    """
    Always include even if not present on discovery doc.
    """
    IGNORE = "IGNORE"
    """
    Always ignore if not present on discovery doc.
    """
    INCLUDE_WITH_WARNINGS = "INCLUDE_WITH_WARNINGS"
    """
    Include on request, but emit a warning.
    """
    IGNORE_WITH_WARNINGS = "IGNORE_WITH_WARNINGS"
    """
    Ignore properties, but emit a warning.
    """
    FAIL = "FAIL"
    """
    Always fail if undeclared properties are present.
    """
