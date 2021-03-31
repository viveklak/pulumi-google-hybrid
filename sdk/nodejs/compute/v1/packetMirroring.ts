// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a PacketMirroring resource in the specified project and region using the data included in the request.
 */
export class PacketMirroring extends pulumi.CustomResource {
    /**
     * Get an existing PacketMirroring resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PacketMirroring {
        return new PacketMirroring(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:compute/v1:PacketMirroring';

    /**
     * Returns true if the given object is an instance of PacketMirroring.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PacketMirroring {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PacketMirroring.__pulumiType;
    }


    /**
     * Create a PacketMirroring resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PacketMirroringArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.packetMirroring === undefined) && !opts.urn) {
                throw new Error("Missing required property 'packetMirroring'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["collectorIlb"] = args ? args.collectorIlb : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enable"] = args ? args.enable : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["mirroredResources"] = args ? args.mirroredResources : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["packetMirroring"] = args ? args.packetMirroring : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PacketMirroring.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PacketMirroring resource.
 */
export interface PacketMirroringArgs {
    /**
     * The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
     */
    readonly collectorIlb?: pulumi.Input<inputs.compute.v1.PacketMirroringForwardingRuleInfo>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network.
     *
     * The default is TRUE.
     */
    readonly enable?: pulumi.Input<string>;
    /**
     * Filter for mirrored traffic. If unspecified, all traffic is mirrored.
     */
    readonly filter?: pulumi.Input<inputs.compute.v1.PacketMirroringFilter>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output Only] Type of the resource. Always compute#packetMirroring for packet mirrorings.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
     */
    readonly mirroredResources?: pulumi.Input<inputs.compute.v1.PacketMirroringMirroredResourceInfo>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
     */
    readonly network?: pulumi.Input<inputs.compute.v1.PacketMirroringNetworkInfo>;
    readonly packetMirroring: pulumi.Input<string>;
    /**
     * The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins.
     *
     * Default value is 1000. Valid range is 0 through 65535.
     */
    readonly priority?: pulumi.Input<number>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] URI of the region where the packetMirroring resides.
     */
    readonly region: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}
