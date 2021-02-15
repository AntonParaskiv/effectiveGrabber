package TaskGenerator

type Generator struct {
	taskOneChan chan struct{}
	taskTwoChan chan struct{}
}

func New() (g *Generator) {
	g = new(Generator)
	return
}

func (g *Generator) SetTaskOneChan(taskChan chan struct{}) *Generator {
	g.taskOneChan = taskChan
	return g
}

func (g *Generator) SetTaskTwoChan(taskChan chan struct{}) *Generator {
	g.taskTwoChan = taskChan
	return g
}

func (g *Generator) Start() {
	// start generate task one
	if g.taskOneChan != nil {
		go g.startTask(g.taskOneChan)
	}

	// start generate task two
	if g.taskTwoChan != nil {
		go g.startTask(g.taskTwoChan)
	}
}

func (g *Generator) startTask(taskChan chan struct{}) {
	// infinite send tasks
	for {
		taskChan <- struct{}{}
	}
}
