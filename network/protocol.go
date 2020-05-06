package network

import (
	"fmt"
)

type IPPackage struct {
	GetHeader *IPPackageHeader
	GetData   []byte
}

// IP数据包头部定义
type IPPackageHeader struct {
	// 版本号4+h首部长度4
	VersionLen int8
	// 区分服务
	Service int8
	//ip数据包总长度
	Length int16
	Logo   int16
	// 标志3+偏移量13
	TagOffset int16
	// 生存时间
	Ttl int8
	// 协议类型
	// 8806 tcp协议
	ProtocolType int8
	//校验和
	CheckSum int16
	//可选字段
	Other int32
}

//ip数据包源地址
func (ipHeader IPPackageHeader) GetSourceIp() []byte {
	panic("IPPackageHeader#GetSourceIp error")
}

// ip数据包目标地址
func (ipHeader IPPackageHeader) GetDestination() []byte {
	panic("IPPackageHeader#GetDestination error")
}

// IP数据包的版本
func (ipHeader IPPackageHeader) GetVersion() int8 {
	return ipHeader.VersionLen >> 4 & 0x0f
}

// ip数据包首长度
func (ipHeader IPPackageHeader) GetLen() int8 {
	return ipHeader.VersionLen & 0x0f
}

// ip数据标志
func (ipHeader IPPackageHeader) GetTag() int8 {
	return ipHeader.VersionLen >> 5 & 0x0f
}

// ip偏移量
func (ipHeader IPPackageHeader) GetOffset() int8 {
	return ipHeader.VersionLen & 0x1f
}

//IPv4数据包定义
type IPV4PackageHeader struct {
	IPPackageHeader
	//源地址
	SourceIP [4]byte
	// 目标地址
	Destination [4]byte
}

func (v4Header IPV4PackageHeader) GetSourceIp() []byte {
	s := make([]byte, 4)
	fmt.Println(copy(s, v4Header.SourceIP[:]))
	return s
}

func (v4Header IPV4PackageHeader) GetDestination() []byte {
	s := make([]byte, 4)
	fmt.Println(copy(s, v4Header.Destination[:]))
	return s
}
