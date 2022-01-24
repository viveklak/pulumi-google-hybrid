// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a metastore service in a project and location.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:metastore/v1beta:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     */
    public /*out*/ readonly artifactGcsUri!: pulumi.Output<string>;
    /**
     * The time when the metastore service was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Immutable. The database type that the Metastore service stores its data.
     */
    public readonly databaseType!: pulumi.Output<string>;
    /**
     * Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
     */
    public readonly encryptionConfig!: pulumi.Output<outputs.metastore.v1beta.EncryptionConfigResponse>;
    /**
     * The URI of the endpoint used to access the metastore service.
     */
    public /*out*/ readonly endpointUri!: pulumi.Output<string>;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     */
    public readonly hiveMetastoreConfig!: pulumi.Output<outputs.metastore.v1beta.HiveMetastoreConfigResponse>;
    /**
     * User-defined labels for the metastore service.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.metastore.v1beta.MaintenanceWindowResponse>;
    /**
     * The setting that defines how metastore metadata should be integrated with external services and systems.
     */
    public readonly metadataIntegration!: pulumi.Output<outputs.metastore.v1beta.MetadataIntegrationResponse>;
    /**
     * The metadata management activities of the metastore service.
     */
    public /*out*/ readonly metadataManagementActivity!: pulumi.Output<outputs.metastore.v1beta.MetadataManagementActivityResponse>;
    /**
     * Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Immutable. The configuration specifying the network settings for the Dataproc Metastore service.
     */
    public readonly networkConfig!: pulumi.Output<outputs.metastore.v1beta.NetworkConfigResponse>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Immutable. The release channel of the service. If unspecified, defaults to STABLE.
     */
    public readonly releaseChannel!: pulumi.Output<string>;
    /**
     * The current state of the metastore service.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current state of the metastore service, if available.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * The tier of the service.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * The globally unique resource identifier of the metastore service.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the metastore service was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["databaseType"] = args ? args.databaseType : undefined;
            resourceInputs["encryptionConfig"] = args ? args.encryptionConfig : undefined;
            resourceInputs["hiveMetastoreConfig"] = args ? args.hiveMetastoreConfig : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["metadataIntegration"] = args ? args.metadataIntegration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["releaseChannel"] = args ? args.releaseChannel : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["artifactGcsUri"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpointUri"] = undefined /*out*/;
            resourceInputs["metadataManagementActivity"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["artifactGcsUri"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["databaseType"] = undefined /*out*/;
            resourceInputs["encryptionConfig"] = undefined /*out*/;
            resourceInputs["endpointUri"] = undefined /*out*/;
            resourceInputs["hiveMetastoreConfig"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["maintenanceWindow"] = undefined /*out*/;
            resourceInputs["metadataIntegration"] = undefined /*out*/;
            resourceInputs["metadataManagementActivity"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["networkConfig"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["releaseChannel"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["tier"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Immutable. The database type that the Metastore service stores its data.
     */
    databaseType?: pulumi.Input<enums.metastore.v1beta.ServiceDatabaseType>;
    /**
     * Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
     */
    encryptionConfig?: pulumi.Input<inputs.metastore.v1beta.EncryptionConfigArgs>;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     */
    hiveMetastoreConfig?: pulumi.Input<inputs.metastore.v1beta.HiveMetastoreConfigArgs>;
    /**
     * User-defined labels for the metastore service.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
     */
    maintenanceWindow?: pulumi.Input<inputs.metastore.v1beta.MaintenanceWindowArgs>;
    /**
     * The setting that defines how metastore metadata should be integrated with external services and systems.
     */
    metadataIntegration?: pulumi.Input<inputs.metastore.v1beta.MetadataIntegrationArgs>;
    /**
     * Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
     */
    network?: pulumi.Input<string>;
    /**
     * Immutable. The configuration specifying the network settings for the Dataproc Metastore service.
     */
    networkConfig?: pulumi.Input<inputs.metastore.v1beta.NetworkConfigArgs>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    port?: pulumi.Input<number>;
    project?: pulumi.Input<string>;
    /**
     * Immutable. The release channel of the service. If unspecified, defaults to STABLE.
     */
    releaseChannel?: pulumi.Input<enums.metastore.v1beta.ServiceReleaseChannel>;
    requestId?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
    /**
     * The tier of the service.
     */
    tier?: pulumi.Input<enums.metastore.v1beta.ServiceTier>;
}
