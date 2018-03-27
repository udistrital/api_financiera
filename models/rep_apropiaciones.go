package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ReporteRubro struct {
	Codigo                           string
	Nombre                           string
	SaldoInicial                     int
	ModificacionMes                  int
	TotalModificaciones              int
	SaldoVigente                     int
	SaldoSuspencion                  int
	SaldoDisponible                  int
	CompromisoMes                    int
	TotalCompromisos                 int
	PorcentajeCompromisosApropiacion float32
	AutorizacionMes                  int
	AutorizacionesTotales            int
	PorcentajeGirosApropiacion       float32
}

type Reporte struct {
	Anio            int
	Mes             int
	UnidadEjecutora int
	Entidad         int
	Rubros          []ReporteRubro
}

func Register() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("kronos").C("repEjecucionGastos")
	Rubro3 := &ReporteRubro{
		"33-01", "pago de servicios públicos", 400, 25, 25, 350, 0, 325, 100, 100, 30.7,
		20, 50, 14.2}

	Rubro4 := &ReporteRubro{
		"33-02", "pago de servicios privados", 500, 0, 0, 500, 0, 500, 100, 100, 20.0, 0,
		0, 0.0}

	var rubros []ReporteRubro
	rubros = append(rubros, *Rubro3)
	rubros = append(rubros, *Rubro4)
	err = c.Insert(&Reporte{2018, 04, 1, 230, rubros})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Se registro bien")
	}

	// result := Reporte{}
	// err = c.Find(bson.M{"anio": "2018", "es": "4"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Vigencia: ", result.Anio, " - ", result.Mes)
}

func GetRepApropiacionesById(anio, mes, unidadejecutora int) (r *Reporte, err error) {
	session, err := mgo.Dial("localhost")
	c := session.DB("kronos").C("repEjecucionGastos")
	result := *Reporte{}
	log.Println(anio, mes, unidadejecutora)
	err = c.Find(bson.M{"anio": anio, "mes": mes, "unidadejecutora": unidadejecutora}).One(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}
