// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a phrase matcher.
 */
export class PhraseMatcher extends pulumi.CustomResource {
    /**
     * Get an existing PhraseMatcher resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PhraseMatcher {
        return new PhraseMatcher(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:contactcenterinsights/v1:PhraseMatcher';

    /**
     * Returns true if the given object is an instance of PhraseMatcher.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PhraseMatcher {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PhraseMatcher.__pulumiType;
    }

    /**
     * The most recent time at which the activation status was updated.
     */
    public /*out*/ readonly activationUpdateTime!: pulumi.Output<string>;
    /**
     * Applies the phrase matcher only when it is active.
     */
    public readonly active!: pulumi.Output<boolean>;
    /**
     * The human-readable name of the phrase matcher.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The resource name of the phrase matcher. Format: projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of phase match rule groups that are included in this matcher.
     */
    public readonly phraseMatchRuleGroups!: pulumi.Output<outputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupResponse[]>;
    /**
     * The timestamp of when the revision was created. It is also the create time when a new matcher is added.
     */
    public /*out*/ readonly revisionCreateTime!: pulumi.Output<string>;
    /**
     * Immutable. The revision ID of the phrase matcher. A new revision is committed whenever the matcher is changed, except when it is activated or deactivated. A server generated random ID will be used. Example: locations/global/phraseMatchers/my-first-matcher@1234567
     */
    public /*out*/ readonly revisionId!: pulumi.Output<string>;
    /**
     * The role whose utterances the phrase matcher should be matched against. If the role is ROLE_UNSPECIFIED it will be matched against any utterances in the transcript.
     */
    public readonly roleMatch!: pulumi.Output<string>;
    /**
     * The type of this phrase matcher.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The most recent time at which the phrase matcher was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The customized version tag to use for the phrase matcher. If not specified, it will default to `revision_id`.
     */
    public readonly versionTag!: pulumi.Output<string>;

    /**
     * Create a PhraseMatcher resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PhraseMatcherArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["phraseMatchRuleGroups"] = args ? args.phraseMatchRuleGroups : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["roleMatch"] = args ? args.roleMatch : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["versionTag"] = args ? args.versionTag : undefined;
            resourceInputs["activationUpdateTime"] = undefined /*out*/;
            resourceInputs["revisionCreateTime"] = undefined /*out*/;
            resourceInputs["revisionId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["activationUpdateTime"] = undefined /*out*/;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["phraseMatchRuleGroups"] = undefined /*out*/;
            resourceInputs["revisionCreateTime"] = undefined /*out*/;
            resourceInputs["revisionId"] = undefined /*out*/;
            resourceInputs["roleMatch"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["versionTag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PhraseMatcher.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PhraseMatcher resource.
 */
export interface PhraseMatcherArgs {
    /**
     * Applies the phrase matcher only when it is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The human-readable name of the phrase matcher.
     */
    displayName?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the phrase matcher. Format: projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
     */
    name?: pulumi.Input<string>;
    /**
     * A list of phase match rule groups that are included in this matcher.
     */
    phraseMatchRuleGroups?: pulumi.Input<pulumi.Input<inputs.contactcenterinsights.v1.GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupArgs>[]>;
    project?: pulumi.Input<string>;
    /**
     * The role whose utterances the phrase matcher should be matched against. If the role is ROLE_UNSPECIFIED it will be matched against any utterances in the transcript.
     */
    roleMatch?: pulumi.Input<enums.contactcenterinsights.v1.PhraseMatcherRoleMatch>;
    /**
     * The type of this phrase matcher.
     */
    type: pulumi.Input<enums.contactcenterinsights.v1.PhraseMatcherType>;
    /**
     * The customized version tag to use for the phrase matcher. If not specified, it will default to `revision_id`.
     */
    versionTag?: pulumi.Input<string>;
}
