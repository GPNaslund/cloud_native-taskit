# PROVISIONING OF MULTIPLE NODES USING LXD ON SINGLE INSTANCE SERVER

This provisioning requires that ansible is preinstalled on the server.




# SETUP NODES

# TAINT CONTROL NODE
kubectl taint nodes <your-control-plane-node-name> node-role.kubernetes.io/control-plane=:NoSchedule

# JOIN WORKERS

# ENABLE DASHBOARD

# CREATE INGRESS RULES TO EXPOSE THE DASHBOARD

# INSTALL GITLAB AGENT


