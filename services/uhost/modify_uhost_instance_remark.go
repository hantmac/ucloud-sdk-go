//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost ModifyUHostInstanceRemark

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ModifyUHostInstanceRemarkRequest is request schema for ModifyUHostInstanceRemark action
type ModifyUHostInstanceRemarkRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// UHost实例ID 参见 [DescribeUHostInstance](describe_uhost_instance.html)
	UHostId *string `required:"true"`

	// 备注
	Remark *string `required:"false"`
}

// ModifyUHostInstanceRemarkResponse is response schema for ModifyUHostInstanceRemark action
type ModifyUHostInstanceRemarkResponse struct {
	response.CommonBase

	// UHost实例ID
	UhostId string
}

// NewModifyUHostInstanceRemarkRequest will create request of ModifyUHostInstanceRemark action.
func (c *UHostClient) NewModifyUHostInstanceRemarkRequest() *ModifyUHostInstanceRemarkRequest {
	req := &ModifyUHostInstanceRemarkRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ModifyUHostInstanceRemark - 修改指定UHost实例备注信息。
func (c *UHostClient) ModifyUHostInstanceRemark(req *ModifyUHostInstanceRemarkRequest) (*ModifyUHostInstanceRemarkResponse, error) {
	var err error
	var res ModifyUHostInstanceRemarkResponse

	err = c.client.InvokeAction("ModifyUHostInstanceRemark", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}