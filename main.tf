#variable "data" {
#  default = {
#    "msg" = "Hello World"
#    "msg2" = [1, 2, 3, 4]
#    "env" = {
#       "test" = "test-value"
#    }
#  }
#}
#
#data "gotemplate" "gotmpl" {
#  template = "${path.module}/file.tmpl"
#  data = "${jsonencode(var.data)}"
#}
#

variable "env" {
  default = { 
     "name"=  "bxbd",
     "region" = "us-east-1" 
  }
}
variable "volume_map"  {
   default = {
      "source_volume" = "destination_volume"
   }
}

variable "data" {
   default = {
       env  = {
          "name"=  "test",
          "region" = "us-east-1"
     }
   }
}

locals {
  compute_data = { env = "${var.env}"}
}

data "gotemplate" "gotmpl" {
  template = "${file("${path.module}/defination.tmpl")}"
  data = "${jsonencode(local.compute_data)}"
}
output "tmpl" {
  value = "${data.gotemplate.gotmpl.rendered}"
}

output "compute_data" {
  value = "${local.compute_data}"
}

