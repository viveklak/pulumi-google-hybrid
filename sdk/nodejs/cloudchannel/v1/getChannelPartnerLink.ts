// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the requested ChannelPartnerLink resource. You must be a distributor to call this method. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: Required request parameters are missing or invalid. * NOT_FOUND: ChannelPartnerLink resource not found because of an invalid channel partner link name. Return value: The ChannelPartnerLink resource.
 */
export function getChannelPartnerLink(args: GetChannelPartnerLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetChannelPartnerLinkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudchannel/v1:getChannelPartnerLink", {
        "accountId": args.accountId,
        "channelPartnerLinkId": args.channelPartnerLinkId,
        "view": args.view,
    }, opts);
}

export interface GetChannelPartnerLinkArgs {
    accountId: string;
    channelPartnerLinkId: string;
    view?: string;
}

export interface GetChannelPartnerLinkResult {
    /**
     * Cloud Identity info of the channel partner (IR).
     */
    readonly channelPartnerCloudIdentityInfo: outputs.cloudchannel.v1.GoogleCloudChannelV1CloudIdentityInfoResponse;
    /**
     * Timestamp of when the channel partner link is created.
     */
    readonly createTime: string;
    /**
     * URI of the web page where partner accepts the link invitation.
     */
    readonly inviteLinkUri: string;
    /**
     * State of the channel partner link.
     */
    readonly linkState: string;
    /**
     * Resource name for the channel partner link, in the format accounts/{account_id}/channelPartnerLinks/{id}.
     */
    readonly name: string;
    /**
     * Public identifier that a customer must use to generate a transfer token to move to this distributor-reseller combination.
     */
    readonly publicId: string;
    /**
     * Cloud Identity ID of the linked reseller.
     */
    readonly resellerCloudIdentityId: string;
    /**
     * Timestamp of when the channel partner link is updated.
     */
    readonly updateTime: string;
}

export function getChannelPartnerLinkOutput(args: GetChannelPartnerLinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChannelPartnerLinkResult> {
    return pulumi.output(args).apply(a => getChannelPartnerLink(a, opts))
}

export interface GetChannelPartnerLinkOutputArgs {
    accountId: pulumi.Input<string>;
    channelPartnerLinkId: pulumi.Input<string>;
    view?: pulumi.Input<string>;
}
