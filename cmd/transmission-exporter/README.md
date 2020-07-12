# transmission-exporter

This is a Prometheus exporter for the Transmission torrent client.

## Installation

```
go get honnef.co/go/transmission/cmd/transmission-exporter
```

## Cardinality

This exporter makes the controversial choice of using the torrent info
hash as a label. For users with many thousands of torrents, this will
result in many more time series. This choice was made because it
allows interesting aggregations that are not otherwise possible, such
as aggregate download rates grouped by tracker. This choice was made
under the assumption that people won't be running thousands of
replicas of torrent clients and that the scrape rate will be fairly
low.
