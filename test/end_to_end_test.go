package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

// This example shows how to navigate to a http://play.golang.org page, input a
// short program, run it, and inspect its output.
//
// If you want to actually run this example:
//
//  1. Ensure the file paths at the top of the function are correct.
//  2. Remove the word "Example" from the comment at the bottom of the
//     function.
//  3. Run:
//     go test -test.run=Example$ github.com/tebeka/selenium
func SendKeysToId(id, key string, wd selenium.WebDriver) error {
	elem, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		return err
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		return err
	}

	// Enter some new code in text box.
	err = elem.SendKeys(key)
	if err != nil {
		return err
	}

	return nil
}

// download selenium server driver from https://www.selenium.dev/downloads/#:~:text=Latest%20stable-,version into /test/vendor/ directory
// https://www.selenium.dev/documentation/grid/getting_started/ here is the tutorial but all that is needed is just the "java -jar selenium-server-<version>.jar standalone"
// download needed browser driver from https://www.selenium.dev/ecosystem/ into /test/vendor/
// you also need the browser the driver for which you downloaded
// change the browser name in selenium.Capabilities method on whatever you downloaded the driver for
// run "java -jar selenium-server-<version>.jar standalone" in the /test/vendor directory in one terminal or background
func TestEndToEndTest(t *testing.T) {

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://localhost:8080/login"); err != nil {
		panic(err)
	}

	if err := SendKeysToId("email", "a@a.a", wd); err != nil {
		panic(err)
	}

	if err := SendKeysToId("passwd", "a", wd); err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByID, "submit")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	if err := wd.Get("http://localhost:8080"); err != nil {
		panic(err)
	}

	var output string
	for {
		outputDiv, err := wd.FindElement(selenium.ByXPATH, "//h2")
		if err != nil {
			panic(err)
		}
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output == "Welcome, a" {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))
}
