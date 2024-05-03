package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

type MenuOption struct {
	abbr byte
	text string
}

func (self Bill) Format() string {
	var output string
	var total float64
	output += fmt.Sprintf("%s's Bill Info\n", self.name)
	output += "----------------\n"

	for k, v := range self.items {
		var line string
		line = fmt.Sprintf("%12v %v\n", k+":", v)
		output += line
		total += v
	}

	output += fmt.Sprintf("Total: $%v", total)
	return output
}

func (self *Bill) ShowMenu() {
	var output string
	var lastIndex int
	options := []MenuOption{
		MenuOption{abbr: 'n', text: "New Item"},
		MenuOption{abbr: 's', text: "Save Bill"},
		MenuOption{abbr: 't', text: "Add tip"},
	}
	output += "Menu Options\n"

	for i, el := range options {
		line := fmt.Sprintf("%d) '%c': %s\n", i+1, el.abbr, el.text)
		output += line
		lastIndex = i + 1
	}

	output += "\n"
	output += fmt.Sprintf("%d) '%c': %s\n", lastIndex+1, 'q', "Quit")

	fmt.Println(output)
}

func (self *Bill) setItem(reader *bufio.Reader) error {
	item, err := GetInput(reader, "Enter item name:")
	if err != nil {
		return err
	}
	p, err := GetInput(reader, "Enter price:")
	if err != nil {
		return err
	}
	price, err := strconv.ParseFloat(p, 64)
	if err != nil {
		return err
	}

	self.items[item] = price
	return nil
}

func (self *Bill) Save() error {
	err := os.WriteFile("./bills/"+self.name, []byte(self.Format()), 0777)
	if err != nil {
		return err
	}
	fmt.Printf("Saved successfully to \"./bills/%s", self.name)
	return nil
}

func (self *Bill) setTip(reader *bufio.Reader) error {
	input, err := GetInput(reader, "How much tip?")
	if err != nil {
		return err
	}
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}

	self.tip = f
	return nil
}

func GetInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Printf("%s ", prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	fmt.Println("You entered", input)
	return input, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	billName, err := GetInput(reader, "Please enter bill's name:")
	if err != nil {
		log.Fatal(err)
	}
	bill := Bill{
		name:  billName,
		items: make(map[string]float64),
		tip:   0,
	}

	for {
		bill.ShowMenu()
		input, err := GetInput(reader, "")
		if err != nil {
			log.Fatal(err)
		}

		switch input {
		case "n":
			err := bill.setItem(reader)
			if err != nil {
				log.Fatal(err)
			}
		case "s":
			err := bill.Save()
			if err != nil {
				log.Fatal(err)
			}
		case "t":
			err := bill.setTip(reader)
			if err != nil {
				log.Fatal(err)
			}
		case "q":
			fmt.Println("Bye Bye")
			break
		default:
			fmt.Println("Command not found")
			continue
		}
	}
}
