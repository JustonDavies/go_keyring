//-- Package Declaration -----------------------------------------------------------------------------------------------
package keyring

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"errors"
	"github.com/JustonDavies/go_keyring/pkg/linux_secret_service"
	"runtime"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------
type Keyring interface {
	GetSecretByAttribute(string, string) (string, error)
}

//-- Exported Functions ------------------------------------------------------------------------------------------------
func New() (Keyring, error) {
	var chain Keyring
	//-- Determine OS-specific Data Path ----------
	{
		switch runtime.GOOS {
		case `linux`:
			chain = new(linuxSecretService.SecretService)
		//case `darwin`:
		//	chain = new()
		//case `windows`:
		//	c.dataPath = CHROME_WINDOWS_DATA_PATH
		default:
			return nil, errors.New(`unsupported operating system`)
		}
	}

	//-- Return ----------
	return chain, nil
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
