// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
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
    public static readonly __pulumiType = 'google-native:compute/v1:NodeTemplate';

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

    public readonly accelerators!: pulumi.Output<outputs.compute.v1.AcceleratorConfigResponse[]>;
    /**
     * CPU overcommit.
     */
    public readonly cpuOvercommitType!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly disks!: pulumi.Output<outputs.compute.v1.LocalDiskResponse[]>;
    /**
     * The type of the resource. Always compute#nodeTemplate for node templates.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
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
     * The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties. This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
     */
    public readonly nodeTypeFlexibility!: pulumi.Output<outputs.compute.v1.NodeTemplateNodeTypeFlexibilityResponse>;
    /**
     * The name of the region where the node template resides, such as us-central1.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
     */
    public readonly serverBinding!: pulumi.Output<outputs.compute.v1.ServerBindingResponse>;
    /**
     * The status of the node template. One of the following values: CREATING, READY, and DELETING.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * An optional, human-readable explanation of the status.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;

    /**
     * Create a NodeTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["accelerators"] = args ? args.accelerators : undefined;
            resourceInputs["cpuOvercommitType"] = args ? args.cpuOvercommitType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disks"] = args ? args.disks : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeAffinityLabels"] = args ? args.nodeAffinityLabels : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["nodeTypeFlexibility"] = args ? args.nodeTypeFlexibility : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["serverBinding"] = args ? args.serverBinding : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
        } else {
            resourceInputs["accelerators"] = undefined /*out*/;
            resourceInputs["cpuOvercommitType"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["disks"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nodeAffinityLabels"] = undefined /*out*/;
            resourceInputs["nodeType"] = undefined /*out*/;
            resourceInputs["nodeTypeFlexibility"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["serverBinding"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NodeTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NodeTemplate resource.
 */
export interface NodeTemplateArgs {
    accelerators?: pulumi.Input<pulumi.Input<inputs.compute.v1.AcceleratorConfigArgs>[]>;
    /**
     * CPU overcommit.
     */
    cpuOvercommitType?: pulumi.Input<enums.compute.v1.NodeTemplateCpuOvercommitType>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    disks?: pulumi.Input<pulumi.Input<inputs.compute.v1.LocalDiskArgs>[]>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * Labels to use for node affinity, which will be used in instance scheduling.
     */
    nodeAffinityLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The node type to use for nodes group that are created from this template.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties. This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
     */
    nodeTypeFlexibility?: pulumi.Input<inputs.compute.v1.NodeTemplateNodeTypeFlexibilityArgs>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
     */
    serverBinding?: pulumi.Input<inputs.compute.v1.ServerBindingArgs>;
}
