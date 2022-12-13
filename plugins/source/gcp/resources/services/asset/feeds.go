// Code generated by codegen; DO NOT EDIT.

package asset

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Feeds() *schema.Table {
	return &schema.Table{
		Name:      "gcp_asset_feeds",
		Resolver:  fetchFeeds,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "asset_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AssetNames"),
			},
			{
				Name:     "asset_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AssetTypes"),
			},
			{
				Name:     "content_type",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("ContentType"),
			},
			{
				Name:     "feed_output_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FeedOutputConfig"),
			},
			{
				Name:     "condition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Condition"),
			},
			{
				Name:     "relationship_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RelationshipTypes"),
			},
		},
	}
}
