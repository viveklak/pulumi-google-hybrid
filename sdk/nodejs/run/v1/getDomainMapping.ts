// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Get information about a domain mapping.
 */
export function getDomainMapping(args: GetDomainMappingArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainMappingResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:run/v1:getDomainMapping", {
        "domainmappingId": args.domainmappingId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetDomainMappingArgs {
    domainmappingId: string;
    location: string;
    project?: string;
}

export interface GetDomainMappingResult {
    /**
     * The API version for this call such as "domains.cloudrun.com/v1".
     */
    readonly apiVersion: string;
    /**
     * The kind of resource, in this case "DomainMapping".
     */
    readonly kind: string;
    /**
     * Metadata associated with this BuildTemplate.
     */
    readonly metadata: outputs.run.v1.ObjectMetaResponse;
    /**
     * The spec for this DomainMapping.
     */
    readonly spec: outputs.run.v1.DomainMappingSpecResponse;
    /**
     * The current status of the DomainMapping.
     */
    readonly status: outputs.run.v1.DomainMappingStatusResponse;
}

export function getDomainMappingOutput(args: GetDomainMappingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainMappingResult> {
    return pulumi.output(args).apply(a => getDomainMapping(a, opts))
}

export interface GetDomainMappingOutputArgs {
    domainmappingId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
