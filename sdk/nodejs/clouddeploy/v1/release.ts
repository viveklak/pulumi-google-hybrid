// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Release in a given project and location.
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Release extends pulumi.CustomResource {
    /**
     * Get an existing Release resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Release {
        return new Release(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:clouddeploy/v1:Release';

    /**
     * Returns true if the given object is an instance of Release.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Release {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Release.__pulumiType;
    }

    /**
     * User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * List of artifacts to pass through to Skaffold command.
     */
    public readonly buildArtifacts!: pulumi.Output<outputs.clouddeploy.v1.BuildArtifactResponse[]>;
    /**
     * Time at which the `Release` was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Snapshot of the parent pipeline taken at release creation time.
     */
    public /*out*/ readonly deliveryPipelineSnapshot!: pulumi.Output<outputs.clouddeploy.v1.DeliveryPipelineResponse>;
    /**
     * Description of the `Release`. Max length is 255 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Optional. Name of the `Release`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/a-z{0,62}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Time at which the render completed.
     */
    public /*out*/ readonly renderEndTime!: pulumi.Output<string>;
    /**
     * Time at which the render began.
     */
    public /*out*/ readonly renderStartTime!: pulumi.Output<string>;
    /**
     * Current state of the render operation.
     */
    public /*out*/ readonly renderState!: pulumi.Output<string>;
    /**
     * Filepath of the Skaffold config inside of the config URI.
     */
    public readonly skaffoldConfigPath!: pulumi.Output<string>;
    /**
     * Cloud Storage URI of tar.gz archive containing Skaffold configuration.
     */
    public readonly skaffoldConfigUri!: pulumi.Output<string>;
    /**
     * The Skaffold version to use when operating on this release, such as "1.20.0". Not all versions are valid; Google Cloud Deploy supports a specific set of versions. If unset, the most recent supported Skaffold version will be used.
     */
    public readonly skaffoldVersion!: pulumi.Output<string>;
    /**
     * Map from target ID to the target artifacts created during the render operation.
     */
    public /*out*/ readonly targetArtifacts!: pulumi.Output<{[key: string]: string}>;
    /**
     * Map from target ID to details of the render operation for that target.
     */
    public /*out*/ readonly targetRenders!: pulumi.Output<{[key: string]: string}>;
    /**
     * Snapshot of the parent pipeline's targets taken at release creation time.
     */
    public /*out*/ readonly targetSnapshots!: pulumi.Output<outputs.clouddeploy.v1.TargetResponse[]>;
    /**
     * Unique identifier of the `Release`.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;

    /**
     * Create a Release resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReleaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deliveryPipelineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deliveryPipelineId'");
            }
            if ((!args || args.releaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'releaseId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["buildArtifacts"] = args ? args.buildArtifacts : undefined;
            resourceInputs["deliveryPipelineId"] = args ? args.deliveryPipelineId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["releaseId"] = args ? args.releaseId : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["skaffoldConfigPath"] = args ? args.skaffoldConfigPath : undefined;
            resourceInputs["skaffoldConfigUri"] = args ? args.skaffoldConfigUri : undefined;
            resourceInputs["skaffoldVersion"] = args ? args.skaffoldVersion : undefined;
            resourceInputs["validateOnly"] = args ? args.validateOnly : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deliveryPipelineSnapshot"] = undefined /*out*/;
            resourceInputs["renderEndTime"] = undefined /*out*/;
            resourceInputs["renderStartTime"] = undefined /*out*/;
            resourceInputs["renderState"] = undefined /*out*/;
            resourceInputs["targetArtifacts"] = undefined /*out*/;
            resourceInputs["targetRenders"] = undefined /*out*/;
            resourceInputs["targetSnapshots"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        } else {
            resourceInputs["annotations"] = undefined /*out*/;
            resourceInputs["buildArtifacts"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deliveryPipelineSnapshot"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["renderEndTime"] = undefined /*out*/;
            resourceInputs["renderStartTime"] = undefined /*out*/;
            resourceInputs["renderState"] = undefined /*out*/;
            resourceInputs["skaffoldConfigPath"] = undefined /*out*/;
            resourceInputs["skaffoldConfigUri"] = undefined /*out*/;
            resourceInputs["skaffoldVersion"] = undefined /*out*/;
            resourceInputs["targetArtifacts"] = undefined /*out*/;
            resourceInputs["targetRenders"] = undefined /*out*/;
            resourceInputs["targetSnapshots"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Release.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Release resource.
 */
export interface ReleaseArgs {
    /**
     * User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of artifacts to pass through to Skaffold command.
     */
    buildArtifacts?: pulumi.Input<pulumi.Input<inputs.clouddeploy.v1.BuildArtifactArgs>[]>;
    deliveryPipelineId: pulumi.Input<string>;
    /**
     * Description of the `Release`. Max length is 255 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Optional. Name of the `Release`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/a-z{0,62}.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    releaseId: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * Filepath of the Skaffold config inside of the config URI.
     */
    skaffoldConfigPath?: pulumi.Input<string>;
    /**
     * Cloud Storage URI of tar.gz archive containing Skaffold configuration.
     */
    skaffoldConfigUri?: pulumi.Input<string>;
    /**
     * The Skaffold version to use when operating on this release, such as "1.20.0". Not all versions are valid; Google Cloud Deploy supports a specific set of versions. If unset, the most recent supported Skaffold version will be used.
     */
    skaffoldVersion?: pulumi.Input<string>;
    validateOnly?: pulumi.Input<string>;
}
