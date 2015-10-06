package lib

import (
	"fmt"
	"os"
	"path"
	"strings"

	c "github.com/revel/config"
)

var (
	production = `production`
	filename   = `app.conf`
	directory  = `conf/`
)

// Models an entity config reader
type Config struct {
	Pwd          string
	Filename     string
	IsProduction bool
	*c.Config
}

// Crea una nueva estructura tipo Config
func NewConfig() (config *Config, err error) {
	config = &Config{Filename: filename}
	if config.Pwd, err = os.Getwd(); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		panic(err)
	}

	var file = config.File()
	//	fmt.Printf("App | Config will be loaded from %v \n", file)
	if config.Config, err = c.ReadDefault(file); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		panic(err)
	}
	//	fmt.Println("App | Config loaded successfully! \n")
	config.IsProduction = strings.EqualFold(config.Default("env"), production)
	return
}

// Regresa la ruta del archivo de configuración
func (self *Config) File() (file string) {
	return path.Join(self.Pwd, directory, self.Filename)
}

// Obtiene una propiedad de la configuración de base de datos
func (self *Config) Database(property string) (result string) {
	result, _ = self.String("database", property)
	return
}

// Obtiene una propiedad de la configuración de la sección default
func (self *Config) Default(property string) (result string) {
	result, _ = self.String("default", property)
	return
}
