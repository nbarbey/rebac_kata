package rebac_kata

type DelegateAdEditionParameters struct {
	adId     int
	userName string
}

type Option func(parameters *DelegateAdEditionParameters)

func ForAd(adId int) Option {
	return func(parameters *DelegateAdEditionParameters) {
		parameters.adId = adId
	}
}

func ToUser(userName string) Option {
	return func(parameters *DelegateAdEditionParameters) {
		parameters.userName = userName
	}
}

func (u *User) DelegateAdEdition(options ...Option) {
	parameters := DelegateAdEditionParameters{}
	for _, option := range options {
		option(&parameters)
	}
}
