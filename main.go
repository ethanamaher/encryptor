package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

const conversionString = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="


func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter string to encrypt: ")
    text, _ := reader.ReadString('\n')
    trimmedString := strings.Trim(text, "\n")
    binary := toBinary(trimmedString)
    sextetSegments := getSextets(binary)

    // '=' padding
    for i := 0; i < len(sextetSegments)%4; i++ {
        sextetSegments = append(sextetSegments, 64)
    }

    var builder strings.Builder
    for _, s := range sextetSegments {
        builder.WriteByte(conversionString[s])
    }

    encryptedString := builder.String()
    fmt.Printf("Encrypted string: %s\n", encryptedString)
}

func toBinary(s string) (b string) {
    for _, c := range s {
        b = fmt.Sprintf("%s%.8b", b, c)
    }
    return
}

func getSextets(s string) []int {
    var segments[] int

    for i := 0; i < len(s); i += 6 {
        endOfSegment := i+6
        if endOfSegment > len(s) {
            endOfSegment = len(s)
        }

        bitSegment := s[i:endOfSegment]
        // if bit sextet needs padding at end
        if (endOfSegment-i)%6 != 0 {
            lengthDifference := 6 - ((endOfSegment-i)%6)

            var builder strings.Builder
            builder.WriteString(bitSegment)
            for j := 0; j < lengthDifference; j++ {
                builder.WriteString("0")
            }
            bitSegment = builder.String()
        }
        sextetSegment, _ := strconv.ParseInt(bitSegment, 2, 64)
        segments = append(segments, int(sextetSegment))

    }
    return segments
}

