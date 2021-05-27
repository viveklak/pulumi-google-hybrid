// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a regional HealthCheckService resource in the specified project and region using the data included in the request.
 */
export class RegionHealthCheckService extends pulumi.CustomResource {
    /**
     * Get an existing RegionHealthCheckService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionHealthCheckService {
        return new RegionHealthCheckService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:RegionHealthCheckService';

    /**
     * Returns true if the given object is an instance of RegionHealthCheckService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionHealthCheckService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionHealthCheckService.__pulumiType;
    }

    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a HealthCheckService. An up-to-date fingerprint must be provided in order to patch/update the HealthCheckService; Otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the HealthCheckService.
     */
    public readonly fingerprint!: pulumi.Output<string>;
    /**
     * List of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks</code? must belong to the same region as zones of NEGs.
     */
    public readonly healthChecks!: pulumi.Output<string[]>;
    /**
     * Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified.  
     * - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service. 
     * - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
     */
    public readonly healthStatusAggregationPolicy!: pulumi.Output<string>;
    /**
     * This field is deprecated. Use health_status_aggregation_policy instead.
     *
     * Policy for how the results from multiple health checks for the same endpoint are aggregated.  
     * - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service. 
     * - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
     */
    public readonly healthStatusAggregationStrategy!: pulumi.Output<string>;
    /**
     * [Output only] Type of the resource. Always compute#healthCheckServicefor health check services.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
     */
    public readonly networkEndpointGroups!: pulumi.Output<string[]>;
    /**
     * List of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
     */
    public readonly notificationEndpoints!: pulumi.Output<string[]>;
    /**
     * [Output Only] URL of the region where the health check service resides. This field is not applicable to global health check services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a RegionHealthCheckService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionHealthCheckServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["fingerprint"] = args ? args.fingerprint : undefined;
            inputs["healthChecks"] = args ? args.healthChecks : undefined;
            inputs["healthStatusAggregationPolicy"] = args ? args.healthStatusAggregationPolicy : undefined;
            inputs["healthStatusAggregationStrategy"] = args ? args.healthStatusAggregationStrategy : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkEndpointGroups"] = args ? args.networkEndpointGroups : undefined;
            inputs["notificationEndpoints"] = args ? args.notificationEndpoints : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["healthChecks"] = undefined /*out*/;
            inputs["healthStatusAggregationPolicy"] = undefined /*out*/;
            inputs["healthStatusAggregationStrategy"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkEndpointGroups"] = undefined /*out*/;
            inputs["notificationEndpoints"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionHealthCheckService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionHealthCheckService resource.
 */
export interface RegionHealthCheckServiceArgs {
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a HealthCheckService. An up-to-date fingerprint must be provided in order to patch/update the HealthCheckService; Otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the HealthCheckService.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * List of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks</code? must belong to the same region as zones of NEGs.
     */
    readonly healthChecks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified.  
     * - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service. 
     * - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
     */
    readonly healthStatusAggregationPolicy?: pulumi.Input<string>;
    /**
     * This field is deprecated. Use health_status_aggregation_policy instead.
     *
     * Policy for how the results from multiple health checks for the same endpoint are aggregated.  
     * - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service. 
     * - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
     */
    readonly healthStatusAggregationStrategy?: pulumi.Input<string>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output only] Type of the resource. Always compute#healthCheckServicefor health check services.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
     */
    readonly networkEndpointGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
     */
    readonly notificationEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] URL of the region where the health check service resides. This field is not applicable to global health check services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}
