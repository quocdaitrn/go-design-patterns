package builder

import "testing"

func TestManufacturer(t *testing.T)  {
	director := Director{}

	homeEditionPCBuilder := &HomeEditionPCBuilder{}
	director.SetBuilder(homeEditionPCBuilder)
	director.ConstructPC()
	homeEditionPC := homeEditionPCBuilder.Build()

	if homeEditionPC.CPU != "Intel Core i5" {
		t.Errorf("Home edition PC cpu should be 'Intel core i5' but found %s", homeEditionPC.CPU)
	}

	if homeEditionPC.OS != "Window 11 Home Edition" {
		t.Errorf("Home edition PC OS shoul be 'Window 11 Home Edition' but found %s", homeEditionPC.OS)
	}

	gameEditionPCBuilder := &GameEditionPCBuilder{}
	director.SetBuilder(gameEditionPCBuilder)
	director.ConstructPC()
	gameEditionPC := gameEditionPCBuilder.Build()

	if gameEditionPC.CPU != "Intel Core i9" {
		t.Errorf("Game edition PC cpu should be 'Intel core i9' but found %s", gameEditionPC.CPU)
	}

	if gameEditionPC.OS != "Window 11 Ultimate" {
		t.Errorf("Game edition PC gpu shoul be 'Window 11 Ultimate' but found %s", gameEditionPC.OS)
	}
}