// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the specified TargetHttpProxy resource in the specified region. Gets a list of available target HTTP proxies by making a list() request.
 */
export function getRegionTargetHttpProxy(args: GetRegionTargetHttpProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionTargetHttpProxyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/beta:getRegionTargetHttpProxy", {
        "project": args.project,
        "region": args.region,
        "targetHttpProxy": args.targetHttpProxy,
    }, opts);
}

export interface GetRegionTargetHttpProxyArgs {
    project?: string;
    region: string;
    targetHttpProxy: string;
}

export interface GetRegionTargetHttpProxyResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
     */
    readonly fingerprint: string;
    /**
     * URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/ locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list. httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
     */
    readonly httpFilters: string[];
    /**
     * Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
     */
    readonly proxyBind: boolean;
    /**
     * URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * URL to the UrlMap resource that defines the mapping from URL to the BackendService.
     */
    readonly urlMap: string;
}

export function getRegionTargetHttpProxyOutput(args: GetRegionTargetHttpProxyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionTargetHttpProxyResult> {
    return pulumi.output(args).apply(a => getRegionTargetHttpProxy(a, opts))
}

export interface GetRegionTargetHttpProxyOutputArgs {
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    targetHttpProxy: pulumi.Input<string>;
}
