//-- Package Declaration -----------------------------------------------------------------------------------------------
package linuxSecretService

//-- C Imports ---------------------------------------------------------------------------------------------------------

// #cgo pkg-config: libsecret-1
// #include "linux_secret_service.h"
import "C"

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"errors"
	"unsafe"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------
type SecretService struct {
}

//-- Exported Functions ------------------------------------------------------------------------------------------------
func (s *SecretService) GetSecretByAttribute(name string, value string) (string, error) {
	var nameC *C.char
	var valueC *C.char
	var secretC *C.char

	defer C.free(unsafe.Pointer(nameC))
	defer C.free(unsafe.Pointer(valueC))
	defer C.free(unsafe.Pointer(secretC))

	nameC = C.CString(name) //These don't seem to be freeing as they should...
	valueC = C.CString(value)

	if err := C.getSecretByAttribute(nameC, valueC, &secretC); err != nil {
		defer C.g_error_free(err)
		errMsg := (*C.char)(unsafe.Pointer(err.message))
		return ``, errors.New(C.GoString(errMsg))
	}

	var secret = C.GoString(secretC)

	return secret, nil
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
