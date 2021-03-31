// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a ExternalVpnGateway in the specified project using the data included in the request.
 */
export class ExternalVpnGateway extends pulumi.CustomResource {
    /**
     * Get an existing ExternalVpnGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExternalVpnGateway {
        return new ExternalVpnGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:compute/beta:ExternalVpnGateway';

    /**
     * Returns true if the given object is an instance of ExternalVpnGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalVpnGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalVpnGateway.__pulumiType;
    }


    /**
     * Create a ExternalVpnGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExternalVpnGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.externalVpnGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalVpnGateway'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["externalVpnGateway"] = args ? args.externalVpnGateway : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["interfaces"] = args ? args.interfaces : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["labelFingerprint"] = args ? args.labelFingerprint : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["redundancyType"] = args ? args.redundancyType : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ExternalVpnGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExternalVpnGateway resource.
 */
export interface ExternalVpnGatewayArgs {
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    readonly externalVpnGateway: pulumi.Input<string>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * List of interfaces for this external VPN gateway.
     */
    readonly interfaces?: pulumi.Input<pulumi.Input<inputs.compute.beta.ExternalVpnGatewayInterface>[]>;
    /**
     * [Output Only] Type of the resource. Always compute#externalVpnGateway for externalVpnGateways.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * A fingerprint for the labels being applied to this ExternalVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an ExternalVpnGateway.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Indicates the user-supplied redundancy type of this external VPN gateway.
     */
    readonly redundancyType?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}
