terraform {
    required_version = ">= 1.5.0"

    required_providers {
        sbercloud = {
            source  = "sbercloud-terraform/sbercloud"
            version = "~> 1.12.0"
        }

        kubernetes = {
            source  = "hashicorp/kubernetes"
            version = "~> 2.23"
        }
        helm = {
            source  = "hashicorp/helm"
            version = "~> 2.11"
        }
        random = {
            source  = "hashicorp/random"
            version = "~> 3.5"
        }
        null = {
            source  = "hashicorp/null"
            version = "~> 3.2"
        }
    }
}
