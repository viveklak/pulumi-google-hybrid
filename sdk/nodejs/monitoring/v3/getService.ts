// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Get the named Service.
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:monitoring/v3:getService", {
        "serviceId": args.serviceId,
        "v3Id": args.v3Id,
        "v3Id1": args.v3Id1,
    }, opts);
}

export interface GetServiceArgs {
    serviceId: string;
    v3Id: string;
    v3Id1: string;
}

export interface GetServiceResult {
    /**
     * Type used for App Engine services.
     */
    readonly appEngine: outputs.monitoring.v3.AppEngineResponse;
    /**
     * Type used for Cloud Endpoints services.
     */
    readonly cloudEndpoints: outputs.monitoring.v3.CloudEndpointsResponse;
    /**
     * Type used for Istio services that live in a Kubernetes cluster.
     */
    readonly clusterIstio: outputs.monitoring.v3.ClusterIstioResponse;
    /**
     * Custom service type.
     */
    readonly custom: outputs.monitoring.v3.CustomResponse;
    /**
     * Name used for UI elements listing this Service.
     */
    readonly displayName: string;
    /**
     * Type used for canonical services scoped to an Istio mesh. Metrics for Istio are documented here (https://istio.io/latest/docs/reference/config/metrics/)
     */
    readonly istioCanonicalService: outputs.monitoring.v3.IstioCanonicalServiceResponse;
    /**
     * Type used for Istio services scoped to an Istio mesh.
     */
    readonly meshIstio: outputs.monitoring.v3.MeshIstioResponse;
    /**
     * Resource name for this Service. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID] 
     */
    readonly name: string;
    /**
     * Configuration for how to query telemetry on a Service.
     */
    readonly telemetry: outputs.monitoring.v3.TelemetryResponse;
    /**
     * Labels which have been used to annotate the service. Label keys must start with a letter. Label keys and values may contain lowercase letters, numbers, underscores, and dashes. Label keys and values have a maximum length of 63 characters, and must be less than 128 bytes in size. Up to 64 label entries may be stored. For labels which do not have a semantic value, the empty string may be supplied for the label value.
     */
    readonly userLabels: {[key: string]: string};
}

export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

export interface GetServiceOutputArgs {
    serviceId: pulumi.Input<string>;
    v3Id: pulumi.Input<string>;
    v3Id1: pulumi.Input<string>;
}
