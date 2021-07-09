// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves a pipeline based on ID. Caller must have READ permission to the project.
 */
export function getPipeline(args: GetPipelineArgs, opts?: pulumi.InvokeOptions): Promise<GetPipelineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:genomics/v1alpha2:getPipeline", {
        "pipelineId": args.pipelineId,
    }, opts);
}

export interface GetPipelineArgs {
    pipelineId: string;
}

export interface GetPipelineResult {
    /**
     * User-specified description.
     */
    readonly description: string;
    /**
     * Specifies the docker run information.
     */
    readonly docker: outputs.genomics.v1alpha2.DockerExecutorResponse;
    /**
     * Input parameters of the pipeline.
     */
    readonly inputParameters: outputs.genomics.v1alpha2.PipelineParameterResponse[];
    /**
     * A user specified pipeline name that does not have to be unique. This name can be used for filtering Pipelines in ListPipelines.
     */
    readonly name: string;
    /**
     * Output parameters of the pipeline.
     */
    readonly outputParameters: outputs.genomics.v1alpha2.PipelineParameterResponse[];
    /**
     * Unique pipeline id that is generated by the service when CreatePipeline is called. Cannot be specified in the Pipeline used in the CreatePipelineRequest, and will be populated in the response to CreatePipeline and all subsequent Get and List calls. Indicates that the service has registered this pipeline.
     */
    readonly pipelineId: string;
    /**
     * The project in which to create the pipeline. The caller must have WRITE access.
     */
    readonly project: string;
    /**
     * Specifies resource requirements for the pipeline run. Required fields: * minimumCpuCores * minimumRamGb
     */
    readonly resources: outputs.genomics.v1alpha2.PipelineResourcesResponse;
}
