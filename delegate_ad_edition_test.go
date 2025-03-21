package rebac_kata

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelegateAdEdition_unauthorized(t *testing.T) {
	t.Skip()
	a := NewApplication(map[string]*User{"Jane": {}, "Joe": {}})

	jane := a.Login("Jane")
	adId := jane.PublishAd(&Ad{Title: "LEGO Space Astronaut", Price: 30})

	user := a.Login("Joe")
	err := user.ChangeAdPrice(adId, 10)
	require.ErrorIs(t, err, ErrUnauthorized)

	assert.Equal(t, 10, user.GetAd(adId).Price)
}

func TestDelegateAdEdition_ok(t *testing.T) {
	t.Skip()
	a := NewApplication(map[string]*User{"Jane": {}, "Joe": {}})

	jane := a.Login("Jane")
	adId := jane.PublishAd(&Ad{Title: "LEGO Space Astronaut", Price: 30})

	jane.DelegateAdEdition(ForAd(adId), ToUser("Joe"))
	user := a.Login("Joe")
	require.NoError(t, user.ChangeAdPrice(adId, 10))

	assert.Equal(t, 10, user.GetAd(adId).Price)
}
