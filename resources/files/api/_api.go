// {{.Comment}}
package api

type Runnable interface {
	Run() error
}

type RunnableFactory interface {
	Create() Runnable
}
