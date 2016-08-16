package main

import (
	"github.com/hailongz/kk-go-task/task"
	"log"
)

type DemoService struct {
	task.Service
	name string
}

func (S *DemoService) Handle(task task.ITask) error {
	return S.ReflectHandle(task, S)
}

func (S *DemoService) HandleTask(task task.ITask) error {
	log.Println("ok")
	return nil
}

func main() {

	log.SetFlags(log.Llongfile | log.LstdFlags)
	var C = task.NewContext()

	C.Service(&DemoService{})(&task.Task{})

	C.Handle(&task.Task{})

}
