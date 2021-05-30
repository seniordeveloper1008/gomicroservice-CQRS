package infrastructures

import (
	sLog "log"

	"bitbucket.org/burhanmubarok/microservice/structures/infrastructures"
	"github.com/fsnotify/fsnotify"
	config "github.com/spf13/viper"
)

// Configuration doc
type Configuration struct {
	Name    string
	Desc    string
	Env     string
	Port    int
	Debug   bool
	LogPath string
}

// Init doc
func (c *Configuration) Init() {
	c.set()
}

func (c *Configuration) set() {
	config.SetConfigName("app")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	errorRead := config.ReadInConfig()
	if errorRead != nil {
		logData := structures.Log{Message: "Failed to read configuration file"}
		log := &Logger{}
		log.Warning(logData)
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		logData := structures.Log{Message: "app.yaml just changed"}
		log := &Logger{}
		log.Warning(logData)
		sLog.Println("Warning: app.yaml just changed")
	})
	c.Name = config.GetString("microservice.name")
	c.Desc = config.GetString("microservice.desc")
	c.Env = config.GetString("microservice.env")
	c.Port = config.GetInt("microservice.port")
	c.Debug = config.GetBool("microservice.debug")
	c.LogPath = config.GetString("microservice.logPath")
}

// Get doc
func (c *Configuration) Get() {

}
