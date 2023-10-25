// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a NS1 Monitoring Job resource. This can be used to create, modify, and delete monitoring jobs.
//
// ## NS1 Documentation
//
// [MonitoringJob Api Doc](https://ns1.com/api#monitoring-jobs)
//
// ## Import
//
// ```sh
//
//	$ pulumi import ns1:index/monitoringJob:MonitoringJob <name> <monitoringjob_id>`
//
// ```
type MonitoringJob struct {
	pulumi.CustomResourceState

	// Indicates if the job is active or temporarily disabled.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Config pulumi.MapOutput `pulumi:"config"`
	// The frequency, in seconds, at which to run the monitoring job in each region.
	Frequency pulumi.IntOutput `pulumi:"frequency"`
	// The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
	JobType pulumi.StringOutput `pulumi:"jobType"`
	// turn off the notifications for the monitoring job.
	Mute pulumi.BoolPtrOutput `pulumi:"mute"`
	// The free-form display name for the monitoring job.
	Name pulumi.StringOutput `pulumi:"name"`
	// Freeform notes to be included in any notifications about this job.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The time in seconds after a failure to wait before sending a notification.
	NotifyDelay pulumi.IntPtrOutput `pulumi:"notifyDelay"`
	// If true, a notification is sent when a job returns to an "up" state.
	NotifyFailback pulumi.BoolPtrOutput   `pulumi:"notifyFailback"`
	NotifyList     pulumi.StringPtrOutput `pulumi:"notifyList"`
	// If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
	NotifyRegional pulumi.BoolPtrOutput `pulumi:"notifyRegional"`
	// The time in seconds between repeat notifications of a failed job.
	NotifyRepeat pulumi.IntPtrOutput `pulumi:"notifyRepeat"`
	// The policy for determining the monitor's global status
	// based on the status of the job in all regions. See NS1 API docs for supported values.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
	RapidRecheck pulumi.BoolPtrOutput `pulumi:"rapidRecheck"`
	// The list of region codes in which to run the monitoring
	// job. See NS1 API docs for supported values.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Rules MonitoringJobRuleArrayOutput `pulumi:"rules"`
}

// NewMonitoringJob registers a new resource with the given unique name, arguments, and options.
func NewMonitoringJob(ctx *pulumi.Context,
	name string, args *MonitoringJobArgs, opts ...pulumi.ResourceOption) (*MonitoringJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.JobType == nil {
		return nil, errors.New("invalid value for required argument 'JobType'")
	}
	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MonitoringJob
	err := ctx.RegisterResource("ns1:index/monitoringJob:MonitoringJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringJob gets an existing MonitoringJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringJobState, opts ...pulumi.ResourceOption) (*MonitoringJob, error) {
	var resource MonitoringJob
	err := ctx.ReadResource("ns1:index/monitoringJob:MonitoringJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringJob resources.
type monitoringJobState struct {
	// Indicates if the job is active or temporarily disabled.
	Active *bool `pulumi:"active"`
	// A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Config map[string]interface{} `pulumi:"config"`
	// The frequency, in seconds, at which to run the monitoring job in each region.
	Frequency *int `pulumi:"frequency"`
	// The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
	JobType *string `pulumi:"jobType"`
	// turn off the notifications for the monitoring job.
	Mute *bool `pulumi:"mute"`
	// The free-form display name for the monitoring job.
	Name *string `pulumi:"name"`
	// Freeform notes to be included in any notifications about this job.
	Notes *string `pulumi:"notes"`
	// The time in seconds after a failure to wait before sending a notification.
	NotifyDelay *int `pulumi:"notifyDelay"`
	// If true, a notification is sent when a job returns to an "up" state.
	NotifyFailback *bool   `pulumi:"notifyFailback"`
	NotifyList     *string `pulumi:"notifyList"`
	// If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
	NotifyRegional *bool `pulumi:"notifyRegional"`
	// The time in seconds between repeat notifications of a failed job.
	NotifyRepeat *int `pulumi:"notifyRepeat"`
	// The policy for determining the monitor's global status
	// based on the status of the job in all regions. See NS1 API docs for supported values.
	Policy *string `pulumi:"policy"`
	// If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
	RapidRecheck *bool `pulumi:"rapidRecheck"`
	// The list of region codes in which to run the monitoring
	// job. See NS1 API docs for supported values.
	Regions []string `pulumi:"regions"`
	// A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Rules []MonitoringJobRule `pulumi:"rules"`
}

type MonitoringJobState struct {
	// Indicates if the job is active or temporarily disabled.
	Active pulumi.BoolPtrInput
	// A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Config pulumi.MapInput
	// The frequency, in seconds, at which to run the monitoring job in each region.
	Frequency pulumi.IntPtrInput
	// The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
	JobType pulumi.StringPtrInput
	// turn off the notifications for the monitoring job.
	Mute pulumi.BoolPtrInput
	// The free-form display name for the monitoring job.
	Name pulumi.StringPtrInput
	// Freeform notes to be included in any notifications about this job.
	Notes pulumi.StringPtrInput
	// The time in seconds after a failure to wait before sending a notification.
	NotifyDelay pulumi.IntPtrInput
	// If true, a notification is sent when a job returns to an "up" state.
	NotifyFailback pulumi.BoolPtrInput
	NotifyList     pulumi.StringPtrInput
	// If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
	NotifyRegional pulumi.BoolPtrInput
	// The time in seconds between repeat notifications of a failed job.
	NotifyRepeat pulumi.IntPtrInput
	// The policy for determining the monitor's global status
	// based on the status of the job in all regions. See NS1 API docs for supported values.
	Policy pulumi.StringPtrInput
	// If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
	RapidRecheck pulumi.BoolPtrInput
	// The list of region codes in which to run the monitoring
	// job. See NS1 API docs for supported values.
	Regions pulumi.StringArrayInput
	// A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Rules MonitoringJobRuleArrayInput
}

func (MonitoringJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringJobState)(nil)).Elem()
}

type monitoringJobArgs struct {
	// Indicates if the job is active or temporarily disabled.
	Active *bool `pulumi:"active"`
	// A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Config map[string]interface{} `pulumi:"config"`
	// The frequency, in seconds, at which to run the monitoring job in each region.
	Frequency int `pulumi:"frequency"`
	// The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
	JobType string `pulumi:"jobType"`
	// turn off the notifications for the monitoring job.
	Mute *bool `pulumi:"mute"`
	// The free-form display name for the monitoring job.
	Name *string `pulumi:"name"`
	// Freeform notes to be included in any notifications about this job.
	Notes *string `pulumi:"notes"`
	// The time in seconds after a failure to wait before sending a notification.
	NotifyDelay *int `pulumi:"notifyDelay"`
	// If true, a notification is sent when a job returns to an "up" state.
	NotifyFailback *bool   `pulumi:"notifyFailback"`
	NotifyList     *string `pulumi:"notifyList"`
	// If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
	NotifyRegional *bool `pulumi:"notifyRegional"`
	// The time in seconds between repeat notifications of a failed job.
	NotifyRepeat *int `pulumi:"notifyRepeat"`
	// The policy for determining the monitor's global status
	// based on the status of the job in all regions. See NS1 API docs for supported values.
	Policy *string `pulumi:"policy"`
	// If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
	RapidRecheck *bool `pulumi:"rapidRecheck"`
	// The list of region codes in which to run the monitoring
	// job. See NS1 API docs for supported values.
	Regions []string `pulumi:"regions"`
	// A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Rules []MonitoringJobRule `pulumi:"rules"`
}

// The set of arguments for constructing a MonitoringJob resource.
type MonitoringJobArgs struct {
	// Indicates if the job is active or temporarily disabled.
	Active pulumi.BoolPtrInput
	// A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Config pulumi.MapInput
	// The frequency, in seconds, at which to run the monitoring job in each region.
	Frequency pulumi.IntInput
	// The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
	JobType pulumi.StringInput
	// turn off the notifications for the monitoring job.
	Mute pulumi.BoolPtrInput
	// The free-form display name for the monitoring job.
	Name pulumi.StringPtrInput
	// Freeform notes to be included in any notifications about this job.
	Notes pulumi.StringPtrInput
	// The time in seconds after a failure to wait before sending a notification.
	NotifyDelay pulumi.IntPtrInput
	// If true, a notification is sent when a job returns to an "up" state.
	NotifyFailback pulumi.BoolPtrInput
	NotifyList     pulumi.StringPtrInput
	// If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
	NotifyRegional pulumi.BoolPtrInput
	// The time in seconds between repeat notifications of a failed job.
	NotifyRepeat pulumi.IntPtrInput
	// The policy for determining the monitor's global status
	// based on the status of the job in all regions. See NS1 API docs for supported values.
	Policy pulumi.StringPtrInput
	// If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
	RapidRecheck pulumi.BoolPtrInput
	// The list of region codes in which to run the monitoring
	// job. See NS1 API docs for supported values.
	Regions pulumi.StringArrayInput
	// A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
	Rules MonitoringJobRuleArrayInput
}

func (MonitoringJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringJobArgs)(nil)).Elem()
}

type MonitoringJobInput interface {
	pulumi.Input

	ToMonitoringJobOutput() MonitoringJobOutput
	ToMonitoringJobOutputWithContext(ctx context.Context) MonitoringJobOutput
}

func (*MonitoringJob) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringJob)(nil)).Elem()
}

func (i *MonitoringJob) ToMonitoringJobOutput() MonitoringJobOutput {
	return i.ToMonitoringJobOutputWithContext(context.Background())
}

func (i *MonitoringJob) ToMonitoringJobOutputWithContext(ctx context.Context) MonitoringJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobOutput)
}

func (i *MonitoringJob) ToOutput(ctx context.Context) pulumix.Output[*MonitoringJob] {
	return pulumix.Output[*MonitoringJob]{
		OutputState: i.ToMonitoringJobOutputWithContext(ctx).OutputState,
	}
}

// MonitoringJobArrayInput is an input type that accepts MonitoringJobArray and MonitoringJobArrayOutput values.
// You can construct a concrete instance of `MonitoringJobArrayInput` via:
//
//	MonitoringJobArray{ MonitoringJobArgs{...} }
type MonitoringJobArrayInput interface {
	pulumi.Input

	ToMonitoringJobArrayOutput() MonitoringJobArrayOutput
	ToMonitoringJobArrayOutputWithContext(context.Context) MonitoringJobArrayOutput
}

type MonitoringJobArray []MonitoringJobInput

func (MonitoringJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitoringJob)(nil)).Elem()
}

func (i MonitoringJobArray) ToMonitoringJobArrayOutput() MonitoringJobArrayOutput {
	return i.ToMonitoringJobArrayOutputWithContext(context.Background())
}

func (i MonitoringJobArray) ToMonitoringJobArrayOutputWithContext(ctx context.Context) MonitoringJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobArrayOutput)
}

func (i MonitoringJobArray) ToOutput(ctx context.Context) pulumix.Output[[]*MonitoringJob] {
	return pulumix.Output[[]*MonitoringJob]{
		OutputState: i.ToMonitoringJobArrayOutputWithContext(ctx).OutputState,
	}
}

// MonitoringJobMapInput is an input type that accepts MonitoringJobMap and MonitoringJobMapOutput values.
// You can construct a concrete instance of `MonitoringJobMapInput` via:
//
//	MonitoringJobMap{ "key": MonitoringJobArgs{...} }
type MonitoringJobMapInput interface {
	pulumi.Input

	ToMonitoringJobMapOutput() MonitoringJobMapOutput
	ToMonitoringJobMapOutputWithContext(context.Context) MonitoringJobMapOutput
}

type MonitoringJobMap map[string]MonitoringJobInput

func (MonitoringJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitoringJob)(nil)).Elem()
}

func (i MonitoringJobMap) ToMonitoringJobMapOutput() MonitoringJobMapOutput {
	return i.ToMonitoringJobMapOutputWithContext(context.Background())
}

func (i MonitoringJobMap) ToMonitoringJobMapOutputWithContext(ctx context.Context) MonitoringJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobMapOutput)
}

func (i MonitoringJobMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MonitoringJob] {
	return pulumix.Output[map[string]*MonitoringJob]{
		OutputState: i.ToMonitoringJobMapOutputWithContext(ctx).OutputState,
	}
}

type MonitoringJobOutput struct{ *pulumi.OutputState }

func (MonitoringJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringJob)(nil)).Elem()
}

func (o MonitoringJobOutput) ToMonitoringJobOutput() MonitoringJobOutput {
	return o
}

func (o MonitoringJobOutput) ToMonitoringJobOutputWithContext(ctx context.Context) MonitoringJobOutput {
	return o
}

func (o MonitoringJobOutput) ToOutput(ctx context.Context) pulumix.Output[*MonitoringJob] {
	return pulumix.Output[*MonitoringJob]{
		OutputState: o.OutputState,
	}
}

// Indicates if the job is active or temporarily disabled.
func (o MonitoringJobOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
func (o MonitoringJobOutput) Config() pulumi.MapOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.MapOutput { return v.Config }).(pulumi.MapOutput)
}

// The frequency, in seconds, at which to run the monitoring job in each region.
func (o MonitoringJobOutput) Frequency() pulumi.IntOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.IntOutput { return v.Frequency }).(pulumi.IntOutput)
}

// The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
func (o MonitoringJobOutput) JobType() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.StringOutput { return v.JobType }).(pulumi.StringOutput)
}

// turn off the notifications for the monitoring job.
func (o MonitoringJobOutput) Mute() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.BoolPtrOutput { return v.Mute }).(pulumi.BoolPtrOutput)
}

// The free-form display name for the monitoring job.
func (o MonitoringJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Freeform notes to be included in any notifications about this job.
func (o MonitoringJobOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The time in seconds after a failure to wait before sending a notification.
func (o MonitoringJobOutput) NotifyDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.IntPtrOutput { return v.NotifyDelay }).(pulumi.IntPtrOutput)
}

// If true, a notification is sent when a job returns to an "up" state.
func (o MonitoringJobOutput) NotifyFailback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.BoolPtrOutput { return v.NotifyFailback }).(pulumi.BoolPtrOutput)
}

func (o MonitoringJobOutput) NotifyList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.StringPtrOutput { return v.NotifyList }).(pulumi.StringPtrOutput)
}

// If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
func (o MonitoringJobOutput) NotifyRegional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.BoolPtrOutput { return v.NotifyRegional }).(pulumi.BoolPtrOutput)
}

// The time in seconds between repeat notifications of a failed job.
func (o MonitoringJobOutput) NotifyRepeat() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.IntPtrOutput { return v.NotifyRepeat }).(pulumi.IntPtrOutput)
}

// The policy for determining the monitor's global status
// based on the status of the job in all regions. See NS1 API docs for supported values.
func (o MonitoringJobOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.StringPtrOutput { return v.Policy }).(pulumi.StringPtrOutput)
}

// If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
func (o MonitoringJobOutput) RapidRecheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.BoolPtrOutput { return v.RapidRecheck }).(pulumi.BoolPtrOutput)
}

// The list of region codes in which to run the monitoring
// job. See NS1 API docs for supported values.
func (o MonitoringJobOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonitoringJob) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
func (o MonitoringJobOutput) Rules() MonitoringJobRuleArrayOutput {
	return o.ApplyT(func(v *MonitoringJob) MonitoringJobRuleArrayOutput { return v.Rules }).(MonitoringJobRuleArrayOutput)
}

type MonitoringJobArrayOutput struct{ *pulumi.OutputState }

func (MonitoringJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitoringJob)(nil)).Elem()
}

func (o MonitoringJobArrayOutput) ToMonitoringJobArrayOutput() MonitoringJobArrayOutput {
	return o
}

func (o MonitoringJobArrayOutput) ToMonitoringJobArrayOutputWithContext(ctx context.Context) MonitoringJobArrayOutput {
	return o
}

func (o MonitoringJobArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MonitoringJob] {
	return pulumix.Output[[]*MonitoringJob]{
		OutputState: o.OutputState,
	}
}

func (o MonitoringJobArrayOutput) Index(i pulumi.IntInput) MonitoringJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MonitoringJob {
		return vs[0].([]*MonitoringJob)[vs[1].(int)]
	}).(MonitoringJobOutput)
}

type MonitoringJobMapOutput struct{ *pulumi.OutputState }

func (MonitoringJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitoringJob)(nil)).Elem()
}

func (o MonitoringJobMapOutput) ToMonitoringJobMapOutput() MonitoringJobMapOutput {
	return o
}

func (o MonitoringJobMapOutput) ToMonitoringJobMapOutputWithContext(ctx context.Context) MonitoringJobMapOutput {
	return o
}

func (o MonitoringJobMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MonitoringJob] {
	return pulumix.Output[map[string]*MonitoringJob]{
		OutputState: o.OutputState,
	}
}

func (o MonitoringJobMapOutput) MapIndex(k pulumi.StringInput) MonitoringJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MonitoringJob {
		return vs[0].(map[string]*MonitoringJob)[vs[1].(string)]
	}).(MonitoringJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringJobInput)(nil)).Elem(), &MonitoringJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringJobArrayInput)(nil)).Elem(), MonitoringJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringJobMapInput)(nil)).Elem(), MonitoringJobMap{})
	pulumi.RegisterOutputType(MonitoringJobOutput{})
	pulumi.RegisterOutputType(MonitoringJobArrayOutput{})
	pulumi.RegisterOutputType(MonitoringJobMapOutput{})
}
