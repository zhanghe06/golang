package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	message.SetString(language.Greek, "%s went to %s.", "%s πήγε στήν %s.")
	message.SetString(language.AmericanEnglish, "%s went to %s.", "%s is in %s.")
	message.SetString(language.Greek, "%s has been stolen.", "%s κλάπηκε.")
	message.SetString(language.AmericanEnglish, "%s has been stolen.", "%s has been stolen.")
	message.SetString(language.Greek, "How are you?", "Πώς είστε?.")
}

func main() {
	p := message.NewPrinter(language.Greek)
	p.Printf("%s went to %s.", "Ο Πέτρος", "Αγγλία")
	fmt.Println()
	p.Printf("%s has been stolen.", "Η πέτρα")
	fmt.Println()
	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("%s went to %s.", "Peter", "England")
	fmt.Println()
	p.Printf("%s has been stolen.", "The Gem")
}
