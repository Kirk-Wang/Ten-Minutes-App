package infrastructure

import (
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/lotteryjs/Ten-Minutes-App/infrastructure/engines"
)

type Runer struct {
	engineCtx engines.EngineContext
}

func New() *Runer {
	return &Runer{engineCtx: engines.EngineContext{}}
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
		v.Init(r.engineCtx)
	}
}

func (r *Runer) setup() {
	log.Info("Setup engines...")
	for _, v := range GetEngines() {
		typ := reflect.TypeOf(v)
		log.Debug("Setup: ", typ.String())
		v.Setup(r.engineCtx)
	}
}

func (r *Runer) start() {
	log.Info("Starting engines...")
	for i, v := range GetEngines() {
		typ := reflect.TypeOf(v)
		log.Debug("Starting: ", typ.String())
		if v.StartBlocking() {
			if i+1 == len(GetEngines()) {
				v.Start(r.engineCtx)
			} else {
				go v.Start(r.engineCtx)
			}
		} else {
			v.Start(r.engineCtx)
		}
	}
}

func (r *Runer) Stop() {
	log.Info("Stoping engines...")
	for _, v := range GetEngines() {
		typ := reflect.TypeOf(v)
		log.Debug("Stoping: ", typ.String())
		v.Shutdown(r.engineCtx)
	}
}
