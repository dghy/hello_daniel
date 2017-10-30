package main

import (
	"fmt"

	"github.com/bunsenapp/go-selenium"

	"time"
)

func maintain(message string, err error,
	          webDriver goselenium.WebDriver) (string) {
	if err != nil {
		webDriver.DeleteSession()
		return "WebDriver deleted.. " + err.Error()
	} else {
		return message
	}
}


func main() {
    realName := "Toby"
    displayName := "Tob_Random_Bot"
    passWord := "tobby_the_best_bot"
	// Create a capabilities object.
	capabilities := goselenium.Capabilities{}

	// Populate it with the browser you wish to use.
	capabilities.SetBrowser(goselenium.ChromeBrowser())

	// Initialise a new web drivers: driver
	driver, err := goselenium.NewSeleniumWebDriver("http://localhost:4444/wd/hub",
		                                           capabilities)
	fmt.Println(maintain("Driver Created", err, driver))
	// Create a session for driver.
	_, err = driver.CreateSession()
	fmt.Println(maintain("Session Created", err, driver))


	// Initialise a new web drivers:driver2
	driver2, err := goselenium.NewSeleniumWebDriver("http://localhost:4444/wd/hub",
		capabilities)
	fmt.Println(maintain("Driver2 Created", err, driver2))

	// Create a session.
	_, err = driver2.CreateSession()
	fmt.Println(maintain("Session2 Created", err, driver2))


	newLink := "https://slack.com/create"
	_, err = driver.Go(newLink)
	fmt.Println(maintain("Navigated to " + newLink, err, driver))


	el, err := driver.FindElement(goselenium.ByCSSSelector("#signup_email"))
	fmt.Println(maintain("Element found!", err, driver))


	elTagName, err := el.TagName()
	time.Sleep(1 * time.Second)

	fmt.Println(maintain("Successful Tag Name", err, driver))
	fmt.Println(elTagName)

	response, err := el.SendKeys("auto_tester_12345@wp.pl")
	// wp
	//login: auto_tester_12345@wp.pl
	//pass: keuthm_856tygh


	fmt.Println(maintain("Response" + response.State, err, driver))

	el, err = driver.FindElement(goselenium.ByCSSSelector("#email_misc"))
	fmt.Println(maintain("Element found!", err, driver))

	// Click the checkbox - don't want new mail!
	_, err = el.Click()
	fmt.Println("Checkbox unchecked", err, driver)

	// select submit button
	el, err = driver.FindElement(goselenium.ByCSSSelector("#submit_btn"))
	fmt.Println(maintain("Element found!", err, driver))
	// Click the submit button
	_, err = el.Click()
	fmt.Println("Clicked! Now log in into email", err, driver)

	// wait for the e-mail to come
    time.Sleep(60 * time.Second)

    ////////////////////////////////////////////////////////////////
    ////////////////////////////////////////////////////////////////
    // 2nd session for finding verification code
	newLink = "https://profil.wp.pl/login.html?url=http%3A%2F%2Fpoczta.wp.pl" +
		      "%2Findexgwt.html%3Fflg%3D1&serwis=nowa_poczta_wp"
	_, err = driver2.Go(newLink)
	fmt.Println(maintain("Navigated to " + newLink, err, driver2))
	time.Sleep(1 * time.Second)

    // login into wp email - get code from slack for new workspace creation
	// select login input
	el, err = driver2.FindElement(goselenium.ByCSSSelector("#login"))
	fmt.Println(maintain("Element found!", err, driver2))
	// write login
	response, err = el.SendKeys("auto_tester_12345@wp.pl")
	fmt.Println(maintain("Response" + response.State, err, driver2))

	// select password input
	el, err = driver2.FindElement(goselenium.ByCSSSelector("#password"))
	fmt.Println(maintain("Element found!", err, driver2))
	// write password
	response, err = el.SendKeys("keuthm_856tygh")
	fmt.Println(maintain("Response" + response.State, err, driver2))

	// select submit button
	el, err = driver2.FindElement(goselenium.ByCSSSelector("#btnSubmit"))
	fmt.Println(maintain("Element found!", err, driver2))
	// Click the submit button
	_, err = el.Click()
	fmt.Println(maintain("Logged In!", err, driver2))
	time.Sleep(1 * time.Second)

	// select recived messages
	el, err = driver2.FindElement(goselenium.ByCSSSelector(".GNNS3A-MK"))
	fmt.Println(maintain("Element found!", err, driver2))
	// Click the submit button
	_, err = el.Click()
	fmt.Println(maintain("Show messages", err, driver2))
	time.Sleep(1 * time.Second)

	// select message
	el, err = driver2.FindElement(goselenium.ByCSSSelector(".normalMsg"))
	fmt.Println(maintain("Element found!", err, driver2))
	// Click the message to open it
	_, err = el.Click()
	fmt.Println(maintain("Open the message!", err, driver2))
    time.Sleep(2 * time.Second)

	// find code in the slack message
	// show message in text mode
	el, err = driver2.FindElement(goselenium.ByCSSSelector("#trescRamka"))
	fmt.Println(maintain("Looking for the code..", err, driver2))

	newLink2, err := el.Attribute("src")
	fmt.Println(maintain("Got the link:", err, driver2))
    fmt.Println(newLink2.Value)
	_, err = driver2.Go(newLink2.Value)
	fmt.Println(maintain("Navigated to " + newLink, err, driver2))
	time.Sleep(1 * time.Second)

	el, err = driver2.FindElement(goselenium.ByCSSSelector(
		"html > body > table > tbody > tr:nth-child(2) > td > table > " +
			"tbody > tr > td > div > div > table > tbody > tr > td"))
	fmt.Println(maintain("Looking for the code..", err, driver2))
    code, err := el.Text()
	fmt.Println(maintain("..", err, driver2))
	fmt.Println(code.Text)

    // At this point got the code required for first session - kill 2nd session
	driver2.DeleteSession()

    //////////////////////////////////////////////////////////////////
    //////////////////////////////////////////////////////////////////
	//".confirmation_code_group:nth-of-type(1) > div:nth-of-type(1) > input"

	el, err = driver.FindElement(goselenium.ByCSSSelector(
		".confirmation_code_group:nth-of-type(1) > " +
				 "div:nth-of-type(1) > input"))
	fmt.Println(maintain("Element found!", err, driver))

	_, _ = el.SendKeys(string(code.Text[0]))

	el, err = driver.FindElement(goselenium.ByCSSSelector(
		".confirmation_code_group:nth-of-type(1) > " +
			"div:nth-of-type(2) > input"))
	fmt.Println(maintain("Element found!", err, driver))

	_, _ = el.SendKeys(string(code.Text[1]))


	el, err = driver.FindElement(goselenium.ByCSSSelector(
		".confirmation_code_group:nth-of-type(1) > " +
			"div:nth-of-type(3) > input"))
	fmt.Println(maintain("Element found!", err, driver))

	_, _ = el.SendKeys(string(code.Text[2]))


	el, err = driver.FindElement(goselenium.ByCSSSelector(
		".confirmation_code_group:nth-of-type(3) > " +
			"div:nth-of-type(1) > input"))
	fmt.Println(maintain("Element found!", err, driver))

	_, _ = el.SendKeys(string(code.Text[4]))


	el, err = driver.FindElement(goselenium.ByCSSSelector(
		".confirmation_code_group:nth-of-type(3) > " +
			"div:nth-of-type(2) > input"))
	fmt.Println(maintain("Element found!", err, driver))

	_, _ = el.SendKeys(string(code.Text[5]))


	el, err = driver.FindElement(goselenium.ByCSSSelector(
		".confirmation_code_group:nth-of-type(3) > " +
			"div:nth-of-type(3) > input"))
	fmt.Println(maintain("Element found!", err, driver))
	_, _ = el.SendKeys(string(code.Text[6]))

	time.Sleep(1 * time.Second)

	el, err = driver.FindElement(goselenium.ByCSSSelector("#signup_real_name"))
	fmt.Println(maintain("Element found", err, driver))
	_, _ = el.SendKeys(realName)

	el, err = driver.FindElement(goselenium.ByCSSSelector("#signup_display_name"))
	fmt.Println(maintain("Element found", err, driver))
	_, _ = el.SendKeys(displayName)

	el, err = driver.FindElement(goselenium.ByCSSSelector("#submit_btn"))
	_, err = el.Click()
	time.Sleep(2 * time.Second)
	fmt.Println(maintain("El Klikto", err, driver))


	el, err = driver.FindElement(goselenium.ByCSSSelector("#signup_password"))
	fmt.Println(maintain("Element found", err, driver))
	_, _ = el.SendKeys(passWord)

	el, err = driver.FindElement(goselenium.ByCSSSelector("#submit_btn"))
	_, err = el.Click()
	time.Sleep(2 * time.Second)
	fmt.Println(maintain("El Klikto", err, driver))


	// delete session
	driver.DeleteSession()

	fmt.Println("Sessions deleted!")


}



////login: auto_tester_12345@wp.pl
////pass: keuthm_856tygh
