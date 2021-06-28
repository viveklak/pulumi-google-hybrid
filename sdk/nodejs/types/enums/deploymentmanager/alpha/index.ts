// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AuditLogConfigLogType = {
    /**
     * Default case. Should never be this.
     */
    LogTypeUnspecified: "LOG_TYPE_UNSPECIFIED",
    /**
     * Admin reads. Example: CloudIAM getIamPolicy
     */
    AdminRead: "ADMIN_READ",
    /**
     * Data writes. Example: CloudSQL Users create
     */
    DataWrite: "DATA_WRITE",
    /**
     * Data reads. Example: CloudSQL Users list
     */
    DataRead: "DATA_READ",
} as const;

/**
 * The log type that this config enables.
 */
export type AuditLogConfigLogType = (typeof AuditLogConfigLogType)[keyof typeof AuditLogConfigLogType];

export const CompositeTypeStatus = {
    UnknownStatus: "UNKNOWN_STATUS",
    Deprecated: "DEPRECATED",
    Experimental: "EXPERIMENTAL",
    Supported: "SUPPORTED",
} as const;

export type CompositeTypeStatus = (typeof CompositeTypeStatus)[keyof typeof CompositeTypeStatus];

export const DiagnosticLevel = {
    Unknown: "UNKNOWN",
    /**
     * If level is informational, it only gets displayed as part of the resource.
     */
    Information: "INFORMATION",
    /**
     * If level is warning, will end up in the resource as a warning.
     */
    Warning: "WARNING",
    /**
     * If level is error, it will indicate an error occurred after finishCondition is set, and this field will populate resource errors and operation errors.
     */
    Error: "ERROR",
} as const;

/**
 * Level to record this diagnostic.
 */
export type DiagnosticLevel = (typeof DiagnosticLevel)[keyof typeof DiagnosticLevel];

export const InputMappingLocation = {
    Unknown: "UNKNOWN",
    Path: "PATH",
    Query: "QUERY",
    Body: "BODY",
    Header: "HEADER",
} as const;

/**
 * The location where this mapping applies.
 */
export type InputMappingLocation = (typeof InputMappingLocation)[keyof typeof InputMappingLocation];

export const TemplateContentsInterpreter = {
    UnknownInterpreter: "UNKNOWN_INTERPRETER",
    Python: "PYTHON",
    Jinja: "JINJA",
} as const;

/**
 * Which interpreter (python or jinja) should be used during expansion.
 */
export type TemplateContentsInterpreter = (typeof TemplateContentsInterpreter)[keyof typeof TemplateContentsInterpreter];

export const ValidationOptionsSchemaValidation = {
    Unknown: "UNKNOWN",
    /**
     * Ignore schema failures.
     */
    Ignore: "IGNORE",
    /**
     * Ignore schema failures but display them as warnings.
     */
    IgnoreWithWarnings: "IGNORE_WITH_WARNINGS",
    /**
     * Fail the resource if the schema is not valid, this is the default behavior.
     */
    Fail: "FAIL",
} as const;

/**
 * Customize how deployment manager will validate the resource against schema errors.
 */
export type ValidationOptionsSchemaValidation = (typeof ValidationOptionsSchemaValidation)[keyof typeof ValidationOptionsSchemaValidation];

export const ValidationOptionsUndeclaredProperties = {
    Unknown: "UNKNOWN",
    /**
     * Always include even if not present on discovery doc.
     */
    Include: "INCLUDE",
    /**
     * Always ignore if not present on discovery doc.
     */
    Ignore: "IGNORE",
    /**
     * Include on request, but emit a warning.
     */
    IncludeWithWarnings: "INCLUDE_WITH_WARNINGS",
    /**
     * Ignore properties, but emit a warning.
     */
    IgnoreWithWarnings: "IGNORE_WITH_WARNINGS",
    /**
     * Always fail if undeclared properties are present.
     */
    Fail: "FAIL",
} as const;

/**
 * Specify what to do with extra properties when executing a request.
 */
export type ValidationOptionsUndeclaredProperties = (typeof ValidationOptionsUndeclaredProperties)[keyof typeof ValidationOptionsUndeclaredProperties];
