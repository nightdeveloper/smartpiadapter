package devices

import (
	"github.com/nightdeveloper/smartpiadapter/structs"
	"github.com/nightdeveloper/smartpiadapter/logger"
	"github.com/nightdeveloper/smartpiadapter/interfaces"
	"net"
	"strings"
)

type IpInfoDevice struct {
	name string
	command string
	isCommand bool
	value string
	rc interfaces.ResultConverter
}


func (spd *IpInfoDevice) GetName() string {
	return spd.name;
}

func (spd *IpInfoDevice) SetProps(name string, command string, isCommand bool, rc interfaces.ResultConverter) {
	spd.name = name
	spd.command = command
	spd.isCommand = isCommand
	spd.rc = rc
}

func (spd *IpInfoDevice) Init() { }

func (spd *IpInfoDevice) GetStatus() structs.DeviceStatus {

	var ds structs.DeviceStatus
	ds.DeviceType = "IpInfoDevice"

	if !spd.isCommand {

		var extIp = "";

		ifaces, err := net.Interfaces()

		if (err != nil) {
			ds.Value, ds.Status = 0, "error while getting ips";
			logger.Error("error while getting ips", err)

		} else {
			for _, i := range ifaces {
				addrs, err := i.Addrs()

				if (err != nil) {
					logger.Error("error while getting ips", err)
					continue;
				}

				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP
					case *net.IPAddr:
						ip = v.IP
					}

					if (!ip.IsLoopback() && ip.String() != "127.0.0.1" &&
						strings.Index(ip.String(), ":") == -1) {
						extIp = extIp + ip.String() + " "
					}
				}
			}
			ds.Value, ds.Status = 1, extIp;
		}

	} else {
		logger.Error("command requests not implemented yet", nil)
	}

	return ds;
}




