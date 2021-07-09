// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastore.V1.Outputs
{

    [OutputType]
    public sealed class GoogleDatastoreAdminV1IndexedPropertyResponse
    {
        /// <summary>
        /// The indexed property's direction. Must not be DIRECTION_UNSPECIFIED.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The property name to index.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GoogleDatastoreAdminV1IndexedPropertyResponse(
            string direction,

            string name)
        {
            Direction = direction;
            Name = name;
        }
    }
}
