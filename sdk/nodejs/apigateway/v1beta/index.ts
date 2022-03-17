// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./api";
export * from "./apiConfigIamPolicy";
export * from "./apiIamPolicy";
export * from "./config";
export * from "./gateway";
export * from "./gatewayIamPolicy";
export * from "./getApi";
export * from "./getApiConfigIamPolicy";
export * from "./getApiIamPolicy";
export * from "./getConfig";
export * from "./getGateway";
export * from "./getGatewayIamPolicy";

// Export enums:
export * from "../../types/enums/apigateway/v1beta";

// Import resources to register:
import { Api } from "./api";
import { ApiConfigIamPolicy } from "./apiConfigIamPolicy";
import { ApiIamPolicy } from "./apiIamPolicy";
import { Config } from "./config";
import { Gateway } from "./gateway";
import { GatewayIamPolicy } from "./gatewayIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:apigateway/v1beta:Api":
                return new Api(name, <any>undefined, { urn })
            case "google-native:apigateway/v1beta:ApiConfigIamPolicy":
                return new ApiConfigIamPolicy(name, <any>undefined, { urn })
            case "google-native:apigateway/v1beta:ApiIamPolicy":
                return new ApiIamPolicy(name, <any>undefined, { urn })
            case "google-native:apigateway/v1beta:Config":
                return new Config(name, <any>undefined, { urn })
            case "google-native:apigateway/v1beta:Gateway":
                return new Gateway(name, <any>undefined, { urn })
            case "google-native:apigateway/v1beta:GatewayIamPolicy":
                return new GatewayIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "apigateway/v1beta", _module)
