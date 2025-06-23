package models

type RegionTable struct {
	RegionID          int    `json:"region_id" gorm:"primaryKey;type:integer" csv:"region_id"`
	RegionName        string `json:"region_name" csv:"region_name"`
	RegionDescription string `json:"region_description" csv:"region_description"`
}

func (RegionTable) TableName() string {
	return "location_schema.region_table"
}

type ProvinceTable struct {
	ProvinceID   int         `json:"province_id" gorm:"primaryKey;type:integer" csv:"province_id"`
	ProvinceName string      `json:"province_name" csv:"province_name"`
	RegionID     int         `json:"region_id" gorm:"type:integer" csv:"region_id"`
	Region       RegionTable `gorm:"foreignKey:RegionID;references:RegionID" json:"-"`
	ProvinceRate float64     `json:"province_rate" csv:"province_rate"`
}

func (ProvinceTable) TableName() string {
	return "location_schema.province_table"
}

type MunicipalityTable struct {
	MunicipalityID   int           `json:"municipality_id" gorm:"primaryKey;type:integer" csv:"municipality_id"`
	MunicipalityName string        `json:"municipality_name" csv:"municipality_name"`
	ProvinceID       int           `json:"province_id" gorm:"type:integer" csv:"province_id"`
	Province         ProvinceTable `gorm:"foreignKey:ProvinceID;references:ProvinceID" json:"-"`
}

func (MunicipalityTable) TableName() string {
	return "location_schema.municipality_table"
}

type BarangayTable struct {
	BarangayID     int               `json:"barangay_id" gorm:"primaryKey;type:integer" csv:"barangay_id"`
	MunicipalityID int               `json:"municipality_id" gorm:"type:integer" csv:"municipality_id"`
	Municipality   MunicipalityTable `gorm:"foreignKey:MunicipalityID;references:MunicipalityID" json:"-"`
	BarangayName   string            `json:"barangay_name" csv:"barangay_name"`
}

func (BarangayTable) TableName() string {
	return "location_schema.barangay_table"
}
