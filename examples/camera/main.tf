terraform {
  required_providers {
    reolink = {
      source = "registery.terraform.io/robbert229/reolink"
    }
  }
}

provider "reolink" {

}

# configure camera
resource "reolink_camera" "RLC-810A" {
  address = "192.168.103.221"
}
