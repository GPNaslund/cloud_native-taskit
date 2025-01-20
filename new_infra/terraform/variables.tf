variable "default_image_id" {
  type        = string
  description = "The default image id"
  default     = "1291b7e9-5c28-4d1a-863e-9dc095c8cbc4"
}

variable "default_flavor_id" {
  type        = string
  description = "The default flavor id"
  default     = "c1-r1-d10"
}

variable "larger_flavor_id" {
  type = string
  description = "Larger compute instance flavor id"
  default = "c2-r2-d20"
}

variable "keypair" {
  type        = string
  description = "The keypair to associate with an instance"
  nullable    = false
}

variable "identity_file" {
  type = string
  description = "The path to the identity file to use for authentication"
  nullable = false
}

variable "default_external_network_id" {
  type = string
  description = "The default id of the external network"
  default = "fd401e50-9484-4883-9672-a2814089528c"
}

variable "default_availability_zone" {
  type = string
  description = "The default availability zone"
  default = "Education"
}
