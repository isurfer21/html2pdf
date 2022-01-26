# html2pdf
A CLI tool to generate PDF from HTML webpage

## Prepare

To clone the repository via `go get` command, run

	$ go get github.com/isurfer21/html2pdf

or, using `git clone` command, run

	$ git clone https://github.com/isurfer21/html2pdf.git

## Setup

To resolve dependencies, run 

	$ go get .

To generate build, run

	$ go build 

## Usage

To see help menu, run

	$ ./html2pdf -h
	Usage of ./html2pdf:
	  -out string
	    	Output filename (default "sample.pdf")
	  -url string
	    	Page URL (default "http://www.example.com")

To convert webpage to pdf, run

	$ ./html2pdf -url="https://www.google.co.in/" -out=foo.pdf

	Took: 4.930184 secs
