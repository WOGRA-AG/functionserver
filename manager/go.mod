module wogra.com/manager

go 1.18

replace wogra.com/executor => ../executor

replace wogra.com/config => ../config

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.1.2
	golang.org/x/exp v0.0.0-20220602145555-4a0574d9293f
	wogra.com/config v0.0.0-00010101000000-000000000000
	wogra.com/executor v0.0.0-00010101000000-000000000000
)

require (
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/robertkrimen/otto v0.0.0-20211024170158-b87d35c0b86f // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)
