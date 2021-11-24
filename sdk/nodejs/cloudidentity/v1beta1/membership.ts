// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a `Membership`.
 * Auto-naming is currently not supported for this resource.
 */
export class Membership extends pulumi.CustomResource {
    /**
     * Get an existing Membership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Membership {
        return new Membership(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudidentity/v1beta1:Membership';

    /**
     * Returns true if the given object is an instance of Membership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Membership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Membership.__pulumiType;
    }

    /**
     * The time when the `Membership` was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
     */
    public readonly memberKey!: pulumi.Output<outputs.cloudidentity.v1beta1.EntityKeyResponse>;
    /**
     * The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Membership`. Shall be of the form `groups/{group_id}/memberships/{membership_id}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
     */
    public readonly preferredMemberKey!: pulumi.Output<outputs.cloudidentity.v1beta1.EntityKeyResponse>;
    /**
     * The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
     */
    public readonly roles!: pulumi.Output<outputs.cloudidentity.v1beta1.MembershipRoleResponse[]>;
    /**
     * The type of the membership.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The time when the `Membership` was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Membership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MembershipArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.preferredMemberKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'preferredMemberKey'");
            }
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["memberKey"] = args ? args.memberKey : undefined;
            resourceInputs["preferredMemberKey"] = args ? args.preferredMemberKey : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["memberKey"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["preferredMemberKey"] = undefined /*out*/;
            resourceInputs["roles"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Membership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Membership resource.
 */
export interface MembershipArgs {
    groupId: pulumi.Input<string>;
    /**
     * Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
     */
    memberKey?: pulumi.Input<inputs.cloudidentity.v1beta1.EntityKeyArgs>;
    /**
     * Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
     */
    preferredMemberKey: pulumi.Input<inputs.cloudidentity.v1beta1.EntityKeyArgs>;
    /**
     * The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.cloudidentity.v1beta1.MembershipRoleArgs>[]>;
}
