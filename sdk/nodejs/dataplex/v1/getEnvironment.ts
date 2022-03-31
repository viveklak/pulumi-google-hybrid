// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Get environment resource.
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dataplex/v1:getEnvironment", {
        "environmentId": args.environmentId,
        "lakeId": args.lakeId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetEnvironmentArgs {
    environmentId: string;
    lakeId: string;
    location: string;
    project?: string;
}

export interface GetEnvironmentResult {
    /**
     * Environment creation time.
     */
    readonly createTime: string;
    /**
     * Optional. Description of the environment.
     */
    readonly description: string;
    /**
     * Optional. User friendly display name.
     */
    readonly displayName: string;
    /**
     * URI Endpoints to access sessions associated with the Environment.
     */
    readonly endpoints: outputs.dataplex.v1.GoogleCloudDataplexV1EnvironmentEndpointsResponse;
    /**
     * Infrastructure specification for the Environment.
     */
    readonly infrastructureSpec: outputs.dataplex.v1.GoogleCloudDataplexV1EnvironmentInfrastructureSpecResponse;
    /**
     * Optional. User defined labels for the environment.
     */
    readonly labels: {[key: string]: string};
    /**
     * The relative resource name of the environment, of the form: projects/{project_id}/locations/{location_id}/lakes/{lake_id}/environment/{environment_id}
     */
    readonly name: string;
    /**
     * Optional. Configuration for sessions created for this environment.
     */
    readonly sessionSpec: outputs.dataplex.v1.GoogleCloudDataplexV1EnvironmentSessionSpecResponse;
    /**
     * Status of sessions created for this environment.
     */
    readonly sessionStatus: outputs.dataplex.v1.GoogleCloudDataplexV1EnvironmentSessionStatusResponse;
    /**
     * Current state of the environment.
     */
    readonly state: string;
    /**
     * System generated globally unique ID for the environment. This ID will be different if the environment is deleted and re-created with the same name.
     */
    readonly uid: string;
    /**
     * The time when the environment was last updated.
     */
    readonly updateTime: string;
}

export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnvironmentResult> {
    return pulumi.output(args).apply(a => getEnvironment(a, opts))
}

export interface GetEnvironmentOutputArgs {
    environmentId: pulumi.Input<string>;
    lakeId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
