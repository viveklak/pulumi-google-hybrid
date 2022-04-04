// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-hybrid:organizations/classic:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Create the 'default' network automatically. Default true. If set to false, the default network will be deleted. Note
     * that, for quota purposes, you will still need to have 1 network slot available to create the project successfully, even
     * if you set auto_create_network to false, since the network will exist momentarily.
     */
    public readonly autoCreateNetwork!: pulumi.Output<boolean | undefined>;
    /**
     * The alphanumeric ID of the billing account this project belongs to. The user or service account performing this
     * operation with Terraform must have Billing Account Administrator privileges (roles/billing.admin) in the organization.
     * See Google Cloud Billing API Access Control for more details.
     */
    public readonly billingAccount!: pulumi.Output<string | undefined>;
    /**
     * The numeric ID of the folder this project should be created under. Only one of org_id or folder_id may be specified. If
     * the folder_id is specified, then the project is created under the specified folder. Changing this forces the project to
     * be migrated to the newly specified folder.
     */
    public readonly folderId!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value label pairs to assign to the project.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The display name of the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The numeric identifier of the project.
     */
    public /*out*/ readonly number!: pulumi.Output<string>;
    /**
     * The numeric ID of the organization this project belongs to. Changing this forces a new project to be created. Only one
     * of org_id or folder_id may be specified. If the org_id is specified then the project is created at the top level.
     * Changing this forces the project to be migrated to the newly specified organization.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * The project ID. Changing this forces a new project to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * If true, the Terraform resource can be deleted without deleting the Project via the Google API.
     */
    public readonly skipDelete!: pulumi.Output<boolean>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["autoCreateNetwork"] = state ? state.autoCreateNetwork : undefined;
            resourceInputs["billingAccount"] = state ? state.billingAccount : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["number"] = state ? state.number : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["skipDelete"] = state ? state.skipDelete : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["autoCreateNetwork"] = args ? args.autoCreateNetwork : undefined;
            resourceInputs["billingAccount"] = args ? args.billingAccount : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["skipDelete"] = args ? args.skipDelete : undefined;
            resourceInputs["number"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Create the 'default' network automatically. Default true. If set to false, the default network will be deleted. Note
     * that, for quota purposes, you will still need to have 1 network slot available to create the project successfully, even
     * if you set auto_create_network to false, since the network will exist momentarily.
     */
    autoCreateNetwork?: pulumi.Input<boolean>;
    /**
     * The alphanumeric ID of the billing account this project belongs to. The user or service account performing this
     * operation with Terraform must have Billing Account Administrator privileges (roles/billing.admin) in the organization.
     * See Google Cloud Billing API Access Control for more details.
     */
    billingAccount?: pulumi.Input<string>;
    /**
     * The numeric ID of the folder this project should be created under. Only one of org_id or folder_id may be specified. If
     * the folder_id is specified, then the project is created under the specified folder. Changing this forces the project to
     * be migrated to the newly specified folder.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the project.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The display name of the project.
     */
    name?: pulumi.Input<string>;
    /**
     * The numeric identifier of the project.
     */
    number?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization this project belongs to. Changing this forces a new project to be created. Only one
     * of org_id or folder_id may be specified. If the org_id is specified then the project is created at the top level.
     * Changing this forces the project to be migrated to the newly specified organization.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The project ID. Changing this forces a new project to be created.
     */
    projectId?: pulumi.Input<string>;
    /**
     * If true, the Terraform resource can be deleted without deleting the Project via the Google API.
     */
    skipDelete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Create the 'default' network automatically. Default true. If set to false, the default network will be deleted. Note
     * that, for quota purposes, you will still need to have 1 network slot available to create the project successfully, even
     * if you set auto_create_network to false, since the network will exist momentarily.
     */
    autoCreateNetwork?: pulumi.Input<boolean>;
    /**
     * The alphanumeric ID of the billing account this project belongs to. The user or service account performing this
     * operation with Terraform must have Billing Account Administrator privileges (roles/billing.admin) in the organization.
     * See Google Cloud Billing API Access Control for more details.
     */
    billingAccount?: pulumi.Input<string>;
    /**
     * The numeric ID of the folder this project should be created under. Only one of org_id or folder_id may be specified. If
     * the folder_id is specified, then the project is created under the specified folder. Changing this forces the project to
     * be migrated to the newly specified folder.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the project.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The display name of the project.
     */
    name?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization this project belongs to. Changing this forces a new project to be created. Only one
     * of org_id or folder_id may be specified. If the org_id is specified then the project is created at the top level.
     * Changing this forces the project to be migrated to the newly specified organization.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The project ID. Changing this forces a new project to be created.
     */
    projectId: pulumi.Input<string>;
    /**
     * If true, the Terraform resource can be deleted without deleting the Project via the Google API.
     */
    skipDelete?: pulumi.Input<boolean>;
}