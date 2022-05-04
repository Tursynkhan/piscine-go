package main
import(
	"io"
	"os"
	"github.com/01-edu/z01"
)

func main(){
	args:=os.Args[1:]
	if len(args)==0{
		s:=os.Stdin
		io.Copy(os.Stdout,s)
	}else{
		for i:=range args{
			content, err:=os.ReadFile(args[i])
			PrintStr(string(content))
			if err!=nil{
				str:="ERROR:"+err.Error()+"\n"
				PrintStr(str)
				os.Exit(1)
			}
		}
	}
}
func PrintStr(str string){
	for_, w:= range str{
		z01.PrintRune(w)
	}
}