package rebac_kata

import "errors"

type User struct {
	a *Application
}

var ErrUnauthorized = errors.New("user is not authorized")

func (u *User) ChangeAdPrice(id int, price int) error {
	ad := u.GetAd(id)
	ad.Price = price
	return nil
}

func (u *User) GetAd(id int) *Ad {
	ad := u.a.ads[id]
	return ad
}

func (u *User) PublishAd(ad *Ad) (id int) {
	id = len(u.a.ads)
	u.a.ads[id] = ad
	return
}
