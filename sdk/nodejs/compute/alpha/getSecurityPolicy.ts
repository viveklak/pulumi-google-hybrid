// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * List all of the ordered rules present in a single specified policy.
 */
export function getSecurityPolicy(args: GetSecurityPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/alpha:getSecurityPolicy", {
        "project": args.project,
        "securityPolicy": args.securityPolicy,
    }, opts);
}

export interface GetSecurityPolicyArgs {
    project?: string;
    securityPolicy: string;
}

export interface GetSecurityPolicyResult {
    readonly adaptiveProtectionConfig: outputs.compute.alpha.SecurityPolicyAdaptiveProtectionConfigResponse;
    readonly advancedOptionsConfig: outputs.compute.alpha.SecurityPolicyAdvancedOptionsConfigResponse;
    /**
     * A list of associations that belong to this policy.
     */
    readonly associations: outputs.compute.alpha.SecurityPolicyAssociationResponse[];
    readonly cloudArmorConfig: outputs.compute.alpha.SecurityPolicyCloudArmorConfigResponse;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    readonly ddosProtectionConfig: outputs.compute.alpha.SecurityPolicyDdosProtectionConfigResponse;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly displayName: string;
    /**
     * Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
     */
    readonly fingerprint: string;
    /**
     * [Output only] Type of the resource. Always compute#securityPolicyfor security policies
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
     */
    readonly labelFingerprint: string;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The parent of the security policy.
     */
    readonly parent: string;
    readonly recaptchaOptionsConfig: outputs.compute.alpha.SecurityPolicyRecaptchaOptionsConfigResponse;
    /**
     * URL of the region where the regional security policy resides. This field is not applicable to global security policies.
     */
    readonly region: string;
    /**
     * Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
     */
    readonly ruleTupleCount: number;
    /**
     * A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
     */
    readonly rules: outputs.compute.alpha.SecurityPolicyRuleResponse[];
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
     */
    readonly type: string;
}

export function getSecurityPolicyOutput(args: GetSecurityPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityPolicyResult> {
    return pulumi.output(args).apply(a => getSecurityPolicy(a, opts))
}

export interface GetSecurityPolicyOutputArgs {
    project?: pulumi.Input<string>;
    securityPolicy: pulumi.Input<string>;
}
