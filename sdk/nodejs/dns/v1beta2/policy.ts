// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Policy.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dns/v1beta2:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
     */
    public readonly alternativeNameServerConfig!: pulumi.Output<outputs.dns.v1beta2.PolicyAlternativeNameServerConfigResponse>;
    /**
     * A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
     */
    public readonly enableInboundForwarding!: pulumi.Output<boolean>;
    /**
     * Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
     */
    public readonly enableLogging!: pulumi.Output<boolean>;
    public readonly kind!: pulumi.Output<string>;
    /**
     * User-assigned name for this policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of network names specifying networks to which this policy is applied.
     */
    public readonly networks!: pulumi.Output<outputs.dns.v1beta2.PolicyNetworkResponse[]>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["alternativeNameServerConfig"] = args ? args.alternativeNameServerConfig : undefined;
            inputs["clientOperationId"] = args ? args.clientOperationId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableInboundForwarding"] = args ? args.enableInboundForwarding : undefined;
            inputs["enableLogging"] = args ? args.enableLogging : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networks"] = args ? args.networks : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["alternativeNameServerConfig"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["enableInboundForwarding"] = undefined /*out*/;
            inputs["enableLogging"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networks"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
     */
    readonly alternativeNameServerConfig?: pulumi.Input<inputs.dns.v1beta2.PolicyAlternativeNameServerConfigArgs>;
    readonly clientOperationId?: pulumi.Input<string>;
    /**
     * A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
     */
    readonly enableInboundForwarding?: pulumi.Input<boolean>;
    /**
     * Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    /**
     * Unique identifier for the resource; defined by the server (output only).
     */
    readonly id?: pulumi.Input<string>;
    readonly kind?: pulumi.Input<string>;
    /**
     * User-assigned name for this policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of network names specifying networks to which this policy is applied.
     */
    readonly networks?: pulumi.Input<pulumi.Input<inputs.dns.v1beta2.PolicyNetworkArgs>[]>;
    readonly project: pulumi.Input<string>;
}
