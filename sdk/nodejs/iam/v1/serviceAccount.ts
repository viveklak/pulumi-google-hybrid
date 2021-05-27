// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a ServiceAccount.
 */
export class ServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceAccount {
        return new ServiceAccount(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:iam/v1:ServiceAccount';

    /**
     * Returns true if the given object is an instance of ServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAccount.__pulumiType;
    }

    /**
     * Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether the service account is disabled.
     */
    public /*out*/ readonly disabled!: pulumi.Output<boolean>;
    /**
     * Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The email address of the service account.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to get the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The OAuth 2.0 client ID for the service account.
     */
    public /*out*/ readonly oauth2ClientId!: pulumi.Output<string>;
    /**
     * The ID of the project that owns the service account.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The unique, stable numeric ID for the service account. Each service account retains its unique ID even if you delete the service account. For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;

    /**
     * Create a ServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["disabled"] = undefined /*out*/;
            inputs["email"] = undefined /*out*/;
            inputs["oauth2ClientId"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["disabled"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["email"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["oauth2ClientId"] = undefined /*out*/;
            inputs["project"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceAccount resource.
 */
export interface ServiceAccountArgs {
    /**
     * Required. The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])` to comply with RFC1035.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * Optional. A user-specified, human-readable description of the service account. The maximum length is 256 UTF-8 bytes.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Optional. A user-specified, human-readable name for the service account. The maximum length is 100 UTF-8 bytes.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The resource name of the service account. Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes. For example, if you try to get the service account `projects/-/serviceAccounts/fake@example.com`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
}
