package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"honnef.co/go/transmission"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO uptime_seconds_total

const namespace = "transmission_"

var (
	telemetryAddr = flag.String("telemetry.addr", ":9742", "address for transmission exporter")
	metricsPath   = flag.String("telemetry.path", "/metrics", "URL path for surfacing collected metrics")

	transmissionRPC      = flag.String("transmission.rpc", "http://localhost:9091/transmission/rpc", "URL of Transmission RPC endpoint")
	transmissionUser     = flag.String("transmission.user", "", "Transmission username")
	transmissionPass     = flag.String("transmission.pass", "", "Transmission password")
	transmissionPassFile = flag.String("transmission.pass-file", "", "File to read Transmission password from")
)

type Collector struct {
	client *transmission.Client

	descTorrentDownloaded    *prometheus.Desc
	descTorrentUploaded      *prometheus.Desc
	descTorrentTrackers      *prometheus.Desc
	descTorrentStatus        *prometheus.Desc
	descTorrentSize          *prometheus.Desc
	descPeers                *prometheus.Desc
	descPeersUploadingTo     *prometheus.Desc
	descPeersDownloadingFrom *prometheus.Desc
}

func NewCollector(client *transmission.Client) *Collector {
	return &Collector{
		client: client,

		descTorrentDownloaded:    prometheus.NewDesc(namespace+"torrent_downloaded_bytes_total", "", []string{"torrent"}, nil),
		descTorrentUploaded:      prometheus.NewDesc(namespace+"torrent_uploaded_bytes_total", "", []string{"torrent"}, nil),
		descTorrentTrackers:      prometheus.NewDesc(namespace+"torrent_trackers", "", []string{"torrent", "tracker"}, nil),
		descTorrentStatus:        prometheus.NewDesc(namespace+"torrent_status", "", []string{"torrent", "status"}, nil),
		descTorrentSize:          prometheus.NewDesc(namespace+"torrent_size_bytes", "", []string{"torrent"}, nil),
		descPeers:                prometheus.NewDesc(namespace+"torrent_peers", "", []string{"torrent"}, nil),
		descPeersUploadingTo:     prometheus.NewDesc(namespace+"torrent_peers_uploading_to", "", []string{"torrent"}, nil),
		descPeersDownloadingFrom: prometheus.NewDesc(namespace+"torrent_peers_downloading_from", "", []string{"torrent"}, nil),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.descTorrentDownloaded
	ch <- c.descTorrentUploaded
	ch <- c.descTorrentTrackers
	ch <- c.descTorrentStatus
	ch <- c.descTorrentSize
	ch <- c.descPeers
	ch <- c.descPeersUploadingTo
	ch <- c.descPeersDownloadingFrom
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	res, err := c.client.TorrentInfo(nil, transmission.AllTorrentFields)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(prometheus.NewInvalidDesc(err), err)
	}

	for _, info := range res {
		for _, tr := range info.TrackerStats {
			ch <- prometheus.MustNewConstMetric(c.descTorrentTrackers, prometheus.GaugeValue, 1, info.Hash, tr.Host)
		}
		ch <- prometheus.MustNewConstMetric(c.descTorrentStatus, prometheus.GaugeValue, 1, info.Hash, info.Status.String())
		ch <- prometheus.MustNewConstMetric(c.descTorrentSize, prometheus.GaugeValue, float64(info.SizeWhenDone), info.Hash)
		ch <- prometheus.MustNewConstMetric(c.descTorrentUploaded, prometheus.CounterValue, float64(info.UploadedEver), info.Hash)
		ch <- prometheus.MustNewConstMetric(c.descTorrentDownloaded, prometheus.CounterValue, float64(info.DownloadedEver), info.Hash)
		ch <- prometheus.MustNewConstMetric(c.descPeersUploadingTo, prometheus.GaugeValue, float64(info.PeersGettingFromUs), info.Hash)
		ch <- prometheus.MustNewConstMetric(c.descPeersDownloadingFrom, prometheus.GaugeValue, float64(info.PeersSendingToUs), info.Hash)
		ch <- prometheus.MustNewConstMetric(c.descPeers, prometheus.GaugeValue, float64(info.PeersConnected), info.Hash)
	}
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if *transmissionPass != "" && *transmissionPassFile != "" {
		log.Fatal("shouldn't specify both -transmission.pass and -transmission.pass-file")
	}

	pass := *transmissionPass
	if *transmissionPassFile != "" {
		b, err := ioutil.ReadFile(*transmissionPassFile)
		if err != nil {
			log.Fatalf("couldn't read password: %s", err)
		}
		pass = string(bytes.TrimRight(b, "\n"))
	}

	cl := &transmission.Client{
		Client:   http.DefaultClient,
		Endpoint: *transmissionRPC,
		Username: *transmissionUser,
		Password: pass,
	}

	prometheus.MustRegister(NewCollector(cl))
	http.Handle(*metricsPath, promhttp.Handler())
	if err := http.ListenAndServe(*telemetryAddr, nil); err != nil {
		log.Fatal(err)
	}
}
