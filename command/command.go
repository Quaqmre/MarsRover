package command

import "fmt"

//Gelen emirleri uygulamak ve ayırmak için kullanılır
//Tek yaptığı iş emir bulmaktır.
type Command struct {
	allCommands         string
	currentCommand      rune
	currentCommandIndex int
}

func NewCommand(com string) *Command {
	cm := Command{
		allCommands:         com,
		currentCommand:      0,
		currentCommandIndex: 0,
	}

	return &cm
}

//Sıradaki emiri alır.
func (cm *Command) GetCurrentCommand() rune {
	return cm.currentCommand
}

// Yerine getirilen emir sonrası yeni emir bulur.
func (cm *Command) GetNextCommand() error {
	if cm.currentCommandIndex == len(cm.allCommands) {
		return fmt.Errorf("end of the plato len(%v) - commandIndex(%v)", len(cm.allCommands), cm.currentCommandIndex)
	}
	cm.currentCommand = rune(cm.allCommands[cm.currentCommandIndex])
	cm.currentCommandIndex++
	return nil

}
