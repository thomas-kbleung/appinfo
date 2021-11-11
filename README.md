# appinfo
Returns demo application meta information for mobile app.  The api binds /v1/app_info/:system which accepts "iOS" or "Android" as the :system parameter.  Minimum version and latest version of the requested system is returned.

# Usage
API takes API_ENPOINT in host:port format as the binding address from environmental variable.
To test locally, runs scripts/run_api.sh from the project root directory.
```console
scripts/run_api.sh
```

Press Ctrl-C to break the program.

# Examples
Sample request for each system can be found in the [example rest file](examples/get_appinfo.rest).

# Build
Build the project with Makefile. Default build target builds the executable.  
```console
Make
```

Clean up the binary file with Make clean.
```console
Make clean
```

