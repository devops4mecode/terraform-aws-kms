<p align="center"> <img src="https://user-images.githubusercontent.com/50652676/62349836-882fef80-b51e-11e9-99e3-7b974309c7e3.png" width="100" height="100"></p>


<h1 align="center">
    Terraform AWS KMS
</h1>

<p align="center" style="font-size: 1.2rem;"> 
    This terraform module creates a KMS Customer Master Key (CMK) and its alias.
     </p>

<p align="center">

<a href="https://www.terraform.io">
  <img src="https://img.shields.io/badge/Terraform-v0.13-green" alt="Terraform">
</a>
<a href="LICENSE.md">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="Licence">
</a>


</p>
<p align="center">

<a href='https://facebook.com/sharer/sharer.php?u=https://github.com/devops4me/terraform-aws-kms'>
  <img title="Share on Facebook" src="https://user-images.githubusercontent.com/50652676/62817743-4f64cb80-bb59-11e9-90c7-b057252ded50.png" />
</a>
<a href='https://www.linkedin.com/shareArticle?mini=true&title=Terraform+AWS+KMS&url=https://github.com/devops4mecode/terraform-aws-kms'>
  <img title="Share on LinkedIn" src="https://user-images.githubusercontent.com/50652676/62817742-4e339e80-bb59-11e9-87b9-a1f68cae1049.png" />
</a>
<a href='https://twitter.com/intent/tweet/?text=Terraform+AWS+KMS&url=https://github.com/devops4mecode/terraform-aws-kms'>
  <img title="Share on Twitter" src="https://user-images.githubusercontent.com/50652676/62817740-4c69db00-bb59-11e9-8a79-3580fbbf6d5c.png" />
</a>

</p>
<hr>
## Prerequisites

This module has a few dependencies: 

- [Terraform 0.13](https://learn.hashicorp.com/terraform/getting-started/install.html)
- [Go](https://golang.org/doc/install)
- [github.com/stretchr/testify/assert](https://github.com/stretchr/testify)
- [github.com/gruntwork-io/terratest/modules/terraform](https://github.com/gruntwork-io/terratest)

## Examples


**IMPORTANT:** Since the `master` branch used in `source` varies based on new modifications, we suggest that you use the release versions [here](https://github.com/devops4mecode/terraform-aws-kms/releases).


### Simple Example
Here is an example of how you can use this module in your inventory structure:
```hcl
  module "kms_key" {
    source      = "devops4mecode/kms/aws"
    version     = "0.13.0"
    name        = "kms"
    application = "devops4me"
    environment = "test"
    label_order = ["environment", "application", "name"]
    enabled     = true
    description             = "KMS key for cloudtrail"
    deletion_window_in_days = 7
    enable_key_rotation     = true
    alias                   = "alias/cloudtrail"
    policy                  = data.aws_iam_policy_document.default.json
  }

  data "aws_iam_policy_document" "default" {
    version = "2012-10-17"
    statement {
      sid    = "Enable IAM User Permissions"
      effect = "Allow"
      principals {
        type        = "AWS"
        identifiers = ["*"]
      }
      actions   = ["kms:*"]
      resources = ["*"]
    }
    statement {
      sid    = "Allow CloudTrail to encrypt logs"
      effect = "Allow"
      principals {
        type        = "Service"
        identifiers = ["cloudtrail.amazonaws.com"]
      }
      actions   = ["kms:GenerateDataKey*"]
      resources = ["*"]
      condition {
        test     = "StringLike"
        variable = "kms:EncryptionContext:aws:cloudtrail:arn"
        values   = ["arn:aws:cloudtrail:*:XXXXXXXXXXXX:trail/*"]
      }
    }

    statement {
      sid    = "Allow CloudTrail to describe key"
      effect = "Allow"
      principals {
        type        = "Service"
        identifiers = ["cloudtrail.amazonaws.com"]
      }
      actions   = ["kms:DescribeKey"]
      resources = ["*"]
    }

    statement {
      sid    = "Allow principals in the account to decrypt log files"
      effect = "Allow"
      principals {
        type        = "AWS"
        identifiers = ["*"]
      }
      actions = [
        "kms:Decrypt",
        "kms:ReEncryptFrom"
      ]
      resources = ["*"]
      condition {
        test     = "StringEquals"
        variable = "kms:CallerAccount"
        values = [
        "XXXXXXXXXXXX"]
      }
      condition {
        test     = "StringLike"
        variable = "kms:EncryptionContext:aws:cloudtrail:arn"
        values   = ["arn:aws:cloudtrail:*:XXXXXXXXXXXX:trail/*"]
      }
    }

    statement {
      sid    = "Allow alias creation during setup"
      effect = "Allow"
      principals {
        type        = "AWS"
        identifiers = ["*"]
      }
      actions   = ["kms:CreateAlias"]
      resources = ["*"]
    }
  }

```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| alias | The display name of the alias. The name must start with the word `alias` followed by a forward slash. | `string` | `""` | no |
| application | Application (e.g. `do4m` or `devops4me`). | `string` | `""` | no |
| attributes | Additional attributes (e.g. `1`). | `list(string)` | `[]` | no |
| customer\_master\_key\_spec | Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports. Valid values: SYMMETRIC\_DEFAULT, RSA\_2048, RSA\_3072, RSA\_4096, ECC\_NIST\_P256, ECC\_NIST\_P384, ECC\_NIST\_P521, or ECC\_SECG\_P256K1. Defaults to SYMMETRIC\_DEFAULT. | `string` | `"SYMMETRIC_DEFAULT"` | no |
| deletion\_window\_in\_days | Duration in days after which the key is deleted after destruction of the resource. | `number` | `10` | no |
| description | The description of the key as viewed in AWS console. | `string` | `"Parameter Store KMS master key"` | no |
| enable\_key\_rotation | Specifies whether key rotation is enabled. | `bool` | `true` | no |
| enabled | Specifies whether the kms is enabled or disabled. | `bool` | `true` | no |
| environment | Environment (e.g. `prod`, `dev`, `staging`). | `string` | `""` | no |
| is\_enabled | Specifies whether the key is enabled. | `bool` | `true` | no |
| key\_usage | Specifies the intended use of the key. Defaults to ENCRYPT\_DECRYPT, and only symmetric encryption and decryption are supported. | `string` | `"ENCRYPT_DECRYPT"` | no |
| label\_order | label order, e.g. `name`,`application`. | `list` | `[]` | no |
| managedby | ManagedBy, eg 'DevOps4Me' or 'NajibRadzuan'. | `string` | `"najibradzuan@devops4me.com"` | no |
| name | Name  (e.g. `app` or `cluster`). | `string` | `""` | no |
| policy | A valid policy JSON document. For more information about building AWS IAM policy documents with Terraform. | `string` | `""` | no |
| tags | Additional tags (e.g. map(`BusinessUnit`,`XYZ`). | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| alias\_arn | Alias ARN. |
| alias\_name | Alias name. |
| key\_arn | Key ARN. |
| key\_id | Key ID. |
| tags | A mapping of tags to assign to the resource. |

## Testing
In this module testing is performed with [terratest](https://github.com/gruntwork-io/terratest) and it creates a small piece of infrastructure, matches the output like ARN, ID and Tags name etc and destroy infrastructure in your AWS account. This testing is written in GO, so you need a [GO environment](https://golang.org/doc/install) in your system. 

You need to run the following command in the testing folder:
```hcl
  go test -run Test
```