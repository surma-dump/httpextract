HTTPExtract Proof-of-Concept
==========================

About
-----
HTTPExtract is a simple HTTP proxy program, which analyzes all requests and saves the response to disk if the filename matches a certain pattern.  Right now, it's hardcoded for the filename `stream.php` (*cough* [Grooveshark](http://www.grooveshark.com) *cough*)

Usage
-----
Compile HTTPExtract using `gb` or `make` and make your browser use `localhost:8080` as HTTP proxy.

Credits
-------
(c) 2011 Alexander “Surma” Surma
