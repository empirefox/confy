# confy xps

Use CLI tool `xps` to generate encrypted tarball, then load it using only one line code.

## Usage

1. get xps
```bash
go get -u github.com/empirefox/confy
```

2. generate xps tarball from `xps-config.json`
```bash
xps -k yourpassword
```

3. generate from `xps-config-dev.json`
```bash
xps -x xps-config-dev.json
```

4. test
```bash
go test
```

5. Extract tarball to prod
```bash
xps -d ./prod [-x xps-config.json] [-k password]
```

6. api to load app config file from xps tarball, equip with env, and validate
```go
package xps
type EnvLoadable interface {
  // load env to these ptrs
	GetEnvPtrs() []interface{}
}

type Validable interface {
  // overwrite default Validate func
	Validate(v interface{}) error
}
func LoadConfig(config interface{}, opts *ConfigOptions) (err error)
```

7. Parse ConfigOptions from env:
```go
import "github.com/empirefox/confy/xps"

config := new(Config)
err := xps.LoadConfig(config, nil)
```