#Description of the example:

This example aim to deploy a "standard" web application.

This application is composed of
- a couple of static web pages, hosted on "web servers", which are "apache based"
- a J2EE applicationn, hosted on a middleware plateform, which is weblogic based
- a database, hosted on a db instance, which is oracle based.

# The topology in term of layers

## Inventory of the machines (Layer0)
- We a an admin host admin01 that can ssh any host
- We have 3 nodes, that are all able to communicate via ssh, web01, web02, web03.
- We have 2 middleware nodes, mdw01, mdw02 that can communicate in ssh
- We have 2 db machines in dataguard db01 and db02

The topoloy is represented in layer0.topology.png (see layer0/layer0.topology.dot for source)

### The nodes.json file
This file will list the nodes and the associated role informations (see layer0/nodes.json for example)

Example:
```json
{
   "nodes":[
      {
         "name":"web01",
         "hostname":"web01.local",
         "user":"useradm",
         "layer3":"WebsServers",
         "layer5":"StaticInstance",
         "layer7":"FancyPics"
      },
  ]
}
```

## Product role (Layer 3)
- web01, web02 and web03 will have the role **WebServer** and run apache
- mdw01 and mdw02 will have the role **Middleware** and run Oracle Weblogic Server
- db01 and db02 will have the role **Database** and run Oracle Database 

We can install the product in a parrallel way, the topology is represented in the layer3.topology.png (see layer3/layer3.topology.dot for source)

Then each role has a task list associated (see layer3/layer3.\*.tasks.dot).
On a given node, **flue** will run locate from Layer0 what is the layer3 role associated, and then execute all the tasks.

## Middle role (layer 5)
- WebServer role is to server static pictures
- Middleware is composed of two role: a producer and a consummer. When the setup task will be finished, we may have two weblogic domains and the weblogic stuff associated to run a portal instance (datasources, jms and co)
- Database, as said previously is a dataguard architecture composed of a master and a slave

The "container for static pictures" is an apache instance that can server static content. It may be installed at any moment.

Due to the "Weblogic portal architecture" we shall setup the database first.

See the layer5.topology.png for a representation


