package unet

/*
EIPAddrSet - DescribeShareBandwidth

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type EIPAddrSet struct {

	// 运营商信息, 枚举值为: Telecom 电信; Unicom: 联通; Duplet: 双线; Bgp: BGP; International: 国际.
	OperatorName string

	// 弹性IP地址
	IP string
}
