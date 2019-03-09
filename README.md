#### Build Status
| windows | macos | linux | 
|-----------|---------|-------|
|[![Build Status](https://jerrywiltse.visualstudio.com/envx/_apis/build/status/envx?branchName=master&jobName=windows_x64)](https://jerrywiltse.visualstudio.com/envx/_build/latest?definitionId=3&branchName=master)|[![Build Status](https://jerrywiltse.visualstudio.com/envx/_apis/build/status/envx?branchName=master&jobName=macos_x64)](https://jerrywiltse.visualstudio.com/envx/_build/latest?definitionId=3&branchName=master)|[![Build Status](https://jerrywiltse.visualstudio.com/envx/_apis/build/status/envx?branchName=master&jobName=linux_x64)](https://jerrywiltse.visualstudio.com/envx/_build/latest?definitionId=3&branchName=master)

#### Latest Version
[ ![Download](https://api.bintray.com/packages/solvingj/public-bin/envx/images/download.svg) ](https://bintray.com/solvingj/public-bin/envx/_latestVersion)


# envx

`envx` aims to provide a native, cross-platform, cross-language command-line utility for managing environment variables during development workflows in a robust way.  It was inspired by the "Profiles" feature of the Conan Package Manager for C and C++.  It is also somewhat inspired by a common convention shared by Python, Docker, and others which is supporting `-e env_var` and `-e env_var=value` as command-line arguments.  It also aims to enable users to make use of the existing `.env` file format which is becoming somewhat in some ecosystems. 
    
# Download Instructions

Precompiled binaries for Windows, Linux, and macOS are hosted on Bintray.com (courtesy of JFrog).  Eventually, we may package them and submit to the various package managers. 

#### Windows:    
    POSH: curl -OutFile envx.exe https://dl.bintray.com/solvingj/public-bin/windows_x64/envx.exe 
    CMD: powershell -command "curl -OutFile envx.exe https://dl.bintray.com/solvingj/public-bin/windows_x64/envx.exe"
	
#### macOS:   
    curl -L "https://dl.bintray.com/solvingj/public-bin/macos_x64/envx" -o envx

#### Linux:    
    curl -L "https://dl.bintray.com/solvingj/public-bin/linux_x64/envx" -o envx
    

# Build Instructions

If you want to contribute to the code, all you need is a recent version of Go (1.10.0+).  With that, you can just run these commands in the root of the repository: 

#### Windows: 
    go build -o envx.exe
    
#### Linux/macOS
    go build -o envx

To run unit tests, use the following standard command: 

    go test ./...
    
    
# Domain Background

Environment variables are one of very few "common coins" used by virtually all operating systems of software devlopment ecosystems. They are used for managing both local development environments, and are essential for automation in CI platforms and cloud services.  In many cases, they are commonly set in `sh` and `bat` scripts prior to launching developer tools.  In CI systems, we see environment variables representing the cornerstone of a "delcarative" instruction set. They are also compositional in nature, allowing variables to be overridden or appended to at any point later in the call stack.  In practice, they function as implicit arguments which can flow through to any tool involved in a scripted process or pipeline, without intermediate tools and scripts needing to know about them.   

This implicit behavior is what makes them essential for automation, but it can become a liability if not used in a disciplined or experienced way. On the positive side, it allows the calls to scripts from within CI systems to be very concise, often something like `python build.py`.  It also allows them to be very flexible, because the behaviors and parameters of the scripts used can often be changed without having to change any of the command-line calls in the pipeline. On the negative side, this can make it extremely challenging to understand how variables are flowing through such scripts and tools, such as when debugging a problem or getting involved in a project. 

The importance of an effective strategy for using environment variables is typically a function of the size of the team and complexity of a project or environment. For small teams or small environments, it may be inconsequential, but for bigger teams and bigger environments, such a strategy can have a massive impact on scalability and maintainability of an entire devops pipeline.  In either case, there seems to be much room for improvement regarding the tools we have for interacting with environment variables. 

# Motivations

While `sh` and `bat` scripts is often sufficient for setting environment variables for small teams with simple environments, it has many characteristics which are undesirable to larger teams with more complex environments, such as cross-platform and cross-ecosystem use-cases.  Thus, many developers reach for "cross-platform" scripting languages such as python or groovy instead of native scripts.  While they do begin to address these issues, they come with significant drawbacks as well. 

The choice of scripting language and strategy has a major impact on the developer and automation experience, as these scripts become the entry-point for most of the developer tools in the environment, such as docker, build systems, unit-test suites, package managers, etc.  So, in most cases, the automation and scripting language represents a second layer of language, dependency, and complexity for working with the tools for a given project or team.  Meanwhile, in many cases, the primary function these scripts serve is simply to set environment variables and then call other tools.  The result is a proliferation of "wrapper" scripts which clutter up repositories and workflows, and make automation really unpleasant.

While `envx` does not aim to eliminate the use of wrapper scripts for development and automatin workflows, it does aim to extract the responsibility of setting environment variables for these workflows.  Extracting this responsibility from the scripts creates many immediate advantages, but also opens the door to something much more significant: inverting the composition paradigm for automation.

#  Inverting Composition Flow

Here is a loose description of how many scripts function in many respositories: 

    scripts_dir
        run_tests.xyz -> sets vars -> calls test suite
        build_all_dev.xyz -> sets vars -> calls build system
        build_all_prod.xyz -> sets  vars -> calls build system
        deploy_nightly.xyz -> ...

With this pattern, it's very difficult to find a good balance between making the scripts "turn-key" and making the scripts "flexible".  When trying to make common operations "turn-key", you end up with a bunch of single-purpose scripts which gets messy and still don't address all the common workflows for the team.  Developers will still often have to modify or copy scripts to suit their specific needs.  When trying to have fewer scripts which are more flexible, you end up writing, exposing, and maintaining worse APIs to each of the tools the script will call. 

With `envx`, the workflow of a given script can look like this instead: 

                envx run --with-env somevars.env --with-env somevars2.env "build system" 
        
As mentioned earlier, we get several immediate benefits with this approach.  We get composition of environment variables which are stored in declarative, non-proprietary way.  We get behavior that is consistent across any platform and ecosystem.  The variables used by our pipelines are now completely decoupled from our current scripting language.  We can now have fewer and slimmer scripts, and more freedom to change our scripting mechanisms.  Furthermore, in many cases, we can call our tools directly without using a script at all.  


Initial Goals include: 
- Provide alternative to `.sh` and `.bat` scripts which  `export` and `set` 
- Enable users to avoid env var setting from script running via a "run --with-env" functionality
- Enable logging of specific environment variable values
- Providing composability of variables from multiple files and CLI args
- Providing flexible storage mechanism including user-profile (default) or an arbitrary path

Future Goals Include
- Enable prepend/append/replace-in for environment variables
- Providing secure storage for credentials as part of environments
- Sharing environments via git repository
- A member of a future collection of cross-platform utilities with similar vision
- Providing first-class mechanisms for variables relating to search paths: LD_LIBRARY_PATH, PATH, etc.
- Providing an extensibility model for various ecosystems 
