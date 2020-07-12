package transmission

import (
	"encoding/json"
	"time"
)

type Response struct {
	Result    string           `json:"result"`
	Arguments *json.RawMessage `json:"arguments"`
	Tag       int              `json:"tag"`
}

type SessionStats struct {
	ActiveTorrents  int             `json:"activeTorrentCount"`
	DownloadSpeed   int             `json:"downloadSpeed"`
	PausedTorrents  int             `json:"pausedTorrentCount"`
	Torrents        int             `json:"torrentCount"`
	UploadSpeed     int             `json:"uploadSpeed"`
	CumulativeStats SessionStatsSub `json:"cumulative-stats"`
	CurrentStats    SessionStatsSub `json:"current-stats"`
}

type SessionStatsSub struct {
	UploadedBytes   int `json:"uploadedBytes"`
	DownloadedBytes int `json:"downloadedBytes"`
	FilesAdded      int `json:"filesAdded"`
	SessionCount    int `json:"sessionCount"`
	SecondsActive   int `json:"secondsActive"`
}

type NewTorrent struct {
	Cookies           string   `json:"cookies,omitempty"`
	DownloadDir       string   `json:"download-dir,omitempty"`
	Filename          string   `json:"filename,omitempty"`
	Metainfo          string   `json:"metainfo,omitempty"`
	Paused            bool     `json:"paused,omitempty"`
	PeerLimit         int      `json:"peer-limit,omitempty"`
	BandwidthPriority Priority `json:"bandwidthPriority,omitempty"`
	FilesWanted       []int    `json:"files-wanted,omitempty"`
	FilesUnwanted     []int    `json:"files-unwanted,omitempty"`
	PriorityHigh      []int    `json:"priority-high,omitempty"`
	PriorityLow       []int    `json:"priority-low,omitempty"`
	PriorityNormal    []int    `json:"priority-normal,omitempty"`
}

type AddedTorrent struct {
	ID   int
	Name string
	Hash string
}

type TorrentInfo struct {
	// The last time we uploaded or downloaded piece data on this torrent.
	ActivityDate time.Time
	// When the torrent was first added.
	AddedDate         time.Time
	BandwidthPriority Priority
	Comment           string
	// Byte count of all the corrupt data you've ever downloaded for
	// this torrent. If you're on a poisoned torrent, this number can
	// grow very large.
	CorruptEver int
	Creator     string
	DateCreated time.Time
	// Byte count of all the piece data we want and don't have yet,
	// but that a connected peer does have. [0...leftUntilDone]
	DesiredAvailable int
	// When the torrent finished downloading.
	DoneDate    time.Time
	DownloadDir string
	// Byte count of all the non-corrupt data you've ever downloaded
	// for this torrent. If you deleted the files and downloaded a second
	// time, this will be 2*totalSize.
	DownloadedEver  int
	DownloadLimit   int
	DownloadLimited bool
	// The last time during this session that a rarely-changing field
	// changed -- e.g. any tr_info field (trackers, filenames, name)
	// or download directory. RPC clients can monitor this to know when
	// to reload fields that rarely change.
	EditDate time.Time
	// Defines what kind of text is in errorString.
	// @see errorString
	Error int
	// A warning or error message regarding the torrent.
	// @see error
	ErrorString string
	// If downloading, estimated number of seconds left until the torrent is done.
	// If seeding, estimated number of seconds left until seed ratio is reached.
	ETA time.Duration
	// If seeding, number of seconds left until the idle time limit is reached.
	ETAIdle   time.Duration
	FileStats []FileStats
	Files     []File
	Hash      string
	// Byte count of all the partial piece data we have for this torrent.
	// As pieces become complete, this value may decrease as portions of it
	// are moved to `corrupt' or `haveValid'.
	HaveUnchecked int
	// Byte count of all the checksum-verified data we have for this torrent.
	HaveValid           int
	HonorsSessionLimits bool
	ID                  int
	IsFinished          bool
	IsPrivate           bool
	// True if the torrent is running, but has been idle for long enough
	// to be considered stalled.  @see tr_sessionGetQueueStalledMinutes()
	IsStalled bool
	Labels    []string
	// Byte count of how much data is left to be downloaded until we've got
	// all the pieces that we want. [0...tr_info.sizeWhenDone]
	LeftUntilDone int
	MagnetLink    string
	// time when one or more of the torrent's trackers will
	// allow you to manually ask for more peers,
	// or 0 if you can't.
	ManualAnnounceTime time.Time
	MaxConnectedPeers  int
	// How much of the metadata the torrent has.
	// For torrents added from a .torrent this will always be 1.
	// For magnet links, this number will from from 0 to 1 as the metadata is downloaded.
	// Range is [0..1]
	MetadataPercentComplete float64
	Name                    string
	PeerLimit               int
	Peers                   []Peer
	// Number of peers that we're connected to.
	PeersConnected int
	// How many peers we found out about from the tracker, or from pex,
	// or from incoming connections, or from our resume file.
	PeersFrom PeersFrom
	// Number of peers that we're sending data to.
	PeersGettingFromUs int
	// Number of peers that are sending data to us.
	PeersSendingToUs int
	// How much has been downloaded of the files the user wants. This differs
	// from percentComplete if the user wants only some of the torrent's files.
	// Range is [0..1]
	// @see tr_stat.leftUntilDone
	PercentDone float64
	PieceCount  int
	PieceSize   int
	Pieces      []byte
	Priorities  []Priority
	// This torrent's queue position.
	// All torrents have a queue position, even if it's not queued.
	QueuePosition int
	RateDownload  int
	RateUpload    int
	// When tr_stat.activity is TR_STATUS_CHECK or TR_STATUS_CHECK_WAIT,
	// this is the percentage of how much of the files has been
	// verified. When it gets to 1, the verify process is done.
	// Range is [0..1]
	// @see tr_stat.activity
	RecheckProgress float64
	// Cumulative seconds the torrent's ever spent downloading.
	SecondsDownloading time.Duration
	// Cumulative seconds the torrent's ever spent seeding.
	SecondsSeeding time.Duration
	SeedIdleLimit  int
	SeedIdleMode   int
	SeedRatioLimit float64
	SeedRatioMode  int
	// Byte count of all the piece data we'll have downloaded when we're done,
	// whether or not we have it yet. This may be less than tr_info.totalSize
	// if only some of the torrent's files are wanted.
	// [0...tr_info.totalSize]
	SizeWhenDone int
	// When the torrent was last started.
	StartDate   time.Time
	Status      TorrentStatus
	TorrentFile string
	// total size of the torrent, in bytes.
	TotalSize     int
	TrackerStats  []TrackerStats
	Trackers      []Tracker
	UploadLimit   int
	UploadLimited bool
	UploadRatio   float64
	// Byte count of all data you've ever uploaded for this torrent.
	UploadedEver int
	Wanted       []bool
	Webseeds     []string
	// Number of webseeds that are sending data to us.
	WebseedsSendingToUs int
}

type FileStats struct {
	BytesCompleted int      `json:"bytesCompleted"`
	Wanted         bool     `json:"wanted"`
	Priority       Priority `json:"priority"`
}

type File struct {
	BytesCompleted int    `json:"bytesCompleted"`
	Length         int    `json:"length"`
	Name           string `json:"name"`
}

type PeersFrom struct {
	FromCache    int `json:"fromCache"`
	FromDHT      int `json:"fromDht"`
	FromIncoming int `json:"fromIncoming"`
	FromLPD      int `json:"fromLpd"`
	FromLTEP     int `json:"fromLtep"`
	FromPEX      int `json:"fromPex"`
	FromTracker  int `json:"fromTracker"`
}

type Tracker struct {
	Announce string `json:"announce"`
	ID       int    `json:"id"`
	Scrape   string `json:"scrape"`
	Tier     int    `json:"tier"`
}

type torrentInfo struct {
	ActivityDate            int            `json:"activityDate"`
	AddedDate               int            `json:"addedDate"`
	BandwidthPriority       Priority       `json:"bandwidthPriority"`
	Comment                 string         `json:"comment"`
	CorruptEver             int            `json:"corruptEver"`
	Creator                 string         `json:"creator"`
	DateCreated             int            `json:"dateCreated"`
	DesiredAvailable        int            `json:"desiredAvailable"`
	DoneDate                int            `json:"doneDate"`
	DownloadDir             string         `json:"downloadDir"`
	DownloadedEver          int            `json:"downloadedEver"`
	DownloadLimit           int            `json:"downloadLimit"`
	DownloadLimited         bool           `json:"downloadLimited"`
	EditDate                int            `json:"editDate"`
	Error                   int            `json:"error"`
	ErrorString             string         `json:"errorString"`
	ETA                     int            `json:"eta"`
	ETAIdle                 int            `json:"etaIdle"`
	FileStats               []FileStats    `json:"fileStats"`
	Files                   []File         `json:"files"`
	Hash                    string         `json:"hashString"`
	HaveUnchecked           int            `json:"haveUnchecked"`
	HaveValid               int            `json:"haveValid"`
	HonorsSessionLimits     bool           `json:"honorsSessionLimits"`
	ID                      int            `json:"id"`
	IsFinished              bool           `json:"isFinished"`
	IsPrivate               bool           `json:"isPrivate"`
	IsStalled               bool           `json:"isStalled"`
	Labels                  []string       `json:"labels"`
	LeftUntilDone           int            `json:"leftUntilDone"`
	MagnetLink              string         `json:"magnetLink"`
	ManualAnnounceTime      int            `json:"manualAnnounceTime"`
	MaxConnectedPeers       int            `json:"maxConnectedPeers"`
	MetadataPercentComplete float64        `json:"metadataPercentComplete"`
	Name                    string         `json:"name"`
	PeerLimit               int            `json:"peer-limit"`
	Peers                   []Peer         `json:"peers"`
	PeersConnected          int            `json:"peersConnected"`
	PeersFrom               PeersFrom      `json:"peersFrom"`
	PeersGettingFromUs      int            `json:"peersGettingFromUs"`
	PeersSendingToUs        int            `json:"peersSendingToUs"`
	PercentDone             float64        `json:"percentDone"`
	PieceCount              int            `json:"pieceCount"`
	PieceSize               int            `json:"pieceSize"`
	Pieces                  []byte         `json:"pieces"`
	Priorities              []Priority     `json:"priorities"`
	QueuePosition           int            `json:"queuePosition"`
	RateDownload            int            `json:"rateDownload (B/s)"`
	RateUpload              int            `json:"rateUpload (B/s)"`
	RecheckProgress         float64        `json:"recheckProgress"`
	SecondsDownloading      int            `json:"secondsDownloading"`
	SecondsSeeding          int            `json:"secondsSeeding"`
	SeedIdleLimit           int            `json:"seedIdleLimit"`
	SeedIdleMode            int            `json:"seedIdleMode"`
	SeedRatioLimit          float64        `json:"seedRatioLimit"`
	SeedRatioMode           int            `json:"seedRatioMode"`
	SizeWhenDone            int            `json:"sizeWhenDone"`
	StartDate               int            `json:"startDate"`
	Status                  TorrentStatus  `json:"status"`
	TorrentFile             string         `json:"torrentFile"`
	TotalSize               int            `json:"totalSize"`
	TrackerStats            []trackerStats `json:"trackerStats"`
	Trackers                []Tracker      `json:"trackers"`
	UploadLimit             int            `json:"uploadLimit"`
	UploadLimited           bool           `json:"uploadLimited"`
	UploadRatio             float64        `json:"uploadRatio"`
	UploadedEver            int            `json:"uploadedEver"`
	Wanted                  []int          `json:"wanted"`
	Webseeds                []string       `json:"webseeds"`
	WebseedsSendingToUs     int            `json:"webseedsSendingToUs"`
}

type Peer struct {
	Address            string  `json:"address"`
	ClientName         string  `json:"clientName"`
	ClientIsChoked     bool    `json:"clientIsChoked"`
	ClientIsInterested bool    `json:"clientIsInterested"`
	FlagStr            string  `json:"flagStr"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom"`
	IsEncrypted        bool    `json:"isEncrypted"`
	IsIncoming         bool    `json:"isIncoming"`
	IsUploadingTo      bool    `json:"isUploadingTo"`
	IsUTP              bool    `json:"isUTP"`
	PeerIsChoked       bool    `json:"peerIsChoked"`
	PeerIsInterested   bool    `json:"peerIsInterested"`
	Port               int     `json:"port"`
	Progress           float64 `json:"progress"`
	RateToClient       int     `json:"rateToClient (B/s)"`
	RateToPeer         int     `json:"rateToPeer (B/s)"`
}

type SessionInfo struct {
	// max global download speed (KBps)
	AltSpeedDown int `json:"alt-speed-down"`
	// true means use the alt speeds
	AltSpeedEnabled bool `json:"alt-speed-enabled"`
	// when to turn on alt speeds (units: minutes after midnight)
	AltSpeedTimeBegin int `json:"alt-speed-time-begin"`
	// true means the scheduled on/off times are used
	AltSpeedTimeEnabled bool `json:"alt-speed-time-enabled"`
	// when to turn off alt speeds (units: same)
	AltSpeedTimeEnd int `json:"alt-speed-time-end"`
	// what day(s) to turn on alt speeds (look at tr_sched_day)
	AltSpeedTimeDay int `json:"alt-speed-time-day"`
	// max global upload speed (KBps)
	AltSpeedUp int `json:"alt-speed-up"`
	// location of the blocklist to use for "blocklist-update"
	BlocklistUrl string `json:"blocklist-url"`
	// true means enabled
	BlocklistEnabled bool `json:"blocklist-enabled"`
	// number of rules in the blocklist
	BlocklistSize int `json:"blocklist-size"`
	// maximum size of the disk cache (MB)
	CacheSizeMB int `json:"cache-size-mb"`
	// location of transmission's configuration directory
	ConfigDir string `json:"config-dir"`
	// default path to download torrents
	DownloadDir string `json:"download-dir"`
	// max number of torrents to download at once (see download-queue-enabled)
	DownloadQueueSize int `json:"download-queue-size"`
	// if true, limit how many torrents can be downloaded at once
	DownloadQueueEnabled bool `json:"download-queue-enabled"`
	// true means allow dht in public torrents
	DhtEnabled bool `json:"dht-enabled"`
	// "required", "preferred", "tolerated"
	Encryption string `json:"encryption"`
	// torrents we're seeding will be stopped if they're idle for this long
	IdleSeedingLimit int `json:"idle-seeding-limit"`
	// true if the seeding inactivity limit is honored by default
	IdleSeedingLimitEnabled bool `json:"idle-seeding-limit-enabled"`
	// path for incomplete torrents, when enabled
	IncompleteDir string `json:"incomplete-dir"`
	// true means keep torrents in incomplete-dir until done
	IncompleteDirEnabled bool `json:"incomplete-dir-enabled"`
	// true means allow Local Peer Discovery in public torrents
	LpdEnabled bool `json:"lpd-enabled"`
	// maximum global number of peers
	PeerLimitGlobal int `json:"peer-limit-global"`
	// maximum global number of peers
	PeerLimitPerTorrent int `json:"peer-limit-per-torrent"`
	// true means allow pex in public torrents
	PexEnabled bool `json:"pex-enabled"`
	// port number
	PeerPort int `json:"peer-port"`
	// true means pick a random peer port on launch
	PeerPortRandomOnStart bool `json:"peer-port-random-on-start"`
	// true means enabled
	PortForwardingEnabled bool `json:"port-forwarding-enabled"`
	// whether or not to consider idle torrents as stalled
	QueueStalledEnabled bool `json:"queue-stalled-enabled"`
	// torrents that are idle for N minuets aren't counted toward seed-queue-size or download-queue-size
	QueueStalledMinutes int `json:"queue-stalled-minutes"`
	// true means append ".part" to incomplete files
	RenamePartialFiles bool `json:"rename-partial-files"`
	// the current RPC API version
	RpcVersion int `json:"rpc-version"`
	// the minimum RPC API version supported
	RpcVersionMinimum int `json:"rpc-version-minimum"`
	// filename of the script to run
	ScriptTorrentDoneFilename string `json:"script-torrent-done-filename"`
	// whether or not to call the "done" script
	ScriptTorrentDoneEnabled bool `json:"script-torrent-done-enabled"`
	// the default seed ratio for torrents to use
	SeedRatioLimit float64 `json:"seedRatioLimit"`
	// true if seedRatioLimit is honored by default
	SeedRatioLimited bool `json:"seedRatioLimited"`
	// max number of torrents to uploaded at once (see seed-queue-enabled)
	SeedQueueSize int `json:"seed-queue-size"`
	// if true, limit how many torrents can be uploaded at once
	SeedQueueEnabled bool `json:"seed-queue-enabled"`
	// max global download speed (KBps)
	SpeedLimitDown int `json:"speed-limit-down"`
	// true means enabled
	SpeedLimitDownEnabled bool `json:"speed-limit-down-enabled"`
	// max global upload speed (KBps)
	SpeedLimitUp int `json:"speed-limit-up"`
	// true means enabled
	SpeedLimitUpEnabled bool `json:"speed-limit-up-enabled"`
	// true means added torrents will be started right away
	StartAddedTorrents bool `json:"start-added-torrents"`
	// true means the .torrent file of added torrents will be deleted
	TrashOriginalTorrentFiles bool `json:"trash-original-torrent-files"`
	//
	Units struct {
		// 4 strings: KB/s, MB/s, GB/s, TB/s
		SpeedUnits []string `json:"speed-units"`
		// number of bytes in a KB (1000 for kB; 1024 for KiB)
		SpeedBytes int `json:"speed-bytes"`
		// 4 strings: KB/s, MB/s, GB/s, TB/s
		SizeUnits []string `json:"size-units"`
		// number of bytes in a KB (1000 for kB; 1024 for KiB)
		SizeBytes int `json:"size-bytes"`
		// 4 strings: KB/s, MB/s, GB/s, TB/s
		MemoryUnits []string `json:"memory-units"`
		// number of bytes in a KB (1000 for kB; 1024 for KiB)
		MemoryBytes int `json:"memory-bytes"`
	} `json:"units"`
	// true means allow utp
	UtpEnabled bool `json:"utp-enabled"`
	// long version string "$version ($revision)"
	Version string `json:"version"`
}

type trackerStats struct {
	Announce              string       `json:"announce"`
	AnnounceState         TrackerState `json:"announceState"`
	DownloadCount         int          `json:"downloadCount"`
	HasAnnounced          bool         `json:"hasAnnounced"`
	HasScraped            bool         `json:"hasScraped"`
	Host                  string       `json:"host"`
	ID                    int          `json:"id"`
	IsBackup              bool         `json:"isBackup"`
	LastAnnouncePeerCount int          `json:"lastAnnouncePeerCount"`
	LastAnnounceResult    string       `json:"lastAnnounceResult"`
	LastAnnounceStartTime int          `json:"lastAnnounceStartTime"`
	LastAnnounceSucceeded bool         `json:"lastAnnounceSucceeded"`
	LastAnnounceTime      int          `json:"lastAnnounceTime"`
	LastAnnounceTimedOut  bool         `json:"lastAnnounceTimedOut"`
	LastScrapeResult      string       `json:"lastScrapeResult"`
	LastScrapeStartTime   int          `json:"lastScrapeStartTime"`
	LastScrapeSucceeded   bool         `json:"lastScrapeSucceeded"`
	LastScrapeTime        int          `json:"lastScrapeTime"`
	LastScrapeTimedOut    bool         `json:"lastScrapeTimedOut"`
	LeecherCount          int          `json:"leecherCount"`
	NextAnnounceTime      int          `json:"nextAnnounceTime"`
	NextScrapeTime        int          `json:"nextScrapeTime"`
	Scrape                string       `json:"scrape"`
	ScrapeState           TrackerState `json:"scrapeState"`
	SeederCount           int          `json:"seederCount"`
	Tier                  int          `json:"tier"`
}

type TrackerStats struct {
	// the full announce URL
	Announce string
	// is the tracker announcing, waiting, queued, etc
	AnnounceState TrackerState
	// how many downloads this tracker knows of (-1 means it does not know)
	DownloadCount int
	// whether or not we've ever sent this tracker an announcement
	HasAnnounced bool
	// whether or not we've ever scraped to this tracker
	HasScraped bool
	// human-readable string identifying the tracker
	Host string
	ID   int
	// Transmission uses one tracker per tier, and the others are kept as backups
	IsBackup bool
	// number of peers the tracker told us about last time.
	// if "lastAnnounceSucceeded" is false, this field is undefined
	LastAnnouncePeerCount int
	// human-readable string with the result of the last announce.
	// if "hasAnnounced" is false, this field is undefined
	LastAnnounceResult string
	// when the last announce was sent to the tracker.
	// if "hasAnnounced" is false, this field is undefined
	LastAnnounceStartTime time.Time
	// whether or not the last announce was a success.
	// if "hasAnnounced" is false, this field is undefined
	LastAnnounceSucceeded bool
	// when the last announce was completed.
	// if "hasAnnounced" is false, this field is undefined
	LastAnnounceTime     time.Time
	LastAnnounceTimedOut bool
	// human-readable string with the result of the last scrape.
	// if "hasScraped" is false, this field is undefined
	LastScrapeResult string
	// when the last scrape was sent to the tracker.
	// if "hasScraped" is false, this field is undefined
	LastScrapeStartTime time.Time
	// whether or not the last scrape was a success.
	// if "hasAnnounced" is false, this field is undefined
	LastScrapeSucceeded bool
	// when the last scrape was completed.
	// if "hasScraped" is false, this field is undefined
	LastScrapeTime time.Time
	// whether or not the last scrape timed out.
	LastScrapeTimedOut bool
	// number of leechers this tracker knows of (-1 means it does not know)
	LeecherCount int
	// when the next periodic announce message will be sent out.
	// if announceState isn't TR_TRACKER_WAITING, this field is undefined
	NextAnnounceTime time.Time
	// when the next periodic scrape message will be sent out.
	// if scrapeState isn't TR_TRACKER_WAITING, this field is undefined
	NextScrapeTime time.Time
	// the full scrape URL
	Scrape string
	// is the tracker scraping, waiting, queued, etc
	ScrapeState TrackerState
	// number of seeders this tracker knows of (-1 means it does not know)
	SeederCount int
	// which tier this tracker is in
	Tier int
}

func convertTorrentInfo(in *torrentInfo, out *TorrentInfo) {
	*out = TorrentInfo{
		ActivityDate:            time.Unix(int64(in.ActivityDate), 0),
		AddedDate:               time.Unix(int64(in.AddedDate), 0),
		BandwidthPriority:       in.BandwidthPriority,
		Comment:                 in.Comment,
		CorruptEver:             in.CorruptEver,
		Creator:                 in.Creator,
		DateCreated:             time.Unix(int64(in.DateCreated), 0),
		DesiredAvailable:        in.DesiredAvailable,
		DoneDate:                time.Unix(int64(in.DoneDate), 0),
		DownloadDir:             in.DownloadDir,
		DownloadedEver:          in.DownloadedEver,
		DownloadLimit:           in.DownloadLimit,
		DownloadLimited:         in.DownloadLimited,
		EditDate:                time.Unix(int64(in.EditDate), 0),
		Error:                   in.Error,
		ErrorString:             in.ErrorString,
		ETA:                     time.Duration(in.ETA) * time.Second,
		ETAIdle:                 time.Duration(in.ETAIdle) * time.Second,
		FileStats:               in.FileStats,
		Files:                   in.Files,
		Hash:                    in.Hash,
		HaveUnchecked:           in.HaveUnchecked,
		HaveValid:               in.HaveValid,
		HonorsSessionLimits:     in.HonorsSessionLimits,
		ID:                      in.ID,
		IsFinished:              in.IsFinished,
		IsPrivate:               in.IsPrivate,
		IsStalled:               in.IsStalled,
		Labels:                  in.Labels,
		LeftUntilDone:           in.LeftUntilDone,
		MagnetLink:              in.MagnetLink,
		ManualAnnounceTime:      time.Unix(int64(in.ManualAnnounceTime), 0),
		MaxConnectedPeers:       in.MaxConnectedPeers,
		MetadataPercentComplete: in.MetadataPercentComplete,
		Name:                    in.Name,
		PeerLimit:               in.PeerLimit,
		Peers:                   in.Peers,
		PeersConnected:          in.PeersConnected,
		PeersFrom:               in.PeersFrom,
		PeersGettingFromUs:      in.PeersGettingFromUs,
		PeersSendingToUs:        in.PeersSendingToUs,
		PercentDone:             in.PercentDone,
		PieceCount:              in.PieceCount,
		PieceSize:               in.PieceSize,
		Pieces:                  in.Pieces,
		Priorities:              in.Priorities,
		QueuePosition:           in.QueuePosition,
		RateDownload:            in.RateDownload,
		RateUpload:              in.RateUpload,
		RecheckProgress:         in.RecheckProgress,
		SecondsDownloading:      time.Duration(in.SecondsDownloading) * time.Second,
		SecondsSeeding:          time.Duration(in.SecondsSeeding) * time.Second,
		SeedIdleLimit:           in.SeedIdleLimit,
		SeedIdleMode:            in.SeedIdleMode,
		SeedRatioLimit:          in.SeedRatioLimit,
		SeedRatioMode:           in.SeedRatioMode,
		SizeWhenDone:            in.SizeWhenDone,
		StartDate:               time.Unix(int64(in.StartDate), 0),
		Status:                  in.Status,
		TorrentFile:             in.TorrentFile,
		TotalSize:               in.TotalSize,
		Trackers:                in.Trackers,
		UploadLimit:             in.UploadLimit,
		UploadLimited:           in.UploadLimited,
		UploadRatio:             in.UploadRatio,
		UploadedEver:            in.UploadedEver,
		Webseeds:                in.Webseeds,
		WebseedsSendingToUs:     in.WebseedsSendingToUs,
	}
	out.Wanted = make([]bool, len(in.Wanted))
	for i := range out.Wanted {
		out.Wanted[i] = in.Wanted[i] == 1
	}
	out.TrackerStats = make([]TrackerStats, len(in.TrackerStats))
	for i := range out.TrackerStats {
		convertTrackerStats(&in.TrackerStats[i], &out.TrackerStats[i])
	}
}

func convertTrackerStats(in *trackerStats, out *TrackerStats) {
	*out = TrackerStats{
		Announce:              in.Announce,
		AnnounceState:         in.AnnounceState,
		DownloadCount:         in.DownloadCount,
		HasAnnounced:          in.HasAnnounced,
		HasScraped:            in.HasScraped,
		Host:                  in.Host,
		ID:                    in.ID,
		IsBackup:              in.IsBackup,
		LastAnnouncePeerCount: in.LastAnnouncePeerCount,
		LastAnnounceResult:    in.LastAnnounceResult,
		LastAnnounceStartTime: time.Unix(int64(in.LastAnnounceStartTime), 0),
		LastAnnounceSucceeded: in.LastAnnounceSucceeded,
		LastAnnounceTime:      time.Unix(int64(in.LastAnnounceTime), 0),
		LastAnnounceTimedOut:  in.LastAnnounceTimedOut,
		LastScrapeResult:      in.LastScrapeResult,
		LastScrapeStartTime:   time.Unix(int64(in.LastScrapeStartTime), 0),
		LastScrapeSucceeded:   in.LastScrapeSucceeded,
		LastScrapeTime:        time.Unix(int64(in.LastScrapeTime), 0),
		LastScrapeTimedOut:    in.LastScrapeTimedOut,
		LeecherCount:          in.LeecherCount,
		NextAnnounceTime:      time.Unix(int64(in.NextAnnounceTime), 0),
		NextScrapeTime:        time.Unix(int64(in.NextScrapeTime), 0),
		Scrape:                in.Scrape,
		ScrapeState:           in.ScrapeState,
		SeederCount:           in.SeederCount,
		Tier:                  in.Tier,
	}
}
