// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified SSL policy resource. Gets a list of available SSL policies by making a list() request.
 */
export class SslPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SslPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SslPolicy {
        return new SslPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:SslPolicy';

    /**
     * Returns true if the given object is an instance of SslPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SslPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SslPolicy.__pulumiType;
    }

    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * A list of features enabled when the selected profile is CUSTOM. The
     * - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
     */
    public readonly customFeatures!: pulumi.Output<string[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * [Output Only] The list of features enabled in the SSL policy.
     */
    public readonly enabledFeatures!: pulumi.Output<string[]>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
     */
    public readonly fingerprint!: pulumi.Output<string>;
    /**
     * [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
     */
    public readonly minTlsVersion!: pulumi.Output<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
     */
    public readonly profile!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for this resource with the resource id.
     */
    public readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Security settings for the proxy. This field is only applicable to a global backend service with the loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     */
    public readonly tlsSettings!: pulumi.Output<outputs.compute.alpha.ServerTlsSettingsResponse>;
    /**
     * [Output Only] If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
     */
    public readonly warnings!: pulumi.Output<outputs.compute.alpha.SslPolicyWarningsItemResponse[]>;

    /**
     * Create a SslPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["customFeatures"] = args ? args.customFeatures : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabledFeatures"] = args ? args.enabledFeatures : undefined;
            inputs["fingerprint"] = args ? args.fingerprint : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["minTlsVersion"] = args ? args.minTlsVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["profile"] = args ? args.profile : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["selfLinkWithId"] = args ? args.selfLinkWithId : undefined;
            inputs["tlsSettings"] = args ? args.tlsSettings : undefined;
            inputs["warnings"] = args ? args.warnings : undefined;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["customFeatures"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["enabledFeatures"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["minTlsVersion"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["profile"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["tlsSettings"] = undefined /*out*/;
            inputs["warnings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SslPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SslPolicy resource.
 */
export interface SslPolicyArgs {
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * A list of features enabled when the selected profile is CUSTOM. The
     * - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
     */
    readonly customFeatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * [Output Only] The list of features enabled in the SSL policy.
     */
    readonly enabledFeatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
     */
    readonly minTlsVersion?: pulumi.Input<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
     */
    readonly profile?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
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
     * Security settings for the proxy. This field is only applicable to a global backend service with the loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     */
    readonly tlsSettings?: pulumi.Input<inputs.compute.alpha.ServerTlsSettingsArgs>;
    /**
     * [Output Only] If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
     */
    readonly warnings?: pulumi.Input<pulumi.Input<inputs.compute.alpha.SslPolicyWarningsItemArgs>[]>;
}
