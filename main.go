package main

import "github.com/drone/drone-plugin-go/plugin"

type Marathon struct {
	//Hostname for the Marathon Master
	Host string `json:"host"`

	// The app config for marathon
	//https://mesosphere.github.io/marathon/docs/rest-api.html#post-v2-apps
	// Examples:
	//    /path/to/file
	//    /path/to/*.txt
	//    /path/to/*/*.txt
	//    /path/to/**
	ConfigFile string               `json:"config_file"`
	Condition  *condition.Condition `json:"when"`
}

func main() {
	vargs := Marathon{}
	plugin.Parse()

}
