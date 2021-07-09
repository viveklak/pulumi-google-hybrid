// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Data Fusion instance in the specified project and location.
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
    public static readonly __pulumiType = 'google-native:datafusion/v1beta1:Instance';

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
     * List of accelerators enabled for this CDF instance.
     */
    public readonly accelerators!: pulumi.Output<outputs.datafusion.v1beta1.AcceleratorResponse[]>;
    /**
     * Endpoint on which the REST APIs is accessible.
     */
    public /*out*/ readonly apiEndpoint!: pulumi.Output<string>;
    /**
     * Available versions that the instance can be upgraded to using UpdateInstanceRequest.
     */
    public readonly availableVersion!: pulumi.Output<outputs.datafusion.v1beta1.VersionResponse[]>;
    /**
     * The time the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
     */
    public readonly dataprocServiceAccount!: pulumi.Output<string>;
    /**
     * A description of this instance.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name for an instance.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Option to enable granular role-based access control.
     */
    public readonly enableRbac!: pulumi.Output<boolean>;
    /**
     * Option to enable Stackdriver Logging.
     */
    public readonly enableStackdriverLogging!: pulumi.Output<boolean>;
    /**
     * Option to enable Stackdriver Monitoring.
     */
    public readonly enableStackdriverMonitoring!: pulumi.Output<boolean>;
    /**
     * Cloud Storage bucket generated by Data Fusion in the customer project.
     */
    public /*out*/ readonly gcsBucket!: pulumi.Output<string>;
    /**
     * The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network configuration options. These are required when a private Data Fusion instance is to be created.
     */
    public readonly networkConfig!: pulumi.Output<outputs.datafusion.v1beta1.NetworkConfigResponse>;
    /**
     * Map of additional options used to configure the behavior of Data Fusion instance.
     */
    public readonly options!: pulumi.Output<{[key: string]: string}>;
    /**
     * P4 service account for the customer project.
     */
    public /*out*/ readonly p4ServiceAccount!: pulumi.Output<string>;
    /**
     * Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
     */
    public readonly privateInstance!: pulumi.Output<boolean>;
    /**
     * Endpoint on which the Data Fusion UI is accessible.
     */
    public /*out*/ readonly serviceEndpoint!: pulumi.Output<string>;
    /**
     * The current state of this Data Fusion instance.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current state of this Data Fusion instance if available.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * The name of the tenant project.
     */
    public /*out*/ readonly tenantProjectId!: pulumi.Output<string>;
    /**
     * Instance type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The time the instance was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Current version of Data Fusion.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
     */
    public readonly zone!: pulumi.Output<string>;

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
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["accelerators"] = args ? args.accelerators : undefined;
            inputs["availableVersion"] = args ? args.availableVersion : undefined;
            inputs["dataprocServiceAccount"] = args ? args.dataprocServiceAccount : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enableRbac"] = args ? args.enableRbac : undefined;
            inputs["enableStackdriverLogging"] = args ? args.enableStackdriverLogging : undefined;
            inputs["enableStackdriverMonitoring"] = args ? args.enableStackdriverMonitoring : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkConfig"] = args ? args.networkConfig : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["privateInstance"] = args ? args.privateInstance : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["apiEndpoint"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["gcsBucket"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["p4ServiceAccount"] = undefined /*out*/;
            inputs["serviceEndpoint"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["stateMessage"] = undefined /*out*/;
            inputs["tenantProjectId"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["accelerators"] = undefined /*out*/;
            inputs["apiEndpoint"] = undefined /*out*/;
            inputs["availableVersion"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["dataprocServiceAccount"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["enableRbac"] = undefined /*out*/;
            inputs["enableStackdriverLogging"] = undefined /*out*/;
            inputs["enableStackdriverMonitoring"] = undefined /*out*/;
            inputs["gcsBucket"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkConfig"] = undefined /*out*/;
            inputs["options"] = undefined /*out*/;
            inputs["p4ServiceAccount"] = undefined /*out*/;
            inputs["privateInstance"] = undefined /*out*/;
            inputs["serviceEndpoint"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["stateMessage"] = undefined /*out*/;
            inputs["tenantProjectId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
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
     * List of accelerators enabled for this CDF instance.
     */
    accelerators?: pulumi.Input<pulumi.Input<inputs.datafusion.v1beta1.AcceleratorArgs>[]>;
    /**
     * Available versions that the instance can be upgraded to using UpdateInstanceRequest.
     */
    availableVersion?: pulumi.Input<pulumi.Input<inputs.datafusion.v1beta1.VersionArgs>[]>;
    /**
     * User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
     */
    dataprocServiceAccount?: pulumi.Input<string>;
    /**
     * A description of this instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Display name for an instance.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Option to enable granular role-based access control.
     */
    enableRbac?: pulumi.Input<boolean>;
    /**
     * Option to enable Stackdriver Logging.
     */
    enableStackdriverLogging?: pulumi.Input<boolean>;
    /**
     * Option to enable Stackdriver Monitoring.
     */
    enableStackdriverMonitoring?: pulumi.Input<boolean>;
    instanceId?: pulumi.Input<string>;
    /**
     * The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location: pulumi.Input<string>;
    /**
     * Network configuration options. These are required when a private Data Fusion instance is to be created.
     */
    networkConfig?: pulumi.Input<inputs.datafusion.v1beta1.NetworkConfigArgs>;
    /**
     * Map of additional options used to configure the behavior of Data Fusion instance.
     */
    options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
     */
    privateInstance?: pulumi.Input<boolean>;
    project: pulumi.Input<string>;
    /**
     * Instance type.
     */
    type: pulumi.Input<enums.datafusion.v1beta1.InstanceType>;
    /**
     * Current version of Data Fusion.
     */
    version?: pulumi.Input<string>;
    /**
     * Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
     */
    zone?: pulumi.Input<string>;
}
