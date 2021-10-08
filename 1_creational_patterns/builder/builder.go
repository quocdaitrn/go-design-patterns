package builder

// PersonalComputer represents a PC.
type PersonalComputer struct {
	CPU    string
	OS     string
	RamCap uint
	HddCap uint
}

// PCBuilder exposes all method of a PC builder.
type PCBuilder interface {
	SetCPU() PCBuilder
	SetOS() PCBuilder
	SetRamCap() PCBuilder
	SetHddCap() PCBuilder
	Build() *PersonalComputer
}

//--------------------------------------------------

// HomeEditionPCBuilder builder for Home Edition PC.
type HomeEditionPCBuilder struct {
	pc PersonalComputer
}

// SetCPU sets CPU for the Home Edition PC.
func (b *HomeEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.CPU = "Intel Core i5"
	return b
}

// SetCPU sets OS for the Home Edition PC.
func (b *HomeEditionPCBuilder) SetOS() PCBuilder {
	b.pc.OS = "Window 11 Home Edition"
	return b
}

// SetCPU sets RAM capacity for the Home Edition PC.
func (b *HomeEditionPCBuilder) SetRamCap() PCBuilder {
	b.pc.RamCap = 16
	return b
}

// SetCPU sets HDD capacity for the Home Edition PC.
func (b *HomeEditionPCBuilder) SetHddCap() PCBuilder {
	b.pc.HddCap = 500
	return b
}

// Build builds the Home Edition PC from builder.
func (b *HomeEditionPCBuilder) Build() *PersonalComputer {
	return &b.pc
}

//--------------------------------------------------

// GameEditionPCBuilder builder for Game Edition PC.
type GameEditionPCBuilder struct {
	pc PersonalComputer
}

// SetCPU sets CPU for the Game Edition PC.
func (b *GameEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.CPU = "Intel Core i9"
	return b
}

// SetCPU sets OS for the Game Edition PC.
func (b *GameEditionPCBuilder) SetOS() PCBuilder {
	b.pc.OS = "Window 11 Ultimate"
	return b
}

// SetCPU sets RAM capacity for the Game Edition PC.
func (b *GameEditionPCBuilder) SetRamCap() PCBuilder {
	b.pc.RamCap = 64
	return b
}

// SetCPU sets HDD capacity for the Game Edition PC.
func (b *GameEditionPCBuilder) SetHddCap() PCBuilder {
	b.pc.HddCap = 1000
	return b
}

// Build builds the Game Edition PC from builder.
func (b *GameEditionPCBuilder) Build() *PersonalComputer {
	return &b.pc
}

//--------------------------------------------------

// Director aware of builder process for builder type.
type Director struct {
	builder PCBuilder
}

func (d *Director) SetBuilder(builder PCBuilder) {
	d.builder = builder
}

func (d *Director) ConstructPC() {
	d.builder.SetCPU().SetOS().SetRamCap().SetHddCap()
}
