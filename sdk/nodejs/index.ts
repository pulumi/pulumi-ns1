// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { APIKeyArgs, APIKeyState } from "./apikey";
export type APIKey = import("./apikey").APIKey;
export const APIKey: typeof import("./apikey").APIKey = null as any;
utilities.lazyLoad(exports, ["APIKey"], () => require("./apikey"));

export { ApplicationArgs, ApplicationState } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;
utilities.lazyLoad(exports, ["Application"], () => require("./application"));

export { DataFeedArgs, DataFeedState } from "./dataFeed";
export type DataFeed = import("./dataFeed").DataFeed;
export const DataFeed: typeof import("./dataFeed").DataFeed = null as any;
utilities.lazyLoad(exports, ["DataFeed"], () => require("./dataFeed"));

export { DataSourceArgs, DataSourceState } from "./dataSource";
export type DataSource = import("./dataSource").DataSource;
export const DataSource: typeof import("./dataSource").DataSource = null as any;
utilities.lazyLoad(exports, ["DataSource"], () => require("./dataSource"));

export { DnsviewArgs, DnsviewState } from "./dnsview";
export type Dnsview = import("./dnsview").Dnsview;
export const Dnsview: typeof import("./dnsview").Dnsview = null as any;
utilities.lazyLoad(exports, ["Dnsview"], () => require("./dnsview"));

export { GetDNSSecArgs, GetDNSSecResult, GetDNSSecOutputArgs } from "./getDNSSec";
export const getDNSSec: typeof import("./getDNSSec").getDNSSec = null as any;
export const getDNSSecOutput: typeof import("./getDNSSec").getDNSSecOutput = null as any;
utilities.lazyLoad(exports, ["getDNSSec","getDNSSecOutput"], () => require("./getDNSSec"));

export { GetRecordArgs, GetRecordResult, GetRecordOutputArgs } from "./getRecord";
export const getRecord: typeof import("./getRecord").getRecord = null as any;
export const getRecordOutput: typeof import("./getRecord").getRecordOutput = null as any;
utilities.lazyLoad(exports, ["getRecord","getRecordOutput"], () => require("./getRecord"));

export { GetZoneArgs, GetZoneResult, GetZoneOutputArgs } from "./getZone";
export const getZone: typeof import("./getZone").getZone = null as any;
export const getZoneOutput: typeof import("./getZone").getZoneOutput = null as any;
utilities.lazyLoad(exports, ["getZone","getZoneOutput"], () => require("./getZone"));

export { MonitoringJobArgs, MonitoringJobState } from "./monitoringJob";
export type MonitoringJob = import("./monitoringJob").MonitoringJob;
export const MonitoringJob: typeof import("./monitoringJob").MonitoringJob = null as any;
utilities.lazyLoad(exports, ["MonitoringJob"], () => require("./monitoringJob"));

export { NotifyListArgs, NotifyListState } from "./notifyList";
export type NotifyList = import("./notifyList").NotifyList;
export const NotifyList: typeof import("./notifyList").NotifyList = null as any;
utilities.lazyLoad(exports, ["NotifyList"], () => require("./notifyList"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { PulsarJobArgs, PulsarJobState } from "./pulsarJob";
export type PulsarJob = import("./pulsarJob").PulsarJob;
export const PulsarJob: typeof import("./pulsarJob").PulsarJob = null as any;
utilities.lazyLoad(exports, ["PulsarJob"], () => require("./pulsarJob"));

export { RecordArgs, RecordState } from "./record";
export type Record = import("./record").Record;
export const Record: typeof import("./record").Record = null as any;
utilities.lazyLoad(exports, ["Record"], () => require("./record"));

export { SubnetArgs, SubnetState } from "./subnet";
export type Subnet = import("./subnet").Subnet;
export const Subnet: typeof import("./subnet").Subnet = null as any;
utilities.lazyLoad(exports, ["Subnet"], () => require("./subnet"));

export { TeamArgs, TeamState } from "./team";
export type Team = import("./team").Team;
export const Team: typeof import("./team").Team = null as any;
utilities.lazyLoad(exports, ["Team"], () => require("./team"));

export { TsigkeyArgs, TsigkeyState } from "./tsigkey";
export type Tsigkey = import("./tsigkey").Tsigkey;
export const Tsigkey: typeof import("./tsigkey").Tsigkey = null as any;
utilities.lazyLoad(exports, ["Tsigkey"], () => require("./tsigkey"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));

export { ZoneArgs, ZoneState } from "./zone";
export type Zone = import("./zone").Zone;
export const Zone: typeof import("./zone").Zone = null as any;
utilities.lazyLoad(exports, ["Zone"], () => require("./zone"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ns1:index/aPIKey:APIKey":
                return new APIKey(name, <any>undefined, { urn })
            case "ns1:index/application:Application":
                return new Application(name, <any>undefined, { urn })
            case "ns1:index/dataFeed:DataFeed":
                return new DataFeed(name, <any>undefined, { urn })
            case "ns1:index/dataSource:DataSource":
                return new DataSource(name, <any>undefined, { urn })
            case "ns1:index/dnsview:Dnsview":
                return new Dnsview(name, <any>undefined, { urn })
            case "ns1:index/monitoringJob:MonitoringJob":
                return new MonitoringJob(name, <any>undefined, { urn })
            case "ns1:index/notifyList:NotifyList":
                return new NotifyList(name, <any>undefined, { urn })
            case "ns1:index/pulsarJob:PulsarJob":
                return new PulsarJob(name, <any>undefined, { urn })
            case "ns1:index/record:Record":
                return new Record(name, <any>undefined, { urn })
            case "ns1:index/subnet:Subnet":
                return new Subnet(name, <any>undefined, { urn })
            case "ns1:index/team:Team":
                return new Team(name, <any>undefined, { urn })
            case "ns1:index/tsigkey:Tsigkey":
                return new Tsigkey(name, <any>undefined, { urn })
            case "ns1:index/user:User":
                return new User(name, <any>undefined, { urn })
            case "ns1:index/zone:Zone":
                return new Zone(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ns1", "index/aPIKey", _module)
pulumi.runtime.registerResourceModule("ns1", "index/application", _module)
pulumi.runtime.registerResourceModule("ns1", "index/dataFeed", _module)
pulumi.runtime.registerResourceModule("ns1", "index/dataSource", _module)
pulumi.runtime.registerResourceModule("ns1", "index/dnsview", _module)
pulumi.runtime.registerResourceModule("ns1", "index/monitoringJob", _module)
pulumi.runtime.registerResourceModule("ns1", "index/notifyList", _module)
pulumi.runtime.registerResourceModule("ns1", "index/pulsarJob", _module)
pulumi.runtime.registerResourceModule("ns1", "index/record", _module)
pulumi.runtime.registerResourceModule("ns1", "index/subnet", _module)
pulumi.runtime.registerResourceModule("ns1", "index/team", _module)
pulumi.runtime.registerResourceModule("ns1", "index/tsigkey", _module)
pulumi.runtime.registerResourceModule("ns1", "index/user", _module)
pulumi.runtime.registerResourceModule("ns1", "index/zone", _module)
pulumi.runtime.registerResourcePackage("ns1", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:ns1") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
