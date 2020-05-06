package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
)

var Filename Basefile = "config.json"

type SettingsContainer interface {
	GetConfigDir() Dir
	GetBasefile() Basefile
	GetAppSlug() Slug
	GetDefaultConfig() SerializedConfig
}

type Config struct {
	ConfigDir Dir       `json:"-"`
	DataDir   Dir       `json:"data_dir"`
	CacheDir  Dir       `json:"cache_dir"`
	SettingsContainer   `json:"settings"`
}


func (me * Config) GetConfigFile() Filepath {
	return fmt.Sprintf("%s%c%s",
		me.GetConfigDir(),
		os.PathSeparator,
		me.GetBasefile(),
	)
}

func LoadConfig(settings SettingsContainer) (config *Config) {
	var err error
	for range Once {

		config = &Config{
			SettingsContainer: settings,
		}

		config.ConfigDir = expandConfigDir(settings.GetConfigDir())
		err = maybeMakeDir(config.ConfigDir,os.ModePerm)
		if err != nil {
			break
		}

		var b []byte
		b,err = config.loadConfigFile()
		if err != nil {
			break
		}

		err = json.Unmarshal(b, &config)
		if err != nil {
			break
		}

		if config.DataDir != "" {
			config.DataDir = maybeExpandDir(config.DataDir)
		} else {
			config.DataDir = getDataDir(settings.GetAppSlug())
		}

		if config.CacheDir != "" {
			config.CacheDir = maybeExpandDir(config.CacheDir)
		} else {
			config.CacheDir = getCacheDir(settings.GetAppSlug())
		}

	}
	if err != nil {
		log.Fatalf("Config file '%s' cannot be processed. It is likely invalid JSON or is not using the correct schema: %s.",
			getFilepath(config.ConfigDir, settings.GetBasefile()),
			err,
		)
	}
	return config
}

func (me *Config) loadConfigFile() (b []byte, err error) {
	var isnew bool
	var f *os.File
	fp := getFilepath(me.ConfigDir,me.SettingsContainer.GetBasefile())
	for range Once {
		if fileExists(fp) {
			b, err = ioutil.ReadFile(fp)
			if err != nil {
				log.Fatalf("Config file '%s' exists but cannot be read: %s.", fp, err)
			}
			if string(b) == me.SettingsContainer.GetDefaultConfig() {
				isnew = true
			}
			break
		}
		if !dirExists(me.ConfigDir) {
			err := os.MkdirAll(me.ConfigDir, os.ModePerm)
			if err != nil {
				log.Fatalf("Cannot make directory '%s'; Check permissions: %s.", me.ConfigDir, err)
			}
		}
		f, err = os.Create(fp)
		if err != nil {
			log.Fatalf("Cannot create config file '%s'; Check permissions: %s.", fp, err)
		}
		var n int
		dc := me.SettingsContainer.GetDefaultConfig()
		n, err = f.WriteString(dc)
		if err != nil || n != len(dc) {
			log.Fatalf("Cannot create config file '%s'; Check permissions: %s.", fp, err)
		}
		var size int64
		size, err = f.Seek(0, 2)
		if err != nil || size != int64(len(dc)) {
			log.Fatalf("Cannot determine length of config file just written '%s'; Check permissions: %s", fp, err)
		}
		var n64 int64
		n64, err = f.Seek(0, 0)
		if err != nil || n64 != 0 {
			log.Fatalf("Cannot reset config file just written '%s'; Check permissions: %s.", fp, err)
		}
		b, err = ioutil.ReadAll(f)
		if err != nil || string(b) != dc {
			log.Fatalf("Config read does not equal config file just written '%s': %s.", fp, err)
		}
		isnew = true
	}
	closeFile(f)
	if isnew {
		fmt.Printf("\nYour config file '%s' is newly initialized.", fp)
		fmt.Printf("\nPlease EDIT to configure appropriate settings and rerun your command.\n")
		os.Exit(1)
	}
	return b,err
}

func closeFile(f *os.File) {
	_ = f.Close()
}

func getDataDir(appslug Slug) (cd Dir) {
	return fmt.Sprintf("%s%c%s",
		HomeDir(),
		os.PathSeparator,
		appslug,
	)
}

func getCacheDir(appslug Slug) (cd Dir) {
	for range Once {
		cd, err := os.UserCacheDir()
		if err != nil {
			if runtime.GOOS == "windows" {
				cd = "C:\\tmp"
			} else {
				cd = "/tmp"
			}
		}
		cd = fmt.Sprintf("%s%c%s",
			cd,
			os.PathSeparator,
			appslug,
		)
	}
	return cd
}

func getFilepath(configdir Dir,basefile Basefile) Filepath {
	if basefile != "" && Filename != basefile {
		Filename =basefile
	}
	return fmt.Sprintf("%s%c%s",
		configdir,
		os.PathSeparator,
		Filename,
	)
}

var dirRegexp *regexp.Regexp

func init() {
	dirRegexp = regexp.MustCompile(`^~/`)
}

func maybeExpandDir(dir Dir) Filepath {
	hd := fmt.Sprintf("%s%c", HomeDir(), os.PathSeparator)
	return dirRegexp.ReplaceAllString(dir, hd)
}

func expandConfigDir(configdir Dir) Filepath {
	hd := fmt.Sprintf("%s%c", HomeDir(), os.PathSeparator)
	cd := dirRegexp.ReplaceAllString(configdir, hd)
	return cd
}

