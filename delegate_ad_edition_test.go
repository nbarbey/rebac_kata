package rebac_kata

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserCannotEditAdOfAnotherUser(t *testing.T) {
	t.Skip()
	a := NewApplication(map[string]*User{"Jane": {}, "Joe": {}})

	jane := a.Login("Jane")
	adId := jane.PublishAd(&Ad{Title: "LEGO Space Astronaut", Price: 30})

	joe := a.Login("Joe")
	err := joe.ChangeAdPrice(adId, 10)
	require.ErrorIs(t, err, ErrUnauthorized)
}

func TestUserCanEditOwnAd(t *testing.T) {
	t.Skip()
	a := NewApplication(map[string]*User{"Jane": {}, "Joe": {}})

	jane := a.Login("Jane")
	adId := jane.PublishAd(&Ad{Title: "LEGO Space Astronaut", Price: 30})

	require.NoError(t, jane.ChangeAdPrice(adId, 10))

	assert.Equal(t, 10, jane.GetAd(adId).Price)
}

func TestDelegateAdEdition_ok(t *testing.T) {
	t.Skip()
	a := NewApplication(map[string]*User{"Jane": {}, "Joe": {}})

	jane := a.Login("Jane")
	adId := jane.PublishAd(&Ad{Title: "LEGO Space Astronaut", Price: 30})

	jane.DelegateAdEdition(adId, "Joe")
	joe := a.Login("Joe")
	require.NoError(t, joe.ChangeAdPrice(adId, 10))

	assert.Equal(t, 10, joe.GetAd(adId).Price)
}
