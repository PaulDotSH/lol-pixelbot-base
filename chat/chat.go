package chat

func SendToChat(word string)  {
	robotgo.KeyTap("enter")
	time.Sleep(20*time.Millisecond)
	for _, v := range strings.Split(word,"\n") {
	robotgo.TypeStr(v)
	time.Sleep(100*time.Millisecond)
	robotgo.KeyTap("enter")
	time.Sleep(50*time.Millisecond)
	robotgo.KeyTap("enter")
	}
}