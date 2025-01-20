resource "local_file" "ansible_inventory" {
  content = templatefile("${path.module}/templates/inventory.tmpl",
    {
      load_balancer_ip  = openstack_networking_floatingip_v2.load_balancer_floating_ip.address
      jump_host_ip      = openstack_networking_floatingip_v2.jump_host_floating_ip.address
      identity_file     = var.identity_file
    })
  filename = "../ansible/inventory.ini"
}


output "load_balancer_ip" {
  value = openstack_networking_floatingip_v2.load_balancer_floating_ip.address
}

output "jump_host_ip" {
  value = openstack_networking_floatingip_v2.jump_host_floating_ip.address
}

