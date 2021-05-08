// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIKeys.V2.Outputs
{

    [OutputType]
    public sealed class V2ApiTargetResponse
    {
        /// <summary>
        /// Optional. List of one or more methods that can be called. If empty, all methods for the service are allowed. A wildcard (*) can be used as the last symbol. Valid examples: `google.cloud.translate.v2.TranslateService.GetSupportedLanguage` `TranslateText` `Get*` `translate.googleapis.com.Get*`
        /// </summary>
        public readonly ImmutableArray<string> Methods;
        /// <summary>
        /// The service for this restriction. It should be the canonical service name, for example: `translate.googleapis.com`. You can use [`gcloud services list`](/sdk/gcloud/reference/services/list) to get a list of services that are enabled in the project.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private V2ApiTargetResponse(
            ImmutableArray<string> methods,

            string service)
        {
            Methods = methods;
            Service = service;
        }
    }
}