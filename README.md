# inspector
detect runtime env by hosts name

## usage :

```shell
//create config files as :
|- conf
    |--- app-dev.ini
	|--- app-product.ini
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
    inspector.Inspect("ini")
}

```
