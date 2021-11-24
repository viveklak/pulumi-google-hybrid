// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a Service Perimeter. The longrunning operation from this RPC will have a successful status once the Service Perimeter has propagated to long-lasting storage. Service Perimeters containing errors will result in an error response for the first error encountered.
 */
export class ServicePerimeter extends pulumi.CustomResource {
    /**
     * Get an existing ServicePerimeter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServicePerimeter {
        return new ServicePerimeter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:accesscontextmanager/v1beta:ServicePerimeter';

    /**
     * Returns true if the given object is an instance of ServicePerimeter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePerimeter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePerimeter.__pulumiType;
    }

    /**
     * Description of the `ServicePerimeter` and its use. Does not affect behavior.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
     */
    public readonly perimeterType!: pulumi.Output<string>;
    /**
     * Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
     */
    public readonly status!: pulumi.Output<outputs.accesscontextmanager.v1beta.ServicePerimeterConfigResponse>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a ServicePerimeter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePerimeterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accessPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPolicyId'");
            }
            resourceInputs["accessPolicyId"] = args ? args.accessPolicyId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["perimeterType"] = args ? args.perimeterType : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["perimeterType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServicePerimeter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServicePerimeter resource.
 */
export interface ServicePerimeterArgs {
    accessPolicyId: pulumi.Input<string>;
    /**
     * Description of the `ServicePerimeter` and its use. Does not affect behavior.
     */
    description?: pulumi.Input<string>;
    /**
     * Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
     */
    name?: pulumi.Input<string>;
    /**
     * Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
     */
    perimeterType?: pulumi.Input<enums.accesscontextmanager.v1beta.ServicePerimeterPerimeterType>;
    /**
     * Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
     */
    status?: pulumi.Input<inputs.accesscontextmanager.v1beta.ServicePerimeterConfigArgs>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    title?: pulumi.Input<string>;
}
