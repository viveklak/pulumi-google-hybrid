// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./alertPolicy";
export * from "./collectdTimeSeries";
export * from "./group";
export * from "./metricDescriptor";
export * from "./notificationChannel";
export * from "./service";
export * from "./serviceServiceLevelObjective";
export * from "./timeSeries";
export * from "./uptimeCheckConfig";

// Import resources to register:
import { AlertPolicy } from "./alertPolicy";
import { CollectdTimeSeries } from "./collectdTimeSeries";
import { Group } from "./group";
import { MetricDescriptor } from "./metricDescriptor";
import { NotificationChannel } from "./notificationChannel";
import { Service } from "./service";
import { ServiceServiceLevelObjective } from "./serviceServiceLevelObjective";
import { TimeSeries } from "./timeSeries";
import { UptimeCheckConfig } from "./uptimeCheckConfig";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:monitoring/v3:AlertPolicy":
                return new AlertPolicy(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:CollectdTimeSeries":
                return new CollectdTimeSeries(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:Group":
                return new Group(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:MetricDescriptor":
                return new MetricDescriptor(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:NotificationChannel":
                return new NotificationChannel(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:Service":
                return new Service(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:ServiceServiceLevelObjective":
                return new ServiceServiceLevelObjective(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:TimeSeries":
                return new TimeSeries(name, <any>undefined, { urn })
            case "google-cloud:monitoring/v3:UptimeCheckConfig":
                return new UptimeCheckConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "monitoring/v3", _module)
