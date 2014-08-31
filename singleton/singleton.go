package singleton

type Obj map[string]interface{}

type Singleton struct {
	Data Obj
}

var instance *Singleton = nil

func NewSingleton() *Singleton {
	if instance == nil {
		instance = &Singleton{};
		instance.Data = make(Obj)
	}
	return instance;
}
