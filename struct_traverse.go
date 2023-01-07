package gotils

import (
	"log"
	"reflect"
)

/*
	Traverse the structs recursively allowing to modify the instances in place if required.
	For example - hiding or updating values selectively just before serializing to JSON.
**/
type StructTraverseTraverser struct {
	Params map[string]interface{}
}

/*
	Interface to be adopted for instances that want to hide / modify fields by StructTraverseTraverser
**/
type StructTraverseVisitor interface {
	StructTraverseVisitor(traverseContext *StructTraverseTraverser) error
}

func CreateStructTraverseContext(params map[string]interface{}) *StructTraverseTraverser {
	return &StructTraverseTraverser{
		Params: params,
	}
}

func (t *StructTraverseTraverser) Traverse(objectToTraverse interface{}) error {
	var err error

	original := reflect.ValueOf(objectToTraverse)
	err = t.traverseRecursive(original, false)
	CheckNotFatal(err)

	return err
}

func (t *StructTraverseTraverser) traverseRecursive(
	original reflect.Value, isPointer bool) error {
	var err error
	switch original.Kind() {
	case reflect.Ptr:
		originalValue := original.Elem()

		if !originalValue.IsValid() {
			return err
		}

		// check if CanInterface:
		if !originalValue.CanInterface() {
			return err
		}

		if !originalValue.CanAddr() {
			return err
		}

		if visitor, ok := (originalValue.Addr().Interface()).(StructTraverseVisitor); ok {
			visitor.StructTraverseVisitor(t)
		}

		// continue recursively:
		t.traverseRecursive(originalValue, true)

	case reflect.Interface:
		if original.IsZero() {
			return err
		}

		originalValue := original.Elem()

		if !originalValue.IsValid() {
			return err
		}

		// check if CanInterface:
		if !originalValue.CanInterface() {
			return err
		}

		if !originalValue.CanAddr() {
			return err
		}

		reflect.New(originalValue.Type()).Elem()
		t.traverseRecursive(originalValue, false)

	case reflect.Struct:
		if !isPointer && original.CanAddr() && original.CanInterface() {
			if _, ok := (original.Addr().Interface()).(StructTraverseVisitor); ok {
				log.Fatalln("ERROR: all StructTraverseVisitor must be sent as a pointer: ", original.Type().Name())
				CheckNotFatal(err)
			}
		}

		for i := 0; i < original.NumField(); i += 1 {
			field := original.Field(i)
			t.traverseRecursive(field, false)
		}

	// TODO:: should reflect.UnsafePointer be handled?

	case reflect.Slice:
		for i := 0; i < original.Len(); i += 1 {
			t.traverseRecursive(original.Index(i), false)
		}

	case reflect.Map:
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			t.traverseRecursive(originalValue, false)
		}

	default:
	}

	return err
}
