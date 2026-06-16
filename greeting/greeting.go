package greeting
import (
	"fmt" 
	"errors"
	"math/rand"
)

func PrintName(name string) ( string, error) {



	if name == "" {
		return "", errors.New("Name is required")
	}

	name2 := fmt.Sprintf(nameformater(),name)

	return name2, nil


}

func MultiplePrints(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, error := PrintName(name)

		if error != nil{
			return nil, error
		}

		messages[name] = message
	}
	return messages, nil
}


func nameformater() string{
	formates := []string{
		"Hi %v, whats up?",
		"Hello %v, how are you?",
		"Hey %v, how's it going?",
		"Good day %v, how are you doing?",
		"Hi %v, how's your day going?",
	}
	formate := formates[rand.Intn(len(formates))]
	return formate
}