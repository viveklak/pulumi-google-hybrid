// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a mute config.
 */
export class OrganizationMuteConfig extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationMuteConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationMuteConfig {
        return new OrganizationMuteConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:securitycenter/v1:OrganizationMuteConfig';

    /**
     * Returns true if the given object is an instance of OrganizationMuteConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationMuteConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationMuteConfig.__pulumiType;
    }

    /**
     * The time at which the mute config was created. This field is set by the server and will be ignored if provided on config creation.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A description of the mute config.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The human readable name to be displayed for the mute config.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * An expression that defines the filter to apply across create/update events of findings. While creating a filter string, be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is created under the project = Y scope, it might not match any findings. The following field and operator combinations are supported: * severity: `=`, `:` * category: `=`, `:` * resource.name: `=`, `:` * resource.project_name: `=`, `:` * resource.project_display_name: `=`, `:` * resource.folders.resource_folder: `=`, `:` * resource.parent_name: `=`, `:` * resource.parent_display_name: `=`, `:` * resource.type: `=`, `:` * finding_class: `=`, `:` * indicator.ip_addresses: `=`, `:` * indicator.domains: `=`, `:`
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * Email address of the user who last edited the mute config. This field is set by the server and will be ignored if provided on config creation or update.
     */
    public /*out*/ readonly mostRecentEditor!: pulumi.Output<string>;
    /**
     * This field will be ignored if provided on config creation. Format "organizations/{organization}/muteConfigs/{mute_config}" "folders/{folder}/muteConfigs/{mute_config}" "projects/{project}/muteConfigs/{mute_config}"
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The most recent time at which the mute config was updated. This field is set by the server and will be ignored if provided on config creation or update.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a OrganizationMuteConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationMuteConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.filter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filter'");
            }
            if ((!args || args.muteConfigId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'muteConfigId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["muteConfigId"] = args ? args.muteConfigId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["mostRecentEditor"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["filter"] = undefined /*out*/;
            resourceInputs["mostRecentEditor"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationMuteConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationMuteConfig resource.
 */
export interface OrganizationMuteConfigArgs {
    /**
     * A description of the mute config.
     */
    description?: pulumi.Input<string>;
    /**
     * The human readable name to be displayed for the mute config.
     */
    displayName?: pulumi.Input<string>;
    /**
     * An expression that defines the filter to apply across create/update events of findings. While creating a filter string, be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is created under the project = Y scope, it might not match any findings. The following field and operator combinations are supported: * severity: `=`, `:` * category: `=`, `:` * resource.name: `=`, `:` * resource.project_name: `=`, `:` * resource.project_display_name: `=`, `:` * resource.folders.resource_folder: `=`, `:` * resource.parent_name: `=`, `:` * resource.parent_display_name: `=`, `:` * resource.type: `=`, `:` * finding_class: `=`, `:` * indicator.ip_addresses: `=`, `:` * indicator.domains: `=`, `:`
     */
    filter: pulumi.Input<string>;
    muteConfigId: pulumi.Input<string>;
    /**
     * This field will be ignored if provided on config creation. Format "organizations/{organization}/muteConfigs/{mute_config}" "folders/{folder}/muteConfigs/{mute_config}" "projects/{project}/muteConfigs/{mute_config}"
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
