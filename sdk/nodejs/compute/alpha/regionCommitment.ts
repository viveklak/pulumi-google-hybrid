// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a commitment in the specified project using the data included in the request.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class RegionCommitment extends pulumi.CustomResource {
    /**
     * Get an existing RegionCommitment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionCommitment {
        return new RegionCommitment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:RegionCommitment';

    /**
     * Returns true if the given object is an instance of RegionCommitment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionCommitment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionCommitment.__pulumiType;
    }

    /**
     * Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
     */
    public readonly autoRenew!: pulumi.Output<boolean>;
    /**
     * The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Commitment end time in RFC3339 text format.
     */
    public /*out*/ readonly endTimestamp!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#commitment for commitments.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The license specification required as part of a license commitment.
     */
    public readonly licenseResource!: pulumi.Output<outputs.compute.alpha.LicenseResourceCommitmentResponse>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
     */
    public readonly plan!: pulumi.Output<string>;
    /**
     * URL of the region where this commitment may be used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * List of reservations in this commitment.
     */
    public readonly reservations!: pulumi.Output<outputs.compute.alpha.ReservationResponse[]>;
    /**
     * A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
     */
    public readonly resources!: pulumi.Output<outputs.compute.alpha.ResourceCommitmentResponse[]>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Commitment start time in RFC3339 text format.
     */
    public /*out*/ readonly startTimestamp!: pulumi.Output<string>;
    /**
     * Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * An optional, human-readable explanation of the status.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a RegionCommitment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionCommitmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["category"] = args ? args.category : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["licenseResource"] = args ? args.licenseResource : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["reservations"] = args ? args.reservations : undefined;
            inputs["resources"] = args ? args.resources : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["endTimestamp"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["startTimestamp"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
        } else {
            inputs["autoRenew"] = undefined /*out*/;
            inputs["category"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["endTimestamp"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["licenseResource"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["plan"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["reservations"] = undefined /*out*/;
            inputs["resources"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["startTimestamp"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionCommitment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionCommitment resource.
 */
export interface RegionCommitmentArgs {
    /**
     * Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
     */
    category?: pulumi.Input<enums.compute.alpha.RegionCommitmentCategory>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The license specification required as part of a license commitment.
     */
    licenseResource?: pulumi.Input<inputs.compute.alpha.LicenseResourceCommitmentArgs>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
     */
    plan?: pulumi.Input<enums.compute.alpha.RegionCommitmentPlan>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * List of reservations in this commitment.
     */
    reservations?: pulumi.Input<pulumi.Input<inputs.compute.alpha.ReservationArgs>[]>;
    /**
     * A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.compute.alpha.ResourceCommitmentArgs>[]>;
    /**
     * The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
     */
    type?: pulumi.Input<enums.compute.alpha.RegionCommitmentType>;
}
