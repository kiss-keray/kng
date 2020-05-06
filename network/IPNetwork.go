package network

// ip数据包转发口
type IPNetwork interface {
	// ip数据包流入处理
	PkgIn(ipPackage *IPPackage)
	// ip数据包流出处理
	PkgOut(ipPackage *IPPackage)
}

func InstanceFactory() IPNetwork {
	return *new(IPNetwork)
}
