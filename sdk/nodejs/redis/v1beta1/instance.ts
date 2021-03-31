// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Redis instance based on the specified tier and memory size. By default, the instance is accessible from the project's [default network](https://cloud.google.com/vpc/docs/vpc). The creation is executed asynchronously and callers may check the returned operation to track its progress. Once the operation is completed the Redis instance will be fully functional. Completed longrunning.Operation will contain the new instance object in the response field. The returned operation is automatically deleted after a few hours, so there is no need to call DeleteOperation.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:redis/v1beta1:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }


    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instancesId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instancesId'");
            }
            if ((!args || args.locationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["alternativeLocationId"] = args ? args.alternativeLocationId : undefined;
            inputs["authEnabled"] = args ? args.authEnabled : undefined;
            inputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            inputs["connectMode"] = args ? args.connectMode : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["currentLocationId"] = args ? args.currentLocationId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["instancesId"] = args ? args.instancesId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["locationId"] = args ? args.locationId : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["memorySizeGb"] = args ? args.memorySizeGb : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["persistenceIamIdentity"] = args ? args.persistenceIamIdentity : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["redisConfigs"] = args ? args.redisConfigs : undefined;
            inputs["redisVersion"] = args ? args.redisVersion : undefined;
            inputs["reservedIpRange"] = args ? args.reservedIpRange : undefined;
            inputs["serverCaCerts"] = args ? args.serverCaCerts : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["statusMessage"] = args ? args.statusMessage : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["transitEncryptionMode"] = args ? args.transitEncryptionMode : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
     */
    readonly alternativeLocationId?: pulumi.Input<string>;
    /**
     * Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
     */
    readonly authEnabled?: pulumi.Input<boolean>;
    /**
     * Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
     */
    readonly authorizedNetwork?: pulumi.Input<string>;
    /**
     * Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
     */
    readonly connectMode?: pulumi.Input<string>;
    /**
     * Output only. The time the instance was created.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Output only. The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the location_id provided by the user at creation time. For Standard Tier instances, this can be either location_id or alternative_location_id and can change after a failover event.
     */
    readonly currentLocationId?: pulumi.Input<string>;
    /**
     * An arbitrary and optional user-provided name for the instance.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Output only. Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
     */
    readonly host?: pulumi.Input<string>;
    readonly instancesId: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
     */
    readonly locationId?: pulumi.Input<string>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * Required. Redis memory size in GiB.
     */
    readonly memorySizeGb?: pulumi.Input<number>;
    /**
     * Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Output only. Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
     */
    readonly persistenceIamIdentity?: pulumi.Input<string>;
    /**
     * Output only. The port number of the exposed Redis endpoint.
     */
    readonly port?: pulumi.Input<number>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
     */
    readonly redisConfigs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility
     */
    readonly redisVersion?: pulumi.Input<string>;
    /**
     * Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
     */
    readonly reservedIpRange?: pulumi.Input<string>;
    /**
     * Output only. List of server CA certificates for the instance.
     */
    readonly serverCaCerts?: pulumi.Input<pulumi.Input<inputs.redis.v1beta1.TlsCertificate>[]>;
    /**
     * Output only. The current state of this instance.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Output only. Additional information about the current status of this instance, if available.
     */
    readonly statusMessage?: pulumi.Input<string>;
    /**
     * Required. The service tier of the instance.
     */
    readonly tier?: pulumi.Input<string>;
    /**
     * Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
     */
    readonly transitEncryptionMode?: pulumi.Input<string>;
}
