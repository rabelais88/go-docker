package log

import "fmt"

// type Logger int
type Logger string

// const (
// 	ConcurrencyRace Logger = iota + 1
// 	ConcurrencyEnv
// 	ConcurrencyChannel
// 	Employee
// 	Myfile
// 	Shape
// )

// const LogTypes = ConcurrencyRace | ConcurrencyEnv | ConcurrencyChannel | Employee | Myfile | Shape

// var loggerDisplay = map[Logger]string{
// 	ConcurrencyRace:    `concurrency/race`,
// 	ConcurrencyEnv:     `concurrency/env`,
// 	ConcurrencyChannel: `concurrency/channel`,
// 	Employee:           `employee`,
// 	Myfile:             `myfile`,
// 	Shape:              `shape`,
// }

func (loggerType Logger) Print(arg ...interface{}) {
	// if _, ok := loggerDisplay[loggerType]; ok == false {
	// 	fmt.Println(`the logger type doesn't exist`, loggerType)
	// 	return
	// }
	// fmt.Println(loggerDisplay[loggerType], `->`, arg)
	fmt.Println(loggerType, `->`, arg)
}
