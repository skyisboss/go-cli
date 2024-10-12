package config

import (
	"go-cli/internal/logs"
	"go-cli/internal/util"
	"io"
	"sort"
	"strings"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/olekukonko/tablewriter"
	"github.com/samber/lo"
)

type Config struct {
	// compile-time parameters
	GitCommit  string
	GitVersion string

	Env       string      `yaml:"env" env:"APP_ENV" env-default:"production" env-description:"Environment [production, local, sandbox]"`
	Debug     bool        `yaml:"debug" env:"APP_DEBUG" env-default:"false" env-description:"Enables debug mode"`
	Logger    logs.Config `yaml:"logger"`
	Providers Providers   `yaml:"providers"`
}
type Providers struct {
	SaltKey string      `yaml:"salt_key" env:"SALT_KEY"`
	Redis   RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Host string `yaml:"host" env:"REDIS_HOST" env-default:"127.0.0.1"`
	Port uint16 `yaml:"port" env:"REDIS_PORT" env-default:"6379"`
	Auth string `yaml:"auth" env:"REDIS_AUTH"`
	DB   int    `yaml:"db" env:"REDIS_DB"`
	Pool int    `yaml:"pool" env:"REDIS_POOL"`
}

var once = sync.Once{}
var cfg = &Config{}
var errCfg error

func New(gitCommit, gitVersion, configPath string) (*Config, error) {
	once.Do(func() {
		cfg = &Config{
			GitCommit:  gitCommit,
			GitVersion: gitVersion,
		}

		// if skipConfig {
		// 	errCfg = cleanenv.ReadEnv(cfg)
		// 	return
		// }

		errCfg = cleanenv.ReadConfig(configPath, cfg)
	})

	return cfg, errCfg
}

func PrintUsage(w io.Writer) error {
	desc, err := cleanenv.GetDescription(&Config{}, nil)
	if err != nil {
		return err
	}

	const delimiter = "||"

	// 1 line == 1 env var
	desc = strings.ReplaceAll(desc, "\n    \t", delimiter)

	lines := strings.Split(desc, "\n")

	// remove header
	lines = lines[1:]

	// hide internal vars
	lines = util.FilterSlice(lines, func(line string) bool {
		return !strings.Contains(strings.ToLower(line), "internal variable")
	})

	// remove duplicates
	lines = lo.Uniq(lines)

	// sort a-z (skip header)
	sort.Strings(lines[1:])

	// write as a table
	t := tablewriter.NewWriter(w)
	t.SetBorder(false)
	t.SetAutoWrapText(false)
	t.SetHeader([]string{"ENV", "Description"})
	t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	for _, line := range lines {
		cells := strings.Split(line, delimiter)
		cells = util.MapSlice(cells, strings.TrimSpace)
		t.Append(cells)
	}

	t.Render()

	return nil
}
