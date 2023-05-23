package schema

type Group struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Scope       string `json:"scope"`
	GiD         uint   `json:"gid"`
	Member      []uint `json:"member"`
	Priv        string `json:"priv"`
}
