package repo

import (
	"devops-tpl/internal/entity"
	"fmt"
	"sync"
)

type MetricRepo struct {
	data  map[string]interface{}
	Mutex *sync.Mutex
}

func New() *MetricRepo {
	metricRepo := MetricRepo{
		Mutex: &sync.Mutex{},
	}
	metricRepo.data = make(map[string]interface{})
	return &metricRepo
}

func (r *MetricRepo) GetMetricNames() []string {
	var list []string
	for name := range r.data {
		list = append(list, name)
	}
	return list
}

func (r *MetricRepo) StoreGauge(name string, value entity.Gauge) error {
	r.Mutex.Lock()
	r.data[name] = value
	r.Mutex.Unlock()
	return nil
}

func (r *MetricRepo) AddCounter(name string, value entity.Counter) error {
	r.Mutex.Lock()
	oldValue, ok := r.data[name]
	r.Mutex.Unlock()
	if ok {
		r.data[name] = value + oldValue.(entity.Counter)
	} else {
		r.data[name] = value
	}
	return nil
}

func (r *MetricRepo) GetMetric(name string) (interface{}, error) {
	value, ok := r.data[name]
	if !ok {
		return nil, fmt.Errorf("not Found (%s)", name)
	}
	return value, nil
}
