package store

type Store struct {
	Uuid     string    `json:"uuid" bson:"_id"`
	Title    string    `json:"title"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Uuid        string       `json:"uuid"`
	Title       string       `json:"title"`
	Subtitle    string       `json:"subtitle"`
	IsOnSale    bool         `json:"isOnSale"`
	Subsections []Subsection `json:"subsections"`
}

type Subsection struct {
	Uuid     string `json:"uuid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Items    []Item `json:"items"`
}

type Item struct {
	Uuid           string          `json:"uuid"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	Price          int             `json:"price"`
	ImageUrl       string          `json:"imageUrl"`
	Customizations []Customization `json:"customization"`
}

type Customization struct {
	Uuid         string   `json:"uuid"`
	Title        string   `json:"title"`
	MinPermitted int      `json:"min_permitted"`
	MaxPermitted int      `json:"max_permitted"`
	Options      []Option `json:"options"`
}

type Option struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
	Price int    `json:"price"`
}