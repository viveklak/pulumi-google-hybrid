// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a saved query in a parent project/folder/organization.
 * Auto-naming is currently not supported for this resource.
 */
export class SavedQuery extends pulumi.CustomResource {
    /**
     * Get an existing SavedQuery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SavedQuery {
        return new SavedQuery(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudasset/v1:SavedQuery';

    /**
     * Returns true if the given object is an instance of SavedQuery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SavedQuery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SavedQuery.__pulumiType;
    }

    /**
     * The query content.
     */
    public readonly content!: pulumi.Output<outputs.cloudasset.v1.QueryContentResponse>;
    /**
     * The create time of this saved query.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The account's email address who has created this saved query.
     */
    public /*out*/ readonly creator!: pulumi.Output<string>;
    /**
     * The description of this saved query. This value should be fewer than 255 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The last update time of this saved query.
     */
    public /*out*/ readonly lastUpdateTime!: pulumi.Output<string>;
    /**
     * The account's email address who has updated this saved query most recently.
     */
    public /*out*/ readonly lastUpdater!: pulumi.Output<string>;
    /**
     * The resource name of the saved query. The format must be: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a SavedQuery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SavedQueryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.savedQueryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'savedQueryId'");
            }
            if ((!args || args.v1Id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'v1Id'");
            }
            if ((!args || args.v1Id1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'v1Id1'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["savedQueryId"] = args ? args.savedQueryId : undefined;
            resourceInputs["v1Id"] = args ? args.v1Id : undefined;
            resourceInputs["v1Id1"] = args ? args.v1Id1 : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["creator"] = undefined /*out*/;
            resourceInputs["lastUpdateTime"] = undefined /*out*/;
            resourceInputs["lastUpdater"] = undefined /*out*/;
        } else {
            resourceInputs["content"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["creator"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["lastUpdateTime"] = undefined /*out*/;
            resourceInputs["lastUpdater"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SavedQuery.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SavedQuery resource.
 */
export interface SavedQueryArgs {
    /**
     * The query content.
     */
    content?: pulumi.Input<inputs.cloudasset.v1.QueryContentArgs>;
    /**
     * The description of this saved query. This value should be fewer than 255 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name of the saved query. The format must be: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id
     */
    name?: pulumi.Input<string>;
    savedQueryId: pulumi.Input<string>;
    v1Id: pulumi.Input<string>;
    v1Id1: pulumi.Input<string>;
}
