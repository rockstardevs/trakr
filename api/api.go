package api


type API struct {
}

func NewAPI() (*API, error) {
  return &API{}, nil
}

func (a *API) Close() {
}
