# BET-435 - Relational Model in UML

Language: /  
Frameworks: /  
Task code: https://github.com/ouahabs/LD-Academy/tree/master/BET-435%20-%20Relational%20Model%20in%20UML

Task details:
> Understand and study how to design relational databases using UML.
> Study in depth this link => `https://sparxsystems.com/downloads/whitepapers/Database_Modeling_In_UML.pdf`
> super important to design first the logical model and then understand what's the translation to physical model

# Introduction
* All vendors have their own solutions for software systems, which results in incompatibility.
* An object oriented class model (UML) on top of a relational database
* An entity is a class or an object

# The Class Model
* Captures both the data requirments and its objects' behaviour
* A class is a template model from which instances/objects can be created
	* Logical models (classes)
	* Sequence and collaboration (dynamic) diagrams (objects/instances)
* A class features encapsulation which is hiding/showing attributes/methods to other entities/children (public, protected, private).
### Relationships & Identity
- An association is a relationship between 2 classes, multiple types of associations exist:
	* **Aggregation**: the collection of one entity within another
	* **Composition**: an entity is compsed of another/others (stronger aggregation)
	* **Inheritance**: derives behaviour from parent/ancestor classes, for re-use purposes as well as complexity

# The Relational Model
* A model that features units called "tables" which are compsed of one or more "columns", which contain data
* A column is defined with a name and all data within a column follow a certain data type (text, number, etc)
### Behaviour
* Rules can be applied to a table for integrity purposes
	* **Constraints**: uniqueness requirments, allowables values, etc
	* **Triggers**: before/after events (updates, deetes, inserts)
* ***Navigation** is primarily through SQL
* **A primary key** is defined within a table, two types of keys
	* A meaningful key that uses columns from the table
	* A unique identifier 'ID' key that is meant to identify rows
* ***A foreign key*** is the primary key of another table implying an association between the two

# The UML Data Model Profile
Extended UML to support relational modeling, including concepts like keys, triggers, constraints, etc.
* **Tables**: classes with a table icon in top-right corner
* **Columns**: Attributes of the class,  fromat is: `name: type`
* **Constraints**: Methods of a class, format is `"constraint_name" constraint_method()`
	* PK/FK: primary key/foreign key constraint
	* Index
	* Trigger
	* Unique
	* Proc (procedure)
	* Check (validity check)
* A relationship is "identifying" if a foreign key includes all parts of a primary key, "non-identifying" otherwise
* Cardinality constraints can also be present between relationships (e.g 1 -> 0..n)

# The Physical Model
* Represents the physical structure of a database, its contents and deployement.
* Uses **components**, `<<schema>>` to represent a schema and its tables.
---
# Mapping the Class Model to the Relational Model
1. Model class
2. Identify persistent objects
3. Assume each persistent class maps to one relational table
4. Select an inheritance strategy
	1. Each hierarchy in one table
		* Best performace
		* Convenient for updates and selects
	2. Each class has a table (private attributes only + inherited)
		* Middle ground
		* Best encapsulation
		* Can produce redundancy
	3. Each generation has a table (private attributes only)
		* Simplest
		* Easier maintenance
		* Can have a heavy runtime
5. For each class add a unique object identifier
6. Map attributes to columns
7. Map associations to foreign keys
8. Map Aggregation and Composition
	* Many-to-many relationships are usually implemented with a mapping table between the keys of two tables
	* Exclusive relationships may work similarly with a mapping table
9. Define relationship roles
10. Model behaviour
	* Object methods: purity of design, portability, maintenance, and flexibility
	* Triggers, stored procedures (DBMS operations): performance and efficiency
11. Produce a physical model
	* Describes how an application will be deployed, including platforms, operating systems, and artifacts, etc
