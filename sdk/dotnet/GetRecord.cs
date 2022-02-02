// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    public static class GetRecord
    {
        /// <summary>
        /// Provides details about a NS1 Record. Use this if you would simply like to read
        /// information from NS1 into your configurations. For read/write operations, you
        /// should use a resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ns1 = Pulumi.Ns1;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Ns1.GetRecord.InvokeAsync(new Ns1.GetRecordArgs
        ///         {
        ///             Domain = "terraform.example.io",
        ///             Type = "A",
        ///             Zone = "example.io",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRecordResult> InvokeAsync(GetRecordArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRecordResult>("ns1:index/getRecord:getRecord", args ?? new GetRecordArgs(), options.WithDefaults());

        /// <summary>
        /// Provides details about a NS1 Record. Use this if you would simply like to read
        /// information from NS1 into your configurations. For read/write operations, you
        /// should use a resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ns1 = Pulumi.Ns1;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Ns1.GetRecord.InvokeAsync(new Ns1.GetRecordArgs
        ///         {
        ///             Domain = "terraform.example.io",
        ///             Type = "A",
        ///             Zone = "example.io",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRecordResult> Invoke(GetRecordInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRecordResult>("ns1:index/getRecord:getRecord", args ?? new GetRecordInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRecordArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The records' domain.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        /// <summary>
        /// The records' RR type.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        /// <summary>
        /// The zone the record belongs to.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetRecordArgs()
        {
        }
    }

    public sealed class GetRecordInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The records' domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The records' RR type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The zone the record belongs to.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetRecordInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRecordResult
    {
        /// <summary>
        /// List of NS1 answers.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRecordAnswerResult> Answers;
        public readonly string Domain;
        /// <summary>
        /// List of NS1 filters.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRecordFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The target record this links to.
        /// </summary>
        public readonly string Link;
        /// <summary>
        /// Map of metadata
        /// </summary>
        public readonly ImmutableDictionary<string, object> Meta;
        /// <summary>
        /// List of regions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRecordRegionResult> Regions;
        public readonly ImmutableArray<string> ShortAnswers;
        /// <summary>
        /// The records' time to live (in seconds).
        /// </summary>
        public readonly int Ttl;
        public readonly string Type;
        /// <summary>
        /// Whether to use EDNS client subnet data when available (in filter chain).
        /// </summary>
        public readonly bool UseClientSubnet;
        public readonly string Zone;

        [OutputConstructor]
        private GetRecordResult(
            ImmutableArray<Outputs.GetRecordAnswerResult> answers,

            string domain,

            ImmutableArray<Outputs.GetRecordFilterResult> filters,

            string id,

            string link,

            ImmutableDictionary<string, object> meta,

            ImmutableArray<Outputs.GetRecordRegionResult> regions,

            ImmutableArray<string> shortAnswers,

            int ttl,

            string type,

            bool useClientSubnet,

            string zone)
        {
            Answers = answers;
            Domain = domain;
            Filters = filters;
            Id = id;
            Link = link;
            Meta = meta;
            Regions = regions;
            ShortAnswers = shortAnswers;
            Ttl = ttl;
            Type = type;
            UseClientSubnet = useClientSubnet;
            Zone = zone;
        }
    }
}
