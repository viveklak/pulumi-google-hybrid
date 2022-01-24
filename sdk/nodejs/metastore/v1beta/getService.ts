// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the details of a single service.
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:metastore/v1beta:getService", {
        "location": args.location,
        "project": args.project,
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetServiceArgs {
    location: string;
    project?: string;
    serviceId: string;
}

export interface GetServiceResult {
    /**
     * A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     */
    readonly artifactGcsUri: string;
    /**
     * The time when the metastore service was created.
     */
    readonly createTime: string;
    /**
     * Immutable. The database type that the Metastore service stores its data.
     */
    readonly databaseType: string;
    /**
     * Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
     */
    readonly encryptionConfig: outputs.metastore.v1beta.EncryptionConfigResponse;
    /**
     * The URI of the endpoint used to access the metastore service.
     */
    readonly endpointUri: string;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     */
    readonly hiveMetastoreConfig: outputs.metastore.v1beta.HiveMetastoreConfigResponse;
    /**
     * User-defined labels for the metastore service.
     */
    readonly labels: {[key: string]: string};
    /**
     * The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
     */
    readonly maintenanceWindow: outputs.metastore.v1beta.MaintenanceWindowResponse;
    /**
     * The setting that defines how metastore metadata should be integrated with external services and systems.
     */
    readonly metadataIntegration: outputs.metastore.v1beta.MetadataIntegrationResponse;
    /**
     * The metadata management activities of the metastore service.
     */
    readonly metadataManagementActivity: outputs.metastore.v1beta.MetadataManagementActivityResponse;
    /**
     * Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
     */
    readonly name: string;
    /**
     * Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
     */
    readonly network: string;
    /**
     * Immutable. The configuration specifying the network settings for the Dataproc Metastore service.
     */
    readonly networkConfig: outputs.metastore.v1beta.NetworkConfigResponse;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    readonly port: number;
    /**
     * Immutable. The release channel of the service. If unspecified, defaults to STABLE.
     */
    readonly releaseChannel: string;
    /**
     * The current state of the metastore service.
     */
    readonly state: string;
    /**
     * Additional information about the current state of the metastore service, if available.
     */
    readonly stateMessage: string;
    /**
     * The tier of the service.
     */
    readonly tier: string;
    /**
     * The globally unique resource identifier of the metastore service.
     */
    readonly uid: string;
    /**
     * The time when the metastore service was last updated.
     */
    readonly updateTime: string;
}

export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

export interface GetServiceOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
}
