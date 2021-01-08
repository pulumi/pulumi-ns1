CHANGELOG
=========

## HEAD (Unreleased)
_(none)_

---

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
