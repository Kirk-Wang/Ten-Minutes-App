package infrastructure

import (
	log "github.com/sirupsen/logrus"
	"reflect"
)

type Runer struct {
	engineCtx EngineContext
}

func New() *Runer {
	e := &Runer{engineCtx: EngineContext{}}
	return e
}

func (r *Runer) Start() {
	r.init()
	r.setup()
	r.start()
}

func (r *Runer) init() {
	log.Info("Initializing engines...")
	for _, v := range GetEngines() {
		typ := reflect.TypeOf(v)
		log.Debugf("Initializing: type=%s", typ.String())
		v.Init(e.engineCtx)
	}
}

func (e *Runer) setup() {
	log.Info("Setup engines...")
	for _, v := range GetEngines() {
		typ := reflect.TypeOf(v)
		log.Debug("Setup: ", typ.String())
		v.Setup(e.engineCtx)
	}
}

func (e *Runer) start() {
	log.Info("Starting engines...")
	for i, v := range GetEngines() {
		typ := reflect.TypeOf(v)
		log.Debug("Starting: ", typ.String())
		if v.StartBlocking() {
			if i+1 == len(GetEngines()) {
				v.Start(e.engineCtx)
			} else {
				go v.Start(e.engineCtx)
			}
		} else {
			v.Start(e.engineCtx)
		}
	}
}

func (e *Runer) Stop() {
	log.Info("Stoping engines...")
	for _, v := range GetEngines() {
		typ := reflect.TypeOf(v)
		log.Debug("Stoping: ", typ.String())
		v.Stop(e.engineCtx)
	}
}
