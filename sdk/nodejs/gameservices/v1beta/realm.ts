// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new realm in a given project and location.
 */
export class Realm extends pulumi.CustomResource {
    /**
     * Get an existing Realm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Realm {
        return new Realm(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gameservices/v1beta:Realm';

    /**
     * Returns true if the given object is an instance of Realm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Realm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Realm.__pulumiType;
    }

    /**
     * The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Human readable description of the realm.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * ETag of the resource.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The labels associated with this realm. Each label is a key-value pair.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name of the realm, in the following form: `projects/{project}/locations/{location}/realms/{realm}`. For example, `projects/my-project/locations/{location}/realms/my-realm`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Time zone where all policies targeting this realm are evaluated. The value of this field must be from the IANA time zone database: https://www.iana.org/time-zones.
     */
    public readonly timeZone!: pulumi.Output<string>;
    /**
     * The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Realm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealmArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.timeZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeZone'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["timeZone"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Realm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Realm resource.
 */
export interface RealmArgs {
    /**
     * Human readable description of the realm.
     */
    description?: pulumi.Input<string>;
    /**
     * ETag of the resource.
     */
    etag?: pulumi.Input<string>;
    /**
     * The labels associated with this realm. Each label is a key-value pair.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the realm, in the following form: `projects/{project}/locations/{location}/realms/{realm}`. For example, `projects/my-project/locations/{location}/realms/my-realm`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
    /**
     * Time zone where all policies targeting this realm are evaluated. The value of this field must be from the IANA time zone database: https://www.iana.org/time-zones.
     */
    timeZone: pulumi.Input<string>;
}
