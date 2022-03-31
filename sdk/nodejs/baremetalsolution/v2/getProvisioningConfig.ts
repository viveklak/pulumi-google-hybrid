// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Get ProvisioningConfig by name.
 */
export function getProvisioningConfig(args: GetProvisioningConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetProvisioningConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:baremetalsolution/v2:getProvisioningConfig", {
        "location": args.location,
        "project": args.project,
        "provisioningConfigId": args.provisioningConfigId,
    }, opts);
}

export interface GetProvisioningConfigArgs {
    location: string;
    project?: string;
    provisioningConfigId: string;
}

export interface GetProvisioningConfigResult {
    /**
     * URI to Cloud Console UI view of this provisioning config.
     */
    readonly cloudConsoleUri: string;
    /**
     * Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
     *
     * @deprecated Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
     */
    readonly email: string;
    /**
     * A service account to enable customers to access instance credentials upon handover.
     */
    readonly handoverServiceAccount: string;
    /**
     * Instances to be created.
     */
    readonly instances: outputs.baremetalsolution.v2.InstanceConfigResponse[];
    /**
     * Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
     */
    readonly location: string;
    /**
     * The name of the provisioning config.
     */
    readonly name: string;
    /**
     * Networks to be created.
     */
    readonly networks: outputs.baremetalsolution.v2.NetworkConfigResponse[];
    /**
     * State of ProvisioningConfig.
     */
    readonly state: string;
    /**
     * A generated buganizer id to track provisioning request.
     */
    readonly ticketId: string;
    /**
     * Last update timestamp.
     */
    readonly updateTime: string;
    /**
     * Volumes to be created.
     */
    readonly volumes: outputs.baremetalsolution.v2.VolumeConfigResponse[];
}

export function getProvisioningConfigOutput(args: GetProvisioningConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProvisioningConfigResult> {
    return pulumi.output(args).apply(a => getProvisioningConfig(a, opts))
}

export interface GetProvisioningConfigOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    provisioningConfigId: pulumi.Input<string>;
}
