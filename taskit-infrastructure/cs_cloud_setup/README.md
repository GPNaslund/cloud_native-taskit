# TASKIT INFRASTRUCTURE
1) pulumi up
2) pulumi stack output ansibleInventory > ansible/inventory.ini
3) ansible-playbook -i ansible/inventory.ini ansible/playbooks/cluster.yml -vvvv
