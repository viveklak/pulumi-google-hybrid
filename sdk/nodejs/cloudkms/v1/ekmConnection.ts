// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new EkmConnection in a given Project and Location.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class EkmConnection extends pulumi.CustomResource {
    /**
     * Get an existing EkmConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EkmConnection {
        return new EkmConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudkms/v1:EkmConnection';

    /**
     * Returns true if the given object is an instance of EkmConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EkmConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EkmConnection.__pulumiType;
    }

    /**
     * The time at which the EkmConnection was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update requests to ensure the client has an up-to-date value before proceeding.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The resource name for the EkmConnection in the format `projects/*&#47;locations/*&#47;ekmConnections/*`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
     */
    public readonly serviceResolvers!: pulumi.Output<outputs.cloudkms.v1.ServiceResolverResponse[]>;

    /**
     * Create a EkmConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EkmConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["ekmConnectionId"] = args ? args.ekmConnectionId : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceResolvers"] = args ? args.serviceResolvers : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceResolvers"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EkmConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EkmConnection resource.
 */
export interface EkmConnectionArgs {
    ekmConnectionId?: pulumi.Input<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
     */
    serviceResolvers?: pulumi.Input<pulumi.Input<inputs.cloudkms.v1.ServiceResolverArgs>[]>;
}