// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Microsoft AD domain.
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
    public static readonly __pulumiType = 'google-native:managedidentities/v1beta1:Domain';

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
     * Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
     */
    public readonly admin!: pulumi.Output<string>;
    /**
     * Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
     */
    public readonly auditLogsEnabled!: pulumi.Output<boolean>;
    /**
     * Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
     */
    public readonly authorizedNetworks!: pulumi.Output<string[]>;
    /**
     * The time the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory set up on an internal network.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * Optional. Resource labels that can contain user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     */
    public readonly locations!: pulumi.Output<string[]>;
    /**
     * The unique name of the domain using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
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
    public /*out*/ readonly trusts!: pulumi.Output<outputs.managedidentities.v1beta1.TrustResponse[]>;
    /**
     * The last update time.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["admin"] = args ? args.admin : undefined;
            inputs["auditLogsEnabled"] = args ? args.auditLogsEnabled : undefined;
            inputs["authorizedNetworks"] = args ? args.authorizedNetworks : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["locations"] = args ? args.locations : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["reservedIpRange"] = args ? args.reservedIpRange : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
            inputs["trusts"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["admin"] = undefined /*out*/;
            inputs["auditLogsEnabled"] = undefined /*out*/;
            inputs["authorizedNetworks"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["locations"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["reservedIpRange"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
            inputs["trusts"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Domain.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
     */
    admin?: pulumi.Input<string>;
    /**
     * Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
     */
    auditLogsEnabled?: pulumi.Input<boolean>;
    /**
     * Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
     */
    authorizedNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    domainName: pulumi.Input<string>;
    /**
     * Optional. Resource labels that can contain user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: string}>;
    /**
     * Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     */
    locations?: pulumi.Input<pulumi.Input<string>[]>;
    project: pulumi.Input<string>;
    /**
     * Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
     */
    reservedIpRange?: pulumi.Input<string>;
}
