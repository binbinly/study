package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func usage() {
	log.Printf("Usage: nats-echo [-s server] [-creds file] [-t] <subject>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[%d] Echoing to [%s]: %q", i, m.Reply, m.Data)
}

// We only want region, country
type geo struct {
	// There are others..
	Region  string
	Country string
}

func main()  {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var userCreds = flag.String("creds", "", "User Credentials File")
	var nkeyFile = flag.String("nkey", "", "NKey Seed File")
	var showTime = flag.Bool("t", false, "Display timestamps")
	var showHelp = flag.Bool("h", false, "Show help message")
	var geoloc = flag.Bool("geo", false, "Display geo location of echo service")

	var geo string
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	
	if *showHelp {
		showUsageAndExit(0)
	}
	
	args := flag.Args()
	fmt.Printf("len args:%v\n", len(args))
	fmt.Println("args", *urls, *showTime)
	if len(args) != 1 {
		//showUsageAndExit(1)
	}
	
	if *geoloc {
		geo = lookupGeo()
	}
	opts := []nats.Option{nats.Name("NATS Echo Service")}
	opts = setupConnOptions(opts)

	if *userCreds != "" && *nkeyFile != "" {
		log.Fatal("specify -seed or -creds")
	}

	if *userCreds != ""{
		opts = append(opts, nats.UserCredentials(*userCreds))
	}

	if *nkeyFile != "" {
		opt, err := nats.NkeyOptionFromSeed(*nkeyFile)
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, opt)
	}

	nc, err := nats.Connect(*urls, opts...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")

	subj, i := "nats-echo", 0

	nc.QueueSubscribe(subj, "echo", func(msg *nats.Msg) {
		i++
		if msg.Reply != "" {
			printMsg(msg, i)

			if geo != "" {
				m := fmt.Sprintf("[%s]: %q", geo, msg.Data)
				nc.Publish(msg.Reply, []byte(m))
			} else {
				nc.Publish(msg.Reply, msg.Data)
			}
		}
	})
	fmt.Println("flush before")
	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal("last err", err)
	}

	log.Printf("Echo Service listening on [%s]\n", subj)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	go func() {
		<-c
		log.Printf("<caught signal - draining>")
		nc.Drain()
	}()

	if *showTime {
		log.SetFlags(log.LstdFlags)
	}

	runtime.Goexit()
}

func lookupGeo() string {
	c := &http.Client{Timeout: 2 * time.Second}
	resp, err := c.Get("https://ipapi.co/json")
	if err != nil || resp == nil {
		log.Fatalf("Could not retrive geo location data: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	g := geo{}
	if err := json.Unmarshal(body, &g); err != nil {
		log.Fatalf("Error unmarshalling geo: %v", err)
	}
	fmt.Println("geo", string(body), g.Region + ", " + g.Country)
	return g.Region + ", " + g.Country
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
		if !conn.IsClosed() {
			log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
		}
	}))
	opts = append(opts, nats.ReconnectHandler(func(conn *nats.Conn) {
		log.Printf("Reconnected [%s[", conn.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(conn *nats.Conn) {
		if !conn.IsClosed() {
			log.Fatal("Exiting: no servers available")
		} else {
			log.Fatal("Exiting")
		}
	}))
	return opts
}