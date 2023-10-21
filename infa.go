package main

import (
	"fmt"
)

// Furniture defines the properties of a piece of furniture.
type Furniture struct {
	name   string
	length int
	width  int
	height int
}

// Appliance defines the properties of a household appliance.
type Appliance struct {
	name   string
	length int
	width  int
	height int
}

// FamilyMember defines the properties of a family member.
type FamilyMember struct {
	name string
	age  int
	height int
	weight int
}

// House aggregates furniture, appliances, and family members.
type House struct {
	furniture    []Furniture
	appliances   []Appliance
	familyMembers []FamilyMember
}

func main() {
	// Define some furniture
	furniture := []Furniture{
		{name: "Couch", length: 200, width: 90, height: 85},
		{name: "Double Bed", length: 190, width: 140, height: 60},
		{name: "Dining Table", length: 150, width: 90, height: 75},
		{name: "Wooden Chair", length: 45, width: 45, height: 90},
		{name: "Bookshelf", length: 80, width: 30, height: 180},
	}

	// Define some appliances
	appliances := []Appliance{
		{name: "Microwave Oven", length: 50, width: 35, height: 30},
		{name: "Gas Stove", length: 60, width: 60, height: 85},
		{name: "Laundry Machine", length: 60, width: 60, height: 85},
		{name: "LED TV", length: 120, width: 10, height: 70},
		{name: "Hand Dryer", length: 10, width: 8, height: 25},
	}

	// Define some family members
	familyMembers := []FamilyMember{
		{name: "Mom", age: 40, height: 170, weight: 60},
		{name: "Dad", age: 42, height: 180, weight: 80},
		{name: "Sister", age: 18, height: 165, weight: 55},
		{name: "Grandma", age: 70, height: 160, weight: 65},
		{name: "Brother", age: 15, height: 175, weight: 70},
	}

	// Aggregate everything into a house
	house := House{
		furniture:     furniture,
		appliances:    appliances,
		familyMembers: familyMembers,
	}

	// Print the details of the house
	fmt.Println("Furniture:")
	for _, f := range house.furniture {
		fmt.Printf("%s: %dcm x %dcm x %dcm\n", f.name, f.length, f.width, f.height)
	}

	fmt.Println("\nAppliances:")
	for _, a := range house.appliances {
		fmt.Printf("%s: %dcm x %dcm x %dcm\n", a.name, a.length, a.width, a.height)
	}

	fmt.Println("\nFamily Members:")
	for _, fm := range house.familyMembers {
		fmt.Printf("%s: %d years old, %dcm tall, %dkg\n", fm.name, fm.age, fm.height, fm.weight)
	}
}
