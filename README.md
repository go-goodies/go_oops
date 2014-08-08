# Summary

<img align="right" src="https://github.com/go-goodies/go_oops/blob/master/golang-gopher-oop.png">

In software engineering, a design pattern is a general reusable solution to a commonly occurring problem within a given context in software design. 

Whereas a design pattern is not a finished design that can be transformed directly into source code, this repo provide implementations for the Go programming language.

Patterns are formalized best practices that the programmer can use to solve common problems when designing an application or system. 

Object-oriented design patterns typically show relationships and interactions between classes or objects, without specifying the final application classes or objects that are involved. 


# Patterns by Type

## Creational

Main article: Creational pattern

* __Creational__ patterns are ones that create objects for you, rather than having you instantiate objects directly. This gives your program more flexibility in deciding which objects need to be created for a given case.
* __Abstract__ Factory groups object factories that have a common theme.
* __Builder__ constructs complex objects by separating construction and representation.
* __Factory__ Method creates objects without specifying the exact class to create.
* __Prototype__ creates objects by cloning an existing object.
* __Singleton__ restricts object creation for a class to only one instance.


## Structural

These concern class and object composition. They use inheritance to compose interfaces and define ways to compose objects to obtain new functionality.

* __Adapter__ allows classes with incompatible interfaces to work together by wrapping its own interface around that of an already existing class.
* __Bridge__ decouples an abstraction from its implementation so that the two can vary independently.
* __Composite__ composes zero-or-more similar objects so that they can be manipulated as one object.
* __Decorator__ dynamically adds/overrides behaviour in an existing method of an object.
* __Facade__ provides a simplified interface to a large body of code.
* __Flyweight__ reduces the cost of creating and manipulating a large number of similar objects.
* __Proxy__ provides a placeholder for another object to control access, reduce cost, and reduce complexity.


## Behavioral

Most of these design patterns are specifically concerned with communication between objects.

* __Chain__ of responsibility delegates commands to a chain of processing objects.
* __Command__ creates objects which encapsulate actions and parameters.
* __Interpreter__ implements a specialized language.
* __Iterator__ accesses the elements of an object sequentially without exposing its underlying representation.
* __Mediator__ allows loose coupling between classes by being the only class that has detailed knowledge of their methods.
* __Memento__ provides the ability to restore an object to its previous state (undo).
* __Observer__ is a publish/subscribe pattern which allows a number of observer objects to see an event.
* __State__ allows an object to alter its behavior when its internal state changes.
* __Strategy__ allows one of a family of algorithms to be selected on-the-fly at runtime.
* __Template__ method defines the skeleton of an algorithm as an abstract class, allowing its subclasses to provide concrete behavior.
* __Visitor__ separates an algorithm from an object structure by moving the hierarchy of methods into one object.

## Notes

See interface documenation at [package go_ops] (http://godoc.org/github.com/go-goodies/go_oops)

See companion article at [Golang Code Examples - Web Server] (http://l3x.github.io/golang-code-examples/2014/08/05/web-server.html)

This is the repository that I will refer to reuse OOP constructs for various Go projects.

I created a go-design-patterns github organization, so if you would like to be a regular contributor, let me know and I'll add you to the team.

Otherwise, feel free to fork it.

In either case, thanks in advance for applying your corrections and/or making contributions.

~ Lex

# References

* [Software design pattern] (http://en.wikipedia.org/wiki/Software_design_pattern)
* [Object-oriented Programming (OOP)] (http://en.wikipedia.org/wiki/Object-oriented)
* [Applicaiton Development with Lex Sheehan] (http://lexsheehan.blogspot.com/)