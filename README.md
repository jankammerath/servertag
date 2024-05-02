# Servertag

This is a proof of concept I always wanted to do in Go. The question is whether it is possible to build a (LiveWire)[https://docs.oracle.com/cd/E19957-01/816-6411-10/contents.htm] and Netscape Enterprise Server compatible server-side JavaScript server with Go and the V8 JavaScript engine.

## Server-side JavaScript

The contents inside the `<server></server>` tag is parsed and sent through V8. After the execution, the output result of the executed JavaScript replaces the `<server>` tag code. Pretty much like early server-side JavaScript. This code thus proves how simple server-side JavaScript or JavaScript in general can be added to extend applications during runtime. 

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello, World!</title>
</head>
<body>
    <h1>Hello, World!</h1>
    5 * 8 = <server>5 * 8</server>

    <div>
        <h2>Server-side rendering</h2>
        <p>3 + 6 = <server>3 + 6</server></p>
    </div>

    <server>"ok, this is the server at " + new Date().toString()</server>
</body>
</html>
```

## Challenges

The main challenge is implementing the interfaces for the JavaScript code to do meaningful things on the server. Since this is plain V8, there are no libraries that you would know from things like Node or Deno. The (v8go)[https://github.com/rogchap/v8go] library however allows for exactly that as it is a binding for the V8 C++ API.