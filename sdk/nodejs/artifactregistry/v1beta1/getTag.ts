// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets a tag.
 */
export function getTag(args: GetTagArgs, opts?: pulumi.InvokeOptions): Promise<GetTagResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:artifactregistry/v1beta1:getTag", {
        "location": args.location,
        "packageId": args.packageId,
        "project": args.project,
        "repositoryId": args.repositoryId,
        "tagId": args.tagId,
    }, opts);
}

export interface GetTagArgs {
    location: string;
    packageId: string;
    project?: string;
    repositoryId: string;
    tagId: string;
}

export interface GetTagResult {
    /**
     * The name of the tag, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/tags/tag1". If the package part contains slashes, the slashes are escaped. The tag part can only have characters in [a-zA-Z0-9\-._~:@], anything else must be URL encoded.
     */
    readonly name: string;
    /**
     * The name of the version the tag refers to, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811" If the package or version ID parts contain slashes, the slashes are escaped.
     */
    readonly version: string;
}

export function getTagOutput(args: GetTagOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTagResult> {
    return pulumi.output(args).apply(a => getTag(a, opts))
}

export interface GetTagOutputArgs {
    location: pulumi.Input<string>;
    packageId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    repositoryId: pulumi.Input<string>;
    tagId: pulumi.Input<string>;
}
