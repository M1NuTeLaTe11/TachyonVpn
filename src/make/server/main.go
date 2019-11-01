package main

import (
	"github.com/tachyon-protocol/udw/udwConsole"
	tachyonVpnClient "tachyonVpnServer"
)

//kmg make sshDeploy -PkgPath make/server -Command server
//kmg make sshDeploy -PkgPath make/server -Command server -UseRelay -RelayServerIp [ip]
func main() {
	server := &tachyonVpnClient.Server{}
	udwConsole.MustRunCommandLineFromFuncV2(server.Run)
}
