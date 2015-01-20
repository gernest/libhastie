# libhastie

A library for static site generation

## Whaaat?

Yup, this is a library for generatig static sites written in pure Go. It is a fork of
[hastie](https://github.com/mkaz/hastie) a simple and cool static site generator. This is compatible with hastie,
which means you can build your hastie website with this library.

I dont't plan to keep compatibility with hastie in the future, so I will keep shaping stuffs as the
library evolve.

## Motivation

I have been frustrated when I was looking for a library to use in my project. Thre are plenty os static site generators
but they aren't libraries. Many are commandline utilities. So, I decided to build one. Rather than reinventing the wheel
I forked the hastie project and hacked my ideas in.

## Warning
This is still work it progress and I don't intend to keep compatibility with hastie project
, and expect API to chance without notice(totally unstable API) so use it at your own
risk, and by the way you can hack it to your needs.


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
project. The file should be `config.json` but you can also load dynamically the configuration 
into the configuration object via `Config.Load(path string)

Please for more details visit the hasite page [here ](https://github.com/mkaz/hastie) there is a good intro into go template
snd static site project.

A sample config.json file looks like this

    {
      
        "SourceDir"  : "posts",
        "LayoutDir"   : "layouts",
        "PublishDir" : "public",       
        "BaseUrl"    : "test/public"
    
    }
`
 
You can use crunch the content of the `test` folder to see what the site may look like.
 
## State

The library lacks many fancy features, as it looks I only intend to make generating static site easier, so you can add
whatever you want on top of it.


## COntributing

Contributions are warmly welcome.

## Authors

geofrey ernest
