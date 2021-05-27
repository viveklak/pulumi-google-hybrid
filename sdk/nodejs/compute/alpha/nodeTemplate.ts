// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a NodeTemplate resource in the specified project using the data included in the request.
 */
export class NodeTemplate extends pulumi.CustomResource {
    /**
     * Get an existing NodeTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NodeTemplate {
        return new NodeTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:NodeTemplate';

    /**
     * Returns true if the given object is an instance of NodeTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeTemplate.__pulumiType;
    }

    public readonly accelerators!: pulumi.Output<outputs.compute.alpha.AcceleratorConfigResponse[]>;
    /**
     * CPU overcommit.
     */
    public readonly cpuOvercommitType!: pulumi.Output<string>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly disks!: pulumi.Output<outputs.compute.alpha.LocalDiskResponse[]>;
    /**
     * [Output Only] The type of the resource. Always compute#nodeTemplate for node templates.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Labels to use for node affinity, which will be used in instance scheduling.
     */
    public readonly nodeAffinityLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The node type to use for nodes group that are created from this template.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties.
     *
     * This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
     */
    public readonly nodeTypeFlexibility!: pulumi.Output<outputs.compute.alpha.NodeTemplateNodeTypeFlexibilityResponse>;
    /**
     * [Output Only] The name of the region where the node template resides, such as us-central1.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for this resource with the resource id.
     */
    public readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Sets the binding properties for the physical server. Valid values include:  
     * - [Default] RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server 
     * - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible  
     *
     * See Sole-tenant node options for more information.
     */
    public readonly serverBinding!: pulumi.Output<outputs.compute.alpha.ServerBindingResponse>;
    /**
     * [Output Only] The status of the node template. One of the following values: CREATING, READY, and DELETING.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * [Output Only] An optional, human-readable explanation of the status.
     */
    public readonly statusMessage!: pulumi.Output<string>;

    /**
     * Create a NodeTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["accelerators"] = args ? args.accelerators : undefined;
            inputs["cpuOvercommitType"] = args ? args.cpuOvercommitType : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disks"] = args ? args.disks : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeAffinityLabels"] = args ? args.nodeAffinityLabels : undefined;
            inputs["nodeType"] = args ? args.nodeType : undefined;
            inputs["nodeTypeFlexibility"] = args ? args.nodeTypeFlexibility : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["selfLinkWithId"] = args ? args.selfLinkWithId : undefined;
            inputs["serverBinding"] = args ? args.serverBinding : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["statusMessage"] = args ? args.statusMessage : undefined;
        } else {
            inputs["accelerators"] = undefined /*out*/;
            inputs["cpuOvercommitType"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["disks"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodeAffinityLabels"] = undefined /*out*/;
            inputs["nodeType"] = undefined /*out*/;
            inputs["nodeTypeFlexibility"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["serverBinding"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NodeTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NodeTemplate resource.
 */
export interface NodeTemplateArgs {
    readonly accelerators?: pulumi.Input<pulumi.Input<inputs.compute.alpha.AcceleratorConfigArgs>[]>;
    /**
     * CPU overcommit.
     */
    readonly cpuOvercommitType?: pulumi.Input<string>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    readonly disks?: pulumi.Input<pulumi.Input<inputs.compute.alpha.LocalDiskArgs>[]>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output Only] The type of the resource. Always compute#nodeTemplate for node templates.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Labels to use for node affinity, which will be used in instance scheduling.
     */
    readonly nodeAffinityLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The node type to use for nodes group that are created from this template.
     */
    readonly nodeType?: pulumi.Input<string>;
    /**
     * The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties.
     *
     * This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
     */
    readonly nodeTypeFlexibility?: pulumi.Input<inputs.compute.alpha.NodeTemplateNodeTypeFlexibilityArgs>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] The name of the region where the node template resides, such as us-central1.
     */
    readonly region: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId?: pulumi.Input<string>;
    /**
     * Sets the binding properties for the physical server. Valid values include:  
     * - [Default] RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server 
     * - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible  
     *
     * See Sole-tenant node options for more information.
     */
    readonly serverBinding?: pulumi.Input<inputs.compute.alpha.ServerBindingArgs>;
    /**
     * [Output Only] The status of the node template. One of the following values: CREATING, READY, and DELETING.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * [Output Only] An optional, human-readable explanation of the status.
     */
    readonly statusMessage?: pulumi.Input<string>;
}
