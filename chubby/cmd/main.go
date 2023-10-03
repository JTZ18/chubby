// Adapted from Leto command interface:
// https://github.com/yongman/leto/blob/master/cmd/main.go

// Copyright (C) 2018 YanMing <yming0221@gmail.com>
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"chubby/chubby/config"
	"chubby/chubby/server"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	listen		string		// Server listen port.
	raftDir 	string		// Raft data directory.
	raftBind 	string		// Raft bus transport bind port.
	nodeId		string		// Node ID.
	join		string		// Address of existing cluster at which to join.
	inmem		bool		// If true, keep log and stable storage in memory.
)

func init() {
	flag.StringVar(&listen, "listen", ":5379", "server listen port")
	flag.StringVar(&raftDir, "raftdir", "./", "raft data directory")
	flag.StringVar(&raftBind, "raftbind", ":15379", "raft bus transport bind port")
	flag.StringVar(&nodeId, "id", "", "node id")
	flag.StringVar(&join, "join", "", "join to existing cluster at this address")
	flag.BoolVar(&inmem, "inmem", false, "log and stable storage in memory")
}

func main() {
	// Parse flags from command line.
	flag.Parse()

	var (
		c *config.Config
	)

	// Create new Chubby config.
	c = config.NewConfig(listen, raftDir, raftBind, nodeId, join, inmem)
	//fmt.Println(c)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Run the server.
	go server.Run(c)
	// Exit on signal.
	<-quitCh
}
