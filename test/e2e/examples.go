package e2e

import (
	"fmt"

	"github.com/xinda/desk/test/e2e/framework"
	"github.com/xinda/desk/test/e2e/framework/consts"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("[Feature: Example]", func() {
	f := framework.NewDefaultFramework()

	Describe("TCP", func() {
		It("Expose a TCP echo server", func() {
			serverConf := consts.DefaultServerConfig
			clientConf := consts.DefaultClientConfig

			remotePort := f.AllocPort()
			clientConf += fmt.Sprintf(`
			[tcp]
			type = tcp
			local_port = {{ .%s }}
			remote_port = %d
			`, framework.TCPEchoServerPort, remotePort)

			f.RunProcesses([]string{serverConf}, []string{clientConf})

			framework.NewRequestExpect(f).Port(remotePort).Ensure()
		})
	})
})
