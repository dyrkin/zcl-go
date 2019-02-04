package cluster

type ResetToFactoryDefaultsCommand struct {
}

type IdentifyCommand struct {
	IdentifyTime uint16
}

type IdentifyQueryCommand struct{}

type TriggerEffectCommand struct {
	EffectIdentifier uint8
	EffectVariant    uint8
}

type IdentifyQueryResponse struct {
	Timeout uint16
}

type MoveToLevelCommand struct {
	Level          uint8
	TransitionTime uint16
}

type MoveCommand struct {
	MoveMode uint8
	Rate     uint8
}

type StepCommand struct {
	StepMode       uint8
	StepSize       uint8
	TransitionTime uint16
}

type StopCommand struct{}

type MoveToLevelOnOffCommand struct {
	Level          uint8
	TransitionTime uint16
}

type MoveOnOffCommand struct {
	MoveMode uint8
	Rate     uint8
}

type StepOnOffCommand struct {
	StepMode       uint8
	StepSize       uint8
	TransitionTime uint16
}

type StopOnOffCommand struct{}
