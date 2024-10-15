# Persian Number Converter

This Golang package converts numeric values to their corresponding Persian alphabetic representation. It supports
numbers.

## Installation

To use this package, run:

```bash
go get github.com/mohammad-safakhou/numalph-go
```

## Usage
First, import the package in your Go file:

```
import "github.com/mohammad-safakhou/numalph-go"
```

Then, you can use the NumToPersian function to convert integers to Persian words:

```persianNumber := numtopersian.NumToPersian(12345)
fmt.Println(persianNumber)
```
## Examples
Here are some examples of how to use the NumToPersian function:

```
package main

import (
    "fmt"

    numtopersian "github.com/mohammad-safakhou/numalph-go"
)

func main() {
    fmt.Println(numtopersian.NumToPersian(0))           // Output: صفر
    fmt.Println(numtopersian.NumToPersian(7))           // Output: هفت
    fmt.Println(numtopersian.NumToPersian(42))          // Output: چهل و دو
    fmt.Println(numtopersian.NumToPersian(1999))        // Output: یک هزار و نهصد و نود و نه
    fmt.Println(numtopersian.NumToPersian(-256))        // Output: منفی دویست و پنجاه و شش
    fmt.Println(numtopersian.NumToPersian(1000000))     // Output: یک میلیون
    fmt.Println(numtopersian.NumToPersian(123456789))   // Output: یکصد و بیست و سه میلیون و چهارصد و پنجاه و شش هزار و هفتصد و هشتاد و نه
}
```