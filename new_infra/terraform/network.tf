resource "openstack_networking_router_v2" "cluster_router" {
  name             = "cluster_router"
  description      = "The router for the cluster network"
  admin_state_up   = true
  external_network_id = var.default_external_network_id
}

resource "openstack_networking_network_v2" "cluster_network" {
  name           = "cluster_network"
  admin_state_up = true
}

resource "openstack_networking_subnet_v2" "cluster_subnet" {
  name       = "cluster_subnet"
  network_id = openstack_networking_network_v2.cluster_network.id
  cidr       = "192.168.1.0/24"
  ip_version = 4
}

resource "openstack_networking_router_interface_v2" "cluster_router_interface" {
  router_id = openstack_networking_router_v2.cluster_router.id
  subnet_id = openstack_networking_subnet_v2.cluster_subnet.id
}

resource "openstack_networking_floatingip_v2" "jump_host_floating_ip" {
  pool = "public"
}

resource "openstack_compute_floatingip_associate_v2" "jump_host_floating_ip_associate" {
  floating_ip = openstack_networking_floatingip_v2.jump_host_floating_ip.address
  instance_id = openstack_compute_instance_v2.jump_host.id
}

resource "openstack_networking_floatingip_v2" "load_balancer_floating_ip" {
  pool = "public"
}

resource "openstack_compute_floatingip_associate_v2" "load_balancer_floating_ip_associate" {
  floating_ip = openstack_networking_floatingip_v2.load_balancer_floating_ip.address
  instance_id = openstack_compute_instance_v2.load_balancer.id
}
