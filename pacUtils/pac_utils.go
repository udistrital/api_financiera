package pacUtils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/udistrital/api_financiera/models"
)

type WorkRequest struct {
	JobParameter interface{}
	Job          func(...interface{}) interface{}
}

var WorkQueue = make(chan WorkRequest, 100)
var WorkerQueue chan chan WorkRequest

func failOnError(err error, msg string) {
	if err != nil {
		beego.Info("%s: %s", msg, err)
		beego.Info(fmt.Sprintf("%s: %s", msg, err))
	}
}

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				work.Job(work.JobParameter)
			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		beego.Info("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				beego.Info("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					beego.Info("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}

func FunctionBeforeStatic(ctx *context.Context) {
	beego.Info("beego.BeforeStatic: Before finding the static file")
}
func FunctionBeforeRouter(ctx *context.Context) {
	beego.Info("beego.BeforeRouter: Executing Before finding router")
}
func FunctionBeforeExec(ctx *context.Context) {

	beego.Info("beego.BeforeExec: After finding router and before executing the matched Controller")
}

func FunctionAfterExecIngresoPac(ctx *context.Context) {
	//beego.Info("Llamada...")
	var u map[string]interface{}
	ingreso := models.Ingreso{}
	FillStruct(ctx.Input.Data()["json"], &u)
	FillStruct(u["Body"], &ingreso)
	//work := WorkRequest{JobParameter: ingreso, Job: FunctionJobExample}
	work := WorkRequest{JobParameter: ingreso, Job: (models.AddIngresoPac)}
	// Push the work onto the queue.
	WorkQueue <- work
	beego.Info("Work request queued")
}

func FunctionJobExample(parameter ...interface{}) (res interface{}) {
	beego.Info("Job's Parameter: ", parameter[0].(models.Ingreso).Id)
	return
}

func FillStruct(m interface{}, s interface{}) (err error) {
	j, _ := json.Marshal(m)
	err = json.Unmarshal(j, s)
	return
}

func FillStructDeep(m map[string]interface{}, fields string, s interface{}) (err error) {
	f := strings.Split(fields, ".")
	if len(f) == 0 {
		err = errors.New("invalid fields.")
		return
	}

	var aux map[string]interface{}
	var load interface{}
	for i, value := range f {

		if i == 0 {
			//fmt.Println(m[value])
			FillStruct(m[value], &load)
		} else {
			FillStruct(load, &aux)
			FillStruct(aux[value], &load)
			//fmt.Println(aux[value])
		}
	}
	j, _ := json.Marshal(load)
	err = json.Unmarshal(j, s)
	return
}

func FunctionFinishRouter(ctx *context.Context) {
	beego.Info("beego.FinishRouter: After finishing router")
}

func Init() {
	StartDispatcher(1)
	beego.InsertFilter("/v1/ingreso/AprobarIngreso", beego.AfterExec, FunctionAfterExecIngresoPac, false)
}

func sendJson(urlp string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		json.NewEncoder(b).Encode(datajson)
	}
	//proxyUrl, err := url.Parse("http://10.20.4.15:3128")
	//http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{}
	req, err := http.NewRequest(trequest, urlp, b)
	r, err := client.Do(req)
	//r, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getJson(urlp string, target interface{}) error {
	//proxyUrl, err := url.Parse("http://10.20.4.15:3128")
	//http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	r, err := http.Get(urlp)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
