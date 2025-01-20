package main

import (
	"log"
	"taskit-infrastructure/instances"
	"taskit-infrastructure/instances/k8"
	"taskit-infrastructure/network"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		imageId := pulumi.String("c71f74af-bd88-4e7d-89d3-d7d3f9f3a0ea")
		networkID := pulumi.String("fd401e50-9484-4883-9672-a2814089528c")
		flavorId := pulumi.String("c1-r1-d10")
		availabilityZone := pulumi.String("Education")
		keyPair := pulumi.String("gn222gq-keypair")

		netStack, err := network.NewNetworkStack(ctx, &network.NetworkStackConfig{
			ExternalNetworkId: networkID,
		}, 2)
		if err != nil {
			return err
		}

		controlPlaneConfig := k8.NodeConfig{
			FlavorId:         flavorId,
			ImageId:          imageId,
			KeyPair:          keyPair,
			NetworkStack:     netStack,
			AvailabilityZone: availabilityZone,
		}
		controlPlaneNodes, err := k8.CreateK8ControlPlaneNodes(ctx, 3, controlPlaneConfig)
		if err != nil {
			log.Printf("Creating control plane nodes caused an error: %s", err.Error())
			return err
		}

		workerNodeConfig := k8.NodeConfig{
			FlavorId:         flavorId,
			ImageId:          imageId,
			KeyPair:          keyPair,
			NetworkStack:     netStack,
			AvailabilityZone: availabilityZone,
		}
		workerNodes, err := k8.CreateK8WorkerNodes(ctx, 3, workerNodeConfig)
		if err != nil {
			log.Printf("Creating worker nodes caused an error: %s", err.Error())
			return err
		}

		jumpHost, err := instances.NewJumpHost(ctx, &instances.JumpHostConfig{
			FlavorId:         flavorId,
			ImageId:          imageId,
			KeyPair:          keyPair,
			NetworkStack:     netStack,
			AssignFloating:   true,
			AvailabilityZone: availabilityZone,
		})
		if err != nil {
			return err
		}

		loadBalancer, err := instances.NewLoadBalancer(ctx, &instances.LoadBalancerConfig{
			FlavorId:         flavorId,
			ImageId:          imageId,
			KeyPair:          keyPair,
			NetworkStack:     netStack,
			AssignFloating:   true,
			AvailabilityZone: availabilityZone,
		})

		if err != nil {
			return err
		}

		network.AllowSSH(ctx, "jumphost", jumpHost.SecurityGroup.ID())
		network.AllowSSH(ctx, "loadbalancer", loadBalancer.SecurityGroup.ID())
		network.AllowHttp(ctx, "loadbalancer", loadBalancer.SecurityGroup.ID())
		network.AllowHttps(ctx, "loadbalancer", loadBalancer.SecurityGroup.ID())

		jumpHostCidr := jumpHost.Instance.Networks.Index(pulumi.Int(0)).FixedIpV4().ApplyT(func(ip *string) string {
			if ip == nil {
				return ""
			}
			return *ip + "/32"
		}).(pulumi.StringOutput)

		loadBalancerCidr := loadBalancer.Instance.Networks.Index(pulumi.Int(0)).FixedIpV4().ApplyT(func(ip *string) string {
			if ip == nil {
				return ""
			}
			return *ip + "/32"
		}).(pulumi.StringOutput)

		networkCidr := netStack.Subnet.Cidr

		network.AllowSSHFromJumpHost(ctx, "control", controlPlaneNodes.SecurityGroup.ID(), jumpHostCidr)
		network.AllowKubernetesAPIServerAccessFromLoadBalancer(ctx, "control", controlPlaneNodes.SecurityGroup.ID(), loadBalancerCidr)
		network.AllowEtcdCommunicationBetweenControlPlaneNodes(ctx, "control", controlPlaneNodes.SecurityGroup.ID(), networkCidr)
		network.AllowIntraClusterCommunication(ctx, "control", controlPlaneNodes.SecurityGroup.ID(), networkCidr)
		network.AllowClusterDNS(ctx, "control", controlPlaneNodes.SecurityGroup.ID(), networkCidr)
		network.AllowJoinServerCommunication(ctx, "worker", controlPlaneNodes.SecurityGroup.ID(), networkCidr)

		network.AllowSSHFromJumpHost(ctx, "worker", workerNodes.SecurityGroup.ID(), jumpHostCidr)
		network.AllowIntraClusterCommunication(ctx, "worker", workerNodes.SecurityGroup.ID(), networkCidr)
		network.AllowClusterDNS(ctx, "worker", workerNodes.SecurityGroup.ID(), networkCidr)
		network.AllowNodePortCommunication(ctx, "worker", workerNodes.SecurityGroup.ID(), networkCidr)
		network.AllowIngress(ctx, "worker", workerNodes.SecurityGroup.ID(), networkCidr)

		controlPlaneIds := []pulumi.StringInput{}
		for _, node := range controlPlaneNodes.Instances {
			controlPlaneIds = append(controlPlaneIds, node.Networks.Index(pulumi.Int(0)).FixedIpV4().ApplyT(func(ip *string) string {
				if ip == nil {
					log.Println("Failed to extract control plane ip")
					return ""
				}
				return *ip
			}).(pulumi.StringOutput))
		}

		workerNodeIds := []pulumi.StringInput{}
		for _, node := range workerNodes.Instances {
			workerNodeIds = append(workerNodeIds, node.Networks.Index(pulumi.Int(0)).FixedIpV4().ApplyT(func(ip *string) string {
				if ip == nil {
					log.Println("Failed to extract worker node ip")
					return ""
				}
				return *ip
			}).(pulumi.StringOutput))
		}

		ctx.Export("controlPlaneIPs", pulumi.StringArray(controlPlaneIds))
		ctx.Export("workerNodeIPs", pulumi.StringArray(workerNodeIds))
		ctx.Export("loadBalancerIP", netStack.FloatingIps[0].Address)
		ctx.Export("jumpHostIP", netStack.FloatingIps[1].Address)
		ctx.Export("metalLBStart", pulumi.String("192.168.1.200"))
		ctx.Export("metalLBEnd", pulumi.String("192.168.1.250"))

		return nil
	})
}
