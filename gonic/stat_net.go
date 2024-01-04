package gonic

import (
	"encoding/json"
	"os"

	"github.com/ahuigo/goos-tools/cli/netstat"
	"github.com/ahuigo/goos-tools/nets"
	"github.com/gin-gonic/gin"
)

// NetStat: show net stat information like shell `netstat`
func NetStat(c *gin.Context) {
	conns, err := netstat.GetAllTcpConnections()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	nets, err := nets.GetStats("")
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	s := formatNetwork(conns, nets)
	c.Data(200, "text/html", s)
	// c.String(200, s)
}

func formatNetwork(conns []netstat.TcpConnection, nets nets.Stats) []byte {
	hostname, _:= os.Hostname()
	tcpCounts := map[string]int{}
	for _, conn := range conns {
		tcpCounts[conn.State]++
	}

	netsBytes, _ := json.MarshalIndent(nets, "", "  ")
	data := map[string]interface{}{
		"title":     "Netstat",
		"hostname":  hostname,
		"conns":    conns,
		"tcpCounts": tcpCounts,
		"nets": nets,
		"netsStr": string(netsBytes),
	}

	s, _ := render("tpl/netstat.tmpl", data)
	return s
}
