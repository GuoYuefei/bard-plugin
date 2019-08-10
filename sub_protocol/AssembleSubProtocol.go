package sub_protocol

import "net"


type ReadDoFuncType = func(net.Conn) ([]byte, int)
type WriteDoFunType = func([]byte) ([]byte, int)

// 可以通过组合AssembleTCSP来拓展它
type AssembleTCSP struct {
	id string
	readDo ReadDoFuncType
	writeDo WriteDoFunType
}

func NewAssembleTCSP(id string, readDo ReadDoFuncType, writeDo WriteDoFunType) AssembleTCSP {
	return AssembleTCSP{
		id:      id,
		readDo:  readDo,
		writeDo: writeDo,
	}
}

func (a AssembleTCSP)ID() string {
	return a.id
}

// 将do注册如AssembleTCSP
func (a AssembleTCSP)RegisterDo(do1 ReadDoFuncType, do2 WriteDoFunType) {
	a.readDo = do1
	a.writeDo = do2
}

func (a AssembleTCSP)ReadDo(conn net.Conn) ([]byte, int) {
	return a.readDo(conn)
}

func (a AssembleTCSP)WriteDo(bs []byte) ([]byte, int) {
	return a.writeDo(bs)
}
