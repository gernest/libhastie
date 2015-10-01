# libhastie [![Build Status](https://travis-ci.org/gernest/libhastie.svg)](https://travis-ci.org/gernest/libhastie)

A library for static site generation

## What?

This is a library for generatig static websites. It is a fork of
[hastie](https://github.com/mkaz/hastie) a simple and cool static site generator. This is compatible with hastie,
which means you can build your hastie website with this library.


## Motivation

I have been frustrated when I was looking for a library to use in my project. Thre are plenty os static site generators
but they aren't libraries. Many are commandline utilities. So, I decided to build one. Rather than reinventing the wheel
I forked the hastie project and hacked my ideas in.

## Warning
This is still a work it progress and I don't intend to keep compatibility with hastie project
, so expect API to change,  use it at your own risk.


## Usage

    import(
      "github.com/gernest/libhastie"      
    )
    
    sitePath:="/path/to/your/site"
    
    site:=libhastie.NewSIte(sitePath)
    site.Build()
    

That is enough to build your site.

Your site structure
-------

Libhastie, follows the footsteps of other static site generators, it assumes you have a configuration file at the root of your
project. The file name should be `config.json` but you can also load dynamically the configuration 
 via `Config.Load(path string)`

For more details about static websites that libhastie can build visit the hastiie page [here ](https://github.com/mkaz/hastie) there is a good intro into go templates snd static website project.

A sample `config.json` file looks like this

    {
      
        "SourceDir"  : "posts",
        "LayoutDir"   : "layouts",
        "PublishDir" : "public",       
        "BaseUrl"    : "test/public"
    
    }
`
 
You can use check the content of the `test` folder to see what the site may look like.
 
## State

The library is not actovely maintained, you can fork it and hack the way you like.


## Contributing

Contributions are warmly welcome.

## Author

Geofrey Ernest

## Licence
MIT
