# TASKIT INFRASTRUCTURE

## PROVISIONING
### Cause of going with my own VPS
Initially my plan was to provision the infrastructure on cs_cloud, using pulumi and ansible. Since i was using different tools than we had been using in
the course, there was quite alot of time spent learning pulumi. When significant progress had been made with pulumi and ansible towards cscloud, the deadline for the assignment grew closer. I have allways had technical issues with cscloud, making it hard to provision on it. I get reoccuring VPN issues, where i cant connect at all or the VPN closes mid-action and ruins the current actions. I also had reoccuring issues with terraform + cscloud (hence the choice to use pulumi, also for the learning value), where provisioning of instances alot of times did not get its cloud-init with my authorized ssh key provided by the school. When i just had a couple of days left to get the cluster and application online, i choose to buy my own VPS due to cscloud being such a big hinderence to the assignment.

I have provided the provisioning i had going for cscloud using pulumi + ansible, it is not 100% done due to the change of provider, but its near and i provide it as a proof of previous setup.

### Current infrastructure
I run the cluster on a single machine VPS from hostinger, with 8 cpu cores 32GB RAM, 33TB bandwith and 400GB memory. I choose their biggest VPS available, to have enough resource to create multiple VM's to create a cluster as similar as possible that i was working with on cscloud.

I provisioned the VM's using LXD (https://canonical.com/lxd), and the cluster itself uses MicroK8s (https://microk8s.io/), due to the resource constraints. Since MicroK8s is a minimal, CNCF certified Kubernetes distribution, i assumed it falls under the term "launch kubernetes", since
microk8s itself runs kubectl and other default kubernetes tools.

## Application infrastructure
