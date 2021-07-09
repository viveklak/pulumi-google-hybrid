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
    public static readonly __pulumiType = 'google-native:accesscontextmanager/v1:ServicePerimeter';

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
     * Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, the restricted service list as well as access level lists must be empty.
     */
    public readonly perimeterType!: pulumi.Output<string>;
    /**
     * Proposed (or dry run) ServicePerimeter configuration. This configuration allows to specify and test ServicePerimeter configuration without enforcing actual access restrictions. Only allowed to be set when the "use_explicit_dry_run_spec" flag is set.
     */
    public readonly spec!: pulumi.Output<outputs.accesscontextmanager.v1.ServicePerimeterConfigResponse>;
    /**
     * Current ServicePerimeter configuration. Specifies sets of resources, restricted services and access levels that determine perimeter content and boundaries.
     */
    public readonly status!: pulumi.Output<outputs.accesscontextmanager.v1.ServicePerimeterConfigResponse>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists for all Service Perimeters, and that spec is identical to the status for those Service Perimeters. When this flag is set, it inhibits the generation of the implicit spec, thereby allowing the user to explicitly provide a configuration ("spec") to use in a dry-run version of the Service Perimeter. This allows the user to test changes to the enforced config ("status") without actually enforcing them. This testing is done through analyzing the differences between currently enforced and suggested restrictions. use_explicit_dry_run_spec must bet set to True if any of the fields in the spec are set to non-default values.
     */
    public readonly useExplicitDryRunSpec!: pulumi.Output<boolean>;

    /**
     * Create a ServicePerimeter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePerimeterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accessPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPolicyId'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["accessPolicyId"] = args ? args.accessPolicyId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["perimeterType"] = args ? args.perimeterType : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["useExplicitDryRunSpec"] = args ? args.useExplicitDryRunSpec : undefined;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["perimeterType"] = undefined /*out*/;
            inputs["spec"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["useExplicitDryRunSpec"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServicePerimeter.__pulumiType, name, inputs, opts);
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
    name: pulumi.Input<string>;
    /**
     * Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, the restricted service list as well as access level lists must be empty.
     */
    perimeterType?: pulumi.Input<enums.accesscontextmanager.v1.ServicePerimeterPerimeterType>;
    /**
     * Proposed (or dry run) ServicePerimeter configuration. This configuration allows to specify and test ServicePerimeter configuration without enforcing actual access restrictions. Only allowed to be set when the "use_explicit_dry_run_spec" flag is set.
     */
    spec?: pulumi.Input<inputs.accesscontextmanager.v1.ServicePerimeterConfigArgs>;
    /**
     * Current ServicePerimeter configuration. Specifies sets of resources, restricted services and access levels that determine perimeter content and boundaries.
     */
    status?: pulumi.Input<inputs.accesscontextmanager.v1.ServicePerimeterConfigArgs>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    title?: pulumi.Input<string>;
    /**
     * Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists for all Service Perimeters, and that spec is identical to the status for those Service Perimeters. When this flag is set, it inhibits the generation of the implicit spec, thereby allowing the user to explicitly provide a configuration ("spec") to use in a dry-run version of the Service Perimeter. This allows the user to test changes to the enforced config ("status") without actually enforcing them. This testing is done through analyzing the differences between currently enforced and suggested restrictions. use_explicit_dry_run_spec must bet set to True if any of the fields in the spec are set to non-default values.
     */
    useExplicitDryRunSpec?: pulumi.Input<boolean>;
}
