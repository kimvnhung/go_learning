package models

type HouseFilterField string

const (
	InvalidField HouseFilterField = "invalid"
	OwnerName    HouseFilterField = "owner_name"
	Phone        HouseFilterField = "phone"
	Address      HouseFilterField = "address"
	Price        HouseFilterField = "price"
	PriceUnit    HouseFilterField = "price_unit"
	BedRoom      HouseFilterField = "bed_room"
	LivingRoom   HouseFilterField = "living_room"
	HouseType    HouseFilterField = "house_type"
)

type HouseFilterType int

const (
	InvalidType HouseFilterType = iota
	Equal
	NotEqual
	Contain
	NotContain
	Greater
	GreaterOrEqual
	Less
	LessOrEqual
)

type House struct {
	ID               int     `json:"id" gorm:"primaryKey;autoIncrement=true"`
	OwnerName        string  `json:"owner_name" gorm:"unqiueIndex=house" default:"anonymous"`
	Phone            string  `json:"phone" gorm:"unqiueIndex=house"`
	Address          string  `json:"address" gorm:"unqiueIndex=house"`
	DistanceToTarget float32 `json:"distance_to_target" default:"0"`
	Price            string  `json:"price" default:"0"`
	PriceUnit        string  `json:"price_unit" default:"month"`
	BedRoom          int     `json:"bed_room"`
	LivingRoom       int     `json:"living_room"`
	HouseType        string  `json:"house_type" gorm:"unqiueIndex=house"`
}
