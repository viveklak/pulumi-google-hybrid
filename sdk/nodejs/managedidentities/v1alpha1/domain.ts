// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Microsoft AD Domain in a given project. Operation
 * Auto-naming is currently not supported for this resource.
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:managedidentities/v1alpha1:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
     */
    public readonly auditLogsEnabled!: pulumi.Output<boolean>;
    /**
     * Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
     */
    public readonly authorizedNetworks!: pulumi.Output<string[]>;
    /**
     * The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * Optional. Resource labels to represent user provided metadata
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     */
    public readonly locations!: pulumi.Output<string[]>;
    /**
     * Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
     */
    public readonly managedIdentitiesAdminName!: pulumi.Output<string>;
    /**
     * Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
     */
    public readonly reservedIpRange!: pulumi.Output<string>;
    /**
     * The current state of this domain.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current status of this domain, if available.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * The current trusts associated with the domain.
     */
    public /*out*/ readonly trusts!: pulumi.Output<outputs.managedidentities.v1alpha1.TrustResponse[]>;
    /**
     * Last update time. Synthetic field is populated automatically by CCFE.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.locations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locations'");
            }
            if ((!args || args.reservedIpRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reservedIpRange'");
            }
            resourceInputs["auditLogsEnabled"] = args ? args.auditLogsEnabled : undefined;
            resourceInputs["authorizedNetworks"] = args ? args.authorizedNetworks : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["managedIdentitiesAdminName"] = args ? args.managedIdentitiesAdminName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["reservedIpRange"] = args ? args.reservedIpRange : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["fqdn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["trusts"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["auditLogsEnabled"] = undefined /*out*/;
            resourceInputs["authorizedNetworks"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["fqdn"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["locations"] = undefined /*out*/;
            resourceInputs["managedIdentitiesAdminName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["reservedIpRange"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["trusts"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
     */
    auditLogsEnabled?: pulumi.Input<boolean>;
    /**
     * Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
     */
    authorizedNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    domainName?: pulumi.Input<string>;
    /**
     * Optional. Resource labels to represent user provided metadata
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     */
    locations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
     */
    managedIdentitiesAdminName?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
     */
    reservedIpRange: pulumi.Input<string>;
}
