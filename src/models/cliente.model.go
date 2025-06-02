package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cliente struct {
    ID                  primitive.ObjectID          `json:"id,omitempty" bson:"_id,omitempty"`
    Clave_Cliente       any                         `json:"clave_cliente" bson:"clave_cliente"`
    Nombre              string                      `json:"nombre" bson:"nombre"`
    Celular             string                      `json:"celular" bson:"celular"`
    Email               string                      `json:"email" bson:"email"`
    Character_Icon      any                         `json:"character_icon" bson:"character_icon"`
}
