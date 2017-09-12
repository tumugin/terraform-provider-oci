// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreRouteTableTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreRouteTableTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
		resource "oci_core_virtual_network" "t" {
			compartment_id = "${var.compartment_id}"
			cidr_block = "10.0.0.0/16"
			display_name = "-tf-vcn"
		}
		resource "oci_core_internet_gateway" "internet-gateway1" {
			compartment_id = "${var.compartment_id}"
			vcn_id = "${oci_core_virtual_network.t.id}"
			display_name = "-tf-internet-gateway"
		}`

	s.ResourceName = "oci_core_route_table.t"
}

func (s *ResourceCoreRouteTableTestSuite) TestAccResourceCoreRouteTable_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						route_rules {
							cidr_block = "0.0.0.0/0"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
					}`,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
				),
			},
			// verify update
			{
				Config: s.Config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "-tf-route-table"
						route_rules {
							cidr_block = "10.0.0.0/8"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-route-table"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "10.0.0.0/8"),
				),
			},
		},
	})
}

func TestResourceCoreRouteTableTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreRouteTableTestSuite))
}
