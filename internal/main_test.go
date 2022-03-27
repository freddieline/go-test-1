package internal

import(
	"os"
	"testing"

)


func TestMain(m *testing.M){
	main = NewMain()
	os.Exit(m.Run)
}