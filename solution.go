package main

import (
    "flag"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "bufio"
)

const (
    bufSize = 8 * 1024 * 1024
)

func walkDir(name string, output io.Writer) {
    entries, err := os.ReadDir(name)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read directory %v: %v\n", name, err)
        return
    }

    fmt.Fprintf(output, "%q:[", name)
    for _, entry := range entries {
        entryName := filepath.Join(name, entry.Name())
        if entry.IsDir() {
            walkDir(entryName, output)
            fmt.Fprintf(output, ",")
        } else {
            fmt.Fprintf(output, "%q,", entryName)
        }
    }

    fmt.Fprintf(output, "]")
}

func main() {
    directoryFlag := flag.String("directory", ".", "directory to walk")
    outputFlag := flag.String("output", "result.txt", "result file")
    flag.Parse()

    output, err := os.Create(*outputFlag)
    if err != nil {
        panic(fmt.Sprintf("Failed to open %v for writing\n", *outputFlag))
    }

    defer output.Close()

    bufferedOuput := bufio.NewWriterSize(output, bufSize)
    defer bufferedOuput.Flush()

    walkDir(*directoryFlag, bufferedOuput)
}
