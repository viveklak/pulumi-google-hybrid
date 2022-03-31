// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a BackendBucket resource in the specified project using the data included in the request.
 */
export class BackendBucket extends pulumi.CustomResource {
    /**
     * Get an existing BackendBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BackendBucket {
        return new BackendBucket(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:BackendBucket';

    /**
     * Returns true if the given object is an instance of BackendBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendBucket.__pulumiType;
    }

    /**
     * Cloud Storage bucket name.
     */
    public readonly bucketName!: pulumi.Output<string>;
    /**
     * Cloud CDN configuration for this BackendBucket.
     */
    public readonly cdnPolicy!: pulumi.Output<outputs.compute.alpha.BackendBucketCdnPolicyResponse>;
    /**
     * Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
     */
    public readonly compressionMode!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied responses.
     */
    public readonly customResponseHeaders!: pulumi.Output<string[]>;
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The resource URL for the edge security policy associated with this backend bucket.
     */
    public /*out*/ readonly edgeSecurityPolicy!: pulumi.Output<string>;
    /**
     * If true, enable Cloud CDN for this BackendBucket.
     */
    public readonly enableCdn!: pulumi.Output<boolean>;
    /**
     * Type of the resource.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;

    /**
     * Create a BackendBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BackendBucketArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            resourceInputs["compressionMode"] = args ? args.compressionMode : undefined;
            resourceInputs["customResponseHeaders"] = args ? args.customResponseHeaders : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableCdn"] = args ? args.enableCdn : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["edgeSecurityPolicy"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
        } else {
            resourceInputs["bucketName"] = undefined /*out*/;
            resourceInputs["cdnPolicy"] = undefined /*out*/;
            resourceInputs["compressionMode"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["customResponseHeaders"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["edgeSecurityPolicy"] = undefined /*out*/;
            resourceInputs["enableCdn"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackendBucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BackendBucket resource.
 */
export interface BackendBucketArgs {
    /**
     * Cloud Storage bucket name.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * Cloud CDN configuration for this BackendBucket.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.alpha.BackendBucketCdnPolicyArgs>;
    /**
     * Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
     */
    compressionMode?: pulumi.Input<enums.compute.alpha.BackendBucketCompressionMode>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied responses.
     */
    customResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    description?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this BackendBucket.
     */
    enableCdn?: pulumi.Input<boolean>;
    /**
     * Type of the resource.
     */
    kind?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
}
