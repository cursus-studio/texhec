package console

type Service interface {
	PrintPermanent(string)
	Print(string)
	Flush()
}
