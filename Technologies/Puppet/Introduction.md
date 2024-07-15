# Puppet Introduction
## Course Link [here](https://www.youtube.com/watch?v=llcjg1R0DdM&list=PLEiEAq2VkUULKeFdfrw25YUayg0omb9wK)

### What is Puppet
Puppet is a configuration management tool and can be used as a deployment tool. Puppet is
an infrastructure as code application similar to ansible or chef. Puppet contains the puppet master
which contains the main configuration files that includes manifests, templates, and files. This
is a single host machine that acts as a source of truth that all of the clients will read from.
This server with act as a certificate authority and all of the clients, when requesting to be updated,
will need to be signed by the CA.

When changes are made on the client, a report is generated and sent to the master so the master is
aware of the current state of all of the nodes.

Manifests are configuration files written out in code that define how we want to configure an
environment. They are compiled into catalogs which are executed on the client. These are written
in Ruby and have a pp extension. The full flow look like the following:

- 1: Create the manifests
- 2: Manifests are compiled into catalogs
- 3: Catalogs are deployed on the clients
- 4: Catalogs are run on the client by the agent
- 5: Clients are configured to desired state