// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Use this method to create a connection profile in a project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class ConnectionProfile extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectionProfile {
        return new ConnectionProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datastream/v1:ConnectionProfile';

    /**
     * Returns true if the given object is an instance of ConnectionProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionProfile.__pulumiType;
    }

    /**
     * The create time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Forward SSH tunnel connectivity.
     */
    public readonly forwardSshConnectivity!: pulumi.Output<outputs.datastream.v1.ForwardSshTunnelConnectivityResponse>;
    /**
     * Cloud Storage ConnectionProfile configuration.
     */
    public readonly gcsProfile!: pulumi.Output<outputs.datastream.v1.GcsProfileResponse>;
    /**
     * Labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * MySQL ConnectionProfile configuration.
     */
    public readonly mysqlProfile!: pulumi.Output<outputs.datastream.v1.MysqlProfileResponse>;
    /**
     * The resource's name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Oracle ConnectionProfile configuration.
     */
    public readonly oracleProfile!: pulumi.Output<outputs.datastream.v1.OracleProfileResponse>;
    /**
     * Private connectivity.
     */
    public readonly privateConnectivity!: pulumi.Output<outputs.datastream.v1.PrivateConnectivityResponse>;
    /**
     * Static Service IP connectivity.
     */
    public readonly staticServiceIpConnectivity!: pulumi.Output<outputs.datastream.v1.StaticServiceIpConnectivityResponse>;
    /**
     * The update time of the resource.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ConnectionProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectionProfileId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionProfileId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["connectionProfileId"] = args ? args.connectionProfileId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["forwardSshConnectivity"] = args ? args.forwardSshConnectivity : undefined;
            resourceInputs["gcsProfile"] = args ? args.gcsProfile : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["mysqlProfile"] = args ? args.mysqlProfile : undefined;
            resourceInputs["oracleProfile"] = args ? args.oracleProfile : undefined;
            resourceInputs["privateConnectivity"] = args ? args.privateConnectivity : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["staticServiceIpConnectivity"] = args ? args.staticServiceIpConnectivity : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["forwardSshConnectivity"] = undefined /*out*/;
            resourceInputs["gcsProfile"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["mysqlProfile"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["oracleProfile"] = undefined /*out*/;
            resourceInputs["privateConnectivity"] = undefined /*out*/;
            resourceInputs["staticServiceIpConnectivity"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConnectionProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectionProfile resource.
 */
export interface ConnectionProfileArgs {
    connectionProfileId: pulumi.Input<string>;
    /**
     * Display name.
     */
    displayName: pulumi.Input<string>;
    force?: pulumi.Input<string>;
    /**
     * Forward SSH tunnel connectivity.
     */
    forwardSshConnectivity?: pulumi.Input<inputs.datastream.v1.ForwardSshTunnelConnectivityArgs>;
    /**
     * Cloud Storage ConnectionProfile configuration.
     */
    gcsProfile?: pulumi.Input<inputs.datastream.v1.GcsProfileArgs>;
    /**
     * Labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * MySQL ConnectionProfile configuration.
     */
    mysqlProfile?: pulumi.Input<inputs.datastream.v1.MysqlProfileArgs>;
    /**
     * Oracle ConnectionProfile configuration.
     */
    oracleProfile?: pulumi.Input<inputs.datastream.v1.OracleProfileArgs>;
    /**
     * Private connectivity.
     */
    privateConnectivity?: pulumi.Input<inputs.datastream.v1.PrivateConnectivityArgs>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * Static Service IP connectivity.
     */
    staticServiceIpConnectivity?: pulumi.Input<inputs.datastream.v1.StaticServiceIpConnectivityArgs>;
}
