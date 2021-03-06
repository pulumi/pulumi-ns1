// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 Monitoring Job resource. This can be used to create, modify, and delete monitoring jobs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-ns1/sdk/v2/go/ns1"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ns1.NewMonitoringJob(ctx, "uswestMonitor", &ns1.MonitoringJobArgs{
// 			Active: pulumi.Bool(true),
// 			Config: pulumi.Map{
// 				"host": pulumi.String("example-elb-uswest.aws.amazon.com"),
// 				"port": pulumi.Float64(443),
// 				"send": pulumi.String(fmt.Sprintf("%v%v%v", "HEAD / HTTP/1.0\n", "\n", "\n")),
// 				"ssl": pulumi.Float64(1),
// 			},
// 			Frequency:    pulumi.Int(60),
// 			JobType:      pulumi.String("tcp"),
// 			Policy:       pulumi.String("quorum"),
// 			RapidRecheck: pulumi.Bool(true),
// 			Regions: pulumi.StringArray{
// 				pulumi.String("sjc"),
// 				pulumi.String("sin"),
// 				pulumi.String("lga"),
// 			},
// 			Rules: ns1.MonitoringJobRuleArray{
// 				&ns1.MonitoringJobRuleArgs{
// 					Comparison: pulumi.String("contains"),
// 					Key:        pulumi.String("output"),
// 					Value:      pulumi.String("200 OK"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## NS1 Documentation
//
// [MonitoringJob Api Doc](https://ns1.com/api#monitoring-jobs)
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
	return reflect.TypeOf((*MonitoringJob)(nil))
}

func (i *MonitoringJob) ToMonitoringJobOutput() MonitoringJobOutput {
	return i.ToMonitoringJobOutputWithContext(context.Background())
}

func (i *MonitoringJob) ToMonitoringJobOutputWithContext(ctx context.Context) MonitoringJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobOutput)
}

func (i *MonitoringJob) ToMonitoringJobPtrOutput() MonitoringJobPtrOutput {
	return i.ToMonitoringJobPtrOutputWithContext(context.Background())
}

func (i *MonitoringJob) ToMonitoringJobPtrOutputWithContext(ctx context.Context) MonitoringJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobPtrOutput)
}

type MonitoringJobPtrInput interface {
	pulumi.Input

	ToMonitoringJobPtrOutput() MonitoringJobPtrOutput
	ToMonitoringJobPtrOutputWithContext(ctx context.Context) MonitoringJobPtrOutput
}

type monitoringJobPtrType MonitoringJobArgs

func (*monitoringJobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringJob)(nil))
}

func (i *monitoringJobPtrType) ToMonitoringJobPtrOutput() MonitoringJobPtrOutput {
	return i.ToMonitoringJobPtrOutputWithContext(context.Background())
}

func (i *monitoringJobPtrType) ToMonitoringJobPtrOutputWithContext(ctx context.Context) MonitoringJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobPtrOutput)
}

// MonitoringJobArrayInput is an input type that accepts MonitoringJobArray and MonitoringJobArrayOutput values.
// You can construct a concrete instance of `MonitoringJobArrayInput` via:
//
//          MonitoringJobArray{ MonitoringJobArgs{...} }
type MonitoringJobArrayInput interface {
	pulumi.Input

	ToMonitoringJobArrayOutput() MonitoringJobArrayOutput
	ToMonitoringJobArrayOutputWithContext(context.Context) MonitoringJobArrayOutput
}

type MonitoringJobArray []MonitoringJobInput

func (MonitoringJobArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MonitoringJob)(nil))
}

func (i MonitoringJobArray) ToMonitoringJobArrayOutput() MonitoringJobArrayOutput {
	return i.ToMonitoringJobArrayOutputWithContext(context.Background())
}

func (i MonitoringJobArray) ToMonitoringJobArrayOutputWithContext(ctx context.Context) MonitoringJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobArrayOutput)
}

// MonitoringJobMapInput is an input type that accepts MonitoringJobMap and MonitoringJobMapOutput values.
// You can construct a concrete instance of `MonitoringJobMapInput` via:
//
//          MonitoringJobMap{ "key": MonitoringJobArgs{...} }
type MonitoringJobMapInput interface {
	pulumi.Input

	ToMonitoringJobMapOutput() MonitoringJobMapOutput
	ToMonitoringJobMapOutputWithContext(context.Context) MonitoringJobMapOutput
}

type MonitoringJobMap map[string]MonitoringJobInput

func (MonitoringJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MonitoringJob)(nil))
}

func (i MonitoringJobMap) ToMonitoringJobMapOutput() MonitoringJobMapOutput {
	return i.ToMonitoringJobMapOutputWithContext(context.Background())
}

func (i MonitoringJobMap) ToMonitoringJobMapOutputWithContext(ctx context.Context) MonitoringJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringJobMapOutput)
}

type MonitoringJobOutput struct {
	*pulumi.OutputState
}

func (MonitoringJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringJob)(nil))
}

func (o MonitoringJobOutput) ToMonitoringJobOutput() MonitoringJobOutput {
	return o
}

func (o MonitoringJobOutput) ToMonitoringJobOutputWithContext(ctx context.Context) MonitoringJobOutput {
	return o
}

func (o MonitoringJobOutput) ToMonitoringJobPtrOutput() MonitoringJobPtrOutput {
	return o.ToMonitoringJobPtrOutputWithContext(context.Background())
}

func (o MonitoringJobOutput) ToMonitoringJobPtrOutputWithContext(ctx context.Context) MonitoringJobPtrOutput {
	return o.ApplyT(func(v MonitoringJob) *MonitoringJob {
		return &v
	}).(MonitoringJobPtrOutput)
}

type MonitoringJobPtrOutput struct {
	*pulumi.OutputState
}

func (MonitoringJobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringJob)(nil))
}

func (o MonitoringJobPtrOutput) ToMonitoringJobPtrOutput() MonitoringJobPtrOutput {
	return o
}

func (o MonitoringJobPtrOutput) ToMonitoringJobPtrOutputWithContext(ctx context.Context) MonitoringJobPtrOutput {
	return o
}

type MonitoringJobArrayOutput struct{ *pulumi.OutputState }

func (MonitoringJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitoringJob)(nil))
}

func (o MonitoringJobArrayOutput) ToMonitoringJobArrayOutput() MonitoringJobArrayOutput {
	return o
}

func (o MonitoringJobArrayOutput) ToMonitoringJobArrayOutputWithContext(ctx context.Context) MonitoringJobArrayOutput {
	return o
}

func (o MonitoringJobArrayOutput) Index(i pulumi.IntInput) MonitoringJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitoringJob {
		return vs[0].([]MonitoringJob)[vs[1].(int)]
	}).(MonitoringJobOutput)
}

type MonitoringJobMapOutput struct{ *pulumi.OutputState }

func (MonitoringJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MonitoringJob)(nil))
}

func (o MonitoringJobMapOutput) ToMonitoringJobMapOutput() MonitoringJobMapOutput {
	return o
}

func (o MonitoringJobMapOutput) ToMonitoringJobMapOutputWithContext(ctx context.Context) MonitoringJobMapOutput {
	return o
}

func (o MonitoringJobMapOutput) MapIndex(k pulumi.StringInput) MonitoringJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MonitoringJob {
		return vs[0].(map[string]MonitoringJob)[vs[1].(string)]
	}).(MonitoringJobOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitoringJobOutput{})
	pulumi.RegisterOutputType(MonitoringJobPtrOutput{})
	pulumi.RegisterOutputType(MonitoringJobArrayOutput{})
	pulumi.RegisterOutputType(MonitoringJobMapOutput{})
}
