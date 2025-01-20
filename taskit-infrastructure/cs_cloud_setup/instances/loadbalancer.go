package instances

import (
	"fmt"
	"taskit-infrastructure/network"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/compute"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/networking"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LoadBalancerConfig struct {
	FlavorId         pulumi.StringInput
	ImageId          pulumi.StringInput
	KeyPair          pulumi.StringInput
	NetworkStack     *network.NetworkStack
	AssignFloating   bool
	AvailabilityZone pulumi.StringInput
}

type LoadBalancer struct {
	Instance      *compute.Instance
	SecurityGroup *networking.SecGroup
}

func NewLoadBalancer(ctx *pulumi.Context, config *LoadBalancerConfig) (*LoadBalancer, error) {
	secGroup, err := networking.NewSecGroup(ctx, "jump-host-sg", &networking.SecGroupArgs{
		Name:        pulumi.String("jump-host-sg"),
		Description: pulumi.String("Security group for the load balancer"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create security group: %w", err)
	}

	port, err := networking.NewPort(ctx, "load-balancer-port", &networking.PortArgs{
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
		return nil, fmt.Errorf("failed to create port: %w", err)
	}

	lbInstance, err := compute.NewInstance(ctx, "load-balancer", &compute.InstanceArgs{
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
		return nil, fmt.Errorf("failed to create load balancer instance: %w", err)
	}

	if config.AssignFloating {
		_, err := networking.NewFloatingIpAssociate(ctx, "load-balancer-floating-ip", &networking.FloatingIpAssociateArgs{
			PortId:     port.ID(),
			FloatingIp: config.NetworkStack.FloatingIps[0].Address,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to associate floating IP: %w", err)
		}

		ctx.Export("loadBalancerFloatingIp", config.NetworkStack.FloatingIps[0].Address)
	}

	return &LoadBalancer{Instance: lbInstance, SecurityGroup: secGroup}, nil
}
