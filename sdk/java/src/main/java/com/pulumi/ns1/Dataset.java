// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.DatasetArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.DatasetState;
import com.pulumi.ns1.outputs.DatasetDatatype;
import com.pulumi.ns1.outputs.DatasetRepeat;
import com.pulumi.ns1.outputs.DatasetReport;
import com.pulumi.ns1.outputs.DatasetTimeframe;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="ns1:index/dataset:Dataset")
public class Dataset extends com.pulumi.resources.CustomResource {
    @Export(name="datatype", refs={DatasetDatatype.class}, tree="[0]")
    private Output<DatasetDatatype> datatype;

    public Output<DatasetDatatype> datatype() {
        return this.datatype;
    }
    @Export(name="exportType", refs={String.class}, tree="[0]")
    private Output<String> exportType;

    public Output<String> exportType() {
        return this.exportType;
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="recipientEmails", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> recipientEmails;

    public Output<Optional<List<String>>> recipientEmails() {
        return Codegen.optional(this.recipientEmails);
    }
    @Export(name="repeat", refs={DatasetRepeat.class}, tree="[0]")
    private Output</* @Nullable */ DatasetRepeat> repeat;

    public Output<Optional<DatasetRepeat>> repeat() {
        return Codegen.optional(this.repeat);
    }
    @Export(name="reports", refs={List.class,DatasetReport.class}, tree="[0,1]")
    private Output<List<DatasetReport>> reports;

    public Output<List<DatasetReport>> reports() {
        return this.reports;
    }
    @Export(name="timeframe", refs={DatasetTimeframe.class}, tree="[0]")
    private Output<DatasetTimeframe> timeframe;

    public Output<DatasetTimeframe> timeframe() {
        return this.timeframe;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Dataset(String name) {
        this(name, DatasetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Dataset(String name, DatasetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Dataset(String name, DatasetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/dataset:Dataset", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private Dataset(String name, Output<String> id, @Nullable DatasetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/dataset:Dataset", name, state, makeResourceOptions(options, id));
    }

    private static DatasetArgs makeArgs(DatasetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DatasetArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Dataset get(String name, Output<String> id, @Nullable DatasetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Dataset(name, id, state, options);
    }
}
