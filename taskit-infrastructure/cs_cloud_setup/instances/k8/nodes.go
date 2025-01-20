package k8

import (
	"fmt"
	"log"
	"taskit-infrastructure/network"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/compute"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/networking"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodeConfig struct {
	FlavorId         pulumi.StringInput
	ImageId          pulumi.StringInput
	KeyPair          pulumi.StringInput
	NetworkStack     *network.NetworkStack
	SecGroup         *networking.SecGroup
	AvailabilityZone pulumi.StringInput
}

type ControlPlaneNodesOutput struct {
	Instances     []*compute.Instance
	SecurityGroup *networking.SecGroup
}

func CreateK8ControlPlaneNodes(ctx *pulumi.Context, numberOfNodes int, config NodeConfig) (*ControlPlaneNodesOutput, error) {

	secGroup, err := networking.NewSecGroup(ctx, "control-plane-sg", &networking.SecGroupArgs{
		Name:        pulumi.String("control-plane-sg"),
		Description: pulumi.String("Security group for Kubernetes control plane nodes"),
	})
	if err != nil {
		log.Printf("Creating security group for Kubernetes control plane nodes caused error: %s", err.Error())
		return nil, err
	}

	instances := []*compute.Instance{}
	for i := 0; i < numberOfNodes; i++ {
		nodeName := fmt.Sprintf("control-plane-node-%d", i+1)
		portName := fmt.Sprintf("control-plane-port-%d", i+1)

		port, err := networking.NewPort(ctx, portName, &networking.PortArgs{
			NetworkId: config.NetworkStack.Network.ID(),
			FixedIps: networking.PortFixedIpArray{
				networking.PortFixedIpArgs{
					SubnetId: config.NetworkStack.Subnet.ID(),
				},
			},
			SecurityGroupIds: pulumi.StringArray{
				secGroup.ID(),
			},
		})
		if err != nil {
			return nil, err
		}

		instance, err := compute.NewInstance(ctx, nodeName, &compute.InstanceArgs{
			FlavorId:         config.FlavorId,
			ImageId:          config.ImageId,
			KeyPair:          config.KeyPair,
			AvailabilityZone: config.AvailabilityZone,
			Networks: compute.InstanceNetworkArray{
				compute.InstanceNetworkArgs{
					Port: port.ID(),
				},
			},
		})
		if err != nil {
			return nil, err
		}

		instances = append(instances, instance)
	}
	return &ControlPlaneNodesOutput{
		Instances:     instances,
		SecurityGroup: secGroup,
	}, nil
}

type WorkerNodesOutput struct {
	Instances     []*compute.Instance
	SecurityGroup *networking.SecGroup
}

func CreateK8WorkerNodes(ctx *pulumi.Context, numberOfNodes int, config NodeConfig) (*WorkerNodesOutput, error) {
	instances := []*compute.Instance{}
	ports := []*networking.Port{}

	secGroup, err := networking.NewSecGroup(ctx, "worker-node-sg", &networking.SecGroupArgs{
		Name:        pulumi.String("worker-node-sg"),
		Description: pulumi.String("Security group for Kubernetes worker nodes"),
	})
	if err != nil {
		log.Printf("Creating security group for Kubernetes worker nodes caused error: %s", err.Error())
		return nil, err
	}

	for i := 0; i < numberOfNodes; i++ {
		nodeName := fmt.Sprintf("worker-node-%d", i+1)
		portName := fmt.Sprintf("worker-node-port-%d", i+1)

		port, err := networking.NewPort(ctx, portName, &networking.PortArgs{
			NetworkId: config.NetworkStack.Network.ID(),
			FixedIps: networking.PortFixedIpArray{
				networking.PortFixedIpArgs{
					SubnetId: config.NetworkStack.Subnet.ID(),
				},
			},
			SecurityGroupIds: pulumi.StringArray{
				secGroup.ID(),
			},
		})

		ports = append(ports, port)

		if err != nil {
			log.Printf("Creating port for worker node %d caused an error: %s", i+1, err.Error())
			return nil, err
		}

		instanceArgs := &compute.InstanceArgs{
			FlavorId:         config.FlavorId,
			ImageId:          config.ImageId,
			KeyPair:          config.KeyPair,
			AvailabilityZone: config.AvailabilityZone,
			Networks: compute.InstanceNetworkArray{
				compute.InstanceNetworkArgs{
					Port: port.ID(),
				},
			},
		}

		instance, err := compute.NewInstance(ctx, nodeName, instanceArgs)
		if err != nil {
			log.Printf("Creating k8 worker node caused an error: %s", err.Error())
			return nil, err
		}
		instances = append(instances, instance)
	}

	return &WorkerNodesOutput{
		Instances:     instances,
		SecurityGroup: secGroup,
	}, nil
}
