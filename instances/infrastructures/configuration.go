package infrastructures

import (
	sLog "log"

	structures "github.com/burubur/microservice/structures/infrastructures"
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
	config.SetDefault("microservice.name", "Microservice")
	config.SetDefault("microservice.desc", "Microservice")
	config.SetDefault("microservice.env", "development")
	config.SetDefault("microservice.port", "8008")
	config.SetDefault("microservice.debug", "false")
	config.SetDefault("microservice.logPath", "./logs")
	errorRead := config.ReadInConfig()
	if errorRead != nil {
		logData := structures.Log{"Failed to read configuration file"}
		log := &Logger{}
		log.Warning(logData)
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		logData := structures.Log{"app.yaml just changed"}
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
