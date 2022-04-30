package ipc

type Request struct {
	Activate int
	Context  int
	Complete int
	Quit     int
	Search   string
}

type Exit struct{}

type Interrupt struct{}

type Activate struct {
	Index int
}

type Complete struct {
	Index int
}

type Context struct {
	Index int
}

type Quit struct {
	Index int
}

type Search struct {
	Name string
}
