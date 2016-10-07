# inspector
detect runtime env by hosts name

## usage :

```shell
//create config files as :
|- conf
    |--- app-dev.conf
	|--- app-product.conf
	|--- hosts.json
```

```json
//hosts.json
{
    "dev": [
	    "xzmRMBP.local"
    ],
	"product": [
	    "dsfds"
	]
}

```


```golang
import (
	"github.com/airylinus/inspector"
)

//...
func main() {
    //...
    beego.AddAPPStartHook(inspector.ConfigHook)
}

```
