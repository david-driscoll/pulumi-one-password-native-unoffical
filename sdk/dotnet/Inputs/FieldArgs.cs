// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword.Inputs
{

    public sealed class FieldArgs : Pulumi.ResourceArgs
    {
        [Input("purpose", required: true)]
        public Input<Pulumi.Onepassword.FieldPurpose> Purpose { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public FieldArgs()
        {
            Purpose = Pulumi.Onepassword.FieldPurpose.Note;
        }
    }
}
