// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v27/core"
)

func init() {
	RegisterDatasource("oci_core_volume_group_backups", CoreVolumeGroupBackupsDataSource())
}

func CoreVolumeGroupBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVolumeGroupBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_group_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreVolumeGroupBackupResource()),
			},
		},
	}
}

func readCoreVolumeGroupBackups(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

type CoreVolumeGroupBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeGroupBackupsResponse
}

func (s *CoreVolumeGroupBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVolumeGroupBackupsDataSourceCrud) Get() error {
	request := oci_core.ListVolumeGroupBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if volumeGroupId, ok := s.D.GetOkExists("volume_group_id"); ok {
		tmp := volumeGroupId.(string)
		request.VolumeGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVolumeGroupBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeGroupBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVolumeGroupBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreVolumeGroupBackupsDataSource-", CoreVolumeGroupBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeGroupBackup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			volumeGroupBackup["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			volumeGroupBackup["display_name"] = *r.DisplayName
		}

		volumeGroupBackup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			volumeGroupBackup["id"] = *r.Id
		}

		if r.SizeInGBs != nil {
			volumeGroupBackup["size_in_gbs"] = strconv.FormatInt(*r.SizeInGBs, 10)
		}

		if r.SizeInMBs != nil {
			volumeGroupBackup["size_in_mbs"] = strconv.FormatInt(*r.SizeInMBs, 10)
		}

		volumeGroupBackup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			volumeGroupBackup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeRequestReceived != nil {
			volumeGroupBackup["time_request_received"] = r.TimeRequestReceived.String()
		}

		volumeGroupBackup["type"] = r.Type

		if r.UniqueSizeInGbs != nil {
			volumeGroupBackup["unique_size_in_gbs"] = strconv.FormatInt(*r.UniqueSizeInGbs, 10)
		}

		if r.UniqueSizeInMbs != nil {
			volumeGroupBackup["unique_size_in_mbs"] = strconv.FormatInt(*r.UniqueSizeInMbs, 10)
		}

		volumeGroupBackup["volume_backup_ids"] = r.VolumeBackupIds

		if r.VolumeGroupId != nil {
			volumeGroupBackup["volume_group_id"] = *r.VolumeGroupId
		}

		resources = append(resources, volumeGroupBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreVolumeGroupBackupsDataSource().Schema["volume_group_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_group_backups", resources); err != nil {
		return err
	}

	return nil
}
