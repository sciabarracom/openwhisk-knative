# Setup for development

For develomnet you just need to start julia and type `include("setup.jl")`

This will add dependencies to the global scope and add the `src` folder to the LOAD_PATH so you can use  `using KnativeWhisk` for development.

When you develop, to do interactive development you should `cd("src")` or `cd("test")` and include the various files to into the REPL the current code.

As long as you have in your PATH `kubectl` configured to access Kubernetes you can work in the same way as the operator

# Build for Kubernetes

You should first commit your stuff to your current Git origin before building as Docker uses sources from there.

Then make will build your code


