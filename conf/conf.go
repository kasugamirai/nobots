package conf

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/kr/pretty"
	"github.com/rs/zerolog/log"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env            string
	TwtIds         Twt            `yaml:"twt"`
	Relays         Relays         `yaml:"relays"`
	Discord        Discord        `yaml:"discord"`
	Nostr          Nostr          `yaml:"nostr"`
	TwitterProfile TwitterProfile `yaml:"twitterProfile"`
	Sqlite         Sqlite         `yaml:"sqlite"`
	Report         Report         `yaml:"report"`
	Postgres       Postgres       `yaml:"postgres"`
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	SSLMode  string
}

type Report struct {
	PublicKeys []string `yaml:"publicKeys"`
	UserNames  []string `yaml:"userNames"`
}

type Sqlite struct {
	DSN string `yaml:"dsn"`
}

type Nostr struct {
	PubKey  string `yaml:"pubKey"`
	EventID string `yaml:"eventID"`
}

type Discord struct {
	Token     string `yaml:"token"`
	ChannelId string `yaml:"channelID"`
}

type BindMainDir struct {
	Dir string `json:"Dir"`
}

type Relays struct {
	Urls []string `yaml:"urls"`
}

type Twt struct {
	UserIds    []string `yaml:"users"`
	VipUserIds []string `yaml:"vipUsers"`
}

type TwitterProfile struct {
	Name      string `json:"name"`
	UserName  string `json:"userName"`
	Website   string `json:"website"`
	Avatar    string `json:"avatar"`
	Banner    string `json:"banner"`
	Biography string `json:"biography"`
	Location  string `json:"location"`
	URL       string `json:"url"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	confFileRelPath := getConfAbsPath()
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		log.
			Error().
			Msgf("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		log.
			Error().
			Msgf("validate config error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()

	pretty.Printf("%+v\n", conf)
}

func getConfAbsPath() string {
	cmd := exec.Command("go", "list", "-m", "-json")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	bindDir := &BindMainDir{}
	if err := json.Unmarshal(out.Bytes(), bindDir); err != nil {
		panic(err)
	}

	prefix := "conf"
	return filepath.Join(bindDir.Dir, prefix, filepath.Join(GetEnv(), "conf.yaml"))
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}
