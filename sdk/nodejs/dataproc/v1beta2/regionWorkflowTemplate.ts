// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates new workflow template.
 */
export class RegionWorkflowTemplate extends pulumi.CustomResource {
    /**
     * Get an existing RegionWorkflowTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionWorkflowTemplate {
        return new RegionWorkflowTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataproc/v1beta2:RegionWorkflowTemplate';

    /**
     * Returns true if the given object is an instance of RegionWorkflowTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionWorkflowTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionWorkflowTemplate.__pulumiType;
    }

    /**
     * The time template was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
     */
    public readonly dagTimeout!: pulumi.Output<string>;
    /**
     * Required. The Directed Acyclic Graph of Jobs to submit.
     */
    public readonly jobs!: pulumi.Output<outputs.dataproc.v1beta2.OrderedJobResponse[]>;
    /**
     * Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
     */
    public readonly parameters!: pulumi.Output<outputs.dataproc.v1beta2.TemplateParameterResponse[]>;
    /**
     * Required. WorkflowTemplate scheduling information.
     */
    public readonly placement!: pulumi.Output<outputs.dataproc.v1beta2.WorkflowTemplatePlacementResponse>;
    /**
     * The time template was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a RegionWorkflowTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionWorkflowTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.regionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regionId'");
            }
            inputs["dagTimeout"] = args ? args.dagTimeout : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["jobs"] = args ? args.jobs : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["placement"] = args ? args.placement : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["regionId"] = args ? args.regionId : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["dagTimeout"] = undefined /*out*/;
            inputs["jobs"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["placement"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionWorkflowTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionWorkflowTemplate resource.
 */
export interface RegionWorkflowTemplateArgs {
    /**
     * Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
     */
    readonly dagTimeout?: pulumi.Input<string>;
    /**
     * Required. The template id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters..
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Required. The Directed Acyclic Graph of Jobs to submit.
     */
    readonly jobs?: pulumi.Input<pulumi.Input<inputs.dataproc.v1beta2.OrderedJobArgs>[]>;
    /**
     * Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.dataproc.v1beta2.TemplateParameterArgs>[]>;
    /**
     * Required. WorkflowTemplate scheduling information.
     */
    readonly placement?: pulumi.Input<inputs.dataproc.v1beta2.WorkflowTemplatePlacementArgs>;
    readonly project: pulumi.Input<string>;
    readonly regionId: pulumi.Input<string>;
    /**
     * Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
     */
    readonly version?: pulumi.Input<number>;
}
