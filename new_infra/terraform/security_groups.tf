resource "openstack_networking_secgroup_v2" "allow_ssh" {
  name        = "allow_ssh"
  description = "Allows ingress SSH"
}

resource "openstack_networking_secgroup_rule_v2" "ingress_ssh" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = openstack_networking_secgroup_v2.allow_ssh.id
}

resource "openstack_networking_secgroup_v2" "allow_http_https" {
  name        = "allow_http_https"
  description = "Allows http and https"
}

resource "openstack_networking_secgroup_rule_v2" "allow_http" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 80
  port_range_max    = 80
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = openstack_networking_secgroup_v2.allow_http_https.id
}

resource "openstack_networking_secgroup_rule_v2" "allow_https" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 443
  port_range_max    = 443
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = openstack_networking_secgroup_v2.allow_http_https.id
}

resource "openstack_networking_secgroup_v2" "allow_ssh_from_jumphost" {
  name = "allow_ssh_from_jumphost"
  description = "Allows ssh from jumphost"
}

resource "openstack_networking_secgroup_rule_v2" "ssh_from_jumphost" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = openstack_compute_instance_v2.jump_host.access_ip_v4
  security_group_id = openstack_networking_secgroup_v2.allow_ssh_from_jumphost.id
}

resource "openstack_networking_secgroup_v2" "k8_node_sg" {
  name        = "k8_node_sg"
  description = "Security group for Kubernetes nodes"
}


resource "openstack_networking_secgroup_rule_v2" "microk8s_api" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 6443
  port_range_max    = 6443
  remote_ip_prefix  = openstack_networking_subnet_v2.cluster_subnet.cidr
  security_group_id = openstack_networking_secgroup_v2.k8_node_sg.id
}

resource "openstack_networking_secgroup_rule_v2" "microk8s_worker" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 10250
  port_range_max    = 10250
  remote_ip_prefix  = openstack_networking_subnet_v2.cluster_subnet.cidr
  security_group_id = openstack_networking_secgroup_v2.k8_node_sg.id
}

resource "openstack_networking_secgroup_rule_v2" "microk8s_flannel" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "udp"
  port_range_min    = 8472
  port_range_max    = 8472
  remote_ip_prefix  = openstack_networking_subnet_v2.cluster_subnet.cidr
  security_group_id = openstack_networking_secgroup_v2.k8_node_sg.id
}

resource "openstack_networking_secgroup_rule_v2" "microk8s_nodeport" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 30000
  port_range_max    = 32767
  security_group_id = openstack_networking_secgroup_v2.k8_node_sg.id
}

resource "openstack_networking_secgroup_rule_v2" "microk8s_node_join" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 25000
  port_range_max    = 25000
  remote_ip_prefix  = openstack_networking_subnet_v2.cluster_subnet.cidr
  security_group_id = openstack_networking_secgroup_v2.k8_node_sg.id
}


resource "openstack_networking_secgroup_rule_v2" "microk8s_ha" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 16443
  port_range_max    = 16443
  remote_ip_prefix  = openstack_networking_subnet_v2.cluster_subnet.cidr
  security_group_id = openstack_networking_secgroup_v2.k8_node_sg.id
}
