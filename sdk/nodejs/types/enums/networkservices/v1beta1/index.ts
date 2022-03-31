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

export const EndpointPolicyType = {
    /**
     * Default value. Must not be used.
     */
    EndpointPolicyTypeUnspecified: "ENDPOINT_POLICY_TYPE_UNSPECIFIED",
    /**
     * Represents a proxy deployed as a sidecar.
     */
    SidecarProxy: "SIDECAR_PROXY",
    /**
     * Represents a proxyless gRPC backend.
     */
    GrpcServer: "GRPC_SERVER",
} as const;

/**
 * Required. The type of endpoint policy. This is primarily used to validate the configuration.
 */
export type EndpointPolicyType = (typeof EndpointPolicyType)[keyof typeof EndpointPolicyType];

export const GatewayType = {
    /**
     * The type of the customer managed gateway is unspecified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * The type of the customer managed gateway is TrafficDirector Open Mesh.
     */
    OpenMesh: "OPEN_MESH",
    /**
     * The type of the customer managed gateway is SecureWebGateway (SWG).
     */
    SecureWebGateway: "SECURE_WEB_GATEWAY",
} as const;

/**
 * Immutable. The type of the customer managed gateway.
 */
export type GatewayType = (typeof GatewayType)[keyof typeof GatewayType];

export const GrpcRouteHeaderMatchType = {
    /**
     * Unspecified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Will only match the exact value provided.
     */
    Exact: "EXACT",
    /**
     * Will match paths conforming to the prefix specified by value. RE2 syntax is supported.
     */
    RegularExpression: "REGULAR_EXPRESSION",
} as const;

/**
 * Optional. Specifies how to match against the value of the header. If not specified, a default value of EXACT is used.
 */
export type GrpcRouteHeaderMatchType = (typeof GrpcRouteHeaderMatchType)[keyof typeof GrpcRouteHeaderMatchType];

export const GrpcRouteMethodMatchType = {
    /**
     * Unspecified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Will only match the exact name provided.
     */
    Exact: "EXACT",
    /**
     * Will interpret grpc_method and grpc_service as regexes. RE2 syntax is supported.
     */
    RegularExpression: "REGULAR_EXPRESSION",
} as const;

/**
 * Optional. Specifies how to match against the name. If not specified, a default value of "EXACT" is used.
 */
export type GrpcRouteMethodMatchType = (typeof GrpcRouteMethodMatchType)[keyof typeof GrpcRouteMethodMatchType];

export const HttpRouteRedirectResponseCode = {
    /**
     * Default value
     */
    ResponseCodeUnspecified: "RESPONSE_CODE_UNSPECIFIED",
    /**
     * Corresponds to 301.
     */
    MovedPermanentlyDefault: "MOVED_PERMANENTLY_DEFAULT",
    /**
     * Corresponds to 302.
     */
    Found: "FOUND",
    /**
     * Corresponds to 303.
     */
    SeeOther: "SEE_OTHER",
    /**
     * Corresponds to 307. In this case, the request method will be retained.
     */
    TemporaryRedirect: "TEMPORARY_REDIRECT",
    /**
     * Corresponds to 308. In this case, the request method will be retained.
     */
    PermanentRedirect: "PERMANENT_REDIRECT",
} as const;

/**
 * The HTTP Status code to use for the redirect.
 */
export type HttpRouteRedirectResponseCode = (typeof HttpRouteRedirectResponseCode)[keyof typeof HttpRouteRedirectResponseCode];

export const MetadataLabelMatcherMetadataLabelMatchCriteria = {
    /**
     * Default value. Should not be used.
     */
    MetadataLabelMatchCriteriaUnspecified: "METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED",
    /**
     * At least one of the Labels specified in the matcher should match the metadata presented by xDS client.
     */
    MatchAny: "MATCH_ANY",
    /**
     * The metadata presented by the xDS client should contain all of the labels specified here.
     */
    MatchAll: "MATCH_ALL",
} as const;

/**
 * Specifies how matching should be done. Supported values are: MATCH_ANY: At least one of the Labels specified in the matcher should match the metadata presented by xDS client. MATCH_ALL: The metadata presented by the xDS client should contain all of the labels specified here. The selection is determined based on the best match. For example, suppose there are three EndpointPolicy resources P1, P2 and P3 and if P1 has a the matcher as MATCH_ANY , P2 has MATCH_ALL , and P3 has MATCH_ALL . If a client with label connects, the config from P1 will be selected. If a client with label connects, the config from P2 will be selected. If a client with label connects, the config from P3 will be selected. If there is more than one best match, (for example, if a config P4 with selector exists and if a client with label connects), an error will be thrown.
 */
export type MetadataLabelMatcherMetadataLabelMatchCriteria = (typeof MetadataLabelMatcherMetadataLabelMatchCriteria)[keyof typeof MetadataLabelMatcherMetadataLabelMatchCriteria];
