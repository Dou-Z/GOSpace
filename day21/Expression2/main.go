package main

type Data struct{}

func (Data) TestValue() {}

func (*Data) TestPointer() {}

func main() {
	var p *Data = nil
	p.TestPointer()

	(*Data)(nil).TestPointer()
	(*Data).TestPointer(nil)

}
