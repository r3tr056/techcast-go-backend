package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Options struct {
	DB db.DB*
	MusicPaths []string
	PodcastPath string
	CachePath string
	CoverCachePath string
	ProxyPrefix string
	GenreSplit string
	HTTPLog bool
	JukeboxEnabled bool
}

type Server struct {
	scanner *scanner.Scanner
	jukebox *jukebox.Jukebox
	router *mux.Router
	sessDB *gormstore.Store
	podcast *podcasts.Podcasts
}

func NewServer(opts Options) (*Server, error) {
	for i, musicPath := range opts.MusicPaths {
		opts.MusicPaths[i] = filepath.Clean(musicPath)
	}
	opts.CachePath = filepath.Clean(opts.CachePath)
	opts.PodcastPath = filepath.Clean(opts.PodcastPath)

	tagger := &tags.TagReader{}

	scanner := scanner.New(opts.MusicPaths, opts.DB, opts.GenreSplit, tagger)
	base := &ctrlbase.Controller{
		DB: opts.DB,
		ProxyPrefix: opts.ProxyPrefix,
		Scanner: scanner,
	}

	// router with common wares for admin and techcasts
	r := mux.NewRouter()
	if opts.HTTPLog {
		r.Use(base.WithLogging)
	}
	r.Use(base.WithCORS)

	sessKey, err := opts.DB.GetSetting("session_key")
	if err != nil {
		return nil, fmt.Errorf("get session key : %w", err)
	}

	if sessKey == "" {
		if err := opts.DB.SetSetting("session_key", string(securecookie.GenerateRandomKey(32))); err != nil {
			return nil, fmt.Errorf("set session key: %w", err)
		}
	}

	sessDB := gormstore.New(opts.DB.DB, []byte(sessKey))
	sessDB.SessionOpts.HttpOnly = true
	sessDB.SessionOpts.SameSite = http.SameSiteLaxMode

	podcast := podcasts.New(opts.DB, opts.PodcastPath, tagger)
	
	cacheTranscoder := transcode.NewCachingTranscoder(
		transcode.NewFFmpegTranscoder(),
		opts.CachePath,
	)

	ctrlAdmin, err := ctrladmin.New(base, sessDB, podcast)
	if err != nil {
		return nil, fmt.Errorf("create admin controller : %w", err)
	}

	ctrlTechcasts := &ctrlTechcasts.Controller{
		Controller: base,
		CachePath: opts.CachePath,
		CoverCachePath: opts.CoverCachePath,
		PodcastPath: opts.PodcastPath,
		MusicPaths: opts.MusicPaths,
		Jukebox: &jukebox.Jukebox{},
		Scrobblers: []scrobble.Scrobbler{&lastfm.Scrobbler{DB: opts.DB}, &listenbrainz.Scrobbler{}},
	}
}