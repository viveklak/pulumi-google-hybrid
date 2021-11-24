// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new EndpointPolicy in a given project and location.
 */
export class EndpointPolicy extends pulumi.CustomResource {
    /**
     * Get an existing EndpointPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EndpointPolicy {
        return new EndpointPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networkservices/v1:EndpointPolicy';

    /**
     * Returns true if the given object is an instance of EndpointPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointPolicy.__pulumiType;
    }

    /**
     * Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
     */
    public readonly authorizationPolicy!: pulumi.Output<string>;
    /**
     * Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
     */
    public readonly clientTlsPolicy!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * A matcher that selects endpoints to which the policies should be applied.
     */
    public readonly endpointMatcher!: pulumi.Output<outputs.networkservices.v1.EndpointMatcherResponse>;
    /**
     * Optional. Set of label tags associated with the EndpointPolicy resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
     */
    public readonly serverTlsPolicy!: pulumi.Output<string>;
    /**
     * Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
     */
    public readonly trafficPortSelector!: pulumi.Output<outputs.networkservices.v1.TrafficPortSelectorResponse>;
    /**
     * The type of endpoint policy. This is primarily used to validate the configuration.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a EndpointPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.endpointMatcher === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointMatcher'");
            }
            if ((!args || args.endpointPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointPolicyId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["authorizationPolicy"] = args ? args.authorizationPolicy : undefined;
            resourceInputs["clientTlsPolicy"] = args ? args.clientTlsPolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointMatcher"] = args ? args.endpointMatcher : undefined;
            resourceInputs["endpointPolicyId"] = args ? args.endpointPolicyId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serverTlsPolicy"] = args ? args.serverTlsPolicy : undefined;
            resourceInputs["trafficPortSelector"] = args ? args.trafficPortSelector : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["authorizationPolicy"] = undefined /*out*/;
            resourceInputs["clientTlsPolicy"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["endpointMatcher"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serverTlsPolicy"] = undefined /*out*/;
            resourceInputs["trafficPortSelector"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EndpointPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EndpointPolicy resource.
 */
export interface EndpointPolicyArgs {
    /**
     * Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
     */
    authorizationPolicy?: pulumi.Input<string>;
    /**
     * Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
     */
    clientTlsPolicy?: pulumi.Input<string>;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * A matcher that selects endpoints to which the policies should be applied.
     */
    endpointMatcher: pulumi.Input<inputs.networkservices.v1.EndpointMatcherArgs>;
    endpointPolicyId: pulumi.Input<string>;
    /**
     * Optional. Set of label tags associated with the EndpointPolicy resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
     */
    serverTlsPolicy?: pulumi.Input<string>;
    /**
     * Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
     */
    trafficPortSelector?: pulumi.Input<inputs.networkservices.v1.TrafficPortSelectorArgs>;
    /**
     * The type of endpoint policy. This is primarily used to validate the configuration.
     */
    type: pulumi.Input<enums.networkservices.v1.EndpointPolicyType>;
}
