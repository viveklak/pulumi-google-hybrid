// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the specified domain mapping.
 */
export function getDomainMapping(args: GetDomainMappingArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainMappingResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:appengine/v1:getDomainMapping", {
        "appId": args.appId,
        "domainMappingId": args.domainMappingId,
    }, opts);
}

export interface GetDomainMappingArgs {
    appId: string;
    domainMappingId: string;
}

export interface GetDomainMappingResult {
    /**
     * Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
     */
    readonly name: string;
    /**
     * The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.
     */
    readonly resourceRecords: outputs.appengine.v1.ResourceRecordResponse[];
    /**
     * SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
     */
    readonly sslSettings: outputs.appengine.v1.SslSettingsResponse;
}

export function getDomainMappingOutput(args: GetDomainMappingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainMappingResult> {
    return pulumi.output(args).apply(a => getDomainMapping(a, opts))
}

export interface GetDomainMappingOutputArgs {
    appId: pulumi.Input<string>;
    domainMappingId: pulumi.Input<string>;
}
