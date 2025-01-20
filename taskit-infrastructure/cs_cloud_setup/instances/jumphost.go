package instances

import (
	"fmt"
	"taskit-infrastructure/network"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/compute"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/networking"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JumpHostConfig struct {
	FlavorId         pulumi.StringInput
	ImageId          pulumi.StringInput
	KeyPair          pulumi.StringInput
	NetworkStack     *network.NetworkStack
	AssignFloating   bool
	AvailabilityZone pulumi.StringInput
}

type JumpHost struct {
	Instance *compute.Instance
	SecurityGroup *networking.SecGroup
}

func NewJumpHost(ctx *pulumi.Context, config *JumpHostConfig) (*JumpHost, error) {
	secGroup, err := networking.NewSecGroup(ctx, "jump-host-sec-group", &networking.SecGroupArgs{
		Name:        pulumi.String("jump-host-sg"),
		Description: pulumi.String("Security group for the jump host"),
	})
	if err != nil {
		return nil, err
	}


	port, err := networking.NewPort(ctx, "jump-host-port", &networking.PortArgs{
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

	jumpHost, err := compute.NewInstance(ctx, "jump-host", &compute.InstanceArgs{
		FlavorId:         config.FlavorId,
		ImageId:          config.ImageId,
		KeyPair:          config.KeyPair,
		AvailabilityZone: config.AvailabilityZone,
		Networks: compute.InstanceNetworkArray{
			compute.InstanceNetworkArgs{
				Port: port.ID(),
			},
		},
	}, pulumi.DependsOn([]pulumi.Resource{
		config.NetworkStack.Subnet,
		config.NetworkStack.RouterInt,
		secGroup,
	}))

	if err != nil {
		return nil, err
	}

	if config.AssignFloating {
		_, err := networking.NewFloatingIpAssociate(ctx, "jump-host-floating-ip", &networking.FloatingIpAssociateArgs{
			PortId:     port.ID(),
			FloatingIp: config.NetworkStack.FloatingIps[1].Address,
		})
		if err != nil {
			return nil, err
		}
		ctx.Export("jumpHostFloatingIp", config.NetworkStack.FloatingIps[1].Address)
	}

	return &JumpHost{ Instance: jumpHost, SecurityGroup: secGroup }, nil
}
