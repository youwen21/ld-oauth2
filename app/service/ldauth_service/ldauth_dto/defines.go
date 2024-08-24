package ldauth_dto

import "net/url"

type AuthForm struct {
	RedirectTo string `json:"redirect_to" form:"redirect_to"`
}

func (a *AuthForm) Check() error {
	if a.RedirectTo == "" {
		return nil
	}

	_, err := url.Parse(a.RedirectTo)
	return err
}

type RediForm struct {
	RedirectTo string `form:"redirect_to" json:"redirect_to"`
}

func (a *RediForm) Check() error {
	if a.RedirectTo == "" {
		return nil
	}

	_, err := url.Parse(a.RedirectTo)
	return err
}
