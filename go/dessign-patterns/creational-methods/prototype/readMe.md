<div>
<h1>Prototype:</h1>
<h2>Definition:</h2>
<p>The Prototype is a creational design pattern that allows an object to create a copy of itself. It is used when the cost of creating a new object is more expensive than copying an existing one. This pattern is particularly useful when creating an object from scratch is resource-intensive or when creating objects with complex configurations.</p>
<h2>Implementation:</h2>
<p>In the Prototype pattern, we create an interface that declares a cloning method. Concrete prototype classes implement this interface to clone themselves. The client code can then use the cloning method to create a new object by copying an existing prototype instance, rather than creating a new instance from scratch.</p>
<h2>How it works:</h2>
<p>The Prototype works by defining an abstract prototype interface with a method for cloning. Concrete prototype classes implement this method to create a copy of the instance. The client code uses the prototype interface to clone objects, ensuring that the cloned objects are independent of the original instance.</p>
<h2>What can we solve with the prototype:</h2>
<h3>A real world problem:</h3>
<p>Imagine a graphic design application where users can create various shapes such as circles, rectangles, and lines. Each shape can have different properties such as color, size, and position. Creating these shapes from scratch every time can be resource-intensive, especially when they have complex configurations.</p>
<h3>Solution implemented with the prototype:</h3>
<p>The Prototype can be used to solve this problem by creating an interface `Shape` with a method `clone()`. Concrete classes such as `Circle`, `Rectangle`, and `Line` would implement this interface to create copies of themselves. The client code would use the `Shape` interface to clone shapes, allowing for efficient creation of new shapes with the same properties as existing ones, without the need for reinitializing all properties from scratch.</p>
<h2>Some pros and cons</h2>
<h3>Pros:</h3>
<ul>
<li>Reduces the cost of creating complex objects by cloning existing instances.</li>
<li>Promotes the reuse of existing objects, which can improve performance and resource utilization.</li>
<li>Encapsulates the cloning logic within the prototype class, making the code more maintainable and easier to understand.</li>
<li>Adheres to the Single Responsibility Principle by allowing the cloning process to be separated from the creation process.</li>
</ul>
<h3>Cons:</h3>
<ul>
<li>Can increase the complexity of the code with the addition of multiple prototype classes and cloning logic.</li>
<li>May lead to issues with deep vs shallow copying, requiring careful implementation to ensure proper cloning behavior.</li>
<li>Sometimes overkill for simple object creation tasks where a simple factory method might suffice.</li>
</ul>
<h2>When to use it and when not:</h2>
<h3>When to use it:</h3>
<ul>
<li>When creating a new object is more expensive than cloning an existing one.</li>
<li>When you need to create multiple instances of a complex object with the same initial configuration.</li>
<li>When you want to decouple the creation of objects from their specific classes, promoting flexibility and scalability.</li>
</ul>
<h3>When not to use it:</h3>
<ul>
<li>When the objects being created are simple and do not require complex configurations or initialization.</li>
<li>When the overhead of implementing and maintaining the prototype pattern outweighs the benefits of cloning.</li>
<li>When issues with deep vs shallow copying are too complex to manage effectively.</li>
</ul>
<h4>Graphic representation: </h4>
<img src="https://refactoring.guru/images/patterns/diagrams/prototype/example.png"/>
</div>
