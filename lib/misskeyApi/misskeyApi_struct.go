package misskeyapi

type GetnotesRequest struct {
	Limit int    `json:"limit"`
	Token string `json:"i"`
}

type NotesResponse struct {
	Text   string `json:"text"`
	Renote Renote `json:"renote"`
	User   User   `json:"user"`
}

type User struct {
	Id string `json:"id"`
}

type Renote struct {
	Text string `json:"text"`
}

type NotesRequest struct {
	Visibility string `json:"visibility"`
	Text       string `json:"text"`
	Token      string `json:"i"`
	LocalOnly  bool   `json:"localOnly"`
}

type GetIDRequest struct {
	Token string `json:"i"`
}

type IDResponse struct {
	ID string `json:"id"`
}
