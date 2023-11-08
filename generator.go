package main

import (
	"gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewKeyboard returns a keyboard with randomized values for the
// Layout and Backlit fields
func NewKeyboard() *gen.Keyboard {
	keyboard := &gen.Keyboard{
		layout:  randomKeyboardLayout(),
		backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *gen.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numCores := int32(randomInt(2, 8))
	numThreads := int32(randomInt(numCores, 12))
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	cpu := &gen.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   numCores,
		NumberThreads: numThreads,
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *gen.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := gen.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  gen.Memory_GIGABYTE,
	}

	return gen.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
}

func NewRAM() *gen.Memory {
	ram := &gen.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  gen.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *gen.Storage {
	ssd := &gen.Storage{
		Driver: gen.Storage_SSD,
		Memory: &gen.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  gen.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *gen.Storage {
	hdd := &gen.Storage{
		Driver: gen.Storage_HDD,
		Memory: &gen.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  gen.Memory_TERABYTE,
		},
	}

	return hdd
}

func NewScreen() *gen.Screen {
	screen := &gen.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *gen.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &gen.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*gen.GPU{NewGPU()},
		Storages: []*gen.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &gen.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}

	return laptop
}
