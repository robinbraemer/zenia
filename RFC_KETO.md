# Abstract

This document contains proposals, design as well as implementation decisions
and will be extended and rewritten in the cycle of development of Keto's underlying
ACL system.

> This document's authors:
> - [robinbraemer](https://github.com/robinbraemer)

Curated list of projects with similar goals:
 - https://github.com/robinbraemer/zenia


# Design decisions	

The ACL system in Keto is based off of the Google Zanzibar
[paper](https://research.google/pubs/pub48190/).

## Clarification of "subject" vs. "user"

> TL;DR
> - Use "User", which is more concrete
> - Do not mislead that UserID and UserSet are a Subject as these are very different
>contexts with different complexities.
> - Rename the Subject interface **and** make it **very internal**
>  - or remove the interface entirely

Right now the code uses the terms "Subject" to interface/abstract a
"UserID" vs a "UserSet". This adds a named third layer on top.

Adding and using such an interface with its name on top may
considerably cause confusion to developers using our package
and result in external or in-tree code comments that refer to
a Subject, but actually mean a UserSet or UserID, or refer to multiple
Subjects, but actually mean a UserSet array/slice
`[]UserSet`, but not a `[]Subject` or vice-versa.

This may hinder the overall design and goal to stay simple and consistent,
e.g. wouldn't a set of subjects not be called a SubjectSet?
Nevertheless, the Subject interface is an abstraction of a concrete user,
the "UserID" vs a set of users that are not even user IDs,
but a "UserSet" in its own is an Object+Relation that expands to more tuples
that are pointers to either an exact UserID or expand to more tuples.

While each one of the types contain "User" in the name,
they should be looked at from a different perspective and be considered
a completely different concept also with a different complexity when
implemented in the system.

One could argue that the term "user" (as in the Zanzibar paper for ACL space)
is shorter as a word, as well as more concrete in what it is and does:
"A user `uses` the object/resource to do something on/with it."
You see there is a more clear separation of the actual "User" and an object,
where a	"Subject" can also mean an object.

## Don't make "things" dependent on the project name

Something I often see is code that has or produces **fixed** digital objects,
such as database tables, that include the project's name.

There are many reasons why this should not be done on most cases
and there are many examples where this is not done where one concrete
example would be the Google Kubernetes API where
there is a [`ClusterManager`](https://github.com/googleapis/googleapis/blob/b1e1e0b13b580d8fb7a641978198fc8de7228b2f/google/container/v1/cluster_service.proto#L35),
not a `KubernetesManager`.

Yes it highly depends on in which context the "thing" is,
but there are other reasons. Developers want to make sure they
use the shortest and abstracted but meaningful term for their "thing"
while staying flexible enough and don't make their "thing"
dependent on the project name in case it changes.

Projects get forked, renamed and code gets re-used in other projects.
Try not to include your project's name in our "things"'s name!

**Keto:**
- replace "keto_" table prefix in ["keto_%s_relation_tuples"](https://github.com/ory/keto/blob/3251ca3c2f5056bcda22b9785324e319d54a68ac/persistence/sql/namespace.go#L26)
  with preferably "relation_tuple_%s" or less preferably "%s_relation_tuple"

## Database

- Table names in singular
  - [Here is why!](https://stackoverflow.com/a/5841297/10937429)
  - [Tables vs RESTful resources](https://softwareengineering.stackexchange.com/questions/290646/why-does-convention-say-db-table-names-should-be-singular-but-restful-resources#:~:text=It's%20a%20pretty%20established%20convention,resource%20names%20should%20be%20plural.)

## Storing a UserID and a UserSet in the data store

Currently, the user (yet subject) in the relation tuple table
gets stored under a single column and either contains the
concrete user id or the `<object>#<relation>` UserSet that points
to inherit other users or UserSet tuples.
This column is included in the primary key.

I propose to add four additional columns to the table that will
allow for efficient querying/filtering by UserSet's relation, object & namespace
as well as improve the efficiency for future analytics jobs for analyzing
trees relationships between tuples.

New table schema:
```sql
CREATE TABLE relation_tuple_%namespace%
(
    shard_id    VARCHAR(64),
    object_id   VARCHAR(64),
    relation    VARCHAR(64),
    "user"      VARCHAR(128),
    commit_time TIMESTAMP,
    
    user_id                  VARCHAR(128),
    userset_object_namespace VARCHAR(128),
    userset_object_id        VARCHAR(128),
    userset_relation         VARCHAR(128),

    PRIMARY KEY (shard_id, object_id, relation, "user", commit_time)
);
```
- `user` can be `object_id#relation` or `user_id`
  - this column will still be used in ACL checks
- Note: The columns in the second block are optional,
  in which the `user_id` column is only ever set when relating to an exact user.
  For a UserSet tuple the `user_id` column **must** be null and the `userset_...` columns must be set.

This provides aforementioned benefits since this data now is
hold by the database in efficient binary encodings.
It will also allow us to support more fine-grained filter parameters to the ACL API.

## Prepared statements

- try to use prepared statements where possible
	- otherwise, parse unpredictable statements using a template engine
	  to efficiently generate the SQL queries
	- https://github.com/valyala/fasttemplate
