package model

type Portfolio struct {
	UID       string `json:"uid" bson:"uid"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	UpdatedBy string `json:"updated_by" bson:"updated_by"`
}

func (Portfolio) IsEntity() {}