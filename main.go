package main

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Config struct {
	DevicesName string
	SnapLen     int32
}

func main() {
	var config = loadConfig()
	var res, err = checkConfig(config)
	if err != nil {
		panic(err)
	}
	fmt.Print("work")
	if res {
		// 开始抓包
		work(config)
	}
}

// 配置检查
func checkConfig(config *Config) (res bool, err error) {

	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return false, err
	}
	for _, d := range devices {
		if d.Name == config.DevicesName {
			return true, nil
		}
	}
	return false, errors.New(fmt.Sprintf("未找到%s网卡", config.DevicesName))
}

// 加载配置文件
func loadConfig() *Config {
	var config = new(Config)
	config.DevicesName = "en0"
	config.SnapLen = 65535
	return config
}

// pacp抓包
func work(config *Config) {
	//打开网络接口，抓取在线数据
	handle, err := pcap.OpenLive(config.DevicesName, config.SnapLen, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("pcap open live failed: %v", err)
		return
	}
	// 设置过滤器
	if err := handle.SetBPFFilter(getFilter()); err != nil {
		fmt.Printf("set bpf filter failed: %v", err)
		return
	}
	defer handle.Close()

	//var work = network.InstanceFactory()
	// 抓包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.NoCopy = true
	for packet := range packetSource.Packets() {
		if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
			fmt.Println("unexpected packet")
			continue
		}

		fmt.Printf("packet:%v\n", packet)
	}
}

//定义过滤器
func getFilter() string {
	filter := fmt.Sprintf("ip")
	return filter
}
