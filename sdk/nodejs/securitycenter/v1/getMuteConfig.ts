// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets a mute config.
 */
export function getMuteConfig(args: GetMuteConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetMuteConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:securitycenter/v1:getMuteConfig", {
        "muteConfigId": args.muteConfigId,
        "project": args.project,
    }, opts);
}

export interface GetMuteConfigArgs {
    muteConfigId: string;
    project?: string;
}

export interface GetMuteConfigResult {
    /**
     * The time at which the mute config was created. This field is set by the server and will be ignored if provided on config creation.
     */
    readonly createTime: string;
    /**
     * A description of the mute config.
     */
    readonly description: string;
    /**
     * The human readable name to be displayed for the mute config.
     */
    readonly displayName: string;
    /**
     * An expression that defines the filter to apply across create/update events of findings. While creating a filter string, be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is created under the project = Y scope, it might not match any findings. The following field and operator combinations are supported: * severity: `=`, `:` * category: `=`, `:` * resource.name: `=`, `:` * resource.project_name: `=`, `:` * resource.project_display_name: `=`, `:` * resource.folders.resource_folder: `=`, `:` * resource.parent_name: `=`, `:` * resource.parent_display_name: `=`, `:` * resource.type: `=`, `:` * finding_class: `=`, `:` * indicator.ip_addresses: `=`, `:` * indicator.domains: `=`, `:`
     */
    readonly filter: string;
    /**
     * Email address of the user who last edited the mute config. This field is set by the server and will be ignored if provided on config creation or update.
     */
    readonly mostRecentEditor: string;
    /**
     * This field will be ignored if provided on config creation. Format "organizations/{organization}/muteConfigs/{mute_config}" "folders/{folder}/muteConfigs/{mute_config}" "projects/{project}/muteConfigs/{mute_config}"
     */
    readonly name: string;
    /**
     * The most recent time at which the mute config was updated. This field is set by the server and will be ignored if provided on config creation or update.
     */
    readonly updateTime: string;
}

export function getMuteConfigOutput(args: GetMuteConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMuteConfigResult> {
    return pulumi.output(args).apply(a => getMuteConfig(a, opts))
}

export interface GetMuteConfigOutputArgs {
    muteConfigId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
