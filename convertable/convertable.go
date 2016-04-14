package convertable

type Convertable interface {
	String() string
	Bytes() []byte
	FromFile(filename string) *Convertable
	FromBytes(bytes []byte) *Convertable
}
