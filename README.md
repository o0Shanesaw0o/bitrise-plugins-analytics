# plugins-analytics

Analytics Plugin for Bitrise CLI

Submitting anonymized usage information.  
This usage helps us identify problems with the integrations.  

The sent data only contains information about steps (id, version, runtime, error), **NO logs or other data is included**.

## How to use this Plugin

Can be run directly with the [bitrise CLI](https://github.com/bitrise-io/bitrise), requires version 1.3.0 or newer.

First install the plugin:

```
bitrise plugin install --source https://github.com/bitrise-core/bitrise-plugins-analytics.git
```

After that, you can use it:

```
bitrise :analytics
```
