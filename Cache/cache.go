package Cache


type Cache interface {
	Set (key string, value interface{})
	Get (key string) interface{}
	Del (key string)
}
