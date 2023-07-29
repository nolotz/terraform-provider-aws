// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package vpclattice

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	vpclattice_sdkv2 "github.com/aws/aws-sdk-go-v2/service/vpclattice"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAuthPolicy,
			TypeName: "aws_vpclattice_auth_policy",
			Name:     "Auth Policy",
		},
		{
			Factory:  DataSourceListener,
			TypeName: "aws_vpclattice_listener",
			Name:     "Listener",
		},
		{
			Factory:  DataSourceResourcePolicy,
			TypeName: "aws_vpclattice_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  DataSourceService,
			TypeName: "aws_vpclattice_service",
		},
		{
			Factory:  DataSourceServiceNetwork,
			TypeName: "aws_vpclattice_service_network",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccessLogSubscription,
			TypeName: "aws_vpclattice_access_log_subscription",
			Name:     "Access Log Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceAuthPolicy,
			TypeName: "aws_vpclattice_auth_policy",
		},
		{
			Factory:  ResourceListener,
			TypeName: "aws_vpclattice_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceListenerRule,
			TypeName: "aws_vpclattice_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_vpclattice_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  ResourceService,
			TypeName: "aws_vpclattice_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceServiceNetwork,
			TypeName: "aws_vpclattice_service_network",
			Name:     "ServiceNetwork",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceServiceNetworkServiceAssociation,
			TypeName: "aws_vpclattice_service_network_service_association",
			Name:     "Service Network Service Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceServiceNetworkVPCAssociation,
			TypeName: "aws_vpclattice_service_network_vpc_association",
			Name:     "Service Network VPC Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTargetGroup,
			TypeName: "aws_vpclattice_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceTargetGroupAttachment,
			TypeName: "aws_vpclattice_target_group_attachment",
			Name:     "Target Group Attachment",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.VPCLattice
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*vpclattice_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return vpclattice_sdkv2.NewFromConfig(cfg, func(o *vpclattice_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = vpclattice_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
