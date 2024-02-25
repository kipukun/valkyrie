package storagetest

import (
	"slices"
	"strconv"
	"time"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *Suite) TestSongCreateAndRetrieve() {
	t := suite.T()
	ss := suite.Storage.Song(suite.ctx)

	song := radio.Song{
		Metadata: "test-song-create-and-retrieve",
		Length:   time.Second * 300,
	}
	song.Hydrate()

	new, err := ss.Create(song)
	if assert.NoError(t, err) {
		assert.NotNil(t, new)
		assert.True(t, song.EqualTo(*new))
		assert.Equal(t, song.Length, new.Length)
	}

	fromHash, err := ss.FromHash(song.Hash)
	if assert.NoError(t, err) {
		assert.NotNil(t, fromHash)
		assert.True(t, song.EqualTo(*fromHash))
		assert.Equal(t, song.Length, fromHash.Length)
	}

	fromMetadata, err := ss.FromMetadata(song.Metadata)
	if assert.NoError(t, err) {
		assert.NotNil(t, fromMetadata)
		assert.True(t, song.EqualTo(*fromMetadata))
		assert.Equal(t, song.Length, fromMetadata.Length)
	}
}

func (suite *Suite) TestSongLastPlayed() {
	t := suite.T()
	ss := suite.Storage.Song(suite.ctx)

	base := radio.Song{
		Metadata: "test-song-last-played-",
	}
	user := radio.User{
		DJ: radio.DJ{
			ID: 10,
		},
	}
	amount := int64(50)

	// create 50 testing songs
	var songs []radio.Song
	for i := int64(0); i < amount; i++ {
		song := base
		song.Length = time.Duration(i*2) * time.Second
		song.Metadata = song.Metadata + strconv.FormatInt(i, 10)
		song.Hydrate()

		new, err := ss.Create(song)
		require.NoError(t, err)
		require.NotNil(t, new)
		assert.True(t, song.EqualTo(*new))

		songs = append(songs, *new)
	}

	// now have them all play
	for i, song := range songs {
		err := ss.AddPlay(song, user, nil)
		require.NoError(t, err)

		if i == 15 || i == 40 { // Artificially wait a second in the middle somewhere
			time.Sleep(time.Second)
		}
	}

	n, err := ss.LastPlayedCount()
	require.NoError(t, err)
	assert.Equal(t, amount, n)

	// test the full list of songs
	lp, err := ss.LastPlayed(0, amount)
	require.NoError(t, err)
	// reverse them since we added them in 0-49 order but we will get them back as 49-0 order
	slices.Reverse(lp)
	for i, original := range songs {
		assert.True(t, original.EqualTo(lp[i]), "all: expected %s got %s", original.Metadata, lp[i].Metadata)
		// we also don't have any database tracks or users associated with these songs
		// so the songs returned by this call of LastPlayed should have the following properties:
		// 		.HasTrack() should be false
		//		.DatabaseTrack should be nil
		//		.LastPlayedBy should be nil
		assert.False(t, lp[i].HasTrack(), "has track should be false")
		assert.Nil(t, lp[i].DatabaseTrack, "database track should be nil")
		assert.Nil(t, lp[i].LastPlayedBy, "last played by should be nil")
	}

	// test a subset of the list
	lp, err = ss.LastPlayed(0, 20)
	require.NoError(t, err)
	slices.Reverse(lp)
	for i, original := range songs[amount-20 : amount] {
		assert.True(t, original.EqualTo(lp[i]), "subset start: expected %s got %s", original.Metadata, lp[i].Metadata)
	}

	// test the other end of the subset
	lp, err = ss.LastPlayed(30, 20)
	require.NoError(t, err)
	slices.Reverse(lp)
	for i, original := range songs[:20] {
		assert.True(t, original.EqualTo(lp[i]), "subset end: expected %s got %s", original.Metadata, lp[i].Metadata)
	}
}
