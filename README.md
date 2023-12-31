# PDF Generation PoC With Puppeteer and Go

<h3>Puppeteer CLI</h3>
<p>puppeteer cli is a cli wrapper for generating pdfs and take a screenshot with <a href="https://developer.chrome.com/docs/puppeteer/">Puppeteer</a>. <br>
To install puppeteer cli globally,
</p>

``npm install -g puppeteer-cli`` <br>
<p>
For more information <a href="https://github.com/JarvusInnovations/puppeteer-cli">Puppeteer-CLI</a>
</p>


<h3>Mustache</h3>
<p>Mustache is a simple, logic-less template engine. due to it's simplicity, I chose it to use in this PoC <br>
For more information <a href="https://mustache.github.io/">Mustache</a> and <a href="https://mustache.github.io/mustache.5.html">Mustache Specs</a> <br>
<a href="https://pkg.go.dev/github.com/cbroglie/mustache">Mustache</a> package used in this PoC
</p>


<h3>Setup</h3>
Step 01: <br>
Install puppeteer globally

Step 02: <br>
Clone the repo

Step 03:
install the relevant packages <br>
``go mod tidy`` <br>
`` go mod vendor``

Step 04: <br>
Run the application 
````
go run main.go