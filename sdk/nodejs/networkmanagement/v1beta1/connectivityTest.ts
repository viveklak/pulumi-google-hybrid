// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Connectivity Test. After you create a test, the reachability analysis is performed as part of the long running operation, which completes when the analysis completes. If the endpoint specifications in `ConnectivityTest` are invalid (for example, containing non-existent resources in the network, or you don't have read permissions to the network configurations of listed projects), then the reachability result returns a value of `UNKNOWN`. If the endpoint specifications in `ConnectivityTest` are incomplete, the reachability result returns a value of AMBIGUOUS. For more information, see the Connectivity Test documentation.
 */
export class ConnectivityTest extends pulumi.CustomResource {
    /**
     * Get an existing ConnectivityTest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectivityTest {
        return new ConnectivityTest(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networkmanagement/v1beta1:ConnectivityTest';

    /**
     * Returns true if the given object is an instance of ConnectivityTest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectivityTest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectivityTest.__pulumiType;
    }

    /**
     * The time the test was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The user-supplied description of the Connectivity Test. Maximum of 512 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
     */
    public readonly destination!: pulumi.Output<outputs.networkmanagement.v1beta1.EndpointResponse>;
    /**
     * The display name of a Connectivity Test.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Required. Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The probing details of this test from the latest run, present for applicable tests only. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
     */
    public /*out*/ readonly probingDetails!: pulumi.Output<outputs.networkmanagement.v1beta1.ProbingDetailsResponse>;
    /**
     * IP Protocol of the test. When not provided, "TCP" is assumed.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The reachability details of this test from the latest run. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
     */
    public /*out*/ readonly reachabilityDetails!: pulumi.Output<outputs.networkmanagement.v1beta1.ReachabilityDetailsResponse>;
    /**
     * Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
     */
    public readonly relatedProjects!: pulumi.Output<string[]>;
    /**
     * Required. Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
     */
    public readonly source!: pulumi.Output<outputs.networkmanagement.v1beta1.EndpointResponse>;
    /**
     * The time the test's configuration was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ConnectivityTest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectivityTestArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.testId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'testId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["relatedProjects"] = args ? args.relatedProjects : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["testId"] = args ? args.testId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["probingDetails"] = undefined /*out*/;
            inputs["reachabilityDetails"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["destination"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["probingDetails"] = undefined /*out*/;
            inputs["protocol"] = undefined /*out*/;
            inputs["reachabilityDetails"] = undefined /*out*/;
            inputs["relatedProjects"] = undefined /*out*/;
            inputs["source"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConnectivityTest.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectivityTest resource.
 */
export interface ConnectivityTestArgs {
    /**
     * The user-supplied description of the Connectivity Test. Maximum of 512 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
     */
    destination?: pulumi.Input<inputs.networkmanagement.v1beta1.EndpointArgs>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: string}>;
    /**
     * Required. Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * IP Protocol of the test. When not provided, "TCP" is assumed.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
     */
    relatedProjects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
     */
    source?: pulumi.Input<inputs.networkmanagement.v1beta1.EndpointArgs>;
    testId: pulumi.Input<string>;
}
