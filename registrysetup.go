package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

func RegisterProtocol() {
	k, _, err := registry.CreateKey(registry.CLASSES_ROOT, "gbp\\shell\\open\\command", registry.ALL_ACCESS) // create the necessary keys in HKEY_CLASSES_ROOT
	if err != nil {
		log.Fatalf("there was an error creating the key :  %s", err.Error())
	}
	gbpkey, err := registry.OpenKey(registry.CLASSES_ROOT, "gbp", registry.ALL_ACCESS) //open the gbp key to set default value to "URL:gbp" and add a string empty value named "URL Protocol"
	if err != nil {
		log.Fatalf("there was an error creating the key :  %s", err.Error())
	}
	gbpkey.SetStringValue("", "URL:gbp")
	gbpkey.SetStringValue("URL Protocol", "")
	defer gbpkey.Close()
	defer k.Close()
	k.SetStringValue("", "\""+GetExecutablePath()+"\" -open "+" \"%1\"") // setting up default value in command subkey, we need executable full path + two arguments : "-open" and an argument for the link passed by the browser

	fmt.Printf("Install operation was a success.")

}
func UnregisterProtocol() {
	//deleting all the keys recursively
	err := registry.DeleteKey(registry.CLASSES_ROOT, "gbp\\shell\\open\\command")
	if err != nil {
		log.Fatalf("An error occured when trying to delete the registry key 'command' : %s", err.Error())
	}
	err = registry.DeleteKey(registry.CLASSES_ROOT, "gbp\\shell\\open")
	if err != nil {
		log.Fatalf("An error occured when trying to delete the registry key 'open' : %s", err.Error())
	}
	err = registry.DeleteKey(registry.CLASSES_ROOT, "gbp\\shell")
	if err != nil {
		log.Fatalf("An error occured when trying to delete the registry key 'shell': %s", err.Error())
	}
	err = registry.DeleteKey(registry.CLASSES_ROOT, "gbp")
	if err != nil {
		log.Fatalf("An error occured when trying to delete the registry key 'gbp' : %s", err.Error())
	}
	fmt.Printf("Uninstall operation was a success.")
}
func GetExecutablePath() string {
	exepath, err := os.Executable() //retriving the full path of the binary
	if err != nil {
		log.Fatalf("Counldn't retrive executable path : %s", err.Error())
	}
	return exepath
}
