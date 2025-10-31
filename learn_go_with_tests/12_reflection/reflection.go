package main

import (
    "reflect"
)

func walk(x interface{}, fn func(string)) {
    val := getVal(x)

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct, reflect.Pointer:
            walk(field.Interface(), fn)
        }
    }
}

func getVal(x interface{}) reflect.Value {
    val := reflect.ValueOf(x)
    if val.Kind() == reflect.Pointer {
        return val.Elem()
    }
    return val
}
