package main

import (
    "fmt"
    "log"
    "time"

    "github.com/panjf2000/gnet/v2"
)

// MyEventHandler implements gnet.EventHandler (v2) 
type MyEventHandler struct{}

// OnBoot fires when the engine is ready to accept connections.
func (h *MyEventHandler) OnBoot(eng gnet.Engine) gnet.Action {
    log.Println("✅ Server is ready to accept connections")
    return gnet.None
}

// OnShutdown fires when the engine is shutting down.
func (h *MyEventHandler) OnShutdown(eng gnet.Engine) {
    log.Println("🛑 Server is shutting down")
}

// OnOpen fires when a new connection is opened.
func (h *MyEventHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    log.Printf("🔌 New connection from %s\n", c.RemoteAddr())
    return nil, gnet.None
}

// OnClose fires when a connection has closed.
func (h *MyEventHandler) OnClose(c gnet.Conn, err error) gnet.Action {
    log.Printf("🔒 Connection closed %s (err: %v)\n", c.RemoteAddr(), err)
    return gnet.None
}

// OnTraffic fires when data arrives.  We read all pending bytes, echo them back, then discard.
func (h *MyEventHandler) OnTraffic(c gnet.Conn) gnet.Action {
    // how many bytes are waiting?
    n := c.InboundBuffered()
    if n == 0 {
        return gnet.None
    }
    // grab them
    buf, _ := c.Next(n)
    msg := string(buf)
    log.Printf("📨 Received: %s\n", msg)

    // echo back asynchronously (nil callback)
    echo := []byte(fmt.Sprintf("Server received: %s", msg))
    c.AsyncWrite(echo, nil) // matches AsyncWrite(buf []byte, cb AsyncCallback) :contentReference[oaicite:3]{index=3}

    return gnet.None
}

// OnTick fires on your chosen interval; here we do nothing and reset to tick again in 10s.
func (h *MyEventHandler) OnTick() (delay time.Duration, action gnet.Action) {
    return 10 * time.Second, gnet.None
}

func main() {
    // Run the server on TCP port 9000, with multicore enabled
    if err := gnet.Run(
        &MyEventHandler{},
        "tcp://0.0.0.0:9000",           // listen on all interfaces:9000 :contentReference[oaicite:4]{index=4}
        gnet.WithMulticore(true),      // enable all CPU cores :contentReference[oaicite:5]{index=5}
    ); err != nil {
        log.Fatalf("❌ gnet.Run failed: %v", err)
    }
}
