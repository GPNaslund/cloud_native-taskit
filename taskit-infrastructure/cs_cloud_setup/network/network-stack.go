package network

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/networking"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkStackConfig struct {
	ExternalNetworkId pulumi.StringInput
}

type NetworkStack struct {
	Router        *networking.Router
	Network       *networking.Network
	Subnet        *networking.Subnet
	FloatingIps   []*networking.FloatingIp
	RouterInt     *networking.RouterInterface
	ExternalNetId pulumi.StringInput
}

func NewNetworkStack(ctx *pulumi.Context, config *NetworkStackConfig, numFloatingIps int) (*NetworkStack, error) {
	router, err := networking.NewRouter(ctx, "taskit-router", &networking.RouterArgs{
		Name:              pulumi.String("taskit-router"),
		AdminStateUp:      pulumi.Bool(true),
		Description:       pulumi.String("Router for taskit"),
		ExternalNetworkId: config.ExternalNetworkId,
	})
	if err != nil {
		log.Printf("Creating a router caused an error: %s", err.Error())
		return nil, err
	}

	network, err := networking.NewNetwork(ctx, "taskit-network", &networking.NetworkArgs{
		AdminStateUp: pulumi.Bool(true),
		Description:  pulumi.String("Network for taskit"),
		Name:         pulumi.String("taskit-network"),
	})
	if err != nil {
		log.Printf("Creating a new network caused an error: %s", err.Error())
		return nil, err
	}

	subnet, err := networking.NewSubnet(ctx, "taskit-subnet", &networking.SubnetArgs{
		Name:      pulumi.String("taskit-subnet"),
		NetworkId: network.ID(),
		IpVersion: pulumi.Int(4),
		Cidr:      pulumi.String("192.168.1.0/24"),
		GatewayIp: pulumi.String("192.168.1.1"),
	})
	if err != nil {
		log.Printf("Creating a subnet caused an error: %s", err.Error())
		return nil, err
	}

	routerInt, err := networking.NewRouterInterface(ctx, "taskit-router-interface", &networking.RouterInterfaceArgs{
		RouterId: router.ID(),
		SubnetId: subnet.ID(),
	})
	if err != nil {
		log.Printf("Creating a router interface caused an error: %s", err.Error())
		return nil, err
	}

	floatingIps := []*networking.FloatingIp{}
	for i := 0; i < numFloatingIps; i++ {
		floatingIpName := fmt.Sprintf("%s-floating-ip-%d", "taskit", i+1)
		floatingIp, err := networking.NewFloatingIp(ctx, floatingIpName, &networking.FloatingIpArgs{
			Pool: pulumi.String("public"),
		}, pulumi.DependsOn([]pulumi.Resource{ routerInt }))
		if err != nil {
			log.Printf("Creating a floating IP caused an error: %s", err.Error())
			return nil, err
		}
		floatingIps = append(floatingIps, floatingIp)
	}

	return &NetworkStack{
		Router:        router,
		Network:       network,
		Subnet:        subnet,
		FloatingIps:   floatingIps,
		RouterInt:     routerInt,
		ExternalNetId: config.ExternalNetworkId,
	}, nil
}
