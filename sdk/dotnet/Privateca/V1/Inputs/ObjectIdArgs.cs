// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Inputs
{

    /// <summary>
    /// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
    /// </summary>
    public sealed class ObjectIdArgs : Pulumi.ResourceArgs
    {
        [Input("objectIdPath", required: true)]
        private InputList<int>? _objectIdPath;

        /// <summary>
        /// The parts of an OID path. The most significant parts of the path come first.
        /// </summary>
        public InputList<int> ObjectIdPath
        {
            get => _objectIdPath ?? (_objectIdPath = new InputList<int>());
            set => _objectIdPath = value;
        }

        public ObjectIdArgs()
        {
        }
    }
}
