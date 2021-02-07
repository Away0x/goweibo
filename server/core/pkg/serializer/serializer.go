package serializer

import (
  "encoding/json"
  "reflect"
)

type Serializer interface {
  Serialize() interface{}
}

type Data map[string]interface{}

func (s Data) ToJSONString() string {
  b, err := json.Marshal(s)
  if err != nil {
    return ""
  }

  return string(b)
}

func Serialize(data interface{}) interface{} {
  return getSerializeData(data)
}

func getSerializeData(val interface{}) (data interface{}) {
  defer func() {
    if err := recover(); err != nil {
      data = val
    }
  }()

  // val 是 Serializer 类型
  if v, ok := val.(Serializer); ok {
    data = v.Serialize()
    return
  }

  // val 可能是 []Serializer 类型
  value := reflect.ValueOf(val)
  kind := value.Kind()

  switch kind {
  case reflect.Slice, reflect.Array:
    length := value.Len()
    l := make([]interface{}, 0)

    for i := 0; i < length; i++ {
      item := value.Index(i).Interface()
      l = append(l, getItemSerializeData(item))
    }

    data = l
  case reflect.Ptr:
    v := value.Elem().Interface()
    if reflect.TypeOf(v).Kind() != reflect.Ptr {
      data = getSerializeData(value.Elem().Interface())
    } else {
      data = val
    }
  default:
    data = val
  }

  return
}

func getItemSerializeData(val interface{}) (data interface{}) {
  switch typed := val.(type) {
  case Serializer:
    data = typed.Serialize()
  case Data:
    data = typed
  default:
    data = typed
  }

  return data
}
