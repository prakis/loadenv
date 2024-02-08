package loadenv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Test(name string) string {

	message := "Test by " + name

	return message
}

func Load(envfiles ...string) error {

	envfiles = defaultCheck(envfiles)

	for _, envfilename := range envfiles {
		if len(envfilename) == 0 {
			continue
		}
		err := loadSingeEnvFile(envfilename)

		if err != nil {
			println(err)
			return err
		}
	}

	return nil
}

func defaultCheck(envfiles []string) []string {
	if len(envfiles) == 0 {
		return []string{".env"}
	}

	return envfiles
}

func loadSingeEnvFile(envfilename string) error {

	envMap := make(map[string]string)

	file, err := os.Open(envfilename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		// check if the line has = sign
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				// assign the map
				envMap[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	// load this map to os env
	for envKey, envVal := range envMap {
		_ = os.Setenv(envKey, envVal)
	}

	return nil
}
