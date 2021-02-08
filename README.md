# UniHTML

[UniDoc](https://unidoc.io) UniHTML is a [UniPDF](https://github.com/unidoc/unipdf) library plugin module with capability 
of extracting and converting HTML files and inject into PDF document.

It requires a UniHTML Server component to work. Visit [https://unidoc.io](https://unidoc.io). 

[![License: UniDoc EULA](https://img.shields.io/badge/license-UniDoc%20EULA-blue)](https://unidoc.io/eula/)
[![GoDoc](https://godoc.org/github.com/unidoc/unipdf?status.svg)](https://godoc.org/github.com/unidoc/unihtml)

## Features

- Create a PDF file from simple HTML file.
- Create PDF file based on the directory that contains HTML along with CSS, JavaScript and Images.
- Create PDF file based on the external web link.
- Define output page dimensions in metric or imperial systems.
- Define output page size based on the [ISO Paper Sizes](https://en.wikipedia.org/wiki/Paper_size#International_paper_sizes) for `A`, `B` series and a `Letter`.
- Define output page orientation.
- Define custom margin sizes for the output document.
- Injection of HTML defined document into existing PDF document.
- CLI that allows to execute HTML -> PDF conversion.


## Preparation 

UniHTML is a plugin for the [UniPDF](https://github.com/unidoc/unipdf) Golang library.

It requires valid UniPDF license with the UniHTML plugin. Visit [https://unidoc.io](https://unidoc.io) to get more information.

This plugin works in a pair with the UniHTML server. It is distributed using Docker images and could be downloaded directly from the `unidoccloud/unihtml` DockerHub repository 


## Installation

1. Get UniHTML-Server Docker image: 
   `docker pull unidoccloud/unihtml:latest`
2. Start UniHTML server with some output port defined: `docker run -p 8080:8080 -e UNIHTML_LICENSE=path/to/license unidoccloud/unihtml`
3. Define environment variable: `UNIPDF_LICENSE_PATH` with the path to the UniDoc license.
4. Define environment variable: `UNIPDF_CUSTOMER_NAME` with your customer name matching your license.
5. Get latest version of the `github.com/unidoc/unipdf/v3` module: `go get github.com/unidoc/unipdf/v3`

## Usage

Following example connects to the UniHTML server, reads the content of the input file and converts it using `github.com/unidoc/unipdf/v3/creator` package.

```go
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unihtml"
	"github.com/unidoc/unipdf/v3/creator"
)

func main() {
	// UniHTML requires two arguments:
	// - Connection path to the unihtml-server i.e.: https://localhost:8080
	// - Input path for the file / directory to convert
	if len(os.Args) != 3 {
		fmt.Println("Err: provided invalid arguments. No UniHTML server path provided")
		os.Exit(1)
	}

	// Establish connection with the UniHTML Server.
	if err := unihtml.Connect(os.Args[1]); err != nil {
		fmt.Printf("Err:  Connect failed: %v\n", err)
		os.Exit(1)
	}

	// Get new PDF Creator.
	c := creator.New()

	// Read the content of the simple.html file and load it to the conversion.
	doc, err := unihtml.NewDocument(os.Args[2])
	if err != nil {
		fmt.Printf("Err: NewDocument failed: %v\n", err)
		os.Exit(1)
	}

	// Draw the html document file in the context of the creator.
	if err = c.Draw(doc); err != nil {
		fmt.Printf("Err: Draw failed: %v\n", err)
		os.Exit(1)
	}

	// Write the result file to PDF.
	if err = c.WriteToFile("document.pdf"); err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}
}
```


## Contributing

[![CLA assistant](https://cla-assistant.io/readme/badge/unidoc/unipdf)](https://cla-assistant.io/unidoc/unipdf)

All contributors must sign a contributor license agreement before their code will be reviewed and merged.

## Support and consulting

Please email us at support@unidoc.io for any queries.

If you have any specific tasks that need to be done, we offer consulting in certain cases.
Please contact us with a brief summary of what you need and we will get back to you with a quote, if appropriate.

## Licensing Information

This software package (unihtml) is a commercial product and requires a license
code to operate.

The use of this software package is governed by the end-user license agreement
(EULA) available at: https://unidoc.io/eula/

To obtain a Trial license code to evaluate the software, please visit
https://unidoc.io/

