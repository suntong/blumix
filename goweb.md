# Goweb Application

Simple application demonstrating how CloudFoundry-based clouds can host web applications written in [Go language](http://golang.org/).  
It does simple *maths* like calculating prime factors of a number and listing the first N prime numbers.  

You can use the following clouds for testing:  

- [IBM Codename BlueMix](https://ng.bluemix.net)  
- [Pivotal](https://run.pivotal.io).  

You can play with an instance of the app running in BlueMix at [this link](http://goweb.ng.bluemix.net/),
as well as getting the code from JazzHub at [this link](https://hub.jazz.net/project/mcrudele/Goweb).  

Once you get the code you can run *Goweb* directly on your desktop: the only requirement is the `Go Programming Language`
binary distribution that you can download from the official site.  

Thereafter the instructions on how to deploy a Go App on CloudFoundry-based clouds. Feel free to take them as example
for deploying your own app.  


## Creating the deployment artifacts

Once you have done with the code, you need to create a couple of files on the root folder of your app before deploying it in the cloud:  

- `.godir`  
  The file should contain the desired name of the final go libary. E.g.:  
  `example.com/goweb`  
  
  NOTE - CloudFoundry may ignore files beginning with `.`. In that case rename
  `.godir` to `_godir`  

- `Procfile`  
  The file contains the command to start the web app. E.g.:  
  `web: goweb`  


## Deploying the app

Unless your cloud provide support for Go language, you need to use a custom
*buildpack* to deploy a Go web application.  

The one used in the example below has been modified to work with CloudFoundry-based clouds by [Michal Jemala](https://github.com/michaljemala).
Though this is an experimental Cloudfoundry buildpack for Go, it does its job well (at least on BlueMix and Pivotal clouds).  

So, login to your cloud and run the `gcf` command below from the root folder of your application:  

```
gcf push goweb -b https://github.com/michaljemala/cloudfoundry-buildpack-go
```

To deploy this *Goweb* application run the script `.push-this.sh` that you should find on the root folder: it performs some preparation steps before
calling the `gcf push`.  

```
./.push-this.sh
```

This is what you should get to your console:  

```
...
0 of 1 instances running, 1 down
0 of 1 instances running, 1 starting
1 of 1 instances running

App started

Showing health and status for app goweb in org your.name-org / space dev as your.name@your.mail...
OK

requested state: started
instances: 1/1
usage: 256M x 1 instances
urls: goweb.ng.bluemix.net

     state     since                    cpu    memory         disk
#0   running   2014-02-04 06:18:38 PM   0.0%   1.5M of 256M   8.7M of 1G
```


## Licensed under the MIT

```
The MIT License (MIT)

Copyright (c) 2014 Michele Crudele

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
