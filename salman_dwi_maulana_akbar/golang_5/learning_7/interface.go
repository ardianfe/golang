package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// CustomWriter implements io.Writer
type CustomWriter struct{}

func (cw CustomWriter) Write(p []byte) (n int, err error) {
    n = len(p)
    fmt.Printf("Custom Writer: %s\n", string(p))
    return n, nil
}

// readAllData reads from an io.Reader until EOF
func readAllData(r io.Reader) {
    buf := make([]byte, 1024)
    for {
        n, err := r.Read(buf)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Error:", err)
            break
        }
        fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
    }
}

func main() {
    // Reader example
    reader := strings.NewReader("Hello, Go!")
    readAllData(reader)

    // Writer example
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    data := []byte("Hello, Go Writer!")
    n, err := file.Write(data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Wrote %d bytes\n", n)

    // io.Copy example
    srcFile, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer srcFile.Close()

    dstFile, err := os.Create("copy_example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer dstFile.Close()

    bytesCopied, err := io.Copy(dstFile, srcFile)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Copied %d bytes\n", bytesCopied)

    // Custom Writer example
    customWriter := CustomWriter{}
    io.WriteString(customWriter, "Hello, Custom Writer!")
}
