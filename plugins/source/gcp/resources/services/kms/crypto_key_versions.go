package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func CryptoKeyVersions() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_crypto_key_versions",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion`,
		Resolver:    fetchCryptoKeyVersions,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.CryptoKeyVersion{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
