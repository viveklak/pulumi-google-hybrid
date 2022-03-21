// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Endpoint in a given project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:ids/v1:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * The create time timestamp.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * User-provided description of the endpoint
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The fully qualified URL of the endpoint's ILB Forwarding Rule.
     */
    public /*out*/ readonly endpointForwardingRule!: pulumi.Output<string>;
    /**
     * The IP address of the IDS Endpoint's ILB.
     */
    public /*out*/ readonly endpointIp!: pulumi.Output<string>;
    /**
     * The labels of the endpoint.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of the endpoint.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The fully qualified URL of the network to which the IDS Endpoint is attached.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Lowest threat severity that this endpoint will alert on.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Current state of the endpoint.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Whether the endpoint should report traffic logs in addition to threat logs.
     */
    public readonly trafficLogs!: pulumi.Output<boolean>;
    /**
     * The update time timestamp.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.endpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointId'");
            }
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            if ((!args || args.severity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'severity'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointId"] = args ? args.endpointId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["trafficLogs"] = args ? args.trafficLogs : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpointForwardingRule"] = undefined /*out*/;
            resourceInputs["endpointIp"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["endpointForwardingRule"] = undefined /*out*/;
            resourceInputs["endpointIp"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["severity"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["trafficLogs"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Endpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * User-provided description of the endpoint
     */
    description?: pulumi.Input<string>;
    /**
     * Required. The endpoint identifier. This will be part of the endpoint's resource name. This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. Values that do not match this pattern will trigger an INVALID_ARGUMENT error.
     */
    endpointId: pulumi.Input<string>;
    /**
     * The labels of the endpoint.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The fully qualified URL of the network to which the IDS Endpoint is attached.
     */
    network: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Lowest threat severity that this endpoint will alert on.
     */
    severity: pulumi.Input<enums.ids.v1.EndpointSeverity>;
    /**
     * Whether the endpoint should report traffic logs in addition to threat logs.
     */
    trafficLogs?: pulumi.Input<boolean>;
}
