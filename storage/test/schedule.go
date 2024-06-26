package storagetest

import (
	"cmp"
	"slices"
	"testing"
	"time"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *Suite) TestScheduleUpdate(t *testing.T) {
	s := suite.Storage(t)
	ss := s.Schedule(suite.ctx)

	user := OneOff[radio.User](genUser())
	user.ID = 0

	uid, err := s.User(suite.ctx).Create(user)
	require.NoError(t, err)
	user.ID = uid

	empty, err := ss.Latest()
	require.NoError(t, err)
	require.Equal(t, 7, len(empty), "latest should have 7 entries")

	should := make(map[radio.ScheduleDay]radio.ScheduleEntry)

	var entries []radio.ScheduleEntry
	for i := radio.ScheduleDay(0); i <= radio.Sunday; i++ {
		entries = append(entries, radio.ScheduleEntry{
			Weekday:      i,
			Text:         "This is " + i.String(),
			UpdatedBy:    user,
			Notification: true,
		})
	}

	updatedMonday := radio.ScheduleEntry{
		Weekday:      radio.Monday,
		Text:         "This is an updated Monday",
		UpdatedBy:    user,
		Notification: false,
	}

	updatedFriday := radio.ScheduleEntry{
		Weekday:      radio.Friday,
		Text:         "And this is an updated Friday",
		Owner:        &user,
		UpdatedBy:    user,
		Notification: true,
	}
	entries = append(entries, updatedMonday, updatedFriday)

	for _, entry := range entries {
		err := ss.Update(entry)
		require.NoError(t, err)
		time.Sleep(time.Second / 4)
		should[entry.Weekday] = entry
	}

	latest, err := ss.Latest()
	require.NoError(t, err)

	// theres only 7 days, so we should only get 7 entries back
	require.Equal(t, 7, len(latest), "latest should have 7 entries")
	require.True(t, slices.IsSortedFunc(latest, func(a, b *radio.ScheduleEntry) int {
		return cmp.Compare(a.Weekday, b.Weekday)
	}), "latest should be sorted by weekday")

	for _, got := range latest {
		s := should[got.Weekday]

		assert.Equal(t, s.Text, got.Text)
		assert.Equal(t, s.Notification, got.Notification)
		assert.Equal(t, s.UpdatedBy.ID, got.UpdatedBy.ID)
		if s.Owner != nil && assert.NotNil(t, got.Owner) {
			assert.Equal(t, s.Owner.ID, got.Owner.ID)
		}
	}

	history, err := ss.History(radio.Friday, 10, 0)
	require.NoError(t, err)
	require.Equal(t, 2, len(history), "history for friday should have two entries")
	assert.Equal(t, updatedFriday.Text, history[0].Text)
	assert.Equal(t, entries[radio.Friday].Text, history[1].Text)
}
