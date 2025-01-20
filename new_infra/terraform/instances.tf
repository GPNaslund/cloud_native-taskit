resource "openstack_compute_instance_v2" "jump_host" {
  name            = "jump_host"
  image_id        = var.default_image_id
  flavor_id       = var.default_flavor_id
  key_pair        = var.keypair
  security_groups = [openstack_networking_secgroup_v2.allow_ssh.name]
  availability_zone = var.default_availability_zone

  network {
    uuid = openstack_networking_network_v2.cluster_network.id
  }
}

resource "openstack_compute_instance_v2" "load_balancer" {
  name            = "load_balancer"
  image_id        = var.default_image_id
  flavor_id       = var.default_flavor_id
  key_pair        = var.keypair
  security_groups = [openstack_networking_secgroup_v2.allow_http_https.name, openstack_networking_secgroup_v2.allow_ssh.name]
  availability_zone = var.default_availability_zone


  network {
    uuid = openstack_networking_network_v2.cluster_network.id
  }
}

# resource "openstack_compute_instance_v2" "k8_control_planes" {
#   count     = 3
#   name      = "k8_control_plane-${count.index}"
#   image_id  = var.default_image_id
#   flavor_id = var.larger_flavor_id
#   key_pair = var.keypair
#   security_groups = [openstack_networking_secgroup_v2.k8_node_sg.name, openstack_networking_secgroup_v2.allow_ssh_from_jumphost.name]
#   availability_zone = var.default_availability_zone


#   network {
#     uuid = openstack_networking_network_v2.cluster_network.id
#   }
# }

# resource "openstack_compute_instance_v2" "k8_worker_nodes" {
#   count = 3
#   name = "k8_worker_node-${count.index}"
#   image_id = var.default_image_id
#   flavor_id = var.default_flavor_id
#   key_pair = var.keypair
#   security_groups = [openstack_networking_secgroup_v2.k8_node_sg.name, openstack_networking_secgroup_v2.allow_ssh_from_jumphost.name]
#   availability_zone = var.default_availability_zone


#   network {
#     uuid = openstack_networking_network_v2.cluster_network.id
#   }
# }
