# Viper (App Config) [[{devops.configuration.viper,architecture.distributed.12factor,01_PM.low_code]]

* Low code application configuration library.
* Features support:
  * Find/load/unmarshal JSON/TOML/YAML/HCL/INI/envfile/Java-properties
  * default values.
  * override through command line flags.
  * alias system (rename parameters without breaking code).
  * Make it easy to tell the difference between when a user has provided
    a command line or config file which is the same as the default.

* PRECEDENCE ORDER
  1. explicit call to Set  
  2. command line flag     
  3. env                   
  4. config
  5. key/value store
  6. default

* Working with "injected" OS ENViroment VARiables:
  * Viper treats ENV VARs variables as case sensitive.
  * Ex.1:
    ```
    SetEnvPrefix("SPF") ←  use "SPF_..."  prefix for ENV.VARs.
    AutomaticEnv()      ←  Alt 1: any viper.Get("SPF_...") will
                                  automatically query ENV.VARs
    BindEnv("id")       ←  Alt 2: viper.Get("SPF_ID") will query ENV.VARs
    BindEnv("id",       ←  Alt 3: Alias/Backward compatibility
       "deprecated",...)          It will query also for SPF_DEPRECATED,...
    ```

* NOTE: ENV.VAR value will be read each time it is accessed
        (vs constant value after first query)

  ```
  | SetEnvKeyReplacer(string...) *strings.Replacer :
  | allows to use strings.Replacer object to rewrite Env keys (to an extent).
  | Use-case: Replace conflictive characters  in Get() with "_" delimiters
  |           in (OS) ENV.VARs.
  | ( replacers can also be established in NewWithOptions function
  |   using EnvKeyReplacer that also accepts a StringReplacer interface
  |   for custom replace-logic ).
  ```

* By default empty ENV.VARs are considered unset, falling back to
  next conf source). AllowEmptyEnv() will consider them set as "".
[[{devops.configuration.viper}]]
