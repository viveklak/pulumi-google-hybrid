// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a knowledge base.
 */
export class KnowledgeBase extends pulumi.CustomResource {
    /**
     * Get an existing KnowledgeBase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KnowledgeBase {
        return new KnowledgeBase(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2:KnowledgeBase';

    /**
     * Returns true if the given object is an instance of KnowledgeBase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KnowledgeBase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KnowledgeBase.__pulumiType;
    }

    /**
     * The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, expect this to be present for non en-us languages. When unspecified, the default language code en-us applies.
     */
    public readonly languageCode!: pulumi.Output<string>;
    /**
     * The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a KnowledgeBase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KnowledgeBaseArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["displayName"] = undefined /*out*/;
            inputs["languageCode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KnowledgeBase.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a KnowledgeBase resource.
 */
export interface KnowledgeBaseArgs {
    /**
     * The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
     */
    displayName: pulumi.Input<string>;
    /**
     * Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, expect this to be present for non en-us languages. When unspecified, the default language code en-us applies.
     */
    languageCode?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    /**
     * The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
}
