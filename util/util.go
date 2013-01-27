package util

import (
    "fmt"
    "os"
    "time"
    "reflect"
    "strings"
    "strconv"
    "unicode/utf8"
    "path/filepath"
    "runtime"
)

// Prompt user to input data
func Prompt(msg string) string {
    fmt.Printf(msg)
    var str string
    fmt.Scanln(&str)
    return str
}

// Returns true if file/directory exists
func FileExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    return false
}

func Mkdir(path string) error {
    dir := filepath.Dir(path)
    if FileExists(dir) {
        return nil
    }
    return os.Mkdir(dir, 0700)
}

// Returns the users home dir
func Homedir() string {
    if runtime.GOOS == "windows" {
        return os.Getenv("APPDATA")
    }
    return os.Getenv("HOME")
}

func FormatBool(b bool) string {
    return strings.Title(strconv.FormatBool(b))
}

func JoinInt(numbers []int, sep string) string {
    var res string
    for i, n := range numbers {
        if i > 0 {
            res += sep
        }
        res += strconv.Itoa(n) //strconv.FormatInt(n, 10)
    }
    return res
}

func FileSizeFormat(bytes int64) string {
    units := []string{"B", "KB", "MB", "GB", "TB", "PB"}

    var i int
    value := bytes

    for value > 1000 {
        value /= 1000
        i++
    }
    return fmt.Sprintf("%d %s", value, units[i])
}

// Truncates string to given max length, and inserts ellipsis into
// the middle of the string to signify that the string has been truncated
func TruncateString(str string, maxRunes int) string {
    indicator := "..."

    // Number of runes in string
    runeCount := utf8.RuneCountInString(str)

    // Return input string if length of input string is less than max length
    // Input string is also returned if max length is less than 9 which is the minmal supported length
    if runeCount <= maxRunes || maxRunes < 9 {
        return str
    }
    
    // Number of remaining runes to be removed
    remaining := (runeCount - maxRunes) + utf8.RuneCountInString(indicator)
    
    var truncated string
    var skip bool
    
    for leftOffset, char := range str {
        rightOffset := runeCount - (leftOffset + remaining)

        // Start skipping chars when the left and right offsets are equal
        // Or in the case where we wont be able to do an even split: when the left offset is larger than the right offset
        if leftOffset == rightOffset || (leftOffset > rightOffset && !skip) {
            skip = true
            truncated += indicator
        }

        if skip && remaining > 0 {
            // Skip char and decrement the remaining skip counter
            remaining--
            continue
        }
        
        // Add char to result string
        truncated += string(char)
    }
    
    // Return truncated string
    return truncated
}

func TimeToLocal(t time.Time) string {
    local := t.Local()
    year, month, day := local.Date()
    hour, min, sec := local.Clock()
    return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec)

}

func ISODateToLocal(iso string) string {
    t, err := time.Parse(time.RFC3339, iso)
    if err != nil {
        return iso
    }
    return TimeToLocal(t)
}

func MeasureTransferRate() func(int64)string {
    start := time.Now()

    return func(bytes int64) string {
        seconds := int64(time.Now().Sub(start).Seconds())
        if seconds < 1 {
            return fmt.Sprintf("%s/s", FileSizeFormat(bytes))
        }
        bps := bytes / seconds
        return fmt.Sprintf("%s/s", FileSizeFormat(bps))
    }
}

// Prints a map in the provided order with one key-value-pair per line
func Print(m map[string]interface{}, keyOrder []string) {
    for _, key := range keyOrder {
        value, ok := m[key]
        if ok && value != "" {
            fmt.Printf("%s: %v\n", key, value)
        }
    }
}

// Prints a struct with one name-value-pair per line
func PrintStruct(s interface{}) {
    printStruct(s, "")
}

func printStruct(s interface{}, indent string) {
    structValue := reflect.Indirect(reflect.ValueOf(s))
    if !structValue.IsValid() {
        return
    }

    structType := structValue.Type()

    for i := 0; i < structType.NumField(); i++ {
        structField := structType.Field(i)
        name := structField.Name
        fieldValue := structValue.Field(i)
        value := fieldValue.Interface()
        fieldKind := fieldValue.Kind().String()
        if fieldKind == "ptr" {
            if !reflect.Indirect(reflect.ValueOf(value)).IsValid() {
                continue
            }

            // Dont add indent or print field name if it is anonymous (embeded)
            if !structField.Anonymous {
                fmt.Printf("%s%s:\n", indent, name)
                newIdent := indent + "    "
                printStruct(value, newIdent)
            } else {
                printStruct(value, indent)
            }
        } else if fieldKind == "slice" {
            switch t := value.(type) {
                case []string:
                    fmt.Printf("%s%s: %v\n", indent, name, strings.Join(value.([]string), ", "))
                case []int:
                    fmt.Printf("%s%s: %v\n", indent, name, JoinInt(value.([]int), ", "))
                default:
                    fmt.Printf("%s%s: <Unhandled type: %v>\n", indent, name, t)
            }
        } else if fieldKind == "map" {
            fmt.Printf("%s%s:\n", indent, name)
            newIndent := indent + "    "
            printMap(value.(map[string]interface{}), newIndent)
        } else if value != nil && value != "" {
            fmt.Printf("%s%s: %v\n", indent, name, value)
        }
    }
}

func printMap(m map[string]interface{}, indent string) {
    for k, v := range m {
        fmt.Printf("%s%s: %v\n", indent, k, v)
    }
}

// Prints items in columns with header and correct padding
func PrintColumns(items []map[string]string, keyOrder []string, columnSpacing int) {
    // Create header
    header := make(map[string]string)
    for _, key := range keyOrder {
        header[key] = key
    }

    // Add header as the first element of items
    items = append([]map[string]string{header}, items...)

    // Get a padding function for each column
    padFns := make(map[string]func(string)string)
    for _, key := range keyOrder {
        padFns[key] = columnPadder(items, key, columnSpacing)
    }
    
    // Loop, pad and print items
    for _, item := range items {
        var line string

        // Add each column to line with correct padding
        for _, key := range keyOrder {
            value, _ := item[key]
            line += padFns[key](value)
        }

        // Print line
        fmt.Println(line)
    }
}

// Returns a padding function, that pads input to the longest string in items
func columnPadder(items []map[string]string, key string, spacing int) func(string)string {
    // Holds length of longest string
    var max int
    
    // Find the longest string of type key in the array
    for _, item := range items {
        str := item[key]
        length := utf8.RuneCountInString(str)
        if length > max {
            max = length
        }
    }

    // Return padding function
    return func(str string) string {
        column := str
        for utf8.RuneCountInString(column) < max + spacing {
            column += " "
        }
        return column
    }
}
