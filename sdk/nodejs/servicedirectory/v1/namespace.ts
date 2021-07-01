// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a namespace, and returns the new namespace.
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:servicedirectory/v1:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * Optional. Resource labels associated with this namespace. No more than 64 user labels can be associated with a given resource. Label keys and values can be no longer than 63 characters.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Immutable. The resource name for the namespace in the format `projects/*&#47;locations/*&#47;namespaces/*`.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.namespaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceId"] = args ? args.namespaceId : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Namespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * Optional. Resource labels associated with this namespace. No more than 64 user labels can be associated with a given resource. Label keys and values can be no longer than 63 characters.
     */
    labels?: pulumi.Input<{[key: string]: string}>;
    location: pulumi.Input<string>;
    /**
     * Immutable. The resource name for the namespace in the format `projects/*&#47;locations/*&#47;namespaces/*`.
     */
    name?: pulumi.Input<string>;
    namespaceId: pulumi.Input<string>;
    project: pulumi.Input<string>;
}
