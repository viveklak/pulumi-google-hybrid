// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Version in the specified Flow.
 */
export class AgentFlowVersion extends pulumi.CustomResource {
    /**
     * Get an existing AgentFlowVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentFlowVersion {
        return new AgentFlowVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:AgentFlowVersion';

    /**
     * Returns true if the given object is an instance of AgentFlowVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentFlowVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentFlowVersion.__pulumiType;
    }

    /**
     * Create time of the version.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. The human-readable name of the version. Limit of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The NLU settings of the flow at version creation.
     */
    public /*out*/ readonly nluSettings!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3NluSettingsResponse>;
    /**
     * The state of this version. This field is read-only and cannot be set by create and update methods.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a AgentFlowVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentFlowVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.flowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["agentId"] = args ? args.agentId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["flowId"] = args ? args.flowId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["nluSettings"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nluSettings"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentFlowVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentFlowVersion resource.
 */
export interface AgentFlowVersionArgs {
    readonly agentId: pulumi.Input<string>;
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Required. The human-readable name of the version. Limit of 64 characters.
     */
    readonly displayName?: pulumi.Input<string>;
    readonly flowId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    /**
     * Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
}
