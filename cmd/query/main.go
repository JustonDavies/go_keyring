//-- Package Declaration -----------------------------------------------------------------------------------------------
package main

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"log"
	"time"

	"github.com/JustonDavies/go_keyring/pkg/keyring"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------

//-- Exported Functions ------------------------------------------------------------------------------------------------
func main() {
	//-- Log nice output ----------
	var start = time.Now().Unix()
	log.Println(`Starting task...`)

	//-- Perform task ----------
	var chain, _ = keyring.New()

	if secret, err := chain.GetSecretByAttribute(`application`, `chrome`); err != nil {
		log.Println(`Error: `, err)
	} else {
		log.Printf(`I got the secret: %s`, secret)
	}

	//-- Log nice output ----------
	log.Printf(`Task complete! It took %d seconds`, time.Now().Unix()-start)
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
