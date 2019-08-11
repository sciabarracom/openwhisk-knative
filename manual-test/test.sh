EXT=${1:?ext}
export WSK_CONFIG_FILE=~/.wskprops.$EXT
F=${2:-}
WSK="wsk $F"

package() {
 wsk $F package list
 wsk $F package delete hello 
 wsk $F package create hello 
 wsk $F package list
 wsk $F package update hello
 wsk $F action delete hello/hello 
 wsk $F action create hello/hello hello.js
 wsk $F package delete hello
 wsk $F action delete hello/hello 
 wsk $F package delete hello
 wsk $F package list
}

single() {
 wsk $F action list
 wsk $F action delete hello
 wsk $F action create hello hello.js
 wsk $F action list
 wsk $F action delete hello
 wsk $F action list
 wsk $F package create hello
 wsk $F action delete hello/world
 wsk $F action create hello/world hello.js
 wsk $F action list
 wsk $F action delete hello/world
 wsk $F package delete hello
 wsk $F action list
}

single
package
