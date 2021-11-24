// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Cloud Dataflow job. To create a job, we recommend using `projects.locations.jobs.create` with a [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints). Using `projects.jobs.create` is not recommended, as your job will always start in `us-central1`.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataflow/v1b3:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * The client's unique identifier of the job, re-used across retried attempts. If this field is set, the service will ensure its uniqueness. The request to create a job will fail if the service has knowledge of a previously submitted job with the same client's ID and job name. The caller may use this field to ensure idempotence of job creation across retried attempts to create a job. By default, the field is empty and, in that case, the service ignores it.
     */
    public readonly clientRequestId!: pulumi.Output<string>;
    /**
     * The timestamp when the job was initially created. Immutable and set by the Cloud Dataflow service.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * If this is specified, the job's initial state is populated from the given snapshot.
     */
    public readonly createdFromSnapshotId!: pulumi.Output<string>;
    /**
     * The current state of the job. Jobs are created in the `JOB_STATE_STOPPED` state unless otherwise specified. A job in the `JOB_STATE_RUNNING` state may asynchronously enter a terminal state. After a job has reached a terminal state, no further state updates may be made. This field may be mutated by the Cloud Dataflow service; callers cannot mutate it.
     */
    public readonly currentState!: pulumi.Output<string>;
    /**
     * The timestamp associated with the current state.
     */
    public readonly currentStateTime!: pulumi.Output<string>;
    /**
     * The environment for the job.
     */
    public readonly environment!: pulumi.Output<outputs.dataflow.v1b3.EnvironmentResponse>;
    /**
     * This field is populated by the Dataflow service to support filtering jobs by the metadata values provided here. Populated for ListJobs and all GetJob views SUMMARY and higher.
     */
    public readonly jobMetadata!: pulumi.Output<outputs.dataflow.v1b3.JobMetadataResponse>;
    /**
     * User-defined labels for this job. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: * Keys must conform to regexp: \p{Ll}\p{Lo}{0,62} * Values must conform to regexp: [\p{Ll}\p{Lo}\p{N}_-]{0,63} * Both keys and values are additionally constrained to be <= 128 bytes in size.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) that contains this job.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The user-specified Cloud Dataflow job name. Only one Job with a given name may exist in a project at any given time. If a caller attempts to create a Job with the same name as an already-existing Job, the attempt returns the existing Job. The name must match the regular expression `[a-z]([-a-z0-9]{0,38}[a-z0-9])?`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Preliminary field: The format of this data may change at any time. A description of the user pipeline and stages through which it is executed. Created by Cloud Dataflow service. Only retrieved with JOB_VIEW_DESCRIPTION or JOB_VIEW_ALL.
     */
    public readonly pipelineDescription!: pulumi.Output<outputs.dataflow.v1b3.PipelineDescriptionResponse>;
    /**
     * The ID of the Cloud Platform project that the job belongs to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * If this job is an update of an existing job, this field is the job ID of the job it replaced. When sending a `CreateJobRequest`, you can update a job by specifying it here. The job named here is stopped, and its intermediate state is transferred to this job.
     */
    public readonly replaceJobId!: pulumi.Output<string>;
    /**
     * If another job is an update of this job (and thus, this job is in `JOB_STATE_UPDATED`), this field contains the ID of that job.
     */
    public readonly replacedByJobId!: pulumi.Output<string>;
    /**
     * The job's requested state. `UpdateJob` may be used to switch between the `JOB_STATE_STOPPED` and `JOB_STATE_RUNNING` states, by setting requested_state. `UpdateJob` may also be used to directly set a job's requested state to `JOB_STATE_CANCELLED` or `JOB_STATE_DONE`, irrevocably terminating the job if it has not already reached a terminal state.
     */
    public readonly requestedState!: pulumi.Output<string>;
    /**
     * Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
     */
    public readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * This field may be mutated by the Cloud Dataflow service; callers cannot mutate it.
     */
    public readonly stageStates!: pulumi.Output<outputs.dataflow.v1b3.ExecutionStageStateResponse[]>;
    /**
     * The timestamp when the job was started (transitioned to JOB_STATE_PENDING). Flexible resource scheduling jobs are started with some delay after job creation, so start_time is unset before start and is updated when the job is started by the Cloud Dataflow service. For other jobs, start_time always equals to create_time and is immutable and set by the Cloud Dataflow service.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * Exactly one of step or steps_location should be specified. The top-level steps that constitute the entire job. Only retrieved with JOB_VIEW_ALL.
     */
    public readonly steps!: pulumi.Output<outputs.dataflow.v1b3.StepResponse[]>;
    /**
     * The Cloud Storage location where the steps are stored.
     */
    public readonly stepsLocation!: pulumi.Output<string>;
    /**
     * A set of files the system should be aware of that are used for temporary storage. These temporary files will be removed on job completion. No duplicates are allowed. No file patterns are supported. The supported files are: Google Cloud Storage: storage.googleapis.com/{bucket}/{object} bucket.storage.googleapis.com/{object}
     */
    public readonly tempFiles!: pulumi.Output<string[]>;
    /**
     * The map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job.
     */
    public readonly transformNameMapping!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of Cloud Dataflow job.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["clientRequestId"] = args ? args.clientRequestId : undefined;
            resourceInputs["createTime"] = args ? args.createTime : undefined;
            resourceInputs["createdFromSnapshotId"] = args ? args.createdFromSnapshotId : undefined;
            resourceInputs["currentState"] = args ? args.currentState : undefined;
            resourceInputs["currentStateTime"] = args ? args.currentStateTime : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["jobMetadata"] = args ? args.jobMetadata : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pipelineDescription"] = args ? args.pipelineDescription : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["replaceJobId"] = args ? args.replaceJobId : undefined;
            resourceInputs["replacedByJobId"] = args ? args.replacedByJobId : undefined;
            resourceInputs["requestedState"] = args ? args.requestedState : undefined;
            resourceInputs["satisfiesPzs"] = args ? args.satisfiesPzs : undefined;
            resourceInputs["stageStates"] = args ? args.stageStates : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["steps"] = args ? args.steps : undefined;
            resourceInputs["stepsLocation"] = args ? args.stepsLocation : undefined;
            resourceInputs["tempFiles"] = args ? args.tempFiles : undefined;
            resourceInputs["transformNameMapping"] = args ? args.transformNameMapping : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["view"] = args ? args.view : undefined;
        } else {
            resourceInputs["clientRequestId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["createdFromSnapshotId"] = undefined /*out*/;
            resourceInputs["currentState"] = undefined /*out*/;
            resourceInputs["currentStateTime"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["jobMetadata"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pipelineDescription"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["replaceJobId"] = undefined /*out*/;
            resourceInputs["replacedByJobId"] = undefined /*out*/;
            resourceInputs["requestedState"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["stageStates"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["steps"] = undefined /*out*/;
            resourceInputs["stepsLocation"] = undefined /*out*/;
            resourceInputs["tempFiles"] = undefined /*out*/;
            resourceInputs["transformNameMapping"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * The client's unique identifier of the job, re-used across retried attempts. If this field is set, the service will ensure its uniqueness. The request to create a job will fail if the service has knowledge of a previously submitted job with the same client's ID and job name. The caller may use this field to ensure idempotence of job creation across retried attempts to create a job. By default, the field is empty and, in that case, the service ignores it.
     */
    clientRequestId?: pulumi.Input<string>;
    /**
     * The timestamp when the job was initially created. Immutable and set by the Cloud Dataflow service.
     */
    createTime?: pulumi.Input<string>;
    /**
     * If this is specified, the job's initial state is populated from the given snapshot.
     */
    createdFromSnapshotId?: pulumi.Input<string>;
    /**
     * The current state of the job. Jobs are created in the `JOB_STATE_STOPPED` state unless otherwise specified. A job in the `JOB_STATE_RUNNING` state may asynchronously enter a terminal state. After a job has reached a terminal state, no further state updates may be made. This field may be mutated by the Cloud Dataflow service; callers cannot mutate it.
     */
    currentState?: pulumi.Input<enums.dataflow.v1b3.JobCurrentState>;
    /**
     * The timestamp associated with the current state.
     */
    currentStateTime?: pulumi.Input<string>;
    /**
     * The environment for the job.
     */
    environment?: pulumi.Input<inputs.dataflow.v1b3.EnvironmentArgs>;
    /**
     * The unique ID of this job. This field is set by the Cloud Dataflow service when the Job is created, and is immutable for the life of the job.
     */
    id?: pulumi.Input<string>;
    /**
     * This field is populated by the Dataflow service to support filtering jobs by the metadata values provided here. Populated for ListJobs and all GetJob views SUMMARY and higher.
     */
    jobMetadata?: pulumi.Input<inputs.dataflow.v1b3.JobMetadataArgs>;
    /**
     * User-defined labels for this job. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: * Keys must conform to regexp: \p{Ll}\p{Lo}{0,62} * Values must conform to regexp: [\p{Ll}\p{Lo}\p{N}_-]{0,63} * Both keys and values are additionally constrained to be <= 128 bytes in size.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) that contains this job.
     */
    location?: pulumi.Input<string>;
    /**
     * The user-specified Cloud Dataflow job name. Only one Job with a given name may exist in a project at any given time. If a caller attempts to create a Job with the same name as an already-existing Job, the attempt returns the existing Job. The name must match the regular expression `[a-z]([-a-z0-9]{0,38}[a-z0-9])?`
     */
    name?: pulumi.Input<string>;
    /**
     * Preliminary field: The format of this data may change at any time. A description of the user pipeline and stages through which it is executed. Created by Cloud Dataflow service. Only retrieved with JOB_VIEW_DESCRIPTION or JOB_VIEW_ALL.
     */
    pipelineDescription?: pulumi.Input<inputs.dataflow.v1b3.PipelineDescriptionArgs>;
    /**
     * The ID of the Cloud Platform project that the job belongs to.
     */
    project?: pulumi.Input<string>;
    /**
     * If this job is an update of an existing job, this field is the job ID of the job it replaced. When sending a `CreateJobRequest`, you can update a job by specifying it here. The job named here is stopped, and its intermediate state is transferred to this job.
     */
    replaceJobId?: pulumi.Input<string>;
    /**
     * If another job is an update of this job (and thus, this job is in `JOB_STATE_UPDATED`), this field contains the ID of that job.
     */
    replacedByJobId?: pulumi.Input<string>;
    /**
     * The job's requested state. `UpdateJob` may be used to switch between the `JOB_STATE_STOPPED` and `JOB_STATE_RUNNING` states, by setting requested_state. `UpdateJob` may also be used to directly set a job's requested state to `JOB_STATE_CANCELLED` or `JOB_STATE_DONE`, irrevocably terminating the job if it has not already reached a terminal state.
     */
    requestedState?: pulumi.Input<enums.dataflow.v1b3.JobRequestedState>;
    /**
     * Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
     */
    satisfiesPzs?: pulumi.Input<boolean>;
    /**
     * This field may be mutated by the Cloud Dataflow service; callers cannot mutate it.
     */
    stageStates?: pulumi.Input<pulumi.Input<inputs.dataflow.v1b3.ExecutionStageStateArgs>[]>;
    /**
     * The timestamp when the job was started (transitioned to JOB_STATE_PENDING). Flexible resource scheduling jobs are started with some delay after job creation, so start_time is unset before start and is updated when the job is started by the Cloud Dataflow service. For other jobs, start_time always equals to create_time and is immutable and set by the Cloud Dataflow service.
     */
    startTime?: pulumi.Input<string>;
    /**
     * Exactly one of step or steps_location should be specified. The top-level steps that constitute the entire job. Only retrieved with JOB_VIEW_ALL.
     */
    steps?: pulumi.Input<pulumi.Input<inputs.dataflow.v1b3.StepArgs>[]>;
    /**
     * The Cloud Storage location where the steps are stored.
     */
    stepsLocation?: pulumi.Input<string>;
    /**
     * A set of files the system should be aware of that are used for temporary storage. These temporary files will be removed on job completion. No duplicates are allowed. No file patterns are supported. The supported files are: Google Cloud Storage: storage.googleapis.com/{bucket}/{object} bucket.storage.googleapis.com/{object}
     */
    tempFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job.
     */
    transformNameMapping?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of Cloud Dataflow job.
     */
    type?: pulumi.Input<enums.dataflow.v1b3.JobType>;
    view?: pulumi.Input<string>;
}
