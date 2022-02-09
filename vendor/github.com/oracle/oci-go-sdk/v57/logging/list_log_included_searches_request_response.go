// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"net/http"
	"strings"
)

// ListLogIncludedSearchesRequest wrapper for the ListLogIncludedSearches operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/logging/ListLogIncludedSearches.go.html to see an example of how to use ListLogIncludedSearchesRequest.
type ListLogIncludedSearchesRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OCID of the LogIncludedSearch
	LogIncludedSearchId *string `mandatory:"false" contributesTo:"query" name:"logIncludedSearchId"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogIncludedSearchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLogIncludedSearchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogIncludedSearchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogIncludedSearchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogIncludedSearchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogIncludedSearchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogIncludedSearchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListLogIncludedSearchesSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogIncludedSearchesSortByEnumStringValues(), ",")))
	}
	if _, ok := mappingListLogIncludedSearchesSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogIncludedSearchesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogIncludedSearchesResponse wrapper for the ListLogIncludedSearches operation
type ListLogIncludedSearchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogIncludedSearchSummaryCollection instances
	LogIncludedSearchSummaryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogIncludedSearchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogIncludedSearchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogIncludedSearchesSortByEnum Enum with underlying type: string
type ListLogIncludedSearchesSortByEnum string

// Set of constants representing the allowable values for ListLogIncludedSearchesSortByEnum
const (
	ListLogIncludedSearchesSortByTimecreated ListLogIncludedSearchesSortByEnum = "timeCreated"
	ListLogIncludedSearchesSortByDisplayname ListLogIncludedSearchesSortByEnum = "displayName"
)

var mappingListLogIncludedSearchesSortByEnum = map[string]ListLogIncludedSearchesSortByEnum{
	"timeCreated": ListLogIncludedSearchesSortByTimecreated,
	"displayName": ListLogIncludedSearchesSortByDisplayname,
}

// GetListLogIncludedSearchesSortByEnumValues Enumerates the set of values for ListLogIncludedSearchesSortByEnum
func GetListLogIncludedSearchesSortByEnumValues() []ListLogIncludedSearchesSortByEnum {
	values := make([]ListLogIncludedSearchesSortByEnum, 0)
	for _, v := range mappingListLogIncludedSearchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogIncludedSearchesSortByEnumStringValues Enumerates the set of values in String for ListLogIncludedSearchesSortByEnum
func GetListLogIncludedSearchesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// ListLogIncludedSearchesSortOrderEnum Enum with underlying type: string
type ListLogIncludedSearchesSortOrderEnum string

// Set of constants representing the allowable values for ListLogIncludedSearchesSortOrderEnum
const (
	ListLogIncludedSearchesSortOrderAsc  ListLogIncludedSearchesSortOrderEnum = "ASC"
	ListLogIncludedSearchesSortOrderDesc ListLogIncludedSearchesSortOrderEnum = "DESC"
)

var mappingListLogIncludedSearchesSortOrderEnum = map[string]ListLogIncludedSearchesSortOrderEnum{
	"ASC":  ListLogIncludedSearchesSortOrderAsc,
	"DESC": ListLogIncludedSearchesSortOrderDesc,
}

// GetListLogIncludedSearchesSortOrderEnumValues Enumerates the set of values for ListLogIncludedSearchesSortOrderEnum
func GetListLogIncludedSearchesSortOrderEnumValues() []ListLogIncludedSearchesSortOrderEnum {
	values := make([]ListLogIncludedSearchesSortOrderEnum, 0)
	for _, v := range mappingListLogIncludedSearchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogIncludedSearchesSortOrderEnumStringValues Enumerates the set of values in String for ListLogIncludedSearchesSortOrderEnum
func GetListLogIncludedSearchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}