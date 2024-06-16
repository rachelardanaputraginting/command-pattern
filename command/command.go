package command

type Interface interface{
	Execute()
}

type FriedRiceCommand struct {}

func (frc FriedRiceCommand) Execute() {
	
}