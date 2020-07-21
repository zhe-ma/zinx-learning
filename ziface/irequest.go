package ziface

type IRequtest interface {
	GetConnection() IConnection
	GetData() []byte
}
