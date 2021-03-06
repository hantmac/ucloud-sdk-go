//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DeleteUMemcacheGroup

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteUMemcacheGroupRequest is request schema for DeleteUMemcacheGroup action
type DeleteUMemcacheGroupRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 组ID
	GroupId *string `required:"true"`
}

// DeleteUMemcacheGroupResponse is response schema for DeleteUMemcacheGroup action
type DeleteUMemcacheGroupResponse struct {
	response.CommonBase
}

// NewDeleteUMemcacheGroupRequest will create request of DeleteUMemcacheGroup action.
func (c *UMemClient) NewDeleteUMemcacheGroupRequest() *DeleteUMemcacheGroupRequest {
	req := &DeleteUMemcacheGroupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteUMemcacheGroup - 删除单机Memcache
func (c *UMemClient) DeleteUMemcacheGroup(req *DeleteUMemcacheGroupRequest) (*DeleteUMemcacheGroupResponse, error) {
	var err error
	var res DeleteUMemcacheGroupResponse

	err = c.client.InvokeAction("DeleteUMemcacheGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
