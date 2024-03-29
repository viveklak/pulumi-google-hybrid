// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CertificateMapEntryMatcher = {
    /**
     * A matcher has't been recognized.
     */
    MatcherUnspecified: "MATCHER_UNSPECIFIED",
    /**
     * A primary certificate that is served when SNI wasn't specified in the request or SNI couldn't be found in the map.
     */
    Primary: "PRIMARY",
} as const;

/**
 * A predefined matcher for particular cases, other than SNI selection.
 */
export type CertificateMapEntryMatcher = (typeof CertificateMapEntryMatcher)[keyof typeof CertificateMapEntryMatcher];

export const CertificateScope = {
    /**
     * Certificates with default scope are served from core Google data centers. If unsure, choose this option.
     */
    Default: "DEFAULT",
    /**
     * Certificates with scope EDGE_CACHE are special-purposed certificates, served from non-core Google data centers.
     */
    EdgeCache: "EDGE_CACHE",
} as const;

/**
 * Immutable. The scope of the certificate.
 */
export type CertificateScope = (typeof CertificateScope)[keyof typeof CertificateScope];

export const ProvisioningIssueReason = {
    ReasonUnspecified: "REASON_UNSPECIFIED",
    /**
     * Certificate provisioning failed due to an issue with one or more of the domains on the certificate. For details of which domains failed, consult the `authorization_attempt_info` field.
     */
    AuthorizationIssue: "AUTHORIZATION_ISSUE",
    /**
     * Exceeded Certificate Authority quotas or internal rate limits of the system. Provisioning may take longer to complete.
     */
    RateLimited: "RATE_LIMITED",
} as const;

/**
 * Reason for provisioning failures.
 */
export type ProvisioningIssueReason = (typeof ProvisioningIssueReason)[keyof typeof ProvisioningIssueReason];
