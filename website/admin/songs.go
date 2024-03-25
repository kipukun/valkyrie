package admin

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/R-a-dio/valkyrie/errors"
	"github.com/R-a-dio/valkyrie/util"
	"github.com/R-a-dio/valkyrie/website/middleware"
	"github.com/R-a-dio/valkyrie/website/shared"
	"github.com/rs/zerolog/hlog"
)

const songsPageSize = 20

type SongsInput struct {
	middleware.Input

	Forms []SongsForm
	Query string
	Page  *shared.Pagination
}

func (SongsInput) TemplateBundle() string {
	return "database"
}

type SongsForm struct {
	HasDelete bool
	HasEdit   bool
	Song      radio.Song
}

func (SongsForm) TemplateName() string {
	return "form_admin_songs"
}

func (SongsForm) TemplateBundle() string {
	return "database"
}

func NewSongsInput(s radio.SearchService, r *http.Request) (*SongsInput, error) {
	const op errors.Op = "website/admin.NewSongInput"
	ctx := r.Context()

	page, offset, err := shared.PageAndOffset(r, songsPageSize)
	if err != nil {
		return nil, errors.E(op, err)
	}

	query := r.FormValue("q")
	searchResult, err := s.Search(ctx, query, songsPageSize, offset)
	if err != nil {
		return nil, errors.E(op, err)
	}

	// generate the input we can so far, since we need some data from it
	input := &SongsInput{
		Input: middleware.InputFromContext(ctx),
		Query: query,
		Page: shared.NewPagination(
			page, shared.PageCount(int64(searchResult.TotalHits), songsPageSize),
			r.URL,
		),
	}

	hasDelete := input.User.UserPermissions.Has(radio.PermDatabaseDelete)
	hasEdit := input.User.UserPermissions.Has(radio.PermDatabaseEdit)
	forms := make([]SongsForm, len(searchResult.Songs))
	for i := range searchResult.Songs {
		forms[i].Song = searchResult.Songs[i]
		forms[i].HasDelete = hasDelete
		forms[i].HasEdit = hasEdit
	}

	input.Forms = forms
	return input, nil
}

func (s *State) GetSongs(w http.ResponseWriter, r *http.Request) {
	input, err := NewSongsInput(s.Search, r)
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("input creation failure")
		return
	}

	err = s.TemplateExecutor.Execute(w, r, input)
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("template failure")
		return
	}
}

func (s *State) PostSongs(w http.ResponseWriter, r *http.Request) {
	form, err := s.postSongs(w, r)
	if err != nil {
		s.errorHandler(w, r, err)
		return
	}

	if form == nil && util.IsHTMX(r) {
		// delete operation that succeeded and htmx, return nothing
		return
	}

	// otherwise just return the new listing
	s.GetSongs(w, r)
}

func (s *State) postSongs(w http.ResponseWriter, r *http.Request) (*SongsForm, error) {
	const op errors.Op = "website/admin.postSongs"
	ctx := r.Context()

	// parse the form explicitly, net/http otherwise eats any errors
	if err := r.ParseForm(); err != nil {
		return nil, errors.E(op, err, errors.InvalidForm)
	}

	ts := s.Storage.Track(r.Context())
	user := middleware.UserFromContext(ctx)
	if user == nil {
		return nil, errors.E(op, errors.AccessDenied)
	}

	// construct the new updated song form from the input
	form, err := NewSongsForm(ts, *user, r.Form)
	if err != nil {
		return nil, errors.E(op, err)
	}

	// delete action is separate from all the others
	if r.Form.Get("action") == "delete" {
		// make sure the user has permission to do this, since the route
		// only checks for PermDatabaseEdit
		if user.UserPermissions.Has(radio.PermDatabaseDelete) {
			err = ts.Delete(form.Song.TrackID)
			if err != nil {
				return nil, errors.E(op, err)
			}

			// successfully deleted the song from the database, now we just
			// need to remove the file we have on-disk
			toRemovePath := form.Song.FilePath
			if !filepath.IsAbs(toRemovePath) {
				toRemovePath = filepath.Join(s.Conf().MusicPath, toRemovePath)
			}

			err = s.FS.Remove(toRemovePath)
			if err != nil {
				return nil, errors.E(op, err, errors.InternalServer)
			}
			return nil, nil
		}
		return form, errors.E(op, errors.AccessDenied)
	}

	// anything but delete is effectively an update
	err = ts.UpdateMetadata(form.Song)
	if err != nil {
		return form, errors.E(op, err, errors.InternalServer)
	}

	return form, nil
}

func NewSongsForm(ts radio.TrackStorage, user radio.User, values url.Values) (*SongsForm, error) {
	const op errors.Op = "website/admin.NewSongsForm"

	var form SongsForm

	id, err := strconv.ParseUint(values.Get("id"), 10, 64)
	if err != nil {
		return nil, errors.E(op, err, errors.InvalidForm, errors.Info("missing id in songs form"))
	}
	tid := radio.TrackID(id)
	song, err := ts.Get(tid)
	if err != nil {
		return nil, errors.E(op, err, errors.InvalidForm)
	}

	song.Artist = values.Get("artist")
	song.Album = values.Get("album")
	song.Title = values.Get("title")
	song.Tags = values.Get("tags")

	if values.Get("action") == "mark-replacement" {
		song.NeedReplacement = true
	} else if values.Get("action") == "unmark-replacement" {
		song.NeedReplacement = false
	}
	song.LastEditor = user.Username

	form.Song = *song
	return &form, nil
}
