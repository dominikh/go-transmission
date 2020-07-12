package transmission

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TrackerState int

const (
	TrackerInactive TrackerState = 0
	TrackerWaiting  TrackerState = 1
	TrackerQueued   TrackerState = 2
	TrackerActive   TrackerState = 3
)

func (ts TrackerState) String() string {
	switch ts {
	case TrackerInactive:
		return "inactive"
	case TrackerWaiting:
		return "waiting"
	case TrackerQueued:
		return "queued"
	case TrackerActive:
		return "active"
	default:
		return fmt.Sprintf("TrackerState(%d)", ts)
	}
}

type Priority int

const (
	PriorityLow    Priority = -1
	PriorityNormal Priority = 0
	PriorityHigh   Priority = 1
)

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "low"
	case PriorityNormal:
		return "normal"
	case PriorityHigh:
		return "high"
	default:
		return fmt.Sprintf("Priority(%d)", p)
	}
}

const csrfHeader = "X-Transmission-Session-Id"

const (
	MethodTorrentStart       = "torrent-start"
	MethodTorrentStartNow    = "torrent-start-now"
	MethodTorrentStop        = "torrent-stop"
	MethodTorrentVerify      = "torrent-verify"
	MethodTorrentReannounce  = "torrent-reannounce"
	MethodTorrentSet         = "torrent-set"
	MethodTorrentGet         = "torrent-get"
	MethodTorrentAdd         = "torrent-add"
	MethodTorrentRemove      = "torrent-remove"
	MethodTorrentSetLocation = "torrent-set-location"
	MethodTorrentRenamePath  = "torrent-rename-path"
	MethodSessionSet         = "session-set"
	MethodSessionGet         = "session-get"
	MethodSessionStats       = "session-stats"
	MethodBlocklistUpdate    = "blocklist-update"
	MethodPortTest           = "port-test"
	MethodSessionClose       = "session-close"
	MethodQueueMoveTop       = "queue-move-top"
	MethodQueueMoveUp        = "queue-move-up"
	MethodQueueMoveDown      = "queue-move-down"
	MethodQueueMoveBottom    = "queue-move-bottom"
	MethodFreeSpace          = "free-space"
)

type TorrentStatus int

const (
	TorrentStatusStopped      TorrentStatus = 0
	TorrentStatusCheckWait    TorrentStatus = 1
	TorrentStatusCheck        TorrentStatus = 2
	TorrentStatusDownloadWait TorrentStatus = 3
	TorrentStatusDownload     TorrentStatus = 4
	TorrentStatusSeedWait     TorrentStatus = 5
	TorrentStatusSeed         TorrentStatus = 6
)

func (s TorrentStatus) String() string {
	switch s {
	case TorrentStatusStopped:
		return "stopped"
	case TorrentStatusCheckWait:
		return "queued to check"
	case TorrentStatusCheck:
		return "checking"
	case TorrentStatusDownloadWait:
		return "queued to download"
	case TorrentStatusDownload:
		return "downloading"
	case TorrentStatusSeedWait:
		return "queued to seed"
	case TorrentStatusSeed:
		return "seeding"
	default:
		return fmt.Sprintf("TorrentStatus(%d)", s)
	}
}

var AllTorrentFields = []string{
	"activityDate", "addedDate", "bandwidthPriority", "comment", "corruptEver", "creator", "dateCreated",
	"desiredAvailable", "doneDate", "downloadDir", "downloadLimit", "downloadLimited", "downloadedEver",
	"editDate", "error", "errorString", "eta", "etaIdle", "fileStats", "files",
	"hashString", "haveUnchecked", "haveValid", "honorsSessionLimits", "id", "isFinished", "isPrivate",
	"isStalled", "labels", "leftUntilDone", "magnetLink", "manualAnnounceTime", "maxConnectedPeers",
	"metadataPercentComplete", "name", "peer-limit", "peers", "peersConnected", "peersFrom",
	"peersGettingFromUs", "peersSendingToUs", "percentDone", "pieceCount", "pieceSize", "pieces",
	"priorities", "queuePosition", "rateDownload (B/s)", "rateUpload (B/s)", "recheckProgress",
	"secondsDownloading", "secondsSeeding", "seedIdleLimit", "seedIdleMode", "seedRatioLimit", "seedRatioMode",
	"sizeWhenDone", "startDate", "status", "torrentFile", "totalSize", "trackerStats",
	"trackers", "uploadLimit", "uploadLimited", "uploadRatio", "uploadedEver", "wanted",
	"webseeds", "webseedsSendingToUs",
}

var AllSessionFields = []string{
	"alt-speed-down", "alt-speed-enabled", "alt-speed-time-begin", "alt-speed-time-enabled", "alt-speed-time-end",
	"alt-speed-time-day", "alt-speed-up", "blocklist-url", "blocklist-enabled", "blocklist-size", "cache-size-mb",
	"config-dir", "download-dir", "download-queue-size", "download-queue-enabled", "dht-enabled", "encryption",
	"idle-seeding-limit", "idle-seeding-limit-enabled", "incomplete-dir", "incomplete-dir-enabled", "lpd-enabled",
	"peer-limit-global", "peer-limit-per-torrent", "pex-enabled", "peer-port", "peer-port-random-on-start",
	"port-forwarding-enabled", "queue-stalled-enabled", "queue-stalled-minutes", "rename-partial-files", "rpc-version",
	"rpc-version-minimum", "script-torrent-done-filename", "script-torrent-done-enabled", "seedRatioLimit",
	"seedRatioLimited", "seed-queue-size", "seed-queue-enabled", "speed-limit-down", "speed-limit-down-enabled",
	"speed-limit-up", "speed-limit-up-enabled", "start-added-torrents", "trash-original-torrent-files", "units",
	"utp-enabled", "version",
}

type Client struct {
	Client   *http.Client
	Endpoint string
	Username string
	Password string
	csrf     string
}

func NewClient(endpoint string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		Client:   client,
		Endpoint: endpoint,
	}
}

func (cl *Client) Request(method string, args interface{}) (Response, error) {
	type request struct {
		Method    string      `json:"method"`
		Arguments interface{} `json:"arguments"`
	}

	b, err := json.Marshal(request{method, args})
	if err != nil {
		return Response{}, err
	}
	hreq, err := http.NewRequest(http.MethodPost, cl.Endpoint, bytes.NewReader(b))
	if err != nil {
		return Response{}, err
	}
	hreq.Header.Set(csrfHeader, cl.csrf)
	if cl.Username != "" {
		hreq.SetBasicAuth(cl.Username, cl.Password)
	}

	resp, err := cl.Client.Do(hreq)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusConflict {
		// XXX don't get stuck in a loop
		cl.csrf = resp.Header.Get(csrfHeader)
		return cl.Request(method, args)
	}
	if resp.StatusCode/100 != 2 {
		return Response{},
			fmt.Errorf("request %s to %s failed: %s", method, cl.Endpoint, http.StatusText(resp.StatusCode))
	}

	var out Response
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return Response{}, err
	}
	if out.Result != "success" {
		return out, fmt.Errorf("request %s to %s failed: %s", method, cl.Endpoint, out.Result)
	}
	return out, nil
}

func (cl *Client) SessionStats() (*SessionStats, error) {
	resp, err := cl.Request(MethodSessionStats, nil)
	if err != nil {
		return nil, err
	}

	var out SessionStats
	err = json.Unmarshal([]byte(*resp.Arguments), &out)
	return &out, err
}

func (cl *Client) StartTorrent(ids []string) error {
	return cl.torrentAction(MethodTorrentStart, ids)
}

func (cl *Client) StartTorrentNow(ids []string) error {
	return cl.torrentAction(MethodTorrentStartNow, ids)
}

func (cl *Client) StopTorrent(ids []string) error {
	return cl.torrentAction(MethodTorrentStop, ids)
}

func (cl *Client) VerifyTorrent(ids []string) error {
	return cl.torrentAction(MethodTorrentVerify, ids)
}

func (cl *Client) ReannounceTorrent(ids []string) error {
	return cl.torrentAction(MethodTorrentReannounce, ids)
}

func (cl *Client) torrentAction(method string, ids []string) error {
	_, err := cl.Request(method, struct {
		IDs []string `json:"ids"`
	}{ids})
	return err
}

// BUG(dh): it is impossible to add a torrent with a peer limit of 0.
func (cl *Client) AddTorrent(torrent *NewTorrent) (info AddedTorrent, duplicate bool, err error) {
	resp, err := cl.Request(MethodTorrentAdd, torrent)
	if err != nil {
		return AddedTorrent{}, false, err
	}

	type fields struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Hash string `json:"hashString"`
	}
	var body struct {
		TorrentAdded     fields `json:"torrent-added"`
		TorrentDuplicate fields `json:"torrent-duplicate"`
	}
	if err := json.Unmarshal([]byte(*resp.Arguments), &body); err != nil {
		return AddedTorrent{}, false, err
	}

	if body.TorrentAdded.Name != "" {
		return AddedTorrent{
			ID:   body.TorrentAdded.ID,
			Name: body.TorrentAdded.Name,
			Hash: body.TorrentAdded.Hash,
		}, false, nil
	} else {
		return AddedTorrent{
			ID:   body.TorrentDuplicate.ID,
			Name: body.TorrentDuplicate.Name,
			Hash: body.TorrentDuplicate.Hash,
		}, true, nil
	}
}

func (cl *Client) TorrentInfo(ids []string, fields []string) ([]TorrentInfo, error) {
	req := struct {
		IDs    []string `json:"ids,omitempty"`
		Fields []string `json:"fields"`
		Format string   `json:"format"`
	}{
		ids,
		fields,
		"objects",
	}
	resp, err := cl.Request(MethodTorrentGet, req)
	if err != nil {
		return nil, err
	}

	var infos struct {
		Torrents []torrentInfo `json:"torrents"`
	}
	if err := json.Unmarshal([]byte(*resp.Arguments), &infos); err != nil {
		return nil, err
	}

	out := make([]TorrentInfo, len(infos.Torrents))
	for i := range out {
		convertTorrentInfo(&infos.Torrents[i], &out[i])
	}
	return out, nil
}

func (cl *Client) RemoveTorrent(ids []string, deleteLocalData bool) error {
	_, err := cl.Request(MethodTorrentRemove, struct {
		IDs             []string `json:"ids"`
		DeleteLocalData bool     `json:"delete-local-data"`
	}{ids, deleteLocalData})
	return err
}

func (cl *Client) MoveTorrent(ids []string, location string, move bool) error {
	_, err := cl.Request(MethodTorrentSetLocation, struct {
		IDs      []string `json:"ids"`
		Location string   `json:"location"`
		Move     bool     `json:"move"`
	}{ids, location, move})
	return err
}

func (cl *Client) RenameTorrentPath(ids []string, path string, name string) error {
	_, err := cl.Request(MethodTorrentSetLocation, struct {
		IDs  []string `json:"ids"`
		Path string   `json:"path"`
		Name string   `json:"name"`
	}{ids, path, name})
	return err
}

func (cl *Client) SessionInfo(fields []string) (*SessionInfo, error) {
	req := struct {
		Fields []string `json:"fields,omitempty"`
	}{fields}
	resp, err := cl.Request(MethodSessionGet, req)
	if err != nil {
		return nil, err
	}

	var out SessionInfo
	if err := json.Unmarshal([]byte(*resp.Arguments), &out); err != nil {
		return nil, err
	}
	return &out, nil
}
