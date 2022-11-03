## What?

This is a simple program I wrote for a [light I bought off of adafruit.](https://www.adafruit.com/product/5125) They had
a sample in python but I wanted to do some stuff with the gokrazy library.

This uses the [serial](https://github.com/bugstr/go-serial) library from bugstr to
signal state changes to the light. Its a little self container and may not match
your use case, but works for what I need it for (basically a commandline on/off 
tool). I won't be accepting enhancements to this, but wanted to publish in case anyone
else finds it useful.
