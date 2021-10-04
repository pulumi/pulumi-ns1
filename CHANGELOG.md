CHANGELOG
=========

## HEAD (Unreleased)
_(none)_

---

## 2.2.1 (2021-10-04)
* Upgrade to v1.12.1 of the NS1 Terraform Provider

## 2.2.0 (2021-09-27)
* Upgrade to v1.12.0 of the NS1 Terraform Provider

## 2.1.0 (2021-08-09)
* Upgrade to v1.11.0 of the NS1 Terraform Provider

## 2.0.3 (2021-06-28)
* Upgrade to v1.10.3 of the NS1 Terraform Provider

## 2.0.2 (2021-05-27)
* Upgrade to v1.10.2 of the NS1 Terraform Provider

## 2.0.1 (2021-04-30)
* Upgrade to v1.10.1 of the NS1 Terraform Provider

## 2.0.0 (2021-04-19)
* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.

## 1.6.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 1.5.2 (2021-04-05)
* Upgrade to v1.9.4 of the NS1 Terraform Provider

## 1.5.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 1.5.0 (2021-03-15)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 1.4.2 (2021-03-09)
* Upgrade to v1.9.3 of the NS1 Terraform Provider

## 1.4.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 1.4.0 (2021-02-01)
* Upgrade to pulumi-terraform-bridge v2.18.1

## 1.3.3 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 1.3.2 (2021-01-08)
* Upgrade to v1.9.1 of the NS1 Terraform Provider

## 1.3.1 (2021-01-05)
* Upgrade to pulumi-terraform-bridge v2.13.2
  * This adds support for import specific examples in documentation

## 1.3.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 1.2.0 (2020-09-08)
* Upgrade to v1.9.0 of the NS1 Terraform Provider

## 1.1.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python  

## 1.0.0 (2020-06-23)
* Initial creation of the provider against v1.8.3 of the NS1 Terraform Provider
