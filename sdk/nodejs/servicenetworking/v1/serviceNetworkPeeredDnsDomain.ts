// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a peered DNS domain which sends requests for records in given namespace originating in the service producer VPC network to the consumer VPC network to be resolved.
 */
export class ServiceNetworkPeeredDnsDomain extends pulumi.CustomResource {
    /**
     * Get an existing ServiceNetworkPeeredDnsDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceNetworkPeeredDnsDomain {
        return new ServiceNetworkPeeredDnsDomain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:servicenetworking/v1:ServiceNetworkPeeredDnsDomain';

    /**
     * Returns true if the given object is an instance of ServiceNetworkPeeredDnsDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceNetworkPeeredDnsDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceNetworkPeeredDnsDomain.__pulumiType;
    }


    /**
     * Create a ServiceNetworkPeeredDnsDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceNetworkPeeredDnsDomainArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.networksId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networksId'");
            }
            if ((!args || args.peeredDnsDomainsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peeredDnsDomainsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            if ((!args || args.servicesId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicesId'");
            }
            inputs["dnsSuffix"] = args ? args.dnsSuffix : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networksId"] = args ? args.networksId : undefined;
            inputs["peeredDnsDomainsId"] = args ? args.peeredDnsDomainsId : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["servicesId"] = args ? args.servicesId : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceNetworkPeeredDnsDomain.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceNetworkPeeredDnsDomain resource.
 */
export interface ServiceNetworkPeeredDnsDomainArgs {
    /**
     * The DNS domain name suffix e.g. `example.com.`.
     */
    readonly dnsSuffix?: pulumi.Input<string>;
    /**
     * User assigned name for this resource. Must be unique within the consumer network. The name must be 1-63 characters long, must begin with a letter, end with a letter or digit, and only contain lowercase letters, digits or dashes.
     */
    readonly name?: pulumi.Input<string>;
    readonly networksId: pulumi.Input<string>;
    readonly peeredDnsDomainsId: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
    readonly servicesId: pulumi.Input<string>;
}