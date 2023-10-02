variable "rg_name" {
  type        = string
  description = "The name of the resource group."
}

variable "location" {
  type        = string
  description = "The location/region where the resource group will be created."
}

variable "prefix" {
  type        = string
  description = "The prefix which should be used for all resources in this example"
}