# go-workable-api

API client in Go for workable.com, please see relevant documentation on the Workable API
here: https://developers.workable.com/

The pre-requisites for this api are:

 * A workable account and API key;
 * Any recent version of go

 You can view the API reference for this module here: https://pkg.go.dev/github.com/thevfxcoop/go-workable-api
 
## Getting Started

To use this API in your own code, import it and create a client:

```go
package main

import (
	"github.com/thevfxcoop/go-workable-api/pkg/client"
)

func main() {
    var key string

    endpoint,err := url.Parse("https://xxx.workable.com/spi/v3")
    if err != nil {
        // ...
    }

    workable, err := client.NewClient(endpoint,client.OptAPIKey(key),client.OptRateLimit(0.5))
    if err != nil {
        // ...
    }

    accounts, err := workable.GetAccounts()
    if err != nil {
        // ...
    }    
}
```

Replace `xxx` in the URL with your workable account URL, or 'www' otherwise.
You can call the `GetAccounts` method in order to get the accounts your API Key (or OAuth token) works for.

## Schema

Some of the Workable schema is implemented and available
in `pkg/schema` for accounts, candidates and jobs. The schema documentation is:

  * [Job](https://github.com/thevfxcoop/go-workable-api/blob/main/pkg/schema/jobs.go)
  * [Candidate](https://github.com/thevfxcoop/go-workable-api/blob/main/pkg/schema/candidates.go)
  * [Account](https://github.com/thevfxcoop/go-workable-api/blob/main/pkg/schema/accounts.go)

## Under Construction

This API is all that is needed for The VFX Cooperative in order to post roles from their website
into workable, and as such it is the bare minimum to get this to work. If you wish to expand on this API, please do send pull requests in order to extend and make this code more universally useful, thank you. And if you have issues, please do file them!

## Command-Line Tool

There is a command-line tool to demonstrate some of the features of the API. In order to
build this, use the following command:

```bash
[bash] git clone https://github.com/thevfxcoop/go-workable-api
[bash] cd go-workable-api && make
```

The tool is placed in the "build" folder within the repository. To use the tool, you can
set the environment variable WORKABLE_ENDPOINT to the endpoint for your web service or
include the endpoint URL on the command line with the `-endpoint` flag. You will also
need to provide an API Key (or OAuth token) with the WORKABLE_APIKEY variable or `-key`
command-line flag. For example,


```bash
[bash] export WORKABLE_ENDPOINT=https://xxx.workable.com/spi/v3
[bash] export WORKABLE_APIKEY=0d7bc22dfeff845537c21e7ce7b106db35a60a7c
[bash] build/workable 
[
  {
    "id": "abcdef",
    "name": "ABC Def",
    "subdomain": "abc-def",
    "website_url": "https://www.abcdef.com/"
  }
]
```

Without any arguments the accounts accessible are returned in JSON format. You can 
add a `-debug` flag to trace traffic to and from your web service. Responses from the 
tool are currently displayed in JSON format. The commands you can use are listed if 
you use the `-help` flag:

```bash
[bash] build/workable -help
Usage of workable:
  workable <flags> <command> (<args>)

Flags:
  -debug
    	Trace request and reponse with API
  -endpoint string
    	Endpoint URL, can be overridden with WORKABLE_ENDPOINT environment variable
  -key string
    	API Key, can be overridden with WORKABLE_APIKEY environment variable
```

## License

Copyright 2021 The VFX Cooperative

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

>http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

