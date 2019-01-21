package ircbot

import (
	"context"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/R-a-dio/valkyrie/rpc"
	"github.com/twitchtv/twirp"
)

func NewHTTPServer(b *Bot) (*http.Server, error) {
	rpcServer := rpc.NewBotServer(b, nil)
	mux := http.NewServeMux()
	// rpc server path
	mux.Handle(rpc.BotPathPrefix, rpcServer)

	// debug symbols
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	conf := b.Conf()
	server := &http.Server{Addr: conf.IRC.Addr, Handler: mux}
	return server, nil
}

func (b *Bot) AnnounceSong(ctx context.Context, song *rpc.Song) (*rpc.Null, error) {
	e := Event{
		Bot:    b,
		Client: b.c,
	}

	fn := nowPlayingMessage(e)
	message, args, err := fn()
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	// only difference between the announce and .np command is that it
	// starts with "Now starting" instead of "Now playing"
	message = strings.Replace(message, "playing", "starting", 1)

	b.c.Cmd.Message(b.Conf().IRC.MainChannel, Fmt(message, args...))
	return new(rpc.Null), nil
}