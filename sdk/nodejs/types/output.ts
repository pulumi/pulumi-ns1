// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetDNSSecDelegation {
    /**
     * (Computed) List of Keys. Key is documented below.
     */
    dnskeys: outputs.GetDNSSecDelegationDnskey[];
    /**
     * (Computed) List of Keys. Key is documented below.
     */
    ds: outputs.GetDNSSecDelegationD[];
    /**
     * (Computed) TTL for the Keys (int).
     */
    ttl: number;
}

export interface GetDNSSecDelegationD {
    /**
     * (Computed) Algorithm of the key.
     */
    algorithm: string;
    /**
     * (Computed) Flags for the key.
     */
    flags: string;
    /**
     * (Computed) Protocol of the key.
     */
    protocol: string;
    /**
     * (Computed) Public key for the key.
     */
    publicKey: string;
}

export interface GetDNSSecDelegationDnskey {
    /**
     * (Computed) Algorithm of the key.
     */
    algorithm: string;
    /**
     * (Computed) Flags for the key.
     */
    flags: string;
    /**
     * (Computed) Protocol of the key.
     */
    protocol: string;
    /**
     * (Computed) Public key for the key.
     */
    publicKey: string;
}

export interface GetDNSSecKeys {
    /**
     * (Computed) List of Keys. Key is documented below.
     */
    dnskeys: outputs.GetDNSSecKeysDnskey[];
    /**
     * (Computed) TTL for the Keys (int).
     */
    ttl: number;
}

export interface GetDNSSecKeysDnskey {
    /**
     * (Computed) Algorithm of the key.
     */
    algorithm: string;
    /**
     * (Computed) Flags for the key.
     */
    flags: string;
    /**
     * (Computed) Protocol of the key.
     */
    protocol: string;
    /**
     * (Computed) Public key for the key.
     */
    publicKey: string;
}

export interface GetZoneSecondary {
    /**
     * IPv4 address of the secondary server.
     */
    ip: string;
    /**
     * List of network IDs (`int`) for which the zone should be made
     * available. Default is network 0, the primary NSONE Global Network.
     */
    networks: number[];
    /**
     * Whether we send `NOTIFY` messages to the secondary host
     * when the zone changes. Default `false`.
     */
    notify: boolean;
    /**
     * Port of the the secondary server. Default `53`.
     */
    port: number;
}

export interface MonitoringJobRule {
    /**
     * The comparison to perform on the the output.
     */
    comparison: string;
    /**
     * The output key.
     */
    key: string;
    /**
     * The value to compare to.
     */
    value: string;
}

export interface NotifyListNotification {
    /**
     * Configuration details for the given notifier type.
     */
    config: {[key: string]: any};
    /**
     * The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
     */
    type: string;
}

export interface RecordAnswer {
    /**
     * Space delimited string of RDATA fields dependent on the record type.
     */
    answer?: string;
    meta?: {[key: string]: any};
    /**
     * The region (Answer Group really) that this answer
     * belongs to. This should be one of the names specified in `regions`. Only a
     * single `region` per answer is currently supported. If you want an answer in
     * multiple regions, duplicating the answer (including metadata) is the correct
     * approach.
     * * ` meta` - (Optional) meta is supported at the `answer` level. Meta
     * is documented below.
     */
    region?: string;
}

export interface RecordFilter {
    /**
     * The filters' configuration. Simple key/value pairs
     * determined by the filter type.
     */
    config?: {[key: string]: any};
    /**
     * Determines whether the filter is applied in the
     * filter chain.
     */
    disabled?: boolean;
    /**
     * The type of filter.
     */
    filter: string;
}

export interface RecordRegion {
    meta?: {[key: string]: any};
    /**
     * Name of the region (or Answer Group).
     */
    name: string;
}

export interface TeamIpWhitelist {
    /**
     * The free form name of the team.
     */
    name: string;
    values: string[];
}

export interface ZoneSecondary {
    /**
     * IPv4 address of the secondary server.
     */
    ip: string;
    /**
     * - List of network IDs (`int`) for which the zone
     * should be made available. Default is network 0, the primary NSONE Global
     * Network. Normally, you should not have to worry about this.
     */
    networks: number[];
    /**
     * Whether we send `NOTIFY` messages to the secondary host
     * when the zone changes. Default `false`.
     */
    notify: boolean;
    /**
     * Port of the the secondary server. Default `53`.
     */
    port: number;
}
