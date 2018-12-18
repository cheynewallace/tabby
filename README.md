# Tabby
A tiny library for super simple GoLang tables

```
go get github.com/cheynewallace/tabby
```

Tabby is a tiny (around 80 lines of code) libary for writing extremely simple table based terminal output in GoLang. 

Many table libraries out there are overly complicated and packed with features you don't need. If you simply want to write clean output to your terminal in table format with minimal effort, Tabby is for you.

Typical examples are writing tables with heading and tab spaced columns, or writing log lines to the terminal with evenly spaced columns

**Example With Heading**
```	
tab := tabby.New()
tab.AddHeader("NAME", "TITLE", "DEPARTMENT")
tab.AddLine("John Smith", "Developer", "Engineering")
tab.Print()
```

**Output**
```
NAME       TITLE     DEPARTMENT
----       -----     ----------
John Smith Developer Engineering
```

**Example Without Heading**
```	
tab := tabby.New()
tab.AddLine("Info:", "WEB", "Success 200")
tab.AddLine("Info:", "API", "Success 201")
tab.AddLine("Error:", "DATABASE", "Connection Established")
tab.Print()
```

**Output**
```
Info:  WEB      Success 200
Info:  API      Success 201
Error: DATABASE Connection Established
```