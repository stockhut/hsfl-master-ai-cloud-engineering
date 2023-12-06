package worker

type Job[R any] interface {
	Do() R
}

type JobFunc[R any] func() R

func (f JobFunc[int]) Do() int {
	return f()
}
