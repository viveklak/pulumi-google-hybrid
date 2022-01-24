// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves information for the specified channel of the specified site.
 */
export function getChannel(args: GetChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetChannelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:firebasehosting/v1beta1:getChannel", {
        "channelId": args.channelId,
        "project": args.project,
        "siteId": args.siteId,
    }, opts);
}

export interface GetChannelArgs {
    channelId: string;
    project?: string;
    siteId: string;
}

export interface GetChannelResult {
    /**
     * The time at which the channel was created.
     */
    readonly createTime: string;
    /**
     * The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
     */
    readonly expireTime: string;
    /**
     * Text labels used for extra metadata and/or filtering.
     */
    readonly labels: {[key: string]: string};
    /**
     * The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
     */
    readonly name: string;
    /**
     * The current release for the channel, if any.
     */
    readonly release: outputs.firebasehosting.v1beta1.ReleaseResponse;
    /**
     * The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
     */
    readonly retainedReleaseCount: number;
    /**
     * Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
     */
    readonly ttl: string;
    /**
     * The time at which the channel was last updated.
     */
    readonly updateTime: string;
    /**
     * The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
     */
    readonly url: string;
}

export function getChannelOutput(args: GetChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChannelResult> {
    return pulumi.output(args).apply(a => getChannel(a, opts))
}

export interface GetChannelOutputArgs {
    channelId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    siteId: pulumi.Input<string>;
}
