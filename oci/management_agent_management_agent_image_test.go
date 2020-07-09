// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managementAgentImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
	}

	ManagementAgentImageResourceConfig = ""
)

func TestManagementAgentManagementAgentImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentImageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_management_agent_management_agent_images.test_management_agent_images"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_images", "test_management_agent_images", Required, Create, managementAgentImageDataSourceRepresentation) +
					compartmentIdVariableStr + ManagementAgentImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.checksum"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.object_url"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.platform_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.platform_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.size"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.version"),
				),
			},
		},
	})
}
