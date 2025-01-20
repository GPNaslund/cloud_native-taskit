package network

import (
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/networking"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func AllowHttp(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput) error {
	_, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-http", &networking.SecGroupRuleArgs{
		Description:     pulumi.String("Allow HTTP traffic"),
		Direction:       pulumi.String("ingress"),
		Ethertype:       pulumi.String("IPv4"),
		Protocol:        pulumi.String("tcp"),
		PortRangeMin:    pulumi.Int(80),
		PortRangeMax:    pulumi.Int(80),
		SecurityGroupId: securityGroupId,
		RemoteIpPrefix:  pulumi.String("0.0.0.0/0"),
	})
	return err
}

func AllowHttps(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput) error {
	_, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-https", &networking.SecGroupRuleArgs{
		Description:     pulumi.String("Allow HTTPS traffic"),
		Direction:       pulumi.String("ingress"),
		Ethertype:       pulumi.String("IPv4"),
		Protocol:        pulumi.String("tcp"),
		PortRangeMin:    pulumi.Int(443),
		PortRangeMax:    pulumi.Int(443),
		SecurityGroupId: securityGroupId,
		RemoteIpPrefix:  pulumi.String("0.0.0.0/0"),
	})
	return err
}

func AllowSSH(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput) error {
	_, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-ssh", &networking.SecGroupRuleArgs{
		Description:     pulumi.String("Allow SSH traffic"),
		Direction:       pulumi.String("ingress"),
		Ethertype:       pulumi.String("IPv4"),
		Protocol:        pulumi.String("tcp"),
		PortRangeMin:    pulumi.Int(22),
		PortRangeMax:    pulumi.Int(22),
		SecurityGroupId: securityGroupId,
		RemoteIpPrefix:  pulumi.String("0.0.0.0/0"),
	})

	return err
}

func AllowSSHFromJumpHost(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, jumpHostCidr pulumi.StringInput) error {
	_, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-ssh-from-jumphost", &networking.SecGroupRuleArgs{
		Direction:       pulumi.String("ingress"),
		Ethertype:       pulumi.String("IPv4"),
		Protocol:        pulumi.String("tcp"),
		PortRangeMin:    pulumi.Int(22),
		PortRangeMax:    pulumi.Int(22),
		SecurityGroupId: securityGroupId,
		RemoteIpPrefix:  jumpHostCidr,
	})
	return err
}

func AllowKubernetesAPIServerAccessFromLoadBalancer(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, loadBalancerCidr pulumi.StringInput) error {
	_, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-k8-api-from-load-balancer", &networking.SecGroupRuleArgs{
		Direction:       pulumi.String("ingress"),
		Ethertype:       pulumi.String("IPv4"),
		Protocol:        pulumi.String("tcp"),
		PortRangeMin:    pulumi.Int(16443),
		PortRangeMax:    pulumi.Int(16443),
		SecurityGroupId: securityGroupId,
		RemoteIpPrefix:  loadBalancerCidr,
	})
	return err
}

func AllowEtcdCommunicationBetweenControlPlaneNodes(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, networkCidr pulumi.StringInput) error {
	_, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-etcd-control-plane-nodes", &networking.SecGroupRuleArgs{
		Direction:       pulumi.String("ingress"),
		Ethertype:       pulumi.String("IPv4"),
		Protocol:        pulumi.String("tcp"),
		PortRangeMin:    pulumi.Int(2379),
		PortRangeMax:    pulumi.Int(2380),
		SecurityGroupId: securityGroupId,
		RemoteIpPrefix:  networkCidr,
	})
	return err
}

func AllowIntraClusterCommunication(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, networkCidr pulumi.StringInput) error {
	_, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-intra-cluster-communication", &networking.SecGroupRuleArgs{
		Direction:       pulumi.String("ingress"),
		Ethertype:       pulumi.String("IPv4"),
		SecurityGroupId: securityGroupId,
		RemoteIpPrefix:  networkCidr,
	})
	return err
}

func AllowClusterDNS(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, networkCidr pulumi.StringInput) error {
    _, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-cluster-dns", &networking.SecGroupRuleArgs{
        Direction:       pulumi.String("ingress"),
        Ethertype:       pulumi.String("IPv4"),
        Protocol:        pulumi.String("udp"),
        PortRangeMin:    pulumi.Int(53),
        PortRangeMax:    pulumi.Int(53),
        SecurityGroupId: securityGroupId,
        RemoteIpPrefix:  networkCidr,
    })
    if err != nil {
        return err
    }

    _, err = networking.NewSecGroupRule(ctx, namePrefix + "-allow-cluster-dns-tcp", &networking.SecGroupRuleArgs{
        Direction:       pulumi.String("ingress"),
        Ethertype:       pulumi.String("IPv4"),
        Protocol:        pulumi.String("tcp"),
        PortRangeMin:    pulumi.Int(53),
        PortRangeMax:    pulumi.Int(53),
        SecurityGroupId: securityGroupId,
        RemoteIpPrefix:  networkCidr,
    })
    return err
}


func AllowNodePortCommunication(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, networkCidr pulumi.StringInput) error {
    _, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-nodeport-communication", &networking.SecGroupRuleArgs{
        Direction:       pulumi.String("ingress"),
        Ethertype:       pulumi.String("IPv4"),
        Protocol:        pulumi.String("tcp"),
        PortRangeMin:    pulumi.Int(30000),
        PortRangeMax:    pulumi.Int(32767),
        SecurityGroupId: securityGroupId,
        RemoteIpPrefix:  networkCidr,
    })
    return err
}

func AllowIngress(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, networkCidr pulumi.StringInput) error {
    _, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-ingress-http", &networking.SecGroupRuleArgs{
        Direction:       pulumi.String("ingress"),
        Ethertype:       pulumi.String("IPv4"),
        Protocol:        pulumi.String("tcp"),
        PortRangeMin:    pulumi.Int(80),
        PortRangeMax:    pulumi.Int(80),
        SecurityGroupId: securityGroupId,
        RemoteIpPrefix:  networkCidr,
    })
    if err != nil {
        return err
    }

    _, err = networking.NewSecGroupRule(ctx, namePrefix + "-allow-ingress-https", &networking.SecGroupRuleArgs{
        Direction:       pulumi.String("ingress"),
        Ethertype:       pulumi.String("IPv4"),
        Protocol:        pulumi.String("tcp"),
        PortRangeMin:    pulumi.Int(443),
        PortRangeMax:    pulumi.Int(443),
        SecurityGroupId: securityGroupId,
        RemoteIpPrefix:  networkCidr,
    })
    return err
}

func AllowJoinServerCommunication(ctx *pulumi.Context, namePrefix string, securityGroupId pulumi.StringInput, networkCidr pulumi.StringInput) error {
    _, err := networking.NewSecGroupRule(ctx, namePrefix + "-allow-join-server", &networking.SecGroupRuleArgs{
        Direction:       pulumi.String("ingress"),
        Ethertype:      pulumi.String("IPv4"),
        Protocol:       pulumi.String("tcp"),
        PortRangeMin:   pulumi.Int(8080),
        PortRangeMax:   pulumi.Int(8080),
        SecurityGroupId: securityGroupId,
        RemoteIpPrefix: networkCidr,
    })
    return err
}
