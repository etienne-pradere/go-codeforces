# Entrenamiento CodeForces

## Crear template en el editor de jetbrains 

- En el menu de **Preferences** 
  - En la opción **File and Code Template**
    - En **Files** crear un nuevo tipo **Name: go-cf** y **Extension: go**. Luego pegar el **Ejemplo de template** y listo.
---
  

### Ejemplo de template:

Este template define funciones utiles para la lectura estándar de los ejercicios de CodeForces utilizando el lenguaje de programación GO.

Si seguiste los pasos para crear un template podras crear un archivo nuevo, solo necesitas hacer clic derecho sobre el proyecto de GO y luego crear un nuevo archivo **go-cf** este tipo de archivo te solicitara la información del **USER** en el cual puedes ingresar tu nombre de usuario en CodeForces, **NAME** en el cual podrás definir el nombre del archivo que vas a crear y por último el campo **URL** donde podrás agrega la url de enlace al ejercicio que vas a resolver.


**HAPPY CODING :)**

```
package ${GO_PACKAGE_NAME}

import (
        "bufio"
        "fmt"
        "os"
        "sort"
        "strconv"
        "strings"
)
type ${NAME} struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *${NAME}) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
  fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *${NAME}) load() {
  if len(in.split) <= in.index {
  in.split = strings.Split(in.GetLine(), in.separator)
  in.index = 0
  }
}

func (in *${NAME}) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *${NAME}) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *${NAME}) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem ${NAME}
Date: ${DATE}
User: ${USER}
URL: ${URL}
Problem: ${NAME}
**/
func (in *${NAME}) Run(){

}

func New${NAME}(r *bufio.Reader) *${NAME} {
  return &${NAME}{
  sc:        r,
  split:     []string{},
  index:     0,
  separator: " ",
  }
}

func main() {
  New${NAME}(bufio.NewReader(os.Stdin)).Run()
}
```
