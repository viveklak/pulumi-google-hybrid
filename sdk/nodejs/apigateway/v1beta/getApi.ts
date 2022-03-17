// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Api.
 */
export function getApi(args: GetApiArgs, opts?: pulumi.InvokeOptions): Promise<GetApiResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:apigateway/v1beta:getApi", {
        "apiId": args.apiId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetApiArgs {
    apiId: string;
    location: string;
    project?: string;
}

export interface GetApiResult {
    /**
     * Created time.
     */
    readonly createTime: string;
    /**
     * Optional. Display name.
     */
    readonly displayName: string;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    readonly labels: {[key: string]: string};
    /**
     * Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
     */
    readonly managedService: string;
    /**
     * Resource name of the API. Format: projects/{project}/locations/global/apis/{api}
     */
    readonly name: string;
    /**
     * State of the API.
     */
    readonly state: string;
    /**
     * Updated time.
     */
    readonly updateTime: string;
}

export function getApiOutput(args: GetApiOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiResult> {
    return pulumi.output(args).apply(a => getApi(a, opts))
}

export interface GetApiOutputArgs {
    apiId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
