package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	lastIn := in
	for _, stage := range stages {
		lastIn = stage(lastIn)
	}

	out := make(Bi)
	if done != nil {
		go func() {
			defer close(out)
			for {
				select {
				case val := <-lastIn:
					out <- val
				case <-done:
					return
				}
			}
		}()
	} else {
		go func() {
			defer close(out)
			for val := range lastIn {
				out <- val
			}
		}()
	}

	return out
}
