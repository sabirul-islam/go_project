package main
import "fmt"

type Email interface{
	Write(string, string, string)
	Send()
	Read()
}

type Test struct{
	To string
	From string
	Subject string
	Message string
}

func (t Test) Write(to string, from string, subject string)  {
	fmt.Println(to, from, subject)
}

func (h Test) Send()  {
	fmt.Println(h.Message, "email sent")
}

func (y Test) Read()  {
	fmt.Println(y.From, "recieved from")
}

func main()  {
	// var e Email
	// fmt.Println(e)

	var tst Test
	tst.To = "sabirul.shimul@gmail.com"
	tst.From ="shimul"
	tst.Subject = "test email"
	tst.Message = "Hello"
	tst.Write(tst.To,tst.From,tst.Subject)
	tst.Send()
	tst.Read()
}