// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.PulsarJobArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.PulsarJobState;
import com.pulumi.ns1.outputs.PulsarJobBlendMetricWeights;
import com.pulumi.ns1.outputs.PulsarJobConfig;
import com.pulumi.ns1.outputs.PulsarJobWeight;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="ns1:index/pulsarJob:PulsarJob")
public class PulsarJob extends com.pulumi.resources.CustomResource {
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> active;

    public Output<Optional<Boolean>> active() {
        return Codegen.optional(this.active);
    }
    @Export(name="appId", refs={String.class}, tree="[0]")
    private Output<String> appId;

    public Output<String> appId() {
        return this.appId;
    }
    @Export(name="blendMetricWeights", refs={PulsarJobBlendMetricWeights.class}, tree="[0]")
    private Output</* @Nullable */ PulsarJobBlendMetricWeights> blendMetricWeights;

    public Output<Optional<PulsarJobBlendMetricWeights>> blendMetricWeights() {
        return Codegen.optional(this.blendMetricWeights);
    }
    @Export(name="community", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> community;

    public Output<Boolean> community() {
        return this.community;
    }
    @Export(name="config", refs={PulsarJobConfig.class}, tree="[0]")
    private Output</* @Nullable */ PulsarJobConfig> config;

    public Output<Optional<PulsarJobConfig>> config() {
        return Codegen.optional(this.config);
    }
    @Export(name="customer", refs={Integer.class}, tree="[0]")
    private Output<Integer> customer;

    public Output<Integer> customer() {
        return this.customer;
    }
    @Export(name="jobId", refs={String.class}, tree="[0]")
    private Output<String> jobId;

    public Output<String> jobId() {
        return this.jobId;
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="shared", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> shared;

    public Output<Optional<Boolean>> shared() {
        return Codegen.optional(this.shared);
    }
    @Export(name="typeId", refs={String.class}, tree="[0]")
    private Output<String> typeId;

    public Output<String> typeId() {
        return this.typeId;
    }
    @Export(name="weights", refs={List.class,PulsarJobWeight.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PulsarJobWeight>> weights;

    public Output<Optional<List<PulsarJobWeight>>> weights() {
        return Codegen.optional(this.weights);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PulsarJob(java.lang.String name) {
        this(name, PulsarJobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PulsarJob(java.lang.String name, PulsarJobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PulsarJob(java.lang.String name, PulsarJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/pulsarJob:PulsarJob", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PulsarJob(java.lang.String name, Output<java.lang.String> id, @Nullable PulsarJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/pulsarJob:PulsarJob", name, state, makeResourceOptions(options, id), false);
    }

    private static PulsarJobArgs makeArgs(PulsarJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PulsarJobArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static PulsarJob get(java.lang.String name, Output<java.lang.String> id, @Nullable PulsarJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PulsarJob(name, id, state, options);
    }
}
