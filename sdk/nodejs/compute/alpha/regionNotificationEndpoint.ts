// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a NotificationEndpoint in the specified project in the given region using the parameters that are included in the request.
 */
export class RegionNotificationEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing RegionNotificationEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionNotificationEndpoint {
        return new RegionNotificationEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:RegionNotificationEndpoint';

    /**
     * Returns true if the given object is an instance of RegionNotificationEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionNotificationEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionNotificationEndpoint.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
     */
    public readonly grpcSettings!: pulumi.Output<outputs.compute.alpha.NotificationEndpointGrpcSettingsResponse>;
    /**
     * Type of the resource. Always compute#notificationEndpoint for notification endpoints.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL of the region where the notification endpoint resides. This field applies only to the regional resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a RegionNotificationEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionNotificationEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["grpcSettings"] = args ? args.grpcSettings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["grpcSettings"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionNotificationEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionNotificationEndpoint resource.
 */
export interface RegionNotificationEndpointArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
     */
    grpcSettings?: pulumi.Input<inputs.compute.alpha.NotificationEndpointGrpcSettingsArgs>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
}
