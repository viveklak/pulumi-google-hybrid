// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new custom Role.
 */
export class OrganizationRole extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationRole {
        return new OrganizationRole(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:iam/v1:OrganizationRole';

    /**
     * Returns true if the given object is an instance of OrganizationRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationRole.__pulumiType;
    }

    /**
     * The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
     */
    public readonly deleted!: pulumi.Output<boolean>;
    /**
     * Optional. A human-readable description for the role.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Used to perform a consistent read-modify-write.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The names of the permissions this role grants when bound in an IAM policy.
     */
    public readonly includedPermissions!: pulumi.Output<string[]>;
    /**
     * The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
     */
    public readonly stage!: pulumi.Output<string>;
    /**
     * Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a OrganizationRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationRoleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["deleted"] = args ? args.deleted : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["includedPermissions"] = args ? args.includedPermissions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
            resourceInputs["stage"] = args ? args.stage : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
        } else {
            resourceInputs["deleted"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["includedPermissions"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["stage"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationRole resource.
 */
export interface OrganizationRoleArgs {
    /**
     * The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
     */
    deleted?: pulumi.Input<boolean>;
    /**
     * Optional. A human-readable description for the role.
     */
    description?: pulumi.Input<string>;
    /**
     * Used to perform a consistent read-modify-write.
     */
    etag?: pulumi.Input<string>;
    /**
     * The names of the permissions this role grants when bound in an IAM policy.
     */
    includedPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * The role ID to use for this role. A role ID may contain alphanumeric characters, underscores (`_`), and periods (`.`). It must contain a minimum of 3 characters and a maximum of 64 characters.
     */
    roleId?: pulumi.Input<string>;
    /**
     * The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
     */
    stage?: pulumi.Input<enums.iam.v1.OrganizationRoleStage>;
    /**
     * Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
     */
    title?: pulumi.Input<string>;
}
