// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./backup";
export * from "./getBackup";
export * from "./getMetadataImport";
export * from "./getService";
export * from "./getServiceBackupIamPolicy";
export * from "./getServiceDatabaseIamPolicy";
export * from "./getServiceDatabaseTableIamPolicy";
export * from "./getServiceIamPolicy";
export * from "./metadataImport";
export * from "./service";
export * from "./serviceBackupIamPolicy";
export * from "./serviceDatabaseIamPolicy";
export * from "./serviceDatabaseTableIamPolicy";
export * from "./serviceIamPolicy";

// Export enums:
export * from "../../types/enums/metastore/v1alpha";

// Import resources to register:
import { Backup } from "./backup";
import { MetadataImport } from "./metadataImport";
import { Service } from "./service";
import { ServiceBackupIamPolicy } from "./serviceBackupIamPolicy";
import { ServiceDatabaseIamPolicy } from "./serviceDatabaseIamPolicy";
import { ServiceDatabaseTableIamPolicy } from "./serviceDatabaseTableIamPolicy";
import { ServiceIamPolicy } from "./serviceIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:metastore/v1alpha:Backup":
                return new Backup(name, <any>undefined, { urn })
            case "google-native:metastore/v1alpha:MetadataImport":
                return new MetadataImport(name, <any>undefined, { urn })
            case "google-native:metastore/v1alpha:Service":
                return new Service(name, <any>undefined, { urn })
            case "google-native:metastore/v1alpha:ServiceBackupIamPolicy":
                return new ServiceBackupIamPolicy(name, <any>undefined, { urn })
            case "google-native:metastore/v1alpha:ServiceDatabaseIamPolicy":
                return new ServiceDatabaseIamPolicy(name, <any>undefined, { urn })
            case "google-native:metastore/v1alpha:ServiceDatabaseTableIamPolicy":
                return new ServiceDatabaseTableIamPolicy(name, <any>undefined, { urn })
            case "google-native:metastore/v1alpha:ServiceIamPolicy":
                return new ServiceIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "metastore/v1alpha", _module)
