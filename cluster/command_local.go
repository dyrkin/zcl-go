package cluster

type ResetToFactoryDefaultsCommand struct {
}

type IdentifyCommand struct {
	IdentifyTime uint16
}

type IdentifyQuery struct{}

type TriggerEffect struct {
	EffectIdentifier uint8
	EffectVariant    uint8
}

type IdentifyQueryResponseCommand struct {
	Timeout uint16
}
