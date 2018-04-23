package models

import "gopkg.in/mgo.v2/bson"

type MongoApropiacion struct {
	Id            bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Vigencia      int           `json:"vigencia"`
	Valor_inicial int           `json:"valor_inicial"`
	Rubro         MongoRubro       	`json:"rubro"`
}

type MongoDisponibilidadApropiacion struct {
	Id bson.ObjectId 					`json:"_id" bson:"_id,omitempty"`
	Valor int 								`json:"valor"`
  Fuente_financiamiento int `json:"fuente_financiamiento"`
  Apropiacion MongoApropiacion 	`json:"apropiacion"`
}

type MongoMovimiento struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Numero string `json:"numero"`
  Estado_movimiento string `json:"estado_movimiento"`
  Fecha_movimiento string `json:"fecha_movimiento"`
  Numero_oficio string `json:"numero_oficio"`
  Fecha_oficio string `json:"fecha_oficio"`
  Descripcion string `json:"descripcion"`
  Unidad_ejecutora int `json:"unidad_ejecutora"`
  RubroDestino string `json:"rubro_destino"`
  RubroOrigen string `json:"rubro_origen"`
  Valor int64 `json:"valor"`
  Tipo_movimiento string `json:"tipo_movimiento"`
  Apropiacion MongoApropiacion `json:"apropiacion"`
}

type MongoRegistroPresupuestal struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Vigencia int `json:"vigencia"`
  FechaRegistro string `json:"fecha_registro"`
  Estado string `json:"estado"`
  Numero_Registro_Presupuestal int `json:"numero_registro_presupuestal"`
  Solicitud int`json:"solicitud"`
	Disponibilidad_apropiacion MongoDisponibilidadApropiacion `json:"disponibilidad_apropiacion"`
}

type MongoOrdenPago struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Vigencia int `json:"vigencia"`
  Valor_base int `json:"valor_base"`
  Unidad_ejecutora int `json:"unidad_ejecutora"`
  Forma_pago int `json:"forma_pago"`
  Registro_presupuestal MongoRegistroPresupuestal `json:"registro_presupuestals"`
}

type MongoRubro struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Codigo string `json:"codigo"`
  Nombre string `json:"nombre"`
  Entidad string `json:"entidad"`
  Descripcion string `json:"descripcion"`
  Unidad_ejecutora int `json:"unidad_ejecutora"`
  Hijos []string `json:"hijos"`
}
