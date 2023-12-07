package worker

type Job[R any] interface {
	Do() R
}

type JobFunc[R any] func() R

func (f JobFunc[R]) Do() R {
	return f()
}

type JobFactory[R any] interface {
	Get() Job[R]
}

type JobFactoryFunc[R any] func() Job[R]

func (f JobFactoryFunc[R]) Get() Job[R] {
	return f()
}
