package EstructurasDeDatos

type IList interface {
	Add(value interface{}) string
	Delete(value interface{}) string
	Insert(value interface{}) string
	Get(value interface{}) string
	Lenght() int
}

type Lista struct{
	len  int
}

func (L *Lista) Add(value interface{}) string {
	return "add"
}
func (L *Lista) Delete(value interface{}) string {
	return "delete"
}
func (L *Lista) Insert(value interface{}) string {
	return "insert"
}
func (L *Lista) Get(value interface{}) string {
	return "get"
}
func (L *Lista) Lenght() int {
	return L.len
}

func CreateLista(L IList) IList {
	return L
}
