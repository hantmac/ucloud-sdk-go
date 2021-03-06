package umem

import (
	"encoding/base64"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateURedisGroupRequest is request schema for CreateURedisGroup action
type CreateURedisGroupRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 请求创建组的名称 (范围[6-63],只能包含英文、数字以及符号-和_)
	Name *string `required:"true"`

	//
	Protocol *string `required:"true"`

	// 是否开启高可用,enable或disable
	HighAvailability *string `required:"true"`

	// 每个节点的内存大小,单位GB,默认1GB,目前仅支持1/2/4/8/16/32,六种
	Size *int `required:"true"`

	// 配置ID,目前仅支持默认配置id 默认配置id:"03f58ca9-b64d-4bdd-abc7-c6b9a46fd801"
	ConfigId *string `required:"true"`

	// Redis版本信息(详见DescribeURedisVersion返回结果),默认版本3.0
	Version *string `required:"true"`

	// 计费模式，Year , Month, Dynamic 默认: Month
	ChargeType *string `required:"true"`

	// 业务组名称
	Tag *string `required:"true"`

	// 是否自动备份,enable或disable，默认disable
	AutoBackup *string `required:"false"`

	// 自动备份开始时间,范围[0-23],默认3点
	BackupTime *int `required:"false"`

	// 购买时长，默认为1
	Quantity *int `required:"false"`

	// 初始化密码
	Password *string `required:"false"`

	// 有此项代表从备份中创建，无代表正常创建
	BackupId *string `required:"false"`

	// 跨机房URedis，slave所在可用区（必须和Zone在同一Region，且不可相同）
	SlaveZone *string `required:"false"`

	// Master Redis Group的ID，创建只读Slave时，必须填写
	MasterGroupId *string `required:"false"`

	// 代金券ID
	CouponId *string `required:"false"`
}

// CreateURedisGroupResponse is response schema for CreateURedisGroup action
type CreateURedisGroupResponse struct {
	response.CommonBase

	// 创建的组ID
	GroupId string
}

// NewCreateURedisGroupRequest will create request of CreateURedisGroup action.
func (c *UMemClient) NewCreateURedisGroupRequest() *CreateURedisGroupRequest {
	req := &CreateURedisGroupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateURedisGroup - 创建主备redis
func (c *UMemClient) CreateURedisGroup(req *CreateURedisGroupRequest) (*CreateURedisGroupResponse, error) {
	var err error
	var res CreateURedisGroupResponse
	req.Password = ucloud.String(base64.StdEncoding.EncodeToString([]byte(ucloud.StringValue(req.Password))))

	err = c.client.InvokeAction("CreateURedisGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
