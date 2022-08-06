package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	FstabConfig map[string]Devices `yaml:"fstab"`
}
type Devices struct {
	MountPoint           string   `yaml:"mount"`
	Export               string   `yaml:"export"`
	FileSystemType       string   `yaml:"type"`
	RootReserve          string   `yaml:"root-reserve"`
	Options              []string `yaml:"options"`
	BackupOperation      int      `yaml:"backup"`
	FileSystemCheckOrder int      `yaml:"fs-check-order"`
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("fstab.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func write(filepath string, c Config) bool {
	// for fstab
	f, err := os.Create(filepath)
	check(err)

	defer f.Close()

	// for bash script
	tune2fs, err := os.Create("tune2fs.sh")
	check(err)
	defer tune2fs.Close()

	for device, parameters := range c.FstabConfig {
		// check options. if empty -> defaults
		var options string = ""
		for _, value := range parameters.Options {
			if options == "" {
				options = value
			} else {
				options = options + "," + value
			}
		}
		if options == "" {
			options = "defaults"
		}

		// check backup exist. if empty -> 0
		if parameters.BackupOperation != 0 || parameters.BackupOperation != 1 {
			parameters.BackupOperation = 0
		}
		// fs-check-order check if in range
		if parameters.FileSystemCheckOrder < 0 || parameters.FileSystemCheckOrder > 2 {
			parameters.FileSystemCheckOrder = 0
		}

		_, err := f.WriteString(
			device + " " + parameters.MountPoint + "\t" + parameters.FileSystemType + "\t" + options + "\t" + strconv.Itoa(parameters.BackupOperation) + "\t" + strconv.Itoa(parameters.FileSystemCheckOrder) + "\n",
		)
		check(err)
		// fmt.Printf("wrote %d bytes\n", config_line)

		// sudo tune2fs -m 3 /dev/sda1
		RootReserve := ""
		if parameters.RootReserve != "" {
			RootReserve = parameters.RootReserve[:strings.IndexByte(parameters.RootReserve, '%')]
			// fmt.Println("RootReserve = ", RootReserve)

			_, err = tune2fs.WriteString("sudo tune2fs -m " + RootReserve + " " + device)
			check(err)
		}
	}

	_ = f.Sync() // Sync content of the file with disk, before it can be in system cash and possible lost (for reboot situation)

	return true
}

func main() {
	var c Config
	c.GetConf()

	result := write("fstab_test", c)

	if result {
		fmt.Println("File fstab written successfully.")
	} else {
		fmt.Println("File fstab fatal.")
	}
}
