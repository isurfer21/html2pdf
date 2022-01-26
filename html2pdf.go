package main

import (
    "context"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "time"

    "github.com/chromedp/cdproto/emulation"
    "github.com/chromedp/cdproto/page"
    "github.com/chromedp/chromedp"
)

func main() {
    pageUrl := flag.String("url", "http://www.example.com", "Page URL")
    outPdf := flag.String("out", "sample.pdf", "Output filename")

    flag.Parse()

    taskCtx, cancel := chromedp.NewContext(
        context.Background(),
        chromedp.WithLogf(log.Printf),
    )
    defer cancel()
    var pdfBuffer []byte
    if err := chromedp.Run(taskCtx, genPDF(*pageUrl, "body", &pdfBuffer)); err != nil {
        log.Fatal(err)
    }
    if err := ioutil.WriteFile(*outPdf, pdfBuffer, 0644); err != nil {
        log.Fatal(err)
    }
}

func genPDF(url string, sel string, res *[]byte) chromedp.Tasks {
    start := time.Now()
    return chromedp.Tasks{
        emulation.SetUserAgentOverride("WebScraper 1.0"),
        chromedp.Navigate(url),
        chromedp.WaitVisible(`body`, chromedp.ByQuery),
        chromedp.ActionFunc(func(ctx context.Context) error {
            buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
            if err != nil {
                return err
            }
            *res = buf
            fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())
            return nil
        }),
    }
}
