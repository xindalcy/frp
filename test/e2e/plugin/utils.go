package plugin

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"

	plugin "github.com/xinda/desk/pkg/plugin/server"
	"github.com/xinda/desk/pkg/util/log"
	"github.com/xinda/desk/test/e2e/mock/server/httpserver"
)

type PluginHandler func(req *plugin.Request) *plugin.Response

type NewPluginRequest func() *plugin.Request

func NewHTTPPluginServer(port int, newFunc NewPluginRequest, handler PluginHandler, tlsConfig *tls.Config) *httpserver.Server {
	return httpserver.New(
		httpserver.WithBindPort(port),
		httpserver.WithTlsConfig(tlsConfig),
		httpserver.WithHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			r := newFunc()
			buf, err := ioutil.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(500)
				return
			}
			log.Trace("plugin request: %s", string(buf))
			err = json.Unmarshal(buf, &r)
			if err != nil {
				w.WriteHeader(500)
				return
			}
			resp := handler(r)
			buf, _ = json.Marshal(resp)
			log.Trace("plugin response: %s", string(buf))
			w.Write(buf)
		})),
	)
}
