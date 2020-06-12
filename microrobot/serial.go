package microrobot

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

func ListSerial() map[string]string {
	result := make(map[string]string)
	//Open the key
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\\DEVICEMAP\\SERIALCOMM`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	//Get the key info
	ki, err := k.Stat()
	if err != nil {
		log.Fatal(err)
	}

	//Get the value count
	s, err := k.ReadValueNames(int(ki.ValueCount))
	if err != nil {
		log.Fatal(err)
	}

	for _, test := range s {
		q, _, err := k.GetStringValue(test)
		if err != nil {
			log.Fatal(err)
		}
		result[q] = test
	}

	return result
}
