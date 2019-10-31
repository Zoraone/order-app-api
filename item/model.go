package item

type Item struct {
	Uuid               string          `json:"uuid" bson:"_id"`
	Title              string          `json:"title"`
	Price              int             `json:"price"`
	ItemDescription    string          `json:"itemDescription"`
	ImageUrl           string          `json:"imageUrl"`
	CustomizationsList []Customization `json:"customizationsList"`
}

type Customization struct {
	Uuid         string   `json:"uuid"`
	Title        string   `json:"title"`
	MinPermitted int      `json:"minPermitted"`
	MaxPermitted int      `json:"maxPermitted"`
	DisplayState string   `json:"displayState"`
	Options      []Option `json:"options"`
}

type Option struct {
	Uuid            string `json:"uuid"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	DefaultQuantity int    `json:"defaultQuantity"`
	MinPermitted    int    `json:"minPermitted"`
	MaxPermitted    int    `json:"maxPermitted"`
	IsSoldOut       bool   `json:"isSoldOut"`
}

type AddResponse struct {
	NewId interface{} `json:"newId"`
}
