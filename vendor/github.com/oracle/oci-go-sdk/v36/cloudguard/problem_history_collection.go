// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v36/common"
)

// ProblemHistoryCollection Collection of Problem History
type ProblemHistoryCollection struct {

	// List of ProblemHistorySummary
	Items []ProblemHistorySummary `mandatory:"true" json:"items"`
}

func (m ProblemHistoryCollection) String() string {
	return common.PointerString(m)
}