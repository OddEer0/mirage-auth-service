package appDto

type (
	PureUser struct {
		Id        string
		Login     string
		Email     string
		IsBanned  bool
		BanReason *string
		Role      string
	}

	RegistrationData struct {
		Login    string
		Password string
		Email    string
	}

	LoginData struct {
	}

	SaveTokenServiceDto struct {
		Id           string
		RefreshToken string
	}
)
