# automaton

Automation of marketing tools via CLI

Some of the tools include creation of guide, automation of applying best practises and settings onto the various marketing platforms etc

# Using the CLI tool

There is currently no binary builds that is loaded up on releases. However, in the upcoming future, some of the releases will make it into the brew and other installable applications.

Some of the example commands available:

## Snapshot

Snapshot command requires you to download a service account minimally. At the same time, there has to be a config file available to be used to get the settings from Google Analytics and download those settings.

Do note that as of now, snapshot command only takes snapshots of the following tools only

- `Google Analytics` (This is not an exhaustive list)

```bash
# Prints the output to command line
automation snapshot

# Save the output to a JSON file
automation snapshot > snapshot.json
```

## Guide

Guide command does not require any initialized file etc. You would first need to run an init command to generate out the initial command that you would use to be able to use it effectively.

Do note that as of now, the guide command only has the capability to create guides for the following tools only:

- `Google Tag Manager`

```bash
# Get the initial configurations
automaton guide init > config.json

# Generate the guide
automaton guide generate
```

## Audit

Audit command does not require any initialized file etc. You would first need to run an init command to generate out the initial command before you would use it effectively.

Do take note that as of now, the audit command only has the capabilty to audit the following tools only:

- `Google Analytics`

```bash
# Get the initial configuration
automaton audit init > config.json

# Audit the property
automaton audit runaudit
```

## Apply

You would first need to run an init command which would generate a config json file based on your tool of choice. This would allow you to fill up the config json file accordingly which would then allow you to manage the tool accordingly.

Do take note that the only tools supported as of now is `Google Tag Manager`

```bash
# Get the initial configuration
automation apply init > config.json

# Apply the settings to a tool
automation apply --config config.json --tool gtm
```

# Contributing to the project

## Quick notes

In order to increase interoperatability between structs esp in the audit object etc, it seems to make sense that the structs follow roughly the same following structure within the same struct:

- Parameters to be used to run the algorithm
- Results to be stored after running the algorithm

With that, it would allow us to do have the following benefits:

- Our parameters can have its own struct of data - certain audits can have a hefty amount of structs
- Our response can have its own struct of data - each audit will implement its own type of data
- These aggregated data will then be passed into the cmd package that will then be used to render the data out
- These prevents out interfaces from going into scenarios of requiring `interface{}` in order to accept certain parameters etc. It won't matter too much if its already dumped into the struct.
- As much as possible, follow the same style throughout the code base:
  - All data needed for the algorithm colocated with algo
  - All result dumped into the same algorithm
  - (Need to ensure this) Any rerunning of the algorithm will always lead to the same result
  - If there is a need to alter the behaviour, then use interfaces to change it. E.g. The audit object would take in differing clients which allows it to switch between actual extraction from API or mocking extraction via tests.
- For any functions that need to connect to external APIs
  - Rather than providing a single struct with the data attached to call and retrieve the data, it would make more sense to actually provide an interface instead.
  - Depending on how we want much of the implementation we would want to externalize, we can either have the user define the data in a struct and use it from there or, we can provide an interface which contains many params - this would then take the data and use it for connecting to the API.
- A tip from going through many changes in the code base - it actually makes plenty of sense to not generalize the interfaces too much. Rather than creating a very generalized interface, smaller and more specialized interfaces work as well

## Interesting libraries to utilize

1. https://github.com/fatih/structs
2. https://github.com/olekukonko/tablewriter
