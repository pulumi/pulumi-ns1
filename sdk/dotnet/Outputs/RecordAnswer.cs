// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Outputs
{

    [OutputType]
    public sealed class RecordAnswer
    {
        /// <summary>
        /// Space delimited string of RDATA fields dependent on the record type.
        /// 
        /// A:
        /// 
        /// answer = "1.2.3.4"
        /// 
        /// CNAME:
        /// 
        /// answer = "www.example.com"
        /// 
        /// MX:
        /// 
        /// answer = "5 mail.example.com"
        /// 
        /// SRV:
        /// 
        /// answer = "10 0 2380 node-1.example.com"
        /// 
        /// SPF:
        /// 
        /// answer = "v=DKIM1; k=rsa; p=XXXXXXXX"
        /// </summary>
        public readonly string? Answer;
        public readonly ImmutableDictionary<string, object>? Meta;
        /// <summary>
        /// The region (Answer Group really) that this answer
        /// belongs to. This should be one of the names specified in `regions`. Only a
        /// single `region` per answer is currently supported. If you want an answer in
        /// multiple regions, duplicating the answer (including metadata) is the correct
        /// approach.
        /// * ` meta` - (Optional) meta is supported at the `answer` level. Meta
        /// is documented below.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private RecordAnswer(
            string? answer,

            ImmutableDictionary<string, object>? meta,

            string? region)
        {
            Answer = answer;
            Meta = meta;
            Region = region;
        }
    }
}