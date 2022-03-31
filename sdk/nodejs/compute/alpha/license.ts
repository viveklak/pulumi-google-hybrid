// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a License resource in the specified project. *Caution* This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
 */
export class License extends pulumi.CustomResource {
    /**
     * Get an existing License resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): License {
        return new License(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:License';

    /**
     * Returns true if the given object is an instance of License.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is License {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === License.__pulumiType;
    }

    /**
     * Deprecated. This field no longer reflects whether a license charges a usage fee.
     *
     * @deprecated [Output Only] Deprecated. This field no longer reflects whether a license charges a usage fee.
     */
    public /*out*/ readonly chargesUseFee!: pulumi.Output<boolean>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Type of resource. Always compute#license for licenses.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The unique code used to attach this license to images, snapshots, and disks.
     */
    public /*out*/ readonly licenseCode!: pulumi.Output<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly resourceRequirements!: pulumi.Output<outputs.compute.alpha.LicenseResourceRequirementsResponse>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
     */
    public readonly transferable!: pulumi.Output<boolean>;

    /**
     * Create a License resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LicenseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["resourceRequirements"] = args ? args.resourceRequirements : undefined;
            resourceInputs["transferable"] = args ? args.transferable : undefined;
            resourceInputs["chargesUseFee"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["licenseCode"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
        } else {
            resourceInputs["chargesUseFee"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["licenseCode"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["resourceRequirements"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["transferable"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(License.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a License resource.
 */
export interface LicenseArgs {
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    resourceRequirements?: pulumi.Input<inputs.compute.alpha.LicenseResourceRequirementsArgs>;
    /**
     * If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
     */
    transferable?: pulumi.Input<boolean>;
}
