package seeweb

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"
)

// Compares two structs in similiar way of reflect.DeepEqual, but with Date
// (time.Time) is limited to the fields at the root of the structs. This
// function was introduced because dates aren't being evaluated in same way
// between Mac and Linux.
func equalStructWithDatesFn(a interface{}, b interface{}) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.NumField() != vb.NumField() {
		return false
	}

	for i := 1; i < va.NumField(); i++ {
		if va.Field(i).Type() != vb.Field(i).Type() {
			return false
		}
		valueOfA := va.Field(i).Interface()
		valueOfB := vb.Field(i).Interface()
		if fmt.Sprint(va.Field(i).Type()) == "time.Time" {
			if !equalForDatesFn(valueOfA, valueOfB) {
				return false
			}
		} else {
			if !reflect.DeepEqual(valueOfA, valueOfB) {
				return false
			}
		}
	}
	return true
}

func equalForDatesFn(a interface{}, b interface{}) bool {
	da := a.(time.Time)
	db := b.(time.Time)

	return da.Equal(db)
}

var logMutex sync.Mutex

func WriteLog(filename, content string) {
	logMutex.Lock()
	defer logMutex.Unlock()

	dirName := "LOG_CLIENT"
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return
	}

	filePath := filepath.Join(dirName, filename)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(content + "\n--------------------------------------------------\n")
}

func PrintStruct(s interface{}) string {
	var sb strings.Builder
	printStructRecursive(reflect.ValueOf(s), "", &sb)
	return sb.String()
}

func printStructRecursive(v reflect.Value, indent string, sb *strings.Builder) {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			sb.WriteString("<nil>\n")
			return
		}
		v = v.Elem()
	}

	// Gestione di Slice e Array (es. []*seeweb.Server)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if v.Len() == 0 {
			sb.WriteString("[] (vuoto)\n")
			return
		}
		sb.WriteString(fmt.Sprintf("(Lista di %d elementi):\n", v.Len()))
		for i := 0; i < v.Len(); i++ {
			sb.WriteString(fmt.Sprintf("%s[%d]:\n", indent+"    ", i))
			printStructRecursive(v.Index(i), indent+"        ", sb)
		}
		return
	}

	if v.Kind() != reflect.Struct {
		sb.WriteString(fmt.Sprintf("%v\n", v.Interface()))
		return
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		// Salta i campi non esportati (minuscoli) per evitare il panic di reflection
		if field.PkgPath != "" {
			continue
		}
		if value.Kind() == reflect.Ptr && value.IsNil() {
			sb.WriteString(fmt.Sprintf("%sCampo: %-15s | Tipo: %-15s | Valore: <nil>\n", indent, field.Name, field.Type))
			continue
		}

		actualVal := value
		if actualVal.Kind() == reflect.Ptr {
			actualVal = actualVal.Elem()
		}

		if actualVal.Kind() == reflect.Struct || actualVal.Kind() == reflect.Slice || actualVal.Kind() == reflect.Array {
			sb.WriteString(fmt.Sprintf("%sCampo: %-15s | Tipo: %-15s |\n", indent, field.Name, field.Type))
			printStructRecursive(value, indent+"    ", sb)
		} else {
			valInterface := value.Interface()
			sb.WriteString(fmt.Sprintf("%sCampo: %-15s | Tipo: %-15s | Valore: %v\n", indent, field.Name, field.Type, valInterface))
		}
	}
}
