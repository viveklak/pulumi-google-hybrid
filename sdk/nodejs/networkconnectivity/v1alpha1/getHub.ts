// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Hub.
 */
export function getHub(args: GetHubArgs, opts?: pulumi.InvokeOptions): Promise<GetHubResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:networkconnectivity/v1alpha1:getHub", {
        "hubId": args.hubId,
        "project": args.project,
    }, opts);
}

export interface GetHubArgs {
    hubId: string;
    project?: string;
}

export interface GetHubResult {
    /**
     * Time when the Hub was created.
     */
    readonly createTime: string;
    /**
     * Short description of the hub resource.
     */
    readonly description: string;
    /**
     * User-defined labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * Immutable. The name of a Hub resource.
     */
    readonly name: string;
    /**
     * The current lifecycle state of this Hub.
     */
    readonly state: string;
    /**
     * Google-generated UUID for this resource. This is unique across all Hub resources. If a Hub resource is deleted and another with the same name is created, it gets a different unique_id.
     */
    readonly uniqueId: string;
    /**
     * Time when the Hub was updated.
     */
    readonly updateTime: string;
}

export function getHubOutput(args: GetHubOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHubResult> {
    return pulumi.output(args).apply(a => getHub(a, opts))
}

export interface GetHubOutputArgs {
    hubId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
