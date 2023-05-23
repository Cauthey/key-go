package schema

type Group struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Scope       string   `json:"scope"`
	GiD         uint64   `json:"gid"`
	Member      []uint64 `json:"member"`
	Priv        string   `json:"priv"`
}
