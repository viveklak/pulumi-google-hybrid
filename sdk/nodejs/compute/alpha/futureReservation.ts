// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Future Reservation.
 */
export class FutureReservation extends pulumi.CustomResource {
    /**
     * Get an existing FutureReservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FutureReservation {
        return new FutureReservation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:FutureReservation';

    /**
     * Returns true if the given object is an instance of FutureReservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FutureReservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FutureReservation.__pulumiType;
    }

    /**
     * The creation timestamp for this future reservation in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the future reservation.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly kind!: pulumi.Output<string>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * List of Projects/Folders to share with.
     */
    public readonly shareSettings!: pulumi.Output<outputs.compute.alpha.ShareSettingsResponse>;
    /**
     * Future Reservation configuration to indicate instance properties and total count.
     */
    public readonly specificSkuProperties!: pulumi.Output<outputs.compute.alpha.FutureReservationSpecificSKUPropertiesResponse>;
    /**
     * [Output only] Status of the Future Reservation
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.compute.alpha.FutureReservationStatusResponse>;
    /**
     * Time window for this Future Reservation.
     */
    public readonly timeWindow!: pulumi.Output<outputs.compute.alpha.FutureReservationTimeWindowResponse>;
    /**
     * URL of the Zone where this future reservation resides.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a FutureReservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FutureReservationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["shareSettings"] = args ? args.shareSettings : undefined;
            resourceInputs["specificSkuProperties"] = args ? args.specificSkuProperties : undefined;
            resourceInputs["timeWindow"] = args ? args.timeWindow : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namePrefix"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["shareSettings"] = undefined /*out*/;
            resourceInputs["specificSkuProperties"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["timeWindow"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FutureReservation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FutureReservation resource.
 */
export interface FutureReservationArgs {
    /**
     * An optional description of this resource. Provide this property when you create the future reservation.
     */
    description?: pulumi.Input<string>;
    kind?: pulumi.Input<string>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
     */
    namePrefix?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * List of Projects/Folders to share with.
     */
    shareSettings?: pulumi.Input<inputs.compute.alpha.ShareSettingsArgs>;
    /**
     * Future Reservation configuration to indicate instance properties and total count.
     */
    specificSkuProperties?: pulumi.Input<inputs.compute.alpha.FutureReservationSpecificSKUPropertiesArgs>;
    /**
     * Time window for this Future Reservation.
     */
    timeWindow?: pulumi.Input<inputs.compute.alpha.FutureReservationTimeWindowArgs>;
    zone?: pulumi.Input<string>;
}
